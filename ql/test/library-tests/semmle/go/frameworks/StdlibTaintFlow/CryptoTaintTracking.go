// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "crypto"

func TaintStepTest_CryptoDecrypterDecrypt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte491` into `intoByte867`.

	// Assume that `sourceCQL` has the underlying type of `fromByte491`:
	fromByte491 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL crypto.Decrypter

	// Call the method that transfers the taint
	// from the parameter `fromByte491` to the result `intoByte867`
	// (`intoByte867` is now tainted).
	intoByte867, _ := mediumObjCQL.Decrypt(nil, fromByte491, nil)

	// Return the tainted `intoByte867`:
	return intoByte867
}

func TaintStepTest_CryptoSignerSign_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte554` into `intoByte331`.

	// Assume that `sourceCQL` has the underlying type of `fromByte554`:
	fromByte554 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL crypto.Signer

	// Call the method that transfers the taint
	// from the parameter `fromByte554` to the result `intoByte331`
	// (`intoByte331` is now tainted).
	intoByte331, _ := mediumObjCQL.Sign(nil, fromByte554, nil)

	// Return the tainted `intoByte331`:
	return intoByte331
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
