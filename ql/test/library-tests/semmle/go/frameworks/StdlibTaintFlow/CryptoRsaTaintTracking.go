package main

import "crypto/rsa"

func TaintStepTest_CryptoRsaDecryptOAEP_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte221` into `intoByte610`.

	// Assume that `sourceCQL` has the underlying type of `fromByte221`:
	fromByte221 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte221` to result `intoByte610`
	// (`intoByte610` is now tainted).
	intoByte610, _ := rsa.DecryptOAEP(nil, nil, nil, fromByte221, nil)

	// Sink the tainted `intoByte610`:
	sink(intoByte610)
}

func TaintStepTest_CryptoRsaDecryptPKCS1V15_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte234` into `intoByte114`.

	// Assume that `sourceCQL` has the underlying type of `fromByte234`:
	fromByte234 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte234` to result `intoByte114`
	// (`intoByte114` is now tainted).
	intoByte114, _ := rsa.DecryptPKCS1v15(nil, nil, fromByte234)

	// Sink the tainted `intoByte114`:
	sink(intoByte114)
}

func TaintStepTest_CryptoRsaEncryptOAEP_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte522` into `intoByte431`.

	// Assume that `sourceCQL` has the underlying type of `fromByte522`:
	fromByte522 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte522` to result `intoByte431`
	// (`intoByte431` is now tainted).
	intoByte431, _ := rsa.EncryptOAEP(nil, nil, nil, fromByte522, nil)

	// Sink the tainted `intoByte431`:
	sink(intoByte431)
}

func TaintStepTest_CryptoRsaEncryptPKCS1V15_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte809` into `intoByte950`.

	// Assume that `sourceCQL` has the underlying type of `fromByte809`:
	fromByte809 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte809` to result `intoByte950`
	// (`intoByte950` is now tainted).
	intoByte950, _ := rsa.EncryptPKCS1v15(nil, nil, fromByte809)

	// Sink the tainted `intoByte950`:
	sink(intoByte950)
}

func TaintStepTest_CryptoRsaSignPKCS1V15_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte271` into `intoByte368`.

	// Assume that `sourceCQL` has the underlying type of `fromByte271`:
	fromByte271 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte271` to result `intoByte368`
	// (`intoByte368` is now tainted).
	intoByte368, _ := rsa.SignPKCS1v15(nil, nil, 0, fromByte271)

	// Sink the tainted `intoByte368`:
	sink(intoByte368)
}

func TaintStepTest_CryptoRsaSignPSS_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte665` into `intoByte257`.

	// Assume that `sourceCQL` has the underlying type of `fromByte665`:
	fromByte665 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte665` to result `intoByte257`
	// (`intoByte257` is now tainted).
	intoByte257, _ := rsa.SignPSS(nil, nil, 0, fromByte665, nil)

	// Sink the tainted `intoByte257`:
	sink(intoByte257)
}

func TaintStepTest_CryptoRsaPrivateKeyDecrypt_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte820` into `intoByte558`.

	// Assume that `sourceCQL` has the underlying type of `fromByte820`:
	fromByte820 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL rsa.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `fromByte820` to the result `intoByte558`
	// (`intoByte558` is now tainted).
	intoByte558, _ := mediumObjCQL.Decrypt(nil, fromByte820, nil)

	// Sink the tainted `intoByte558`:
	sink(intoByte558)
}

func TaintStepTest_CryptoRsaPrivateKeySign_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte238` into `intoByte878`.

	// Assume that `sourceCQL` has the underlying type of `fromByte238`:
	fromByte238 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL rsa.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `fromByte238` to the result `intoByte878`
	// (`intoByte878` is now tainted).
	intoByte878, _ := mediumObjCQL.Sign(nil, fromByte238, nil)

	// Sink the tainted `intoByte878`:
	sink(intoByte878)
}

func RunAllTaints_CryptoRsa() {
	{
		source := newSource()
		TaintStepTest_CryptoRsaDecryptOAEP_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoRsaDecryptPKCS1V15_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoRsaEncryptOAEP_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoRsaEncryptPKCS1V15_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoRsaSignPKCS1V15_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoRsaSignPSS_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoRsaPrivateKeyDecrypt_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoRsaPrivateKeySign_B0I0O0(source)
	}
}
