package main

import (
	"io"
	"mime/multipart"
)

func TaintStepTest_MimeMultipartNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader944` into `intoReader260`.

	// Assume that `sourceCQL` has the underlying type of `fromReader944`:
	fromReader944 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader944` to result `intoReader260`
	// (`intoReader260` is now tainted).
	intoReader260 := multipart.NewReader(fromReader944, "")

	// Return the tainted `intoReader260`:
	return intoReader260
}

func TaintStepTest_MimeMultipartNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter420` into `intoWriter576`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter420`:
	fromWriter420 := sourceCQL.(*multipart.Writer)

	// Declare `intoWriter576` variable:
	var intoWriter576 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter576`:
	intermediateCQL := multipart.NewWriter(intoWriter576)

	// Extra step (`fromWriter420` taints `intermediateCQL`, which taints `intoWriter576`:
	link(fromWriter420, intermediateCQL)

	// Return the tainted `intoWriter576`:
	return intoWriter576
}

func TaintStepTest_MimeMultipartFileHeaderOpen_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFileHeader322` into `intoFile450`.

	// Assume that `sourceCQL` has the underlying type of `fromFileHeader322`:
	fromFileHeader322 := sourceCQL.(multipart.FileHeader)

	// Call the method that transfers the taint
	// from the receiver `fromFileHeader322` to the result `intoFile450`
	// (`intoFile450` is now tainted).
	intoFile450, _ := fromFileHeader322.Open()

	// Return the tainted `intoFile450`:
	return intoFile450
}

func TaintStepTest_MimeMultipartPartRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPart862` into `intoByte926`.

	// Assume that `sourceCQL` has the underlying type of `fromPart862`:
	fromPart862 := sourceCQL.(multipart.Part)

	// Declare `intoByte926` variable:
	var intoByte926 []byte

	// Call the method that transfers the taint
	// from the receiver `fromPart862` to the argument `intoByte926`
	// (`intoByte926` is now tainted).
	fromPart862.Read(intoByte926)

	// Return the tainted `intoByte926`:
	return intoByte926
}

func TaintStepTest_MimeMultipartReaderNextPart_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader924` into `intoPart702`.

	// Assume that `sourceCQL` has the underlying type of `fromReader924`:
	fromReader924 := sourceCQL.(multipart.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader924` to the result `intoPart702`
	// (`intoPart702` is now tainted).
	intoPart702, _ := fromReader924.NextPart()

	// Return the tainted `intoPart702`:
	return intoPart702
}

func TaintStepTest_MimeMultipartReaderNextRawPart_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader460` into `intoPart683`.

	// Assume that `sourceCQL` has the underlying type of `fromReader460`:
	fromReader460 := sourceCQL.(multipart.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader460` to the result `intoPart683`
	// (`intoPart683` is now tainted).
	intoPart683, _ := fromReader460.NextRawPart()

	// Return the tainted `intoPart683`:
	return intoPart683
}

func TaintStepTest_MimeMultipartReaderReadForm_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader529` into `intoForm502`.

	// Assume that `sourceCQL` has the underlying type of `fromReader529`:
	fromReader529 := sourceCQL.(multipart.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader529` to the result `intoForm502`
	// (`intoForm502` is now tainted).
	intoForm502, _ := fromReader529.ReadForm(0)

	// Return the tainted `intoForm502`:
	return intoForm502
}

func TaintStepTest_MimeMultipartWriterCreateFormField_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter304` into `intoWriter563`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter304`:
	fromWriter304 := sourceCQL.(io.Writer)

	// Declare `intoWriter563` variable:
	var intoWriter563 multipart.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter563`:
	intermediateCQL, _ := intoWriter563.CreateFormField("")

	// Extra step (`fromWriter304` taints `intermediateCQL`, which taints `intoWriter563`:
	link(fromWriter304, intermediateCQL)

	// Return the tainted `intoWriter563`:
	return intoWriter563
}

func TaintStepTest_MimeMultipartWriterCreateFormFile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter242` into `intoWriter836`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter242`:
	fromWriter242 := sourceCQL.(io.Writer)

	// Declare `intoWriter836` variable:
	var intoWriter836 multipart.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter836`:
	intermediateCQL, _ := intoWriter836.CreateFormFile("", "")

	// Extra step (`fromWriter242` taints `intermediateCQL`, which taints `intoWriter836`:
	link(fromWriter242, intermediateCQL)

	// Return the tainted `intoWriter836`:
	return intoWriter836
}

func TaintStepTest_MimeMultipartWriterCreatePart_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter920` into `intoWriter147`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter920`:
	fromWriter920 := sourceCQL.(io.Writer)

	// Declare `intoWriter147` variable:
	var intoWriter147 multipart.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter147`:
	intermediateCQL, _ := intoWriter147.CreatePart(nil)

	// Extra step (`fromWriter920` taints `intermediateCQL`, which taints `intoWriter147`:
	link(fromWriter920, intermediateCQL)

	// Return the tainted `intoWriter147`:
	return intoWriter147
}

func TaintStepTest_MimeMultipartWriterWriteField_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString693` into `intoWriter669`.

	// Assume that `sourceCQL` has the underlying type of `fromString693`:
	fromString693 := sourceCQL.(string)

	// Declare `intoWriter669` variable:
	var intoWriter669 multipart.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString693` to the receiver `intoWriter669`
	// (`intoWriter669` is now tainted).
	intoWriter669.WriteField(fromString693, "")

	// Return the tainted `intoWriter669`:
	return intoWriter669
}

func TaintStepTest_MimeMultipartWriterWriteField_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString725` into `intoWriter741`.

	// Assume that `sourceCQL` has the underlying type of `fromString725`:
	fromString725 := sourceCQL.(string)

	// Declare `intoWriter741` variable:
	var intoWriter741 multipart.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString725` to the receiver `intoWriter741`
	// (`intoWriter741` is now tainted).
	intoWriter741.WriteField("", fromString725)

	// Return the tainted `intoWriter741`:
	return intoWriter741
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
