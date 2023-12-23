package keeper_test

import (
	"testing"

	testkeeper "github.com/airchains-network/airsettle/testutil/keeper"
	"github.com/airchains-network/airsettle/x/airsettle/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AirsettleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
