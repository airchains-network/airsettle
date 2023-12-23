package keeper

import (
	"encoding/json"
	"github.com/airchains-network/airsettle/x/airsettle/types"
	bls12381 "github.com/airchains-network/gnark/backend/groth16/bls12-381"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
)

func (k Keeper) AddBatchHelper(ctx sdk.Context, msg *types.MsgAddBatch) *sdkerrors.Error {

	var witness fr.Vector
	witnessByte := []byte(msg.Witness)
	witnessErr := json.Unmarshal(witnessByte, &witness)
	if witnessErr != nil {
		return sdkerrors.ErrInvalidRequest
	}

	exLayer, found := k.GetExelayerById(ctx, msg.ChainId)
	if !found {
		return sdkerrors.ErrKeyNotFound
	}

	isAdmin := false
	for _, value := range exLayer.Validator {
		if value == msg.Creator {
			isAdmin = true
		}
	}

	if !isAdmin {
		return sdkerrors.ErrInvalidRequest
	}

	if exLayer.LatestBatch+1 != msg.BatchNumber {
		return sdkerrors.ErrInvalidRequest
	}

	batchMin := types.BatchMin{
		MerkleRootHash: "",
		ZkProof:        "",
		Witness:        msg.Witness,
		Verified:       "false",
		BatchSubmitter: msg.Creator,
		BatchVerifier:  "",
	}
	dynamicBatchStoreKey := "batches/" + msg.ChainId
	batchStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(
		dynamicBatchStoreKey,
	))
	batchBinary := k.cdc.MustMarshal(&batchMin)
	batchNumberByte := []byte(strconv.FormatUint(msg.BatchNumber, 10))
	batchStore.Set(batchNumberByte, batchBinary)

	return nil
}

func (k Keeper) VerifyBatchHelper(ctx sdk.Context, msg *types.MsgVerifyBatch) *sdkerrors.Error {

	exLayer, found := k.GetExelayerById(ctx, msg.ChainId)
	if !found {
		Log("Verifier: Execution layer not found")
		return sdkerrors.ErrKeyNotFound
	}

	isAdmin := false
	for _, value := range exLayer.Validator {
		if value == msg.Creator {
			isAdmin = true
		}
	}
	if !isAdmin {
		Log("Verifier: Not Admin")
		return sdkerrors.ErrInvalidRequest
	}

	if exLayer.LatestBatch+1 != msg.BatchNumber {

		Log("Verifier: wrong batch number ")
		return sdkerrors.ErrInvalidRequest
	}

	dynamicBatchStoreKey := "batches/" + msg.ChainId
	batchStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(
		dynamicBatchStoreKey,
	))

	batchBinary := batchStore.Get([]byte(strconv.FormatUint(msg.BatchNumber, 10)))
	if batchBinary == nil {

		Log("Verifier: batch_binary error")
		return sdkerrors.ErrKeyNotFound
	}

	var batchMin types.BatchMin
	k.cdc.MustUnmarshal(batchBinary, &batchMin)
	if batchMin.Verified == "true" {

		Log("Verifier: Batch already verified ")
		return sdkerrors.ErrInvalidRequest
	}

	if msg.BatchNumber > 1 {
		batchNumberBytePrevious := []byte(strconv.FormatUint(msg.BatchNumber-1, 10))
		batchBinaryPrevious := batchStore.Get(batchNumberBytePrevious)
		var batchMinPrevious types.BatchMin
		k.cdc.MustUnmarshal(batchBinaryPrevious, &batchMinPrevious)
		if batchMinPrevious.MerkleRootHash != msg.PrevMerkleRoot {
			Log("Verifier: Previous merkle root hash not correct")
			return sdkerrors.ErrInvalidRequest
		}
	}

	var proof *bls12381.Proof
	var vk bls12381.VerifyingKey
	var witness fr.Vector

	proofByte := []byte(msg.ZkProof)
	proofErr := json.Unmarshal(proofByte, &proof)
	if proofErr != nil {
		Log("Verifier: Error in proof unmarshal")
		return sdkerrors.ErrInvalidRequest
	}

	witnessByte := []byte(batchMin.Witness)
	witnessErr := json.Unmarshal(witnessByte, &witness)
	if witnessErr != nil {

		Log("Verifier: Error in witness unmarshal")
		return sdkerrors.ErrInvalidRequest
	}

	chainId := msg.ChainId
	vKey, found := k.GetVerificationKeyById(ctx, chainId)
	if !found {
		Log("Verifier: Verification key not found")

		return sdkerrors.ErrKeyNotFound
	}

	vkByte := []byte(vKey.VerificationKey)
	vkErr := json.Unmarshal(vkByte, &vk)
	if vkErr != nil {
		Log("Verifier: Error in vkey unmarshal")

		return sdkerrors.ErrInvalidRequest
	}

	verifyErr := bls12381.Verify(proof, &vk, witness)
	if verifyErr != nil {
		Log("Verifier: Verification failed")

		return sdkerrors.ErrInvalidRequest
	}

	batchMin.Verified = "true"
	batchMin.MerkleRootHash = msg.MerkleRootHash
	batchMin.ZkProof = msg.ZkProof
	batchMin.BatchVerifier = msg.Creator

	batchBinary = k.cdc.MustMarshal(&batchMin)
	batchNumberByte := []byte(strconv.FormatUint(msg.BatchNumber, 10))
	batchStore.Set(batchNumberByte, batchBinary)

	exLayer.LatestBatch = msg.BatchNumber
	exLayer.LatestMerkleRootHash = msg.MerkleRootHash
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutionLayerKey))
	b := k.cdc.MustMarshal(&exLayer)
	store.Set([]byte(exLayer.Id), b)

	return nil
}

func (k Keeper) QueryVerifyBatchHelper(ctx sdk.Context, request *types.QueryQVerifyBatchRequest) *sdkerrors.Error {

	dynamicBatchStoreKey := "batches/" + request.ChainId
	batchStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(
		dynamicBatchStoreKey,
	))

	batchBinary := batchStore.Get([]byte(strconv.FormatUint(request.BatchNumber, 10)))
	if batchBinary == nil {

		Log("Verifier: batch_binary error")
		return sdkerrors.ErrKeyNotFound
	}

	var batchMin types.BatchMin
	k.cdc.MustUnmarshal(batchBinary, &batchMin)

	if request.BatchNumber > 1 {
		batchNumberPrevious := request.BatchNumber - 1
		batchNumberBytePrevious := []byte(strconv.FormatUint(batchNumberPrevious, 10))
		batchBinaryPrevious := batchStore.Get(batchNumberBytePrevious)
		var batchMinPrevious types.BatchMin
		k.cdc.MustUnmarshal(batchBinaryPrevious, &batchMinPrevious)
		if batchMinPrevious.MerkleRootHash != request.PrevMerkleRoot {
			Log("Verifier: Previous merkle root hash not correct")
			return sdkerrors.ErrInvalidRequest
		}
	}

	var proof *bls12381.Proof
	var vk bls12381.VerifyingKey
	var witness fr.Vector

	proofByte := []byte(request.ZkProof)
	proofErr := json.Unmarshal(proofByte, &proof)
	if proofErr != nil {
		Log("Verifier: Error in proof unmarshal")
		return sdkerrors.ErrInvalidRequest
	}

	witnessByte := []byte(batchMin.Witness)
	witnessErr := json.Unmarshal(witnessByte, &witness)
	if witnessErr != nil {

		Log("Verifier: Error in witness unmarshal")
		return sdkerrors.ErrInvalidRequest
	}

	chainId := request.ChainId
	vKey, found := k.GetVerificationKeyById(ctx, chainId)
	if !found {
		Log("Verifier: Verification key not found")

		return sdkerrors.ErrKeyNotFound
	}

	vkByte := []byte(vKey.VerificationKey)
	vkErr := json.Unmarshal(vkByte, &vk)
	if vkErr != nil {
		Log("Verifier: Error in vkey unmarshal")

		return sdkerrors.ErrInvalidRequest
	}

	verifyErr := bls12381.Verify(proof, &vk, witness)
	if verifyErr != nil {
		Log("Verifier: Verification failed")

		return sdkerrors.ErrInvalidRequest
	}

	return nil
}

func (k Keeper) GetBatchHelper(ctx sdk.Context, chainId string, batchNumber uint64) (*types.BatchMax, bool) {

	dynamicBatchStoreKey := "batches/" + chainId
	batchStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(
		dynamicBatchStoreKey,
	))

	batchNumberByte := []byte(strconv.FormatUint(batchNumber, 10))
	batchBinary := batchStore.Get(batchNumberByte)
	if batchBinary == nil {
		return nil, false
	}
	var batchMin types.BatchMin
	k.cdc.MustUnmarshal(batchBinary, &batchMin)

	batchMax := types.BatchMax{
		BatchNumber:        batchNumber,
		ChainId:            chainId,
		PrevMerkleRootHash: "",
		MerkleRootHash:     batchMin.MerkleRootHash,
		ZkProof:            batchMin.ZkProof,
		Witness:            batchMin.Witness,
		Verified:           batchMin.Verified,
		BatchSubmitter:     batchMin.BatchSubmitter,
		BatchVerifier:      batchMin.BatchVerifier,
	}

	if batchNumber == 1 {
		return &batchMax, true
	} else {
		batchNumberPrevious := batchNumber - 1
		batchNumberBytePrevious := []byte(strconv.FormatUint(batchNumberPrevious, 10))
		batchBinaryPrevious := batchStore.Get(batchNumberBytePrevious)
		if batchBinaryPrevious == nil {
			return nil, false
		}
		var batchMinPrevious types.BatchMin
		k.cdc.MustUnmarshal(batchBinaryPrevious, &batchMinPrevious)

		batchMax.PrevMerkleRootHash = batchMinPrevious.MerkleRootHash

		return &batchMax, true
	}
}
