package main

import "errors"

func TaintStepTest_ErrorsAs(sourceCQL interface{}) {
	// The flow is from `err` into `target`.

	// Assume that `sourceCQL` has the underlying type of `err`:
	err := sourceCQL.(error)

	// Declare `target` variable:
	var target interface{}

	// Call the function that transfers the taint
	// from the parameter `err` to parameter `target`;
	// `target` is now tainted.
	errors.As(err, target)

	// Sink the tainted `target`:
	sink(target)
}

func TaintStepTest_ErrorsNew(sourceCQL interface{}) {
	// The flow is from `text` into `into586`.

	// Assume that `sourceCQL` has the underlying type of `text`:
	text := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `text` to result `into586`
	// (`into586` is now tainted).
	into586 := errors.New(text)

	// Sink the tainted `into586`:
	sink(into586)
}

func TaintStepTest_ErrorsUnwrap(sourceCQL interface{}) {
	// The flow is from `err` into `into361`.

	// Assume that `sourceCQL` has the underlying type of `err`:
	err := sourceCQL.(error)

	// Call the function that transfers the taint
	// from the parameter `err` to result `into361`
	// (`into361` is now tainted).
	into361 := errors.Unwrap(err)

	// Sink the tainted `into361`:
	sink(into361)
}

func RunAllTaints_Errors(v interface{}) {
	{
		source := newSource()
		TaintStepTest_ErrorsAs(source)
	}
	{
		source := newSource()
		TaintStepTest_ErrorsNew(source)
	}
	{
		source := newSource()
		TaintStepTest_ErrorsUnwrap(source)
	}
}
