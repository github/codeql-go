// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"bufio"
	"io"
)

func TaintStepTest_BufioNewReadWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader723` into `intoReadWriter292`.

	// Assume that `sourceCQL` has the underlying type of `fromReader723`:
	fromReader723 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader723` to result `intoReadWriter292`
	// (`intoReadWriter292` is now tainted).
	intoReadWriter292 := bufio.NewReadWriter(fromReader723, nil)

	// Return the tainted `intoReadWriter292`:
	return intoReadWriter292
}

func TaintStepTest_BufioNewReadWriter_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReadWriter152` into `intoWriter600`.

	// Assume that `sourceCQL` has the underlying type of `fromReadWriter152`:
	fromReadWriter152 := sourceCQL.(*bufio.ReadWriter)

	// Declare `intoWriter600` variable:
	var intoWriter600 *bufio.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter600`:
	intermediateCQL := bufio.NewReadWriter(nil, intoWriter600)

	// Extra step (`fromReadWriter152` taints `intermediateCQL`, which taints `intoWriter600`:
	link(fromReadWriter152, intermediateCQL)

	// Return the tainted `intoWriter600`:
	return intoWriter600
}

func TaintStepTest_BufioNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader572` into `intoReader520`.

	// Assume that `sourceCQL` has the underlying type of `fromReader572`:
	fromReader572 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader572` to result `intoReader520`
	// (`intoReader520` is now tainted).
	intoReader520 := bufio.NewReader(fromReader572)

	// Return the tainted `intoReader520`:
	return intoReader520
}

func TaintStepTest_BufioNewReaderSize_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader969` into `intoReader327`.

	// Assume that `sourceCQL` has the underlying type of `fromReader969`:
	fromReader969 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader969` to result `intoReader327`
	// (`intoReader327` is now tainted).
	intoReader327 := bufio.NewReaderSize(fromReader969, 0)

	// Return the tainted `intoReader327`:
	return intoReader327
}

func TaintStepTest_BufioNewScanner_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader837` into `intoScanner555`.

	// Assume that `sourceCQL` has the underlying type of `fromReader837`:
	fromReader837 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader837` to result `intoScanner555`
	// (`intoScanner555` is now tainted).
	intoScanner555 := bufio.NewScanner(fromReader837)

	// Return the tainted `intoScanner555`:
	return intoScanner555
}

func TaintStepTest_BufioNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter500` into `intoWriter957`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter500`:
	fromWriter500 := sourceCQL.(*bufio.Writer)

	// Declare `intoWriter957` variable:
	var intoWriter957 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter957`:
	intermediateCQL := bufio.NewWriter(intoWriter957)

	// Extra step (`fromWriter500` taints `intermediateCQL`, which taints `intoWriter957`:
	link(fromWriter500, intermediateCQL)

	// Return the tainted `intoWriter957`:
	return intoWriter957
}

func TaintStepTest_BufioNewWriterSize_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter807` into `intoWriter800`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter807`:
	fromWriter807 := sourceCQL.(*bufio.Writer)

	// Declare `intoWriter800` variable:
	var intoWriter800 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter800`:
	intermediateCQL := bufio.NewWriterSize(intoWriter800, 0)

	// Extra step (`fromWriter807` taints `intermediateCQL`, which taints `intoWriter800`:
	link(fromWriter807, intermediateCQL)

	// Return the tainted `intoWriter800`:
	return intoWriter800
}

func TaintStepTest_BufioScanBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte638` into `intoByte963`.

	// Assume that `sourceCQL` has the underlying type of `fromByte638`:
	fromByte638 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte638` to result `intoByte963`
	// (`intoByte963` is now tainted).
	_, intoByte963, _ := bufio.ScanBytes(fromByte638, false)

	// Return the tainted `intoByte963`:
	return intoByte963
}

func TaintStepTest_BufioScanLines_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte259` into `intoByte408`.

	// Assume that `sourceCQL` has the underlying type of `fromByte259`:
	fromByte259 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte259` to result `intoByte408`
	// (`intoByte408` is now tainted).
	_, intoByte408, _ := bufio.ScanLines(fromByte259, false)

	// Return the tainted `intoByte408`:
	return intoByte408
}

func TaintStepTest_BufioScanRunes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte168` into `intoByte728`.

	// Assume that `sourceCQL` has the underlying type of `fromByte168`:
	fromByte168 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte168` to result `intoByte728`
	// (`intoByte728` is now tainted).
	_, intoByte728, _ := bufio.ScanRunes(fromByte168, false)

	// Return the tainted `intoByte728`:
	return intoByte728
}

func TaintStepTest_BufioScanWords_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte979` into `intoByte867`.

	// Assume that `sourceCQL` has the underlying type of `fromByte979`:
	fromByte979 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte979` to result `intoByte867`
	// (`intoByte867` is now tainted).
	_, intoByte867, _ := bufio.ScanWords(fromByte979, false)

	// Return the tainted `intoByte867`:
	return intoByte867
}

func TaintStepTest_BufioReaderPeek_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader224` into `intoByte617`.

	// Assume that `sourceCQL` has the underlying type of `fromReader224`:
	fromReader224 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader224` to the result `intoByte617`
	// (`intoByte617` is now tainted).
	intoByte617, _ := fromReader224.Peek(0)

	// Return the tainted `intoByte617`:
	return intoByte617
}

func TaintStepTest_BufioReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader849` into `intoByte369`.

	// Assume that `sourceCQL` has the underlying type of `fromReader849`:
	fromReader849 := sourceCQL.(bufio.Reader)

	// Declare `intoByte369` variable:
	var intoByte369 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader849` to the argument `intoByte369`
	// (`intoByte369` is now tainted).
	fromReader849.Read(intoByte369)

	// Return the tainted `intoByte369`:
	return intoByte369
}

func TaintStepTest_BufioReaderReadByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader887` into `intoByte286`.

	// Assume that `sourceCQL` has the underlying type of `fromReader887`:
	fromReader887 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader887` to the result `intoByte286`
	// (`intoByte286` is now tainted).
	intoByte286, _ := fromReader887.ReadByte()

	// Return the tainted `intoByte286`:
	return intoByte286
}

func TaintStepTest_BufioReaderReadBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader309` into `intoByte362`.

	// Assume that `sourceCQL` has the underlying type of `fromReader309`:
	fromReader309 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader309` to the result `intoByte362`
	// (`intoByte362` is now tainted).
	intoByte362, _ := fromReader309.ReadBytes(0)

	// Return the tainted `intoByte362`:
	return intoByte362
}

func TaintStepTest_BufioReaderReadLine_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader134` into `intoByte794`.

	// Assume that `sourceCQL` has the underlying type of `fromReader134`:
	fromReader134 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader134` to the result `intoByte794`
	// (`intoByte794` is now tainted).
	intoByte794, _, _ := fromReader134.ReadLine()

	// Return the tainted `intoByte794`:
	return intoByte794
}

func TaintStepTest_BufioReaderReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader127` into `intoRune123`.

	// Assume that `sourceCQL` has the underlying type of `fromReader127`:
	fromReader127 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader127` to the result `intoRune123`
	// (`intoRune123` is now tainted).
	intoRune123, _, _ := fromReader127.ReadRune()

	// Return the tainted `intoRune123`:
	return intoRune123
}

func TaintStepTest_BufioReaderReadSlice_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader802` into `intoByte569`.

	// Assume that `sourceCQL` has the underlying type of `fromReader802`:
	fromReader802 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader802` to the result `intoByte569`
	// (`intoByte569` is now tainted).
	intoByte569, _ := fromReader802.ReadSlice(0)

	// Return the tainted `intoByte569`:
	return intoByte569
}

func TaintStepTest_BufioReaderReadString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader295` into `intoString921`.

	// Assume that `sourceCQL` has the underlying type of `fromReader295`:
	fromReader295 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader295` to the result `intoString921`
	// (`intoString921` is now tainted).
	intoString921, _ := fromReader295.ReadString(0)

	// Return the tainted `intoString921`:
	return intoString921
}

func TaintStepTest_BufioReaderReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader297` into `intoReader758`.

	// Assume that `sourceCQL` has the underlying type of `fromReader297`:
	fromReader297 := sourceCQL.(io.Reader)

	// Declare `intoReader758` variable:
	var intoReader758 bufio.Reader

	// Call the method that transfers the taint
	// from the parameter `fromReader297` to the receiver `intoReader758`
	// (`intoReader758` is now tainted).
	intoReader758.Reset(fromReader297)

	// Return the tainted `intoReader758`:
	return intoReader758
}

func TaintStepTest_BufioReaderWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader896` into `intoWriter340`.

	// Assume that `sourceCQL` has the underlying type of `fromReader896`:
	fromReader896 := sourceCQL.(bufio.Reader)

	// Declare `intoWriter340` variable:
	var intoWriter340 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromReader896` to the argument `intoWriter340`
	// (`intoWriter340` is now tainted).
	fromReader896.WriteTo(intoWriter340)

	// Return the tainted `intoWriter340`:
	return intoWriter340
}

func TaintStepTest_BufioScannerBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner627` into `intoByte673`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner627`:
	fromScanner627 := sourceCQL.(bufio.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner627` to the result `intoByte673`
	// (`intoByte673` is now tainted).
	intoByte673 := fromScanner627.Bytes()

	// Return the tainted `intoByte673`:
	return intoByte673
}

func TaintStepTest_BufioScannerText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner248` into `intoString381`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner248`:
	fromScanner248 := sourceCQL.(bufio.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner248` to the result `intoString381`
	// (`intoString381` is now tainted).
	intoString381 := fromScanner248.Text()

	// Return the tainted `intoString381`:
	return intoString381
}

func TaintStepTest_BufioWriterReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader529` into `intoWriter506`.

	// Assume that `sourceCQL` has the underlying type of `fromReader529`:
	fromReader529 := sourceCQL.(io.Reader)

	// Declare `intoWriter506` variable:
	var intoWriter506 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromReader529` to the receiver `intoWriter506`
	// (`intoWriter506` is now tainted).
	intoWriter506.ReadFrom(fromReader529)

	// Return the tainted `intoWriter506`:
	return intoWriter506
}

func TaintStepTest_BufioWriterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter664` into `intoWriter466`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter664`:
	fromWriter664 := sourceCQL.(bufio.Writer)

	// Declare `intoWriter466` variable:
	var intoWriter466 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter664` to the argument `intoWriter466`
	// (`intoWriter466` is now tainted).
	fromWriter664.Reset(intoWriter466)

	// Return the tainted `intoWriter466`:
	return intoWriter466
}

func TaintStepTest_BufioWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte267` into `intoWriter879`.

	// Assume that `sourceCQL` has the underlying type of `fromByte267`:
	fromByte267 := sourceCQL.([]byte)

	// Declare `intoWriter879` variable:
	var intoWriter879 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte267` to the receiver `intoWriter879`
	// (`intoWriter879` is now tainted).
	intoWriter879.Write(fromByte267)

	// Return the tainted `intoWriter879`:
	return intoWriter879
}

func TaintStepTest_BufioWriterWriteByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte706` into `intoWriter537`.

	// Assume that `sourceCQL` has the underlying type of `fromByte706`:
	fromByte706 := sourceCQL.(byte)

	// Declare `intoWriter537` variable:
	var intoWriter537 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte706` to the receiver `intoWriter537`
	// (`intoWriter537` is now tainted).
	intoWriter537.WriteByte(fromByte706)

	// Return the tainted `intoWriter537`:
	return intoWriter537
}

func TaintStepTest_BufioWriterWriteRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune287` into `intoWriter690`.

	// Assume that `sourceCQL` has the underlying type of `fromRune287`:
	fromRune287 := sourceCQL.(rune)

	// Declare `intoWriter690` variable:
	var intoWriter690 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromRune287` to the receiver `intoWriter690`
	// (`intoWriter690` is now tainted).
	intoWriter690.WriteRune(fromRune287)

	// Return the tainted `intoWriter690`:
	return intoWriter690
}

func TaintStepTest_BufioWriterWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString886` into `intoWriter637`.

	// Assume that `sourceCQL` has the underlying type of `fromString886`:
	fromString886 := sourceCQL.(string)

	// Declare `intoWriter637` variable:
	var intoWriter637 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString886` to the receiver `intoWriter637`
	// (`intoWriter637` is now tainted).
	intoWriter637.WriteString(fromString886)

	// Return the tainted `intoWriter637`:
	return intoWriter637
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
