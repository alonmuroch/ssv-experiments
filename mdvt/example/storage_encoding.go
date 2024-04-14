// Code generated by fastssz. DO NOT EDIT.
// Hash: b229ff513834a2b5997c1a63cf3cbe922915d7d8ce69a5caebd1810b87a5d5f8
// Version: 0.1.2
package example

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the ExtraData object
func (e *ExtraData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(e)
}

// MarshalSSZTo ssz marshals the ExtraData object to a target array
func (e *ExtraData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'ValidatorPK'
	dst = append(dst, e.ValidatorPK[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the ExtraData object
func (e *ExtraData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 48 {
		return ssz.ErrSize
	}

	// Field (0) 'ValidatorPK'
	copy(e.ValidatorPK[:], buf[0:48])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ExtraData object
func (e *ExtraData) SizeSSZ() (size int) {
	size = 48
	return
}

// HashTreeRoot ssz hashes the ExtraData object
func (e *ExtraData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the ExtraData object with a hasher
func (e *ExtraData) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ValidatorPK'
	hh.PutBytes(e.ValidatorPK[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ExtraData object
func (e *ExtraData) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(e)
}
