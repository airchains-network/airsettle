package keeper

import (
	"context"
	"github.com/airchains-network/airsettle/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListPolls(goCtx context.Context, req *types.QueryListPollsRequest) (*types.QueryListPollsResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	ChainId := req.ChainId
	dynamicPollIdKeyPrefix := types.PollKeyPrefix + ChainId + "/"
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(dynamicPollIdKeyPrefix))

	var polls []*types.Poll

	_, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error { // Make sure req.Pagination is defined
		var poll types.Poll
		if err := k.cdc.Unmarshal(value, &poll); err != nil {
			return err
		}
		if !poll.IsComplete {
			polls = append(polls, &poll)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListPollsResponse{
		Poll: polls,
	}, nil
}
