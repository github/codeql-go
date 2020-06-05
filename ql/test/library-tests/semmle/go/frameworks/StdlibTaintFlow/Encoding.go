package main

import "encoding"

func TaintStepTest_EncodingBinaryMarshalerMarshalBinary(sourceCQL interface{}) {
	// The flow is from `from483` into `data`.

	// Assume that `sourceCQL` has the underlying type of `from483`:
	from483 := sourceCQL.(encoding.BinaryMarshaler)

	// Call the method that transfers the taint
	// from the receiver `from483` to the result `data`
	// (`data` is now tainted).
	data, _ := from483.MarshalBinary()

	// Sink the tainted `data`:
	sink(data)
}

func TaintStepTest_EncodingBinaryUnmarshalerUnmarshalBinary(sourceCQL interface{}) {
	// The flow is from `data` into `into208`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.([]byte)

	// Declare `into208` variable:
	var into208 encoding.BinaryUnmarshaler

	// Call the method that transfers the taint
	// from the parameter `data` to the receiver `into208`
	// (`into208` is now tainted).
	into208.UnmarshalBinary(data)

	// Sink the tainted `into208`:
	sink(into208)
}

func TaintStepTest_EncodingTextMarshalerMarshalText(sourceCQL interface{}) {
	// The flow is from `from685` into `text`.

	// Assume that `sourceCQL` has the underlying type of `from685`:
	from685 := sourceCQL.(encoding.TextMarshaler)

	// Call the method that transfers the taint
	// from the receiver `from685` to the result `text`
	// (`text` is now tainted).
	text, _ := from685.MarshalText()

	// Sink the tainted `text`:
	sink(text)
}

func TaintStepTest_EncodingTextUnmarshalerUnmarshalText(sourceCQL interface{}) {
	// The flow is from `text` into `into548`.

	// Assume that `sourceCQL` has the underlying type of `text`:
	text := sourceCQL.([]byte)

	// Declare `into548` variable:
	var into548 encoding.TextUnmarshaler

	// Call the method that transfers the taint
	// from the parameter `text` to the receiver `into548`
	// (`into548` is now tainted).
	into548.UnmarshalText(text)

	// Sink the tainted `into548`:
	sink(into548)
}

func RunAllTaints_Encoding(v interface{}) {
	{
		source := newSource()
		TaintStepTest_EncodingBinaryMarshalerMarshalBinary(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryUnmarshalerUnmarshalBinary(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingTextMarshalerMarshalText(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingTextUnmarshalerUnmarshalText(source)
	}
}
