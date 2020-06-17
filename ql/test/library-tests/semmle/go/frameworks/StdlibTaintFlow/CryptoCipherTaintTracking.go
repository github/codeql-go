// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "crypto/cipher"

func TaintStepTest_CryptoCipherStreamReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStreamReader256` into `intoByte656`.

	// Assume that `sourceCQL` has the underlying type of `fromStreamReader256`:
	fromStreamReader256 := sourceCQL.(cipher.StreamReader)

	// Declare `intoByte656` variable:
	var intoByte656 []byte

	// Call the method that transfers the taint
	// from the receiver `fromStreamReader256` to the argument `intoByte656`
	// (`intoByte656` is now tainted).
	fromStreamReader256.Read(intoByte656)

	// Return the tainted `intoByte656`:
	return intoByte656
}

func TaintStepTest_CryptoCipherStreamWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte947` into `intoStreamWriter233`.

	// Assume that `sourceCQL` has the underlying type of `fromByte947`:
	fromByte947 := sourceCQL.([]byte)

	// Declare `intoStreamWriter233` variable:
	var intoStreamWriter233 cipher.StreamWriter

	// Call the method that transfers the taint
	// from the parameter `fromByte947` to the receiver `intoStreamWriter233`
	// (`intoStreamWriter233` is now tainted).
	intoStreamWriter233.Write(fromByte947)

	// Return the tainted `intoStreamWriter233`:
	return intoStreamWriter233
}

func TaintStepTest_CryptoCipherBlockModeCryptBlocks_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte661` into `intoByte870`.

	// Assume that `sourceCQL` has the underlying type of `fromByte661`:
	fromByte661 := sourceCQL.([]byte)

	// Declare `intoByte870` variable:
	var intoByte870 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.BlockMode

	// Call the method that transfers the taint
	// from the parameter `fromByte661` to the parameter `intoByte870`
	// (`intoByte870` is now tainted).
	mediumObjCQL.CryptBlocks(intoByte870, fromByte661)

	// Return the tainted `intoByte870`:
	return intoByte870
}

func TaintStepTest_CryptoCipherBlockDecrypt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte141` into `intoByte386`.

	// Assume that `sourceCQL` has the underlying type of `fromByte141`:
	fromByte141 := sourceCQL.([]byte)

	// Declare `intoByte386` variable:
	var intoByte386 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.Block

	// Call the method that transfers the taint
	// from the parameter `fromByte141` to the parameter `intoByte386`
	// (`intoByte386` is now tainted).
	mediumObjCQL.Decrypt(intoByte386, fromByte141)

	// Return the tainted `intoByte386`:
	return intoByte386
}

func TaintStepTest_CryptoCipherBlockEncrypt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte556` into `intoByte137`.

	// Assume that `sourceCQL` has the underlying type of `fromByte556`:
	fromByte556 := sourceCQL.([]byte)

	// Declare `intoByte137` variable:
	var intoByte137 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.Block

	// Call the method that transfers the taint
	// from the parameter `fromByte556` to the parameter `intoByte137`
	// (`intoByte137` is now tainted).
	mediumObjCQL.Encrypt(intoByte137, fromByte556)

	// Return the tainted `intoByte137`:
	return intoByte137
}

func TaintStepTest_CryptoCipherAEADOpen_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte513` into `intoByte142`.

	// Assume that `sourceCQL` has the underlying type of `fromByte513`:
	fromByte513 := sourceCQL.([]byte)

	// Declare `intoByte142` variable:
	var intoByte142 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.AEAD

	// Call the method that transfers the taint
	// from the parameter `fromByte513` to the parameter `intoByte142`
	// (`intoByte142` is now tainted).
	mediumObjCQL.Open(intoByte142, nil, fromByte513, nil)

	// Return the tainted `intoByte142`:
	return intoByte142
}

func TaintStepTest_CryptoCipherAEADOpen_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte436` into `intoByte992`.

	// Assume that `sourceCQL` has the underlying type of `fromByte436`:
	fromByte436 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL cipher.AEAD

	// Call the method that transfers the taint
	// from the parameter `fromByte436` to the result `intoByte992`
	// (`intoByte992` is now tainted).
	intoByte992, _ := mediumObjCQL.Open(nil, nil, fromByte436, nil)

	// Return the tainted `intoByte992`:
	return intoByte992
}

func TaintStepTest_CryptoCipherAEADSeal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte224` into `intoByte278`.

	// Assume that `sourceCQL` has the underlying type of `fromByte224`:
	fromByte224 := sourceCQL.([]byte)

	// Declare `intoByte278` variable:
	var intoByte278 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.AEAD

	// Call the method that transfers the taint
	// from the parameter `fromByte224` to the parameter `intoByte278`
	// (`intoByte278` is now tainted).
	mediumObjCQL.Seal(intoByte278, nil, fromByte224, nil)

	// Return the tainted `intoByte278`:
	return intoByte278
}

func TaintStepTest_CryptoCipherAEADSeal_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte320` into `intoByte900`.

	// Assume that `sourceCQL` has the underlying type of `fromByte320`:
	fromByte320 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL cipher.AEAD

	// Call the method that transfers the taint
	// from the parameter `fromByte320` to the result `intoByte900`
	// (`intoByte900` is now tainted).
	intoByte900 := mediumObjCQL.Seal(nil, nil, fromByte320, nil)

	// Return the tainted `intoByte900`:
	return intoByte900
}

func TaintStepTest_CryptoCipherStreamXORKeyStream_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte899` into `intoByte823`.

	// Assume that `sourceCQL` has the underlying type of `fromByte899`:
	fromByte899 := sourceCQL.([]byte)

	// Declare `intoByte823` variable:
	var intoByte823 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.Stream

	// Call the method that transfers the taint
	// from the parameter `fromByte899` to the parameter `intoByte823`
	// (`intoByte823` is now tainted).
	mediumObjCQL.XORKeyStream(intoByte823, fromByte899)

	// Return the tainted `intoByte823`:
	return intoByte823
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
