package clerk

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/maticnetwork/heimdall/clerk/types"
	hmTypes "github.com/maticnetwork/heimdall/types"
)

// NewQuerier creates a querier for auth REST endpoints
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case types.QueryRecord:
			return handleQueryRecord(ctx, req, keeper)
		case types.QueryRecordList:
			return handleQueryRecordList(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown auth query endpoint")
		}
	}
}

func handleQueryRecord(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var params types.QueryRecordParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	// get state record by record id
	record, err := keeper.GetEventRecord(ctx, params.RecordID)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not get state record", err.Error()))
	}

	// return error if record doesn't exist
	if record == nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("record %v does not exist", params.RecordID))
	}

	// json record
	bz, err := json.Marshal(record)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}

func handleQueryRecordList(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var params hmTypes.QueryPaginationParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	res, err := keeper.GetEventRecordList(ctx, params.Page, params.Limit)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr(fmt.Sprintf("could not fetch record list with page %v and limit %v", params.Page, params.Limit), err.Error()))
	}

	bz, err := json.Marshal(res)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
