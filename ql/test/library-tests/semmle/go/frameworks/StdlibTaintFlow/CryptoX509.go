package main

import (
	"crypto/x509"
	"encoding/pem"
)

func TaintStepTest_CryptoX509DecryptPEMBlock(sourceCQL interface{}) {
	// The flow is from `b` into `into798`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.(*pem.Block)

	// Call the function that transfers the taint
	// from the parameter `b` to result `into798`
	// (`into798` is now tainted).
	into798, _ := x509.DecryptPEMBlock(b, nil)

	// Sink the tainted `into798`:
	sink(into798)
}

func TaintStepTest_CryptoX509EncryptPEMBlock(sourceCQL interface{}) {
	// The flow is from `data` into `into686`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `data` to result `into686`
	// (`into686` is now tainted).
	into686, _ := x509.EncryptPEMBlock(nil, "", data, nil, 0)

	// Sink the tainted `into686`:
	sink(into686)
}

func RunAllTaints_CryptoX509(v interface{}) {
	{
		source := newSource()
		TaintStepTest_CryptoX509DecryptPEMBlock(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoX509EncryptPEMBlock(source)
	}
}
