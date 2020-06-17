package main

import (
	"archive/tar"
	"io"
	"os"
)

func TaintStepTest_ArchiveTarFileInfoHeader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFileInfo908` into `intoHeader792`.

	// Assume that `sourceCQL` has the underlying type of `fromFileInfo908`:
	fromFileInfo908 := sourceCQL.(os.FileInfo)

	// Call the function that transfers the taint
	// from the parameter `fromFileInfo908` to result `intoHeader792`
	// (`intoHeader792` is now tainted).
	intoHeader792, _ := tar.FileInfoHeader(fromFileInfo908, "")

	// Return the tainted `intoHeader792`:
	return intoHeader792
}

func TaintStepTest_ArchiveTarNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader622` into `intoReader413`.

	// Assume that `sourceCQL` has the underlying type of `fromReader622`:
	fromReader622 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader622` to result `intoReader413`
	// (`intoReader413` is now tainted).
	intoReader413 := tar.NewReader(fromReader622)

	// Return the tainted `intoReader413`:
	return intoReader413
}

func TaintStepTest_ArchiveTarNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter551` into `intoWriter953`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter551`:
	fromWriter551 := sourceCQL.(*tar.Writer)

	// Declare `intoWriter953` variable:
	var intoWriter953 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter953`:
	intermediateCQL := tar.NewWriter(intoWriter953)

	// Extra step (`fromWriter551` taints `intermediateCQL`, which taints `intoWriter953`:
	link(fromWriter551, intermediateCQL)

	// Return the tainted `intoWriter953`:
	return intoWriter953
}

func TaintStepTest_ArchiveTarHeaderFileInfo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader451` into `intoFileInfo277`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader451`:
	fromHeader451 := sourceCQL.(tar.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader451` to the result `intoFileInfo277`
	// (`intoFileInfo277` is now tainted).
	intoFileInfo277 := fromHeader451.FileInfo()

	// Return the tainted `intoFileInfo277`:
	return intoFileInfo277
}

func TaintStepTest_ArchiveTarReaderNext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader617` into `intoHeader820`.

	// Assume that `sourceCQL` has the underlying type of `fromReader617`:
	fromReader617 := sourceCQL.(tar.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader617` to the result `intoHeader820`
	// (`intoHeader820` is now tainted).
	intoHeader820, _ := fromReader617.Next()

	// Return the tainted `intoHeader820`:
	return intoHeader820
}

func TaintStepTest_ArchiveTarReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader558` into `intoByte994`.

	// Assume that `sourceCQL` has the underlying type of `fromReader558`:
	fromReader558 := sourceCQL.(tar.Reader)

	// Declare `intoByte994` variable:
	var intoByte994 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader558` to the argument `intoByte994`
	// (`intoByte994` is now tainted).
	fromReader558.Read(intoByte994)

	// Return the tainted `intoByte994`:
	return intoByte994
}

func TaintStepTest_ArchiveTarWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte639` into `intoWriter922`.

	// Assume that `sourceCQL` has the underlying type of `fromByte639`:
	fromByte639 := sourceCQL.([]byte)

	// Declare `intoWriter922` variable:
	var intoWriter922 tar.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte639` to the receiver `intoWriter922`
	// (`intoWriter922` is now tainted).
	intoWriter922.Write(fromByte639)

	// Return the tainted `intoWriter922`:
	return intoWriter922
}

func TaintStepTest_ArchiveTarWriterWriteHeader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader883` into `intoWriter395`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader883`:
	fromHeader883 := sourceCQL.(*tar.Header)

	// Declare `intoWriter395` variable:
	var intoWriter395 tar.Writer

	// Call the method that transfers the taint
	// from the parameter `fromHeader883` to the receiver `intoWriter395`
	// (`intoWriter395` is now tainted).
	intoWriter395.WriteHeader(fromHeader883)

	// Return the tainted `intoWriter395`:
	return intoWriter395
}

func RunAllTaints_ArchiveTar() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveTarFileInfoHeader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveTarNewReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveTarNewWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveTarHeaderFileInfo_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveTarReaderNext_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveTarReaderRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveTarWriterWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveTarWriterWriteHeader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
