// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "crypto/ed25519"

func TaintStepTest_CryptoEd25519Sign_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte197` into `intoByte939`.

	// Assume that `sourceCQL` has the underlying type of `fromByte197`:
	fromByte197 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte197` to result `intoByte939`
	// (`intoByte939` is now tainted).
	intoByte939 := ed25519.Sign(nil, fromByte197)

	// Return the tainted `intoByte939`:
	return intoByte939
}

func TaintStepTest_CryptoEd25519PrivateKeySign_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte967` into `intoByte570`.

	// Assume that `sourceCQL` has the underlying type of `fromByte967`:
	fromByte967 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL ed25519.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `fromByte967` to the result `intoByte570`
	// (`intoByte570` is now tainted).
	intoByte570, _ := mediumObjCQL.Sign(nil, fromByte967, nil)

	// Return the tainted `intoByte570`:
	return intoByte570
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
