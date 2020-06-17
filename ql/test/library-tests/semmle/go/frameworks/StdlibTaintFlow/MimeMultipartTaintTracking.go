// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"mime/multipart"
)

func TaintStepTest_MimeMultipartNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader656` into `intoReader414`.

	// Assume that `sourceCQL` has the underlying type of `fromReader656`:
	fromReader656 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader656` to result `intoReader414`
	// (`intoReader414` is now tainted).
	intoReader414 := multipart.NewReader(fromReader656, "")

	// Return the tainted `intoReader414`:
	return intoReader414
}

func TaintStepTest_MimeMultipartNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter518` into `intoWriter650`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter518`:
	fromWriter518 := sourceCQL.(*multipart.Writer)

	// Declare `intoWriter650` variable:
	var intoWriter650 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter650`:
	intermediateCQL := multipart.NewWriter(intoWriter650)

	// Extra step (`fromWriter518` taints `intermediateCQL`, which taints `intoWriter650`:
	link(fromWriter518, intermediateCQL)

	// Return the tainted `intoWriter650`:
	return intoWriter650
}

func TaintStepTest_MimeMultipartFileHeaderOpen_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFileHeader784` into `intoFile957`.

	// Assume that `sourceCQL` has the underlying type of `fromFileHeader784`:
	fromFileHeader784 := sourceCQL.(multipart.FileHeader)

	// Call the method that transfers the taint
	// from the receiver `fromFileHeader784` to the result `intoFile957`
	// (`intoFile957` is now tainted).
	intoFile957, _ := fromFileHeader784.Open()

	// Return the tainted `intoFile957`:
	return intoFile957
}

func TaintStepTest_MimeMultipartPartRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPart520` into `intoByte443`.

	// Assume that `sourceCQL` has the underlying type of `fromPart520`:
	fromPart520 := sourceCQL.(multipart.Part)

	// Declare `intoByte443` variable:
	var intoByte443 []byte

	// Call the method that transfers the taint
	// from the receiver `fromPart520` to the argument `intoByte443`
	// (`intoByte443` is now tainted).
	fromPart520.Read(intoByte443)

	// Return the tainted `intoByte443`:
	return intoByte443
}

func TaintStepTest_MimeMultipartReaderNextPart_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader127` into `intoPart483`.

	// Assume that `sourceCQL` has the underlying type of `fromReader127`:
	fromReader127 := sourceCQL.(multipart.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader127` to the result `intoPart483`
	// (`intoPart483` is now tainted).
	intoPart483, _ := fromReader127.NextPart()

	// Return the tainted `intoPart483`:
	return intoPart483
}

func TaintStepTest_MimeMultipartReaderNextRawPart_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader989` into `intoPart982`.

	// Assume that `sourceCQL` has the underlying type of `fromReader989`:
	fromReader989 := sourceCQL.(multipart.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader989` to the result `intoPart982`
	// (`intoPart982` is now tainted).
	intoPart982, _ := fromReader989.NextRawPart()

	// Return the tainted `intoPart982`:
	return intoPart982
}

func TaintStepTest_MimeMultipartReaderReadForm_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader417` into `intoForm584`.

	// Assume that `sourceCQL` has the underlying type of `fromReader417`:
	fromReader417 := sourceCQL.(multipart.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader417` to the result `intoForm584`
	// (`intoForm584` is now tainted).
	intoForm584, _ := fromReader417.ReadForm(0)

	// Return the tainted `intoForm584`:
	return intoForm584
}

func TaintStepTest_MimeMultipartWriterCreateFormField_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter991` into `intoWriter881`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter991`:
	fromWriter991 := sourceCQL.(io.Writer)

	// Declare `intoWriter881` variable:
	var intoWriter881 multipart.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter881`:
	intermediateCQL, _ := intoWriter881.CreateFormField("")

	// Extra step (`fromWriter991` taints `intermediateCQL`, which taints `intoWriter881`:
	link(fromWriter991, intermediateCQL)

	// Return the tainted `intoWriter881`:
	return intoWriter881
}

func TaintStepTest_MimeMultipartWriterCreateFormFile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter186` into `intoWriter284`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter186`:
	fromWriter186 := sourceCQL.(io.Writer)

	// Declare `intoWriter284` variable:
	var intoWriter284 multipart.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter284`:
	intermediateCQL, _ := intoWriter284.CreateFormFile("", "")

	// Extra step (`fromWriter186` taints `intermediateCQL`, which taints `intoWriter284`:
	link(fromWriter186, intermediateCQL)

	// Return the tainted `intoWriter284`:
	return intoWriter284
}

func TaintStepTest_MimeMultipartWriterCreatePart_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter908` into `intoWriter137`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter908`:
	fromWriter908 := sourceCQL.(io.Writer)

	// Declare `intoWriter137` variable:
	var intoWriter137 multipart.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter137`:
	intermediateCQL, _ := intoWriter137.CreatePart(nil)

	// Extra step (`fromWriter908` taints `intermediateCQL`, which taints `intoWriter137`:
	link(fromWriter908, intermediateCQL)

	// Return the tainted `intoWriter137`:
	return intoWriter137
}

func TaintStepTest_MimeMultipartWriterWriteField_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString494` into `intoWriter873`.

	// Assume that `sourceCQL` has the underlying type of `fromString494`:
	fromString494 := sourceCQL.(string)

	// Declare `intoWriter873` variable:
	var intoWriter873 multipart.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString494` to the receiver `intoWriter873`
	// (`intoWriter873` is now tainted).
	intoWriter873.WriteField(fromString494, "")

	// Return the tainted `intoWriter873`:
	return intoWriter873
}

func TaintStepTest_MimeMultipartWriterWriteField_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString599` into `intoWriter409`.

	// Assume that `sourceCQL` has the underlying type of `fromString599`:
	fromString599 := sourceCQL.(string)

	// Declare `intoWriter409` variable:
	var intoWriter409 multipart.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString599` to the receiver `intoWriter409`
	// (`intoWriter409` is now tainted).
	intoWriter409.WriteField("", fromString599)

	// Return the tainted `intoWriter409`:
	return intoWriter409
}

func RunAllTaints_MimeMultipart() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeMultipartNewReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeMultipartNewWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeMultipartFileHeaderOpen_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeMultipartPartRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeMultipartReaderNextPart_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeMultipartReaderNextRawPart_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeMultipartReaderReadForm_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeMultipartWriterCreateFormField_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeMultipartWriterCreateFormFile_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeMultipartWriterCreatePart_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeMultipartWriterWriteField_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_MimeMultipartWriterWriteField_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
