// Code generated by fastssz. DO NOT EDIT.
// Hash: 0c11aaf72c110db980be735890075cca69eab00e3a6a9d09e5d35fc8b146942d
// Version: 0.1.3
package qbft

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/new_arch/types"
)

// MarshalSSZ ssz marshals the Instance object
func (i *Instance) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(i)
}

// MarshalSSZTo ssz marshals the Instance object to a target array
func (i *Instance) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(140)

	// Offset (0) 'State'
	dst = ssz.WriteOffset(dst, offset)
	if i.State == nil {
		i.State = new(State)
	}
	offset += i.State.SizeSSZ()

	// Field (1) 'Share'
	if i.Share == nil {
		i.Share = new(types.Share)
	}
	if dst, err = i.Share.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'Identifier'
	dst = append(dst, i.Identifier[:]...)

	// Offset (3) 'StartValue'
	dst = ssz.WriteOffset(dst, offset)
	if i.StartValue == nil {
		i.StartValue = new(types.ConsensusData)
	}
	offset += i.StartValue.SizeSSZ()

	// Field (0) 'State'
	if dst, err = i.State.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (3) 'StartValue'
	if dst, err = i.StartValue.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Instance object
func (i *Instance) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 140 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o3 uint64

	// Offset (0) 'State'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 140 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'Share'
	if i.Share == nil {
		i.Share = new(types.Share)
	}
	if err = i.Share.UnmarshalSSZ(buf[4:80]); err != nil {
		return err
	}

	// Field (2) 'Identifier'
	copy(i.Identifier[:], buf[80:136])

	// Offset (3) 'StartValue'
	if o3 = ssz.ReadOffset(buf[136:140]); o3 > size || o0 > o3 {
		return ssz.ErrOffset
	}

	// Field (0) 'State'
	{
		buf = tail[o0:o3]
		if i.State == nil {
			i.State = new(State)
		}
		if err = i.State.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (3) 'StartValue'
	{
		buf = tail[o3:]
		if i.StartValue == nil {
			i.StartValue = new(types.ConsensusData)
		}
		if err = i.StartValue.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Instance object
func (i *Instance) SizeSSZ() (size int) {
	size = 140

	// Field (0) 'State'
	if i.State == nil {
		i.State = new(State)
	}
	size += i.State.SizeSSZ()

	// Field (3) 'StartValue'
	if i.StartValue == nil {
		i.StartValue = new(types.ConsensusData)
	}
	size += i.StartValue.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the Instance object
func (i *Instance) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(i)
}

// HashTreeRootWith ssz hashes the Instance object with a hasher
func (i *Instance) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'State'
	if err = i.State.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Share'
	if i.Share == nil {
		i.Share = new(types.Share)
	}
	if err = i.Share.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Identifier'
	hh.PutBytes(i.Identifier[:])

	// Field (3) 'StartValue'
	if err = i.StartValue.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Instance object
func (i *Instance) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(i)
}
