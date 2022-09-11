// Code generated by fastssz. DO NOT EDIT.
// Hash: 9a59802e6b47ac0790fa7b09764147ef834ec53184727d702dbbc8e2c4ab8f16
package qbft

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Message object
func (m *Message) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the Message object to a target array
func (m *Message) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	return
}

// UnmarshalSSZ ssz unmarshals the Message object
func (m *Message) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Message object
func (m *Message) SizeSSZ() (size int) {
	size = 4
	return
}

// HashTreeRoot ssz hashes the Message object
func (m *Message) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the Message object with a hasher
func (m *Message) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Message object
func (m *Message) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(m)
}

// MarshalSSZ ssz marshals the SignedMessage object
func (s *SignedMessage) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SignedMessage object to a target array
func (s *SignedMessage) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	return
}

// UnmarshalSSZ ssz unmarshals the SignedMessage object
func (s *SignedMessage) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedMessage object
func (s *SignedMessage) SizeSSZ() (size int) {
	size = 4
	return
}

// HashTreeRoot ssz hashes the SignedMessage object
func (s *SignedMessage) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SignedMessage object with a hasher
func (s *SignedMessage) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SignedMessage object
func (s *SignedMessage) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}

// MarshalSSZ ssz marshals the MessageHeader object
func (m *MessageHeader) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the MessageHeader object to a target array
func (m *MessageHeader) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	return
}

// UnmarshalSSZ ssz unmarshals the MessageHeader object
func (m *MessageHeader) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the MessageHeader object
func (m *MessageHeader) SizeSSZ() (size int) {
	size = 4
	return
}

// HashTreeRoot ssz hashes the MessageHeader object
func (m *MessageHeader) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the MessageHeader object with a hasher
func (m *MessageHeader) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the MessageHeader object
func (m *MessageHeader) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(m)
}

// MarshalSSZ ssz marshals the SignedMessageHeader object
func (s *SignedMessageHeader) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SignedMessageHeader object to a target array
func (s *SignedMessageHeader) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	return
}

// UnmarshalSSZ ssz unmarshals the SignedMessageHeader object
func (s *SignedMessageHeader) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedMessageHeader object
func (s *SignedMessageHeader) SizeSSZ() (size int) {
	size = 4
	return
}

// HashTreeRoot ssz hashes the SignedMessageHeader object
func (s *SignedMessageHeader) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SignedMessageHeader object with a hasher
func (s *SignedMessageHeader) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SignedMessageHeader object
func (s *SignedMessageHeader) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
