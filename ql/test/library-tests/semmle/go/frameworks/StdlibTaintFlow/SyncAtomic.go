package main

import (
	"sync/atomic"
	"unsafe"
)

func TaintStepTest_SyncAtomicAddUintptr(sourceCQL interface{}) {
	// The flow is from `delta` into `addr`.

	// Assume that `sourceCQL` has the underlying type of `delta`:
	delta := sourceCQL.(uintptr)

	// Declare `addr` variable:
	var addr *uintptr

	// Call the function that transfers the taint
	// from the parameter `delta` to parameter `addr`;
	// `addr` is now tainted.
	atomic.AddUintptr(addr, delta)

	// Sink the tainted `addr`:
	sink(addr)
}

func TaintStepTest_SyncAtomicCompareAndSwapPointer(sourceCQL interface{}) {
	// The flow is from `new` into `addr`.

	// Assume that `sourceCQL` has the underlying type of `new`:
	new := sourceCQL.(unsafe.Pointer)

	// Declare `addr` variable:
	var addr *unsafe.Pointer

	// Call the function that transfers the taint
	// from the parameter `new` to parameter `addr`;
	// `addr` is now tainted.
	atomic.CompareAndSwapPointer(addr, nil, new)

	// Sink the tainted `addr`:
	sink(addr)
}

func TaintStepTest_SyncAtomicCompareAndSwapUintptr(sourceCQL interface{}) {
	// The flow is from `new` into `addr`.

	// Assume that `sourceCQL` has the underlying type of `new`:
	new := sourceCQL.(uintptr)

	// Declare `addr` variable:
	var addr *uintptr

	// Call the function that transfers the taint
	// from the parameter `new` to parameter `addr`;
	// `addr` is now tainted.
	atomic.CompareAndSwapUintptr(addr, 0, new)

	// Sink the tainted `addr`:
	sink(addr)
}

func TaintStepTest_SyncAtomicLoadPointer(sourceCQL interface{}) {
	// The flow is from `addr` into `val`.

	// Assume that `sourceCQL` has the underlying type of `addr`:
	addr := sourceCQL.(*unsafe.Pointer)

	// Call the function that transfers the taint
	// from the parameter `addr` to result `val`
	// (`val` is now tainted).
	val := atomic.LoadPointer(addr)

	// Sink the tainted `val`:
	sink(val)
}

func TaintStepTest_SyncAtomicLoadUintptr(sourceCQL interface{}) {
	// The flow is from `addr` into `val`.

	// Assume that `sourceCQL` has the underlying type of `addr`:
	addr := sourceCQL.(*uintptr)

	// Call the function that transfers the taint
	// from the parameter `addr` to result `val`
	// (`val` is now tainted).
	val := atomic.LoadUintptr(addr)

	// Sink the tainted `val`:
	sink(val)
}

func TaintStepTest_SyncAtomicStorePointer(sourceCQL interface{}) {
	// The flow is from `val` into `addr`.

	// Assume that `sourceCQL` has the underlying type of `val`:
	val := sourceCQL.(unsafe.Pointer)

	// Declare `addr` variable:
	var addr *unsafe.Pointer

	// Call the function that transfers the taint
	// from the parameter `val` to parameter `addr`;
	// `addr` is now tainted.
	atomic.StorePointer(addr, val)

	// Sink the tainted `addr`:
	sink(addr)
}

func TaintStepTest_SyncAtomicStoreUintptr(sourceCQL interface{}) {
	// The flow is from `val` into `addr`.

	// Assume that `sourceCQL` has the underlying type of `val`:
	val := sourceCQL.(uintptr)

	// Declare `addr` variable:
	var addr *uintptr

	// Call the function that transfers the taint
	// from the parameter `val` to parameter `addr`;
	// `addr` is now tainted.
	atomic.StoreUintptr(addr, val)

	// Sink the tainted `addr`:
	sink(addr)
}

func TaintStepTest_SyncAtomicSwapPointer(sourceCQL interface{}) {
	// The flow is from `new` into `addr`.

	// Assume that `sourceCQL` has the underlying type of `new`:
	new := sourceCQL.(unsafe.Pointer)

	// Declare `addr` variable:
	var addr *unsafe.Pointer

	// Call the function that transfers the taint
	// from the parameter `new` to parameter `addr`;
	// `addr` is now tainted.
	atomic.SwapPointer(addr, new)

	// Sink the tainted `addr`:
	sink(addr)
}

func TaintStepTest_SyncAtomicSwapUintptr(sourceCQL interface{}) {
	// The flow is from `new` into `addr`.

	// Assume that `sourceCQL` has the underlying type of `new`:
	new := sourceCQL.(uintptr)

	// Declare `addr` variable:
	var addr *uintptr

	// Call the function that transfers the taint
	// from the parameter `new` to parameter `addr`;
	// `addr` is now tainted.
	atomic.SwapUintptr(addr, new)

	// Sink the tainted `addr`:
	sink(addr)
}

func TaintStepTest_SyncAtomicValueLoad(sourceCQL interface{}) {
	// The flow is from `from607` into `x`.

	// Assume that `sourceCQL` has the underlying type of `from607`:
	from607 := sourceCQL.(atomic.Value)

	// Call the method that transfers the taint
	// from the receiver `from607` to the result `x`
	// (`x` is now tainted).
	x := from607.Load()

	// Sink the tainted `x`:
	sink(x)
}

func TaintStepTest_SyncAtomicValueStore(sourceCQL interface{}) {
	// The flow is from `x` into `into462`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.(interface{})

	// Declare `into462` variable:
	var into462 atomic.Value

	// Call the method that transfers the taint
	// from the parameter `x` to the receiver `into462`
	// (`into462` is now tainted).
	into462.Store(x)

	// Sink the tainted `into462`:
	sink(into462)
}

func RunAllTaints_SyncAtomic(v interface{}) {
	{
		source := newSource()
		TaintStepTest_SyncAtomicAddUintptr(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicCompareAndSwapPointer(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicCompareAndSwapUintptr(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicLoadPointer(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicLoadUintptr(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicStorePointer(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicStoreUintptr(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicSwapPointer(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicSwapUintptr(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicValueLoad(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicValueStore(source)
	}
}
