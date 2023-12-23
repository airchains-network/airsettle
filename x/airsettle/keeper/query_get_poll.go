package keeper

import (
	"context"
	"github.com/airchains-network/airsettle/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetPoll(goCtx context.Context, req *types.QueryGetPollRequest) (*types.QueryGetPollResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	ChainId := req.ChainId
	PollId := req.PollId

	dynamicPollIdKeyPrefix := types.PollKeyPrefix + ChainId + "/"
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(dynamicPollIdKeyPrefix))

	b := store.Get([]byte(PollId))
	if b == nil {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	var poll types.Poll
	k.cdc.MustUnmarshal(b, &poll)

	return &types.QueryGetPollResponse{
		Poll: &poll,
	}, nil
}
