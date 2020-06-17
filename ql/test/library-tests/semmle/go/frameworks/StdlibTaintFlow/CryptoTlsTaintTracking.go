// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"crypto/tls"
	"net"
)

func TaintStepTest_CryptoTlsClient_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn950` into `intoConn947`.

	// Assume that `sourceCQL` has the underlying type of `fromConn950`:
	fromConn950 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `fromConn950` to result `intoConn947`
	// (`intoConn947` is now tainted).
	intoConn947 := tls.Client(fromConn950, nil)

	// Return the tainted `intoConn947`:
	return intoConn947
}

func TaintStepTest_CryptoTlsClient_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn833` into `intoConn612`.

	// Assume that `sourceCQL` has the underlying type of `fromConn833`:
	fromConn833 := sourceCQL.(*tls.Conn)

	// Declare `intoConn612` variable:
	var intoConn612 net.Conn

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoConn612`:
	intermediateCQL := tls.Client(intoConn612, nil)

	// Extra step (`fromConn833` taints `intermediateCQL`, which taints `intoConn612`:
	link(fromConn833, intermediateCQL)

	// Return the tainted `intoConn612`:
	return intoConn612
}

func TaintStepTest_CryptoTlsNewListener_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromListener763` into `intoListener571`.

	// Assume that `sourceCQL` has the underlying type of `fromListener763`:
	fromListener763 := sourceCQL.(net.Listener)

	// Call the function that transfers the taint
	// from the parameter `fromListener763` to result `intoListener571`
	// (`intoListener571` is now tainted).
	intoListener571 := tls.NewListener(fromListener763, nil)

	// Return the tainted `intoListener571`:
	return intoListener571
}

func TaintStepTest_CryptoTlsServer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn211` into `intoConn662`.

	// Assume that `sourceCQL` has the underlying type of `fromConn211`:
	fromConn211 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the parameter `fromConn211` to result `intoConn662`
	// (`intoConn662` is now tainted).
	intoConn662 := tls.Server(fromConn211, nil)

	// Return the tainted `intoConn662`:
	return intoConn662
}

func TaintStepTest_CryptoTlsServer_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn565` into `intoConn138`.

	// Assume that `sourceCQL` has the underlying type of `fromConn565`:
	fromConn565 := sourceCQL.(*tls.Conn)

	// Declare `intoConn138` variable:
	var intoConn138 net.Conn

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoConn138`:
	intermediateCQL := tls.Server(intoConn138, nil)

	// Extra step (`fromConn565` taints `intermediateCQL`, which taints `intoConn138`:
	link(fromConn565, intermediateCQL)

	// Return the tainted `intoConn138`:
	return intoConn138
}

func TaintStepTest_CryptoTlsConnRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn127` into `intoByte467`.

	// Assume that `sourceCQL` has the underlying type of `fromConn127`:
	fromConn127 := sourceCQL.(tls.Conn)

	// Declare `intoByte467` variable:
	var intoByte467 []byte

	// Call the method that transfers the taint
	// from the receiver `fromConn127` to the argument `intoByte467`
	// (`intoByte467` is now tainted).
	fromConn127.Read(intoByte467)

	// Return the tainted `intoByte467`:
	return intoByte467
}

func TaintStepTest_CryptoTlsConnWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte748` into `intoConn123`.

	// Assume that `sourceCQL` has the underlying type of `fromByte748`:
	fromByte748 := sourceCQL.([]byte)

	// Declare `intoConn123` variable:
	var intoConn123 tls.Conn

	// Call the method that transfers the taint
	// from the parameter `fromByte748` to the receiver `intoConn123`
	// (`intoConn123` is now tainted).
	intoConn123.Write(fromByte748)

	// Return the tainted `intoConn123`:
	return intoConn123
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
