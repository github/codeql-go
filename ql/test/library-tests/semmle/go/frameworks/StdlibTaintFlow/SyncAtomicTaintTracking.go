// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"sync/atomic"
	"unsafe"
)

func TaintStepTest_SyncAtomicAddUintptr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr656` into `intoUintptr414`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr656`:
	fromUintptr656 := sourceCQL.(uintptr)

	// Declare `intoUintptr414` variable:
	var intoUintptr414 *uintptr

	// Call the function that transfers the taint
	// from the parameter `fromUintptr656` to parameter `intoUintptr414`;
	// `intoUintptr414` is now tainted.
	atomic.AddUintptr(intoUintptr414, fromUintptr656)

	// Return the tainted `intoUintptr414`:
	return intoUintptr414
}

func TaintStepTest_SyncAtomicAddUintptr_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr518` into `intoUintptr650`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr518`:
	fromUintptr518 := sourceCQL.(uintptr)

	// Call the function that transfers the taint
	// from the parameter `fromUintptr518` to result `intoUintptr650`
	// (`intoUintptr650` is now tainted).
	intoUintptr650 := atomic.AddUintptr(nil, fromUintptr518)

	// Return the tainted `intoUintptr650`:
	return intoUintptr650
}

func TaintStepTest_SyncAtomicCompareAndSwapPointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer784` into `intoPointer957`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer784`:
	fromPointer784 := sourceCQL.(unsafe.Pointer)

	// Declare `intoPointer957` variable:
	var intoPointer957 *unsafe.Pointer

	// Call the function that transfers the taint
	// from the parameter `fromPointer784` to parameter `intoPointer957`;
	// `intoPointer957` is now tainted.
	atomic.CompareAndSwapPointer(intoPointer957, nil, fromPointer784)

	// Return the tainted `intoPointer957`:
	return intoPointer957
}

func TaintStepTest_SyncAtomicCompareAndSwapUintptr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr520` into `intoUintptr443`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr520`:
	fromUintptr520 := sourceCQL.(uintptr)

	// Declare `intoUintptr443` variable:
	var intoUintptr443 *uintptr

	// Call the function that transfers the taint
	// from the parameter `fromUintptr520` to parameter `intoUintptr443`;
	// `intoUintptr443` is now tainted.
	atomic.CompareAndSwapUintptr(intoUintptr443, 0, fromUintptr520)

	// Return the tainted `intoUintptr443`:
	return intoUintptr443
}

func TaintStepTest_SyncAtomicLoadPointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer127` into `intoPointer483`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer127`:
	fromPointer127 := sourceCQL.(*unsafe.Pointer)

	// Call the function that transfers the taint
	// from the parameter `fromPointer127` to result `intoPointer483`
	// (`intoPointer483` is now tainted).
	intoPointer483 := atomic.LoadPointer(fromPointer127)

	// Return the tainted `intoPointer483`:
	return intoPointer483
}

func TaintStepTest_SyncAtomicLoadUintptr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr989` into `intoUintptr982`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr989`:
	fromUintptr989 := sourceCQL.(*uintptr)

	// Call the function that transfers the taint
	// from the parameter `fromUintptr989` to result `intoUintptr982`
	// (`intoUintptr982` is now tainted).
	intoUintptr982 := atomic.LoadUintptr(fromUintptr989)

	// Return the tainted `intoUintptr982`:
	return intoUintptr982
}

func TaintStepTest_SyncAtomicStorePointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer417` into `intoPointer584`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer417`:
	fromPointer417 := sourceCQL.(unsafe.Pointer)

	// Declare `intoPointer584` variable:
	var intoPointer584 *unsafe.Pointer

	// Call the function that transfers the taint
	// from the parameter `fromPointer417` to parameter `intoPointer584`;
	// `intoPointer584` is now tainted.
	atomic.StorePointer(intoPointer584, fromPointer417)

	// Return the tainted `intoPointer584`:
	return intoPointer584
}

func TaintStepTest_SyncAtomicStoreUintptr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr991` into `intoUintptr881`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr991`:
	fromUintptr991 := sourceCQL.(uintptr)

	// Declare `intoUintptr881` variable:
	var intoUintptr881 *uintptr

	// Call the function that transfers the taint
	// from the parameter `fromUintptr991` to parameter `intoUintptr881`;
	// `intoUintptr881` is now tainted.
	atomic.StoreUintptr(intoUintptr881, fromUintptr991)

	// Return the tainted `intoUintptr881`:
	return intoUintptr881
}

func TaintStepTest_SyncAtomicSwapPointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer186` into `intoPointer284`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer186`:
	fromPointer186 := sourceCQL.(unsafe.Pointer)

	// Declare `intoPointer284` variable:
	var intoPointer284 *unsafe.Pointer

	// Call the function that transfers the taint
	// from the parameter `fromPointer186` to parameter `intoPointer284`;
	// `intoPointer284` is now tainted.
	atomic.SwapPointer(intoPointer284, fromPointer186)

	// Return the tainted `intoPointer284`:
	return intoPointer284
}

func TaintStepTest_SyncAtomicSwapPointer_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer908` into `intoPointer137`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer908`:
	fromPointer908 := sourceCQL.(*unsafe.Pointer)

	// Call the function that transfers the taint
	// from the parameter `fromPointer908` to result `intoPointer137`
	// (`intoPointer137` is now tainted).
	intoPointer137 := atomic.SwapPointer(fromPointer908, nil)

	// Return the tainted `intoPointer137`:
	return intoPointer137
}

func TaintStepTest_SyncAtomicSwapUintptr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr494` into `intoUintptr873`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr494`:
	fromUintptr494 := sourceCQL.(uintptr)

	// Declare `intoUintptr873` variable:
	var intoUintptr873 *uintptr

	// Call the function that transfers the taint
	// from the parameter `fromUintptr494` to parameter `intoUintptr873`;
	// `intoUintptr873` is now tainted.
	atomic.SwapUintptr(intoUintptr873, fromUintptr494)

	// Return the tainted `intoUintptr873`:
	return intoUintptr873
}

func TaintStepTest_SyncAtomicSwapUintptr_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr599` into `intoUintptr409`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr599`:
	fromUintptr599 := sourceCQL.(*uintptr)

	// Call the function that transfers the taint
	// from the parameter `fromUintptr599` to result `intoUintptr409`
	// (`intoUintptr409` is now tainted).
	intoUintptr409 := atomic.SwapUintptr(fromUintptr599, 0)

	// Return the tainted `intoUintptr409`:
	return intoUintptr409
}

func TaintStepTest_SyncAtomicValueLoad_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue246` into `intoInterface898`.

	// Assume that `sourceCQL` has the underlying type of `fromValue246`:
	fromValue246 := sourceCQL.(atomic.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue246` to the result `intoInterface898`
	// (`intoInterface898` is now tainted).
	intoInterface898 := fromValue246.Load()

	// Return the tainted `intoInterface898`:
	return intoInterface898
}

func TaintStepTest_SyncAtomicValueStore_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface598` into `intoValue631`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface598`:
	fromInterface598 := sourceCQL.(interface{})

	// Declare `intoValue631` variable:
	var intoValue631 atomic.Value

	// Call the method that transfers the taint
	// from the parameter `fromInterface598` to the receiver `intoValue631`
	// (`intoValue631` is now tainted).
	intoValue631.Store(fromInterface598)

	// Return the tainted `intoValue631`:
	return intoValue631
}

func RunAllTaints_SyncAtomic() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncAtomicAddUintptr_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncAtomicAddUintptr_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncAtomicCompareAndSwapPointer_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncAtomicCompareAndSwapUintptr_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncAtomicLoadPointer_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncAtomicLoadUintptr_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncAtomicStorePointer_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncAtomicStoreUintptr_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncAtomicSwapPointer_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncAtomicSwapPointer_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncAtomicSwapUintptr_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncAtomicSwapUintptr_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncAtomicValueLoad_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncAtomicValueStore_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
