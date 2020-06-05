package main

import (
	"context"
	"time"
)

func TaintStepTest_ContextWithCancel(sourceCQL interface{}) {
	// The flow is from `parent` into `ctx`.

	// Assume that `sourceCQL` has the underlying type of `parent`:
	parent := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `parent` to result `ctx`
	// (`ctx` is now tainted).
	ctx, _ := context.WithCancel(parent)

	// Sink the tainted `ctx`:
	sink(ctx)
}

func TaintStepTest_ContextWithDeadline(sourceCQL interface{}) {
	// The flow is from `parent` into `into564`.

	// Assume that `sourceCQL` has the underlying type of `parent`:
	parent := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `parent` to result `into564`
	// (`into564` is now tainted).
	into564, _ := context.WithDeadline(parent, time.Time{})

	// Sink the tainted `into564`:
	sink(into564)
}

func TaintStepTest_ContextWithTimeout(sourceCQL interface{}) {
	// The flow is from `parent` into `into540`.

	// Assume that `sourceCQL` has the underlying type of `parent`:
	parent := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `parent` to result `into540`
	// (`into540` is now tainted).
	into540, _ := context.WithTimeout(parent, 0)

	// Sink the tainted `into540`:
	sink(into540)
}

func TaintStepTest_ContextWithValue(sourceCQL interface{}) {
	// The flow is from `parent` into `into145`.

	// Assume that `sourceCQL` has the underlying type of `parent`:
	parent := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `parent` to result `into145`
	// (`into145` is now tainted).
	into145 := context.WithValue(parent, nil, nil)

	// Sink the tainted `into145`:
	sink(into145)
}

func TaintStepTest_ContextContextValue(sourceCQL interface{}) {
	// The flow is from `from948` into `into731`.

	// Assume that `sourceCQL` has the underlying type of `from948`:
	from948 := sourceCQL.(context.Context)

	// Call the method that transfers the taint
	// from the receiver `from948` to the result `into731`
	// (`into731` is now tainted).
	into731 := from948.Value(nil)

	// Sink the tainted `into731`:
	sink(into731)
}

func RunAllTaints_Context(v interface{}) {
	{
		source := newSource()
		TaintStepTest_ContextWithCancel(source)
	}
	{
		source := newSource()
		TaintStepTest_ContextWithDeadline(source)
	}
	{
		source := newSource()
		TaintStepTest_ContextWithTimeout(source)
	}
	{
		source := newSource()
		TaintStepTest_ContextWithValue(source)
	}
	{
		source := newSource()
		TaintStepTest_ContextContextValue(source)
	}
}
