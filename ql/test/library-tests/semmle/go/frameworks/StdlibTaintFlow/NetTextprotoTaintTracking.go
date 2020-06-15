package main

import (
	"bufio"
	"io"
	"net/textproto"
)

func TaintStepTest_NetTextprotoCanonicalMIMEHeaderKey_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString332` into `intoString455`.

	// Assume that `sourceCQL` has the underlying type of `fromString332`:
	fromString332 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString332` to result `intoString455`
	// (`intoString455` is now tainted).
	intoString455 := textproto.CanonicalMIMEHeaderKey(fromString332)

	// Sink the tainted `intoString455`:
	sink(intoString455)
}

func TaintStepTest_NetTextprotoNewConn_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReadWriteCloser775` into `intoConn273`.

	// Assume that `sourceCQL` has the underlying type of `fromReadWriteCloser775`:
	fromReadWriteCloser775 := sourceCQL.(io.ReadWriteCloser)

	// Call the function that transfers the taint
	// from the parameter `fromReadWriteCloser775` to result `intoConn273`
	// (`intoConn273` is now tainted).
	intoConn273 := textproto.NewConn(fromReadWriteCloser775)

	// Sink the tainted `intoConn273`:
	sink(intoConn273)
}

func TaintStepTest_NetTextprotoNewConn_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromConn523` into `intoReadWriteCloser338`.

	// Assume that `sourceCQL` has the underlying type of `fromConn523`:
	fromConn523 := sourceCQL.(*textproto.Conn)

	// Declare `intoReadWriteCloser338` variable:
	var intoReadWriteCloser338 io.ReadWriteCloser

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoReadWriteCloser338`:
	intermediateCQL := textproto.NewConn(intoReadWriteCloser338)

	// Extra step (`fromConn523` taints `intermediateCQL`, which taints `intoReadWriteCloser338`:
	link(fromConn523, intermediateCQL)

	// Sink the tainted `intoReadWriteCloser338`:
	sink(intoReadWriteCloser338)
}

func TaintStepTest_NetTextprotoNewReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader246` into `intoReader257`.

	// Assume that `sourceCQL` has the underlying type of `fromReader246`:
	fromReader246 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader246` to result `intoReader257`
	// (`intoReader257` is now tainted).
	intoReader257 := textproto.NewReader(fromReader246)

	// Sink the tainted `intoReader257`:
	sink(intoReader257)
}

func TaintStepTest_NetTextprotoNewWriter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter392` into `intoWriter234`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter392`:
	fromWriter392 := sourceCQL.(*textproto.Writer)

	// Declare `intoWriter234` variable:
	var intoWriter234 *bufio.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter234`:
	intermediateCQL := textproto.NewWriter(intoWriter234)

	// Extra step (`fromWriter392` taints `intermediateCQL`, which taints `intoWriter234`:
	link(fromWriter392, intermediateCQL)

	// Sink the tainted `intoWriter234`:
	sink(intoWriter234)
}

func TaintStepTest_NetTextprotoTrimBytes_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte335` into `intoByte187`.

	// Assume that `sourceCQL` has the underlying type of `fromByte335`:
	fromByte335 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte335` to result `intoByte187`
	// (`intoByte187` is now tainted).
	intoByte187 := textproto.TrimBytes(fromByte335)

	// Sink the tainted `intoByte187`:
	sink(intoByte187)
}

func TaintStepTest_NetTextprotoTrimString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString120` into `intoString474`.

	// Assume that `sourceCQL` has the underlying type of `fromString120`:
	fromString120 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString120` to result `intoString474`
	// (`intoString474` is now tainted).
	intoString474 := textproto.TrimString(fromString120)

	// Sink the tainted `intoString474`:
	sink(intoString474)
}

func TaintStepTest_NetTextprotoMIMEHeaderAdd_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString245` into `intoMIMEHeader162`.

	// Assume that `sourceCQL` has the underlying type of `fromString245`:
	fromString245 := sourceCQL.(string)

	// Declare `intoMIMEHeader162` variable:
	var intoMIMEHeader162 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `fromString245` to the receiver `intoMIMEHeader162`
	// (`intoMIMEHeader162` is now tainted).
	intoMIMEHeader162.Add(fromString245, "")

	// Sink the tainted `intoMIMEHeader162`:
	sink(intoMIMEHeader162)
}

func TaintStepTest_NetTextprotoMIMEHeaderAdd_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString968` into `intoMIMEHeader128`.

	// Assume that `sourceCQL` has the underlying type of `fromString968`:
	fromString968 := sourceCQL.(string)

	// Declare `intoMIMEHeader128` variable:
	var intoMIMEHeader128 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `fromString968` to the receiver `intoMIMEHeader128`
	// (`intoMIMEHeader128` is now tainted).
	intoMIMEHeader128.Add("", fromString968)

	// Sink the tainted `intoMIMEHeader128`:
	sink(intoMIMEHeader128)
}

func TaintStepTest_NetTextprotoMIMEHeaderGet_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromMIMEHeader217` into `intoString430`.

	// Assume that `sourceCQL` has the underlying type of `fromMIMEHeader217`:
	fromMIMEHeader217 := sourceCQL.(textproto.MIMEHeader)

	// Call the method that transfers the taint
	// from the receiver `fromMIMEHeader217` to the result `intoString430`
	// (`intoString430` is now tainted).
	intoString430 := fromMIMEHeader217.Get("")

	// Sink the tainted `intoString430`:
	sink(intoString430)
}

func TaintStepTest_NetTextprotoMIMEHeaderSet_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString268` into `intoMIMEHeader400`.

	// Assume that `sourceCQL` has the underlying type of `fromString268`:
	fromString268 := sourceCQL.(string)

	// Declare `intoMIMEHeader400` variable:
	var intoMIMEHeader400 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `fromString268` to the receiver `intoMIMEHeader400`
	// (`intoMIMEHeader400` is now tainted).
	intoMIMEHeader400.Set(fromString268, "")

	// Sink the tainted `intoMIMEHeader400`:
	sink(intoMIMEHeader400)
}

func TaintStepTest_NetTextprotoMIMEHeaderSet_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString645` into `intoMIMEHeader223`.

	// Assume that `sourceCQL` has the underlying type of `fromString645`:
	fromString645 := sourceCQL.(string)

	// Declare `intoMIMEHeader223` variable:
	var intoMIMEHeader223 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `fromString645` to the receiver `intoMIMEHeader223`
	// (`intoMIMEHeader223` is now tainted).
	intoMIMEHeader223.Set("", fromString645)

	// Sink the tainted `intoMIMEHeader223`:
	sink(intoMIMEHeader223)
}

func TaintStepTest_NetTextprotoMIMEHeaderValues_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromMIMEHeader870` into `intoString814`.

	// Assume that `sourceCQL` has the underlying type of `fromMIMEHeader870`:
	fromMIMEHeader870 := sourceCQL.(textproto.MIMEHeader)

	// Call the method that transfers the taint
	// from the receiver `fromMIMEHeader870` to the result `intoString814`
	// (`intoString814` is now tainted).
	intoString814 := fromMIMEHeader870.Values("")

	// Sink the tainted `intoString814`:
	sink(intoString814)
}

func TaintStepTest_NetTextprotoReaderDotReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader885` into `intoReader436`.

	// Assume that `sourceCQL` has the underlying type of `fromReader885`:
	fromReader885 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader885` to the result `intoReader436`
	// (`intoReader436` is now tainted).
	intoReader436 := fromReader885.DotReader()

	// Sink the tainted `intoReader436`:
	sink(intoReader436)
}

func TaintStepTest_NetTextprotoReaderReadCodeLine_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader705` into `intoString574`.

	// Assume that `sourceCQL` has the underlying type of `fromReader705`:
	fromReader705 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader705` to the result `intoString574`
	// (`intoString574` is now tainted).
	_, intoString574, _ := fromReader705.ReadCodeLine(0)

	// Sink the tainted `intoString574`:
	sink(intoString574)
}

func TaintStepTest_NetTextprotoReaderReadContinuedLine_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader728` into `intoString158`.

	// Assume that `sourceCQL` has the underlying type of `fromReader728`:
	fromReader728 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader728` to the result `intoString158`
	// (`intoString158` is now tainted).
	intoString158, _ := fromReader728.ReadContinuedLine()

	// Sink the tainted `intoString158`:
	sink(intoString158)
}

func TaintStepTest_NetTextprotoReaderReadContinuedLineBytes_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader889` into `intoByte921`.

	// Assume that `sourceCQL` has the underlying type of `fromReader889`:
	fromReader889 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader889` to the result `intoByte921`
	// (`intoByte921` is now tainted).
	intoByte921, _ := fromReader889.ReadContinuedLineBytes()

	// Sink the tainted `intoByte921`:
	sink(intoByte921)
}

func TaintStepTest_NetTextprotoReaderReadDotBytes_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader875` into `intoByte879`.

	// Assume that `sourceCQL` has the underlying type of `fromReader875`:
	fromReader875 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader875` to the result `intoByte879`
	// (`intoByte879` is now tainted).
	intoByte879, _ := fromReader875.ReadDotBytes()

	// Sink the tainted `intoByte879`:
	sink(intoByte879)
}

func TaintStepTest_NetTextprotoReaderReadDotLines_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader832` into `intoString356`.

	// Assume that `sourceCQL` has the underlying type of `fromReader832`:
	fromReader832 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader832` to the result `intoString356`
	// (`intoString356` is now tainted).
	intoString356, _ := fromReader832.ReadDotLines()

	// Sink the tainted `intoString356`:
	sink(intoString356)
}

func TaintStepTest_NetTextprotoReaderReadLine_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader232` into `intoString484`.

	// Assume that `sourceCQL` has the underlying type of `fromReader232`:
	fromReader232 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader232` to the result `intoString484`
	// (`intoString484` is now tainted).
	intoString484, _ := fromReader232.ReadLine()

	// Sink the tainted `intoString484`:
	sink(intoString484)
}

func TaintStepTest_NetTextprotoReaderReadLineBytes_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader663` into `intoByte677`.

	// Assume that `sourceCQL` has the underlying type of `fromReader663`:
	fromReader663 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader663` to the result `intoByte677`
	// (`intoByte677` is now tainted).
	intoByte677, _ := fromReader663.ReadLineBytes()

	// Sink the tainted `intoByte677`:
	sink(intoByte677)
}

func TaintStepTest_NetTextprotoReaderReadMIMEHeader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader856` into `intoMIMEHeader917`.

	// Assume that `sourceCQL` has the underlying type of `fromReader856`:
	fromReader856 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader856` to the result `intoMIMEHeader917`
	// (`intoMIMEHeader917` is now tainted).
	intoMIMEHeader917, _ := fromReader856.ReadMIMEHeader()

	// Sink the tainted `intoMIMEHeader917`:
	sink(intoMIMEHeader917)
}

func TaintStepTest_NetTextprotoReaderReadResponse_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader316` into `intoString919`.

	// Assume that `sourceCQL` has the underlying type of `fromReader316`:
	fromReader316 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader316` to the result `intoString919`
	// (`intoString919` is now tainted).
	_, intoString919, _ := fromReader316.ReadResponse(0)

	// Sink the tainted `intoString919`:
	sink(intoString919)
}

func TaintStepTest_NetTextprotoWriterDotWriter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriteCloser493` into `intoWriter738`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser493`:
	fromWriteCloser493 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter738` variable:
	var intoWriter738 textproto.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter738`:
	intermediateCQL := intoWriter738.DotWriter()

	// Extra step (`fromWriteCloser493` taints `intermediateCQL`, which taints `intoWriter738`:
	link(fromWriteCloser493, intermediateCQL)

	// Sink the tainted `intoWriter738`:
	sink(intoWriter738)
}

func TaintStepTest_NetTextprotoWriterPrintfLine_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString750` into `intoWriter272`.

	// Assume that `sourceCQL` has the underlying type of `fromString750`:
	fromString750 := sourceCQL.(string)

	// Declare `intoWriter272` variable:
	var intoWriter272 textproto.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString750` to the receiver `intoWriter272`
	// (`intoWriter272` is now tainted).
	intoWriter272.PrintfLine(fromString750, nil)

	// Sink the tainted `intoWriter272`:
	sink(intoWriter272)
}

func TaintStepTest_NetTextprotoWriterPrintfLine_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromInterface795` into `intoWriter114`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface795`:
	fromInterface795 := sourceCQL.([]interface{})

	// Declare `intoWriter114` variable:
	var intoWriter114 textproto.Writer

	// Call the method that transfers the taint
	// from the parameter `fromInterface795` to the receiver `intoWriter114`
	// (`intoWriter114` is now tainted).
	intoWriter114.PrintfLine("", fromInterface795)

	// Sink the tainted `intoWriter114`:
	sink(intoWriter114)
}

func RunAllTaints_NetTextproto() {
	{
		source := newSource()
		TaintStepTest_NetTextprotoCanonicalMIMEHeaderKey_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoNewConn_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoNewConn_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoNewReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoNewWriter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoTrimBytes_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoTrimString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoMIMEHeaderAdd_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoMIMEHeaderAdd_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoMIMEHeaderGet_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoMIMEHeaderSet_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoMIMEHeaderSet_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoMIMEHeaderValues_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderDotReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadCodeLine_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadContinuedLine_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadContinuedLineBytes_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadDotBytes_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadDotLines_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadLine_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadLineBytes_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadMIMEHeader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadResponse_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoWriterDotWriter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoWriterPrintfLine_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoWriterPrintfLine_B0I1O0(source)
	}
}
