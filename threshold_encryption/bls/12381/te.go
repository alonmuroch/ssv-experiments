package _2381

import (
	"crypto/sha256"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/pkg/errors"
	"math"
)

type CipherData struct {
	// V cipher text
	V []byte
	// U ephemeral public key
	U *bls.PublicKey
	// W ephemeral key * HashToGroup(U,V)
	W *bls.G2
}

func hash(y *bls.G1) [32]byte {
	return sha256.Sum256(y.Serialize())
}

func hashToGroup(v []byte, u *bls.PublicKey) (*bls.G2, error) {
	ret := &bls.G2{}
	if err := ret.HashAndMapTo(append(v, u.Serialize()...)); err != nil {
		return nil, err
	}
	return ret, nil
}

// padKeyAndData returns padded key and data with equal lengths
func padKeyAndData(key, data []byte) (lhsToHash []byte, rhsToHash []byte, xored []byte) {
	size := int(math.Max(float64(len(data)), float64(len(key))))
	lhsToHash = make([]byte, size)
	rhsToHash = make([]byte, size)
	xored = make([]byte, size)
	for i := 0; i < size; i++ {
		if i < len(key) {
			lhsToHash[i] = key[i]
		} else {
			lhsToHash[i] = 0
		}

		if i < len(data) {
			rhsToHash[i] = data[i]
		} else {
			rhsToHash[i] = 0
		}

		xored[i] = lhsToHash[i] ^ rhsToHash[i]
	}
	return
}

func GetCipherText(msg []byte, pk *bls.PublicKey) (*CipherData, error) {
	if len(msg) < 32 {
		return nil, errors.New("plain text must be >= 32 bytes")
	}

	r := &bls.SecretKey{}
	r.SetByCSPRNG()

	// U
	U := r.GetPublicKey()

	// Y
	Y := &bls.G1{}
	bls.G1Mul(Y, bls.CastFromPublicKey(pk), bls.CastFromSecretKey(r))

	hash := hash(Y)

	_, _, V := padKeyAndData(hash[:], msg)

	H, err := hashToGroup(V, U)
	if err != nil {
		return nil, err
	}
	W := &bls.G2{}
	bls.G2Mul(W, H, bls.CastFromSecretKey(r))

	return &CipherData{
		U: U,
		V: V,
		W: W,
	}, nil
}

func GetDecryptionShare(cipherData *CipherData, ski *bls.SecretKey) (*bls.G1, error) {
	H, err := hashToGroup(cipherData.V, cipherData.U)
	if err != nil {
		return nil, err
	}

	if !bls.VerifyPairing(bls.CastToSign(cipherData.W), bls.CastToSign(H), cipherData.U) {
		return nil, errors.New("cannot decrypt data")
	}

	ret := &bls.G1{}
	bls.G1Mul(ret, bls.CastFromPublicKey(cipherData.U), bls.CastFromSecretKey(ski))
	return ret, nil
}

func Verify(cipherData *CipherData, decryptionShare *bls.G1, pki *bls.PublicKey) error {
	H, err := hashToGroup(cipherData.V, cipherData.U)
	if err != nil {
		return err
	}

	if !bls.VerifyPairing(bls.CastToSign(cipherData.W), bls.CastToSign(H), cipherData.U) {
		return errors.New("cannot verify decryption data")
	}

	// pp1 = libff::alt_bn128_ate_reduced_pairing( W, public_key );
	// pp2 = libff::alt_bn128_ate_reduced_pairing( H, decryptionShare );
	pp1 := &bls.GT{}
	pp2 := &bls.GT{}

	bls.Pairing(pp1, bls.CastFromPublicKey(pki), cipherData.W)
	bls.Pairing(pp2, decryptionShare, H)

	if !pp1.IsEqual(pp2) {
		return errors.New("cannot verify decryption share")
	}

	return nil
}

type Share struct {
	DecryptionShare *bls.G1
	ID              int
}

func combineSharesIntoAESKey(shares []*Share, t int) ([32]byte, error) {
	idx := make([]bls.Fr, t)
	points := make([]bls.G1, t)
	for i := 0; i < t; i++ {
		idx[i] = bls.Fr{}
		idx[i].SetInt64(int64(shares[i].ID))

		points[i] = *shares[i].DecryptionShare
	}

	zeroPosition := &bls.G1{}
	if err := bls.G1LagrangeInterpolation(zeroPosition, idx, points); err != nil {
		return [32]byte{}, err
	}

	return hash(zeroPosition), nil
}

func CombineShares(cipherData *CipherData, decryptionShare []*bls.G1, t int) ([]byte, error) {
	H, err := hashToGroup(cipherData.V, cipherData.U)
	if err != nil {
		return nil, err
	}

	if !bls.VerifyPairing(bls.CastToSign(cipherData.W), bls.CastToSign(H), cipherData.U) {
		return nil, errors.New("cannot verify decryption data")
	}

	shares := make([]*Share, len(decryptionShare))
	for i, ds := range decryptionShare {
		shares[i] = &Share{
			DecryptionShare: ds,
			ID:              i + 1,
		}
	}
	aes, err := combineSharesIntoAESKey(shares, t)
	if err != nil {
		return nil, err
	}

	_, _, xored := padKeyAndData(aes[:], cipherData.V)

	return xored, nil
}
