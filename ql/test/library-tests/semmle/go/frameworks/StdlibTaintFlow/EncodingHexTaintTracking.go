package main

import (
	"encoding/hex"
	"io"
)

func TaintStepTest_EncodingHexDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte708` into `intoByte185`.

	// Assume that `sourceCQL` has the underlying type of `fromByte708`:
	fromByte708 := sourceCQL.([]byte)

	// Declare `intoByte185` variable:
	var intoByte185 []byte

	// Call the function that transfers the taint
	// from the parameter `fromByte708` to parameter `intoByte185`;
	// `intoByte185` is now tainted.
	hex.Decode(intoByte185, fromByte708)

	// Return the tainted `intoByte185`:
	return intoByte185
}

func TaintStepTest_EncodingHexDecodeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString124` into `intoByte637`.

	// Assume that `sourceCQL` has the underlying type of `fromString124`:
	fromString124 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString124` to result `intoByte637`
	// (`intoByte637` is now tainted).
	intoByte637, _ := hex.DecodeString(fromString124)

	// Return the tainted `intoByte637`:
	return intoByte637
}

func TaintStepTest_EncodingHexDump_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte935` into `intoString520`.

	// Assume that `sourceCQL` has the underlying type of `fromByte935`:
	fromByte935 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte935` to result `intoString520`
	// (`intoString520` is now tainted).
	intoString520 := hex.Dump(fromByte935)

	// Return the tainted `intoString520`:
	return intoString520
}

func TaintStepTest_EncodingHexDumper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser700` into `intoWriter486`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser700`:
	fromWriteCloser700 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter486` variable:
	var intoWriter486 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter486`:
	intermediateCQL := hex.Dumper(intoWriter486)

	// Extra step (`fromWriteCloser700` taints `intermediateCQL`, which taints `intoWriter486`:
	link(fromWriteCloser700, intermediateCQL)

	// Return the tainted `intoWriter486`:
	return intoWriter486
}

func TaintStepTest_EncodingHexEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte835` into `intoByte208`.

	// Assume that `sourceCQL` has the underlying type of `fromByte835`:
	fromByte835 := sourceCQL.([]byte)

	// Declare `intoByte208` variable:
	var intoByte208 []byte

	// Call the function that transfers the taint
	// from the parameter `fromByte835` to parameter `intoByte208`;
	// `intoByte208` is now tainted.
	hex.Encode(intoByte208, fromByte835)

	// Return the tainted `intoByte208`:
	return intoByte208
}

func TaintStepTest_EncodingHexEncodeToString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte375` into `intoString166`.

	// Assume that `sourceCQL` has the underlying type of `fromByte375`:
	fromByte375 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte375` to result `intoString166`
	// (`intoString166` is now tainted).
	intoString166 := hex.EncodeToString(fromByte375)

	// Return the tainted `intoString166`:
	return intoString166
}

func TaintStepTest_EncodingHexNewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader916` into `intoReader561`.

	// Assume that `sourceCQL` has the underlying type of `fromReader916`:
	fromReader916 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader916` to result `intoReader561`
	// (`intoReader561` is now tainted).
	intoReader561 := hex.NewDecoder(fromReader916)

	// Return the tainted `intoReader561`:
	return intoReader561
}

func TaintStepTest_EncodingHexNewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter667` into `intoWriter458`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter667`:
	fromWriter667 := sourceCQL.(io.Writer)

	// Declare `intoWriter458` variable:
	var intoWriter458 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter458`:
	intermediateCQL := hex.NewEncoder(intoWriter458)

	// Extra step (`fromWriter667` taints `intermediateCQL`, which taints `intoWriter458`:
	link(fromWriter667, intermediateCQL)

	// Return the tainted `intoWriter458`:
	return intoWriter458
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
