package main

import (
	"io"
	"text/tabwriter"
)

func TaintStepTest_TextTabwriterNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter270` into `intoWriter143`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter270`:
	fromWriter270 := sourceCQL.(*tabwriter.Writer)

	// Declare `intoWriter143` variable:
	var intoWriter143 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter143`:
	intermediateCQL := tabwriter.NewWriter(intoWriter143, 0, 0, 0, 0, 0)

	// Extra step (`fromWriter270` taints `intermediateCQL`, which taints `intoWriter143`:
	link(fromWriter270, intermediateCQL)

	// Return the tainted `intoWriter143`:
	return intoWriter143
}

func TaintStepTest_TextTabwriterWriterInit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter424` into `intoWriter591`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter424`:
	fromWriter424 := sourceCQL.(tabwriter.Writer)

	// Declare `intoWriter591` variable:
	var intoWriter591 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter424` to the argument `intoWriter591`
	// (`intoWriter591` is now tainted).
	fromWriter424.Init(intoWriter591, 0, 0, 0, 0, 0)

	// Return the tainted `intoWriter591`:
	return intoWriter591
}

func TaintStepTest_TextTabwriterWriterInit_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter525` into `intoWriter956`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter525`:
	fromWriter525 := sourceCQL.(*tabwriter.Writer)

	// Declare `intoWriter956` variable:
	var intoWriter956 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL tabwriter.Writer

	// Call the method that transfers the taint
	// from the result `fromWriter525` to the parameter `intoWriter956`
	// (`intoWriter956` is now tainted).
	intermediateCQL := mediumObjCQL.Init(intoWriter956, 0, 0, 0, 0, 0)

	// Extra step (`fromWriter525` taints `intermediateCQL`, which taints `intoWriter956`:
	link(fromWriter525, intermediateCQL)

	// Return the tainted `intoWriter956`:
	return intoWriter956
}

func TaintStepTest_TextTabwriterWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte475` into `intoWriter665`.

	// Assume that `sourceCQL` has the underlying type of `fromByte475`:
	fromByte475 := sourceCQL.([]byte)

	// Declare `intoWriter665` variable:
	var intoWriter665 tabwriter.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte475` to the receiver `intoWriter665`
	// (`intoWriter665` is now tainted).
	intoWriter665.Write(fromByte475)

	// Return the tainted `intoWriter665`:
	return intoWriter665
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
