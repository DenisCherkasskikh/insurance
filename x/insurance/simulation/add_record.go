package simulation

import (
	"math/rand"

	"github.com/DenisCherkasskikh/insurance/x/insurance/keeper"
	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAddRecord(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddRecord{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddRecord simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AddRecord simulation not implemented"), nil, nil
	}
}
