// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "sort"

func TaintStepTest_SortReverse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface646` into `intoInterface488`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface646`:
	fromInterface646 := sourceCQL.(sort.Interface)

	// Call the function that transfers the taint
	// from the parameter `fromInterface646` to result `intoInterface488`
	// (`intoInterface488` is now tainted).
	intoInterface488 := sort.Reverse(fromInterface646)

	// Return the tainted `intoInterface488`:
	return intoInterface488
}

func RunAllTaints_Sort() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SortReverse_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
