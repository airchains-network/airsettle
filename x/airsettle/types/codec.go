package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAddExecutionLayer{}, "airsettle/AddExecutionLayer", nil)
	cdc.RegisterConcrete(&MsgDeleteExecutionLayer{}, "airsettle/DeleteExecutionLayer", nil)
	cdc.RegisterConcrete(&MsgAddBatch{}, "airsettle/AddBatch", nil)
	cdc.RegisterConcrete(&MsgVerifyBatch{}, "airsettle/VerifyBatch", nil)
	cdc.RegisterConcrete(&MsgAddValidator{}, "airsettle/AddValidator", nil)
	cdc.RegisterConcrete(&MsgVotePoll{}, "airsettle/VotePoll", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddExecutionLayer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteExecutionLayer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddBatch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVerifyBatch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddValidator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVotePoll{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
