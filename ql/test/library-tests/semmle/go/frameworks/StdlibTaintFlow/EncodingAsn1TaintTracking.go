// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "encoding/asn1"

func TaintStepTest_EncodingAsn1Marshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface656` into `intoByte414`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface656`:
	fromInterface656 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface656` to result `intoByte414`
	// (`intoByte414` is now tainted).
	intoByte414, _ := asn1.Marshal(fromInterface656)

	// Return the tainted `intoByte414`:
	return intoByte414
}

func TaintStepTest_EncodingAsn1MarshalWithParams_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface518` into `intoByte650`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface518`:
	fromInterface518 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface518` to result `intoByte650`
	// (`intoByte650` is now tainted).
	intoByte650, _ := asn1.MarshalWithParams(fromInterface518, "")

	// Return the tainted `intoByte650`:
	return intoByte650
}

func TaintStepTest_EncodingAsn1MarshalWithParams_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString784` into `intoByte957`.

	// Assume that `sourceCQL` has the underlying type of `fromString784`:
	fromString784 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString784` to result `intoByte957`
	// (`intoByte957` is now tainted).
	intoByte957, _ := asn1.MarshalWithParams(nil, fromString784)

	// Return the tainted `intoByte957`:
	return intoByte957
}

func TaintStepTest_EncodingAsn1Unmarshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte520` into `intoInterface443`.

	// Assume that `sourceCQL` has the underlying type of `fromByte520`:
	fromByte520 := sourceCQL.([]byte)

	// Declare `intoInterface443` variable:
	var intoInterface443 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromByte520` to parameter `intoInterface443`;
	// `intoInterface443` is now tainted.
	asn1.Unmarshal(fromByte520, intoInterface443)

	// Return the tainted `intoInterface443`:
	return intoInterface443
}

func TaintStepTest_EncodingAsn1Unmarshal_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte127` into `intoByte483`.

	// Assume that `sourceCQL` has the underlying type of `fromByte127`:
	fromByte127 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte127` to result `intoByte483`
	// (`intoByte483` is now tainted).
	intoByte483, _ := asn1.Unmarshal(fromByte127, nil)

	// Return the tainted `intoByte483`:
	return intoByte483
}

func TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte989` into `intoInterface982`.

	// Assume that `sourceCQL` has the underlying type of `fromByte989`:
	fromByte989 := sourceCQL.([]byte)

	// Declare `intoInterface982` variable:
	var intoInterface982 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromByte989` to parameter `intoInterface982`;
	// `intoInterface982` is now tainted.
	asn1.UnmarshalWithParams(fromByte989, intoInterface982, "")

	// Return the tainted `intoInterface982`:
	return intoInterface982
}

func TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte417` into `intoByte584`.

	// Assume that `sourceCQL` has the underlying type of `fromByte417`:
	fromByte417 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte417` to result `intoByte584`
	// (`intoByte584` is now tainted).
	intoByte584, _ := asn1.UnmarshalWithParams(fromByte417, nil, "")

	// Return the tainted `intoByte584`:
	return intoByte584
}

func TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString991` into `intoInterface881`.

	// Assume that `sourceCQL` has the underlying type of `fromString991`:
	fromString991 := sourceCQL.(string)

	// Declare `intoInterface881` variable:
	var intoInterface881 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString991` to parameter `intoInterface881`;
	// `intoInterface881` is now tainted.
	asn1.UnmarshalWithParams(nil, intoInterface881, fromString991)

	// Return the tainted `intoInterface881`:
	return intoInterface881
}

func TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I1O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString186` into `intoByte284`.

	// Assume that `sourceCQL` has the underlying type of `fromString186`:
	fromString186 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString186` to result `intoByte284`
	// (`intoByte284` is now tainted).
	intoByte284, _ := asn1.UnmarshalWithParams(nil, nil, fromString186)

	// Return the tainted `intoByte284`:
	return intoByte284
}

func TaintStepTest_EncodingAsn1ObjectIdentifierString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromObjectIdentifier908` into `intoString137`.

	// Assume that `sourceCQL` has the underlying type of `fromObjectIdentifier908`:
	fromObjectIdentifier908 := sourceCQL.(asn1.ObjectIdentifier)

	// Call the method that transfers the taint
	// from the receiver `fromObjectIdentifier908` to the result `intoString137`
	// (`intoString137` is now tainted).
	intoString137 := fromObjectIdentifier908.String()

	// Return the tainted `intoString137`:
	return intoString137
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
