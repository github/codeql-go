// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"encoding/csv"
	"io"
)

func TaintStepTest_EncodingCsvNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader497` into `intoReader556`.

	// Assume that `sourceCQL` has the underlying type of `fromReader497`:
	fromReader497 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader497` to result `intoReader556`
	// (`intoReader556` is now tainted).
	intoReader556 := csv.NewReader(fromReader497)

	// Return the tainted `intoReader556`:
	return intoReader556
}

func TaintStepTest_EncodingCsvNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter397` into `intoWriter253`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter397`:
	fromWriter397 := sourceCQL.(*csv.Writer)

	// Declare `intoWriter253` variable:
	var intoWriter253 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter253`:
	intermediateCQL := csv.NewWriter(intoWriter253)

	// Extra step (`fromWriter397` taints `intermediateCQL`, which taints `intoWriter253`:
	link(fromWriter397, intermediateCQL)

	// Return the tainted `intoWriter253`:
	return intoWriter253
}

func TaintStepTest_EncodingCsvReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader757` into `intoString473`.

	// Assume that `sourceCQL` has the underlying type of `fromReader757`:
	fromReader757 := sourceCQL.(csv.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader757` to the result `intoString473`
	// (`intoString473` is now tainted).
	intoString473, _ := fromReader757.Read()

	// Return the tainted `intoString473`:
	return intoString473
}

func TaintStepTest_EncodingCsvReaderReadAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader657` into `intoString948`.

	// Assume that `sourceCQL` has the underlying type of `fromReader657`:
	fromReader657 := sourceCQL.(csv.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader657` to the result `intoString948`
	// (`intoString948` is now tainted).
	intoString948, _ := fromReader657.ReadAll()

	// Return the tainted `intoString948`:
	return intoString948
}

func TaintStepTest_EncodingCsvWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString127` into `intoWriter313`.

	// Assume that `sourceCQL` has the underlying type of `fromString127`:
	fromString127 := sourceCQL.([]string)

	// Declare `intoWriter313` variable:
	var intoWriter313 csv.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString127` to the receiver `intoWriter313`
	// (`intoWriter313` is now tainted).
	intoWriter313.Write(fromString127)

	// Return the tainted `intoWriter313`:
	return intoWriter313
}

func TaintStepTest_EncodingCsvWriterWriteAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString497` into `intoWriter788`.

	// Assume that `sourceCQL` has the underlying type of `fromString497`:
	fromString497 := sourceCQL.([][]string)

	// Declare `intoWriter788` variable:
	var intoWriter788 csv.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString497` to the receiver `intoWriter788`
	// (`intoWriter788` is now tainted).
	intoWriter788.WriteAll(fromString497)

	// Return the tainted `intoWriter788`:
	return intoWriter788
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
