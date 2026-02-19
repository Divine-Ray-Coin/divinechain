package keeper

import (
	"context"

	"divine/x/collectibles/types"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) Balance(goCtx context.Context, req *types.QueryBalanceRequest) (*types.QueryBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	count := uint64(0)
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := q.k.Collectibles.Walk(ctx, nil, func(key collections.Pair[string, string], value types.Collectible) (bool, error) {
		if key.K1() == req.ClassId && value.Owner == req.Owner {
			count++
		}
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	return &types.QueryBalanceResponse{
		Amount: count,
	}, nil
}
