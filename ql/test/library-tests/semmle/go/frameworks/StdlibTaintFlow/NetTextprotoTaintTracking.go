// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"bufio"
	"io"
	"net/textproto"
)

func TaintStepTest_NetTextprotoCanonicalMIMEHeaderKey_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString404` into `intoString427`.

	// Assume that `sourceCQL` has the underlying type of `fromString404`:
	fromString404 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString404` to result `intoString427`
	// (`intoString427` is now tainted).
	intoString427 := textproto.CanonicalMIMEHeaderKey(fromString404)

	// Return the tainted `intoString427`:
	return intoString427
}

func TaintStepTest_NetTextprotoNewConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReadWriteCloser962` into `intoConn614`.

	// Assume that `sourceCQL` has the underlying type of `fromReadWriteCloser962`:
	fromReadWriteCloser962 := sourceCQL.(io.ReadWriteCloser)

	// Call the function that transfers the taint
	// from the parameter `fromReadWriteCloser962` to result `intoConn614`
	// (`intoConn614` is now tainted).
	intoConn614 := textproto.NewConn(fromReadWriteCloser962)

	// Return the tainted `intoConn614`:
	return intoConn614
}

func TaintStepTest_NetTextprotoNewConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn769` into `intoReadWriteCloser332`.

	// Assume that `sourceCQL` has the underlying type of `fromConn769`:
	fromConn769 := sourceCQL.(*textproto.Conn)

	// Declare `intoReadWriteCloser332` variable:
	var intoReadWriteCloser332 io.ReadWriteCloser

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoReadWriteCloser332`:
	intermediateCQL := textproto.NewConn(intoReadWriteCloser332)

	// Extra step (`fromConn769` taints `intermediateCQL`, which taints `intoReadWriteCloser332`:
	link(fromConn769, intermediateCQL)

	// Return the tainted `intoReadWriteCloser332`:
	return intoReadWriteCloser332
}

func TaintStepTest_NetTextprotoNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader423` into `intoReader702`.

	// Assume that `sourceCQL` has the underlying type of `fromReader423`:
	fromReader423 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader423` to result `intoReader702`
	// (`intoReader702` is now tainted).
	intoReader702 := textproto.NewReader(fromReader423)

	// Return the tainted `intoReader702`:
	return intoReader702
}

func TaintStepTest_NetTextprotoNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter141` into `intoWriter988`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter141`:
	fromWriter141 := sourceCQL.(*textproto.Writer)

	// Declare `intoWriter988` variable:
	var intoWriter988 *bufio.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter988`:
	intermediateCQL := textproto.NewWriter(intoWriter988)

	// Extra step (`fromWriter141` taints `intermediateCQL`, which taints `intoWriter988`:
	link(fromWriter141, intermediateCQL)

	// Return the tainted `intoWriter988`:
	return intoWriter988
}

func TaintStepTest_NetTextprotoTrimBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte177` into `intoByte919`.

	// Assume that `sourceCQL` has the underlying type of `fromByte177`:
	fromByte177 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte177` to result `intoByte919`
	// (`intoByte919` is now tainted).
	intoByte919 := textproto.TrimBytes(fromByte177)

	// Return the tainted `intoByte919`:
	return intoByte919
}

func TaintStepTest_NetTextprotoTrimString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString986` into `intoString438`.

	// Assume that `sourceCQL` has the underlying type of `fromString986`:
	fromString986 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString986` to result `intoString438`
	// (`intoString438` is now tainted).
	intoString438 := textproto.TrimString(fromString986)

	// Return the tainted `intoString438`:
	return intoString438
}

func TaintStepTest_NetTextprotoMIMEHeaderAdd_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString977` into `intoMIMEHeader912`.

	// Assume that `sourceCQL` has the underlying type of `fromString977`:
	fromString977 := sourceCQL.(string)

	// Declare `intoMIMEHeader912` variable:
	var intoMIMEHeader912 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `fromString977` to the receiver `intoMIMEHeader912`
	// (`intoMIMEHeader912` is now tainted).
	intoMIMEHeader912.Add(fromString977, "")

	// Return the tainted `intoMIMEHeader912`:
	return intoMIMEHeader912
}

func TaintStepTest_NetTextprotoMIMEHeaderAdd_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString520` into `intoMIMEHeader558`.

	// Assume that `sourceCQL` has the underlying type of `fromString520`:
	fromString520 := sourceCQL.(string)

	// Declare `intoMIMEHeader558` variable:
	var intoMIMEHeader558 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `fromString520` to the receiver `intoMIMEHeader558`
	// (`intoMIMEHeader558` is now tainted).
	intoMIMEHeader558.Add("", fromString520)

	// Return the tainted `intoMIMEHeader558`:
	return intoMIMEHeader558
}

func TaintStepTest_NetTextprotoMIMEHeaderGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMIMEHeader313` into `intoString695`.

	// Assume that `sourceCQL` has the underlying type of `fromMIMEHeader313`:
	fromMIMEHeader313 := sourceCQL.(textproto.MIMEHeader)

	// Call the method that transfers the taint
	// from the receiver `fromMIMEHeader313` to the result `intoString695`
	// (`intoString695` is now tainted).
	intoString695 := fromMIMEHeader313.Get("")

	// Return the tainted `intoString695`:
	return intoString695
}

func TaintStepTest_NetTextprotoMIMEHeaderSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString516` into `intoMIMEHeader858`.

	// Assume that `sourceCQL` has the underlying type of `fromString516`:
	fromString516 := sourceCQL.(string)

	// Declare `intoMIMEHeader858` variable:
	var intoMIMEHeader858 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `fromString516` to the receiver `intoMIMEHeader858`
	// (`intoMIMEHeader858` is now tainted).
	intoMIMEHeader858.Set(fromString516, "")

	// Return the tainted `intoMIMEHeader858`:
	return intoMIMEHeader858
}

func TaintStepTest_NetTextprotoMIMEHeaderSet_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString525` into `intoMIMEHeader323`.

	// Assume that `sourceCQL` has the underlying type of `fromString525`:
	fromString525 := sourceCQL.(string)

	// Declare `intoMIMEHeader323` variable:
	var intoMIMEHeader323 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `fromString525` to the receiver `intoMIMEHeader323`
	// (`intoMIMEHeader323` is now tainted).
	intoMIMEHeader323.Set("", fromString525)

	// Return the tainted `intoMIMEHeader323`:
	return intoMIMEHeader323
}

func TaintStepTest_NetTextprotoMIMEHeaderValues_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMIMEHeader706` into `intoString353`.

	// Assume that `sourceCQL` has the underlying type of `fromMIMEHeader706`:
	fromMIMEHeader706 := sourceCQL.(textproto.MIMEHeader)

	// Call the method that transfers the taint
	// from the receiver `fromMIMEHeader706` to the result `intoString353`
	// (`intoString353` is now tainted).
	intoString353 := fromMIMEHeader706.Values("")

	// Return the tainted `intoString353`:
	return intoString353
}

func TaintStepTest_NetTextprotoReaderDotReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader683` into `intoReader536`.

	// Assume that `sourceCQL` has the underlying type of `fromReader683`:
	fromReader683 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader683` to the result `intoReader536`
	// (`intoReader536` is now tainted).
	intoReader536 := fromReader683.DotReader()

	// Return the tainted `intoReader536`:
	return intoReader536
}

func TaintStepTest_NetTextprotoReaderReadCodeLine_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader306` into `intoString557`.

	// Assume that `sourceCQL` has the underlying type of `fromReader306`:
	fromReader306 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader306` to the result `intoString557`
	// (`intoString557` is now tainted).
	_, intoString557, _ := fromReader306.ReadCodeLine(0)

	// Return the tainted `intoString557`:
	return intoString557
}

func TaintStepTest_NetTextprotoReaderReadContinuedLine_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader658` into `intoString721`.

	// Assume that `sourceCQL` has the underlying type of `fromReader658`:
	fromReader658 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader658` to the result `intoString721`
	// (`intoString721` is now tainted).
	intoString721, _ := fromReader658.ReadContinuedLine()

	// Return the tainted `intoString721`:
	return intoString721
}

func TaintStepTest_NetTextprotoReaderReadContinuedLineBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader980` into `intoByte570`.

	// Assume that `sourceCQL` has the underlying type of `fromReader980`:
	fromReader980 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader980` to the result `intoByte570`
	// (`intoByte570` is now tainted).
	intoByte570, _ := fromReader980.ReadContinuedLineBytes()

	// Return the tainted `intoByte570`:
	return intoByte570
}

func TaintStepTest_NetTextprotoReaderReadDotBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader756` into `intoByte149`.

	// Assume that `sourceCQL` has the underlying type of `fromReader756`:
	fromReader756 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader756` to the result `intoByte149`
	// (`intoByte149` is now tainted).
	intoByte149, _ := fromReader756.ReadDotBytes()

	// Return the tainted `intoByte149`:
	return intoByte149
}

func TaintStepTest_NetTextprotoReaderReadDotLines_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader988` into `intoString966`.

	// Assume that `sourceCQL` has the underlying type of `fromReader988`:
	fromReader988 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader988` to the result `intoString966`
	// (`intoString966` is now tainted).
	intoString966, _ := fromReader988.ReadDotLines()

	// Return the tainted `intoString966`:
	return intoString966
}

func TaintStepTest_NetTextprotoReaderReadLine_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader554` into `intoString558`.

	// Assume that `sourceCQL` has the underlying type of `fromReader554`:
	fromReader554 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader554` to the result `intoString558`
	// (`intoString558` is now tainted).
	intoString558, _ := fromReader554.ReadLine()

	// Return the tainted `intoString558`:
	return intoString558
}

func TaintStepTest_NetTextprotoReaderReadLineBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader396` into `intoByte457`.

	// Assume that `sourceCQL` has the underlying type of `fromReader396`:
	fromReader396 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader396` to the result `intoByte457`
	// (`intoByte457` is now tainted).
	intoByte457, _ := fromReader396.ReadLineBytes()

	// Return the tainted `intoByte457`:
	return intoByte457
}

func TaintStepTest_NetTextprotoReaderReadMIMEHeader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader139` into `intoMIMEHeader115`.

	// Assume that `sourceCQL` has the underlying type of `fromReader139`:
	fromReader139 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader139` to the result `intoMIMEHeader115`
	// (`intoMIMEHeader115` is now tainted).
	intoMIMEHeader115, _ := fromReader139.ReadMIMEHeader()

	// Return the tainted `intoMIMEHeader115`:
	return intoMIMEHeader115
}

func TaintStepTest_NetTextprotoReaderReadResponse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader598` into `intoString922`.

	// Assume that `sourceCQL` has the underlying type of `fromReader598`:
	fromReader598 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader598` to the result `intoString922`
	// (`intoString922` is now tainted).
	_, intoString922, _ := fromReader598.ReadResponse(0)

	// Return the tainted `intoString922`:
	return intoString922
}

func TaintStepTest_NetTextprotoWriterDotWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser490` into `intoWriter234`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser490`:
	fromWriteCloser490 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter234` variable:
	var intoWriter234 textproto.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter234`:
	intermediateCQL := intoWriter234.DotWriter()

	// Extra step (`fromWriteCloser490` taints `intermediateCQL`, which taints `intoWriter234`:
	link(fromWriteCloser490, intermediateCQL)

	// Return the tainted `intoWriter234`:
	return intoWriter234
}

func TaintStepTest_NetTextprotoWriterPrintfLine_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString421` into `intoWriter154`.

	// Assume that `sourceCQL` has the underlying type of `fromString421`:
	fromString421 := sourceCQL.(string)

	// Declare `intoWriter154` variable:
	var intoWriter154 textproto.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString421` to the receiver `intoWriter154`
	// (`intoWriter154` is now tainted).
	intoWriter154.PrintfLine(fromString421, nil)

	// Return the tainted `intoWriter154`:
	return intoWriter154
}

func TaintStepTest_NetTextprotoWriterPrintfLine_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface897` into `intoWriter795`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface897`:
	fromInterface897 := sourceCQL.(interface{})

	// Declare `intoWriter795` variable:
	var intoWriter795 textproto.Writer

	// Call the method that transfers the taint
	// from the parameter `fromInterface897` to the receiver `intoWriter795`
	// (`intoWriter795` is now tainted).
	intoWriter795.PrintfLine("", fromInterface897)

	// Return the tainted `intoWriter795`:
	return intoWriter795
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
