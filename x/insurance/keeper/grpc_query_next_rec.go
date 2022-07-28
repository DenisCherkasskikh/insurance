package keeper

import (
	"context"

	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NextRec(c context.Context, req *types.QueryGetNextRecRequest) (*types.QueryGetNextRecResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNextRec(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNextRecResponse{NextRec: val}, nil
}
