package keeper

import (
	"context"

	"ledger/x/ledger/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListLedger(goCtx context.Context, req *types.QueryListLedgerRequest) (*types.QueryListLedgerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LedgerKey))

	var ledgers []types.Ledger
	ledgerRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var ledger types.Ledger
		if err := k.cdc.Unmarshal(value, &ledger); err != nil {
			return err
		}

		ledgers = append(ledgers, ledger)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListLedgerResponse{Ledger: ledgers, Pagination: ledgerRes}, nil
}
