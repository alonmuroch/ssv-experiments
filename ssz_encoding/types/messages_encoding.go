// Code generated by fastssz. DO NOT EDIT.
// Hash: 2a5a9c79e7d4f46a15cfaa2180d8d272420f66b59bff17eaeb1350b0f1250ac1
package types

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
	offset := int(36)

	// Field (0) 'ID'
	dst = append(dst, m.ID[:]...)

	// Offset (1) 'DataSSZSnappy'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(m.DataSSZSnappy)

	// Field (1) 'DataSSZSnappy'
	if size := len(m.DataSSZSnappy); size > 2048 {
		err = ssz.ErrBytesLengthFn("Message.DataSSZSnappy", size, 2048)
		return
	}
	dst = append(dst, m.DataSSZSnappy...)

	return
}

// UnmarshalSSZ ssz unmarshals the Message object
func (m *Message) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 36 {
		return ssz.ErrSize
	}

	tail := buf
	var o1 uint64

	// Field (0) 'ID'
	copy(m.ID[:], buf[0:32])

	// Offset (1) 'DataSSZSnappy'
	if o1 = ssz.ReadOffset(buf[32:36]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 36 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'DataSSZSnappy'
	{
		buf = tail[o1:]
		if len(buf) > 2048 {
			return ssz.ErrBytesLength
		}
		if cap(m.DataSSZSnappy) == 0 {
			m.DataSSZSnappy = make([]byte, 0, len(buf))
		}
		m.DataSSZSnappy = append(m.DataSSZSnappy, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Message object
func (m *Message) SizeSSZ() (size int) {
	size = 36

	// Field (1) 'DataSSZSnappy'
	size += len(m.DataSSZSnappy)

	return
}

// HashTreeRoot ssz hashes the Message object
func (m *Message) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the Message object with a hasher
func (m *Message) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ID'
	hh.PutBytes(m.ID[:])

	// Field (1) 'DataSSZSnappy'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(m.DataSSZSnappy))
		if byteLen > 2048 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(m.DataSSZSnappy)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (2048+31)/32)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Message object
func (m *Message) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(m)
}
