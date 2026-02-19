package keeper

import (
	"context"

	"divine/x/collectibles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) TransferCollectible(goCtx context.Context, msg *types.MsgTransferCollectible) (*types.MsgTransferCollectibleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// ----------------------------
	// 1️⃣ Validate sender address
	// ----------------------------
	_, err := k.addressCodec.StringToBytes(msg.Creator)
	if err != nil {
		return nil, err
	}

	collectible, err := k.GetCollectible(ctx, msg.ClassId, msg.Id)
	if err != nil {
		return nil, types.ErrCollectibleNotFound
	}

	// ----------------------------
	// 3️⃣ Ownership check
	// ----------------------------
	if collectible.Owner != msg.Creator {
		return nil, types.ErrNotOwner
	}

	// ----------------------------
	// 4️⃣ Prevent transfer if locked (IBC escrow etc.)
	// ----------------------------
	if collectible.Locked {
		return nil, types.ErrCollectibleLocked
	}

	// ----------------------------
	// 5️⃣ Update owner
	// ----------------------------
	collectible.Owner = msg.Receiver

	if err := k.SetCollectible(ctx, collectible); err != nil {
		return nil, err
	}

	// ----------------------------
	// 6️⃣ Emit transfer event
	// ----------------------------
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTransferCollectible,
			sdk.NewAttribute(types.AttributeKeyClassID, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeyCollectibleID, msg.Id),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
		),
	)

	return &types.MsgTransferCollectibleResponse{}, nil
}
