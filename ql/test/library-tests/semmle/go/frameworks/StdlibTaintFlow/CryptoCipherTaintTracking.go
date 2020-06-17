// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "crypto/cipher"

func TaintStepTest_CryptoCipherStreamReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStreamReader656` into `intoByte414`.

	// Assume that `sourceCQL` has the underlying type of `fromStreamReader656`:
	fromStreamReader656 := sourceCQL.(cipher.StreamReader)

	// Declare `intoByte414` variable:
	var intoByte414 []byte

	// Call the method that transfers the taint
	// from the receiver `fromStreamReader656` to the argument `intoByte414`
	// (`intoByte414` is now tainted).
	fromStreamReader656.Read(intoByte414)

	// Return the tainted `intoByte414`:
	return intoByte414
}

func TaintStepTest_CryptoCipherStreamWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte518` into `intoStreamWriter650`.

	// Assume that `sourceCQL` has the underlying type of `fromByte518`:
	fromByte518 := sourceCQL.([]byte)

	// Declare `intoStreamWriter650` variable:
	var intoStreamWriter650 cipher.StreamWriter

	// Call the method that transfers the taint
	// from the parameter `fromByte518` to the receiver `intoStreamWriter650`
	// (`intoStreamWriter650` is now tainted).
	intoStreamWriter650.Write(fromByte518)

	// Return the tainted `intoStreamWriter650`:
	return intoStreamWriter650
}

func TaintStepTest_CryptoCipherBlockModeCryptBlocks_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte784` into `intoByte957`.

	// Assume that `sourceCQL` has the underlying type of `fromByte784`:
	fromByte784 := sourceCQL.([]byte)

	// Declare `intoByte957` variable:
	var intoByte957 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.BlockMode

	// Call the method that transfers the taint
	// from the parameter `fromByte784` to the parameter `intoByte957`
	// (`intoByte957` is now tainted).
	mediumObjCQL.CryptBlocks(intoByte957, fromByte784)

	// Return the tainted `intoByte957`:
	return intoByte957
}

func TaintStepTest_CryptoCipherBlockDecrypt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte520` into `intoByte443`.

	// Assume that `sourceCQL` has the underlying type of `fromByte520`:
	fromByte520 := sourceCQL.([]byte)

	// Declare `intoByte443` variable:
	var intoByte443 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.Block

	// Call the method that transfers the taint
	// from the parameter `fromByte520` to the parameter `intoByte443`
	// (`intoByte443` is now tainted).
	mediumObjCQL.Decrypt(intoByte443, fromByte520)

	// Return the tainted `intoByte443`:
	return intoByte443
}

func TaintStepTest_CryptoCipherBlockEncrypt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte127` into `intoByte483`.

	// Assume that `sourceCQL` has the underlying type of `fromByte127`:
	fromByte127 := sourceCQL.([]byte)

	// Declare `intoByte483` variable:
	var intoByte483 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.Block

	// Call the method that transfers the taint
	// from the parameter `fromByte127` to the parameter `intoByte483`
	// (`intoByte483` is now tainted).
	mediumObjCQL.Encrypt(intoByte483, fromByte127)

	// Return the tainted `intoByte483`:
	return intoByte483
}

func TaintStepTest_CryptoCipherAEADOpen_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte989` into `intoByte982`.

	// Assume that `sourceCQL` has the underlying type of `fromByte989`:
	fromByte989 := sourceCQL.([]byte)

	// Declare `intoByte982` variable:
	var intoByte982 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.AEAD

	// Call the method that transfers the taint
	// from the parameter `fromByte989` to the parameter `intoByte982`
	// (`intoByte982` is now tainted).
	mediumObjCQL.Open(intoByte982, nil, fromByte989, nil)

	// Return the tainted `intoByte982`:
	return intoByte982
}

func TaintStepTest_CryptoCipherAEADOpen_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte417` into `intoByte584`.

	// Assume that `sourceCQL` has the underlying type of `fromByte417`:
	fromByte417 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL cipher.AEAD

	// Call the method that transfers the taint
	// from the parameter `fromByte417` to the result `intoByte584`
	// (`intoByte584` is now tainted).
	intoByte584, _ := mediumObjCQL.Open(nil, nil, fromByte417, nil)

	// Return the tainted `intoByte584`:
	return intoByte584
}

func TaintStepTest_CryptoCipherAEADSeal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte991` into `intoByte881`.

	// Assume that `sourceCQL` has the underlying type of `fromByte991`:
	fromByte991 := sourceCQL.([]byte)

	// Declare `intoByte881` variable:
	var intoByte881 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.AEAD

	// Call the method that transfers the taint
	// from the parameter `fromByte991` to the parameter `intoByte881`
	// (`intoByte881` is now tainted).
	mediumObjCQL.Seal(intoByte881, nil, fromByte991, nil)

	// Return the tainted `intoByte881`:
	return intoByte881
}

func TaintStepTest_CryptoCipherAEADSeal_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte186` into `intoByte284`.

	// Assume that `sourceCQL` has the underlying type of `fromByte186`:
	fromByte186 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL cipher.AEAD

	// Call the method that transfers the taint
	// from the parameter `fromByte186` to the result `intoByte284`
	// (`intoByte284` is now tainted).
	intoByte284 := mediumObjCQL.Seal(nil, nil, fromByte186, nil)

	// Return the tainted `intoByte284`:
	return intoByte284
}

func TaintStepTest_CryptoCipherStreamXORKeyStream_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte908` into `intoByte137`.

	// Assume that `sourceCQL` has the underlying type of `fromByte908`:
	fromByte908 := sourceCQL.([]byte)

	// Declare `intoByte137` variable:
	var intoByte137 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.Stream

	// Call the method that transfers the taint
	// from the parameter `fromByte908` to the parameter `intoByte137`
	// (`intoByte137` is now tainted).
	mediumObjCQL.XORKeyStream(intoByte137, fromByte908)

	// Return the tainted `intoByte137`:
	return intoByte137
}

func RunAllTaints_CryptoCipher() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoCipherStreamReaderRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoCipherStreamWriterWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoCipherBlockModeCryptBlocks_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoCipherBlockDecrypt_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoCipherBlockEncrypt_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoCipherAEADOpen_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoCipherAEADOpen_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoCipherAEADSeal_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoCipherAEADSeal_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoCipherStreamXORKeyStream_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
