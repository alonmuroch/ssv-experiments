package p2p

import (
	"bytes"
	"encoding/binary"
)

type Identifier [56]byte

func NewIdentifier(slot uint64, validatorPK [48]byte, beaconRole uint64) Identifier {
	slotByts := make([]byte, 4)
	binary.LittleEndian.PutUint32(slotByts, uint32(slot))

	roleByts := make([]byte, 4)
	binary.LittleEndian.PutUint32(roleByts, uint32(beaconRole))

	ret := Identifier{}
	copy(ret[:4], slotByts)
	copy(ret[4:52], validatorPK[:])
	copy(ret[52:], roleByts)
	return ret
}

func (id Identifier) Equal(other Identifier) bool {
	return bytes.Equal(id[:], other[:])
}

type MsgType uint64

const (
	// SSVConsensusMsgType are all QBFT consensus related messages
	SSVConsensusMsgType MsgType = iota
	// SSVPartialSignatureMsgType are all partial signatures msgs over beacon chain specific signatures
	SSVPartialSignatureMsgType
)

type Message struct {
	// Data max size is qbft SignedMessage max ~= 2^22 + 2^20 + 96 + 13 + 2^20 ~= 2^23
	Data    []byte `ssz-max:"8388608"` // 2^23
	MsgType MsgType
}

func (msg *Message) QuickLookIdentifier() Identifier {
	ret := Identifier{}
	copy(ret[:], msg.Data[:len(ret)])
	return ret
}
