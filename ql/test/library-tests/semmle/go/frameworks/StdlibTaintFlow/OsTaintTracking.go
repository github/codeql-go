package main

import (
	"os"
	"syscall"
)

func TaintStepTest_OsExpand_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString353` into `intoString726`.

	// Assume that `sourceCQL` has the underlying type of `fromString353`:
	fromString353 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString353` to result `intoString726`
	// (`intoString726` is now tainted).
	intoString726 := os.Expand(fromString353, nil)

	// Return the tainted `intoString726`:
	return intoString726
}

func TaintStepTest_OsExpandEnv_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString332` into `intoString740`.

	// Assume that `sourceCQL` has the underlying type of `fromString332`:
	fromString332 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString332` to result `intoString740`
	// (`intoString740` is now tainted).
	intoString740 := os.ExpandEnv(fromString332)

	// Return the tainted `intoString740`:
	return intoString740
}

func TaintStepTest_OsNewFile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUintptr182` into `intoFile857`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr182`:
	fromUintptr182 := sourceCQL.(uintptr)

	// Call the function that transfers the taint
	// from the parameter `fromUintptr182` to result `intoFile857`
	// (`intoFile857` is now tainted).
	intoFile857 := os.NewFile(fromUintptr182, "")

	// Return the tainted `intoFile857`:
	return intoFile857
}

func TaintStepTest_OsPipe_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile699` into `intoFile324`.

	// Assume that `sourceCQL` has the underlying type of `fromFile699`:
	fromFile699 := sourceCQL.(*os.File)

	// Call the function that transfers the taint
	// from the result `fromFile699` to result `intoFile324`
	// (extra steps needed)
	intoFile324, intermediateCQL, _ := os.Pipe()

	// Extra step (`fromFile699` taints `intermediateCQL`, which taints `intoFile324`:
	link(fromFile699, intermediateCQL)

	// Return the tainted `intoFile324`:
	return intoFile324
}

func TaintStepTest_OsFileFd_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile206` into `intoUintptr804`.

	// Assume that `sourceCQL` has the underlying type of `fromFile206`:
	fromFile206 := sourceCQL.(os.File)

	// Call the method that transfers the taint
	// from the receiver `fromFile206` to the result `intoUintptr804`
	// (`intoUintptr804` is now tainted).
	intoUintptr804 := fromFile206.Fd()

	// Return the tainted `intoUintptr804`:
	return intoUintptr804
}

func TaintStepTest_OsFileRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile229` into `intoByte908`.

	// Assume that `sourceCQL` has the underlying type of `fromFile229`:
	fromFile229 := sourceCQL.(os.File)

	// Declare `intoByte908` variable:
	var intoByte908 []byte

	// Call the method that transfers the taint
	// from the receiver `fromFile229` to the argument `intoByte908`
	// (`intoByte908` is now tainted).
	fromFile229.Read(intoByte908)

	// Return the tainted `intoByte908`:
	return intoByte908
}

func TaintStepTest_OsFileReadAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile122` into `intoByte368`.

	// Assume that `sourceCQL` has the underlying type of `fromFile122`:
	fromFile122 := sourceCQL.(os.File)

	// Declare `intoByte368` variable:
	var intoByte368 []byte

	// Call the method that transfers the taint
	// from the receiver `fromFile122` to the argument `intoByte368`
	// (`intoByte368` is now tainted).
	fromFile122.ReadAt(intoByte368, 0)

	// Return the tainted `intoByte368`:
	return intoByte368
}

func TaintStepTest_OsFileSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile847` into `intoRawConn313`.

	// Assume that `sourceCQL` has the underlying type of `fromFile847`:
	fromFile847 := sourceCQL.(os.File)

	// Call the method that transfers the taint
	// from the receiver `fromFile847` to the result `intoRawConn313`
	// (`intoRawConn313` is now tainted).
	intoRawConn313, _ := fromFile847.SyscallConn()

	// Return the tainted `intoRawConn313`:
	return intoRawConn313
}

func TaintStepTest_OsFileSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn662` into `intoFile399`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn662`:
	fromRawConn662 := sourceCQL.(syscall.RawConn)

	// Declare `intoFile399` variable:
	var intoFile399 os.File

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoFile399`:
	intermediateCQL, _ := intoFile399.SyscallConn()

	// Extra step (`fromRawConn662` taints `intermediateCQL`, which taints `intoFile399`:
	link(fromRawConn662, intermediateCQL)

	// Return the tainted `intoFile399`:
	return intoFile399
}

func TaintStepTest_OsFileWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte198` into `intoFile963`.

	// Assume that `sourceCQL` has the underlying type of `fromByte198`:
	fromByte198 := sourceCQL.([]byte)

	// Declare `intoFile963` variable:
	var intoFile963 os.File

	// Call the method that transfers the taint
	// from the parameter `fromByte198` to the receiver `intoFile963`
	// (`intoFile963` is now tainted).
	intoFile963.Write(fromByte198)

	// Return the tainted `intoFile963`:
	return intoFile963
}

func TaintStepTest_OsFileWriteAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte382` into `intoFile321`.

	// Assume that `sourceCQL` has the underlying type of `fromByte382`:
	fromByte382 := sourceCQL.([]byte)

	// Declare `intoFile321` variable:
	var intoFile321 os.File

	// Call the method that transfers the taint
	// from the parameter `fromByte382` to the receiver `intoFile321`
	// (`intoFile321` is now tainted).
	intoFile321.WriteAt(fromByte382, 0)

	// Return the tainted `intoFile321`:
	return intoFile321
}

func TaintStepTest_OsFileWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString417` into `intoFile855`.

	// Assume that `sourceCQL` has the underlying type of `fromString417`:
	fromString417 := sourceCQL.(string)

	// Declare `intoFile855` variable:
	var intoFile855 os.File

	// Call the method that transfers the taint
	// from the parameter `fromString417` to the receiver `intoFile855`
	// (`intoFile855` is now tainted).
	intoFile855.WriteString(fromString417)

	// Return the tainted `intoFile855`:
	return intoFile855
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
