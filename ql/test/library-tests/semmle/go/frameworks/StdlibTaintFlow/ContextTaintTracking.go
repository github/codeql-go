package main

import (
	"context"
	"time"
)

func TaintStepTest_ContextWithCancel_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext587` into `intoContext566`.

	// Assume that `sourceCQL` has the underlying type of `fromContext587`:
	fromContext587 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext587` to result `intoContext566`
	// (`intoContext566` is now tainted).
	intoContext566, _ := context.WithCancel(fromContext587)

	// Return the tainted `intoContext566`:
	return intoContext566
}

func TaintStepTest_ContextWithDeadline_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext877` into `intoContext206`.

	// Assume that `sourceCQL` has the underlying type of `fromContext877`:
	fromContext877 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext877` to result `intoContext206`
	// (`intoContext206` is now tainted).
	intoContext206, _ := context.WithDeadline(fromContext877, time.Time{})

	// Return the tainted `intoContext206`:
	return intoContext206
}

func TaintStepTest_ContextWithTimeout_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext848` into `intoContext136`.

	// Assume that `sourceCQL` has the underlying type of `fromContext848`:
	fromContext848 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext848` to result `intoContext136`
	// (`intoContext136` is now tainted).
	intoContext136, _ := context.WithTimeout(fromContext848, 0)

	// Return the tainted `intoContext136`:
	return intoContext136
}

func TaintStepTest_ContextWithValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext142` into `intoContext283`.

	// Assume that `sourceCQL` has the underlying type of `fromContext142`:
	fromContext142 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext142` to result `intoContext283`
	// (`intoContext283` is now tainted).
	intoContext283 := context.WithValue(fromContext142, nil, nil)

	// Return the tainted `intoContext283`:
	return intoContext283
}

func TaintStepTest_ContextWithValue_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface762` into `intoContext214`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface762`:
	fromInterface762 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface762` to result `intoContext214`
	// (`intoContext214` is now tainted).
	intoContext214 := context.WithValue(nil, fromInterface762, nil)

	// Return the tainted `intoContext214`:
	return intoContext214
}

func TaintStepTest_ContextWithValue_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface822` into `intoContext183`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface822`:
	fromInterface822 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface822` to result `intoContext183`
	// (`intoContext183` is now tainted).
	intoContext183 := context.WithValue(nil, nil, fromInterface822)

	// Return the tainted `intoContext183`:
	return intoContext183
}

func TaintStepTest_ContextContextValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext712` into `intoInterface772`.

	// Assume that `sourceCQL` has the underlying type of `fromContext712`:
	fromContext712 := sourceCQL.(context.Context)

	// Call the method that transfers the taint
	// from the receiver `fromContext712` to the result `intoInterface772`
	// (`intoInterface772` is now tainted).
	intoInterface772 := fromContext712.Value(nil)

	// Return the tainted `intoInterface772`:
	return intoInterface772
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
