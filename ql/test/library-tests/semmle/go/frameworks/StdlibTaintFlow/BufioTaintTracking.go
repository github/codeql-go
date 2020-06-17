package main

import (
	"bufio"
	"io"
)

func TaintStepTest_BufioNewReadWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader455` into `intoReadWriter828`.

	// Assume that `sourceCQL` has the underlying type of `fromReader455`:
	fromReader455 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader455` to result `intoReadWriter828`
	// (`intoReadWriter828` is now tainted).
	intoReadWriter828 := bufio.NewReadWriter(fromReader455, nil)

	// Return the tainted `intoReadWriter828`:
	return intoReadWriter828
}

func TaintStepTest_BufioNewReadWriter_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReadWriter167` into `intoWriter548`.

	// Assume that `sourceCQL` has the underlying type of `fromReadWriter167`:
	fromReadWriter167 := sourceCQL.(*bufio.ReadWriter)

	// Declare `intoWriter548` variable:
	var intoWriter548 *bufio.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter548`:
	intermediateCQL := bufio.NewReadWriter(nil, intoWriter548)

	// Extra step (`fromReadWriter167` taints `intermediateCQL`, which taints `intoWriter548`:
	link(fromReadWriter167, intermediateCQL)

	// Return the tainted `intoWriter548`:
	return intoWriter548
}

func TaintStepTest_BufioNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader883` into `intoReader526`.

	// Assume that `sourceCQL` has the underlying type of `fromReader883`:
	fromReader883 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader883` to result `intoReader526`
	// (`intoReader526` is now tainted).
	intoReader526 := bufio.NewReader(fromReader883)

	// Return the tainted `intoReader526`:
	return intoReader526
}

func TaintStepTest_BufioNewReaderSize_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader831` into `intoReader619`.

	// Assume that `sourceCQL` has the underlying type of `fromReader831`:
	fromReader831 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader831` to result `intoReader619`
	// (`intoReader619` is now tainted).
	intoReader619 := bufio.NewReaderSize(fromReader831, 0)

	// Return the tainted `intoReader619`:
	return intoReader619
}

func TaintStepTest_BufioNewScanner_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader174` into `intoScanner736`.

	// Assume that `sourceCQL` has the underlying type of `fromReader174`:
	fromReader174 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader174` to result `intoScanner736`
	// (`intoScanner736` is now tainted).
	intoScanner736 := bufio.NewScanner(fromReader174)

	// Return the tainted `intoScanner736`:
	return intoScanner736
}

func TaintStepTest_BufioNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter212` into `intoWriter365`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter212`:
	fromWriter212 := sourceCQL.(*bufio.Writer)

	// Declare `intoWriter365` variable:
	var intoWriter365 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter365`:
	intermediateCQL := bufio.NewWriter(intoWriter365)

	// Extra step (`fromWriter212` taints `intermediateCQL`, which taints `intoWriter365`:
	link(fromWriter212, intermediateCQL)

	// Return the tainted `intoWriter365`:
	return intoWriter365
}

func TaintStepTest_BufioNewWriterSize_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter414` into `intoWriter417`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter414`:
	fromWriter414 := sourceCQL.(*bufio.Writer)

	// Declare `intoWriter417` variable:
	var intoWriter417 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter417`:
	intermediateCQL := bufio.NewWriterSize(intoWriter417, 0)

	// Extra step (`fromWriter414` taints `intermediateCQL`, which taints `intoWriter417`:
	link(fromWriter414, intermediateCQL)

	// Return the tainted `intoWriter417`:
	return intoWriter417
}

func TaintStepTest_BufioScanBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte755` into `intoByte763`.

	// Assume that `sourceCQL` has the underlying type of `fromByte755`:
	fromByte755 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte755` to result `intoByte763`
	// (`intoByte763` is now tainted).
	_, intoByte763, _ := bufio.ScanBytes(fromByte755, false)

	// Return the tainted `intoByte763`:
	return intoByte763
}

func TaintStepTest_BufioScanLines_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte612` into `intoByte693`.

	// Assume that `sourceCQL` has the underlying type of `fromByte612`:
	fromByte612 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte612` to result `intoByte693`
	// (`intoByte693` is now tainted).
	_, intoByte693, _ := bufio.ScanLines(fromByte612, false)

	// Return the tainted `intoByte693`:
	return intoByte693
}

func TaintStepTest_BufioScanRunes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte516` into `intoByte182`.

	// Assume that `sourceCQL` has the underlying type of `fromByte516`:
	fromByte516 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte516` to result `intoByte182`
	// (`intoByte182` is now tainted).
	_, intoByte182, _ := bufio.ScanRunes(fromByte516, false)

	// Return the tainted `intoByte182`:
	return intoByte182
}

func TaintStepTest_BufioScanWords_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte784` into `intoByte862`.

	// Assume that `sourceCQL` has the underlying type of `fromByte784`:
	fromByte784 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte784` to result `intoByte862`
	// (`intoByte862` is now tainted).
	_, intoByte862, _ := bufio.ScanWords(fromByte784, false)

	// Return the tainted `intoByte862`:
	return intoByte862
}

func TaintStepTest_BufioReaderPeek_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader493` into `intoByte538`.

	// Assume that `sourceCQL` has the underlying type of `fromReader493`:
	fromReader493 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader493` to the result `intoByte538`
	// (`intoByte538` is now tainted).
	intoByte538, _ := fromReader493.Peek(0)

	// Return the tainted `intoByte538`:
	return intoByte538
}

func TaintStepTest_BufioReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader577` into `intoByte144`.

	// Assume that `sourceCQL` has the underlying type of `fromReader577`:
	fromReader577 := sourceCQL.(bufio.Reader)

	// Declare `intoByte144` variable:
	var intoByte144 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader577` to the argument `intoByte144`
	// (`intoByte144` is now tainted).
	fromReader577.Read(intoByte144)

	// Return the tainted `intoByte144`:
	return intoByte144
}

func TaintStepTest_BufioReaderReadByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader339` into `intoByte180`.

	// Assume that `sourceCQL` has the underlying type of `fromReader339`:
	fromReader339 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader339` to the result `intoByte180`
	// (`intoByte180` is now tainted).
	intoByte180, _ := fromReader339.ReadByte()

	// Return the tainted `intoByte180`:
	return intoByte180
}

func TaintStepTest_BufioReaderReadBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader744` into `intoByte160`.

	// Assume that `sourceCQL` has the underlying type of `fromReader744`:
	fromReader744 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader744` to the result `intoByte160`
	// (`intoByte160` is now tainted).
	intoByte160, _ := fromReader744.ReadBytes(0)

	// Return the tainted `intoByte160`:
	return intoByte160
}

func TaintStepTest_BufioReaderReadLine_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader651` into `intoByte759`.

	// Assume that `sourceCQL` has the underlying type of `fromReader651`:
	fromReader651 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader651` to the result `intoByte759`
	// (`intoByte759` is now tainted).
	intoByte759, _, _ := fromReader651.ReadLine()

	// Return the tainted `intoByte759`:
	return intoByte759
}

func TaintStepTest_BufioReaderReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader974` into `intoRune768`.

	// Assume that `sourceCQL` has the underlying type of `fromReader974`:
	fromReader974 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader974` to the result `intoRune768`
	// (`intoRune768` is now tainted).
	intoRune768, _, _ := fromReader974.ReadRune()

	// Return the tainted `intoRune768`:
	return intoRune768
}

func TaintStepTest_BufioReaderReadSlice_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader631` into `intoByte679`.

	// Assume that `sourceCQL` has the underlying type of `fromReader631`:
	fromReader631 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader631` to the result `intoByte679`
	// (`intoByte679` is now tainted).
	intoByte679, _ := fromReader631.ReadSlice(0)

	// Return the tainted `intoByte679`:
	return intoByte679
}

func TaintStepTest_BufioReaderReadString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader689` into `intoString670`.

	// Assume that `sourceCQL` has the underlying type of `fromReader689`:
	fromReader689 := sourceCQL.(bufio.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader689` to the result `intoString670`
	// (`intoString670` is now tainted).
	intoString670, _ := fromReader689.ReadString(0)

	// Return the tainted `intoString670`:
	return intoString670
}

func TaintStepTest_BufioReaderReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader294` into `intoReader907`.

	// Assume that `sourceCQL` has the underlying type of `fromReader294`:
	fromReader294 := sourceCQL.(io.Reader)

	// Declare `intoReader907` variable:
	var intoReader907 bufio.Reader

	// Call the method that transfers the taint
	// from the parameter `fromReader294` to the receiver `intoReader907`
	// (`intoReader907` is now tainted).
	intoReader907.Reset(fromReader294)

	// Return the tainted `intoReader907`:
	return intoReader907
}

func TaintStepTest_BufioReaderWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader429` into `intoWriter382`.

	// Assume that `sourceCQL` has the underlying type of `fromReader429`:
	fromReader429 := sourceCQL.(bufio.Reader)

	// Declare `intoWriter382` variable:
	var intoWriter382 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromReader429` to the argument `intoWriter382`
	// (`intoWriter382` is now tainted).
	fromReader429.WriteTo(intoWriter382)

	// Return the tainted `intoWriter382`:
	return intoWriter382
}

func TaintStepTest_BufioScannerBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner979` into `intoByte932`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner979`:
	fromScanner979 := sourceCQL.(bufio.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner979` to the result `intoByte932`
	// (`intoByte932` is now tainted).
	intoByte932 := fromScanner979.Bytes()

	// Return the tainted `intoByte932`:
	return intoByte932
}

func TaintStepTest_BufioScannerText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner188` into `intoString397`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner188`:
	fromScanner188 := sourceCQL.(bufio.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner188` to the result `intoString397`
	// (`intoString397` is now tainted).
	intoString397 := fromScanner188.Text()

	// Return the tainted `intoString397`:
	return intoString397
}

func TaintStepTest_BufioWriterReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader286` into `intoWriter297`.

	// Assume that `sourceCQL` has the underlying type of `fromReader286`:
	fromReader286 := sourceCQL.(io.Reader)

	// Declare `intoWriter297` variable:
	var intoWriter297 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromReader286` to the receiver `intoWriter297`
	// (`intoWriter297` is now tainted).
	intoWriter297.ReadFrom(fromReader286)

	// Return the tainted `intoWriter297`:
	return intoWriter297
}

func TaintStepTest_BufioWriterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter827` into `intoWriter323`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter827`:
	fromWriter827 := sourceCQL.(bufio.Writer)

	// Declare `intoWriter323` variable:
	var intoWriter323 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter827` to the argument `intoWriter323`
	// (`intoWriter323` is now tainted).
	fromWriter827.Reset(intoWriter323)

	// Return the tainted `intoWriter323`:
	return intoWriter323
}

func TaintStepTest_BufioWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte493` into `intoWriter194`.

	// Assume that `sourceCQL` has the underlying type of `fromByte493`:
	fromByte493 := sourceCQL.([]byte)

	// Declare `intoWriter194` variable:
	var intoWriter194 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte493` to the receiver `intoWriter194`
	// (`intoWriter194` is now tainted).
	intoWriter194.Write(fromByte493)

	// Return the tainted `intoWriter194`:
	return intoWriter194
}

func TaintStepTest_BufioWriterWriteByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte652` into `intoWriter733`.

	// Assume that `sourceCQL` has the underlying type of `fromByte652`:
	fromByte652 := sourceCQL.(byte)

	// Declare `intoWriter733` variable:
	var intoWriter733 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte652` to the receiver `intoWriter733`
	// (`intoWriter733` is now tainted).
	intoWriter733.WriteByte(fromByte652)

	// Return the tainted `intoWriter733`:
	return intoWriter733
}

func TaintStepTest_BufioWriterWriteRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune157` into `intoWriter898`.

	// Assume that `sourceCQL` has the underlying type of `fromRune157`:
	fromRune157 := sourceCQL.(rune)

	// Declare `intoWriter898` variable:
	var intoWriter898 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromRune157` to the receiver `intoWriter898`
	// (`intoWriter898` is now tainted).
	intoWriter898.WriteRune(fromRune157)

	// Return the tainted `intoWriter898`:
	return intoWriter898
}

func TaintStepTest_BufioWriterWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString802` into `intoWriter439`.

	// Assume that `sourceCQL` has the underlying type of `fromString802`:
	fromString802 := sourceCQL.(string)

	// Declare `intoWriter439` variable:
	var intoWriter439 bufio.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString802` to the receiver `intoWriter439`
	// (`intoWriter439` is now tainted).
	intoWriter439.WriteString(fromString802)

	// Return the tainted `intoWriter439`:
	return intoWriter439
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
