# New SSV-Spec Architecture

The current [SSV-Spec](https://github.com/bloxapp/ssv-spec) is the reference code for which SSV clients can be built off. 
The Spec has been in development since early 2020, carrying a lot of legacy code and complexity.

This new spec architecture has the following improvements/ focus:
1) Minimize implementation details (networking/ storage interfaces, signing, etc) from spec
2) Focus on small functions that are easily testable
3) standardize SSZ encoding throughout 
4) Standardize spec tests, reduce duplicate code

## The ethereum spec and tests as inspiration
* Written in python, an assembly of structures and functions manipulating the states
* The main state mutating function is [state_transition](https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#beacon-chain-state-transition-function)
* All relevant functions for the above function are included in the spec (math, crypto, containers, etc)
* Per fork changes are seperated into different "specs"
* The spec is not runnable
* The tests package is runnable, producing spec tests, ssz encoded, with pre- and post-states
* No implementation details in spec (storage, network, threading, etc), pure logic

## Overview

### Node & Client
To simplify the spec and minimize all "implementation" specific code, the new spec architecture supports a node & client approach (like eth2).

There is a single function call for all message processing called [ProcessMessage](./spec/asgard/process.go) which takes a [P2P Message](./spec/asgard/types/p2p_message.go) and calls subsequent qbft/ runner process messages. [ProcessMessage](./spec/asgard/process.go) is the entry point for the spec.
All state changes happen within [ProcessMessage](./spec/asgard/process.go), or it returns error. 

SSV reacts to incoming messages, broadcasting responses as needed. In the old spec, that was embedded as part of processing messages, resulting in various interfaces within the spec to later be implemented within the actual node. That created a lot of clutter.
With this new design, when [ProcessMessage](./spec/asgard/process.go) is called, only the state changes. No outbound message broadcasting happens.

The responsibility to react to state changes and broadcast responses is up to a [client](./spec/asgard/client.go) code which looks at the state and acts upon it.  
This separation enables flexibility in implementation without coupling spec and implementation.
