// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"compress/flate"
	"io"
)

func TaintStepTest_CompressFlateNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader656` into `intoReadCloser414`.

	// Assume that `sourceCQL` has the underlying type of `fromReader656`:
	fromReader656 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader656` to result `intoReadCloser414`
	// (`intoReadCloser414` is now tainted).
	intoReadCloser414 := flate.NewReader(fromReader656)

	// Return the tainted `intoReadCloser414`:
	return intoReadCloser414
}

func TaintStepTest_CompressFlateNewReaderDict_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader518` into `intoReadCloser650`.

	// Assume that `sourceCQL` has the underlying type of `fromReader518`:
	fromReader518 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader518` to result `intoReadCloser650`
	// (`intoReadCloser650` is now tainted).
	intoReadCloser650 := flate.NewReaderDict(fromReader518, nil)

	// Return the tainted `intoReadCloser650`:
	return intoReadCloser650
}

func TaintStepTest_CompressFlateNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter784` into `intoWriter957`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter784`:
	fromWriter784 := sourceCQL.(*flate.Writer)

	// Declare `intoWriter957` variable:
	var intoWriter957 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter957`:
	intermediateCQL, _ := flate.NewWriter(intoWriter957, 0)

	// Extra step (`fromWriter784` taints `intermediateCQL`, which taints `intoWriter957`:
	link(fromWriter784, intermediateCQL)

	// Return the tainted `intoWriter957`:
	return intoWriter957
}

func TaintStepTest_CompressFlateNewWriterDict_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter520` into `intoWriter443`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter520`:
	fromWriter520 := sourceCQL.(*flate.Writer)

	// Declare `intoWriter443` variable:
	var intoWriter443 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter443`:
	intermediateCQL, _ := flate.NewWriterDict(intoWriter443, 0, nil)

	// Extra step (`fromWriter520` taints `intermediateCQL`, which taints `intoWriter443`:
	link(fromWriter520, intermediateCQL)

	// Return the tainted `intoWriter443`:
	return intoWriter443
}

func TaintStepTest_CompressFlateWriterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter127` into `intoWriter483`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter127`:
	fromWriter127 := sourceCQL.(flate.Writer)

	// Declare `intoWriter483` variable:
	var intoWriter483 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter127` to the argument `intoWriter483`
	// (`intoWriter483` is now tainted).
	fromWriter127.Reset(intoWriter483)

	// Return the tainted `intoWriter483`:
	return intoWriter483
}

func TaintStepTest_CompressFlateWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte989` into `intoWriter982`.

	// Assume that `sourceCQL` has the underlying type of `fromByte989`:
	fromByte989 := sourceCQL.([]byte)

	// Declare `intoWriter982` variable:
	var intoWriter982 flate.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte989` to the receiver `intoWriter982`
	// (`intoWriter982` is now tainted).
	intoWriter982.Write(fromByte989)

	// Return the tainted `intoWriter982`:
	return intoWriter982
}

func TaintStepTest_CompressFlateResetterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader417` into `intoResetter584`.

	// Assume that `sourceCQL` has the underlying type of `fromReader417`:
	fromReader417 := sourceCQL.(io.Reader)

	// Declare `intoResetter584` variable:
	var intoResetter584 flate.Resetter

	// Call the method that transfers the taint
	// from the parameter `fromReader417` to the receiver `intoResetter584`
	// (`intoResetter584` is now tainted).
	intoResetter584.Reset(fromReader417, nil)

	// Return the tainted `intoResetter584`:
	return intoResetter584
}

func RunAllTaints_CompressFlate() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressFlateNewReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressFlateNewReaderDict_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressFlateNewWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressFlateNewWriterDict_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressFlateWriterReset_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressFlateWriterWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressFlateResetterReset_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
