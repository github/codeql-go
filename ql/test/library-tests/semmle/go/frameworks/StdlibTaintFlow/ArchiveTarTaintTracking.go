package main

import (
	"archive/tar"
	"io"
	"os"
)

func TaintStepTest_ArchiveTarFileInfoHeader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFileInfo149` into `intoHeader348`.

	// Assume that `sourceCQL` has the underlying type of `fromFileInfo149`:
	fromFileInfo149 := sourceCQL.(os.FileInfo)

	// Call the function that transfers the taint
	// from the parameter `fromFileInfo149` to result `intoHeader348`
	// (`intoHeader348` is now tainted).
	intoHeader348, _ := tar.FileInfoHeader(fromFileInfo149, "")

	// Sink the tainted `intoHeader348`:
	sink(intoHeader348)
}

func TaintStepTest_ArchiveTarNewReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader458` into `intoReader708`.

	// Assume that `sourceCQL` has the underlying type of `fromReader458`:
	fromReader458 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader458` to result `intoReader708`
	// (`intoReader708` is now tainted).
	intoReader708 := tar.NewReader(fromReader458)

	// Sink the tainted `intoReader708`:
	sink(intoReader708)
}

func TaintStepTest_ArchiveTarNewWriter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter130` into `intoWriter839`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter130`:
	fromWriter130 := sourceCQL.(*tar.Writer)

	// Declare `intoWriter839` variable:
	var intoWriter839 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter839`:
	intermediateCQL := tar.NewWriter(intoWriter839)

	// Extra step (`fromWriter130` taints `intermediateCQL`, which taints `intoWriter839`:
	link(fromWriter130, intermediateCQL)

	// Sink the tainted `intoWriter839`:
	sink(intoWriter839)
}

func TaintStepTest_ArchiveTarHeaderFileInfo_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromHeader292` into `intoFileInfo779`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader292`:
	fromHeader292 := sourceCQL.(tar.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader292` to the result `intoFileInfo779`
	// (`intoFileInfo779` is now tainted).
	intoFileInfo779 := fromHeader292.FileInfo()

	// Sink the tainted `intoFileInfo779`:
	sink(intoFileInfo779)
}

func TaintStepTest_ArchiveTarReaderNext_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader216` into `intoHeader237`.

	// Assume that `sourceCQL` has the underlying type of `fromReader216`:
	fromReader216 := sourceCQL.(tar.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader216` to the result `intoHeader237`
	// (`intoHeader237` is now tainted).
	intoHeader237, _ := fromReader216.Next()

	// Sink the tainted `intoHeader237`:
	sink(intoHeader237)
}

func TaintStepTest_ArchiveTarReaderRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader366` into `intoByte295`.

	// Assume that `sourceCQL` has the underlying type of `fromReader366`:
	fromReader366 := sourceCQL.(tar.Reader)

	// Declare `intoByte295` variable:
	var intoByte295 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader366` to the argument `intoByte295`
	// (`intoByte295` is now tainted).
	fromReader366.Read(intoByte295)

	// Sink the tainted `intoByte295`:
	sink(intoByte295)
}

func TaintStepTest_ArchiveTarWriterWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte367` into `intoWriter847`.

	// Assume that `sourceCQL` has the underlying type of `fromByte367`:
	fromByte367 := sourceCQL.([]byte)

	// Declare `intoWriter847` variable:
	var intoWriter847 tar.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte367` to the receiver `intoWriter847`
	// (`intoWriter847` is now tainted).
	intoWriter847.Write(fromByte367)

	// Sink the tainted `intoWriter847`:
	sink(intoWriter847)
}

func TaintStepTest_ArchiveTarWriterWriteHeader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromHeader257` into `intoWriter439`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader257`:
	fromHeader257 := sourceCQL.(*tar.Header)

	// Declare `intoWriter439` variable:
	var intoWriter439 tar.Writer

	// Call the method that transfers the taint
	// from the parameter `fromHeader257` to the receiver `intoWriter439`
	// (`intoWriter439` is now tainted).
	intoWriter439.WriteHeader(fromHeader257)

	// Sink the tainted `intoWriter439`:
	sink(intoWriter439)
}

func RunAllTaints_ArchiveTar() {
	{
		source := newSource()
		TaintStepTest_ArchiveTarFileInfoHeader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveTarNewReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveTarNewWriter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveTarHeaderFileInfo_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveTarReaderNext_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveTarReaderRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveTarWriterWrite_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ArchiveTarWriterWriteHeader_B0I0O0(source)
	}
}
