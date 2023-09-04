package types

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
	"github.com/herumi/bls-eth-go-binary/bls"
)

type Signer struct {
	Signer uint64
	PubKey []byte `ssz-size:"48"`
}

func VerifyObjectSignature(
	signature [96]byte,
	rootObj ssz.HashRoot,
	domain Domain,
	signatureType SignatureType,
	signers []*Signer) error {
	r, err := rootObj.HashTreeRoot()
	if err != nil {
		return err
	}
	return VerifySignature(signature, r, domain, signatureType, signers)
}

func VerifySignature(
	signature [96]byte,
	root [32]byte,
	domain Domain,
	signatureType SignatureType,
	signers []*Signer) error {
	panic("implement")
}

// SignToBLSSignature converts bls.Sign to spec.BLSSignature
func SignToBLSSignature(in *bls.Sign) phase0.BLSSignature {
	ret := phase0.BLSSignature{}
	copy(ret[:], in.Serialize())
	return ret
}
