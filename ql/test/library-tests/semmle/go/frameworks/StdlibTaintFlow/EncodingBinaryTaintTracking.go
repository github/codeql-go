package main

import (
	"encoding/binary"
	"io"
)

func TaintStepTest_EncodingBinaryPutUvarint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUint64898` into `intoByte709`.

	// Assume that `sourceCQL` has the underlying type of `fromUint64898`:
	fromUint64898 := sourceCQL.(uint64)

	// Declare `intoByte709` variable:
	var intoByte709 []byte

	// Call the function that transfers the taint
	// from the parameter `fromUint64898` to parameter `intoByte709`;
	// `intoByte709` is now tainted.
	binary.PutUvarint(intoByte709, fromUint64898)

	// Return the tainted `intoByte709`:
	return intoByte709
}

func TaintStepTest_EncodingBinaryRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader816` into `intoInterface817`.

	// Assume that `sourceCQL` has the underlying type of `fromReader816`:
	fromReader816 := sourceCQL.(io.Reader)

	// Declare `intoInterface817` variable:
	var intoInterface817 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromReader816` to parameter `intoInterface817`;
	// `intoInterface817` is now tainted.
	binary.Read(fromReader816, nil, intoInterface817)

	// Return the tainted `intoInterface817`:
	return intoInterface817
}

func TaintStepTest_EncodingBinaryReadUvarint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByteReader542` into `intoUint64566`.

	// Assume that `sourceCQL` has the underlying type of `fromByteReader542`:
	fromByteReader542 := sourceCQL.(io.ByteReader)

	// Call the function that transfers the taint
	// from the parameter `fromByteReader542` to result `intoUint64566`
	// (`intoUint64566` is now tainted).
	intoUint64566, _ := binary.ReadUvarint(fromByteReader542)

	// Return the tainted `intoUint64566`:
	return intoUint64566
}

func TaintStepTest_EncodingBinaryWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface243` into `intoWriter828`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface243`:
	fromInterface243 := sourceCQL.(interface{})

	// Declare `intoWriter828` variable:
	var intoWriter828 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromInterface243` to parameter `intoWriter828`;
	// `intoWriter828` is now tainted.
	binary.Write(intoWriter828, nil, fromInterface243)

	// Return the tainted `intoWriter828`:
	return intoWriter828
}

func TaintStepTest_EncodingBinaryByteOrderString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByteOrder704` into `intoString742`.

	// Assume that `sourceCQL` has the underlying type of `fromByteOrder704`:
	fromByteOrder704 := sourceCQL.(binary.ByteOrder)

	// Call the method that transfers the taint
	// from the receiver `fromByteOrder704` to the result `intoString742`
	// (`intoString742` is now tainted).
	intoString742 := fromByteOrder704.String()

	// Return the tainted `intoString742`:
	return intoString742
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
