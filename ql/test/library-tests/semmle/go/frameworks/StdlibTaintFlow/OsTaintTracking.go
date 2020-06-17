// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"os"
	"syscall"
)

func TaintStepTest_OsExpand_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString656` into `intoString414`.

	// Assume that `sourceCQL` has the underlying type of `fromString656`:
	fromString656 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString656` to result `intoString414`
	// (`intoString414` is now tainted).
	intoString414 := os.Expand(fromString656, nil)

	// Return the tainted `intoString414`:
	return intoString414
}

func TaintStepTest_OsExpandEnv_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString518` into `intoString650`.

	// Assume that `sourceCQL` has the underlying type of `fromString518`:
	fromString518 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString518` to result `intoString650`
	// (`intoString650` is now tainted).
	intoString650 := os.ExpandEnv(fromString518)

	// Return the tainted `intoString650`:
	return intoString650
}

func TaintStepTest_OsNewFile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr784` into `intoFile957`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr784`:
	fromUintptr784 := sourceCQL.(uintptr)

	// Call the function that transfers the taint
	// from the parameter `fromUintptr784` to result `intoFile957`
	// (`intoFile957` is now tainted).
	intoFile957 := os.NewFile(fromUintptr784, "")

	// Return the tainted `intoFile957`:
	return intoFile957
}

func TaintStepTest_OsPipe_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile520` into `intoFile443`.

	// Assume that `sourceCQL` has the underlying type of `fromFile520`:
	fromFile520 := sourceCQL.(*os.File)

	// Call the function that transfers the taint
	// from the result `fromFile520` to result `intoFile443`
	// (extra steps needed)
	intoFile443, intermediateCQL, _ := os.Pipe()

	// Extra step (`fromFile520` taints `intermediateCQL`, which taints `intoFile443`:
	link(fromFile520, intermediateCQL)

	// Return the tainted `intoFile443`:
	return intoFile443
}

func TaintStepTest_OsFileFd_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile127` into `intoUintptr483`.

	// Assume that `sourceCQL` has the underlying type of `fromFile127`:
	fromFile127 := sourceCQL.(os.File)

	// Call the method that transfers the taint
	// from the receiver `fromFile127` to the result `intoUintptr483`
	// (`intoUintptr483` is now tainted).
	intoUintptr483 := fromFile127.Fd()

	// Return the tainted `intoUintptr483`:
	return intoUintptr483
}

func TaintStepTest_OsFileRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile989` into `intoByte982`.

	// Assume that `sourceCQL` has the underlying type of `fromFile989`:
	fromFile989 := sourceCQL.(os.File)

	// Declare `intoByte982` variable:
	var intoByte982 []byte

	// Call the method that transfers the taint
	// from the receiver `fromFile989` to the argument `intoByte982`
	// (`intoByte982` is now tainted).
	fromFile989.Read(intoByte982)

	// Return the tainted `intoByte982`:
	return intoByte982
}

func TaintStepTest_OsFileReadAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile417` into `intoByte584`.

	// Assume that `sourceCQL` has the underlying type of `fromFile417`:
	fromFile417 := sourceCQL.(os.File)

	// Declare `intoByte584` variable:
	var intoByte584 []byte

	// Call the method that transfers the taint
	// from the receiver `fromFile417` to the argument `intoByte584`
	// (`intoByte584` is now tainted).
	fromFile417.ReadAt(intoByte584, 0)

	// Return the tainted `intoByte584`:
	return intoByte584
}

func TaintStepTest_OsFileSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile991` into `intoRawConn881`.

	// Assume that `sourceCQL` has the underlying type of `fromFile991`:
	fromFile991 := sourceCQL.(os.File)

	// Call the method that transfers the taint
	// from the receiver `fromFile991` to the result `intoRawConn881`
	// (`intoRawConn881` is now tainted).
	intoRawConn881, _ := fromFile991.SyscallConn()

	// Return the tainted `intoRawConn881`:
	return intoRawConn881
}

func TaintStepTest_OsFileSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn186` into `intoFile284`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn186`:
	fromRawConn186 := sourceCQL.(syscall.RawConn)

	// Declare `intoFile284` variable:
	var intoFile284 os.File

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoFile284`:
	intermediateCQL, _ := intoFile284.SyscallConn()

	// Extra step (`fromRawConn186` taints `intermediateCQL`, which taints `intoFile284`:
	link(fromRawConn186, intermediateCQL)

	// Return the tainted `intoFile284`:
	return intoFile284
}

func TaintStepTest_OsFileWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte908` into `intoFile137`.

	// Assume that `sourceCQL` has the underlying type of `fromByte908`:
	fromByte908 := sourceCQL.([]byte)

	// Declare `intoFile137` variable:
	var intoFile137 os.File

	// Call the method that transfers the taint
	// from the parameter `fromByte908` to the receiver `intoFile137`
	// (`intoFile137` is now tainted).
	intoFile137.Write(fromByte908)

	// Return the tainted `intoFile137`:
	return intoFile137
}

func TaintStepTest_OsFileWriteAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte494` into `intoFile873`.

	// Assume that `sourceCQL` has the underlying type of `fromByte494`:
	fromByte494 := sourceCQL.([]byte)

	// Declare `intoFile873` variable:
	var intoFile873 os.File

	// Call the method that transfers the taint
	// from the parameter `fromByte494` to the receiver `intoFile873`
	// (`intoFile873` is now tainted).
	intoFile873.WriteAt(fromByte494, 0)

	// Return the tainted `intoFile873`:
	return intoFile873
}

func TaintStepTest_OsFileWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString599` into `intoFile409`.

	// Assume that `sourceCQL` has the underlying type of `fromString599`:
	fromString599 := sourceCQL.(string)

	// Declare `intoFile409` variable:
	var intoFile409 os.File

	// Call the method that transfers the taint
	// from the parameter `fromString599` to the receiver `intoFile409`
	// (`intoFile409` is now tainted).
	intoFile409.WriteString(fromString599)

	// Return the tainted `intoFile409`:
	return intoFile409
}

func RunAllTaints_Os() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_OsExpand_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_OsExpandEnv_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_OsNewFile_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_OsPipe_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_OsFileFd_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_OsFileRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_OsFileReadAt_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_OsFileSyscallConn_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_OsFileSyscallConn_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_OsFileWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_OsFileWriteAt_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_OsFileWriteString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
