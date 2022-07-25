// Code generated by fastssz. DO NOT EDIT.
// Hash: 469f5470938b23275dfc67721365de440dfe23c88c69ad48c7c19695fa4607ea
package ssv

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/ssz_encoding/qbft"
)

// MarshalSSZ ssz marshals the PartialSignature object
func (p *PartialSignature) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(p)
}

// MarshalSSZTo ssz marshals the PartialSignature object to a target array
func (p *PartialSignature) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	return
}

// UnmarshalSSZ ssz unmarshals the PartialSignature object
func (p *PartialSignature) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 0 {
		return ssz.ErrSize
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the PartialSignature object
func (p *PartialSignature) SizeSSZ() (size int) {
	size = 0
	return
}

// HashTreeRoot ssz hashes the PartialSignature object
func (p *PartialSignature) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(p)
}

// HashTreeRootWith ssz hashes the PartialSignature object with a hasher
func (p *PartialSignature) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the PartialSignature object
func (p *PartialSignature) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(p)
}

// MarshalSSZ ssz marshals the SignedPartialSignatures object
func (s *SignedPartialSignatures) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SignedPartialSignatures object to a target array
func (s *SignedPartialSignatures) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(60)

	// Field (0) 'ID'
	dst = append(dst, s.ID[:]...)

	// Offset (1) 'PartialSignatures'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(s.PartialSignatures) * 0

	// Offset (2) 'Justification'
	dst = ssz.WriteOffset(dst, offset)
	if s.Justification == nil {
		s.Justification = new(qbft.SignedMessageHeader)
	}
	offset += s.Justification.SizeSSZ()

	// Field (1) 'PartialSignatures'
	if size := len(s.PartialSignatures); size > 13 {
		err = ssz.ErrListTooBigFn("SignedPartialSignatures.PartialSignatures", size, 13)
		return
	}
	for ii := 0; ii < len(s.PartialSignatures); ii++ {
		if dst, err = s.PartialSignatures[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (2) 'Justification'
	if dst, err = s.Justification.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the SignedPartialSignatures object
func (s *SignedPartialSignatures) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 60 {
		return ssz.ErrSize
	}

	tail := buf
	var o1, o2 uint64

	// Field (0) 'ID'
	copy(s.ID[:], buf[0:52])

	// Offset (1) 'PartialSignatures'
	if o1 = ssz.ReadOffset(buf[52:56]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 60 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (2) 'Justification'
	if o2 = ssz.ReadOffset(buf[56:60]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Field (1) 'PartialSignatures'
	{
		buf = tail[o1:o2]
		num, err := ssz.DivideInt2(len(buf), 0, 13)
		if err != nil {
			return err
		}
		s.PartialSignatures = make([]*PartialSignature, num)
		for ii := 0; ii < num; ii++ {
			if s.PartialSignatures[ii] == nil {
				s.PartialSignatures[ii] = new(PartialSignature)
			}
			if err = s.PartialSignatures[ii].UnmarshalSSZ(buf[ii*0 : (ii+1)*0]); err != nil {
				return err
			}
		}
	}

	// Field (2) 'Justification'
	{
		buf = tail[o2:]
		if s.Justification == nil {
			s.Justification = new(qbft.SignedMessageHeader)
		}
		if err = s.Justification.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedPartialSignatures object
func (s *SignedPartialSignatures) SizeSSZ() (size int) {
	size = 60

	// Field (1) 'PartialSignatures'
	size += len(s.PartialSignatures) * 0

	// Field (2) 'Justification'
	if s.Justification == nil {
		s.Justification = new(qbft.SignedMessageHeader)
	}
	size += s.Justification.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the SignedPartialSignatures object
func (s *SignedPartialSignatures) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SignedPartialSignatures object with a hasher
func (s *SignedPartialSignatures) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ID'
	hh.PutBytes(s.ID[:])

	// Field (1) 'PartialSignatures'
	{
		subIndx := hh.Index()
		num := uint64(len(s.PartialSignatures))
		if num > 13 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.PartialSignatures {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 13)
	}

	// Field (2) 'Justification'
	if err = s.Justification.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SignedPartialSignatures object
func (s *SignedPartialSignatures) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}

// MarshalSSZ ssz marshals the SignedPartialSignatureHeader object
func (s *SignedPartialSignatureHeader) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SignedPartialSignatureHeader object to a target array
func (s *SignedPartialSignatureHeader) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(88)

	// Field (0) 'ID'
	dst = append(dst, s.ID[:]...)

	// Offset (1) 'PartialSignatures'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(s.PartialSignatures) * 0

	// Field (2) 'JustificationRoot'
	dst = append(dst, s.JustificationRoot[:]...)

	// Field (1) 'PartialSignatures'
	if size := len(s.PartialSignatures); size > 13 {
		err = ssz.ErrListTooBigFn("SignedPartialSignatureHeader.PartialSignatures", size, 13)
		return
	}
	for ii := 0; ii < len(s.PartialSignatures); ii++ {
		if dst, err = s.PartialSignatures[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the SignedPartialSignatureHeader object
func (s *SignedPartialSignatureHeader) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 88 {
		return ssz.ErrSize
	}

	tail := buf
	var o1 uint64

	// Field (0) 'ID'
	copy(s.ID[:], buf[0:52])

	// Offset (1) 'PartialSignatures'
	if o1 = ssz.ReadOffset(buf[52:56]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 88 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (2) 'JustificationRoot'
	copy(s.JustificationRoot[:], buf[56:88])

	// Field (1) 'PartialSignatures'
	{
		buf = tail[o1:]
		num, err := ssz.DivideInt2(len(buf), 0, 13)
		if err != nil {
			return err
		}
		s.PartialSignatures = make([]*PartialSignature, num)
		for ii := 0; ii < num; ii++ {
			if s.PartialSignatures[ii] == nil {
				s.PartialSignatures[ii] = new(PartialSignature)
			}
			if err = s.PartialSignatures[ii].UnmarshalSSZ(buf[ii*0 : (ii+1)*0]); err != nil {
				return err
			}
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedPartialSignatureHeader object
func (s *SignedPartialSignatureHeader) SizeSSZ() (size int) {
	size = 88

	// Field (1) 'PartialSignatures'
	size += len(s.PartialSignatures) * 0

	return
}

// HashTreeRoot ssz hashes the SignedPartialSignatureHeader object
func (s *SignedPartialSignatureHeader) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SignedPartialSignatureHeader object with a hasher
func (s *SignedPartialSignatureHeader) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ID'
	hh.PutBytes(s.ID[:])

	// Field (1) 'PartialSignatures'
	{
		subIndx := hh.Index()
		num := uint64(len(s.PartialSignatures))
		if num > 13 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.PartialSignatures {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 13)
	}

	// Field (2) 'JustificationRoot'
	hh.PutBytes(s.JustificationRoot[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SignedPartialSignatureHeader object
func (s *SignedPartialSignatureHeader) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
