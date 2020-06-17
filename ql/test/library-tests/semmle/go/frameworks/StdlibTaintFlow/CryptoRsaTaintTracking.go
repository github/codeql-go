// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "crypto/rsa"

func TaintStepTest_CryptoRsaDecryptOAEP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte481` into `intoByte617`.

	// Assume that `sourceCQL` has the underlying type of `fromByte481`:
	fromByte481 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte481` to result `intoByte617`
	// (`intoByte617` is now tainted).
	intoByte617, _ := rsa.DecryptOAEP(nil, nil, nil, fromByte481, nil)

	// Return the tainted `intoByte617`:
	return intoByte617
}

func TaintStepTest_CryptoRsaDecryptPKCS1V15_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte344` into `intoByte884`.

	// Assume that `sourceCQL` has the underlying type of `fromByte344`:
	fromByte344 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte344` to result `intoByte884`
	// (`intoByte884` is now tainted).
	intoByte884, _ := rsa.DecryptPKCS1v15(nil, nil, fromByte344)

	// Return the tainted `intoByte884`:
	return intoByte884
}

func TaintStepTest_CryptoRsaEncryptOAEP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte360` into `intoByte873`.

	// Assume that `sourceCQL` has the underlying type of `fromByte360`:
	fromByte360 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte360` to result `intoByte873`
	// (`intoByte873` is now tainted).
	intoByte873, _ := rsa.EncryptOAEP(nil, nil, nil, fromByte360, nil)

	// Return the tainted `intoByte873`:
	return intoByte873
}

func TaintStepTest_CryptoRsaEncryptPKCS1V15_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte318` into `intoByte253`.

	// Assume that `sourceCQL` has the underlying type of `fromByte318`:
	fromByte318 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte318` to result `intoByte253`
	// (`intoByte253` is now tainted).
	intoByte253, _ := rsa.EncryptPKCS1v15(nil, nil, fromByte318)

	// Return the tainted `intoByte253`:
	return intoByte253
}

func TaintStepTest_CryptoRsaSignPKCS1V15_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte228` into `intoByte530`.

	// Assume that `sourceCQL` has the underlying type of `fromByte228`:
	fromByte228 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte228` to result `intoByte530`
	// (`intoByte530` is now tainted).
	intoByte530, _ := rsa.SignPKCS1v15(nil, nil, 0, fromByte228)

	// Return the tainted `intoByte530`:
	return intoByte530
}

func TaintStepTest_CryptoRsaSignPSS_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte853` into `intoByte355`.

	// Assume that `sourceCQL` has the underlying type of `fromByte853`:
	fromByte853 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte853` to result `intoByte355`
	// (`intoByte355` is now tainted).
	intoByte355, _ := rsa.SignPSS(nil, nil, 0, fromByte853, nil)

	// Return the tainted `intoByte355`:
	return intoByte355
}

func TaintStepTest_CryptoRsaPrivateKeyDecrypt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte349` into `intoByte402`.

	// Assume that `sourceCQL` has the underlying type of `fromByte349`:
	fromByte349 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL rsa.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `fromByte349` to the result `intoByte402`
	// (`intoByte402` is now tainted).
	intoByte402, _ := mediumObjCQL.Decrypt(nil, fromByte349, nil)

	// Return the tainted `intoByte402`:
	return intoByte402
}

func TaintStepTest_CryptoRsaPrivateKeySign_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte452` into `intoByte766`.

	// Assume that `sourceCQL` has the underlying type of `fromByte452`:
	fromByte452 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL rsa.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `fromByte452` to the result `intoByte766`
	// (`intoByte766` is now tainted).
	intoByte766, _ := mediumObjCQL.Sign(nil, fromByte452, nil)

	// Return the tainted `intoByte766`:
	return intoByte766
}

func RunAllTaints_CryptoRsa() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoRsaDecryptOAEP_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoRsaDecryptPKCS1V15_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoRsaEncryptOAEP_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoRsaEncryptPKCS1V15_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoRsaSignPKCS1V15_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoRsaSignPSS_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoRsaPrivateKeyDecrypt_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoRsaPrivateKeySign_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
