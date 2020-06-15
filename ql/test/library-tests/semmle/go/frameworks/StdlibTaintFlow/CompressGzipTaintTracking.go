package main

import (
	"compress/gzip"
	"io"
)

func TaintStepTest_CompressGzipNewReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader273` into `intoReader351`.

	// Assume that `sourceCQL` has the underlying type of `fromReader273`:
	fromReader273 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader273` to result `intoReader351`
	// (`intoReader351` is now tainted).
	intoReader351, _ := gzip.NewReader(fromReader273)

	// Sink the tainted `intoReader351`:
	sink(intoReader351)
}

func TaintStepTest_CompressGzipNewWriter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter837` into `intoWriter692`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter837`:
	fromWriter837 := sourceCQL.(*gzip.Writer)

	// Declare `intoWriter692` variable:
	var intoWriter692 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter692`:
	intermediateCQL := gzip.NewWriter(intoWriter692)

	// Extra step (`fromWriter837` taints `intermediateCQL`, which taints `intoWriter692`:
	link(fromWriter837, intermediateCQL)

	// Sink the tainted `intoWriter692`:
	sink(intoWriter692)
}

func TaintStepTest_CompressGzipNewWriterLevel_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter144` into `intoWriter193`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter144`:
	fromWriter144 := sourceCQL.(*gzip.Writer)

	// Declare `intoWriter193` variable:
	var intoWriter193 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter193`:
	intermediateCQL, _ := gzip.NewWriterLevel(intoWriter193, 0)

	// Extra step (`fromWriter144` taints `intermediateCQL`, which taints `intoWriter193`:
	link(fromWriter144, intermediateCQL)

	// Sink the tainted `intoWriter193`:
	sink(intoWriter193)
}

func TaintStepTest_CompressGzipReaderRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader238` into `intoByte275`.

	// Assume that `sourceCQL` has the underlying type of `fromReader238`:
	fromReader238 := sourceCQL.(gzip.Reader)

	// Declare `intoByte275` variable:
	var intoByte275 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader238` to the argument `intoByte275`
	// (`intoByte275` is now tainted).
	fromReader238.Read(intoByte275)

	// Sink the tainted `intoByte275`:
	sink(intoByte275)
}

func TaintStepTest_CompressGzipReaderReset_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader297` into `intoReader146`.

	// Assume that `sourceCQL` has the underlying type of `fromReader297`:
	fromReader297 := sourceCQL.(io.Reader)

	// Declare `intoReader146` variable:
	var intoReader146 gzip.Reader

	// Call the method that transfers the taint
	// from the parameter `fromReader297` to the receiver `intoReader146`
	// (`intoReader146` is now tainted).
	intoReader146.Reset(fromReader297)

	// Sink the tainted `intoReader146`:
	sink(intoReader146)
}

func TaintStepTest_CompressGzipWriterReset_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter112` into `intoWriter829`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter112`:
	fromWriter112 := sourceCQL.(gzip.Writer)

	// Declare `intoWriter829` variable:
	var intoWriter829 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter112` to the argument `intoWriter829`
	// (`intoWriter829` is now tainted).
	fromWriter112.Reset(intoWriter829)

	// Sink the tainted `intoWriter829`:
	sink(intoWriter829)
}

func TaintStepTest_CompressGzipWriterWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte294` into `intoWriter914`.

	// Assume that `sourceCQL` has the underlying type of `fromByte294`:
	fromByte294 := sourceCQL.([]byte)

	// Declare `intoWriter914` variable:
	var intoWriter914 gzip.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte294` to the receiver `intoWriter914`
	// (`intoWriter914` is now tainted).
	intoWriter914.Write(fromByte294)

	// Sink the tainted `intoWriter914`:
	sink(intoWriter914)
}

func RunAllTaints_CompressGzip() {
	{
		source := newSource()
		TaintStepTest_CompressGzipNewReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressGzipNewWriter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressGzipNewWriterLevel_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressGzipReaderRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressGzipReaderReset_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressGzipWriterReset_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressGzipWriterWrite_B0I0O0(source)
	}
}
