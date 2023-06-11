package ssv

import (
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/types"
)

// Config is the SSV's config structure holding various structs as configuration, not strictly part of the runner or state.
// The config structure varies between runners for different validators or beacon role types
type Config struct {
	// share is the share for the runner with which messages are verified and signed
	Share *types.Share
	// identifier identifies this particular runner
	Identifier p2p.Identifier
}
