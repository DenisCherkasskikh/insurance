package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/DenisCherkasskikh/insurance/testutil/keeper"
	"github.com/DenisCherkasskikh/insurance/testutil/nullify"
	"github.com/DenisCherkasskikh/insurance/x/insurance/keeper"
	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
)

func createTestNextRec(keeper *keeper.Keeper, ctx sdk.Context) types.NextRec {
	item := types.NextRec{}
	keeper.SetNextRec(ctx, item)
	return item
}

func TestNextRecGet(t *testing.T) {
	keeper, ctx := keepertest.InsuranceKeeper(t)
	item := createTestNextRec(keeper, ctx)
	rst, found := keeper.GetNextRec(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestNextRecRemove(t *testing.T) {
	keeper, ctx := keepertest.InsuranceKeeper(t)
	createTestNextRec(keeper, ctx)
	keeper.RemoveNextRec(ctx)
	_, found := keeper.GetNextRec(ctx)
	require.False(t, found)
}
