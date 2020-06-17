// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "unicode/utf8"

func TaintStepTest_UnicodeUtf8DecodeLastRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte417` into `intoRune976`.

	// Assume that `sourceCQL` has the underlying type of `fromByte417`:
	fromByte417 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte417` to result `intoRune976`
	// (`intoRune976` is now tainted).
	intoRune976, _ := utf8.DecodeLastRune(fromByte417)

	// Return the tainted `intoRune976`:
	return intoRune976
}

func TaintStepTest_UnicodeUtf8DecodeLastRuneInString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString426` into `intoRune969`.

	// Assume that `sourceCQL` has the underlying type of `fromString426`:
	fromString426 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString426` to result `intoRune969`
	// (`intoRune969` is now tainted).
	intoRune969, _ := utf8.DecodeLastRuneInString(fromString426)

	// Return the tainted `intoRune969`:
	return intoRune969
}

func TaintStepTest_UnicodeUtf8DecodeRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte153` into `intoRune423`.

	// Assume that `sourceCQL` has the underlying type of `fromByte153`:
	fromByte153 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte153` to result `intoRune423`
	// (`intoRune423` is now tainted).
	intoRune423, _ := utf8.DecodeRune(fromByte153)

	// Return the tainted `intoRune423`:
	return intoRune423
}

func TaintStepTest_UnicodeUtf8DecodeRuneInString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString229` into `intoRune970`.

	// Assume that `sourceCQL` has the underlying type of `fromString229`:
	fromString229 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString229` to result `intoRune970`
	// (`intoRune970` is now tainted).
	intoRune970, _ := utf8.DecodeRuneInString(fromString229)

	// Return the tainted `intoRune970`:
	return intoRune970
}

func TaintStepTest_UnicodeUtf8EncodeRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune866` into `intoByte562`.

	// Assume that `sourceCQL` has the underlying type of `fromRune866`:
	fromRune866 := sourceCQL.(rune)

	// Declare `intoByte562` variable:
	var intoByte562 []byte

	// Call the function that transfers the taint
	// from the parameter `fromRune866` to parameter `intoByte562`;
	// `intoByte562` is now tainted.
	utf8.EncodeRune(intoByte562, fromRune866)

	// Return the tainted `intoByte562`:
	return intoByte562
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
