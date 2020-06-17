// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"compress/bzip2"
	"io"
)

func TaintStepTest_CompressBzip2NewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader815` into `intoReader389`.

	// Assume that `sourceCQL` has the underlying type of `fromReader815`:
	fromReader815 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader815` to result `intoReader389`
	// (`intoReader389` is now tainted).
	intoReader389 := bzip2.NewReader(fromReader815)

	// Return the tainted `intoReader389`:
	return intoReader389
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
