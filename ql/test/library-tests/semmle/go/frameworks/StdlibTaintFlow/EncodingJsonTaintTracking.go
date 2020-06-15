package main

import (
	"bytes"
	"encoding/json"
	"io"
)

func TaintStepTest_EncodingJsonCompact_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte984` into `intoBuffer386`.

	// Assume that `sourceCQL` has the underlying type of `fromByte984`:
	fromByte984 := sourceCQL.([]byte)

	// Declare `intoBuffer386` variable:
	var intoBuffer386 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromByte984` to parameter `intoBuffer386`;
	// `intoBuffer386` is now tainted.
	json.Compact(intoBuffer386, fromByte984)

	// Sink the tainted `intoBuffer386`:
	sink(intoBuffer386)
}

func TaintStepTest_EncodingJsonHTMLEscape_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte952` into `intoBuffer780`.

	// Assume that `sourceCQL` has the underlying type of `fromByte952`:
	fromByte952 := sourceCQL.([]byte)

	// Declare `intoBuffer780` variable:
	var intoBuffer780 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromByte952` to parameter `intoBuffer780`;
	// `intoBuffer780` is now tainted.
	json.HTMLEscape(intoBuffer780, fromByte952)

	// Sink the tainted `intoBuffer780`:
	sink(intoBuffer780)
}

func TaintStepTest_EncodingJsonIndent_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte175` into `intoBuffer356`.

	// Assume that `sourceCQL` has the underlying type of `fromByte175`:
	fromByte175 := sourceCQL.([]byte)

	// Declare `intoBuffer356` variable:
	var intoBuffer356 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromByte175` to parameter `intoBuffer356`;
	// `intoBuffer356` is now tainted.
	json.Indent(intoBuffer356, fromByte175, "", "")

	// Sink the tainted `intoBuffer356`:
	sink(intoBuffer356)
}

func TaintStepTest_EncodingJsonIndent_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString223` into `intoBuffer988`.

	// Assume that `sourceCQL` has the underlying type of `fromString223`:
	fromString223 := sourceCQL.(string)

	// Declare `intoBuffer988` variable:
	var intoBuffer988 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromString223` to parameter `intoBuffer988`;
	// `intoBuffer988` is now tainted.
	json.Indent(intoBuffer988, nil, fromString223, "")

	// Sink the tainted `intoBuffer988`:
	sink(intoBuffer988)
}

func TaintStepTest_EncodingJsonIndent_B0I2O0(sourceCQL interface{}) {
	// The flow is from `fromString303` into `intoBuffer285`.

	// Assume that `sourceCQL` has the underlying type of `fromString303`:
	fromString303 := sourceCQL.(string)

	// Declare `intoBuffer285` variable:
	var intoBuffer285 *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `fromString303` to parameter `intoBuffer285`;
	// `intoBuffer285` is now tainted.
	json.Indent(intoBuffer285, nil, "", fromString303)

	// Sink the tainted `intoBuffer285`:
	sink(intoBuffer285)
}

func TaintStepTest_EncodingJsonMarshal_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface443` into `intoByte658`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface443`:
	fromInterface443 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface443` to result `intoByte658`
	// (`intoByte658` is now tainted).
	intoByte658, _ := json.Marshal(fromInterface443)

	// Sink the tainted `intoByte658`:
	sink(intoByte658)
}

func TaintStepTest_EncodingJsonMarshalIndent_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface704` into `intoByte418`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface704`:
	fromInterface704 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface704` to result `intoByte418`
	// (`intoByte418` is now tainted).
	intoByte418, _ := json.MarshalIndent(fromInterface704, "", "")

	// Sink the tainted `intoByte418`:
	sink(intoByte418)
}

func TaintStepTest_EncodingJsonMarshalIndent_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString457` into `intoByte902`.

	// Assume that `sourceCQL` has the underlying type of `fromString457`:
	fromString457 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString457` to result `intoByte902`
	// (`intoByte902` is now tainted).
	intoByte902, _ := json.MarshalIndent(nil, fromString457, "")

	// Sink the tainted `intoByte902`:
	sink(intoByte902)
}

func TaintStepTest_EncodingJsonMarshalIndent_B0I2O0(sourceCQL interface{}) {
	// The flow is from `fromString507` into `intoByte192`.

	// Assume that `sourceCQL` has the underlying type of `fromString507`:
	fromString507 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString507` to result `intoByte192`
	// (`intoByte192` is now tainted).
	intoByte192, _ := json.MarshalIndent(nil, "", fromString507)

	// Sink the tainted `intoByte192`:
	sink(intoByte192)
}

func TaintStepTest_EncodingJsonNewDecoder_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader114` into `intoDecoder207`.

	// Assume that `sourceCQL` has the underlying type of `fromReader114`:
	fromReader114 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader114` to result `intoDecoder207`
	// (`intoDecoder207` is now tainted).
	intoDecoder207 := json.NewDecoder(fromReader114)

	// Sink the tainted `intoDecoder207`:
	sink(intoDecoder207)
}

func TaintStepTest_EncodingJsonNewEncoder_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromEncoder428` into `intoWriter717`.

	// Assume that `sourceCQL` has the underlying type of `fromEncoder428`:
	fromEncoder428 := sourceCQL.(*json.Encoder)

	// Declare `intoWriter717` variable:
	var intoWriter717 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter717`:
	intermediateCQL := json.NewEncoder(intoWriter717)

	// Extra step (`fromEncoder428` taints `intermediateCQL`, which taints `intoWriter717`:
	link(fromEncoder428, intermediateCQL)

	// Sink the tainted `intoWriter717`:
	sink(intoWriter717)
}

func TaintStepTest_EncodingJsonUnmarshal_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte872` into `intoInterface827`.

	// Assume that `sourceCQL` has the underlying type of `fromByte872`:
	fromByte872 := sourceCQL.([]byte)

	// Declare `intoInterface827` variable:
	var intoInterface827 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromByte872` to parameter `intoInterface827`;
	// `intoInterface827` is now tainted.
	json.Unmarshal(fromByte872, intoInterface827)

	// Sink the tainted `intoInterface827`:
	sink(intoInterface827)
}

func TaintStepTest_EncodingJsonDecoderBuffered_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromDecoder346` into `intoReader131`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder346`:
	fromDecoder346 := sourceCQL.(json.Decoder)

	// Call the method that transfers the taint
	// from the receiver `fromDecoder346` to the result `intoReader131`
	// (`intoReader131` is now tainted).
	intoReader131 := fromDecoder346.Buffered()

	// Sink the tainted `intoReader131`:
	sink(intoReader131)
}

func TaintStepTest_EncodingJsonDecoderDecode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromDecoder492` into `intoInterface678`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder492`:
	fromDecoder492 := sourceCQL.(json.Decoder)

	// Declare `intoInterface678` variable:
	var intoInterface678 interface{}

	// Call the method that transfers the taint
	// from the receiver `fromDecoder492` to the argument `intoInterface678`
	// (`intoInterface678` is now tainted).
	fromDecoder492.Decode(intoInterface678)

	// Sink the tainted `intoInterface678`:
	sink(intoInterface678)
}

func TaintStepTest_EncodingJsonDecoderToken_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromDecoder545` into `intoToken197`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder545`:
	fromDecoder545 := sourceCQL.(json.Decoder)

	// Call the method that transfers the taint
	// from the receiver `fromDecoder545` to the result `intoToken197`
	// (`intoToken197` is now tainted).
	intoToken197, _ := fromDecoder545.Token()

	// Sink the tainted `intoToken197`:
	sink(intoToken197)
}

func TaintStepTest_EncodingJsonEncoderEncode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface283` into `intoEncoder585`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface283`:
	fromInterface283 := sourceCQL.(interface{})

	// Declare `intoEncoder585` variable:
	var intoEncoder585 json.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromInterface283` to the receiver `intoEncoder585`
	// (`intoEncoder585` is now tainted).
	intoEncoder585.Encode(fromInterface283)

	// Sink the tainted `intoEncoder585`:
	sink(intoEncoder585)
}

func TaintStepTest_EncodingJsonEncoderSetIndent_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString915` into `intoEncoder272`.

	// Assume that `sourceCQL` has the underlying type of `fromString915`:
	fromString915 := sourceCQL.(string)

	// Declare `intoEncoder272` variable:
	var intoEncoder272 json.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromString915` to the receiver `intoEncoder272`
	// (`intoEncoder272` is now tainted).
	intoEncoder272.SetIndent(fromString915, "")

	// Sink the tainted `intoEncoder272`:
	sink(intoEncoder272)
}

func TaintStepTest_EncodingJsonEncoderSetIndent_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString149` into `intoEncoder776`.

	// Assume that `sourceCQL` has the underlying type of `fromString149`:
	fromString149 := sourceCQL.(string)

	// Declare `intoEncoder776` variable:
	var intoEncoder776 json.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromString149` to the receiver `intoEncoder776`
	// (`intoEncoder776` is now tainted).
	intoEncoder776.SetIndent("", fromString149)

	// Sink the tainted `intoEncoder776`:
	sink(intoEncoder776)
}

func TaintStepTest_EncodingJsonRawMessageMarshalJSON_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRawMessage522` into `intoByte644`.

	// Assume that `sourceCQL` has the underlying type of `fromRawMessage522`:
	fromRawMessage522 := sourceCQL.(json.RawMessage)

	// Call the method that transfers the taint
	// from the receiver `fromRawMessage522` to the result `intoByte644`
	// (`intoByte644` is now tainted).
	intoByte644, _ := fromRawMessage522.MarshalJSON()

	// Sink the tainted `intoByte644`:
	sink(intoByte644)
}

func TaintStepTest_EncodingJsonRawMessageUnmarshalJSON_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte630` into `intoRawMessage331`.

	// Assume that `sourceCQL` has the underlying type of `fromByte630`:
	fromByte630 := sourceCQL.([]byte)

	// Declare `intoRawMessage331` variable:
	var intoRawMessage331 json.RawMessage

	// Call the method that transfers the taint
	// from the parameter `fromByte630` to the receiver `intoRawMessage331`
	// (`intoRawMessage331` is now tainted).
	intoRawMessage331.UnmarshalJSON(fromByte630)

	// Sink the tainted `intoRawMessage331`:
	sink(intoRawMessage331)
}

func TaintStepTest_EncodingJsonMarshalerMarshalJSON_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromMarshaler224` into `intoByte577`.

	// Assume that `sourceCQL` has the underlying type of `fromMarshaler224`:
	fromMarshaler224 := sourceCQL.(json.Marshaler)

	// Call the method that transfers the taint
	// from the receiver `fromMarshaler224` to the result `intoByte577`
	// (`intoByte577` is now tainted).
	intoByte577, _ := fromMarshaler224.MarshalJSON()

	// Sink the tainted `intoByte577`:
	sink(intoByte577)
}

func TaintStepTest_EncodingJsonUnmarshalerUnmarshalJSON_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte913` into `intoUnmarshaler199`.

	// Assume that `sourceCQL` has the underlying type of `fromByte913`:
	fromByte913 := sourceCQL.([]byte)

	// Declare `intoUnmarshaler199` variable:
	var intoUnmarshaler199 json.Unmarshaler

	// Call the method that transfers the taint
	// from the parameter `fromByte913` to the receiver `intoUnmarshaler199`
	// (`intoUnmarshaler199` is now tainted).
	intoUnmarshaler199.UnmarshalJSON(fromByte913)

	// Sink the tainted `intoUnmarshaler199`:
	sink(intoUnmarshaler199)
}

func RunAllTaints_EncodingJson() {
	{
		source := newSource()
		TaintStepTest_EncodingJsonCompact_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonHTMLEscape_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonIndent_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonIndent_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonIndent_B0I2O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonMarshal_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonMarshalIndent_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonMarshalIndent_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonMarshalIndent_B0I2O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonNewDecoder_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonNewEncoder_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonUnmarshal_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonDecoderBuffered_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonDecoderDecode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonDecoderToken_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonEncoderEncode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonEncoderSetIndent_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonEncoderSetIndent_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonRawMessageMarshalJSON_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonRawMessageUnmarshalJSON_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonMarshalerMarshalJSON_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonUnmarshalerUnmarshalJSON_B0I0O0(source)
	}
}
