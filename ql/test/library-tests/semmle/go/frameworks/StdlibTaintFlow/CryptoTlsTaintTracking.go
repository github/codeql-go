// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"crypto/tls"
	"net"
)

func TaintStepTest_CryptoTlsClient_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn656` into `intoConn414`.

	// Assume that `sourceCQL` has the underlying type of `fromConn656`:
	fromConn656 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `fromConn656` to result `intoConn414`
	// (`intoConn414` is now tainted).
	intoConn414 := tls.Client(fromConn656, nil)

	// Return the tainted `intoConn414`:
	return intoConn414
}

func TaintStepTest_CryptoTlsClient_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn518` into `intoConn650`.

	// Assume that `sourceCQL` has the underlying type of `fromConn518`:
	fromConn518 := sourceCQL.(*tls.Conn)

	// Declare `intoConn650` variable:
	var intoConn650 net.Conn

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoConn650`:
	intermediateCQL := tls.Client(intoConn650, nil)

	// Extra step (`fromConn518` taints `intermediateCQL`, which taints `intoConn650`:
	link(fromConn518, intermediateCQL)

	// Return the tainted `intoConn650`:
	return intoConn650
}

func TaintStepTest_CryptoTlsNewListener_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromListener784` into `intoListener957`.

	// Assume that `sourceCQL` has the underlying type of `fromListener784`:
	fromListener784 := sourceCQL.(net.Listener)

	// Call the function that transfers the taint
	// from the parameter `fromListener784` to result `intoListener957`
	// (`intoListener957` is now tainted).
	intoListener957 := tls.NewListener(fromListener784, nil)

	// Return the tainted `intoListener957`:
	return intoListener957
}

func TaintStepTest_CryptoTlsServer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn520` into `intoConn443`.

	// Assume that `sourceCQL` has the underlying type of `fromConn520`:
	fromConn520 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `fromConn520` to result `intoConn443`
	// (`intoConn443` is now tainted).
	intoConn443 := tls.Server(fromConn520, nil)

	// Return the tainted `intoConn443`:
	return intoConn443
}

func TaintStepTest_CryptoTlsServer_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn127` into `intoConn483`.

	// Assume that `sourceCQL` has the underlying type of `fromConn127`:
	fromConn127 := sourceCQL.(*tls.Conn)

	// Declare `intoConn483` variable:
	var intoConn483 net.Conn

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoConn483`:
	intermediateCQL := tls.Server(intoConn483, nil)

	// Extra step (`fromConn127` taints `intermediateCQL`, which taints `intoConn483`:
	link(fromConn127, intermediateCQL)

	// Return the tainted `intoConn483`:
	return intoConn483
}

func TaintStepTest_CryptoTlsConnRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn989` into `intoByte982`.

	// Assume that `sourceCQL` has the underlying type of `fromConn989`:
	fromConn989 := sourceCQL.(tls.Conn)

	// Declare `intoByte982` variable:
	var intoByte982 []byte

	// Call the method that transfers the taint
	// from the receiver `fromConn989` to the argument `intoByte982`
	// (`intoByte982` is now tainted).
	fromConn989.Read(intoByte982)

	// Return the tainted `intoByte982`:
	return intoByte982
}

func TaintStepTest_CryptoTlsConnWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte417` into `intoConn584`.

	// Assume that `sourceCQL` has the underlying type of `fromByte417`:
	fromByte417 := sourceCQL.([]byte)

	// Declare `intoConn584` variable:
	var intoConn584 tls.Conn

	// Call the method that transfers the taint
	// from the parameter `fromByte417` to the receiver `intoConn584`
	// (`intoConn584` is now tainted).
	intoConn584.Write(fromByte417)

	// Return the tainted `intoConn584`:
	return intoConn584
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
