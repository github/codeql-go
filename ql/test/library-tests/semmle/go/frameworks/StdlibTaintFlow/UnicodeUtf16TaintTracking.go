package main

import "unicode/utf16"

func TaintStepTest_UnicodeUtf16Decode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUint16190` into `intoRune647`.

	// Assume that `sourceCQL` has the underlying type of `fromUint16190`:
	fromUint16190 := sourceCQL.([]uint16)

	// Call the function that transfers the taint
	// from the parameter `fromUint16190` to result `intoRune647`
	// (`intoRune647` is now tainted).
	intoRune647 := utf16.Decode(fromUint16190)

	// Return the tainted `intoRune647`:
	return intoRune647
}

func TaintStepTest_UnicodeUtf16DecodeRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune880` into `intoRune992`.

	// Assume that `sourceCQL` has the underlying type of `fromRune880`:
	fromRune880 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune880` to result `intoRune992`
	// (`intoRune992` is now tainted).
	intoRune992 := utf16.DecodeRune(fromRune880, 0)

	// Return the tainted `intoRune992`:
	return intoRune992
}

func TaintStepTest_UnicodeUtf16DecodeRune_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune306` into `intoRune255`.

	// Assume that `sourceCQL` has the underlying type of `fromRune306`:
	fromRune306 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune306` to result `intoRune255`
	// (`intoRune255` is now tainted).
	intoRune255 := utf16.DecodeRune(0, fromRune306)

	// Return the tainted `intoRune255`:
	return intoRune255
}

func TaintStepTest_UnicodeUtf16Encode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune333` into `intoUint16676`.

	// Assume that `sourceCQL` has the underlying type of `fromRune333`:
	fromRune333 := sourceCQL.([]rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune333` to result `intoUint16676`
	// (`intoUint16676` is now tainted).
	intoUint16676 := utf16.Encode(fromRune333)

	// Return the tainted `intoUint16676`:
	return intoUint16676
}

func TaintStepTest_UnicodeUtf16EncodeRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune849` into `intoRune894`.

	// Assume that `sourceCQL` has the underlying type of `fromRune849`:
	fromRune849 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune849` to result `intoRune894`
	// (`intoRune894` is now tainted).
	intoRune894, _ := utf16.EncodeRune(fromRune849)

	// Return the tainted `intoRune894`:
	return intoRune894
}

func TaintStepTest_UnicodeUtf16EncodeRune_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune579` into `intoRune335`.

	// Assume that `sourceCQL` has the underlying type of `fromRune579`:
	fromRune579 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune579` to result `intoRune335`
	// (`intoRune335` is now tainted).
	_, intoRune335 := utf16.EncodeRune(fromRune579)

	// Return the tainted `intoRune335`:
	return intoRune335
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
