package main

import (
	"crypto/x509"
	"encoding/pem"
)

func TaintStepTest_CryptoX509DecryptPEMBlock_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBlock744` into `intoByte890`.

	// Assume that `sourceCQL` has the underlying type of `fromBlock744`:
	fromBlock744 := sourceCQL.(*pem.Block)

	// Call the function that transfers the taint
	// from the parameter `fromBlock744` to result `intoByte890`
	// (`intoByte890` is now tainted).
	intoByte890, _ := x509.DecryptPEMBlock(fromBlock744, nil)

	// Sink the tainted `intoByte890`:
	sink(intoByte890)
}

func TaintStepTest_CryptoX509EncryptPEMBlock_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte898` into `intoBlock113`.

	// Assume that `sourceCQL` has the underlying type of `fromByte898`:
	fromByte898 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte898` to result `intoBlock113`
	// (`intoBlock113` is now tainted).
	intoBlock113, _ := x509.EncryptPEMBlock(nil, "", fromByte898, nil, 0)

	// Sink the tainted `intoBlock113`:
	sink(intoBlock113)
}

func TaintStepTest_CryptoX509CertPoolAddCert_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromCertificate310` into `intoCertPool799`.

	// Assume that `sourceCQL` has the underlying type of `fromCertificate310`:
	fromCertificate310 := sourceCQL.(*x509.Certificate)

	// Declare `intoCertPool799` variable:
	var intoCertPool799 x509.CertPool

	// Call the method that transfers the taint
	// from the parameter `fromCertificate310` to the receiver `intoCertPool799`
	// (`intoCertPool799` is now tainted).
	intoCertPool799.AddCert(fromCertificate310)

	// Sink the tainted `intoCertPool799`:
	sink(intoCertPool799)
}

func TaintStepTest_CryptoX509CertPoolAppendCertsFromPEM_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte793` into `intoCertPool180`.

	// Assume that `sourceCQL` has the underlying type of `fromByte793`:
	fromByte793 := sourceCQL.([]byte)

	// Declare `intoCertPool180` variable:
	var intoCertPool180 x509.CertPool

	// Call the method that transfers the taint
	// from the parameter `fromByte793` to the receiver `intoCertPool180`
	// (`intoCertPool180` is now tainted).
	intoCertPool180.AppendCertsFromPEM(fromByte793)

	// Sink the tainted `intoCertPool180`:
	sink(intoCertPool180)
}

func RunAllTaints_CryptoX509() {
	{
		source := newSource()
		TaintStepTest_CryptoX509DecryptPEMBlock_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoX509EncryptPEMBlock_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoX509CertPoolAddCert_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoX509CertPoolAppendCertsFromPEM_B0I0O0(source)
	}
}
