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

func TaintStepTest_SyscallPread_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInt784` into `intoByte957`.

	// Assume that `sourceCQL` has the underlying type of `fromInt784`:
	fromInt784 := sourceCQL.(int)

	// Declare `intoByte957` variable:
	var intoByte957 []byte

	// Call the function that transfers the taint
	// from the parameter `fromInt784` to parameter `intoByte957`;
	// `intoByte957` is now tainted.
	syscall.Pread(fromInt784, intoByte957, 0)

	// Return the tainted `intoByte957`:
	return intoByte957
}

func TaintStepTest_SyscallSlicePtrFromStrings_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString520` into `intoByte443`.

	// Assume that `sourceCQL` has the underlying type of `fromString520`:
	fromString520 := sourceCQL.([]string)

	// Call the function that transfers the taint
	// from the parameter `fromString520` to result `intoByte443`
	// (`intoByte443` is now tainted).
	intoByte443, _ := syscall.SlicePtrFromStrings(fromString520)

	// Return the tainted `intoByte443`:
	return intoByte443
}

func TaintStepTest_SyscallStringBytePtr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString127` into `intoByte483`.

	// Assume that `sourceCQL` has the underlying type of `fromString127`:
	fromString127 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString127` to result `intoByte483`
	// (`intoByte483` is now tainted).
	intoByte483 := syscall.StringBytePtr(fromString127)

	// Return the tainted `intoByte483`:
	return intoByte483
}

func TaintStepTest_SyscallStringByteSlice_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString989` into `intoByte982`.

	// Assume that `sourceCQL` has the underlying type of `fromString989`:
	fromString989 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString989` to result `intoByte982`
	// (`intoByte982` is now tainted).
	intoByte982 := syscall.StringByteSlice(fromString989)

	// Return the tainted `intoByte982`:
	return intoByte982
}

func TaintStepTest_SyscallStringSlicePtr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString417` into `intoByte584`.

	// Assume that `sourceCQL` has the underlying type of `fromString417`:
	fromString417 := sourceCQL.([]string)

	// Call the function that transfers the taint
	// from the parameter `fromString417` to result `intoByte584`
	// (`intoByte584` is now tainted).
	intoByte584 := syscall.StringSlicePtr(fromString417)

	// Return the tainted `intoByte584`:
	return intoByte584
}

func TaintStepTest_SyscallUnixCredentials_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUcred991` into `intoByte881`.

	// Assume that `sourceCQL` has the underlying type of `fromUcred991`:
	fromUcred991 := sourceCQL.(*syscall.Ucred)

	// Call the function that transfers the taint
	// from the parameter `fromUcred991` to result `intoByte881`
	// (`intoByte881` is now tainted).
	intoByte881 := syscall.UnixCredentials(fromUcred991)

	// Return the tainted `intoByte881`:
	return intoByte881
}

func TaintStepTest_SyscallRawConnRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn186` into `intoFuncfdUintptrdoneBool284`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn186`:
	fromRawConn186 := sourceCQL.(syscall.RawConn)

	// Declare `intoFuncfdUintptrdoneBool284` variable:
	var intoFuncfdUintptrdoneBool284 func(uintptr) bool

	// Call the method that transfers the taint
	// from the receiver `fromRawConn186` to the argument `intoFuncfdUintptrdoneBool284`
	// (`intoFuncfdUintptrdoneBool284` is now tainted).
	fromRawConn186.Read(intoFuncfdUintptrdoneBool284)

	// Return the tainted `intoFuncfdUintptrdoneBool284`:
	return intoFuncfdUintptrdoneBool284
}

func TaintStepTest_SyscallConnSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn908` into `intoRawConn137`.

	// Assume that `sourceCQL` has the underlying type of `fromConn908`:
	fromConn908 := sourceCQL.(syscall.Conn)

	// Call the method that transfers the taint
	// from the receiver `fromConn908` to the result `intoRawConn137`
	// (`intoRawConn137` is now tainted).
	intoRawConn137, _ := fromConn908.SyscallConn()

	// Return the tainted `intoRawConn137`:
	return intoRawConn137
}

func TaintStepTest_SyscallConnSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn494` into `intoConn873`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn494`:
	fromRawConn494 := sourceCQL.(syscall.RawConn)

	// Declare `intoConn873` variable:
	var intoConn873 syscall.Conn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoConn873`:
	intermediateCQL, _ := intoConn873.SyscallConn()

	// Extra step (`fromRawConn494` taints `intermediateCQL`, which taints `intoConn873`:
	link(fromRawConn494, intermediateCQL)

	// Return the tainted `intoConn873`:
	return intoConn873
}

func TaintStepTest_SyscallRawConnWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFuncfdUintptrdoneBool599` into `intoRawConn409`.

	// Assume that `sourceCQL` has the underlying type of `fromFuncfdUintptrdoneBool599`:
	fromFuncfdUintptrdoneBool599 := sourceCQL.(func(uintptr) bool)

	// Declare `intoRawConn409` variable:
	var intoRawConn409 syscall.RawConn

	// Call the method that transfers the taint
	// from the parameter `fromFuncfdUintptrdoneBool599` to the receiver `intoRawConn409`
	// (`intoRawConn409` is now tainted).
	intoRawConn409.Write(fromFuncfdUintptrdoneBool599)

	// Return the tainted `intoRawConn409`:
	return intoRawConn409
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
