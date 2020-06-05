package main

import (
	"io"
	"mime/quotedprintable"
)

func TaintStepTest_MimeQuotedprintableNewReader(sourceCQL interface{}) {
	// The flow is from `r` into `into366`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into366`
	// (`into366` is now tainted).
	into366 := quotedprintable.NewReader(r)

	// Sink the tainted `into366`:
	sink(into366)
}

func TaintStepTest_MimeQuotedprintableNewWriter(sourceCQL interface{}) {
	// The flow is from `from912` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from912`:
	from912 := sourceCQL.(*quotedprintable.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := quotedprintable.NewWriter(w)

	// Extra step (`from912` taints `intermediateCQL`, which taints `w`:
	link(from912, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_MimeQuotedprintableReaderRead(sourceCQL interface{}) {
	// The flow is from `from834` into `p`.

	// Assume that `sourceCQL` has the underlying type of `from834`:
	from834 := sourceCQL.(quotedprintable.Reader)

	// Declare `p` variable:
	var p []byte

	// Call the method that transfers the taint
	// from the receiver `from834` to the argument `p`
	// (`p` is now tainted).
	from834.Read(p)

	// Sink the tainted `p`:
	sink(p)
}

func TaintStepTest_MimeQuotedprintableWriterWrite(sourceCQL interface{}) {
	// The flow is from `p` into `into494`.

	// Assume that `sourceCQL` has the underlying type of `p`:
	p := sourceCQL.([]byte)

	// Declare `into494` variable:
	var into494 quotedprintable.Writer

	// Call the method that transfers the taint
	// from the parameter `p` to the receiver `into494`
	// (`into494` is now tainted).
	into494.Write(p)

	// Sink the tainted `into494`:
	sink(into494)
}

func RunAllTaints_MimeQuotedprintable(v interface{}) {
	{
		source := newSource()
		TaintStepTest_MimeQuotedprintableNewReader(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeQuotedprintableNewWriter(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeQuotedprintableReaderRead(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeQuotedprintableWriterWrite(source)
	}
}
