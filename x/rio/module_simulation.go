package rio

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"rio/testutil/sample"
	riosimulation "rio/x/rio/simulation"
	"rio/x/rio/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = riosimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateCert = "op_weight_msg_create_cert"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCert int = 100

	opWeightMsgSendCert = "op_weight_msg_send_cert"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendCert int = 100

	opWeightMsgCreateResume = "op_weight_msg_create_resume"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateResume int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	rioGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&rioGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateCert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateCert, &weightMsgCreateCert, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCert = defaultWeightMsgCreateCert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCert,
		riosimulation.SimulateMsgCreateCert(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSendCert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSendCert, &weightMsgSendCert, nil,
		func(_ *rand.Rand) {
			weightMsgSendCert = defaultWeightMsgSendCert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSendCert,
		riosimulation.SimulateMsgSendCert(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateResume int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateResume, &weightMsgCreateResume, nil,
		func(_ *rand.Rand) {
			weightMsgCreateResume = defaultWeightMsgCreateResume
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateResume,
		riosimulation.SimulateMsgCreateResume(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
