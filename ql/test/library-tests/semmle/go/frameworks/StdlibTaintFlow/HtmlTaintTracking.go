// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "html"

func TaintStepTest_HtmlEscapeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString761` into `intoString784`.

	// Assume that `sourceCQL` has the underlying type of `fromString761`:
	fromString761 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString761` to result `intoString784`
	// (`intoString784` is now tainted).
	intoString784 := html.EscapeString(fromString761)

	// Return the tainted `intoString784`:
	return intoString784
}

func TaintStepTest_HtmlUnescapeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString394` into `intoString703`.

	// Assume that `sourceCQL` has the underlying type of `fromString394`:
	fromString394 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString394` to result `intoString703`
	// (`intoString703` is now tainted).
	intoString703 := html.UnescapeString(fromString394)

	// Return the tainted `intoString703`:
	return intoString703
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
