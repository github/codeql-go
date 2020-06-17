// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "mime"

func TaintStepTest_MimeFormatMediaType_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString656` into `intoString414`.

	// Assume that `sourceCQL` has the underlying type of `fromString656`:
	fromString656 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString656` to result `intoString414`
	// (`intoString414` is now tainted).
	intoString414 := mime.FormatMediaType(fromString656, nil)

	// Return the tainted `intoString414`:
	return intoString414
}

func TaintStepTest_MimeFormatMediaType_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMapstringstring518` into `intoString650`.

	// Assume that `sourceCQL` has the underlying type of `fromMapstringstring518`:
	fromMapstringstring518 := sourceCQL.(map[string]string)

	// Call the function that transfers the taint
	// from the parameter `fromMapstringstring518` to result `intoString650`
	// (`intoString650` is now tainted).
	intoString650 := mime.FormatMediaType("", fromMapstringstring518)

	// Return the tainted `intoString650`:
	return intoString650
}

func TaintStepTest_MimeParseMediaType_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString784` into `intoString957`.

	// Assume that `sourceCQL` has the underlying type of `fromString784`:
	fromString784 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString784` to result `intoString957`
	// (`intoString957` is now tainted).
	intoString957, _, _ := mime.ParseMediaType(fromString784)

	// Return the tainted `intoString957`:
	return intoString957
}

func TaintStepTest_MimeParseMediaType_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString520` into `intoMapstringstring443`.

	// Assume that `sourceCQL` has the underlying type of `fromString520`:
	fromString520 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString520` to result `intoMapstringstring443`
	// (`intoMapstringstring443` is now tainted).
	_, intoMapstringstring443, _ := mime.ParseMediaType(fromString520)

	// Return the tainted `intoMapstringstring443`:
	return intoMapstringstring443
}

func TaintStepTest_MimeWordDecoderDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString127` into `intoString483`.

	// Assume that `sourceCQL` has the underlying type of `fromString127`:
	fromString127 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mime.WordDecoder

	// Call the method that transfers the taint
	// from the parameter `fromString127` to the result `intoString483`
	// (`intoString483` is now tainted).
	intoString483, _ := mediumObjCQL.Decode(fromString127)

	// Return the tainted `intoString483`:
	return intoString483
}

func TaintStepTest_MimeWordDecoderDecodeHeader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString989` into `intoString982`.

	// Assume that `sourceCQL` has the underlying type of `fromString989`:
	fromString989 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mime.WordDecoder

	// Call the method that transfers the taint
	// from the parameter `fromString989` to the result `intoString982`
	// (`intoString982` is now tainted).
	intoString982, _ := mediumObjCQL.DecodeHeader(fromString989)

	// Return the tainted `intoString982`:
	return intoString982
}

func TaintStepTest_MimeWordEncoderEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString417` into `intoString584`.

	// Assume that `sourceCQL` has the underlying type of `fromString417`:
	fromString417 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mime.WordEncoder

	// Call the method that transfers the taint
	// from the parameter `fromString417` to the result `intoString584`
	// (`intoString584` is now tainted).
	intoString584 := mediumObjCQL.Encode("", fromString417)

	// Return the tainted `intoString584`:
	return intoString584
}

func RunAllTaints_Mime() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeFormatMediaType_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeFormatMediaType_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeParseMediaType_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeParseMediaType_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeWordDecoderDecode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeWordDecoderDecodeHeader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeWordEncoderEncode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
