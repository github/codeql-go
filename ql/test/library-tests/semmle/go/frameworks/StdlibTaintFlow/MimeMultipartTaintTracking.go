package main

import (
	"io"
	"mime/multipart"
)

func TaintStepTest_MimeMultipartNewReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader460` into `intoReader877`.

	// Assume that `sourceCQL` has the underlying type of `fromReader460`:
	fromReader460 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader460` to result `intoReader877`
	// (`intoReader877` is now tainted).
	intoReader877 := multipart.NewReader(fromReader460, "")

	// Sink the tainted `intoReader877`:
	sink(intoReader877)
}

func TaintStepTest_MimeMultipartNewWriter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter413` into `intoWriter186`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter413`:
	fromWriter413 := sourceCQL.(*multipart.Writer)

	// Declare `intoWriter186` variable:
	var intoWriter186 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter186`:
	intermediateCQL := multipart.NewWriter(intoWriter186)

	// Extra step (`fromWriter413` taints `intermediateCQL`, which taints `intoWriter186`:
	link(fromWriter413, intermediateCQL)

	// Sink the tainted `intoWriter186`:
	sink(intoWriter186)
}

func TaintStepTest_MimeMultipartFileHeaderOpen_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFileHeader128` into `intoFile858`.

	// Assume that `sourceCQL` has the underlying type of `fromFileHeader128`:
	fromFileHeader128 := sourceCQL.(multipart.FileHeader)

	// Call the method that transfers the taint
	// from the receiver `fromFileHeader128` to the result `intoFile858`
	// (`intoFile858` is now tainted).
	intoFile858, _ := fromFileHeader128.Open()

	// Sink the tainted `intoFile858`:
	sink(intoFile858)
}

func TaintStepTest_MimeMultipartPartRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromPart502` into `intoByte613`.

	// Assume that `sourceCQL` has the underlying type of `fromPart502`:
	fromPart502 := sourceCQL.(multipart.Part)

	// Declare `intoByte613` variable:
	var intoByte613 []byte

	// Call the method that transfers the taint
	// from the receiver `fromPart502` to the argument `intoByte613`
	// (`intoByte613` is now tainted).
	fromPart502.Read(intoByte613)

	// Sink the tainted `intoByte613`:
	sink(intoByte613)
}

func TaintStepTest_MimeMultipartReaderNextPart_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader613` into `intoPart315`.

	// Assume that `sourceCQL` has the underlying type of `fromReader613`:
	fromReader613 := sourceCQL.(multipart.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader613` to the result `intoPart315`
	// (`intoPart315` is now tainted).
	intoPart315, _ := fromReader613.NextPart()

	// Sink the tainted `intoPart315`:
	sink(intoPart315)
}

func TaintStepTest_MimeMultipartReaderNextRawPart_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader806` into `intoPart705`.

	// Assume that `sourceCQL` has the underlying type of `fromReader806`:
	fromReader806 := sourceCQL.(multipart.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader806` to the result `intoPart705`
	// (`intoPart705` is now tainted).
	intoPart705, _ := fromReader806.NextRawPart()

	// Sink the tainted `intoPart705`:
	sink(intoPart705)
}

func TaintStepTest_MimeMultipartReaderReadForm_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader732` into `intoForm472`.

	// Assume that `sourceCQL` has the underlying type of `fromReader732`:
	fromReader732 := sourceCQL.(multipart.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader732` to the result `intoForm472`
	// (`intoForm472` is now tainted).
	intoForm472, _ := fromReader732.ReadForm(0)

	// Sink the tainted `intoForm472`:
	sink(intoForm472)
}

func TaintStepTest_MimeMultipartWriterCreateFormField_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter322` into `intoWriter642`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter322`:
	fromWriter322 := sourceCQL.(io.Writer)

	// Declare `intoWriter642` variable:
	var intoWriter642 multipart.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter642`:
	intermediateCQL, _ := intoWriter642.CreateFormField("")

	// Extra step (`fromWriter322` taints `intermediateCQL`, which taints `intoWriter642`:
	link(fromWriter322, intermediateCQL)

	// Sink the tainted `intoWriter642`:
	sink(intoWriter642)
}

func TaintStepTest_MimeMultipartWriterCreateFormFile_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter923` into `intoWriter539`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter923`:
	fromWriter923 := sourceCQL.(io.Writer)

	// Declare `intoWriter539` variable:
	var intoWriter539 multipart.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter539`:
	intermediateCQL, _ := intoWriter539.CreateFormFile("", "")

	// Extra step (`fromWriter923` taints `intermediateCQL`, which taints `intoWriter539`:
	link(fromWriter923, intermediateCQL)

	// Sink the tainted `intoWriter539`:
	sink(intoWriter539)
}

func TaintStepTest_MimeMultipartWriterCreatePart_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter766` into `intoWriter583`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter766`:
	fromWriter766 := sourceCQL.(io.Writer)

	// Declare `intoWriter583` variable:
	var intoWriter583 multipart.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter583`:
	intermediateCQL, _ := intoWriter583.CreatePart(nil)

	// Extra step (`fromWriter766` taints `intermediateCQL`, which taints `intoWriter583`:
	link(fromWriter766, intermediateCQL)

	// Sink the tainted `intoWriter583`:
	sink(intoWriter583)
}

func TaintStepTest_MimeMultipartWriterWriteField_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString191` into `intoWriter217`.

	// Assume that `sourceCQL` has the underlying type of `fromString191`:
	fromString191 := sourceCQL.(string)

	// Declare `intoWriter217` variable:
	var intoWriter217 multipart.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString191` to the receiver `intoWriter217`
	// (`intoWriter217` is now tainted).
	intoWriter217.WriteField(fromString191, "")

	// Sink the tainted `intoWriter217`:
	sink(intoWriter217)
}

func TaintStepTest_MimeMultipartWriterWriteField_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString157` into `intoWriter152`.

	// Assume that `sourceCQL` has the underlying type of `fromString157`:
	fromString157 := sourceCQL.(string)

	// Declare `intoWriter152` variable:
	var intoWriter152 multipart.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString157` to the receiver `intoWriter152`
	// (`intoWriter152` is now tainted).
	intoWriter152.WriteField("", fromString157)

	// Sink the tainted `intoWriter152`:
	sink(intoWriter152)
}

func RunAllTaints_MimeMultipart() {
	{
		source := newSource()
		TaintStepTest_MimeMultipartNewReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartNewWriter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartFileHeaderOpen_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartPartRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartReaderNextPart_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartReaderNextRawPart_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartReaderReadForm_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartWriterCreateFormField_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartWriterCreateFormFile_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartWriterCreatePart_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartWriterWriteField_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartWriterWriteField_B0I1O0(source)
	}
}
