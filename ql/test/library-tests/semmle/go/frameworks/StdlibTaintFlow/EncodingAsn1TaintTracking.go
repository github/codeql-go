package main

import "encoding/asn1"

func TaintStepTest_EncodingAsn1Marshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface498` into `intoByte546`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface498`:
	fromInterface498 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface498` to result `intoByte546`
	// (`intoByte546` is now tainted).
	intoByte546, _ := asn1.Marshal(fromInterface498)

	// Return the tainted `intoByte546`:
	return intoByte546
}

func TaintStepTest_EncodingAsn1MarshalWithParams_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface274` into `intoByte503`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface274`:
	fromInterface274 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface274` to result `intoByte503`
	// (`intoByte503` is now tainted).
	intoByte503, _ := asn1.MarshalWithParams(fromInterface274, "")

	// Return the tainted `intoByte503`:
	return intoByte503
}

func TaintStepTest_EncodingAsn1MarshalWithParams_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString334` into `intoByte694`.

	// Assume that `sourceCQL` has the underlying type of `fromString334`:
	fromString334 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString334` to result `intoByte694`
	// (`intoByte694` is now tainted).
	intoByte694, _ := asn1.MarshalWithParams(nil, fromString334)

	// Return the tainted `intoByte694`:
	return intoByte694
}

func TaintStepTest_EncodingAsn1Unmarshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte470` into `intoInterface570`.

	// Assume that `sourceCQL` has the underlying type of `fromByte470`:
	fromByte470 := sourceCQL.([]byte)

	// Declare `intoInterface570` variable:
	var intoInterface570 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromByte470` to parameter `intoInterface570`;
	// `intoInterface570` is now tainted.
	asn1.Unmarshal(fromByte470, intoInterface570)

	// Return the tainted `intoInterface570`:
	return intoInterface570
}

func TaintStepTest_EncodingAsn1Unmarshal_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte657` into `intoByte694`.

	// Assume that `sourceCQL` has the underlying type of `fromByte657`:
	fromByte657 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte657` to result `intoByte694`
	// (`intoByte694` is now tainted).
	intoByte694, _ := asn1.Unmarshal(fromByte657, nil)

	// Return the tainted `intoByte694`:
	return intoByte694
}

func TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte716` into `intoInterface520`.

	// Assume that `sourceCQL` has the underlying type of `fromByte716`:
	fromByte716 := sourceCQL.([]byte)

	// Declare `intoInterface520` variable:
	var intoInterface520 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromByte716` to parameter `intoInterface520`;
	// `intoInterface520` is now tainted.
	asn1.UnmarshalWithParams(fromByte716, intoInterface520, "")

	// Return the tainted `intoInterface520`:
	return intoInterface520
}

func TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte616` into `intoByte489`.

	// Assume that `sourceCQL` has the underlying type of `fromByte616`:
	fromByte616 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte616` to result `intoByte489`
	// (`intoByte489` is now tainted).
	intoByte489, _ := asn1.UnmarshalWithParams(fromByte616, nil, "")

	// Return the tainted `intoByte489`:
	return intoByte489
}

func TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString517` into `intoInterface820`.

	// Assume that `sourceCQL` has the underlying type of `fromString517`:
	fromString517 := sourceCQL.(string)

	// Declare `intoInterface820` variable:
	var intoInterface820 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString517` to parameter `intoInterface820`;
	// `intoInterface820` is now tainted.
	asn1.UnmarshalWithParams(nil, intoInterface820, fromString517)

	// Return the tainted `intoInterface820`:
	return intoInterface820
}

func TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I1O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString891` into `intoByte514`.

	// Assume that `sourceCQL` has the underlying type of `fromString891`:
	fromString891 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString891` to result `intoByte514`
	// (`intoByte514` is now tainted).
	intoByte514, _ := asn1.UnmarshalWithParams(nil, nil, fromString891)

	// Return the tainted `intoByte514`:
	return intoByte514
}

func TaintStepTest_EncodingAsn1ObjectIdentifierString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromObjectIdentifier352` into `intoString813`.

	// Assume that `sourceCQL` has the underlying type of `fromObjectIdentifier352`:
	fromObjectIdentifier352 := sourceCQL.(asn1.ObjectIdentifier)

	// Call the method that transfers the taint
	// from the receiver `fromObjectIdentifier352` to the result `intoString813`
	// (`intoString813` is now tainted).
	intoString813 := fromObjectIdentifier352.String()

	// Return the tainted `intoString813`:
	return intoString813
}

func RunAllTaints_EncodingAsn1() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingAsn1Marshal_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingAsn1MarshalWithParams_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingAsn1MarshalWithParams_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingAsn1Unmarshal_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingAsn1Unmarshal_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I1O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingAsn1ObjectIdentifierString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
