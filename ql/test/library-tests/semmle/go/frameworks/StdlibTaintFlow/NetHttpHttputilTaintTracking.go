package main

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
)

func TaintStepTest_NetHttpHttputilDumpRequest_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRequest723` into `intoByte881`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest723`:
	fromRequest723 := sourceCQL.(*http.Request)

	// Call the function that transfers the taint
	// from the parameter `fromRequest723` to result `intoByte881`
	// (`intoByte881` is now tainted).
	intoByte881, _ := httputil.DumpRequest(fromRequest723, false)

	// Sink the tainted `intoByte881`:
	sink(intoByte881)
}

func TaintStepTest_NetHttpHttputilDumpRequestOut_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRequest722` into `intoByte916`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest722`:
	fromRequest722 := sourceCQL.(*http.Request)

	// Call the function that transfers the taint
	// from the parameter `fromRequest722` to result `intoByte916`
	// (`intoByte916` is now tainted).
	intoByte916, _ := httputil.DumpRequestOut(fromRequest722, false)

	// Sink the tainted `intoByte916`:
	sink(intoByte916)
}

func TaintStepTest_NetHttpHttputilDumpResponse_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromResponse934` into `intoByte302`.

	// Assume that `sourceCQL` has the underlying type of `fromResponse934`:
	fromResponse934 := sourceCQL.(*http.Response)

	// Call the function that transfers the taint
	// from the parameter `fromResponse934` to result `intoByte302`
	// (`intoByte302` is now tainted).
	intoByte302, _ := httputil.DumpResponse(fromResponse934, false)

	// Sink the tainted `intoByte302`:
	sink(intoByte302)
}

func TaintStepTest_NetHttpHttputilNewChunkedReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader469` into `intoReader430`.

	// Assume that `sourceCQL` has the underlying type of `fromReader469`:
	fromReader469 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader469` to result `intoReader430`
	// (`intoReader430` is now tainted).
	intoReader430 := httputil.NewChunkedReader(fromReader469)

	// Sink the tainted `intoReader430`:
	sink(intoReader430)
}

func TaintStepTest_NetHttpHttputilNewChunkedWriter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriteCloser901` into `intoWriter540`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser901`:
	fromWriteCloser901 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter540` variable:
	var intoWriter540 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter540`:
	intermediateCQL := httputil.NewChunkedWriter(intoWriter540)

	// Extra step (`fromWriteCloser901` taints `intermediateCQL`, which taints `intoWriter540`:
	link(fromWriteCloser901, intermediateCQL)

	// Sink the tainted `intoWriter540`:
	sink(intoWriter540)
}

func TaintStepTest_NetHttpHttputilNewClientConn_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromConn464` into `intoClientConn956`.

	// Assume that `sourceCQL` has the underlying type of `fromConn464`:
	fromConn464 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `fromConn464` to result `intoClientConn956`
	// (`intoClientConn956` is now tainted).
	intoClientConn956 := httputil.NewClientConn(fromConn464, nil)

	// Sink the tainted `intoClientConn956`:
	sink(intoClientConn956)
}

func TaintStepTest_NetHttpHttputilNewClientConn_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromReader345` into `intoClientConn697`.

	// Assume that `sourceCQL` has the underlying type of `fromReader345`:
	fromReader345 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader345` to result `intoClientConn697`
	// (`intoClientConn697` is now tainted).
	intoClientConn697 := httputil.NewClientConn(nil, fromReader345)

	// Sink the tainted `intoClientConn697`:
	sink(intoClientConn697)
}

func TaintStepTest_NetHttpHttputilNewClientConn_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromClientConn114` into `intoConn278`.

	// Assume that `sourceCQL` has the underlying type of `fromClientConn114`:
	fromClientConn114 := sourceCQL.(*httputil.ClientConn)

	// Declare `intoConn278` variable:
	var intoConn278 net.Conn

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoConn278`:
	intermediateCQL := httputil.NewClientConn(intoConn278, nil)

	// Extra step (`fromClientConn114` taints `intermediateCQL`, which taints `intoConn278`:
	link(fromClientConn114, intermediateCQL)

	// Sink the tainted `intoConn278`:
	sink(intoConn278)
}

func TaintStepTest_NetHttpHttputilNewProxyClientConn_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromConn403` into `intoClientConn550`.

	// Assume that `sourceCQL` has the underlying type of `fromConn403`:
	fromConn403 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `fromConn403` to result `intoClientConn550`
	// (`intoClientConn550` is now tainted).
	intoClientConn550 := httputil.NewProxyClientConn(fromConn403, nil)

	// Sink the tainted `intoClientConn550`:
	sink(intoClientConn550)
}

func TaintStepTest_NetHttpHttputilNewProxyClientConn_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromReader232` into `intoClientConn791`.

	// Assume that `sourceCQL` has the underlying type of `fromReader232`:
	fromReader232 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader232` to result `intoClientConn791`
	// (`intoClientConn791` is now tainted).
	intoClientConn791 := httputil.NewProxyClientConn(nil, fromReader232)

	// Sink the tainted `intoClientConn791`:
	sink(intoClientConn791)
}

func TaintStepTest_NetHttpHttputilNewProxyClientConn_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromClientConn548` into `intoConn944`.

	// Assume that `sourceCQL` has the underlying type of `fromClientConn548`:
	fromClientConn548 := sourceCQL.(*httputil.ClientConn)

	// Declare `intoConn944` variable:
	var intoConn944 net.Conn

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoConn944`:
	intermediateCQL := httputil.NewProxyClientConn(intoConn944, nil)

	// Extra step (`fromClientConn548` taints `intermediateCQL`, which taints `intoConn944`:
	link(fromClientConn548, intermediateCQL)

	// Sink the tainted `intoConn944`:
	sink(intoConn944)
}

func TaintStepTest_NetHttpHttputilClientConnHijack_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromClientConn437` into `intoConn341`.

	// Assume that `sourceCQL` has the underlying type of `fromClientConn437`:
	fromClientConn437 := sourceCQL.(httputil.ClientConn)

	// Call the method that transfers the taint
	// from the receiver `fromClientConn437` to the result `intoConn341`
	// (`intoConn341` is now tainted).
	intoConn341, _ := fromClientConn437.Hijack()

	// Sink the tainted `intoConn341`:
	sink(intoConn341)
}

func TaintStepTest_NetHttpHttputilClientConnHijack_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromClientConn913` into `intoReader514`.

	// Assume that `sourceCQL` has the underlying type of `fromClientConn913`:
	fromClientConn913 := sourceCQL.(httputil.ClientConn)

	// Call the method that transfers the taint
	// from the receiver `fromClientConn913` to the result `intoReader514`
	// (`intoReader514` is now tainted).
	_, intoReader514 := fromClientConn913.Hijack()

	// Sink the tainted `intoReader514`:
	sink(intoReader514)
}

func TaintStepTest_NetHttpHttputilServerConnHijack_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromServerConn731` into `intoConn762`.

	// Assume that `sourceCQL` has the underlying type of `fromServerConn731`:
	fromServerConn731 := sourceCQL.(httputil.ServerConn)

	// Call the method that transfers the taint
	// from the receiver `fromServerConn731` to the result `intoConn762`
	// (`intoConn762` is now tainted).
	intoConn762, _ := fromServerConn731.Hijack()

	// Sink the tainted `intoConn762`:
	sink(intoConn762)
}

func TaintStepTest_NetHttpHttputilServerConnHijack_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromServerConn250` into `intoReader877`.

	// Assume that `sourceCQL` has the underlying type of `fromServerConn250`:
	fromServerConn250 := sourceCQL.(httputil.ServerConn)

	// Call the method that transfers the taint
	// from the receiver `fromServerConn250` to the result `intoReader877`
	// (`intoReader877` is now tainted).
	_, intoReader877 := fromServerConn250.Hijack()

	// Sink the tainted `intoReader877`:
	sink(intoReader877)
}

func TaintStepTest_NetHttpHttputilBufferPoolGet_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBufferPool156` into `intoByte953`.

	// Assume that `sourceCQL` has the underlying type of `fromBufferPool156`:
	fromBufferPool156 := sourceCQL.(httputil.BufferPool)

	// Call the method that transfers the taint
	// from the receiver `fromBufferPool156` to the result `intoByte953`
	// (`intoByte953` is now tainted).
	intoByte953 := fromBufferPool156.Get()

	// Sink the tainted `intoByte953`:
	sink(intoByte953)
}

func TaintStepTest_NetHttpHttputilBufferPoolPut_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte933` into `intoBufferPool562`.

	// Assume that `sourceCQL` has the underlying type of `fromByte933`:
	fromByte933 := sourceCQL.([]byte)

	// Declare `intoBufferPool562` variable:
	var intoBufferPool562 httputil.BufferPool

	// Call the method that transfers the taint
	// from the parameter `fromByte933` to the receiver `intoBufferPool562`
	// (`intoBufferPool562` is now tainted).
	intoBufferPool562.Put(fromByte933)

	// Sink the tainted `intoBufferPool562`:
	sink(intoBufferPool562)
}

func RunAllTaints_NetHttpHttputil() {
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilDumpRequest_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilDumpRequestOut_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilDumpResponse_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilNewChunkedReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilNewChunkedWriter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilNewClientConn_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilNewClientConn_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilNewClientConn_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilNewProxyClientConn_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilNewProxyClientConn_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilNewProxyClientConn_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilClientConnHijack_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilClientConnHijack_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilServerConnHijack_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilServerConnHijack_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilBufferPoolGet_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHttputilBufferPoolPut_B0I0O0(source)
	}
}
