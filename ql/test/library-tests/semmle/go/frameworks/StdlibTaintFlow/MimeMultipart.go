package main

import (
	"io"
	"mime/multipart"
)

func TaintStepTest_MimeMultipartNewReader(sourceCQL interface{}) {
	// The flow is from `r` into `into269`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into269`
	// (`into269` is now tainted).
	into269 := multipart.NewReader(r, "")

	// Sink the tainted `into269`:
	sink(into269)
}

func TaintStepTest_MimeMultipartNewWriter(sourceCQL interface{}) {
	// The flow is from `from905` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from905`:
	from905 := sourceCQL.(*multipart.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := multipart.NewWriter(w)

	// Extra step (`from905` taints `intermediateCQL`, which taints `w`:
	link(from905, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_MimeMultipartFileHeaderOpen(sourceCQL interface{}) {
	// The flow is from `from950` into `into889`.

	// Assume that `sourceCQL` has the underlying type of `from950`:
	from950 := sourceCQL.(multipart.FileHeader)

	// Call the method that transfers the taint
	// from the receiver `from950` to the result `into889`
	// (`into889` is now tainted).
	into889, _ := from950.Open()

	// Sink the tainted `into889`:
	sink(into889)
}

func TaintStepTest_MimeMultipartPartRead(sourceCQL interface{}) {
	// The flow is from `from226` into `d`.

	// Assume that `sourceCQL` has the underlying type of `from226`:
	from226 := sourceCQL.(multipart.Part)

	// Declare `d` variable:
	var d []byte

	// Call the method that transfers the taint
	// from the receiver `from226` to the argument `d`
	// (`d` is now tainted).
	from226.Read(d)

	// Sink the tainted `d`:
	sink(d)
}

func TaintStepTest_MimeMultipartReaderNextPart(sourceCQL interface{}) {
	// The flow is from `from962` into `into116`.

	// Assume that `sourceCQL` has the underlying type of `from962`:
	from962 := sourceCQL.(multipart.Reader)

	// Call the method that transfers the taint
	// from the receiver `from962` to the result `into116`
	// (`into116` is now tainted).
	into116, _ := from962.NextPart()

	// Sink the tainted `into116`:
	sink(into116)
}

func TaintStepTest_MimeMultipartReaderNextRawPart(sourceCQL interface{}) {
	// The flow is from `from495` into `into194`.

	// Assume that `sourceCQL` has the underlying type of `from495`:
	from495 := sourceCQL.(multipart.Reader)

	// Call the method that transfers the taint
	// from the receiver `from495` to the result `into194`
	// (`into194` is now tainted).
	into194, _ := from495.NextRawPart()

	// Sink the tainted `into194`:
	sink(into194)
}

func TaintStepTest_MimeMultipartReaderReadForm(sourceCQL interface{}) {
	// The flow is from `from247` into `into720`.

	// Assume that `sourceCQL` has the underlying type of `from247`:
	from247 := sourceCQL.(multipart.Reader)

	// Call the method that transfers the taint
	// from the receiver `from247` to the result `into720`
	// (`into720` is now tainted).
	into720, _ := from247.ReadForm(0)

	// Sink the tainted `into720`:
	sink(into720)
}

func TaintStepTest_MimeMultipartWriterCreateFormField(sourceCQL interface{}) {
	// The flow is from `from546` into `into174`.

	// Assume that `sourceCQL` has the underlying type of `from546`:
	from546 := sourceCQL.(io.Writer)

	// Declare `into174` variable:
	var into174 multipart.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `into174`:
	intermediateCQL, _ := into174.CreateFormField("")

	// Extra step (`from546` taints `intermediateCQL`, which taints `into174`:
	link(from546, intermediateCQL)

	// Sink the tainted `into174`:
	sink(into174)
}

func TaintStepTest_MimeMultipartWriterCreateFormFile(sourceCQL interface{}) {
	// The flow is from `from209` into `into952`.

	// Assume that `sourceCQL` has the underlying type of `from209`:
	from209 := sourceCQL.(io.Writer)

	// Declare `into952` variable:
	var into952 multipart.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `into952`:
	intermediateCQL, _ := into952.CreateFormFile("", "")

	// Extra step (`from209` taints `intermediateCQL`, which taints `into952`:
	link(from209, intermediateCQL)

	// Sink the tainted `into952`:
	sink(into952)
}

func TaintStepTest_MimeMultipartWriterCreatePart(sourceCQL interface{}) {
	// The flow is from `from889` into `into274`.

	// Assume that `sourceCQL` has the underlying type of `from889`:
	from889 := sourceCQL.(io.Writer)

	// Declare `into274` variable:
	var into274 multipart.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `into274`:
	intermediateCQL, _ := into274.CreatePart(nil)

	// Extra step (`from889` taints `intermediateCQL`, which taints `into274`:
	link(from889, intermediateCQL)

	// Sink the tainted `into274`:
	sink(into274)
}

func TaintStepTest_MimeMultipartWriterWriteField(sourceCQL interface{}) {
	// The flow is from `value` into `into235`.

	// Assume that `sourceCQL` has the underlying type of `value`:
	value := sourceCQL.(string)

	// Declare `into235` variable:
	var into235 multipart.Writer

	// Call the method that transfers the taint
	// from the parameter `value` to the receiver `into235`
	// (`into235` is now tainted).
	into235.WriteField("", value)

	// Sink the tainted `into235`:
	sink(into235)
}

func RunAllTaints_MimeMultipart(v interface{}) {
	{
		source := newSource()
		TaintStepTest_MimeMultipartNewReader(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartNewWriter(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartFileHeaderOpen(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartPartRead(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartReaderNextPart(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartReaderNextRawPart(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartReaderReadForm(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartWriterCreateFormField(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartWriterCreateFormFile(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartWriterCreatePart(source)
	}
	{
		source := newSource()
		TaintStepTest_MimeMultipartWriterWriteField(source)
	}
}
