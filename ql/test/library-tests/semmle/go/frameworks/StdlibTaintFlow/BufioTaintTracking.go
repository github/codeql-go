// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"bufio"
	"io"
)

func TaintStepTest_BufioNewReadWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader656` into `intoReadWriter414`.

	// Assume that `sourceCQL` has the underlying type of `fromReader656`:
	fromReader656 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader656` to result `intoReadWriter414`
	// (`intoReadWriter414` is now tainted).
	intoReadWriter414 := bufio.NewReadWriter(fromReader656, nil)

	// Return the tainted `intoReadWriter414`:
	return intoReadWriter414
}

func TaintStepTest_BufioNewReadWriter_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReadWriter518` into `intoWriter650`.

	// Assume that `sourceCQL` has the underlying type of `fromReadWriter518`:
	fromReadWriter518 := sourceCQL.(*bufio.ReadWriter)

	// Declare `intoWriter650` variable:
	var intoWriter650 *bufio.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter650`:
	intermediateCQL := bufio.NewReadWriter(nil, intoWriter650)

	// Extra step (`fromReadWriter518` taints `intermediateCQL`, which taints `intoWriter650`:
	link(fromReadWriter518, intermediateCQL)

	// Return the tainted `intoWriter650`:
	return intoWriter650
}

func TaintStepTest_BufioNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader784` into `intoReader957`.

	// Assume that `sourceCQL` has the underlying type of `fromReader784`:
	fromReader784 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader784` to result `intoReader957`
	// (`intoReader957` is now tainted).
	intoReader957 := bufio.NewReader(fromReader784)

	// Return the tainted `intoReader957`:
	return intoReader957
}

func TaintStepTest_BufioNewReaderSize_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader520` into `intoReader443`.

	// Assume that `sourceCQL` has the underlying type of `fromReader520`:
	fromReader520 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader520` to result `intoReader443`
	// (`intoReader443` is now tainted).
	intoReader443 := bufio.NewReaderSize(fromReader520, 0)

	// Return the tainted `intoReader443`:
	return intoReader443
}

func TaintStepTest_BufioNewScanner_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader127` into `intoScanner483`.

	// Assume that `sourceCQL` has the underlying type of `fromReader127`:
	fromReader127 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader127` to result `intoScanner483`
	// (`intoScanner483` is now tainted).
	intoScanner483 := bufio.NewScanner(fromReader127)

	// Return the tainted `intoScanner483`:
	return intoScanner483
}

func TaintStepTest_BufioNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter989` into `intoWriter982`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter989`:
	fromWriter989 := sourceCQL.(*bufio.Writer)

	// Declare `intoWriter982` variable:
	var intoWriter982 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter982`:
	intermediateCQL := bufio.NewWriter(intoWriter982)

	// Extra step (`fromWriter989` taints `intermediateCQL`, which taints `intoWriter982`:
	link(fromWriter989, intermediateCQL)

	// Return the tainted `intoWriter982`:
	return intoWriter982
}

func TaintStepTest_BufioNewWriterSize_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter417` into `intoWriter584`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter417`:
	fromWriter417 := sourceCQL.(*bufio.Writer)

	// Declare `intoWriter584` variable:
	var intoWriter584 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter584`:
	intermediateCQL := bufio.NewWriterSize(intoWriter584, 0)

	// Extra step (`fromWriter417` taints `intermediateCQL`, which taints `intoWriter584`:
	link(fromWriter417, intermediateCQL)

	// Return the tainted `intoWriter584`:
	return intoWriter584
}

func TaintStepTest_BufioScanBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte991` into `intoByte881`.

	// Assume that `sourceCQL` has the underlying type of `fromByte991`:
	fromByte991 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte991` to result `intoByte881`
	// (`intoByte881` is now tainted).
	_, intoByte881, _ := bufio.ScanBytes(fromByte991, false)

	// Return the tainted `intoByte881`:
	return intoByte881
}

func TaintStepTest_BufioScanLines_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte186` into `intoByte284`.

	// Assume that `sourceCQL` has the underlying type of `fromByte186`:
	fromByte186 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte186` to result `intoByte284`
	// (`intoByte284` is now tainted).
	_, intoByte284, _ := bufio.ScanLines(fromByte186, false)

	// Return the tainted `intoByte284`:
	return intoByte284
}

func TaintStepTest_BufioScanRunes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte908` into `intoByte137`.

	// Assume that `sourceCQL` has the underlying type of `fromByte908`:
	fromByte908 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte908` to result `intoByte137`
	// (`intoByte137` is now tainted).
	_, intoByte137, _ := bufio.ScanRunes(fromByte908, false)

	// Return the tainted `intoByte137`:
	return intoByte137
}

func TaintStepTest_BufioScanWords_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte494` into `intoByte873`.

	// Assume that `sourceCQL` has the underlying type of `fromByte494`:
	fromByte494 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte494` to result `intoByte873`
	// (`intoByte873` is now tainted).
	_, intoByte873, _ := bufio.ScanWords(fromByte494, false)

	// Return the tainted `intoByte873`:
	return intoByte873
}

func TaintStepTest_BufioReaderPeek_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader599` into `intoByte409`.

	// Assume that `sourceCQL` has the underlying type of `fromReader599`:
	fromReader599 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader599` to the result `intoByte409`
	// (`intoByte409` is now tainted).
	intoByte409, _ := fromReader599.Peek(0)

	// Return the tainted `intoByte409`:
	return intoByte409
}

func TaintStepTest_BufioReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader246` into `intoByte898`.

	// Assume that `sourceCQL` has the underlying type of `fromReader246`:
	fromReader246 := sourceCQL.(bufio.Reader)

	// Declare `intoByte898` variable:
	var intoByte898 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader246` to the argument `intoByte898`
	// (`intoByte898` is now tainted).
	fromReader246.Read(intoByte898)

	// Return the tainted `intoByte898`:
	return intoByte898
}

func TaintStepTest_BufioReaderReadByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader598` into `intoByte631`.

	// Assume that `sourceCQL` has the underlying type of `fromReader598`:
	fromReader598 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader598` to the result `intoByte631`
	// (`intoByte631` is now tainted).
	intoByte631, _ := fromReader598.ReadByte()

	// Return the tainted `intoByte631`:
	return intoByte631
}

func TaintStepTest_BufioReaderReadBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader165` into `intoByte150`.

	// Assume that `sourceCQL` has the underlying type of `fromReader165`:
	fromReader165 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader165` to the result `intoByte150`
	// (`intoByte150` is now tainted).
	intoByte150, _ := fromReader165.ReadBytes(0)

	// Return the tainted `intoByte150`:
	return intoByte150
}

func TaintStepTest_BufioReaderReadLine_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader340` into `intoByte471`.

	// Assume that `sourceCQL` has the underlying type of `fromReader340`:
	fromReader340 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader340` to the result `intoByte471`
	// (`intoByte471` is now tainted).
	intoByte471, _, _ := fromReader340.ReadLine()

	// Return the tainted `intoByte471`:
	return intoByte471
}

func TaintStepTest_BufioReaderReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader290` into `intoRune758`.

	// Assume that `sourceCQL` has the underlying type of `fromReader290`:
	fromReader290 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader290` to the result `intoRune758`
	// (`intoRune758` is now tainted).
	intoRune758, _, _ := fromReader290.ReadRune()

	// Return the tainted `intoRune758`:
	return intoRune758
}

func TaintStepTest_BufioReaderReadSlice_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader396` into `intoByte707`.

	// Assume that `sourceCQL` has the underlying type of `fromReader396`:
	fromReader396 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader396` to the result `intoByte707`
	// (`intoByte707` is now tainted).
	intoByte707, _ := fromReader396.ReadSlice(0)

	// Return the tainted `intoByte707`:
	return intoByte707
}

func TaintStepTest_BufioReaderReadString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader912` into `intoString718`.

	// Assume that `sourceCQL` has the underlying type of `fromReader912`:
	fromReader912 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader912` to the result `intoString718`
	// (`intoString718` is now tainted).
	intoString718, _ := fromReader912.ReadString(0)

	// Return the tainted `intoString718`:
	return intoString718
}

func TaintStepTest_BufioReaderReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader972` into `intoReader633`.

	// Assume that `sourceCQL` has the underlying type of `fromReader972`:
	fromReader972 := sourceCQL.(io.Reader)

	// Declare `intoReader633` variable:
	var intoReader633 bufio.Reader

	// Call the method that transfers the taint
	// from the parameter `fromReader972` to the receiver `intoReader633`
	// (`intoReader633` is now tainted).
	intoReader633.Reset(fromReader972)

	// Return the tainted `intoReader633`:
	return intoReader633
}

func TaintStepTest_BufioReaderWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader316` into `intoWriter145`.

	// Assume that `sourceCQL` has the underlying type of `fromReader316`:
	fromReader316 := sourceCQL.(bufio.Reader)

	// Declare `intoWriter145` variable:
	var intoWriter145 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromReader316` to the argument `intoWriter145`
	// (`intoWriter145` is now tainted).
	fromReader316.WriteTo(intoWriter145)

	// Return the tainted `intoWriter145`:
	return intoWriter145
}

func TaintStepTest_BufioScannerBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner817` into `intoByte474`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner817`:
	fromScanner817 := sourceCQL.(bufio.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner817` to the result `intoByte474`
	// (`intoByte474` is now tainted).
	intoByte474 := fromScanner817.Bytes()

	// Return the tainted `intoByte474`:
	return intoByte474
}

func TaintStepTest_BufioScannerText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner832` into `intoString378`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner832`:
	fromScanner832 := sourceCQL.(bufio.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner832` to the result `intoString378`
	// (`intoString378` is now tainted).
	intoString378 := fromScanner832.Text()

	// Return the tainted `intoString378`:
	return intoString378
}

func TaintStepTest_BufioWriterReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader541` into `intoWriter139`.

	// Assume that `sourceCQL` has the underlying type of `fromReader541`:
	fromReader541 := sourceCQL.(io.Reader)

	// Declare `intoWriter139` variable:
	var intoWriter139 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromReader541` to the receiver `intoWriter139`
	// (`intoWriter139` is now tainted).
	intoWriter139.ReadFrom(fromReader541)

	// Return the tainted `intoWriter139`:
	return intoWriter139
}

func TaintStepTest_BufioWriterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter814` into `intoWriter768`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter814`:
	fromWriter814 := sourceCQL.(bufio.Writer)

	// Declare `intoWriter768` variable:
	var intoWriter768 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter814` to the argument `intoWriter768`
	// (`intoWriter768` is now tainted).
	fromWriter814.Reset(intoWriter768)

	// Return the tainted `intoWriter768`:
	return intoWriter768
}

func TaintStepTest_BufioWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte468` into `intoWriter736`.

	// Assume that `sourceCQL` has the underlying type of `fromByte468`:
	fromByte468 := sourceCQL.([]byte)

	// Declare `intoWriter736` variable:
	var intoWriter736 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte468` to the receiver `intoWriter736`
	// (`intoWriter736` is now tainted).
	intoWriter736.Write(fromByte468)

	// Return the tainted `intoWriter736`:
	return intoWriter736
}

func TaintStepTest_BufioWriterWriteByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte516` into `intoWriter246`.

	// Assume that `sourceCQL` has the underlying type of `fromByte516`:
	fromByte516 := sourceCQL.(byte)

	// Declare `intoWriter246` variable:
	var intoWriter246 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte516` to the receiver `intoWriter246`
	// (`intoWriter246` is now tainted).
	intoWriter246.WriteByte(fromByte516)

	// Return the tainted `intoWriter246`:
	return intoWriter246
}

func TaintStepTest_BufioWriterWriteRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune679` into `intoWriter736`.

	// Assume that `sourceCQL` has the underlying type of `fromRune679`:
	fromRune679 := sourceCQL.(rune)

	// Declare `intoWriter736` variable:
	var intoWriter736 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromRune679` to the receiver `intoWriter736`
	// (`intoWriter736` is now tainted).
	intoWriter736.WriteRune(fromRune679)

	// Return the tainted `intoWriter736`:
	return intoWriter736
}

func TaintStepTest_BufioWriterWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString839` into `intoWriter273`.

	// Assume that `sourceCQL` has the underlying type of `fromString839`:
	fromString839 := sourceCQL.(string)

	// Declare `intoWriter273` variable:
	var intoWriter273 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString839` to the receiver `intoWriter273`
	// (`intoWriter273` is now tainted).
	intoWriter273.WriteString(fromString839)

	// Return the tainted `intoWriter273`:
	return intoWriter273
}

func RunAllTaints_Bufio() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioNewReadWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioNewReadWriter_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioNewReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioNewReaderSize_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioNewScanner_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioNewWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioNewWriterSize_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioScanBytes_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioScanLines_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioScanRunes_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioScanWords_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioReaderPeek_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioReaderRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioReaderReadByte_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioReaderReadBytes_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioReaderReadLine_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioReaderReadRune_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioReaderReadSlice_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioReaderReadString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioReaderReset_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioReaderWriteTo_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioScannerBytes_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioScannerText_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioWriterReadFrom_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioWriterReset_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioWriterWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioWriterWriteByte_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioWriterWriteRune_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BufioWriterWriteString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
