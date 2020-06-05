package main

import (
	"bytes"
	"encoding/json"
	"io"
)

func TaintStepTest_EncodingJsonCompact(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare `dst` variable:
	var dst *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `src` to parameter `dst`;
	// `dst` is now tainted.
	json.Compact(dst, src)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_EncodingJsonHTMLEscape(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare `dst` variable:
	var dst *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `src` to parameter `dst`;
	// `dst` is now tainted.
	json.HTMLEscape(dst, src)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_EncodingJsonIndent(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare `dst` variable:
	var dst *bytes.Buffer

	// Call the function that transfers the taint
	// from the parameter `src` to parameter `dst`;
	// `dst` is now tainted.
	json.Indent(dst, src, "", "")

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_EncodingJsonMarshal(sourceCQL interface{}) {
	// The flow is from `v` into `into838`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `v` to result `into838`
	// (`into838` is now tainted).
	into838, _ := json.Marshal(v)

	// Sink the tainted `into838`:
	sink(into838)
}

func TaintStepTest_EncodingJsonMarshalIndent(sourceCQL interface{}) {
	// The flow is from `v` into `into612`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `v` to result `into612`
	// (`into612` is now tainted).
	into612, _ := json.MarshalIndent(v, "", "")

	// Sink the tainted `into612`:
	sink(into612)
}

func TaintStepTest_EncodingJsonNewDecoder(sourceCQL interface{}) {
	// The flow is from `r` into `into574`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into574`
	// (`into574` is now tainted).
	into574 := json.NewDecoder(r)

	// Sink the tainted `into574`:
	sink(into574)
}

func TaintStepTest_EncodingJsonNewEncoder(sourceCQL interface{}) {
	// The flow is from `from351` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from351`:
	from351 := sourceCQL.(*json.Encoder)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := json.NewEncoder(w)

	// Extra step (`from351` taints `intermediateCQL`, which taints `w`:
	link(from351, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_EncodingJsonUnmarshal(sourceCQL interface{}) {
	// The flow is from `data` into `v`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.([]byte)

	// Declare `v` variable:
	var v interface{}

	// Call the function that transfers the taint
	// from the parameter `data` to parameter `v`;
	// `v` is now tainted.
	json.Unmarshal(data, v)

	// Sink the tainted `v`:
	sink(v)
}

func TaintStepTest_EncodingJsonDecoderBuffered(sourceCQL interface{}) {
	// The flow is from `from953` into `into669`.

	// Assume that `sourceCQL` has the underlying type of `from953`:
	from953 := sourceCQL.(json.Decoder)

	// Call the method that transfers the taint
	// from the receiver `from953` to the result `into669`
	// (`into669` is now tainted).
	into669 := from953.Buffered()

	// Sink the tainted `into669`:
	sink(into669)
}

func TaintStepTest_EncodingJsonDecoderDecode(sourceCQL interface{}) {
	// The flow is from `from230` into `v`.

	// Assume that `sourceCQL` has the underlying type of `from230`:
	from230 := sourceCQL.(json.Decoder)

	// Declare `v` variable:
	var v interface{}

	// Call the method that transfers the taint
	// from the receiver `from230` to the argument `v`
	// (`v` is now tainted).
	from230.Decode(v)

	// Sink the tainted `v`:
	sink(v)
}

func TaintStepTest_EncodingJsonDecoderToken(sourceCQL interface{}) {
	// The flow is from `from501` into `into923`.

	// Assume that `sourceCQL` has the underlying type of `from501`:
	from501 := sourceCQL.(json.Decoder)

	// Call the method that transfers the taint
	// from the receiver `from501` to the result `into923`
	// (`into923` is now tainted).
	into923, _ := from501.Token()

	// Sink the tainted `into923`:
	sink(into923)
}

func TaintStepTest_EncodingJsonEncoderEncode(sourceCQL interface{}) {
	// The flow is from `v` into `into742`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.(interface{})

	// Declare `into742` variable:
	var into742 json.Encoder

	// Call the method that transfers the taint
	// from the parameter `v` to the receiver `into742`
	// (`into742` is now tainted).
	into742.Encode(v)

	// Sink the tainted `into742`:
	sink(into742)
}

func TaintStepTest_EncodingJsonRawMessageMarshalJSON(sourceCQL interface{}) {
	// The flow is from `from220` into `into745`.

	// Assume that `sourceCQL` has the underlying type of `from220`:
	from220 := sourceCQL.(json.RawMessage)

	// Call the method that transfers the taint
	// from the receiver `from220` to the result `into745`
	// (`into745` is now tainted).
	into745, _ := from220.MarshalJSON()

	// Sink the tainted `into745`:
	sink(into745)
}

func TaintStepTest_EncodingJsonRawMessageUnmarshalJSON(sourceCQL interface{}) {
	// The flow is from `data` into `into448`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.([]byte)

	// Declare `into448` variable:
	var into448 json.RawMessage

	// Call the method that transfers the taint
	// from the parameter `data` to the receiver `into448`
	// (`into448` is now tainted).
	into448.UnmarshalJSON(data)

	// Sink the tainted `into448`:
	sink(into448)
}

func TaintStepTest_EncodingJsonMarshalerMarshalJSON(sourceCQL interface{}) {
	// The flow is from `from746` into `into661`.

	// Assume that `sourceCQL` has the underlying type of `from746`:
	from746 := sourceCQL.(json.Marshaler)

	// Call the method that transfers the taint
	// from the receiver `from746` to the result `into661`
	// (`into661` is now tainted).
	into661, _ := from746.MarshalJSON()

	// Sink the tainted `into661`:
	sink(into661)
}

func TaintStepTest_EncodingJsonUnmarshalerUnmarshalJSON(sourceCQL interface{}) {
	// The flow is from `from681` into `into516`.

	// Assume that `sourceCQL` has the underlying type of `from681`:
	from681 := sourceCQL.([]byte)

	// Declare `into516` variable:
	var into516 json.Unmarshaler

	// Call the method that transfers the taint
	// from the parameter `from681` to the receiver `into516`
	// (`into516` is now tainted).
	into516.UnmarshalJSON(from681)

	// Sink the tainted `into516`:
	sink(into516)
}

func RunAllTaints_EncodingJson(v interface{}) {
	{
		source := newSource()
		TaintStepTest_EncodingJsonCompact(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonHTMLEscape(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonIndent(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonMarshal(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonMarshalIndent(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonNewDecoder(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonNewEncoder(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonUnmarshal(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonDecoderBuffered(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonDecoderDecode(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonDecoderToken(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonEncoderEncode(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonRawMessageMarshalJSON(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonRawMessageUnmarshalJSON(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonMarshalerMarshalJSON(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingJsonUnmarshalerUnmarshalJSON(source)
	}
}
