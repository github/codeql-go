// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"compress/gzip"
	"io"
)

func TaintStepTest_CompressGzipNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader210` into `intoReader552`.

	// Assume that `sourceCQL` has the underlying type of `fromReader210`:
	fromReader210 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader210` to result `intoReader552`
	// (`intoReader552` is now tainted).
	intoReader552, _ := gzip.NewReader(fromReader210)

	// Return the tainted `intoReader552`:
	return intoReader552
}

func TaintStepTest_CompressGzipNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter379` into `intoWriter289`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter379`:
	fromWriter379 := sourceCQL.(*gzip.Writer)

	// Declare `intoWriter289` variable:
	var intoWriter289 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter289`:
	intermediateCQL := gzip.NewWriter(intoWriter289)

	// Extra step (`fromWriter379` taints `intermediateCQL`, which taints `intoWriter289`:
	link(fromWriter379, intermediateCQL)

	// Return the tainted `intoWriter289`:
	return intoWriter289
}

func TaintStepTest_CompressGzipNewWriterLevel_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter124` into `intoWriter535`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter124`:
	fromWriter124 := sourceCQL.(*gzip.Writer)

	// Declare `intoWriter535` variable:
	var intoWriter535 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter535`:
	intermediateCQL, _ := gzip.NewWriterLevel(intoWriter535, 0)

	// Extra step (`fromWriter124` taints `intermediateCQL`, which taints `intoWriter535`:
	link(fromWriter124, intermediateCQL)

	// Return the tainted `intoWriter535`:
	return intoWriter535
}

func TaintStepTest_CompressGzipReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader289` into `intoByte985`.

	// Assume that `sourceCQL` has the underlying type of `fromReader289`:
	fromReader289 := sourceCQL.(gzip.Reader)

	// Declare `intoByte985` variable:
	var intoByte985 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader289` to the argument `intoByte985`
	// (`intoByte985` is now tainted).
	fromReader289.Read(intoByte985)

	// Return the tainted `intoByte985`:
	return intoByte985
}

func TaintStepTest_CompressGzipReaderReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader832` into `intoReader417`.

	// Assume that `sourceCQL` has the underlying type of `fromReader832`:
	fromReader832 := sourceCQL.(io.Reader)

	// Declare `intoReader417` variable:
	var intoReader417 gzip.Reader

	// Call the method that transfers the taint
	// from the parameter `fromReader832` to the receiver `intoReader417`
	// (`intoReader417` is now tainted).
	intoReader417.Reset(fromReader832)

	// Return the tainted `intoReader417`:
	return intoReader417
}

func TaintStepTest_CompressGzipWriterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter585` into `intoWriter238`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter585`:
	fromWriter585 := sourceCQL.(gzip.Writer)

	// Declare `intoWriter238` variable:
	var intoWriter238 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter585` to the argument `intoWriter238`
	// (`intoWriter238` is now tainted).
	fromWriter585.Reset(intoWriter238)

	// Return the tainted `intoWriter238`:
	return intoWriter238
}

func TaintStepTest_CompressGzipWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte350` into `intoWriter142`.

	// Assume that `sourceCQL` has the underlying type of `fromByte350`:
	fromByte350 := sourceCQL.([]byte)

	// Declare `intoWriter142` variable:
	var intoWriter142 gzip.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte350` to the receiver `intoWriter142`
	// (`intoWriter142` is now tainted).
	intoWriter142.Write(fromByte350)

	// Return the tainted `intoWriter142`:
	return intoWriter142
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
