package main

import (
	"encoding/binary"
	"io"
)

func TaintStepTest_EncodingBinaryPutUvarint_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUint64946` into `intoByte695`.

	// Assume that `sourceCQL` has the underlying type of `fromUint64946`:
	fromUint64946 := sourceCQL.(uint64)

	// Declare `intoByte695` variable:
	var intoByte695 []byte

	// Call the function that transfers the taint
	// from the parameter `fromUint64946` to parameter `intoByte695`;
	// `intoByte695` is now tainted.
	binary.PutUvarint(intoByte695, fromUint64946)

	// Sink the tainted `intoByte695`:
	sink(intoByte695)
}

func TaintStepTest_EncodingBinaryRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader454` into `intoInterface584`.

	// Assume that `sourceCQL` has the underlying type of `fromReader454`:
	fromReader454 := sourceCQL.(io.Reader)

	// Declare `intoInterface584` variable:
	var intoInterface584 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromReader454` to parameter `intoInterface584`;
	// `intoInterface584` is now tainted.
	binary.Read(fromReader454, nil, intoInterface584)

	// Sink the tainted `intoInterface584`:
	sink(intoInterface584)
}

func TaintStepTest_EncodingBinaryReadUvarint_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByteReader841` into `intoUint64739`.

	// Assume that `sourceCQL` has the underlying type of `fromByteReader841`:
	fromByteReader841 := sourceCQL.(io.ByteReader)

	// Call the function that transfers the taint
	// from the parameter `fromByteReader841` to result `intoUint64739`
	// (`intoUint64739` is now tainted).
	intoUint64739, _ := binary.ReadUvarint(fromByteReader841)

	// Sink the tainted `intoUint64739`:
	sink(intoUint64739)
}

func TaintStepTest_EncodingBinaryWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface278` into `intoWriter868`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface278`:
	fromInterface278 := sourceCQL.(interface{})

	// Declare `intoWriter868` variable:
	var intoWriter868 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromInterface278` to parameter `intoWriter868`;
	// `intoWriter868` is now tainted.
	binary.Write(intoWriter868, nil, fromInterface278)

	// Sink the tainted `intoWriter868`:
	sink(intoWriter868)
}

func TaintStepTest_EncodingBinaryByteOrderString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByteOrder822` into `intoString816`.

	// Assume that `sourceCQL` has the underlying type of `fromByteOrder822`:
	fromByteOrder822 := sourceCQL.(binary.ByteOrder)

	// Call the method that transfers the taint
	// from the receiver `fromByteOrder822` to the result `intoString816`
	// (`intoString816` is now tainted).
	intoString816 := fromByteOrder822.String()

	// Sink the tainted `intoString816`:
	sink(intoString816)
}

func RunAllTaints_EncodingBinary() {
	{
		source := newSource()
		TaintStepTest_EncodingBinaryPutUvarint_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryReadUvarint_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryWrite_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryByteOrderString_B0I0O0(source)
	}
}
