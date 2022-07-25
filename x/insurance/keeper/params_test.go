package keeper_test

import (
	"testing"

	testkeeper "github.com/DenisCherkasskikh/insurance/testutil/keeper"
	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.InsuranceKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
