package insurance_test

import (
	"testing"

	keepertest "github.com/DenisCherkasskikh/insurance/testutil/keeper"
	"github.com/DenisCherkasskikh/insurance/testutil/nullify"
	"github.com/DenisCherkasskikh/insurance/x/insurance"
	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CrashRecordList: []types.CrashRecord{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.InsuranceKeeper(t)
	insurance.InitGenesis(ctx, *k, genesisState)
	got := insurance.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CrashRecordList, got.CrashRecordList)
	// this line is used by starport scaffolding # genesis/test/assert
}
