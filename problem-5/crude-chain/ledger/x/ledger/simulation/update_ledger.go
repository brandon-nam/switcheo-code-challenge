package simulation

import (
	"math/rand"

	"ledger/x/ledger/keeper"
	"ledger/x/ledger/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgUpdateLedger(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUpdateLedger{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UpdateLedger simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "UpdateLedger simulation not implemented"), nil, nil
	}
}
