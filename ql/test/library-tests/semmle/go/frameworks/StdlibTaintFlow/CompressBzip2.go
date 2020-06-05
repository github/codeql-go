package main

import (
	"compress/bzip2"
	"io"
)

func TaintStepTest_CompressBzip2NewReader(sourceCQL interface{}) {
	// The flow is from `r` into `into231`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into231`
	// (`into231` is now tainted).
	into231 := bzip2.NewReader(r)

	// Sink the tainted `into231`:
	sink(into231)
}

func RunAllTaints_CompressBzip2(v interface{}) {
	{
		source := newSource()
		TaintStepTest_CompressBzip2NewReader(source)
	}
}
