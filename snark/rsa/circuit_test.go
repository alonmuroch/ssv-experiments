package rsa

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/test"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDebug(t *testing.T) {
	c := &Circuit{}

	p, _ := (&fr.Element{}).SetString("64135289477071580278790190170577389084825014742943447208116859632024532344630238623598752668347708737661925585694639798853367")
	q, _ := (&fr.Element{}).SetString("33372027594978156556226010605355114227940760344767554666784520987023841729210037080257448673296881877565718986258036932062711")
	rsa, _ := (&fr.Element{}).SetString("2140324650240744961264423072839333563008614715144755017797754920881418023447140136643345519095804679610992851872470914587687396261921557363047454770520805119056493106687691590019759405693457452230589325976697471681738069364894699871578494975937497937")

	w := &Circuit{
		P:   p,
		Q:   q,
		RSA: rsa,
	}

	// 0. compile
	//var circuit Circuit
	//_, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	//require.NoError(t, err)

	//witness, err := frontend.NewWitness(c, ecc.BN254.ScalarField())
	//require.NoError(t, err)

	require.NoError(t, test.IsSolved(c, w, ecc.BN254.ScalarField()))
}

func TestCircuit_Define(t *testing.T) {
	p, _ := (&fr.Element{}).SetString("64135289477071580278790190170577389084825014742943447208116859632024532344630238623598752668347708737661925585694639798853367")
	q, _ := (&fr.Element{}).SetString("33372027594978156556226010605355114227940760344767554666784520987023841729210037080257448673296881877565718986258036932062711")
	rsa, _ := (&fr.Element{}).SetString("2140324650240744961264423072839333563008614715144755017797754920881418023447140136643345519095804679610992851872470914587687396261921557363047454770520805119056493106687691590019759405693457452230589325976697471681738069364894699871578494975937497937")

	c := &Circuit{
		P:   p,
		Q:   q,
		RSA: rsa,
	}

	// 0. compile
	var myCircuit Circuit
	r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &myCircuit)
	require.NoError(t, err)

	// 1. One time setup
	pk, vk, err := groth16.Setup(r1cs)

	// 2. witness
	witness, err := frontend.NewWitness(c, ecc.BN254.ScalarField())
	require.NoError(t, err)
	witnessPublic, err := frontend.NewWitness(c, ecc.BN254.ScalarField(), frontend.PublicOnly())

	// 3. Proof creation
	proof, err := groth16.Prove(r1cs, pk, witness)

	// 4. Proof verification
	require.NoError(t, groth16.Verify(proof, vk, witnessPublic))
}
