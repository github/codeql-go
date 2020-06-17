// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"context"
	"time"
)

func TaintStepTest_ContextWithCancel_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext656` into `intoContext414`.

	// Assume that `sourceCQL` has the underlying type of `fromContext656`:
	fromContext656 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext656` to result `intoContext414`
	// (`intoContext414` is now tainted).
	intoContext414, _ := context.WithCancel(fromContext656)

	// Return the tainted `intoContext414`:
	return intoContext414
}

func TaintStepTest_ContextWithDeadline_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext518` into `intoContext650`.

	// Assume that `sourceCQL` has the underlying type of `fromContext518`:
	fromContext518 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext518` to result `intoContext650`
	// (`intoContext650` is now tainted).
	intoContext650, _ := context.WithDeadline(fromContext518, time.Time{})

	// Return the tainted `intoContext650`:
	return intoContext650
}

func TaintStepTest_ContextWithTimeout_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext784` into `intoContext957`.

	// Assume that `sourceCQL` has the underlying type of `fromContext784`:
	fromContext784 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext784` to result `intoContext957`
	// (`intoContext957` is now tainted).
	intoContext957, _ := context.WithTimeout(fromContext784, 0)

	// Return the tainted `intoContext957`:
	return intoContext957
}

func TaintStepTest_ContextWithValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext520` into `intoContext443`.

	// Assume that `sourceCQL` has the underlying type of `fromContext520`:
	fromContext520 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext520` to result `intoContext443`
	// (`intoContext443` is now tainted).
	intoContext443 := context.WithValue(fromContext520, nil, nil)

	// Return the tainted `intoContext443`:
	return intoContext443
}

func TaintStepTest_ContextWithValue_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface127` into `intoContext483`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface127`:
	fromInterface127 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface127` to result `intoContext483`
	// (`intoContext483` is now tainted).
	intoContext483 := context.WithValue(nil, fromInterface127, nil)

	// Return the tainted `intoContext483`:
	return intoContext483
}

func TaintStepTest_ContextWithValue_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface989` into `intoContext982`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface989`:
	fromInterface989 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface989` to result `intoContext982`
	// (`intoContext982` is now tainted).
	intoContext982 := context.WithValue(nil, nil, fromInterface989)

	// Return the tainted `intoContext982`:
	return intoContext982
}

func TaintStepTest_ContextContextValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext417` into `intoInterface584`.

	// Assume that `sourceCQL` has the underlying type of `fromContext417`:
	fromContext417 := sourceCQL.(context.Context)

	// Call the method that transfers the taint
	// from the receiver `fromContext417` to the result `intoInterface584`
	// (`intoInterface584` is now tainted).
	intoInterface584 := fromContext417.Value(nil)

	// Return the tainted `intoInterface584`:
	return intoInterface584
}

func RunAllTaints_Context() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContextWithCancel_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContextWithDeadline_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContextWithTimeout_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContextWithValue_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContextWithValue_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContextWithValue_B0I2O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContextContextValue_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
