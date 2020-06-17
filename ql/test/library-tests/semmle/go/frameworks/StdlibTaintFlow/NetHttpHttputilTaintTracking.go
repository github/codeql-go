package main

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
)

func TaintStepTest_NetHttpHttputilDumpRequest_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRequest115` into `intoByte212`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest115`:
	fromRequest115 := sourceCQL.(*http.Request)

	// Call the function that transfers the taint
	// from the parameter `fromRequest115` to result `intoByte212`
	// (`intoByte212` is now tainted).
	intoByte212, _ := httputil.DumpRequest(fromRequest115, false)

	// Return the tainted `intoByte212`:
	return intoByte212
}

func TaintStepTest_NetHttpHttputilDumpRequestOut_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRequest984` into `intoByte942`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest984`:
	fromRequest984 := sourceCQL.(*http.Request)

	// Call the function that transfers the taint
	// from the parameter `fromRequest984` to result `intoByte942`
	// (`intoByte942` is now tainted).
	intoByte942, _ := httputil.DumpRequestOut(fromRequest984, false)

	// Return the tainted `intoByte942`:
	return intoByte942
}

func TaintStepTest_NetHttpHttputilDumpResponse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromResponse959` into `intoByte263`.

	// Assume that `sourceCQL` has the underlying type of `fromResponse959`:
	fromResponse959 := sourceCQL.(*http.Response)

	// Call the function that transfers the taint
	// from the parameter `fromResponse959` to result `intoByte263`
	// (`intoByte263` is now tainted).
	intoByte263, _ := httputil.DumpResponse(fromResponse959, false)

	// Return the tainted `intoByte263`:
	return intoByte263
}

func TaintStepTest_NetHttpHttputilNewChunkedReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader868` into `intoReader630`.

	// Assume that `sourceCQL` has the underlying type of `fromReader868`:
	fromReader868 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader868` to result `intoReader630`
	// (`intoReader630` is now tainted).
	intoReader630 := httputil.NewChunkedReader(fromReader868)

	// Return the tainted `intoReader630`:
	return intoReader630
}

func TaintStepTest_NetHttpHttputilNewChunkedWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser640` into `intoWriter462`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser640`:
	fromWriteCloser640 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter462` variable:
	var intoWriter462 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter462`:
	intermediateCQL := httputil.NewChunkedWriter(intoWriter462)

	// Extra step (`fromWriteCloser640` taints `intermediateCQL`, which taints `intoWriter462`:
	link(fromWriteCloser640, intermediateCQL)

	// Return the tainted `intoWriter462`:
	return intoWriter462
}

func TaintStepTest_NetHttpHttputilNewClientConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn148` into `intoClientConn976`.

	// Assume that `sourceCQL` has the underlying type of `fromConn148`:
	fromConn148 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `fromConn148` to result `intoClientConn976`
	// (`intoClientConn976` is now tainted).
	intoClientConn976 := httputil.NewClientConn(fromConn148, nil)

	// Return the tainted `intoClientConn976`:
	return intoClientConn976
}

func TaintStepTest_NetHttpHttputilNewClientConn_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader940` into `intoClientConn639`.

	// Assume that `sourceCQL` has the underlying type of `fromReader940`:
	fromReader940 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader940` to result `intoClientConn639`
	// (`intoClientConn639` is now tainted).
	intoClientConn639 := httputil.NewClientConn(nil, fromReader940)

	// Return the tainted `intoClientConn639`:
	return intoClientConn639
}

func TaintStepTest_NetHttpHttputilNewClientConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromClientConn899` into `intoConn824`.

	// Assume that `sourceCQL` has the underlying type of `fromClientConn899`:
	fromClientConn899 := sourceCQL.(*httputil.ClientConn)

	// Declare `intoConn824` variable:
	var intoConn824 net.Conn

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoConn824`:
	intermediateCQL := httputil.NewClientConn(intoConn824, nil)

	// Extra step (`fromClientConn899` taints `intermediateCQL`, which taints `intoConn824`:
	link(fromClientConn899, intermediateCQL)

	// Return the tainted `intoConn824`:
	return intoConn824
}

func TaintStepTest_NetHttpHttputilNewProxyClientConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn875` into `intoClientConn885`.

	// Assume that `sourceCQL` has the underlying type of `fromConn875`:
	fromConn875 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `fromConn875` to result `intoClientConn885`
	// (`intoClientConn885` is now tainted).
	intoClientConn885 := httputil.NewProxyClientConn(fromConn875, nil)

	// Return the tainted `intoClientConn885`:
	return intoClientConn885
}

func TaintStepTest_NetHttpHttputilNewProxyClientConn_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader882` into `intoClientConn232`.

	// Assume that `sourceCQL` has the underlying type of `fromReader882`:
	fromReader882 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader882` to result `intoClientConn232`
	// (`intoClientConn232` is now tainted).
	intoClientConn232 := httputil.NewProxyClientConn(nil, fromReader882)

	// Return the tainted `intoClientConn232`:
	return intoClientConn232
}

func TaintStepTest_NetHttpHttputilNewProxyClientConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromClientConn291` into `intoConn338`.

	// Assume that `sourceCQL` has the underlying type of `fromClientConn291`:
	fromClientConn291 := sourceCQL.(*httputil.ClientConn)

	// Declare `intoConn338` variable:
	var intoConn338 net.Conn

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoConn338`:
	intermediateCQL := httputil.NewProxyClientConn(intoConn338, nil)

	// Extra step (`fromClientConn291` taints `intermediateCQL`, which taints `intoConn338`:
	link(fromClientConn291, intermediateCQL)

	// Return the tainted `intoConn338`:
	return intoConn338
}

func TaintStepTest_NetHttpHttputilClientConnHijack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromClientConn609` into `intoConn782`.

	// Assume that `sourceCQL` has the underlying type of `fromClientConn609`:
	fromClientConn609 := sourceCQL.(httputil.ClientConn)

	// Call the method that transfers the taint
	// from the receiver `fromClientConn609` to the result `intoConn782`
	// (`intoConn782` is now tainted).
	intoConn782, _ := fromClientConn609.Hijack()

	// Return the tainted `intoConn782`:
	return intoConn782
}

func TaintStepTest_NetHttpHttputilClientConnHijack_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromClientConn748` into `intoReader520`.

	// Assume that `sourceCQL` has the underlying type of `fromClientConn748`:
	fromClientConn748 := sourceCQL.(httputil.ClientConn)

	// Call the method that transfers the taint
	// from the receiver `fromClientConn748` to the result `intoReader520`
	// (`intoReader520` is now tainted).
	_, intoReader520 := fromClientConn748.Hijack()

	// Return the tainted `intoReader520`:
	return intoReader520
}

func TaintStepTest_NetHttpHttputilServerConnHijack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromServerConn635` into `intoConn985`.

	// Assume that `sourceCQL` has the underlying type of `fromServerConn635`:
	fromServerConn635 := sourceCQL.(httputil.ServerConn)

	// Call the method that transfers the taint
	// from the receiver `fromServerConn635` to the result `intoConn985`
	// (`intoConn985` is now tainted).
	intoConn985, _ := fromServerConn635.Hijack()

	// Return the tainted `intoConn985`:
	return intoConn985
}

func TaintStepTest_NetHttpHttputilServerConnHijack_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromServerConn828` into `intoReader555`.

	// Assume that `sourceCQL` has the underlying type of `fromServerConn828`:
	fromServerConn828 := sourceCQL.(httputil.ServerConn)

	// Call the method that transfers the taint
	// from the receiver `fromServerConn828` to the result `intoReader555`
	// (`intoReader555` is now tainted).
	_, intoReader555 := fromServerConn828.Hijack()

	// Return the tainted `intoReader555`:
	return intoReader555
}

func TaintStepTest_NetHttpHttputilBufferPoolGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBufferPool906` into `intoByte675`.

	// Assume that `sourceCQL` has the underlying type of `fromBufferPool906`:
	fromBufferPool906 := sourceCQL.(httputil.BufferPool)

	// Call the method that transfers the taint
	// from the receiver `fromBufferPool906` to the result `intoByte675`
	// (`intoByte675` is now tainted).
	intoByte675 := fromBufferPool906.Get()

	// Return the tainted `intoByte675`:
	return intoByte675
}

func TaintStepTest_NetHttpHttputilBufferPoolPut_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte650` into `intoBufferPool472`.

	// Assume that `sourceCQL` has the underlying type of `fromByte650`:
	fromByte650 := sourceCQL.([]byte)

	// Declare `intoBufferPool472` variable:
	var intoBufferPool472 httputil.BufferPool

	// Call the method that transfers the taint
	// from the parameter `fromByte650` to the receiver `intoBufferPool472`
	// (`intoBufferPool472` is now tainted).
	intoBufferPool472.Put(fromByte650)

	// Return the tainted `intoBufferPool472`:
	return intoBufferPool472
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
