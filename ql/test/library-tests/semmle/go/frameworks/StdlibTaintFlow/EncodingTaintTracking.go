// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "encoding"

func TaintStepTest_EncodingBinaryMarshalerMarshalBinary_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBinaryMarshaler656` into `intoByte414`.

	// Assume that `sourceCQL` has the underlying type of `fromBinaryMarshaler656`:
	fromBinaryMarshaler656 := sourceCQL.(encoding.BinaryMarshaler)

	// Call the method that transfers the taint
	// from the receiver `fromBinaryMarshaler656` to the result `intoByte414`
	// (`intoByte414` is now tainted).
	intoByte414, _ := fromBinaryMarshaler656.MarshalBinary()

	// Return the tainted `intoByte414`:
	return intoByte414
}

func TaintStepTest_EncodingTextMarshalerMarshalText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTextMarshaler518` into `intoByte650`.

	// Assume that `sourceCQL` has the underlying type of `fromTextMarshaler518`:
	fromTextMarshaler518 := sourceCQL.(encoding.TextMarshaler)

	// Call the method that transfers the taint
	// from the receiver `fromTextMarshaler518` to the result `intoByte650`
	// (`intoByte650` is now tainted).
	intoByte650, _ := fromTextMarshaler518.MarshalText()

	// Return the tainted `intoByte650`:
	return intoByte650
}

func TaintStepTest_EncodingBinaryUnmarshalerUnmarshalBinary_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte784` into `intoBinaryUnmarshaler957`.

	// Assume that `sourceCQL` has the underlying type of `fromByte784`:
	fromByte784 := sourceCQL.([]byte)

	// Declare `intoBinaryUnmarshaler957` variable:
	var intoBinaryUnmarshaler957 encoding.BinaryUnmarshaler

	// Call the method that transfers the taint
	// from the parameter `fromByte784` to the receiver `intoBinaryUnmarshaler957`
	// (`intoBinaryUnmarshaler957` is now tainted).
	intoBinaryUnmarshaler957.UnmarshalBinary(fromByte784)

	// Return the tainted `intoBinaryUnmarshaler957`:
	return intoBinaryUnmarshaler957
}

func TaintStepTest_EncodingTextUnmarshalerUnmarshalText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte520` into `intoTextUnmarshaler443`.

	// Assume that `sourceCQL` has the underlying type of `fromByte520`:
	fromByte520 := sourceCQL.([]byte)

	// Declare `intoTextUnmarshaler443` variable:
	var intoTextUnmarshaler443 encoding.TextUnmarshaler

	// Call the method that transfers the taint
	// from the parameter `fromByte520` to the receiver `intoTextUnmarshaler443`
	// (`intoTextUnmarshaler443` is now tainted).
	intoTextUnmarshaler443.UnmarshalText(fromByte520)

	// Return the tainted `intoTextUnmarshaler443`:
	return intoTextUnmarshaler443
}

func RunAllTaints_Encoding() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBinaryMarshalerMarshalBinary_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingTextMarshalerMarshalText_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBinaryUnmarshalerUnmarshalBinary_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingTextUnmarshalerUnmarshalText_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
