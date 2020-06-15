package main

import (
	"io"
	"text/tabwriter"
)

func TaintStepTest_TextTabwriterNewWriter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter517` into `intoWriter244`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter517`:
	fromWriter517 := sourceCQL.(*tabwriter.Writer)

	// Declare `intoWriter244` variable:
	var intoWriter244 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter244`:
	intermediateCQL := tabwriter.NewWriter(intoWriter244, 0, 0, 0, 0, 0)

	// Extra step (`fromWriter517` taints `intermediateCQL`, which taints `intoWriter244`:
	link(fromWriter517, intermediateCQL)

	// Sink the tainted `intoWriter244`:
	sink(intoWriter244)
}

func TaintStepTest_TextTabwriterWriterInit_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter756` into `intoWriter271`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter756`:
	fromWriter756 := sourceCQL.(tabwriter.Writer)

	// Declare `intoWriter271` variable:
	var intoWriter271 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter756` to the argument `intoWriter271`
	// (`intoWriter271` is now tainted).
	fromWriter756.Init(intoWriter271, 0, 0, 0, 0, 0)

	// Sink the tainted `intoWriter271`:
	sink(intoWriter271)
}

func TaintStepTest_TextTabwriterWriterInit_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromWriter421` into `intoWriter683`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter421`:
	fromWriter421 := sourceCQL.(*tabwriter.Writer)

	// Declare `intoWriter683` variable:
	var intoWriter683 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL tabwriter.Writer

	// Call the method that transfers the taint
	// from the result `fromWriter421` to the parameter `intoWriter683`
	// (`intoWriter683` is now tainted).
	intermediateCQL := mediumObjCQL.Init(intoWriter683, 0, 0, 0, 0, 0)

	// Extra step (`fromWriter421` taints `intermediateCQL`, which taints `intoWriter683`:
	link(fromWriter421, intermediateCQL)

	// Sink the tainted `intoWriter683`:
	sink(intoWriter683)
}

func TaintStepTest_TextTabwriterWriterWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte198` into `intoWriter441`.

	// Assume that `sourceCQL` has the underlying type of `fromByte198`:
	fromByte198 := sourceCQL.([]byte)

	// Declare `intoWriter441` variable:
	var intoWriter441 tabwriter.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte198` to the receiver `intoWriter441`
	// (`intoWriter441` is now tainted).
	intoWriter441.Write(fromByte198)

	// Sink the tainted `intoWriter441`:
	sink(intoWriter441)
}

func RunAllTaints_TextTabwriter() {
	{
		source := newSource()
		TaintStepTest_TextTabwriterNewWriter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTabwriterWriterInit_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTabwriterWriterInit_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTabwriterWriterWrite_B0I0O0(source)
	}
}
