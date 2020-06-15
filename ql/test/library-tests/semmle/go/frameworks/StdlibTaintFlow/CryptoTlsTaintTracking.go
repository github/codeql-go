package main

import (
	"crypto/tls"
	"net"
)

func TaintStepTest_CryptoTlsClient_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromConn672` into `intoConn872`.

	// Assume that `sourceCQL` has the underlying type of `fromConn672`:
	fromConn672 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `fromConn672` to result `intoConn872`
	// (`intoConn872` is now tainted).
	intoConn872 := tls.Client(fromConn672, nil)

	// Sink the tainted `intoConn872`:
	sink(intoConn872)
}

func TaintStepTest_CryptoTlsClient_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromConn923` into `intoConn793`.

	// Assume that `sourceCQL` has the underlying type of `fromConn923`:
	fromConn923 := sourceCQL.(*tls.Conn)

	// Declare `intoConn793` variable:
	var intoConn793 net.Conn

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoConn793`:
	intermediateCQL := tls.Client(intoConn793, nil)

	// Extra step (`fromConn923` taints `intermediateCQL`, which taints `intoConn793`:
	link(fromConn923, intermediateCQL)

	// Sink the tainted `intoConn793`:
	sink(intoConn793)
}

func TaintStepTest_CryptoTlsNewListener_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromListener971` into `intoListener681`.

	// Assume that `sourceCQL` has the underlying type of `fromListener971`:
	fromListener971 := sourceCQL.(net.Listener)

	// Call the function that transfers the taint
	// from the parameter `fromListener971` to result `intoListener681`
	// (`intoListener681` is now tainted).
	intoListener681 := tls.NewListener(fromListener971, nil)

	// Sink the tainted `intoListener681`:
	sink(intoListener681)
}

func TaintStepTest_CryptoTlsServer_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromConn440` into `intoConn444`.

	// Assume that `sourceCQL` has the underlying type of `fromConn440`:
	fromConn440 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `fromConn440` to result `intoConn444`
	// (`intoConn444` is now tainted).
	intoConn444 := tls.Server(fromConn440, nil)

	// Sink the tainted `intoConn444`:
	sink(intoConn444)
}

func TaintStepTest_CryptoTlsServer_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromConn606` into `intoConn573`.

	// Assume that `sourceCQL` has the underlying type of `fromConn606`:
	fromConn606 := sourceCQL.(*tls.Conn)

	// Declare `intoConn573` variable:
	var intoConn573 net.Conn

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoConn573`:
	intermediateCQL := tls.Server(intoConn573, nil)

	// Extra step (`fromConn606` taints `intermediateCQL`, which taints `intoConn573`:
	link(fromConn606, intermediateCQL)

	// Sink the tainted `intoConn573`:
	sink(intoConn573)
}

func TaintStepTest_CryptoTlsConnRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromConn850` into `intoByte256`.

	// Assume that `sourceCQL` has the underlying type of `fromConn850`:
	fromConn850 := sourceCQL.(tls.Conn)

	// Declare `intoByte256` variable:
	var intoByte256 []byte

	// Call the method that transfers the taint
	// from the receiver `fromConn850` to the argument `intoByte256`
	// (`intoByte256` is now tainted).
	fromConn850.Read(intoByte256)

	// Sink the tainted `intoByte256`:
	sink(intoByte256)
}

func TaintStepTest_CryptoTlsConnWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte198` into `intoConn120`.

	// Assume that `sourceCQL` has the underlying type of `fromByte198`:
	fromByte198 := sourceCQL.([]byte)

	// Declare `intoConn120` variable:
	var intoConn120 tls.Conn

	// Call the method that transfers the taint
	// from the parameter `fromByte198` to the receiver `intoConn120`
	// (`intoConn120` is now tainted).
	intoConn120.Write(fromByte198)

	// Sink the tainted `intoConn120`:
	sink(intoConn120)
}

func RunAllTaints_CryptoTls() {
	{
		source := newSource()
		TaintStepTest_CryptoTlsClient_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoTlsClient_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoTlsNewListener_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoTlsServer_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoTlsServer_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoTlsConnRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoTlsConnWrite_B0I0O0(source)
	}
}
