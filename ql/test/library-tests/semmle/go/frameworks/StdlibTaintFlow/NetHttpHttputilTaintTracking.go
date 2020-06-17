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
	// The flow is from `fromRequest656` into `intoByte414`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest656`:
	fromRequest656 := sourceCQL.(*http.Request)

	// Call the function that transfers the taint
	// from the parameter `fromRequest656` to result `intoByte414`
	// (`intoByte414` is now tainted).
	intoByte414, _ := httputil.DumpRequest(fromRequest656, false)

	// Return the tainted `intoByte414`:
	return intoByte414
}

func TaintStepTest_NetHttpHttputilDumpRequestOut_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRequest518` into `intoByte650`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest518`:
	fromRequest518 := sourceCQL.(*http.Request)

	// Call the function that transfers the taint
	// from the parameter `fromRequest518` to result `intoByte650`
	// (`intoByte650` is now tainted).
	intoByte650, _ := httputil.DumpRequestOut(fromRequest518, false)

	// Return the tainted `intoByte650`:
	return intoByte650
}

func TaintStepTest_NetHttpHttputilDumpResponse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromResponse784` into `intoByte957`.

	// Assume that `sourceCQL` has the underlying type of `fromResponse784`:
	fromResponse784 := sourceCQL.(*http.Response)

	// Call the function that transfers the taint
	// from the parameter `fromResponse784` to result `intoByte957`
	// (`intoByte957` is now tainted).
	intoByte957, _ := httputil.DumpResponse(fromResponse784, false)

	// Return the tainted `intoByte957`:
	return intoByte957
}

func TaintStepTest_NetHttpHttputilNewChunkedReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader520` into `intoReader443`.

	// Assume that `sourceCQL` has the underlying type of `fromReader520`:
	fromReader520 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader520` to result `intoReader443`
	// (`intoReader443` is now tainted).
	intoReader443 := httputil.NewChunkedReader(fromReader520)

	// Return the tainted `intoReader443`:
	return intoReader443
}

func TaintStepTest_NetHttpHttputilNewChunkedWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser127` into `intoWriter483`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser127`:
	fromWriteCloser127 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter483` variable:
	var intoWriter483 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter483`:
	intermediateCQL := httputil.NewChunkedWriter(intoWriter483)

	// Extra step (`fromWriteCloser127` taints `intermediateCQL`, which taints `intoWriter483`:
	link(fromWriteCloser127, intermediateCQL)

	// Return the tainted `intoWriter483`:
	return intoWriter483
}

func TaintStepTest_NetHttpHttputilNewClientConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn989` into `intoClientConn982`.

	// Assume that `sourceCQL` has the underlying type of `fromConn989`:
	fromConn989 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `fromConn989` to result `intoClientConn982`
	// (`intoClientConn982` is now tainted).
	intoClientConn982 := httputil.NewClientConn(fromConn989, nil)

	// Return the tainted `intoClientConn982`:
	return intoClientConn982
}

func TaintStepTest_NetHttpHttputilNewClientConn_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader417` into `intoClientConn584`.

	// Assume that `sourceCQL` has the underlying type of `fromReader417`:
	fromReader417 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader417` to result `intoClientConn584`
	// (`intoClientConn584` is now tainted).
	intoClientConn584 := httputil.NewClientConn(nil, fromReader417)

	// Return the tainted `intoClientConn584`:
	return intoClientConn584
}

func TaintStepTest_NetHttpHttputilNewClientConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromClientConn991` into `intoConn881`.

	// Assume that `sourceCQL` has the underlying type of `fromClientConn991`:
	fromClientConn991 := sourceCQL.(*httputil.ClientConn)

	// Declare `intoConn881` variable:
	var intoConn881 net.Conn

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoConn881`:
	intermediateCQL := httputil.NewClientConn(intoConn881, nil)

	// Extra step (`fromClientConn991` taints `intermediateCQL`, which taints `intoConn881`:
	link(fromClientConn991, intermediateCQL)

	// Return the tainted `intoConn881`:
	return intoConn881
}

func TaintStepTest_NetHttpHttputilNewProxyClientConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn186` into `intoClientConn284`.

	// Assume that `sourceCQL` has the underlying type of `fromConn186`:
	fromConn186 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `fromConn186` to result `intoClientConn284`
	// (`intoClientConn284` is now tainted).
	intoClientConn284 := httputil.NewProxyClientConn(fromConn186, nil)

	// Return the tainted `intoClientConn284`:
	return intoClientConn284
}

func TaintStepTest_NetHttpHttputilNewProxyClientConn_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader908` into `intoClientConn137`.

	// Assume that `sourceCQL` has the underlying type of `fromReader908`:
	fromReader908 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader908` to result `intoClientConn137`
	// (`intoClientConn137` is now tainted).
	intoClientConn137 := httputil.NewProxyClientConn(nil, fromReader908)

	// Return the tainted `intoClientConn137`:
	return intoClientConn137
}

func TaintStepTest_NetHttpHttputilNewProxyClientConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromClientConn494` into `intoConn873`.

	// Assume that `sourceCQL` has the underlying type of `fromClientConn494`:
	fromClientConn494 := sourceCQL.(*httputil.ClientConn)

	// Declare `intoConn873` variable:
	var intoConn873 net.Conn

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoConn873`:
	intermediateCQL := httputil.NewProxyClientConn(intoConn873, nil)

	// Extra step (`fromClientConn494` taints `intermediateCQL`, which taints `intoConn873`:
	link(fromClientConn494, intermediateCQL)

	// Return the tainted `intoConn873`:
	return intoConn873
}

func TaintStepTest_NetHttpHttputilClientConnHijack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromClientConn599` into `intoConn409`.

	// Assume that `sourceCQL` has the underlying type of `fromClientConn599`:
	fromClientConn599 := sourceCQL.(httputil.ClientConn)

	// Call the method that transfers the taint
	// from the receiver `fromClientConn599` to the result `intoConn409`
	// (`intoConn409` is now tainted).
	intoConn409, _ := fromClientConn599.Hijack()

	// Return the tainted `intoConn409`:
	return intoConn409
}

func TaintStepTest_NetHttpHttputilClientConnHijack_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromClientConn246` into `intoReader898`.

	// Assume that `sourceCQL` has the underlying type of `fromClientConn246`:
	fromClientConn246 := sourceCQL.(httputil.ClientConn)

	// Call the method that transfers the taint
	// from the receiver `fromClientConn246` to the result `intoReader898`
	// (`intoReader898` is now tainted).
	_, intoReader898 := fromClientConn246.Hijack()

	// Return the tainted `intoReader898`:
	return intoReader898
}

func TaintStepTest_NetHttpHttputilServerConnHijack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromServerConn598` into `intoConn631`.

	// Assume that `sourceCQL` has the underlying type of `fromServerConn598`:
	fromServerConn598 := sourceCQL.(httputil.ServerConn)

	// Call the method that transfers the taint
	// from the receiver `fromServerConn598` to the result `intoConn631`
	// (`intoConn631` is now tainted).
	intoConn631, _ := fromServerConn598.Hijack()

	// Return the tainted `intoConn631`:
	return intoConn631
}

func TaintStepTest_NetHttpHttputilServerConnHijack_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromServerConn165` into `intoReader150`.

	// Assume that `sourceCQL` has the underlying type of `fromServerConn165`:
	fromServerConn165 := sourceCQL.(httputil.ServerConn)

	// Call the method that transfers the taint
	// from the receiver `fromServerConn165` to the result `intoReader150`
	// (`intoReader150` is now tainted).
	_, intoReader150 := fromServerConn165.Hijack()

	// Return the tainted `intoReader150`:
	return intoReader150
}

func TaintStepTest_NetHttpHttputilBufferPoolGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBufferPool340` into `intoByte471`.

	// Assume that `sourceCQL` has the underlying type of `fromBufferPool340`:
	fromBufferPool340 := sourceCQL.(httputil.BufferPool)

	// Call the method that transfers the taint
	// from the receiver `fromBufferPool340` to the result `intoByte471`
	// (`intoByte471` is now tainted).
	intoByte471 := fromBufferPool340.Get()

	// Return the tainted `intoByte471`:
	return intoByte471
}

func TaintStepTest_NetHttpHttputilBufferPoolPut_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte290` into `intoBufferPool758`.

	// Assume that `sourceCQL` has the underlying type of `fromByte290`:
	fromByte290 := sourceCQL.([]byte)

	// Declare `intoBufferPool758` variable:
	var intoBufferPool758 httputil.BufferPool

	// Call the method that transfers the taint
	// from the parameter `fromByte290` to the receiver `intoBufferPool758`
	// (`intoBufferPool758` is now tainted).
	intoBufferPool758.Put(fromByte290)

	// Return the tainted `intoBufferPool758`:
	return intoBufferPool758
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
