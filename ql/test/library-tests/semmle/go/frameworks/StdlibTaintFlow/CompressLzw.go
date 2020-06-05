package main

import (
	"compress/lzw"
	"io"
)

func TaintStepTest_CompressLzwNewReader(sourceCQL interface{}) {
	// The flow is from `r` into `into222`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into222`
	// (`into222` is now tainted).
	into222 := lzw.NewReader(r, 0, 0)

	// Sink the tainted `into222`:
	sink(into222)
}

func TaintStepTest_CompressLzwNewWriter(sourceCQL interface{}) {
	// The flow is from `from499` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from499`:
	from499 := sourceCQL.(io.WriteCloser)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := lzw.NewWriter(w, 0, 0)

	// Extra step (`from499` taints `intermediateCQL`, which taints `w`:
	link(from499, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func RunAllTaints_CompressLzw(v interface{}) {
	{
		source := newSource()
		TaintStepTest_CompressLzwNewReader(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressLzwNewWriter(source)
	}
}
