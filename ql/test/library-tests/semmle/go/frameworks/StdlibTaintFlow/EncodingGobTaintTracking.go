package main

import (
	"encoding/gob"
	"io"
	"reflect"
)

func TaintStepTest_EncodingGobNewDecoder_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader554` into `intoDecoder586`.

	// Assume that `sourceCQL` has the underlying type of `fromReader554`:
	fromReader554 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader554` to result `intoDecoder586`
	// (`intoDecoder586` is now tainted).
	intoDecoder586 := gob.NewDecoder(fromReader554)

	// Sink the tainted `intoDecoder586`:
	sink(intoDecoder586)
}

func TaintStepTest_EncodingGobNewEncoder_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromEncoder323` into `intoWriter214`.

	// Assume that `sourceCQL` has the underlying type of `fromEncoder323`:
	fromEncoder323 := sourceCQL.(*gob.Encoder)

	// Declare `intoWriter214` variable:
	var intoWriter214 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter214`:
	intermediateCQL := gob.NewEncoder(intoWriter214)

	// Extra step (`fromEncoder323` taints `intermediateCQL`, which taints `intoWriter214`:
	link(fromEncoder323, intermediateCQL)

	// Sink the tainted `intoWriter214`:
	sink(intoWriter214)
}

func TaintStepTest_EncodingGobDecoderDecode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromDecoder469` into `intoInterface932`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder469`:
	fromDecoder469 := sourceCQL.(gob.Decoder)

	// Declare `intoInterface932` variable:
	var intoInterface932 interface{}

	// Call the method that transfers the taint
	// from the receiver `fromDecoder469` to the argument `intoInterface932`
	// (`intoInterface932` is now tainted).
	fromDecoder469.Decode(intoInterface932)

	// Sink the tainted `intoInterface932`:
	sink(intoInterface932)
}

func TaintStepTest_EncodingGobDecoderDecodeValue_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromDecoder397` into `intoValue235`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder397`:
	fromDecoder397 := sourceCQL.(gob.Decoder)

	// Declare `intoValue235` variable:
	var intoValue235 reflect.Value

	// Call the method that transfers the taint
	// from the receiver `fromDecoder397` to the argument `intoValue235`
	// (`intoValue235` is now tainted).
	fromDecoder397.DecodeValue(intoValue235)

	// Sink the tainted `intoValue235`:
	sink(intoValue235)
}

func TaintStepTest_EncodingGobEncoderEncode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface667` into `intoEncoder736`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface667`:
	fromInterface667 := sourceCQL.(interface{})

	// Declare `intoEncoder736` variable:
	var intoEncoder736 gob.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromInterface667` to the receiver `intoEncoder736`
	// (`intoEncoder736` is now tainted).
	intoEncoder736.Encode(fromInterface667)

	// Sink the tainted `intoEncoder736`:
	sink(intoEncoder736)
}

func TaintStepTest_EncodingGobEncoderEncodeValue_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue574` into `intoEncoder435`.

	// Assume that `sourceCQL` has the underlying type of `fromValue574`:
	fromValue574 := sourceCQL.(reflect.Value)

	// Declare `intoEncoder435` variable:
	var intoEncoder435 gob.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromValue574` to the receiver `intoEncoder435`
	// (`intoEncoder435` is now tainted).
	intoEncoder435.EncodeValue(fromValue574)

	// Sink the tainted `intoEncoder435`:
	sink(intoEncoder435)
}

func TaintStepTest_EncodingGobGobDecoderGobDecode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte171` into `intoGobDecoder202`.

	// Assume that `sourceCQL` has the underlying type of `fromByte171`:
	fromByte171 := sourceCQL.([]byte)

	// Declare `intoGobDecoder202` variable:
	var intoGobDecoder202 gob.GobDecoder

	// Call the method that transfers the taint
	// from the parameter `fromByte171` to the receiver `intoGobDecoder202`
	// (`intoGobDecoder202` is now tainted).
	intoGobDecoder202.GobDecode(fromByte171)

	// Sink the tainted `intoGobDecoder202`:
	sink(intoGobDecoder202)
}

func TaintStepTest_EncodingGobGobEncoderGobEncode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromGobEncoder712` into `intoByte174`.

	// Assume that `sourceCQL` has the underlying type of `fromGobEncoder712`:
	fromGobEncoder712 := sourceCQL.(gob.GobEncoder)

	// Call the method that transfers the taint
	// from the receiver `fromGobEncoder712` to the result `intoByte174`
	// (`intoByte174` is now tainted).
	intoByte174, _ := fromGobEncoder712.GobEncode()

	// Sink the tainted `intoByte174`:
	sink(intoByte174)
}

func RunAllTaints_EncodingGob() {
	{
		source := newSource()
		TaintStepTest_EncodingGobNewDecoder_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingGobNewEncoder_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingGobDecoderDecode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingGobDecoderDecodeValue_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingGobEncoderEncode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingGobEncoderEncodeValue_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingGobGobDecoderGobDecode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingGobGobEncoderGobEncode_B0I0O0(source)
	}
}
