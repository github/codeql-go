package main

import "container/heap"

func TaintStepTest_ContainerHeapPop(sourceCQL interface{}) {
	// The flow is from `h` into `into198`.

	// Assume that `sourceCQL` has the underlying type of `h`:
	h := sourceCQL.(heap.Interface)

	// Call the function that transfers the taint
	// from the parameter `h` to result `into198`
	// (`into198` is now tainted).
	into198 := heap.Pop(h)

	// Sink the tainted `into198`:
	sink(into198)
}

func TaintStepTest_ContainerHeapPush(sourceCQL interface{}) {
	// The flow is from `x` into `h`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.(interface{})

	// Declare `h` variable:
	var h heap.Interface

	// Call the function that transfers the taint
	// from the parameter `x` to parameter `h`;
	// `h` is now tainted.
	heap.Push(h, x)

	// Sink the tainted `h`:
	sink(h)
}

func TaintStepTest_ContainerHeapRemove(sourceCQL interface{}) {
	// The flow is from `h` into `into432`.

	// Assume that `sourceCQL` has the underlying type of `h`:
	h := sourceCQL.(heap.Interface)

	// Call the function that transfers the taint
	// from the parameter `h` to result `into432`
	// (`into432` is now tainted).
	into432 := heap.Remove(h, 0)

	// Sink the tainted `into432`:
	sink(into432)
}

func TaintStepTest_ContainerHeapInterfacePop(sourceCQL interface{}) {
	// The flow is from `from179` into `into212`.

	// Assume that `sourceCQL` has the underlying type of `from179`:
	from179 := sourceCQL.(heap.Interface)

	// Call the method that transfers the taint
	// from the receiver `from179` to the result `into212`
	// (`into212` is now tainted).
	into212 := from179.Pop()

	// Sink the tainted `into212`:
	sink(into212)
}

func TaintStepTest_ContainerHeapInterfacePush(sourceCQL interface{}) {
	// The flow is from `x` into `into432`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.(interface{})

	// Declare `into432` variable:
	var into432 heap.Interface

	// Call the method that transfers the taint
	// from the parameter `x` to the receiver `into432`
	// (`into432` is now tainted).
	into432.Push(x)

	// Sink the tainted `into432`:
	sink(into432)
}

func RunAllTaints_ContainerHeap(v interface{}) {
	{
		source := newSource()
		TaintStepTest_ContainerHeapPop(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerHeapPush(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerHeapRemove(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerHeapInterfacePop(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerHeapInterfacePush(source)
	}
}
