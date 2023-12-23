package keeper

import (
	"encoding/json"
	"github.com/airchains-network/airsettle/x/airsettle/types"
	bls12381 "github.com/airchains-network/gnark/backend/groth16/bls12-381"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
)

func (k Keeper) AddExecutionLayerHelper(ctx sdk.Context, exelayer types.Exelayer, creator string) *sdkerrors.Error {

	var vk bls12381.VerifyingKey
	err := json.Unmarshal([]byte(exelayer.VerificationKey), &vk)
	if err != nil {
		return sdkerrors.ErrInvalidRequest
	}

	adminStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainAdminKey))
	exeLayerId := []byte(exelayer.Id)
	byteChainId := adminStore.Get([]byte(creator))
	if byteChainId != nil {
		return sdkerrors.ErrInvalidRequest
	}

	adminStore.Set([]byte(creator), []byte(exeLayerId))

	vkStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerificationKey))
	verificationFormat := types.Vkey{
		Id:              exelayer.Id,
		VerificationKey: exelayer.VerificationKey,
	}
	vkBinary := k.cdc.MustMarshal(&verificationFormat)
	vkStore.Set([]byte(exelayer.Id), vkBinary)

	exelayer.VerificationKey = "/verificationKey/" + exelayer.Id + "/"

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutionLayerKey))
	b := k.cdc.MustMarshal(&exelayer)
	store.Set([]byte(exelayer.Id), b)

	countStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CounterStoreKey))
	countByte := countStore.Get([]byte("exelayers"))
	if countByte == nil {
		countStore.Set([]byte("exelayers"), []byte("1"))
	} else {
		count := string(countByte)

		countUint64, err := strconv.ParseUint(count, 10, 64)
		if err != nil {
			return sdkerrors.ErrInvalidRequest
		}
		countUint64++

		count = strconv.FormatUint(countUint64, 10)
		countStore.Set([]byte("exelayers"), []byte(count))
	}

	return nil
}

func (k Keeper) UpdateExecutionLayer(ctx sdk.Context, executionLayer types.Exelayer) *sdkerrors.Error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutionLayerKey))
	b := k.cdc.MustMarshal(&executionLayer)
	store.Set([]byte(executionLayer.Id), b)
	return nil
}

func (k Keeper) GetExelayerById(ctx sdk.Context, id string) (val types.Exelayer, found bool) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutionLayerKey))

	b := store.Get([]byte(id))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetExelayerIdByAddress(ctx sdk.Context, address string) (val string, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainAdminKey))

	b := store.Get([]byte(address))
	if b == nil {
		return val, false
	}

	return string(b), true
}

func (k Keeper) GetExecutionlayers(ctx sdk.Context, id string) (val types.Exelayer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutionLayerKey))

	b := store.Get(types.ExeLayerKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetVerificationKeyById(ctx sdk.Context, id string) (val types.Vkey, found bool) {
	vk_store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerificationKey))

	b := vk_store.Get([]byte(id))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) DeleteExecutionLayerHelper(ctx sdk.Context, address string, chainId string) {

	adminStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainAdminKey))
	adminStore.Delete([]byte(address))

	chainStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutionLayerKey))
	chainStore.Delete([]byte(chainId))

	verificationKeyStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerificationKey))
	verificationKeyStore.Delete([]byte(chainId))

	countStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CounterStoreKey))
	countByte := countStore.Get([]byte("exelayers"))
	if countByte == nil {
		countStore.Set([]byte("exelayers"), []byte("1"))
	} else {

		count := string(countByte)

		countUint64, err := strconv.ParseUint(count, 10, 64)
		if err != nil {

			panic(err)
		}

		countUint64--

		count = strconv.FormatUint(countUint64, 10)
		countStore.Set([]byte("exelayers"), []byte(count))
	}
}

func (k Keeper) DecrementExecutionLayerCount(ctx sdk.Context) {

	countStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CounterStoreKey))
	countByte := countStore.Get([]byte("exelayers"))
	if countByte == nil {
		countStore.Set([]byte("exelayers"), []byte("0"))
	} else {

		count := string(countByte)

		countUint64, err := strconv.ParseUint(count, 10, 64)
		if err != nil {

			panic(err)
		}

		countUint64--

		count = strconv.FormatUint(countUint64, 10)
		countStore.Set([]byte("exelayers"), []byte(count))
	}
}
