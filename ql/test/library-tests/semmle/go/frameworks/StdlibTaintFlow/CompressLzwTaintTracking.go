package main

import (
	"compress/lzw"
	"io"
)

func TaintStepTest_CompressLzwNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader548` into `intoReadCloser786`.

	// Assume that `sourceCQL` has the underlying type of `fromReader548`:
	fromReader548 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader548` to result `intoReadCloser786`
	// (`intoReadCloser786` is now tainted).
	intoReadCloser786 := lzw.NewReader(fromReader548, 0, 0)

	// Return the tainted `intoReadCloser786`:
	return intoReadCloser786
}

func TaintStepTest_CompressLzwNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser308` into `intoWriter185`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser308`:
	fromWriteCloser308 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter185` variable:
	var intoWriter185 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter185`:
	intermediateCQL := lzw.NewWriter(intoWriter185, 0, 0)

	// Extra step (`fromWriteCloser308` taints `intermediateCQL`, which taints `intoWriter185`:
	link(fromWriteCloser308, intermediateCQL)

	// Return the tainted `intoWriter185`:
	return intoWriter185
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
