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

func createNCrashRecord(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CrashRecord {
	items := make([]types.CrashRecord, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetCrashRecord(ctx, items[i])
	}
	return items
}

func TestCrashRecordGet(t *testing.T) {
	keeper, ctx := keepertest.InsuranceKeeper(t)
	items := createNCrashRecord(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCrashRecord(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCrashRecordRemove(t *testing.T) {
	keeper, ctx := keepertest.InsuranceKeeper(t)
	items := createNCrashRecord(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCrashRecord(ctx,
			item.Index,
		)
		_, found := keeper.GetCrashRecord(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestCrashRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.InsuranceKeeper(t)
	items := createNCrashRecord(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCrashRecord(ctx)),
	)
}
