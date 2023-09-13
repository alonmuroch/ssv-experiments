// Code generated by fastssz. DO NOT EDIT.
// Hash: 10e8feeefa63660b2a7c6adea05db1884d5628aab04ecf0d5ecae0e590636e87
// Version: 0.1.2
package types

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the QBFT object
func (q *QBFT) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(q)
}

// MarshalSSZTo ssz marshals the QBFT object to a target array
func (q *QBFT) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(33)

	// Field (0) 'Round'
	dst = ssz.MarshalUint64(dst, q.Round)

	// Field (1) 'Height'
	dst = ssz.MarshalUint64(dst, q.Height)

	// Field (2) 'PreparedRound'
	dst = ssz.MarshalUint64(dst, q.PreparedRound)

	// Offset (3) 'ProposalAcceptedForCurrentRound'
	dst = ssz.WriteOffset(dst, offset)
	if q.ProposalAcceptedForCurrentRound == nil {
		q.ProposalAcceptedForCurrentRound = new(QBFTSignedMessage)
	}
	offset += q.ProposalAcceptedForCurrentRound.SizeSSZ()

	// Offset (4) 'Messages'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(q.Messages); ii++ {
		offset += 4
		offset += q.Messages[ii].SizeSSZ()
	}

	// Field (5) 'Stopped'
	dst = ssz.MarshalBool(dst, q.Stopped)

	// Field (3) 'ProposalAcceptedForCurrentRound'
	if dst, err = q.ProposalAcceptedForCurrentRound.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (4) 'Messages'
	if size := len(q.Messages); size > 256 {
		err = ssz.ErrListTooBigFn("QBFT.Messages", size, 256)
		return
	}
	{
		offset = 4 * len(q.Messages)
		for ii := 0; ii < len(q.Messages); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += q.Messages[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(q.Messages); ii++ {
		if dst, err = q.Messages[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the QBFT object
func (q *QBFT) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 33 {
		return ssz.ErrSize
	}

	tail := buf
	var o3, o4 uint64

	// Field (0) 'Round'
	q.Round = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Height'
	q.Height = ssz.UnmarshallUint64(buf[8:16])

	// Field (2) 'PreparedRound'
	q.PreparedRound = ssz.UnmarshallUint64(buf[16:24])

	// Offset (3) 'ProposalAcceptedForCurrentRound'
	if o3 = ssz.ReadOffset(buf[24:28]); o3 > size {
		return ssz.ErrOffset
	}

	if o3 < 33 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (4) 'Messages'
	if o4 = ssz.ReadOffset(buf[28:32]); o4 > size || o3 > o4 {
		return ssz.ErrOffset
	}

	// Field (5) 'Stopped'
	q.Stopped = ssz.UnmarshalBool(buf[32:33])

	// Field (3) 'ProposalAcceptedForCurrentRound'
	{
		buf = tail[o3:o4]
		if q.ProposalAcceptedForCurrentRound == nil {
			q.ProposalAcceptedForCurrentRound = new(QBFTSignedMessage)
		}
		if err = q.ProposalAcceptedForCurrentRound.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (4) 'Messages'
	{
		buf = tail[o4:]
		num, err := ssz.DecodeDynamicLength(buf, 256)
		if err != nil {
			return err
		}
		q.Messages = make([]*QBFTSignedMessage, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if q.Messages[indx] == nil {
				q.Messages[indx] = new(QBFTSignedMessage)
			}
			if err = q.Messages[indx].UnmarshalSSZ(buf); err != nil {
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

// SizeSSZ returns the ssz encoded size in bytes for the QBFT object
func (q *QBFT) SizeSSZ() (size int) {
	size = 33

	// Field (3) 'ProposalAcceptedForCurrentRound'
	if q.ProposalAcceptedForCurrentRound == nil {
		q.ProposalAcceptedForCurrentRound = new(QBFTSignedMessage)
	}
	size += q.ProposalAcceptedForCurrentRound.SizeSSZ()

	// Field (4) 'Messages'
	for ii := 0; ii < len(q.Messages); ii++ {
		size += 4
		size += q.Messages[ii].SizeSSZ()
	}

	return
}

// HashTreeRoot ssz hashes the QBFT object
func (q *QBFT) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(q)
}

// HashTreeRootWith ssz hashes the QBFT object with a hasher
func (q *QBFT) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Round'
	hh.PutUint64(q.Round)

	// Field (1) 'Height'
	hh.PutUint64(q.Height)

	// Field (2) 'PreparedRound'
	hh.PutUint64(q.PreparedRound)

	// Field (3) 'ProposalAcceptedForCurrentRound'
	if err = q.ProposalAcceptedForCurrentRound.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (4) 'Messages'
	{
		subIndx := hh.Index()
		num := uint64(len(q.Messages))
		if num > 256 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range q.Messages {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 256)
	}

	// Field (5) 'Stopped'
	hh.PutBool(q.Stopped)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the QBFT object
func (q *QBFT) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(q)
}

// MarshalSSZ ssz marshals the State object
func (s *State) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the State object to a target array
func (s *State) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(12)

	// Offset (0) 'PartialSignatures'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(s.PartialSignatures); ii++ {
		offset += 4
		offset += s.PartialSignatures[ii].SizeSSZ()
	}

	// Offset (1) 'QBFT'
	dst = ssz.WriteOffset(dst, offset)
	if s.QBFT == nil {
		s.QBFT = new(QBFT)
	}
	offset += s.QBFT.SizeSSZ()

	// Offset (2) 'StartingDuty'
	dst = ssz.WriteOffset(dst, offset)
	if s.StartingDuty == nil {
		s.StartingDuty = new(Duty)
	}
	offset += s.StartingDuty.SizeSSZ()

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

	// Field (1) 'QBFT'
	if dst, err = s.QBFT.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'StartingDuty'
	if dst, err = s.StartingDuty.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the State object
func (s *State) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 12 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1, o2 uint64

	// Offset (0) 'PartialSignatures'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 12 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'QBFT'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Offset (2) 'StartingDuty'
	if o2 = ssz.ReadOffset(buf[8:12]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Field (0) 'PartialSignatures'
	{
		buf = tail[o0:o1]
		num, err := ssz.DecodeDynamicLength(buf, 256)
		if err != nil {
			return err
		}
		s.PartialSignatures = make([]*SignedPartialSignatureMessages, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if s.PartialSignatures[indx] == nil {
				s.PartialSignatures[indx] = new(SignedPartialSignatureMessages)
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

	// Field (1) 'QBFT'
	{
		buf = tail[o1:o2]
		if s.QBFT == nil {
			s.QBFT = new(QBFT)
		}
		if err = s.QBFT.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (2) 'StartingDuty'
	{
		buf = tail[o2:]
		if s.StartingDuty == nil {
			s.StartingDuty = new(Duty)
		}
		if err = s.StartingDuty.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the State object
func (s *State) SizeSSZ() (size int) {
	size = 12

	// Field (0) 'PartialSignatures'
	for ii := 0; ii < len(s.PartialSignatures); ii++ {
		size += 4
		size += s.PartialSignatures[ii].SizeSSZ()
	}

	// Field (1) 'QBFT'
	if s.QBFT == nil {
		s.QBFT = new(QBFT)
	}
	size += s.QBFT.SizeSSZ()

	// Field (2) 'StartingDuty'
	if s.StartingDuty == nil {
		s.StartingDuty = new(Duty)
	}
	size += s.StartingDuty.SizeSSZ()

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

	// Field (1) 'QBFT'
	if err = s.QBFT.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'StartingDuty'
	if err = s.StartingDuty.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the State object
func (s *State) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
