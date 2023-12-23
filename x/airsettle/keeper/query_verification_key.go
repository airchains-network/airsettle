package keeper

import (
	"context"
	"github.com/ComputerKeeda/airsettle/x/airsettle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VerificationKey(goCtx context.Context, req *types.QueryVerificationKeyRequest) (*types.QueryVerificationKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	chainId := req.Id
	if chainId == "" {
		// no execution layer on this address
		return nil, sdkerrors.ErrKeyNotFound
	}

	vKey, found := k.GetVerificationKeyById(ctx, chainId)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryVerificationKeyResponse{
		Vkey: vKey.VerificationKey,
	}, nil
}
