// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
)

func TaintStepTest_NetHttpHttputilDumpRequest_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRequest717` into `intoByte539`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest717`:
	fromRequest717 := sourceCQL.(*http.Request)

	// Call the function that transfers the taint
	// from the parameter `fromRequest717` to result `intoByte539`
	// (`intoByte539` is now tainted).
	intoByte539, _ := httputil.DumpRequest(fromRequest717, false)

	// Return the tainted `intoByte539`:
	return intoByte539
}

func TaintStepTest_NetHttpHttputilDumpRequestOut_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRequest284` into `intoByte956`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest284`:
	fromRequest284 := sourceCQL.(*http.Request)

	// Call the function that transfers the taint
	// from the parameter `fromRequest284` to result `intoByte956`
	// (`intoByte956` is now tainted).
	intoByte956, _ := httputil.DumpRequestOut(fromRequest284, false)

	// Return the tainted `intoByte956`:
	return intoByte956
}

func TaintStepTest_NetHttpHttputilDumpResponse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromResponse882` into `intoByte299`.

	// Assume that `sourceCQL` has the underlying type of `fromResponse882`:
	fromResponse882 := sourceCQL.(*http.Response)

	// Call the function that transfers the taint
	// from the parameter `fromResponse882` to result `intoByte299`
	// (`intoByte299` is now tainted).
	intoByte299, _ := httputil.DumpResponse(fromResponse882, false)

	// Return the tainted `intoByte299`:
	return intoByte299
}

func TaintStepTest_NetHttpHttputilNewChunkedReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader531` into `intoReader675`.

	// Assume that `sourceCQL` has the underlying type of `fromReader531`:
	fromReader531 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader531` to result `intoReader675`
	// (`intoReader675` is now tainted).
	intoReader675 := httputil.NewChunkedReader(fromReader531)

	// Return the tainted `intoReader675`:
	return intoReader675
}

func TaintStepTest_NetHttpHttputilNewChunkedWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser761` into `intoWriter953`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser761`:
	fromWriteCloser761 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter953` variable:
	var intoWriter953 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter953`:
	intermediateCQL := httputil.NewChunkedWriter(intoWriter953)

	// Extra step (`fromWriteCloser761` taints `intermediateCQL`, which taints `intoWriter953`:
	link(fromWriteCloser761, intermediateCQL)

	// Return the tainted `intoWriter953`:
	return intoWriter953
}

func TaintStepTest_NetHttpHttputilNewClientConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn970` into `intoClientConn537`.

	// Assume that `sourceCQL` has the underlying type of `fromConn970`:
	fromConn970 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `fromConn970` to result `intoClientConn537`
	// (`intoClientConn537` is now tainted).
	intoClientConn537 := httputil.NewClientConn(fromConn970, nil)

	// Return the tainted `intoClientConn537`:
	return intoClientConn537
}

func TaintStepTest_NetHttpHttputilNewClientConn_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader671` into `intoClientConn346`.

	// Assume that `sourceCQL` has the underlying type of `fromReader671`:
	fromReader671 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader671` to result `intoClientConn346`
	// (`intoClientConn346` is now tainted).
	intoClientConn346 := httputil.NewClientConn(nil, fromReader671)

	// Return the tainted `intoClientConn346`:
	return intoClientConn346
}

func TaintStepTest_NetHttpHttputilNewClientConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromClientConn233` into `intoConn638`.

	// Assume that `sourceCQL` has the underlying type of `fromClientConn233`:
	fromClientConn233 := sourceCQL.(*httputil.ClientConn)

	// Declare `intoConn638` variable:
	var intoConn638 net.Conn

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoConn638`:
	intermediateCQL := httputil.NewClientConn(intoConn638, nil)

	// Extra step (`fromClientConn233` taints `intermediateCQL`, which taints `intoConn638`:
	link(fromClientConn233, intermediateCQL)

	// Return the tainted `intoConn638`:
	return intoConn638
}

func TaintStepTest_NetHttpHttputilNewProxyClientConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn372` into `intoClientConn167`.

	// Assume that `sourceCQL` has the underlying type of `fromConn372`:
	fromConn372 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `fromConn372` to result `intoClientConn167`
	// (`intoClientConn167` is now tainted).
	intoClientConn167 := httputil.NewProxyClientConn(fromConn372, nil)

	// Return the tainted `intoClientConn167`:
	return intoClientConn167
}

func TaintStepTest_NetHttpHttputilNewProxyClientConn_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader218` into `intoClientConn813`.

	// Assume that `sourceCQL` has the underlying type of `fromReader218`:
	fromReader218 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader218` to result `intoClientConn813`
	// (`intoClientConn813` is now tainted).
	intoClientConn813 := httputil.NewProxyClientConn(nil, fromReader218)

	// Return the tainted `intoClientConn813`:
	return intoClientConn813
}

func TaintStepTest_NetHttpHttputilNewProxyClientConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromClientConn420` into `intoConn473`.

	// Assume that `sourceCQL` has the underlying type of `fromClientConn420`:
	fromClientConn420 := sourceCQL.(*httputil.ClientConn)

	// Declare `intoConn473` variable:
	var intoConn473 net.Conn

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoConn473`:
	intermediateCQL := httputil.NewProxyClientConn(intoConn473, nil)

	// Extra step (`fromClientConn420` taints `intermediateCQL`, which taints `intoConn473`:
	link(fromClientConn420, intermediateCQL)

	// Return the tainted `intoConn473`:
	return intoConn473
}

func TaintStepTest_NetHttpHttputilClientConnHijack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromClientConn458` into `intoConn189`.

	// Assume that `sourceCQL` has the underlying type of `fromClientConn458`:
	fromClientConn458 := sourceCQL.(httputil.ClientConn)

	// Call the method that transfers the taint
	// from the receiver `fromClientConn458` to the result `intoConn189`
	// (`intoConn189` is now tainted).
	intoConn189, _ := fromClientConn458.Hijack()

	// Return the tainted `intoConn189`:
	return intoConn189
}

func TaintStepTest_NetHttpHttputilClientConnHijack_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromClientConn684` into `intoReader486`.

	// Assume that `sourceCQL` has the underlying type of `fromClientConn684`:
	fromClientConn684 := sourceCQL.(httputil.ClientConn)

	// Call the method that transfers the taint
	// from the receiver `fromClientConn684` to the result `intoReader486`
	// (`intoReader486` is now tainted).
	_, intoReader486 := fromClientConn684.Hijack()

	// Return the tainted `intoReader486`:
	return intoReader486
}

func TaintStepTest_NetHttpHttputilServerConnHijack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromServerConn774` into `intoConn656`.

	// Assume that `sourceCQL` has the underlying type of `fromServerConn774`:
	fromServerConn774 := sourceCQL.(httputil.ServerConn)

	// Call the method that transfers the taint
	// from the receiver `fromServerConn774` to the result `intoConn656`
	// (`intoConn656` is now tainted).
	intoConn656, _ := fromServerConn774.Hijack()

	// Return the tainted `intoConn656`:
	return intoConn656
}

func TaintStepTest_NetHttpHttputilServerConnHijack_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromServerConn129` into `intoReader253`.

	// Assume that `sourceCQL` has the underlying type of `fromServerConn129`:
	fromServerConn129 := sourceCQL.(httputil.ServerConn)

	// Call the method that transfers the taint
	// from the receiver `fromServerConn129` to the result `intoReader253`
	// (`intoReader253` is now tainted).
	_, intoReader253 := fromServerConn129.Hijack()

	// Return the tainted `intoReader253`:
	return intoReader253
}

func TaintStepTest_NetHttpHttputilBufferPoolGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBufferPool831` into `intoByte411`.

	// Assume that `sourceCQL` has the underlying type of `fromBufferPool831`:
	fromBufferPool831 := sourceCQL.(httputil.BufferPool)

	// Call the method that transfers the taint
	// from the receiver `fromBufferPool831` to the result `intoByte411`
	// (`intoByte411` is now tainted).
	intoByte411 := fromBufferPool831.Get()

	// Return the tainted `intoByte411`:
	return intoByte411
}

func TaintStepTest_NetHttpHttputilBufferPoolPut_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte564` into `intoBufferPool754`.

	// Assume that `sourceCQL` has the underlying type of `fromByte564`:
	fromByte564 := sourceCQL.([]byte)

	// Declare `intoBufferPool754` variable:
	var intoBufferPool754 httputil.BufferPool

	// Call the method that transfers the taint
	// from the parameter `fromByte564` to the receiver `intoBufferPool754`
	// (`intoBufferPool754` is now tainted).
	intoBufferPool754.Put(fromByte564)

	// Return the tainted `intoBufferPool754`:
	return intoBufferPool754
}

func RunAllTaints_NetHttpHttputil() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHttputilDumpRequest_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHttputilDumpRequestOut_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHttputilDumpResponse_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHttputilNewChunkedReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHttputilNewChunkedWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHttputilNewClientConn_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHttputilNewClientConn_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHttputilNewClientConn_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHttputilNewProxyClientConn_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHttputilNewProxyClientConn_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHttputilNewProxyClientConn_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHttputilClientConnHijack_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHttputilClientConnHijack_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHttputilServerConnHijack_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHttputilServerConnHijack_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHttputilBufferPoolGet_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHttputilBufferPoolPut_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
