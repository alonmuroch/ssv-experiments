// Code generated by fastssz. DO NOT EDIT.
// Hash: ee3fcaf624433aaebd200f172b5a1026de7e7dd2cbd0716e0fea3c397a1c5f30
// Version: 0.1.2
package types

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the State object
func (s *State) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the State object to a target array
func (s *State) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(64)

	// Field (0) 'Domain'
	dst = append(dst, s.Domain[:]...)

	// Field (1) 'BlockHeight'
	dst = ssz.MarshalUint64(dst, s.BlockHeight)

	// Field (2) 'LatestBlockHeaderHash'
	if size := len(s.LatestBlockHeaderHash); size != 32 {
		err = ssz.ErrBytesLengthFn("State.LatestBlockHeaderHash", size, 32)
		return
	}
	dst = append(dst, s.LatestBlockHeaderHash...)

	// Offset (3) 'Validators'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(s.Validators); ii++ {
		offset += 4
		offset += s.Validators[ii].SizeSSZ()
	}

	// Offset (4) 'Accounts'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(s.Accounts); ii++ {
		offset += 4
		offset += s.Accounts[ii].SizeSSZ()
	}

	// Offset (5) 'Clusters'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(s.Clusters); ii++ {
		offset += 4
		offset += s.Clusters[ii].SizeSSZ()
	}

	// Offset (6) 'Operators'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(s.Operators); ii++ {
		offset += 4
		offset += s.Operators[ii].SizeSSZ()
	}

	// Offset (7) 'Modules'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(s.Modules); ii++ {
		offset += 4
		offset += s.Modules[ii].SizeSSZ()
	}

	// Field (3) 'Validators'
	if size := len(s.Validators); size > 128 {
		err = ssz.ErrListTooBigFn("State.Validators", size, 128)
		return
	}
	{
		offset = 4 * len(s.Validators)
		for ii := 0; ii < len(s.Validators); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += s.Validators[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(s.Validators); ii++ {
		if dst, err = s.Validators[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (4) 'Accounts'
	if size := len(s.Accounts); size > 65536 {
		err = ssz.ErrListTooBigFn("State.Accounts", size, 65536)
		return
	}
	{
		offset = 4 * len(s.Accounts)
		for ii := 0; ii < len(s.Accounts); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += s.Accounts[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(s.Accounts); ii++ {
		if dst, err = s.Accounts[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (5) 'Clusters'
	if size := len(s.Clusters); size > 1048576 {
		err = ssz.ErrListTooBigFn("State.Clusters", size, 1048576)
		return
	}
	{
		offset = 4 * len(s.Clusters)
		for ii := 0; ii < len(s.Clusters); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += s.Clusters[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(s.Clusters); ii++ {
		if dst, err = s.Clusters[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (6) 'Operators'
	if size := len(s.Operators); size > 65536 {
		err = ssz.ErrListTooBigFn("State.Operators", size, 65536)
		return
	}
	{
		offset = 4 * len(s.Operators)
		for ii := 0; ii < len(s.Operators); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += s.Operators[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(s.Operators); ii++ {
		if dst, err = s.Operators[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (7) 'Modules'
	if size := len(s.Modules); size > 65536 {
		err = ssz.ErrListTooBigFn("State.Modules", size, 65536)
		return
	}
	{
		offset = 4 * len(s.Modules)
		for ii := 0; ii < len(s.Modules); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += s.Modules[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(s.Modules); ii++ {
		if dst, err = s.Modules[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the State object
func (s *State) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 64 {
		return ssz.ErrSize
	}

	tail := buf
	var o3, o4, o5, o6, o7 uint64

	// Field (0) 'Domain'
	copy(s.Domain[:], buf[0:4])

	// Field (1) 'BlockHeight'
	s.BlockHeight = ssz.UnmarshallUint64(buf[4:12])

	// Field (2) 'LatestBlockHeaderHash'
	if cap(s.LatestBlockHeaderHash) == 0 {
		s.LatestBlockHeaderHash = make([]byte, 0, len(buf[12:44]))
	}
	s.LatestBlockHeaderHash = append(s.LatestBlockHeaderHash, buf[12:44]...)

	// Offset (3) 'Validators'
	if o3 = ssz.ReadOffset(buf[44:48]); o3 > size {
		return ssz.ErrOffset
	}

	if o3 < 64 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (4) 'Accounts'
	if o4 = ssz.ReadOffset(buf[48:52]); o4 > size || o3 > o4 {
		return ssz.ErrOffset
	}

	// Offset (5) 'Clusters'
	if o5 = ssz.ReadOffset(buf[52:56]); o5 > size || o4 > o5 {
		return ssz.ErrOffset
	}

	// Offset (6) 'Operators'
	if o6 = ssz.ReadOffset(buf[56:60]); o6 > size || o5 > o6 {
		return ssz.ErrOffset
	}

	// Offset (7) 'Modules'
	if o7 = ssz.ReadOffset(buf[60:64]); o7 > size || o6 > o7 {
		return ssz.ErrOffset
	}

	// Field (3) 'Validators'
	{
		buf = tail[o3:o4]
		num, err := ssz.DecodeDynamicLength(buf, 128)
		if err != nil {
			return err
		}
		s.Validators = make([]*Validator, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if s.Validators[indx] == nil {
				s.Validators[indx] = new(Validator)
			}
			if err = s.Validators[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (4) 'Accounts'
	{
		buf = tail[o4:o5]
		num, err := ssz.DecodeDynamicLength(buf, 65536)
		if err != nil {
			return err
		}
		s.Accounts = make([]*Account, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if s.Accounts[indx] == nil {
				s.Accounts[indx] = new(Account)
			}
			if err = s.Accounts[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (5) 'Clusters'
	{
		buf = tail[o5:o6]
		num, err := ssz.DecodeDynamicLength(buf, 1048576)
		if err != nil {
			return err
		}
		s.Clusters = make([]*Cluster, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if s.Clusters[indx] == nil {
				s.Clusters[indx] = new(Cluster)
			}
			if err = s.Clusters[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (6) 'Operators'
	{
		buf = tail[o6:o7]
		num, err := ssz.DecodeDynamicLength(buf, 65536)
		if err != nil {
			return err
		}
		s.Operators = make([]*Operator, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if s.Operators[indx] == nil {
				s.Operators[indx] = new(Operator)
			}
			if err = s.Operators[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (7) 'Modules'
	{
		buf = tail[o7:]
		num, err := ssz.DecodeDynamicLength(buf, 65536)
		if err != nil {
			return err
		}
		s.Modules = make([]*Module, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if s.Modules[indx] == nil {
				s.Modules[indx] = new(Module)
			}
			if err = s.Modules[indx].UnmarshalSSZ(buf); err != nil {
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

// SizeSSZ returns the ssz encoded size in bytes for the State object
func (s *State) SizeSSZ() (size int) {
	size = 64

	// Field (3) 'Validators'
	for ii := 0; ii < len(s.Validators); ii++ {
		size += 4
		size += s.Validators[ii].SizeSSZ()
	}

	// Field (4) 'Accounts'
	for ii := 0; ii < len(s.Accounts); ii++ {
		size += 4
		size += s.Accounts[ii].SizeSSZ()
	}

	// Field (5) 'Clusters'
	for ii := 0; ii < len(s.Clusters); ii++ {
		size += 4
		size += s.Clusters[ii].SizeSSZ()
	}

	// Field (6) 'Operators'
	for ii := 0; ii < len(s.Operators); ii++ {
		size += 4
		size += s.Operators[ii].SizeSSZ()
	}

	// Field (7) 'Modules'
	for ii := 0; ii < len(s.Modules); ii++ {
		size += 4
		size += s.Modules[ii].SizeSSZ()
	}

	return
}

// HashTreeRoot ssz hashes the State object
func (s *State) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the State object with a hasher
func (s *State) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Domain'
	hh.PutBytes(s.Domain[:])

	// Field (1) 'BlockHeight'
	hh.PutUint64(s.BlockHeight)

	// Field (2) 'LatestBlockHeaderHash'
	if size := len(s.LatestBlockHeaderHash); size != 32 {
		err = ssz.ErrBytesLengthFn("State.LatestBlockHeaderHash", size, 32)
		return
	}
	hh.PutBytes(s.LatestBlockHeaderHash)

	// Field (3) 'Validators'
	{
		subIndx := hh.Index()
		num := uint64(len(s.Validators))
		if num > 128 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.Validators {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 128)
	}

	// Field (4) 'Accounts'
	{
		subIndx := hh.Index()
		num := uint64(len(s.Accounts))
		if num > 65536 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.Accounts {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 65536)
	}

	// Field (5) 'Clusters'
	{
		subIndx := hh.Index()
		num := uint64(len(s.Clusters))
		if num > 1048576 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.Clusters {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 1048576)
	}

	// Field (6) 'Operators'
	{
		subIndx := hh.Index()
		num := uint64(len(s.Operators))
		if num > 65536 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.Operators {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 65536)
	}

	// Field (7) 'Modules'
	{
		subIndx := hh.Index()
		num := uint64(len(s.Modules))
		if num > 65536 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.Modules {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 65536)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the State object
func (s *State) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
