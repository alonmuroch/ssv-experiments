// Code generated by fastssz. DO NOT EDIT.
// Hash: 57d2a45a3903274c0ec458a60afcada9a6c8f74ed6889cd015f65fe18d11fa1d
// Version: 0.1.2
package operator

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/ssv_chain/common"
	"ssv-experiments/ssv_chain/types"
)

// MarshalSSZ ssz marshals the AddOperatorV0 object
func (a *AddOperatorV0) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(a)
}

// MarshalSSZTo ssz marshals the AddOperatorV0 object to a target array
func (a *AddOperatorV0) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(16)

	// Offset (0) 'PublicKey'
	dst = ssz.WriteOffset(dst, offset)
	if a.PublicKey == nil {
		a.PublicKey = new(common.CryptoKey)
	}
	offset += a.PublicKey.SizeSSZ()

	// Field (1) 'ModuleID'
	dst = ssz.MarshalUint64(dst, a.ModuleID)

	// Offset (2) 'Tiers'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(a.Tiers); ii++ {
		offset += 4
		offset += a.Tiers[ii].SizeSSZ()
	}

	// Field (0) 'PublicKey'
	if dst, err = a.PublicKey.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'Tiers'
	if size := len(a.Tiers); size > 16 {
		err = ssz.ErrListTooBigFn("AddOperatorV0.Tiers", size, 16)
		return
	}
	{
		offset = 4 * len(a.Tiers)
		for ii := 0; ii < len(a.Tiers); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += a.Tiers[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(a.Tiers); ii++ {
		if dst, err = a.Tiers[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the AddOperatorV0 object
func (a *AddOperatorV0) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 16 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o2 uint64

	// Offset (0) 'PublicKey'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 16 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'ModuleID'
	a.ModuleID = ssz.UnmarshallUint64(buf[4:12])

	// Offset (2) 'Tiers'
	if o2 = ssz.ReadOffset(buf[12:16]); o2 > size || o0 > o2 {
		return ssz.ErrOffset
	}

	// Field (0) 'PublicKey'
	{
		buf = tail[o0:o2]
		if a.PublicKey == nil {
			a.PublicKey = new(common.CryptoKey)
		}
		if err = a.PublicKey.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (2) 'Tiers'
	{
		buf = tail[o2:]
		num, err := ssz.DecodeDynamicLength(buf, 16)
		if err != nil {
			return err
		}
		a.Tiers = make([]*types.PriceTier, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if a.Tiers[indx] == nil {
				a.Tiers[indx] = new(types.PriceTier)
			}
			if err = a.Tiers[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the AddOperatorV0 object
func (a *AddOperatorV0) SizeSSZ() (size int) {
	size = 16

	// Field (0) 'PublicKey'
	if a.PublicKey == nil {
		a.PublicKey = new(common.CryptoKey)
	}
	size += a.PublicKey.SizeSSZ()

	// Field (2) 'Tiers'
	for ii := 0; ii < len(a.Tiers); ii++ {
		size += 4
		size += a.Tiers[ii].SizeSSZ()
	}

	return
}

// HashTreeRoot ssz hashes the AddOperatorV0 object
func (a *AddOperatorV0) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(a)
}

// HashTreeRootWith ssz hashes the AddOperatorV0 object with a hasher
func (a *AddOperatorV0) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'PublicKey'
	if err = a.PublicKey.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'ModuleID'
	hh.PutUint64(a.ModuleID)

	// Field (2) 'Tiers'
	{
		subIndx := hh.Index()
		num := uint64(len(a.Tiers))
		if num > 16 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range a.Tiers {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 16)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the AddOperatorV0 object
func (a *AddOperatorV0) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(a)
}
