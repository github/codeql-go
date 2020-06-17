// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"encoding/ascii85"
	"io"
)

func TaintStepTest_EncodingAscii85Decode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte269` into `intoByte120`.

	// Assume that `sourceCQL` has the underlying type of `fromByte269`:
	fromByte269 := sourceCQL.([]byte)

	// Declare `intoByte120` variable:
	var intoByte120 []byte

	// Call the function that transfers the taint
	// from the parameter `fromByte269` to parameter `intoByte120`;
	// `intoByte120` is now tainted.
	ascii85.Decode(intoByte120, fromByte269, false)

	// Return the tainted `intoByte120`:
	return intoByte120
}

func TaintStepTest_EncodingAscii85Encode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte846` into `intoByte322`.

	// Assume that `sourceCQL` has the underlying type of `fromByte846`:
	fromByte846 := sourceCQL.([]byte)

	// Declare `intoByte322` variable:
	var intoByte322 []byte

	// Call the function that transfers the taint
	// from the parameter `fromByte846` to parameter `intoByte322`;
	// `intoByte322` is now tainted.
	ascii85.Encode(intoByte322, fromByte846)

	// Return the tainted `intoByte322`:
	return intoByte322
}

func TaintStepTest_EncodingAscii85NewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader241` into `intoReader836`.

	// Assume that `sourceCQL` has the underlying type of `fromReader241`:
	fromReader241 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader241` to result `intoReader836`
	// (`intoReader836` is now tainted).
	intoReader836 := ascii85.NewDecoder(fromReader241)

	// Return the tainted `intoReader836`:
	return intoReader836
}

func TaintStepTest_EncodingAscii85NewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser111` into `intoWriter171`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser111`:
	fromWriteCloser111 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter171` variable:
	var intoWriter171 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter171`:
	intermediateCQL := ascii85.NewEncoder(intoWriter171)

	// Extra step (`fromWriteCloser111` taints `intermediateCQL`, which taints `intoWriter171`:
	link(fromWriteCloser111, intermediateCQL)

	// Return the tainted `intoWriter171`:
	return intoWriter171
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
