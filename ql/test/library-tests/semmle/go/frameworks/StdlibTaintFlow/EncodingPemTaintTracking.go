package main

import (
	"encoding/pem"
	"io"
)

func TaintStepTest_EncodingPemDecode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte346` into `intoBlock285`.

	// Assume that `sourceCQL` has the underlying type of `fromByte346`:
	fromByte346 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte346` to result `intoBlock285`
	// (`intoBlock285` is now tainted).
	intoBlock285, _ := pem.Decode(fromByte346)

	// Sink the tainted `intoBlock285`:
	sink(intoBlock285)
}

func TaintStepTest_EncodingPemDecode_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromByte936` into `intoByte589`.

	// Assume that `sourceCQL` has the underlying type of `fromByte936`:
	fromByte936 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte936` to result `intoByte589`
	// (`intoByte589` is now tainted).
	_, intoByte589 := pem.Decode(fromByte936)

	// Sink the tainted `intoByte589`:
	sink(intoByte589)
}

func TaintStepTest_EncodingPemEncode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBlock417` into `intoWriter197`.

	// Assume that `sourceCQL` has the underlying type of `fromBlock417`:
	fromBlock417 := sourceCQL.(*pem.Block)

	// Declare `intoWriter197` variable:
	var intoWriter197 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromBlock417` to parameter `intoWriter197`;
	// `intoWriter197` is now tainted.
	pem.Encode(intoWriter197, fromBlock417)

	// Sink the tainted `intoWriter197`:
	sink(intoWriter197)
}

func TaintStepTest_EncodingPemEncodeToMemory_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBlock477` into `intoByte232`.

	// Assume that `sourceCQL` has the underlying type of `fromBlock477`:
	fromBlock477 := sourceCQL.(*pem.Block)

	// Call the function that transfers the taint
	// from the parameter `fromBlock477` to result `intoByte232`
	// (`intoByte232` is now tainted).
	intoByte232 := pem.EncodeToMemory(fromBlock477)

	// Sink the tainted `intoByte232`:
	sink(intoByte232)
}

func RunAllTaints_EncodingPem() {
	{
		source := newSource()
		TaintStepTest_EncodingPemDecode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingPemDecode_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingPemEncode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingPemEncodeToMemory_B0I0O0(source)
	}
}
