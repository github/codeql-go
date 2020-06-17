// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"encoding/pem"
	"io"
)

func TaintStepTest_EncodingPemDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte489` into `intoBlock311`.

	// Assume that `sourceCQL` has the underlying type of `fromByte489`:
	fromByte489 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte489` to result `intoBlock311`
	// (`intoBlock311` is now tainted).
	intoBlock311, _ := pem.Decode(fromByte489)

	// Return the tainted `intoBlock311`:
	return intoBlock311
}

func TaintStepTest_EncodingPemDecode_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte272` into `intoByte579`.

	// Assume that `sourceCQL` has the underlying type of `fromByte272`:
	fromByte272 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte272` to result `intoByte579`
	// (`intoByte579` is now tainted).
	_, intoByte579 := pem.Decode(fromByte272)

	// Return the tainted `intoByte579`:
	return intoByte579
}

func TaintStepTest_EncodingPemEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBlock409` into `intoWriter652`.

	// Assume that `sourceCQL` has the underlying type of `fromBlock409`:
	fromBlock409 := sourceCQL.(*pem.Block)

	// Declare `intoWriter652` variable:
	var intoWriter652 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromBlock409` to parameter `intoWriter652`;
	// `intoWriter652` is now tainted.
	pem.Encode(intoWriter652, fromBlock409)

	// Return the tainted `intoWriter652`:
	return intoWriter652
}

func TaintStepTest_EncodingPemEncodeToMemory_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBlock867` into `intoByte466`.

	// Assume that `sourceCQL` has the underlying type of `fromBlock867`:
	fromBlock867 := sourceCQL.(*pem.Block)

	// Call the function that transfers the taint
	// from the parameter `fromBlock867` to result `intoByte466`
	// (`intoByte466` is now tainted).
	intoByte466 := pem.EncodeToMemory(fromBlock867)

	// Return the tainted `intoByte466`:
	return intoByte466
}

func RunAllTaints_EncodingPem() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingPemDecode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingPemDecode_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingPemEncode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingPemEncodeToMemory_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
