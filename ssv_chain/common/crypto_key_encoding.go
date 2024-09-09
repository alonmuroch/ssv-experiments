// Code generated by fastssz. DO NOT EDIT.
// Hash: b6f2106dd3f4291614ca5ca0df6cf07e845c5f779f78fb16749bc22460b7cbef
// Version: 0.1.2
package common

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the CryptoKey object
func (c *CryptoKey) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the CryptoKey object to a target array
func (c *CryptoKey) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(6)

	// Field (0) 'Type'
	dst = append(dst, c.Type[:]...)

	// Offset (1) 'Key'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(c.Key)

	// Field (1) 'Key'
	if size := len(c.Key); size > 1024 {
		err = ssz.ErrBytesLengthFn("CryptoKey.Key", size, 1024)
		return
	}
	dst = append(dst, c.Key...)

	return
}

// UnmarshalSSZ ssz unmarshals the CryptoKey object
func (c *CryptoKey) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 6 {
		return ssz.ErrSize
	}

	tail := buf
	var o1 uint64

	// Field (0) 'Type'
	copy(c.Type[:], buf[0:2])

	// Offset (1) 'Key'
	if o1 = ssz.ReadOffset(buf[2:6]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 6 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'Key'
	{
		buf = tail[o1:]
		if len(buf) > 1024 {
			return ssz.ErrBytesLength
		}
		if cap(c.Key) == 0 {
			c.Key = make([]byte, 0, len(buf))
		}
		c.Key = append(c.Key, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the CryptoKey object
func (c *CryptoKey) SizeSSZ() (size int) {
	size = 6

	// Field (1) 'Key'
	size += len(c.Key)

	return
}

// HashTreeRoot ssz hashes the CryptoKey object
func (c *CryptoKey) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the CryptoKey object with a hasher
func (c *CryptoKey) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Type'
	hh.PutBytes(c.Type[:])

	// Field (1) 'Key'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(c.Key))
		if byteLen > 1024 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(c.Key)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (1024+31)/32)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the CryptoKey object
func (c *CryptoKey) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(c)
}
