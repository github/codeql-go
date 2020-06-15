package main

import (
	"compress/flate"
	"io"
)

func TaintStepTest_CompressFlateNewReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader454` into `intoReadCloser518`.

	// Assume that `sourceCQL` has the underlying type of `fromReader454`:
	fromReader454 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader454` to result `intoReadCloser518`
	// (`intoReadCloser518` is now tainted).
	intoReadCloser518 := flate.NewReader(fromReader454)

	// Sink the tainted `intoReadCloser518`:
	sink(intoReadCloser518)
}

func TaintStepTest_CompressFlateNewReaderDict_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader327` into `intoReadCloser903`.

	// Assume that `sourceCQL` has the underlying type of `fromReader327`:
	fromReader327 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader327` to result `intoReadCloser903`
	// (`intoReadCloser903` is now tainted).
	intoReadCloser903 := flate.NewReaderDict(fromReader327, nil)

	// Sink the tainted `intoReadCloser903`:
	sink(intoReadCloser903)
}

func TaintStepTest_CompressFlateNewWriter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter775` into `intoWriter251`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter775`:
	fromWriter775 := sourceCQL.(*flate.Writer)

	// Declare `intoWriter251` variable:
	var intoWriter251 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter251`:
	intermediateCQL, _ := flate.NewWriter(intoWriter251, 0)

	// Extra step (`fromWriter775` taints `intermediateCQL`, which taints `intoWriter251`:
	link(fromWriter775, intermediateCQL)

	// Sink the tainted `intoWriter251`:
	sink(intoWriter251)
}

func TaintStepTest_CompressFlateNewWriterDict_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter408` into `intoWriter727`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter408`:
	fromWriter408 := sourceCQL.(*flate.Writer)

	// Declare `intoWriter727` variable:
	var intoWriter727 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter727`:
	intermediateCQL, _ := flate.NewWriterDict(intoWriter727, 0, nil)

	// Extra step (`fromWriter408` taints `intermediateCQL`, which taints `intoWriter727`:
	link(fromWriter408, intermediateCQL)

	// Sink the tainted `intoWriter727`:
	sink(intoWriter727)
}

func TaintStepTest_CompressFlateWriterReset_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter815` into `intoWriter237`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter815`:
	fromWriter815 := sourceCQL.(flate.Writer)

	// Declare `intoWriter237` variable:
	var intoWriter237 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter815` to the argument `intoWriter237`
	// (`intoWriter237` is now tainted).
	fromWriter815.Reset(intoWriter237)

	// Sink the tainted `intoWriter237`:
	sink(intoWriter237)
}

func TaintStepTest_CompressFlateWriterWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte602` into `intoWriter559`.

	// Assume that `sourceCQL` has the underlying type of `fromByte602`:
	fromByte602 := sourceCQL.([]byte)

	// Declare `intoWriter559` variable:
	var intoWriter559 flate.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte602` to the receiver `intoWriter559`
	// (`intoWriter559` is now tainted).
	intoWriter559.Write(fromByte602)

	// Sink the tainted `intoWriter559`:
	sink(intoWriter559)
}

func TaintStepTest_CompressFlateResetterReset_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader361` into `intoResetter223`.

	// Assume that `sourceCQL` has the underlying type of `fromReader361`:
	fromReader361 := sourceCQL.(io.Reader)

	// Declare `intoResetter223` variable:
	var intoResetter223 flate.Resetter

	// Call the method that transfers the taint
	// from the parameter `fromReader361` to the receiver `intoResetter223`
	// (`intoResetter223` is now tainted).
	intoResetter223.Reset(fromReader361, nil)

	// Sink the tainted `intoResetter223`:
	sink(intoResetter223)
}

func RunAllTaints_CompressFlate() {
	{
		source := newSource()
		TaintStepTest_CompressFlateNewReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressFlateNewReaderDict_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressFlateNewWriter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressFlateNewWriterDict_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressFlateWriterReset_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressFlateWriterWrite_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressFlateResetterReset_B0I0O0(source)
	}
}
