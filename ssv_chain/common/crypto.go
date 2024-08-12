package common

const (
	PublicKey           = 0x0
	EncryptedPrivateKey = 0x1
)

const (
	RSA   = 0x0
	ECDSA = 0x1
	BLS   = 0x2
)

type CryptoKey struct {
	// Type represents the cryptography standard for the key: byte[0] cryptography standard type (RSA, BLS, etc), byte[1] type (public key, encrypted private key, etc)
	Type [2]byte `ssz-size:"2"`
	PK   []byte  `ssz-max:"1024"`
}

func (key *CryptoKey) VerifySignature(sig []byte) error {
	panic("implement")
}
