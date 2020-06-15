package main

import (
	"os"
	"syscall"
)

func TaintStepTest_OsExpand_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString327` into `intoString475`.

	// Assume that `sourceCQL` has the underlying type of `fromString327`:
	fromString327 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString327` to result `intoString475`
	// (`intoString475` is now tainted).
	intoString475 := os.Expand(fromString327, nil)

	// Sink the tainted `intoString475`:
	sink(intoString475)
}

func TaintStepTest_OsExpandEnv_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString162` into `intoString120`.

	// Assume that `sourceCQL` has the underlying type of `fromString162`:
	fromString162 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString162` to result `intoString120`
	// (`intoString120` is now tainted).
	intoString120 := os.ExpandEnv(fromString162)

	// Sink the tainted `intoString120`:
	sink(intoString120)
}

func TaintStepTest_OsNewFile_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUintptr625` into `intoFile308`.

	// Assume that `sourceCQL` has the underlying type of `fromUintptr625`:
	fromUintptr625 := sourceCQL.(uintptr)

	// Call the function that transfers the taint
	// from the parameter `fromUintptr625` to result `intoFile308`
	// (`intoFile308` is now tainted).
	intoFile308 := os.NewFile(fromUintptr625, "")

	// Sink the tainted `intoFile308`:
	sink(intoFile308)
}

func TaintStepTest_OsPipe_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFile533` into `intoFile796`.

	// Assume that `sourceCQL` has the underlying type of `fromFile533`:
	fromFile533 := sourceCQL.(*os.File)

	// Call the function that transfers the taint
	// from the result `fromFile533` to result `intoFile796`
	// (extra steps needed)
	intoFile796, intermediateCQL, _ := os.Pipe()

	// Extra step (`fromFile533` taints `intermediateCQL`, which taints `intoFile796`:
	link(fromFile533, intermediateCQL)

	// Sink the tainted `intoFile796`:
	sink(intoFile796)
}

func TaintStepTest_OsFileFd_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFile793` into `intoUintptr528`.

	// Assume that `sourceCQL` has the underlying type of `fromFile793`:
	fromFile793 := sourceCQL.(os.File)

	// Call the method that transfers the taint
	// from the receiver `fromFile793` to the result `intoUintptr528`
	// (`intoUintptr528` is now tainted).
	intoUintptr528 := fromFile793.Fd()

	// Sink the tainted `intoUintptr528`:
	sink(intoUintptr528)
}

func TaintStepTest_OsFileRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFile984` into `intoByte513`.

	// Assume that `sourceCQL` has the underlying type of `fromFile984`:
	fromFile984 := sourceCQL.(os.File)

	// Declare `intoByte513` variable:
	var intoByte513 []byte

	// Call the method that transfers the taint
	// from the receiver `fromFile984` to the argument `intoByte513`
	// (`intoByte513` is now tainted).
	fromFile984.Read(intoByte513)

	// Sink the tainted `intoByte513`:
	sink(intoByte513)
}

func TaintStepTest_OsFileReadAt_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFile629` into `intoByte542`.

	// Assume that `sourceCQL` has the underlying type of `fromFile629`:
	fromFile629 := sourceCQL.(os.File)

	// Declare `intoByte542` variable:
	var intoByte542 []byte

	// Call the method that transfers the taint
	// from the receiver `fromFile629` to the argument `intoByte542`
	// (`intoByte542` is now tainted).
	fromFile629.ReadAt(intoByte542, 0)

	// Sink the tainted `intoByte542`:
	sink(intoByte542)
}

func TaintStepTest_OsFileSyscallConn_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFile773` into `intoRawConn250`.

	// Assume that `sourceCQL` has the underlying type of `fromFile773`:
	fromFile773 := sourceCQL.(os.File)

	// Call the method that transfers the taint
	// from the receiver `fromFile773` to the result `intoRawConn250`
	// (`intoRawConn250` is now tainted).
	intoRawConn250, _ := fromFile773.SyscallConn()

	// Sink the tainted `intoRawConn250`:
	sink(intoRawConn250)
}

func TaintStepTest_OsFileSyscallConn_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromRawConn282` into `intoFile262`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn282`:
	fromRawConn282 := sourceCQL.(syscall.RawConn)

	// Declare `intoFile262` variable:
	var intoFile262 os.File

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoFile262`:
	intermediateCQL, _ := intoFile262.SyscallConn()

	// Extra step (`fromRawConn282` taints `intermediateCQL`, which taints `intoFile262`:
	link(fromRawConn282, intermediateCQL)

	// Sink the tainted `intoFile262`:
	sink(intoFile262)
}

func TaintStepTest_OsFileWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte175` into `intoFile345`.

	// Assume that `sourceCQL` has the underlying type of `fromByte175`:
	fromByte175 := sourceCQL.([]byte)

	// Declare `intoFile345` variable:
	var intoFile345 os.File

	// Call the method that transfers the taint
	// from the parameter `fromByte175` to the receiver `intoFile345`
	// (`intoFile345` is now tainted).
	intoFile345.Write(fromByte175)

	// Sink the tainted `intoFile345`:
	sink(intoFile345)
}

func TaintStepTest_OsFileWriteAt_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte355` into `intoFile709`.

	// Assume that `sourceCQL` has the underlying type of `fromByte355`:
	fromByte355 := sourceCQL.([]byte)

	// Declare `intoFile709` variable:
	var intoFile709 os.File

	// Call the method that transfers the taint
	// from the parameter `fromByte355` to the receiver `intoFile709`
	// (`intoFile709` is now tainted).
	intoFile709.WriteAt(fromByte355, 0)

	// Sink the tainted `intoFile709`:
	sink(intoFile709)
}

func TaintStepTest_OsFileWriteString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString930` into `intoFile605`.

	// Assume that `sourceCQL` has the underlying type of `fromString930`:
	fromString930 := sourceCQL.(string)

	// Declare `intoFile605` variable:
	var intoFile605 os.File

	// Call the method that transfers the taint
	// from the parameter `fromString930` to the receiver `intoFile605`
	// (`intoFile605` is now tainted).
	intoFile605.WriteString(fromString930)

	// Sink the tainted `intoFile605`:
	sink(intoFile605)
}

func RunAllTaints_Os() {
	{
		source := newSource()
		TaintStepTest_OsExpand_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_OsExpandEnv_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_OsNewFile_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_OsPipe_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_OsFileFd_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_OsFileRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_OsFileReadAt_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_OsFileSyscallConn_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_OsFileSyscallConn_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_OsFileWrite_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_OsFileWriteAt_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_OsFileWriteString_B0I0O0(source)
	}
}
