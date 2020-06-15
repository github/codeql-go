package main

import "crypto/ecdsa"

func TaintStepTest_CryptoEcdsaPrivateKeySign_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte157` into `intoByte912`.

	// Assume that `sourceCQL` has the underlying type of `fromByte157`:
	fromByte157 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL ecdsa.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `fromByte157` to the result `intoByte912`
	// (`intoByte912` is now tainted).
	intoByte912, _ := mediumObjCQL.Sign(nil, fromByte157, nil)

	// Sink the tainted `intoByte912`:
	sink(intoByte912)
}

func RunAllTaints_CryptoEcdsa() {
	{
		source := newSource()
		TaintStepTest_CryptoEcdsaPrivateKeySign_B0I0O0(source)
	}
}
