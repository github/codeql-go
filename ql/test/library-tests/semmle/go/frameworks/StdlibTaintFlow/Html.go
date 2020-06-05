package main

import "html"

func TaintStepTest_HtmlEscapeString(sourceCQL interface{}) {
	// The flow is from `s` into `into751`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into751`
	// (`into751` is now tainted).
	into751 := html.EscapeString(s)

	// Sink the tainted `into751`:
	sink(into751)
}

func TaintStepTest_HtmlUnescapeString(sourceCQL interface{}) {
	// The flow is from `s` into `into455`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into455`
	// (`into455` is now tainted).
	into455 := html.UnescapeString(s)

	// Sink the tainted `into455`:
	sink(into455)
}

func RunAllTaints_Html(v interface{}) {
	{
		source := newSource()
		TaintStepTest_HtmlEscapeString(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlUnescapeString(source)
	}
}
