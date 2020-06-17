// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"encoding/gob"
	"io"
	"reflect"
)

func TaintStepTest_EncodingGobNewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader212` into `intoDecoder950`.

	// Assume that `sourceCQL` has the underlying type of `fromReader212`:
	fromReader212 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader212` to result `intoDecoder950`
	// (`intoDecoder950` is now tainted).
	intoDecoder950 := gob.NewDecoder(fromReader212)

	// Return the tainted `intoDecoder950`:
	return intoDecoder950
}

func TaintStepTest_EncodingGobNewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromEncoder502` into `intoWriter580`.

	// Assume that `sourceCQL` has the underlying type of `fromEncoder502`:
	fromEncoder502 := sourceCQL.(*gob.Encoder)

	// Declare `intoWriter580` variable:
	var intoWriter580 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter580`:
	intermediateCQL := gob.NewEncoder(intoWriter580)

	// Extra step (`fromEncoder502` taints `intermediateCQL`, which taints `intoWriter580`:
	link(fromEncoder502, intermediateCQL)

	// Return the tainted `intoWriter580`:
	return intoWriter580
}

func TaintStepTest_EncodingGobDecoderDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder318` into `intoInterface920`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder318`:
	fromDecoder318 := sourceCQL.(gob.Decoder)

	// Declare `intoInterface920` variable:
	var intoInterface920 interface{}

	// Call the method that transfers the taint
	// from the receiver `fromDecoder318` to the argument `intoInterface920`
	// (`intoInterface920` is now tainted).
	fromDecoder318.Decode(intoInterface920)

	// Return the tainted `intoInterface920`:
	return intoInterface920
}

func TaintStepTest_EncodingGobDecoderDecodeValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder214` into `intoValue857`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder214`:
	fromDecoder214 := sourceCQL.(gob.Decoder)

	// Declare `intoValue857` variable:
	var intoValue857 reflect.Value

	// Call the method that transfers the taint
	// from the receiver `fromDecoder214` to the argument `intoValue857`
	// (`intoValue857` is now tainted).
	fromDecoder214.DecodeValue(intoValue857)

	// Return the tainted `intoValue857`:
	return intoValue857
}

func TaintStepTest_EncodingGobEncoderEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface163` into `intoEncoder711`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface163`:
	fromInterface163 := sourceCQL.(interface{})

	// Declare `intoEncoder711` variable:
	var intoEncoder711 gob.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromInterface163` to the receiver `intoEncoder711`
	// (`intoEncoder711` is now tainted).
	intoEncoder711.Encode(fromInterface163)

	// Return the tainted `intoEncoder711`:
	return intoEncoder711
}

func TaintStepTest_EncodingGobEncoderEncodeValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue896` into `intoEncoder601`.

	// Assume that `sourceCQL` has the underlying type of `fromValue896`:
	fromValue896 := sourceCQL.(reflect.Value)

	// Declare `intoEncoder601` variable:
	var intoEncoder601 gob.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromValue896` to the receiver `intoEncoder601`
	// (`intoEncoder601` is now tainted).
	intoEncoder601.EncodeValue(fromValue896)

	// Return the tainted `intoEncoder601`:
	return intoEncoder601
}

func TaintStepTest_EncodingGobGobDecoderGobDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte229` into `intoGobDecoder376`.

	// Assume that `sourceCQL` has the underlying type of `fromByte229`:
	fromByte229 := sourceCQL.([]byte)

	// Declare `intoGobDecoder376` variable:
	var intoGobDecoder376 gob.GobDecoder

	// Call the method that transfers the taint
	// from the parameter `fromByte229` to the receiver `intoGobDecoder376`
	// (`intoGobDecoder376` is now tainted).
	intoGobDecoder376.GobDecode(fromByte229)

	// Return the tainted `intoGobDecoder376`:
	return intoGobDecoder376
}

func TaintStepTest_EncodingGobGobEncoderGobEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromGobEncoder321` into `intoByte914`.

	// Assume that `sourceCQL` has the underlying type of `fromGobEncoder321`:
	fromGobEncoder321 := sourceCQL.(gob.GobEncoder)

	// Call the method that transfers the taint
	// from the receiver `fromGobEncoder321` to the result `intoByte914`
	// (`intoByte914` is now tainted).
	intoByte914, _ := fromGobEncoder321.GobEncode()

	// Return the tainted `intoByte914`:
	return intoByte914
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
