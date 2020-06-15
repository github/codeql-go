package main

import "container/heap"

func TaintStepTest_ContainerHeapPop_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface687` into `intoInterface340`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface687`:
	fromInterface687 := sourceCQL.(heap.Interface)

	// Call the function that transfers the taint
	// from the parameter `fromInterface687` to result `intoInterface340`
	// (`intoInterface340` is now tainted).
	intoInterface340 := heap.Pop(fromInterface687)

	// Sink the tainted `intoInterface340`:
	sink(intoInterface340)
}

func TaintStepTest_ContainerHeapPush_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface174` into `intoInterface271`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface174`:
	fromInterface174 := sourceCQL.(interface{})

	// Declare `intoInterface271` variable:
	var intoInterface271 heap.Interface

	// Call the function that transfers the taint
	// from the parameter `fromInterface174` to parameter `intoInterface271`;
	// `intoInterface271` is now tainted.
	heap.Push(intoInterface271, fromInterface174)

	// Sink the tainted `intoInterface271`:
	sink(intoInterface271)
}

func TaintStepTest_ContainerHeapRemove_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface415` into `intoInterface843`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface415`:
	fromInterface415 := sourceCQL.(heap.Interface)

	// Call the function that transfers the taint
	// from the parameter `fromInterface415` to result `intoInterface843`
	// (`intoInterface843` is now tainted).
	intoInterface843 := heap.Remove(fromInterface415, 0)

	// Sink the tainted `intoInterface843`:
	sink(intoInterface843)
}

func TaintStepTest_ContainerHeapInterfacePop_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface940` into `intoInterface932`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface940`:
	fromInterface940 := sourceCQL.(heap.Interface)

	// Call the method that transfers the taint
	// from the receiver `fromInterface940` to the result `intoInterface932`
	// (`intoInterface932` is now tainted).
	intoInterface932 := fromInterface940.Pop()

	// Sink the tainted `intoInterface932`:
	sink(intoInterface932)
}

func TaintStepTest_ContainerHeapInterfacePush_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface741` into `intoInterface237`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface741`:
	fromInterface741 := sourceCQL.(interface{})

	// Declare `intoInterface237` variable:
	var intoInterface237 heap.Interface

	// Call the method that transfers the taint
	// from the parameter `fromInterface741` to the receiver `intoInterface237`
	// (`intoInterface237` is now tainted).
	intoInterface237.Push(fromInterface741)

	// Sink the tainted `intoInterface237`:
	sink(intoInterface237)
}

func RunAllTaints_ContainerHeap() {
	{
		source := newSource()
		TaintStepTest_ContainerHeapPop_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerHeapPush_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerHeapRemove_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerHeapInterfacePop_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerHeapInterfacePush_B0I0O0(source)
	}
}
