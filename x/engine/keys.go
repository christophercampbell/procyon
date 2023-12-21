package engine

const (
	// ModuleName is the name of the engine module
	ModuleName = "engine"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// RouterKey is the msg router key for the engine module
	RouterKey = ModuleName
)

var (
	stateStorageKey = []byte("state")
)
