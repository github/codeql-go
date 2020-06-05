package main

import (
	"compress/flate"
	"io"
)

func TaintStepTest_CompressFlateNewReader(sourceCQL interface{}) {
	// The flow is from `r` into `into798`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into798`
	// (`into798` is now tainted).
	into798 := flate.NewReader(r)

	// Sink the tainted `into798`:
	sink(into798)
}

func TaintStepTest_CompressFlateNewReaderDict(sourceCQL interface{}) {
	// The flow is from `r` into `into643`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into643`
	// (`into643` is now tainted).
	into643 := flate.NewReaderDict(r, nil)

	// Sink the tainted `into643`:
	sink(into643)
}

func TaintStepTest_CompressFlateNewWriter(sourceCQL interface{}) {
	// The flow is from `from413` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from413`:
	from413 := sourceCQL.(*flate.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL, _ := flate.NewWriter(w, 0)

	// Extra step (`from413` taints `intermediateCQL`, which taints `w`:
	link(from413, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_CompressFlateNewWriterDict(sourceCQL interface{}) {
	// The flow is from `from529` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from529`:
	from529 := sourceCQL.(*flate.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL, _ := flate.NewWriterDict(w, 0, nil)

	// Extra step (`from529` taints `intermediateCQL`, which taints `w`:
	link(from529, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_CompressFlateWriterReset(sourceCQL interface{}) {
	// The flow is from `from955` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `from955`:
	from955 := sourceCQL.(flate.Writer)

	// Declare `dst` variable:
	var dst io.Writer

	// Call the method that transfers the taint
	// from the receiver `from955` to the argument `dst`
	// (`dst` is now tainted).
	from955.Reset(dst)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_CompressFlateWriterWrite(sourceCQL interface{}) {
	// The flow is from `data` into `into410`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.([]byte)

	// Declare `into410` variable:
	var into410 flate.Writer

	// Call the method that transfers the taint
	// from the parameter `data` to the receiver `into410`
	// (`into410` is now tainted).
	into410.Write(data)

	// Sink the tainted `into410`:
	sink(into410)
}

func TaintStepTest_CompressFlateResetterReset(sourceCQL interface{}) {
	// The flow is from `r` into `into309`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Declare `into309` variable:
	var into309 flate.Resetter

	// Call the method that transfers the taint
	// from the parameter `r` to the receiver `into309`
	// (`into309` is now tainted).
	into309.Reset(r, nil)

	// Sink the tainted `into309`:
	sink(into309)
}

func RunAllTaints_CompressFlate(v interface{}) {
	{
		source := newSource()
		TaintStepTest_CompressFlateNewReader(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressFlateNewReaderDict(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressFlateNewWriter(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressFlateNewWriterDict(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressFlateWriterReset(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressFlateWriterWrite(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressFlateResetterReset(source)
	}
}
