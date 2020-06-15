package main

import (
	"encoding/hex"
	"io"
)

func TaintStepTest_EncodingHexDecode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte768` into `intoByte535`.

	// Assume that `sourceCQL` has the underlying type of `fromByte768`:
	fromByte768 := sourceCQL.([]byte)

	// Declare `intoByte535` variable:
	var intoByte535 []byte

	// Call the function that transfers the taint
	// from the parameter `fromByte768` to parameter `intoByte535`;
	// `intoByte535` is now tainted.
	hex.Decode(intoByte535, fromByte768)

	// Sink the tainted `intoByte535`:
	sink(intoByte535)
}

func TaintStepTest_EncodingHexDecodeString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString777` into `intoByte696`.

	// Assume that `sourceCQL` has the underlying type of `fromString777`:
	fromString777 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString777` to result `intoByte696`
	// (`intoByte696` is now tainted).
	intoByte696, _ := hex.DecodeString(fromString777)

	// Sink the tainted `intoByte696`:
	sink(intoByte696)
}

func TaintStepTest_EncodingHexDump_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte616` into `intoString457`.

	// Assume that `sourceCQL` has the underlying type of `fromByte616`:
	fromByte616 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte616` to result `intoString457`
	// (`intoString457` is now tainted).
	intoString457 := hex.Dump(fromByte616)

	// Sink the tainted `intoString457`:
	sink(intoString457)
}

func TaintStepTest_EncodingHexDumper_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriteCloser472` into `intoWriter173`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser472`:
	fromWriteCloser472 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter173` variable:
	var intoWriter173 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter173`:
	intermediateCQL := hex.Dumper(intoWriter173)

	// Extra step (`fromWriteCloser472` taints `intermediateCQL`, which taints `intoWriter173`:
	link(fromWriteCloser472, intermediateCQL)

	// Sink the tainted `intoWriter173`:
	sink(intoWriter173)
}

func TaintStepTest_EncodingHexEncode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte808` into `intoByte456`.

	// Assume that `sourceCQL` has the underlying type of `fromByte808`:
	fromByte808 := sourceCQL.([]byte)

	// Declare `intoByte456` variable:
	var intoByte456 []byte

	// Call the function that transfers the taint
	// from the parameter `fromByte808` to parameter `intoByte456`;
	// `intoByte456` is now tainted.
	hex.Encode(intoByte456, fromByte808)

	// Sink the tainted `intoByte456`:
	sink(intoByte456)
}

func TaintStepTest_EncodingHexEncodeToString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte815` into `intoString326`.

	// Assume that `sourceCQL` has the underlying type of `fromByte815`:
	fromByte815 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte815` to result `intoString326`
	// (`intoString326` is now tainted).
	intoString326 := hex.EncodeToString(fromByte815)

	// Sink the tainted `intoString326`:
	sink(intoString326)
}

func TaintStepTest_EncodingHexNewDecoder_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader212` into `intoReader606`.

	// Assume that `sourceCQL` has the underlying type of `fromReader212`:
	fromReader212 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader212` to result `intoReader606`
	// (`intoReader606` is now tainted).
	intoReader606 := hex.NewDecoder(fromReader212)

	// Sink the tainted `intoReader606`:
	sink(intoReader606)
}

func TaintStepTest_EncodingHexNewEncoder_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter543` into `intoWriter307`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter543`:
	fromWriter543 := sourceCQL.(io.Writer)

	// Declare `intoWriter307` variable:
	var intoWriter307 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter307`:
	intermediateCQL := hex.NewEncoder(intoWriter307)

	// Extra step (`fromWriter543` taints `intermediateCQL`, which taints `intoWriter307`:
	link(fromWriter543, intermediateCQL)

	// Sink the tainted `intoWriter307`:
	sink(intoWriter307)
}

func RunAllTaints_EncodingHex() {
	{
		source := newSource()
		TaintStepTest_EncodingHexDecode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingHexDecodeString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingHexDump_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingHexDumper_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingHexEncode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingHexEncodeToString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingHexNewDecoder_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingHexNewEncoder_B0I0O0(source)
	}
}
