package keeper

import (
	"context"

	"divine/x/collectibles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MintCollectible(goCtx context.Context, msg *types.MsgMintCollectible) (*types.MsgMintCollectibleResponse, error) {
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
	// 3️⃣ Prevent duplicate URI
	// ----------------------------
	_, err = k.GetCollectibleByURI(ctx, msg.Uri)
	if err == nil {
		return nil, types.ErrURIAlreadyExists
	}

	// ----------------------------
	// 4️⃣ Generate auto ID (per class)
	// ----------------------------
	id, err := k.NextCollectibleID(ctx, msg.ClassId)
	if err != nil {
		return nil, err
	}

	// ----------------------------
	// 5️⃣ Build Collectible object
	// ----------------------------
	collectible := types.Collectible{
		ClassId: msg.ClassId,
		Id:      id,
		Uri:     msg.Uri,
		Owner:   msg.Receiver,
		Locked:  false,
	}

	// ----------------------------
	// 6️⃣ Store collectible + URI index
	// ----------------------------
	if err := k.SetCollectible(ctx, collectible); err != nil {
		return nil, err
	}

	// ----------------------------
	// 6️⃣ Increment Supply
	// ----------------------------
	supply, err := k.GetSupply(ctx, collectible.ClassId)
	if err != nil {
		supply = 0
	}
	if err := k.SetSupply(ctx, collectible.ClassId, supply+1); err != nil {
		return nil, err
	}

	// ----------------------------
	// 7️⃣ Emit event
	// ----------------------------
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMintCollectible,
			sdk.NewAttribute(types.AttributeKeyClassID, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeyCollectibleID, id),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Receiver),
			sdk.NewAttribute(types.AttributeKeyURI, msg.Uri),
		),
	)

	return &types.MsgMintCollectibleResponse{
		ClassId: msg.ClassId,
		TokenId: id,
	}, nil
}
