package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteExecutionLayer = "delete_execution_layer"

var _ sdk.Msg = &MsgDeleteExecutionLayer{}

func NewMsgDeleteExecutionLayer(creator string) *MsgDeleteExecutionLayer {
	return &MsgDeleteExecutionLayer{
		Creator: creator,
	}
}

func (msg *MsgDeleteExecutionLayer) Route() string {
	return RouterKey
}

func (msg *MsgDeleteExecutionLayer) Type() string {
	return TypeMsgDeleteExecutionLayer
}

func (msg *MsgDeleteExecutionLayer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteExecutionLayer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteExecutionLayer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
