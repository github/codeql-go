// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"compress/bzip2"
	"io"
)

func TaintStepTest_CompressBzip2NewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader656` into `intoReader414`.

	// Assume that `sourceCQL` has the underlying type of `fromReader656`:
	fromReader656 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader656` to result `intoReader414`
	// (`intoReader414` is now tainted).
	intoReader414 := bzip2.NewReader(fromReader656)

	// Return the tainted `intoReader414`:
	return intoReader414
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
