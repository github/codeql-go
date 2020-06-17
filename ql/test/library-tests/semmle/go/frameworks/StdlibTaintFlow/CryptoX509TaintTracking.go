// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"crypto/x509"
	"encoding/pem"
)

func TaintStepTest_CryptoX509DecryptPEMBlock_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBlock656` into `intoByte414`.

	// Assume that `sourceCQL` has the underlying type of `fromBlock656`:
	fromBlock656 := sourceCQL.(*pem.Block)

	// Call the function that transfers the taint
	// from the parameter `fromBlock656` to result `intoByte414`
	// (`intoByte414` is now tainted).
	intoByte414, _ := x509.DecryptPEMBlock(fromBlock656, nil)

	// Return the tainted `intoByte414`:
	return intoByte414
}

func TaintStepTest_CryptoX509EncryptPEMBlock_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte518` into `intoBlock650`.

	// Assume that `sourceCQL` has the underlying type of `fromByte518`:
	fromByte518 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte518` to result `intoBlock650`
	// (`intoBlock650` is now tainted).
	intoBlock650, _ := x509.EncryptPEMBlock(nil, "", fromByte518, nil, 0)

	// Return the tainted `intoBlock650`:
	return intoBlock650
}

func TaintStepTest_CryptoX509CertPoolAddCert_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromCertificate784` into `intoCertPool957`.

	// Assume that `sourceCQL` has the underlying type of `fromCertificate784`:
	fromCertificate784 := sourceCQL.(*x509.Certificate)

	// Declare `intoCertPool957` variable:
	var intoCertPool957 x509.CertPool

	// Call the method that transfers the taint
	// from the parameter `fromCertificate784` to the receiver `intoCertPool957`
	// (`intoCertPool957` is now tainted).
	intoCertPool957.AddCert(fromCertificate784)

	// Return the tainted `intoCertPool957`:
	return intoCertPool957
}

func TaintStepTest_CryptoX509CertPoolAppendCertsFromPEM_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte520` into `intoCertPool443`.

	// Assume that `sourceCQL` has the underlying type of `fromByte520`:
	fromByte520 := sourceCQL.([]byte)

	// Declare `intoCertPool443` variable:
	var intoCertPool443 x509.CertPool

	// Call the method that transfers the taint
	// from the parameter `fromByte520` to the receiver `intoCertPool443`
	// (`intoCertPool443` is now tainted).
	intoCertPool443.AppendCertsFromPEM(fromByte520)

	// Return the tainted `intoCertPool443`:
	return intoCertPool443
}

func RunAllTaints_CryptoX509() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoX509DecryptPEMBlock_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoX509EncryptPEMBlock_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoX509CertPoolAddCert_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoX509CertPoolAppendCertsFromPEM_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
