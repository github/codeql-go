// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"mime/multipart"
)

func TaintStepTest_MimeMultipartNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader893` into `intoReader251`.

	// Assume that `sourceCQL` has the underlying type of `fromReader893`:
	fromReader893 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader893` to result `intoReader251`
	// (`intoReader251` is now tainted).
	intoReader251 := multipart.NewReader(fromReader893, "")

	// Return the tainted `intoReader251`:
	return intoReader251
}

func TaintStepTest_MimeMultipartNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter561` into `intoWriter269`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter561`:
	fromWriter561 := sourceCQL.(*multipart.Writer)

	// Declare `intoWriter269` variable:
	var intoWriter269 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter269`:
	intermediateCQL := multipart.NewWriter(intoWriter269)

	// Extra step (`fromWriter561` taints `intermediateCQL`, which taints `intoWriter269`:
	link(fromWriter561, intermediateCQL)

	// Return the tainted `intoWriter269`:
	return intoWriter269
}

func TaintStepTest_MimeMultipartFileHeaderOpen_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFileHeader432` into `intoFile667`.

	// Assume that `sourceCQL` has the underlying type of `fromFileHeader432`:
	fromFileHeader432 := sourceCQL.(multipart.FileHeader)

	// Call the method that transfers the taint
	// from the receiver `fromFileHeader432` to the result `intoFile667`
	// (`intoFile667` is now tainted).
	intoFile667, _ := fromFileHeader432.Open()

	// Return the tainted `intoFile667`:
	return intoFile667
}

func TaintStepTest_MimeMultipartPartRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPart911` into `intoByte522`.

	// Assume that `sourceCQL` has the underlying type of `fromPart911`:
	fromPart911 := sourceCQL.(multipart.Part)

	// Declare `intoByte522` variable:
	var intoByte522 []byte

	// Call the method that transfers the taint
	// from the receiver `fromPart911` to the argument `intoByte522`
	// (`intoByte522` is now tainted).
	fromPart911.Read(intoByte522)

	// Return the tainted `intoByte522`:
	return intoByte522
}

func TaintStepTest_MimeMultipartReaderNextPart_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader130` into `intoPart706`.

	// Assume that `sourceCQL` has the underlying type of `fromReader130`:
	fromReader130 := sourceCQL.(multipart.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader130` to the result `intoPart706`
	// (`intoPart706` is now tainted).
	intoPart706, _ := fromReader130.NextPart()

	// Return the tainted `intoPart706`:
	return intoPart706
}

func TaintStepTest_MimeMultipartReaderNextRawPart_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader749` into `intoPart595`.

	// Assume that `sourceCQL` has the underlying type of `fromReader749`:
	fromReader749 := sourceCQL.(multipart.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader749` to the result `intoPart595`
	// (`intoPart595` is now tainted).
	intoPart595, _ := fromReader749.NextRawPart()

	// Return the tainted `intoPart595`:
	return intoPart595
}

func TaintStepTest_MimeMultipartReaderReadForm_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader204` into `intoForm718`.

	// Assume that `sourceCQL` has the underlying type of `fromReader204`:
	fromReader204 := sourceCQL.(multipart.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader204` to the result `intoForm718`
	// (`intoForm718` is now tainted).
	intoForm718, _ := fromReader204.ReadForm(0)

	// Return the tainted `intoForm718`:
	return intoForm718
}

func TaintStepTest_MimeMultipartWriterCreateFormField_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter363` into `intoWriter596`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter363`:
	fromWriter363 := sourceCQL.(io.Writer)

	// Declare `intoWriter596` variable:
	var intoWriter596 multipart.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter596`:
	intermediateCQL, _ := intoWriter596.CreateFormField("")

	// Extra step (`fromWriter363` taints `intermediateCQL`, which taints `intoWriter596`:
	link(fromWriter363, intermediateCQL)

	// Return the tainted `intoWriter596`:
	return intoWriter596
}

func TaintStepTest_MimeMultipartWriterCreateFormFile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter576` into `intoWriter599`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter576`:
	fromWriter576 := sourceCQL.(io.Writer)

	// Declare `intoWriter599` variable:
	var intoWriter599 multipart.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter599`:
	intermediateCQL, _ := intoWriter599.CreateFormFile("", "")

	// Extra step (`fromWriter576` taints `intermediateCQL`, which taints `intoWriter599`:
	link(fromWriter576, intermediateCQL)

	// Return the tainted `intoWriter599`:
	return intoWriter599
}

func TaintStepTest_MimeMultipartWriterCreatePart_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter634` into `intoWriter548`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter634`:
	fromWriter634 := sourceCQL.(io.Writer)

	// Declare `intoWriter548` variable:
	var intoWriter548 multipart.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter548`:
	intermediateCQL, _ := intoWriter548.CreatePart(nil)

	// Extra step (`fromWriter634` taints `intermediateCQL`, which taints `intoWriter548`:
	link(fromWriter634, intermediateCQL)

	// Return the tainted `intoWriter548`:
	return intoWriter548
}

func TaintStepTest_MimeMultipartWriterWriteField_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString480` into `intoWriter370`.

	// Assume that `sourceCQL` has the underlying type of `fromString480`:
	fromString480 := sourceCQL.(string)

	// Declare `intoWriter370` variable:
	var intoWriter370 multipart.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString480` to the receiver `intoWriter370`
	// (`intoWriter370` is now tainted).
	intoWriter370.WriteField(fromString480, "")

	// Return the tainted `intoWriter370`:
	return intoWriter370
}

func TaintStepTest_MimeMultipartWriterWriteField_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString711` into `intoWriter501`.

	// Assume that `sourceCQL` has the underlying type of `fromString711`:
	fromString711 := sourceCQL.(string)

	// Declare `intoWriter501` variable:
	var intoWriter501 multipart.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString711` to the receiver `intoWriter501`
	// (`intoWriter501` is now tainted).
	intoWriter501.WriteField("", fromString711)

	// Return the tainted `intoWriter501`:
	return intoWriter501
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
