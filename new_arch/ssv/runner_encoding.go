// Code generated by fastssz. DO NOT EDIT.
// Hash: b8c2a699450887ffe1987e34672545cf07e77a590cee3c2cd0da54bcd5a5d0eb
// Version: 0.1.3
package ssv

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/new_arch/types"
)

// MarshalSSZ ssz marshals the Runner object
func (r *Runner) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(r)
}

// MarshalSSZTo ssz marshals the Runner object to a target array
func (r *Runner) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(136)

	// Offset (0) 'State'
	dst = ssz.WriteOffset(dst, offset)
	if r.State == nil {
		r.State = new(State)
	}
	offset += r.State.SizeSSZ()

	// Field (1) 'Share'
	if r.Share == nil {
		r.Share = new(types.Share)
	}
	if dst, err = r.Share.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'Identifier'
	dst = append(dst, r.Identifier[:]...)

	// Field (0) 'State'
	if dst, err = r.State.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Runner object
func (r *Runner) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 136 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'State'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 136 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'Share'
	if r.Share == nil {
		r.Share = new(types.Share)
	}
	if err = r.Share.UnmarshalSSZ(buf[4:80]); err != nil {
		return err
	}

	// Field (2) 'Identifier'
	copy(r.Identifier[:], buf[80:136])

	// Field (0) 'State'
	{
		buf = tail[o0:]
		if r.State == nil {
			r.State = new(State)
		}
		if err = r.State.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Runner object
func (r *Runner) SizeSSZ() (size int) {
	size = 136

	// Field (0) 'State'
	if r.State == nil {
		r.State = new(State)
	}
	size += r.State.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the Runner object
func (r *Runner) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(r)
}

// HashTreeRootWith ssz hashes the Runner object with a hasher
func (r *Runner) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'State'
	if err = r.State.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Share'
	if r.Share == nil {
		r.Share = new(types.Share)
	}
	if err = r.Share.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Identifier'
	hh.PutBytes(r.Identifier[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Runner object
func (r *Runner) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(r)
}
