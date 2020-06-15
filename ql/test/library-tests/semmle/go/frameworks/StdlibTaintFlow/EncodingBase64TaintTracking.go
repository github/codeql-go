package main

import (
	"encoding/base64"
	"io"
)

func TaintStepTest_EncodingBase64NewDecoder_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader156` into `intoReader953`.

	// Assume that `sourceCQL` has the underlying type of `fromReader156`:
	fromReader156 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader156` to result `intoReader953`
	// (`intoReader953` is now tainted).
	intoReader953 := base64.NewDecoder(nil, fromReader156)

	// Sink the tainted `intoReader953`:
	sink(intoReader953)
}

func TaintStepTest_EncodingBase64NewEncoder_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriteCloser199` into `intoWriter209`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser199`:
	fromWriteCloser199 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter209` variable:
	var intoWriter209 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter209`:
	intermediateCQL := base64.NewEncoder(nil, intoWriter209)

	// Extra step (`fromWriteCloser199` taints `intermediateCQL`, which taints `intoWriter209`:
	link(fromWriteCloser199, intermediateCQL)

	// Sink the tainted `intoWriter209`:
	sink(intoWriter209)
}

func TaintStepTest_EncodingBase64EncodingDecode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte713` into `intoByte380`.

	// Assume that `sourceCQL` has the underlying type of `fromByte713`:
	fromByte713 := sourceCQL.([]byte)

	// Declare `intoByte380` variable:
	var intoByte380 []byte

	// Declare medium object/interface:
	var mediumObjCQL base64.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte713` to the parameter `intoByte380`
	// (`intoByte380` is now tainted).
	mediumObjCQL.Decode(intoByte380, fromByte713)

	// Sink the tainted `intoByte380`:
	sink(intoByte380)
}

func TaintStepTest_EncodingBase64EncodingDecodeString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString344` into `intoByte619`.

	// Assume that `sourceCQL` has the underlying type of `fromString344`:
	fromString344 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL base64.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromString344` to the result `intoByte619`
	// (`intoByte619` is now tainted).
	intoByte619, _ := mediumObjCQL.DecodeString(fromString344)

	// Sink the tainted `intoByte619`:
	sink(intoByte619)
}

func TaintStepTest_EncodingBase64EncodingEncode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte899` into `intoByte678`.

	// Assume that `sourceCQL` has the underlying type of `fromByte899`:
	fromByte899 := sourceCQL.([]byte)

	// Declare `intoByte678` variable:
	var intoByte678 []byte

	// Declare medium object/interface:
	var mediumObjCQL base64.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte899` to the parameter `intoByte678`
	// (`intoByte678` is now tainted).
	mediumObjCQL.Encode(intoByte678, fromByte899)

	// Sink the tainted `intoByte678`:
	sink(intoByte678)
}

func TaintStepTest_EncodingBase64EncodingEncodeToString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte398` into `intoString982`.

	// Assume that `sourceCQL` has the underlying type of `fromByte398`:
	fromByte398 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL base64.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte398` to the result `intoString982`
	// (`intoString982` is now tainted).
	intoString982 := mediumObjCQL.EncodeToString(fromByte398)

	// Sink the tainted `intoString982`:
	sink(intoString982)
}

func RunAllTaints_EncodingBase64() {
	{
		source := newSource()
		TaintStepTest_EncodingBase64NewDecoder_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase64NewEncoder_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase64EncodingDecode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase64EncodingDecodeString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase64EncodingEncode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase64EncodingEncodeToString_B0I0O0(source)
	}
}
