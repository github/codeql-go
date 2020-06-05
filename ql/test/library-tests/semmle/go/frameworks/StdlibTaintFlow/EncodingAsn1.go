package main

import "encoding/asn1"

func TaintStepTest_EncodingAsn1Marshal(sourceCQL interface{}) {
	// The flow is from `val` into `into409`.

	// Assume that `sourceCQL` has the underlying type of `val`:
	val := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `val` to result `into409`
	// (`into409` is now tainted).
	into409, _ := asn1.Marshal(val)

	// Sink the tainted `into409`:
	sink(into409)
}

func TaintStepTest_EncodingAsn1MarshalWithParams(sourceCQL interface{}) {
	// The flow is from `val` into `into980`.

	// Assume that `sourceCQL` has the underlying type of `val`:
	val := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `val` to result `into980`
	// (`into980` is now tainted).
	into980, _ := asn1.MarshalWithParams(val, "")

	// Sink the tainted `into980`:
	sink(into980)
}

func TaintStepTest_EncodingAsn1Unmarshal(sourceCQL interface{}) {
	// The flow is from `b` into `val`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare `val` variable:
	var val interface{}

	// Call the function that transfers the taint
	// from the parameter `b` to parameter `val`;
	// `val` is now tainted.
	asn1.Unmarshal(b, val)

	// Sink the tainted `val`:
	sink(val)
}

func TaintStepTest_EncodingAsn1UnmarshalWithParams(sourceCQL interface{}) {
	// The flow is from `b` into `val`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare `val` variable:
	var val interface{}

	// Call the function that transfers the taint
	// from the parameter `b` to parameter `val`;
	// `val` is now tainted.
	asn1.UnmarshalWithParams(b, val, "")

	// Sink the tainted `val`:
	sink(val)
}

func TaintStepTest_EncodingAsn1ObjectIdentifierString(sourceCQL interface{}) {
	// The flow is from `from731` into `into594`.

	// Assume that `sourceCQL` has the underlying type of `from731`:
	from731 := sourceCQL.(asn1.ObjectIdentifier)

	// Call the method that transfers the taint
	// from the receiver `from731` to the result `into594`
	// (`into594` is now tainted).
	into594 := from731.String()

	// Sink the tainted `into594`:
	sink(into594)
}

func RunAllTaints_EncodingAsn1(v interface{}) {
	{
		source := newSource()
		TaintStepTest_EncodingAsn1Marshal(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAsn1MarshalWithParams(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAsn1Unmarshal(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAsn1UnmarshalWithParams(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAsn1ObjectIdentifierString(source)
	}
}
