// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "strconv"

func TaintStepTest_StrconvAppendQuote_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte307` into `intoByte111`.

	// Assume that `sourceCQL` has the underlying type of `fromByte307`:
	fromByte307 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte307` to result `intoByte111`
	// (`intoByte111` is now tainted).
	intoByte111 := strconv.AppendQuote(fromByte307, "")

	// Return the tainted `intoByte111`:
	return intoByte111
}

func TaintStepTest_StrconvAppendQuote_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString271` into `intoByte825`.

	// Assume that `sourceCQL` has the underlying type of `fromString271`:
	fromString271 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString271` to result `intoByte825`
	// (`intoByte825` is now tainted).
	intoByte825 := strconv.AppendQuote(nil, fromString271)

	// Return the tainted `intoByte825`:
	return intoByte825
}

func TaintStepTest_StrconvAppendQuoteRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte692` into `intoByte352`.

	// Assume that `sourceCQL` has the underlying type of `fromByte692`:
	fromByte692 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte692` to result `intoByte352`
	// (`intoByte352` is now tainted).
	intoByte352 := strconv.AppendQuoteRune(fromByte692, 0)

	// Return the tainted `intoByte352`:
	return intoByte352
}

func TaintStepTest_StrconvAppendQuoteRune_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune429` into `intoByte524`.

	// Assume that `sourceCQL` has the underlying type of `fromRune429`:
	fromRune429 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune429` to result `intoByte524`
	// (`intoByte524` is now tainted).
	intoByte524 := strconv.AppendQuoteRune(nil, fromRune429)

	// Return the tainted `intoByte524`:
	return intoByte524
}

func TaintStepTest_StrconvAppendQuoteRuneToASCII_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte384` into `intoByte836`.

	// Assume that `sourceCQL` has the underlying type of `fromByte384`:
	fromByte384 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte384` to result `intoByte836`
	// (`intoByte836` is now tainted).
	intoByte836 := strconv.AppendQuoteRuneToASCII(fromByte384, 0)

	// Return the tainted `intoByte836`:
	return intoByte836
}

func TaintStepTest_StrconvAppendQuoteRuneToASCII_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune632` into `intoByte614`.

	// Assume that `sourceCQL` has the underlying type of `fromRune632`:
	fromRune632 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune632` to result `intoByte614`
	// (`intoByte614` is now tainted).
	intoByte614 := strconv.AppendQuoteRuneToASCII(nil, fromRune632)

	// Return the tainted `intoByte614`:
	return intoByte614
}

func TaintStepTest_StrconvAppendQuoteRuneToGraphic_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte117` into `intoByte285`.

	// Assume that `sourceCQL` has the underlying type of `fromByte117`:
	fromByte117 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte117` to result `intoByte285`
	// (`intoByte285` is now tainted).
	intoByte285 := strconv.AppendQuoteRuneToGraphic(fromByte117, 0)

	// Return the tainted `intoByte285`:
	return intoByte285
}

func TaintStepTest_StrconvAppendQuoteRuneToGraphic_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune427` into `intoByte882`.

	// Assume that `sourceCQL` has the underlying type of `fromRune427`:
	fromRune427 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune427` to result `intoByte882`
	// (`intoByte882` is now tainted).
	intoByte882 := strconv.AppendQuoteRuneToGraphic(nil, fromRune427)

	// Return the tainted `intoByte882`:
	return intoByte882
}

func TaintStepTest_StrconvAppendQuoteToASCII_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte382` into `intoByte583`.

	// Assume that `sourceCQL` has the underlying type of `fromByte382`:
	fromByte382 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte382` to result `intoByte583`
	// (`intoByte583` is now tainted).
	intoByte583 := strconv.AppendQuoteToASCII(fromByte382, "")

	// Return the tainted `intoByte583`:
	return intoByte583
}

func TaintStepTest_StrconvAppendQuoteToASCII_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString752` into `intoByte454`.

	// Assume that `sourceCQL` has the underlying type of `fromString752`:
	fromString752 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString752` to result `intoByte454`
	// (`intoByte454` is now tainted).
	intoByte454 := strconv.AppendQuoteToASCII(nil, fromString752)

	// Return the tainted `intoByte454`:
	return intoByte454
}

func TaintStepTest_StrconvAppendQuoteToGraphic_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte475` into `intoByte192`.

	// Assume that `sourceCQL` has the underlying type of `fromByte475`:
	fromByte475 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte475` to result `intoByte192`
	// (`intoByte192` is now tainted).
	intoByte192 := strconv.AppendQuoteToGraphic(fromByte475, "")

	// Return the tainted `intoByte192`:
	return intoByte192
}

func TaintStepTest_StrconvAppendQuoteToGraphic_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString252` into `intoByte433`.

	// Assume that `sourceCQL` has the underlying type of `fromString252`:
	fromString252 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString252` to result `intoByte433`
	// (`intoByte433` is now tainted).
	intoByte433 := strconv.AppendQuoteToGraphic(nil, fromString252)

	// Return the tainted `intoByte433`:
	return intoByte433
}

func TaintStepTest_StrconvQuote_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString707` into `intoString186`.

	// Assume that `sourceCQL` has the underlying type of `fromString707`:
	fromString707 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString707` to result `intoString186`
	// (`intoString186` is now tainted).
	intoString186 := strconv.Quote(fromString707)

	// Return the tainted `intoString186`:
	return intoString186
}

func TaintStepTest_StrconvQuoteRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune417` into `intoString919`.

	// Assume that `sourceCQL` has the underlying type of `fromRune417`:
	fromRune417 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune417` to result `intoString919`
	// (`intoString919` is now tainted).
	intoString919 := strconv.QuoteRune(fromRune417)

	// Return the tainted `intoString919`:
	return intoString919
}

func TaintStepTest_StrconvQuoteRuneToASCII_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune994` into `intoString545`.

	// Assume that `sourceCQL` has the underlying type of `fromRune994`:
	fromRune994 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune994` to result `intoString545`
	// (`intoString545` is now tainted).
	intoString545 := strconv.QuoteRuneToASCII(fromRune994)

	// Return the tainted `intoString545`:
	return intoString545
}

func TaintStepTest_StrconvQuoteRuneToGraphic_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune667` into `intoString467`.

	// Assume that `sourceCQL` has the underlying type of `fromRune667`:
	fromRune667 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune667` to result `intoString467`
	// (`intoString467` is now tainted).
	intoString467 := strconv.QuoteRuneToGraphic(fromRune667)

	// Return the tainted `intoString467`:
	return intoString467
}

func TaintStepTest_StrconvQuoteToASCII_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString850` into `intoString593`.

	// Assume that `sourceCQL` has the underlying type of `fromString850`:
	fromString850 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString850` to result `intoString593`
	// (`intoString593` is now tainted).
	intoString593 := strconv.QuoteToASCII(fromString850)

	// Return the tainted `intoString593`:
	return intoString593
}

func TaintStepTest_StrconvQuoteToGraphic_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString321` into `intoString155`.

	// Assume that `sourceCQL` has the underlying type of `fromString321`:
	fromString321 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString321` to result `intoString155`
	// (`intoString155` is now tainted).
	intoString155 := strconv.QuoteToGraphic(fromString321)

	// Return the tainted `intoString155`:
	return intoString155
}

func TaintStepTest_StrconvUnquote_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString285` into `intoString115`.

	// Assume that `sourceCQL` has the underlying type of `fromString285`:
	fromString285 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString285` to result `intoString115`
	// (`intoString115` is now tainted).
	intoString115, _ := strconv.Unquote(fromString285)

	// Return the tainted `intoString115`:
	return intoString115
}

func TaintStepTest_StrconvUnquoteChar_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString303` into `intoRune961`.

	// Assume that `sourceCQL` has the underlying type of `fromString303`:
	fromString303 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString303` to result `intoRune961`
	// (`intoRune961` is now tainted).
	intoRune961, _, _, _ := strconv.UnquoteChar(fromString303, 0)

	// Return the tainted `intoRune961`:
	return intoRune961
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
