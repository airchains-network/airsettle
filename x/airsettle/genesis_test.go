package airsettle_test

import (
	"testing"

	keepertest "github.com/ComputerKeeda/airsettle/testutil/keeper"
	"github.com/ComputerKeeda/airsettle/testutil/nullify"
	"github.com/ComputerKeeda/airsettle/x/airsettle"
	"github.com/ComputerKeeda/airsettle/x/airsettle/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AirsettleKeeper(t)
	airsettle.InitGenesis(ctx, *k, genesisState)
	got := airsettle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
