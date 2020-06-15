package main

import (
	"encoding/csv"
	"io"
)

func TaintStepTest_EncodingCsvNewReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader211` into `intoReader232`.

	// Assume that `sourceCQL` has the underlying type of `fromReader211`:
	fromReader211 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader211` to result `intoReader232`
	// (`intoReader232` is now tainted).
	intoReader232 := csv.NewReader(fromReader211)

	// Sink the tainted `intoReader232`:
	sink(intoReader232)
}

func TaintStepTest_EncodingCsvNewWriter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter724` into `intoWriter514`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter724`:
	fromWriter724 := sourceCQL.(*csv.Writer)

	// Declare `intoWriter514` variable:
	var intoWriter514 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter514`:
	intermediateCQL := csv.NewWriter(intoWriter514)

	// Extra step (`fromWriter724` taints `intermediateCQL`, which taints `intoWriter514`:
	link(fromWriter724, intermediateCQL)

	// Sink the tainted `intoWriter514`:
	sink(intoWriter514)
}

func TaintStepTest_EncodingCsvReaderRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader326` into `intoString190`.

	// Assume that `sourceCQL` has the underlying type of `fromReader326`:
	fromReader326 := sourceCQL.(csv.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader326` to the result `intoString190`
	// (`intoString190` is now tainted).
	intoString190, _ := fromReader326.Read()

	// Sink the tainted `intoString190`:
	sink(intoString190)
}

func TaintStepTest_EncodingCsvReaderReadAll_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader758` into `intoString842`.

	// Assume that `sourceCQL` has the underlying type of `fromReader758`:
	fromReader758 := sourceCQL.(csv.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader758` to the result `intoString842`
	// (`intoString842` is now tainted).
	intoString842, _ := fromReader758.ReadAll()

	// Sink the tainted `intoString842`:
	sink(intoString842)
}

func TaintStepTest_EncodingCsvWriterWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString476` into `intoWriter695`.

	// Assume that `sourceCQL` has the underlying type of `fromString476`:
	fromString476 := sourceCQL.([]string)

	// Declare `intoWriter695` variable:
	var intoWriter695 csv.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString476` to the receiver `intoWriter695`
	// (`intoWriter695` is now tainted).
	intoWriter695.Write(fromString476)

	// Sink the tainted `intoWriter695`:
	sink(intoWriter695)
}

func TaintStepTest_EncodingCsvWriterWriteAll_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString268` into `intoWriter939`.

	// Assume that `sourceCQL` has the underlying type of `fromString268`:
	fromString268 := sourceCQL.([][]string)

	// Declare `intoWriter939` variable:
	var intoWriter939 csv.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString268` to the receiver `intoWriter939`
	// (`intoWriter939` is now tainted).
	intoWriter939.WriteAll(fromString268)

	// Sink the tainted `intoWriter939`:
	sink(intoWriter939)
}

func RunAllTaints_EncodingCsv() {
	{
		source := newSource()
		TaintStepTest_EncodingCsvNewReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingCsvNewWriter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingCsvReaderRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingCsvReaderReadAll_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingCsvWriterWrite_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingCsvWriterWriteAll_B0I0O0(source)
	}
}
