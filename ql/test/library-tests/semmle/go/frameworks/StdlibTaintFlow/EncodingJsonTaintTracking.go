// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"bytes"
	"encoding/json"
	"io"
)

func TaintStepTest_EncodingJsonCompact_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte936` into `intoBuffer157`.

	// Assume that `sourceCQL` has the underlying type of `fromByte936`:
	fromByte936 := sourceCQL.([]byte)

	// Declare `intoBuffer157` variable:
	var intoBuffer157 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromByte936` to parameter `intoBuffer157`;
	// `intoBuffer157` is now tainted.
	json.Compact(intoBuffer157, fromByte936)

	// Return the tainted `intoBuffer157`:
	return intoBuffer157
}

func TaintStepTest_EncodingJsonHTMLEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte578` into `intoBuffer427`.

	// Assume that `sourceCQL` has the underlying type of `fromByte578`:
	fromByte578 := sourceCQL.([]byte)

	// Declare `intoBuffer427` variable:
	var intoBuffer427 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromByte578` to parameter `intoBuffer427`;
	// `intoBuffer427` is now tainted.
	json.HTMLEscape(intoBuffer427, fromByte578)

	// Return the tainted `intoBuffer427`:
	return intoBuffer427
}

func TaintStepTest_EncodingJsonIndent_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte898` into `intoBuffer486`.

	// Assume that `sourceCQL` has the underlying type of `fromByte898`:
	fromByte898 := sourceCQL.([]byte)

	// Declare `intoBuffer486` variable:
	var intoBuffer486 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromByte898` to parameter `intoBuffer486`;
	// `intoBuffer486` is now tainted.
	json.Indent(intoBuffer486, fromByte898, "", "")

	// Return the tainted `intoBuffer486`:
	return intoBuffer486
}

func TaintStepTest_EncodingJsonIndent_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString244` into `intoBuffer784`.

	// Assume that `sourceCQL` has the underlying type of `fromString244`:
	fromString244 := sourceCQL.(string)

	// Declare `intoBuffer784` variable:
	var intoBuffer784 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromString244` to parameter `intoBuffer784`;
	// `intoBuffer784` is now tainted.
	json.Indent(intoBuffer784, nil, fromString244, "")

	// Return the tainted `intoBuffer784`:
	return intoBuffer784
}

func TaintStepTest_EncodingJsonIndent_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString174` into `intoBuffer157`.

	// Assume that `sourceCQL` has the underlying type of `fromString174`:
	fromString174 := sourceCQL.(string)

	// Declare `intoBuffer157` variable:
	var intoBuffer157 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromString174` to parameter `intoBuffer157`;
	// `intoBuffer157` is now tainted.
	json.Indent(intoBuffer157, nil, "", fromString174)

	// Return the tainted `intoBuffer157`:
	return intoBuffer157
}

func TaintStepTest_EncodingJsonMarshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface197` into `intoByte769`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface197`:
	fromInterface197 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface197` to result `intoByte769`
	// (`intoByte769` is now tainted).
	intoByte769, _ := json.Marshal(fromInterface197)

	// Return the tainted `intoByte769`:
	return intoByte769
}

func TaintStepTest_EncodingJsonMarshalIndent_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface942` into `intoByte853`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface942`:
	fromInterface942 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface942` to result `intoByte853`
	// (`intoByte853` is now tainted).
	intoByte853, _ := json.MarshalIndent(fromInterface942, "", "")

	// Return the tainted `intoByte853`:
	return intoByte853
}

func TaintStepTest_EncodingJsonMarshalIndent_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString838` into `intoByte607`.

	// Assume that `sourceCQL` has the underlying type of `fromString838`:
	fromString838 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString838` to result `intoByte607`
	// (`intoByte607` is now tainted).
	intoByte607, _ := json.MarshalIndent(nil, fromString838, "")

	// Return the tainted `intoByte607`:
	return intoByte607
}

func TaintStepTest_EncodingJsonMarshalIndent_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString389` into `intoByte326`.

	// Assume that `sourceCQL` has the underlying type of `fromString389`:
	fromString389 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString389` to result `intoByte326`
	// (`intoByte326` is now tainted).
	intoByte326, _ := json.MarshalIndent(nil, "", fromString389)

	// Return the tainted `intoByte326`:
	return intoByte326
}

func TaintStepTest_EncodingJsonNewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader539` into `intoDecoder964`.

	// Assume that `sourceCQL` has the underlying type of `fromReader539`:
	fromReader539 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader539` to result `intoDecoder964`
	// (`intoDecoder964` is now tainted).
	intoDecoder964 := json.NewDecoder(fromReader539)

	// Return the tainted `intoDecoder964`:
	return intoDecoder964
}

func TaintStepTest_EncodingJsonNewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromEncoder644` into `intoWriter687`.

	// Assume that `sourceCQL` has the underlying type of `fromEncoder644`:
	fromEncoder644 := sourceCQL.(*json.Encoder)

	// Declare `intoWriter687` variable:
	var intoWriter687 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter687`:
	intermediateCQL := json.NewEncoder(intoWriter687)

	// Extra step (`fromEncoder644` taints `intermediateCQL`, which taints `intoWriter687`:
	link(fromEncoder644, intermediateCQL)

	// Return the tainted `intoWriter687`:
	return intoWriter687
}

func TaintStepTest_EncodingJsonUnmarshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte316` into `intoInterface781`.

	// Assume that `sourceCQL` has the underlying type of `fromByte316`:
	fromByte316 := sourceCQL.([]byte)

	// Declare `intoInterface781` variable:
	var intoInterface781 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromByte316` to parameter `intoInterface781`;
	// `intoInterface781` is now tainted.
	json.Unmarshal(fromByte316, intoInterface781)

	// Return the tainted `intoInterface781`:
	return intoInterface781
}

func TaintStepTest_EncodingJsonDecoderBuffered_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder444` into `intoReader131`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder444`:
	fromDecoder444 := sourceCQL.(json.Decoder)

	// Call the method that transfers the taint
	// from the receiver `fromDecoder444` to the result `intoReader131`
	// (`intoReader131` is now tainted).
	intoReader131 := fromDecoder444.Buffered()

	// Return the tainted `intoReader131`:
	return intoReader131
}

func TaintStepTest_EncodingJsonDecoderDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder891` into `intoInterface228`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder891`:
	fromDecoder891 := sourceCQL.(json.Decoder)

	// Declare `intoInterface228` variable:
	var intoInterface228 interface{}

	// Call the method that transfers the taint
	// from the receiver `fromDecoder891` to the argument `intoInterface228`
	// (`intoInterface228` is now tainted).
	fromDecoder891.Decode(intoInterface228)

	// Return the tainted `intoInterface228`:
	return intoInterface228
}

func TaintStepTest_EncodingJsonDecoderToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder645` into `intoToken297`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder645`:
	fromDecoder645 := sourceCQL.(json.Decoder)

	// Call the method that transfers the taint
	// from the receiver `fromDecoder645` to the result `intoToken297`
	// (`intoToken297` is now tainted).
	intoToken297, _ := fromDecoder645.Token()

	// Return the tainted `intoToken297`:
	return intoToken297
}

func TaintStepTest_EncodingJsonEncoderEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface482` into `intoEncoder280`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface482`:
	fromInterface482 := sourceCQL.(interface{})

	// Declare `intoEncoder280` variable:
	var intoEncoder280 json.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromInterface482` to the receiver `intoEncoder280`
	// (`intoEncoder280` is now tainted).
	intoEncoder280.Encode(fromInterface482)

	// Return the tainted `intoEncoder280`:
	return intoEncoder280
}

func TaintStepTest_EncodingJsonEncoderSetIndent_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString120` into `intoEncoder274`.

	// Assume that `sourceCQL` has the underlying type of `fromString120`:
	fromString120 := sourceCQL.(string)

	// Declare `intoEncoder274` variable:
	var intoEncoder274 json.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromString120` to the receiver `intoEncoder274`
	// (`intoEncoder274` is now tainted).
	intoEncoder274.SetIndent(fromString120, "")

	// Return the tainted `intoEncoder274`:
	return intoEncoder274
}

func TaintStepTest_EncodingJsonEncoderSetIndent_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString443` into `intoEncoder508`.

	// Assume that `sourceCQL` has the underlying type of `fromString443`:
	fromString443 := sourceCQL.(string)

	// Declare `intoEncoder508` variable:
	var intoEncoder508 json.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromString443` to the receiver `intoEncoder508`
	// (`intoEncoder508` is now tainted).
	intoEncoder508.SetIndent("", fromString443)

	// Return the tainted `intoEncoder508`:
	return intoEncoder508
}

func TaintStepTest_EncodingJsonRawMessageMarshalJSON_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawMessage760` into `intoByte600`.

	// Assume that `sourceCQL` has the underlying type of `fromRawMessage760`:
	fromRawMessage760 := sourceCQL.(json.RawMessage)

	// Call the method that transfers the taint
	// from the receiver `fromRawMessage760` to the result `intoByte600`
	// (`intoByte600` is now tainted).
	intoByte600, _ := fromRawMessage760.MarshalJSON()

	// Return the tainted `intoByte600`:
	return intoByte600
}

func TaintStepTest_EncodingJsonRawMessageUnmarshalJSON_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte841` into `intoRawMessage249`.

	// Assume that `sourceCQL` has the underlying type of `fromByte841`:
	fromByte841 := sourceCQL.([]byte)

	// Declare `intoRawMessage249` variable:
	var intoRawMessage249 json.RawMessage

	// Call the method that transfers the taint
	// from the parameter `fromByte841` to the receiver `intoRawMessage249`
	// (`intoRawMessage249` is now tainted).
	intoRawMessage249.UnmarshalJSON(fromByte841)

	// Return the tainted `intoRawMessage249`:
	return intoRawMessage249
}

func TaintStepTest_EncodingJsonMarshalerMarshalJSON_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMarshaler623` into `intoByte610`.

	// Assume that `sourceCQL` has the underlying type of `fromMarshaler623`:
	fromMarshaler623 := sourceCQL.(json.Marshaler)

	// Call the method that transfers the taint
	// from the receiver `fromMarshaler623` to the result `intoByte610`
	// (`intoByte610` is now tainted).
	intoByte610, _ := fromMarshaler623.MarshalJSON()

	// Return the tainted `intoByte610`:
	return intoByte610
}

func TaintStepTest_EncodingJsonUnmarshalerUnmarshalJSON_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte876` into `intoUnmarshaler171`.

	// Assume that `sourceCQL` has the underlying type of `fromByte876`:
	fromByte876 := sourceCQL.([]byte)

	// Declare `intoUnmarshaler171` variable:
	var intoUnmarshaler171 json.Unmarshaler

	// Call the method that transfers the taint
	// from the parameter `fromByte876` to the receiver `intoUnmarshaler171`
	// (`intoUnmarshaler171` is now tainted).
	intoUnmarshaler171.UnmarshalJSON(fromByte876)

	// Return the tainted `intoUnmarshaler171`:
	return intoUnmarshaler171
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
