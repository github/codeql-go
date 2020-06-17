// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"encoding/base64"
	"io"
)

func TaintStepTest_EncodingBase64NewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader573` into `intoReader749`.

	// Assume that `sourceCQL` has the underlying type of `fromReader573`:
	fromReader573 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader573` to result `intoReader749`
	// (`intoReader749` is now tainted).
	intoReader749 := base64.NewDecoder(nil, fromReader573)

	// Return the tainted `intoReader749`:
	return intoReader749
}

func TaintStepTest_EncodingBase64NewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser388` into `intoWriter602`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser388`:
	fromWriteCloser388 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter602` variable:
	var intoWriter602 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter602`:
	intermediateCQL := base64.NewEncoder(nil, intoWriter602)

	// Extra step (`fromWriteCloser388` taints `intermediateCQL`, which taints `intoWriter602`:
	link(fromWriteCloser388, intermediateCQL)

	// Return the tainted `intoWriter602`:
	return intoWriter602
}

func TaintStepTest_EncodingBase64EncodingDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte971` into `intoByte158`.

	// Assume that `sourceCQL` has the underlying type of `fromByte971`:
	fromByte971 := sourceCQL.([]byte)

	// Declare `intoByte158` variable:
	var intoByte158 []byte

	// Declare medium object/interface:
	var mediumObjCQL base64.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte971` to the parameter `intoByte158`
	// (`intoByte158` is now tainted).
	mediumObjCQL.Decode(intoByte158, fromByte971)

	// Return the tainted `intoByte158`:
	return intoByte158
}

func TaintStepTest_EncodingBase64EncodingDecodeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString965` into `intoByte558`.

	// Assume that `sourceCQL` has the underlying type of `fromString965`:
	fromString965 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL base64.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromString965` to the result `intoByte558`
	// (`intoByte558` is now tainted).
	intoByte558, _ := mediumObjCQL.DecodeString(fromString965)

	// Return the tainted `intoByte558`:
	return intoByte558
}

func TaintStepTest_EncodingBase64EncodingEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte269` into `intoByte510`.

	// Assume that `sourceCQL` has the underlying type of `fromByte269`:
	fromByte269 := sourceCQL.([]byte)

	// Declare `intoByte510` variable:
	var intoByte510 []byte

	// Declare medium object/interface:
	var mediumObjCQL base64.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte269` to the parameter `intoByte510`
	// (`intoByte510` is now tainted).
	mediumObjCQL.Encode(intoByte510, fromByte269)

	// Return the tainted `intoByte510`:
	return intoByte510
}

func TaintStepTest_EncodingBase64EncodingEncodeToString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte379` into `intoString448`.

	// Assume that `sourceCQL` has the underlying type of `fromByte379`:
	fromByte379 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL base64.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte379` to the result `intoString448`
	// (`intoString448` is now tainted).
	intoString448 := mediumObjCQL.EncodeToString(fromByte379)

	// Return the tainted `intoString448`:
	return intoString448
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
