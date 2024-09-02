package types

import (
	"bytes"
	"fmt"
	"ssv-experiments/ssv_chain/common"
)

type ClusterInstance struct {
	// Keys holds operator shares per Cluster.Operators
	Keys []*common.CryptoKey `ssz-max:"13"`
	// Metadata is an open meta data container for application specific cluster instance info (e.g. validator pubkey, ID, etc.)
	Metadata []byte `ssz-max:"2048"`
}

type Cluster struct {
	// ID is unique for each cluster
	ID uint64
	// ModuleID for which the cluster belongs
	ModuleID uint64
	// Address that controls the account
	Address []byte `ssz-max:"128"`
	// Operators id's belonging to the cluster
	Operators []uint64 `ssz-max:"13"`
	// FaultyNodes represents the number of faulty nodes a cluster can sustain (fault tolerance)
	FaultyNodes uint64
	// Instances represents cluster instances for a specific module and account (e.g. staking module)
	Instances []*ClusterInstance `ssz-max:"500"`
}

// FindClusterInstance returns instance index or error if not found
func (c *Cluster) FindClusterInstance(instance *ClusterInstance) (uint64, error) {
	root, err := instance.HashTreeRoot()
	if err != nil {
		return 0, err
	}

	for i, inst := range c.Instances {
		root2, err := inst.HashTreeRoot()
		if err != nil {
			return 0, err
		}

		if bytes.Equal(root[:], root2[:]) {
			return uint64(i), nil
		}
	}
	return 0, fmt.Errorf("not found")
}

// RemoveInstance removes cluster instance if found
func (c *Cluster) RemoveInstance(instance *ClusterInstance) {
	panic("implement")
}
