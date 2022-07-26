// Code generated by fastssz. DO NOT EDIT.
// Hash: ae1672a1897051bd8608c254fc2aed1ab9f56557b63e6ab653444a548f98ddae
package qbft

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/ssz_encoding/types"
)

// MarshalSSZ ssz marshals the Message object
func (m *Message) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the Message object to a target array
func (m *Message) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(88)

	// Field (0) 'ID'
	dst = append(dst, m.ID[:]...)

	// Field (1) 'Type'
	dst = ssz.MarshalUint64(dst, uint64(m.Type))

	// Field (2) 'Height'
	dst = ssz.MarshalUint64(dst, m.Height)

	// Field (3) 'Round'
	dst = ssz.MarshalUint64(dst, m.Round)

	// Offset (4) 'Input'
	dst = ssz.WriteOffset(dst, offset)
	offset += m.Input.SizeSSZ()

	// Field (5) 'PreparedRound'
	dst = ssz.MarshalUint64(dst, m.PreparedRound)

	// Field (4) 'Input'
	if dst, err = m.Input.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Message object
func (m *Message) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 88 {
		return ssz.ErrSize
	}

	tail := buf
	var o4 uint64

	// Field (0) 'ID'
	copy(m.ID[:], buf[0:52])

	// Field (1) 'Type'
	m.Type = Type(ssz.UnmarshallUint64(buf[52:60]))

	// Field (2) 'Height'
	m.Height = ssz.UnmarshallUint64(buf[60:68])

	// Field (3) 'Round'
	m.Round = ssz.UnmarshallUint64(buf[68:76])

	// Offset (4) 'Input'
	if o4 = ssz.ReadOffset(buf[76:80]); o4 > size {
		return ssz.ErrOffset
	}

	if o4 < 88 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (5) 'PreparedRound'
	m.PreparedRound = ssz.UnmarshallUint64(buf[80:88])

	// Field (4) 'Input'
	{
		buf = tail[o4:]
		if err = m.Input.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Message object
func (m *Message) SizeSSZ() (size int) {
	size = 88

	// Field (4) 'Input'
	size += m.Input.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the Message object
func (m *Message) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the Message object with a hasher
func (m *Message) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ID'
	hh.PutBytes(m.ID[:])

	// Field (1) 'Type'
	hh.PutUint64(uint64(m.Type))

	// Field (2) 'Height'
	hh.PutUint64(m.Height)

	// Field (3) 'Round'
	hh.PutUint64(m.Round)

	// Field (4) 'Input'
	if err = m.Input.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (5) 'PreparedRound'
	hh.PutUint64(m.PreparedRound)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Message object
func (m *Message) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(m)
}

// MarshalSSZ ssz marshals the SignedMessage object
func (s *SignedMessage) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SignedMessage object to a target array
func (s *SignedMessage) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(112)

	// Offset (0) 'Message'
	dst = ssz.WriteOffset(dst, offset)
	offset += s.Message.SizeSSZ()

	// Offset (1) 'Signers'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(s.Signers) * 8

	// Field (2) 'Signature'
	dst = append(dst, s.Signature[:]...)

	// Offset (3) 'RoundChangeJustifications'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(s.RoundChangeJustifications); ii++ {
		offset += 4
		offset += s.RoundChangeJustifications[ii].SizeSSZ()
	}

	// Offset (4) 'ProposalJustifications'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(s.ProposalJustifications); ii++ {
		offset += 4
		offset += s.ProposalJustifications[ii].SizeSSZ()
	}

	// Field (0) 'Message'
	if dst, err = s.Message.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Signers'
	if size := len(s.Signers); size > 13 {
		err = ssz.ErrListTooBigFn("SignedMessage.Signers", size, 13)
		return
	}
	for ii := 0; ii < len(s.Signers); ii++ {
		dst = ssz.MarshalUint64(dst, s.Signers[ii])
	}

	// Field (3) 'RoundChangeJustifications'
	if size := len(s.RoundChangeJustifications); size > 13 {
		err = ssz.ErrListTooBigFn("SignedMessage.RoundChangeJustifications", size, 13)
		return
	}
	{
		offset = 4 * len(s.RoundChangeJustifications)
		for ii := 0; ii < len(s.RoundChangeJustifications); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += s.RoundChangeJustifications[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(s.RoundChangeJustifications); ii++ {
		if dst, err = s.RoundChangeJustifications[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (4) 'ProposalJustifications'
	if size := len(s.ProposalJustifications); size > 13 {
		err = ssz.ErrListTooBigFn("SignedMessage.ProposalJustifications", size, 13)
		return
	}
	{
		offset = 4 * len(s.ProposalJustifications)
		for ii := 0; ii < len(s.ProposalJustifications); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += s.ProposalJustifications[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(s.ProposalJustifications); ii++ {
		if dst, err = s.ProposalJustifications[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the SignedMessage object
func (s *SignedMessage) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 112 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1, o3, o4 uint64

	// Offset (0) 'Message'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 112 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'Signers'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (2) 'Signature'
	copy(s.Signature[:], buf[8:104])

	// Offset (3) 'RoundChangeJustifications'
	if o3 = ssz.ReadOffset(buf[104:108]); o3 > size || o1 > o3 {
		return ssz.ErrOffset
	}

	// Offset (4) 'ProposalJustifications'
	if o4 = ssz.ReadOffset(buf[108:112]); o4 > size || o3 > o4 {
		return ssz.ErrOffset
	}

	// Field (0) 'Message'
	{
		buf = tail[o0:o1]
		if err = s.Message.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'Signers'
	{
		buf = tail[o1:o3]
		num, err := ssz.DivideInt2(len(buf), 8, 13)
		if err != nil {
			return err
		}
		s.Signers = ssz.ExtendUint64(s.Signers, num)
		for ii := 0; ii < num; ii++ {
			s.Signers[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}

	// Field (3) 'RoundChangeJustifications'
	{
		buf = tail[o3:o4]
		num, err := ssz.DecodeDynamicLength(buf, 13)
		if err != nil {
			return err
		}
		s.RoundChangeJustifications = make([]*SignedMessageHeader, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if s.RoundChangeJustifications[indx] == nil {
				s.RoundChangeJustifications[indx] = new(SignedMessageHeader)
			}
			if err = s.RoundChangeJustifications[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (4) 'ProposalJustifications'
	{
		buf = tail[o4:]
		num, err := ssz.DecodeDynamicLength(buf, 13)
		if err != nil {
			return err
		}
		s.ProposalJustifications = make([]*SignedMessageHeader, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if s.ProposalJustifications[indx] == nil {
				s.ProposalJustifications[indx] = new(SignedMessageHeader)
			}
			if err = s.ProposalJustifications[indx].UnmarshalSSZ(buf); err != nil {
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

// SizeSSZ returns the ssz encoded size in bytes for the SignedMessage object
func (s *SignedMessage) SizeSSZ() (size int) {
	size = 112

	// Field (0) 'Message'
	size += s.Message.SizeSSZ()

	// Field (1) 'Signers'
	size += len(s.Signers) * 8

	// Field (3) 'RoundChangeJustifications'
	for ii := 0; ii < len(s.RoundChangeJustifications); ii++ {
		size += 4
		size += s.RoundChangeJustifications[ii].SizeSSZ()
	}

	// Field (4) 'ProposalJustifications'
	for ii := 0; ii < len(s.ProposalJustifications); ii++ {
		size += 4
		size += s.ProposalJustifications[ii].SizeSSZ()
	}

	return
}

// HashTreeRoot ssz hashes the SignedMessage object
func (s *SignedMessage) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SignedMessage object with a hasher
func (s *SignedMessage) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Message'
	if err = s.Message.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Signers'
	{
		if size := len(s.Signers); size > 13 {
			err = ssz.ErrListTooBigFn("SignedMessage.Signers", size, 13)
			return
		}
		subIndx := hh.Index()
		for _, i := range s.Signers {
			hh.AppendUint64(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(s.Signers))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(13, numItems, 8))
	}

	// Field (2) 'Signature'
	hh.PutBytes(s.Signature[:])

	// Field (3) 'RoundChangeJustifications'
	{
		subIndx := hh.Index()
		num := uint64(len(s.RoundChangeJustifications))
		if num > 13 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.RoundChangeJustifications {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 13)
	}

	// Field (4) 'ProposalJustifications'
	{
		subIndx := hh.Index()
		num := uint64(len(s.ProposalJustifications))
		if num > 13 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.ProposalJustifications {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 13)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SignedMessage object
func (s *SignedMessage) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}

// MarshalSSZ ssz marshals the MessageHeader object
func (m *MessageHeader) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the MessageHeader object to a target array
func (m *MessageHeader) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'ID'
	dst = append(dst, m.ID[:]...)

	// Field (1) 'Type'
	dst = ssz.MarshalUint64(dst, uint64(m.Type))

	// Field (2) 'Height'
	dst = ssz.MarshalUint64(dst, m.Height)

	// Field (3) 'Round'
	dst = ssz.MarshalUint64(dst, m.Round)

	// Field (4) 'InputRoot'
	dst = append(dst, m.InputRoot[:]...)

	// Field (5) 'PreparedRound'
	dst = ssz.MarshalUint64(dst, m.PreparedRound)

	return
}

// UnmarshalSSZ ssz unmarshals the MessageHeader object
func (m *MessageHeader) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 116 {
		return ssz.ErrSize
	}

	// Field (0) 'ID'
	copy(m.ID[:], buf[0:52])

	// Field (1) 'Type'
	m.Type = Type(ssz.UnmarshallUint64(buf[52:60]))

	// Field (2) 'Height'
	m.Height = ssz.UnmarshallUint64(buf[60:68])

	// Field (3) 'Round'
	m.Round = ssz.UnmarshallUint64(buf[68:76])

	// Field (4) 'InputRoot'
	copy(m.InputRoot[:], buf[76:108])

	// Field (5) 'PreparedRound'
	m.PreparedRound = ssz.UnmarshallUint64(buf[108:116])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the MessageHeader object
func (m *MessageHeader) SizeSSZ() (size int) {
	size = 116
	return
}

// HashTreeRoot ssz hashes the MessageHeader object
func (m *MessageHeader) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the MessageHeader object with a hasher
func (m *MessageHeader) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ID'
	hh.PutBytes(m.ID[:])

	// Field (1) 'Type'
	hh.PutUint64(uint64(m.Type))

	// Field (2) 'Height'
	hh.PutUint64(m.Height)

	// Field (3) 'Round'
	hh.PutUint64(m.Round)

	// Field (4) 'InputRoot'
	hh.PutBytes(m.InputRoot[:])

	// Field (5) 'PreparedRound'
	hh.PutUint64(m.PreparedRound)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the MessageHeader object
func (m *MessageHeader) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(m)
}

// MarshalSSZ ssz marshals the SignedMessageHeader object
func (s *SignedMessageHeader) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SignedMessageHeader object to a target array
func (s *SignedMessageHeader) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(216)

	// Field (0) 'Message'
	if dst, err = s.Message.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (1) 'Signers'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(s.Signers) * 8

	// Field (2) 'Signature'
	dst = append(dst, s.Signature[:]...)

	// Field (1) 'Signers'
	if size := len(s.Signers); size > 13 {
		err = ssz.ErrListTooBigFn("SignedMessageHeader.Signers", size, 13)
		return
	}
	for ii := 0; ii < len(s.Signers); ii++ {
		dst = ssz.MarshalUint64(dst, s.Signers[ii])
	}

	return
}

// UnmarshalSSZ ssz unmarshals the SignedMessageHeader object
func (s *SignedMessageHeader) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 216 {
		return ssz.ErrSize
	}

	tail := buf
	var o1 uint64

	// Field (0) 'Message'
	if err = s.Message.UnmarshalSSZ(buf[0:116]); err != nil {
		return err
	}

	// Offset (1) 'Signers'
	if o1 = ssz.ReadOffset(buf[116:120]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 216 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (2) 'Signature'
	copy(s.Signature[:], buf[120:216])

	// Field (1) 'Signers'
	{
		buf = tail[o1:]
		num, err := ssz.DivideInt2(len(buf), 8, 13)
		if err != nil {
			return err
		}
		s.Signers = ssz.ExtendUint64(s.Signers, num)
		for ii := 0; ii < num; ii++ {
			s.Signers[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedMessageHeader object
func (s *SignedMessageHeader) SizeSSZ() (size int) {
	size = 216

	// Field (1) 'Signers'
	size += len(s.Signers) * 8

	return
}

// HashTreeRoot ssz hashes the SignedMessageHeader object
func (s *SignedMessageHeader) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SignedMessageHeader object with a hasher
func (s *SignedMessageHeader) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Message'
	if err = s.Message.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Signers'
	{
		if size := len(s.Signers); size > 13 {
			err = ssz.ErrListTooBigFn("SignedMessageHeader.Signers", size, 13)
			return
		}
		subIndx := hh.Index()
		for _, i := range s.Signers {
			hh.AppendUint64(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(s.Signers))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(13, numItems, 8))
	}

	// Field (2) 'Signature'
	hh.PutBytes(s.Signature[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SignedMessageHeader object
func (s *SignedMessageHeader) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
