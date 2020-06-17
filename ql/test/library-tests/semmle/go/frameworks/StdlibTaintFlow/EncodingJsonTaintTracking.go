// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"bytes"
	"encoding/json"
	"io"
)

func TaintStepTest_EncodingJsonCompact_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte656` into `intoBuffer414`.

	// Assume that `sourceCQL` has the underlying type of `fromByte656`:
	fromByte656 := sourceCQL.([]byte)

	// Declare `intoBuffer414` variable:
	var intoBuffer414 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromByte656` to parameter `intoBuffer414`;
	// `intoBuffer414` is now tainted.
	json.Compact(intoBuffer414, fromByte656)

	// Return the tainted `intoBuffer414`:
	return intoBuffer414
}

func TaintStepTest_EncodingJsonHTMLEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte518` into `intoBuffer650`.

	// Assume that `sourceCQL` has the underlying type of `fromByte518`:
	fromByte518 := sourceCQL.([]byte)

	// Declare `intoBuffer650` variable:
	var intoBuffer650 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromByte518` to parameter `intoBuffer650`;
	// `intoBuffer650` is now tainted.
	json.HTMLEscape(intoBuffer650, fromByte518)

	// Return the tainted `intoBuffer650`:
	return intoBuffer650
}

func TaintStepTest_EncodingJsonIndent_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte784` into `intoBuffer957`.

	// Assume that `sourceCQL` has the underlying type of `fromByte784`:
	fromByte784 := sourceCQL.([]byte)

	// Declare `intoBuffer957` variable:
	var intoBuffer957 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromByte784` to parameter `intoBuffer957`;
	// `intoBuffer957` is now tainted.
	json.Indent(intoBuffer957, fromByte784, "", "")

	// Return the tainted `intoBuffer957`:
	return intoBuffer957
}

func TaintStepTest_EncodingJsonIndent_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString520` into `intoBuffer443`.

	// Assume that `sourceCQL` has the underlying type of `fromString520`:
	fromString520 := sourceCQL.(string)

	// Declare `intoBuffer443` variable:
	var intoBuffer443 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromString520` to parameter `intoBuffer443`;
	// `intoBuffer443` is now tainted.
	json.Indent(intoBuffer443, nil, fromString520, "")

	// Return the tainted `intoBuffer443`:
	return intoBuffer443
}

func TaintStepTest_EncodingJsonIndent_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString127` into `intoBuffer483`.

	// Assume that `sourceCQL` has the underlying type of `fromString127`:
	fromString127 := sourceCQL.(string)

	// Declare `intoBuffer483` variable:
	var intoBuffer483 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromString127` to parameter `intoBuffer483`;
	// `intoBuffer483` is now tainted.
	json.Indent(intoBuffer483, nil, "", fromString127)

	// Return the tainted `intoBuffer483`:
	return intoBuffer483
}

func TaintStepTest_EncodingJsonMarshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface989` into `intoByte982`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface989`:
	fromInterface989 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface989` to result `intoByte982`
	// (`intoByte982` is now tainted).
	intoByte982, _ := json.Marshal(fromInterface989)

	// Return the tainted `intoByte982`:
	return intoByte982
}

func TaintStepTest_EncodingJsonMarshalIndent_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface417` into `intoByte584`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface417`:
	fromInterface417 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface417` to result `intoByte584`
	// (`intoByte584` is now tainted).
	intoByte584, _ := json.MarshalIndent(fromInterface417, "", "")

	// Return the tainted `intoByte584`:
	return intoByte584
}

func TaintStepTest_EncodingJsonMarshalIndent_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString991` into `intoByte881`.

	// Assume that `sourceCQL` has the underlying type of `fromString991`:
	fromString991 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString991` to result `intoByte881`
	// (`intoByte881` is now tainted).
	intoByte881, _ := json.MarshalIndent(nil, fromString991, "")

	// Return the tainted `intoByte881`:
	return intoByte881
}

func TaintStepTest_EncodingJsonMarshalIndent_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString186` into `intoByte284`.

	// Assume that `sourceCQL` has the underlying type of `fromString186`:
	fromString186 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString186` to result `intoByte284`
	// (`intoByte284` is now tainted).
	intoByte284, _ := json.MarshalIndent(nil, "", fromString186)

	// Return the tainted `intoByte284`:
	return intoByte284
}

func TaintStepTest_EncodingJsonNewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader908` into `intoDecoder137`.

	// Assume that `sourceCQL` has the underlying type of `fromReader908`:
	fromReader908 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader908` to result `intoDecoder137`
	// (`intoDecoder137` is now tainted).
	intoDecoder137 := json.NewDecoder(fromReader908)

	// Return the tainted `intoDecoder137`:
	return intoDecoder137
}

func TaintStepTest_EncodingJsonNewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromEncoder494` into `intoWriter873`.

	// Assume that `sourceCQL` has the underlying type of `fromEncoder494`:
	fromEncoder494 := sourceCQL.(*json.Encoder)

	// Declare `intoWriter873` variable:
	var intoWriter873 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter873`:
	intermediateCQL := json.NewEncoder(intoWriter873)

	// Extra step (`fromEncoder494` taints `intermediateCQL`, which taints `intoWriter873`:
	link(fromEncoder494, intermediateCQL)

	// Return the tainted `intoWriter873`:
	return intoWriter873
}

func TaintStepTest_EncodingJsonUnmarshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte599` into `intoInterface409`.

	// Assume that `sourceCQL` has the underlying type of `fromByte599`:
	fromByte599 := sourceCQL.([]byte)

	// Declare `intoInterface409` variable:
	var intoInterface409 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromByte599` to parameter `intoInterface409`;
	// `intoInterface409` is now tainted.
	json.Unmarshal(fromByte599, intoInterface409)

	// Return the tainted `intoInterface409`:
	return intoInterface409
}

func TaintStepTest_EncodingJsonDecoderBuffered_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder246` into `intoReader898`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder246`:
	fromDecoder246 := sourceCQL.(json.Decoder)

	// Call the method that transfers the taint
	// from the receiver `fromDecoder246` to the result `intoReader898`
	// (`intoReader898` is now tainted).
	intoReader898 := fromDecoder246.Buffered()

	// Return the tainted `intoReader898`:
	return intoReader898
}

func TaintStepTest_EncodingJsonDecoderDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder598` into `intoInterface631`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder598`:
	fromDecoder598 := sourceCQL.(json.Decoder)

	// Declare `intoInterface631` variable:
	var intoInterface631 interface{}

	// Call the method that transfers the taint
	// from the receiver `fromDecoder598` to the argument `intoInterface631`
	// (`intoInterface631` is now tainted).
	fromDecoder598.Decode(intoInterface631)

	// Return the tainted `intoInterface631`:
	return intoInterface631
}

func TaintStepTest_EncodingJsonDecoderToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder165` into `intoToken150`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder165`:
	fromDecoder165 := sourceCQL.(json.Decoder)

	// Call the method that transfers the taint
	// from the receiver `fromDecoder165` to the result `intoToken150`
	// (`intoToken150` is now tainted).
	intoToken150, _ := fromDecoder165.Token()

	// Return the tainted `intoToken150`:
	return intoToken150
}

func TaintStepTest_EncodingJsonEncoderEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface340` into `intoEncoder471`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface340`:
	fromInterface340 := sourceCQL.(interface{})

	// Declare `intoEncoder471` variable:
	var intoEncoder471 json.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromInterface340` to the receiver `intoEncoder471`
	// (`intoEncoder471` is now tainted).
	intoEncoder471.Encode(fromInterface340)

	// Return the tainted `intoEncoder471`:
	return intoEncoder471
}

func TaintStepTest_EncodingJsonEncoderSetIndent_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString290` into `intoEncoder758`.

	// Assume that `sourceCQL` has the underlying type of `fromString290`:
	fromString290 := sourceCQL.(string)

	// Declare `intoEncoder758` variable:
	var intoEncoder758 json.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromString290` to the receiver `intoEncoder758`
	// (`intoEncoder758` is now tainted).
	intoEncoder758.SetIndent(fromString290, "")

	// Return the tainted `intoEncoder758`:
	return intoEncoder758
}

func TaintStepTest_EncodingJsonEncoderSetIndent_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString396` into `intoEncoder707`.

	// Assume that `sourceCQL` has the underlying type of `fromString396`:
	fromString396 := sourceCQL.(string)

	// Declare `intoEncoder707` variable:
	var intoEncoder707 json.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromString396` to the receiver `intoEncoder707`
	// (`intoEncoder707` is now tainted).
	intoEncoder707.SetIndent("", fromString396)

	// Return the tainted `intoEncoder707`:
	return intoEncoder707
}

func TaintStepTest_EncodingJsonRawMessageMarshalJSON_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawMessage912` into `intoByte718`.

	// Assume that `sourceCQL` has the underlying type of `fromRawMessage912`:
	fromRawMessage912 := sourceCQL.(json.RawMessage)

	// Call the method that transfers the taint
	// from the receiver `fromRawMessage912` to the result `intoByte718`
	// (`intoByte718` is now tainted).
	intoByte718, _ := fromRawMessage912.MarshalJSON()

	// Return the tainted `intoByte718`:
	return intoByte718
}

func TaintStepTest_EncodingJsonRawMessageUnmarshalJSON_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte972` into `intoRawMessage633`.

	// Assume that `sourceCQL` has the underlying type of `fromByte972`:
	fromByte972 := sourceCQL.([]byte)

	// Declare `intoRawMessage633` variable:
	var intoRawMessage633 json.RawMessage

	// Call the method that transfers the taint
	// from the parameter `fromByte972` to the receiver `intoRawMessage633`
	// (`intoRawMessage633` is now tainted).
	intoRawMessage633.UnmarshalJSON(fromByte972)

	// Return the tainted `intoRawMessage633`:
	return intoRawMessage633
}

func TaintStepTest_EncodingJsonMarshalerMarshalJSON_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMarshaler316` into `intoByte145`.

	// Assume that `sourceCQL` has the underlying type of `fromMarshaler316`:
	fromMarshaler316 := sourceCQL.(json.Marshaler)

	// Call the method that transfers the taint
	// from the receiver `fromMarshaler316` to the result `intoByte145`
	// (`intoByte145` is now tainted).
	intoByte145, _ := fromMarshaler316.MarshalJSON()

	// Return the tainted `intoByte145`:
	return intoByte145
}

func TaintStepTest_EncodingJsonUnmarshalerUnmarshalJSON_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte817` into `intoUnmarshaler474`.

	// Assume that `sourceCQL` has the underlying type of `fromByte817`:
	fromByte817 := sourceCQL.([]byte)

	// Declare `intoUnmarshaler474` variable:
	var intoUnmarshaler474 json.Unmarshaler

	// Call the method that transfers the taint
	// from the parameter `fromByte817` to the receiver `intoUnmarshaler474`
	// (`intoUnmarshaler474` is now tainted).
	intoUnmarshaler474.UnmarshalJSON(fromByte817)

	// Return the tainted `intoUnmarshaler474`:
	return intoUnmarshaler474
}

func RunAllTaints_EncodingJson() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonCompact_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonHTMLEscape_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonIndent_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonIndent_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonIndent_B0I2O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonMarshal_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonMarshalIndent_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonMarshalIndent_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonMarshalIndent_B0I2O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonNewDecoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonNewEncoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonUnmarshal_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonDecoderBuffered_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonDecoderDecode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonDecoderToken_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonEncoderEncode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonEncoderSetIndent_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonEncoderSetIndent_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonRawMessageMarshalJSON_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonRawMessageUnmarshalJSON_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonMarshalerMarshalJSON_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingJsonUnmarshalerUnmarshalJSON_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
