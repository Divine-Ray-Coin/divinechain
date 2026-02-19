package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(registrar codectypes.InterfaceRegistry) {
	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateClassAdmin{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferCollectible{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurnCollectible{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintCollectible{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateClass{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registrar, &_Msg_serviceDesc)
}
