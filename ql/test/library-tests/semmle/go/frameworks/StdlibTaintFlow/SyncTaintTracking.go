// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "sync"

func TaintStepTest_SyncMapLoad_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMap656` into `intoInterface414`.

	// Assume that `sourceCQL` has the underlying type of `fromMap656`:
	fromMap656 := sourceCQL.(sync.Map)

	// Call the method that transfers the taint
	// from the receiver `fromMap656` to the result `intoInterface414`
	// (`intoInterface414` is now tainted).
	intoInterface414, _ := fromMap656.Load(nil)

	// Return the tainted `intoInterface414`:
	return intoInterface414
}

func TaintStepTest_SyncMapLoadOrStore_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMap518` into `intoInterface650`.

	// Assume that `sourceCQL` has the underlying type of `fromMap518`:
	fromMap518 := sourceCQL.(sync.Map)

	// Call the method that transfers the taint
	// from the receiver `fromMap518` to the result `intoInterface650`
	// (`intoInterface650` is now tainted).
	intoInterface650, _ := fromMap518.LoadOrStore(nil, nil)

	// Return the tainted `intoInterface650`:
	return intoInterface650
}

func TaintStepTest_SyncMapLoadOrStore_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface784` into `intoMap957`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface784`:
	fromInterface784 := sourceCQL.(interface{})

	// Declare `intoMap957` variable:
	var intoMap957 sync.Map

	// Call the method that transfers the taint
	// from the parameter `fromInterface784` to the receiver `intoMap957`
	// (`intoMap957` is now tainted).
	intoMap957.LoadOrStore(nil, fromInterface784)

	// Return the tainted `intoMap957`:
	return intoMap957
}

func TaintStepTest_SyncMapLoadOrStore_B1I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface520` into `intoInterface443`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface520`:
	fromInterface520 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL sync.Map

	// Call the method that transfers the taint
	// from the parameter `fromInterface520` to the result `intoInterface443`
	// (`intoInterface443` is now tainted).
	intoInterface443, _ := mediumObjCQL.LoadOrStore(nil, fromInterface520)

	// Return the tainted `intoInterface443`:
	return intoInterface443
}

func TaintStepTest_SyncMapRange_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMap127` into `intoFunckeyInterfaceValueInterfaceBool483`.

	// Assume that `sourceCQL` has the underlying type of `fromMap127`:
	fromMap127 := sourceCQL.(sync.Map)

	// Declare `intoFunckeyInterfaceValueInterfaceBool483` variable:
	var intoFunckeyInterfaceValueInterfaceBool483 func(interface{}, interface{}) bool

	// Call the method that transfers the taint
	// from the receiver `fromMap127` to the argument `intoFunckeyInterfaceValueInterfaceBool483`
	// (`intoFunckeyInterfaceValueInterfaceBool483` is now tainted).
	fromMap127.Range(intoFunckeyInterfaceValueInterfaceBool483)

	// Return the tainted `intoFunckeyInterfaceValueInterfaceBool483`:
	return intoFunckeyInterfaceValueInterfaceBool483
}

func TaintStepTest_SyncMapStore_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface989` into `intoMap982`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface989`:
	fromInterface989 := sourceCQL.(interface{})

	// Declare `intoMap982` variable:
	var intoMap982 sync.Map

	// Call the method that transfers the taint
	// from the parameter `fromInterface989` to the receiver `intoMap982`
	// (`intoMap982` is now tainted).
	intoMap982.Store(fromInterface989, nil)

	// Return the tainted `intoMap982`:
	return intoMap982
}

func TaintStepTest_SyncMapStore_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface417` into `intoMap584`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface417`:
	fromInterface417 := sourceCQL.(interface{})

	// Declare `intoMap584` variable:
	var intoMap584 sync.Map

	// Call the method that transfers the taint
	// from the parameter `fromInterface417` to the receiver `intoMap584`
	// (`intoMap584` is now tainted).
	intoMap584.Store(nil, fromInterface417)

	// Return the tainted `intoMap584`:
	return intoMap584
}

func TaintStepTest_SyncPoolGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPool991` into `intoInterface881`.

	// Assume that `sourceCQL` has the underlying type of `fromPool991`:
	fromPool991 := sourceCQL.(sync.Pool)

	// Call the method that transfers the taint
	// from the receiver `fromPool991` to the result `intoInterface881`
	// (`intoInterface881` is now tainted).
	intoInterface881 := fromPool991.Get()

	// Return the tainted `intoInterface881`:
	return intoInterface881
}

func TaintStepTest_SyncPoolPut_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface186` into `intoPool284`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface186`:
	fromInterface186 := sourceCQL.(interface{})

	// Declare `intoPool284` variable:
	var intoPool284 sync.Pool

	// Call the method that transfers the taint
	// from the parameter `fromInterface186` to the receiver `intoPool284`
	// (`intoPool284` is now tainted).
	intoPool284.Put(fromInterface186)

	// Return the tainted `intoPool284`:
	return intoPool284
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
