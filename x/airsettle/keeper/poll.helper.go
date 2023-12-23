package keeper

import (
	"github.com/airchains-network/airsettle/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetPollHelper(ctx sdk.Context, chainId string, pollId string) (types.Poll, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKeyPrefix+chainId+"/"))
	var poll types.Poll
	byteKey := []byte(pollId)
	err := k.cdc.Unmarshal(store.Get(byteKey), &poll)
	if err != nil {
		return poll, err
	}
	return poll, nil
}
