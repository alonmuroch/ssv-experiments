package module

// Module holds information about an mDVT module as parsed and captured by the SSV contract
type Module struct {
	// ID for the module
	ID uint
	// Share for the operator
	Share Share
	// Clusters registered for the entire module
	Clusters []Cluster `ssz-max:"65536,128"`
}
