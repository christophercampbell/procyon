package engine

import "time"

const DefaultRPCTimeout = 10 * time.Second

const (
	EXECUTION_CODESPACE = "Execution"
	EXECUTION_SUCCESS   = uint32(0)
	EXECUTION_ERROR     = uint32(1)
)

const (
	RESET_EXPONENTIAL_BACKOFF = 1 * time.Minute
)
