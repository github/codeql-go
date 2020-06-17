// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"encoding/base32"
	"io"
)

func TaintStepTest_EncodingBase32NewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader656` into `intoReader414`.

	// Assume that `sourceCQL` has the underlying type of `fromReader656`:
	fromReader656 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader656` to result `intoReader414`
	// (`intoReader414` is now tainted).
	intoReader414 := base32.NewDecoder(nil, fromReader656)

	// Return the tainted `intoReader414`:
	return intoReader414
}

func TaintStepTest_EncodingBase32NewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser518` into `intoWriter650`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser518`:
	fromWriteCloser518 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter650` variable:
	var intoWriter650 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter650`:
	intermediateCQL := base32.NewEncoder(nil, intoWriter650)

	// Extra step (`fromWriteCloser518` taints `intermediateCQL`, which taints `intoWriter650`:
	link(fromWriteCloser518, intermediateCQL)

	// Return the tainted `intoWriter650`:
	return intoWriter650
}

func TaintStepTest_EncodingBase32EncodingDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte784` into `intoByte957`.

	// Assume that `sourceCQL` has the underlying type of `fromByte784`:
	fromByte784 := sourceCQL.([]byte)

	// Declare `intoByte957` variable:
	var intoByte957 []byte

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte784` to the parameter `intoByte957`
	// (`intoByte957` is now tainted).
	mediumObjCQL.Decode(intoByte957, fromByte784)

	// Return the tainted `intoByte957`:
	return intoByte957
}

func TaintStepTest_EncodingBase32EncodingDecodeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString520` into `intoByte443`.

	// Assume that `sourceCQL` has the underlying type of `fromString520`:
	fromString520 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromString520` to the result `intoByte443`
	// (`intoByte443` is now tainted).
	intoByte443, _ := mediumObjCQL.DecodeString(fromString520)

	// Return the tainted `intoByte443`:
	return intoByte443
}

func TaintStepTest_EncodingBase32EncodingEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte127` into `intoByte483`.

	// Assume that `sourceCQL` has the underlying type of `fromByte127`:
	fromByte127 := sourceCQL.([]byte)

	// Declare `intoByte483` variable:
	var intoByte483 []byte

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte127` to the parameter `intoByte483`
	// (`intoByte483` is now tainted).
	mediumObjCQL.Encode(intoByte483, fromByte127)

	// Return the tainted `intoByte483`:
	return intoByte483
}

func TaintStepTest_EncodingBase32EncodingEncodeToString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte989` into `intoString982`.

	// Assume that `sourceCQL` has the underlying type of `fromByte989`:
	fromByte989 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte989` to the result `intoString982`
	// (`intoString982` is now tainted).
	intoString982 := mediumObjCQL.EncodeToString(fromByte989)

	// Return the tainted `intoString982`:
	return intoString982
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
