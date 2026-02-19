package keeper

import (
	"context"

	"divine/x/collectibles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ClassAdmin(goCtx context.Context, req *types.QueryClassAdminRequest) (*types.QueryClassAdminResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	class, err := q.k.GetClass(ctx, req.ClassId)
	if err != nil {
		return nil, err
	}

	return &types.QueryClassAdminResponse{
		Admin: class.Admin,
	}, nil
}
