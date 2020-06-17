// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "mime"

func TaintStepTest_MimeFormatMediaType_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString589` into `intoString758`.

	// Assume that `sourceCQL` has the underlying type of `fromString589`:
	fromString589 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString589` to result `intoString758`
	// (`intoString758` is now tainted).
	intoString758 := mime.FormatMediaType(fromString589, nil)

	// Return the tainted `intoString758`:
	return intoString758
}

func TaintStepTest_MimeFormatMediaType_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMapstringstring160` into `intoString120`.

	// Assume that `sourceCQL` has the underlying type of `fromMapstringstring160`:
	fromMapstringstring160 := sourceCQL.(map[string]string)

	// Call the function that transfers the taint
	// from the parameter `fromMapstringstring160` to result `intoString120`
	// (`intoString120` is now tainted).
	intoString120 := mime.FormatMediaType("", fromMapstringstring160)

	// Return the tainted `intoString120`:
	return intoString120
}

func TaintStepTest_MimeParseMediaType_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString818` into `intoString221`.

	// Assume that `sourceCQL` has the underlying type of `fromString818`:
	fromString818 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString818` to result `intoString221`
	// (`intoString221` is now tainted).
	intoString221, _, _ := mime.ParseMediaType(fromString818)

	// Return the tainted `intoString221`:
	return intoString221
}

func TaintStepTest_MimeParseMediaType_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString323` into `intoMapstringstring254`.

	// Assume that `sourceCQL` has the underlying type of `fromString323`:
	fromString323 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString323` to result `intoMapstringstring254`
	// (`intoMapstringstring254` is now tainted).
	_, intoMapstringstring254, _ := mime.ParseMediaType(fromString323)

	// Return the tainted `intoMapstringstring254`:
	return intoMapstringstring254
}

func TaintStepTest_MimeWordDecoderDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString956` into `intoString213`.

	// Assume that `sourceCQL` has the underlying type of `fromString956`:
	fromString956 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mime.WordDecoder

	// Call the method that transfers the taint
	// from the parameter `fromString956` to the result `intoString213`
	// (`intoString213` is now tainted).
	intoString213, _ := mediumObjCQL.Decode(fromString956)

	// Return the tainted `intoString213`:
	return intoString213
}

func TaintStepTest_MimeWordDecoderDecodeHeader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString797` into `intoString410`.

	// Assume that `sourceCQL` has the underlying type of `fromString797`:
	fromString797 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mime.WordDecoder

	// Call the method that transfers the taint
	// from the parameter `fromString797` to the result `intoString410`
	// (`intoString410` is now tainted).
	intoString410, _ := mediumObjCQL.DecodeHeader(fromString797)

	// Return the tainted `intoString410`:
	return intoString410
}

func TaintStepTest_MimeWordEncoderEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString503` into `intoString990`.

	// Assume that `sourceCQL` has the underlying type of `fromString503`:
	fromString503 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mime.WordEncoder

	// Call the method that transfers the taint
	// from the parameter `fromString503` to the result `intoString990`
	// (`intoString990` is now tainted).
	intoString990 := mediumObjCQL.Encode("", fromString503)

	// Return the tainted `intoString990`:
	return intoString990
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
