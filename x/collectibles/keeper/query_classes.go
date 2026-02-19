package keeper

import (
	"context"

	"divine/x/collectibles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) Classes(goCtx context.Context, req *types.QueryClassesRequest) (*types.QueryClassesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	classes, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Classes,
		req.Pagination,
		func(key string, value types.Class) (*types.Class, error) {
			return &value, nil
		},
	)
	if err != nil {
		return nil, err
	}

	return &types.QueryClassesResponse{
		Classes:    classes,
		Pagination: pageRes,
	}, nil
}
