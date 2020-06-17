// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "crypto/rsa"

func TaintStepTest_CryptoRsaDecryptOAEP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte656` into `intoByte414`.

	// Assume that `sourceCQL` has the underlying type of `fromByte656`:
	fromByte656 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte656` to result `intoByte414`
	// (`intoByte414` is now tainted).
	intoByte414, _ := rsa.DecryptOAEP(nil, nil, nil, fromByte656, nil)

	// Return the tainted `intoByte414`:
	return intoByte414
}

func TaintStepTest_CryptoRsaDecryptPKCS1V15_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte518` into `intoByte650`.

	// Assume that `sourceCQL` has the underlying type of `fromByte518`:
	fromByte518 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte518` to result `intoByte650`
	// (`intoByte650` is now tainted).
	intoByte650, _ := rsa.DecryptPKCS1v15(nil, nil, fromByte518)

	// Return the tainted `intoByte650`:
	return intoByte650
}

func TaintStepTest_CryptoRsaEncryptOAEP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte784` into `intoByte957`.

	// Assume that `sourceCQL` has the underlying type of `fromByte784`:
	fromByte784 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte784` to result `intoByte957`
	// (`intoByte957` is now tainted).
	intoByte957, _ := rsa.EncryptOAEP(nil, nil, nil, fromByte784, nil)

	// Return the tainted `intoByte957`:
	return intoByte957
}

func TaintStepTest_CryptoRsaEncryptPKCS1V15_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte520` into `intoByte443`.

	// Assume that `sourceCQL` has the underlying type of `fromByte520`:
	fromByte520 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte520` to result `intoByte443`
	// (`intoByte443` is now tainted).
	intoByte443, _ := rsa.EncryptPKCS1v15(nil, nil, fromByte520)

	// Return the tainted `intoByte443`:
	return intoByte443
}

func TaintStepTest_CryptoRsaSignPKCS1V15_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte127` into `intoByte483`.

	// Assume that `sourceCQL` has the underlying type of `fromByte127`:
	fromByte127 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte127` to result `intoByte483`
	// (`intoByte483` is now tainted).
	intoByte483, _ := rsa.SignPKCS1v15(nil, nil, 0, fromByte127)

	// Return the tainted `intoByte483`:
	return intoByte483
}

func TaintStepTest_CryptoRsaSignPSS_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte989` into `intoByte982`.

	// Assume that `sourceCQL` has the underlying type of `fromByte989`:
	fromByte989 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte989` to result `intoByte982`
	// (`intoByte982` is now tainted).
	intoByte982, _ := rsa.SignPSS(nil, nil, 0, fromByte989, nil)

	// Return the tainted `intoByte982`:
	return intoByte982
}

func TaintStepTest_CryptoRsaPrivateKeyDecrypt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte417` into `intoByte584`.

	// Assume that `sourceCQL` has the underlying type of `fromByte417`:
	fromByte417 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL rsa.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `fromByte417` to the result `intoByte584`
	// (`intoByte584` is now tainted).
	intoByte584, _ := mediumObjCQL.Decrypt(nil, fromByte417, nil)

	// Return the tainted `intoByte584`:
	return intoByte584
}

func TaintStepTest_CryptoRsaPrivateKeySign_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte991` into `intoByte881`.

	// Assume that `sourceCQL` has the underlying type of `fromByte991`:
	fromByte991 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL rsa.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `fromByte991` to the result `intoByte881`
	// (`intoByte881` is now tainted).
	intoByte881, _ := mediumObjCQL.Sign(nil, fromByte991, nil)

	// Return the tainted `intoByte881`:
	return intoByte881
}

func RunAllTaints_CryptoRsa() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoRsaDecryptOAEP_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoRsaDecryptPKCS1V15_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoRsaEncryptOAEP_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoRsaEncryptPKCS1V15_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoRsaSignPKCS1V15_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoRsaSignPSS_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoRsaPrivateKeyDecrypt_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoRsaPrivateKeySign_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
