// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"encoding/pem"
	"io"
)

func TaintStepTest_EncodingPemDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte656` into `intoBlock414`.

	// Assume that `sourceCQL` has the underlying type of `fromByte656`:
	fromByte656 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte656` to result `intoBlock414`
	// (`intoBlock414` is now tainted).
	intoBlock414, _ := pem.Decode(fromByte656)

	// Return the tainted `intoBlock414`:
	return intoBlock414
}

func TaintStepTest_EncodingPemDecode_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte518` into `intoByte650`.

	// Assume that `sourceCQL` has the underlying type of `fromByte518`:
	fromByte518 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte518` to result `intoByte650`
	// (`intoByte650` is now tainted).
	_, intoByte650 := pem.Decode(fromByte518)

	// Return the tainted `intoByte650`:
	return intoByte650
}

func TaintStepTest_EncodingPemEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBlock784` into `intoWriter957`.

	// Assume that `sourceCQL` has the underlying type of `fromBlock784`:
	fromBlock784 := sourceCQL.(*pem.Block)

	// Declare `intoWriter957` variable:
	var intoWriter957 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromBlock784` to parameter `intoWriter957`;
	// `intoWriter957` is now tainted.
	pem.Encode(intoWriter957, fromBlock784)

	// Return the tainted `intoWriter957`:
	return intoWriter957
}

func TaintStepTest_EncodingPemEncodeToMemory_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBlock520` into `intoByte443`.

	// Assume that `sourceCQL` has the underlying type of `fromBlock520`:
	fromBlock520 := sourceCQL.(*pem.Block)

	// Call the function that transfers the taint
	// from the parameter `fromBlock520` to result `intoByte443`
	// (`intoByte443` is now tainted).
	intoByte443 := pem.EncodeToMemory(fromBlock520)

	// Return the tainted `intoByte443`:
	return intoByte443
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
