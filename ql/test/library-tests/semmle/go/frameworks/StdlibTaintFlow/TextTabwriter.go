package main

import (
	"io"
	"text/tabwriter"
)

func TaintStepTest_TextTabwriterNewWriter(sourceCQL interface{}) {
	// The flow is from `from415` into `output`.

	// Assume that `sourceCQL` has the underlying type of `from415`:
	from415 := sourceCQL.(*tabwriter.Writer)

	// Declare `output` variable:
	var output io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `output`:
	intermediateCQL := tabwriter.NewWriter(output, 0, 0, 0, 0, 0)

	// Extra step (`from415` taints `intermediateCQL`, which taints `output`:
	link(from415, intermediateCQL)

	// Sink the tainted `output`:
	sink(output)
}

func TaintStepTest_TextTabwriterWriterInit(sourceCQL interface{}) {
	// The flow is from `from501` into `output`.

	// Assume that `sourceCQL` has the underlying type of `from501`:
	from501 := sourceCQL.(*tabwriter.Writer)

	// Declare `output` variable:
	var output io.Writer

	// Declare medium object/interface:
	var mediumObjCQL tabwriter.Writer

	// Call the method that transfers the taint
	// from the result `from501` to the parameter `output`
	// (`output` is now tainted).
	intermediateCQL := mediumObjCQL.Init(output, 0, 0, 0, 0, 0)

	// Extra step (`from501` taints `intermediateCQL`, which taints `output`:
	link(from501, intermediateCQL)

	// Sink the tainted `output`:
	sink(output)
}

func TaintStepTest_TextTabwriterWriterWrite(sourceCQL interface{}) {
	// The flow is from `buf` into `into308`.

	// Assume that `sourceCQL` has the underlying type of `buf`:
	buf := sourceCQL.([]byte)

	// Declare `into308` variable:
	var into308 tabwriter.Writer

	// Call the method that transfers the taint
	// from the parameter `buf` to the receiver `into308`
	// (`into308` is now tainted).
	into308.Write(buf)

	// Sink the tainted `into308`:
	sink(into308)
}

func RunAllTaints_TextTabwriter(v interface{}) {
	{
		source := newSource()
		TaintStepTest_TextTabwriterNewWriter(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTabwriterWriterInit(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTabwriterWriterWrite(source)
	}
}
