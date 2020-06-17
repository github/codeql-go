package main

import (
	"compress/flate"
	"io"
)

func TaintStepTest_CompressFlateNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader607` into `intoReadCloser311`.

	// Assume that `sourceCQL` has the underlying type of `fromReader607`:
	fromReader607 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader607` to result `intoReadCloser311`
	// (`intoReadCloser311` is now tainted).
	intoReadCloser311 := flate.NewReader(fromReader607)

	// Return the tainted `intoReadCloser311`:
	return intoReadCloser311
}

func TaintStepTest_CompressFlateNewReaderDict_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader456` into `intoReadCloser208`.

	// Assume that `sourceCQL` has the underlying type of `fromReader456`:
	fromReader456 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader456` to result `intoReadCloser208`
	// (`intoReadCloser208` is now tainted).
	intoReadCloser208 := flate.NewReaderDict(fromReader456, nil)

	// Return the tainted `intoReadCloser208`:
	return intoReadCloser208
}

func TaintStepTest_CompressFlateNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter770` into `intoWriter646`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter770`:
	fromWriter770 := sourceCQL.(*flate.Writer)

	// Declare `intoWriter646` variable:
	var intoWriter646 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter646`:
	intermediateCQL, _ := flate.NewWriter(intoWriter646, 0)

	// Extra step (`fromWriter770` taints `intermediateCQL`, which taints `intoWriter646`:
	link(fromWriter770, intermediateCQL)

	// Return the tainted `intoWriter646`:
	return intoWriter646
}

func TaintStepTest_CompressFlateNewWriterDict_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter619` into `intoWriter909`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter619`:
	fromWriter619 := sourceCQL.(*flate.Writer)

	// Declare `intoWriter909` variable:
	var intoWriter909 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter909`:
	intermediateCQL, _ := flate.NewWriterDict(intoWriter909, 0, nil)

	// Extra step (`fromWriter619` taints `intermediateCQL`, which taints `intoWriter909`:
	link(fromWriter619, intermediateCQL)

	// Return the tainted `intoWriter909`:
	return intoWriter909
}

func TaintStepTest_CompressFlateWriterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter758` into `intoWriter304`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter758`:
	fromWriter758 := sourceCQL.(flate.Writer)

	// Declare `intoWriter304` variable:
	var intoWriter304 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter758` to the argument `intoWriter304`
	// (`intoWriter304` is now tainted).
	fromWriter758.Reset(intoWriter304)

	// Return the tainted `intoWriter304`:
	return intoWriter304
}

func TaintStepTest_CompressFlateWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte479` into `intoWriter465`.

	// Assume that `sourceCQL` has the underlying type of `fromByte479`:
	fromByte479 := sourceCQL.([]byte)

	// Declare `intoWriter465` variable:
	var intoWriter465 flate.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte479` to the receiver `intoWriter465`
	// (`intoWriter465` is now tainted).
	intoWriter465.Write(fromByte479)

	// Return the tainted `intoWriter465`:
	return intoWriter465
}

func TaintStepTest_CompressFlateResetterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader313` into `intoResetter651`.

	// Assume that `sourceCQL` has the underlying type of `fromReader313`:
	fromReader313 := sourceCQL.(io.Reader)

	// Declare `intoResetter651` variable:
	var intoResetter651 flate.Resetter

	// Call the method that transfers the taint
	// from the parameter `fromReader313` to the receiver `intoResetter651`
	// (`intoResetter651` is now tainted).
	intoResetter651.Reset(fromReader313, nil)

	// Return the tainted `intoResetter651`:
	return intoResetter651
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
