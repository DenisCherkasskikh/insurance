package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/DenisCherkasskikh/insurance/testutil/keeper"
	"github.com/DenisCherkasskikh/insurance/testutil/nullify"
	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
)

func TestNextRecQuery(t *testing.T) {
	keeper, ctx := keepertest.InsuranceKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestNextRec(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNextRecRequest
		response *types.QueryGetNextRecResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetNextRecRequest{},
			response: &types.QueryGetNextRecResponse{NextRec: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.NextRec(wctx, tc.request)
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
