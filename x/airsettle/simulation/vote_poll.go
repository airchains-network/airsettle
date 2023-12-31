package simulation

import (
	"math/rand"

	"github.com/airchains-network/airsettle/x/airsettle/keeper"
	"github.com/airchains-network/airsettle/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgVotePoll(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgVotePoll{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the VotePoll simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "VotePoll simulation not implemented"), nil, nil
	}
}
