// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"mime/quotedprintable"
)

func TaintStepTest_MimeQuotedprintableNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader200` into `intoReader525`.

	// Assume that `sourceCQL` has the underlying type of `fromReader200`:
	fromReader200 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader200` to result `intoReader525`
	// (`intoReader525` is now tainted).
	intoReader525 := quotedprintable.NewReader(fromReader200)

	// Return the tainted `intoReader525`:
	return intoReader525
}

func TaintStepTest_MimeQuotedprintableNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter935` into `intoWriter337`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter935`:
	fromWriter935 := sourceCQL.(*quotedprintable.Writer)

	// Declare `intoWriter337` variable:
	var intoWriter337 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter337`:
	intermediateCQL := quotedprintable.NewWriter(intoWriter337)

	// Extra step (`fromWriter935` taints `intermediateCQL`, which taints `intoWriter337`:
	link(fromWriter935, intermediateCQL)

	// Return the tainted `intoWriter337`:
	return intoWriter337
}

func TaintStepTest_MimeQuotedprintableReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader882` into `intoByte328`.

	// Assume that `sourceCQL` has the underlying type of `fromReader882`:
	fromReader882 := sourceCQL.(quotedprintable.Reader)

	// Declare `intoByte328` variable:
	var intoByte328 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader882` to the argument `intoByte328`
	// (`intoByte328` is now tainted).
	fromReader882.Read(intoByte328)

	// Return the tainted `intoByte328`:
	return intoByte328
}

func TaintStepTest_MimeQuotedprintableWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte757` into `intoWriter750`.

	// Assume that `sourceCQL` has the underlying type of `fromByte757`:
	fromByte757 := sourceCQL.([]byte)

	// Declare `intoWriter750` variable:
	var intoWriter750 quotedprintable.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte757` to the receiver `intoWriter750`
	// (`intoWriter750` is now tainted).
	intoWriter750.Write(fromByte757)

	// Return the tainted `intoWriter750`:
	return intoWriter750
}

func RunAllTaints_MimeQuotedprintable() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeQuotedprintableNewReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeQuotedprintableNewWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeQuotedprintableReaderRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeQuotedprintableWriterWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
