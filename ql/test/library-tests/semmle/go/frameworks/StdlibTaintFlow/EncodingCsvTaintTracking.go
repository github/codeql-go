package main

import (
	"encoding/csv"
	"io"
)

func TaintStepTest_EncodingCsvNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader134` into `intoReader222`.

	// Assume that `sourceCQL` has the underlying type of `fromReader134`:
	fromReader134 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader134` to result `intoReader222`
	// (`intoReader222` is now tainted).
	intoReader222 := csv.NewReader(fromReader134)

	// Return the tainted `intoReader222`:
	return intoReader222
}

func TaintStepTest_EncodingCsvNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter794` into `intoWriter666`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter794`:
	fromWriter794 := sourceCQL.(*csv.Writer)

	// Declare `intoWriter666` variable:
	var intoWriter666 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter666`:
	intermediateCQL := csv.NewWriter(intoWriter666)

	// Extra step (`fromWriter794` taints `intermediateCQL`, which taints `intoWriter666`:
	link(fromWriter794, intermediateCQL)

	// Return the tainted `intoWriter666`:
	return intoWriter666
}

func TaintStepTest_EncodingCsvReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader603` into `intoString867`.

	// Assume that `sourceCQL` has the underlying type of `fromReader603`:
	fromReader603 := sourceCQL.(csv.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader603` to the result `intoString867`
	// (`intoString867` is now tainted).
	intoString867, _ := fromReader603.Read()

	// Return the tainted `intoString867`:
	return intoString867
}

func TaintStepTest_EncodingCsvReaderReadAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader545` into `intoString614`.

	// Assume that `sourceCQL` has the underlying type of `fromReader545`:
	fromReader545 := sourceCQL.(csv.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader545` to the result `intoString614`
	// (`intoString614` is now tainted).
	intoString614, _ := fromReader545.ReadAll()

	// Return the tainted `intoString614`:
	return intoString614
}

func TaintStepTest_EncodingCsvWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString712` into `intoWriter772`.

	// Assume that `sourceCQL` has the underlying type of `fromString712`:
	fromString712 := sourceCQL.([]string)

	// Declare `intoWriter772` variable:
	var intoWriter772 csv.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString712` to the receiver `intoWriter772`
	// (`intoWriter772` is now tainted).
	intoWriter772.Write(fromString712)

	// Return the tainted `intoWriter772`:
	return intoWriter772
}

func TaintStepTest_EncodingCsvWriterWriteAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString917` into `intoWriter489`.

	// Assume that `sourceCQL` has the underlying type of `fromString917`:
	fromString917 := sourceCQL.([][]string)

	// Declare `intoWriter489` variable:
	var intoWriter489 csv.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString917` to the receiver `intoWriter489`
	// (`intoWriter489` is now tainted).
	intoWriter489.WriteAll(fromString917)

	// Return the tainted `intoWriter489`:
	return intoWriter489
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
