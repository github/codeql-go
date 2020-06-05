package main

import (
	"bufio"
	"io"
	"net/textproto"
)

func TaintStepTest_NetTextprotoCanonicalMIMEHeaderKey(sourceCQL interface{}) {
	// The flow is from `s` into `into489`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into489`
	// (`into489` is now tainted).
	into489 := textproto.CanonicalMIMEHeaderKey(s)

	// Sink the tainted `into489`:
	sink(into489)
}

func TaintStepTest_NetTextprotoNewReader(sourceCQL interface{}) {
	// The flow is from `r` into `into836`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into836`
	// (`into836` is now tainted).
	into836 := textproto.NewReader(r)

	// Sink the tainted `into836`:
	sink(into836)
}

func TaintStepTest_NetTextprotoNewWriter(sourceCQL interface{}) {
	// The flow is from `from912` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from912`:
	from912 := sourceCQL.(*textproto.Writer)

	// Declare `w` variable:
	var w bufio.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := textproto.NewWriter(&w)

	// Extra step (`from912` taints `intermediateCQL`, which taints `w`:
	link(from912, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_NetTextprotoTrimBytes(sourceCQL interface{}) {
	// The flow is from `b` into `into329`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `b` to result `into329`
	// (`into329` is now tainted).
	into329 := textproto.TrimBytes(b)

	// Sink the tainted `into329`:
	sink(into329)
}

func TaintStepTest_NetTextprotoTrimString(sourceCQL interface{}) {
	// The flow is from `s` into `into760`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into760`
	// (`into760` is now tainted).
	into760 := textproto.TrimString(s)

	// Sink the tainted `into760`:
	sink(into760)
}

func TaintStepTest_NetTextprotoMIMEHeaderAdd(sourceCQL interface{}) {
	// The flow is from `value` into `into197`.

	// Assume that `sourceCQL` has the underlying type of `value`:
	value := sourceCQL.(string)

	// Declare `into197` variable:
	var into197 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `value` to the receiver `into197`
	// (`into197` is now tainted).
	into197.Add("", value)

	// Sink the tainted `into197`:
	sink(into197)
}

func TaintStepTest_NetTextprotoMIMEHeaderGet(sourceCQL interface{}) {
	// The flow is from `from519` into `into295`.

	// Assume that `sourceCQL` has the underlying type of `from519`:
	from519 := sourceCQL.(textproto.MIMEHeader)

	// Call the method that transfers the taint
	// from the receiver `from519` to the result `into295`
	// (`into295` is now tainted).
	into295 := from519.Get("")

	// Sink the tainted `into295`:
	sink(into295)
}

func TaintStepTest_NetTextprotoMIMEHeaderSet(sourceCQL interface{}) {
	// The flow is from `value` into `into772`.

	// Assume that `sourceCQL` has the underlying type of `value`:
	value := sourceCQL.(string)

	// Declare `into772` variable:
	var into772 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `value` to the receiver `into772`
	// (`into772` is now tainted).
	into772.Set("", value)

	// Sink the tainted `into772`:
	sink(into772)
}

func TaintStepTest_NetTextprotoMIMEHeaderValues(sourceCQL interface{}) {
	// The flow is from `from283` into `into339`.

	// Assume that `sourceCQL` has the underlying type of `from283`:
	from283 := sourceCQL.(textproto.MIMEHeader)

	// Call the method that transfers the taint
	// from the receiver `from283` to the result `into339`
	// (`into339` is now tainted).
	into339 := from283.Values("")

	// Sink the tainted `into339`:
	sink(into339)
}

func TaintStepTest_NetTextprotoReaderDotReader(sourceCQL interface{}) {
	// The flow is from `from316` into `into282`.

	// Assume that `sourceCQL` has the underlying type of `from316`:
	from316 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `from316` to the result `into282`
	// (`into282` is now tainted).
	into282 := from316.DotReader()

	// Sink the tainted `into282`:
	sink(into282)
}

func TaintStepTest_NetTextprotoReaderReadCodeLine(sourceCQL interface{}) {
	// The flow is from `from657` into `message`.

	// Assume that `sourceCQL` has the underlying type of `from657`:
	from657 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `from657` to the result `message`
	// (`message` is now tainted).
	_, message, _ := from657.ReadCodeLine(0)

	// Sink the tainted `message`:
	sink(message)
}

func TaintStepTest_NetTextprotoReaderReadContinuedLine(sourceCQL interface{}) {
	// The flow is from `from537` into `into802`.

	// Assume that `sourceCQL` has the underlying type of `from537`:
	from537 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `from537` to the result `into802`
	// (`into802` is now tainted).
	into802, _ := from537.ReadContinuedLine()

	// Sink the tainted `into802`:
	sink(into802)
}

func TaintStepTest_NetTextprotoReaderReadContinuedLineBytes(sourceCQL interface{}) {
	// The flow is from `from495` into `into277`.

	// Assume that `sourceCQL` has the underlying type of `from495`:
	from495 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `from495` to the result `into277`
	// (`into277` is now tainted).
	into277, _ := from495.ReadContinuedLineBytes()

	// Sink the tainted `into277`:
	sink(into277)
}

func TaintStepTest_NetTextprotoReaderReadDotBytes(sourceCQL interface{}) {
	// The flow is from `from499` into `into289`.

	// Assume that `sourceCQL` has the underlying type of `from499`:
	from499 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `from499` to the result `into289`
	// (`into289` is now tainted).
	into289, _ := from499.ReadDotBytes()

	// Sink the tainted `into289`:
	sink(into289)
}

func TaintStepTest_NetTextprotoReaderReadDotLines(sourceCQL interface{}) {
	// The flow is from `from877` into `into580`.

	// Assume that `sourceCQL` has the underlying type of `from877`:
	from877 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `from877` to the result `into580`
	// (`into580` is now tainted).
	into580, _ := from877.ReadDotLines()

	// Sink the tainted `into580`:
	sink(into580)
}

func TaintStepTest_NetTextprotoReaderReadLine(sourceCQL interface{}) {
	// The flow is from `from742` into `into474`.

	// Assume that `sourceCQL` has the underlying type of `from742`:
	from742 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `from742` to the result `into474`
	// (`into474` is now tainted).
	into474, _ := from742.ReadLine()

	// Sink the tainted `into474`:
	sink(into474)
}

func TaintStepTest_NetTextprotoReaderReadLineBytes(sourceCQL interface{}) {
	// The flow is from `from599` into `into697`.

	// Assume that `sourceCQL` has the underlying type of `from599`:
	from599 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `from599` to the result `into697`
	// (`into697` is now tainted).
	into697, _ := from599.ReadLineBytes()

	// Sink the tainted `into697`:
	sink(into697)
}

func TaintStepTest_NetTextprotoReaderReadMIMEHeader(sourceCQL interface{}) {
	// The flow is from `from211` into `into821`.

	// Assume that `sourceCQL` has the underlying type of `from211`:
	from211 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `from211` to the result `into821`
	// (`into821` is now tainted).
	into821, _ := from211.ReadMIMEHeader()

	// Sink the tainted `into821`:
	sink(into821)
}

func TaintStepTest_NetTextprotoReaderReadResponse(sourceCQL interface{}) {
	// The flow is from `from795` into `message`.

	// Assume that `sourceCQL` has the underlying type of `from795`:
	from795 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `from795` to the result `message`
	// (`message` is now tainted).
	_, message, _ := from795.ReadResponse(0)

	// Sink the tainted `message`:
	sink(message)
}

func TaintStepTest_NetTextprotoWriterDotWriter(sourceCQL interface{}) {
	// The flow is from `from188` into `into199`.

	// Assume that `sourceCQL` has the underlying type of `from188`:
	from188 := sourceCQL.(io.WriteCloser)

	// Declare `into199` variable:
	var into199 textproto.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `into199`:
	intermediateCQL := into199.DotWriter()

	// Extra step (`from188` taints `intermediateCQL`, which taints `into199`:
	link(from188, intermediateCQL)

	// Sink the tainted `into199`:
	sink(into199)
}

func TaintStepTest_NetTextprotoWriterPrintfLine(sourceCQL interface{}) {
	// The flow is from `format` into `into566`.

	// Assume that `sourceCQL` has the underlying type of `format`:
	format := sourceCQL.(string)

	// Declare `into566` variable:
	var into566 textproto.Writer

	// Call the method that transfers the taint
	// from the parameter `format` to the receiver `into566`
	// (`into566` is now tainted).
	into566.PrintfLine(format, nil)

	// Sink the tainted `into566`:
	sink(into566)
}

func RunAllTaints_NetTextproto(v interface{}) {
	{
		source := newSource()
		TaintStepTest_NetTextprotoCanonicalMIMEHeaderKey(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoNewReader(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoNewWriter(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoTrimBytes(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoTrimString(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoMIMEHeaderAdd(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoMIMEHeaderGet(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoMIMEHeaderSet(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoMIMEHeaderValues(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderDotReader(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadCodeLine(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadContinuedLine(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadContinuedLineBytes(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadDotBytes(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadDotLines(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadLine(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadLineBytes(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadMIMEHeader(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoReaderReadResponse(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoWriterDotWriter(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTextprotoWriterPrintfLine(source)
	}
}
