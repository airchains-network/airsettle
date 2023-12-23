package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/airchains-network/airsettle/testutil/keeper"
	"github.com/airchains-network/airsettle/x/airsettle/keeper"
	"github.com/airchains-network/airsettle/x/airsettle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AirsettleKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
