package main

import (
	"bufio"
	"io"
)

func TaintStepTest_BufioNewReadWriter(sourceCQL interface{}) {
	// The flow is from `r` into `into500`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into500`
	// (`into500` is now tainted).
	into500 := bufio.NewReadWriter(r, nil)

	// Sink the tainted `into500`:
	sink(into500)
}

func TaintStepTest_BufioNewReader(sourceCQL interface{}) {
	// The flow is from `rd` into `into659`.

	// Assume that `sourceCQL` has the underlying type of `rd`:
	rd := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `rd` to result `into659`
	// (`into659` is now tainted).
	into659 := bufio.NewReader(rd)

	// Sink the tainted `into659`:
	sink(into659)
}

func TaintStepTest_BufioNewReaderSize(sourceCQL interface{}) {
	// The flow is from `rd` into `into976`.

	// Assume that `sourceCQL` has the underlying type of `rd`:
	rd := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `rd` to result `into976`
	// (`into976` is now tainted).
	into976 := bufio.NewReaderSize(rd, 0)

	// Sink the tainted `into976`:
	sink(into976)
}

func TaintStepTest_BufioNewScanner(sourceCQL interface{}) {
	// The flow is from `r` into `into315`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into315`
	// (`into315` is now tainted).
	into315 := bufio.NewScanner(r)

	// Sink the tainted `into315`:
	sink(into315)
}

func TaintStepTest_BufioNewWriter(sourceCQL interface{}) {
	// The flow is from `from792` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from792`:
	from792 := sourceCQL.(*bufio.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := bufio.NewWriter(w)

	// Extra step (`from792` taints `intermediateCQL`, which taints `w`:
	link(from792, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_BufioNewWriterSize(sourceCQL interface{}) {
	// The flow is from `from444` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from444`:
	from444 := sourceCQL.(*bufio.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := bufio.NewWriterSize(w, 0)

	// Extra step (`from444` taints `intermediateCQL`, which taints `w`:
	link(from444, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_BufioScanBytes(sourceCQL interface{}) {
	// The flow is from `data` into `token`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `data` to result `token`
	// (`token` is now tainted).
	_, token, _ := bufio.ScanBytes(data, false)

	// Sink the tainted `token`:
	sink(token)
}

func TaintStepTest_BufioScanLines(sourceCQL interface{}) {
	// The flow is from `data` into `token`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `data` to result `token`
	// (`token` is now tainted).
	_, token, _ := bufio.ScanLines(data, false)

	// Sink the tainted `token`:
	sink(token)
}

func TaintStepTest_BufioScanRunes(sourceCQL interface{}) {
	// The flow is from `data` into `token`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `data` to result `token`
	// (`token` is now tainted).
	_, token, _ := bufio.ScanRunes(data, false)

	// Sink the tainted `token`:
	sink(token)
}

func TaintStepTest_BufioScanWords(sourceCQL interface{}) {
	// The flow is from `data` into `token`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `data` to result `token`
	// (`token` is now tainted).
	_, token, _ := bufio.ScanWords(data, false)

	// Sink the tainted `token`:
	sink(token)
}

func TaintStepTest_BufioReaderPeek(sourceCQL interface{}) {
	// The flow is from `from971` into `into277`.

	// Assume that `sourceCQL` has the underlying type of `from971`:
	from971 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `from971` to the result `into277`
	// (`into277` is now tainted).
	into277, _ := from971.Peek(0)

	// Sink the tainted `into277`:
	sink(into277)
}

func TaintStepTest_BufioReaderRead(sourceCQL interface{}) {
	// The flow is from `from563` into `p`.

	// Assume that `sourceCQL` has the underlying type of `from563`:
	from563 := sourceCQL.(bufio.Reader)

	// Declare `p` variable:
	var p []byte

	// Call the method that transfers the taint
	// from the receiver `from563` to the argument `p`
	// (`p` is now tainted).
	from563.Read(p)

	// Sink the tainted `p`:
	sink(p)
}

func TaintStepTest_BufioReaderReadByte(sourceCQL interface{}) {
	// The flow is from `from970` into `into533`.

	// Assume that `sourceCQL` has the underlying type of `from970`:
	from970 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `from970` to the result `into533`
	// (`into533` is now tainted).
	into533, _ := from970.ReadByte()

	// Sink the tainted `into533`:
	sink(into533)
}

func TaintStepTest_BufioReaderReadBytes(sourceCQL interface{}) {
	// The flow is from `from694` into `into872`.

	// Assume that `sourceCQL` has the underlying type of `from694`:
	from694 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `from694` to the result `into872`
	// (`into872` is now tainted).
	into872, _ := from694.ReadBytes(0)

	// Sink the tainted `into872`:
	sink(into872)
}

func TaintStepTest_BufioReaderReadLine(sourceCQL interface{}) {
	// The flow is from `from768` into `line`.

	// Assume that `sourceCQL` has the underlying type of `from768`:
	from768 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `from768` to the result `line`
	// (`line` is now tainted).
	line, _, _ := from768.ReadLine()

	// Sink the tainted `line`:
	sink(line)
}

func TaintStepTest_BufioReaderReadRune(sourceCQL interface{}) {
	// The flow is from `from715` into `r`.

	// Assume that `sourceCQL` has the underlying type of `from715`:
	from715 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `from715` to the result `r`
	// (`r` is now tainted).
	r, _, _ := from715.ReadRune()

	// Sink the tainted `r`:
	sink(r)
}

func TaintStepTest_BufioReaderReadSlice(sourceCQL interface{}) {
	// The flow is from `from335` into `line`.

	// Assume that `sourceCQL` has the underlying type of `from335`:
	from335 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `from335` to the result `line`
	// (`line` is now tainted).
	line, _ := from335.ReadSlice(0)

	// Sink the tainted `line`:
	sink(line)
}

func TaintStepTest_BufioReaderReadString(sourceCQL interface{}) {
	// The flow is from `from658` into `into387`.

	// Assume that `sourceCQL` has the underlying type of `from658`:
	from658 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `from658` to the result `into387`
	// (`into387` is now tainted).
	into387, _ := from658.ReadString(0)

	// Sink the tainted `into387`:
	sink(into387)
}

func TaintStepTest_BufioReaderReset(sourceCQL interface{}) {
	// The flow is from `r` into `into570`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Declare `into570` variable:
	var into570 bufio.Reader

	// Call the method that transfers the taint
	// from the parameter `r` to the receiver `into570`
	// (`into570` is now tainted).
	into570.Reset(r)

	// Sink the tainted `into570`:
	sink(into570)
}

func TaintStepTest_BufioReaderWriteTo(sourceCQL interface{}) {
	// The flow is from `from803` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from803`:
	from803 := sourceCQL.(bufio.Reader)

	// Declare `w` variable:
	var w io.Writer

	// Call the method that transfers the taint
	// from the receiver `from803` to the argument `w`
	// (`w` is now tainted).
	from803.WriteTo(w)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_BufioScannerBytes(sourceCQL interface{}) {
	// The flow is from `from143` into `into643`.

	// Assume that `sourceCQL` has the underlying type of `from143`:
	from143 := sourceCQL.(bufio.Scanner)

	// Call the method that transfers the taint
	// from the receiver `from143` to the result `into643`
	// (`into643` is now tainted).
	into643 := from143.Bytes()

	// Sink the tainted `into643`:
	sink(into643)
}

func TaintStepTest_BufioScannerText(sourceCQL interface{}) {
	// The flow is from `from423` into `into429`.

	// Assume that `sourceCQL` has the underlying type of `from423`:
	from423 := sourceCQL.(bufio.Scanner)

	// Call the method that transfers the taint
	// from the receiver `from423` to the result `into429`
	// (`into429` is now tainted).
	into429 := from423.Text()

	// Sink the tainted `into429`:
	sink(into429)
}

func TaintStepTest_BufioWriterReadFrom(sourceCQL interface{}) {
	// The flow is from `r` into `into221`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Declare `into221` variable:
	var into221 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `r` to the receiver `into221`
	// (`into221` is now tainted).
	into221.ReadFrom(r)

	// Sink the tainted `into221`:
	sink(into221)
}

func TaintStepTest_BufioWriterReset(sourceCQL interface{}) {
	// The flow is from `from998` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from998`:
	from998 := sourceCQL.(bufio.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the method that transfers the taint
	// from the receiver `from998` to the argument `w`
	// (`w` is now tainted).
	from998.Reset(w)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_BufioWriterWrite(sourceCQL interface{}) {
	// The flow is from `p` into `into406`.

	// Assume that `sourceCQL` has the underlying type of `p`:
	p := sourceCQL.([]byte)

	// Declare `into406` variable:
	var into406 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `p` to the receiver `into406`
	// (`into406` is now tainted).
	into406.Write(p)

	// Sink the tainted `into406`:
	sink(into406)
}

func TaintStepTest_BufioWriterWriteByte(sourceCQL interface{}) {
	// The flow is from `c` into `into973`.

	// Assume that `sourceCQL` has the underlying type of `c`:
	c := sourceCQL.(byte)

	// Declare `into973` variable:
	var into973 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `c` to the receiver `into973`
	// (`into973` is now tainted).
	into973.WriteByte(c)

	// Sink the tainted `into973`:
	sink(into973)
}

func TaintStepTest_BufioWriterWriteRune(sourceCQL interface{}) {
	// The flow is from `r` into `into566`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(rune)

	// Declare `into566` variable:
	var into566 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `r` to the receiver `into566`
	// (`into566` is now tainted).
	into566.WriteRune(r)

	// Sink the tainted `into566`:
	sink(into566)
}

func TaintStepTest_BufioWriterWriteString(sourceCQL interface{}) {
	// The flow is from `s` into `into512`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare `into512` variable:
	var into512 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `s` to the receiver `into512`
	// (`into512` is now tainted).
	into512.WriteString(s)

	// Sink the tainted `into512`:
	sink(into512)
}

func RunAllTaints_Bufio(v interface{}) {
	{
		source := newSource()
		TaintStepTest_BufioNewReadWriter(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioNewReader(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioNewReaderSize(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioNewScanner(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioNewWriter(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioNewWriterSize(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioScanBytes(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioScanLines(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioScanRunes(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioScanWords(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderPeek(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderRead(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderReadByte(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderReadBytes(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderReadLine(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderReadRune(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderReadSlice(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderReadString(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderReset(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderWriteTo(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioScannerBytes(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioScannerText(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioWriterReadFrom(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioWriterReset(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioWriterWrite(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioWriterWriteByte(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioWriterWriteRune(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioWriterWriteString(source)
	}
}
