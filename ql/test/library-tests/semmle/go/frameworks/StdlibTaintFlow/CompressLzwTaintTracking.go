// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"compress/lzw"
	"io"
)

func TaintStepTest_CompressLzwNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader656` into `intoReadCloser414`.

	// Assume that `sourceCQL` has the underlying type of `fromReader656`:
	fromReader656 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader656` to result `intoReadCloser414`
	// (`intoReadCloser414` is now tainted).
	intoReadCloser414 := lzw.NewReader(fromReader656, 0, 0)

	// Return the tainted `intoReadCloser414`:
	return intoReadCloser414
}

func TaintStepTest_CompressLzwNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser518` into `intoWriter650`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser518`:
	fromWriteCloser518 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter650` variable:
	var intoWriter650 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter650`:
	intermediateCQL := lzw.NewWriter(intoWriter650, 0, 0)

	// Extra step (`fromWriteCloser518` taints `intermediateCQL`, which taints `intoWriter650`:
	link(fromWriteCloser518, intermediateCQL)

	// Return the tainted `intoWriter650`:
	return intoWriter650
}

func RunAllTaints_CompressLzw() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressLzwNewReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressLzwNewWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
