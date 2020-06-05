package main

import "crypto/rsa"

func TaintStepTest_CryptoRsaDecryptOAEP(sourceCQL interface{}) {
	// The flow is from `ciphertext` into `into492`.

	// Assume that `sourceCQL` has the underlying type of `ciphertext`:
	ciphertext := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `ciphertext` to result `into492`
	// (`into492` is now tainted).
	into492, _ := rsa.DecryptOAEP(nil, nil, nil, ciphertext, nil)

	// Sink the tainted `into492`:
	sink(into492)
}

func TaintStepTest_CryptoRsaDecryptPKCS1V15(sourceCQL interface{}) {
	// The flow is from `ciphertext` into `into768`.

	// Assume that `sourceCQL` has the underlying type of `ciphertext`:
	ciphertext := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `ciphertext` to result `into768`
	// (`into768` is now tainted).
	into768, _ := rsa.DecryptPKCS1v15(nil, nil, ciphertext)

	// Sink the tainted `into768`:
	sink(into768)
}

func TaintStepTest_CryptoRsaDecryptPKCS1V15SessionKey(sourceCQL interface{}) {
	// The flow is from `ciphertext` into `key`.

	// Assume that `sourceCQL` has the underlying type of `ciphertext`:
	ciphertext := sourceCQL.([]byte)

	// Declare `key` variable:
	var key []byte

	// Call the function that transfers the taint
	// from the parameter `ciphertext` to parameter `key`;
	// `key` is now tainted.
	rsa.DecryptPKCS1v15SessionKey(nil, nil, ciphertext, key)

	// Sink the tainted `key`:
	sink(key)
}

func TaintStepTest_CryptoRsaEncryptOAEP(sourceCQL interface{}) {
	// The flow is from `msg` into `into739`.

	// Assume that `sourceCQL` has the underlying type of `msg`:
	msg := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `msg` to result `into739`
	// (`into739` is now tainted).
	into739, _ := rsa.EncryptOAEP(nil, nil, nil, msg, nil)

	// Sink the tainted `into739`:
	sink(into739)
}

func TaintStepTest_CryptoRsaEncryptPKCS1V15(sourceCQL interface{}) {
	// The flow is from `msg` into `into234`.

	// Assume that `sourceCQL` has the underlying type of `msg`:
	msg := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `msg` to result `into234`
	// (`into234` is now tainted).
	into234, _ := rsa.EncryptPKCS1v15(nil, nil, msg)

	// Sink the tainted `into234`:
	sink(into234)
}

func TaintStepTest_CryptoRsaSignPKCS1V15(sourceCQL interface{}) {
	// The flow is from `hashed` into `into852`.

	// Assume that `sourceCQL` has the underlying type of `hashed`:
	hashed := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `hashed` to result `into852`
	// (`into852` is now tainted).
	into852, _ := rsa.SignPKCS1v15(nil, nil, 0, hashed)

	// Sink the tainted `into852`:
	sink(into852)
}

func TaintStepTest_CryptoRsaPrivateKeyDecrypt(sourceCQL interface{}) {
	// The flow is from `ciphertext` into `plaintext`.

	// Assume that `sourceCQL` has the underlying type of `ciphertext`:
	ciphertext := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL rsa.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `ciphertext` to the result `plaintext`
	// (`plaintext` is now tainted).
	plaintext, _ := mediumObjCQL.Decrypt(nil, ciphertext, nil)

	// Sink the tainted `plaintext`:
	sink(plaintext)
}

func TaintStepTest_CryptoRsaPrivateKeySign(sourceCQL interface{}) {
	// The flow is from `digest` into `into128`.

	// Assume that `sourceCQL` has the underlying type of `digest`:
	digest := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL rsa.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `digest` to the result `into128`
	// (`into128` is now tainted).
	into128, _ := mediumObjCQL.Sign(nil, digest, nil)

	// Sink the tainted `into128`:
	sink(into128)
}

func RunAllTaints_CryptoRsa(v interface{}) {
	{
		source := newSource()
		TaintStepTest_CryptoRsaDecryptOAEP(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoRsaDecryptPKCS1V15(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoRsaDecryptPKCS1V15SessionKey(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoRsaEncryptOAEP(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoRsaEncryptPKCS1V15(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoRsaSignPKCS1V15(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoRsaPrivateKeyDecrypt(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoRsaPrivateKeySign(source)
	}
}
