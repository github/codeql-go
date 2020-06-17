package main

import (
	"crypto/tls"
	"net"
)

func TaintStepTest_CryptoTlsClient_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn302` into `intoConn731`.

	// Assume that `sourceCQL` has the underlying type of `fromConn302`:
	fromConn302 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `fromConn302` to result `intoConn731`
	// (`intoConn731` is now tainted).
	intoConn731 := tls.Client(fromConn302, nil)

	// Return the tainted `intoConn731`:
	return intoConn731
}

func TaintStepTest_CryptoTlsClient_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn248` into `intoConn257`.

	// Assume that `sourceCQL` has the underlying type of `fromConn248`:
	fromConn248 := sourceCQL.(*tls.Conn)

	// Declare `intoConn257` variable:
	var intoConn257 net.Conn

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoConn257`:
	intermediateCQL := tls.Client(intoConn257, nil)

	// Extra step (`fromConn248` taints `intermediateCQL`, which taints `intoConn257`:
	link(fromConn248, intermediateCQL)

	// Return the tainted `intoConn257`:
	return intoConn257
}

func TaintStepTest_CryptoTlsNewListener_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromListener564` into `intoListener777`.

	// Assume that `sourceCQL` has the underlying type of `fromListener564`:
	fromListener564 := sourceCQL.(net.Listener)

	// Call the function that transfers the taint
	// from the parameter `fromListener564` to result `intoListener777`
	// (`intoListener777` is now tainted).
	intoListener777 := tls.NewListener(fromListener564, nil)

	// Return the tainted `intoListener777`:
	return intoListener777
}

func TaintStepTest_CryptoTlsServer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn483` into `intoConn708`.

	// Assume that `sourceCQL` has the underlying type of `fromConn483`:
	fromConn483 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `fromConn483` to result `intoConn708`
	// (`intoConn708` is now tainted).
	intoConn708 := tls.Server(fromConn483, nil)

	// Return the tainted `intoConn708`:
	return intoConn708
}

func TaintStepTest_CryptoTlsServer_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn466` into `intoConn297`.

	// Assume that `sourceCQL` has the underlying type of `fromConn466`:
	fromConn466 := sourceCQL.(*tls.Conn)

	// Declare `intoConn297` variable:
	var intoConn297 net.Conn

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoConn297`:
	intermediateCQL := tls.Server(intoConn297, nil)

	// Extra step (`fromConn466` taints `intermediateCQL`, which taints `intoConn297`:
	link(fromConn466, intermediateCQL)

	// Return the tainted `intoConn297`:
	return intoConn297
}

func TaintStepTest_CryptoTlsConnRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn458` into `intoByte489`.

	// Assume that `sourceCQL` has the underlying type of `fromConn458`:
	fromConn458 := sourceCQL.(tls.Conn)

	// Declare `intoByte489` variable:
	var intoByte489 []byte

	// Call the method that transfers the taint
	// from the receiver `fromConn458` to the argument `intoByte489`
	// (`intoByte489` is now tainted).
	fromConn458.Read(intoByte489)

	// Return the tainted `intoByte489`:
	return intoByte489
}

func TaintStepTest_CryptoTlsConnWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte558` into `intoConn443`.

	// Assume that `sourceCQL` has the underlying type of `fromByte558`:
	fromByte558 := sourceCQL.([]byte)

	// Declare `intoConn443` variable:
	var intoConn443 tls.Conn

	// Call the method that transfers the taint
	// from the parameter `fromByte558` to the receiver `intoConn443`
	// (`intoConn443` is now tainted).
	intoConn443.Write(fromByte558)

	// Return the tainted `intoConn443`:
	return intoConn443
}

func RunAllTaints_CryptoTls() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoTlsClient_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoTlsClient_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoTlsNewListener_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoTlsServer_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoTlsServer_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoTlsConnRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CryptoTlsConnWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
