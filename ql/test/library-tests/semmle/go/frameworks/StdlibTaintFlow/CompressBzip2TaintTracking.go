package main

import (
	"compress/bzip2"
	"io"
)

func TaintStepTest_CompressBzip2NewReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader995` into `intoReader635`.

	// Assume that `sourceCQL` has the underlying type of `fromReader995`:
	fromReader995 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader995` to result `intoReader635`
	// (`intoReader635` is now tainted).
	intoReader635 := bzip2.NewReader(fromReader995)

	// Sink the tainted `intoReader635`:
	sink(intoReader635)
}

func RunAllTaints_CompressBzip2() {
	{
		source := newSource()
		TaintStepTest_CompressBzip2NewReader_B0I0O0(source)
	}
}
