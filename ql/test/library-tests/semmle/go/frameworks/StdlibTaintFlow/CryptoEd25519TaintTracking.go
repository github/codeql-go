package main

import "crypto/ed25519"

func TaintStepTest_CryptoEd25519Sign_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte906` into `intoByte735`.

	// Assume that `sourceCQL` has the underlying type of `fromByte906`:
	fromByte906 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte906` to result `intoByte735`
	// (`intoByte735` is now tainted).
	intoByte735 := ed25519.Sign(nil, fromByte906)

	// Sink the tainted `intoByte735`:
	sink(intoByte735)
}

func TaintStepTest_CryptoEd25519PrivateKeySign_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte312` into `intoByte301`.

	// Assume that `sourceCQL` has the underlying type of `fromByte312`:
	fromByte312 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL ed25519.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `fromByte312` to the result `intoByte301`
	// (`intoByte301` is now tainted).
	intoByte301, _ := mediumObjCQL.Sign(nil, fromByte312, nil)

	// Sink the tainted `intoByte301`:
	sink(intoByte301)
}

func RunAllTaints_CryptoEd25519() {
	{
		source := newSource()
		TaintStepTest_CryptoEd25519Sign_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoEd25519PrivateKeySign_B0I0O0(source)
	}
}
