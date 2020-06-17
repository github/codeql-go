package main

import "unicode/utf8"

func TaintStepTest_UnicodeUtf8DecodeLastRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte924` into `intoRune679`.

	// Assume that `sourceCQL` has the underlying type of `fromByte924`:
	fromByte924 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte924` to result `intoRune679`
	// (`intoRune679` is now tainted).
	intoRune679, _ := utf8.DecodeLastRune(fromByte924)

	// Return the tainted `intoRune679`:
	return intoRune679
}

func TaintStepTest_UnicodeUtf8DecodeLastRuneInString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString535` into `intoRune969`.

	// Assume that `sourceCQL` has the underlying type of `fromString535`:
	fromString535 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString535` to result `intoRune969`
	// (`intoRune969` is now tainted).
	intoRune969, _ := utf8.DecodeLastRuneInString(fromString535)

	// Return the tainted `intoRune969`:
	return intoRune969
}

func TaintStepTest_UnicodeUtf8DecodeRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte717` into `intoRune736`.

	// Assume that `sourceCQL` has the underlying type of `fromByte717`:
	fromByte717 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte717` to result `intoRune736`
	// (`intoRune736` is now tainted).
	intoRune736, _ := utf8.DecodeRune(fromByte717)

	// Return the tainted `intoRune736`:
	return intoRune736
}

func TaintStepTest_UnicodeUtf8DecodeRuneInString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString135` into `intoRune774`.

	// Assume that `sourceCQL` has the underlying type of `fromString135`:
	fromString135 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString135` to result `intoRune774`
	// (`intoRune774` is now tainted).
	intoRune774, _ := utf8.DecodeRuneInString(fromString135)

	// Return the tainted `intoRune774`:
	return intoRune774
}

func TaintStepTest_UnicodeUtf8EncodeRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune609` into `intoByte560`.

	// Assume that `sourceCQL` has the underlying type of `fromRune609`:
	fromRune609 := sourceCQL.(rune)

	// Declare `intoByte560` variable:
	var intoByte560 []byte

	// Call the function that transfers the taint
	// from the parameter `fromRune609` to parameter `intoByte560`;
	// `intoByte560` is now tainted.
	utf8.EncodeRune(intoByte560, fromRune609)

	// Return the tainted `intoByte560`:
	return intoByte560
}

func RunAllTaints_UnicodeUtf8() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeUtf8DecodeLastRune_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeUtf8DecodeLastRuneInString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeUtf8DecodeRune_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeUtf8DecodeRuneInString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeUtf8EncodeRune_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
