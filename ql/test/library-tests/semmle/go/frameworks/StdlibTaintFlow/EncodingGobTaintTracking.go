// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"encoding/gob"
	"io"
	"reflect"
)

func TaintStepTest_EncodingGobNewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader656` into `intoDecoder414`.

	// Assume that `sourceCQL` has the underlying type of `fromReader656`:
	fromReader656 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader656` to result `intoDecoder414`
	// (`intoDecoder414` is now tainted).
	intoDecoder414 := gob.NewDecoder(fromReader656)

	// Return the tainted `intoDecoder414`:
	return intoDecoder414
}

func TaintStepTest_EncodingGobNewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromEncoder518` into `intoWriter650`.

	// Assume that `sourceCQL` has the underlying type of `fromEncoder518`:
	fromEncoder518 := sourceCQL.(*gob.Encoder)

	// Declare `intoWriter650` variable:
	var intoWriter650 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter650`:
	intermediateCQL := gob.NewEncoder(intoWriter650)

	// Extra step (`fromEncoder518` taints `intermediateCQL`, which taints `intoWriter650`:
	link(fromEncoder518, intermediateCQL)

	// Return the tainted `intoWriter650`:
	return intoWriter650
}

func TaintStepTest_EncodingGobDecoderDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder784` into `intoInterface957`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder784`:
	fromDecoder784 := sourceCQL.(gob.Decoder)

	// Declare `intoInterface957` variable:
	var intoInterface957 interface{}

	// Call the method that transfers the taint
	// from the receiver `fromDecoder784` to the argument `intoInterface957`
	// (`intoInterface957` is now tainted).
	fromDecoder784.Decode(intoInterface957)

	// Return the tainted `intoInterface957`:
	return intoInterface957
}

func TaintStepTest_EncodingGobDecoderDecodeValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder520` into `intoValue443`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder520`:
	fromDecoder520 := sourceCQL.(gob.Decoder)

	// Declare `intoValue443` variable:
	var intoValue443 reflect.Value

	// Call the method that transfers the taint
	// from the receiver `fromDecoder520` to the argument `intoValue443`
	// (`intoValue443` is now tainted).
	fromDecoder520.DecodeValue(intoValue443)

	// Return the tainted `intoValue443`:
	return intoValue443
}

func TaintStepTest_EncodingGobEncoderEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface127` into `intoEncoder483`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface127`:
	fromInterface127 := sourceCQL.(interface{})

	// Declare `intoEncoder483` variable:
	var intoEncoder483 gob.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromInterface127` to the receiver `intoEncoder483`
	// (`intoEncoder483` is now tainted).
	intoEncoder483.Encode(fromInterface127)

	// Return the tainted `intoEncoder483`:
	return intoEncoder483
}

func TaintStepTest_EncodingGobEncoderEncodeValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue989` into `intoEncoder982`.

	// Assume that `sourceCQL` has the underlying type of `fromValue989`:
	fromValue989 := sourceCQL.(reflect.Value)

	// Declare `intoEncoder982` variable:
	var intoEncoder982 gob.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromValue989` to the receiver `intoEncoder982`
	// (`intoEncoder982` is now tainted).
	intoEncoder982.EncodeValue(fromValue989)

	// Return the tainted `intoEncoder982`:
	return intoEncoder982
}

func TaintStepTest_EncodingGobGobDecoderGobDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte417` into `intoGobDecoder584`.

	// Assume that `sourceCQL` has the underlying type of `fromByte417`:
	fromByte417 := sourceCQL.([]byte)

	// Declare `intoGobDecoder584` variable:
	var intoGobDecoder584 gob.GobDecoder

	// Call the method that transfers the taint
	// from the parameter `fromByte417` to the receiver `intoGobDecoder584`
	// (`intoGobDecoder584` is now tainted).
	intoGobDecoder584.GobDecode(fromByte417)

	// Return the tainted `intoGobDecoder584`:
	return intoGobDecoder584
}

func TaintStepTest_EncodingGobGobEncoderGobEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromGobEncoder991` into `intoByte881`.

	// Assume that `sourceCQL` has the underlying type of `fromGobEncoder991`:
	fromGobEncoder991 := sourceCQL.(gob.GobEncoder)

	// Call the method that transfers the taint
	// from the receiver `fromGobEncoder991` to the result `intoByte881`
	// (`intoByte881` is now tainted).
	intoByte881, _ := fromGobEncoder991.GobEncode()

	// Return the tainted `intoByte881`:
	return intoByte881
}

func RunAllTaints_EncodingGob() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingGobNewDecoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingGobNewEncoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingGobDecoderDecode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingGobDecoderDecodeValue_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingGobEncoderEncode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingGobEncoderEncodeValue_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingGobGobDecoderGobDecode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingGobGobEncoderGobEncode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
