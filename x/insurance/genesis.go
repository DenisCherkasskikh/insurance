package insurance

import (
	"github.com/DenisCherkasskikh/insurance/x/insurance/keeper"
	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.NextRec != nil {
		k.SetNextRec(ctx, *genState.NextRec)
	}
	// Set all the crashRec
	for _, elem := range genState.CrashRecList {
		k.SetCrashRec(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all nextRec
	nextRec, found := k.GetNextRec(ctx)
	if found {
		genesis.NextRec = &nextRec
	}
	genesis.CrashRecList = k.GetAllCrashRec(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
