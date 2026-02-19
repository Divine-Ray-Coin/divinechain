package collectibles

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	collectiblessimulation "divine/x/collectibles/simulation"
	"divine/x/collectibles/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	collectiblesGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&collectiblesGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreateClass          = "op_weight_msg_collectibles"
		defaultWeightMsgCreateClass int = 100
	)

	var weightMsgCreateClass int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateClass, &weightMsgCreateClass, nil,
		func(_ *rand.Rand) {
			weightMsgCreateClass = defaultWeightMsgCreateClass
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateClass,
		collectiblessimulation.SimulateMsgCreateClass(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgMintCollectible          = "op_weight_msg_collectibles"
		defaultWeightMsgMintCollectible int = 100
	)

	var weightMsgMintCollectible int
	simState.AppParams.GetOrGenerate(opWeightMsgMintCollectible, &weightMsgMintCollectible, nil,
		func(_ *rand.Rand) {
			weightMsgMintCollectible = defaultWeightMsgMintCollectible
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintCollectible,
		collectiblessimulation.SimulateMsgMintCollectible(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgBurnCollectible          = "op_weight_msg_collectibles"
		defaultWeightMsgBurnCollectible int = 100
	)

	var weightMsgBurnCollectible int
	simState.AppParams.GetOrGenerate(opWeightMsgBurnCollectible, &weightMsgBurnCollectible, nil,
		func(_ *rand.Rand) {
			weightMsgBurnCollectible = defaultWeightMsgBurnCollectible
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBurnCollectible,
		collectiblessimulation.SimulateMsgBurnCollectible(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgTransferCollectible          = "op_weight_msg_collectibles"
		defaultWeightMsgTransferCollectible int = 100
	)

	var weightMsgTransferCollectible int
	simState.AppParams.GetOrGenerate(opWeightMsgTransferCollectible, &weightMsgTransferCollectible, nil,
		func(_ *rand.Rand) {
			weightMsgTransferCollectible = defaultWeightMsgTransferCollectible
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferCollectible,
		collectiblessimulation.SimulateMsgTransferCollectible(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateClassAdmin          = "op_weight_msg_collectibles"
		defaultWeightMsgUpdateClassAdmin int = 100
	)

	var weightMsgUpdateClassAdmin int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateClassAdmin, &weightMsgUpdateClassAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClassAdmin = defaultWeightMsgUpdateClassAdmin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateClassAdmin,
		collectiblessimulation.SimulateMsgUpdateClassAdmin(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
