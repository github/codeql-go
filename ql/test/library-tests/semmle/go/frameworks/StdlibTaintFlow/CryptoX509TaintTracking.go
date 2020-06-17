package main

import (
	"crypto/x509"
	"encoding/pem"
)

func TaintStepTest_CryptoX509DecryptPEMBlock_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBlock129` into `intoByte456`.

	// Assume that `sourceCQL` has the underlying type of `fromBlock129`:
	fromBlock129 := sourceCQL.(*pem.Block)

	// Call the function that transfers the taint
	// from the parameter `fromBlock129` to result `intoByte456`
	// (`intoByte456` is now tainted).
	intoByte456, _ := x509.DecryptPEMBlock(fromBlock129, nil)

	// Return the tainted `intoByte456`:
	return intoByte456
}

func TaintStepTest_CryptoX509EncryptPEMBlock_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte375` into `intoBlock851`.

	// Assume that `sourceCQL` has the underlying type of `fromByte375`:
	fromByte375 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte375` to result `intoBlock851`
	// (`intoBlock851` is now tainted).
	intoBlock851, _ := x509.EncryptPEMBlock(nil, "", fromByte375, nil, 0)

	// Return the tainted `intoBlock851`:
	return intoBlock851
}

func TaintStepTest_CryptoX509CertPoolAddCert_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromCertificate432` into `intoCertPool128`.

	// Assume that `sourceCQL` has the underlying type of `fromCertificate432`:
	fromCertificate432 := sourceCQL.(*x509.Certificate)

	// Declare `intoCertPool128` variable:
	var intoCertPool128 x509.CertPool

	// Call the method that transfers the taint
	// from the parameter `fromCertificate432` to the receiver `intoCertPool128`
	// (`intoCertPool128` is now tainted).
	intoCertPool128.AddCert(fromCertificate432)

	// Return the tainted `intoCertPool128`:
	return intoCertPool128
}

func TaintStepTest_CryptoX509CertPoolAppendCertsFromPEM_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte258` into `intoCertPool911`.

	// Assume that `sourceCQL` has the underlying type of `fromByte258`:
	fromByte258 := sourceCQL.([]byte)

	// Declare `intoCertPool911` variable:
	var intoCertPool911 x509.CertPool

	// Call the method that transfers the taint
	// from the parameter `fromByte258` to the receiver `intoCertPool911`
	// (`intoCertPool911` is now tainted).
	intoCertPool911.AppendCertsFromPEM(fromByte258)

	// Return the tainted `intoCertPool911`:
	return intoCertPool911
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
