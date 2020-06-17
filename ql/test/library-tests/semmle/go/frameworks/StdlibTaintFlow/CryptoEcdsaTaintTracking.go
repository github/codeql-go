// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "crypto/ecdsa"

func TaintStepTest_CryptoEcdsaPrivateKeySign_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte163` into `intoByte978`.

	// Assume that `sourceCQL` has the underlying type of `fromByte163`:
	fromByte163 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL ecdsa.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `fromByte163` to the result `intoByte978`
	// (`intoByte978` is now tainted).
	intoByte978, _ := mediumObjCQL.Sign(nil, fromByte163, nil)

	// Return the tainted `intoByte978`:
	return intoByte978
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
