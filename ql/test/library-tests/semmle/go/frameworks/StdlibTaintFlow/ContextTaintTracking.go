package main

import (
	"context"
	"time"
)

func TaintStepTest_ContextWithCancel_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromContext387` into `intoContext115`.

	// Assume that `sourceCQL` has the underlying type of `fromContext387`:
	fromContext387 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext387` to result `intoContext115`
	// (`intoContext115` is now tainted).
	intoContext115, _ := context.WithCancel(fromContext387)

	// Sink the tainted `intoContext115`:
	sink(intoContext115)
}

func TaintStepTest_ContextWithDeadline_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromContext264` into `intoContext572`.

	// Assume that `sourceCQL` has the underlying type of `fromContext264`:
	fromContext264 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext264` to result `intoContext572`
	// (`intoContext572` is now tainted).
	intoContext572, _ := context.WithDeadline(fromContext264, time.Time{})

	// Sink the tainted `intoContext572`:
	sink(intoContext572)
}

func TaintStepTest_ContextWithTimeout_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromContext402` into `intoContext921`.

	// Assume that `sourceCQL` has the underlying type of `fromContext402`:
	fromContext402 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext402` to result `intoContext921`
	// (`intoContext921` is now tainted).
	intoContext921, _ := context.WithTimeout(fromContext402, 0)

	// Sink the tainted `intoContext921`:
	sink(intoContext921)
}

func TaintStepTest_ContextWithValue_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromContext772` into `intoContext180`.

	// Assume that `sourceCQL` has the underlying type of `fromContext772`:
	fromContext772 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext772` to result `intoContext180`
	// (`intoContext180` is now tainted).
	intoContext180 := context.WithValue(fromContext772, nil, nil)

	// Sink the tainted `intoContext180`:
	sink(intoContext180)
}

func TaintStepTest_ContextWithValue_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromInterface146` into `intoContext339`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface146`:
	fromInterface146 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface146` to result `intoContext339`
	// (`intoContext339` is now tainted).
	intoContext339 := context.WithValue(nil, fromInterface146, nil)

	// Sink the tainted `intoContext339`:
	sink(intoContext339)
}

func TaintStepTest_ContextWithValue_B0I2O0(sourceCQL interface{}) {
	// The flow is from `fromInterface651` into `intoContext426`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface651`:
	fromInterface651 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface651` to result `intoContext426`
	// (`intoContext426` is now tainted).
	intoContext426 := context.WithValue(nil, nil, fromInterface651)

	// Sink the tainted `intoContext426`:
	sink(intoContext426)
}

func TaintStepTest_ContextContextValue_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromContext812` into `intoInterface649`.

	// Assume that `sourceCQL` has the underlying type of `fromContext812`:
	fromContext812 := sourceCQL.(context.Context)

	// Call the method that transfers the taint
	// from the receiver `fromContext812` to the result `intoInterface649`
	// (`intoInterface649` is now tainted).
	intoInterface649 := fromContext812.Value(nil)

	// Sink the tainted `intoInterface649`:
	sink(intoInterface649)
}

func RunAllTaints_Context() {
	{
		source := newSource()
		TaintStepTest_ContextWithCancel_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContextWithDeadline_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContextWithTimeout_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContextWithValue_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContextWithValue_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContextWithValue_B0I2O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContextContextValue_B0I0O0(source)
	}
}
