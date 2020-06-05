package main

import (
	"encoding/gob"
	"io"
	"reflect"
)

func TaintStepTest_EncodingGobNewDecoder(sourceCQL interface{}) {
	// The flow is from `r` into `into964`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into964`
	// (`into964` is now tainted).
	into964 := gob.NewDecoder(r)

	// Sink the tainted `into964`:
	sink(into964)
}

func TaintStepTest_EncodingGobNewEncoder(sourceCQL interface{}) {
	// The flow is from `from118` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from118`:
	from118 := sourceCQL.(*gob.Encoder)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := gob.NewEncoder(w)

	// Extra step (`from118` taints `intermediateCQL`, which taints `w`:
	link(from118, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_EncodingGobDecoderDecode(sourceCQL interface{}) {
	// The flow is from `from709` into `e`.

	// Assume that `sourceCQL` has the underlying type of `from709`:
	from709 := sourceCQL.(gob.Decoder)

	// Declare `e` variable:
	var e interface{}

	// Call the method that transfers the taint
	// from the receiver `from709` to the argument `e`
	// (`e` is now tainted).
	from709.Decode(e)

	// Sink the tainted `e`:
	sink(e)
}

func TaintStepTest_EncodingGobDecoderDecodeValue(sourceCQL interface{}) {
	// The flow is from `from772` into `v`.

	// Assume that `sourceCQL` has the underlying type of `from772`:
	from772 := sourceCQL.(gob.Decoder)

	// Declare `v` variable:
	var v reflect.Value

	// Call the method that transfers the taint
	// from the receiver `from772` to the argument `v`
	// (`v` is now tainted).
	from772.DecodeValue(v)

	// Sink the tainted `v`:
	sink(v)
}

func TaintStepTest_EncodingGobEncoderEncode(sourceCQL interface{}) {
	// The flow is from `e` into `into599`.

	// Assume that `sourceCQL` has the underlying type of `e`:
	e := sourceCQL.(interface{})

	// Declare `into599` variable:
	var into599 gob.Encoder

	// Call the method that transfers the taint
	// from the parameter `e` to the receiver `into599`
	// (`into599` is now tainted).
	into599.Encode(e)

	// Sink the tainted `into599`:
	sink(into599)
}

func TaintStepTest_EncodingGobEncoderEncodeValue(sourceCQL interface{}) {
	// The flow is from `value` into `into956`.

	// Assume that `sourceCQL` has the underlying type of `value`:
	value := sourceCQL.(reflect.Value)

	// Declare `into956` variable:
	var into956 gob.Encoder

	// Call the method that transfers the taint
	// from the parameter `value` to the receiver `into956`
	// (`into956` is now tainted).
	into956.EncodeValue(value)

	// Sink the tainted `into956`:
	sink(into956)
}

func TaintStepTest_EncodingGobGobDecoderGobDecode(sourceCQL interface{}) {
	// The flow is from `from394` into `into783`.

	// Assume that `sourceCQL` has the underlying type of `from394`:
	from394 := sourceCQL.([]byte)

	// Declare `into783` variable:
	var into783 gob.GobDecoder

	// Call the method that transfers the taint
	// from the parameter `from394` to the receiver `into783`
	// (`into783` is now tainted).
	into783.GobDecode(from394)

	// Sink the tainted `into783`:
	sink(into783)
}

func TaintStepTest_EncodingGobGobEncoderGobEncode(sourceCQL interface{}) {
	// The flow is from `from340` into `into387`.

	// Assume that `sourceCQL` has the underlying type of `from340`:
	from340 := sourceCQL.(gob.GobEncoder)

	// Call the method that transfers the taint
	// from the receiver `from340` to the result `into387`
	// (`into387` is now tainted).
	into387, _ := from340.GobEncode()

	// Sink the tainted `into387`:
	sink(into387)
}

func RunAllTaints_EncodingGob(v interface{}) {
	{
		source := newSource()
		TaintStepTest_EncodingGobNewDecoder(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingGobNewEncoder(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingGobDecoderDecode(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingGobDecoderDecodeValue(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingGobEncoderEncode(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingGobEncoderEncodeValue(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingGobGobDecoderGobDecode(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingGobGobEncoderGobEncode(source)
	}
}
