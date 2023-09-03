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

### Pipeline - Separating functionality and message processing flow
The current SSV spec is a mini program that has as input P2P messages and as output P2P messages. In the middle we have state mutation. This is also known as Message-State-Message(MSM) paradigm.

**P2P Message -> State Mutation -> P2P Message**

The problem it creates is introducing implementation details into the spec which are otherwise irrelevant.

This new architecture removes that coupling with the introduction of the [Pipeline](../new_arch/pipeline/pipeline.go).
The Pipeline is a flow of functions which have standardized input and output, together forming the full MSM flow. The Pipeline is the implementation, simplistic pipelines can be used for spec tests and more sophisticated pipelines can be a full working node. 

The Pipeline has the goal of manipulating the state. 2 seemingly different pipelines producing the same stat for the same set of messages can be considered equal