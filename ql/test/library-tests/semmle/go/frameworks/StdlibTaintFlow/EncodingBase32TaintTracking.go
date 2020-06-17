// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"encoding/base32"
	"io"
)

func TaintStepTest_EncodingBase32NewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader112` into `intoReader228`.

	// Assume that `sourceCQL` has the underlying type of `fromReader112`:
	fromReader112 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader112` to result `intoReader228`
	// (`intoReader228` is now tainted).
	intoReader228 := base32.NewDecoder(nil, fromReader112)

	// Return the tainted `intoReader228`:
	return intoReader228
}

func TaintStepTest_EncodingBase32NewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser876` into `intoWriter333`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser876`:
	fromWriteCloser876 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter333` variable:
	var intoWriter333 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter333`:
	intermediateCQL := base32.NewEncoder(nil, intoWriter333)

	// Extra step (`fromWriteCloser876` taints `intermediateCQL`, which taints `intoWriter333`:
	link(fromWriteCloser876, intermediateCQL)

	// Return the tainted `intoWriter333`:
	return intoWriter333
}

func TaintStepTest_EncodingBase32EncodingDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte753` into `intoByte126`.

	// Assume that `sourceCQL` has the underlying type of `fromByte753`:
	fromByte753 := sourceCQL.([]byte)

	// Declare `intoByte126` variable:
	var intoByte126 []byte

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte753` to the parameter `intoByte126`
	// (`intoByte126` is now tainted).
	mediumObjCQL.Decode(intoByte126, fromByte753)

	// Return the tainted `intoByte126`:
	return intoByte126
}

func TaintStepTest_EncodingBase32EncodingDecodeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString403` into `intoByte540`.

	// Assume that `sourceCQL` has the underlying type of `fromString403`:
	fromString403 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromString403` to the result `intoByte540`
	// (`intoByte540` is now tainted).
	intoByte540, _ := mediumObjCQL.DecodeString(fromString403)

	// Return the tainted `intoByte540`:
	return intoByte540
}

func TaintStepTest_EncodingBase32EncodingEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte922` into `intoByte216`.

	// Assume that `sourceCQL` has the underlying type of `fromByte922`:
	fromByte922 := sourceCQL.([]byte)

	// Declare `intoByte216` variable:
	var intoByte216 []byte

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte922` to the parameter `intoByte216`
	// (`intoByte216` is now tainted).
	mediumObjCQL.Encode(intoByte216, fromByte922)

	// Return the tainted `intoByte216`:
	return intoByte216
}

func TaintStepTest_EncodingBase32EncodingEncodeToString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte628` into `intoString976`.

	// Assume that `sourceCQL` has the underlying type of `fromByte628`:
	fromByte628 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte628` to the result `intoString976`
	// (`intoString976` is now tainted).
	intoString976 := mediumObjCQL.EncodeToString(fromByte628)

	// Return the tainted `intoString976`:
	return intoString976
}

func RunAllTaints_EncodingBase32() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase32NewDecoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase32NewEncoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase32EncodingDecode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase32EncodingDecodeString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase32EncodingEncode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase32EncodingEncodeToString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
