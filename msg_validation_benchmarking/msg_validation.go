package main

import (
	"github.com/herumi/bls-eth-go-binary/bls"
	"ssv-experiments/msg_validation_benchmarking/p2p"
	"ssv-experiments/msg_validation_benchmarking/qbft"
)

func MessageValidation(pk *bls.PublicKey) func(msgByts []byte) {
	return func(msgByts []byte) {
		p2pMsg := &p2p.Message{}
		p2pMsg.UnmarshalSSZ(msgByts)

		qbftMsg := &qbft.SignedMessage{}
		qbftMsg.UnmarshalSSZ(p2pMsg.Data)

		r, _ := qbftMsg.Message.HashTreeRoot()

		// solves some weird cgo issue
		sigByts := make([]byte, 96)
		copy(sigByts, qbftMsg.Signature[:])

		sig := bls.Sign{}
		if err := sig.Deserialize(sigByts); err != nil {
			panic(err.Error())
		}
		if !sig.VerifyByte(pk, r[:]) {
			panic("sig invalid")
		}
	}
}
