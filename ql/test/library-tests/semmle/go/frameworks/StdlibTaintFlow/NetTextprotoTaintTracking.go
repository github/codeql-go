package main

import (
	"bufio"
	"io"
	"net/textproto"
)

func TaintStepTest_NetTextprotoCanonicalMIMEHeaderKey_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString885` into `intoString414`.

	// Assume that `sourceCQL` has the underlying type of `fromString885`:
	fromString885 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString885` to result `intoString414`
	// (`intoString414` is now tainted).
	intoString414 := textproto.CanonicalMIMEHeaderKey(fromString885)

	// Return the tainted `intoString414`:
	return intoString414
}

func TaintStepTest_NetTextprotoNewConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReadWriteCloser454` into `intoConn420`.

	// Assume that `sourceCQL` has the underlying type of `fromReadWriteCloser454`:
	fromReadWriteCloser454 := sourceCQL.(io.ReadWriteCloser)

	// Call the function that transfers the taint
	// from the parameter `fromReadWriteCloser454` to result `intoConn420`
	// (`intoConn420` is now tainted).
	intoConn420 := textproto.NewConn(fromReadWriteCloser454)

	// Return the tainted `intoConn420`:
	return intoConn420
}

func TaintStepTest_NetTextprotoNewConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn252` into `intoReadWriteCloser768`.

	// Assume that `sourceCQL` has the underlying type of `fromConn252`:
	fromConn252 := sourceCQL.(*textproto.Conn)

	// Declare `intoReadWriteCloser768` variable:
	var intoReadWriteCloser768 io.ReadWriteCloser

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoReadWriteCloser768`:
	intermediateCQL := textproto.NewConn(intoReadWriteCloser768)

	// Extra step (`fromConn252` taints `intermediateCQL`, which taints `intoReadWriteCloser768`:
	link(fromConn252, intermediateCQL)

	// Return the tainted `intoReadWriteCloser768`:
	return intoReadWriteCloser768
}

func TaintStepTest_NetTextprotoNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader720` into `intoReader875`.

	// Assume that `sourceCQL` has the underlying type of `fromReader720`:
	fromReader720 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader720` to result `intoReader875`
	// (`intoReader875` is now tainted).
	intoReader875 := textproto.NewReader(fromReader720)

	// Return the tainted `intoReader875`:
	return intoReader875
}

func TaintStepTest_NetTextprotoNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter898` into `intoWriter761`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter898`:
	fromWriter898 := sourceCQL.(*textproto.Writer)

	// Declare `intoWriter761` variable:
	var intoWriter761 *bufio.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter761`:
	intermediateCQL := textproto.NewWriter(intoWriter761)

	// Extra step (`fromWriter898` taints `intermediateCQL`, which taints `intoWriter761`:
	link(fromWriter898, intermediateCQL)

	// Return the tainted `intoWriter761`:
	return intoWriter761
}

func TaintStepTest_NetTextprotoTrimBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte436` into `intoByte300`.

	// Assume that `sourceCQL` has the underlying type of `fromByte436`:
	fromByte436 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte436` to result `intoByte300`
	// (`intoByte300` is now tainted).
	intoByte300 := textproto.TrimBytes(fromByte436)

	// Return the tainted `intoByte300`:
	return intoByte300
}

func TaintStepTest_NetTextprotoTrimString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString773` into `intoString353`.

	// Assume that `sourceCQL` has the underlying type of `fromString773`:
	fromString773 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString773` to result `intoString353`
	// (`intoString353` is now tainted).
	intoString353 := textproto.TrimString(fromString773)

	// Return the tainted `intoString353`:
	return intoString353
}

func TaintStepTest_NetTextprotoMIMEHeaderAdd_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString378` into `intoMIMEHeader414`.

	// Assume that `sourceCQL` has the underlying type of `fromString378`:
	fromString378 := sourceCQL.(string)

	// Declare `intoMIMEHeader414` variable:
	var intoMIMEHeader414 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `fromString378` to the receiver `intoMIMEHeader414`
	// (`intoMIMEHeader414` is now tainted).
	intoMIMEHeader414.Add(fromString378, "")

	// Return the tainted `intoMIMEHeader414`:
	return intoMIMEHeader414
}

func TaintStepTest_NetTextprotoMIMEHeaderAdd_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString769` into `intoMIMEHeader454`.

	// Assume that `sourceCQL` has the underlying type of `fromString769`:
	fromString769 := sourceCQL.(string)

	// Declare `intoMIMEHeader454` variable:
	var intoMIMEHeader454 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `fromString769` to the receiver `intoMIMEHeader454`
	// (`intoMIMEHeader454` is now tainted).
	intoMIMEHeader454.Add("", fromString769)

	// Return the tainted `intoMIMEHeader454`:
	return intoMIMEHeader454
}

func TaintStepTest_NetTextprotoMIMEHeaderGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMIMEHeader511` into `intoString506`.

	// Assume that `sourceCQL` has the underlying type of `fromMIMEHeader511`:
	fromMIMEHeader511 := sourceCQL.(textproto.MIMEHeader)

	// Call the method that transfers the taint
	// from the receiver `fromMIMEHeader511` to the result `intoString506`
	// (`intoString506` is now tainted).
	intoString506 := fromMIMEHeader511.Get("")

	// Return the tainted `intoString506`:
	return intoString506
}

func TaintStepTest_NetTextprotoMIMEHeaderSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString980` into `intoMIMEHeader137`.

	// Assume that `sourceCQL` has the underlying type of `fromString980`:
	fromString980 := sourceCQL.(string)

	// Declare `intoMIMEHeader137` variable:
	var intoMIMEHeader137 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `fromString980` to the receiver `intoMIMEHeader137`
	// (`intoMIMEHeader137` is now tainted).
	intoMIMEHeader137.Set(fromString980, "")

	// Return the tainted `intoMIMEHeader137`:
	return intoMIMEHeader137
}

func TaintStepTest_NetTextprotoMIMEHeaderSet_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString525` into `intoMIMEHeader713`.

	// Assume that `sourceCQL` has the underlying type of `fromString525`:
	fromString525 := sourceCQL.(string)

	// Declare `intoMIMEHeader713` variable:
	var intoMIMEHeader713 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `fromString525` to the receiver `intoMIMEHeader713`
	// (`intoMIMEHeader713` is now tainted).
	intoMIMEHeader713.Set("", fromString525)

	// Return the tainted `intoMIMEHeader713`:
	return intoMIMEHeader713
}

func TaintStepTest_NetTextprotoMIMEHeaderValues_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMIMEHeader695` into `intoString862`.

	// Assume that `sourceCQL` has the underlying type of `fromMIMEHeader695`:
	fromMIMEHeader695 := sourceCQL.(textproto.MIMEHeader)

	// Call the method that transfers the taint
	// from the receiver `fromMIMEHeader695` to the result `intoString862`
	// (`intoString862` is now tainted).
	intoString862 := fromMIMEHeader695.Values("")

	// Return the tainted `intoString862`:
	return intoString862
}

func TaintStepTest_NetTextprotoReaderDotReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader707` into `intoReader609`.

	// Assume that `sourceCQL` has the underlying type of `fromReader707`:
	fromReader707 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader707` to the result `intoReader609`
	// (`intoReader609` is now tainted).
	intoReader609 := fromReader707.DotReader()

	// Return the tainted `intoReader609`:
	return intoReader609
}

func TaintStepTest_NetTextprotoReaderReadCodeLine_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader712` into `intoString440`.

	// Assume that `sourceCQL` has the underlying type of `fromReader712`:
	fromReader712 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader712` to the result `intoString440`
	// (`intoString440` is now tainted).
	_, intoString440, _ := fromReader712.ReadCodeLine(0)

	// Return the tainted `intoString440`:
	return intoString440
}

func TaintStepTest_NetTextprotoReaderReadContinuedLine_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader244` into `intoString808`.

	// Assume that `sourceCQL` has the underlying type of `fromReader244`:
	fromReader244 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader244` to the result `intoString808`
	// (`intoString808` is now tainted).
	intoString808, _ := fromReader244.ReadContinuedLine()

	// Return the tainted `intoString808`:
	return intoString808
}

func TaintStepTest_NetTextprotoReaderReadContinuedLineBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader139` into `intoByte525`.

	// Assume that `sourceCQL` has the underlying type of `fromReader139`:
	fromReader139 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader139` to the result `intoByte525`
	// (`intoByte525` is now tainted).
	intoByte525, _ := fromReader139.ReadContinuedLineBytes()

	// Return the tainted `intoByte525`:
	return intoByte525
}

func TaintStepTest_NetTextprotoReaderReadDotBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader308` into `intoByte397`.

	// Assume that `sourceCQL` has the underlying type of `fromReader308`:
	fromReader308 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader308` to the result `intoByte397`
	// (`intoByte397` is now tainted).
	intoByte397, _ := fromReader308.ReadDotBytes()

	// Return the tainted `intoByte397`:
	return intoByte397
}

func TaintStepTest_NetTextprotoReaderReadDotLines_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader269` into `intoString990`.

	// Assume that `sourceCQL` has the underlying type of `fromReader269`:
	fromReader269 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader269` to the result `intoString990`
	// (`intoString990` is now tainted).
	intoString990, _ := fromReader269.ReadDotLines()

	// Return the tainted `intoString990`:
	return intoString990
}

func TaintStepTest_NetTextprotoReaderReadLine_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader194` into `intoString511`.

	// Assume that `sourceCQL` has the underlying type of `fromReader194`:
	fromReader194 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader194` to the result `intoString511`
	// (`intoString511` is now tainted).
	intoString511, _ := fromReader194.ReadLine()

	// Return the tainted `intoString511`:
	return intoString511
}

func TaintStepTest_NetTextprotoReaderReadLineBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader989` into `intoByte478`.

	// Assume that `sourceCQL` has the underlying type of `fromReader989`:
	fromReader989 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader989` to the result `intoByte478`
	// (`intoByte478` is now tainted).
	intoByte478, _ := fromReader989.ReadLineBytes()

	// Return the tainted `intoByte478`:
	return intoByte478
}

func TaintStepTest_NetTextprotoReaderReadMIMEHeader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader900` into `intoMIMEHeader160`.

	// Assume that `sourceCQL` has the underlying type of `fromReader900`:
	fromReader900 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader900` to the result `intoMIMEHeader160`
	// (`intoMIMEHeader160` is now tainted).
	intoMIMEHeader160, _ := fromReader900.ReadMIMEHeader()

	// Return the tainted `intoMIMEHeader160`:
	return intoMIMEHeader160
}

func TaintStepTest_NetTextprotoReaderReadResponse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader313` into `intoString149`.

	// Assume that `sourceCQL` has the underlying type of `fromReader313`:
	fromReader313 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader313` to the result `intoString149`
	// (`intoString149` is now tainted).
	_, intoString149, _ := fromReader313.ReadResponse(0)

	// Return the tainted `intoString149`:
	return intoString149
}

func TaintStepTest_NetTextprotoWriterDotWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser767` into `intoWriter374`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser767`:
	fromWriteCloser767 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter374` variable:
	var intoWriter374 textproto.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter374`:
	intermediateCQL := intoWriter374.DotWriter()

	// Extra step (`fromWriteCloser767` taints `intermediateCQL`, which taints `intoWriter374`:
	link(fromWriteCloser767, intermediateCQL)

	// Return the tainted `intoWriter374`:
	return intoWriter374
}

func TaintStepTest_NetTextprotoWriterPrintfLine_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString794` into `intoWriter216`.

	// Assume that `sourceCQL` has the underlying type of `fromString794`:
	fromString794 := sourceCQL.(string)

	// Declare `intoWriter216` variable:
	var intoWriter216 textproto.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString794` to the receiver `intoWriter216`
	// (`intoWriter216` is now tainted).
	intoWriter216.PrintfLine(fromString794, nil)

	// Return the tainted `intoWriter216`:
	return intoWriter216
}

func TaintStepTest_NetTextprotoWriterPrintfLine_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface813` into `intoWriter451`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface813`:
	fromInterface813 := sourceCQL.(interface{})

	// Declare `intoWriter451` variable:
	var intoWriter451 textproto.Writer

	// Call the method that transfers the taint
	// from the parameter `fromInterface813` to the receiver `intoWriter451`
	// (`intoWriter451` is now tainted).
	intoWriter451.PrintfLine("", fromInterface813)

	// Return the tainted `intoWriter451`:
	return intoWriter451
}

func RunAllTaints_NetTextproto() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoCanonicalMIMEHeaderKey_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoNewConn_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoNewConn_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoNewReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoNewWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoTrimBytes_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoTrimString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoMIMEHeaderAdd_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoMIMEHeaderAdd_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoMIMEHeaderGet_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoMIMEHeaderSet_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoMIMEHeaderSet_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoMIMEHeaderValues_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoReaderDotReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoReaderReadCodeLine_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoReaderReadContinuedLine_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoReaderReadContinuedLineBytes_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoReaderReadDotBytes_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoReaderReadDotLines_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoReaderReadLine_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoReaderReadLineBytes_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoReaderReadMIMEHeader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoReaderReadResponse_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoWriterDotWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoWriterPrintfLine_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTextprotoWriterPrintfLine_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
