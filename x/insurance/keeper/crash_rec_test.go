package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/DenisCherkasskikh/insurance/testutil/keeper"
	"github.com/DenisCherkasskikh/insurance/testutil/nullify"
	"github.com/DenisCherkasskikh/insurance/x/insurance/keeper"
	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNCrashRec(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CrashRec {
	items := make([]types.CrashRec, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetCrashRec(ctx, items[i])
	}
	return items
}

func TestCrashRecGet(t *testing.T) {
	keeper, ctx := keepertest.InsuranceKeeper(t)
	items := createNCrashRec(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCrashRec(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCrashRecRemove(t *testing.T) {
	keeper, ctx := keepertest.InsuranceKeeper(t)
	items := createNCrashRec(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCrashRec(ctx,
			item.Index,
		)
		_, found := keeper.GetCrashRec(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestCrashRecGetAll(t *testing.T) {
	keeper, ctx := keepertest.InsuranceKeeper(t)
	items := createNCrashRec(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCrashRec(ctx)),
	)
}
