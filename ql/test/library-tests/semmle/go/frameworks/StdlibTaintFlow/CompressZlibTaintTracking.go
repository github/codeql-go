package main

import (
	"compress/zlib"
	"io"
)

func TaintStepTest_CompressZlibNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader331` into `intoReadCloser213`.

	// Assume that `sourceCQL` has the underlying type of `fromReader331`:
	fromReader331 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader331` to result `intoReadCloser213`
	// (`intoReadCloser213` is now tainted).
	intoReadCloser213, _ := zlib.NewReader(fromReader331)

	// Return the tainted `intoReadCloser213`:
	return intoReadCloser213
}

func TaintStepTest_CompressZlibNewReaderDict_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader956` into `intoReadCloser245`.

	// Assume that `sourceCQL` has the underlying type of `fromReader956`:
	fromReader956 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader956` to result `intoReadCloser245`
	// (`intoReadCloser245` is now tainted).
	intoReadCloser245, _ := zlib.NewReaderDict(fromReader956, nil)

	// Return the tainted `intoReadCloser245`:
	return intoReadCloser245
}

func TaintStepTest_CompressZlibNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter689` into `intoWriter302`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter689`:
	fromWriter689 := sourceCQL.(*zlib.Writer)

	// Declare `intoWriter302` variable:
	var intoWriter302 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter302`:
	intermediateCQL := zlib.NewWriter(intoWriter302)

	// Extra step (`fromWriter689` taints `intermediateCQL`, which taints `intoWriter302`:
	link(fromWriter689, intermediateCQL)

	// Return the tainted `intoWriter302`:
	return intoWriter302
}

func TaintStepTest_CompressZlibNewWriterLevel_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter973` into `intoWriter586`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter973`:
	fromWriter973 := sourceCQL.(*zlib.Writer)

	// Declare `intoWriter586` variable:
	var intoWriter586 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter586`:
	intermediateCQL, _ := zlib.NewWriterLevel(intoWriter586, 0)

	// Extra step (`fromWriter973` taints `intermediateCQL`, which taints `intoWriter586`:
	link(fromWriter973, intermediateCQL)

	// Return the tainted `intoWriter586`:
	return intoWriter586
}

func TaintStepTest_CompressZlibNewWriterLevelDict_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter900` into `intoWriter205`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter900`:
	fromWriter900 := sourceCQL.(*zlib.Writer)

	// Declare `intoWriter205` variable:
	var intoWriter205 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter205`:
	intermediateCQL, _ := zlib.NewWriterLevelDict(intoWriter205, 0, nil)

	// Extra step (`fromWriter900` taints `intermediateCQL`, which taints `intoWriter205`:
	link(fromWriter900, intermediateCQL)

	// Return the tainted `intoWriter205`:
	return intoWriter205
}

func TaintStepTest_CompressZlibWriterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter113` into `intoWriter123`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter113`:
	fromWriter113 := sourceCQL.(zlib.Writer)

	// Declare `intoWriter123` variable:
	var intoWriter123 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter113` to the argument `intoWriter123`
	// (`intoWriter123` is now tainted).
	fromWriter113.Reset(intoWriter123)

	// Return the tainted `intoWriter123`:
	return intoWriter123
}

func TaintStepTest_CompressZlibWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte994` into `intoWriter977`.

	// Assume that `sourceCQL` has the underlying type of `fromByte994`:
	fromByte994 := sourceCQL.([]byte)

	// Declare `intoWriter977` variable:
	var intoWriter977 zlib.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte994` to the receiver `intoWriter977`
	// (`intoWriter977` is now tainted).
	intoWriter977.Write(fromByte994)

	// Return the tainted `intoWriter977`:
	return intoWriter977
}

func TaintStepTest_CompressZlibResetterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader788` into `intoResetter754`.

	// Assume that `sourceCQL` has the underlying type of `fromReader788`:
	fromReader788 := sourceCQL.(io.Reader)

	// Declare `intoResetter754` variable:
	var intoResetter754 zlib.Resetter

	// Call the method that transfers the taint
	// from the parameter `fromReader788` to the receiver `intoResetter754`
	// (`intoResetter754` is now tainted).
	intoResetter754.Reset(fromReader788, nil)

	// Return the tainted `intoResetter754`:
	return intoResetter754
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
