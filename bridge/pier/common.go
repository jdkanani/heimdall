package pier

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	cliContext "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/libs/log"
	httpClient "github.com/tendermint/tendermint/rpc/client"
	tmTypes "github.com/tendermint/tendermint/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/maticnetwork/heimdall/helper"
	hmtypes "github.com/maticnetwork/heimdall/types"
	rest "github.com/maticnetwork/heimdall/types/rest"
)

const (
	ChainSyncer          = "chain-syncer"
	HeimdallCheckpointer = "heimdall-checkpointer"
	NoackService         = "checkpoint-no-ack"
	SpanServiceStr       = "span-service"
	ClerkServiceStr      = "clerk-service"
	AMQPConsumerService  = "amqp-consumer-service"

	// TxsURL represents txs url
	TxsURL = "/txs"

	AccountDetailsURL     = "/auth/accounts/%v"
	LastNoAckURL          = "/checkpoint/last-no-ack"
	ProposersURL          = "/staking/proposer/%v"
	StateSyncerURL        = "/clerk/state-syncer-list"
	BufferedCheckpointURL = "/checkpoint/buffer"
	LatestCheckpointURL   = "/checkpoint/latest-checkpoint"
	CurrentProposerURL    = "/staking/current-proposer"
	LatestSpanURL         = "/bor/latest-span"
	SpanProposerURL       = "/bor/span-proposer"
	NextSpanInfoURL       = "/bor/prepare-next-span"
	InitialRewardRootURL  = "/staking/initial-reward-root"

	TransactionTimeout = 1 * time.Minute
	CommitTimeout      = 2 * time.Minute

	BridgeDBFlag = "bridge-db"
)

// Global logger for bridge
var Logger log.Logger

func init() {
	Logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
}

// checks if we are proposer
func isProposer(cliCtx cliContext.CLIContext) bool {
	var proposers []hmtypes.Validator
	count := uint64(1)
	result, err := FetchFromAPI(cliCtx,
		GetHeimdallServerEndpoint(fmt.Sprintf(ProposersURL, strconv.FormatUint(count, 10))),
	)
	if err != nil {
		Logger.Error("Error fetching proposers", "error", err)
		return false
	}

	err = json.Unmarshal(result.Result, &proposers)
	if err != nil {
		Logger.Error("error unmarshalling proposer slice", "error", err)
		return false
	}

	if bytes.Equal(proposers[0].Signer.Bytes(), helper.GetAddress()) {
		return true
	}

	return false
}

// IsStateSyncer returns if current user is state syncer or not
func IsStateSyncer(cliCtx cliContext.CLIContext) bool {
	var stateSyncerList []common.Address
	isStateSyncer := false

	result, err := FetchFromAPI(cliCtx, GetHeimdallServerEndpoint(fmt.Sprintf(StateSyncerURL)))
	if err != nil {
		Logger.Error("Error fetching State syncer", "error", err)
		return false
	}

	err = json.Unmarshal(result.Result, &stateSyncerList)
	if err != nil {
		Logger.Error("Error unmarshalling state syncer", "error", err)
		return false
	}

	for _, syncerAddr := range stateSyncerList {
		if bytes.Equal(syncerAddr.Bytes(), helper.GetAddress()) {
			isStateSyncer = true
			break
		}
	}
	return isStateSyncer
}

// GetHeimdallServerEndpoint returns heimdall server endpoint
func GetHeimdallServerEndpoint(endpoint string) string {
	u, _ := url.Parse(helper.GetConfig().HeimdallServerURL)
	u.Path = path.Join(u.Path, endpoint)
	return u.String()
}

// FetchFromAPI fetches data from any URL
func FetchFromAPI(cliCtx cliContext.CLIContext, URL string) (result rest.ResponseWithHeight, err error) {
	resp, err := http.Get(URL)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// response
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return result, err
		}
		// unmarshall data from buffer
		// var proposers []hmtypes.Validator
		var response rest.ResponseWithHeight
		if err := cliCtx.Codec.UnmarshalJSON(body, &response); err != nil {
			return result, err
		}
		return response, nil
	}

	Logger.Debug("Error while fetching data from URL", "status", resp.StatusCode, "URL", URL)
	return result, fmt.Errorf("Error while fetching data from url: %v, status: %v", URL, resp.StatusCode)
}

// WaitForOneEvent subscribes to a websocket event for the given
// event time and returns upon receiving it one time, or
// when the timeout duration has expired.
//
// This handles subscribing and unsubscribing under the hood
func WaitForOneEvent(tx tmTypes.Tx, client *httpClient.HTTP) (tmTypes.TMEventData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CommitTimeout)
	defer cancel()

	// subscriber
	subscriber := hex.EncodeToString(tx.Hash())

	// query
	query := tmTypes.EventQueryTxFor(tx).String()

	// register for the next event of this type
	eventCh, err := client.Subscribe(ctx, subscriber, query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to subscribe")
	}

	// make sure to unregister after the test is over
	defer client.UnsubscribeAll(ctx, subscriber)
	select {
	case event := <-eventCh:
		return event.Data.(tmTypes.TMEventData), nil
	case <-ctx.Done():
		return nil, errors.New("timed out waiting for event")
	}
}

// FetchVotes fetches votes and extracts sigs from it
func FetchVotes(
	height int64,
	client *httpClient.HTTP,
) (votes []*tmTypes.CommitSig, sigs []byte, chainID string, err error) {
	// get block client
	blockDetails, err := helper.GetBlockWithClient(client, height+1)

	if err != nil {
		return nil, nil, "", err
	}

	// extract votes from response
	preCommits := blockDetails.LastCommit.Precommits

	// extract signs from votes
	valSigs := helper.GetSigs(preCommits)

	// extract chainID
	chainID = blockDetails.ChainID

	// return
	return preCommits, valSigs, chainID, nil
}
