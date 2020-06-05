package main

import "crypto/cipher"

func TaintStepTest_CryptoCipherStreamReaderRead(sourceCQL interface{}) {
	// The flow is from `from593` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `from593`:
	from593 := sourceCQL.(cipher.StreamReader)

	// Declare `dst` variable:
	var dst []byte

	// Call the method that transfers the taint
	// from the receiver `from593` to the argument `dst`
	// (`dst` is now tainted).
	from593.Read(dst)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_CryptoCipherStreamWriterWrite(sourceCQL interface{}) {
	// The flow is from `src` into `into197`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare `into197` variable:
	var into197 cipher.StreamWriter

	// Call the method that transfers the taint
	// from the parameter `src` to the receiver `into197`
	// (`into197` is now tainted).
	into197.Write(src)

	// Sink the tainted `into197`:
	sink(into197)
}

func TaintStepTest_CryptoCipherAEADOpen(sourceCQL interface{}) {
	// The flow is from `ciphertext` into `into909`.

	// Assume that `sourceCQL` has the underlying type of `ciphertext`:
	ciphertext := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL cipher.AEAD

	// Call the method that transfers the taint
	// from the parameter `ciphertext` to the result `into909`
	// (`into909` is now tainted).
	into909, _ := mediumObjCQL.Open(nil, nil, ciphertext, nil)

	// Sink the tainted `into909`:
	sink(into909)
}

func TaintStepTest_CryptoCipherBlockDecrypt(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare `dst` variable:
	var dst []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.Block

	// Call the method that transfers the taint
	// from the parameter `src` to the parameter `dst`
	// (`dst` is now tainted).
	mediumObjCQL.Decrypt(dst, src)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_CryptoCipherBlockEncrypt(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare `dst` variable:
	var dst []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.Block

	// Call the method that transfers the taint
	// from the parameter `src` to the parameter `dst`
	// (`dst` is now tainted).
	mediumObjCQL.Encrypt(dst, src)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_CryptoCipherBlockModeCryptBlocks(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare `dst` variable:
	var dst []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.BlockMode

	// Call the method that transfers the taint
	// from the parameter `src` to the parameter `dst`
	// (`dst` is now tainted).
	mediumObjCQL.CryptBlocks(dst, src)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_CryptoCipherStreamXORKeyStream(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare `dst` variable:
	var dst []byte

	// Declare medium object/interface:
	var mediumObjCQL cipher.Stream

	// Call the method that transfers the taint
	// from the parameter `src` to the parameter `dst`
	// (`dst` is now tainted).
	mediumObjCQL.XORKeyStream(dst, src)

	// Sink the tainted `dst`:
	sink(dst)
}

func RunAllTaints_CryptoCipher(v interface{}) {
	{
		source := newSource()
		TaintStepTest_CryptoCipherStreamReaderRead(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoCipherStreamWriterWrite(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoCipherAEADOpen(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoCipherBlockDecrypt(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoCipherBlockEncrypt(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoCipherBlockModeCryptBlocks(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoCipherStreamXORKeyStream(source)
	}
}
