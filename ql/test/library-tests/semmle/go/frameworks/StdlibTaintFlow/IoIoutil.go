package main

import (
	"io"
	"io/ioutil"
)

func TaintStepTest_IoIoutilNopCloser(sourceCQL interface{}) {
	// The flow is from `r` into `into969`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into969`
	// (`into969` is now tainted).
	into969 := ioutil.NopCloser(r)

	// Sink the tainted `into969`:
	sink(into969)
}

func TaintStepTest_IoIoutilReadAll(sourceCQL interface{}) {
	// The flow is from `r` into `into205`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into205`
	// (`into205` is now tainted).
	into205, _ := ioutil.ReadAll(r)

	// Sink the tainted `into205`:
	sink(into205)
}

func RunAllTaints_IoIoutil(v interface{}) {
	{
		source := newSource()
		TaintStepTest_IoIoutilNopCloser(source)
	}
	{
		source := newSource()
		TaintStepTest_IoIoutilReadAll(source)
	}
}
