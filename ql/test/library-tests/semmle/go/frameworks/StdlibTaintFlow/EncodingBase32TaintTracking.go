package main

import (
	"encoding/base32"
	"io"
)

func TaintStepTest_EncodingBase32NewDecoder_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader829` into `intoReader444`.

	// Assume that `sourceCQL` has the underlying type of `fromReader829`:
	fromReader829 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader829` to result `intoReader444`
	// (`intoReader444` is now tainted).
	intoReader444 := base32.NewDecoder(nil, fromReader829)

	// Sink the tainted `intoReader444`:
	sink(intoReader444)
}

func TaintStepTest_EncodingBase32NewEncoder_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriteCloser914` into `intoWriter941`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser914`:
	fromWriteCloser914 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter941` variable:
	var intoWriter941 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter941`:
	intermediateCQL := base32.NewEncoder(nil, intoWriter941)

	// Extra step (`fromWriteCloser914` taints `intermediateCQL`, which taints `intoWriter941`:
	link(fromWriteCloser914, intermediateCQL)

	// Sink the tainted `intoWriter941`:
	sink(intoWriter941)
}

func TaintStepTest_EncodingBase32EncodingDecode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte312` into `intoByte628`.

	// Assume that `sourceCQL` has the underlying type of `fromByte312`:
	fromByte312 := sourceCQL.([]byte)

	// Declare `intoByte628` variable:
	var intoByte628 []byte

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte312` to the parameter `intoByte628`
	// (`intoByte628` is now tainted).
	mediumObjCQL.Decode(intoByte628, fromByte312)

	// Sink the tainted `intoByte628`:
	sink(intoByte628)
}

func TaintStepTest_EncodingBase32EncodingDecodeString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString136` into `intoByte366`.

	// Assume that `sourceCQL` has the underlying type of `fromString136`:
	fromString136 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromString136` to the result `intoByte366`
	// (`intoByte366` is now tainted).
	intoByte366, _ := mediumObjCQL.DecodeString(fromString136)

	// Sink the tainted `intoByte366`:
	sink(intoByte366)
}

func TaintStepTest_EncodingBase32EncodingEncode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte168` into `intoByte836`.

	// Assume that `sourceCQL` has the underlying type of `fromByte168`:
	fromByte168 := sourceCQL.([]byte)

	// Declare `intoByte836` variable:
	var intoByte836 []byte

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte168` to the parameter `intoByte836`
	// (`intoByte836` is now tainted).
	mediumObjCQL.Encode(intoByte836, fromByte168)

	// Sink the tainted `intoByte836`:
	sink(intoByte836)
}

func TaintStepTest_EncodingBase32EncodingEncodeToString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte764` into `intoString885`.

	// Assume that `sourceCQL` has the underlying type of `fromByte764`:
	fromByte764 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte764` to the result `intoString885`
	// (`intoString885` is now tainted).
	intoString885 := mediumObjCQL.EncodeToString(fromByte764)

	// Sink the tainted `intoString885`:
	sink(intoString885)
}

func RunAllTaints_EncodingBase32() {
	{
		source := newSource()
		TaintStepTest_EncodingBase32NewDecoder_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase32NewEncoder_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase32EncodingDecode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase32EncodingDecodeString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase32EncodingEncode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase32EncodingEncodeToString_B0I0O0(source)
	}
}
