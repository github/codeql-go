package main

import (
	"sync/atomic"
	"unsafe"
)

func TaintStepTest_SyncAtomicAddUintptr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr582` into `intoUintptr753`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr582`:
	fromUintptr582 := sourceCQL.(uintptr)

	// Declare `intoUintptr753` variable:
	var intoUintptr753 *uintptr

	// Call the function that transfers the taint
	// from the parameter `fromUintptr582` to parameter `intoUintptr753`;
	// `intoUintptr753` is now tainted.
	atomic.AddUintptr(intoUintptr753, fromUintptr582)

	// Return the tainted `intoUintptr753`:
	return intoUintptr753
}

func TaintStepTest_SyncAtomicAddUintptr_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr187` into `intoUintptr980`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr187`:
	fromUintptr187 := sourceCQL.(uintptr)

	// Call the function that transfers the taint
	// from the parameter `fromUintptr187` to result `intoUintptr980`
	// (`intoUintptr980` is now tainted).
	intoUintptr980 := atomic.AddUintptr(nil, fromUintptr187)

	// Return the tainted `intoUintptr980`:
	return intoUintptr980
}

func TaintStepTest_SyncAtomicCompareAndSwapPointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer240` into `intoPointer455`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer240`:
	fromPointer240 := sourceCQL.(unsafe.Pointer)

	// Declare `intoPointer455` variable:
	var intoPointer455 *unsafe.Pointer

	// Call the function that transfers the taint
	// from the parameter `fromPointer240` to parameter `intoPointer455`;
	// `intoPointer455` is now tainted.
	atomic.CompareAndSwapPointer(intoPointer455, nil, fromPointer240)

	// Return the tainted `intoPointer455`:
	return intoPointer455
}

func TaintStepTest_SyncAtomicCompareAndSwapUintptr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr962` into `intoUintptr323`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr962`:
	fromUintptr962 := sourceCQL.(uintptr)

	// Declare `intoUintptr323` variable:
	var intoUintptr323 *uintptr

	// Call the function that transfers the taint
	// from the parameter `fromUintptr962` to parameter `intoUintptr323`;
	// `intoUintptr323` is now tainted.
	atomic.CompareAndSwapUintptr(intoUintptr323, 0, fromUintptr962)

	// Return the tainted `intoUintptr323`:
	return intoUintptr323
}

func TaintStepTest_SyncAtomicLoadPointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer244` into `intoPointer748`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer244`:
	fromPointer244 := sourceCQL.(*unsafe.Pointer)

	// Call the function that transfers the taint
	// from the parameter `fromPointer244` to result `intoPointer748`
	// (`intoPointer748` is now tainted).
	intoPointer748 := atomic.LoadPointer(fromPointer244)

	// Return the tainted `intoPointer748`:
	return intoPointer748
}

func TaintStepTest_SyncAtomicLoadUintptr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr527` into `intoUintptr127`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr527`:
	fromUintptr527 := sourceCQL.(*uintptr)

	// Call the function that transfers the taint
	// from the parameter `fromUintptr527` to result `intoUintptr127`
	// (`intoUintptr127` is now tainted).
	intoUintptr127 := atomic.LoadUintptr(fromUintptr527)

	// Return the tainted `intoUintptr127`:
	return intoUintptr127
}

func TaintStepTest_SyncAtomicStorePointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer283` into `intoPointer281`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer283`:
	fromPointer283 := sourceCQL.(unsafe.Pointer)

	// Declare `intoPointer281` variable:
	var intoPointer281 *unsafe.Pointer

	// Call the function that transfers the taint
	// from the parameter `fromPointer283` to parameter `intoPointer281`;
	// `intoPointer281` is now tainted.
	atomic.StorePointer(intoPointer281, fromPointer283)

	// Return the tainted `intoPointer281`:
	return intoPointer281
}

func TaintStepTest_SyncAtomicStoreUintptr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr607` into `intoUintptr348`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr607`:
	fromUintptr607 := sourceCQL.(uintptr)

	// Declare `intoUintptr348` variable:
	var intoUintptr348 *uintptr

	// Call the function that transfers the taint
	// from the parameter `fromUintptr607` to parameter `intoUintptr348`;
	// `intoUintptr348` is now tainted.
	atomic.StoreUintptr(intoUintptr348, fromUintptr607)

	// Return the tainted `intoUintptr348`:
	return intoUintptr348
}

func TaintStepTest_SyncAtomicSwapPointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer530` into `intoPointer750`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer530`:
	fromPointer530 := sourceCQL.(unsafe.Pointer)

	// Declare `intoPointer750` variable:
	var intoPointer750 *unsafe.Pointer

	// Call the function that transfers the taint
	// from the parameter `fromPointer530` to parameter `intoPointer750`;
	// `intoPointer750` is now tainted.
	atomic.SwapPointer(intoPointer750, fromPointer530)

	// Return the tainted `intoPointer750`:
	return intoPointer750
}

func TaintStepTest_SyncAtomicSwapPointer_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer792` into `intoPointer237`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer792`:
	fromPointer792 := sourceCQL.(*unsafe.Pointer)

	// Call the function that transfers the taint
	// from the parameter `fromPointer792` to result `intoPointer237`
	// (`intoPointer237` is now tainted).
	intoPointer237 := atomic.SwapPointer(fromPointer792, nil)

	// Return the tainted `intoPointer237`:
	return intoPointer237
}

func TaintStepTest_SyncAtomicSwapUintptr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr814` into `intoUintptr577`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr814`:
	fromUintptr814 := sourceCQL.(uintptr)

	// Declare `intoUintptr577` variable:
	var intoUintptr577 *uintptr

	// Call the function that transfers the taint
	// from the parameter `fromUintptr814` to parameter `intoUintptr577`;
	// `intoUintptr577` is now tainted.
	atomic.SwapUintptr(intoUintptr577, fromUintptr814)

	// Return the tainted `intoUintptr577`:
	return intoUintptr577
}

func TaintStepTest_SyncAtomicSwapUintptr_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr192` into `intoUintptr169`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr192`:
	fromUintptr192 := sourceCQL.(*uintptr)

	// Call the function that transfers the taint
	// from the parameter `fromUintptr192` to result `intoUintptr169`
	// (`intoUintptr169` is now tainted).
	intoUintptr169 := atomic.SwapUintptr(fromUintptr192, 0)

	// Return the tainted `intoUintptr169`:
	return intoUintptr169
}

func TaintStepTest_SyncAtomicValueLoad_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue132` into `intoInterface600`.

	// Assume that `sourceCQL` has the underlying type of `fromValue132`:
	fromValue132 := sourceCQL.(atomic.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue132` to the result `intoInterface600`
	// (`intoInterface600` is now tainted).
	intoInterface600 := fromValue132.Load()

	// Return the tainted `intoInterface600`:
	return intoInterface600
}

func TaintStepTest_SyncAtomicValueStore_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface119` into `intoValue275`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface119`:
	fromInterface119 := sourceCQL.(interface{})

	// Declare `intoValue275` variable:
	var intoValue275 atomic.Value

	// Call the method that transfers the taint
	// from the parameter `fromInterface119` to the receiver `intoValue275`
	// (`intoValue275` is now tainted).
	intoValue275.Store(fromInterface119)

	// Return the tainted `intoValue275`:
	return intoValue275
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
