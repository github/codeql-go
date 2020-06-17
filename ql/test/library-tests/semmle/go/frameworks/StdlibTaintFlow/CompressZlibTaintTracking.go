// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"compress/zlib"
	"io"
)

func TaintStepTest_CompressZlibNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader656` into `intoReadCloser414`.

	// Assume that `sourceCQL` has the underlying type of `fromReader656`:
	fromReader656 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader656` to result `intoReadCloser414`
	// (`intoReadCloser414` is now tainted).
	intoReadCloser414, _ := zlib.NewReader(fromReader656)

	// Return the tainted `intoReadCloser414`:
	return intoReadCloser414
}

func TaintStepTest_CompressZlibNewReaderDict_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader518` into `intoReadCloser650`.

	// Assume that `sourceCQL` has the underlying type of `fromReader518`:
	fromReader518 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader518` to result `intoReadCloser650`
	// (`intoReadCloser650` is now tainted).
	intoReadCloser650, _ := zlib.NewReaderDict(fromReader518, nil)

	// Return the tainted `intoReadCloser650`:
	return intoReadCloser650
}

func TaintStepTest_CompressZlibNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter784` into `intoWriter957`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter784`:
	fromWriter784 := sourceCQL.(*zlib.Writer)

	// Declare `intoWriter957` variable:
	var intoWriter957 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter957`:
	intermediateCQL := zlib.NewWriter(intoWriter957)

	// Extra step (`fromWriter784` taints `intermediateCQL`, which taints `intoWriter957`:
	link(fromWriter784, intermediateCQL)

	// Return the tainted `intoWriter957`:
	return intoWriter957
}

func TaintStepTest_CompressZlibNewWriterLevel_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter520` into `intoWriter443`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter520`:
	fromWriter520 := sourceCQL.(*zlib.Writer)

	// Declare `intoWriter443` variable:
	var intoWriter443 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter443`:
	intermediateCQL, _ := zlib.NewWriterLevel(intoWriter443, 0)

	// Extra step (`fromWriter520` taints `intermediateCQL`, which taints `intoWriter443`:
	link(fromWriter520, intermediateCQL)

	// Return the tainted `intoWriter443`:
	return intoWriter443
}

func TaintStepTest_CompressZlibNewWriterLevelDict_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter127` into `intoWriter483`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter127`:
	fromWriter127 := sourceCQL.(*zlib.Writer)

	// Declare `intoWriter483` variable:
	var intoWriter483 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter483`:
	intermediateCQL, _ := zlib.NewWriterLevelDict(intoWriter483, 0, nil)

	// Extra step (`fromWriter127` taints `intermediateCQL`, which taints `intoWriter483`:
	link(fromWriter127, intermediateCQL)

	// Return the tainted `intoWriter483`:
	return intoWriter483
}

func TaintStepTest_CompressZlibWriterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter989` into `intoWriter982`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter989`:
	fromWriter989 := sourceCQL.(zlib.Writer)

	// Declare `intoWriter982` variable:
	var intoWriter982 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter989` to the argument `intoWriter982`
	// (`intoWriter982` is now tainted).
	fromWriter989.Reset(intoWriter982)

	// Return the tainted `intoWriter982`:
	return intoWriter982
}

func TaintStepTest_CompressZlibWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte417` into `intoWriter584`.

	// Assume that `sourceCQL` has the underlying type of `fromByte417`:
	fromByte417 := sourceCQL.([]byte)

	// Declare `intoWriter584` variable:
	var intoWriter584 zlib.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte417` to the receiver `intoWriter584`
	// (`intoWriter584` is now tainted).
	intoWriter584.Write(fromByte417)

	// Return the tainted `intoWriter584`:
	return intoWriter584
}

func TaintStepTest_CompressZlibResetterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader991` into `intoResetter881`.

	// Assume that `sourceCQL` has the underlying type of `fromReader991`:
	fromReader991 := sourceCQL.(io.Reader)

	// Declare `intoResetter881` variable:
	var intoResetter881 zlib.Resetter

	// Call the method that transfers the taint
	// from the parameter `fromReader991` to the receiver `intoResetter881`
	// (`intoResetter881` is now tainted).
	intoResetter881.Reset(fromReader991, nil)

	// Return the tainted `intoResetter881`:
	return intoResetter881
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
