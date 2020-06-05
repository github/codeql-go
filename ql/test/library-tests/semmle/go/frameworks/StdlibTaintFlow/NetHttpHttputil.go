package main

import (
	"io"
	"net/http"
	"net/http/httputil"
)

func TaintStepTest_NetHttpHttputilDumpRequest(sourceCQL interface{}) {
	// The flow is from `req` into `into225`.

	// Assume that `sourceCQL` has the underlying type of `req`:
	req := sourceCQL.(*http.Request)

	// Call the function that transfers the taint
	// from the parameter `req` to result `into225`
	// (`into225` is now tainted).
	into225, _ := httputil.DumpRequest(req, false)

	// Sink the tainted `into225`:
	sink(into225)
}

func TaintStepTest_NetHttpHttputilDumpRequestOut(sourceCQL interface{}) {
	// The flow is from `req` into `into652`.

	// Assume that `sourceCQL` has the underlying type of `req`:
	req := sourceCQL.(*http.Request)

	// Call the function that transfers the taint
	// from the parameter `req` to result `into652`
	// (`into652` is now tainted).
	into652, _ := httputil.DumpRequestOut(req, false)

	// Sink the tainted `into652`:
	sink(into652)
}

func TaintStepTest_NetHttpHttputilDumpResponse(sourceCQL interface{}) {
	// The flow is from `resp` into `into803`.

	// Assume that `sourceCQL` has the underlying type of `resp`:
	resp := sourceCQL.(*http.Response)

	// Call the function that transfers the taint
	// from the parameter `resp` to result `into803`
	// (`into803` is now tainted).
	into803, _ := httputil.DumpResponse(resp, false)

	// Sink the tainted `into803`:
	sink(into803)
}

func TaintStepTest_NetHttpHttputilNewChunkedReader(sourceCQL interface{}) {
	// The flow is from `r` into `into785`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into785`
	// (`into785` is now tainted).
	into785 := httputil.NewChunkedReader(r)

	// Sink the tainted `into785`:
	sink(into785)
}

func TaintStepTest_NetHttpHttputilNewChunkedWriter(sourceCQL interface{}) {
	// The flow is from `from703` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from703`:
	from703 := sourceCQL.(io.WriteCloser)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := httputil.NewChunkedWriter(w)

	// Extra step (`from703` taints `intermediateCQL`, which taints `w`:
	link(from703, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_NetHttpHttputilBufferPoolGet(sourceCQL interface{}) {
	// The flow is from `from317` into `into786`.

	// Assume that `sourceCQL` has the underlying type of `from317`:
	from317 := sourceCQL.(httputil.BufferPool)

	// Call the method that transfers the taint
	// from the receiver `from317` to the result `into786`
	// (`into786` is now tainted).
	into786 := from317.Get()

	// Sink the tainted `into786`:
	sink(into786)
}

func TaintStepTest_NetHttpHttputilBufferPoolPut(sourceCQL interface{}) {
	// The flow is from `from399` into `into320`.

	// Assume that `sourceCQL` has the underlying type of `from399`:
	from399 := sourceCQL.([]byte)

	// Declare `into320` variable:
	var into320 httputil.BufferPool

	// Call the method that transfers the taint
	// from the parameter `from399` to the receiver `into320`
	// (`into320` is now tainted).
	into320.Put(from399)

	// Sink the tainted `into320`:
	sink(into320)
}

func RunAllTaints_NetHttpHttputil(v interface{}) {
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilDumpRequest(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilDumpRequestOut(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilDumpResponse(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilNewChunkedReader(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilNewChunkedWriter(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilBufferPoolGet(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilBufferPoolPut(source)
	}
}
