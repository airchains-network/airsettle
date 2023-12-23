package airsettle

import (
	"math/rand"

	"github.com/ComputerKeeda/airsettle/testutil/sample"
	airsettlesimulation "github.com/ComputerKeeda/airsettle/x/airsettle/simulation"
	"github.com/ComputerKeeda/airsettle/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = airsettlesimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgAddExecutionLayer = "op_weight_msg_add_execution_layer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddExecutionLayer int = 100

	opWeightMsgDeleteExecutionLayer = "op_weight_msg_delete_execution_layer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteExecutionLayer int = 100

	opWeightMsgAddBatch = "op_weight_msg_add_batch"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddBatch int = 100

	opWeightMsgVerifyBatch = "op_weight_msg_verify_batch"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVerifyBatch int = 100

	opWeightMsgAddValidator = "op_weight_msg_add_validator"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddValidator int = 100

	opWeightMsgVotePoll = "op_weight_msg_vote_poll"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVotePoll int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	airsettleGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&airsettleGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAddExecutionLayer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddExecutionLayer, &weightMsgAddExecutionLayer, nil,
		func(_ *rand.Rand) {
			weightMsgAddExecutionLayer = defaultWeightMsgAddExecutionLayer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddExecutionLayer,
		airsettlesimulation.SimulateMsgAddExecutionLayer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteExecutionLayer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteExecutionLayer, &weightMsgDeleteExecutionLayer, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteExecutionLayer = defaultWeightMsgDeleteExecutionLayer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteExecutionLayer,
		airsettlesimulation.SimulateMsgDeleteExecutionLayer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddBatch int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddBatch, &weightMsgAddBatch, nil,
		func(_ *rand.Rand) {
			weightMsgAddBatch = defaultWeightMsgAddBatch
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddBatch,
		airsettlesimulation.SimulateMsgAddBatch(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVerifyBatch int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgVerifyBatch, &weightMsgVerifyBatch, nil,
		func(_ *rand.Rand) {
			weightMsgVerifyBatch = defaultWeightMsgVerifyBatch
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVerifyBatch,
		airsettlesimulation.SimulateMsgVerifyBatch(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddValidator int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddValidator, &weightMsgAddValidator, nil,
		func(_ *rand.Rand) {
			weightMsgAddValidator = defaultWeightMsgAddValidator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddValidator,
		airsettlesimulation.SimulateMsgAddValidator(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVotePoll int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgVotePoll, &weightMsgVotePoll, nil,
		func(_ *rand.Rand) {
			weightMsgVotePoll = defaultWeightMsgVotePoll
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVotePoll,
		airsettlesimulation.SimulateMsgVotePoll(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgAddExecutionLayer,
			defaultWeightMsgAddExecutionLayer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				airsettlesimulation.SimulateMsgAddExecutionLayer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteExecutionLayer,
			defaultWeightMsgDeleteExecutionLayer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				airsettlesimulation.SimulateMsgDeleteExecutionLayer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAddBatch,
			defaultWeightMsgAddBatch,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				airsettlesimulation.SimulateMsgAddBatch(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgVerifyBatch,
			defaultWeightMsgVerifyBatch,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				airsettlesimulation.SimulateMsgVerifyBatch(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAddValidator,
			defaultWeightMsgAddValidator,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				airsettlesimulation.SimulateMsgAddValidator(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgVotePoll,
			defaultWeightMsgVotePoll,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				airsettlesimulation.SimulateMsgVotePoll(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
