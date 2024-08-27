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