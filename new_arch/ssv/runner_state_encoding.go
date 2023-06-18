// Code generated by fastssz. DO NOT EDIT.
// Hash: f607ea79036333cb2ddc9220c85bdb3a37877f51a87737cfd5c84e1729a5ef99
// Version: 0.1.2
package ssv

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/new_arch/types"
)

// MarshalSSZ ssz marshals the State object
func (s *State) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the State object to a target array
func (s *State) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(72)

	// Offset (0) 'PartialSignatures'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(s.PartialSignatures); ii++ {
		offset += 4
		offset += s.PartialSignatures[ii].SizeSSZ()
	}

	// Field (1) 'StartingDuty'
	if s.StartingDuty == nil {
		s.StartingDuty = new(types.Duty)
	}
	if dst, err = s.StartingDuty.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (2) 'DecidedData'
	dst = ssz.WriteOffset(dst, offset)
	if s.DecidedData == nil {
		s.DecidedData = new(types.ConsensusData)
	}
	offset += s.DecidedData.SizeSSZ()

	// Field (0) 'PartialSignatures'
	if size := len(s.PartialSignatures); size > 256 {
		err = ssz.ErrListTooBigFn("State.PartialSignatures", size, 256)
		return
	}
	{
		offset = 4 * len(s.PartialSignatures)
		for ii := 0; ii < len(s.PartialSignatures); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += s.PartialSignatures[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(s.PartialSignatures); ii++ {
		if dst, err = s.PartialSignatures[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (2) 'DecidedData'
	if dst, err = s.DecidedData.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the State object
func (s *State) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 72 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o2 uint64

	// Offset (0) 'PartialSignatures'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 72 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'StartingDuty'
	if s.StartingDuty == nil {
		s.StartingDuty = new(types.Duty)
	}
	if err = s.StartingDuty.UnmarshalSSZ(buf[4:68]); err != nil {
		return err
	}

	// Offset (2) 'DecidedData'
	if o2 = ssz.ReadOffset(buf[68:72]); o2 > size || o0 > o2 {
		return ssz.ErrOffset
	}

	// Field (0) 'PartialSignatures'
	{
		buf = tail[o0:o2]
		num, err := ssz.DecodeDynamicLength(buf, 256)
		if err != nil {
			return err
		}
		s.PartialSignatures = make([]*types.SignedPartialSignatureMessages, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if s.PartialSignatures[indx] == nil {
				s.PartialSignatures[indx] = new(types.SignedPartialSignatureMessages)
			}
			if err = s.PartialSignatures[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (2) 'DecidedData'
	{
		buf = tail[o2:]
		if s.DecidedData == nil {
			s.DecidedData = new(types.ConsensusData)
		}
		if err = s.DecidedData.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the State object
func (s *State) SizeSSZ() (size int) {
	size = 72

	// Field (0) 'PartialSignatures'
	for ii := 0; ii < len(s.PartialSignatures); ii++ {
		size += 4
		size += s.PartialSignatures[ii].SizeSSZ()
	}

	// Field (2) 'DecidedData'
	if s.DecidedData == nil {
		s.DecidedData = new(types.ConsensusData)
	}
	size += s.DecidedData.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the State object
func (s *State) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the State object with a hasher
func (s *State) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'PartialSignatures'
	{
		subIndx := hh.Index()
		num := uint64(len(s.PartialSignatures))
		if num > 256 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.PartialSignatures {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 256)
	}

	// Field (1) 'StartingDuty'
	if s.StartingDuty == nil {
		s.StartingDuty = new(types.Duty)
	}
	if err = s.StartingDuty.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'DecidedData'
	if err = s.DecidedData.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the State object
func (s *State) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
