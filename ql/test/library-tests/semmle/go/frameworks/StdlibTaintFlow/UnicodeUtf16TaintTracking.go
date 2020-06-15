package main

import "unicode/utf16"

func TaintStepTest_UnicodeUtf16Decode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUint16661` into `intoRune293`.

	// Assume that `sourceCQL` has the underlying type of `fromUint16661`:
	fromUint16661 := sourceCQL.([]uint16)

	// Call the function that transfers the taint
	// from the parameter `fromUint16661` to result `intoRune293`
	// (`intoRune293` is now tainted).
	intoRune293 := utf16.Decode(fromUint16661)

	// Sink the tainted `intoRune293`:
	sink(intoRune293)
}

func TaintStepTest_UnicodeUtf16DecodeRune_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune192` into `intoRune580`.

	// Assume that `sourceCQL` has the underlying type of `fromRune192`:
	fromRune192 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune192` to result `intoRune580`
	// (`intoRune580` is now tainted).
	intoRune580 := utf16.DecodeRune(fromRune192, 0)

	// Sink the tainted `intoRune580`:
	sink(intoRune580)
}

func TaintStepTest_UnicodeUtf16DecodeRune_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromRune213` into `intoRune848`.

	// Assume that `sourceCQL` has the underlying type of `fromRune213`:
	fromRune213 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune213` to result `intoRune848`
	// (`intoRune848` is now tainted).
	intoRune848 := utf16.DecodeRune(0, fromRune213)

	// Sink the tainted `intoRune848`:
	sink(intoRune848)
}

func TaintStepTest_UnicodeUtf16Encode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune987` into `intoUint16354`.

	// Assume that `sourceCQL` has the underlying type of `fromRune987`:
	fromRune987 := sourceCQL.([]rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune987` to result `intoUint16354`
	// (`intoUint16354` is now tainted).
	intoUint16354 := utf16.Encode(fromRune987)

	// Sink the tainted `intoUint16354`:
	sink(intoUint16354)
}

func TaintStepTest_UnicodeUtf16EncodeRune_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune966` into `intoRune143`.

	// Assume that `sourceCQL` has the underlying type of `fromRune966`:
	fromRune966 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune966` to result `intoRune143`
	// (`intoRune143` is now tainted).
	intoRune143, _ := utf16.EncodeRune(fromRune966)

	// Sink the tainted `intoRune143`:
	sink(intoRune143)
}

func TaintStepTest_UnicodeUtf16EncodeRune_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromRune336` into `intoRune178`.

	// Assume that `sourceCQL` has the underlying type of `fromRune336`:
	fromRune336 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune336` to result `intoRune178`
	// (`intoRune178` is now tainted).
	_, intoRune178 := utf16.EncodeRune(fromRune336)

	// Sink the tainted `intoRune178`:
	sink(intoRune178)
}

func RunAllTaints_UnicodeUtf16() {
	{
		source := newSource()
		TaintStepTest_UnicodeUtf16Decode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeUtf16DecodeRune_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeUtf16DecodeRune_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeUtf16Encode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeUtf16EncodeRune_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeUtf16EncodeRune_B0I0O1(source)
	}
}
