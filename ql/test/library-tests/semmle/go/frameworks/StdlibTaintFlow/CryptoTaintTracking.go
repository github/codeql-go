package main

import "crypto"

func TaintStepTest_CryptoDecrypterDecrypt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte982` into `intoByte248`.

	// Assume that `sourceCQL` has the underlying type of `fromByte982`:
	fromByte982 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL crypto.Decrypter

	// Call the method that transfers the taint
	// from the parameter `fromByte982` to the result `intoByte248`
	// (`intoByte248` is now tainted).
	intoByte248, _ := mediumObjCQL.Decrypt(nil, fromByte982, nil)

	// Return the tainted `intoByte248`:
	return intoByte248
}

func TaintStepTest_CryptoSignerSign_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte892` into `intoByte431`.

	// Assume that `sourceCQL` has the underlying type of `fromByte892`:
	fromByte892 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL crypto.Signer

	// Call the method that transfers the taint
	// from the parameter `fromByte892` to the result `intoByte431`
	// (`intoByte431` is now tainted).
	intoByte431, _ := mediumObjCQL.Sign(nil, fromByte892, nil)

	// Return the tainted `intoByte431`:
	return intoByte431
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
