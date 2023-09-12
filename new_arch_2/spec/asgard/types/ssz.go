package types

import (
	"encoding/binary"
	ssz "github.com/ferranbt/fastssz"
)

// SSZBytes --
type SSZBytes []byte

func (b SSZBytes) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

func (b SSZBytes) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}

func (b SSZBytes) HashTreeRootWith(hh ssz.HashWalker) error {
	indx := hh.Index()
	hh.PutBytes(b)
	hh.Merkleize(indx)
	return nil
}

type SSZUint64 uint64

func (s SSZUint64) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}

func (s SSZUint64) HashTreeRootWith(hh ssz.HashWalker) error {
	indx := hh.Index()
	hh.PutUint64(uint64(s))
	hh.Merkleize(indx)
	return nil
}

// HashTreeRoot --
func (s SSZUint64) HashTreeRoot() ([32]byte, error) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(s))
	var root [32]byte
	copy(root[:], buf)
	return root, nil
}
