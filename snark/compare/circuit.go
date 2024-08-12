package compare

import (
	"github.com/consensys/gnark/frontend"
)

// Circuit defines a simple circuit
// X<=Y
type Circuit struct {
	// struct tags on a variable is optional
	// default uses variable name and secret visibility.
	X frontend.Variable `gnark:"x,public"`
	Y frontend.Variable `gnark:"y"`
}

// Define declares the circuit constraints
// prove x < y
func (circuit *Circuit) Define(api frontend.API) error {
	api.Println(api.Cmp(circuit.X, circuit.Y))
	api.AssertIsEqual(api.Cmp(circuit.X, circuit.Y), -1)
	return nil
}
