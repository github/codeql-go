package main

import (
	"bufio"
	"io"
)

func TaintStepTest_BufioNewReadWriter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader776` into `intoReadWriter377`.

	// Assume that `sourceCQL` has the underlying type of `fromReader776`:
	fromReader776 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader776` to result `intoReadWriter377`
	// (`intoReadWriter377` is now tainted).
	intoReadWriter377 := bufio.NewReadWriter(fromReader776, nil)

	// Sink the tainted `intoReadWriter377`:
	sink(intoReadWriter377)
}

func TaintStepTest_BufioNewReadWriter_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromReadWriter425` into `intoWriter207`.

	// Assume that `sourceCQL` has the underlying type of `fromReadWriter425`:
	fromReadWriter425 := sourceCQL.(*bufio.ReadWriter)

	// Declare `intoWriter207` variable:
	var intoWriter207 *bufio.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter207`:
	intermediateCQL := bufio.NewReadWriter(nil, intoWriter207)

	// Extra step (`fromReadWriter425` taints `intermediateCQL`, which taints `intoWriter207`:
	link(fromReadWriter425, intermediateCQL)

	// Sink the tainted `intoWriter207`:
	sink(intoWriter207)
}

func TaintStepTest_BufioNewReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader174` into `intoReader616`.

	// Assume that `sourceCQL` has the underlying type of `fromReader174`:
	fromReader174 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader174` to result `intoReader616`
	// (`intoReader616` is now tainted).
	intoReader616 := bufio.NewReader(fromReader174)

	// Sink the tainted `intoReader616`:
	sink(intoReader616)
}

func TaintStepTest_BufioNewReaderSize_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader325` into `intoReader751`.

	// Assume that `sourceCQL` has the underlying type of `fromReader325`:
	fromReader325 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader325` to result `intoReader751`
	// (`intoReader751` is now tainted).
	intoReader751 := bufio.NewReaderSize(fromReader325, 0)

	// Sink the tainted `intoReader751`:
	sink(intoReader751)
}

func TaintStepTest_BufioNewScanner_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader392` into `intoScanner290`.

	// Assume that `sourceCQL` has the underlying type of `fromReader392`:
	fromReader392 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader392` to result `intoScanner290`
	// (`intoScanner290` is now tainted).
	intoScanner290 := bufio.NewScanner(fromReader392)

	// Sink the tainted `intoScanner290`:
	sink(intoScanner290)
}

func TaintStepTest_BufioNewWriter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter900` into `intoWriter599`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter900`:
	fromWriter900 := sourceCQL.(*bufio.Writer)

	// Declare `intoWriter599` variable:
	var intoWriter599 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter599`:
	intermediateCQL := bufio.NewWriter(intoWriter599)

	// Extra step (`fromWriter900` taints `intermediateCQL`, which taints `intoWriter599`:
	link(fromWriter900, intermediateCQL)

	// Sink the tainted `intoWriter599`:
	sink(intoWriter599)
}

func TaintStepTest_BufioNewWriterSize_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter728` into `intoWriter972`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter728`:
	fromWriter728 := sourceCQL.(*bufio.Writer)

	// Declare `intoWriter972` variable:
	var intoWriter972 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter972`:
	intermediateCQL := bufio.NewWriterSize(intoWriter972, 0)

	// Extra step (`fromWriter728` taints `intermediateCQL`, which taints `intoWriter972`:
	link(fromWriter728, intermediateCQL)

	// Sink the tainted `intoWriter972`:
	sink(intoWriter972)
}

func TaintStepTest_BufioScanBytes_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte529` into `intoByte927`.

	// Assume that `sourceCQL` has the underlying type of `fromByte529`:
	fromByte529 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte529` to result `intoByte927`
	// (`intoByte927` is now tainted).
	_, intoByte927, _ := bufio.ScanBytes(fromByte529, false)

	// Sink the tainted `intoByte927`:
	sink(intoByte927)
}

func TaintStepTest_BufioScanLines_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte260` into `intoByte561`.

	// Assume that `sourceCQL` has the underlying type of `fromByte260`:
	fromByte260 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte260` to result `intoByte561`
	// (`intoByte561` is now tainted).
	_, intoByte561, _ := bufio.ScanLines(fromByte260, false)

	// Sink the tainted `intoByte561`:
	sink(intoByte561)
}

func TaintStepTest_BufioScanRunes_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte780` into `intoByte794`.

	// Assume that `sourceCQL` has the underlying type of `fromByte780`:
	fromByte780 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte780` to result `intoByte794`
	// (`intoByte794` is now tainted).
	_, intoByte794, _ := bufio.ScanRunes(fromByte780, false)

	// Sink the tainted `intoByte794`:
	sink(intoByte794)
}

func TaintStepTest_BufioScanWords_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte245` into `intoByte377`.

	// Assume that `sourceCQL` has the underlying type of `fromByte245`:
	fromByte245 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte245` to result `intoByte377`
	// (`intoByte377` is now tainted).
	_, intoByte377, _ := bufio.ScanWords(fromByte245, false)

	// Sink the tainted `intoByte377`:
	sink(intoByte377)
}

func TaintStepTest_BufioReaderPeek_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader515` into `intoByte563`.

	// Assume that `sourceCQL` has the underlying type of `fromReader515`:
	fromReader515 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader515` to the result `intoByte563`
	// (`intoByte563` is now tainted).
	intoByte563, _ := fromReader515.Peek(0)

	// Sink the tainted `intoByte563`:
	sink(intoByte563)
}

func TaintStepTest_BufioReaderRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader792` into `intoByte500`.

	// Assume that `sourceCQL` has the underlying type of `fromReader792`:
	fromReader792 := sourceCQL.(bufio.Reader)

	// Declare `intoByte500` variable:
	var intoByte500 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader792` to the argument `intoByte500`
	// (`intoByte500` is now tainted).
	fromReader792.Read(intoByte500)

	// Sink the tainted `intoByte500`:
	sink(intoByte500)
}

func TaintStepTest_BufioReaderReadByte_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader198` into `intoByte480`.

	// Assume that `sourceCQL` has the underlying type of `fromReader198`:
	fromReader198 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader198` to the result `intoByte480`
	// (`intoByte480` is now tainted).
	intoByte480, _ := fromReader198.ReadByte()

	// Sink the tainted `intoByte480`:
	sink(intoByte480)
}

func TaintStepTest_BufioReaderReadBytes_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader683` into `intoByte613`.

	// Assume that `sourceCQL` has the underlying type of `fromReader683`:
	fromReader683 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader683` to the result `intoByte613`
	// (`intoByte613` is now tainted).
	intoByte613, _ := fromReader683.ReadBytes(0)

	// Sink the tainted `intoByte613`:
	sink(intoByte613)
}

func TaintStepTest_BufioReaderReadLine_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader655` into `intoByte200`.

	// Assume that `sourceCQL` has the underlying type of `fromReader655`:
	fromReader655 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader655` to the result `intoByte200`
	// (`intoByte200` is now tainted).
	intoByte200, _, _ := fromReader655.ReadLine()

	// Sink the tainted `intoByte200`:
	sink(intoByte200)
}

func TaintStepTest_BufioReaderReadRune_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader463` into `intoRune480`.

	// Assume that `sourceCQL` has the underlying type of `fromReader463`:
	fromReader463 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader463` to the result `intoRune480`
	// (`intoRune480` is now tainted).
	intoRune480, _, _ := fromReader463.ReadRune()

	// Sink the tainted `intoRune480`:
	sink(intoRune480)
}

func TaintStepTest_BufioReaderReadSlice_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader749` into `intoByte802`.

	// Assume that `sourceCQL` has the underlying type of `fromReader749`:
	fromReader749 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader749` to the result `intoByte802`
	// (`intoByte802` is now tainted).
	intoByte802, _ := fromReader749.ReadSlice(0)

	// Sink the tainted `intoByte802`:
	sink(intoByte802)
}

func TaintStepTest_BufioReaderReadString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader428` into `intoString859`.

	// Assume that `sourceCQL` has the underlying type of `fromReader428`:
	fromReader428 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader428` to the result `intoString859`
	// (`intoString859` is now tainted).
	intoString859, _ := fromReader428.ReadString(0)

	// Sink the tainted `intoString859`:
	sink(intoString859)
}

func TaintStepTest_BufioReaderReset_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader154` into `intoReader501`.

	// Assume that `sourceCQL` has the underlying type of `fromReader154`:
	fromReader154 := sourceCQL.(io.Reader)

	// Declare `intoReader501` variable:
	var intoReader501 bufio.Reader

	// Call the method that transfers the taint
	// from the parameter `fromReader154` to the receiver `intoReader501`
	// (`intoReader501` is now tainted).
	intoReader501.Reset(fromReader154)

	// Sink the tainted `intoReader501`:
	sink(intoReader501)
}

func TaintStepTest_BufioReaderWriteTo_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader673` into `intoWriter248`.

	// Assume that `sourceCQL` has the underlying type of `fromReader673`:
	fromReader673 := sourceCQL.(bufio.Reader)

	// Declare `intoWriter248` variable:
	var intoWriter248 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromReader673` to the argument `intoWriter248`
	// (`intoWriter248` is now tainted).
	fromReader673.WriteTo(intoWriter248)

	// Sink the tainted `intoWriter248`:
	sink(intoWriter248)
}

func TaintStepTest_BufioScannerBytes_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromScanner357` into `intoByte197`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner357`:
	fromScanner357 := sourceCQL.(bufio.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner357` to the result `intoByte197`
	// (`intoByte197` is now tainted).
	intoByte197 := fromScanner357.Bytes()

	// Sink the tainted `intoByte197`:
	sink(intoByte197)
}

func TaintStepTest_BufioScannerText_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromScanner220` into `intoString563`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner220`:
	fromScanner220 := sourceCQL.(bufio.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner220` to the result `intoString563`
	// (`intoString563` is now tainted).
	intoString563 := fromScanner220.Text()

	// Sink the tainted `intoString563`:
	sink(intoString563)
}

func TaintStepTest_BufioWriterReadFrom_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader401` into `intoWriter914`.

	// Assume that `sourceCQL` has the underlying type of `fromReader401`:
	fromReader401 := sourceCQL.(io.Reader)

	// Declare `intoWriter914` variable:
	var intoWriter914 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromReader401` to the receiver `intoWriter914`
	// (`intoWriter914` is now tainted).
	intoWriter914.ReadFrom(fromReader401)

	// Sink the tainted `intoWriter914`:
	sink(intoWriter914)
}

func TaintStepTest_BufioWriterReset_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter500` into `intoWriter504`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter500`:
	fromWriter500 := sourceCQL.(bufio.Writer)

	// Declare `intoWriter504` variable:
	var intoWriter504 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter500` to the argument `intoWriter504`
	// (`intoWriter504` is now tainted).
	fromWriter500.Reset(intoWriter504)

	// Sink the tainted `intoWriter504`:
	sink(intoWriter504)
}

func TaintStepTest_BufioWriterWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte237` into `intoWriter733`.

	// Assume that `sourceCQL` has the underlying type of `fromByte237`:
	fromByte237 := sourceCQL.([]byte)

	// Declare `intoWriter733` variable:
	var intoWriter733 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte237` to the receiver `intoWriter733`
	// (`intoWriter733` is now tainted).
	intoWriter733.Write(fromByte237)

	// Sink the tainted `intoWriter733`:
	sink(intoWriter733)
}

func TaintStepTest_BufioWriterWriteByte_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte597` into `intoWriter332`.

	// Assume that `sourceCQL` has the underlying type of `fromByte597`:
	fromByte597 := sourceCQL.(byte)

	// Declare `intoWriter332` variable:
	var intoWriter332 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte597` to the receiver `intoWriter332`
	// (`intoWriter332` is now tainted).
	intoWriter332.WriteByte(fromByte597)

	// Sink the tainted `intoWriter332`:
	sink(intoWriter332)
}

func TaintStepTest_BufioWriterWriteRune_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune851` into `intoWriter369`.

	// Assume that `sourceCQL` has the underlying type of `fromRune851`:
	fromRune851 := sourceCQL.(rune)

	// Declare `intoWriter369` variable:
	var intoWriter369 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromRune851` to the receiver `intoWriter369`
	// (`intoWriter369` is now tainted).
	intoWriter369.WriteRune(fromRune851)

	// Sink the tainted `intoWriter369`:
	sink(intoWriter369)
}

func TaintStepTest_BufioWriterWriteString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString154` into `intoWriter647`.

	// Assume that `sourceCQL` has the underlying type of `fromString154`:
	fromString154 := sourceCQL.(string)

	// Declare `intoWriter647` variable:
	var intoWriter647 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString154` to the receiver `intoWriter647`
	// (`intoWriter647` is now tainted).
	intoWriter647.WriteString(fromString154)

	// Sink the tainted `intoWriter647`:
	sink(intoWriter647)
}

func RunAllTaints_Bufio() {
	{
		source := newSource()
		TaintStepTest_BufioNewReadWriter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioNewReadWriter_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioNewReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioNewReaderSize_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioNewScanner_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioNewWriter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioNewWriterSize_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioScanBytes_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioScanLines_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioScanRunes_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioScanWords_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderPeek_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderReadByte_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderReadBytes_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderReadLine_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderReadRune_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderReadSlice_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderReadString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderReset_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioReaderWriteTo_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioScannerBytes_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioScannerText_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioWriterReadFrom_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioWriterReset_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioWriterWrite_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioWriterWriteByte_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioWriterWriteRune_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BufioWriterWriteString_B0I0O0(source)
	}
}
