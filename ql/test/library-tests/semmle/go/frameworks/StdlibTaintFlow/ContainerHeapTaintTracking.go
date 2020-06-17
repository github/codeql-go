package main

import "container/heap"

func TaintStepTest_ContainerHeapPop_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface537` into `intoInterface902`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface537`:
	fromInterface537 := sourceCQL.(heap.Interface)

	// Call the function that transfers the taint
	// from the parameter `fromInterface537` to result `intoInterface902`
	// (`intoInterface902` is now tainted).
	intoInterface902 := heap.Pop(fromInterface537)

	// Return the tainted `intoInterface902`:
	return intoInterface902
}

func TaintStepTest_ContainerHeapPush_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface216` into `intoInterface503`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface216`:
	fromInterface216 := sourceCQL.(interface{})

	// Declare `intoInterface503` variable:
	var intoInterface503 heap.Interface

	// Call the function that transfers the taint
	// from the parameter `fromInterface216` to parameter `intoInterface503`;
	// `intoInterface503` is now tainted.
	heap.Push(intoInterface503, fromInterface216)

	// Return the tainted `intoInterface503`:
	return intoInterface503
}

func TaintStepTest_ContainerHeapRemove_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface608` into `intoInterface503`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface608`:
	fromInterface608 := sourceCQL.(heap.Interface)

	// Call the function that transfers the taint
	// from the parameter `fromInterface608` to result `intoInterface503`
	// (`intoInterface503` is now tainted).
	intoInterface503 := heap.Remove(fromInterface608, 0)

	// Return the tainted `intoInterface503`:
	return intoInterface503
}

func TaintStepTest_ContainerHeapInterfacePop_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface845` into `intoInterface669`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface845`:
	fromInterface845 := sourceCQL.(heap.Interface)

	// Call the method that transfers the taint
	// from the receiver `fromInterface845` to the result `intoInterface669`
	// (`intoInterface669` is now tainted).
	intoInterface669 := fromInterface845.Pop()

	// Return the tainted `intoInterface669`:
	return intoInterface669
}

func TaintStepTest_ContainerHeapInterfacePush_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface477` into `intoInterface622`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface477`:
	fromInterface477 := sourceCQL.(interface{})

	// Declare `intoInterface622` variable:
	var intoInterface622 heap.Interface

	// Call the method that transfers the taint
	// from the parameter `fromInterface477` to the receiver `intoInterface622`
	// (`intoInterface622` is now tainted).
	intoInterface622.Push(fromInterface477)

	// Return the tainted `intoInterface622`:
	return intoInterface622
}

func RunAllTaints_ContainerHeap() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerHeapPop_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerHeapPush_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerHeapRemove_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerHeapInterfacePop_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerHeapInterfacePush_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
