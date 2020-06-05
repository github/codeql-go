package main

import "strconv"

func TaintStepTest_StrconvAppendQuote(sourceCQL interface{}) {
	// The flow is from `dst` into `into797`.

	// Assume that `sourceCQL` has the underlying type of `dst`:
	dst := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `dst` to result `into797`
	// (`into797` is now tainted).
	into797 := strconv.AppendQuote(dst, "")

	// Sink the tainted `into797`:
	sink(into797)
}

func TaintStepTest_StrconvAppendQuoteRune(sourceCQL interface{}) {
	// The flow is from `dst` into `into364`.

	// Assume that `sourceCQL` has the underlying type of `dst`:
	dst := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `dst` to result `into364`
	// (`into364` is now tainted).
	into364 := strconv.AppendQuoteRune(dst, 0)

	// Sink the tainted `into364`:
	sink(into364)
}

func TaintStepTest_StrconvAppendQuoteRuneToASCII(sourceCQL interface{}) {
	// The flow is from `dst` into `into785`.

	// Assume that `sourceCQL` has the underlying type of `dst`:
	dst := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `dst` to result `into785`
	// (`into785` is now tainted).
	into785 := strconv.AppendQuoteRuneToASCII(dst, 0)

	// Sink the tainted `into785`:
	sink(into785)
}

func TaintStepTest_StrconvAppendQuoteRuneToGraphic(sourceCQL interface{}) {
	// The flow is from `dst` into `into989`.

	// Assume that `sourceCQL` has the underlying type of `dst`:
	dst := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `dst` to result `into989`
	// (`into989` is now tainted).
	into989 := strconv.AppendQuoteRuneToGraphic(dst, 0)

	// Sink the tainted `into989`:
	sink(into989)
}

func TaintStepTest_StrconvAppendQuoteToASCII(sourceCQL interface{}) {
	// The flow is from `dst` into `into220`.

	// Assume that `sourceCQL` has the underlying type of `dst`:
	dst := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `dst` to result `into220`
	// (`into220` is now tainted).
	into220 := strconv.AppendQuoteToASCII(dst, "")

	// Sink the tainted `into220`:
	sink(into220)
}

func TaintStepTest_StrconvAppendQuoteToGraphic(sourceCQL interface{}) {
	// The flow is from `dst` into `into449`.

	// Assume that `sourceCQL` has the underlying type of `dst`:
	dst := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `dst` to result `into449`
	// (`into449` is now tainted).
	into449 := strconv.AppendQuoteToGraphic(dst, "")

	// Sink the tainted `into449`:
	sink(into449)
}

func TaintStepTest_StrconvQuote(sourceCQL interface{}) {
	// The flow is from `s` into `into384`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into384`
	// (`into384` is now tainted).
	into384 := strconv.Quote(s)

	// Sink the tainted `into384`:
	sink(into384)
}

func TaintStepTest_StrconvQuoteRune(sourceCQL interface{}) {
	// The flow is from `r` into `into938`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into938`
	// (`into938` is now tainted).
	into938 := strconv.QuoteRune(r)

	// Sink the tainted `into938`:
	sink(into938)
}

func TaintStepTest_StrconvQuoteRuneToASCII(sourceCQL interface{}) {
	// The flow is from `r` into `into611`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into611`
	// (`into611` is now tainted).
	into611 := strconv.QuoteRuneToASCII(r)

	// Sink the tainted `into611`:
	sink(into611)
}

func TaintStepTest_StrconvQuoteRuneToGraphic(sourceCQL interface{}) {
	// The flow is from `r` into `into602`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into602`
	// (`into602` is now tainted).
	into602 := strconv.QuoteRuneToGraphic(r)

	// Sink the tainted `into602`:
	sink(into602)
}

func TaintStepTest_StrconvQuoteToASCII(sourceCQL interface{}) {
	// The flow is from `s` into `into219`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into219`
	// (`into219` is now tainted).
	into219 := strconv.QuoteToASCII(s)

	// Sink the tainted `into219`:
	sink(into219)
}

func TaintStepTest_StrconvQuoteToGraphic(sourceCQL interface{}) {
	// The flow is from `s` into `into203`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into203`
	// (`into203` is now tainted).
	into203 := strconv.QuoteToGraphic(s)

	// Sink the tainted `into203`:
	sink(into203)
}

func TaintStepTest_StrconvUnquote(sourceCQL interface{}) {
	// The flow is from `s` into `into858`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into858`
	// (`into858` is now tainted).
	into858, _ := strconv.Unquote(s)

	// Sink the tainted `into858`:
	sink(into858)
}

func TaintStepTest_StrconvUnquoteChar(sourceCQL interface{}) {
	// The flow is from `s` into `value`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `value`
	// (`value` is now tainted).
	value, _, _, _ := strconv.UnquoteChar(s, 0)

	// Sink the tainted `value`:
	sink(value)
}

func RunAllTaints_Strconv(v interface{}) {
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuote(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuoteRune(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuoteRuneToASCII(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuoteRuneToGraphic(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuoteToASCII(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuoteToGraphic(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvQuote(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvQuoteRune(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvQuoteRuneToASCII(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvQuoteRuneToGraphic(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvQuoteToASCII(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvQuoteToGraphic(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvUnquote(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvUnquoteChar(source)
	}
}
