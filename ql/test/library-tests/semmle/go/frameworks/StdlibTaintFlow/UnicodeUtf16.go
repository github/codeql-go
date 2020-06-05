package main

import "unicode/utf16"

func TaintStepTest_UnicodeUtf16Decode(sourceCQL interface{}) {
	// The flow is from `s` into `into795`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]uint16)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into795`
	// (`into795` is now tainted).
	into795 := utf16.Decode(s)

	// Sink the tainted `into795`:
	sink(into795)
}

func TaintStepTest_UnicodeUtf16DecodeRune(sourceCQL interface{}) {
	// The flow is from `r1` into `into628`.

	// Assume that `sourceCQL` has the underlying type of `r1`:
	r1 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `r1` to result `into628`
	// (`into628` is now tainted).
	into628 := utf16.DecodeRune(r1, 0)

	// Sink the tainted `into628`:
	sink(into628)
}

func TaintStepTest_UnicodeUtf16Encode(sourceCQL interface{}) {
	// The flow is from `s` into `into365`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]rune)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into365`
	// (`into365` is now tainted).
	into365 := utf16.Encode(s)

	// Sink the tainted `into365`:
	sink(into365)
}

func TaintStepTest_UnicodeUtf16EncodeRune(sourceCQL interface{}) {
	// The flow is from `r` into `r1`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `r` to result `r1`
	// (`r1` is now tainted).
	r1, _ := utf16.EncodeRune(r)

	// Sink the tainted `r1`:
	sink(r1)
}

func RunAllTaints_UnicodeUtf16(v interface{}) {
	{
		source := newSource()
		TaintStepTest_UnicodeUtf16Decode(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeUtf16DecodeRune(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeUtf16Encode(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeUtf16EncodeRune(source)
	}
}
