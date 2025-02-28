package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"ledger/x/ledger/types"
)

func (k Keeper) AppendLedger(ctx sdk.Context, ledger types.Ledger) uint64 {
	count := k.GetLedgerCount(ctx)
	ledger.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LedgerKey))
	appendedValue := k.cdc.MustMarshal(&ledger)
	store.Set(GetLedgerIdBytes(ledger.Id), appendedValue)
	k.SetLedgerCount(ctx, count+1)
	return count
}

func (k Keeper) GetLedgerCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.LedgerCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetLedgerIdBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetLedgerCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.LedgerCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetLedger(ctx sdk.Context, id uint64) (val types.Ledger, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LedgerKey))
	b := store.Get(GetLedgerIdBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetLedger(ctx sdk.Context, ledger types.Ledger) {
    storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LedgerKey))
    b := k.cdc.MustMarshal(&ledger)
    store.Set(GetLedgerIdBytes(ledger.Id), b)
}