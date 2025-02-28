package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "ledger/testutil/keeper"
	"ledger/x/ledger/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.LedgerKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
