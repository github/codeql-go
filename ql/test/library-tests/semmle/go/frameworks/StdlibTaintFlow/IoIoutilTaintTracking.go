// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"io/ioutil"
)

func TaintStepTest_IoIoutilNopCloser_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader656` into `intoReadCloser414`.

	// Assume that `sourceCQL` has the underlying type of `fromReader656`:
	fromReader656 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader656` to result `intoReadCloser414`
	// (`intoReadCloser414` is now tainted).
	intoReadCloser414 := ioutil.NopCloser(fromReader656)

	// Return the tainted `intoReadCloser414`:
	return intoReadCloser414
}

func TaintStepTest_IoIoutilReadAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader518` into `intoByte650`.

	// Assume that `sourceCQL` has the underlying type of `fromReader518`:
	fromReader518 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader518` to result `intoByte650`
	// (`intoByte650` is now tainted).
	intoByte650, _ := ioutil.ReadAll(fromReader518)

	// Return the tainted `intoByte650`:
	return intoByte650
}

func RunAllTaints_IoIoutil() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoIoutilNopCloser_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoIoutilReadAll_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
