// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "crypto/ecdsa"

func TaintStepTest_CryptoEcdsaPrivateKeySign_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte656` into `intoByte414`.

	// Assume that `sourceCQL` has the underlying type of `fromByte656`:
	fromByte656 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL ecdsa.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `fromByte656` to the result `intoByte414`
	// (`intoByte414` is now tainted).
	intoByte414, _ := mediumObjCQL.Sign(nil, fromByte656, nil)

	// Return the tainted `intoByte414`:
	return intoByte414
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
