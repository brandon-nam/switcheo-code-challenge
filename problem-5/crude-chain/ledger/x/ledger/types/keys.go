package types

const (
	// ModuleName defines the module name
	ModuleName = "ledger"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ledger"

	// Ledger Key
	LedgerKey = "Ledger/value/"

	// Ledger Count Key
	LedgerCountKey = "Ledger/count/"
)

var (
	ParamsKey = []byte("p_ledger")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
