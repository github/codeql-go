package main

import (
	"archive/zip"
	"io"
	"os"
)

func TaintStepTest_ArchiveZipFileInfoHeader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFileInfo645` into `intoFileHeader452`.

	// Assume that `sourceCQL` has the underlying type of `fromFileInfo645`:
	fromFileInfo645 := sourceCQL.(os.FileInfo)

	// Call the function that transfers the taint
	// from the parameter `fromFileInfo645` to result `intoFileHeader452`
	// (`intoFileHeader452` is now tainted).
	intoFileHeader452, _ := zip.FileInfoHeader(fromFileInfo645)

	// Return the tainted `intoFileHeader452`:
	return intoFileHeader452
}

func TaintStepTest_ArchiveZipNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReaderAt535` into `intoReader974`.

	// Assume that `sourceCQL` has the underlying type of `fromReaderAt535`:
	fromReaderAt535 := sourceCQL.(io.ReaderAt)

	// Call the function that transfers the taint
	// from the parameter `fromReaderAt535` to result `intoReader974`
	// (`intoReader974` is now tainted).
	intoReader974, _ := zip.NewReader(fromReaderAt535, 0)

	// Return the tainted `intoReader974`:
	return intoReader974
}

func TaintStepTest_ArchiveZipNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter374` into `intoWriter769`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter374`:
	fromWriter374 := sourceCQL.(*zip.Writer)

	// Declare `intoWriter769` variable:
	var intoWriter769 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter769`:
	intermediateCQL := zip.NewWriter(intoWriter769)

	// Extra step (`fromWriter374` taints `intermediateCQL`, which taints `intoWriter769`:
	link(fromWriter374, intermediateCQL)

	// Return the tainted `intoWriter769`:
	return intoWriter769
}

func TaintStepTest_ArchiveZipOpenReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString418` into `intoReadCloser901`.

	// Assume that `sourceCQL` has the underlying type of `fromString418`:
	fromString418 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString418` to result `intoReadCloser901`
	// (`intoReadCloser901` is now tainted).
	intoReadCloser901, _ := zip.OpenReader(fromString418)

	// Return the tainted `intoReadCloser901`:
	return intoReadCloser901
}

func TaintStepTest_ArchiveZipFileOpen_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile554` into `intoReadCloser623`.

	// Assume that `sourceCQL` has the underlying type of `fromFile554`:
	fromFile554 := sourceCQL.(zip.File)

	// Call the method that transfers the taint
	// from the receiver `fromFile554` to the result `intoReadCloser623`
	// (`intoReadCloser623` is now tainted).
	intoReadCloser623, _ := fromFile554.Open()

	// Return the tainted `intoReadCloser623`:
	return intoReadCloser623
}

func TaintStepTest_ArchiveZipFileHeaderFileInfo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFileHeader260` into `intoFileInfo977`.

	// Assume that `sourceCQL` has the underlying type of `fromFileHeader260`:
	fromFileHeader260 := sourceCQL.(zip.FileHeader)

	// Call the method that transfers the taint
	// from the receiver `fromFileHeader260` to the result `intoFileInfo977`
	// (`intoFileInfo977` is now tainted).
	intoFileInfo977 := fromFileHeader260.FileInfo()

	// Return the tainted `intoFileInfo977`:
	return intoFileInfo977
}

func TaintStepTest_ArchiveZipFileHeaderMode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFileHeader341` into `intoFileMode534`.

	// Assume that `sourceCQL` has the underlying type of `fromFileHeader341`:
	fromFileHeader341 := sourceCQL.(zip.FileHeader)

	// Call the method that transfers the taint
	// from the receiver `fromFileHeader341` to the result `intoFileMode534`
	// (`intoFileMode534` is now tainted).
	intoFileMode534 := fromFileHeader341.Mode()

	// Return the tainted `intoFileMode534`:
	return intoFileMode534
}

func TaintStepTest_ArchiveZipFileHeaderSetMode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFileMode854` into `intoFileHeader764`.

	// Assume that `sourceCQL` has the underlying type of `fromFileMode854`:
	fromFileMode854 := sourceCQL.(os.FileMode)

	// Declare `intoFileHeader764` variable:
	var intoFileHeader764 zip.FileHeader

	// Call the method that transfers the taint
	// from the parameter `fromFileMode854` to the receiver `intoFileHeader764`
	// (`intoFileHeader764` is now tainted).
	intoFileHeader764.SetMode(fromFileMode854)

	// Return the tainted `intoFileHeader764`:
	return intoFileHeader764
}

func TaintStepTest_ArchiveZipReaderRegisterDecompressor_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecompressor946` into `intoReader679`.

	// Assume that `sourceCQL` has the underlying type of `fromDecompressor946`:
	fromDecompressor946 := sourceCQL.(zip.Decompressor)

	// Declare `intoReader679` variable:
	var intoReader679 zip.Reader

	// Call the method that transfers the taint
	// from the parameter `fromDecompressor946` to the receiver `intoReader679`
	// (`intoReader679` is now tainted).
	intoReader679.RegisterDecompressor(0, fromDecompressor946)

	// Return the tainted `intoReader679`:
	return intoReader679
}

func TaintStepTest_ArchiveZipWriterCreate_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter262` into `intoWriter934`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter262`:
	fromWriter262 := sourceCQL.(io.Writer)

	// Declare `intoWriter934` variable:
	var intoWriter934 zip.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter934`:
	intermediateCQL, _ := intoWriter934.Create("")

	// Extra step (`fromWriter262` taints `intermediateCQL`, which taints `intoWriter934`:
	link(fromWriter262, intermediateCQL)

	// Return the tainted `intoWriter934`:
	return intoWriter934
}

func TaintStepTest_ArchiveZipWriterCreateHeader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFileHeader995` into `intoWriter561`.

	// Assume that `sourceCQL` has the underlying type of `fromFileHeader995`:
	fromFileHeader995 := sourceCQL.(*zip.FileHeader)

	// Declare `intoWriter561` variable:
	var intoWriter561 zip.Writer

	// Call the method that transfers the taint
	// from the parameter `fromFileHeader995` to the receiver `intoWriter561`
	// (`intoWriter561` is now tainted).
	intoWriter561.CreateHeader(fromFileHeader995)

	// Return the tainted `intoWriter561`:
	return intoWriter561
}

func TaintStepTest_ArchiveZipWriterCreateHeader_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter365` into `intoWriter832`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter365`:
	fromWriter365 := sourceCQL.(io.Writer)

	// Declare `intoWriter832` variable:
	var intoWriter832 zip.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter832`:
	intermediateCQL, _ := intoWriter832.CreateHeader(nil)

	// Extra step (`fromWriter365` taints `intermediateCQL`, which taints `intoWriter832`:
	link(fromWriter365, intermediateCQL)

	// Return the tainted `intoWriter832`:
	return intoWriter832
}

func TaintStepTest_ArchiveZipWriterRegisterCompressor_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromCompressor406` into `intoWriter903`.

	// Assume that `sourceCQL` has the underlying type of `fromCompressor406`:
	fromCompressor406 := sourceCQL.(zip.Compressor)

	// Declare `intoWriter903` variable:
	var intoWriter903 zip.Writer

	// Call the method that transfers the taint
	// from the parameter `fromCompressor406` to the receiver `intoWriter903`
	// (`intoWriter903` is now tainted).
	intoWriter903.RegisterCompressor(0, fromCompressor406)

	// Return the tainted `intoWriter903`:
	return intoWriter903
}

func TaintStepTest_ArchiveZipWriterSetComment_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString992` into `intoWriter299`.

	// Assume that `sourceCQL` has the underlying type of `fromString992`:
	fromString992 := sourceCQL.(string)

	// Declare `intoWriter299` variable:
	var intoWriter299 zip.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString992` to the receiver `intoWriter299`
	// (`intoWriter299` is now tainted).
	intoWriter299.SetComment(fromString992)

	// Return the tainted `intoWriter299`:
	return intoWriter299
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
