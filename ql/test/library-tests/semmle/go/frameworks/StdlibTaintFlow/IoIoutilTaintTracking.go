// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"io/ioutil"
)

func TaintStepTest_IoIoutilNopCloser_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader893` into `intoReadCloser497`.

	// Assume that `sourceCQL` has the underlying type of `fromReader893`:
	fromReader893 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader893` to result `intoReadCloser497`
	// (`intoReadCloser497` is now tainted).
	intoReadCloser497 := ioutil.NopCloser(fromReader893)

	// Return the tainted `intoReadCloser497`:
	return intoReadCloser497
}

func TaintStepTest_IoIoutilReadAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader497` into `intoByte620`.

	// Assume that `sourceCQL` has the underlying type of `fromReader497`:
	fromReader497 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader497` to result `intoByte620`
	// (`intoByte620` is now tainted).
	intoByte620, _ := ioutil.ReadAll(fromReader497)

	// Return the tainted `intoByte620`:
	return intoByte620
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
