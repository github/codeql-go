package main

import "sort"

func TaintStepTest_SortReverse(sourceCQL interface{}) {
	// The flow is from `data` into `into211`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.(sort.Interface)

	// Call the function that transfers the taint
	// from the parameter `data` to result `into211`
	// (`into211` is now tainted).
	into211 := sort.Reverse(data)

	// Sink the tainted `into211`:
	sink(into211)
}

func RunAllTaints_Sort(v interface{}) {
	{
		source := newSource()
		TaintStepTest_SortReverse(source)
	}
}
