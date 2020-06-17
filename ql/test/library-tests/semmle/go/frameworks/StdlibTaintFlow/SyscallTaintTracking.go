// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "syscall"

func TaintStepTest_SyscallBytePtrFromString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString656` into `intoByte414`.

	// Assume that `sourceCQL` has the underlying type of `fromString656`:
	fromString656 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString656` to result `intoByte414`
	// (`intoByte414` is now tainted).
	intoByte414, _ := syscall.BytePtrFromString(fromString656)

	// Return the tainted `intoByte414`:
	return intoByte414
}

func TaintStepTest_SyscallByteSliceFromString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString518` into `intoByte650`.

	// Assume that `sourceCQL` has the underlying type of `fromString518`:
	fromString518 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString518` to result `intoByte650`
	// (`intoByte650` is now tainted).
	intoByte650, _ := syscall.ByteSliceFromString(fromString518)

	// Return the tainted `intoByte650`:
	return intoByte650
}

func TaintStepTest_SyscallStringBytePtr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString784` into `intoByte957`.

	// Assume that `sourceCQL` has the underlying type of `fromString784`:
	fromString784 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString784` to result `intoByte957`
	// (`intoByte957` is now tainted).
	intoByte957 := syscall.StringBytePtr(fromString784)

	// Return the tainted `intoByte957`:
	return intoByte957
}

func TaintStepTest_SyscallStringByteSlice_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString520` into `intoByte443`.

	// Assume that `sourceCQL` has the underlying type of `fromString520`:
	fromString520 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString520` to result `intoByte443`
	// (`intoByte443` is now tainted).
	intoByte443 := syscall.StringByteSlice(fromString520)

	// Return the tainted `intoByte443`:
	return intoByte443
}

func TaintStepTest_SyscallRawConnRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn127` into `intoFuncfdUintptrdoneBool483`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn127`:
	fromRawConn127 := sourceCQL.(syscall.RawConn)

	// Declare `intoFuncfdUintptrdoneBool483` variable:
	var intoFuncfdUintptrdoneBool483 func(uintptr) bool

	// Call the method that transfers the taint
	// from the receiver `fromRawConn127` to the argument `intoFuncfdUintptrdoneBool483`
	// (`intoFuncfdUintptrdoneBool483` is now tainted).
	fromRawConn127.Read(intoFuncfdUintptrdoneBool483)

	// Return the tainted `intoFuncfdUintptrdoneBool483`:
	return intoFuncfdUintptrdoneBool483
}

func TaintStepTest_SyscallConnSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn989` into `intoRawConn982`.

	// Assume that `sourceCQL` has the underlying type of `fromConn989`:
	fromConn989 := sourceCQL.(syscall.Conn)

	// Call the method that transfers the taint
	// from the receiver `fromConn989` to the result `intoRawConn982`
	// (`intoRawConn982` is now tainted).
	intoRawConn982, _ := fromConn989.SyscallConn()

	// Return the tainted `intoRawConn982`:
	return intoRawConn982
}

func TaintStepTest_SyscallConnSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn417` into `intoConn584`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn417`:
	fromRawConn417 := sourceCQL.(syscall.RawConn)

	// Declare `intoConn584` variable:
	var intoConn584 syscall.Conn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoConn584`:
	intermediateCQL, _ := intoConn584.SyscallConn()

	// Extra step (`fromRawConn417` taints `intermediateCQL`, which taints `intoConn584`:
	link(fromRawConn417, intermediateCQL)

	// Return the tainted `intoConn584`:
	return intoConn584
}

func TaintStepTest_SyscallRawConnWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFuncfdUintptrdoneBool991` into `intoRawConn881`.

	// Assume that `sourceCQL` has the underlying type of `fromFuncfdUintptrdoneBool991`:
	fromFuncfdUintptrdoneBool991 := sourceCQL.(func(uintptr) bool)

	// Declare `intoRawConn881` variable:
	var intoRawConn881 syscall.RawConn

	// Call the method that transfers the taint
	// from the parameter `fromFuncfdUintptrdoneBool991` to the receiver `intoRawConn881`
	// (`intoRawConn881` is now tainted).
	intoRawConn881.Write(fromFuncfdUintptrdoneBool991)

	// Return the tainted `intoRawConn881`:
	return intoRawConn881
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
