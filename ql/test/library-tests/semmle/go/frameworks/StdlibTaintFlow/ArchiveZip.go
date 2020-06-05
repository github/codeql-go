package main

import (
	"archive/zip"
	"io"
	"os"
)

func TaintStepTest_ArchiveZipFileInfoHeader(sourceCQL interface{}) {
	// The flow is from `fi` into `into329`.

	// Assume that `sourceCQL` has the underlying type of `fi`:
	fi := sourceCQL.(os.FileInfo)

	// Call the function that transfers the taint
	// from the parameter `fi` to result `into329`
	// (`into329` is now tainted).
	into329, _ := zip.FileInfoHeader(fi)

	// Sink the tainted `into329`:
	sink(into329)
}

func TaintStepTest_ArchiveZipNewReader(sourceCQL interface{}) {
	// The flow is from `r` into `into464`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.ReaderAt)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into464`
	// (`into464` is now tainted).
	into464, _ := zip.NewReader(r, 0)

	// Sink the tainted `into464`:
	sink(into464)
}

func TaintStepTest_ArchiveZipNewWriter(sourceCQL interface{}) {
	// The flow is from `from164` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from164`:
	from164 := sourceCQL.(*zip.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := zip.NewWriter(w)

	// Extra step (`from164` taints `intermediateCQL`, which taints `w`:
	link(from164, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_ArchiveZipOpenReader(sourceCQL interface{}) {
	// The flow is from `name` into `into829`.

	// Assume that `sourceCQL` has the underlying type of `name`:
	name := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `name` to result `into829`
	// (`into829` is now tainted).
	into829, _ := zip.OpenReader(name)

	// Sink the tainted `into829`:
	sink(into829)
}

func TaintStepTest_ArchiveZipFileOpen(sourceCQL interface{}) {
	// The flow is from `from611` into `into334`.

	// Assume that `sourceCQL` has the underlying type of `from611`:
	from611 := sourceCQL.(zip.File)

	// Call the method that transfers the taint
	// from the receiver `from611` to the result `into334`
	// (`into334` is now tainted).
	into334, _ := from611.Open()

	// Sink the tainted `into334`:
	sink(into334)
}

func TaintStepTest_ArchiveZipFileHeaderFileInfo(sourceCQL interface{}) {
	// The flow is from `from450` into `into538`.

	// Assume that `sourceCQL` has the underlying type of `from450`:
	from450 := sourceCQL.(zip.FileHeader)

	// Call the method that transfers the taint
	// from the receiver `from450` to the result `into538`
	// (`into538` is now tainted).
	into538 := from450.FileInfo()

	// Sink the tainted `into538`:
	sink(into538)
}

func TaintStepTest_ArchiveZipFileHeaderMode(sourceCQL interface{}) {
	// The flow is from `from674` into `mode`.

	// Assume that `sourceCQL` has the underlying type of `from674`:
	from674 := sourceCQL.(zip.FileHeader)

	// Call the method that transfers the taint
	// from the receiver `from674` to the result `mode`
	// (`mode` is now tainted).
	mode := from674.Mode()

	// Sink the tainted `mode`:
	sink(mode)
}

func TaintStepTest_ArchiveZipFileHeaderSetMode(sourceCQL interface{}) {
	// The flow is from `mode` into `from160`.

	// Assume that `sourceCQL` has the underlying type of `mode`:
	mode := sourceCQL.(os.FileMode)

	// Declare `from160` variable:
	var from160 zip.FileHeader

	// Call the method that transfers the taint
	// from the parameter `mode` to the receiver `from160`
	// (`from160` is now tainted).
	from160.SetMode(mode)

	// Sink the tainted `from160`:
	sink(from160)
}

func TaintStepTest_ArchiveZipReaderRegisterDecompressor(sourceCQL interface{}) {
	// The flow is from `dcomp` into `into931`.

	// Assume that `sourceCQL` has the underlying type of `dcomp`:
	dcomp := sourceCQL.(zip.Decompressor)

	// Declare `into931` variable:
	var into931 zip.Reader

	// Call the method that transfers the taint
	// from the parameter `dcomp` to the receiver `into931`
	// (`into931` is now tainted).
	into931.RegisterDecompressor(0, dcomp)

	// Sink the tainted `into931`:
	sink(into931)
}

func TaintStepTest_ArchiveZipWriterCreate(sourceCQL interface{}) {
	// The flow is from `into485` into `from878`.

	// Assume that `sourceCQL` has the underlying type of `into485`:
	into485 := sourceCQL.(io.Writer)

	// Declare `from878` variable:
	var from878 zip.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `from878`:
	intermediateCQL, _ := from878.Create("")

	// Extra step (`into485` taints `intermediateCQL`, which taints `from878`:
	link(into485, intermediateCQL)

	// Sink the tainted `from878`:
	sink(from878)
}

func TaintStepTest_ArchiveZipWriterCreateHeader(sourceCQL interface{}) {
	// The flow is from `from869` into `into133`.

	// Assume that `sourceCQL` has the underlying type of `from869`:
	from869 := sourceCQL.(io.Writer)

	// Declare `into133` variable:
	var into133 zip.Writer

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `into133`:
	intermediateCQL, _ := into133.CreateHeader(nil)

	// Extra step (`from869` taints `intermediateCQL`, which taints `into133`:
	link(from869, intermediateCQL)

	// Sink the tainted `into133`:
	sink(into133)
}

func TaintStepTest_ArchiveZipWriterRegisterCompressor(sourceCQL interface{}) {
	// The flow is from `comp` into `into137`.

	// Assume that `sourceCQL` has the underlying type of `comp`:
	comp := sourceCQL.(zip.Compressor)

	// Declare `into137` variable:
	var into137 zip.Writer

	// Call the method that transfers the taint
	// from the parameter `comp` to the receiver `into137`
	// (`into137` is now tainted).
	into137.RegisterCompressor(0, comp)

	// Sink the tainted `into137`:
	sink(into137)
}

func TaintStepTest_ArchiveZipWriterSetComment(sourceCQL interface{}) {
	// The flow is from `comment` into `into163`.

	// Assume that `sourceCQL` has the underlying type of `comment`:
	comment := sourceCQL.(string)

	// Declare `into163` variable:
	var into163 zip.Writer

	// Call the method that transfers the taint
	// from the parameter `comment` to the receiver `into163`
	// (`into163` is now tainted).
	into163.SetComment(comment)

	// Sink the tainted `into163`:
	sink(into163)
}

func RunAllTaints_ArchiveZip(v interface{}) {
	{
		source := newSource()
		TaintStepTest_ArchiveZipFileInfoHeader(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipNewReader(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipNewWriter(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipOpenReader(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipFileOpen(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipFileHeaderFileInfo(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipFileHeaderMode(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipFileHeaderSetMode(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipReaderRegisterDecompressor(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipWriterCreate(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipWriterCreateHeader(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipWriterRegisterCompressor(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveZipWriterSetComment(source)
	}
}
