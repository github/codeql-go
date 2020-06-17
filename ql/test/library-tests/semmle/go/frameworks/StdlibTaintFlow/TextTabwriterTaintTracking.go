// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"text/tabwriter"
)

func TaintStepTest_TextTabwriterNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter584` into `intoWriter217`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter584`:
	fromWriter584 := sourceCQL.(*tabwriter.Writer)

	// Declare `intoWriter217` variable:
	var intoWriter217 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter217`:
	intermediateCQL := tabwriter.NewWriter(intoWriter217, 0, 0, 0, 0, 0)

	// Extra step (`fromWriter584` taints `intermediateCQL`, which taints `intoWriter217`:
	link(fromWriter584, intermediateCQL)

	// Return the tainted `intoWriter217`:
	return intoWriter217
}

func TaintStepTest_TextTabwriterWriterInit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter475` into `intoWriter939`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter475`:
	fromWriter475 := sourceCQL.(tabwriter.Writer)

	// Declare `intoWriter939` variable:
	var intoWriter939 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter475` to the argument `intoWriter939`
	// (`intoWriter939` is now tainted).
	fromWriter475.Init(intoWriter939, 0, 0, 0, 0, 0)

	// Return the tainted `intoWriter939`:
	return intoWriter939
}

func TaintStepTest_TextTabwriterWriterInit_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter425` into `intoWriter476`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter425`:
	fromWriter425 := sourceCQL.(*tabwriter.Writer)

	// Declare `intoWriter476` variable:
	var intoWriter476 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL tabwriter.Writer

	// Call the method that transfers the taint
	// from the result `fromWriter425` to the parameter `intoWriter476`
	// (`intoWriter476` is now tainted).
	intermediateCQL := mediumObjCQL.Init(intoWriter476, 0, 0, 0, 0, 0)

	// Extra step (`fromWriter425` taints `intermediateCQL`, which taints `intoWriter476`:
	link(fromWriter425, intermediateCQL)

	// Return the tainted `intoWriter476`:
	return intoWriter476
}

func TaintStepTest_TextTabwriterWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte568` into `intoWriter118`.

	// Assume that `sourceCQL` has the underlying type of `fromByte568`:
	fromByte568 := sourceCQL.([]byte)

	// Declare `intoWriter118` variable:
	var intoWriter118 tabwriter.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte568` to the receiver `intoWriter118`
	// (`intoWriter118` is now tainted).
	intoWriter118.Write(fromByte568)

	// Return the tainted `intoWriter118`:
	return intoWriter118
}

func RunAllTaints_TextTabwriter() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTabwriterNewWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTabwriterWriterInit_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTabwriterWriterInit_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTabwriterWriterWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
