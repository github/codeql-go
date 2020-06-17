package main

import (
	"encoding/gob"
	"io"
	"reflect"
)

func TaintStepTest_EncodingGobNewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader165` into `intoDecoder735`.

	// Assume that `sourceCQL` has the underlying type of `fromReader165`:
	fromReader165 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader165` to result `intoDecoder735`
	// (`intoDecoder735` is now tainted).
	intoDecoder735 := gob.NewDecoder(fromReader165)

	// Return the tainted `intoDecoder735`:
	return intoDecoder735
}

func TaintStepTest_EncodingGobNewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromEncoder449` into `intoWriter736`.

	// Assume that `sourceCQL` has the underlying type of `fromEncoder449`:
	fromEncoder449 := sourceCQL.(*gob.Encoder)

	// Declare `intoWriter736` variable:
	var intoWriter736 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter736`:
	intermediateCQL := gob.NewEncoder(intoWriter736)

	// Extra step (`fromEncoder449` taints `intermediateCQL`, which taints `intoWriter736`:
	link(fromEncoder449, intermediateCQL)

	// Return the tainted `intoWriter736`:
	return intoWriter736
}

func TaintStepTest_EncodingGobDecoderDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder946` into `intoInterface397`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder946`:
	fromDecoder946 := sourceCQL.(gob.Decoder)

	// Declare `intoInterface397` variable:
	var intoInterface397 interface{}

	// Call the method that transfers the taint
	// from the receiver `fromDecoder946` to the argument `intoInterface397`
	// (`intoInterface397` is now tainted).
	fromDecoder946.Decode(intoInterface397)

	// Return the tainted `intoInterface397`:
	return intoInterface397
}

func TaintStepTest_EncodingGobDecoderDecodeValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder507` into `intoValue302`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder507`:
	fromDecoder507 := sourceCQL.(gob.Decoder)

	// Declare `intoValue302` variable:
	var intoValue302 reflect.Value

	// Call the method that transfers the taint
	// from the receiver `fromDecoder507` to the argument `intoValue302`
	// (`intoValue302` is now tainted).
	fromDecoder507.DecodeValue(intoValue302)

	// Return the tainted `intoValue302`:
	return intoValue302
}

func TaintStepTest_EncodingGobEncoderEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface878` into `intoEncoder864`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface878`:
	fromInterface878 := sourceCQL.(interface{})

	// Declare `intoEncoder864` variable:
	var intoEncoder864 gob.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromInterface878` to the receiver `intoEncoder864`
	// (`intoEncoder864` is now tainted).
	intoEncoder864.Encode(fromInterface878)

	// Return the tainted `intoEncoder864`:
	return intoEncoder864
}

func TaintStepTest_EncodingGobEncoderEncodeValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue235` into `intoEncoder213`.

	// Assume that `sourceCQL` has the underlying type of `fromValue235`:
	fromValue235 := sourceCQL.(reflect.Value)

	// Declare `intoEncoder213` variable:
	var intoEncoder213 gob.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromValue235` to the receiver `intoEncoder213`
	// (`intoEncoder213` is now tainted).
	intoEncoder213.EncodeValue(fromValue235)

	// Return the tainted `intoEncoder213`:
	return intoEncoder213
}

func TaintStepTest_EncodingGobGobDecoderGobDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte998` into `intoGobDecoder569`.

	// Assume that `sourceCQL` has the underlying type of `fromByte998`:
	fromByte998 := sourceCQL.([]byte)

	// Declare `intoGobDecoder569` variable:
	var intoGobDecoder569 gob.GobDecoder

	// Call the method that transfers the taint
	// from the parameter `fromByte998` to the receiver `intoGobDecoder569`
	// (`intoGobDecoder569` is now tainted).
	intoGobDecoder569.GobDecode(fromByte998)

	// Return the tainted `intoGobDecoder569`:
	return intoGobDecoder569
}

func TaintStepTest_EncodingGobGobEncoderGobEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromGobEncoder353` into `intoByte893`.

	// Assume that `sourceCQL` has the underlying type of `fromGobEncoder353`:
	fromGobEncoder353 := sourceCQL.(gob.GobEncoder)

	// Call the method that transfers the taint
	// from the receiver `fromGobEncoder353` to the result `intoByte893`
	// (`intoByte893` is now tainted).
	intoByte893, _ := fromGobEncoder353.GobEncode()

	// Return the tainted `intoByte893`:
	return intoByte893
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
