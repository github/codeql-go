package main

import "sync"

func TaintStepTest_SyncMapLoad_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMap468` into `intoInterface772`.

	// Assume that `sourceCQL` has the underlying type of `fromMap468`:
	fromMap468 := sourceCQL.(sync.Map)

	// Call the method that transfers the taint
	// from the receiver `fromMap468` to the result `intoInterface772`
	// (`intoInterface772` is now tainted).
	intoInterface772, _ := fromMap468.Load(nil)

	// Return the tainted `intoInterface772`:
	return intoInterface772
}

func TaintStepTest_SyncMapLoadOrStore_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMap660` into `intoInterface654`.

	// Assume that `sourceCQL` has the underlying type of `fromMap660`:
	fromMap660 := sourceCQL.(sync.Map)

	// Call the method that transfers the taint
	// from the receiver `fromMap660` to the result `intoInterface654`
	// (`intoInterface654` is now tainted).
	intoInterface654, _ := fromMap660.LoadOrStore(nil, nil)

	// Return the tainted `intoInterface654`:
	return intoInterface654
}

func TaintStepTest_SyncMapLoadOrStore_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface945` into `intoMap365`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface945`:
	fromInterface945 := sourceCQL.(interface{})

	// Declare `intoMap365` variable:
	var intoMap365 sync.Map

	// Call the method that transfers the taint
	// from the parameter `fromInterface945` to the receiver `intoMap365`
	// (`intoMap365` is now tainted).
	intoMap365.LoadOrStore(nil, fromInterface945)

	// Return the tainted `intoMap365`:
	return intoMap365
}

func TaintStepTest_SyncMapLoadOrStore_B1I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface842` into `intoInterface765`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface842`:
	fromInterface842 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL sync.Map

	// Call the method that transfers the taint
	// from the parameter `fromInterface842` to the result `intoInterface765`
	// (`intoInterface765` is now tainted).
	intoInterface765, _ := mediumObjCQL.LoadOrStore(nil, fromInterface842)

	// Return the tainted `intoInterface765`:
	return intoInterface765
}

func TaintStepTest_SyncMapRange_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMap318` into `intoFunckeyInterfaceValueInterfaceBool499`.

	// Assume that `sourceCQL` has the underlying type of `fromMap318`:
	fromMap318 := sourceCQL.(sync.Map)

	// Declare `intoFunckeyInterfaceValueInterfaceBool499` variable:
	var intoFunckeyInterfaceValueInterfaceBool499 func(interface{}, interface{}) bool

	// Call the method that transfers the taint
	// from the receiver `fromMap318` to the argument `intoFunckeyInterfaceValueInterfaceBool499`
	// (`intoFunckeyInterfaceValueInterfaceBool499` is now tainted).
	fromMap318.Range(intoFunckeyInterfaceValueInterfaceBool499)

	// Return the tainted `intoFunckeyInterfaceValueInterfaceBool499`:
	return intoFunckeyInterfaceValueInterfaceBool499
}

func TaintStepTest_SyncMapStore_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface471` into `intoMap536`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface471`:
	fromInterface471 := sourceCQL.(interface{})

	// Declare `intoMap536` variable:
	var intoMap536 sync.Map

	// Call the method that transfers the taint
	// from the parameter `fromInterface471` to the receiver `intoMap536`
	// (`intoMap536` is now tainted).
	intoMap536.Store(fromInterface471, nil)

	// Return the tainted `intoMap536`:
	return intoMap536
}

func TaintStepTest_SyncMapStore_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface667` into `intoMap630`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface667`:
	fromInterface667 := sourceCQL.(interface{})

	// Declare `intoMap630` variable:
	var intoMap630 sync.Map

	// Call the method that transfers the taint
	// from the parameter `fromInterface667` to the receiver `intoMap630`
	// (`intoMap630` is now tainted).
	intoMap630.Store(nil, fromInterface667)

	// Return the tainted `intoMap630`:
	return intoMap630
}

func TaintStepTest_SyncPoolGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPool774` into `intoInterface163`.

	// Assume that `sourceCQL` has the underlying type of `fromPool774`:
	fromPool774 := sourceCQL.(sync.Pool)

	// Call the method that transfers the taint
	// from the receiver `fromPool774` to the result `intoInterface163`
	// (`intoInterface163` is now tainted).
	intoInterface163 := fromPool774.Get()

	// Return the tainted `intoInterface163`:
	return intoInterface163
}

func TaintStepTest_SyncPoolPut_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface951` into `intoPool184`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface951`:
	fromInterface951 := sourceCQL.(interface{})

	// Declare `intoPool184` variable:
	var intoPool184 sync.Pool

	// Call the method that transfers the taint
	// from the parameter `fromInterface951` to the receiver `intoPool184`
	// (`intoPool184` is now tainted).
	intoPool184.Put(fromInterface951)

	// Return the tainted `intoPool184`:
	return intoPool184
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
