package common

import ssz "github.com/ferranbt/fastssz"

// VerifySignature returns true if signature is verified for <network, data, address>
func VerifySignature(network [4]byte, data ssz.HashRoot, signature []byte, address []byte) bool {
	panic("implement")
}
