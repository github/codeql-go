// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "unicode/utf16"

func TaintStepTest_UnicodeUtf16Decode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUint16656` into `intoRune414`.

	// Assume that `sourceCQL` has the underlying type of `fromUint16656`:
	fromUint16656 := sourceCQL.([]uint16)

	// Call the function that transfers the taint
	// from the parameter `fromUint16656` to result `intoRune414`
	// (`intoRune414` is now tainted).
	intoRune414 := utf16.Decode(fromUint16656)

	// Return the tainted `intoRune414`:
	return intoRune414
}

func TaintStepTest_UnicodeUtf16DecodeRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune518` into `intoRune650`.

	// Assume that `sourceCQL` has the underlying type of `fromRune518`:
	fromRune518 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune518` to result `intoRune650`
	// (`intoRune650` is now tainted).
	intoRune650 := utf16.DecodeRune(fromRune518, 0)

	// Return the tainted `intoRune650`:
	return intoRune650
}

func TaintStepTest_UnicodeUtf16DecodeRune_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune784` into `intoRune957`.

	// Assume that `sourceCQL` has the underlying type of `fromRune784`:
	fromRune784 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune784` to result `intoRune957`
	// (`intoRune957` is now tainted).
	intoRune957 := utf16.DecodeRune(0, fromRune784)

	// Return the tainted `intoRune957`:
	return intoRune957
}

func TaintStepTest_UnicodeUtf16Encode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune520` into `intoUint16443`.

	// Assume that `sourceCQL` has the underlying type of `fromRune520`:
	fromRune520 := sourceCQL.([]rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune520` to result `intoUint16443`
	// (`intoUint16443` is now tainted).
	intoUint16443 := utf16.Encode(fromRune520)

	// Return the tainted `intoUint16443`:
	return intoUint16443
}

func TaintStepTest_UnicodeUtf16EncodeRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune127` into `intoRune483`.

	// Assume that `sourceCQL` has the underlying type of `fromRune127`:
	fromRune127 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune127` to result `intoRune483`
	// (`intoRune483` is now tainted).
	intoRune483, _ := utf16.EncodeRune(fromRune127)

	// Return the tainted `intoRune483`:
	return intoRune483
}

func TaintStepTest_UnicodeUtf16EncodeRune_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune989` into `intoRune982`.

	// Assume that `sourceCQL` has the underlying type of `fromRune989`:
	fromRune989 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune989` to result `intoRune982`
	// (`intoRune982` is now tainted).
	_, intoRune982 := utf16.EncodeRune(fromRune989)

	// Return the tainted `intoRune982`:
	return intoRune982
}

func RunAllTaints_UnicodeUtf16() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeUtf16Decode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeUtf16DecodeRune_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeUtf16DecodeRune_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeUtf16Encode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeUtf16EncodeRune_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeUtf16EncodeRune_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
