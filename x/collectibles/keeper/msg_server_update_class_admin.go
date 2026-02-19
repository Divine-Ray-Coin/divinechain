package keeper

import (
	"context"

	"divine/x/collectibles/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateClassAdmin(goCtx context.Context, msg *types.MsgUpdateClassAdmin) (*types.MsgUpdateClassAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// ----------------------------
	// 2️⃣ Ensure class exists
	// ----------------------------
	class, err := k.GetClass(ctx, msg.ClassId)
	if err != nil {
		return nil, types.ErrClassNotFound
	}
	// ----------------------------
	// 1️⃣ Authority check (custodial create)
	// ----------------------------
	signer, err := k.addressCodec.BytesToString(k.GetAuthority())
	if err != nil {
		return nil, types.ErrUnauthorized
	}
	if msg.Creator != class.Admin && msg.Creator != signer {
		return nil, types.ErrUnauthorized
	}
	class.Admin = msg.NewAdmin
	// ----------------------------
	// 6️⃣ Store Class
	// ----------------------------
	if err := k.SetClass(ctx, class); err != nil {
		return nil, err
	}
	// ----------------------------
	// 7️⃣ Emit event
	// ----------------------------
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateClassAdmin,
			sdk.NewAttribute(types.AttributeKeyClassID, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeyClassAdmin, msg.NewAdmin),
		),
	)

	return &types.MsgUpdateClassAdminResponse{}, nil
}
