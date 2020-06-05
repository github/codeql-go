package main

import (
	"archive/tar"
	"io"
	"os"
)

func TaintStepTest_ArchiveTarFileInfoHeader(sourceCQL interface{}) {
	// The flow is from `fi` into `into723`.

	// Assume that `sourceCQL` has the underlying type of `fi`:
	fi := sourceCQL.(os.FileInfo)

	// Call the function that transfers the taint
	// from the parameter `fi` to result `into723`
	// (`into723` is now tainted).
	into723, _ := tar.FileInfoHeader(fi, "")

	// Sink the tainted `into723`:
	sink(into723)
}

func TaintStepTest_ArchiveTarNewReader(sourceCQL interface{}) {
	// The flow is from `r` into `into528`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into528`
	// (`into528` is now tainted).
	into528 := tar.NewReader(r)

	// Sink the tainted `into528`:
	sink(into528)
}

func TaintStepTest_ArchiveTarNewWriter(sourceCQL interface{}) {
	// The flow is from `from332` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from332`:
	from332 := sourceCQL.(*tar.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := tar.NewWriter(w)

	// Extra step (`from332` taints `intermediateCQL`, which taints `w`:
	link(from332, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_ArchiveTarFormatString(sourceCQL interface{}) {
	// The flow is from `from774` into `into180`.

	// Assume that `sourceCQL` has the underlying type of `from774`:
	from774 := sourceCQL.(tar.Format)

	// Call the method that transfers the taint
	// from the receiver `from774` to the result `into180`
	// (`into180` is now tainted).
	into180 := from774.String()

	// Sink the tainted `into180`:
	sink(into180)
}

func TaintStepTest_ArchiveTarHeaderFileInfo(sourceCQL interface{}) {
	// The flow is from `from835` into `into204`.

	// Assume that `sourceCQL` has the underlying type of `from835`:
	from835 := sourceCQL.(tar.Header)

	// Call the method that transfers the taint
	// from the receiver `from835` to the result `into204`
	// (`into204` is now tainted).
	into204 := from835.FileInfo()

	// Sink the tainted `into204`:
	sink(into204)
}

func TaintStepTest_ArchiveTarReaderNext(sourceCQL interface{}) {
	// The flow is from `from526` into `into869`.

	// Assume that `sourceCQL` has the underlying type of `from526`:
	from526 := sourceCQL.(tar.Reader)

	// Call the method that transfers the taint
	// from the receiver `from526` to the result `into869`
	// (`into869` is now tainted).
	into869, _ := from526.Next()

	// Sink the tainted `into869`:
	sink(into869)
}

func TaintStepTest_ArchiveTarReaderRead(sourceCQL interface{}) {
	// The flow is from `from775` into `b`.

	// Assume that `sourceCQL` has the underlying type of `from775`:
	from775 := sourceCQL.(tar.Reader)

	// Declare `b` variable:
	var b []byte

	// Call the method that transfers the taint
	// from the receiver `from775` to the argument `b`
	// (`b` is now tainted).
	from775.Read(b)

	// Sink the tainted `b`:
	sink(b)
}

func TaintStepTest_ArchiveTarWriterWrite(sourceCQL interface{}) {
	// The flow is from `b` into `into765`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare `into765` variable:
	var into765 tar.Writer

	// Call the method that transfers the taint
	// from the parameter `b` to the receiver `into765`
	// (`into765` is now tainted).
	into765.Write(b)

	// Sink the tainted `into765`:
	sink(into765)
}

func TaintStepTest_ArchiveTarWriterWriteHeader(sourceCQL interface{}) {
	// The flow is from `hdr` into `into490`.

	// Assume that `sourceCQL` has the underlying type of `hdr`:
	hdr := sourceCQL.(*tar.Header)

	// Declare `into490` variable:
	var into490 tar.Writer

	// Call the method that transfers the taint
	// from the parameter `hdr` to the receiver `into490`
	// (`into490` is now tainted).
	into490.WriteHeader(hdr)

	// Sink the tainted `into490`:
	sink(into490)
}

func RunAllTaints_ArchiveTar(v interface{}) {
	{
		source := newSource()
		TaintStepTest_ArchiveTarFileInfoHeader(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveTarNewReader(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveTarNewWriter(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveTarFormatString(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveTarHeaderFileInfo(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveTarReaderNext(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveTarReaderRead(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveTarWriterWrite(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveTarWriterWriteHeader(source)
	}
}
