// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "crypto"

func TaintStepTest_CryptoDecrypterDecrypt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte656` into `intoByte414`.

	// Assume that `sourceCQL` has the underlying type of `fromByte656`:
	fromByte656 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL crypto.Decrypter

	// Call the method that transfers the taint
	// from the parameter `fromByte656` to the result `intoByte414`
	// (`intoByte414` is now tainted).
	intoByte414, _ := mediumObjCQL.Decrypt(nil, fromByte656, nil)

	// Return the tainted `intoByte414`:
	return intoByte414
}

func TaintStepTest_CryptoSignerSign_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte518` into `intoByte650`.

	// Assume that `sourceCQL` has the underlying type of `fromByte518`:
	fromByte518 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL crypto.Signer

	// Call the method that transfers the taint
	// from the parameter `fromByte518` to the result `intoByte650`
	// (`intoByte650` is now tainted).
	intoByte650, _ := mediumObjCQL.Sign(nil, fromByte518, nil)

	// Return the tainted `intoByte650`:
	return intoByte650
}

func RunAllTaints_Crypto() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoDecrypterDecrypt_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoSignerSign_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
