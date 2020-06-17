// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "syscall"

func TaintStepTest_SyscallBytePtrFromString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString711` into `intoByte975`.

	// Assume that `sourceCQL` has the underlying type of `fromString711`:
	fromString711 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString711` to result `intoByte975`
	// (`intoByte975` is now tainted).
	intoByte975, _ := syscall.BytePtrFromString(fromString711)

	// Return the tainted `intoByte975`:
	return intoByte975
}

func TaintStepTest_SyscallByteSliceFromString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString781` into `intoByte466`.

	// Assume that `sourceCQL` has the underlying type of `fromString781`:
	fromString781 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString781` to result `intoByte466`
	// (`intoByte466` is now tainted).
	intoByte466, _ := syscall.ByteSliceFromString(fromString781)

	// Return the tainted `intoByte466`:
	return intoByte466
}

func TaintStepTest_SyscallPread_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInt237` into `intoByte429`.

	// Assume that `sourceCQL` has the underlying type of `fromInt237`:
	fromInt237 := sourceCQL.(int)

	// Declare `intoByte429` variable:
	var intoByte429 []byte

	// Call the function that transfers the taint
	// from the parameter `fromInt237` to parameter `intoByte429`;
	// `intoByte429` is now tainted.
	syscall.Pread(fromInt237, intoByte429, 0)

	// Return the tainted `intoByte429`:
	return intoByte429
}

func TaintStepTest_SyscallSlicePtrFromStrings_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString135` into `intoByte664`.

	// Assume that `sourceCQL` has the underlying type of `fromString135`:
	fromString135 := sourceCQL.([]string)

	// Call the function that transfers the taint
	// from the parameter `fromString135` to result `intoByte664`
	// (`intoByte664` is now tainted).
	intoByte664, _ := syscall.SlicePtrFromStrings(fromString135)

	// Return the tainted `intoByte664`:
	return intoByte664
}

func TaintStepTest_SyscallStringBytePtr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString514` into `intoByte120`.

	// Assume that `sourceCQL` has the underlying type of `fromString514`:
	fromString514 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString514` to result `intoByte120`
	// (`intoByte120` is now tainted).
	intoByte120 := syscall.StringBytePtr(fromString514)

	// Return the tainted `intoByte120`:
	return intoByte120
}

func TaintStepTest_SyscallStringByteSlice_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString796` into `intoByte178`.

	// Assume that `sourceCQL` has the underlying type of `fromString796`:
	fromString796 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString796` to result `intoByte178`
	// (`intoByte178` is now tainted).
	intoByte178 := syscall.StringByteSlice(fromString796)

	// Return the tainted `intoByte178`:
	return intoByte178
}

func TaintStepTest_SyscallStringSlicePtr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString227` into `intoByte707`.

	// Assume that `sourceCQL` has the underlying type of `fromString227`:
	fromString227 := sourceCQL.([]string)

	// Call the function that transfers the taint
	// from the parameter `fromString227` to result `intoByte707`
	// (`intoByte707` is now tainted).
	intoByte707 := syscall.StringSlicePtr(fromString227)

	// Return the tainted `intoByte707`:
	return intoByte707
}

func TaintStepTest_SyscallUnixCredentials_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUcred694` into `intoByte303`.

	// Assume that `sourceCQL` has the underlying type of `fromUcred694`:
	fromUcred694 := sourceCQL.(*syscall.Ucred)

	// Call the function that transfers the taint
	// from the parameter `fromUcred694` to result `intoByte303`
	// (`intoByte303` is now tainted).
	intoByte303 := syscall.UnixCredentials(fromUcred694)

	// Return the tainted `intoByte303`:
	return intoByte303
}

func TaintStepTest_SyscallRawConnRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn962` into `intoFuncfdUintptrdoneBool681`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn962`:
	fromRawConn962 := sourceCQL.(syscall.RawConn)

	// Declare `intoFuncfdUintptrdoneBool681` variable:
	var intoFuncfdUintptrdoneBool681 func(uintptr) bool

	// Call the method that transfers the taint
	// from the receiver `fromRawConn962` to the argument `intoFuncfdUintptrdoneBool681`
	// (`intoFuncfdUintptrdoneBool681` is now tainted).
	fromRawConn962.Read(intoFuncfdUintptrdoneBool681)

	// Return the tainted `intoFuncfdUintptrdoneBool681`:
	return intoFuncfdUintptrdoneBool681
}

func TaintStepTest_SyscallConnSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn315` into `intoRawConn754`.

	// Assume that `sourceCQL` has the underlying type of `fromConn315`:
	fromConn315 := sourceCQL.(syscall.Conn)

	// Call the method that transfers the taint
	// from the receiver `fromConn315` to the result `intoRawConn754`
	// (`intoRawConn754` is now tainted).
	intoRawConn754, _ := fromConn315.SyscallConn()

	// Return the tainted `intoRawConn754`:
	return intoRawConn754
}

func TaintStepTest_SyscallConnSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn598` into `intoConn377`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn598`:
	fromRawConn598 := sourceCQL.(syscall.RawConn)

	// Declare `intoConn377` variable:
	var intoConn377 syscall.Conn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoConn377`:
	intermediateCQL, _ := intoConn377.SyscallConn()

	// Extra step (`fromRawConn598` taints `intermediateCQL`, which taints `intoConn377`:
	link(fromRawConn598, intermediateCQL)

	// Return the tainted `intoConn377`:
	return intoConn377
}

func TaintStepTest_SyscallRawConnWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFuncfdUintptrdoneBool925` into `intoRawConn643`.

	// Assume that `sourceCQL` has the underlying type of `fromFuncfdUintptrdoneBool925`:
	fromFuncfdUintptrdoneBool925 := sourceCQL.(func(uintptr) bool)

	// Declare `intoRawConn643` variable:
	var intoRawConn643 syscall.RawConn

	// Call the method that transfers the taint
	// from the parameter `fromFuncfdUintptrdoneBool925` to the receiver `intoRawConn643`
	// (`intoRawConn643` is now tainted).
	intoRawConn643.Write(fromFuncfdUintptrdoneBool925)

	// Return the tainted `intoRawConn643`:
	return intoRawConn643
}

func RunAllTaints_Syscall() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyscallBytePtrFromString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyscallByteSliceFromString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyscallPread_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyscallSlicePtrFromStrings_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyscallStringBytePtr_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyscallStringByteSlice_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyscallStringSlicePtr_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyscallUnixCredentials_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyscallRawConnRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyscallConnSyscallConn_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyscallConnSyscallConn_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_SyscallRawConnWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
