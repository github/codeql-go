package main

import "encoding"

func TaintStepTest_EncodingBinaryMarshalerMarshalBinary_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBinaryMarshaler608` into `intoByte302`.

	// Assume that `sourceCQL` has the underlying type of `fromBinaryMarshaler608`:
	fromBinaryMarshaler608 := sourceCQL.(encoding.BinaryMarshaler)

	// Call the method that transfers the taint
	// from the receiver `fromBinaryMarshaler608` to the result `intoByte302`
	// (`intoByte302` is now tainted).
	intoByte302, _ := fromBinaryMarshaler608.MarshalBinary()

	// Return the tainted `intoByte302`:
	return intoByte302
}

func TaintStepTest_EncodingTextMarshalerMarshalText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTextMarshaler485` into `intoByte668`.

	// Assume that `sourceCQL` has the underlying type of `fromTextMarshaler485`:
	fromTextMarshaler485 := sourceCQL.(encoding.TextMarshaler)

	// Call the method that transfers the taint
	// from the receiver `fromTextMarshaler485` to the result `intoByte668`
	// (`intoByte668` is now tainted).
	intoByte668, _ := fromTextMarshaler485.MarshalText()

	// Return the tainted `intoByte668`:
	return intoByte668
}

func TaintStepTest_EncodingBinaryUnmarshalerUnmarshalBinary_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte524` into `intoBinaryUnmarshaler591`.

	// Assume that `sourceCQL` has the underlying type of `fromByte524`:
	fromByte524 := sourceCQL.([]byte)

	// Declare `intoBinaryUnmarshaler591` variable:
	var intoBinaryUnmarshaler591 encoding.BinaryUnmarshaler

	// Call the method that transfers the taint
	// from the parameter `fromByte524` to the receiver `intoBinaryUnmarshaler591`
	// (`intoBinaryUnmarshaler591` is now tainted).
	intoBinaryUnmarshaler591.UnmarshalBinary(fromByte524)

	// Return the tainted `intoBinaryUnmarshaler591`:
	return intoBinaryUnmarshaler591
}

func TaintStepTest_EncodingTextUnmarshalerUnmarshalText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte528` into `intoTextUnmarshaler811`.

	// Assume that `sourceCQL` has the underlying type of `fromByte528`:
	fromByte528 := sourceCQL.([]byte)

	// Declare `intoTextUnmarshaler811` variable:
	var intoTextUnmarshaler811 encoding.TextUnmarshaler

	// Call the method that transfers the taint
	// from the parameter `fromByte528` to the receiver `intoTextUnmarshaler811`
	// (`intoTextUnmarshaler811` is now tainted).
	intoTextUnmarshaler811.UnmarshalText(fromByte528)

	// Return the tainted `intoTextUnmarshaler811`:
	return intoTextUnmarshaler811
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
