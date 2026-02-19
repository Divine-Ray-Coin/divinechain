package keeper

import (
	"context"

	"divine/x/collectibles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) Collectible(goCtx context.Context, req *types.QueryCollectibleRequest) (*types.QueryCollectibleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	collectible, err := q.k.GetCollectible(ctx, req.ClassId, req.TokenId)
	if err != nil {
		return nil, err
	}

	return &types.QueryCollectibleResponse{
		Collectible: &collectible,
	}, nil
}
