// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"encoding/csv"
	"io"
)

func TaintStepTest_EncodingCsvNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader656` into `intoReader414`.

	// Assume that `sourceCQL` has the underlying type of `fromReader656`:
	fromReader656 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader656` to result `intoReader414`
	// (`intoReader414` is now tainted).
	intoReader414 := csv.NewReader(fromReader656)

	// Return the tainted `intoReader414`:
	return intoReader414
}

func TaintStepTest_EncodingCsvNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter518` into `intoWriter650`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter518`:
	fromWriter518 := sourceCQL.(*csv.Writer)

	// Declare `intoWriter650` variable:
	var intoWriter650 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter650`:
	intermediateCQL := csv.NewWriter(intoWriter650)

	// Extra step (`fromWriter518` taints `intermediateCQL`, which taints `intoWriter650`:
	link(fromWriter518, intermediateCQL)

	// Return the tainted `intoWriter650`:
	return intoWriter650
}

func TaintStepTest_EncodingCsvReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader784` into `intoString957`.

	// Assume that `sourceCQL` has the underlying type of `fromReader784`:
	fromReader784 := sourceCQL.(csv.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader784` to the result `intoString957`
	// (`intoString957` is now tainted).
	intoString957, _ := fromReader784.Read()

	// Return the tainted `intoString957`:
	return intoString957
}

func TaintStepTest_EncodingCsvReaderReadAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader520` into `intoString443`.

	// Assume that `sourceCQL` has the underlying type of `fromReader520`:
	fromReader520 := sourceCQL.(csv.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader520` to the result `intoString443`
	// (`intoString443` is now tainted).
	intoString443, _ := fromReader520.ReadAll()

	// Return the tainted `intoString443`:
	return intoString443
}

func TaintStepTest_EncodingCsvWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString127` into `intoWriter483`.

	// Assume that `sourceCQL` has the underlying type of `fromString127`:
	fromString127 := sourceCQL.([]string)

	// Declare `intoWriter483` variable:
	var intoWriter483 csv.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString127` to the receiver `intoWriter483`
	// (`intoWriter483` is now tainted).
	intoWriter483.Write(fromString127)

	// Return the tainted `intoWriter483`:
	return intoWriter483
}

func TaintStepTest_EncodingCsvWriterWriteAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString989` into `intoWriter982`.

	// Assume that `sourceCQL` has the underlying type of `fromString989`:
	fromString989 := sourceCQL.([][]string)

	// Declare `intoWriter982` variable:
	var intoWriter982 csv.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString989` to the receiver `intoWriter982`
	// (`intoWriter982` is now tainted).
	intoWriter982.WriteAll(fromString989)

	// Return the tainted `intoWriter982`:
	return intoWriter982
}

func RunAllTaints_EncodingCsv() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingCsvNewReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingCsvNewWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingCsvReaderRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingCsvReaderReadAll_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingCsvWriterWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingCsvWriterWriteAll_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
