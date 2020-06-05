package main

import (
	"encoding/binary"
	"io"
)

func TaintStepTest_EncodingBinaryPutUvarint(sourceCQL interface{}) {
	// The flow is from `x` into `buf`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.(uint64)

	// Declare `buf` variable:
	var buf []byte

	// Call the function that transfers the taint
	// from the parameter `x` to parameter `buf`;
	// `buf` is now tainted.
	binary.PutUvarint(buf, x)

	// Sink the tainted `buf`:
	sink(buf)
}

func TaintStepTest_EncodingBinaryPutVarint(sourceCQL interface{}) {
	// The flow is from `x` into `buf`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.(int64)

	// Declare `buf` variable:
	var buf []byte

	// Call the function that transfers the taint
	// from the parameter `x` to parameter `buf`;
	// `buf` is now tainted.
	binary.PutVarint(buf, x)

	// Sink the tainted `buf`:
	sink(buf)
}

func TaintStepTest_EncodingBinaryRead(sourceCQL interface{}) {
	// The flow is from `r` into `data`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Declare `data` variable:
	var data interface{}

	// Call the function that transfers the taint
	// from the parameter `r` to parameter `data`;
	// `data` is now tainted.
	binary.Read(r, nil, data)

	// Sink the tainted `data`:
	sink(data)
}

func TaintStepTest_EncodingBinaryReadUvarint(sourceCQL interface{}) {
	// The flow is from `r` into `into386`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.ByteReader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into386`
	// (`into386` is now tainted).
	into386, _ := binary.ReadUvarint(r)

	// Sink the tainted `into386`:
	sink(into386)
}

func TaintStepTest_EncodingBinaryReadVarint(sourceCQL interface{}) {
	// The flow is from `r` into `into804`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.ByteReader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into804`
	// (`into804` is now tainted).
	into804, _ := binary.ReadVarint(r)

	// Sink the tainted `into804`:
	sink(into804)
}

func TaintStepTest_EncodingBinaryUvarint(sourceCQL interface{}) {
	// The flow is from `buf` into `into125`.

	// Assume that `sourceCQL` has the underlying type of `buf`:
	buf := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `buf` to result `into125`
	// (`into125` is now tainted).
	into125, _ := binary.Uvarint(buf)

	// Sink the tainted `into125`:
	sink(into125)
}

func TaintStepTest_EncodingBinaryVarint(sourceCQL interface{}) {
	// The flow is from `buf` into `into205`.

	// Assume that `sourceCQL` has the underlying type of `buf`:
	buf := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `buf` to result `into205`
	// (`into205` is now tainted).
	into205, _ := binary.Varint(buf)

	// Sink the tainted `into205`:
	sink(into205)
}

func TaintStepTest_EncodingBinaryWrite(sourceCQL interface{}) {
	// The flow is from `data` into `w`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.(interface{})

	// Declare `w` variable:
	var w io.Writer

	// Call the function that transfers the taint
	// from the parameter `data` to parameter `w`;
	// `w` is now tainted.
	binary.Write(w, nil, data)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_EncodingBinaryByteOrderPutUint16(sourceCQL interface{}) {
	// The flow is from `from152` into `into358`.

	// Assume that `sourceCQL` has the underlying type of `from152`:
	from152 := sourceCQL.(uint16)

	// Declare `into358` variable:
	var into358 []byte

	// Declare medium object/interface:
	var mediumObjCQL binary.ByteOrder

	// Call the method that transfers the taint
	// from the parameter `from152` to the parameter `into358`
	// (`into358` is now tainted).
	mediumObjCQL.PutUint16(into358, from152)

	// Sink the tainted `into358`:
	sink(into358)
}

func TaintStepTest_EncodingBinaryByteOrderPutUint32(sourceCQL interface{}) {
	// The flow is from `from906` into `into814`.

	// Assume that `sourceCQL` has the underlying type of `from906`:
	from906 := sourceCQL.(uint32)

	// Declare `into814` variable:
	var into814 []byte

	// Declare medium object/interface:
	var mediumObjCQL binary.ByteOrder

	// Call the method that transfers the taint
	// from the parameter `from906` to the parameter `into814`
	// (`into814` is now tainted).
	mediumObjCQL.PutUint32(into814, from906)

	// Sink the tainted `into814`:
	sink(into814)
}

func TaintStepTest_EncodingBinaryByteOrderPutUint64(sourceCQL interface{}) {
	// The flow is from `from936` into `into566`.

	// Assume that `sourceCQL` has the underlying type of `from936`:
	from936 := sourceCQL.(uint64)

	// Declare `into566` variable:
	var into566 []byte

	// Declare medium object/interface:
	var mediumObjCQL binary.ByteOrder

	// Call the method that transfers the taint
	// from the parameter `from936` to the parameter `into566`
	// (`into566` is now tainted).
	mediumObjCQL.PutUint64(into566, from936)

	// Sink the tainted `into566`:
	sink(into566)
}

func TaintStepTest_EncodingBinaryByteOrderString(sourceCQL interface{}) {
	// The flow is from `from140` into `into966`.

	// Assume that `sourceCQL` has the underlying type of `from140`:
	from140 := sourceCQL.(binary.ByteOrder)

	// Call the method that transfers the taint
	// from the receiver `from140` to the result `into966`
	// (`into966` is now tainted).
	into966 := from140.String()

	// Sink the tainted `into966`:
	sink(into966)
}

func TaintStepTest_EncodingBinaryByteOrderUint16(sourceCQL interface{}) {
	// The flow is from `from422` into `into238`.

	// Assume that `sourceCQL` has the underlying type of `from422`:
	from422 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL binary.ByteOrder

	// Call the method that transfers the taint
	// from the parameter `from422` to the result `into238`
	// (`into238` is now tainted).
	into238 := mediumObjCQL.Uint16(from422)

	// Sink the tainted `into238`:
	sink(into238)
}

func TaintStepTest_EncodingBinaryByteOrderUint32(sourceCQL interface{}) {
	// The flow is from `from149` into `into846`.

	// Assume that `sourceCQL` has the underlying type of `from149`:
	from149 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL binary.ByteOrder

	// Call the method that transfers the taint
	// from the parameter `from149` to the result `into846`
	// (`into846` is now tainted).
	into846 := mediumObjCQL.Uint32(from149)

	// Sink the tainted `into846`:
	sink(into846)
}

func TaintStepTest_EncodingBinaryByteOrderUint64(sourceCQL interface{}) {
	// The flow is from `from808` into `into922`.

	// Assume that `sourceCQL` has the underlying type of `from808`:
	from808 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL binary.ByteOrder

	// Call the method that transfers the taint
	// from the parameter `from808` to the result `into922`
	// (`into922` is now tainted).
	into922 := mediumObjCQL.Uint64(from808)

	// Sink the tainted `into922`:
	sink(into922)
}

func RunAllTaints_EncodingBinary(v interface{}) {
	{
		source := newSource()
		TaintStepTest_EncodingBinaryPutUvarint(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryPutVarint(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryRead(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryReadUvarint(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryReadVarint(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryUvarint(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryVarint(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryWrite(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryByteOrderPutUint16(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryByteOrderPutUint32(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryByteOrderPutUint64(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryByteOrderString(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryByteOrderUint16(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryByteOrderUint32(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBinaryByteOrderUint64(source)
	}
}
