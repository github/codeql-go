// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"compress/lzw"
	"io"
)

func TaintStepTest_CompressLzwNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader264` into `intoReadCloser669`.

	// Assume that `sourceCQL` has the underlying type of `fromReader264`:
	fromReader264 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader264` to result `intoReadCloser669`
	// (`intoReadCloser669` is now tainted).
	intoReadCloser669 := lzw.NewReader(fromReader264, 0, 0)

	// Return the tainted `intoReadCloser669`:
	return intoReadCloser669
}

func TaintStepTest_CompressLzwNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser390` into `intoWriter342`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser390`:
	fromWriteCloser390 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter342` variable:
	var intoWriter342 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter342`:
	intermediateCQL := lzw.NewWriter(intoWriter342, 0, 0)

	// Extra step (`fromWriteCloser390` taints `intermediateCQL`, which taints `intoWriter342`:
	link(fromWriteCloser390, intermediateCQL)

	// Return the tainted `intoWriter342`:
	return intoWriter342
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
