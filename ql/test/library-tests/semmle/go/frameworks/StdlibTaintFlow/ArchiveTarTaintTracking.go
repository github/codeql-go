// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"archive/tar"
	"io"
	"os"
)

func TaintStepTest_ArchiveTarFileInfoHeader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFileInfo193` into `intoHeader385`.

	// Assume that `sourceCQL` has the underlying type of `fromFileInfo193`:
	fromFileInfo193 := sourceCQL.(os.FileInfo)

	// Call the function that transfers the taint
	// from the parameter `fromFileInfo193` to result `intoHeader385`
	// (`intoHeader385` is now tainted).
	intoHeader385, _ := tar.FileInfoHeader(fromFileInfo193, "")

	// Return the tainted `intoHeader385`:
	return intoHeader385
}

func TaintStepTest_ArchiveTarNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader668` into `intoReader352`.

	// Assume that `sourceCQL` has the underlying type of `fromReader668`:
	fromReader668 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader668` to result `intoReader352`
	// (`intoReader352` is now tainted).
	intoReader352 := tar.NewReader(fromReader668)

	// Return the tainted `intoReader352`:
	return intoReader352
}

func TaintStepTest_ArchiveTarNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter921` into `intoWriter471`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter921`:
	fromWriter921 := sourceCQL.(*tar.Writer)

	// Declare `intoWriter471` variable:
	var intoWriter471 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter471`:
	intermediateCQL := tar.NewWriter(intoWriter471)

	// Extra step (`fromWriter921` taints `intermediateCQL`, which taints `intoWriter471`:
	link(fromWriter921, intermediateCQL)

	// Return the tainted `intoWriter471`:
	return intoWriter471
}

func TaintStepTest_ArchiveTarHeaderFileInfo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader182` into `intoFileInfo300`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader182`:
	fromHeader182 := sourceCQL.(tar.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader182` to the result `intoFileInfo300`
	// (`intoFileInfo300` is now tainted).
	intoFileInfo300 := fromHeader182.FileInfo()

	// Return the tainted `intoFileInfo300`:
	return intoFileInfo300
}

func TaintStepTest_ArchiveTarReaderNext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader598` into `intoHeader416`.

	// Assume that `sourceCQL` has the underlying type of `fromReader598`:
	fromReader598 := sourceCQL.(tar.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader598` to the result `intoHeader416`
	// (`intoHeader416` is now tainted).
	intoHeader416, _ := fromReader598.Next()

	// Return the tainted `intoHeader416`:
	return intoHeader416
}

func TaintStepTest_ArchiveTarReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader166` into `intoByte997`.

	// Assume that `sourceCQL` has the underlying type of `fromReader166`:
	fromReader166 := sourceCQL.(tar.Reader)

	// Declare `intoByte997` variable:
	var intoByte997 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader166` to the argument `intoByte997`
	// (`intoByte997` is now tainted).
	fromReader166.Read(intoByte997)

	// Return the tainted `intoByte997`:
	return intoByte997
}

func TaintStepTest_ArchiveTarWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte501` into `intoWriter894`.

	// Assume that `sourceCQL` has the underlying type of `fromByte501`:
	fromByte501 := sourceCQL.([]byte)

	// Declare `intoWriter894` variable:
	var intoWriter894 tar.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte501` to the receiver `intoWriter894`
	// (`intoWriter894` is now tainted).
	intoWriter894.Write(fromByte501)

	// Return the tainted `intoWriter894`:
	return intoWriter894
}

func TaintStepTest_ArchiveTarWriterWriteHeader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader973` into `intoWriter946`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader973`:
	fromHeader973 := sourceCQL.(*tar.Header)

	// Declare `intoWriter946` variable:
	var intoWriter946 tar.Writer

	// Call the method that transfers the taint
	// from the parameter `fromHeader973` to the receiver `intoWriter946`
	// (`intoWriter946` is now tainted).
	intoWriter946.WriteHeader(fromHeader973)

	// Return the tainted `intoWriter946`:
	return intoWriter946
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
