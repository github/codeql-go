package main

import "html"

func TaintStepTest_HtmlEscapeString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString173` into `intoString602`.

	// Assume that `sourceCQL` has the underlying type of `fromString173`:
	fromString173 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString173` to result `intoString602`
	// (`intoString602` is now tainted).
	intoString602 := html.EscapeString(fromString173)

	// Sink the tainted `intoString602`:
	sink(intoString602)
}

func TaintStepTest_HtmlUnescapeString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString652` into `intoString712`.

	// Assume that `sourceCQL` has the underlying type of `fromString652`:
	fromString652 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString652` to result `intoString712`
	// (`intoString712` is now tainted).
	intoString712 := html.UnescapeString(fromString652)

	// Sink the tainted `intoString712`:
	sink(intoString712)
}

func RunAllTaints_Html() {
	{
		source := newSource()
		TaintStepTest_HtmlEscapeString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlUnescapeString_B0I0O0(source)
	}
}
