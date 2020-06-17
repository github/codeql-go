package main

import "crypto/rsa"

func TaintStepTest_CryptoRsaDecryptOAEP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte691` into `intoByte976`.

	// Assume that `sourceCQL` has the underlying type of `fromByte691`:
	fromByte691 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte691` to result `intoByte976`
	// (`intoByte976` is now tainted).
	intoByte976, _ := rsa.DecryptOAEP(nil, nil, nil, fromByte691, nil)

	// Return the tainted `intoByte976`:
	return intoByte976
}

func TaintStepTest_CryptoRsaDecryptPKCS1V15_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte480` into `intoByte970`.

	// Assume that `sourceCQL` has the underlying type of `fromByte480`:
	fromByte480 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte480` to result `intoByte970`
	// (`intoByte970` is now tainted).
	intoByte970, _ := rsa.DecryptPKCS1v15(nil, nil, fromByte480)

	// Return the tainted `intoByte970`:
	return intoByte970
}

func TaintStepTest_CryptoRsaEncryptOAEP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte655` into `intoByte327`.

	// Assume that `sourceCQL` has the underlying type of `fromByte655`:
	fromByte655 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte655` to result `intoByte327`
	// (`intoByte327` is now tainted).
	intoByte327, _ := rsa.EncryptOAEP(nil, nil, nil, fromByte655, nil)

	// Return the tainted `intoByte327`:
	return intoByte327
}

func TaintStepTest_CryptoRsaEncryptPKCS1V15_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte181` into `intoByte953`.

	// Assume that `sourceCQL` has the underlying type of `fromByte181`:
	fromByte181 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte181` to result `intoByte953`
	// (`intoByte953` is now tainted).
	intoByte953, _ := rsa.EncryptPKCS1v15(nil, nil, fromByte181)

	// Return the tainted `intoByte953`:
	return intoByte953
}

func TaintStepTest_CryptoRsaSignPKCS1V15_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte459` into `intoByte740`.

	// Assume that `sourceCQL` has the underlying type of `fromByte459`:
	fromByte459 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte459` to result `intoByte740`
	// (`intoByte740` is now tainted).
	intoByte740, _ := rsa.SignPKCS1v15(nil, nil, 0, fromByte459)

	// Return the tainted `intoByte740`:
	return intoByte740
}

func TaintStepTest_CryptoRsaSignPSS_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte742` into `intoByte514`.

	// Assume that `sourceCQL` has the underlying type of `fromByte742`:
	fromByte742 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte742` to result `intoByte514`
	// (`intoByte514` is now tainted).
	intoByte514, _ := rsa.SignPSS(nil, nil, 0, fromByte742, nil)

	// Return the tainted `intoByte514`:
	return intoByte514
}

func TaintStepTest_CryptoRsaPrivateKeyDecrypt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte115` into `intoByte823`.

	// Assume that `sourceCQL` has the underlying type of `fromByte115`:
	fromByte115 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL rsa.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `fromByte115` to the result `intoByte823`
	// (`intoByte823` is now tainted).
	intoByte823, _ := mediumObjCQL.Decrypt(nil, fromByte115, nil)

	// Return the tainted `intoByte823`:
	return intoByte823
}

func TaintStepTest_CryptoRsaPrivateKeySign_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte699` into `intoByte819`.

	// Assume that `sourceCQL` has the underlying type of `fromByte699`:
	fromByte699 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL rsa.PrivateKey

	// Call the method that transfers the taint
	// from the parameter `fromByte699` to the result `intoByte819`
	// (`intoByte819` is now tainted).
	intoByte819, _ := mediumObjCQL.Sign(nil, fromByte699, nil)

	// Return the tainted `intoByte819`:
	return intoByte819
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
