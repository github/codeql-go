package main

import (
	"io"
	"mime/quotedprintable"
)

func TaintStepTest_MimeQuotedprintableNewReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader606` into `intoReader606`.

	// Assume that `sourceCQL` has the underlying type of `fromReader606`:
	fromReader606 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader606` to result `intoReader606`
	// (`intoReader606` is now tainted).
	intoReader606 := quotedprintable.NewReader(fromReader606)

	// Sink the tainted `intoReader606`:
	sink(intoReader606)
}

func TaintStepTest_MimeQuotedprintableNewWriter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter623` into `intoWriter201`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter623`:
	fromWriter623 := sourceCQL.(*quotedprintable.Writer)

	// Declare `intoWriter201` variable:
	var intoWriter201 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter201`:
	intermediateCQL := quotedprintable.NewWriter(intoWriter201)

	// Extra step (`fromWriter623` taints `intermediateCQL`, which taints `intoWriter201`:
	link(fromWriter623, intermediateCQL)

	// Sink the tainted `intoWriter201`:
	sink(intoWriter201)
}

func TaintStepTest_MimeQuotedprintableReaderRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader841` into `intoByte209`.

	// Assume that `sourceCQL` has the underlying type of `fromReader841`:
	fromReader841 := sourceCQL.(quotedprintable.Reader)

	// Declare `intoByte209` variable:
	var intoByte209 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader841` to the argument `intoByte209`
	// (`intoByte209` is now tainted).
	fromReader841.Read(intoByte209)

	// Sink the tainted `intoByte209`:
	sink(intoByte209)
}

func TaintStepTest_MimeQuotedprintableWriterWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte393` into `intoWriter820`.

	// Assume that `sourceCQL` has the underlying type of `fromByte393`:
	fromByte393 := sourceCQL.([]byte)

	// Declare `intoWriter820` variable:
	var intoWriter820 quotedprintable.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte393` to the receiver `intoWriter820`
	// (`intoWriter820` is now tainted).
	intoWriter820.Write(fromByte393)

	// Sink the tainted `intoWriter820`:
	sink(intoWriter820)
}

func RunAllTaints_MimeQuotedprintable() {
	{
		source := newSource()
		TaintStepTest_MimeQuotedprintableNewReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeQuotedprintableNewWriter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeQuotedprintableReaderRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeQuotedprintableWriterWrite_B0I0O0(source)
	}
}
