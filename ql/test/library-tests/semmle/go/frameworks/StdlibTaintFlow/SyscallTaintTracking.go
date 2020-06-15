package main

import "syscall"

func TaintStepTest_SyscallBytePtrFromString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString966` into `intoByte924`.

	// Assume that `sourceCQL` has the underlying type of `fromString966`:
	fromString966 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString966` to result `intoByte924`
	// (`intoByte924` is now tainted).
	intoByte924, _ := syscall.BytePtrFromString(fromString966)

	// Sink the tainted `intoByte924`:
	sink(intoByte924)
}

func TaintStepTest_SyscallByteSliceFromString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString506` into `intoByte715`.

	// Assume that `sourceCQL` has the underlying type of `fromString506`:
	fromString506 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString506` to result `intoByte715`
	// (`intoByte715` is now tainted).
	intoByte715, _ := syscall.ByteSliceFromString(fromString506)

	// Sink the tainted `intoByte715`:
	sink(intoByte715)
}

func TaintStepTest_SyscallPread_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInt866` into `intoByte990`.

	// Assume that `sourceCQL` has the underlying type of `fromInt866`:
	fromInt866 := sourceCQL.(int)

	// Declare `intoByte990` variable:
	var intoByte990 []byte

	// Call the function that transfers the taint
	// from the parameter `fromInt866` to parameter `intoByte990`;
	// `intoByte990` is now tainted.
	syscall.Pread(fromInt866, intoByte990, 0)

	// Sink the tainted `intoByte990`:
	sink(intoByte990)
}

func TaintStepTest_SyscallSlicePtrFromStrings_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString934` into `intoByte649`.

	// Assume that `sourceCQL` has the underlying type of `fromString934`:
	fromString934 := sourceCQL.([]string)

	// Call the function that transfers the taint
	// from the parameter `fromString934` to result `intoByte649`
	// (`intoByte649` is now tainted).
	intoByte649, _ := syscall.SlicePtrFromStrings(fromString934)

	// Sink the tainted `intoByte649`:
	sink(intoByte649)
}

func TaintStepTest_SyscallStringBytePtr_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString875` into `intoByte692`.

	// Assume that `sourceCQL` has the underlying type of `fromString875`:
	fromString875 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString875` to result `intoByte692`
	// (`intoByte692` is now tainted).
	intoByte692 := syscall.StringBytePtr(fromString875)

	// Sink the tainted `intoByte692`:
	sink(intoByte692)
}

func TaintStepTest_SyscallStringByteSlice_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString139` into `intoByte616`.

	// Assume that `sourceCQL` has the underlying type of `fromString139`:
	fromString139 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString139` to result `intoByte616`
	// (`intoByte616` is now tainted).
	intoByte616 := syscall.StringByteSlice(fromString139)

	// Sink the tainted `intoByte616`:
	sink(intoByte616)
}

func TaintStepTest_SyscallStringSlicePtr_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString117` into `intoByte115`.

	// Assume that `sourceCQL` has the underlying type of `fromString117`:
	fromString117 := sourceCQL.([]string)

	// Call the function that transfers the taint
	// from the parameter `fromString117` to result `intoByte115`
	// (`intoByte115` is now tainted).
	intoByte115 := syscall.StringSlicePtr(fromString117)

	// Sink the tainted `intoByte115`:
	sink(intoByte115)
}

func TaintStepTest_SyscallUnixCredentials_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUcred853` into `intoByte868`.

	// Assume that `sourceCQL` has the underlying type of `fromUcred853`:
	fromUcred853 := sourceCQL.(*syscall.Ucred)

	// Call the function that transfers the taint
	// from the parameter `fromUcred853` to result `intoByte868`
	// (`intoByte868` is now tainted).
	intoByte868 := syscall.UnixCredentials(fromUcred853)

	// Sink the tainted `intoByte868`:
	sink(intoByte868)
}

func TaintStepTest_SyscallRawConnRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRawConn119` into `intoFuncfdUintptrdoneBool205`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn119`:
	fromRawConn119 := sourceCQL.(syscall.RawConn)

	// Declare `intoFuncfdUintptrdoneBool205` variable:
	var intoFuncfdUintptrdoneBool205 func(uintptr) bool

	// Call the method that transfers the taint
	// from the receiver `fromRawConn119` to the argument `intoFuncfdUintptrdoneBool205`
	// (`intoFuncfdUintptrdoneBool205` is now tainted).
	fromRawConn119.Read(intoFuncfdUintptrdoneBool205)

	// Sink the tainted `intoFuncfdUintptrdoneBool205`:
	sink(intoFuncfdUintptrdoneBool205)
}

func TaintStepTest_SyscallConnSyscallConn_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromConn874` into `intoRawConn285`.

	// Assume that `sourceCQL` has the underlying type of `fromConn874`:
	fromConn874 := sourceCQL.(syscall.Conn)

	// Call the method that transfers the taint
	// from the receiver `fromConn874` to the result `intoRawConn285`
	// (`intoRawConn285` is now tainted).
	intoRawConn285, _ := fromConn874.SyscallConn()

	// Sink the tainted `intoRawConn285`:
	sink(intoRawConn285)
}

func TaintStepTest_SyscallConnSyscallConn_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromRawConn852` into `intoConn682`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn852`:
	fromRawConn852 := sourceCQL.(syscall.RawConn)

	// Declare `intoConn682` variable:
	var intoConn682 syscall.Conn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoConn682`:
	intermediateCQL, _ := intoConn682.SyscallConn()

	// Extra step (`fromRawConn852` taints `intermediateCQL`, which taints `intoConn682`:
	link(fromRawConn852, intermediateCQL)

	// Sink the tainted `intoConn682`:
	sink(intoConn682)
}

func TaintStepTest_SyscallRawConnWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFuncfdUintptrdoneBool787` into `intoRawConn349`.

	// Assume that `sourceCQL` has the underlying type of `fromFuncfdUintptrdoneBool787`:
	fromFuncfdUintptrdoneBool787 := sourceCQL.(func(uintptr) bool)

	// Declare `intoRawConn349` variable:
	var intoRawConn349 syscall.RawConn

	// Call the method that transfers the taint
	// from the parameter `fromFuncfdUintptrdoneBool787` to the receiver `intoRawConn349`
	// (`intoRawConn349` is now tainted).
	intoRawConn349.Write(fromFuncfdUintptrdoneBool787)

	// Sink the tainted `intoRawConn349`:
	sink(intoRawConn349)
}

func RunAllTaints_Syscall() {
	{
		source := newSource()
		TaintStepTest_SyscallBytePtrFromString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallByteSliceFromString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallPread_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallSlicePtrFromStrings_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallStringBytePtr_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallStringByteSlice_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallStringSlicePtr_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallUnixCredentials_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallRawConnRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallConnSyscallConn_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallConnSyscallConn_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallRawConnWrite_B0I0O0(source)
	}
}
