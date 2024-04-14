package mdvt

import (
	"context"
	"ssv-experiments/mdvt/module"
)

type Justification struct {
}

// DAM interface for node -> DAM requests
type DAM interface {
	AddedToCluster(id uint64, share *module.Share, encryptedShare []byte)
	RemovedFromCluster(id uint64)
}

// API for DAM -> node requests
type API interface {
	// Decide will start a consensus instance, returning decided value and justifications OR error otherwise (e.g. timeout)
	Decide(context context.Context, id [32]byte, in []byte, validate func(in []byte) error) (decidedValue []byte, justifications []*Justification, err error)
	// CollectSignatures broadcasts a partial signature (by the operator's share) and returns a quorum of partial signatures from other operators
	CollectSignatures(context context.Context, id [32]byte, sig []byte) ([][]byte, error)
	// Sign will return a signature for the provided hash
	Sign(hash []byte) (signature []byte)
	RegisterModule()
}
