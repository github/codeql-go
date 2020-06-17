package main

import "errors"

func TaintStepTest_ErrorsAs_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromError773` into `intoInterface485`.

	// Assume that `sourceCQL` has the underlying type of `fromError773`:
	fromError773 := sourceCQL.(error)

	// Declare `intoInterface485` variable:
	var intoInterface485 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromError773` to parameter `intoInterface485`;
	// `intoInterface485` is now tainted.
	errors.As(fromError773, intoInterface485)

	// Return the tainted `intoInterface485`:
	return intoInterface485
}

func TaintStepTest_ErrorsNew_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString273` into `intoError164`.

	// Assume that `sourceCQL` has the underlying type of `fromString273`:
	fromString273 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString273` to result `intoError164`
	// (`intoError164` is now tainted).
	intoError164 := errors.New(fromString273)

	// Return the tainted `intoError164`:
	return intoError164
}

func TaintStepTest_ErrorsUnwrap_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromError130` into `intoError541`.

	// Assume that `sourceCQL` has the underlying type of `fromError130`:
	fromError130 := sourceCQL.(error)

	// Call the function that transfers the taint
	// from the parameter `fromError130` to result `intoError541`
	// (`intoError541` is now tainted).
	intoError541 := errors.Unwrap(fromError130)

	// Return the tainted `intoError541`:
	return intoError541
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
