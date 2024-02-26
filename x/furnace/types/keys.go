package types

const (
	// ModuleName defines the module name
	ModuleName = "furnace"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_furnace"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	ParamsKey = []byte("p_furnace")
)

const (
	EpochBurnConfigKey = "EpochBurnConfig/value/"
)
