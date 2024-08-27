// Code generated by fastssz. DO NOT EDIT.
// Hash: 23202b62f33ea0404900792ccbd6737948f9db47174ba369d9c4885761d0276f
// Version: 0.1.2
package types

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/ssv_chain/common"
)

// MarshalSSZ ssz marshals the PriceTier object
func (p *PriceTier) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(p)
}

// MarshalSSZTo ssz marshals the PriceTier object to a target array
func (p *PriceTier) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(22)

	// Field (0) 'Network'
	dst = append(dst, p.Network[:]...)

	// Field (1) 'Capacity'
	dst = ssz.MarshalUint16(dst, p.Capacity)

	// Field (2) 'PricePerToken'
	dst = ssz.MarshalUint64(dst, p.PricePerToken)

	// Offset (3) 'PayableTokenAddress'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(p.PayableTokenAddress)

	// Offset (4) 'WhitelistedAddress'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(p.WhitelistedAddress); ii++ {
		offset += 4
		offset += len(p.WhitelistedAddress[ii])
	}

	// Field (3) 'PayableTokenAddress'
	if size := len(p.PayableTokenAddress); size > 64 {
		err = ssz.ErrBytesLengthFn("PriceTier.PayableTokenAddress", size, 64)
		return
	}
	dst = append(dst, p.PayableTokenAddress...)

	// Field (4) 'WhitelistedAddress'
	if size := len(p.WhitelistedAddress); size > 64 {
		err = ssz.ErrListTooBigFn("PriceTier.WhitelistedAddress", size, 64)
		return
	}
	{
		offset = 4 * len(p.WhitelistedAddress)
		for ii := 0; ii < len(p.WhitelistedAddress); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += len(p.WhitelistedAddress[ii])
		}
	}
	for ii := 0; ii < len(p.WhitelistedAddress); ii++ {
		if size := len(p.WhitelistedAddress[ii]); size > 128 {
			err = ssz.ErrBytesLengthFn("PriceTier.WhitelistedAddress[ii]", size, 128)
			return
		}
		dst = append(dst, p.WhitelistedAddress[ii]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the PriceTier object
func (p *PriceTier) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 22 {
		return ssz.ErrSize
	}

	tail := buf
	var o3, o4 uint64

	// Field (0) 'Network'
	copy(p.Network[:], buf[0:4])

	// Field (1) 'Capacity'
	p.Capacity = ssz.UnmarshallUint16(buf[4:6])

	// Field (2) 'PricePerToken'
	p.PricePerToken = ssz.UnmarshallUint64(buf[6:14])

	// Offset (3) 'PayableTokenAddress'
	if o3 = ssz.ReadOffset(buf[14:18]); o3 > size {
		return ssz.ErrOffset
	}

	if o3 < 22 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (4) 'WhitelistedAddress'
	if o4 = ssz.ReadOffset(buf[18:22]); o4 > size || o3 > o4 {
		return ssz.ErrOffset
	}

	// Field (3) 'PayableTokenAddress'
	{
		buf = tail[o3:o4]
		if len(buf) > 64 {
			return ssz.ErrBytesLength
		}
		if cap(p.PayableTokenAddress) == 0 {
			p.PayableTokenAddress = make([]byte, 0, len(buf))
		}
		p.PayableTokenAddress = append(p.PayableTokenAddress, buf...)
	}

	// Field (4) 'WhitelistedAddress'
	{
		buf = tail[o4:]
		num, err := ssz.DecodeDynamicLength(buf, 64)
		if err != nil {
			return err
		}
		p.WhitelistedAddress = make([][]byte, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if len(buf) > 128 {
				return ssz.ErrBytesLength
			}
			if cap(p.WhitelistedAddress[indx]) == 0 {
				p.WhitelistedAddress[indx] = make([]byte, 0, len(buf))
			}
			p.WhitelistedAddress[indx] = append(p.WhitelistedAddress[indx], buf...)
			return nil
		})
		if err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the PriceTier object
func (p *PriceTier) SizeSSZ() (size int) {
	size = 22

	// Field (3) 'PayableTokenAddress'
	size += len(p.PayableTokenAddress)

	// Field (4) 'WhitelistedAddress'
	for ii := 0; ii < len(p.WhitelistedAddress); ii++ {
		size += 4
		size += len(p.WhitelistedAddress[ii])
	}

	return
}

// HashTreeRoot ssz hashes the PriceTier object
func (p *PriceTier) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(p)
}

// HashTreeRootWith ssz hashes the PriceTier object with a hasher
func (p *PriceTier) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Network'
	hh.PutBytes(p.Network[:])

	// Field (1) 'Capacity'
	hh.PutUint16(p.Capacity)

	// Field (2) 'PricePerToken'
	hh.PutUint64(p.PricePerToken)

	// Field (3) 'PayableTokenAddress'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(p.PayableTokenAddress))
		if byteLen > 64 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(p.PayableTokenAddress)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (64+31)/32)
	}

	// Field (4) 'WhitelistedAddress'
	{
		subIndx := hh.Index()
		num := uint64(len(p.WhitelistedAddress))
		if num > 64 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range p.WhitelistedAddress {
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
		hh.MerkleizeWithMixin(subIndx, num, 64)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the PriceTier object
func (p *PriceTier) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(p)
}

// MarshalSSZ ssz marshals the Operator object
func (o *Operator) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(o)
}

// MarshalSSZTo ssz marshals the Operator object to a target array
func (o *Operator) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(28)

	// Offset (0) 'Address'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(o.Address)

	// Field (1) 'ID'
	dst = ssz.MarshalUint64(dst, o.ID)

	// Offset (2) 'PublicKey'
	dst = ssz.WriteOffset(dst, offset)
	if o.PublicKey == nil {
		o.PublicKey = new(common.CryptoKey)
	}
	offset += o.PublicKey.SizeSSZ()

	// Field (3) 'Module'
	dst = ssz.MarshalUint64(dst, o.Module)

	// Offset (4) 'Tiers'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(o.Tiers); ii++ {
		offset += 4
		offset += o.Tiers[ii].SizeSSZ()
	}

	// Field (0) 'Address'
	if size := len(o.Address); size > 128 {
		err = ssz.ErrBytesLengthFn("Operator.Address", size, 128)
		return
	}
	dst = append(dst, o.Address...)

	// Field (2) 'PublicKey'
	if dst, err = o.PublicKey.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (4) 'Tiers'
	if size := len(o.Tiers); size > 16 {
		err = ssz.ErrListTooBigFn("Operator.Tiers", size, 16)
		return
	}
	{
		offset = 4 * len(o.Tiers)
		for ii := 0; ii < len(o.Tiers); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += o.Tiers[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(o.Tiers); ii++ {
		if dst, err = o.Tiers[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Operator object
func (o *Operator) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 28 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o2, o4 uint64

	// Offset (0) 'Address'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 28 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'ID'
	o.ID = ssz.UnmarshallUint64(buf[4:12])

	// Offset (2) 'PublicKey'
	if o2 = ssz.ReadOffset(buf[12:16]); o2 > size || o0 > o2 {
		return ssz.ErrOffset
	}

	// Field (3) 'Module'
	o.Module = ssz.UnmarshallUint64(buf[16:24])

	// Offset (4) 'Tiers'
	if o4 = ssz.ReadOffset(buf[24:28]); o4 > size || o2 > o4 {
		return ssz.ErrOffset
	}

	// Field (0) 'Address'
	{
		buf = tail[o0:o2]
		if len(buf) > 128 {
			return ssz.ErrBytesLength
		}
		if cap(o.Address) == 0 {
			o.Address = make([]byte, 0, len(buf))
		}
		o.Address = append(o.Address, buf...)
	}

	// Field (2) 'PublicKey'
	{
		buf = tail[o2:o4]
		if o.PublicKey == nil {
			o.PublicKey = new(common.CryptoKey)
		}
		if err = o.PublicKey.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (4) 'Tiers'
	{
		buf = tail[o4:]
		num, err := ssz.DecodeDynamicLength(buf, 16)
		if err != nil {
			return err
		}
		o.Tiers = make([]*PriceTier, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if o.Tiers[indx] == nil {
				o.Tiers[indx] = new(PriceTier)
			}
			if err = o.Tiers[indx].UnmarshalSSZ(buf); err != nil {
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

// SizeSSZ returns the ssz encoded size in bytes for the Operator object
func (o *Operator) SizeSSZ() (size int) {
	size = 28

	// Field (0) 'Address'
	size += len(o.Address)

	// Field (2) 'PublicKey'
	if o.PublicKey == nil {
		o.PublicKey = new(common.CryptoKey)
	}
	size += o.PublicKey.SizeSSZ()

	// Field (4) 'Tiers'
	for ii := 0; ii < len(o.Tiers); ii++ {
		size += 4
		size += o.Tiers[ii].SizeSSZ()
	}

	return
}

// HashTreeRoot ssz hashes the Operator object
func (o *Operator) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(o)
}

// HashTreeRootWith ssz hashes the Operator object with a hasher
func (o *Operator) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Address'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(o.Address))
		if byteLen > 128 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(o.Address)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (128+31)/32)
	}

	// Field (1) 'ID'
	hh.PutUint64(o.ID)

	// Field (2) 'PublicKey'
	if err = o.PublicKey.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (3) 'Module'
	hh.PutUint64(o.Module)

	// Field (4) 'Tiers'
	{
		subIndx := hh.Index()
		num := uint64(len(o.Tiers))
		if num > 16 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range o.Tiers {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 16)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Operator object
func (o *Operator) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(o)
}
