// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "errors"

func TaintStepTest_ErrorsAs_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromError656` into `intoInterface414`.

	// Assume that `sourceCQL` has the underlying type of `fromError656`:
	fromError656 := sourceCQL.(error)

	// Declare `intoInterface414` variable:
	var intoInterface414 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromError656` to parameter `intoInterface414`;
	// `intoInterface414` is now tainted.
	errors.As(fromError656, intoInterface414)

	// Return the tainted `intoInterface414`:
	return intoInterface414
}

func TaintStepTest_ErrorsNew_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString518` into `intoError650`.

	// Assume that `sourceCQL` has the underlying type of `fromString518`:
	fromString518 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString518` to result `intoError650`
	// (`intoError650` is now tainted).
	intoError650 := errors.New(fromString518)

	// Return the tainted `intoError650`:
	return intoError650
}

func TaintStepTest_ErrorsUnwrap_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromError784` into `intoError957`.

	// Assume that `sourceCQL` has the underlying type of `fromError784`:
	fromError784 := sourceCQL.(error)

	// Call the function that transfers the taint
	// from the parameter `fromError784` to result `intoError957`
	// (`intoError957` is now tainted).
	intoError957 := errors.Unwrap(fromError784)

	// Return the tainted `intoError957`:
	return intoError957
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
