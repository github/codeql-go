// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"encoding/binary"
	"io"
)

func TaintStepTest_EncodingBinaryPutUvarint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUint64656` into `intoByte414`.

	// Assume that `sourceCQL` has the underlying type of `fromUint64656`:
	fromUint64656 := sourceCQL.(uint64)

	// Declare `intoByte414` variable:
	var intoByte414 []byte

	// Call the function that transfers the taint
	// from the parameter `fromUint64656` to parameter `intoByte414`;
	// `intoByte414` is now tainted.
	binary.PutUvarint(intoByte414, fromUint64656)

	// Return the tainted `intoByte414`:
	return intoByte414
}

func TaintStepTest_EncodingBinaryRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader518` into `intoInterface650`.

	// Assume that `sourceCQL` has the underlying type of `fromReader518`:
	fromReader518 := sourceCQL.(io.Reader)

	// Declare `intoInterface650` variable:
	var intoInterface650 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromReader518` to parameter `intoInterface650`;
	// `intoInterface650` is now tainted.
	binary.Read(fromReader518, nil, intoInterface650)

	// Return the tainted `intoInterface650`:
	return intoInterface650
}

func TaintStepTest_EncodingBinaryReadUvarint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByteReader784` into `intoUint64957`.

	// Assume that `sourceCQL` has the underlying type of `fromByteReader784`:
	fromByteReader784 := sourceCQL.(io.ByteReader)

	// Call the function that transfers the taint
	// from the parameter `fromByteReader784` to result `intoUint64957`
	// (`intoUint64957` is now tainted).
	intoUint64957, _ := binary.ReadUvarint(fromByteReader784)

	// Return the tainted `intoUint64957`:
	return intoUint64957
}

func TaintStepTest_EncodingBinaryWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface520` into `intoWriter443`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface520`:
	fromInterface520 := sourceCQL.(interface{})

	// Declare `intoWriter443` variable:
	var intoWriter443 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromInterface520` to parameter `intoWriter443`;
	// `intoWriter443` is now tainted.
	binary.Write(intoWriter443, nil, fromInterface520)

	// Return the tainted `intoWriter443`:
	return intoWriter443
}

func TaintStepTest_EncodingBinaryByteOrderString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByteOrder127` into `intoString483`.

	// Assume that `sourceCQL` has the underlying type of `fromByteOrder127`:
	fromByteOrder127 := sourceCQL.(binary.ByteOrder)

	// Call the method that transfers the taint
	// from the receiver `fromByteOrder127` to the result `intoString483`
	// (`intoString483` is now tainted).
	intoString483 := fromByteOrder127.String()

	// Return the tainted `intoString483`:
	return intoString483
}

func RunAllTaints_EncodingBinary() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBinaryPutUvarint_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBinaryRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBinaryReadUvarint_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBinaryWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBinaryByteOrderString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
