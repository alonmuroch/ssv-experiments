package types

import "ssv-experiments/ssv_chain/common"

type ClusterInstance struct {
	Keys     []*common.CryptoKey `ssz-max:"13"`
	Metadata []byte              `ssz-max:"2048"`
}

type Cluster struct {
	// ID is unique for each cluster
	ID uint64
	// ModuleID for which the cluster belongs
	ModuleID uint64
	// Account that controls the cluster
	Account uint64
	// Operators id's belonging to the cluster
	Operators []uint64 `ssz-max:"13"`
	// FaultyNodes represents the number of faulty nodes a cluster can sustain (fault tolerance)
	FaultyNodes uint64
	// Instances represents cluster instances for a specific module and account (e.g. staking module)
	Instances []*ClusterInstance `ssz-max:"500"`
}
