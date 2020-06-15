package main

import (
	"compress/lzw"
	"io"
)

func TaintStepTest_CompressLzwNewReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader169` into `intoReadCloser320`.

	// Assume that `sourceCQL` has the underlying type of `fromReader169`:
	fromReader169 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader169` to result `intoReadCloser320`
	// (`intoReadCloser320` is now tainted).
	intoReadCloser320 := lzw.NewReader(fromReader169, 0, 0)

	// Sink the tainted `intoReadCloser320`:
	sink(intoReadCloser320)
}

func TaintStepTest_CompressLzwNewWriter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriteCloser259` into `intoWriter440`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser259`:
	fromWriteCloser259 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter440` variable:
	var intoWriter440 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter440`:
	intermediateCQL := lzw.NewWriter(intoWriter440, 0, 0)

	// Extra step (`fromWriteCloser259` taints `intermediateCQL`, which taints `intoWriter440`:
	link(fromWriteCloser259, intermediateCQL)

	// Sink the tainted `intoWriter440`:
	sink(intoWriter440)
}

func RunAllTaints_CompressLzw() {
	{
		source := newSource()
		TaintStepTest_CompressLzwNewReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressLzwNewWriter_B0I0O0(source)
	}
}
