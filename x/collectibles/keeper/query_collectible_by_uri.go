package keeper

import (
	"context"

	"divine/x/collectibles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) CollectibleByUri(goCtx context.Context, req *types.QueryCollectibleByUriRequest) (*types.QueryCollectibleByUriResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	collectible, err := q.k.GetCollectibleByURI(ctx, req.Uri)
	if err != nil {
		return nil, err
	}

	return &types.QueryCollectibleByUriResponse{
		Collectible: &collectible,
	}, nil
}
