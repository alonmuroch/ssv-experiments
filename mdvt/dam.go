package mdvt

type Justification struct {
}

// DAM interface for node -> DAM requests
type DAM interface {
	// Validate returns nil if provided data blob is valid, error otherwise
	Validate(in []byte) error
	// Process will process incoming message from peer DAM (counterpart for the broadcast func)
	Process(message []byte)
}

// API for DAM -> node requests
type API interface {
	// Decide will start a consensus instance, returning decided value and justifications OR error otherwise (e.g. timeout)
	Decide(id [32]byte, in []byte) (decidedValue []byte, justifications []*Justification, err error)
	// Sign will return a signature for the provided hash
	Sign(hash []byte) (signature []byte)
	// Broadcast will broadcast provided message to peer DAM
	Broadcast(message []byte, topic []byte) error
}
