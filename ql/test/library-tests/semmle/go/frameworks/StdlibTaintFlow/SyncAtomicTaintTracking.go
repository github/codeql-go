package main

import (
	"sync/atomic"
	"unsafe"
)

func TaintStepTest_SyncAtomicAddUintptr_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUintptr580` into `intoUintptr459`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr580`:
	fromUintptr580 := sourceCQL.(uintptr)

	// Declare `intoUintptr459` variable:
	var intoUintptr459 *uintptr

	// Call the function that transfers the taint
	// from the parameter `fromUintptr580` to parameter `intoUintptr459`;
	// `intoUintptr459` is now tainted.
	atomic.AddUintptr(intoUintptr459, fromUintptr580)

	// Sink the tainted `intoUintptr459`:
	sink(intoUintptr459)
}

func TaintStepTest_SyncAtomicAddUintptr_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromUintptr408` into `intoUintptr960`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr408`:
	fromUintptr408 := sourceCQL.(uintptr)

	// Call the function that transfers the taint
	// from the parameter `fromUintptr408` to result `intoUintptr960`
	// (`intoUintptr960` is now tainted).
	intoUintptr960 := atomic.AddUintptr(nil, fromUintptr408)

	// Sink the tainted `intoUintptr960`:
	sink(intoUintptr960)
}

func TaintStepTest_SyncAtomicCompareAndSwapPointer_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromPointer149` into `intoPointer398`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer149`:
	fromPointer149 := sourceCQL.(unsafe.Pointer)

	// Declare `intoPointer398` variable:
	var intoPointer398 *unsafe.Pointer

	// Call the function that transfers the taint
	// from the parameter `fromPointer149` to parameter `intoPointer398`;
	// `intoPointer398` is now tainted.
	atomic.CompareAndSwapPointer(intoPointer398, nil, fromPointer149)

	// Sink the tainted `intoPointer398`:
	sink(intoPointer398)
}

func TaintStepTest_SyncAtomicCompareAndSwapUintptr_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUintptr180` into `intoUintptr181`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr180`:
	fromUintptr180 := sourceCQL.(uintptr)

	// Declare `intoUintptr181` variable:
	var intoUintptr181 *uintptr

	// Call the function that transfers the taint
	// from the parameter `fromUintptr180` to parameter `intoUintptr181`;
	// `intoUintptr181` is now tainted.
	atomic.CompareAndSwapUintptr(intoUintptr181, 0, fromUintptr180)

	// Sink the tainted `intoUintptr181`:
	sink(intoUintptr181)
}

func TaintStepTest_SyncAtomicLoadPointer_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromPointer412` into `intoPointer502`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer412`:
	fromPointer412 := sourceCQL.(*unsafe.Pointer)

	// Call the function that transfers the taint
	// from the parameter `fromPointer412` to result `intoPointer502`
	// (`intoPointer502` is now tainted).
	intoPointer502 := atomic.LoadPointer(fromPointer412)

	// Sink the tainted `intoPointer502`:
	sink(intoPointer502)
}

func TaintStepTest_SyncAtomicLoadUintptr_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUintptr548` into `intoUintptr755`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr548`:
	fromUintptr548 := sourceCQL.(*uintptr)

	// Call the function that transfers the taint
	// from the parameter `fromUintptr548` to result `intoUintptr755`
	// (`intoUintptr755` is now tainted).
	intoUintptr755 := atomic.LoadUintptr(fromUintptr548)

	// Sink the tainted `intoUintptr755`:
	sink(intoUintptr755)
}

func TaintStepTest_SyncAtomicStorePointer_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromPointer804` into `intoPointer737`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer804`:
	fromPointer804 := sourceCQL.(unsafe.Pointer)

	// Declare `intoPointer737` variable:
	var intoPointer737 *unsafe.Pointer

	// Call the function that transfers the taint
	// from the parameter `fromPointer804` to parameter `intoPointer737`;
	// `intoPointer737` is now tainted.
	atomic.StorePointer(intoPointer737, fromPointer804)

	// Sink the tainted `intoPointer737`:
	sink(intoPointer737)
}

func TaintStepTest_SyncAtomicStoreUintptr_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUintptr270` into `intoUintptr607`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr270`:
	fromUintptr270 := sourceCQL.(uintptr)

	// Declare `intoUintptr607` variable:
	var intoUintptr607 *uintptr

	// Call the function that transfers the taint
	// from the parameter `fromUintptr270` to parameter `intoUintptr607`;
	// `intoUintptr607` is now tainted.
	atomic.StoreUintptr(intoUintptr607, fromUintptr270)

	// Sink the tainted `intoUintptr607`:
	sink(intoUintptr607)
}

func TaintStepTest_SyncAtomicSwapPointer_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromPointer419` into `intoPointer229`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer419`:
	fromPointer419 := sourceCQL.(unsafe.Pointer)

	// Declare `intoPointer229` variable:
	var intoPointer229 *unsafe.Pointer

	// Call the function that transfers the taint
	// from the parameter `fromPointer419` to parameter `intoPointer229`;
	// `intoPointer229` is now tainted.
	atomic.SwapPointer(intoPointer229, fromPointer419)

	// Sink the tainted `intoPointer229`:
	sink(intoPointer229)
}

func TaintStepTest_SyncAtomicSwapPointer_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromPointer816` into `intoPointer630`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer816`:
	fromPointer816 := sourceCQL.(*unsafe.Pointer)

	// Call the function that transfers the taint
	// from the parameter `fromPointer816` to result `intoPointer630`
	// (`intoPointer630` is now tainted).
	intoPointer630 := atomic.SwapPointer(fromPointer816, nil)

	// Sink the tainted `intoPointer630`:
	sink(intoPointer630)
}

func TaintStepTest_SyncAtomicSwapUintptr_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUintptr363` into `intoUintptr566`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr363`:
	fromUintptr363 := sourceCQL.(uintptr)

	// Declare `intoUintptr566` variable:
	var intoUintptr566 *uintptr

	// Call the function that transfers the taint
	// from the parameter `fromUintptr363` to parameter `intoUintptr566`;
	// `intoUintptr566` is now tainted.
	atomic.SwapUintptr(intoUintptr566, fromUintptr363)

	// Sink the tainted `intoUintptr566`:
	sink(intoUintptr566)
}

func TaintStepTest_SyncAtomicSwapUintptr_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromUintptr650` into `intoUintptr423`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr650`:
	fromUintptr650 := sourceCQL.(*uintptr)

	// Call the function that transfers the taint
	// from the parameter `fromUintptr650` to result `intoUintptr423`
	// (`intoUintptr423` is now tainted).
	intoUintptr423 := atomic.SwapUintptr(fromUintptr650, 0)

	// Sink the tainted `intoUintptr423`:
	sink(intoUintptr423)
}

func TaintStepTest_SyncAtomicValueLoad_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue933` into `intoInterface143`.

	// Assume that `sourceCQL` has the underlying type of `fromValue933`:
	fromValue933 := sourceCQL.(atomic.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue933` to the result `intoInterface143`
	// (`intoInterface143` is now tainted).
	intoInterface143 := fromValue933.Load()

	// Sink the tainted `intoInterface143`:
	sink(intoInterface143)
}

func TaintStepTest_SyncAtomicValueStore_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface138` into `intoValue614`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface138`:
	fromInterface138 := sourceCQL.(interface{})

	// Declare `intoValue614` variable:
	var intoValue614 atomic.Value

	// Call the method that transfers the taint
	// from the parameter `fromInterface138` to the receiver `intoValue614`
	// (`intoValue614` is now tainted).
	intoValue614.Store(fromInterface138)

	// Sink the tainted `intoValue614`:
	sink(intoValue614)
}

func RunAllTaints_SyncAtomic() {
	{
		source := newSource()
		TaintStepTest_SyncAtomicAddUintptr_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicAddUintptr_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicCompareAndSwapPointer_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicCompareAndSwapUintptr_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicLoadPointer_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicLoadUintptr_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicStorePointer_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicStoreUintptr_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicSwapPointer_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicSwapPointer_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicSwapUintptr_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicSwapUintptr_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicValueLoad_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncAtomicValueStore_B0I0O0(source)
	}
}
