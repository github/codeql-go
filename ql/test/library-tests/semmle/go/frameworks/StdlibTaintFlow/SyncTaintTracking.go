// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "sync"

func TaintStepTest_SyncMapLoad_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMap276` into `intoInterface697`.

	// Assume that `sourceCQL` has the underlying type of `fromMap276`:
	fromMap276 := sourceCQL.(sync.Map)

	// Call the method that transfers the taint
	// from the receiver `fromMap276` to the result `intoInterface697`
	// (`intoInterface697` is now tainted).
	intoInterface697, _ := fromMap276.Load(nil)

	// Return the tainted `intoInterface697`:
	return intoInterface697
}

func TaintStepTest_SyncMapLoadOrStore_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMap165` into `intoInterface211`.

	// Assume that `sourceCQL` has the underlying type of `fromMap165`:
	fromMap165 := sourceCQL.(sync.Map)

	// Call the method that transfers the taint
	// from the receiver `fromMap165` to the result `intoInterface211`
	// (`intoInterface211` is now tainted).
	intoInterface211, _ := fromMap165.LoadOrStore(nil, nil)

	// Return the tainted `intoInterface211`:
	return intoInterface211
}

func TaintStepTest_SyncMapLoadOrStore_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface843` into `intoMap938`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface843`:
	fromInterface843 := sourceCQL.(interface{})

	// Declare `intoMap938` variable:
	var intoMap938 sync.Map

	// Call the method that transfers the taint
	// from the parameter `fromInterface843` to the receiver `intoMap938`
	// (`intoMap938` is now tainted).
	intoMap938.LoadOrStore(nil, fromInterface843)

	// Return the tainted `intoMap938`:
	return intoMap938
}

func TaintStepTest_SyncMapLoadOrStore_B1I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface428` into `intoInterface251`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface428`:
	fromInterface428 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL sync.Map

	// Call the method that transfers the taint
	// from the parameter `fromInterface428` to the result `intoInterface251`
	// (`intoInterface251` is now tainted).
	intoInterface251, _ := mediumObjCQL.LoadOrStore(nil, fromInterface428)

	// Return the tainted `intoInterface251`:
	return intoInterface251
}

func TaintStepTest_SyncMapRange_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMap959` into `intoFunckeyInterfaceValueInterfaceBool408`.

	// Assume that `sourceCQL` has the underlying type of `fromMap959`:
	fromMap959 := sourceCQL.(sync.Map)

	// Declare `intoFunckeyInterfaceValueInterfaceBool408` variable:
	var intoFunckeyInterfaceValueInterfaceBool408 func(interface{}, interface{}) bool

	// Call the method that transfers the taint
	// from the receiver `fromMap959` to the argument `intoFunckeyInterfaceValueInterfaceBool408`
	// (`intoFunckeyInterfaceValueInterfaceBool408` is now tainted).
	fromMap959.Range(intoFunckeyInterfaceValueInterfaceBool408)

	// Return the tainted `intoFunckeyInterfaceValueInterfaceBool408`:
	return intoFunckeyInterfaceValueInterfaceBool408
}

func TaintStepTest_SyncMapStore_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface306` into `intoMap831`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface306`:
	fromInterface306 := sourceCQL.(interface{})

	// Declare `intoMap831` variable:
	var intoMap831 sync.Map

	// Call the method that transfers the taint
	// from the parameter `fromInterface306` to the receiver `intoMap831`
	// (`intoMap831` is now tainted).
	intoMap831.Store(fromInterface306, nil)

	// Return the tainted `intoMap831`:
	return intoMap831
}

func TaintStepTest_SyncMapStore_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface136` into `intoMap705`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface136`:
	fromInterface136 := sourceCQL.(interface{})

	// Declare `intoMap705` variable:
	var intoMap705 sync.Map

	// Call the method that transfers the taint
	// from the parameter `fromInterface136` to the receiver `intoMap705`
	// (`intoMap705` is now tainted).
	intoMap705.Store(nil, fromInterface136)

	// Return the tainted `intoMap705`:
	return intoMap705
}

func TaintStepTest_SyncPoolGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPool649` into `intoInterface952`.

	// Assume that `sourceCQL` has the underlying type of `fromPool649`:
	fromPool649 := sourceCQL.(sync.Pool)

	// Call the method that transfers the taint
	// from the receiver `fromPool649` to the result `intoInterface952`
	// (`intoInterface952` is now tainted).
	intoInterface952 := fromPool649.Get()

	// Return the tainted `intoInterface952`:
	return intoInterface952
}

func TaintStepTest_SyncPoolPut_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface186` into `intoPool649`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface186`:
	fromInterface186 := sourceCQL.(interface{})

	// Declare `intoPool649` variable:
	var intoPool649 sync.Pool

	// Call the method that transfers the taint
	// from the parameter `fromInterface186` to the receiver `intoPool649`
	// (`intoPool649` is now tainted).
	intoPool649.Put(fromInterface186)

	// Return the tainted `intoPool649`:
	return intoPool649
}

func RunAllTaints_Sync() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncMapLoad_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncMapLoadOrStore_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncMapLoadOrStore_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncMapLoadOrStore_B1I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncMapRange_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncMapStore_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncMapStore_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncPoolGet_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyncPoolPut_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
