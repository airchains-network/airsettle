package keeper

import (
	"context"

	"github.com/airchains-network/airsettle/x/airsettle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) VerifyBatch(goCtx context.Context, msg *types.MsgVerifyBatch) (*types.MsgVerifyBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	Error := k.VerifyBatchHelper(
		ctx,
		msg,
	)

	if Error != nil {
		Log("verification of batch failed, error; " + Error.Error())
		return nil, Error
	}

	return &types.MsgVerifyBatchResponse{}, nil
}
