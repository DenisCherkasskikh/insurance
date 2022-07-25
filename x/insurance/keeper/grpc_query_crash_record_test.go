package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/DenisCherkasskikh/insurance/testutil/keeper"
	"github.com/DenisCherkasskikh/insurance/testutil/nullify"
	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCrashRecordQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.InsuranceKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCrashRecord(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCrashRecordRequest
		response *types.QueryGetCrashRecordResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetCrashRecordRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetCrashRecordResponse{CrashRecord: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetCrashRecordRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetCrashRecordResponse{CrashRecord: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetCrashRecordRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.CrashRecord(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestCrashRecordQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.InsuranceKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCrashRecord(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCrashRecordRequest {
		return &types.QueryAllCrashRecordRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CrashRecordAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.CrashRecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.CrashRecord),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CrashRecordAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.CrashRecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.CrashRecord),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.CrashRecordAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.CrashRecord),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.CrashRecordAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
