package main

import "sync"

func TaintStepTest_SyncMapLoad_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromMap920` into `intoInterface615`.

	// Assume that `sourceCQL` has the underlying type of `fromMap920`:
	fromMap920 := sourceCQL.(sync.Map)

	// Call the method that transfers the taint
	// from the receiver `fromMap920` to the result `intoInterface615`
	// (`intoInterface615` is now tainted).
	intoInterface615, _ := fromMap920.Load(nil)

	// Sink the tainted `intoInterface615`:
	sink(intoInterface615)
}

func TaintStepTest_SyncMapLoadOrStore_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromMap377` into `intoInterface588`.

	// Assume that `sourceCQL` has the underlying type of `fromMap377`:
	fromMap377 := sourceCQL.(sync.Map)

	// Call the method that transfers the taint
	// from the receiver `fromMap377` to the result `intoInterface588`
	// (`intoInterface588` is now tainted).
	intoInterface588, _ := fromMap377.LoadOrStore(nil, nil)

	// Sink the tainted `intoInterface588`:
	sink(intoInterface588)
}

func TaintStepTest_SyncMapLoadOrStore_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface660` into `intoMap879`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface660`:
	fromInterface660 := sourceCQL.(interface{})

	// Declare `intoMap879` variable:
	var intoMap879 sync.Map

	// Call the method that transfers the taint
	// from the parameter `fromInterface660` to the receiver `intoMap879`
	// (`intoMap879` is now tainted).
	intoMap879.LoadOrStore(nil, fromInterface660)

	// Sink the tainted `intoMap879`:
	sink(intoMap879)
}

func TaintStepTest_SyncMapLoadOrStore_B1I0O1(sourceCQL interface{}) {
	// The flow is from `fromInterface689` into `intoInterface117`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface689`:
	fromInterface689 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL sync.Map

	// Call the method that transfers the taint
	// from the parameter `fromInterface689` to the result `intoInterface117`
	// (`intoInterface117` is now tainted).
	intoInterface117, _ := mediumObjCQL.LoadOrStore(nil, fromInterface689)

	// Sink the tainted `intoInterface117`:
	sink(intoInterface117)
}

func TaintStepTest_SyncMapRange_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromMap560` into `intoFunckeyInterfaceValueInterfaceBool936`.

	// Assume that `sourceCQL` has the underlying type of `fromMap560`:
	fromMap560 := sourceCQL.(sync.Map)

	// Declare `intoFunckeyInterfaceValueInterfaceBool936` variable:
	var intoFunckeyInterfaceValueInterfaceBool936 func(interface{}, interface{}) bool

	// Call the method that transfers the taint
	// from the receiver `fromMap560` to the argument `intoFunckeyInterfaceValueInterfaceBool936`
	// (`intoFunckeyInterfaceValueInterfaceBool936` is now tainted).
	fromMap560.Range(intoFunckeyInterfaceValueInterfaceBool936)

	// Sink the tainted `intoFunckeyInterfaceValueInterfaceBool936`:
	sink(intoFunckeyInterfaceValueInterfaceBool936)
}

func TaintStepTest_SyncMapStore_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface449` into `intoMap944`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface449`:
	fromInterface449 := sourceCQL.(interface{})

	// Declare `intoMap944` variable:
	var intoMap944 sync.Map

	// Call the method that transfers the taint
	// from the parameter `fromInterface449` to the receiver `intoMap944`
	// (`intoMap944` is now tainted).
	intoMap944.Store(fromInterface449, nil)

	// Sink the tainted `intoMap944`:
	sink(intoMap944)
}

func TaintStepTest_SyncMapStore_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromInterface616` into `intoMap451`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface616`:
	fromInterface616 := sourceCQL.(interface{})

	// Declare `intoMap451` variable:
	var intoMap451 sync.Map

	// Call the method that transfers the taint
	// from the parameter `fromInterface616` to the receiver `intoMap451`
	// (`intoMap451` is now tainted).
	intoMap451.Store(nil, fromInterface616)

	// Sink the tainted `intoMap451`:
	sink(intoMap451)
}

func TaintStepTest_SyncPoolGet_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromPool497` into `intoInterface554`.

	// Assume that `sourceCQL` has the underlying type of `fromPool497`:
	fromPool497 := sourceCQL.(sync.Pool)

	// Call the method that transfers the taint
	// from the receiver `fromPool497` to the result `intoInterface554`
	// (`intoInterface554` is now tainted).
	intoInterface554 := fromPool497.Get()

	// Sink the tainted `intoInterface554`:
	sink(intoInterface554)
}

func TaintStepTest_SyncPoolPut_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface157` into `intoPool784`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface157`:
	fromInterface157 := sourceCQL.(interface{})

	// Declare `intoPool784` variable:
	var intoPool784 sync.Pool

	// Call the method that transfers the taint
	// from the parameter `fromInterface157` to the receiver `intoPool784`
	// (`intoPool784` is now tainted).
	intoPool784.Put(fromInterface157)

	// Sink the tainted `intoPool784`:
	sink(intoPool784)
}

func RunAllTaints_Sync() {
	{
		source := newSource()
		TaintStepTest_SyncMapLoad_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncMapLoadOrStore_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncMapLoadOrStore_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncMapLoadOrStore_B1I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncMapRange_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncMapStore_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncMapStore_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncPoolGet_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncPoolPut_B0I0O0(source)
	}
}
