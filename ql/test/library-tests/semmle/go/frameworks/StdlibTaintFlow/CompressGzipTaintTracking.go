package main

import (
	"compress/gzip"
	"io"
)

func TaintStepTest_CompressGzipNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader853` into `intoReader620`.

	// Assume that `sourceCQL` has the underlying type of `fromReader853`:
	fromReader853 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader853` to result `intoReader620`
	// (`intoReader620` is now tainted).
	intoReader620, _ := gzip.NewReader(fromReader853)

	// Return the tainted `intoReader620`:
	return intoReader620
}

func TaintStepTest_CompressGzipNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter749` into `intoWriter546`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter749`:
	fromWriter749 := sourceCQL.(*gzip.Writer)

	// Declare `intoWriter546` variable:
	var intoWriter546 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter546`:
	intermediateCQL := gzip.NewWriter(intoWriter546)

	// Extra step (`fromWriter749` taints `intermediateCQL`, which taints `intoWriter546`:
	link(fromWriter749, intermediateCQL)

	// Return the tainted `intoWriter546`:
	return intoWriter546
}

func TaintStepTest_CompressGzipNewWriterLevel_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter583` into `intoWriter937`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter583`:
	fromWriter583 := sourceCQL.(*gzip.Writer)

	// Declare `intoWriter937` variable:
	var intoWriter937 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter937`:
	intermediateCQL, _ := gzip.NewWriterLevel(intoWriter937, 0)

	// Extra step (`fromWriter583` taints `intermediateCQL`, which taints `intoWriter937`:
	link(fromWriter583, intermediateCQL)

	// Return the tainted `intoWriter937`:
	return intoWriter937
}

func TaintStepTest_CompressGzipReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader331` into `intoByte493`.

	// Assume that `sourceCQL` has the underlying type of `fromReader331`:
	fromReader331 := sourceCQL.(gzip.Reader)

	// Declare `intoByte493` variable:
	var intoByte493 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader331` to the argument `intoByte493`
	// (`intoByte493` is now tainted).
	fromReader331.Read(intoByte493)

	// Return the tainted `intoByte493`:
	return intoByte493
}

func TaintStepTest_CompressGzipReaderReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader426` into `intoReader886`.

	// Assume that `sourceCQL` has the underlying type of `fromReader426`:
	fromReader426 := sourceCQL.(io.Reader)

	// Declare `intoReader886` variable:
	var intoReader886 gzip.Reader

	// Call the method that transfers the taint
	// from the parameter `fromReader426` to the receiver `intoReader886`
	// (`intoReader886` is now tainted).
	intoReader886.Reset(fromReader426)

	// Return the tainted `intoReader886`:
	return intoReader886
}

func TaintStepTest_CompressGzipWriterReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter875` into `intoWriter147`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter875`:
	fromWriter875 := sourceCQL.(gzip.Writer)

	// Declare `intoWriter147` variable:
	var intoWriter147 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriter875` to the argument `intoWriter147`
	// (`intoWriter147` is now tainted).
	fromWriter875.Reset(intoWriter147)

	// Return the tainted `intoWriter147`:
	return intoWriter147
}

func TaintStepTest_CompressGzipWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte514` into `intoWriter822`.

	// Assume that `sourceCQL` has the underlying type of `fromByte514`:
	fromByte514 := sourceCQL.([]byte)

	// Declare `intoWriter822` variable:
	var intoWriter822 gzip.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte514` to the receiver `intoWriter822`
	// (`intoWriter822` is now tainted).
	intoWriter822.Write(fromByte514)

	// Return the tainted `intoWriter822`:
	return intoWriter822
}

func RunAllTaints_CompressGzip() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressGzipNewReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressGzipNewWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressGzipNewWriterLevel_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressGzipReaderRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressGzipReaderReset_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressGzipWriterReset_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_CompressGzipWriterWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
