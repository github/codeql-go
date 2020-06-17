// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"os"
	"syscall"
)

func TaintStepTest_OsExpand_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString157` into `intoString792`.

	// Assume that `sourceCQL` has the underlying type of `fromString157`:
	fromString157 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString157` to result `intoString792`
	// (`intoString792` is now tainted).
	intoString792 := os.Expand(fromString157, nil)

	// Return the tainted `intoString792`:
	return intoString792
}

func TaintStepTest_OsExpandEnv_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString384` into `intoString885`.

	// Assume that `sourceCQL` has the underlying type of `fromString384`:
	fromString384 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString384` to result `intoString885`
	// (`intoString885` is now tainted).
	intoString885 := os.ExpandEnv(fromString384)

	// Return the tainted `intoString885`:
	return intoString885
}

func TaintStepTest_OsNewFile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr535` into `intoFile607`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr535`:
	fromUintptr535 := sourceCQL.(uintptr)

	// Call the function that transfers the taint
	// from the parameter `fromUintptr535` to result `intoFile607`
	// (`intoFile607` is now tainted).
	intoFile607 := os.NewFile(fromUintptr535, "")

	// Return the tainted `intoFile607`:
	return intoFile607
}

func TaintStepTest_OsPipe_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile551` into `intoFile856`.

	// Assume that `sourceCQL` has the underlying type of `fromFile551`:
	fromFile551 := sourceCQL.(*os.File)

	// Call the function that transfers the taint
	// from the result `fromFile551` to result `intoFile856`
	// (extra steps needed)
	intoFile856, intermediateCQL, _ := os.Pipe()

	// Extra step (`fromFile551` taints `intermediateCQL`, which taints `intoFile856`:
	link(fromFile551, intermediateCQL)

	// Return the tainted `intoFile856`:
	return intoFile856
}

func TaintStepTest_OsFileFd_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile623` into `intoUintptr336`.

	// Assume that `sourceCQL` has the underlying type of `fromFile623`:
	fromFile623 := sourceCQL.(os.File)

	// Call the method that transfers the taint
	// from the receiver `fromFile623` to the result `intoUintptr336`
	// (`intoUintptr336` is now tainted).
	intoUintptr336 := fromFile623.Fd()

	// Return the tainted `intoUintptr336`:
	return intoUintptr336
}

func TaintStepTest_OsFileRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile944` into `intoByte827`.

	// Assume that `sourceCQL` has the underlying type of `fromFile944`:
	fromFile944 := sourceCQL.(os.File)

	// Declare `intoByte827` variable:
	var intoByte827 []byte

	// Call the method that transfers the taint
	// from the receiver `fromFile944` to the argument `intoByte827`
	// (`intoByte827` is now tainted).
	fromFile944.Read(intoByte827)

	// Return the tainted `intoByte827`:
	return intoByte827
}

func TaintStepTest_OsFileReadAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile724` into `intoByte451`.

	// Assume that `sourceCQL` has the underlying type of `fromFile724`:
	fromFile724 := sourceCQL.(os.File)

	// Declare `intoByte451` variable:
	var intoByte451 []byte

	// Call the method that transfers the taint
	// from the receiver `fromFile724` to the argument `intoByte451`
	// (`intoByte451` is now tainted).
	fromFile724.ReadAt(intoByte451, 0)

	// Return the tainted `intoByte451`:
	return intoByte451
}

func TaintStepTest_OsFileSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile119` into `intoRawConn913`.

	// Assume that `sourceCQL` has the underlying type of `fromFile119`:
	fromFile119 := sourceCQL.(os.File)

	// Call the method that transfers the taint
	// from the receiver `fromFile119` to the result `intoRawConn913`
	// (`intoRawConn913` is now tainted).
	intoRawConn913, _ := fromFile119.SyscallConn()

	// Return the tainted `intoRawConn913`:
	return intoRawConn913
}

func TaintStepTest_OsFileSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn600` into `intoFile389`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn600`:
	fromRawConn600 := sourceCQL.(syscall.RawConn)

	// Declare `intoFile389` variable:
	var intoFile389 os.File

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoFile389`:
	intermediateCQL, _ := intoFile389.SyscallConn()

	// Extra step (`fromRawConn600` taints `intermediateCQL`, which taints `intoFile389`:
	link(fromRawConn600, intermediateCQL)

	// Return the tainted `intoFile389`:
	return intoFile389
}

func TaintStepTest_OsFileWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte716` into `intoFile593`.

	// Assume that `sourceCQL` has the underlying type of `fromByte716`:
	fromByte716 := sourceCQL.([]byte)

	// Declare `intoFile593` variable:
	var intoFile593 os.File

	// Call the method that transfers the taint
	// from the parameter `fromByte716` to the receiver `intoFile593`
	// (`intoFile593` is now tainted).
	intoFile593.Write(fromByte716)

	// Return the tainted `intoFile593`:
	return intoFile593
}

func TaintStepTest_OsFileWriteAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte136` into `intoFile601`.

	// Assume that `sourceCQL` has the underlying type of `fromByte136`:
	fromByte136 := sourceCQL.([]byte)

	// Declare `intoFile601` variable:
	var intoFile601 os.File

	// Call the method that transfers the taint
	// from the parameter `fromByte136` to the receiver `intoFile601`
	// (`intoFile601` is now tainted).
	intoFile601.WriteAt(fromByte136, 0)

	// Return the tainted `intoFile601`:
	return intoFile601
}

func TaintStepTest_OsFileWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString760` into `intoFile335`.

	// Assume that `sourceCQL` has the underlying type of `fromString760`:
	fromString760 := sourceCQL.(string)

	// Declare `intoFile335` variable:
	var intoFile335 os.File

	// Call the method that transfers the taint
	// from the parameter `fromString760` to the receiver `intoFile335`
	// (`intoFile335` is now tainted).
	intoFile335.WriteString(fromString760)

	// Return the tainted `intoFile335`:
	return intoFile335
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
