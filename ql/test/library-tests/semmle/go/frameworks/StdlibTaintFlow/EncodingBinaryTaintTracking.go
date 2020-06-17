// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"encoding/binary"
	"io"
)

func TaintStepTest_EncodingBinaryPutUvarint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUint64620` into `intoByte980`.

	// Assume that `sourceCQL` has the underlying type of `fromUint64620`:
	fromUint64620 := sourceCQL.(uint64)

	// Declare `intoByte980` variable:
	var intoByte980 []byte

	// Call the function that transfers the taint
	// from the parameter `fromUint64620` to parameter `intoByte980`;
	// `intoByte980` is now tainted.
	binary.PutUvarint(intoByte980, fromUint64620)

	// Return the tainted `intoByte980`:
	return intoByte980
}

func TaintStepTest_EncodingBinaryRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader902` into `intoInterface418`.

	// Assume that `sourceCQL` has the underlying type of `fromReader902`:
	fromReader902 := sourceCQL.(io.Reader)

	// Declare `intoInterface418` variable:
	var intoInterface418 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromReader902` to parameter `intoInterface418`;
	// `intoInterface418` is now tainted.
	binary.Read(fromReader902, nil, intoInterface418)

	// Return the tainted `intoInterface418`:
	return intoInterface418
}

func TaintStepTest_EncodingBinaryReadUvarint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByteReader703` into `intoUint64256`.

	// Assume that `sourceCQL` has the underlying type of `fromByteReader703`:
	fromByteReader703 := sourceCQL.(io.ByteReader)

	// Call the function that transfers the taint
	// from the parameter `fromByteReader703` to result `intoUint64256`
	// (`intoUint64256` is now tainted).
	intoUint64256, _ := binary.ReadUvarint(fromByteReader703)

	// Return the tainted `intoUint64256`:
	return intoUint64256
}

func TaintStepTest_EncodingBinaryWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface503` into `intoWriter150`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface503`:
	fromInterface503 := sourceCQL.(interface{})

	// Declare `intoWriter150` variable:
	var intoWriter150 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromInterface503` to parameter `intoWriter150`;
	// `intoWriter150` is now tainted.
	binary.Write(intoWriter150, nil, fromInterface503)

	// Return the tainted `intoWriter150`:
	return intoWriter150
}

func TaintStepTest_EncodingBinaryByteOrderString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByteOrder563` into `intoString810`.

	// Assume that `sourceCQL` has the underlying type of `fromByteOrder563`:
	fromByteOrder563 := sourceCQL.(binary.ByteOrder)

	// Call the method that transfers the taint
	// from the receiver `fromByteOrder563` to the result `intoString810`
	// (`intoString810` is now tainted).
	intoString810 := fromByteOrder563.String()

	// Return the tainted `intoString810`:
	return intoString810
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
