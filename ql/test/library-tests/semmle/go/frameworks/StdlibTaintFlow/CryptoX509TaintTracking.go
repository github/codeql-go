// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"crypto/x509"
	"encoding/pem"
)

func TaintStepTest_CryptoX509DecryptPEMBlock_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBlock611` into `intoByte115`.

	// Assume that `sourceCQL` has the underlying type of `fromBlock611`:
	fromBlock611 := sourceCQL.(*pem.Block)

	// Call the function that transfers the taint
	// from the parameter `fromBlock611` to result `intoByte115`
	// (`intoByte115` is now tainted).
	intoByte115, _ := x509.DecryptPEMBlock(fromBlock611, nil)

	// Return the tainted `intoByte115`:
	return intoByte115
}

func TaintStepTest_CryptoX509EncryptPEMBlock_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte497` into `intoBlock176`.

	// Assume that `sourceCQL` has the underlying type of `fromByte497`:
	fromByte497 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte497` to result `intoBlock176`
	// (`intoBlock176` is now tainted).
	intoBlock176, _ := x509.EncryptPEMBlock(nil, "", fromByte497, nil, 0)

	// Return the tainted `intoBlock176`:
	return intoBlock176
}

func TaintStepTest_CryptoX509CertPoolAddCert_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromCertificate789` into `intoCertPool537`.

	// Assume that `sourceCQL` has the underlying type of `fromCertificate789`:
	fromCertificate789 := sourceCQL.(*x509.Certificate)

	// Declare `intoCertPool537` variable:
	var intoCertPool537 x509.CertPool

	// Call the method that transfers the taint
	// from the parameter `fromCertificate789` to the receiver `intoCertPool537`
	// (`intoCertPool537` is now tainted).
	intoCertPool537.AddCert(fromCertificate789)

	// Return the tainted `intoCertPool537`:
	return intoCertPool537
}

func TaintStepTest_CryptoX509CertPoolAppendCertsFromPEM_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte570` into `intoCertPool344`.

	// Assume that `sourceCQL` has the underlying type of `fromByte570`:
	fromByte570 := sourceCQL.([]byte)

	// Declare `intoCertPool344` variable:
	var intoCertPool344 x509.CertPool

	// Call the method that transfers the taint
	// from the parameter `fromByte570` to the receiver `intoCertPool344`
	// (`intoCertPool344` is now tainted).
	intoCertPool344.AppendCertsFromPEM(fromByte570)

	// Return the tainted `intoCertPool344`:
	return intoCertPool344
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
