package _2381

import (
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/stretchr/testify/require"
	"testing"
)

func sharedSecret(t, n int) (*bls.SecretKey, map[int]*bls.SecretKey) {
	var sk *bls.SecretKey
	coefficients := make([]bls.Fr, t+1)
	for i := 0; i < t; i++ {
		coefficients[i] = bls.Fr{}
		coefficients[i].SetByCSPRNG()

		if i == 0 {
			sk = bls.CastToSecretKey(&coefficients[i])
		}
	}

	ret := make(map[int]*bls.SecretKey, n)
	for i := 1; i <= n; i++ {
		x := &bls.Fr{}
		x.SetInt64(int64(i))

		out := &bls.Fr{}

		if err := bls.FrEvaluatePolynomial(out, coefficients, x); err != nil {
			panic(err.Error())
		}

		ret[i-1] = bls.CastToSecretKey(out)
	}

	return sk, ret
}

func TestThresholdEncryptionBLS12381(t *testing.T) {
	_ = bls.Init(bls.BLS12_381)
	_ = bls.SetETHmode(bls.EthModeDraft07)

	plainText := []byte("hello world first try i need 32 bytes at least here!")

	sk, sks := sharedSecret(3, 4)
	data, err := GetCipherText(plainText, sk.GetPublicKey())
	require.NoError(t, err)

	decryptionShares := make([]*bls.G1, len(sks))
	for i, ski := range sks {
		dsi, err := GetDecryptionShare(data, ski)
		require.NoError(t, err)

		require.NoError(t, Verify(data, dsi, ski.GetPublicKey()))

		decryptionShares[i] = dsi
	}

	decrypted, err := CombineShares(data, decryptionShares, 3)
	require.NoError(t, err)

	require.EqualValues(t, plainText, decrypted)
}
