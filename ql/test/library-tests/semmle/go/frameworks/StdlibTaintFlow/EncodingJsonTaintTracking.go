package main

import (
	"bytes"
	"encoding/json"
	"io"
)

func TaintStepTest_EncodingJsonCompact_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte672` into `intoBuffer898`.

	// Assume that `sourceCQL` has the underlying type of `fromByte672`:
	fromByte672 := sourceCQL.([]byte)

	// Declare `intoBuffer898` variable:
	var intoBuffer898 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromByte672` to parameter `intoBuffer898`;
	// `intoBuffer898` is now tainted.
	json.Compact(intoBuffer898, fromByte672)

	// Return the tainted `intoBuffer898`:
	return intoBuffer898
}

func TaintStepTest_EncodingJsonHTMLEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte578` into `intoBuffer908`.

	// Assume that `sourceCQL` has the underlying type of `fromByte578`:
	fromByte578 := sourceCQL.([]byte)

	// Declare `intoBuffer908` variable:
	var intoBuffer908 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromByte578` to parameter `intoBuffer908`;
	// `intoBuffer908` is now tainted.
	json.HTMLEscape(intoBuffer908, fromByte578)

	// Return the tainted `intoBuffer908`:
	return intoBuffer908
}

func TaintStepTest_EncodingJsonIndent_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte636` into `intoBuffer575`.

	// Assume that `sourceCQL` has the underlying type of `fromByte636`:
	fromByte636 := sourceCQL.([]byte)

	// Declare `intoBuffer575` variable:
	var intoBuffer575 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromByte636` to parameter `intoBuffer575`;
	// `intoBuffer575` is now tainted.
	json.Indent(intoBuffer575, fromByte636, "", "")

	// Return the tainted `intoBuffer575`:
	return intoBuffer575
}

func TaintStepTest_EncodingJsonIndent_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString985` into `intoBuffer156`.

	// Assume that `sourceCQL` has the underlying type of `fromString985`:
	fromString985 := sourceCQL.(string)

	// Declare `intoBuffer156` variable:
	var intoBuffer156 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromString985` to parameter `intoBuffer156`;
	// `intoBuffer156` is now tainted.
	json.Indent(intoBuffer156, nil, fromString985, "")

	// Return the tainted `intoBuffer156`:
	return intoBuffer156
}

func TaintStepTest_EncodingJsonIndent_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString366` into `intoBuffer664`.

	// Assume that `sourceCQL` has the underlying type of `fromString366`:
	fromString366 := sourceCQL.(string)

	// Declare `intoBuffer664` variable:
	var intoBuffer664 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromString366` to parameter `intoBuffer664`;
	// `intoBuffer664` is now tainted.
	json.Indent(intoBuffer664, nil, "", fromString366)

	// Return the tainted `intoBuffer664`:
	return intoBuffer664
}

func TaintStepTest_EncodingJsonMarshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface231` into `intoByte844`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface231`:
	fromInterface231 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface231` to result `intoByte844`
	// (`intoByte844` is now tainted).
	intoByte844, _ := json.Marshal(fromInterface231)

	// Return the tainted `intoByte844`:
	return intoByte844
}

func TaintStepTest_EncodingJsonMarshalIndent_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface402` into `intoByte602`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface402`:
	fromInterface402 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface402` to result `intoByte602`
	// (`intoByte602` is now tainted).
	intoByte602, _ := json.MarshalIndent(fromInterface402, "", "")

	// Return the tainted `intoByte602`:
	return intoByte602
}

func TaintStepTest_EncodingJsonMarshalIndent_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString646` into `intoByte338`.

	// Assume that `sourceCQL` has the underlying type of `fromString646`:
	fromString646 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString646` to result `intoByte338`
	// (`intoByte338` is now tainted).
	intoByte338, _ := json.MarshalIndent(nil, fromString646, "")

	// Return the tainted `intoByte338`:
	return intoByte338
}

func TaintStepTest_EncodingJsonMarshalIndent_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString726` into `intoByte260`.

	// Assume that `sourceCQL` has the underlying type of `fromString726`:
	fromString726 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString726` to result `intoByte260`
	// (`intoByte260` is now tainted).
	intoByte260, _ := json.MarshalIndent(nil, "", fromString726)

	// Return the tainted `intoByte260`:
	return intoByte260
}

func TaintStepTest_EncodingJsonNewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader476` into `intoDecoder695`.

	// Assume that `sourceCQL` has the underlying type of `fromReader476`:
	fromReader476 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader476` to result `intoDecoder695`
	// (`intoDecoder695` is now tainted).
	intoDecoder695 := json.NewDecoder(fromReader476)

	// Return the tainted `intoDecoder695`:
	return intoDecoder695
}

func TaintStepTest_EncodingJsonNewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromEncoder549` into `intoWriter940`.

	// Assume that `sourceCQL` has the underlying type of `fromEncoder549`:
	fromEncoder549 := sourceCQL.(*json.Encoder)

	// Declare `intoWriter940` variable:
	var intoWriter940 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter940`:
	intermediateCQL := json.NewEncoder(intoWriter940)

	// Extra step (`fromEncoder549` taints `intermediateCQL`, which taints `intoWriter940`:
	link(fromEncoder549, intermediateCQL)

	// Return the tainted `intoWriter940`:
	return intoWriter940
}

func TaintStepTest_EncodingJsonUnmarshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte644` into `intoInterface787`.

	// Assume that `sourceCQL` has the underlying type of `fromByte644`:
	fromByte644 := sourceCQL.([]byte)

	// Declare `intoInterface787` variable:
	var intoInterface787 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromByte644` to parameter `intoInterface787`;
	// `intoInterface787` is now tainted.
	json.Unmarshal(fromByte644, intoInterface787)

	// Return the tainted `intoInterface787`:
	return intoInterface787
}

func TaintStepTest_EncodingJsonDecoderBuffered_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder907` into `intoReader260`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder907`:
	fromDecoder907 := sourceCQL.(json.Decoder)

	// Call the method that transfers the taint
	// from the receiver `fromDecoder907` to the result `intoReader260`
	// (`intoReader260` is now tainted).
	intoReader260 := fromDecoder907.Buffered()

	// Return the tainted `intoReader260`:
	return intoReader260
}

func TaintStepTest_EncodingJsonDecoderDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder540` into `intoInterface365`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder540`:
	fromDecoder540 := sourceCQL.(json.Decoder)

	// Declare `intoInterface365` variable:
	var intoInterface365 interface{}

	// Call the method that transfers the taint
	// from the receiver `fromDecoder540` to the argument `intoInterface365`
	// (`intoInterface365` is now tainted).
	fromDecoder540.Decode(intoInterface365)

	// Return the tainted `intoInterface365`:
	return intoInterface365
}

func TaintStepTest_EncodingJsonDecoderToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder369` into `intoToken392`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder369`:
	fromDecoder369 := sourceCQL.(json.Decoder)

	// Call the method that transfers the taint
	// from the receiver `fromDecoder369` to the result `intoToken392`
	// (`intoToken392` is now tainted).
	intoToken392, _ := fromDecoder369.Token()

	// Return the tainted `intoToken392`:
	return intoToken392
}

func TaintStepTest_EncodingJsonEncoderEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface279` into `intoEncoder927`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface279`:
	fromInterface279 := sourceCQL.(interface{})

	// Declare `intoEncoder927` variable:
	var intoEncoder927 json.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromInterface279` to the receiver `intoEncoder927`
	// (`intoEncoder927` is now tainted).
	intoEncoder927.Encode(fromInterface279)

	// Return the tainted `intoEncoder927`:
	return intoEncoder927
}

func TaintStepTest_EncodingJsonEncoderSetIndent_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString911` into `intoEncoder816`.

	// Assume that `sourceCQL` has the underlying type of `fromString911`:
	fromString911 := sourceCQL.(string)

	// Declare `intoEncoder816` variable:
	var intoEncoder816 json.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromString911` to the receiver `intoEncoder816`
	// (`intoEncoder816` is now tainted).
	intoEncoder816.SetIndent(fromString911, "")

	// Return the tainted `intoEncoder816`:
	return intoEncoder816
}

func TaintStepTest_EncodingJsonEncoderSetIndent_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString305` into `intoEncoder298`.

	// Assume that `sourceCQL` has the underlying type of `fromString305`:
	fromString305 := sourceCQL.(string)

	// Declare `intoEncoder298` variable:
	var intoEncoder298 json.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromString305` to the receiver `intoEncoder298`
	// (`intoEncoder298` is now tainted).
	intoEncoder298.SetIndent("", fromString305)

	// Return the tainted `intoEncoder298`:
	return intoEncoder298
}

func TaintStepTest_EncodingJsonRawMessageMarshalJSON_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawMessage132` into `intoByte503`.

	// Assume that `sourceCQL` has the underlying type of `fromRawMessage132`:
	fromRawMessage132 := sourceCQL.(json.RawMessage)

	// Call the method that transfers the taint
	// from the receiver `fromRawMessage132` to the result `intoByte503`
	// (`intoByte503` is now tainted).
	intoByte503, _ := fromRawMessage132.MarshalJSON()

	// Return the tainted `intoByte503`:
	return intoByte503
}

func TaintStepTest_EncodingJsonRawMessageUnmarshalJSON_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte489` into `intoRawMessage933`.

	// Assume that `sourceCQL` has the underlying type of `fromByte489`:
	fromByte489 := sourceCQL.([]byte)

	// Declare `intoRawMessage933` variable:
	var intoRawMessage933 json.RawMessage

	// Call the method that transfers the taint
	// from the parameter `fromByte489` to the receiver `intoRawMessage933`
	// (`intoRawMessage933` is now tainted).
	intoRawMessage933.UnmarshalJSON(fromByte489)

	// Return the tainted `intoRawMessage933`:
	return intoRawMessage933
}

func TaintStepTest_EncodingJsonMarshalerMarshalJSON_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMarshaler805` into `intoByte752`.

	// Assume that `sourceCQL` has the underlying type of `fromMarshaler805`:
	fromMarshaler805 := sourceCQL.(json.Marshaler)

	// Call the method that transfers the taint
	// from the receiver `fromMarshaler805` to the result `intoByte752`
	// (`intoByte752` is now tainted).
	intoByte752, _ := fromMarshaler805.MarshalJSON()

	// Return the tainted `intoByte752`:
	return intoByte752
}

func TaintStepTest_EncodingJsonUnmarshalerUnmarshalJSON_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte694` into `intoUnmarshaler572`.

	// Assume that `sourceCQL` has the underlying type of `fromByte694`:
	fromByte694 := sourceCQL.([]byte)

	// Declare `intoUnmarshaler572` variable:
	var intoUnmarshaler572 json.Unmarshaler

	// Call the method that transfers the taint
	// from the parameter `fromByte694` to the receiver `intoUnmarshaler572`
	// (`intoUnmarshaler572` is now tainted).
	intoUnmarshaler572.UnmarshalJSON(fromByte694)

	// Return the tainted `intoUnmarshaler572`:
	return intoUnmarshaler572
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
