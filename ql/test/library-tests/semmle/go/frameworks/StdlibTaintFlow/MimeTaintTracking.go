package main

import "mime"

func TaintStepTest_MimeFormatMediaType_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString687` into `intoString419`.

	// Assume that `sourceCQL` has the underlying type of `fromString687`:
	fromString687 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString687` to result `intoString419`
	// (`intoString419` is now tainted).
	intoString419 := mime.FormatMediaType(fromString687, nil)

	// Sink the tainted `intoString419`:
	sink(intoString419)
}

func TaintStepTest_MimeFormatMediaType_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromMapstringstring894` into `intoString410`.

	// Assume that `sourceCQL` has the underlying type of `fromMapstringstring894`:
	fromMapstringstring894 := sourceCQL.(map[string]string)

	// Call the function that transfers the taint
	// from the parameter `fromMapstringstring894` to result `intoString410`
	// (`intoString410` is now tainted).
	intoString410 := mime.FormatMediaType("", fromMapstringstring894)

	// Sink the tainted `intoString410`:
	sink(intoString410)
}

func TaintStepTest_MimeParseMediaType_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString298` into `intoString986`.

	// Assume that `sourceCQL` has the underlying type of `fromString298`:
	fromString298 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString298` to result `intoString986`
	// (`intoString986` is now tainted).
	intoString986, _, _ := mime.ParseMediaType(fromString298)

	// Sink the tainted `intoString986`:
	sink(intoString986)
}

func TaintStepTest_MimeParseMediaType_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromString310` into `intoMapstringstring600`.

	// Assume that `sourceCQL` has the underlying type of `fromString310`:
	fromString310 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString310` to result `intoMapstringstring600`
	// (`intoMapstringstring600` is now tainted).
	_, intoMapstringstring600, _ := mime.ParseMediaType(fromString310)

	// Sink the tainted `intoMapstringstring600`:
	sink(intoMapstringstring600)
}

func TaintStepTest_MimeWordDecoderDecode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString438` into `intoString587`.

	// Assume that `sourceCQL` has the underlying type of `fromString438`:
	fromString438 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mime.WordDecoder

	// Call the method that transfers the taint
	// from the parameter `fromString438` to the result `intoString587`
	// (`intoString587` is now tainted).
	intoString587, _ := mediumObjCQL.Decode(fromString438)

	// Sink the tainted `intoString587`:
	sink(intoString587)
}

func TaintStepTest_MimeWordDecoderDecodeHeader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString567` into `intoString479`.

	// Assume that `sourceCQL` has the underlying type of `fromString567`:
	fromString567 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mime.WordDecoder

	// Call the method that transfers the taint
	// from the parameter `fromString567` to the result `intoString479`
	// (`intoString479` is now tainted).
	intoString479, _ := mediumObjCQL.DecodeHeader(fromString567)

	// Sink the tainted `intoString479`:
	sink(intoString479)
}

func TaintStepTest_MimeWordEncoderEncode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString357` into `intoString426`.

	// Assume that `sourceCQL` has the underlying type of `fromString357`:
	fromString357 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mime.WordEncoder

	// Call the method that transfers the taint
	// from the parameter `fromString357` to the result `intoString426`
	// (`intoString426` is now tainted).
	intoString426 := mediumObjCQL.Encode("", fromString357)

	// Sink the tainted `intoString426`:
	sink(intoString426)
}

func RunAllTaints_Mime() {
	{
		source := newSource()
		TaintStepTest_MimeFormatMediaType_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeFormatMediaType_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeParseMediaType_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeParseMediaType_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeWordDecoderDecode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeWordDecoderDecodeHeader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeWordEncoderEncode_B0I0O0(source)
	}
}
