package main

import "crypto"

func TaintStepTest_CryptoHashHashFunc(sourceCQL interface{}) {
	// The flow is from `from579` into `into370`.

	// Assume that `sourceCQL` has the underlying type of `from579`:
	from579 := sourceCQL.(crypto.Hash)

	// Call the method that transfers the taint
	// from the receiver `from579` to the result `into370`
	// (`into370` is now tainted).
	into370 := from579.HashFunc()

	// Sink the tainted `into370`:
	sink(into370)
}

func TaintStepTest_CryptoHashNew(sourceCQL interface{}) {
	// The flow is from `from228` into `into283`.

	// Assume that `sourceCQL` has the underlying type of `from228`:
	from228 := sourceCQL.(crypto.Hash)

	// Call the method that transfers the taint
	// from the receiver `from228` to the result `into283`
	// (`into283` is now tainted).
	into283 := from228.New()

	// Sink the tainted `into283`:
	sink(into283)
}

func TaintStepTest_CryptoDecrypterDecrypt(sourceCQL interface{}) {
	// The flow is from `msg` into `plaintext`.

	// Assume that `sourceCQL` has the underlying type of `msg`:
	msg := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL crypto.Decrypter

	// Call the method that transfers the taint
	// from the parameter `msg` to the result `plaintext`
	// (`plaintext` is now tainted).
	plaintext, _ := mediumObjCQL.Decrypt(nil, msg, nil)

	// Sink the tainted `plaintext`:
	sink(plaintext)
}

func TaintStepTest_CryptoSignerPublic(sourceCQL interface{}) {
	// The flow is from `from389` into `into581`.

	// Assume that `sourceCQL` has the underlying type of `from389`:
	from389 := sourceCQL.(crypto.Signer)

	// Call the method that transfers the taint
	// from the receiver `from389` to the result `into581`
	// (`into581` is now tainted).
	into581 := from389.Public()

	// Sink the tainted `into581`:
	sink(into581)
}

func TaintStepTest_CryptoSignerSign(sourceCQL interface{}) {
	// The flow is from `digest` into `signature`.

	// Assume that `sourceCQL` has the underlying type of `digest`:
	digest := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL crypto.Signer

	// Call the method that transfers the taint
	// from the parameter `digest` to the result `signature`
	// (`signature` is now tainted).
	signature, _ := mediumObjCQL.Sign(nil, digest, nil)

	// Sink the tainted `signature`:
	sink(signature)
}

func RunAllTaints_Crypto(v interface{}) {
	{
		source := newSource()
		TaintStepTest_CryptoHashHashFunc(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoHashNew(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoDecrypterDecrypt(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoSignerPublic(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoSignerSign(source)
	}
}
