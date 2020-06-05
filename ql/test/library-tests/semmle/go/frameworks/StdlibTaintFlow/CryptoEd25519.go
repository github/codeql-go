package main

import "crypto/ed25519"

func TaintStepTest_CryptoEd25519Sign(sourceCQL interface{}) {
	// The flow is from `message` into `into222`.

	// Assume that `sourceCQL` has the underlying type of `message`:
	message := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `message` to result `into222`
	// (`into222` is now tainted).
	into222 := ed25519.Sign(nil, message)

	// Sink the tainted `into222`:
	sink(into222)
}

func TaintStepTest_CryptoEd25519PrivateKeySign(sourceCQL interface{}) {
	// The flow is from `message` into `signature`.

	// Assume that `sourceCQL` has the underlying type of `message`:
	message := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL ed25519.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `message` to the result `signature`
	// (`signature` is now tainted).
	signature, _ := mediumObjCQL.Sign(nil, message, nil)

	// Sink the tainted `signature`:
	sink(signature)
}

func RunAllTaints_CryptoEd25519(v interface{}) {
	{
		source := newSource()
		TaintStepTest_CryptoEd25519Sign(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoEd25519PrivateKeySign(source)
	}
}
