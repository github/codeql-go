package main

import (
	"encoding/csv"
	"io"
)

func TaintStepTest_EncodingCsvNewReader(sourceCQL interface{}) {
	// The flow is from `r` into `into364`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into364`
	// (`into364` is now tainted).
	into364 := csv.NewReader(r)

	// Sink the tainted `into364`:
	sink(into364)
}

func TaintStepTest_EncodingCsvNewWriter(sourceCQL interface{}) {
	// The flow is from `from563` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from563`:
	from563 := sourceCQL.(*csv.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := csv.NewWriter(w)

	// Extra step (`from563` taints `intermediateCQL`, which taints `w`:
	link(from563, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_EncodingCsvReaderRead(sourceCQL interface{}) {
	// The flow is from `from282` into `record`.

	// Assume that `sourceCQL` has the underlying type of `from282`:
	from282 := sourceCQL.(csv.Reader)

	// Call the method that transfers the taint
	// from the receiver `from282` to the result `record`
	// (`record` is now tainted).
	record, _ := from282.Read()

	// Sink the tainted `record`:
	sink(record)
}

func TaintStepTest_EncodingCsvReaderReadAll(sourceCQL interface{}) {
	// The flow is from `from164` into `records`.

	// Assume that `sourceCQL` has the underlying type of `from164`:
	from164 := sourceCQL.(csv.Reader)

	// Call the method that transfers the taint
	// from the receiver `from164` to the result `records`
	// (`records` is now tainted).
	records, _ := from164.ReadAll()

	// Sink the tainted `records`:
	sink(records)
}

func TaintStepTest_EncodingCsvWriterWrite(sourceCQL interface{}) {
	// The flow is from `record` into `into801`.

	// Assume that `sourceCQL` has the underlying type of `record`:
	record := sourceCQL.([]string)

	// Declare `into801` variable:
	var into801 csv.Writer

	// Call the method that transfers the taint
	// from the parameter `record` to the receiver `into801`
	// (`into801` is now tainted).
	into801.Write(record)

	// Sink the tainted `into801`:
	sink(into801)
}

func TaintStepTest_EncodingCsvWriterWriteAll(sourceCQL interface{}) {
	// The flow is from `records` into `into725`.

	// Assume that `sourceCQL` has the underlying type of `records`:
	records := sourceCQL.([][]string)

	// Declare `into725` variable:
	var into725 csv.Writer

	// Call the method that transfers the taint
	// from the parameter `records` to the receiver `into725`
	// (`into725` is now tainted).
	into725.WriteAll(records)

	// Sink the tainted `into725`:
	sink(into725)
}

func RunAllTaints_EncodingCsv(v interface{}) {
	{
		source := newSource()
		TaintStepTest_EncodingCsvNewReader(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingCsvNewWriter(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingCsvReaderRead(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingCsvReaderReadAll(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingCsvWriterWrite(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingCsvWriterWriteAll(source)
	}
}
