// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "crypto/ed25519"

func TaintStepTest_CryptoEd25519Sign_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte656` into `intoByte414`.

	// Assume that `sourceCQL` has the underlying type of `fromByte656`:
	fromByte656 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte656` to result `intoByte414`
	// (`intoByte414` is now tainted).
	intoByte414 := ed25519.Sign(nil, fromByte656)

	// Return the tainted `intoByte414`:
	return intoByte414
}

func TaintStepTest_CryptoEd25519PrivateKeySign_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte518` into `intoByte650`.

	// Assume that `sourceCQL` has the underlying type of `fromByte518`:
	fromByte518 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL ed25519.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `fromByte518` to the result `intoByte650`
	// (`intoByte650` is now tainted).
	intoByte650, _ := mediumObjCQL.Sign(nil, fromByte518, nil)

	// Return the tainted `intoByte650`:
	return intoByte650
}

func RunAllTaints_CryptoEd25519() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoEd25519Sign_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoEd25519PrivateKeySign_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
