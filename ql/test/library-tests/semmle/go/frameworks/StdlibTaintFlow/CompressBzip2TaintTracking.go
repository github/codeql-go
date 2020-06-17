package main

import (
	"compress/bzip2"
	"io"
)

func TaintStepTest_CompressBzip2NewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader996` into `intoReader955`.

	// Assume that `sourceCQL` has the underlying type of `fromReader996`:
	fromReader996 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader996` to result `intoReader955`
	// (`intoReader955` is now tainted).
	intoReader955 := bzip2.NewReader(fromReader996)

	// Return the tainted `intoReader955`:
	return intoReader955
}

func RunAllTaints_CompressBzip2() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressBzip2NewReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
