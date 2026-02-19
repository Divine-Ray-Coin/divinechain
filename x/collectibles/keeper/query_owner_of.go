package keeper

import (
	"context"

	"divine/x/collectibles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) OwnerOf(goCtx context.Context, req *types.QueryOwnerOfRequest) (*types.QueryOwnerOfResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	collectible, err := q.k.GetCollectible(ctx, req.ClassId, req.TokenId)
	if err != nil {
		return nil, err
	}

	return &types.QueryOwnerOfResponse{
		Owner: collectible.Owner,
	}, nil
}
