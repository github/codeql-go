// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "encoding"

func TaintStepTest_EncodingBinaryMarshalerMarshalBinary_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBinaryMarshaler415` into `intoByte947`.

	// Assume that `sourceCQL` has the underlying type of `fromBinaryMarshaler415`:
	fromBinaryMarshaler415 := sourceCQL.(encoding.BinaryMarshaler)

	// Call the method that transfers the taint
	// from the receiver `fromBinaryMarshaler415` to the result `intoByte947`
	// (`intoByte947` is now tainted).
	intoByte947, _ := fromBinaryMarshaler415.MarshalBinary()

	// Return the tainted `intoByte947`:
	return intoByte947
}

func TaintStepTest_EncodingTextMarshalerMarshalText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTextMarshaler711` into `intoByte897`.

	// Assume that `sourceCQL` has the underlying type of `fromTextMarshaler711`:
	fromTextMarshaler711 := sourceCQL.(encoding.TextMarshaler)

	// Call the method that transfers the taint
	// from the receiver `fromTextMarshaler711` to the result `intoByte897`
	// (`intoByte897` is now tainted).
	intoByte897, _ := fromTextMarshaler711.MarshalText()

	// Return the tainted `intoByte897`:
	return intoByte897
}

func TaintStepTest_EncodingBinaryUnmarshalerUnmarshalBinary_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte595` into `intoBinaryUnmarshaler887`.

	// Assume that `sourceCQL` has the underlying type of `fromByte595`:
	fromByte595 := sourceCQL.([]byte)

	// Declare `intoBinaryUnmarshaler887` variable:
	var intoBinaryUnmarshaler887 encoding.BinaryUnmarshaler

	// Call the method that transfers the taint
	// from the parameter `fromByte595` to the receiver `intoBinaryUnmarshaler887`
	// (`intoBinaryUnmarshaler887` is now tainted).
	intoBinaryUnmarshaler887.UnmarshalBinary(fromByte595)

	// Return the tainted `intoBinaryUnmarshaler887`:
	return intoBinaryUnmarshaler887
}

func TaintStepTest_EncodingTextUnmarshalerUnmarshalText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte871` into `intoTextUnmarshaler567`.

	// Assume that `sourceCQL` has the underlying type of `fromByte871`:
	fromByte871 := sourceCQL.([]byte)

	// Declare `intoTextUnmarshaler567` variable:
	var intoTextUnmarshaler567 encoding.TextUnmarshaler

	// Call the method that transfers the taint
	// from the parameter `fromByte871` to the receiver `intoTextUnmarshaler567`
	// (`intoTextUnmarshaler567` is now tainted).
	intoTextUnmarshaler567.UnmarshalText(fromByte871)

	// Return the tainted `intoTextUnmarshaler567`:
	return intoTextUnmarshaler567
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
