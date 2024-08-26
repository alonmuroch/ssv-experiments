// Code generated by fastssz. DO NOT EDIT.
// Hash: d0ff1d3459bc10c140126a891ae47d5c35a121b60a3b423a24ef218240948fe1
// Version: 0.1.2
package types

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/ssv_chain/common"
)

// MarshalSSZ ssz marshals the Operator object
func (o *Operator) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(o)
}

// MarshalSSZTo ssz marshals the Operator object to a target array
func (o *Operator) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(28)

	// Field (0) 'Account'
	dst = ssz.MarshalUint64(dst, o.Account)

	// Field (1) 'ID'
	dst = ssz.MarshalUint64(dst, o.ID)

	// Offset (2) 'PublicKey'
	dst = ssz.WriteOffset(dst, offset)
	if o.PublicKey == nil {
		o.PublicKey = new(common.CryptoKey)
	}
	offset += o.PublicKey.SizeSSZ()

	// Field (3) 'Module'
	dst = ssz.MarshalUint64(dst, o.Module)

	// Field (2) 'PublicKey'
	if dst, err = o.PublicKey.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Operator object
func (o *Operator) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 28 {
		return ssz.ErrSize
	}

	tail := buf
	var o2 uint64

	// Field (0) 'Account'
	o.Account = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'ID'
	o.ID = ssz.UnmarshallUint64(buf[8:16])

	// Offset (2) 'PublicKey'
	if o2 = ssz.ReadOffset(buf[16:20]); o2 > size {
		return ssz.ErrOffset
	}

	if o2 < 28 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (3) 'Module'
	o.Module = ssz.UnmarshallUint64(buf[20:28])

	// Field (2) 'PublicKey'
	{
		buf = tail[o2:]
		if o.PublicKey == nil {
			o.PublicKey = new(common.CryptoKey)
		}
		if err = o.PublicKey.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Operator object
func (o *Operator) SizeSSZ() (size int) {
	size = 28

	// Field (2) 'PublicKey'
	if o.PublicKey == nil {
		o.PublicKey = new(common.CryptoKey)
	}
	size += o.PublicKey.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the Operator object
func (o *Operator) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(o)
}

// HashTreeRootWith ssz hashes the Operator object with a hasher
func (o *Operator) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Account'
	hh.PutUint64(o.Account)

	// Field (1) 'ID'
	hh.PutUint64(o.ID)

	// Field (2) 'PublicKey'
	if err = o.PublicKey.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (3) 'Module'
	hh.PutUint64(o.Module)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Operator object
func (o *Operator) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(o)
}
