// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "sort"

func TaintStepTest_SortReverse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface656` into `intoInterface414`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface656`:
	fromInterface656 := sourceCQL.(sort.Interface)

	// Call the function that transfers the taint
	// from the parameter `fromInterface656` to result `intoInterface414`
	// (`intoInterface414` is now tainted).
	intoInterface414 := sort.Reverse(fromInterface656)

	// Return the tainted `intoInterface414`:
	return intoInterface414
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
