package main

import "unicode/utf8"

func TaintStepTest_UnicodeUtf8DecodeLastRune_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte730` into `intoRune943`.

	// Assume that `sourceCQL` has the underlying type of `fromByte730`:
	fromByte730 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte730` to result `intoRune943`
	// (`intoRune943` is now tainted).
	intoRune943, _ := utf8.DecodeLastRune(fromByte730)

	// Sink the tainted `intoRune943`:
	sink(intoRune943)
}

func TaintStepTest_UnicodeUtf8DecodeLastRuneInString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString702` into `intoRune800`.

	// Assume that `sourceCQL` has the underlying type of `fromString702`:
	fromString702 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString702` to result `intoRune800`
	// (`intoRune800` is now tainted).
	intoRune800, _ := utf8.DecodeLastRuneInString(fromString702)

	// Sink the tainted `intoRune800`:
	sink(intoRune800)
}

func TaintStepTest_UnicodeUtf8DecodeRune_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte528` into `intoRune134`.

	// Assume that `sourceCQL` has the underlying type of `fromByte528`:
	fromByte528 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte528` to result `intoRune134`
	// (`intoRune134` is now tainted).
	intoRune134, _ := utf8.DecodeRune(fromByte528)

	// Sink the tainted `intoRune134`:
	sink(intoRune134)
}

func TaintStepTest_UnicodeUtf8DecodeRuneInString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString486` into `intoRune863`.

	// Assume that `sourceCQL` has the underlying type of `fromString486`:
	fromString486 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString486` to result `intoRune863`
	// (`intoRune863` is now tainted).
	intoRune863, _ := utf8.DecodeRuneInString(fromString486)

	// Sink the tainted `intoRune863`:
	sink(intoRune863)
}

func TaintStepTest_UnicodeUtf8EncodeRune_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune624` into `intoByte703`.

	// Assume that `sourceCQL` has the underlying type of `fromRune624`:
	fromRune624 := sourceCQL.(rune)

	// Declare `intoByte703` variable:
	var intoByte703 []byte

	// Call the function that transfers the taint
	// from the parameter `fromRune624` to parameter `intoByte703`;
	// `intoByte703` is now tainted.
	utf8.EncodeRune(intoByte703, fromRune624)

	// Sink the tainted `intoByte703`:
	sink(intoByte703)
}

func RunAllTaints_UnicodeUtf8() {
	{
		source := newSource()
		TaintStepTest_UnicodeUtf8DecodeLastRune_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeUtf8DecodeLastRuneInString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeUtf8DecodeRune_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeUtf8DecodeRuneInString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeUtf8EncodeRune_B0I0O0(source)
	}
}
