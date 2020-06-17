// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"encoding/hex"
	"io"
)

func TaintStepTest_EncodingHexDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte325` into `intoByte463`.

	// Assume that `sourceCQL` has the underlying type of `fromByte325`:
	fromByte325 := sourceCQL.([]byte)

	// Declare `intoByte463` variable:
	var intoByte463 []byte

	// Call the function that transfers the taint
	// from the parameter `fromByte325` to parameter `intoByte463`;
	// `intoByte463` is now tainted.
	hex.Decode(intoByte463, fromByte325)

	// Return the tainted `intoByte463`:
	return intoByte463
}

func TaintStepTest_EncodingHexDecodeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString140` into `intoByte410`.

	// Assume that `sourceCQL` has the underlying type of `fromString140`:
	fromString140 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString140` to result `intoByte410`
	// (`intoByte410` is now tainted).
	intoByte410, _ := hex.DecodeString(fromString140)

	// Return the tainted `intoByte410`:
	return intoByte410
}

func TaintStepTest_EncodingHexDump_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte663` into `intoString665`.

	// Assume that `sourceCQL` has the underlying type of `fromByte663`:
	fromByte663 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte663` to result `intoString665`
	// (`intoString665` is now tainted).
	intoString665 := hex.Dump(fromByte663)

	// Return the tainted `intoString665`:
	return intoString665
}

func TaintStepTest_EncodingHexDumper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser893` into `intoWriter940`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser893`:
	fromWriteCloser893 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter940` variable:
	var intoWriter940 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter940`:
	intermediateCQL := hex.Dumper(intoWriter940)

	// Extra step (`fromWriteCloser893` taints `intermediateCQL`, which taints `intoWriter940`:
	link(fromWriteCloser893, intermediateCQL)

	// Return the tainted `intoWriter940`:
	return intoWriter940
}

func TaintStepTest_EncodingHexEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte878` into `intoByte906`.

	// Assume that `sourceCQL` has the underlying type of `fromByte878`:
	fromByte878 := sourceCQL.([]byte)

	// Declare `intoByte906` variable:
	var intoByte906 []byte

	// Call the function that transfers the taint
	// from the parameter `fromByte878` to parameter `intoByte906`;
	// `intoByte906` is now tainted.
	hex.Encode(intoByte906, fromByte878)

	// Return the tainted `intoByte906`:
	return intoByte906
}

func TaintStepTest_EncodingHexEncodeToString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte613` into `intoString878`.

	// Assume that `sourceCQL` has the underlying type of `fromByte613`:
	fromByte613 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte613` to result `intoString878`
	// (`intoString878` is now tainted).
	intoString878 := hex.EncodeToString(fromByte613)

	// Return the tainted `intoString878`:
	return intoString878
}

func TaintStepTest_EncodingHexNewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader701` into `intoReader830`.

	// Assume that `sourceCQL` has the underlying type of `fromReader701`:
	fromReader701 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader701` to result `intoReader830`
	// (`intoReader830` is now tainted).
	intoReader830 := hex.NewDecoder(fromReader701)

	// Return the tainted `intoReader830`:
	return intoReader830
}

func TaintStepTest_EncodingHexNewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter905` into `intoWriter686`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter905`:
	fromWriter905 := sourceCQL.(io.Writer)

	// Declare `intoWriter686` variable:
	var intoWriter686 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter686`:
	intermediateCQL := hex.NewEncoder(intoWriter686)

	// Extra step (`fromWriter905` taints `intermediateCQL`, which taints `intoWriter686`:
	link(fromWriter905, intermediateCQL)

	// Return the tainted `intoWriter686`:
	return intoWriter686
}

func RunAllTaints_EncodingHex() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingHexDecode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingHexDecodeString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingHexDump_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingHexDumper_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingHexEncode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingHexEncodeToString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingHexNewDecoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingHexNewEncoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
