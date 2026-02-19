package keeper

import (
	"context"

	"divine/x/collectibles/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) Authority(goCtx context.Context, req *types.QueryAuthorityRequest) (*types.QueryAuthorityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	authorityStr, err := q.k.addressCodec.BytesToString(q.k.GetAuthority())
	if err != nil {
		return nil, err
	}

	return &types.QueryAuthorityResponse{
		Address: authorityStr,
	}, nil
}
