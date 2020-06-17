package main

import (
	"encoding/pem"
	"io"
)

func TaintStepTest_EncodingPemDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte141` into `intoBlock691`.

	// Assume that `sourceCQL` has the underlying type of `fromByte141`:
	fromByte141 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte141` to result `intoBlock691`
	// (`intoBlock691` is now tainted).
	intoBlock691, _ := pem.Decode(fromByte141)

	// Return the tainted `intoBlock691`:
	return intoBlock691
}

func TaintStepTest_EncodingPemDecode_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte985` into `intoByte555`.

	// Assume that `sourceCQL` has the underlying type of `fromByte985`:
	fromByte985 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte985` to result `intoByte555`
	// (`intoByte555` is now tainted).
	_, intoByte555 := pem.Decode(fromByte985)

	// Return the tainted `intoByte555`:
	return intoByte555
}

func TaintStepTest_EncodingPemEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBlock613` into `intoWriter517`.

	// Assume that `sourceCQL` has the underlying type of `fromBlock613`:
	fromBlock613 := sourceCQL.(*pem.Block)

	// Declare `intoWriter517` variable:
	var intoWriter517 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromBlock613` to parameter `intoWriter517`;
	// `intoWriter517` is now tainted.
	pem.Encode(intoWriter517, fromBlock613)

	// Return the tainted `intoWriter517`:
	return intoWriter517
}

func TaintStepTest_EncodingPemEncodeToMemory_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBlock623` into `intoByte871`.

	// Assume that `sourceCQL` has the underlying type of `fromBlock623`:
	fromBlock623 := sourceCQL.(*pem.Block)

	// Call the function that transfers the taint
	// from the parameter `fromBlock623` to result `intoByte871`
	// (`intoByte871` is now tainted).
	intoByte871 := pem.EncodeToMemory(fromBlock623)

	// Return the tainted `intoByte871`:
	return intoByte871
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
