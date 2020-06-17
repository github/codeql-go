// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "encoding/asn1"

func TaintStepTest_EncodingAsn1Marshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface253` into `intoByte218`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface253`:
	fromInterface253 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface253` to result `intoByte218`
	// (`intoByte218` is now tainted).
	intoByte218, _ := asn1.Marshal(fromInterface253)

	// Return the tainted `intoByte218`:
	return intoByte218
}

func TaintStepTest_EncodingAsn1MarshalWithParams_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface633` into `intoByte976`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface633`:
	fromInterface633 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface633` to result `intoByte976`
	// (`intoByte976` is now tainted).
	intoByte976, _ := asn1.MarshalWithParams(fromInterface633, "")

	// Return the tainted `intoByte976`:
	return intoByte976
}

func TaintStepTest_EncodingAsn1MarshalWithParams_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString452` into `intoByte221`.

	// Assume that `sourceCQL` has the underlying type of `fromString452`:
	fromString452 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString452` to result `intoByte221`
	// (`intoByte221` is now tainted).
	intoByte221, _ := asn1.MarshalWithParams(nil, fromString452)

	// Return the tainted `intoByte221`:
	return intoByte221
}

func TaintStepTest_EncodingAsn1Unmarshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte975` into `intoInterface376`.

	// Assume that `sourceCQL` has the underlying type of `fromByte975`:
	fromByte975 := sourceCQL.([]byte)

	// Declare `intoInterface376` variable:
	var intoInterface376 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromByte975` to parameter `intoInterface376`;
	// `intoInterface376` is now tainted.
	asn1.Unmarshal(fromByte975, intoInterface376)

	// Return the tainted `intoInterface376`:
	return intoInterface376
}

func TaintStepTest_EncodingAsn1Unmarshal_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte901` into `intoByte391`.

	// Assume that `sourceCQL` has the underlying type of `fromByte901`:
	fromByte901 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte901` to result `intoByte391`
	// (`intoByte391` is now tainted).
	intoByte391, _ := asn1.Unmarshal(fromByte901, nil)

	// Return the tainted `intoByte391`:
	return intoByte391
}

func TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte724` into `intoInterface572`.

	// Assume that `sourceCQL` has the underlying type of `fromByte724`:
	fromByte724 := sourceCQL.([]byte)

	// Declare `intoInterface572` variable:
	var intoInterface572 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromByte724` to parameter `intoInterface572`;
	// `intoInterface572` is now tainted.
	asn1.UnmarshalWithParams(fromByte724, intoInterface572, "")

	// Return the tainted `intoInterface572`:
	return intoInterface572
}

func TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte800` into `intoByte637`.

	// Assume that `sourceCQL` has the underlying type of `fromByte800`:
	fromByte800 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte800` to result `intoByte637`
	// (`intoByte637` is now tainted).
	intoByte637, _ := asn1.UnmarshalWithParams(fromByte800, nil, "")

	// Return the tainted `intoByte637`:
	return intoByte637
}

func TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString116` into `intoInterface640`.

	// Assume that `sourceCQL` has the underlying type of `fromString116`:
	fromString116 := sourceCQL.(string)

	// Declare `intoInterface640` variable:
	var intoInterface640 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString116` to parameter `intoInterface640`;
	// `intoInterface640` is now tainted.
	asn1.UnmarshalWithParams(nil, intoInterface640, fromString116)

	// Return the tainted `intoInterface640`:
	return intoInterface640
}

func TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I1O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString303` into `intoByte784`.

	// Assume that `sourceCQL` has the underlying type of `fromString303`:
	fromString303 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString303` to result `intoByte784`
	// (`intoByte784` is now tainted).
	intoByte784, _ := asn1.UnmarshalWithParams(nil, nil, fromString303)

	// Return the tainted `intoByte784`:
	return intoByte784
}

func TaintStepTest_EncodingAsn1ObjectIdentifierString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromObjectIdentifier271` into `intoString833`.

	// Assume that `sourceCQL` has the underlying type of `fromObjectIdentifier271`:
	fromObjectIdentifier271 := sourceCQL.(asn1.ObjectIdentifier)

	// Call the method that transfers the taint
	// from the receiver `fromObjectIdentifier271` to the result `intoString833`
	// (`intoString833` is now tainted).
	intoString833 := fromObjectIdentifier271.String()

	// Return the tainted `intoString833`:
	return intoString833
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
