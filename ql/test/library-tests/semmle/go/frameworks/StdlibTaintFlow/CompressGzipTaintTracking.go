// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"compress/gzip"
	"io"
)

func TaintStepTest_CompressGzipNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader656` into `intoReader414`.

	// Assume that `sourceCQL` has the underlying type of `fromReader656`:
	fromReader656 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader656` to result `intoReader414`
	// (`intoReader414` is now tainted).
	intoReader414, _ := gzip.NewReader(fromReader656)

	// Return the tainted `intoReader414`:
	return intoReader414
}

func TaintStepTest_CompressGzipNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter518` into `intoWriter650`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter518`:
	fromWriter518 := sourceCQL.(*gzip.Writer)

	// Declare `intoWriter650` variable:
	var intoWriter650 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter650`:
	intermediateCQL := gzip.NewWriter(intoWriter650)

	// Extra step (`fromWriter518` taints `intermediateCQL`, which taints `intoWriter650`:
	link(fromWriter518, intermediateCQL)

	// Return the tainted `intoWriter650`:
	return intoWriter650
}

func TaintStepTest_CompressGzipNewWriterLevel_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter784` into `intoWriter957`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter784`:
	fromWriter784 := sourceCQL.(*gzip.Writer)

	// Declare `intoWriter957` variable:
	var intoWriter957 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter957`:
	intermediateCQL, _ := gzip.NewWriterLevel(intoWriter957, 0)

	// Extra step (`fromWriter784` taints `intermediateCQL`, which taints `intoWriter957`:
	link(fromWriter784, intermediateCQL)

	// Return the tainted `intoWriter957`:
	return intoWriter957
}

func TaintStepTest_CompressGzipReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader520` into `intoByte443`.

	// Assume that `sourceCQL` has the underlying type of `fromReader520`:
	fromReader520 := sourceCQL.(gzip.Reader)

	// Declare `intoByte443` variable:
	var intoByte443 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader520` to the argument `intoByte443`
	// (`intoByte443` is now tainted).
	fromReader520.Read(intoByte443)

	// Return the tainted `intoByte443`:
	return intoByte443
}

func TaintStepTest_CompressGzipReaderReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader127` into `intoReader483`.

	// Assume that `sourceCQL` has the underlying type of `fromReader127`:
	fromReader127 := sourceCQL.(io.Reader)

	// Declare `intoReader483` variable:
	var intoReader483 gzip.Reader

	// Call the method that transfers the taint
	// from the parameter `fromReader127` to the receiver `intoReader483`
	// (`intoReader483` is now tainted).
	intoReader483.Reset(fromReader127)

	// Return the tainted `intoReader483`:
	return intoReader483
}

func TaintStepTest_CompressGzipWriterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter989` into `intoWriter982`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter989`:
	fromWriter989 := sourceCQL.(gzip.Writer)

	// Declare `intoWriter982` variable:
	var intoWriter982 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter989` to the argument `intoWriter982`
	// (`intoWriter982` is now tainted).
	fromWriter989.Reset(intoWriter982)

	// Return the tainted `intoWriter982`:
	return intoWriter982
}

func TaintStepTest_CompressGzipWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte417` into `intoWriter584`.

	// Assume that `sourceCQL` has the underlying type of `fromByte417`:
	fromByte417 := sourceCQL.([]byte)

	// Declare `intoWriter584` variable:
	var intoWriter584 gzip.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte417` to the receiver `intoWriter584`
	// (`intoWriter584` is now tainted).
	intoWriter584.Write(fromByte417)

	// Return the tainted `intoWriter584`:
	return intoWriter584
}

func RunAllTaints_CompressGzip() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressGzipNewReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressGzipNewWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressGzipNewWriterLevel_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressGzipReaderRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressGzipReaderReset_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressGzipWriterReset_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressGzipWriterWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
