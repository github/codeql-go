// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"encoding/hex"
	"io"
)

func TaintStepTest_EncodingHexDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte656` into `intoByte414`.

	// Assume that `sourceCQL` has the underlying type of `fromByte656`:
	fromByte656 := sourceCQL.([]byte)

	// Declare `intoByte414` variable:
	var intoByte414 []byte

	// Call the function that transfers the taint
	// from the parameter `fromByte656` to parameter `intoByte414`;
	// `intoByte414` is now tainted.
	hex.Decode(intoByte414, fromByte656)

	// Return the tainted `intoByte414`:
	return intoByte414
}

func TaintStepTest_EncodingHexDecodeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString518` into `intoByte650`.

	// Assume that `sourceCQL` has the underlying type of `fromString518`:
	fromString518 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString518` to result `intoByte650`
	// (`intoByte650` is now tainted).
	intoByte650, _ := hex.DecodeString(fromString518)

	// Return the tainted `intoByte650`:
	return intoByte650
}

func TaintStepTest_EncodingHexDump_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte784` into `intoString957`.

	// Assume that `sourceCQL` has the underlying type of `fromByte784`:
	fromByte784 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte784` to result `intoString957`
	// (`intoString957` is now tainted).
	intoString957 := hex.Dump(fromByte784)

	// Return the tainted `intoString957`:
	return intoString957
}

func TaintStepTest_EncodingHexDumper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser520` into `intoWriter443`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser520`:
	fromWriteCloser520 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter443` variable:
	var intoWriter443 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter443`:
	intermediateCQL := hex.Dumper(intoWriter443)

	// Extra step (`fromWriteCloser520` taints `intermediateCQL`, which taints `intoWriter443`:
	link(fromWriteCloser520, intermediateCQL)

	// Return the tainted `intoWriter443`:
	return intoWriter443
}

func TaintStepTest_EncodingHexEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte127` into `intoByte483`.

	// Assume that `sourceCQL` has the underlying type of `fromByte127`:
	fromByte127 := sourceCQL.([]byte)

	// Declare `intoByte483` variable:
	var intoByte483 []byte

	// Call the function that transfers the taint
	// from the parameter `fromByte127` to parameter `intoByte483`;
	// `intoByte483` is now tainted.
	hex.Encode(intoByte483, fromByte127)

	// Return the tainted `intoByte483`:
	return intoByte483
}

func TaintStepTest_EncodingHexEncodeToString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte989` into `intoString982`.

	// Assume that `sourceCQL` has the underlying type of `fromByte989`:
	fromByte989 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte989` to result `intoString982`
	// (`intoString982` is now tainted).
	intoString982 := hex.EncodeToString(fromByte989)

	// Return the tainted `intoString982`:
	return intoString982
}

func TaintStepTest_EncodingHexNewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader417` into `intoReader584`.

	// Assume that `sourceCQL` has the underlying type of `fromReader417`:
	fromReader417 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader417` to result `intoReader584`
	// (`intoReader584` is now tainted).
	intoReader584 := hex.NewDecoder(fromReader417)

	// Return the tainted `intoReader584`:
	return intoReader584
}

func TaintStepTest_EncodingHexNewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter991` into `intoWriter881`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter991`:
	fromWriter991 := sourceCQL.(io.Writer)

	// Declare `intoWriter881` variable:
	var intoWriter881 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter881`:
	intermediateCQL := hex.NewEncoder(intoWriter881)

	// Extra step (`fromWriter991` taints `intermediateCQL`, which taints `intoWriter881`:
	link(fromWriter991, intermediateCQL)

	// Return the tainted `intoWriter881`:
	return intoWriter881
}

func RunAllTaints_EncodingHex() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingHexDecode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingHexDecodeString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingHexDump_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingHexDumper_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingHexEncode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingHexEncodeToString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingHexNewDecoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingHexNewEncoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
