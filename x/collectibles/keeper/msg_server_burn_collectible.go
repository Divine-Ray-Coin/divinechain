package keeper

import (
	"context"

	"divine/x/collectibles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) BurnCollectible(goCtx context.Context, msg *types.MsgBurnCollectible) (*types.MsgBurnCollectibleResponse, error) {
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
	// ----------------------------
	// 1️⃣ Authority check (custodial create)
	// ----------------------------
	if msg.Creator != string(k.GetAuthority()) {
		return nil, types.ErrUnauthorized
	}

	// ----------------------------
	// 2️⃣ Ensure collectible exists
	// ----------------------------
	collectible, err := k.GetCollectible(ctx, msg.ClassId, msg.Id)
	if err != nil {
		return nil, types.ErrCollectibleNotFound
	}

	// ----------------------------
	// 3️⃣ Prevent burn if locked (IBC escrow etc.)
	// ----------------------------
	if collectible.Locked {
		return nil, types.ErrCollectibleLocked
	}

	// ----------------------------
	// 4️⃣ Remove from store + URI index
	// ----------------------------
	if err := k.RemoveCollectible(ctx, msg.ClassId, msg.Id); err != nil {
		return nil, err
	}

	// ----------------------------
	// 6️⃣ Decrement Supply
	// ----------------------------
	supply, err := k.GetSupply(ctx, collectible.ClassId)
	if err != nil {
		supply = 0
	}
	if err := k.SetSupply(ctx, collectible.ClassId, supply-1); err != nil {
		return nil, err
	}

	// ----------------------------
	// 5️⃣ Emit burn event
	// ----------------------------
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBurnCollectible,
			sdk.NewAttribute(types.AttributeKeyClassID, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeyCollectibleID, msg.Id),
			sdk.NewAttribute(types.AttributeKeyOwner, collectible.Owner),
		),
	)

	return &types.MsgBurnCollectibleResponse{}, nil
}
