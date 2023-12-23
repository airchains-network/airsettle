package keeper

import (
	"context"
	"github.com/ComputerKeeda/airsettle/x/airsettle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteExecutionLayer(goCtx context.Context, msg *types.MsgDeleteExecutionLayer) (*types.MsgDeleteExecutionLayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	address := msg.Creator

	chainId, found := k.GetExelayerIdByAddress(ctx, address)
	if chainId == "" || !found {

		return nil, sdkerrors.ErrKeyNotFound
	}

	executionLayer, found := k.GetExelayerById(ctx, chainId)
	if !found {

		return nil, sdkerrors.ErrKeyNotFound
	}

	if executionLayer.LatestBatch > 10 {

		return nil, sdkerrors.ErrInvalidRequest
	}

	k.DeleteExecutionLayerHelper(ctx, address, executionLayer.Id)

	k.DecrementExecutionLayerCount(ctx)

	return &types.MsgDeleteExecutionLayerResponse{}, nil
}
