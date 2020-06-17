// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "strconv"

func TaintStepTest_StrconvAppendQuote_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte656` into `intoByte414`.

	// Assume that `sourceCQL` has the underlying type of `fromByte656`:
	fromByte656 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte656` to result `intoByte414`
	// (`intoByte414` is now tainted).
	intoByte414 := strconv.AppendQuote(fromByte656, "")

	// Return the tainted `intoByte414`:
	return intoByte414
}

func TaintStepTest_StrconvAppendQuote_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString518` into `intoByte650`.

	// Assume that `sourceCQL` has the underlying type of `fromString518`:
	fromString518 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString518` to result `intoByte650`
	// (`intoByte650` is now tainted).
	intoByte650 := strconv.AppendQuote(nil, fromString518)

	// Return the tainted `intoByte650`:
	return intoByte650
}

func TaintStepTest_StrconvAppendQuoteRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte784` into `intoByte957`.

	// Assume that `sourceCQL` has the underlying type of `fromByte784`:
	fromByte784 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte784` to result `intoByte957`
	// (`intoByte957` is now tainted).
	intoByte957 := strconv.AppendQuoteRune(fromByte784, 0)

	// Return the tainted `intoByte957`:
	return intoByte957
}

func TaintStepTest_StrconvAppendQuoteRune_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune520` into `intoByte443`.

	// Assume that `sourceCQL` has the underlying type of `fromRune520`:
	fromRune520 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune520` to result `intoByte443`
	// (`intoByte443` is now tainted).
	intoByte443 := strconv.AppendQuoteRune(nil, fromRune520)

	// Return the tainted `intoByte443`:
	return intoByte443
}

func TaintStepTest_StrconvAppendQuoteRuneToASCII_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte127` into `intoByte483`.

	// Assume that `sourceCQL` has the underlying type of `fromByte127`:
	fromByte127 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte127` to result `intoByte483`
	// (`intoByte483` is now tainted).
	intoByte483 := strconv.AppendQuoteRuneToASCII(fromByte127, 0)

	// Return the tainted `intoByte483`:
	return intoByte483
}

func TaintStepTest_StrconvAppendQuoteRuneToASCII_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune989` into `intoByte982`.

	// Assume that `sourceCQL` has the underlying type of `fromRune989`:
	fromRune989 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune989` to result `intoByte982`
	// (`intoByte982` is now tainted).
	intoByte982 := strconv.AppendQuoteRuneToASCII(nil, fromRune989)

	// Return the tainted `intoByte982`:
	return intoByte982
}

func TaintStepTest_StrconvAppendQuoteRuneToGraphic_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte417` into `intoByte584`.

	// Assume that `sourceCQL` has the underlying type of `fromByte417`:
	fromByte417 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte417` to result `intoByte584`
	// (`intoByte584` is now tainted).
	intoByte584 := strconv.AppendQuoteRuneToGraphic(fromByte417, 0)

	// Return the tainted `intoByte584`:
	return intoByte584
}

func TaintStepTest_StrconvAppendQuoteRuneToGraphic_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune991` into `intoByte881`.

	// Assume that `sourceCQL` has the underlying type of `fromRune991`:
	fromRune991 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune991` to result `intoByte881`
	// (`intoByte881` is now tainted).
	intoByte881 := strconv.AppendQuoteRuneToGraphic(nil, fromRune991)

	// Return the tainted `intoByte881`:
	return intoByte881
}

func TaintStepTest_StrconvAppendQuoteToASCII_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte186` into `intoByte284`.

	// Assume that `sourceCQL` has the underlying type of `fromByte186`:
	fromByte186 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte186` to result `intoByte284`
	// (`intoByte284` is now tainted).
	intoByte284 := strconv.AppendQuoteToASCII(fromByte186, "")

	// Return the tainted `intoByte284`:
	return intoByte284
}

func TaintStepTest_StrconvAppendQuoteToASCII_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString908` into `intoByte137`.

	// Assume that `sourceCQL` has the underlying type of `fromString908`:
	fromString908 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString908` to result `intoByte137`
	// (`intoByte137` is now tainted).
	intoByte137 := strconv.AppendQuoteToASCII(nil, fromString908)

	// Return the tainted `intoByte137`:
	return intoByte137
}

func TaintStepTest_StrconvAppendQuoteToGraphic_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte494` into `intoByte873`.

	// Assume that `sourceCQL` has the underlying type of `fromByte494`:
	fromByte494 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte494` to result `intoByte873`
	// (`intoByte873` is now tainted).
	intoByte873 := strconv.AppendQuoteToGraphic(fromByte494, "")

	// Return the tainted `intoByte873`:
	return intoByte873
}

func TaintStepTest_StrconvAppendQuoteToGraphic_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString599` into `intoByte409`.

	// Assume that `sourceCQL` has the underlying type of `fromString599`:
	fromString599 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString599` to result `intoByte409`
	// (`intoByte409` is now tainted).
	intoByte409 := strconv.AppendQuoteToGraphic(nil, fromString599)

	// Return the tainted `intoByte409`:
	return intoByte409
}

func TaintStepTest_StrconvQuote_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString246` into `intoString898`.

	// Assume that `sourceCQL` has the underlying type of `fromString246`:
	fromString246 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString246` to result `intoString898`
	// (`intoString898` is now tainted).
	intoString898 := strconv.Quote(fromString246)

	// Return the tainted `intoString898`:
	return intoString898
}

func TaintStepTest_StrconvQuoteRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune598` into `intoString631`.

	// Assume that `sourceCQL` has the underlying type of `fromRune598`:
	fromRune598 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune598` to result `intoString631`
	// (`intoString631` is now tainted).
	intoString631 := strconv.QuoteRune(fromRune598)

	// Return the tainted `intoString631`:
	return intoString631
}

func TaintStepTest_StrconvQuoteRuneToASCII_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune165` into `intoString150`.

	// Assume that `sourceCQL` has the underlying type of `fromRune165`:
	fromRune165 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune165` to result `intoString150`
	// (`intoString150` is now tainted).
	intoString150 := strconv.QuoteRuneToASCII(fromRune165)

	// Return the tainted `intoString150`:
	return intoString150
}

func TaintStepTest_StrconvQuoteRuneToGraphic_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune340` into `intoString471`.

	// Assume that `sourceCQL` has the underlying type of `fromRune340`:
	fromRune340 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune340` to result `intoString471`
	// (`intoString471` is now tainted).
	intoString471 := strconv.QuoteRuneToGraphic(fromRune340)

	// Return the tainted `intoString471`:
	return intoString471
}

func TaintStepTest_StrconvQuoteToASCII_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString290` into `intoString758`.

	// Assume that `sourceCQL` has the underlying type of `fromString290`:
	fromString290 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString290` to result `intoString758`
	// (`intoString758` is now tainted).
	intoString758 := strconv.QuoteToASCII(fromString290)

	// Return the tainted `intoString758`:
	return intoString758
}

func TaintStepTest_StrconvQuoteToGraphic_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString396` into `intoString707`.

	// Assume that `sourceCQL` has the underlying type of `fromString396`:
	fromString396 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString396` to result `intoString707`
	// (`intoString707` is now tainted).
	intoString707 := strconv.QuoteToGraphic(fromString396)

	// Return the tainted `intoString707`:
	return intoString707
}

func TaintStepTest_StrconvUnquote_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString912` into `intoString718`.

	// Assume that `sourceCQL` has the underlying type of `fromString912`:
	fromString912 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString912` to result `intoString718`
	// (`intoString718` is now tainted).
	intoString718, _ := strconv.Unquote(fromString912)

	// Return the tainted `intoString718`:
	return intoString718
}

func TaintStepTest_StrconvUnquoteChar_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString972` into `intoRune633`.

	// Assume that `sourceCQL` has the underlying type of `fromString972`:
	fromString972 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString972` to result `intoRune633`
	// (`intoRune633` is now tainted).
	intoRune633, _, _, _ := strconv.UnquoteChar(fromString972, 0)

	// Return the tainted `intoRune633`:
	return intoRune633
}

func RunAllTaints_Strconv() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvAppendQuote_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvAppendQuote_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvAppendQuoteRune_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvAppendQuoteRune_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvAppendQuoteRuneToASCII_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvAppendQuoteRuneToASCII_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvAppendQuoteRuneToGraphic_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvAppendQuoteRuneToGraphic_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvAppendQuoteToASCII_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvAppendQuoteToASCII_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvAppendQuoteToGraphic_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvAppendQuoteToGraphic_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvQuote_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvQuoteRune_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvQuoteRuneToASCII_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvQuoteRuneToGraphic_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvQuoteToASCII_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvQuoteToGraphic_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvUnquote_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StrconvUnquoteChar_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
