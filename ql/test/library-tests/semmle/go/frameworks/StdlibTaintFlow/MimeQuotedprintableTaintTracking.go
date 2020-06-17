package main

import (
	"io"
	"mime/quotedprintable"
)

func TaintStepTest_MimeQuotedprintableNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader491` into `intoReader482`.

	// Assume that `sourceCQL` has the underlying type of `fromReader491`:
	fromReader491 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader491` to result `intoReader482`
	// (`intoReader482` is now tainted).
	intoReader482 := quotedprintable.NewReader(fromReader491)

	// Return the tainted `intoReader482`:
	return intoReader482
}

func TaintStepTest_MimeQuotedprintableNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter849` into `intoWriter986`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter849`:
	fromWriter849 := sourceCQL.(*quotedprintable.Writer)

	// Declare `intoWriter986` variable:
	var intoWriter986 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter986`:
	intermediateCQL := quotedprintable.NewWriter(intoWriter986)

	// Extra step (`fromWriter849` taints `intermediateCQL`, which taints `intoWriter986`:
	link(fromWriter849, intermediateCQL)

	// Return the tainted `intoWriter986`:
	return intoWriter986
}

func TaintStepTest_MimeQuotedprintableReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader643` into `intoByte771`.

	// Assume that `sourceCQL` has the underlying type of `fromReader643`:
	fromReader643 := sourceCQL.(quotedprintable.Reader)

	// Declare `intoByte771` variable:
	var intoByte771 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader643` to the argument `intoByte771`
	// (`intoByte771` is now tainted).
	fromReader643.Read(intoByte771)

	// Return the tainted `intoByte771`:
	return intoByte771
}

func TaintStepTest_MimeQuotedprintableWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte602` into `intoWriter827`.

	// Assume that `sourceCQL` has the underlying type of `fromByte602`:
	fromByte602 := sourceCQL.([]byte)

	// Declare `intoWriter827` variable:
	var intoWriter827 quotedprintable.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte602` to the receiver `intoWriter827`
	// (`intoWriter827` is now tainted).
	intoWriter827.Write(fromByte602)

	// Return the tainted `intoWriter827`:
	return intoWriter827
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
