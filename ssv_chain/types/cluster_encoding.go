// Code generated by fastssz. DO NOT EDIT.
// Hash: 23c92cd26da6471eac3a9bd166a322fadba24f5edcfcfe0e69bed8fc23943ea6
// Version: 0.1.2
package types

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/ssv_chain/common"
)

// MarshalSSZ ssz marshals the Share object
func (s *Share) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the Share object to a target array
func (s *Share) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(12)

	// Offset (0) 'Key'
	dst = ssz.WriteOffset(dst, offset)
	if s.Key == nil {
		s.Key = new(common.CryptoKey)
	}
	offset += s.Key.SizeSSZ()

	// Field (1) 'OperatorID'
	dst = ssz.MarshalUint64(dst, s.OperatorID)

	// Field (0) 'Key'
	if dst, err = s.Key.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Share object
func (s *Share) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 12 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Key'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 12 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'OperatorID'
	s.OperatorID = ssz.UnmarshallUint64(buf[4:12])

	// Field (0) 'Key'
	{
		buf = tail[o0:]
		if s.Key == nil {
			s.Key = new(common.CryptoKey)
		}
		if err = s.Key.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Share object
func (s *Share) SizeSSZ() (size int) {
	size = 12

	// Field (0) 'Key'
	if s.Key == nil {
		s.Key = new(common.CryptoKey)
	}
	size += s.Key.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the Share object
func (s *Share) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the Share object with a hasher
func (s *Share) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Key'
	if err = s.Key.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'OperatorID'
	hh.PutUint64(s.OperatorID)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Share object
func (s *Share) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}

// MarshalSSZ ssz marshals the ClusterInstance object
func (c *ClusterInstance) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the ClusterInstance object to a target array
func (c *ClusterInstance) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(8)

	// Offset (0) 'Keys'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(c.Keys); ii++ {
		offset += 4
		offset += c.Keys[ii].SizeSSZ()
	}

	// Offset (1) 'Metadata'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(c.Metadata)

	// Field (0) 'Keys'
	if size := len(c.Keys); size > 13 {
		err = ssz.ErrListTooBigFn("ClusterInstance.Keys", size, 13)
		return
	}
	{
		offset = 4 * len(c.Keys)
		for ii := 0; ii < len(c.Keys); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += c.Keys[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(c.Keys); ii++ {
		if dst, err = c.Keys[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (1) 'Metadata'
	if size := len(c.Metadata); size > 2048 {
		err = ssz.ErrBytesLengthFn("ClusterInstance.Metadata", size, 2048)
		return
	}
	dst = append(dst, c.Metadata...)

	return
}

// UnmarshalSSZ ssz unmarshals the ClusterInstance object
func (c *ClusterInstance) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 8 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1 uint64

	// Offset (0) 'Keys'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 8 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'Metadata'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (0) 'Keys'
	{
		buf = tail[o0:o1]
		num, err := ssz.DecodeDynamicLength(buf, 13)
		if err != nil {
			return err
		}
		c.Keys = make([]*common.CryptoKey, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if c.Keys[indx] == nil {
				c.Keys[indx] = new(common.CryptoKey)
			}
			if err = c.Keys[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (1) 'Metadata'
	{
		buf = tail[o1:]
		if len(buf) > 2048 {
			return ssz.ErrBytesLength
		}
		if cap(c.Metadata) == 0 {
			c.Metadata = make([]byte, 0, len(buf))
		}
		c.Metadata = append(c.Metadata, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ClusterInstance object
func (c *ClusterInstance) SizeSSZ() (size int) {
	size = 8

	// Field (0) 'Keys'
	for ii := 0; ii < len(c.Keys); ii++ {
		size += 4
		size += c.Keys[ii].SizeSSZ()
	}

	// Field (1) 'Metadata'
	size += len(c.Metadata)

	return
}

// HashTreeRoot ssz hashes the ClusterInstance object
func (c *ClusterInstance) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the ClusterInstance object with a hasher
func (c *ClusterInstance) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Keys'
	{
		subIndx := hh.Index()
		num := uint64(len(c.Keys))
		if num > 13 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range c.Keys {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 13)
	}

	// Field (1) 'Metadata'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(c.Metadata))
		if byteLen > 2048 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(c.Metadata)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (2048+31)/32)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ClusterInstance object
func (c *ClusterInstance) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(c)
}

// MarshalSSZ ssz marshals the Cluster object
func (c *Cluster) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the Cluster object to a target array
func (c *Cluster) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(36)

	// Field (0) 'ID'
	dst = ssz.MarshalUint64(dst, c.ID)

	// Field (1) 'ModuleID'
	dst = ssz.MarshalUint64(dst, c.ModuleID)

	// Offset (2) 'Address'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(c.Address)

	// Offset (3) 'Operators'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(c.Operators) * 8

	// Field (4) 'FaultyNodes'
	dst = ssz.MarshalUint64(dst, c.FaultyNodes)

	// Offset (5) 'Instances'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(c.Instances); ii++ {
		offset += 4
		offset += c.Instances[ii].SizeSSZ()
	}

	// Field (2) 'Address'
	if size := len(c.Address); size > 128 {
		err = ssz.ErrBytesLengthFn("Cluster.Address", size, 128)
		return
	}
	dst = append(dst, c.Address...)

	// Field (3) 'Operators'
	if size := len(c.Operators); size > 13 {
		err = ssz.ErrListTooBigFn("Cluster.Operators", size, 13)
		return
	}
	for ii := 0; ii < len(c.Operators); ii++ {
		dst = ssz.MarshalUint64(dst, c.Operators[ii])
	}

	// Field (5) 'Instances'
	if size := len(c.Instances); size > 500 {
		err = ssz.ErrListTooBigFn("Cluster.Instances", size, 500)
		return
	}
	{
		offset = 4 * len(c.Instances)
		for ii := 0; ii < len(c.Instances); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += c.Instances[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(c.Instances); ii++ {
		if dst, err = c.Instances[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Cluster object
func (c *Cluster) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 36 {
		return ssz.ErrSize
	}

	tail := buf
	var o2, o3, o5 uint64

	// Field (0) 'ID'
	c.ID = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'ModuleID'
	c.ModuleID = ssz.UnmarshallUint64(buf[8:16])

	// Offset (2) 'Address'
	if o2 = ssz.ReadOffset(buf[16:20]); o2 > size {
		return ssz.ErrOffset
	}

	if o2 < 36 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (3) 'Operators'
	if o3 = ssz.ReadOffset(buf[20:24]); o3 > size || o2 > o3 {
		return ssz.ErrOffset
	}

	// Field (4) 'FaultyNodes'
	c.FaultyNodes = ssz.UnmarshallUint64(buf[24:32])

	// Offset (5) 'Instances'
	if o5 = ssz.ReadOffset(buf[32:36]); o5 > size || o3 > o5 {
		return ssz.ErrOffset
	}

	// Field (2) 'Address'
	{
		buf = tail[o2:o3]
		if len(buf) > 128 {
			return ssz.ErrBytesLength
		}
		if cap(c.Address) == 0 {
			c.Address = make([]byte, 0, len(buf))
		}
		c.Address = append(c.Address, buf...)
	}

	// Field (3) 'Operators'
	{
		buf = tail[o3:o5]
		num, err := ssz.DivideInt2(len(buf), 8, 13)
		if err != nil {
			return err
		}
		c.Operators = ssz.ExtendUint64(c.Operators, num)
		for ii := 0; ii < num; ii++ {
			c.Operators[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}

	// Field (5) 'Instances'
	{
		buf = tail[o5:]
		num, err := ssz.DecodeDynamicLength(buf, 500)
		if err != nil {
			return err
		}
		c.Instances = make([]*ClusterInstance, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if c.Instances[indx] == nil {
				c.Instances[indx] = new(ClusterInstance)
			}
			if err = c.Instances[indx].UnmarshalSSZ(buf); err != nil {
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

// SizeSSZ returns the ssz encoded size in bytes for the Cluster object
func (c *Cluster) SizeSSZ() (size int) {
	size = 36

	// Field (2) 'Address'
	size += len(c.Address)

	// Field (3) 'Operators'
	size += len(c.Operators) * 8

	// Field (5) 'Instances'
	for ii := 0; ii < len(c.Instances); ii++ {
		size += 4
		size += c.Instances[ii].SizeSSZ()
	}

	return
}

// HashTreeRoot ssz hashes the Cluster object
func (c *Cluster) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the Cluster object with a hasher
func (c *Cluster) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ID'
	hh.PutUint64(c.ID)

	// Field (1) 'ModuleID'
	hh.PutUint64(c.ModuleID)

	// Field (2) 'Address'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(c.Address))
		if byteLen > 128 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(c.Address)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (128+31)/32)
	}

	// Field (3) 'Operators'
	{
		if size := len(c.Operators); size > 13 {
			err = ssz.ErrListTooBigFn("Cluster.Operators", size, 13)
			return
		}
		subIndx := hh.Index()
		for _, i := range c.Operators {
			hh.AppendUint64(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(c.Operators))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(13, numItems, 8))
	}

	// Field (4) 'FaultyNodes'
	hh.PutUint64(c.FaultyNodes)

	// Field (5) 'Instances'
	{
		subIndx := hh.Index()
		num := uint64(len(c.Instances))
		if num > 500 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range c.Instances {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 500)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Cluster object
func (c *Cluster) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(c)
}
