package utils

import (
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
	"math/big"
)

var (
	curveOrder = new(big.Int)
)

// Init initializes BLS
func Init() {
	_ = bls.Init(bls.BLS12_381)
	_ = bls.SetETHmode(bls.EthModeDraft07)

	curveOrder, _ = curveOrder.SetString(bls.GetCurveOrder(), 10)
}

func GenerateShares(t, n uint64) (map[uint64]*bls.SecretKey, *bls.PublicKey) {
	Init()

	// master key Polynomial
	msk := make([]bls.SecretKey, t)

	sk := &bls.SecretKey{}
	sk.SetByCSPRNG()
	msk[0] = *sk

	// construct poly
	for i := uint64(1); i < t; i++ {
		sk := bls.SecretKey{}
		sk.SetByCSPRNG()
		msk[i] = sk
	}

	// evaluate shares - starting from 1 because 0 is master key
	shares := make(map[uint64]*bls.SecretKey)
	for i := uint64(1); i <= n; i++ {
		id := i
		blsID := bls.ID{}
		err := blsID.SetDecString(fmt.Sprintf("%d", id))
		if err != nil {
			panic(err)
		}

		sk := bls.SecretKey{}

		err = sk.Set(msk, &blsID)
		if err != nil {
			panic(err)
		}

		shares[id] = &sk
	}

	return shares, sk.GetPublicKey()
}
