// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"context"
	"time"
)

func TaintStepTest_ContextWithCancel_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext496` into `intoContext728`.

	// Assume that `sourceCQL` has the underlying type of `fromContext496`:
	fromContext496 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext496` to result `intoContext728`
	// (`intoContext728` is now tainted).
	intoContext728, _ := context.WithCancel(fromContext496)

	// Return the tainted `intoContext728`:
	return intoContext728
}

func TaintStepTest_ContextWithDeadline_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext132` into `intoContext344`.

	// Assume that `sourceCQL` has the underlying type of `fromContext132`:
	fromContext132 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext132` to result `intoContext344`
	// (`intoContext344` is now tainted).
	intoContext344, _ := context.WithDeadline(fromContext132, time.Time{})

	// Return the tainted `intoContext344`:
	return intoContext344
}

func TaintStepTest_ContextWithTimeout_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext511` into `intoContext162`.

	// Assume that `sourceCQL` has the underlying type of `fromContext511`:
	fromContext511 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext511` to result `intoContext162`
	// (`intoContext162` is now tainted).
	intoContext162, _ := context.WithTimeout(fromContext511, 0)

	// Return the tainted `intoContext162`:
	return intoContext162
}

func TaintStepTest_ContextWithValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext235` into `intoContext220`.

	// Assume that `sourceCQL` has the underlying type of `fromContext235`:
	fromContext235 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext235` to result `intoContext220`
	// (`intoContext220` is now tainted).
	intoContext220 := context.WithValue(fromContext235, nil, nil)

	// Return the tainted `intoContext220`:
	return intoContext220
}

func TaintStepTest_ContextWithValue_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface151` into `intoContext600`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface151`:
	fromInterface151 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface151` to result `intoContext600`
	// (`intoContext600` is now tainted).
	intoContext600 := context.WithValue(nil, fromInterface151, nil)

	// Return the tainted `intoContext600`:
	return intoContext600
}

func TaintStepTest_ContextWithValue_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface914` into `intoContext279`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface914`:
	fromInterface914 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface914` to result `intoContext279`
	// (`intoContext279` is now tainted).
	intoContext279 := context.WithValue(nil, nil, fromInterface914)

	// Return the tainted `intoContext279`:
	return intoContext279
}

func TaintStepTest_ContextContextValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext232` into `intoInterface952`.

	// Assume that `sourceCQL` has the underlying type of `fromContext232`:
	fromContext232 := sourceCQL.(context.Context)

	// Call the method that transfers the taint
	// from the receiver `fromContext232` to the result `intoInterface952`
	// (`intoInterface952` is now tainted).
	intoInterface952 := fromContext232.Value(nil)

	// Return the tainted `intoInterface952`:
	return intoInterface952
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
