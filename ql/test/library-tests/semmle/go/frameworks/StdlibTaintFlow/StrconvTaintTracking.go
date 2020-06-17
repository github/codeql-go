package main

import "strconv"

func TaintStepTest_StrconvAppendQuote_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte277` into `intoByte472`.

	// Assume that `sourceCQL` has the underlying type of `fromByte277`:
	fromByte277 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte277` to result `intoByte472`
	// (`intoByte472` is now tainted).
	intoByte472 := strconv.AppendQuote(fromByte277, "")

	// Return the tainted `intoByte472`:
	return intoByte472
}

func TaintStepTest_StrconvAppendQuote_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString664` into `intoByte959`.

	// Assume that `sourceCQL` has the underlying type of `fromString664`:
	fromString664 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString664` to result `intoByte959`
	// (`intoByte959` is now tainted).
	intoByte959 := strconv.AppendQuote(nil, fromString664)

	// Return the tainted `intoByte959`:
	return intoByte959
}

func TaintStepTest_StrconvAppendQuoteRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte719` into `intoByte497`.

	// Assume that `sourceCQL` has the underlying type of `fromByte719`:
	fromByte719 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte719` to result `intoByte497`
	// (`intoByte497` is now tainted).
	intoByte497 := strconv.AppendQuoteRune(fromByte719, 0)

	// Return the tainted `intoByte497`:
	return intoByte497
}

func TaintStepTest_StrconvAppendQuoteRune_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune990` into `intoByte236`.

	// Assume that `sourceCQL` has the underlying type of `fromRune990`:
	fromRune990 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune990` to result `intoByte236`
	// (`intoByte236` is now tainted).
	intoByte236 := strconv.AppendQuoteRune(nil, fromRune990)

	// Return the tainted `intoByte236`:
	return intoByte236
}

func TaintStepTest_StrconvAppendQuoteRuneToASCII_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte290` into `intoByte921`.

	// Assume that `sourceCQL` has the underlying type of `fromByte290`:
	fromByte290 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte290` to result `intoByte921`
	// (`intoByte921` is now tainted).
	intoByte921 := strconv.AppendQuoteRuneToASCII(fromByte290, 0)

	// Return the tainted `intoByte921`:
	return intoByte921
}

func TaintStepTest_StrconvAppendQuoteRuneToASCII_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune230` into `intoByte437`.

	// Assume that `sourceCQL` has the underlying type of `fromRune230`:
	fromRune230 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune230` to result `intoByte437`
	// (`intoByte437` is now tainted).
	intoByte437 := strconv.AppendQuoteRuneToASCII(nil, fromRune230)

	// Return the tainted `intoByte437`:
	return intoByte437
}

func TaintStepTest_StrconvAppendQuoteRuneToGraphic_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte447` into `intoByte433`.

	// Assume that `sourceCQL` has the underlying type of `fromByte447`:
	fromByte447 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte447` to result `intoByte433`
	// (`intoByte433` is now tainted).
	intoByte433 := strconv.AppendQuoteRuneToGraphic(fromByte447, 0)

	// Return the tainted `intoByte433`:
	return intoByte433
}

func TaintStepTest_StrconvAppendQuoteRuneToGraphic_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune465` into `intoByte773`.

	// Assume that `sourceCQL` has the underlying type of `fromRune465`:
	fromRune465 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune465` to result `intoByte773`
	// (`intoByte773` is now tainted).
	intoByte773 := strconv.AppendQuoteRuneToGraphic(nil, fromRune465)

	// Return the tainted `intoByte773`:
	return intoByte773
}

func TaintStepTest_StrconvAppendQuoteToASCII_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte692` into `intoByte825`.

	// Assume that `sourceCQL` has the underlying type of `fromByte692`:
	fromByte692 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte692` to result `intoByte825`
	// (`intoByte825` is now tainted).
	intoByte825 := strconv.AppendQuoteToASCII(fromByte692, "")

	// Return the tainted `intoByte825`:
	return intoByte825
}

func TaintStepTest_StrconvAppendQuoteToASCII_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString963` into `intoByte542`.

	// Assume that `sourceCQL` has the underlying type of `fromString963`:
	fromString963 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString963` to result `intoByte542`
	// (`intoByte542` is now tainted).
	intoByte542 := strconv.AppendQuoteToASCII(nil, fromString963)

	// Return the tainted `intoByte542`:
	return intoByte542
}

func TaintStepTest_StrconvAppendQuoteToGraphic_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte450` into `intoByte902`.

	// Assume that `sourceCQL` has the underlying type of `fromByte450`:
	fromByte450 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte450` to result `intoByte902`
	// (`intoByte902` is now tainted).
	intoByte902 := strconv.AppendQuoteToGraphic(fromByte450, "")

	// Return the tainted `intoByte902`:
	return intoByte902
}

func TaintStepTest_StrconvAppendQuoteToGraphic_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString969` into `intoByte845`.

	// Assume that `sourceCQL` has the underlying type of `fromString969`:
	fromString969 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString969` to result `intoByte845`
	// (`intoByte845` is now tainted).
	intoByte845 := strconv.AppendQuoteToGraphic(nil, fromString969)

	// Return the tainted `intoByte845`:
	return intoByte845
}

func TaintStepTest_StrconvQuote_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString336` into `intoString827`.

	// Assume that `sourceCQL` has the underlying type of `fromString336`:
	fromString336 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString336` to result `intoString827`
	// (`intoString827` is now tainted).
	intoString827 := strconv.Quote(fromString336)

	// Return the tainted `intoString827`:
	return intoString827
}

func TaintStepTest_StrconvQuoteRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune289` into `intoString693`.

	// Assume that `sourceCQL` has the underlying type of `fromRune289`:
	fromRune289 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune289` to result `intoString693`
	// (`intoString693` is now tainted).
	intoString693 := strconv.QuoteRune(fromRune289)

	// Return the tainted `intoString693`:
	return intoString693
}

func TaintStepTest_StrconvQuoteRuneToASCII_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune505` into `intoString373`.

	// Assume that `sourceCQL` has the underlying type of `fromRune505`:
	fromRune505 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune505` to result `intoString373`
	// (`intoString373` is now tainted).
	intoString373 := strconv.QuoteRuneToASCII(fromRune505)

	// Return the tainted `intoString373`:
	return intoString373
}

func TaintStepTest_StrconvQuoteRuneToGraphic_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune399` into `intoString802`.

	// Assume that `sourceCQL` has the underlying type of `fromRune399`:
	fromRune399 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune399` to result `intoString802`
	// (`intoString802` is now tainted).
	intoString802 := strconv.QuoteRuneToGraphic(fromRune399)

	// Return the tainted `intoString802`:
	return intoString802
}

func TaintStepTest_StrconvQuoteToASCII_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString283` into `intoString909`.

	// Assume that `sourceCQL` has the underlying type of `fromString283`:
	fromString283 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString283` to result `intoString909`
	// (`intoString909` is now tainted).
	intoString909 := strconv.QuoteToASCII(fromString283)

	// Return the tainted `intoString909`:
	return intoString909
}

func TaintStepTest_StrconvQuoteToGraphic_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString388` into `intoString406`.

	// Assume that `sourceCQL` has the underlying type of `fromString388`:
	fromString388 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString388` to result `intoString406`
	// (`intoString406` is now tainted).
	intoString406 := strconv.QuoteToGraphic(fromString388)

	// Return the tainted `intoString406`:
	return intoString406
}

func TaintStepTest_StrconvUnquote_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString541` into `intoString165`.

	// Assume that `sourceCQL` has the underlying type of `fromString541`:
	fromString541 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString541` to result `intoString165`
	// (`intoString165` is now tainted).
	intoString165, _ := strconv.Unquote(fromString541)

	// Return the tainted `intoString165`:
	return intoString165
}

func TaintStepTest_StrconvUnquoteChar_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString911` into `intoRune561`.

	// Assume that `sourceCQL` has the underlying type of `fromString911`:
	fromString911 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString911` to result `intoRune561`
	// (`intoRune561` is now tainted).
	intoRune561, _, _, _ := strconv.UnquoteChar(fromString911, 0)

	// Return the tainted `intoRune561`:
	return intoRune561
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
