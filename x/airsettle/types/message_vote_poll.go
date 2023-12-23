package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgVotePoll = "vote_poll"

var _ sdk.Msg = &MsgVotePoll{}

func NewMsgVotePoll(creator string, chainId string, pollId string, vote bool) *MsgVotePoll {
	return &MsgVotePoll{
		Creator: creator,
		ChainId: chainId,
		PollId:  pollId,
		Vote:    vote,
	}
}

func (msg *MsgVotePoll) Route() string {
	return RouterKey
}

func (msg *MsgVotePoll) Type() string {
	return TypeMsgVotePoll
}

func (msg *MsgVotePoll) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVotePoll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVotePoll) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
