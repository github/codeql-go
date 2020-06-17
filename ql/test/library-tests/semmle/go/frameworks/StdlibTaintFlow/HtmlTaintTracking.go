package main

import "html"

func TaintStepTest_HtmlEscapeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString894` into `intoString347`.

	// Assume that `sourceCQL` has the underlying type of `fromString894`:
	fromString894 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString894` to result `intoString347`
	// (`intoString347` is now tainted).
	intoString347 := html.EscapeString(fromString894)

	// Return the tainted `intoString347`:
	return intoString347
}

func TaintStepTest_HtmlUnescapeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString801` into `intoString860`.

	// Assume that `sourceCQL` has the underlying type of `fromString801`:
	fromString801 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString801` to result `intoString860`
	// (`intoString860` is now tainted).
	intoString860 := html.UnescapeString(fromString801)

	// Return the tainted `intoString860`:
	return intoString860
}

func RunAllTaints_Html() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlEscapeString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlUnescapeString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
