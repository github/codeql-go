// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"compress/zlib"
	"io"
)

func TaintStepTest_CompressZlibNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader652` into `intoReadCloser713`.

	// Assume that `sourceCQL` has the underlying type of `fromReader652`:
	fromReader652 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader652` to result `intoReadCloser713`
	// (`intoReadCloser713` is now tainted).
	intoReadCloser713, _ := zlib.NewReader(fromReader652)

	// Return the tainted `intoReadCloser713`:
	return intoReadCloser713
}

func TaintStepTest_CompressZlibNewReaderDict_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader612` into `intoReadCloser160`.

	// Assume that `sourceCQL` has the underlying type of `fromReader612`:
	fromReader612 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader612` to result `intoReadCloser160`
	// (`intoReadCloser160` is now tainted).
	intoReadCloser160, _ := zlib.NewReaderDict(fromReader612, nil)

	// Return the tainted `intoReadCloser160`:
	return intoReadCloser160
}

func TaintStepTest_CompressZlibNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter616` into `intoWriter631`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter616`:
	fromWriter616 := sourceCQL.(*zlib.Writer)

	// Declare `intoWriter631` variable:
	var intoWriter631 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter631`:
	intermediateCQL := zlib.NewWriter(intoWriter631)

	// Extra step (`fromWriter616` taints `intermediateCQL`, which taints `intoWriter631`:
	link(fromWriter616, intermediateCQL)

	// Return the tainted `intoWriter631`:
	return intoWriter631
}

func TaintStepTest_CompressZlibNewWriterLevel_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter400` into `intoWriter438`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter400`:
	fromWriter400 := sourceCQL.(*zlib.Writer)

	// Declare `intoWriter438` variable:
	var intoWriter438 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter438`:
	intermediateCQL, _ := zlib.NewWriterLevel(intoWriter438, 0)

	// Extra step (`fromWriter400` taints `intermediateCQL`, which taints `intoWriter438`:
	link(fromWriter400, intermediateCQL)

	// Return the tainted `intoWriter438`:
	return intoWriter438
}

func TaintStepTest_CompressZlibNewWriterLevelDict_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter807` into `intoWriter834`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter807`:
	fromWriter807 := sourceCQL.(*zlib.Writer)

	// Declare `intoWriter834` variable:
	var intoWriter834 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter834`:
	intermediateCQL, _ := zlib.NewWriterLevelDict(intoWriter834, 0, nil)

	// Extra step (`fromWriter807` taints `intermediateCQL`, which taints `intoWriter834`:
	link(fromWriter807, intermediateCQL)

	// Return the tainted `intoWriter834`:
	return intoWriter834
}

func TaintStepTest_CompressZlibWriterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter723` into `intoWriter827`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter723`:
	fromWriter723 := sourceCQL.(zlib.Writer)

	// Declare `intoWriter827` variable:
	var intoWriter827 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter723` to the argument `intoWriter827`
	// (`intoWriter827` is now tainted).
	fromWriter723.Reset(intoWriter827)

	// Return the tainted `intoWriter827`:
	return intoWriter827
}

func TaintStepTest_CompressZlibWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte636` into `intoWriter561`.

	// Assume that `sourceCQL` has the underlying type of `fromByte636`:
	fromByte636 := sourceCQL.([]byte)

	// Declare `intoWriter561` variable:
	var intoWriter561 zlib.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte636` to the receiver `intoWriter561`
	// (`intoWriter561` is now tainted).
	intoWriter561.Write(fromByte636)

	// Return the tainted `intoWriter561`:
	return intoWriter561
}

func TaintStepTest_CompressZlibResetterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader150` into `intoResetter243`.

	// Assume that `sourceCQL` has the underlying type of `fromReader150`:
	fromReader150 := sourceCQL.(io.Reader)

	// Declare `intoResetter243` variable:
	var intoResetter243 zlib.Resetter

	// Call the method that transfers the taint
	// from the parameter `fromReader150` to the receiver `intoResetter243`
	// (`intoResetter243` is now tainted).
	intoResetter243.Reset(fromReader150, nil)

	// Return the tainted `intoResetter243`:
	return intoResetter243
}

func RunAllTaints_CompressZlib() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressZlibNewReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressZlibNewReaderDict_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressZlibNewWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressZlibNewWriterLevel_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressZlibNewWriterLevelDict_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressZlibWriterReset_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressZlibWriterWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressZlibResetterReset_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
