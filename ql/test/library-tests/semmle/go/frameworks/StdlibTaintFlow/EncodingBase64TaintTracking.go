package main

import (
	"encoding/base64"
	"io"
)

func TaintStepTest_EncodingBase64NewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader915` into `intoReader338`.

	// Assume that `sourceCQL` has the underlying type of `fromReader915`:
	fromReader915 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader915` to result `intoReader338`
	// (`intoReader338` is now tainted).
	intoReader338 := base64.NewDecoder(nil, fromReader915)

	// Return the tainted `intoReader338`:
	return intoReader338
}

func TaintStepTest_EncodingBase64NewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser370` into `intoWriter568`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser370`:
	fromWriteCloser370 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter568` variable:
	var intoWriter568 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter568`:
	intermediateCQL := base64.NewEncoder(nil, intoWriter568)

	// Extra step (`fromWriteCloser370` taints `intermediateCQL`, which taints `intoWriter568`:
	link(fromWriteCloser370, intermediateCQL)

	// Return the tainted `intoWriter568`:
	return intoWriter568
}

func TaintStepTest_EncodingBase64EncodingDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte237` into `intoByte952`.

	// Assume that `sourceCQL` has the underlying type of `fromByte237`:
	fromByte237 := sourceCQL.([]byte)

	// Declare `intoByte952` variable:
	var intoByte952 []byte

	// Declare medium object/interface:
	var mediumObjCQL base64.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte237` to the parameter `intoByte952`
	// (`intoByte952` is now tainted).
	mediumObjCQL.Decode(intoByte952, fromByte237)

	// Return the tainted `intoByte952`:
	return intoByte952
}

func TaintStepTest_EncodingBase64EncodingDecodeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString772` into `intoByte303`.

	// Assume that `sourceCQL` has the underlying type of `fromString772`:
	fromString772 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL base64.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromString772` to the result `intoByte303`
	// (`intoByte303` is now tainted).
	intoByte303, _ := mediumObjCQL.DecodeString(fromString772)

	// Return the tainted `intoByte303`:
	return intoByte303
}

func TaintStepTest_EncodingBase64EncodingEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte246` into `intoByte899`.

	// Assume that `sourceCQL` has the underlying type of `fromByte246`:
	fromByte246 := sourceCQL.([]byte)

	// Declare `intoByte899` variable:
	var intoByte899 []byte

	// Declare medium object/interface:
	var mediumObjCQL base64.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte246` to the parameter `intoByte899`
	// (`intoByte899` is now tainted).
	mediumObjCQL.Encode(intoByte899, fromByte246)

	// Return the tainted `intoByte899`:
	return intoByte899
}

func TaintStepTest_EncodingBase64EncodingEncodeToString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte709` into `intoString155`.

	// Assume that `sourceCQL` has the underlying type of `fromByte709`:
	fromByte709 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL base64.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte709` to the result `intoString155`
	// (`intoString155` is now tainted).
	intoString155 := mediumObjCQL.EncodeToString(fromByte709)

	// Return the tainted `intoString155`:
	return intoString155
}

func RunAllTaints_EncodingBase64() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase64NewDecoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase64NewEncoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase64EncodingDecode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase64EncodingDecodeString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase64EncodingEncode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase64EncodingEncodeToString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
