package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddBatch = "add_batch"

var _ sdk.Msg = &MsgAddBatch{}

func NewMsgAddBatch(creator string, batchNumber uint64, chainId string, witness string) *MsgAddBatch {
	return &MsgAddBatch{
		Creator:     creator,
		BatchNumber: batchNumber,
		ChainId:     chainId,
		Witness:     witness,
	}
}

func (msg *MsgAddBatch) Route() string {
	return RouterKey
}

func (msg *MsgAddBatch) Type() string {
	return TypeMsgAddBatch
}

func (msg *MsgAddBatch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddBatch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddBatch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
