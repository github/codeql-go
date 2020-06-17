package main

import "crypto/cipher"

func TaintStepTest_CryptoCipherStreamReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStreamReader625` into `intoByte492`.

	// Assume that `sourceCQL` has the underlying type of `fromStreamReader625`:
	fromStreamReader625 := sourceCQL.(cipher.StreamReader)

	// Declare `intoByte492` variable:
	var intoByte492 []byte

	// Call the method that transfers the taint
	// from the receiver `fromStreamReader625` to the argument `intoByte492`
	// (`intoByte492` is now tainted).
	fromStreamReader625.Read(intoByte492)

	// Return the tainted `intoByte492`:
	return intoByte492
}

func TaintStepTest_CryptoCipherStreamWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte141` into `intoStreamWriter750`.

	// Assume that `sourceCQL` has the underlying type of `fromByte141`:
	fromByte141 := sourceCQL.([]byte)

	// Declare `intoStreamWriter750` variable:
	var intoStreamWriter750 cipher.StreamWriter

	// Call the method that transfers the taint
	// from the parameter `fromByte141` to the receiver `intoStreamWriter750`
	// (`intoStreamWriter750` is now tainted).
	intoStreamWriter750.Write(fromByte141)

	// Return the tainted `intoStreamWriter750`:
	return intoStreamWriter750
}

func TaintStepTest_CryptoCipherBlockModeCryptBlocks_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte358` into `intoByte427`.

	// Assume that `sourceCQL` has the underlying type of `fromByte358`:
	fromByte358 := sourceCQL.([]byte)

	// Declare `intoByte427` variable:
	var intoByte427 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.BlockMode

	// Call the method that transfers the taint
	// from the parameter `fromByte358` to the parameter `intoByte427`
	// (`intoByte427` is now tainted).
	mediumObjCQL.CryptBlocks(intoByte427, fromByte358)

	// Return the tainted `intoByte427`:
	return intoByte427
}

func TaintStepTest_CryptoCipherBlockDecrypt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte448` into `intoByte671`.

	// Assume that `sourceCQL` has the underlying type of `fromByte448`:
	fromByte448 := sourceCQL.([]byte)

	// Declare `intoByte671` variable:
	var intoByte671 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.Block

	// Call the method that transfers the taint
	// from the parameter `fromByte448` to the parameter `intoByte671`
	// (`intoByte671` is now tainted).
	mediumObjCQL.Decrypt(intoByte671, fromByte448)

	// Return the tainted `intoByte671`:
	return intoByte671
}

func TaintStepTest_CryptoCipherBlockEncrypt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte764` into `intoByte883`.

	// Assume that `sourceCQL` has the underlying type of `fromByte764`:
	fromByte764 := sourceCQL.([]byte)

	// Declare `intoByte883` variable:
	var intoByte883 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.Block

	// Call the method that transfers the taint
	// from the parameter `fromByte764` to the parameter `intoByte883`
	// (`intoByte883` is now tainted).
	mediumObjCQL.Encrypt(intoByte883, fromByte764)

	// Return the tainted `intoByte883`:
	return intoByte883
}

func TaintStepTest_CryptoCipherAEADOpen_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte959` into `intoByte506`.

	// Assume that `sourceCQL` has the underlying type of `fromByte959`:
	fromByte959 := sourceCQL.([]byte)

	// Declare `intoByte506` variable:
	var intoByte506 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.AEAD

	// Call the method that transfers the taint
	// from the parameter `fromByte959` to the parameter `intoByte506`
	// (`intoByte506` is now tainted).
	mediumObjCQL.Open(intoByte506, nil, fromByte959, nil)

	// Return the tainted `intoByte506`:
	return intoByte506
}

func TaintStepTest_CryptoCipherAEADOpen_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte512` into `intoByte401`.

	// Assume that `sourceCQL` has the underlying type of `fromByte512`:
	fromByte512 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL cipher.AEAD

	// Call the method that transfers the taint
	// from the parameter `fromByte512` to the result `intoByte401`
	// (`intoByte401` is now tainted).
	intoByte401, _ := mediumObjCQL.Open(nil, nil, fromByte512, nil)

	// Return the tainted `intoByte401`:
	return intoByte401
}

func TaintStepTest_CryptoCipherAEADSeal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte274` into `intoByte351`.

	// Assume that `sourceCQL` has the underlying type of `fromByte274`:
	fromByte274 := sourceCQL.([]byte)

	// Declare `intoByte351` variable:
	var intoByte351 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.AEAD

	// Call the method that transfers the taint
	// from the parameter `fromByte274` to the parameter `intoByte351`
	// (`intoByte351` is now tainted).
	mediumObjCQL.Seal(intoByte351, nil, fromByte274, nil)

	// Return the tainted `intoByte351`:
	return intoByte351
}

func TaintStepTest_CryptoCipherAEADSeal_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte389` into `intoByte527`.

	// Assume that `sourceCQL` has the underlying type of `fromByte389`:
	fromByte389 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL cipher.AEAD

	// Call the method that transfers the taint
	// from the parameter `fromByte389` to the result `intoByte527`
	// (`intoByte527` is now tainted).
	intoByte527 := mediumObjCQL.Seal(nil, nil, fromByte389, nil)

	// Return the tainted `intoByte527`:
	return intoByte527
}

func TaintStepTest_CryptoCipherStreamXORKeyStream_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte406` into `intoByte511`.

	// Assume that `sourceCQL` has the underlying type of `fromByte406`:
	fromByte406 := sourceCQL.([]byte)

	// Declare `intoByte511` variable:
	var intoByte511 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.Stream

	// Call the method that transfers the taint
	// from the parameter `fromByte406` to the parameter `intoByte511`
	// (`intoByte511` is now tainted).
	mediumObjCQL.XORKeyStream(intoByte511, fromByte406)

	// Return the tainted `intoByte511`:
	return intoByte511
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
