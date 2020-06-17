// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "unicode/utf16"

func TaintStepTest_UnicodeUtf16Decode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUint16865` into `intoRune342`.

	// Assume that `sourceCQL` has the underlying type of `fromUint16865`:
	fromUint16865 := sourceCQL.([]uint16)

	// Call the function that transfers the taint
	// from the parameter `fromUint16865` to result `intoRune342`
	// (`intoRune342` is now tainted).
	intoRune342 := utf16.Decode(fromUint16865)

	// Return the tainted `intoRune342`:
	return intoRune342
}

func TaintStepTest_UnicodeUtf16DecodeRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune823` into `intoRune288`.

	// Assume that `sourceCQL` has the underlying type of `fromRune823`:
	fromRune823 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune823` to result `intoRune288`
	// (`intoRune288` is now tainted).
	intoRune288 := utf16.DecodeRune(fromRune823, 0)

	// Return the tainted `intoRune288`:
	return intoRune288
}

func TaintStepTest_UnicodeUtf16DecodeRune_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune248` into `intoRune662`.

	// Assume that `sourceCQL` has the underlying type of `fromRune248`:
	fromRune248 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune248` to result `intoRune662`
	// (`intoRune662` is now tainted).
	intoRune662 := utf16.DecodeRune(0, fromRune248)

	// Return the tainted `intoRune662`:
	return intoRune662
}

func TaintStepTest_UnicodeUtf16Encode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune704` into `intoUint16461`.

	// Assume that `sourceCQL` has the underlying type of `fromRune704`:
	fromRune704 := sourceCQL.([]rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune704` to result `intoUint16461`
	// (`intoUint16461` is now tainted).
	intoUint16461 := utf16.Encode(fromRune704)

	// Return the tainted `intoUint16461`:
	return intoUint16461
}

func TaintStepTest_UnicodeUtf16EncodeRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune168` into `intoRune704`.

	// Assume that `sourceCQL` has the underlying type of `fromRune168`:
	fromRune168 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune168` to result `intoRune704`
	// (`intoRune704` is now tainted).
	intoRune704, _ := utf16.EncodeRune(fromRune168)

	// Return the tainted `intoRune704`:
	return intoRune704
}

func TaintStepTest_UnicodeUtf16EncodeRune_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune698` into `intoRune845`.

	// Assume that `sourceCQL` has the underlying type of `fromRune698`:
	fromRune698 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune698` to result `intoRune845`
	// (`intoRune845` is now tainted).
	_, intoRune845 := utf16.EncodeRune(fromRune698)

	// Return the tainted `intoRune845`:
	return intoRune845
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
