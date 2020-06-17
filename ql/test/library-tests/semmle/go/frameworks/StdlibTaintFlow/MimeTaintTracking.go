package main

import "mime"

func TaintStepTest_MimeFormatMediaType_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString884` into `intoString437`.

	// Assume that `sourceCQL` has the underlying type of `fromString884`:
	fromString884 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString884` to result `intoString437`
	// (`intoString437` is now tainted).
	intoString437 := mime.FormatMediaType(fromString884, nil)

	// Return the tainted `intoString437`:
	return intoString437
}

func TaintStepTest_MimeFormatMediaType_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMapstringstring280` into `intoString354`.

	// Assume that `sourceCQL` has the underlying type of `fromMapstringstring280`:
	fromMapstringstring280 := sourceCQL.(map[string]string)

	// Call the function that transfers the taint
	// from the parameter `fromMapstringstring280` to result `intoString354`
	// (`intoString354` is now tainted).
	intoString354 := mime.FormatMediaType("", fromMapstringstring280)

	// Return the tainted `intoString354`:
	return intoString354
}

func TaintStepTest_MimeParseMediaType_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString508` into `intoString806`.

	// Assume that `sourceCQL` has the underlying type of `fromString508`:
	fromString508 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString508` to result `intoString806`
	// (`intoString806` is now tainted).
	intoString806, _, _ := mime.ParseMediaType(fromString508)

	// Return the tainted `intoString806`:
	return intoString806
}

func TaintStepTest_MimeParseMediaType_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString521` into `intoMapstringstring871`.

	// Assume that `sourceCQL` has the underlying type of `fromString521`:
	fromString521 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString521` to result `intoMapstringstring871`
	// (`intoMapstringstring871` is now tainted).
	_, intoMapstringstring871, _ := mime.ParseMediaType(fromString521)

	// Return the tainted `intoMapstringstring871`:
	return intoMapstringstring871
}

func TaintStepTest_MimeWordDecoderDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString112` into `intoString524`.

	// Assume that `sourceCQL` has the underlying type of `fromString112`:
	fromString112 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mime.WordDecoder

	// Call the method that transfers the taint
	// from the parameter `fromString112` to the result `intoString524`
	// (`intoString524` is now tainted).
	intoString524, _ := mediumObjCQL.Decode(fromString112)

	// Return the tainted `intoString524`:
	return intoString524
}

func TaintStepTest_MimeWordDecoderDecodeHeader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString586` into `intoString976`.

	// Assume that `sourceCQL` has the underlying type of `fromString586`:
	fromString586 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mime.WordDecoder

	// Call the method that transfers the taint
	// from the parameter `fromString586` to the result `intoString976`
	// (`intoString976` is now tainted).
	intoString976, _ := mediumObjCQL.DecodeHeader(fromString586)

	// Return the tainted `intoString976`:
	return intoString976
}

func TaintStepTest_MimeWordEncoderEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString315` into `intoString262`.

	// Assume that `sourceCQL` has the underlying type of `fromString315`:
	fromString315 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mime.WordEncoder

	// Call the method that transfers the taint
	// from the parameter `fromString315` to the result `intoString262`
	// (`intoString262` is now tainted).
	intoString262 := mediumObjCQL.Encode("", fromString315)

	// Return the tainted `intoString262`:
	return intoString262
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
