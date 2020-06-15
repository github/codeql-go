package main

import "crypto/cipher"

func TaintStepTest_CryptoCipherStreamReaderRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromStreamReader600` into `intoByte757`.

	// Assume that `sourceCQL` has the underlying type of `fromStreamReader600`:
	fromStreamReader600 := sourceCQL.(cipher.StreamReader)

	// Declare `intoByte757` variable:
	var intoByte757 []byte

	// Call the method that transfers the taint
	// from the receiver `fromStreamReader600` to the argument `intoByte757`
	// (`intoByte757` is now tainted).
	fromStreamReader600.Read(intoByte757)

	// Sink the tainted `intoByte757`:
	sink(intoByte757)
}

func TaintStepTest_CryptoCipherStreamWriterWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte349` into `intoStreamWriter619`.

	// Assume that `sourceCQL` has the underlying type of `fromByte349`:
	fromByte349 := sourceCQL.([]byte)

	// Declare `intoStreamWriter619` variable:
	var intoStreamWriter619 cipher.StreamWriter

	// Call the method that transfers the taint
	// from the parameter `fromByte349` to the receiver `intoStreamWriter619`
	// (`intoStreamWriter619` is now tainted).
	intoStreamWriter619.Write(fromByte349)

	// Sink the tainted `intoStreamWriter619`:
	sink(intoStreamWriter619)
}

func TaintStepTest_CryptoCipherBlockModeCryptBlocks_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte804` into `intoByte359`.

	// Assume that `sourceCQL` has the underlying type of `fromByte804`:
	fromByte804 := sourceCQL.([]byte)

	// Declare `intoByte359` variable:
	var intoByte359 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.BlockMode

	// Call the method that transfers the taint
	// from the parameter `fromByte804` to the parameter `intoByte359`
	// (`intoByte359` is now tainted).
	mediumObjCQL.CryptBlocks(intoByte359, fromByte804)

	// Sink the tainted `intoByte359`:
	sink(intoByte359)
}

func TaintStepTest_CryptoCipherBlockDecrypt_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte925` into `intoByte262`.

	// Assume that `sourceCQL` has the underlying type of `fromByte925`:
	fromByte925 := sourceCQL.([]byte)

	// Declare `intoByte262` variable:
	var intoByte262 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.Block

	// Call the method that transfers the taint
	// from the parameter `fromByte925` to the parameter `intoByte262`
	// (`intoByte262` is now tainted).
	mediumObjCQL.Decrypt(intoByte262, fromByte925)

	// Sink the tainted `intoByte262`:
	sink(intoByte262)
}

func TaintStepTest_CryptoCipherBlockEncrypt_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte147` into `intoByte445`.

	// Assume that `sourceCQL` has the underlying type of `fromByte147`:
	fromByte147 := sourceCQL.([]byte)

	// Declare `intoByte445` variable:
	var intoByte445 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.Block

	// Call the method that transfers the taint
	// from the parameter `fromByte147` to the parameter `intoByte445`
	// (`intoByte445` is now tainted).
	mediumObjCQL.Encrypt(intoByte445, fromByte147)

	// Sink the tainted `intoByte445`:
	sink(intoByte445)
}

func TaintStepTest_CryptoCipherAEADOpen_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte217` into `intoByte172`.

	// Assume that `sourceCQL` has the underlying type of `fromByte217`:
	fromByte217 := sourceCQL.([]byte)

	// Declare `intoByte172` variable:
	var intoByte172 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.AEAD

	// Call the method that transfers the taint
	// from the parameter `fromByte217` to the parameter `intoByte172`
	// (`intoByte172` is now tainted).
	mediumObjCQL.Open(intoByte172, nil, fromByte217, nil)

	// Sink the tainted `intoByte172`:
	sink(intoByte172)
}

func TaintStepTest_CryptoCipherAEADOpen_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromByte472` into `intoByte754`.

	// Assume that `sourceCQL` has the underlying type of `fromByte472`:
	fromByte472 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL cipher.AEAD

	// Call the method that transfers the taint
	// from the parameter `fromByte472` to the result `intoByte754`
	// (`intoByte754` is now tainted).
	intoByte754, _ := mediumObjCQL.Open(nil, nil, fromByte472, nil)

	// Sink the tainted `intoByte754`:
	sink(intoByte754)
}

func TaintStepTest_CryptoCipherAEADSeal_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte800` into `intoByte820`.

	// Assume that `sourceCQL` has the underlying type of `fromByte800`:
	fromByte800 := sourceCQL.([]byte)

	// Declare `intoByte820` variable:
	var intoByte820 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.AEAD

	// Call the method that transfers the taint
	// from the parameter `fromByte800` to the parameter `intoByte820`
	// (`intoByte820` is now tainted).
	mediumObjCQL.Seal(intoByte820, nil, fromByte800, nil)

	// Sink the tainted `intoByte820`:
	sink(intoByte820)
}

func TaintStepTest_CryptoCipherAEADSeal_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromByte746` into `intoByte552`.

	// Assume that `sourceCQL` has the underlying type of `fromByte746`:
	fromByte746 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL cipher.AEAD

	// Call the method that transfers the taint
	// from the parameter `fromByte746` to the result `intoByte552`
	// (`intoByte552` is now tainted).
	intoByte552 := mediumObjCQL.Seal(nil, nil, fromByte746, nil)

	// Sink the tainted `intoByte552`:
	sink(intoByte552)
}

func TaintStepTest_CryptoCipherStreamXORKeyStream_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte705` into `intoByte522`.

	// Assume that `sourceCQL` has the underlying type of `fromByte705`:
	fromByte705 := sourceCQL.([]byte)

	// Declare `intoByte522` variable:
	var intoByte522 []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.Stream

	// Call the method that transfers the taint
	// from the parameter `fromByte705` to the parameter `intoByte522`
	// (`intoByte522` is now tainted).
	mediumObjCQL.XORKeyStream(intoByte522, fromByte705)

	// Sink the tainted `intoByte522`:
	sink(intoByte522)
}

func RunAllTaints_CryptoCipher() {
	{
		source := newSource()
		TaintStepTest_CryptoCipherStreamReaderRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoCipherStreamWriterWrite_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoCipherBlockModeCryptBlocks_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoCipherBlockDecrypt_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoCipherBlockEncrypt_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoCipherAEADOpen_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoCipherAEADOpen_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoCipherAEADSeal_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoCipherAEADSeal_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoCipherStreamXORKeyStream_B0I0O0(source)
	}
}
