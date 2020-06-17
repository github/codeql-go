// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"compress/flate"
	"io"
)

func TaintStepTest_CompressFlateNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader459` into `intoReadCloser566`.

	// Assume that `sourceCQL` has the underlying type of `fromReader459`:
	fromReader459 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader459` to result `intoReadCloser566`
	// (`intoReadCloser566` is now tainted).
	intoReadCloser566 := flate.NewReader(fromReader459)

	// Return the tainted `intoReadCloser566`:
	return intoReadCloser566
}

func TaintStepTest_CompressFlateNewReaderDict_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader180` into `intoReadCloser279`.

	// Assume that `sourceCQL` has the underlying type of `fromReader180`:
	fromReader180 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader180` to result `intoReadCloser279`
	// (`intoReadCloser279` is now tainted).
	intoReadCloser279 := flate.NewReaderDict(fromReader180, nil)

	// Return the tainted `intoReadCloser279`:
	return intoReadCloser279
}

func TaintStepTest_CompressFlateNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter923` into `intoWriter829`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter923`:
	fromWriter923 := sourceCQL.(*flate.Writer)

	// Declare `intoWriter829` variable:
	var intoWriter829 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter829`:
	intermediateCQL, _ := flate.NewWriter(intoWriter829, 0)

	// Extra step (`fromWriter923` taints `intermediateCQL`, which taints `intoWriter829`:
	link(fromWriter923, intermediateCQL)

	// Return the tainted `intoWriter829`:
	return intoWriter829
}

func TaintStepTest_CompressFlateNewWriterDict_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter216` into `intoWriter811`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter216`:
	fromWriter216 := sourceCQL.(*flate.Writer)

	// Declare `intoWriter811` variable:
	var intoWriter811 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter811`:
	intermediateCQL, _ := flate.NewWriterDict(intoWriter811, 0, nil)

	// Extra step (`fromWriter216` taints `intermediateCQL`, which taints `intoWriter811`:
	link(fromWriter216, intermediateCQL)

	// Return the tainted `intoWriter811`:
	return intoWriter811
}

func TaintStepTest_CompressFlateWriterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter690` into `intoWriter558`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter690`:
	fromWriter690 := sourceCQL.(flate.Writer)

	// Declare `intoWriter558` variable:
	var intoWriter558 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter690` to the argument `intoWriter558`
	// (`intoWriter558` is now tainted).
	fromWriter690.Reset(intoWriter558)

	// Return the tainted `intoWriter558`:
	return intoWriter558
}

func TaintStepTest_CompressFlateWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte750` into `intoWriter834`.

	// Assume that `sourceCQL` has the underlying type of `fromByte750`:
	fromByte750 := sourceCQL.([]byte)

	// Declare `intoWriter834` variable:
	var intoWriter834 flate.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte750` to the receiver `intoWriter834`
	// (`intoWriter834` is now tainted).
	intoWriter834.Write(fromByte750)

	// Return the tainted `intoWriter834`:
	return intoWriter834
}

func TaintStepTest_CompressFlateResetterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader879` into `intoResetter590`.

	// Assume that `sourceCQL` has the underlying type of `fromReader879`:
	fromReader879 := sourceCQL.(io.Reader)

	// Declare `intoResetter590` variable:
	var intoResetter590 flate.Resetter

	// Call the method that transfers the taint
	// from the parameter `fromReader879` to the receiver `intoResetter590`
	// (`intoResetter590` is now tainted).
	intoResetter590.Reset(fromReader879, nil)

	// Return the tainted `intoResetter590`:
	return intoResetter590
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
