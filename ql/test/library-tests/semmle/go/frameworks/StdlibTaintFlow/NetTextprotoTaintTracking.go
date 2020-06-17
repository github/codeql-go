// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"bufio"
	"io"
	"net/textproto"
)

func TaintStepTest_NetTextprotoCanonicalMIMEHeaderKey_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString656` into `intoString414`.

	// Assume that `sourceCQL` has the underlying type of `fromString656`:
	fromString656 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString656` to result `intoString414`
	// (`intoString414` is now tainted).
	intoString414 := textproto.CanonicalMIMEHeaderKey(fromString656)

	// Return the tainted `intoString414`:
	return intoString414
}

func TaintStepTest_NetTextprotoNewConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReadWriteCloser518` into `intoConn650`.

	// Assume that `sourceCQL` has the underlying type of `fromReadWriteCloser518`:
	fromReadWriteCloser518 := sourceCQL.(io.ReadWriteCloser)

	// Call the function that transfers the taint
	// from the parameter `fromReadWriteCloser518` to result `intoConn650`
	// (`intoConn650` is now tainted).
	intoConn650 := textproto.NewConn(fromReadWriteCloser518)

	// Return the tainted `intoConn650`:
	return intoConn650
}

func TaintStepTest_NetTextprotoNewConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn784` into `intoReadWriteCloser957`.

	// Assume that `sourceCQL` has the underlying type of `fromConn784`:
	fromConn784 := sourceCQL.(*textproto.Conn)

	// Declare `intoReadWriteCloser957` variable:
	var intoReadWriteCloser957 io.ReadWriteCloser

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoReadWriteCloser957`:
	intermediateCQL := textproto.NewConn(intoReadWriteCloser957)

	// Extra step (`fromConn784` taints `intermediateCQL`, which taints `intoReadWriteCloser957`:
	link(fromConn784, intermediateCQL)

	// Return the tainted `intoReadWriteCloser957`:
	return intoReadWriteCloser957
}

func TaintStepTest_NetTextprotoNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader520` into `intoReader443`.

	// Assume that `sourceCQL` has the underlying type of `fromReader520`:
	fromReader520 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader520` to result `intoReader443`
	// (`intoReader443` is now tainted).
	intoReader443 := textproto.NewReader(fromReader520)

	// Return the tainted `intoReader443`:
	return intoReader443
}

func TaintStepTest_NetTextprotoNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter127` into `intoWriter483`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter127`:
	fromWriter127 := sourceCQL.(*textproto.Writer)

	// Declare `intoWriter483` variable:
	var intoWriter483 *bufio.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter483`:
	intermediateCQL := textproto.NewWriter(intoWriter483)

	// Extra step (`fromWriter127` taints `intermediateCQL`, which taints `intoWriter483`:
	link(fromWriter127, intermediateCQL)

	// Return the tainted `intoWriter483`:
	return intoWriter483
}

func TaintStepTest_NetTextprotoTrimBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte989` into `intoByte982`.

	// Assume that `sourceCQL` has the underlying type of `fromByte989`:
	fromByte989 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte989` to result `intoByte982`
	// (`intoByte982` is now tainted).
	intoByte982 := textproto.TrimBytes(fromByte989)

	// Return the tainted `intoByte982`:
	return intoByte982
}

func TaintStepTest_NetTextprotoTrimString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString417` into `intoString584`.

	// Assume that `sourceCQL` has the underlying type of `fromString417`:
	fromString417 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString417` to result `intoString584`
	// (`intoString584` is now tainted).
	intoString584 := textproto.TrimString(fromString417)

	// Return the tainted `intoString584`:
	return intoString584
}

func TaintStepTest_NetTextprotoMIMEHeaderAdd_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString991` into `intoMIMEHeader881`.

	// Assume that `sourceCQL` has the underlying type of `fromString991`:
	fromString991 := sourceCQL.(string)

	// Declare `intoMIMEHeader881` variable:
	var intoMIMEHeader881 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `fromString991` to the receiver `intoMIMEHeader881`
	// (`intoMIMEHeader881` is now tainted).
	intoMIMEHeader881.Add(fromString991, "")

	// Return the tainted `intoMIMEHeader881`:
	return intoMIMEHeader881
}

func TaintStepTest_NetTextprotoMIMEHeaderAdd_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString186` into `intoMIMEHeader284`.

	// Assume that `sourceCQL` has the underlying type of `fromString186`:
	fromString186 := sourceCQL.(string)

	// Declare `intoMIMEHeader284` variable:
	var intoMIMEHeader284 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `fromString186` to the receiver `intoMIMEHeader284`
	// (`intoMIMEHeader284` is now tainted).
	intoMIMEHeader284.Add("", fromString186)

	// Return the tainted `intoMIMEHeader284`:
	return intoMIMEHeader284
}

func TaintStepTest_NetTextprotoMIMEHeaderGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMIMEHeader908` into `intoString137`.

	// Assume that `sourceCQL` has the underlying type of `fromMIMEHeader908`:
	fromMIMEHeader908 := sourceCQL.(textproto.MIMEHeader)

	// Call the method that transfers the taint
	// from the receiver `fromMIMEHeader908` to the result `intoString137`
	// (`intoString137` is now tainted).
	intoString137 := fromMIMEHeader908.Get("")

	// Return the tainted `intoString137`:
	return intoString137
}

func TaintStepTest_NetTextprotoMIMEHeaderSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString494` into `intoMIMEHeader873`.

	// Assume that `sourceCQL` has the underlying type of `fromString494`:
	fromString494 := sourceCQL.(string)

	// Declare `intoMIMEHeader873` variable:
	var intoMIMEHeader873 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `fromString494` to the receiver `intoMIMEHeader873`
	// (`intoMIMEHeader873` is now tainted).
	intoMIMEHeader873.Set(fromString494, "")

	// Return the tainted `intoMIMEHeader873`:
	return intoMIMEHeader873
}

func TaintStepTest_NetTextprotoMIMEHeaderSet_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString599` into `intoMIMEHeader409`.

	// Assume that `sourceCQL` has the underlying type of `fromString599`:
	fromString599 := sourceCQL.(string)

	// Declare `intoMIMEHeader409` variable:
	var intoMIMEHeader409 textproto.MIMEHeader

	// Call the method that transfers the taint
	// from the parameter `fromString599` to the receiver `intoMIMEHeader409`
	// (`intoMIMEHeader409` is now tainted).
	intoMIMEHeader409.Set("", fromString599)

	// Return the tainted `intoMIMEHeader409`:
	return intoMIMEHeader409
}

func TaintStepTest_NetTextprotoMIMEHeaderValues_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMIMEHeader246` into `intoString898`.

	// Assume that `sourceCQL` has the underlying type of `fromMIMEHeader246`:
	fromMIMEHeader246 := sourceCQL.(textproto.MIMEHeader)

	// Call the method that transfers the taint
	// from the receiver `fromMIMEHeader246` to the result `intoString898`
	// (`intoString898` is now tainted).
	intoString898 := fromMIMEHeader246.Values("")

	// Return the tainted `intoString898`:
	return intoString898
}

func TaintStepTest_NetTextprotoReaderDotReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader598` into `intoReader631`.

	// Assume that `sourceCQL` has the underlying type of `fromReader598`:
	fromReader598 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader598` to the result `intoReader631`
	// (`intoReader631` is now tainted).
	intoReader631 := fromReader598.DotReader()

	// Return the tainted `intoReader631`:
	return intoReader631
}

func TaintStepTest_NetTextprotoReaderReadCodeLine_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader165` into `intoString150`.

	// Assume that `sourceCQL` has the underlying type of `fromReader165`:
	fromReader165 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader165` to the result `intoString150`
	// (`intoString150` is now tainted).
	_, intoString150, _ := fromReader165.ReadCodeLine(0)

	// Return the tainted `intoString150`:
	return intoString150
}

func TaintStepTest_NetTextprotoReaderReadContinuedLine_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader340` into `intoString471`.

	// Assume that `sourceCQL` has the underlying type of `fromReader340`:
	fromReader340 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader340` to the result `intoString471`
	// (`intoString471` is now tainted).
	intoString471, _ := fromReader340.ReadContinuedLine()

	// Return the tainted `intoString471`:
	return intoString471
}

func TaintStepTest_NetTextprotoReaderReadContinuedLineBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader290` into `intoByte758`.

	// Assume that `sourceCQL` has the underlying type of `fromReader290`:
	fromReader290 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader290` to the result `intoByte758`
	// (`intoByte758` is now tainted).
	intoByte758, _ := fromReader290.ReadContinuedLineBytes()

	// Return the tainted `intoByte758`:
	return intoByte758
}

func TaintStepTest_NetTextprotoReaderReadDotBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader396` into `intoByte707`.

	// Assume that `sourceCQL` has the underlying type of `fromReader396`:
	fromReader396 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader396` to the result `intoByte707`
	// (`intoByte707` is now tainted).
	intoByte707, _ := fromReader396.ReadDotBytes()

	// Return the tainted `intoByte707`:
	return intoByte707
}

func TaintStepTest_NetTextprotoReaderReadDotLines_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader912` into `intoString718`.

	// Assume that `sourceCQL` has the underlying type of `fromReader912`:
	fromReader912 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader912` to the result `intoString718`
	// (`intoString718` is now tainted).
	intoString718, _ := fromReader912.ReadDotLines()

	// Return the tainted `intoString718`:
	return intoString718
}

func TaintStepTest_NetTextprotoReaderReadLine_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader972` into `intoString633`.

	// Assume that `sourceCQL` has the underlying type of `fromReader972`:
	fromReader972 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader972` to the result `intoString633`
	// (`intoString633` is now tainted).
	intoString633, _ := fromReader972.ReadLine()

	// Return the tainted `intoString633`:
	return intoString633
}

func TaintStepTest_NetTextprotoReaderReadLineBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader316` into `intoByte145`.

	// Assume that `sourceCQL` has the underlying type of `fromReader316`:
	fromReader316 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader316` to the result `intoByte145`
	// (`intoByte145` is now tainted).
	intoByte145, _ := fromReader316.ReadLineBytes()

	// Return the tainted `intoByte145`:
	return intoByte145
}

func TaintStepTest_NetTextprotoReaderReadMIMEHeader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader817` into `intoMIMEHeader474`.

	// Assume that `sourceCQL` has the underlying type of `fromReader817`:
	fromReader817 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader817` to the result `intoMIMEHeader474`
	// (`intoMIMEHeader474` is now tainted).
	intoMIMEHeader474, _ := fromReader817.ReadMIMEHeader()

	// Return the tainted `intoMIMEHeader474`:
	return intoMIMEHeader474
}

func TaintStepTest_NetTextprotoReaderReadResponse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader832` into `intoString378`.

	// Assume that `sourceCQL` has the underlying type of `fromReader832`:
	fromReader832 := sourceCQL.(textproto.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader832` to the result `intoString378`
	// (`intoString378` is now tainted).
	_, intoString378, _ := fromReader832.ReadResponse(0)

	// Return the tainted `intoString378`:
	return intoString378
}

func TaintStepTest_NetTextprotoWriterDotWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser541` into `intoWriter139`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser541`:
	fromWriteCloser541 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter139` variable:
	var intoWriter139 textproto.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter139`:
	intermediateCQL := intoWriter139.DotWriter()

	// Extra step (`fromWriteCloser541` taints `intermediateCQL`, which taints `intoWriter139`:
	link(fromWriteCloser541, intermediateCQL)

	// Return the tainted `intoWriter139`:
	return intoWriter139
}

func TaintStepTest_NetTextprotoWriterPrintfLine_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString814` into `intoWriter768`.

	// Assume that `sourceCQL` has the underlying type of `fromString814`:
	fromString814 := sourceCQL.(string)

	// Declare `intoWriter768` variable:
	var intoWriter768 textproto.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString814` to the receiver `intoWriter768`
	// (`intoWriter768` is now tainted).
	intoWriter768.PrintfLine(fromString814, nil)

	// Return the tainted `intoWriter768`:
	return intoWriter768
}

func TaintStepTest_NetTextprotoWriterPrintfLine_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface468` into `intoWriter736`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface468`:
	fromInterface468 := sourceCQL.(interface{})

	// Declare `intoWriter736` variable:
	var intoWriter736 textproto.Writer

	// Call the method that transfers the taint
	// from the parameter `fromInterface468` to the receiver `intoWriter736`
	// (`intoWriter736` is now tainted).
	intoWriter736.PrintfLine("", fromInterface468)

	// Return the tainted `intoWriter736`:
	return intoWriter736
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
