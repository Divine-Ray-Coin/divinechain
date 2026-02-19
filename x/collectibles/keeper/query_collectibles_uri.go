package keeper

import (
	"context"

	"divine/x/collectibles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) CollectiblesUri(goCtx context.Context, req *types.QueryCollectiblesUriRequest) (*types.QueryCollectiblesUriResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	collectible, err := q.k.GetCollectible(ctx, req.ClassId, req.TokenId)
	if err != nil {
		return nil, err
	}

	return &types.QueryCollectiblesUriResponse{
		Uri: collectible.Uri,
	}, nil
}
