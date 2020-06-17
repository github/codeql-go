// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "container/heap"

func TaintStepTest_ContainerHeapPop_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface767` into `intoInterface725`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface767`:
	fromInterface767 := sourceCQL.(heap.Interface)

	// Call the function that transfers the taint
	// from the parameter `fromInterface767` to result `intoInterface725`
	// (`intoInterface725` is now tainted).
	intoInterface725 := heap.Pop(fromInterface767)

	// Return the tainted `intoInterface725`:
	return intoInterface725
}

func TaintStepTest_ContainerHeapPush_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface907` into `intoInterface414`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface907`:
	fromInterface907 := sourceCQL.(interface{})

	// Declare `intoInterface414` variable:
	var intoInterface414 heap.Interface

	// Call the function that transfers the taint
	// from the parameter `fromInterface907` to parameter `intoInterface414`;
	// `intoInterface414` is now tainted.
	heap.Push(intoInterface414, fromInterface907)

	// Return the tainted `intoInterface414`:
	return intoInterface414
}

func TaintStepTest_ContainerHeapRemove_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface569` into `intoInterface697`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface569`:
	fromInterface569 := sourceCQL.(heap.Interface)

	// Call the function that transfers the taint
	// from the parameter `fromInterface569` to result `intoInterface697`
	// (`intoInterface697` is now tainted).
	intoInterface697 := heap.Remove(fromInterface569, 0)

	// Return the tainted `intoInterface697`:
	return intoInterface697
}

func TaintStepTest_ContainerHeapInterfacePop_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface814` into `intoInterface646`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface814`:
	fromInterface814 := sourceCQL.(heap.Interface)

	// Call the method that transfers the taint
	// from the receiver `fromInterface814` to the result `intoInterface646`
	// (`intoInterface646` is now tainted).
	intoInterface646 := fromInterface814.Pop()

	// Return the tainted `intoInterface646`:
	return intoInterface646
}

func TaintStepTest_ContainerHeapInterfacePush_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface159` into `intoInterface843`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface159`:
	fromInterface159 := sourceCQL.(interface{})

	// Declare `intoInterface843` variable:
	var intoInterface843 heap.Interface

	// Call the method that transfers the taint
	// from the parameter `fromInterface159` to the receiver `intoInterface843`
	// (`intoInterface843` is now tainted).
	intoInterface843.Push(fromInterface159)

	// Return the tainted `intoInterface843`:
	return intoInterface843
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
