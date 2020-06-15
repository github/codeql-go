package main

import "crypto"

func TaintStepTest_CryptoDecrypterDecrypt_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte811` into `intoByte501`.

	// Assume that `sourceCQL` has the underlying type of `fromByte811`:
	fromByte811 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL crypto.Decrypter

	// Call the method that transfers the taint
	// from the parameter `fromByte811` to the result `intoByte501`
	// (`intoByte501` is now tainted).
	intoByte501, _ := mediumObjCQL.Decrypt(nil, fromByte811, nil)

	// Sink the tainted `intoByte501`:
	sink(intoByte501)
}

func TaintStepTest_CryptoSignerSign_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte634` into `intoByte359`.

	// Assume that `sourceCQL` has the underlying type of `fromByte634`:
	fromByte634 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL crypto.Signer

	// Call the method that transfers the taint
	// from the parameter `fromByte634` to the result `intoByte359`
	// (`intoByte359` is now tainted).
	intoByte359, _ := mediumObjCQL.Sign(nil, fromByte634, nil)

	// Sink the tainted `intoByte359`:
	sink(intoByte359)
}

func RunAllTaints_Crypto() {
	{
		source := newSource()
		TaintStepTest_CryptoDecrypterDecrypt_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoSignerSign_B0I0O0(source)
	}
}
