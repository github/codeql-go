// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"encoding/xml"
	"io"
)

func TaintStepTest_EncodingXmlCopyToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromToken656` into `intoToken414`.

	// Assume that `sourceCQL` has the underlying type of `fromToken656`:
	fromToken656 := sourceCQL.(xml.Token)

	// Call the function that transfers the taint
	// from the parameter `fromToken656` to result `intoToken414`
	// (`intoToken414` is now tainted).
	intoToken414 := xml.CopyToken(fromToken656)

	// Return the tainted `intoToken414`:
	return intoToken414
}

func TaintStepTest_EncodingXmlEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte518` into `intoWriter650`.

	// Assume that `sourceCQL` has the underlying type of `fromByte518`:
	fromByte518 := sourceCQL.([]byte)

	// Declare `intoWriter650` variable:
	var intoWriter650 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte518` to parameter `intoWriter650`;
	// `intoWriter650` is now tainted.
	xml.Escape(intoWriter650, fromByte518)

	// Return the tainted `intoWriter650`:
	return intoWriter650
}

func TaintStepTest_EncodingXmlEscapeText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte784` into `intoWriter957`.

	// Assume that `sourceCQL` has the underlying type of `fromByte784`:
	fromByte784 := sourceCQL.([]byte)

	// Declare `intoWriter957` variable:
	var intoWriter957 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte784` to parameter `intoWriter957`;
	// `intoWriter957` is now tainted.
	xml.EscapeText(intoWriter957, fromByte784)

	// Return the tainted `intoWriter957`:
	return intoWriter957
}

func TaintStepTest_EncodingXmlMarshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface520` into `intoByte443`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface520`:
	fromInterface520 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface520` to result `intoByte443`
	// (`intoByte443` is now tainted).
	intoByte443, _ := xml.Marshal(fromInterface520)

	// Return the tainted `intoByte443`:
	return intoByte443
}

func TaintStepTest_EncodingXmlMarshalIndent_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface127` into `intoByte483`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface127`:
	fromInterface127 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface127` to result `intoByte483`
	// (`intoByte483` is now tainted).
	intoByte483, _ := xml.MarshalIndent(fromInterface127, "", "")

	// Return the tainted `intoByte483`:
	return intoByte483
}

func TaintStepTest_EncodingXmlMarshalIndent_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString989` into `intoByte982`.

	// Assume that `sourceCQL` has the underlying type of `fromString989`:
	fromString989 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString989` to result `intoByte982`
	// (`intoByte982` is now tainted).
	intoByte982, _ := xml.MarshalIndent(nil, fromString989, "")

	// Return the tainted `intoByte982`:
	return intoByte982
}

func TaintStepTest_EncodingXmlMarshalIndent_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString417` into `intoByte584`.

	// Assume that `sourceCQL` has the underlying type of `fromString417`:
	fromString417 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString417` to result `intoByte584`
	// (`intoByte584` is now tainted).
	intoByte584, _ := xml.MarshalIndent(nil, "", fromString417)

	// Return the tainted `intoByte584`:
	return intoByte584
}

func TaintStepTest_EncodingXmlNewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader991` into `intoDecoder881`.

	// Assume that `sourceCQL` has the underlying type of `fromReader991`:
	fromReader991 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader991` to result `intoDecoder881`
	// (`intoDecoder881` is now tainted).
	intoDecoder881 := xml.NewDecoder(fromReader991)

	// Return the tainted `intoDecoder881`:
	return intoDecoder881
}

func TaintStepTest_EncodingXmlNewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromEncoder186` into `intoWriter284`.

	// Assume that `sourceCQL` has the underlying type of `fromEncoder186`:
	fromEncoder186 := sourceCQL.(*xml.Encoder)

	// Declare `intoWriter284` variable:
	var intoWriter284 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter284`:
	intermediateCQL := xml.NewEncoder(intoWriter284)

	// Extra step (`fromEncoder186` taints `intermediateCQL`, which taints `intoWriter284`:
	link(fromEncoder186, intermediateCQL)

	// Return the tainted `intoWriter284`:
	return intoWriter284
}

func TaintStepTest_EncodingXmlNewTokenDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTokenReader908` into `intoDecoder137`.

	// Assume that `sourceCQL` has the underlying type of `fromTokenReader908`:
	fromTokenReader908 := sourceCQL.(xml.TokenReader)

	// Call the function that transfers the taint
	// from the parameter `fromTokenReader908` to result `intoDecoder137`
	// (`intoDecoder137` is now tainted).
	intoDecoder137 := xml.NewTokenDecoder(fromTokenReader908)

	// Return the tainted `intoDecoder137`:
	return intoDecoder137
}

func TaintStepTest_EncodingXmlUnmarshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte494` into `intoInterface873`.

	// Assume that `sourceCQL` has the underlying type of `fromByte494`:
	fromByte494 := sourceCQL.([]byte)

	// Declare `intoInterface873` variable:
	var intoInterface873 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromByte494` to parameter `intoInterface873`;
	// `intoInterface873` is now tainted.
	xml.Unmarshal(fromByte494, intoInterface873)

	// Return the tainted `intoInterface873`:
	return intoInterface873
}

func TaintStepTest_EncodingXmlCharDataCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromCharData599` into `intoCharData409`.

	// Assume that `sourceCQL` has the underlying type of `fromCharData599`:
	fromCharData599 := sourceCQL.(xml.CharData)

	// Call the method that transfers the taint
	// from the receiver `fromCharData599` to the result `intoCharData409`
	// (`intoCharData409` is now tainted).
	intoCharData409 := fromCharData599.Copy()

	// Return the tainted `intoCharData409`:
	return intoCharData409
}

func TaintStepTest_EncodingXmlCommentCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromComment246` into `intoComment898`.

	// Assume that `sourceCQL` has the underlying type of `fromComment246`:
	fromComment246 := sourceCQL.(xml.Comment)

	// Call the method that transfers the taint
	// from the receiver `fromComment246` to the result `intoComment898`
	// (`intoComment898` is now tainted).
	intoComment898 := fromComment246.Copy()

	// Return the tainted `intoComment898`:
	return intoComment898
}

func TaintStepTest_EncodingXmlDecoderDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder598` into `intoInterface631`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder598`:
	fromDecoder598 := sourceCQL.(xml.Decoder)

	// Declare `intoInterface631` variable:
	var intoInterface631 interface{}

	// Call the method that transfers the taint
	// from the receiver `fromDecoder598` to the argument `intoInterface631`
	// (`intoInterface631` is now tainted).
	fromDecoder598.Decode(intoInterface631)

	// Return the tainted `intoInterface631`:
	return intoInterface631
}

func TaintStepTest_EncodingXmlDecoderDecodeElement_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder165` into `intoInterface150`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder165`:
	fromDecoder165 := sourceCQL.(xml.Decoder)

	// Declare `intoInterface150` variable:
	var intoInterface150 interface{}

	// Call the method that transfers the taint
	// from the receiver `fromDecoder165` to the argument `intoInterface150`
	// (`intoInterface150` is now tainted).
	fromDecoder165.DecodeElement(intoInterface150, nil)

	// Return the tainted `intoInterface150`:
	return intoInterface150
}

func TaintStepTest_EncodingXmlDecoderRawToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder340` into `intoToken471`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder340`:
	fromDecoder340 := sourceCQL.(xml.Decoder)

	// Call the method that transfers the taint
	// from the receiver `fromDecoder340` to the result `intoToken471`
	// (`intoToken471` is now tainted).
	intoToken471, _ := fromDecoder340.RawToken()

	// Return the tainted `intoToken471`:
	return intoToken471
}

func TaintStepTest_EncodingXmlDecoderToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder290` into `intoToken758`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder290`:
	fromDecoder290 := sourceCQL.(xml.Decoder)

	// Call the method that transfers the taint
	// from the receiver `fromDecoder290` to the result `intoToken758`
	// (`intoToken758` is now tainted).
	intoToken758, _ := fromDecoder290.Token()

	// Return the tainted `intoToken758`:
	return intoToken758
}

func TaintStepTest_EncodingXmlDirectiveCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDirective396` into `intoDirective707`.

	// Assume that `sourceCQL` has the underlying type of `fromDirective396`:
	fromDirective396 := sourceCQL.(xml.Directive)

	// Call the method that transfers the taint
	// from the receiver `fromDirective396` to the result `intoDirective707`
	// (`intoDirective707` is now tainted).
	intoDirective707 := fromDirective396.Copy()

	// Return the tainted `intoDirective707`:
	return intoDirective707
}

func TaintStepTest_EncodingXmlEncoderEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface912` into `intoEncoder718`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface912`:
	fromInterface912 := sourceCQL.(interface{})

	// Declare `intoEncoder718` variable:
	var intoEncoder718 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromInterface912` to the receiver `intoEncoder718`
	// (`intoEncoder718` is now tainted).
	intoEncoder718.Encode(fromInterface912)

	// Return the tainted `intoEncoder718`:
	return intoEncoder718
}

func TaintStepTest_EncodingXmlEncoderEncodeElement_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface972` into `intoEncoder633`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface972`:
	fromInterface972 := sourceCQL.(interface{})

	// Declare `intoEncoder633` variable:
	var intoEncoder633 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromInterface972` to the receiver `intoEncoder633`
	// (`intoEncoder633` is now tainted).
	intoEncoder633.EncodeElement(fromInterface972, xml.StartElement{})

	// Return the tainted `intoEncoder633`:
	return intoEncoder633
}

func TaintStepTest_EncodingXmlEncoderEncodeToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromToken316` into `intoEncoder145`.

	// Assume that `sourceCQL` has the underlying type of `fromToken316`:
	fromToken316 := sourceCQL.(xml.Token)

	// Declare `intoEncoder145` variable:
	var intoEncoder145 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromToken316` to the receiver `intoEncoder145`
	// (`intoEncoder145` is now tainted).
	intoEncoder145.EncodeToken(fromToken316)

	// Return the tainted `intoEncoder145`:
	return intoEncoder145
}

func TaintStepTest_EncodingXmlEncoderIndent_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString817` into `intoEncoder474`.

	// Assume that `sourceCQL` has the underlying type of `fromString817`:
	fromString817 := sourceCQL.(string)

	// Declare `intoEncoder474` variable:
	var intoEncoder474 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromString817` to the receiver `intoEncoder474`
	// (`intoEncoder474` is now tainted).
	intoEncoder474.Indent(fromString817, "")

	// Return the tainted `intoEncoder474`:
	return intoEncoder474
}

func TaintStepTest_EncodingXmlEncoderIndent_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString832` into `intoEncoder378`.

	// Assume that `sourceCQL` has the underlying type of `fromString832`:
	fromString832 := sourceCQL.(string)

	// Declare `intoEncoder378` variable:
	var intoEncoder378 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromString832` to the receiver `intoEncoder378`
	// (`intoEncoder378` is now tainted).
	intoEncoder378.Indent("", fromString832)

	// Return the tainted `intoEncoder378`:
	return intoEncoder378
}

func TaintStepTest_EncodingXmlProcInstCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromProcInst541` into `intoProcInst139`.

	// Assume that `sourceCQL` has the underlying type of `fromProcInst541`:
	fromProcInst541 := sourceCQL.(xml.ProcInst)

	// Call the method that transfers the taint
	// from the receiver `fromProcInst541` to the result `intoProcInst139`
	// (`intoProcInst139` is now tainted).
	intoProcInst139 := fromProcInst541.Copy()

	// Return the tainted `intoProcInst139`:
	return intoProcInst139
}

func TaintStepTest_EncodingXmlStartElementCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStartElement814` into `intoStartElement768`.

	// Assume that `sourceCQL` has the underlying type of `fromStartElement814`:
	fromStartElement814 := sourceCQL.(xml.StartElement)

	// Call the method that transfers the taint
	// from the receiver `fromStartElement814` to the result `intoStartElement768`
	// (`intoStartElement768` is now tainted).
	intoStartElement768 := fromStartElement814.Copy()

	// Return the tainted `intoStartElement768`:
	return intoStartElement768
}

func TaintStepTest_EncodingXmlMarshalerMarshalXML_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMarshaler468` into `intoEncoder736`.

	// Assume that `sourceCQL` has the underlying type of `fromMarshaler468`:
	fromMarshaler468 := sourceCQL.(xml.Marshaler)

	// Declare `intoEncoder736` variable:
	var intoEncoder736 *xml.Encoder

	// Call the method that transfers the taint
	// from the receiver `fromMarshaler468` to the argument `intoEncoder736`
	// (`intoEncoder736` is now tainted).
	fromMarshaler468.MarshalXML(intoEncoder736, xml.StartElement{})

	// Return the tainted `intoEncoder736`:
	return intoEncoder736
}

func TaintStepTest_EncodingXmlTokenReaderToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTokenReader516` into `intoToken246`.

	// Assume that `sourceCQL` has the underlying type of `fromTokenReader516`:
	fromTokenReader516 := sourceCQL.(xml.TokenReader)

	// Call the method that transfers the taint
	// from the receiver `fromTokenReader516` to the result `intoToken246`
	// (`intoToken246` is now tainted).
	intoToken246, _ := fromTokenReader516.Token()

	// Return the tainted `intoToken246`:
	return intoToken246
}

func TaintStepTest_EncodingXmlUnmarshalerUnmarshalXML_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder679` into `intoUnmarshaler736`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder679`:
	fromDecoder679 := sourceCQL.(*xml.Decoder)

	// Declare `intoUnmarshaler736` variable:
	var intoUnmarshaler736 xml.Unmarshaler

	// Call the method that transfers the taint
	// from the parameter `fromDecoder679` to the receiver `intoUnmarshaler736`
	// (`intoUnmarshaler736` is now tainted).
	intoUnmarshaler736.UnmarshalXML(fromDecoder679, xml.StartElement{})

	// Return the tainted `intoUnmarshaler736`:
	return intoUnmarshaler736
}

func RunAllTaints_EncodingXml() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlCopyToken_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlEscape_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlEscapeText_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlMarshal_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlMarshalIndent_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlMarshalIndent_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlMarshalIndent_B0I2O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlNewDecoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlNewEncoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlNewTokenDecoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlUnmarshal_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlCharDataCopy_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlCommentCopy_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlDecoderDecode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlDecoderDecodeElement_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlDecoderRawToken_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlDecoderToken_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlDirectiveCopy_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlEncoderEncode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlEncoderEncodeElement_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlEncoderEncodeToken_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlEncoderIndent_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlEncoderIndent_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlProcInstCopy_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlStartElementCopy_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlMarshalerMarshalXML_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlTokenReaderToken_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingXmlUnmarshalerUnmarshalXML_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
