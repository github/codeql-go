package main

import "mime"

func TaintStepTest_MimeFormatMediaType(sourceCQL interface{}) {
	// The flow is from `t` into `into112`.

	// Assume that `sourceCQL` has the underlying type of `t`:
	t := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `t` to result `into112`
	// (`into112` is now tainted).
	into112 := mime.FormatMediaType(t, nil)

	// Sink the tainted `into112`:
	sink(into112)
}

func TaintStepTest_MimeParseMediaType(sourceCQL interface{}) {
	// The flow is from `v` into `mediatype`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `v` to result `mediatype`
	// (`mediatype` is now tainted).
	mediatype, _, _ := mime.ParseMediaType(v)

	// Sink the tainted `mediatype`:
	sink(mediatype)
}

func TaintStepTest_MimeWordDecoderDecode(sourceCQL interface{}) {
	// The flow is from `word` into `into551`.

	// Assume that `sourceCQL` has the underlying type of `word`:
	word := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mime.WordDecoder

	// Call the method that transfers the taint
	// from the parameter `word` to the result `into551`
	// (`into551` is now tainted).
	into551, _ := mediumObjCQL.Decode(word)

	// Sink the tainted `into551`:
	sink(into551)
}

func TaintStepTest_MimeWordDecoderDecodeHeader(sourceCQL interface{}) {
	// The flow is from `header` into `into177`.

	// Assume that `sourceCQL` has the underlying type of `header`:
	header := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mime.WordDecoder

	// Call the method that transfers the taint
	// from the parameter `header` to the result `into177`
	// (`into177` is now tainted).
	into177, _ := mediumObjCQL.DecodeHeader(header)

	// Sink the tainted `into177`:
	sink(into177)
}

func RunAllTaints_Mime(v interface{}) {
	{
		source := newSource()
		TaintStepTest_MimeFormatMediaType(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeParseMediaType(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeWordDecoderDecode(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeWordDecoderDecodeHeader(source)
	}
}
