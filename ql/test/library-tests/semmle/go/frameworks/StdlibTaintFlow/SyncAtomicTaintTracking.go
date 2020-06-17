// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"sync/atomic"
	"unsafe"
)

func TaintStepTest_SyncAtomicAddUintptr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr883` into `intoUintptr239`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr883`:
	fromUintptr883 := sourceCQL.(uintptr)

	// Declare `intoUintptr239` variable:
	var intoUintptr239 *uintptr

	// Call the function that transfers the taint
	// from the parameter `fromUintptr883` to parameter `intoUintptr239`;
	// `intoUintptr239` is now tainted.
	atomic.AddUintptr(intoUintptr239, fromUintptr883)

	// Return the tainted `intoUintptr239`:
	return intoUintptr239
}

func TaintStepTest_SyncAtomicAddUintptr_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr900` into `intoUintptr565`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr900`:
	fromUintptr900 := sourceCQL.(uintptr)

	// Call the function that transfers the taint
	// from the parameter `fromUintptr900` to result `intoUintptr565`
	// (`intoUintptr565` is now tainted).
	intoUintptr565 := atomic.AddUintptr(nil, fromUintptr900)

	// Return the tainted `intoUintptr565`:
	return intoUintptr565
}

func TaintStepTest_SyncAtomicCompareAndSwapPointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer503` into `intoPointer259`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer503`:
	fromPointer503 := sourceCQL.(unsafe.Pointer)

	// Declare `intoPointer259` variable:
	var intoPointer259 *unsafe.Pointer

	// Call the function that transfers the taint
	// from the parameter `fromPointer503` to parameter `intoPointer259`;
	// `intoPointer259` is now tainted.
	atomic.CompareAndSwapPointer(intoPointer259, nil, fromPointer503)

	// Return the tainted `intoPointer259`:
	return intoPointer259
}

func TaintStepTest_SyncAtomicCompareAndSwapUintptr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr680` into `intoUintptr571`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr680`:
	fromUintptr680 := sourceCQL.(uintptr)

	// Declare `intoUintptr571` variable:
	var intoUintptr571 *uintptr

	// Call the function that transfers the taint
	// from the parameter `fromUintptr680` to parameter `intoUintptr571`;
	// `intoUintptr571` is now tainted.
	atomic.CompareAndSwapUintptr(intoUintptr571, 0, fromUintptr680)

	// Return the tainted `intoUintptr571`:
	return intoUintptr571
}

func TaintStepTest_SyncAtomicLoadPointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer323` into `intoPointer225`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer323`:
	fromPointer323 := sourceCQL.(*unsafe.Pointer)

	// Call the function that transfers the taint
	// from the parameter `fromPointer323` to result `intoPointer225`
	// (`intoPointer225` is now tainted).
	intoPointer225 := atomic.LoadPointer(fromPointer323)

	// Return the tainted `intoPointer225`:
	return intoPointer225
}

func TaintStepTest_SyncAtomicLoadUintptr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr361` into `intoUintptr665`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr361`:
	fromUintptr361 := sourceCQL.(*uintptr)

	// Call the function that transfers the taint
	// from the parameter `fromUintptr361` to result `intoUintptr665`
	// (`intoUintptr665` is now tainted).
	intoUintptr665 := atomic.LoadUintptr(fromUintptr361)

	// Return the tainted `intoUintptr665`:
	return intoUintptr665
}

func TaintStepTest_SyncAtomicStorePointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer385` into `intoPointer858`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer385`:
	fromPointer385 := sourceCQL.(unsafe.Pointer)

	// Declare `intoPointer858` variable:
	var intoPointer858 *unsafe.Pointer

	// Call the function that transfers the taint
	// from the parameter `fromPointer385` to parameter `intoPointer858`;
	// `intoPointer858` is now tainted.
	atomic.StorePointer(intoPointer858, fromPointer385)

	// Return the tainted `intoPointer858`:
	return intoPointer858
}

func TaintStepTest_SyncAtomicStoreUintptr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr592` into `intoUintptr184`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr592`:
	fromUintptr592 := sourceCQL.(uintptr)

	// Declare `intoUintptr184` variable:
	var intoUintptr184 *uintptr

	// Call the function that transfers the taint
	// from the parameter `fromUintptr592` to parameter `intoUintptr184`;
	// `intoUintptr184` is now tainted.
	atomic.StoreUintptr(intoUintptr184, fromUintptr592)

	// Return the tainted `intoUintptr184`:
	return intoUintptr184
}

func TaintStepTest_SyncAtomicSwapPointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer722` into `intoPointer591`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer722`:
	fromPointer722 := sourceCQL.(unsafe.Pointer)

	// Declare `intoPointer591` variable:
	var intoPointer591 *unsafe.Pointer

	// Call the function that transfers the taint
	// from the parameter `fromPointer722` to parameter `intoPointer591`;
	// `intoPointer591` is now tainted.
	atomic.SwapPointer(intoPointer591, fromPointer722)

	// Return the tainted `intoPointer591`:
	return intoPointer591
}

func TaintStepTest_SyncAtomicSwapPointer_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer703` into `intoPointer341`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer703`:
	fromPointer703 := sourceCQL.(*unsafe.Pointer)

	// Call the function that transfers the taint
	// from the parameter `fromPointer703` to result `intoPointer341`
	// (`intoPointer341` is now tainted).
	intoPointer341 := atomic.SwapPointer(fromPointer703, nil)

	// Return the tainted `intoPointer341`:
	return intoPointer341
}

func TaintStepTest_SyncAtomicSwapUintptr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr550` into `intoUintptr395`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr550`:
	fromUintptr550 := sourceCQL.(uintptr)

	// Declare `intoUintptr395` variable:
	var intoUintptr395 *uintptr

	// Call the function that transfers the taint
	// from the parameter `fromUintptr550` to parameter `intoUintptr395`;
	// `intoUintptr395` is now tainted.
	atomic.SwapUintptr(intoUintptr395, fromUintptr550)

	// Return the tainted `intoUintptr395`:
	return intoUintptr395
}

func TaintStepTest_SyncAtomicSwapUintptr_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr432` into `intoUintptr491`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr432`:
	fromUintptr432 := sourceCQL.(*uintptr)

	// Call the function that transfers the taint
	// from the parameter `fromUintptr432` to result `intoUintptr491`
	// (`intoUintptr491` is now tainted).
	intoUintptr491 := atomic.SwapUintptr(fromUintptr432, 0)

	// Return the tainted `intoUintptr491`:
	return intoUintptr491
}

func TaintStepTest_SyncAtomicValueLoad_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue514` into `intoInterface757`.

	// Assume that `sourceCQL` has the underlying type of `fromValue514`:
	fromValue514 := sourceCQL.(atomic.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue514` to the result `intoInterface757`
	// (`intoInterface757` is now tainted).
	intoInterface757 := fromValue514.Load()

	// Return the tainted `intoInterface757`:
	return intoInterface757
}

func TaintStepTest_SyncAtomicValueStore_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface489` into `intoValue985`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface489`:
	fromInterface489 := sourceCQL.(interface{})

	// Declare `intoValue985` variable:
	var intoValue985 atomic.Value

	// Call the method that transfers the taint
	// from the parameter `fromInterface489` to the receiver `intoValue985`
	// (`intoValue985` is now tainted).
	intoValue985.Store(fromInterface489)

	// Return the tainted `intoValue985`:
	return intoValue985
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
