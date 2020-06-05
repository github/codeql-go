package main

import (
	"encoding/xml"
	"io"
)

func TaintStepTest_EncodingXmlCopyToken(sourceCQL interface{}) {
	// The flow is from `t` into `into212`.

	// Assume that `sourceCQL` has the underlying type of `t`:
	t := sourceCQL.(xml.Token)

	// Call the function that transfers the taint
	// from the parameter `t` to result `into212`
	// (`into212` is now tainted).
	into212 := xml.CopyToken(t)

	// Sink the tainted `into212`:
	sink(into212)
}

func TaintStepTest_EncodingXmlEscape(sourceCQL interface{}) {
	// The flow is from `s` into `w`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that transfers the taint
	// from the parameter `s` to parameter `w`;
	// `w` is now tainted.
	xml.Escape(w, s)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_EncodingXmlEscapeText(sourceCQL interface{}) {
	// The flow is from `s` into `w`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that transfers the taint
	// from the parameter `s` to parameter `w`;
	// `w` is now tainted.
	xml.EscapeText(w, s)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_EncodingXmlMarshal(sourceCQL interface{}) {
	// The flow is from `v` into `into562`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `v` to result `into562`
	// (`into562` is now tainted).
	into562, _ := xml.Marshal(v)

	// Sink the tainted `into562`:
	sink(into562)
}

func TaintStepTest_EncodingXmlMarshalIndent(sourceCQL interface{}) {
	// The flow is from `v` into `into293`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `v` to result `into293`
	// (`into293` is now tainted).
	into293, _ := xml.MarshalIndent(v, "", "")

	// Sink the tainted `into293`:
	sink(into293)
}

func TaintStepTest_EncodingXmlNewDecoder(sourceCQL interface{}) {
	// The flow is from `r` into `into734`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into734`
	// (`into734` is now tainted).
	into734 := xml.NewDecoder(r)

	// Sink the tainted `into734`:
	sink(into734)
}

func TaintStepTest_EncodingXmlNewEncoder(sourceCQL interface{}) {
	// The flow is from `from321` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from321`:
	from321 := sourceCQL.(*xml.Encoder)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := xml.NewEncoder(w)

	// Extra step (`from321` taints `intermediateCQL`, which taints `w`:
	link(from321, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_EncodingXmlNewTokenDecoder(sourceCQL interface{}) {
	// The flow is from `t` into `into849`.

	// Assume that `sourceCQL` has the underlying type of `t`:
	t := sourceCQL.(xml.TokenReader)

	// Call the function that transfers the taint
	// from the parameter `t` to result `into849`
	// (`into849` is now tainted).
	into849 := xml.NewTokenDecoder(t)

	// Sink the tainted `into849`:
	sink(into849)
}

func TaintStepTest_EncodingXmlUnmarshal(sourceCQL interface{}) {
	// The flow is from `data` into `v`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.([]byte)

	// Declare `v` variable:
	var v interface{}

	// Call the function that transfers the taint
	// from the parameter `data` to parameter `v`;
	// `v` is now tainted.
	xml.Unmarshal(data, v)

	// Sink the tainted `v`:
	sink(v)
}

func TaintStepTest_EncodingXmlCharDataCopy(sourceCQL interface{}) {
	// The flow is from `from166` into `into136`.

	// Assume that `sourceCQL` has the underlying type of `from166`:
	from166 := sourceCQL.(xml.CharData)

	// Call the method that transfers the taint
	// from the receiver `from166` to the result `into136`
	// (`into136` is now tainted).
	into136 := from166.Copy()

	// Sink the tainted `into136`:
	sink(into136)
}

func TaintStepTest_EncodingXmlCommentCopy(sourceCQL interface{}) {
	// The flow is from `from757` into `into765`.

	// Assume that `sourceCQL` has the underlying type of `from757`:
	from757 := sourceCQL.(xml.Comment)

	// Call the method that transfers the taint
	// from the receiver `from757` to the result `into765`
	// (`into765` is now tainted).
	into765 := from757.Copy()

	// Sink the tainted `into765`:
	sink(into765)
}

func TaintStepTest_EncodingXmlDecoderDecode(sourceCQL interface{}) {
	// The flow is from `from111` into `v`.

	// Assume that `sourceCQL` has the underlying type of `from111`:
	from111 := sourceCQL.(xml.Decoder)

	// Declare `v` variable:
	var v interface{}

	// Call the method that transfers the taint
	// from the receiver `from111` to the argument `v`
	// (`v` is now tainted).
	from111.Decode(v)

	// Sink the tainted `v`:
	sink(v)
}

func TaintStepTest_EncodingXmlDecoderDecodeElement(sourceCQL interface{}) {
	// The flow is from `from528` into `v`.

	// Assume that `sourceCQL` has the underlying type of `from528`:
	from528 := sourceCQL.(xml.Decoder)

	// Declare `v` variable:
	var v interface{}

	// Call the method that transfers the taint
	// from the receiver `from528` to the argument `v`
	// (`v` is now tainted).
	from528.DecodeElement(v, nil)

	// Sink the tainted `v`:
	sink(v)
}

func TaintStepTest_EncodingXmlDecoderRawToken(sourceCQL interface{}) {
	// The flow is from `from708` into `into971`.

	// Assume that `sourceCQL` has the underlying type of `from708`:
	from708 := sourceCQL.(xml.Decoder)

	// Call the method that transfers the taint
	// from the receiver `from708` to the result `into971`
	// (`into971` is now tainted).
	into971, _ := from708.RawToken()

	// Sink the tainted `into971`:
	sink(into971)
}

func TaintStepTest_EncodingXmlDecoderToken(sourceCQL interface{}) {
	// The flow is from `from856` into `into830`.

	// Assume that `sourceCQL` has the underlying type of `from856`:
	from856 := sourceCQL.(xml.Decoder)

	// Call the method that transfers the taint
	// from the receiver `from856` to the result `into830`
	// (`into830` is now tainted).
	into830, _ := from856.Token()

	// Sink the tainted `into830`:
	sink(into830)
}

func TaintStepTest_EncodingXmlDirectiveCopy(sourceCQL interface{}) {
	// The flow is from `from136` into `into398`.

	// Assume that `sourceCQL` has the underlying type of `from136`:
	from136 := sourceCQL.(xml.Directive)

	// Call the method that transfers the taint
	// from the receiver `from136` to the result `into398`
	// (`into398` is now tainted).
	into398 := from136.Copy()

	// Sink the tainted `into398`:
	sink(into398)
}

func TaintStepTest_EncodingXmlEncoderEncode(sourceCQL interface{}) {
	// The flow is from `v` into `into644`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.(interface{})

	// Declare `into644` variable:
	var into644 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `v` to the receiver `into644`
	// (`into644` is now tainted).
	into644.Encode(v)

	// Sink the tainted `into644`:
	sink(into644)
}

func TaintStepTest_EncodingXmlEncoderEncodeElement(sourceCQL interface{}) {
	// The flow is from `v` into `into584`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.(interface{})

	// Declare `into584` variable:
	var into584 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `v` to the receiver `into584`
	// (`into584` is now tainted).
	into584.EncodeElement(v, xml.StartElement{})

	// Sink the tainted `into584`:
	sink(into584)
}

func TaintStepTest_EncodingXmlEncoderEncodeToken(sourceCQL interface{}) {
	// The flow is from `t` into `into340`.

	// Assume that `sourceCQL` has the underlying type of `t`:
	t := sourceCQL.(xml.Token)

	// Declare `into340` variable:
	var into340 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `t` to the receiver `into340`
	// (`into340` is now tainted).
	into340.EncodeToken(t)

	// Sink the tainted `into340`:
	sink(into340)
}

func TaintStepTest_EncodingXmlProcInstCopy(sourceCQL interface{}) {
	// The flow is from `from255` into `into763`.

	// Assume that `sourceCQL` has the underlying type of `from255`:
	from255 := sourceCQL.(xml.ProcInst)

	// Call the method that transfers the taint
	// from the receiver `from255` to the result `into763`
	// (`into763` is now tainted).
	into763 := from255.Copy()

	// Sink the tainted `into763`:
	sink(into763)
}

func TaintStepTest_EncodingXmlStartElementCopy(sourceCQL interface{}) {
	// The flow is from `from676` into `into130`.

	// Assume that `sourceCQL` has the underlying type of `from676`:
	from676 := sourceCQL.(xml.StartElement)

	// Call the method that transfers the taint
	// from the receiver `from676` to the result `into130`
	// (`into130` is now tainted).
	into130 := from676.Copy()

	// Sink the tainted `into130`:
	sink(into130)
}

func TaintStepTest_EncodingXmlMarshalerMarshalXML(sourceCQL interface{}) {
	// The flow is from `from671` into `e`.

	// Assume that `sourceCQL` has the underlying type of `from671`:
	from671 := sourceCQL.(xml.Marshaler)

	// Declare `e` variable:
	var e *xml.Encoder

	// Call the method that transfers the taint
	// from the receiver `from671` to the argument `e`
	// (`e` is now tainted).
	from671.MarshalXML(e, xml.StartElement{})

	// Sink the tainted `e`:
	sink(e)
}

func TaintStepTest_EncodingXmlTokenReaderToken(sourceCQL interface{}) {
	// The flow is from `from914` into `into415`.

	// Assume that `sourceCQL` has the underlying type of `from914`:
	from914 := sourceCQL.(xml.TokenReader)

	// Call the method that transfers the taint
	// from the receiver `from914` to the result `into415`
	// (`into415` is now tainted).
	into415, _ := from914.Token()

	// Sink the tainted `into415`:
	sink(into415)
}

func TaintStepTest_EncodingXmlUnmarshalerUnmarshalXML(sourceCQL interface{}) {
	// The flow is from `d` into `into259`.

	// Assume that `sourceCQL` has the underlying type of `d`:
	d := sourceCQL.(*xml.Decoder)

	// Declare `into259` variable:
	var into259 xml.Unmarshaler

	// Call the method that transfers the taint
	// from the parameter `d` to the receiver `into259`
	// (`into259` is now tainted).
	into259.UnmarshalXML(d, xml.StartElement{})

	// Sink the tainted `into259`:
	sink(into259)
}

func RunAllTaints_EncodingXml(v interface{}) {
	{
		source := newSource()
		TaintStepTest_EncodingXmlCopyToken(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlEscape(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlEscapeText(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlMarshal(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlMarshalIndent(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlNewDecoder(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlNewEncoder(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlNewTokenDecoder(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlUnmarshal(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlCharDataCopy(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlCommentCopy(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlDecoderDecode(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlDecoderDecodeElement(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlDecoderRawToken(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlDecoderToken(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlDirectiveCopy(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlEncoderEncode(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlEncoderEncodeElement(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlEncoderEncodeToken(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlProcInstCopy(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlStartElementCopy(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlMarshalerMarshalXML(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlTokenReaderToken(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlUnmarshalerUnmarshalXML(source)
	}
}
