// Code generated by fastssz. DO NOT EDIT.
// Hash: 6ee701d07d4d5de618606a297aadbc30c4556ff9453b6d2617e03e24869e8658
// Version: 0.1.3
package runner

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/ssv"
)

// MarshalSSZ ssz marshals the SpecTest object
func (s *SpecTest) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SpecTest object to a target array
func (s *SpecTest) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(20)

	// Offset (0) 'Pre'
	dst = ssz.WriteOffset(dst, offset)
	if s.Pre == nil {
		s.Pre = new(ssv.Runner)
	}
	offset += s.Pre.SizeSSZ()

	// Offset (1) 'Post'
	dst = ssz.WriteOffset(dst, offset)
	if s.Post == nil {
		s.Post = new(ssv.Runner)
	}
	offset += s.Post.SizeSSZ()

	// Field (2) 'Role'
	dst = ssz.MarshalUint64(dst, s.Role)

	// Offset (3) 'Messages'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(s.Messages); ii++ {
		offset += 4
		offset += s.Messages[ii].SizeSSZ()
	}

	// Field (0) 'Pre'
	if dst, err = s.Pre.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Post'
	if dst, err = s.Post.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (3) 'Messages'
	if size := len(s.Messages); size > 256 {
		err = ssz.ErrListTooBigFn("SpecTest.Messages", size, 256)
		return
	}
	{
		offset = 4 * len(s.Messages)
		for ii := 0; ii < len(s.Messages); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += s.Messages[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(s.Messages); ii++ {
		if dst, err = s.Messages[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the SpecTest object
func (s *SpecTest) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 20 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1, o3 uint64

	// Offset (0) 'Pre'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 20 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'Post'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (2) 'Role'
	s.Role = ssz.UnmarshallUint64(buf[8:16])

	// Offset (3) 'Messages'
	if o3 = ssz.ReadOffset(buf[16:20]); o3 > size || o1 > o3 {
		return ssz.ErrOffset
	}

	// Field (0) 'Pre'
	{
		buf = tail[o0:o1]
		if s.Pre == nil {
			s.Pre = new(ssv.Runner)
		}
		if err = s.Pre.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'Post'
	{
		buf = tail[o1:o3]
		if s.Post == nil {
			s.Post = new(ssv.Runner)
		}
		if err = s.Post.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (3) 'Messages'
	{
		buf = tail[o3:]
		num, err := ssz.DecodeDynamicLength(buf, 256)
		if err != nil {
			return err
		}
		s.Messages = make([]*p2p.Message, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if s.Messages[indx] == nil {
				s.Messages[indx] = new(p2p.Message)
			}
			if err = s.Messages[indx].UnmarshalSSZ(buf); err != nil {
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

// SizeSSZ returns the ssz encoded size in bytes for the SpecTest object
func (s *SpecTest) SizeSSZ() (size int) {
	size = 20

	// Field (0) 'Pre'
	if s.Pre == nil {
		s.Pre = new(ssv.Runner)
	}
	size += s.Pre.SizeSSZ()

	// Field (1) 'Post'
	if s.Post == nil {
		s.Post = new(ssv.Runner)
	}
	size += s.Post.SizeSSZ()

	// Field (3) 'Messages'
	for ii := 0; ii < len(s.Messages); ii++ {
		size += 4
		size += s.Messages[ii].SizeSSZ()
	}

	return
}

// HashTreeRoot ssz hashes the SpecTest object
func (s *SpecTest) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SpecTest object with a hasher
func (s *SpecTest) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Pre'
	if err = s.Pre.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Post'
	if err = s.Post.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Role'
	hh.PutUint64(s.Role)

	// Field (3) 'Messages'
	{
		subIndx := hh.Index()
		num := uint64(len(s.Messages))
		if num > 256 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.Messages {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 256)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SpecTest object
func (s *SpecTest) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
