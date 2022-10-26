package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateLock{}, "lockenomics/CreateLock", nil)
	cdc.RegisterConcrete(&MsgCreateDelegatedAmount{}, "lockenomics/CreateDelegatedAmount", nil)
	cdc.RegisterConcrete(&MsgUpdateDelegatedAmount{}, "lockenomics/UpdateDelegatedAmount", nil)
	cdc.RegisterConcrete(&MsgDeleteDelegatedAmount{}, "lockenomics/DeleteDelegatedAmount", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateLock{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDelegatedAmount{},
		&MsgUpdateDelegatedAmount{},
		&MsgDeleteDelegatedAmount{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
