package main

import "syscall"

func TaintStepTest_SyscallBytePtrFromString(sourceCQL interface{}) {
	// The flow is from `s` into `into807`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into807`
	// (`into807` is now tainted).
	into807, _ := syscall.BytePtrFromString(s)

	// Sink the tainted `into807`:
	sink(into807)
}

func TaintStepTest_SyscallByteSliceFromString(sourceCQL interface{}) {
	// The flow is from `s` into `into273`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into273`
	// (`into273` is now tainted).
	into273, _ := syscall.ByteSliceFromString(s)

	// Sink the tainted `into273`:
	sink(into273)
}

func TaintStepTest_SyscallPread(sourceCQL interface{}) {
	// The flow is from `fd` into `p`.

	// Assume that `sourceCQL` has the underlying type of `fd`:
	fd := sourceCQL.(int)

	// Declare `p` variable:
	var p []byte

	// Call the function that transfers the taint
	// from the parameter `fd` to parameter `p`;
	// `p` is now tainted.
	syscall.Pread(fd, p, 0)

	// Sink the tainted `p`:
	sink(p)
}

func TaintStepTest_SyscallRead(sourceCQL interface{}) {
	// The flow is from `fd` into `p`.

	// Assume that `sourceCQL` has the underlying type of `fd`:
	fd := sourceCQL.(int)

	// Declare `p` variable:
	var p []byte

	// Call the function that transfers the taint
	// from the parameter `fd` to parameter `p`;
	// `p` is now tainted.
	syscall.Read(fd, p)

	// Sink the tainted `p`:
	sink(p)
}

func TaintStepTest_SyscallReadDirent(sourceCQL interface{}) {
	// The flow is from `fd` into `buf`.

	// Assume that `sourceCQL` has the underlying type of `fd`:
	fd := sourceCQL.(int)

	// Declare `buf` variable:
	var buf []byte

	// Call the function that transfers the taint
	// from the parameter `fd` to parameter `buf`;
	// `buf` is now tainted.
	syscall.ReadDirent(fd, buf)

	// Sink the tainted `buf`:
	sink(buf)
}

func TaintStepTest_SyscallSlicePtrFromStrings(sourceCQL interface{}) {
	// The flow is from `ss` into `into767`.

	// Assume that `sourceCQL` has the underlying type of `ss`:
	ss := sourceCQL.([]string)

	// Call the function that transfers the taint
	// from the parameter `ss` to result `into767`
	// (`into767` is now tainted).
	into767, _ := syscall.SlicePtrFromStrings(ss)

	// Sink the tainted `into767`:
	sink(into767)
}

func TaintStepTest_SyscallStringBytePtr(sourceCQL interface{}) {
	// The flow is from `s` into `into862`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into862`
	// (`into862` is now tainted).
	into862 := syscall.StringBytePtr(s)

	// Sink the tainted `into862`:
	sink(into862)
}

func TaintStepTest_SyscallStringByteSlice(sourceCQL interface{}) {
	// The flow is from `s` into `into759`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into759`
	// (`into759` is now tainted).
	into759 := syscall.StringByteSlice(s)

	// Sink the tainted `into759`:
	sink(into759)
}

func TaintStepTest_SyscallStringSlicePtr(sourceCQL interface{}) {
	// The flow is from `ss` into `into370`.

	// Assume that `sourceCQL` has the underlying type of `ss`:
	ss := sourceCQL.([]string)

	// Call the function that transfers the taint
	// from the parameter `ss` to result `into370`
	// (`into370` is now tainted).
	into370 := syscall.StringSlicePtr(ss)

	// Sink the tainted `into370`:
	sink(into370)
}

func TaintStepTest_SyscallUnixCredentials(sourceCQL interface{}) {
	// The flow is from `ucred` into `into519`.

	// Assume that `sourceCQL` has the underlying type of `ucred`:
	ucred := sourceCQL.(*syscall.Ucred)

	// Call the function that transfers the taint
	// from the parameter `ucred` to result `into519`
	// (`into519` is now tainted).
	into519 := syscall.UnixCredentials(ucred)

	// Sink the tainted `into519`:
	sink(into519)
}

func TaintStepTest_SyscallUnixRights(sourceCQL interface{}) {
	// The flow is from `fds` into `into120`.

	// Assume that `sourceCQL` has the underlying type of `fds`:
	fds := sourceCQL.(int)

	// Call the function that transfers the taint
	// from the parameter `fds` to result `into120`
	// (`into120` is now tainted).
	into120 := syscall.UnixRights(fds)

	// Sink the tainted `into120`:
	sink(into120)
}

func TaintStepTest_SyscallRawConnRead(sourceCQL interface{}) {
	// The flow is from `from611` into `f`.

	// Assume that `sourceCQL` has the underlying type of `from611`:
	from611 := sourceCQL.(syscall.RawConn)

	// Declare `f` variable:
	var f func(uintptr) bool

	// Call the method that transfers the taint
	// from the receiver `from611` to the argument `f`
	// (`f` is now tainted).
	from611.Read(f)

	// Sink the tainted `f`:
	sink(f)
}

func TaintStepTest_SyscallRawConnWrite(sourceCQL interface{}) {
	// The flow is from `f` into `into293`.

	// Assume that `sourceCQL` has the underlying type of `f`:
	f := sourceCQL.(func(uintptr) bool)

	// Declare `into293` variable:
	var into293 syscall.RawConn

	// Call the method that transfers the taint
	// from the parameter `f` to the receiver `into293`
	// (`into293` is now tainted).
	into293.Write(f)

	// Sink the tainted `into293`:
	sink(into293)
}

func RunAllTaints_Syscall(v interface{}) {
	{
		source := newSource()
		TaintStepTest_SyscallBytePtrFromString(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallByteSliceFromString(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallPread(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallRead(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallReadDirent(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallSlicePtrFromStrings(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallStringBytePtr(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallStringByteSlice(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallStringSlicePtr(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallUnixCredentials(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallUnixRights(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallRawConnRead(source)
	}
	{
		source := newSource()
		TaintStepTest_SyscallRawConnWrite(source)
	}
}
