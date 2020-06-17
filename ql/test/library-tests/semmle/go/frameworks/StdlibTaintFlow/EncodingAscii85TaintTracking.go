package main

import (
	"encoding/ascii85"
	"io"
)

func TaintStepTest_EncodingAscii85Decode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte604` into `intoByte333`.

	// Assume that `sourceCQL` has the underlying type of `fromByte604`:
	fromByte604 := sourceCQL.([]byte)

	// Declare `intoByte333` variable:
	var intoByte333 []byte

	// Call the function that transfers the taint
	// from the parameter `fromByte604` to parameter `intoByte333`;
	// `intoByte333` is now tainted.
	ascii85.Decode(intoByte333, fromByte604, false)

	// Return the tainted `intoByte333`:
	return intoByte333
}

func TaintStepTest_EncodingAscii85Encode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte216` into `intoByte193`.

	// Assume that `sourceCQL` has the underlying type of `fromByte216`:
	fromByte216 := sourceCQL.([]byte)

	// Declare `intoByte193` variable:
	var intoByte193 []byte

	// Call the function that transfers the taint
	// from the parameter `fromByte216` to parameter `intoByte193`;
	// `intoByte193` is now tainted.
	ascii85.Encode(intoByte193, fromByte216)

	// Return the tainted `intoByte193`:
	return intoByte193
}

func TaintStepTest_EncodingAscii85NewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader299` into `intoReader157`.

	// Assume that `sourceCQL` has the underlying type of `fromReader299`:
	fromReader299 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader299` to result `intoReader157`
	// (`intoReader157` is now tainted).
	intoReader157 := ascii85.NewDecoder(fromReader299)

	// Return the tainted `intoReader157`:
	return intoReader157
}

func TaintStepTest_EncodingAscii85NewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser875` into `intoWriter894`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser875`:
	fromWriteCloser875 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter894` variable:
	var intoWriter894 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter894`:
	intermediateCQL := ascii85.NewEncoder(intoWriter894)

	// Extra step (`fromWriteCloser875` taints `intermediateCQL`, which taints `intoWriter894`:
	link(fromWriteCloser875, intermediateCQL)

	// Return the tainted `intoWriter894`:
	return intoWriter894
}

func RunAllTaints_EncodingAscii85() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingAscii85Decode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingAscii85Encode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingAscii85NewDecoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingAscii85NewEncoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
