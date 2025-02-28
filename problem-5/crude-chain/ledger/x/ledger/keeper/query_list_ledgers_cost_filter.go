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

func (k Keeper) ListLedgersCostFilter(goCtx context.Context, req *types.QueryListLedgersCostFilterRequest) (*types.QueryListLedgersCostFilterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	
    var ledgers []types.Ledger

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	ledgerStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LedgerKey))


    ledgerRes, err := query.Paginate(ledgerStore, req.Pagination, func(key []byte, value []byte) error {
        var ledger types.Ledger
        if err := k.cdc.Unmarshal(value, &ledger); err != nil {
            return err
        }

        // Apply cost filter - smaller than max_cost and bigger than min_cost
        if req.MaxCost >= ledger.Cost && req.MinCost <= ledger.Cost  {
            ledgers = append(ledgers, ledger)
        } else {

			return nil
		}

        return nil
    })

    if err != nil {
        return nil, status.Errorf(codes.Internal, "pagination error: %v", err)
    }

    return &types.QueryListLedgersCostFilterResponse{Ledger: ledgers, Pagination: ledgerRes}, nil
}
