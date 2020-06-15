package main

import "encoding"

func TaintStepTest_EncodingBinaryMarshalerMarshalBinary_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBinaryMarshaler889` into `intoByte136`.

	// Assume that `sourceCQL` has the underlying type of `fromBinaryMarshaler889`:
	fromBinaryMarshaler889 := sourceCQL.(encoding.BinaryMarshaler)

	// Call the method that transfers the taint
	// from the receiver `fromBinaryMarshaler889` to the result `intoByte136`
	// (`intoByte136` is now tainted).
	intoByte136, _ := fromBinaryMarshaler889.MarshalBinary()

	// Sink the tainted `intoByte136`:
	sink(intoByte136)
}

func TaintStepTest_EncodingTextMarshalerMarshalText_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTextMarshaler402` into `intoByte925`.

	// Assume that `sourceCQL` has the underlying type of `fromTextMarshaler402`:
	fromTextMarshaler402 := sourceCQL.(encoding.TextMarshaler)

	// Call the method that transfers the taint
	// from the receiver `fromTextMarshaler402` to the result `intoByte925`
	// (`intoByte925` is now tainted).
	intoByte925, _ := fromTextMarshaler402.MarshalText()

	// Sink the tainted `intoByte925`:
	sink(intoByte925)
}

func TaintStepTest_EncodingBinaryUnmarshalerUnmarshalBinary_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte534` into `intoBinaryUnmarshaler729`.

	// Assume that `sourceCQL` has the underlying type of `fromByte534`:
	fromByte534 := sourceCQL.([]byte)

	// Declare `intoBinaryUnmarshaler729` variable:
	var intoBinaryUnmarshaler729 encoding.BinaryUnmarshaler

	// Call the method that transfers the taint
	// from the parameter `fromByte534` to the receiver `intoBinaryUnmarshaler729`
	// (`intoBinaryUnmarshaler729` is now tainted).
	intoBinaryUnmarshaler729.UnmarshalBinary(fromByte534)

	// Sink the tainted `intoBinaryUnmarshaler729`:
	sink(intoBinaryUnmarshaler729)
}

func TaintStepTest_EncodingTextUnmarshalerUnmarshalText_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte845` into `intoTextUnmarshaler124`.

	// Assume that `sourceCQL` has the underlying type of `fromByte845`:
	fromByte845 := sourceCQL.([]byte)

	// Declare `intoTextUnmarshaler124` variable:
	var intoTextUnmarshaler124 encoding.TextUnmarshaler

	// Call the method that transfers the taint
	// from the parameter `fromByte845` to the receiver `intoTextUnmarshaler124`
	// (`intoTextUnmarshaler124` is now tainted).
	intoTextUnmarshaler124.UnmarshalText(fromByte845)

	// Sink the tainted `intoTextUnmarshaler124`:
	sink(intoTextUnmarshaler124)
}

func RunAllTaints_Encoding() {
	{
		source := newSource()
		TaintStepTest_EncodingBinaryMarshalerMarshalBinary_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingTextMarshalerMarshalText_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryUnmarshalerUnmarshalBinary_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingTextUnmarshalerUnmarshalText_B0I0O0(source)
	}
}
