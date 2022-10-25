// Code generated by fastssz. DO NOT EDIT.
// Hash: edcd88c4c085f5409d3fcc233db2a11db4362330bb111f9314d3cf61c4c38211
package qbft

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/ssz_encoding/types"
)

// MarshalSSZ ssz marshals the State object
func (s *State) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the State object to a target array
func (s *State) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(89)

	// Offset (0) 'Share'
	dst = ssz.WriteOffset(dst, offset)
	offset += s.Share.SizeSSZ()

	// Field (1) 'ID'
	dst = append(dst, s.ID[:]...)

	// Field (2) 'Round'
	dst = ssz.MarshalUint64(dst, s.Round)

	// Field (3) 'Height'
	dst = ssz.MarshalUint64(dst, s.Height)

	// Field (4) 'LastPreparedRound'
	dst = ssz.MarshalUint64(dst, s.LastPreparedRound)

	// Offset (5) 'LastPreparedValue'
	dst = ssz.WriteOffset(dst, offset)
	if s.LastPreparedValue == nil {
		s.LastPreparedValue = new(types.ConsensusInput)
	}
	offset += s.LastPreparedValue.SizeSSZ()

	// Offset (6) 'ProposalAcceptedForCurrentRound'
	dst = ssz.WriteOffset(dst, offset)
	if s.ProposalAcceptedForCurrentRound == nil {
		s.ProposalAcceptedForCurrentRound = new(SignedMessage)
	}
	offset += s.ProposalAcceptedForCurrentRound.SizeSSZ()

	// Field (7) 'Decided'
	dst = ssz.MarshalBool(dst, s.Decided)

	// Offset (8) 'DecidedValue'
	dst = ssz.WriteOffset(dst, offset)
	if s.DecidedValue == nil {
		s.DecidedValue = new(types.ConsensusInput)
	}
	offset += s.DecidedValue.SizeSSZ()

	// Offset (9) 'ProposeContainer'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(s.ProposeContainer); ii++ {
		offset += 4
		offset += s.ProposeContainer[ii].SizeSSZ()
	}

	// Offset (10) 'PrepareContainer'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(s.PrepareContainer); ii++ {
		offset += 4
		offset += s.PrepareContainer[ii].SizeSSZ()
	}

	// Offset (11) 'CommitContainer'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(s.CommitContainer); ii++ {
		offset += 4
		offset += s.CommitContainer[ii].SizeSSZ()
	}

	// Offset (12) 'RoundChangeContainer'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(s.RoundChangeContainer); ii++ {
		offset += 4
		offset += s.RoundChangeContainer[ii].SizeSSZ()
	}

	// Field (0) 'Share'
	if dst, err = s.Share.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (5) 'LastPreparedValue'
	if dst, err = s.LastPreparedValue.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (6) 'ProposalAcceptedForCurrentRound'
	if dst, err = s.ProposalAcceptedForCurrentRound.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (8) 'DecidedValue'
	if dst, err = s.DecidedValue.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (9) 'ProposeContainer'
	if size := len(s.ProposeContainer); size > 256 {
		err = ssz.ErrListTooBigFn("State.ProposeContainer", size, 256)
		return
	}
	{
		offset = 4 * len(s.ProposeContainer)
		for ii := 0; ii < len(s.ProposeContainer); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += s.ProposeContainer[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(s.ProposeContainer); ii++ {
		if dst, err = s.ProposeContainer[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (10) 'PrepareContainer'
	if size := len(s.PrepareContainer); size > 256 {
		err = ssz.ErrListTooBigFn("State.PrepareContainer", size, 256)
		return
	}
	{
		offset = 4 * len(s.PrepareContainer)
		for ii := 0; ii < len(s.PrepareContainer); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += s.PrepareContainer[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(s.PrepareContainer); ii++ {
		if dst, err = s.PrepareContainer[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (11) 'CommitContainer'
	if size := len(s.CommitContainer); size > 256 {
		err = ssz.ErrListTooBigFn("State.CommitContainer", size, 256)
		return
	}
	{
		offset = 4 * len(s.CommitContainer)
		for ii := 0; ii < len(s.CommitContainer); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += s.CommitContainer[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(s.CommitContainer); ii++ {
		if dst, err = s.CommitContainer[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (12) 'RoundChangeContainer'
	if size := len(s.RoundChangeContainer); size > 256 {
		err = ssz.ErrListTooBigFn("State.RoundChangeContainer", size, 256)
		return
	}
	{
		offset = 4 * len(s.RoundChangeContainer)
		for ii := 0; ii < len(s.RoundChangeContainer); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += s.RoundChangeContainer[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(s.RoundChangeContainer); ii++ {
		if dst, err = s.RoundChangeContainer[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the State object
func (s *State) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 89 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o5, o6, o8, o9, o10, o11, o12 uint64

	// Offset (0) 'Share'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 89 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'ID'
	copy(s.ID[:], buf[4:36])

	// Field (2) 'Round'
	s.Round = ssz.UnmarshallUint64(buf[36:44])

	// Field (3) 'Height'
	s.Height = ssz.UnmarshallUint64(buf[44:52])

	// Field (4) 'LastPreparedRound'
	s.LastPreparedRound = ssz.UnmarshallUint64(buf[52:60])

	// Offset (5) 'LastPreparedValue'
	if o5 = ssz.ReadOffset(buf[60:64]); o5 > size || o0 > o5 {
		return ssz.ErrOffset
	}

	// Offset (6) 'ProposalAcceptedForCurrentRound'
	if o6 = ssz.ReadOffset(buf[64:68]); o6 > size || o5 > o6 {
		return ssz.ErrOffset
	}

	// Field (7) 'Decided'
	s.Decided = ssz.UnmarshalBool(buf[68:69])

	// Offset (8) 'DecidedValue'
	if o8 = ssz.ReadOffset(buf[69:73]); o8 > size || o6 > o8 {
		return ssz.ErrOffset
	}

	// Offset (9) 'ProposeContainer'
	if o9 = ssz.ReadOffset(buf[73:77]); o9 > size || o8 > o9 {
		return ssz.ErrOffset
	}

	// Offset (10) 'PrepareContainer'
	if o10 = ssz.ReadOffset(buf[77:81]); o10 > size || o9 > o10 {
		return ssz.ErrOffset
	}

	// Offset (11) 'CommitContainer'
	if o11 = ssz.ReadOffset(buf[81:85]); o11 > size || o10 > o11 {
		return ssz.ErrOffset
	}

	// Offset (12) 'RoundChangeContainer'
	if o12 = ssz.ReadOffset(buf[85:89]); o12 > size || o11 > o12 {
		return ssz.ErrOffset
	}

	// Field (0) 'Share'
	{
		buf = tail[o0:o5]
		if err = s.Share.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (5) 'LastPreparedValue'
	{
		buf = tail[o5:o6]
		if s.LastPreparedValue == nil {
			s.LastPreparedValue = new(types.ConsensusInput)
		}
		if err = s.LastPreparedValue.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (6) 'ProposalAcceptedForCurrentRound'
	{
		buf = tail[o6:o8]
		if s.ProposalAcceptedForCurrentRound == nil {
			s.ProposalAcceptedForCurrentRound = new(SignedMessage)
		}
		if err = s.ProposalAcceptedForCurrentRound.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (8) 'DecidedValue'
	{
		buf = tail[o8:o9]
		if s.DecidedValue == nil {
			s.DecidedValue = new(types.ConsensusInput)
		}
		if err = s.DecidedValue.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (9) 'ProposeContainer'
	{
		buf = tail[o9:o10]
		num, err := ssz.DecodeDynamicLength(buf, 256)
		if err != nil {
			return err
		}
		s.ProposeContainer = make([]*SignedMessage, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if s.ProposeContainer[indx] == nil {
				s.ProposeContainer[indx] = new(SignedMessage)
			}
			if err = s.ProposeContainer[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (10) 'PrepareContainer'
	{
		buf = tail[o10:o11]
		num, err := ssz.DecodeDynamicLength(buf, 256)
		if err != nil {
			return err
		}
		s.PrepareContainer = make([]*SignedMessage, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if s.PrepareContainer[indx] == nil {
				s.PrepareContainer[indx] = new(SignedMessage)
			}
			if err = s.PrepareContainer[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (11) 'CommitContainer'
	{
		buf = tail[o11:o12]
		num, err := ssz.DecodeDynamicLength(buf, 256)
		if err != nil {
			return err
		}
		s.CommitContainer = make([]*SignedMessage, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if s.CommitContainer[indx] == nil {
				s.CommitContainer[indx] = new(SignedMessage)
			}
			if err = s.CommitContainer[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (12) 'RoundChangeContainer'
	{
		buf = tail[o12:]
		num, err := ssz.DecodeDynamicLength(buf, 256)
		if err != nil {
			return err
		}
		s.RoundChangeContainer = make([]*SignedMessage, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if s.RoundChangeContainer[indx] == nil {
				s.RoundChangeContainer[indx] = new(SignedMessage)
			}
			if err = s.RoundChangeContainer[indx].UnmarshalSSZ(buf); err != nil {
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
	size = 89

	// Field (0) 'Share'
	size += s.Share.SizeSSZ()

	// Field (5) 'LastPreparedValue'
	if s.LastPreparedValue == nil {
		s.LastPreparedValue = new(types.ConsensusInput)
	}
	size += s.LastPreparedValue.SizeSSZ()

	// Field (6) 'ProposalAcceptedForCurrentRound'
	if s.ProposalAcceptedForCurrentRound == nil {
		s.ProposalAcceptedForCurrentRound = new(SignedMessage)
	}
	size += s.ProposalAcceptedForCurrentRound.SizeSSZ()

	// Field (8) 'DecidedValue'
	if s.DecidedValue == nil {
		s.DecidedValue = new(types.ConsensusInput)
	}
	size += s.DecidedValue.SizeSSZ()

	// Field (9) 'ProposeContainer'
	for ii := 0; ii < len(s.ProposeContainer); ii++ {
		size += 4
		size += s.ProposeContainer[ii].SizeSSZ()
	}

	// Field (10) 'PrepareContainer'
	for ii := 0; ii < len(s.PrepareContainer); ii++ {
		size += 4
		size += s.PrepareContainer[ii].SizeSSZ()
	}

	// Field (11) 'CommitContainer'
	for ii := 0; ii < len(s.CommitContainer); ii++ {
		size += 4
		size += s.CommitContainer[ii].SizeSSZ()
	}

	// Field (12) 'RoundChangeContainer'
	for ii := 0; ii < len(s.RoundChangeContainer); ii++ {
		size += 4
		size += s.RoundChangeContainer[ii].SizeSSZ()
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

	// Field (0) 'Share'
	if err = s.Share.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'ID'
	hh.PutBytes(s.ID[:])

	// Field (2) 'Round'
	hh.PutUint64(s.Round)

	// Field (3) 'Height'
	hh.PutUint64(s.Height)

	// Field (4) 'LastPreparedRound'
	hh.PutUint64(s.LastPreparedRound)

	// Field (5) 'LastPreparedValue'
	if err = s.LastPreparedValue.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (6) 'ProposalAcceptedForCurrentRound'
	if err = s.ProposalAcceptedForCurrentRound.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (7) 'Decided'
	hh.PutBool(s.Decided)

	// Field (8) 'DecidedValue'
	if err = s.DecidedValue.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (9) 'ProposeContainer'
	{
		subIndx := hh.Index()
		num := uint64(len(s.ProposeContainer))
		if num > 256 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.ProposeContainer {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 256)
	}

	// Field (10) 'PrepareContainer'
	{
		subIndx := hh.Index()
		num := uint64(len(s.PrepareContainer))
		if num > 256 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.PrepareContainer {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 256)
	}

	// Field (11) 'CommitContainer'
	{
		subIndx := hh.Index()
		num := uint64(len(s.CommitContainer))
		if num > 256 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.CommitContainer {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 256)
	}

	// Field (12) 'RoundChangeContainer'
	{
		subIndx := hh.Index()
		num := uint64(len(s.RoundChangeContainer))
		if num > 256 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.RoundChangeContainer {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 256)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the State object
func (s *State) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}

// MarshalSSZ ssz marshals the Instance object
func (i *Instance) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(i)
}

// MarshalSSZTo ssz marshals the Instance object to a target array
func (i *Instance) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(8)

	// Offset (0) 'State'
	dst = ssz.WriteOffset(dst, offset)
	offset += i.State.SizeSSZ()

	// Offset (1) 'StartValue'
	dst = ssz.WriteOffset(dst, offset)
	offset += i.StartValue.SizeSSZ()

	// Field (0) 'State'
	if dst, err = i.State.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'StartValue'
	if dst, err = i.StartValue.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Instance object
func (i *Instance) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 8 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1 uint64

	// Offset (0) 'State'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 8 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'StartValue'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (0) 'State'
	{
		buf = tail[o0:o1]
		if err = i.State.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'StartValue'
	{
		buf = tail[o1:]
		if err = i.StartValue.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Instance object
func (i *Instance) SizeSSZ() (size int) {
	size = 8

	// Field (0) 'State'
	size += i.State.SizeSSZ()

	// Field (1) 'StartValue'
	size += i.StartValue.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the Instance object
func (i *Instance) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(i)
}

// HashTreeRootWith ssz hashes the Instance object with a hasher
func (i *Instance) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'State'
	if err = i.State.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'StartValue'
	if err = i.StartValue.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Instance object
func (i *Instance) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(i)
}

// MarshalSSZ ssz marshals the Controller object
func (c *Controller) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the Controller object to a target array
func (c *Controller) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(56)

	// Field (0) 'ID'
	dst = append(dst, c.ID[:]...)

	// Field (1) 'Height'
	dst = ssz.MarshalUint64(dst, c.Height)

	// Offset (2) 'ActiveInstances'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(c.ActiveInstances); ii++ {
		offset += 4
		offset += c.ActiveInstances[ii].SizeSSZ()
	}

	// Offset (3) 'FutureMsgContainer'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(c.FutureMsgContainer) * 8

	// Field (4) 'Domain'
	dst = append(dst, c.Domain[:]...)

	// Offset (5) 'Share'
	dst = ssz.WriteOffset(dst, offset)
	offset += c.Share.SizeSSZ()

	// Field (2) 'ActiveInstances'
	if size := len(c.ActiveInstances); size > 5 {
		err = ssz.ErrListTooBigFn("Controller.ActiveInstances", size, 5)
		return
	}
	{
		offset = 4 * len(c.ActiveInstances)
		for ii := 0; ii < len(c.ActiveInstances); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += c.ActiveInstances[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(c.ActiveInstances); ii++ {
		if dst, err = c.ActiveInstances[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (3) 'FutureMsgContainer'
	if size := len(c.FutureMsgContainer); size > 13 {
		err = ssz.ErrListTooBigFn("Controller.FutureMsgContainer", size, 13)
		return
	}
	for ii := 0; ii < len(c.FutureMsgContainer); ii++ {
		dst = ssz.MarshalUint64(dst, c.FutureMsgContainer[ii])
	}

	// Field (5) 'Share'
	if dst, err = c.Share.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Controller object
func (c *Controller) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 56 {
		return ssz.ErrSize
	}

	tail := buf
	var o2, o3, o5 uint64

	// Field (0) 'ID'
	copy(c.ID[:], buf[0:32])

	// Field (1) 'Height'
	c.Height = ssz.UnmarshallUint64(buf[32:40])

	// Offset (2) 'ActiveInstances'
	if o2 = ssz.ReadOffset(buf[40:44]); o2 > size {
		return ssz.ErrOffset
	}

	if o2 < 56 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (3) 'FutureMsgContainer'
	if o3 = ssz.ReadOffset(buf[44:48]); o3 > size || o2 > o3 {
		return ssz.ErrOffset
	}

	// Field (4) 'Domain'
	copy(c.Domain[:], buf[48:52])

	// Offset (5) 'Share'
	if o5 = ssz.ReadOffset(buf[52:56]); o5 > size || o3 > o5 {
		return ssz.ErrOffset
	}

	// Field (2) 'ActiveInstances'
	{
		buf = tail[o2:o3]
		num, err := ssz.DecodeDynamicLength(buf, 5)
		if err != nil {
			return err
		}
		c.ActiveInstances = make([]*Instance, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if c.ActiveInstances[indx] == nil {
				c.ActiveInstances[indx] = new(Instance)
			}
			if err = c.ActiveInstances[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (3) 'FutureMsgContainer'
	{
		buf = tail[o3:o5]
		num, err := ssz.DivideInt2(len(buf), 8, 13)
		if err != nil {
			return err
		}
		c.FutureMsgContainer = ssz.ExtendUint64(c.FutureMsgContainer, num)
		for ii := 0; ii < num; ii++ {
			c.FutureMsgContainer[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}

	// Field (5) 'Share'
	{
		buf = tail[o5:]
		if err = c.Share.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Controller object
func (c *Controller) SizeSSZ() (size int) {
	size = 56

	// Field (2) 'ActiveInstances'
	for ii := 0; ii < len(c.ActiveInstances); ii++ {
		size += 4
		size += c.ActiveInstances[ii].SizeSSZ()
	}

	// Field (3) 'FutureMsgContainer'
	size += len(c.FutureMsgContainer) * 8

	// Field (5) 'Share'
	size += c.Share.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the Controller object
func (c *Controller) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the Controller object with a hasher
func (c *Controller) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ID'
	hh.PutBytes(c.ID[:])

	// Field (1) 'Height'
	hh.PutUint64(c.Height)

	// Field (2) 'ActiveInstances'
	{
		subIndx := hh.Index()
		num := uint64(len(c.ActiveInstances))
		if num > 5 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range c.ActiveInstances {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 5)
	}

	// Field (3) 'FutureMsgContainer'
	{
		if size := len(c.FutureMsgContainer); size > 13 {
			err = ssz.ErrListTooBigFn("Controller.FutureMsgContainer", size, 13)
			return
		}
		subIndx := hh.Index()
		for _, i := range c.FutureMsgContainer {
			hh.AppendUint64(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(c.FutureMsgContainer))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(13, numItems, 8))
	}

	// Field (4) 'Domain'
	hh.PutBytes(c.Domain[:])

	// Field (5) 'Share'
	if err = c.Share.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Controller object
func (c *Controller) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(c)
}
