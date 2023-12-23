package keeper

import (
	"context"
	"fmt"

	"github.com/airchains-network/airsettle/x/airsettle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddBatch(goCtx context.Context, msg *types.MsgAddBatch) (*types.MsgAddBatchResponse, error) {

	Log("add batch called with batch number, and witness:")
	Log(fmt.Sprint(msg.BatchNumber))
	Log(msg.Witness)

	ctx := sdk.UnwrapSDKContext(goCtx)

	Error := k.AddBatchHelper(
		ctx,
		msg,
	)

	if Error != nil {
		Log("adding batch failed, error; " + Error.Error())
		return nil, Error
	}

	Log("Batch added successfully")

	return &types.MsgAddBatchResponse{}, nil

}
