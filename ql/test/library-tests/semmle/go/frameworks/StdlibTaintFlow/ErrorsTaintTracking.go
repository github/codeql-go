// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "errors"

func TaintStepTest_ErrorsAs_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromError637` into `intoInterface320`.

	// Assume that `sourceCQL` has the underlying type of `fromError637`:
	fromError637 := sourceCQL.(error)

	// Declare `intoInterface320` variable:
	var intoInterface320 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromError637` to parameter `intoInterface320`;
	// `intoInterface320` is now tainted.
	errors.As(fromError637, intoInterface320)

	// Return the tainted `intoInterface320`:
	return intoInterface320
}

func TaintStepTest_ErrorsNew_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString379` into `intoError522`.

	// Assume that `sourceCQL` has the underlying type of `fromString379`:
	fromString379 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString379` to result `intoError522`
	// (`intoError522` is now tainted).
	intoError522 := errors.New(fromString379)

	// Return the tainted `intoError522`:
	return intoError522
}

func TaintStepTest_ErrorsUnwrap_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromError804` into `intoError872`.

	// Assume that `sourceCQL` has the underlying type of `fromError804`:
	fromError804 := sourceCQL.(error)

	// Call the function that transfers the taint
	// from the parameter `fromError804` to result `intoError872`
	// (`intoError872` is now tainted).
	intoError872 := errors.Unwrap(fromError804)

	// Return the tainted `intoError872`:
	return intoError872
}

func RunAllTaints_Errors() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ErrorsAs_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ErrorsNew_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ErrorsUnwrap_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
