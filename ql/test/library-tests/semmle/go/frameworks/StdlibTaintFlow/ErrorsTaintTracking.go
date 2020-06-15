package main

import "errors"

func TaintStepTest_ErrorsAs_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromError223` into `intoInterface786`.

	// Assume that `sourceCQL` has the underlying type of `fromError223`:
	fromError223 := sourceCQL.(error)

	// Declare `intoInterface786` variable:
	var intoInterface786 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromError223` to parameter `intoInterface786`;
	// `intoInterface786` is now tainted.
	errors.As(fromError223, intoInterface786)

	// Sink the tainted `intoInterface786`:
	sink(intoInterface786)
}

func TaintStepTest_ErrorsNew_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString251` into `intoError950`.

	// Assume that `sourceCQL` has the underlying type of `fromString251`:
	fromString251 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString251` to result `intoError950`
	// (`intoError950` is now tainted).
	intoError950 := errors.New(fromString251)

	// Sink the tainted `intoError950`:
	sink(intoError950)
}

func TaintStepTest_ErrorsUnwrap_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromError362` into `intoError720`.

	// Assume that `sourceCQL` has the underlying type of `fromError362`:
	fromError362 := sourceCQL.(error)

	// Call the function that transfers the taint
	// from the parameter `fromError362` to result `intoError720`
	// (`intoError720` is now tainted).
	intoError720 := errors.Unwrap(fromError362)

	// Sink the tainted `intoError720`:
	sink(intoError720)
}

func RunAllTaints_Errors() {
	{
		source := newSource()
		TaintStepTest_ErrorsAs_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ErrorsNew_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ErrorsUnwrap_B0I0O0(source)
	}
}
