// Code generated by fastssz. DO NOT EDIT.
// Hash: 4638b4f8beb4a4369831c32f3741b6b971c49a7c1c342aa2f9ba1e570a67e31b
// Version: 0.1.2
package types

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Configure object
func (c *Configure) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the Configure object to a target array
func (c *Configure) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(28)

	// Offset (0) 'SupportedNetworks'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(c.SupportedNetworks); ii++ {
		offset += 4
		offset += len(c.SupportedNetworks[ii])
	}

	// Offset (1) 'SystemTxSigner'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(c.SystemTxSigner)

	// Offset (2) 'SSVTokenAddressByNetwork'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(c.SSVTokenAddressByNetwork); ii++ {
		offset += 4
		offset += len(c.SSVTokenAddressByNetwork[ii])
	}

	// Field (3) 'MainSSVTokenNetwork'
	dst = append(dst, c.MainSSVTokenNetwork[:]...)

	// Offset (4) 'MainSSVTokenAddress'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(c.MainSSVTokenAddress)

	// Field (5) 'MissedValidationPenalty'
	dst = ssz.MarshalUint64(dst, c.MissedValidationPenalty)

	// Field (0) 'SupportedNetworks'
	if size := len(c.SupportedNetworks); size > 12 {
		err = ssz.ErrListTooBigFn("Configure.SupportedNetworks", size, 12)
		return
	}
	{
		offset = 4 * len(c.SupportedNetworks)
		for ii := 0; ii < len(c.SupportedNetworks); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += len(c.SupportedNetworks[ii])
		}
	}
	for ii := 0; ii < len(c.SupportedNetworks); ii++ {
		if size := len(c.SupportedNetworks[ii]); size > 4 {
			err = ssz.ErrBytesLengthFn("Configure.SupportedNetworks[ii]", size, 4)
			return
		}
		dst = append(dst, c.SupportedNetworks[ii]...)
	}

	// Field (1) 'SystemTxSigner'
	if size := len(c.SystemTxSigner); size > 128 {
		err = ssz.ErrBytesLengthFn("Configure.SystemTxSigner", size, 128)
		return
	}
	dst = append(dst, c.SystemTxSigner...)

	// Field (2) 'SSVTokenAddressByNetwork'
	if size := len(c.SSVTokenAddressByNetwork); size > 12 {
		err = ssz.ErrListTooBigFn("Configure.SSVTokenAddressByNetwork", size, 12)
		return
	}
	{
		offset = 4 * len(c.SSVTokenAddressByNetwork)
		for ii := 0; ii < len(c.SSVTokenAddressByNetwork); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += len(c.SSVTokenAddressByNetwork[ii])
		}
	}
	for ii := 0; ii < len(c.SSVTokenAddressByNetwork); ii++ {
		if size := len(c.SSVTokenAddressByNetwork[ii]); size > 128 {
			err = ssz.ErrBytesLengthFn("Configure.SSVTokenAddressByNetwork[ii]", size, 128)
			return
		}
		dst = append(dst, c.SSVTokenAddressByNetwork[ii]...)
	}

	// Field (4) 'MainSSVTokenAddress'
	if size := len(c.MainSSVTokenAddress); size > 128 {
		err = ssz.ErrBytesLengthFn("Configure.MainSSVTokenAddress", size, 128)
		return
	}
	dst = append(dst, c.MainSSVTokenAddress...)

	return
}

// UnmarshalSSZ ssz unmarshals the Configure object
func (c *Configure) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 28 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1, o2, o4 uint64

	// Offset (0) 'SupportedNetworks'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 28 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'SystemTxSigner'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Offset (2) 'SSVTokenAddressByNetwork'
	if o2 = ssz.ReadOffset(buf[8:12]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Field (3) 'MainSSVTokenNetwork'
	copy(c.MainSSVTokenNetwork[:], buf[12:16])

	// Offset (4) 'MainSSVTokenAddress'
	if o4 = ssz.ReadOffset(buf[16:20]); o4 > size || o2 > o4 {
		return ssz.ErrOffset
	}

	// Field (5) 'MissedValidationPenalty'
	c.MissedValidationPenalty = ssz.UnmarshallUint64(buf[20:28])

	// Field (0) 'SupportedNetworks'
	{
		buf = tail[o0:o1]
		num, err := ssz.DecodeDynamicLength(buf, 12)
		if err != nil {
			return err
		}
		c.SupportedNetworks = make([][]byte, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if len(buf) > 4 {
				return ssz.ErrBytesLength
			}
			if cap(c.SupportedNetworks[indx]) == 0 {
				c.SupportedNetworks[indx] = make([]byte, 0, len(buf))
			}
			c.SupportedNetworks[indx] = append(c.SupportedNetworks[indx], buf...)
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (1) 'SystemTxSigner'
	{
		buf = tail[o1:o2]
		if len(buf) > 128 {
			return ssz.ErrBytesLength
		}
		if cap(c.SystemTxSigner) == 0 {
			c.SystemTxSigner = make([]byte, 0, len(buf))
		}
		c.SystemTxSigner = append(c.SystemTxSigner, buf...)
	}

	// Field (2) 'SSVTokenAddressByNetwork'
	{
		buf = tail[o2:o4]
		num, err := ssz.DecodeDynamicLength(buf, 12)
		if err != nil {
			return err
		}
		c.SSVTokenAddressByNetwork = make([][]byte, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if len(buf) > 128 {
				return ssz.ErrBytesLength
			}
			if cap(c.SSVTokenAddressByNetwork[indx]) == 0 {
				c.SSVTokenAddressByNetwork[indx] = make([]byte, 0, len(buf))
			}
			c.SSVTokenAddressByNetwork[indx] = append(c.SSVTokenAddressByNetwork[indx], buf...)
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (4) 'MainSSVTokenAddress'
	{
		buf = tail[o4:]
		if len(buf) > 128 {
			return ssz.ErrBytesLength
		}
		if cap(c.MainSSVTokenAddress) == 0 {
			c.MainSSVTokenAddress = make([]byte, 0, len(buf))
		}
		c.MainSSVTokenAddress = append(c.MainSSVTokenAddress, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Configure object
func (c *Configure) SizeSSZ() (size int) {
	size = 28

	// Field (0) 'SupportedNetworks'
	for ii := 0; ii < len(c.SupportedNetworks); ii++ {
		size += 4
		size += len(c.SupportedNetworks[ii])
	}

	// Field (1) 'SystemTxSigner'
	size += len(c.SystemTxSigner)

	// Field (2) 'SSVTokenAddressByNetwork'
	for ii := 0; ii < len(c.SSVTokenAddressByNetwork); ii++ {
		size += 4
		size += len(c.SSVTokenAddressByNetwork[ii])
	}

	// Field (4) 'MainSSVTokenAddress'
	size += len(c.MainSSVTokenAddress)

	return
}

// HashTreeRoot ssz hashes the Configure object
func (c *Configure) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the Configure object with a hasher
func (c *Configure) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'SupportedNetworks'
	{
		subIndx := hh.Index()
		num := uint64(len(c.SupportedNetworks))
		if num > 12 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range c.SupportedNetworks {
			{
				elemIndx := hh.Index()
				byteLen := uint64(len(elem))
				if byteLen > 4 {
					err = ssz.ErrIncorrectListSize
					return
				}
				hh.AppendBytes32(elem)
				hh.MerkleizeWithMixin(elemIndx, byteLen, (4+31)/32)
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 12)
	}

	// Field (1) 'SystemTxSigner'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(c.SystemTxSigner))
		if byteLen > 128 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(c.SystemTxSigner)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (128+31)/32)
	}

	// Field (2) 'SSVTokenAddressByNetwork'
	{
		subIndx := hh.Index()
		num := uint64(len(c.SSVTokenAddressByNetwork))
		if num > 12 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range c.SSVTokenAddressByNetwork {
			{
				elemIndx := hh.Index()
				byteLen := uint64(len(elem))
				if byteLen > 128 {
					err = ssz.ErrIncorrectListSize
					return
				}
				hh.AppendBytes32(elem)
				hh.MerkleizeWithMixin(elemIndx, byteLen, (128+31)/32)
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 12)
	}

	// Field (3) 'MainSSVTokenNetwork'
	hh.PutBytes(c.MainSSVTokenNetwork[:])

	// Field (4) 'MainSSVTokenAddress'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(c.MainSSVTokenAddress))
		if byteLen > 128 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(c.MainSSVTokenAddress)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (128+31)/32)
	}

	// Field (5) 'MissedValidationPenalty'
	hh.PutUint64(c.MissedValidationPenalty)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Configure object
func (c *Configure) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(c)
}
