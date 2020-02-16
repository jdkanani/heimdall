package app

import (
	"encoding/json"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/params"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	"github.com/maticnetwork/heimdall/auth"
	authTypes "github.com/maticnetwork/heimdall/auth/types"
	"github.com/maticnetwork/heimdall/bank"
	bankTypes "github.com/maticnetwork/heimdall/bank/types"
	"github.com/maticnetwork/heimdall/bor"
	borTypes "github.com/maticnetwork/heimdall/bor/types"
	"github.com/maticnetwork/heimdall/checkpoint"
	checkpointTypes "github.com/maticnetwork/heimdall/checkpoint/types"
	"github.com/maticnetwork/heimdall/clerk"
	clerkTypes "github.com/maticnetwork/heimdall/clerk/types"
	"github.com/maticnetwork/heimdall/common"
	"github.com/maticnetwork/heimdall/helper"
	"github.com/maticnetwork/heimdall/staking"
	stakingTypes "github.com/maticnetwork/heimdall/staking/types"
	"github.com/maticnetwork/heimdall/supply"
	supplyTypes "github.com/maticnetwork/heimdall/supply/types"
	"github.com/maticnetwork/heimdall/types"
	"github.com/maticnetwork/heimdall/version"
)

const (
	// AppName denotes app name
	AppName = "Heimdall"
	// ABCIPubKeyTypeSecp256k1 denotes pub key type
	ABCIPubKeyTypeSecp256k1 = "secp256k1"
	// internals
	maxGasPerBlock   int64 = 10000000 // 10 Million
	maxBytesPerBlock int64 = 22020096 // 21 MB
)

var (
	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		supply.AppModuleBasic{},
		staking.AppModuleBasic{},
		checkpoint.AppModuleBasic{},
		bor.AppModuleBasic{},
		clerk.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authTypes.FeeCollectorName: nil,
		// mint.ModuleName:           {supply.Minter},
		// staking.BondedPoolName:    {supply.Burner, supply.Staking},
		// staking.NotBondedPoolName: {supply.Burner, supply.Staking},
		// gov.ModuleName:            {supply.Burner},
	}
)

// HeimdallApp main heimdall app
type HeimdallApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	// keys to access the substores
	keys  map[string]*sdk.KVStoreKey
	tkeys map[string]*sdk.TransientStoreKey

	// subspaces
	subspaces map[string]params.Subspace

	// keepers
	AccountKeeper    auth.AccountKeeper
	BankKeeper       bank.Keeper
	SupplyKeeper     supply.Keeper
	GovKeeper        gov.Keeper
	CheckpointKeeper checkpoint.Keeper
	StakingKeeper    staking.Keeper
	BorKeeper        bor.Keeper
	ClerkKeeper      clerk.Keeper

	// param keeper
	ParamsKeeper params.Keeper

	// masterKeeper common.Keeper
	caller helper.ContractCaller

	//  total coins supply
	TotalCoinsSupply types.Coins

	// the module manager
	mm *module.Manager
}

var logger = helper.Logger.With("module", "app")

//
// Cross communicator
//

// CrossCommunicator retriever
type CrossCommunicator struct {
	App *HeimdallApp
}

// GetACKCount returns ack count
func (d CrossCommunicator) GetACKCount(ctx sdk.Context) uint64 {
	return d.App.CheckpointKeeper.GetACKCount(ctx)
}

// IsCurrentValidatorByAddress check if validator is current validator
func (d CrossCommunicator) IsCurrentValidatorByAddress(ctx sdk.Context, address []byte) bool {
	return d.App.StakingKeeper.IsCurrentValidatorByAddress(ctx, address)
}

//
// Heimdall app
//

// NewHeimdallApp creates heimdall app
func NewHeimdallApp(logger log.Logger, db dbm.DB, baseAppOptions ...func(*bam.BaseApp)) *HeimdallApp {
	// create and register app-level codec for TXs and accounts
	cdc := MakeCodec()

	// create and register pulp codec
	pulp := authTypes.GetPulpInstance()

	// set prefix
	config := sdk.GetConfig()
	config.Seal()

	// base app
	bApp := bam.NewBaseApp(AppName, logger, db, authTypes.RLPTxDecoder(pulp), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(nil)
	bApp.SetAppVersion(version.Version)

	// keys
	keys := sdk.NewKVStoreKeys(
		bam.MainStoreKey,
		authTypes.StoreKey,
		bankTypes.StoreKey,
		supplyTypes.StoreKey,
		// gov.StoreKey,
		stakingTypes.StoreKey,
		checkpointTypes.StoreKey,
		borTypes.StoreKey,
		clerkTypes.StoreKey,
		params.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(params.TStoreKey)

	// create heimdall app
	var app = &HeimdallApp{
		cdc:       cdc,
		BaseApp:   bApp,
		keys:      keys,
		tkeys:     tkeys,
		subspaces: make(map[string]params.Subspace),
	}

	// init params keeper and subspaces
	app.ParamsKeeper = params.NewKeeper(app.cdc, keys[params.StoreKey], tkeys[params.TStoreKey], params.DefaultCodespace)
	app.subspaces[authTypes.ModuleName] = app.ParamsKeeper.Subspace(authTypes.DefaultParamspace)
	app.subspaces[bankTypes.ModuleName] = app.ParamsKeeper.Subspace(bankTypes.DefaultParamspace)
	app.subspaces[supplyTypes.ModuleName] = app.ParamsKeeper.Subspace(supplyTypes.DefaultParamspace)
	// app.subspaces[gov.ModuleName] = app.ParamsKeeper.Subspace(gov.DefaultParamspace).WithKeyTable(gov.ParamKeyTable())
	app.subspaces[stakingTypes.ModuleName] = app.ParamsKeeper.Subspace(stakingTypes.DefaultParamspace)
	app.subspaces[checkpointTypes.ModuleName] = app.ParamsKeeper.Subspace(checkpointTypes.DefaultParamspace)
	app.subspaces[borTypes.ModuleName] = app.ParamsKeeper.Subspace(borTypes.DefaultParamspace)
	app.subspaces[clerkTypes.ModuleName] = app.ParamsKeeper.Subspace(clerkTypes.DefaultParamspace)

	//
	// Contract caller
	//

	contractCallerObj, err := helper.NewContractCaller()
	if err != nil {
		tmos.Exit(err.Error())
	}

	app.caller = contractCallerObj

	//
	// cross communicator
	//

	crossCommunicator := CrossCommunicator{App: app}

	//
	// keepers
	//

	// account keeper
	app.AccountKeeper = auth.NewAccountKeeper(
		app.cdc,
		keys[authTypes.StoreKey], // target store
		app.subspaces[authTypes.ModuleName],
		authTypes.ProtoBaseAccount, // prototype
	)

	app.StakingKeeper = staking.NewKeeper(
		app.cdc,
		keys[stakingTypes.StoreKey], // target store
		app.subspaces[stakingTypes.ModuleName],
		common.DefaultCodespace,
		crossCommunicator,
	)

	// bank keeper
	app.BankKeeper = bank.NewKeeper(
		app.cdc,
		keys[bankTypes.StoreKey], // target store
		app.subspaces[bankTypes.ModuleName],
		bankTypes.DefaultCodespace,
		app.AccountKeeper,
		app.StakingKeeper,
	)

	// bank keeper
	app.SupplyKeeper = supply.NewKeeper(
		app.cdc,
		keys[supplyTypes.StoreKey], // target store
		app.subspaces[supplyTypes.ModuleName],
		maccPerms,
		app.AccountKeeper,
		app.BankKeeper,
	)

	// app.GovKeeper = gov.NewKeeper(
	// 	app.cdc,
	// 	keys[govTypes.StoreKey],
	// 	app.ParamsKeeper,
	// 	app.paramsKeeper.Subspace(gov.DefaultParamspace),
	// 	app.bankKeeper,
	// 	&stakingKeeper,
	// 	gov.DefaultCodespace,
	// )

	app.CheckpointKeeper = checkpoint.NewKeeper(
		app.cdc,
		keys[checkpointTypes.StoreKey], // target store
		app.subspaces[checkpointTypes.ModuleName],
		common.DefaultCodespace,
		app.StakingKeeper,
	)

	app.BorKeeper = bor.NewKeeper(
		app.cdc,
		keys[borTypes.StoreKey], // target store
		app.subspaces[borTypes.ModuleName],
		common.DefaultCodespace,
		app.StakingKeeper,
		app.caller,
	)

	app.ClerkKeeper = clerk.NewKeeper(
		app.cdc,
		keys[clerkTypes.StoreKey], // target store
		app.subspaces[clerkTypes.ModuleName],
		common.DefaultCodespace,
	)

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		auth.NewAppModule(app.AccountKeeper, &app.caller),
		bank.NewAppModule(app.BankKeeper, &app.caller),
		supply.NewAppModule(app.SupplyKeeper, &app.caller),
		staking.NewAppModule(app.StakingKeeper, &app.caller),
		checkpoint.NewAppModule(app.CheckpointKeeper, &app.caller),
		bor.NewAppModule(app.BorKeeper, &app.caller),
		clerk.NewAppModule(app.ClerkKeeper, &app.caller),
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	app.mm.SetOrderInitGenesis(
		authTypes.ModuleName,
		bankTypes.ModuleName,
		supplyTypes.ModuleName,
		stakingTypes.ModuleName,
		checkpointTypes.ModuleName,
		borTypes.ModuleName,
		clerkTypes.ModuleName,
	)

	// register message routes and query routes
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())

	// register message routes
	// app.Router().
	// 	AddRoute(bankTypes.RouterKey, bank.NewHandler(app.bankKeeper, &app.caller)).
	// 	AddRoute(checkpointTypes.RouterKey, checkpoint.NewHandler(app.checkpointKeeper, &app.caller)).
	// 	AddRoute(stakingTypes.RouterKey, staking.NewHandler(app.stakingKeeper, &app.caller)).
	// 	AddRoute(borTypes.RouterKey, bor.NewHandler(app.borKeeper)).
	// 	AddRoute(clerkTypes.RouterKey, clerk.NewHandler(app.clerkKeeper, &app.caller))

	// app.QueryRouter().
	// 	AddRoute(authTypes.QuerierRoute, auth.NewQuerier(app.AccountKeeper)).
	// 	AddRoute(bankTypes.QuerierRoute, bank.NewQuerier(app.bankKeeper)).
	// 	AddRoute(supplyTypes.QuerierRoute, supply.NewQuerier(app.supplyKeeper)).
	// 	AddRoute(stakingTypes.QuerierRoute, staking.NewQuerier(app.stakingKeeper)).
	// 	AddRoute(checkpointTypes.QuerierRoute, checkpoint.NewQuerier(app.checkpointKeeper)).
	// 	AddRoute(borTypes.QuerierRoute, bor.NewQuerier(app.borKeeper)).
	// 	AddRoute(clerkTypes.QuerierRoute, clerk.NewQuerier(app.clerkKeeper))

	// mount the multistore and load the latest state
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)

	// perform initialization logic
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	app.SetAnteHandler(
		auth.NewAnteHandler(
			app.AccountKeeper,
			app.SupplyKeeper,
			&app.caller,
			auth.DefaultSigVerificationGasConsumer,
		),
	)

	// load latest version
	err = app.LoadLatestVersion(app.keys[bam.MainStoreKey])
	if err != nil {
		tmos.Exit(err.Error())
	}

	app.Seal()
	return app
}

// MakeCodec create codec
func MakeCodec() *codec.Codec {
	cdc := codec.New()

	codec.RegisterCrypto(cdc)
	sdk.RegisterCodec(cdc)

	authTypes.RegisterCodec(cdc)
	bankTypes.RegisterCodec(cdc)
	supplyTypes.RegisterCodec(cdc)
	checkpointTypes.RegisterCodec(cdc)
	stakingTypes.RegisterCodec(cdc)
	borTypes.RegisterCodec(cdc)
	clerkTypes.RegisterCodec(cdc)

	cdc.Seal()
	return cdc
}

// MakePulp creates pulp codec and registers custom types for decoder
func MakePulp() *authTypes.Pulp {
	pulp := authTypes.GetPulpInstance()

	// register custom type
	bankTypes.RegisterPulp(pulp)
	stakingTypes.RegisterPulp(pulp)
	checkpointTypes.RegisterPulp(pulp)
	borTypes.RegisterPulp(pulp)
	clerkTypes.RegisterPulp(pulp)

	return pulp
}

// Name returns the name of the App
func (app *HeimdallApp) Name() string { return app.BaseApp.Name() }

// InitChainer initializes chain
func (app *HeimdallApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	err := json.Unmarshal(req.AppStateBytes, &genesisState)
	if err != nil {
		panic(err)
	}

	// get validator updates
	app.mm.InitGenesis(ctx, genesisState)

	stakingState := stakingTypes.GetGenesisStateFromAppState(genesisState)
	checkpointState := checkpointTypes.GetGenesisStateFromAppState(genesisState)

	// check if validator is current validator
	// add to val updates else skip
	var valUpdates []abci.ValidatorUpdate
	for _, validator := range stakingState.Validators {
		if validator.IsCurrentValidator(checkpointState.AckCount) {
			// convert to Validator Update
			updateVal := abci.ValidatorUpdate{
				Power:  int64(validator.VotingPower),
				PubKey: validator.PubKey.ABCIPubKey(),
			}
			// Add validator to validator updated to be processed below
			valUpdates = append(valUpdates, updateVal)
		}
	}

	// TODO make sure old validtors dont go in validator updates ie deactivated validators have to be removed
	// udpate validators
	return abci.ResponseInitChain{
		// validator updates
		Validators: valUpdates,

		// consensus params
		ConsensusParams: &abci.ConsensusParams{
			Block: &abci.BlockParams{
				MaxBytes: maxBytesPerBlock,
				MaxGas:   maxGasPerBlock,
			},
			Evidence:  &abci.EvidenceParams{},
			Validator: &abci.ValidatorParams{PubKeyTypes: []string{ABCIPubKeyTypeSecp256k1}},
		},
	}
}

// BeginBlocker application updates every begin block
func (app *HeimdallApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	app.AccountKeeper.SetBlockProposer(
		ctx,
		types.BytesToHeimdallAddress(req.Header.GetProposerAddress()),
	)
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker executes on each end block
func (app *HeimdallApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	// transfer fees to current proposer
	if proposer, ok := app.AccountKeeper.GetBlockProposer(ctx); ok {
		moduleAccount := app.SupplyKeeper.GetModuleAccount(ctx, authTypes.FeeCollectorName)
		amount := moduleAccount.GetCoins().AmountOf(authTypes.FeeToken)
		if !amount.IsZero() {
			coins := types.Coins{types.Coin{Denom: authTypes.FeeToken, Amount: amount}}
			app.SupplyKeeper.SendCoinsFromModuleToAccount(ctx, authTypes.FeeCollectorName, proposer, coins)
		}

		// remove block proposer
		app.AccountKeeper.RemoveBlockProposer(ctx)
	}

	var tmValUpdates []abci.ValidatorUpdate
	if ctx.BlockHeader().NumTxs > 0 {
		// --- Start update to new validators
		currentValidatorSet := app.StakingKeeper.GetValidatorSet(ctx)
		allValidators := app.StakingKeeper.GetAllValidators(ctx)
		ackCount := app.CheckpointKeeper.GetACKCount(ctx)

		// get validator updates
		setUpdates := helper.GetUpdatedValidators(
			&currentValidatorSet, // pointer to current validator set -- UpdateValidators will modify it
			allValidators,        // All validators
			ackCount,             // ack count
		)

		// create new validator set
		if err := currentValidatorSet.UpdateWithChangeSet(setUpdates); err != nil {
			// return with nothing
			logger.Error("Unable to update current validator set", "Error", err)
			return abci.ResponseEndBlock{}
		}

		// save set in store
		if err := app.StakingKeeper.UpdateValidatorSetInStore(ctx, currentValidatorSet); err != nil {
			// return with nothing
			logger.Error("Unable to update current validator set in state", "Error", err)
			return abci.ResponseEndBlock{}
		}

		// convert updates from map to array
		for _, v := range setUpdates {
			tmValUpdates = append(tmValUpdates, abci.ValidatorUpdate{
				Power:  int64(v.VotingPower),
				PubKey: v.PubKey.ABCIPubKey(),
			})
		}
	}

	// send validator updates to peppermint
	return abci.ResponseEndBlock{
		ValidatorUpdates: tmValUpdates,
	}
}

// initialize store from a genesis state
// func (app *HeimdallApp) initFromGenesisState(ctx sdk.Context, genesisState GenesisState) []abci.ValidatorUpdate {

// 	// Load the genesis accounts
// 	for _, genacc := range genesisState.Accounts {
// 		acc := app.accountKeeper.NewAccountWithAddress(ctx, types.BytesToHeimdallAddress(genacc.Address.Bytes()))
// 		acc.SetCoins(genacc.Coins)
// 		acc.SetSequence(genacc.Sequence)
// 		app.accountKeeper.SetAccount(ctx, acc)
// 	}

// 	//
// 	// InitGenesis
// 	//
// 	auth.InitGenesis(ctx, app.accountKeeper, genesisState.AuthData)
// 	bank.InitGenesis(ctx, app.bankKeeper, genesisState.BankData)
// 	supply.InitGenesis(ctx, app.supplyKeeper, app.accountKeeper, genesisState.SupplyData)
// 	bor.InitGenesis(ctx, app.borKeeper, genesisState.BorData)
// 	// staking should be initialized before checkpoint as checkpoint genesis initialization may depend on staking genesis. [eg.. rewardroot calculation]
// 	staking.InitGenesis(ctx, app.stakingKeeper, genesisState.StakingData)
// 	checkpoint.InitGenesis(ctx, app.checkpointKeeper, genesisState.CheckpointData)
// 	clerk.InitGenesis(ctx, app.clerkKeeper, genesisState.ClerkData)
// 	// validate genesis state
// 	if err := ValidateGenesisState(genesisState); err != nil {
// 		panic(err) // TODO find a way to do this w/o panics
// 	}

// 	// increment accumulator if starting from genesis
// 	if isGenesis {
// 		app.StakingKeeper.IncrementAccum(ctx, 1)
// 	}

// 	//
// 	// get val updates
// 	//

// 	var valUpdates []abci.ValidatorUpdate

// 	// check if validator is current validator
// 	// add to val updates else skip
// 	for _, validator := range genesisState.StakingData.Validators {
// 		if validator.IsCurrentValidator(genesisState.CheckpointData.AckCount) {
// 			// convert to Validator Update
// 			updateVal := abci.ValidatorUpdate{
// 				Power:  int64(validator.VotingPower),
// 				PubKey: validator.PubKey.ABCIPubKey(),
// 			}
// 			// Add validator to validator updated to be processed below
// 			valUpdates = append(valUpdates, updateVal)
// 		}
// 	}
// 	return valUpdates
// }

// LoadHeight loads a particular height
func (app *HeimdallApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keys[bam.MainStoreKey])
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *HeimdallApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[supplyTypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// Codec returns HeimdallApp's codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *HeimdallApp) Codec() *codec.Codec {
	return app.cdc
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *HeimdallApp) GetKey(storeKey string) *sdk.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *HeimdallApp) GetTKey(storeKey string) *sdk.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *HeimdallApp) GetSubspace(moduleName string) params.Subspace {
	return app.subspaces[moduleName]
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}
