package main

import (
	"compress/zlib"
	"io"
)

func TaintStepTest_CompressZlibNewReader(sourceCQL interface{}) {
	// The flow is from `r` into `into207`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into207`
	// (`into207` is now tainted).
	into207, _ := zlib.NewReader(r)

	// Sink the tainted `into207`:
	sink(into207)
}

func TaintStepTest_CompressZlibNewReaderDict(sourceCQL interface{}) {
	// The flow is from `r` into `into940`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into940`
	// (`into940` is now tainted).
	into940, _ := zlib.NewReaderDict(r, nil)

	// Sink the tainted `into940`:
	sink(into940)
}

func TaintStepTest_CompressZlibNewWriter(sourceCQL interface{}) {
	// The flow is from `from201` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from201`:
	from201 := sourceCQL.(*zlib.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := zlib.NewWriter(w)

	// Extra step (`from201` taints `intermediateCQL`, which taints `w`:
	link(from201, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_CompressZlibNewWriterLevel(sourceCQL interface{}) {
	// The flow is from `from607` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from607`:
	from607 := sourceCQL.(*zlib.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL, _ := zlib.NewWriterLevel(w, 0)

	// Extra step (`from607` taints `intermediateCQL`, which taints `w`:
	link(from607, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_CompressZlibNewWriterLevelDict(sourceCQL interface{}) {
	// The flow is from `from551` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from551`:
	from551 := sourceCQL.(*zlib.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL, _ := zlib.NewWriterLevelDict(w, 0, nil)

	// Extra step (`from551` taints `intermediateCQL`, which taints `w`:
	link(from551, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_CompressZlibWriterReset(sourceCQL interface{}) {
	// The flow is from `from549` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from549`:
	from549 := sourceCQL.(zlib.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the method that transfers the taint
	// from the receiver `from549` to the argument `w`
	// (`w` is now tainted).
	from549.Reset(w)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_CompressZlibWriterWrite(sourceCQL interface{}) {
	// The flow is from `p` into `into535`.

	// Assume that `sourceCQL` has the underlying type of `p`:
	p := sourceCQL.([]byte)

	// Declare `into535` variable:
	var into535 zlib.Writer

	// Call the method that transfers the taint
	// from the parameter `p` to the receiver `into535`
	// (`into535` is now tainted).
	into535.Write(p)

	// Sink the tainted `into535`:
	sink(into535)
}

func TaintStepTest_CompressZlibResetterReset(sourceCQL interface{}) {
	// The flow is from `r` into `into716`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Declare `into716` variable:
	var into716 zlib.Resetter

	// Call the method that transfers the taint
	// from the parameter `r` to the receiver `into716`
	// (`into716` is now tainted).
	into716.Reset(r, nil)

	// Sink the tainted `into716`:
	sink(into716)
}

func RunAllTaints_CompressZlib(v interface{}) {
	{
		source := newSource()
		TaintStepTest_CompressZlibNewReader(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressZlibNewReaderDict(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressZlibNewWriter(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressZlibNewWriterLevel(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressZlibNewWriterLevelDict(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressZlibWriterReset(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressZlibWriterWrite(source)
	}
	{
		source := newSource()
		TaintStepTest_CompressZlibResetterReset(source)
	}
}
