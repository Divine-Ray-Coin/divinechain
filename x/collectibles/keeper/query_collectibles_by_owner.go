package keeper

import (
	"context"

	"divine/x/collectibles/types"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) CollectiblesByOwner(goCtx context.Context, req *types.QueryCollectiblesByOwnerRequest) (*types.QueryCollectiblesByOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	collectibles, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Collectibles,
		req.Pagination,
		func(key collections.Pair[string, string], value types.Collectible) (*types.Collectible, error) {
			if value.Owner == req.Owner {
				return &value, nil
			}
			return nil, nil
		},
	)
	if err != nil {
		return nil, err
	}

	return &types.QueryCollectiblesByOwnerResponse{
		Collectibles: collectibles,
		Pagination:   pageRes,
	}, nil
}
