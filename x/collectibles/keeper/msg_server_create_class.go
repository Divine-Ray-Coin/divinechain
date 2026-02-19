package keeper

import (
	"context"

	"divine/x/collectibles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateClass(goCtx context.Context, msg *types.MsgCreateClass) (*types.MsgCreateClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// ----------------------------
	// 1️⃣ Authority check (custodial create)
	// ----------------------------
	signer, err := k.addressCodec.BytesToString(k.GetAuthority())
	if err != nil {
		return nil, types.ErrUnauthorized
	}

	if msg.Creator != signer {
		return nil, types.ErrUnauthorized
	}

	// ----------------------------
	// 2️⃣ Ensure class does not exist
	// ----------------------------
	_, err = k.GetClass(ctx, msg.Symbol)
	if err == nil {
		return nil, types.ErrClassExists
	}

	// ----------------------------
	// 5️⃣ Build Class object
	// ----------------------------
	class := types.Class{
		ClassId: msg.Symbol,
		Name:    msg.Name,
		Symbol:  msg.Symbol,
		Uri:     msg.Uri,
		Admin:   msg.Admin,
	}

	// ----------------------------
	// 6️⃣ Store Class
	// ----------------------------
	if err := k.SetClass(ctx, class); err != nil {
		return nil, err
	}

	// Initialize supply
	if err := k.SetSupply(ctx, msg.Symbol, 0); err != nil {
		return nil, err
	}

	// ----------------------------
	// 7️⃣ Emit event
	// ----------------------------
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreateClass,
			sdk.NewAttribute(types.AttributeKeyClassID, msg.Symbol),
			sdk.NewAttribute(types.AttributeKeyClassName, msg.Name),
			sdk.NewAttribute(types.AttributeKeyClassAdmin, msg.Admin),
			sdk.NewAttribute(types.AttributeKeyURI, msg.Uri),
		),
	)

	return &types.MsgCreateClassResponse{
		ClassId: msg.Symbol,
	}, nil
}
