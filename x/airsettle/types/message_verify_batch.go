package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgVerifyBatch = "verify_batch"

var _ sdk.Msg = &MsgVerifyBatch{}

func NewMsgVerifyBatch(creator string, batchNumber uint64, chainId string, merkleRootHash string, prevMerkleRoot string, zkProof string) *MsgVerifyBatch {
	return &MsgVerifyBatch{
		Creator:        creator,
		BatchNumber:    batchNumber,
		ChainId:        chainId,
		MerkleRootHash: merkleRootHash,
		PrevMerkleRoot: prevMerkleRoot,
		ZkProof:        zkProof,
	}
}

func (msg *MsgVerifyBatch) Route() string {
	return RouterKey
}

func (msg *MsgVerifyBatch) Type() string {
	return TypeMsgVerifyBatch
}

func (msg *MsgVerifyBatch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVerifyBatch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVerifyBatch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
