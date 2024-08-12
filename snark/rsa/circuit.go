package rsa

import (
	"github.com/consensys/gnark/frontend"
)

// Circuit defines a simple circuit
// RSA verification
type Circuit struct {
	P   frontend.Variable // p  --> secret visibility (default)
	Q   frontend.Variable `gnark:"q,secret"` // q  --> secret visibility
	RSA frontend.Variable `gnark:",public"`  // rsa  --> public visibility
}

// Define declares the circuit constraints
// prove x < y
func (circuit *Circuit) Define(api frontend.API) error {
	api.Println("P: ", circuit.P)
	api.Println("Q: ", circuit.Q)

	// ensure we don't accept RSA * 1 == RSA
	api.AssertIsDifferent(circuit.P, 1)
	api.AssertIsDifferent(circuit.Q, 1)

	// compute P * Q and store it in the local variable res.
	rsa := api.Mul(circuit.P, circuit.Q)
	api.Println("RSA: ", rsa)

	// assert that the statement P * Q == RSA is true.
	api.AssertIsEqual(circuit.RSA, rsa)
	return nil
}
