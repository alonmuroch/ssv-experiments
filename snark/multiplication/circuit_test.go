package multiplication

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCircuit_Define(t *testing.T) {
	c := &Circuit{
		X: 3,
		Y: 35,
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
