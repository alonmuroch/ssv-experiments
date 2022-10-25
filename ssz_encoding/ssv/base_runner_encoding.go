// Code generated by fastssz. DO NOT EDIT.
// Hash: fb2502da76b6d23914aa4bf7db02c1390da5a1fce8785c22a8ea603aee78198c
package ssv

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BaseRunner object
func (b *BaseRunner) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BaseRunner object to a target array
func (b *BaseRunner) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(20)

	// Offset (0) 'State'
	dst = ssz.WriteOffset(dst, offset)
	if b.State == nil {
		b.State = new(State)
	}
	offset += b.State.SizeSSZ()

	// Offset (1) 'Share'
	dst = ssz.WriteOffset(dst, offset)
	offset += b.Share.SizeSSZ()

	// Offset (2) 'QBFTController'
	dst = ssz.WriteOffset(dst, offset)
	offset += b.QBFTController.SizeSSZ()

	// Field (3) 'BeaconNetwork'
	dst = append(dst, b.BeaconNetwork[:]...)

	// Field (4) 'BeaconRole'
	dst = append(dst, b.BeaconRole[:]...)

	// Field (0) 'State'
	if dst, err = b.State.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Share'
	if dst, err = b.Share.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'QBFTController'
	if dst, err = b.QBFTController.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BaseRunner object
func (b *BaseRunner) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 20 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1, o2 uint64

	// Offset (0) 'State'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 20 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'Share'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Offset (2) 'QBFTController'
	if o2 = ssz.ReadOffset(buf[8:12]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Field (3) 'BeaconNetwork'
	copy(b.BeaconNetwork[:], buf[12:16])

	// Field (4) 'BeaconRole'
	copy(b.BeaconRole[:], buf[16:20])

	// Field (0) 'State'
	{
		buf = tail[o0:o1]
		if b.State == nil {
			b.State = new(State)
		}
		if err = b.State.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'Share'
	{
		buf = tail[o1:o2]
		if err = b.Share.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (2) 'QBFTController'
	{
		buf = tail[o2:]
		if err = b.QBFTController.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BaseRunner object
func (b *BaseRunner) SizeSSZ() (size int) {
	size = 20

	// Field (0) 'State'
	if b.State == nil {
		b.State = new(State)
	}
	size += b.State.SizeSSZ()

	// Field (1) 'Share'
	size += b.Share.SizeSSZ()

	// Field (2) 'QBFTController'
	size += b.QBFTController.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the BaseRunner object
func (b *BaseRunner) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BaseRunner object with a hasher
func (b *BaseRunner) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'State'
	if err = b.State.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Share'
	if err = b.Share.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'QBFTController'
	if err = b.QBFTController.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (3) 'BeaconNetwork'
	hh.PutBytes(b.BeaconNetwork[:])

	// Field (4) 'BeaconRole'
	hh.PutBytes(b.BeaconRole[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BaseRunner object
func (b *BaseRunner) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
