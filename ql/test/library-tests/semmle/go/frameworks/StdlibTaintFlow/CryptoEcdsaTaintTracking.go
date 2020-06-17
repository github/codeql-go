package main

import "crypto/ecdsa"

func TaintStepTest_CryptoEcdsaPrivateKeySign_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte908` into `intoByte201`.

	// Assume that `sourceCQL` has the underlying type of `fromByte908`:
	fromByte908 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL ecdsa.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `fromByte908` to the result `intoByte201`
	// (`intoByte201` is now tainted).
	intoByte201, _ := mediumObjCQL.Sign(nil, fromByte908, nil)

	// Return the tainted `intoByte201`:
	return intoByte201
}

func RunAllTaints_CryptoEcdsa() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoEcdsaPrivateKeySign_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
