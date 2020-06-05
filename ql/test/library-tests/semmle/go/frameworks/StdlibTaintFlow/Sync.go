package main

import "sync"

func TaintStepTest_SyncMapLoad(sourceCQL interface{}) {
	// The flow is from `from963` into `value`.

	// Assume that `sourceCQL` has the underlying type of `from963`:
	from963 := sourceCQL.(sync.Map)

	// Call the method that transfers the taint
	// from the receiver `from963` to the result `value`
	// (`value` is now tainted).
	value, _ := from963.Load(nil)

	// Sink the tainted `value`:
	sink(value)
}

func TaintStepTest_SyncMapLoadOrStore(sourceCQL interface{}) {
	// The flow is from `from261` into `actual`.

	// Assume that `sourceCQL` has the underlying type of `from261`:
	from261 := sourceCQL.(sync.Map)

	// Call the method that transfers the taint
	// from the receiver `from261` to the result `actual`
	// (`actual` is now tainted).
	actual, _ := from261.LoadOrStore(nil, nil)

	// Sink the tainted `actual`:
	sink(actual)
}

func TaintStepTest_SyncMapRange(sourceCQL interface{}) {
	// The flow is from `from919` into `f`.

	// Assume that `sourceCQL` has the underlying type of `from919`:
	from919 := sourceCQL.(sync.Map)

	// Declare `f` variable:
	var f func(interface{}, interface{}) bool

	// Call the method that transfers the taint
	// from the receiver `from919` to the argument `f`
	// (`f` is now tainted).
	from919.Range(f)

	// Sink the tainted `f`:
	sink(f)
}

func TaintStepTest_SyncMapStore(sourceCQL interface{}) {
	// The flow is from `key` into `into555`.

	// Assume that `sourceCQL` has the underlying type of `key`:
	key := sourceCQL.(interface{})

	// Declare `into555` variable:
	var into555 sync.Map

	// Call the method that transfers the taint
	// from the parameter `key` to the receiver `into555`
	// (`into555` is now tainted).
	into555.Store(key, nil)

	// Sink the tainted `into555`:
	sink(into555)
}

func TaintStepTest_SyncPoolGet(sourceCQL interface{}) {
	// The flow is from `from577` into `into611`.

	// Assume that `sourceCQL` has the underlying type of `from577`:
	from577 := sourceCQL.(sync.Pool)

	// Call the method that transfers the taint
	// from the receiver `from577` to the result `into611`
	// (`into611` is now tainted).
	into611 := from577.Get()

	// Sink the tainted `into611`:
	sink(into611)
}

func TaintStepTest_SyncPoolPut(sourceCQL interface{}) {
	// The flow is from `x` into `into264`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.(interface{})

	// Declare `into264` variable:
	var into264 sync.Pool

	// Call the method that transfers the taint
	// from the parameter `x` to the receiver `into264`
	// (`into264` is now tainted).
	into264.Put(x)

	// Sink the tainted `into264`:
	sink(into264)
}

func RunAllTaints_Sync(v interface{}) {
	{
		source := newSource()
		TaintStepTest_SyncMapLoad(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncMapLoadOrStore(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncMapRange(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncMapStore(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncPoolGet(source)
	}
	{
		source := newSource()
		TaintStepTest_SyncPoolPut(source)
	}
}
