// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"mime/quotedprintable"
)

func TaintStepTest_MimeQuotedprintableNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader656` into `intoReader414`.

	// Assume that `sourceCQL` has the underlying type of `fromReader656`:
	fromReader656 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader656` to result `intoReader414`
	// (`intoReader414` is now tainted).
	intoReader414 := quotedprintable.NewReader(fromReader656)

	// Return the tainted `intoReader414`:
	return intoReader414
}

func TaintStepTest_MimeQuotedprintableNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter518` into `intoWriter650`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter518`:
	fromWriter518 := sourceCQL.(*quotedprintable.Writer)

	// Declare `intoWriter650` variable:
	var intoWriter650 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter650`:
	intermediateCQL := quotedprintable.NewWriter(intoWriter650)

	// Extra step (`fromWriter518` taints `intermediateCQL`, which taints `intoWriter650`:
	link(fromWriter518, intermediateCQL)

	// Return the tainted `intoWriter650`:
	return intoWriter650
}

func TaintStepTest_MimeQuotedprintableReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader784` into `intoByte957`.

	// Assume that `sourceCQL` has the underlying type of `fromReader784`:
	fromReader784 := sourceCQL.(quotedprintable.Reader)

	// Declare `intoByte957` variable:
	var intoByte957 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader784` to the argument `intoByte957`
	// (`intoByte957` is now tainted).
	fromReader784.Read(intoByte957)

	// Return the tainted `intoByte957`:
	return intoByte957
}

func TaintStepTest_MimeQuotedprintableWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte520` into `intoWriter443`.

	// Assume that `sourceCQL` has the underlying type of `fromByte520`:
	fromByte520 := sourceCQL.([]byte)

	// Declare `intoWriter443` variable:
	var intoWriter443 quotedprintable.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte520` to the receiver `intoWriter443`
	// (`intoWriter443` is now tainted).
	intoWriter443.Write(fromByte520)

	// Return the tainted `intoWriter443`:
	return intoWriter443
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
