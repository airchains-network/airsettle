package keeper

import (
	"context"
	"github.com/ComputerKeeda/airsettle/x/airsettle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowExecutionLayerByAddress(goCtx context.Context, req *types.QueryShowExecutionLayerByAddressRequest) (*types.QueryShowExecutionLayerByAddressResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	address := req.Address
	chainId, found := k.GetExelayerIdByAddress(ctx, address)
	if chainId == "" || found == false { // any one condition is enough too
		// no execution layer on this address
		return nil, sdkerrors.ErrKeyNotFound
	}

	executionLayer, found := k.GetExelayerById(ctx, chainId)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryShowExecutionLayerByAddressResponse{Exelayer: &executionLayer}, nil

}
