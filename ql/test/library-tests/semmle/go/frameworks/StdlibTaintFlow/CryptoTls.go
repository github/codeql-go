package main

import (
	"crypto/tls"
	"net"
)

func TaintStepTest_CryptoTlsClient(sourceCQL interface{}) {
	// The flow is from `conn` into `into366`.

	// Assume that `sourceCQL` has the underlying type of `conn`:
	conn := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `conn` to result `into366`
	// (`into366` is now tainted).
	into366 := tls.Client(conn, nil)

	// Sink the tainted `into366`:
	sink(into366)
}

func TaintStepTest_CryptoTlsDial(sourceCQL interface{}) {
	// The flow is from `addr` into `into507`.

	// Assume that `sourceCQL` has the underlying type of `addr`:
	addr := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `addr` to result `into507`
	// (`into507` is now tainted).
	into507, _ := tls.Dial("", addr, nil)

	// Sink the tainted `into507`:
	sink(into507)
}

func TaintStepTest_CryptoTlsDialWithDialer(sourceCQL interface{}) {
	// The flow is from `addr` into `into419`.

	// Assume that `sourceCQL` has the underlying type of `addr`:
	addr := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `addr` to result `into419`
	// (`into419` is now tainted).
	into419, _ := tls.DialWithDialer(nil, "", addr, nil)

	// Sink the tainted `into419`:
	sink(into419)
}

func TaintStepTest_CryptoTlsNewListener(sourceCQL interface{}) {
	// The flow is from `inner` into `into778`.

	// Assume that `sourceCQL` has the underlying type of `inner`:
	inner := sourceCQL.(net.Listener)

	// Call the function that transfers the taint
	// from the parameter `inner` to result `into778`
	// (`into778` is now tainted).
	into778 := tls.NewListener(inner, nil)

	// Sink the tainted `into778`:
	sink(into778)
}

func TaintStepTest_CryptoTlsServer(sourceCQL interface{}) {
	// The flow is from `conn` into `into901`.

	// Assume that `sourceCQL` has the underlying type of `conn`:
	conn := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `conn` to result `into901`
	// (`into901` is now tainted).
	into901 := tls.Server(conn, nil)

	// Sink the tainted `into901`:
	sink(into901)
}

func TaintStepTest_CryptoTlsConnRead(sourceCQL interface{}) {
	// The flow is from `from748` into `b`.

	// Assume that `sourceCQL` has the underlying type of `from748`:
	from748 := sourceCQL.(tls.Conn)

	// Declare `b` variable:
	var b []byte

	// Call the method that transfers the taint
	// from the receiver `from748` to the argument `b`
	// (`b` is now tainted).
	from748.Read(b)

	// Sink the tainted `b`:
	sink(b)
}

func TaintStepTest_CryptoTlsConnWrite(sourceCQL interface{}) {
	// The flow is from `b` into `into454`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare `into454` variable:
	var into454 tls.Conn

	// Call the method that transfers the taint
	// from the parameter `b` to the receiver `into454`
	// (`into454` is now tainted).
	into454.Write(b)

	// Sink the tainted `into454`:
	sink(into454)
}

func RunAllTaints_CryptoTls(v interface{}) {
	{
		source := newSource()
		TaintStepTest_CryptoTlsClient(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoTlsDial(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoTlsDialWithDialer(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoTlsNewListener(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoTlsServer(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoTlsConnRead(source)
	}
	{
		source := newSource()
		TaintStepTest_CryptoTlsConnWrite(source)
	}
}
