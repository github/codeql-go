package main

import (
	"compress/zlib"
	"io"
)

func TaintStepTest_CompressZlibNewReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader174` into `intoReadCloser122`.

	// Assume that `sourceCQL` has the underlying type of `fromReader174`:
	fromReader174 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader174` to result `intoReadCloser122`
	// (`intoReadCloser122` is now tainted).
	intoReadCloser122, _ := zlib.NewReader(fromReader174)

	// Sink the tainted `intoReadCloser122`:
	sink(intoReadCloser122)
}

func TaintStepTest_CompressZlibNewReaderDict_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader876` into `intoReadCloser290`.

	// Assume that `sourceCQL` has the underlying type of `fromReader876`:
	fromReader876 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader876` to result `intoReadCloser290`
	// (`intoReadCloser290` is now tainted).
	intoReadCloser290, _ := zlib.NewReaderDict(fromReader876, nil)

	// Sink the tainted `intoReadCloser290`:
	sink(intoReadCloser290)
}

func TaintStepTest_CompressZlibNewWriter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter358` into `intoWriter362`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter358`:
	fromWriter358 := sourceCQL.(*zlib.Writer)

	// Declare `intoWriter362` variable:
	var intoWriter362 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter362`:
	intermediateCQL := zlib.NewWriter(intoWriter362)

	// Extra step (`fromWriter358` taints `intermediateCQL`, which taints `intoWriter362`:
	link(fromWriter358, intermediateCQL)

	// Sink the tainted `intoWriter362`:
	sink(intoWriter362)
}

func TaintStepTest_CompressZlibNewWriterLevel_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter294` into `intoWriter384`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter294`:
	fromWriter294 := sourceCQL.(*zlib.Writer)

	// Declare `intoWriter384` variable:
	var intoWriter384 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter384`:
	intermediateCQL, _ := zlib.NewWriterLevel(intoWriter384, 0)

	// Extra step (`fromWriter294` taints `intermediateCQL`, which taints `intoWriter384`:
	link(fromWriter294, intermediateCQL)

	// Sink the tainted `intoWriter384`:
	sink(intoWriter384)
}

func TaintStepTest_CompressZlibNewWriterLevelDict_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter432` into `intoWriter204`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter432`:
	fromWriter432 := sourceCQL.(*zlib.Writer)

	// Declare `intoWriter204` variable:
	var intoWriter204 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter204`:
	intermediateCQL, _ := zlib.NewWriterLevelDict(intoWriter204, 0, nil)

	// Extra step (`fromWriter432` taints `intermediateCQL`, which taints `intoWriter204`:
	link(fromWriter432, intermediateCQL)

	// Sink the tainted `intoWriter204`:
	sink(intoWriter204)
}

func TaintStepTest_CompressZlibWriterReset_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter156` into `intoWriter563`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter156`:
	fromWriter156 := sourceCQL.(zlib.Writer)

	// Declare `intoWriter563` variable:
	var intoWriter563 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter156` to the argument `intoWriter563`
	// (`intoWriter563` is now tainted).
	fromWriter156.Reset(intoWriter563)

	// Sink the tainted `intoWriter563`:
	sink(intoWriter563)
}

func TaintStepTest_CompressZlibWriterWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte794` into `intoWriter584`.

	// Assume that `sourceCQL` has the underlying type of `fromByte794`:
	fromByte794 := sourceCQL.([]byte)

	// Declare `intoWriter584` variable:
	var intoWriter584 zlib.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte794` to the receiver `intoWriter584`
	// (`intoWriter584` is now tainted).
	intoWriter584.Write(fromByte794)

	// Sink the tainted `intoWriter584`:
	sink(intoWriter584)
}

func TaintStepTest_CompressZlibResetterReset_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader177` into `intoResetter511`.

	// Assume that `sourceCQL` has the underlying type of `fromReader177`:
	fromReader177 := sourceCQL.(io.Reader)

	// Declare `intoResetter511` variable:
	var intoResetter511 zlib.Resetter

	// Call the method that transfers the taint
	// from the parameter `fromReader177` to the receiver `intoResetter511`
	// (`intoResetter511` is now tainted).
	intoResetter511.Reset(fromReader177, nil)

	// Sink the tainted `intoResetter511`:
	sink(intoResetter511)
}

func RunAllTaints_CompressZlib() {
	{
		source := newSource()
		TaintStepTest_CompressZlibNewReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressZlibNewReaderDict_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressZlibNewWriter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressZlibNewWriterLevel_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressZlibNewWriterLevelDict_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressZlibWriterReset_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressZlibWriterWrite_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressZlibResetterReset_B0I0O0(source)
	}
}
