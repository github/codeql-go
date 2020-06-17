// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "unicode/utf8"

func TaintStepTest_UnicodeUtf8DecodeLastRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte656` into `intoRune414`.

	// Assume that `sourceCQL` has the underlying type of `fromByte656`:
	fromByte656 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte656` to result `intoRune414`
	// (`intoRune414` is now tainted).
	intoRune414, _ := utf8.DecodeLastRune(fromByte656)

	// Return the tainted `intoRune414`:
	return intoRune414
}

func TaintStepTest_UnicodeUtf8DecodeLastRuneInString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString518` into `intoRune650`.

	// Assume that `sourceCQL` has the underlying type of `fromString518`:
	fromString518 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString518` to result `intoRune650`
	// (`intoRune650` is now tainted).
	intoRune650, _ := utf8.DecodeLastRuneInString(fromString518)

	// Return the tainted `intoRune650`:
	return intoRune650
}

func TaintStepTest_UnicodeUtf8DecodeRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte784` into `intoRune957`.

	// Assume that `sourceCQL` has the underlying type of `fromByte784`:
	fromByte784 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte784` to result `intoRune957`
	// (`intoRune957` is now tainted).
	intoRune957, _ := utf8.DecodeRune(fromByte784)

	// Return the tainted `intoRune957`:
	return intoRune957
}

func TaintStepTest_UnicodeUtf8DecodeRuneInString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString520` into `intoRune443`.

	// Assume that `sourceCQL` has the underlying type of `fromString520`:
	fromString520 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString520` to result `intoRune443`
	// (`intoRune443` is now tainted).
	intoRune443, _ := utf8.DecodeRuneInString(fromString520)

	// Return the tainted `intoRune443`:
	return intoRune443
}

func TaintStepTest_UnicodeUtf8EncodeRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune127` into `intoByte483`.

	// Assume that `sourceCQL` has the underlying type of `fromRune127`:
	fromRune127 := sourceCQL.(rune)

	// Declare `intoByte483` variable:
	var intoByte483 []byte

	// Call the function that transfers the taint
	// from the parameter `fromRune127` to parameter `intoByte483`;
	// `intoByte483` is now tainted.
	utf8.EncodeRune(intoByte483, fromRune127)

	// Return the tainted `intoByte483`:
	return intoByte483
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
