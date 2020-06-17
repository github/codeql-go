package main

import "crypto/ed25519"

func TaintStepTest_CryptoEd25519Sign_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte136` into `intoByte588`.

	// Assume that `sourceCQL` has the underlying type of `fromByte136`:
	fromByte136 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte136` to result `intoByte588`
	// (`intoByte588` is now tainted).
	intoByte588 := ed25519.Sign(nil, fromByte136)

	// Return the tainted `intoByte588`:
	return intoByte588
}

func TaintStepTest_CryptoEd25519PrivateKeySign_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte350` into `intoByte459`.

	// Assume that `sourceCQL` has the underlying type of `fromByte350`:
	fromByte350 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL ed25519.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `fromByte350` to the result `intoByte459`
	// (`intoByte459` is now tainted).
	intoByte459, _ := mediumObjCQL.Sign(nil, fromByte350, nil)

	// Return the tainted `intoByte459`:
	return intoByte459
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
