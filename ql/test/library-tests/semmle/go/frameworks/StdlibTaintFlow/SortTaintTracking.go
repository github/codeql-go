package main

import "sort"

func TaintStepTest_SortReverse_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface657` into `intoInterface279`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface657`:
	fromInterface657 := sourceCQL.(sort.Interface)

	// Call the function that transfers the taint
	// from the parameter `fromInterface657` to result `intoInterface279`
	// (`intoInterface279` is now tainted).
	intoInterface279 := sort.Reverse(fromInterface657)

	// Sink the tainted `intoInterface279`:
	sink(intoInterface279)
}

func RunAllTaints_Sort() {
	{
		source := newSource()
		TaintStepTest_SortReverse_B0I0O0(source)
	}
}
