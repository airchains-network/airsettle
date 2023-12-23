package keeper

import (
	"context"

	"github.com/ComputerKeeda/airsettle/x/airsettle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QVerifyBatch(goCtx context.Context, req *types.QueryQVerifyBatchRequest) (*types.QueryQVerifyBatchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	Error := k.QueryVerifyBatchHelper(
		ctx,
		req,
	)

	if Error != nil {
		return &types.QueryQVerifyBatchResponse{
			Verified: false,
		}, nil
	}

	return &types.QueryQVerifyBatchResponse{
		Verified: true,
	}, nil

}
