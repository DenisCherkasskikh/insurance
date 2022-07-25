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

func (k Keeper) CrashRecordAll(c context.Context, req *types.QueryAllCrashRecordRequest) (*types.QueryAllCrashRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var crashRecords []types.CrashRecord
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	crashRecordStore := prefix.NewStore(store, types.KeyPrefix(types.CrashRecordKeyPrefix))

	pageRes, err := query.Paginate(crashRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var crashRecord types.CrashRecord
		if err := k.cdc.Unmarshal(value, &crashRecord); err != nil {
			return err
		}

		crashRecords = append(crashRecords, crashRecord)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCrashRecordResponse{CrashRecord: crashRecords, Pagination: pageRes}, nil
}

func (k Keeper) CrashRecord(c context.Context, req *types.QueryGetCrashRecordRequest) (*types.QueryGetCrashRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCrashRecord(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCrashRecordResponse{CrashRecord: val}, nil
}
