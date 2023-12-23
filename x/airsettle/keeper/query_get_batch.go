package keeper

import (
	"context"

	"github.com/ComputerKeeda/airsettle/x/airsettle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetBatch(goCtx context.Context, req *types.QueryGetBatchRequest) (*types.QueryGetBatchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	batch, found := k.GetBatchHelper(ctx, req.ChainId, req.BatchNumber)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetBatchResponse{
		Batch: batch,
	}, nil
}
