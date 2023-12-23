package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddExecutionLayer = "add_execution_layer"

var _ sdk.Msg = &MsgAddExecutionLayer{}

func NewMsgAddExecutionLayer(creator string, verificationKey string, chainInfo string) *MsgAddExecutionLayer {
	return &MsgAddExecutionLayer{
		Creator:         creator,
		VerificationKey: verificationKey,
		ChainInfo:       chainInfo,
	}
}

func (msg *MsgAddExecutionLayer) Route() string {
	return RouterKey
}

func (msg *MsgAddExecutionLayer) Type() string {
	return TypeMsgAddExecutionLayer
}

func (msg *MsgAddExecutionLayer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddExecutionLayer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddExecutionLayer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
