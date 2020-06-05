package main

import (
	"compress/gzip"
	"io"
)

func TaintStepTest_CompressGzipNewReader(sourceCQL interface{}) {
	// The flow is from `r` into `into423`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into423`
	// (`into423` is now tainted).
	into423, _ := gzip.NewReader(r)

	// Sink the tainted `into423`:
	sink(into423)
}

func TaintStepTest_CompressGzipNewWriter(sourceCQL interface{}) {
	// The flow is from `from706` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from706`:
	from706 := sourceCQL.(*gzip.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := gzip.NewWriter(w)

	// Extra step (`from706` taints `intermediateCQL`, which taints `w`:
	link(from706, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_CompressGzipNewWriterLevel(sourceCQL interface{}) {
	// The flow is from `from682` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from682`:
	from682 := sourceCQL.(*gzip.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL, _ := gzip.NewWriterLevel(w, 0)

	// Extra step (`from682` taints `intermediateCQL`, which taints `w`:
	link(from682, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_CompressGzipReaderRead(sourceCQL interface{}) {
	// The flow is from `from661` into `p`.

	// Assume that `sourceCQL` has the underlying type of `from661`:
	from661 := sourceCQL.(gzip.Reader)

	// Declare `p` variable:
	var p []byte

	// Call the method that transfers the taint
	// from the receiver `from661` to the argument `p`
	// (`p` is now tainted).
	from661.Read(p)

	// Sink the tainted `p`:
	sink(p)
}

func TaintStepTest_CompressGzipReaderReset(sourceCQL interface{}) {
	// The flow is from `r` into `into719`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Declare `into719` variable:
	var into719 gzip.Reader

	// Call the method that transfers the taint
	// from the parameter `r` to the receiver `into719`
	// (`into719` is now tainted).
	into719.Reset(r)

	// Sink the tainted `into719`:
	sink(into719)
}

func TaintStepTest_CompressGzipWriterReset(sourceCQL interface{}) {
	// The flow is from `into457` into `w`.

	// Assume that `sourceCQL` has the underlying type of `into457`:
	into457 := sourceCQL.(gzip.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the method that transfers the taint
	// from the receiver `into457` to the argument `w`
	// (`w` is now tainted).
	into457.Reset(w)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_CompressGzipWriterWrite(sourceCQL interface{}) {
	// The flow is from `p` into `into857`.

	// Assume that `sourceCQL` has the underlying type of `p`:
	p := sourceCQL.([]byte)

	// Declare `into857` variable:
	var into857 gzip.Writer

	// Call the method that transfers the taint
	// from the parameter `p` to the receiver `into857`
	// (`into857` is now tainted).
	into857.Write(p)

	// Sink the tainted `into857`:
	sink(into857)
}

func RunAllTaints_CompressGzip(v interface{}) {
	{
		source := newSource()
		TaintStepTest_CompressGzipNewReader(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressGzipNewWriter(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressGzipNewWriterLevel(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressGzipReaderRead(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressGzipReaderReset(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressGzipWriterReset(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressGzipWriterWrite(source)
	}
}
