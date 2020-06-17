package main

import "syscall"

func TaintStepTest_SyscallBytePtrFromString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString506` into `intoByte294`.

	// Assume that `sourceCQL` has the underlying type of `fromString506`:
	fromString506 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString506` to result `intoByte294`
	// (`intoByte294` is now tainted).
	intoByte294, _ := syscall.BytePtrFromString(fromString506)

	// Return the tainted `intoByte294`:
	return intoByte294
}

func TaintStepTest_SyscallByteSliceFromString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString301` into `intoByte693`.

	// Assume that `sourceCQL` has the underlying type of `fromString301`:
	fromString301 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString301` to result `intoByte693`
	// (`intoByte693` is now tainted).
	intoByte693, _ := syscall.ByteSliceFromString(fromString301)

	// Return the tainted `intoByte693`:
	return intoByte693
}

func TaintStepTest_SyscallPread_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInt793` into `intoByte920`.

	// Assume that `sourceCQL` has the underlying type of `fromInt793`:
	fromInt793 := sourceCQL.(int)

	// Declare `intoByte920` variable:
	var intoByte920 []byte

	// Call the function that transfers the taint
	// from the parameter `fromInt793` to parameter `intoByte920`;
	// `intoByte920` is now tainted.
	syscall.Pread(fromInt793, intoByte920, 0)

	// Return the tainted `intoByte920`:
	return intoByte920
}

func TaintStepTest_SyscallSlicePtrFromStrings_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString276` into `intoByte846`.

	// Assume that `sourceCQL` has the underlying type of `fromString276`:
	fromString276 := sourceCQL.([]string)

	// Call the function that transfers the taint
	// from the parameter `fromString276` to result `intoByte846`
	// (`intoByte846` is now tainted).
	intoByte846, _ := syscall.SlicePtrFromStrings(fromString276)

	// Return the tainted `intoByte846`:
	return intoByte846
}

func TaintStepTest_SyscallStringBytePtr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString636` into `intoByte655`.

	// Assume that `sourceCQL` has the underlying type of `fromString636`:
	fromString636 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString636` to result `intoByte655`
	// (`intoByte655` is now tainted).
	intoByte655 := syscall.StringBytePtr(fromString636)

	// Return the tainted `intoByte655`:
	return intoByte655
}

func TaintStepTest_SyscallStringByteSlice_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString545` into `intoByte614`.

	// Assume that `sourceCQL` has the underlying type of `fromString545`:
	fromString545 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString545` to result `intoByte614`
	// (`intoByte614` is now tainted).
	intoByte614 := syscall.StringByteSlice(fromString545)

	// Return the tainted `intoByte614`:
	return intoByte614
}

func TaintStepTest_SyscallStringSlicePtr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString729` into `intoByte943`.

	// Assume that `sourceCQL` has the underlying type of `fromString729`:
	fromString729 := sourceCQL.([]string)

	// Call the function that transfers the taint
	// from the parameter `fromString729` to result `intoByte943`
	// (`intoByte943` is now tainted).
	intoByte943 := syscall.StringSlicePtr(fromString729)

	// Return the tainted `intoByte943`:
	return intoByte943
}

func TaintStepTest_SyscallUnixCredentials_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUcred940` into `intoByte660`.

	// Assume that `sourceCQL` has the underlying type of `fromUcred940`:
	fromUcred940 := sourceCQL.(*syscall.Ucred)

	// Call the function that transfers the taint
	// from the parameter `fromUcred940` to result `intoByte660`
	// (`intoByte660` is now tainted).
	intoByte660 := syscall.UnixCredentials(fromUcred940)

	// Return the tainted `intoByte660`:
	return intoByte660
}

func TaintStepTest_SyscallRawConnRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn718` into `intoFuncfdUintptrdoneBool883`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn718`:
	fromRawConn718 := sourceCQL.(syscall.RawConn)

	// Declare `intoFuncfdUintptrdoneBool883` variable:
	var intoFuncfdUintptrdoneBool883 func(uintptr) bool

	// Call the method that transfers the taint
	// from the receiver `fromRawConn718` to the argument `intoFuncfdUintptrdoneBool883`
	// (`intoFuncfdUintptrdoneBool883` is now tainted).
	fromRawConn718.Read(intoFuncfdUintptrdoneBool883)

	// Return the tainted `intoFuncfdUintptrdoneBool883`:
	return intoFuncfdUintptrdoneBool883
}

func TaintStepTest_SyscallConnSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn564` into `intoRawConn694`.

	// Assume that `sourceCQL` has the underlying type of `fromConn564`:
	fromConn564 := sourceCQL.(syscall.Conn)

	// Call the method that transfers the taint
	// from the receiver `fromConn564` to the result `intoRawConn694`
	// (`intoRawConn694` is now tainted).
	intoRawConn694, _ := fromConn564.SyscallConn()

	// Return the tainted `intoRawConn694`:
	return intoRawConn694
}

func TaintStepTest_SyscallConnSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn121` into `intoConn137`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn121`:
	fromRawConn121 := sourceCQL.(syscall.RawConn)

	// Declare `intoConn137` variable:
	var intoConn137 syscall.Conn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoConn137`:
	intermediateCQL, _ := intoConn137.SyscallConn()

	// Extra step (`fromRawConn121` taints `intermediateCQL`, which taints `intoConn137`:
	link(fromRawConn121, intermediateCQL)

	// Return the tainted `intoConn137`:
	return intoConn137
}

func TaintStepTest_SyscallRawConnWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFuncfdUintptrdoneBool243` into `intoRawConn216`.

	// Assume that `sourceCQL` has the underlying type of `fromFuncfdUintptrdoneBool243`:
	fromFuncfdUintptrdoneBool243 := sourceCQL.(func(uintptr) bool)

	// Declare `intoRawConn216` variable:
	var intoRawConn216 syscall.RawConn

	// Call the method that transfers the taint
	// from the parameter `fromFuncfdUintptrdoneBool243` to the receiver `intoRawConn216`
	// (`intoRawConn216` is now tainted).
	intoRawConn216.Write(fromFuncfdUintptrdoneBool243)

	// Return the tainted `intoRawConn216`:
	return intoRawConn216
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
