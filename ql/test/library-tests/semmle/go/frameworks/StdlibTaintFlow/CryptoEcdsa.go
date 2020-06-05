package main

import "crypto/ecdsa"

func TaintStepTest_CryptoEcdsaPrivateKeySign(sourceCQL interface{}) {
	// The flow is from `digest` into `into486`.

	// Assume that `sourceCQL` has the underlying type of `digest`:
	digest := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL ecdsa.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `digest` to the result `into486`
	// (`into486` is now tainted).
	into486, _ := mediumObjCQL.Sign(nil, digest, nil)

	// Sink the tainted `into486`:
	sink(into486)
}

func RunAllTaints_CryptoEcdsa(v interface{}) {
	{
		source := newSource()
		TaintStepTest_CryptoEcdsaPrivateKeySign(source)
	}
}
