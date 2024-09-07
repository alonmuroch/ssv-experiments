# SSV-Chain spec

## Overview
This spec includes: types, state, operations and state transition for an ssv-chain

### Asset bridge
Deposits:   ssv-chain validators look at L1 deposit events and generate a deposit operation with the following type: [4]byte{0x0,0x3,0x3,<version byte>}. 
            Validators validate deposit operations have corresponding L1 event, then construct a block with those operations.

Withdrawals: User initiates a withdrawal operations with the following type: [4]byte{0x1,0x3,0x4,<version byte>}.
                Validators process and include in a block the operation.
                Validators generate some (TBD) withdrawable merkle root + proofs and submit it periodically to L1.
                User withdraws assets with proof.

### Operators
An operator needs to register from a valid account, specifiying a crypto key used for message propagation and the module for which the operator registers.
Operators must specify at least 1 price tier of service. Price tiers are what an operator wants to charge from a specific group of users. A simple all-encompassing price tier can be open to everyone and charge some amount of tokens per block.

```go
type PriceTier struct {
    Network [4]byte `ssz-size:"4"`
    // Capacity limits how many validators can be assigned to this price tier
    Capacity uint16
    // Registered marks how many cluster instances use this price tier
    Registered uint16
    // Price is how many payable tokens will be paid, per block
    Price uint64
    // PayableTokenAddress is the L1 address of the token paid for this tier
    PayableTokenAddress []byte `ssz-max:"64"`
    // WhitelistedAddress that can register to this tier, if empty any address can
    WhitelistedAddress [][]byte `ssz-max:"64,128"`
}

type Operator struct {
	// Address that controls the operator
	Address []byte `ssz-max:"128"`
	// ID is unique for each operator
	ID uint64
	// PublicKey the operator uses to send messages
	PublicKey *common.CryptoKey
	// Modules IDs registered to
	Module uint64
	// Tiers represent pricing tiers
	Tiers []*PriceTier `ssz-max:"16"`
}
```

### Clusters
Clusters are defined per the below struct, similar to how they are currently defined in the ssv.network contracts.
The major difference is the use of cluster instances which represent 1 validator using the cluster. 1 cluster has many instances it can hold.
Another crucial difference is that a cluster is defined by a specific account creating it.

```go
type ClusterInstance struct {
    // PriceTierIndexes holds for each operator the price tier this cluster instance uses
    PriceTierIndexes []uint64 `ssz-max:"13"`
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
	// Active is true when cluster is active, if false all operators should not execute per this cluster
	Active bool
	// Instances represents cluster instances for a specific module and account (e.g. staking module)
	Instances []*ClusterInstance `ssz-max:"500"`
}
```

Cluster instances hold the necessary information to execute validator duties (valid shares and meta-data) and charge registering account for per the operator price tier selected.