package main

import (
	"archive/zip"
	"io"
	"os"
)

func TaintStepTest_ArchiveZipFileInfoHeader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFileInfo824` into `intoFileHeader455`.

	// Assume that `sourceCQL` has the underlying type of `fromFileInfo824`:
	fromFileInfo824 := sourceCQL.(os.FileInfo)

	// Call the function that transfers the taint
	// from the parameter `fromFileInfo824` to result `intoFileHeader455`
	// (`intoFileHeader455` is now tainted).
	intoFileHeader455, _ := zip.FileInfoHeader(fromFileInfo824)

	// Sink the tainted `intoFileHeader455`:
	sink(intoFileHeader455)
}

func TaintStepTest_ArchiveZipNewReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReaderAt577` into `intoReader861`.

	// Assume that `sourceCQL` has the underlying type of `fromReaderAt577`:
	fromReaderAt577 := sourceCQL.(io.ReaderAt)

	// Call the function that transfers the taint
	// from the parameter `fromReaderAt577` to result `intoReader861`
	// (`intoReader861` is now tainted).
	intoReader861, _ := zip.NewReader(fromReaderAt577, 0)

	// Sink the tainted `intoReader861`:
	sink(intoReader861)
}

func TaintStepTest_ArchiveZipNewWriter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter643` into `intoWriter482`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter643`:
	fromWriter643 := sourceCQL.(*zip.Writer)

	// Declare `intoWriter482` variable:
	var intoWriter482 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter482`:
	intermediateCQL := zip.NewWriter(intoWriter482)

	// Extra step (`fromWriter643` taints `intermediateCQL`, which taints `intoWriter482`:
	link(fromWriter643, intermediateCQL)

	// Sink the tainted `intoWriter482`:
	sink(intoWriter482)
}

func TaintStepTest_ArchiveZipOpenReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString137` into `intoReadCloser205`.

	// Assume that `sourceCQL` has the underlying type of `fromString137`:
	fromString137 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString137` to result `intoReadCloser205`
	// (`intoReadCloser205` is now tainted).
	intoReadCloser205, _ := zip.OpenReader(fromString137)

	// Sink the tainted `intoReadCloser205`:
	sink(intoReadCloser205)
}

func TaintStepTest_ArchiveZipFileOpen_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFile432` into `intoReadCloser406`.

	// Assume that `sourceCQL` has the underlying type of `fromFile432`:
	fromFile432 := sourceCQL.(zip.File)

	// Call the method that transfers the taint
	// from the receiver `fromFile432` to the result `intoReadCloser406`
	// (`intoReadCloser406` is now tainted).
	intoReadCloser406, _ := fromFile432.Open()

	// Sink the tainted `intoReadCloser406`:
	sink(intoReadCloser406)
}

func TaintStepTest_ArchiveZipFileHeaderFileInfo_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFileHeader614` into `intoFileInfo617`.

	// Assume that `sourceCQL` has the underlying type of `fromFileHeader614`:
	fromFileHeader614 := sourceCQL.(zip.FileHeader)

	// Call the method that transfers the taint
	// from the receiver `fromFileHeader614` to the result `intoFileInfo617`
	// (`intoFileInfo617` is now tainted).
	intoFileInfo617 := fromFileHeader614.FileInfo()

	// Sink the tainted `intoFileInfo617`:
	sink(intoFileInfo617)
}

func TaintStepTest_ArchiveZipFileHeaderMode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFileHeader428` into `intoFileMode138`.

	// Assume that `sourceCQL` has the underlying type of `fromFileHeader428`:
	fromFileHeader428 := sourceCQL.(zip.FileHeader)

	// Call the method that transfers the taint
	// from the receiver `fromFileHeader428` to the result `intoFileMode138`
	// (`intoFileMode138` is now tainted).
	intoFileMode138 := fromFileHeader428.Mode()

	// Sink the tainted `intoFileMode138`:
	sink(intoFileMode138)
}

func TaintStepTest_ArchiveZipFileHeaderSetMode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFileMode871` into `intoFileHeader415`.

	// Assume that `sourceCQL` has the underlying type of `fromFileMode871`:
	fromFileMode871 := sourceCQL.(os.FileMode)

	// Declare `intoFileHeader415` variable:
	var intoFileHeader415 zip.FileHeader

	// Call the method that transfers the taint
	// from the parameter `fromFileMode871` to the receiver `intoFileHeader415`
	// (`intoFileHeader415` is now tainted).
	intoFileHeader415.SetMode(fromFileMode871)

	// Sink the tainted `intoFileHeader415`:
	sink(intoFileHeader415)
}

func TaintStepTest_ArchiveZipReaderRegisterDecompressor_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromDecompressor868` into `intoReader151`.

	// Assume that `sourceCQL` has the underlying type of `fromDecompressor868`:
	fromDecompressor868 := sourceCQL.(zip.Decompressor)

	// Declare `intoReader151` variable:
	var intoReader151 zip.Reader

	// Call the method that transfers the taint
	// from the parameter `fromDecompressor868` to the receiver `intoReader151`
	// (`intoReader151` is now tainted).
	intoReader151.RegisterDecompressor(0, fromDecompressor868)

	// Sink the tainted `intoReader151`:
	sink(intoReader151)
}

func TaintStepTest_ArchiveZipWriterCreate_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter379` into `intoWriter536`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter379`:
	fromWriter379 := sourceCQL.(io.Writer)

	// Declare `intoWriter536` variable:
	var intoWriter536 zip.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter536`:
	intermediateCQL, _ := intoWriter536.Create("")

	// Extra step (`fromWriter379` taints `intermediateCQL`, which taints `intoWriter536`:
	link(fromWriter379, intermediateCQL)

	// Sink the tainted `intoWriter536`:
	sink(intoWriter536)
}

func TaintStepTest_ArchiveZipWriterCreateHeader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFileHeader759` into `intoWriter313`.

	// Assume that `sourceCQL` has the underlying type of `fromFileHeader759`:
	fromFileHeader759 := sourceCQL.(*zip.FileHeader)

	// Declare `intoWriter313` variable:
	var intoWriter313 zip.Writer

	// Call the method that transfers the taint
	// from the parameter `fromFileHeader759` to the receiver `intoWriter313`
	// (`intoWriter313` is now tainted).
	intoWriter313.CreateHeader(fromFileHeader759)

	// Sink the tainted `intoWriter313`:
	sink(intoWriter313)
}

func TaintStepTest_ArchiveZipWriterCreateHeader_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromWriter658` into `intoWriter155`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter658`:
	fromWriter658 := sourceCQL.(io.Writer)

	// Declare `intoWriter155` variable:
	var intoWriter155 zip.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoWriter155`:
	intermediateCQL, _ := intoWriter155.CreateHeader(nil)

	// Extra step (`fromWriter658` taints `intermediateCQL`, which taints `intoWriter155`:
	link(fromWriter658, intermediateCQL)

	// Sink the tainted `intoWriter155`:
	sink(intoWriter155)
}

func TaintStepTest_ArchiveZipWriterRegisterCompressor_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromCompressor967` into `intoWriter713`.

	// Assume that `sourceCQL` has the underlying type of `fromCompressor967`:
	fromCompressor967 := sourceCQL.(zip.Compressor)

	// Declare `intoWriter713` variable:
	var intoWriter713 zip.Writer

	// Call the method that transfers the taint
	// from the parameter `fromCompressor967` to the receiver `intoWriter713`
	// (`intoWriter713` is now tainted).
	intoWriter713.RegisterCompressor(0, fromCompressor967)

	// Sink the tainted `intoWriter713`:
	sink(intoWriter713)
}

func TaintStepTest_ArchiveZipWriterSetComment_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString773` into `intoWriter180`.

	// Assume that `sourceCQL` has the underlying type of `fromString773`:
	fromString773 := sourceCQL.(string)

	// Declare `intoWriter180` variable:
	var intoWriter180 zip.Writer

	// Call the method that transfers the taint
	// from the parameter `fromString773` to the receiver `intoWriter180`
	// (`intoWriter180` is now tainted).
	intoWriter180.SetComment(fromString773)

	// Sink the tainted `intoWriter180`:
	sink(intoWriter180)
}

func RunAllTaints_ArchiveZip() {
	{
		source := newSource()
		TaintStepTest_ArchiveZipFileInfoHeader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipNewReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipNewWriter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipOpenReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipFileOpen_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipFileHeaderFileInfo_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipFileHeaderMode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipFileHeaderSetMode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipReaderRegisterDecompressor_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipWriterCreate_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipWriterCreateHeader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipWriterCreateHeader_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipWriterRegisterCompressor_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipWriterSetComment_B0I0O0(source)
	}
}
