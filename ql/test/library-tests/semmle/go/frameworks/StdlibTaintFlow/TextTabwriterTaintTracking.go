// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"text/tabwriter"
)

func TaintStepTest_TextTabwriterNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter656` into `intoWriter414`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter656`:
	fromWriter656 := sourceCQL.(*tabwriter.Writer)

	// Declare `intoWriter414` variable:
	var intoWriter414 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter414`:
	intermediateCQL := tabwriter.NewWriter(intoWriter414, 0, 0, 0, 0, 0)

	// Extra step (`fromWriter656` taints `intermediateCQL`, which taints `intoWriter414`:
	link(fromWriter656, intermediateCQL)

	// Return the tainted `intoWriter414`:
	return intoWriter414
}

func TaintStepTest_TextTabwriterWriterInit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter518` into `intoWriter650`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter518`:
	fromWriter518 := sourceCQL.(tabwriter.Writer)

	// Declare `intoWriter650` variable:
	var intoWriter650 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter518` to the argument `intoWriter650`
	// (`intoWriter650` is now tainted).
	fromWriter518.Init(intoWriter650, 0, 0, 0, 0, 0)

	// Return the tainted `intoWriter650`:
	return intoWriter650
}

func TaintStepTest_TextTabwriterWriterInit_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter784` into `intoWriter957`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter784`:
	fromWriter784 := sourceCQL.(*tabwriter.Writer)

	// Declare `intoWriter957` variable:
	var intoWriter957 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL tabwriter.Writer

	// Call the method that transfers the taint
	// from the result `fromWriter784` to the parameter `intoWriter957`
	// (`intoWriter957` is now tainted).
	intermediateCQL := mediumObjCQL.Init(intoWriter957, 0, 0, 0, 0, 0)

	// Extra step (`fromWriter784` taints `intermediateCQL`, which taints `intoWriter957`:
	link(fromWriter784, intermediateCQL)

	// Return the tainted `intoWriter957`:
	return intoWriter957
}

func TaintStepTest_TextTabwriterWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte520` into `intoWriter443`.

	// Assume that `sourceCQL` has the underlying type of `fromByte520`:
	fromByte520 := sourceCQL.([]byte)

	// Declare `intoWriter443` variable:
	var intoWriter443 tabwriter.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte520` to the receiver `intoWriter443`
	// (`intoWriter443` is now tainted).
	intoWriter443.Write(fromByte520)

	// Return the tainted `intoWriter443`:
	return intoWriter443
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
