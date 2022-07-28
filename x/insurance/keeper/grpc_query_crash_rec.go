package keeper

import (
	"context"

	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CrashRecAll(c context.Context, req *types.QueryAllCrashRecRequest) (*types.QueryAllCrashRecResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var crashRecs []types.CrashRec
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	crashRecStore := prefix.NewStore(store, types.KeyPrefix(types.CrashRecKeyPrefix))

	pageRes, err := query.Paginate(crashRecStore, req.Pagination, func(key []byte, value []byte) error {
		var crashRec types.CrashRec
		if err := k.cdc.Unmarshal(value, &crashRec); err != nil {
			return err
		}

		crashRecs = append(crashRecs, crashRec)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCrashRecResponse{CrashRec: crashRecs, Pagination: pageRes}, nil
}

func (k Keeper) CrashRec(c context.Context, req *types.QueryGetCrashRecRequest) (*types.QueryGetCrashRecResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCrashRec(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCrashRecResponse{CrashRec: val}, nil
}
