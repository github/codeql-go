// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"encoding/ascii85"
	"io"
)

func TaintStepTest_EncodingAscii85Decode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte656` into `intoByte414`.

	// Assume that `sourceCQL` has the underlying type of `fromByte656`:
	fromByte656 := sourceCQL.([]byte)

	// Declare `intoByte414` variable:
	var intoByte414 []byte

	// Call the function that transfers the taint
	// from the parameter `fromByte656` to parameter `intoByte414`;
	// `intoByte414` is now tainted.
	ascii85.Decode(intoByte414, fromByte656, false)

	// Return the tainted `intoByte414`:
	return intoByte414
}

func TaintStepTest_EncodingAscii85Encode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte518` into `intoByte650`.

	// Assume that `sourceCQL` has the underlying type of `fromByte518`:
	fromByte518 := sourceCQL.([]byte)

	// Declare `intoByte650` variable:
	var intoByte650 []byte

	// Call the function that transfers the taint
	// from the parameter `fromByte518` to parameter `intoByte650`;
	// `intoByte650` is now tainted.
	ascii85.Encode(intoByte650, fromByte518)

	// Return the tainted `intoByte650`:
	return intoByte650
}

func TaintStepTest_EncodingAscii85NewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader784` into `intoReader957`.

	// Assume that `sourceCQL` has the underlying type of `fromReader784`:
	fromReader784 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader784` to result `intoReader957`
	// (`intoReader957` is now tainted).
	intoReader957 := ascii85.NewDecoder(fromReader784)

	// Return the tainted `intoReader957`:
	return intoReader957
}

func TaintStepTest_EncodingAscii85NewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser520` into `intoWriter443`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser520`:
	fromWriteCloser520 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter443` variable:
	var intoWriter443 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter443`:
	intermediateCQL := ascii85.NewEncoder(intoWriter443)

	// Extra step (`fromWriteCloser520` taints `intermediateCQL`, which taints `intoWriter443`:
	link(fromWriteCloser520, intermediateCQL)

	// Return the tainted `intoWriter443`:
	return intoWriter443
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
