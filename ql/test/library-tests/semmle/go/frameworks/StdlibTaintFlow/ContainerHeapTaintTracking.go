// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "container/heap"

func TaintStepTest_ContainerHeapPop_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface656` into `intoInterface414`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface656`:
	fromInterface656 := sourceCQL.(heap.Interface)

	// Call the function that transfers the taint
	// from the parameter `fromInterface656` to result `intoInterface414`
	// (`intoInterface414` is now tainted).
	intoInterface414 := heap.Pop(fromInterface656)

	// Return the tainted `intoInterface414`:
	return intoInterface414
}

func TaintStepTest_ContainerHeapPush_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface518` into `intoInterface650`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface518`:
	fromInterface518 := sourceCQL.(interface{})

	// Declare `intoInterface650` variable:
	var intoInterface650 heap.Interface

	// Call the function that transfers the taint
	// from the parameter `fromInterface518` to parameter `intoInterface650`;
	// `intoInterface650` is now tainted.
	heap.Push(intoInterface650, fromInterface518)

	// Return the tainted `intoInterface650`:
	return intoInterface650
}

func TaintStepTest_ContainerHeapRemove_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface784` into `intoInterface957`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface784`:
	fromInterface784 := sourceCQL.(heap.Interface)

	// Call the function that transfers the taint
	// from the parameter `fromInterface784` to result `intoInterface957`
	// (`intoInterface957` is now tainted).
	intoInterface957 := heap.Remove(fromInterface784, 0)

	// Return the tainted `intoInterface957`:
	return intoInterface957
}

func TaintStepTest_ContainerHeapInterfacePop_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface520` into `intoInterface443`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface520`:
	fromInterface520 := sourceCQL.(heap.Interface)

	// Call the method that transfers the taint
	// from the receiver `fromInterface520` to the result `intoInterface443`
	// (`intoInterface443` is now tainted).
	intoInterface443 := fromInterface520.Pop()

	// Return the tainted `intoInterface443`:
	return intoInterface443
}

func TaintStepTest_ContainerHeapInterfacePush_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface127` into `intoInterface483`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface127`:
	fromInterface127 := sourceCQL.(interface{})

	// Declare `intoInterface483` variable:
	var intoInterface483 heap.Interface

	// Call the method that transfers the taint
	// from the parameter `fromInterface127` to the receiver `intoInterface483`
	// (`intoInterface483` is now tainted).
	intoInterface483.Push(fromInterface127)

	// Return the tainted `intoInterface483`:
	return intoInterface483
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
