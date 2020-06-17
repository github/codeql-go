// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "html"

func TaintStepTest_HtmlEscapeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString656` into `intoString414`.

	// Assume that `sourceCQL` has the underlying type of `fromString656`:
	fromString656 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString656` to result `intoString414`
	// (`intoString414` is now tainted).
	intoString414 := html.EscapeString(fromString656)

	// Return the tainted `intoString414`:
	return intoString414
}

func TaintStepTest_HtmlUnescapeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString518` into `intoString650`.

	// Assume that `sourceCQL` has the underlying type of `fromString518`:
	fromString518 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString518` to result `intoString650`
	// (`intoString650` is now tainted).
	intoString650 := html.UnescapeString(fromString518)

	// Return the tainted `intoString650`:
	return intoString650
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
