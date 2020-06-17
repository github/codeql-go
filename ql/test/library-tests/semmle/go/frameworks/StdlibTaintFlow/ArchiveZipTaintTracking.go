// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"archive/zip"
	"io"
	"os"
)

func TaintStepTest_ArchiveZipFileInfoHeader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFileInfo645` into `intoFileHeader483`.

	// Assume that `sourceCQL` has the underlying type of `fromFileInfo645`:
	fromFileInfo645 := sourceCQL.(os.FileInfo)

	// Call the function that transfers the taint
	// from the parameter `fromFileInfo645` to result `intoFileHeader483`
	// (`intoFileHeader483` is now tainted).
	intoFileHeader483, _ := zip.FileInfoHeader(fromFileInfo645)

	// Return the tainted `intoFileHeader483`:
	return intoFileHeader483
}

func TaintStepTest_ArchiveZipNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReaderAt602` into `intoReader565`.

	// Assume that `sourceCQL` has the underlying type of `fromReaderAt602`:
	fromReaderAt602 := sourceCQL.(io.ReaderAt)

	// Call the function that transfers the taint
	// from the parameter `fromReaderAt602` to result `intoReader565`
	// (`intoReader565` is now tainted).
	intoReader565, _ := zip.NewReader(fromReaderAt602, 0)

	// Return the tainted `intoReader565`:
	return intoReader565
}

func TaintStepTest_ArchiveZipNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter267` into `intoWriter128`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter267`:
	fromWriter267 := sourceCQL.(*zip.Writer)

	// Declare `intoWriter128` variable:
	var intoWriter128 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter128`:
	intermediateCQL := zip.NewWriter(intoWriter128)

	// Extra step (`fromWriter267` taints `intermediateCQL`, which taints `intoWriter128`:
	link(fromWriter267, intermediateCQL)

	// Return the tainted `intoWriter128`:
	return intoWriter128
}

func TaintStepTest_ArchiveZipOpenReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString904` into `intoReadCloser321`.

	// Assume that `sourceCQL` has the underlying type of `fromString904`:
	fromString904 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString904` to result `intoReadCloser321`
	// (`intoReadCloser321` is now tainted).
	intoReadCloser321, _ := zip.OpenReader(fromString904)

	// Return the tainted `intoReadCloser321`:
	return intoReadCloser321
}

func TaintStepTest_ArchiveZipFileOpen_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile342` into `intoReadCloser585`.

	// Assume that `sourceCQL` has the underlying type of `fromFile342`:
	fromFile342 := sourceCQL.(zip.File)

	// Call the method that transfers the taint
	// from the receiver `fromFile342` to the result `intoReadCloser585`
	// (`intoReadCloser585` is now tainted).
	intoReadCloser585, _ := fromFile342.Open()

	// Return the tainted `intoReadCloser585`:
	return intoReadCloser585
}

func TaintStepTest_ArchiveZipFileHeaderFileInfo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFileHeader770` into `intoFileInfo666`.

	// Assume that `sourceCQL` has the underlying type of `fromFileHeader770`:
	fromFileHeader770 := sourceCQL.(zip.FileHeader)

	// Call the method that transfers the taint
	// from the receiver `fromFileHeader770` to the result `intoFileInfo666`
	// (`intoFileInfo666` is now tainted).
	intoFileInfo666 := fromFileHeader770.FileInfo()

	// Return the tainted `intoFileInfo666`:
	return intoFileInfo666
}

func TaintStepTest_ArchiveZipFileHeaderMode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFileHeader760` into `intoFileMode416`.

	// Assume that `sourceCQL` has the underlying type of `fromFileHeader760`:
	fromFileHeader760 := sourceCQL.(zip.FileHeader)

	// Call the method that transfers the taint
	// from the receiver `fromFileHeader760` to the result `intoFileMode416`
	// (`intoFileMode416` is now tainted).
	intoFileMode416 := fromFileHeader760.Mode()

	// Return the tainted `intoFileMode416`:
	return intoFileMode416
}

func TaintStepTest_ArchiveZipFileHeaderSetMode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFileMode285` into `intoFileHeader708`.

	// Assume that `sourceCQL` has the underlying type of `fromFileMode285`:
	fromFileMode285 := sourceCQL.(os.FileMode)

	// Declare `intoFileHeader708` variable:
	var intoFileHeader708 zip.FileHeader

	// Call the method that transfers the taint
	// from the parameter `fromFileMode285` to the receiver `intoFileHeader708`
	// (`intoFileHeader708` is now tainted).
	intoFileHeader708.SetMode(fromFileMode285)

	// Return the tainted `intoFileHeader708`:
	return intoFileHeader708
}

func TaintStepTest_ArchiveZipReaderRegisterDecompressor_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecompressor575` into `intoReader214`.

	// Assume that `sourceCQL` has the underlying type of `fromDecompressor575`:
	fromDecompressor575 := sourceCQL.(zip.Decompressor)

	// Declare `intoReader214` variable:
	var intoReader214 zip.Reader

	// Call the method that transfers the taint
	// from the parameter `fromDecompressor575` to the receiver `intoReader214`
	// (`intoReader214` is now tainted).
	intoReader214.RegisterDecompressor(0, fromDecompressor575)

	// Return the tainted `intoReader214`:
	return intoReader214
}

func TaintStepTest_ArchiveZipWriterCreate_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter208` into `intoWriter698`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter208`:
	fromWriter208 := sourceCQL.(io.Writer)

	// Declare `intoWriter698` variable:
	var intoWriter698 zip.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter698`:
	intermediateCQL, _ := intoWriter698.Create("")

	// Extra step (`fromWriter208` taints `intermediateCQL`, which taints `intoWriter698`:
	link(fromWriter208, intermediateCQL)

	// Return the tainted `intoWriter698`:
	return intoWriter698
}

func TaintStepTest_ArchiveZipWriterCreateHeader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFileHeader896` into `intoWriter175`.

	// Assume that `sourceCQL` has the underlying type of `fromFileHeader896`:
	fromFileHeader896 := sourceCQL.(*zip.FileHeader)

	// Declare `intoWriter175` variable:
	var intoWriter175 zip.Writer

	// Call the method that transfers the taint
	// from the parameter `fromFileHeader896` to the receiver `intoWriter175`
	// (`intoWriter175` is now tainted).
	intoWriter175.CreateHeader(fromFileHeader896)

	// Return the tainted `intoWriter175`:
	return intoWriter175
}

func TaintStepTest_ArchiveZipWriterCreateHeader_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter754` into `intoWriter502`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter754`:
	fromWriter754 := sourceCQL.(io.Writer)

	// Declare `intoWriter502` variable:
	var intoWriter502 zip.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter502`:
	intermediateCQL, _ := intoWriter502.CreateHeader(nil)

	// Extra step (`fromWriter754` taints `intermediateCQL`, which taints `intoWriter502`:
	link(fromWriter754, intermediateCQL)

	// Return the tainted `intoWriter502`:
	return intoWriter502
}

func TaintStepTest_ArchiveZipWriterRegisterCompressor_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromCompressor414` into `intoWriter975`.

	// Assume that `sourceCQL` has the underlying type of `fromCompressor414`:
	fromCompressor414 := sourceCQL.(zip.Compressor)

	// Declare `intoWriter975` variable:
	var intoWriter975 zip.Writer

	// Call the method that transfers the taint
	// from the parameter `fromCompressor414` to the receiver `intoWriter975`
	// (`intoWriter975` is now tainted).
	intoWriter975.RegisterCompressor(0, fromCompressor414)

	// Return the tainted `intoWriter975`:
	return intoWriter975
}

func TaintStepTest_ArchiveZipWriterSetComment_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString251` into `intoWriter753`.

	// Assume that `sourceCQL` has the underlying type of `fromString251`:
	fromString251 := sourceCQL.(string)

	// Declare `intoWriter753` variable:
	var intoWriter753 zip.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString251` to the receiver `intoWriter753`
	// (`intoWriter753` is now tainted).
	intoWriter753.SetComment(fromString251)

	// Return the tainted `intoWriter753`:
	return intoWriter753
}

func RunAllTaints_ArchiveZip() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveZipFileInfoHeader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveZipNewReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveZipNewWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveZipOpenReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveZipFileOpen_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveZipFileHeaderFileInfo_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveZipFileHeaderMode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveZipFileHeaderSetMode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveZipReaderRegisterDecompressor_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveZipWriterCreate_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveZipWriterCreateHeader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveZipWriterCreateHeader_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveZipWriterRegisterCompressor_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ArchiveZipWriterSetComment_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
