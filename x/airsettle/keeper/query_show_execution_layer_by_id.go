package keeper

import (
	"context"

	"github.com/ComputerKeeda/airsettle/x/airsettle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowExecutionLayerById(goCtx context.Context, req *types.QueryShowExecutionLayerByIdRequest) (*types.QueryShowExecutionLayerByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	chainId := req.Id
	if chainId == "" {
		// no execution layer on this address
		return nil, sdkerrors.ErrKeyNotFound
	}

	executionLayer, found := k.GetExelayerById(ctx, chainId)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryShowExecutionLayerByIdResponse{Exelayer: &executionLayer}, nil
}
