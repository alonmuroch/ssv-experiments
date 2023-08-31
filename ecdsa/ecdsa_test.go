package ecdsa

import (
	"crypto/rand"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"testing"
)

func Test1(t *testing.T) {
	sk0, err := ecies.GenerateKey(rand.Reader, crypto.S256(), nil)
	if err != nil {
		t.Fatal(err)
	}

	pk0 := sk0.PublicKey

	sk1, err := ecies.GenerateKey(rand.Reader, crypto.S256(), nil)
	if err != nil {
		t.Fatal(err)
	}

	// d*G + sk_0*G
	sk0.D.Add(sk0.D, sk1.D)
}
