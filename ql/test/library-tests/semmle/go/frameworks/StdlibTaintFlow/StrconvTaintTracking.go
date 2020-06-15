package main

import "strconv"

func TaintStepTest_StrconvAppendQuote_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte526` into `intoByte436`.

	// Assume that `sourceCQL` has the underlying type of `fromByte526`:
	fromByte526 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte526` to result `intoByte436`
	// (`intoByte436` is now tainted).
	intoByte436 := strconv.AppendQuote(fromByte526, "")

	// Sink the tainted `intoByte436`:
	sink(intoByte436)
}

func TaintStepTest_StrconvAppendQuote_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString889` into `intoByte131`.

	// Assume that `sourceCQL` has the underlying type of `fromString889`:
	fromString889 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString889` to result `intoByte131`
	// (`intoByte131` is now tainted).
	intoByte131 := strconv.AppendQuote(nil, fromString889)

	// Sink the tainted `intoByte131`:
	sink(intoByte131)
}

func TaintStepTest_StrconvAppendQuoteRune_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte230` into `intoByte347`.

	// Assume that `sourceCQL` has the underlying type of `fromByte230`:
	fromByte230 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte230` to result `intoByte347`
	// (`intoByte347` is now tainted).
	intoByte347 := strconv.AppendQuoteRune(fromByte230, 0)

	// Sink the tainted `intoByte347`:
	sink(intoByte347)
}

func TaintStepTest_StrconvAppendQuoteRune_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromRune926` into `intoByte812`.

	// Assume that `sourceCQL` has the underlying type of `fromRune926`:
	fromRune926 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune926` to result `intoByte812`
	// (`intoByte812` is now tainted).
	intoByte812 := strconv.AppendQuoteRune(nil, fromRune926)

	// Sink the tainted `intoByte812`:
	sink(intoByte812)
}

func TaintStepTest_StrconvAppendQuoteRuneToASCII_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte366` into `intoByte435`.

	// Assume that `sourceCQL` has the underlying type of `fromByte366`:
	fromByte366 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte366` to result `intoByte435`
	// (`intoByte435` is now tainted).
	intoByte435 := strconv.AppendQuoteRuneToASCII(fromByte366, 0)

	// Sink the tainted `intoByte435`:
	sink(intoByte435)
}

func TaintStepTest_StrconvAppendQuoteRuneToASCII_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromRune297` into `intoByte268`.

	// Assume that `sourceCQL` has the underlying type of `fromRune297`:
	fromRune297 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune297` to result `intoByte268`
	// (`intoByte268` is now tainted).
	intoByte268 := strconv.AppendQuoteRuneToASCII(nil, fromRune297)

	// Sink the tainted `intoByte268`:
	sink(intoByte268)
}

func TaintStepTest_StrconvAppendQuoteRuneToGraphic_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte750` into `intoByte338`.

	// Assume that `sourceCQL` has the underlying type of `fromByte750`:
	fromByte750 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte750` to result `intoByte338`
	// (`intoByte338` is now tainted).
	intoByte338 := strconv.AppendQuoteRuneToGraphic(fromByte750, 0)

	// Sink the tainted `intoByte338`:
	sink(intoByte338)
}

func TaintStepTest_StrconvAppendQuoteRuneToGraphic_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromRune346` into `intoByte296`.

	// Assume that `sourceCQL` has the underlying type of `fromRune346`:
	fromRune346 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune346` to result `intoByte296`
	// (`intoByte296` is now tainted).
	intoByte296 := strconv.AppendQuoteRuneToGraphic(nil, fromRune346)

	// Sink the tainted `intoByte296`:
	sink(intoByte296)
}

func TaintStepTest_StrconvAppendQuoteToASCII_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte477` into `intoByte192`.

	// Assume that `sourceCQL` has the underlying type of `fromByte477`:
	fromByte477 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte477` to result `intoByte192`
	// (`intoByte192` is now tainted).
	intoByte192 := strconv.AppendQuoteToASCII(fromByte477, "")

	// Sink the tainted `intoByte192`:
	sink(intoByte192)
}

func TaintStepTest_StrconvAppendQuoteToASCII_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString927` into `intoByte801`.

	// Assume that `sourceCQL` has the underlying type of `fromString927`:
	fromString927 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString927` to result `intoByte801`
	// (`intoByte801` is now tainted).
	intoByte801 := strconv.AppendQuoteToASCII(nil, fromString927)

	// Sink the tainted `intoByte801`:
	sink(intoByte801)
}

func TaintStepTest_StrconvAppendQuoteToGraphic_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte868` into `intoByte235`.

	// Assume that `sourceCQL` has the underlying type of `fromByte868`:
	fromByte868 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte868` to result `intoByte235`
	// (`intoByte235` is now tainted).
	intoByte235 := strconv.AppendQuoteToGraphic(fromByte868, "")

	// Sink the tainted `intoByte235`:
	sink(intoByte235)
}

func TaintStepTest_StrconvAppendQuoteToGraphic_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString919` into `intoByte677`.

	// Assume that `sourceCQL` has the underlying type of `fromString919`:
	fromString919 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString919` to result `intoByte677`
	// (`intoByte677` is now tainted).
	intoByte677 := strconv.AppendQuoteToGraphic(nil, fromString919)

	// Sink the tainted `intoByte677`:
	sink(intoByte677)
}

func TaintStepTest_StrconvQuote_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString515` into `intoString470`.

	// Assume that `sourceCQL` has the underlying type of `fromString515`:
	fromString515 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString515` to result `intoString470`
	// (`intoString470` is now tainted).
	intoString470 := strconv.Quote(fromString515)

	// Sink the tainted `intoString470`:
	sink(intoString470)
}

func TaintStepTest_StrconvQuoteRune_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune210` into `intoString950`.

	// Assume that `sourceCQL` has the underlying type of `fromRune210`:
	fromRune210 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune210` to result `intoString950`
	// (`intoString950` is now tainted).
	intoString950 := strconv.QuoteRune(fromRune210)

	// Sink the tainted `intoString950`:
	sink(intoString950)
}

func TaintStepTest_StrconvQuoteRuneToASCII_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune267` into `intoString285`.

	// Assume that `sourceCQL` has the underlying type of `fromRune267`:
	fromRune267 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune267` to result `intoString285`
	// (`intoString285` is now tainted).
	intoString285 := strconv.QuoteRuneToASCII(fromRune267)

	// Sink the tainted `intoString285`:
	sink(intoString285)
}

func TaintStepTest_StrconvQuoteRuneToGraphic_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune458` into `intoString911`.

	// Assume that `sourceCQL` has the underlying type of `fromRune458`:
	fromRune458 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune458` to result `intoString911`
	// (`intoString911` is now tainted).
	intoString911 := strconv.QuoteRuneToGraphic(fromRune458)

	// Sink the tainted `intoString911`:
	sink(intoString911)
}

func TaintStepTest_StrconvQuoteToASCII_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString981` into `intoString859`.

	// Assume that `sourceCQL` has the underlying type of `fromString981`:
	fromString981 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString981` to result `intoString859`
	// (`intoString859` is now tainted).
	intoString859 := strconv.QuoteToASCII(fromString981)

	// Sink the tainted `intoString859`:
	sink(intoString859)
}

func TaintStepTest_StrconvQuoteToGraphic_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString854` into `intoString896`.

	// Assume that `sourceCQL` has the underlying type of `fromString854`:
	fromString854 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString854` to result `intoString896`
	// (`intoString896` is now tainted).
	intoString896 := strconv.QuoteToGraphic(fromString854)

	// Sink the tainted `intoString896`:
	sink(intoString896)
}

func TaintStepTest_StrconvUnquote_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString710` into `intoString951`.

	// Assume that `sourceCQL` has the underlying type of `fromString710`:
	fromString710 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString710` to result `intoString951`
	// (`intoString951` is now tainted).
	intoString951, _ := strconv.Unquote(fromString710)

	// Sink the tainted `intoString951`:
	sink(intoString951)
}

func TaintStepTest_StrconvUnquoteChar_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString163` into `intoRune609`.

	// Assume that `sourceCQL` has the underlying type of `fromString163`:
	fromString163 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString163` to result `intoRune609`
	// (`intoRune609` is now tainted).
	intoRune609, _, _, _ := strconv.UnquoteChar(fromString163, 0)

	// Sink the tainted `intoRune609`:
	sink(intoRune609)
}

func RunAllTaints_Strconv() {
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuote_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuote_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuoteRune_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuoteRune_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuoteRuneToASCII_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuoteRuneToASCII_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuoteRuneToGraphic_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuoteRuneToGraphic_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuoteToASCII_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuoteToASCII_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuoteToGraphic_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvAppendQuoteToGraphic_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvQuote_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvQuoteRune_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvQuoteRuneToASCII_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvQuoteRuneToGraphic_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvQuoteToASCII_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvQuoteToGraphic_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvUnquote_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StrconvUnquoteChar_B0I0O0(source)
	}
}
