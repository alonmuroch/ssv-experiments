package rsa_signature

import (
	"github.com/consensys/gnark/frontend"
	"math/big"
)

// Circuit defines a simple circuit
// RSA verification
type Circuit struct {
	S  frontend.Variable `gnark:",private"` // Signature
	E  frontend.Variable `gnark:",public"`  // Public exponent
	N  frontend.Variable `gnark:",public"`  // Modulus
	Hm frontend.Variable `gnark:",public"`  // Hashed message
}

// Define declares the circuit constraints
// prove x < y
func (circuit *Circuit) Define(api frontend.API) error {
	api.Println("S: ", circuit.S)
	api.Println("E: ", circuit.E)
	api.Println("N: ", circuit.N)
	api.Println("Hm: ", circuit.Hm)
	// s^e mod n
	modResult, err := Exp(api, circuit.S, circuit.E)
	if err != nil {
		return err
	}
	api.Println("result: ", modResult)

	// s^e mod n
	//modResult := api.Mod(expResult, circuit.N)
	//
	// Check if s^e mod n == H(m)
	//api.AssertIsEqual(modResult, circuit.Hm)

	return nil
}

func Exp(api frontend.API, base, exp frontend.Variable) (frontend.Variable, error) {
	var result frontend.Variable
	result = (&big.Int{}).SetInt64(1)

	// Exponentiation by repeated multiplication
	Loop(api, exp, func() error {
		result = api.Mul(result, base)
		return nil
	})

	return result, nil
}

//func Exp(api frontend.API, base, exp frontend.Variable) (frontend.Variable, error) {
//	var result frontend.Variable
//	result = base //(&fr.Element{}).SetOne()
//
//	/**
//	1) convert exponent to binary
//	2) iterate bits such that for each set we multiply result by the variable mul
//	3) binary counting: result = 2^i
//	4) mul starts set to 1 and gets multiplied by base with each iteration
//	5) example: exp = 7 (binary: [1,1,1, .. <all zeros> .. ])
//
//		5.1) mul 	-> 1 	-> 1*base -> (1*base)*base 	-> ((1*base)*base)*base
//		5.1) result -> base -> base^2 -> base^4 		-> base^7
//	*/
//
//	bits := bits.ToBinary(api, 2, bits.WithNbDigits(256))
//	var mul frontend.Variable
//	mul = base // (&fr.Element{}).SetOne()
//
//	//api.Println("bits: ", bits)
//
//	// Exponentiation by repeated multiplication
//	for i := 0; i < len(bits); i++ { // assuming Exponent is a 256-bit integer
//		api.Println("result "+fmt.Sprintf("%d", i)+" interm: ", result)
//		// If the i-th bit of Exponent is set, multiply result by base
//
//		result = api.Mul(result,
//			api.Select(bits[i], mul, 1),
//		)
//
//		mul = api.Mul(mul, base)
//
//	}
//
//	return result, nil
//}

func Loop(api frontend.API, stop frontend.Variable, f func() error) error {
	i := (&big.Int{}).SetInt64(1)
	for {
		// stop if i > stop
		if api.Cmp(i, stop).(*big.Int).Cmp((&big.Int{}).SetInt64(1)) == 0 {
			break
		}

		if err := f(); err != nil {
			return err
		}

		i = api.Add(i, 1).(*big.Int)
	}
	return nil
}
