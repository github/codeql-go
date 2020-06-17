package main

import (
	"io"
	"io/ioutil"
)

func TaintStepTest_IoIoutilNopCloser_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader434` into `intoReadCloser983`.

	// Assume that `sourceCQL` has the underlying type of `fromReader434`:
	fromReader434 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader434` to result `intoReadCloser983`
	// (`intoReadCloser983` is now tainted).
	intoReadCloser983 := ioutil.NopCloser(fromReader434)

	// Return the tainted `intoReadCloser983`:
	return intoReadCloser983
}

func TaintStepTest_IoIoutilReadAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader898` into `intoByte834`.

	// Assume that `sourceCQL` has the underlying type of `fromReader898`:
	fromReader898 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader898` to result `intoByte834`
	// (`intoByte834` is now tainted).
	intoByte834, _ := ioutil.ReadAll(fromReader898)

	// Return the tainted `intoByte834`:
	return intoByte834
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
