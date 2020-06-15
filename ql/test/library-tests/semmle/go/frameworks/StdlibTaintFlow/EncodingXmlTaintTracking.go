package main

import (
	"encoding/xml"
	"io"
)

func TaintStepTest_EncodingXmlCopyToken_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromToken260` into `intoToken694`.

	// Assume that `sourceCQL` has the underlying type of `fromToken260`:
	fromToken260 := sourceCQL.(xml.Token)

	// Call the function that transfers the taint
	// from the parameter `fromToken260` to result `intoToken694`
	// (`intoToken694` is now tainted).
	intoToken694 := xml.CopyToken(fromToken260)

	// Sink the tainted `intoToken694`:
	sink(intoToken694)
}

func TaintStepTest_EncodingXmlEscape_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte214` into `intoWriter508`.

	// Assume that `sourceCQL` has the underlying type of `fromByte214`:
	fromByte214 := sourceCQL.([]byte)

	// Declare `intoWriter508` variable:
	var intoWriter508 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte214` to parameter `intoWriter508`;
	// `intoWriter508` is now tainted.
	xml.Escape(intoWriter508, fromByte214)

	// Sink the tainted `intoWriter508`:
	sink(intoWriter508)
}

func TaintStepTest_EncodingXmlEscapeText_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte742` into `intoWriter626`.

	// Assume that `sourceCQL` has the underlying type of `fromByte742`:
	fromByte742 := sourceCQL.([]byte)

	// Declare `intoWriter626` variable:
	var intoWriter626 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte742` to parameter `intoWriter626`;
	// `intoWriter626` is now tainted.
	xml.EscapeText(intoWriter626, fromByte742)

	// Sink the tainted `intoWriter626`:
	sink(intoWriter626)
}

func TaintStepTest_EncodingXmlMarshal_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface156` into `intoByte410`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface156`:
	fromInterface156 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface156` to result `intoByte410`
	// (`intoByte410` is now tainted).
	intoByte410, _ := xml.Marshal(fromInterface156)

	// Sink the tainted `intoByte410`:
	sink(intoByte410)
}

func TaintStepTest_EncodingXmlMarshalIndent_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface451` into `intoByte986`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface451`:
	fromInterface451 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface451` to result `intoByte986`
	// (`intoByte986` is now tainted).
	intoByte986, _ := xml.MarshalIndent(fromInterface451, "", "")

	// Sink the tainted `intoByte986`:
	sink(intoByte986)
}

func TaintStepTest_EncodingXmlMarshalIndent_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString478` into `intoByte705`.

	// Assume that `sourceCQL` has the underlying type of `fromString478`:
	fromString478 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString478` to result `intoByte705`
	// (`intoByte705` is now tainted).
	intoByte705, _ := xml.MarshalIndent(nil, fromString478, "")

	// Sink the tainted `intoByte705`:
	sink(intoByte705)
}

func TaintStepTest_EncodingXmlMarshalIndent_B0I2O0(sourceCQL interface{}) {
	// The flow is from `fromString665` into `intoByte906`.

	// Assume that `sourceCQL` has the underlying type of `fromString665`:
	fromString665 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString665` to result `intoByte906`
	// (`intoByte906` is now tainted).
	intoByte906, _ := xml.MarshalIndent(nil, "", fromString665)

	// Sink the tainted `intoByte906`:
	sink(intoByte906)
}

func TaintStepTest_EncodingXmlNewDecoder_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader988` into `intoDecoder606`.

	// Assume that `sourceCQL` has the underlying type of `fromReader988`:
	fromReader988 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader988` to result `intoDecoder606`
	// (`intoDecoder606` is now tainted).
	intoDecoder606 := xml.NewDecoder(fromReader988)

	// Sink the tainted `intoDecoder606`:
	sink(intoDecoder606)
}

func TaintStepTest_EncodingXmlNewEncoder_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromEncoder881` into `intoWriter427`.

	// Assume that `sourceCQL` has the underlying type of `fromEncoder881`:
	fromEncoder881 := sourceCQL.(*xml.Encoder)

	// Declare `intoWriter427` variable:
	var intoWriter427 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter427`:
	intermediateCQL := xml.NewEncoder(intoWriter427)

	// Extra step (`fromEncoder881` taints `intermediateCQL`, which taints `intoWriter427`:
	link(fromEncoder881, intermediateCQL)

	// Sink the tainted `intoWriter427`:
	sink(intoWriter427)
}

func TaintStepTest_EncodingXmlNewTokenDecoder_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTokenReader469` into `intoDecoder237`.

	// Assume that `sourceCQL` has the underlying type of `fromTokenReader469`:
	fromTokenReader469 := sourceCQL.(xml.TokenReader)

	// Call the function that transfers the taint
	// from the parameter `fromTokenReader469` to result `intoDecoder237`
	// (`intoDecoder237` is now tainted).
	intoDecoder237 := xml.NewTokenDecoder(fromTokenReader469)

	// Sink the tainted `intoDecoder237`:
	sink(intoDecoder237)
}

func TaintStepTest_EncodingXmlUnmarshal_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte936` into `intoInterface115`.

	// Assume that `sourceCQL` has the underlying type of `fromByte936`:
	fromByte936 := sourceCQL.([]byte)

	// Declare `intoInterface115` variable:
	var intoInterface115 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromByte936` to parameter `intoInterface115`;
	// `intoInterface115` is now tainted.
	xml.Unmarshal(fromByte936, intoInterface115)

	// Sink the tainted `intoInterface115`:
	sink(intoInterface115)
}

func TaintStepTest_EncodingXmlCharDataCopy_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromCharData258` into `intoCharData414`.

	// Assume that `sourceCQL` has the underlying type of `fromCharData258`:
	fromCharData258 := sourceCQL.(xml.CharData)

	// Call the method that transfers the taint
	// from the receiver `fromCharData258` to the result `intoCharData414`
	// (`intoCharData414` is now tainted).
	intoCharData414 := fromCharData258.Copy()

	// Sink the tainted `intoCharData414`:
	sink(intoCharData414)
}

func TaintStepTest_EncodingXmlCommentCopy_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromComment336` into `intoComment205`.

	// Assume that `sourceCQL` has the underlying type of `fromComment336`:
	fromComment336 := sourceCQL.(xml.Comment)

	// Call the method that transfers the taint
	// from the receiver `fromComment336` to the result `intoComment205`
	// (`intoComment205` is now tainted).
	intoComment205 := fromComment336.Copy()

	// Sink the tainted `intoComment205`:
	sink(intoComment205)
}

func TaintStepTest_EncodingXmlDecoderDecode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromDecoder753` into `intoInterface666`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder753`:
	fromDecoder753 := sourceCQL.(xml.Decoder)

	// Declare `intoInterface666` variable:
	var intoInterface666 interface{}

	// Call the method that transfers the taint
	// from the receiver `fromDecoder753` to the argument `intoInterface666`
	// (`intoInterface666` is now tainted).
	fromDecoder753.Decode(intoInterface666)

	// Sink the tainted `intoInterface666`:
	sink(intoInterface666)
}

func TaintStepTest_EncodingXmlDecoderDecodeElement_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromDecoder687` into `intoInterface153`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder687`:
	fromDecoder687 := sourceCQL.(xml.Decoder)

	// Declare `intoInterface153` variable:
	var intoInterface153 interface{}

	// Call the method that transfers the taint
	// from the receiver `fromDecoder687` to the argument `intoInterface153`
	// (`intoInterface153` is now tainted).
	fromDecoder687.DecodeElement(intoInterface153, nil)

	// Sink the tainted `intoInterface153`:
	sink(intoInterface153)
}

func TaintStepTest_EncodingXmlDecoderRawToken_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromDecoder278` into `intoToken559`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder278`:
	fromDecoder278 := sourceCQL.(xml.Decoder)

	// Call the method that transfers the taint
	// from the receiver `fromDecoder278` to the result `intoToken559`
	// (`intoToken559` is now tainted).
	intoToken559, _ := fromDecoder278.RawToken()

	// Sink the tainted `intoToken559`:
	sink(intoToken559)
}

func TaintStepTest_EncodingXmlDecoderToken_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromDecoder574` into `intoToken460`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder574`:
	fromDecoder574 := sourceCQL.(xml.Decoder)

	// Call the method that transfers the taint
	// from the receiver `fromDecoder574` to the result `intoToken460`
	// (`intoToken460` is now tainted).
	intoToken460, _ := fromDecoder574.Token()

	// Sink the tainted `intoToken460`:
	sink(intoToken460)
}

func TaintStepTest_EncodingXmlDirectiveCopy_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromDirective440` into `intoDirective675`.

	// Assume that `sourceCQL` has the underlying type of `fromDirective440`:
	fromDirective440 := sourceCQL.(xml.Directive)

	// Call the method that transfers the taint
	// from the receiver `fromDirective440` to the result `intoDirective675`
	// (`intoDirective675` is now tainted).
	intoDirective675 := fromDirective440.Copy()

	// Sink the tainted `intoDirective675`:
	sink(intoDirective675)
}

func TaintStepTest_EncodingXmlEncoderEncode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface814` into `intoEncoder943`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface814`:
	fromInterface814 := sourceCQL.(interface{})

	// Declare `intoEncoder943` variable:
	var intoEncoder943 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromInterface814` to the receiver `intoEncoder943`
	// (`intoEncoder943` is now tainted).
	intoEncoder943.Encode(fromInterface814)

	// Sink the tainted `intoEncoder943`:
	sink(intoEncoder943)
}

func TaintStepTest_EncodingXmlEncoderEncodeElement_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface541` into `intoEncoder276`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface541`:
	fromInterface541 := sourceCQL.(interface{})

	// Declare `intoEncoder276` variable:
	var intoEncoder276 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromInterface541` to the receiver `intoEncoder276`
	// (`intoEncoder276` is now tainted).
	intoEncoder276.EncodeElement(fromInterface541, xml.StartElement{})

	// Sink the tainted `intoEncoder276`:
	sink(intoEncoder276)
}

func TaintStepTest_EncodingXmlEncoderEncodeToken_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromToken957` into `intoEncoder249`.

	// Assume that `sourceCQL` has the underlying type of `fromToken957`:
	fromToken957 := sourceCQL.(xml.Token)

	// Declare `intoEncoder249` variable:
	var intoEncoder249 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromToken957` to the receiver `intoEncoder249`
	// (`intoEncoder249` is now tainted).
	intoEncoder249.EncodeToken(fromToken957)

	// Sink the tainted `intoEncoder249`:
	sink(intoEncoder249)
}

func TaintStepTest_EncodingXmlEncoderIndent_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString428` into `intoEncoder634`.

	// Assume that `sourceCQL` has the underlying type of `fromString428`:
	fromString428 := sourceCQL.(string)

	// Declare `intoEncoder634` variable:
	var intoEncoder634 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromString428` to the receiver `intoEncoder634`
	// (`intoEncoder634` is now tainted).
	intoEncoder634.Indent(fromString428, "")

	// Sink the tainted `intoEncoder634`:
	sink(intoEncoder634)
}

func TaintStepTest_EncodingXmlEncoderIndent_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString874` into `intoEncoder357`.

	// Assume that `sourceCQL` has the underlying type of `fromString874`:
	fromString874 := sourceCQL.(string)

	// Declare `intoEncoder357` variable:
	var intoEncoder357 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromString874` to the receiver `intoEncoder357`
	// (`intoEncoder357` is now tainted).
	intoEncoder357.Indent("", fromString874)

	// Sink the tainted `intoEncoder357`:
	sink(intoEncoder357)
}

func TaintStepTest_EncodingXmlProcInstCopy_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromProcInst890` into `intoProcInst366`.

	// Assume that `sourceCQL` has the underlying type of `fromProcInst890`:
	fromProcInst890 := sourceCQL.(xml.ProcInst)

	// Call the method that transfers the taint
	// from the receiver `fromProcInst890` to the result `intoProcInst366`
	// (`intoProcInst366` is now tainted).
	intoProcInst366 := fromProcInst890.Copy()

	// Sink the tainted `intoProcInst366`:
	sink(intoProcInst366)
}

func TaintStepTest_EncodingXmlStartElementCopy_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromStartElement307` into `intoStartElement756`.

	// Assume that `sourceCQL` has the underlying type of `fromStartElement307`:
	fromStartElement307 := sourceCQL.(xml.StartElement)

	// Call the method that transfers the taint
	// from the receiver `fromStartElement307` to the result `intoStartElement756`
	// (`intoStartElement756` is now tainted).
	intoStartElement756 := fromStartElement307.Copy()

	// Sink the tainted `intoStartElement756`:
	sink(intoStartElement756)
}

func TaintStepTest_EncodingXmlMarshalerMarshalXML_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromMarshaler841` into `intoEncoder322`.

	// Assume that `sourceCQL` has the underlying type of `fromMarshaler841`:
	fromMarshaler841 := sourceCQL.(xml.Marshaler)

	// Declare `intoEncoder322` variable:
	var intoEncoder322 *xml.Encoder

	// Call the method that transfers the taint
	// from the receiver `fromMarshaler841` to the argument `intoEncoder322`
	// (`intoEncoder322` is now tainted).
	fromMarshaler841.MarshalXML(intoEncoder322, xml.StartElement{})

	// Sink the tainted `intoEncoder322`:
	sink(intoEncoder322)
}

func TaintStepTest_EncodingXmlTokenReaderToken_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTokenReader992` into `intoToken141`.

	// Assume that `sourceCQL` has the underlying type of `fromTokenReader992`:
	fromTokenReader992 := sourceCQL.(xml.TokenReader)

	// Call the method that transfers the taint
	// from the receiver `fromTokenReader992` to the result `intoToken141`
	// (`intoToken141` is now tainted).
	intoToken141, _ := fromTokenReader992.Token()

	// Sink the tainted `intoToken141`:
	sink(intoToken141)
}

func TaintStepTest_EncodingXmlUnmarshalerUnmarshalXML_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromDecoder598` into `intoUnmarshaler864`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder598`:
	fromDecoder598 := sourceCQL.(*xml.Decoder)

	// Declare `intoUnmarshaler864` variable:
	var intoUnmarshaler864 xml.Unmarshaler

	// Call the method that transfers the taint
	// from the parameter `fromDecoder598` to the receiver `intoUnmarshaler864`
	// (`intoUnmarshaler864` is now tainted).
	intoUnmarshaler864.UnmarshalXML(fromDecoder598, xml.StartElement{})

	// Sink the tainted `intoUnmarshaler864`:
	sink(intoUnmarshaler864)
}

func RunAllTaints_EncodingXml() {
	{
		source := newSource()
		TaintStepTest_EncodingXmlCopyToken_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlEscape_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlEscapeText_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlMarshal_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlMarshalIndent_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlMarshalIndent_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlMarshalIndent_B0I2O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlNewDecoder_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlNewEncoder_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlNewTokenDecoder_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlUnmarshal_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlCharDataCopy_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlCommentCopy_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlDecoderDecode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlDecoderDecodeElement_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlDecoderRawToken_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlDecoderToken_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlDirectiveCopy_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlEncoderEncode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlEncoderEncodeElement_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlEncoderEncodeToken_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlEncoderIndent_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlEncoderIndent_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlProcInstCopy_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlStartElementCopy_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlMarshalerMarshalXML_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlTokenReaderToken_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingXmlUnmarshalerUnmarshalXML_B0I0O0(source)
	}
}
