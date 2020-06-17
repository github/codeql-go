// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"encoding/xml"
	"io"
)

func TaintStepTest_EncodingXmlCopyToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromToken820` into `intoToken646`.

	// Assume that `sourceCQL` has the underlying type of `fromToken820`:
	fromToken820 := sourceCQL.(xml.Token)

	// Call the function that transfers the taint
	// from the parameter `fromToken820` to result `intoToken646`
	// (`intoToken646` is now tainted).
	intoToken646 := xml.CopyToken(fromToken820)

	// Return the tainted `intoToken646`:
	return intoToken646
}

func TaintStepTest_EncodingXmlEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte802` into `intoWriter161`.

	// Assume that `sourceCQL` has the underlying type of `fromByte802`:
	fromByte802 := sourceCQL.([]byte)

	// Declare `intoWriter161` variable:
	var intoWriter161 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte802` to parameter `intoWriter161`;
	// `intoWriter161` is now tainted.
	xml.Escape(intoWriter161, fromByte802)

	// Return the tainted `intoWriter161`:
	return intoWriter161
}

func TaintStepTest_EncodingXmlEscapeText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte793` into `intoWriter966`.

	// Assume that `sourceCQL` has the underlying type of `fromByte793`:
	fromByte793 := sourceCQL.([]byte)

	// Declare `intoWriter966` variable:
	var intoWriter966 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte793` to parameter `intoWriter966`;
	// `intoWriter966` is now tainted.
	xml.EscapeText(intoWriter966, fromByte793)

	// Return the tainted `intoWriter966`:
	return intoWriter966
}

func TaintStepTest_EncodingXmlMarshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface112` into `intoByte569`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface112`:
	fromInterface112 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface112` to result `intoByte569`
	// (`intoByte569` is now tainted).
	intoByte569, _ := xml.Marshal(fromInterface112)

	// Return the tainted `intoByte569`:
	return intoByte569
}

func TaintStepTest_EncodingXmlMarshalIndent_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface113` into `intoByte187`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface113`:
	fromInterface113 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface113` to result `intoByte187`
	// (`intoByte187` is now tainted).
	intoByte187, _ := xml.MarshalIndent(fromInterface113, "", "")

	// Return the tainted `intoByte187`:
	return intoByte187
}

func TaintStepTest_EncodingXmlMarshalIndent_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString158` into `intoByte946`.

	// Assume that `sourceCQL` has the underlying type of `fromString158`:
	fromString158 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString158` to result `intoByte946`
	// (`intoByte946` is now tainted).
	intoByte946, _ := xml.MarshalIndent(nil, fromString158, "")

	// Return the tainted `intoByte946`:
	return intoByte946
}

func TaintStepTest_EncodingXmlMarshalIndent_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString857` into `intoByte611`.

	// Assume that `sourceCQL` has the underlying type of `fromString857`:
	fromString857 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString857` to result `intoByte611`
	// (`intoByte611` is now tainted).
	intoByte611, _ := xml.MarshalIndent(nil, "", fromString857)

	// Return the tainted `intoByte611`:
	return intoByte611
}

func TaintStepTest_EncodingXmlNewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader656` into `intoDecoder598`.

	// Assume that `sourceCQL` has the underlying type of `fromReader656`:
	fromReader656 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader656` to result `intoDecoder598`
	// (`intoDecoder598` is now tainted).
	intoDecoder598 := xml.NewDecoder(fromReader656)

	// Return the tainted `intoDecoder598`:
	return intoDecoder598
}

func TaintStepTest_EncodingXmlNewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromEncoder413` into `intoWriter586`.

	// Assume that `sourceCQL` has the underlying type of `fromEncoder413`:
	fromEncoder413 := sourceCQL.(*xml.Encoder)

	// Declare `intoWriter586` variable:
	var intoWriter586 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter586`:
	intermediateCQL := xml.NewEncoder(intoWriter586)

	// Extra step (`fromEncoder413` taints `intermediateCQL`, which taints `intoWriter586`:
	link(fromEncoder413, intermediateCQL)

	// Return the tainted `intoWriter586`:
	return intoWriter586
}

func TaintStepTest_EncodingXmlNewTokenDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTokenReader926` into `intoDecoder815`.

	// Assume that `sourceCQL` has the underlying type of `fromTokenReader926`:
	fromTokenReader926 := sourceCQL.(xml.TokenReader)

	// Call the function that transfers the taint
	// from the parameter `fromTokenReader926` to result `intoDecoder815`
	// (`intoDecoder815` is now tainted).
	intoDecoder815 := xml.NewTokenDecoder(fromTokenReader926)

	// Return the tainted `intoDecoder815`:
	return intoDecoder815
}

func TaintStepTest_EncodingXmlUnmarshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte903` into `intoInterface939`.

	// Assume that `sourceCQL` has the underlying type of `fromByte903`:
	fromByte903 := sourceCQL.([]byte)

	// Declare `intoInterface939` variable:
	var intoInterface939 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromByte903` to parameter `intoInterface939`;
	// `intoInterface939` is now tainted.
	xml.Unmarshal(fromByte903, intoInterface939)

	// Return the tainted `intoInterface939`:
	return intoInterface939
}

func TaintStepTest_EncodingXmlCharDataCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromCharData335` into `intoCharData699`.

	// Assume that `sourceCQL` has the underlying type of `fromCharData335`:
	fromCharData335 := sourceCQL.(xml.CharData)

	// Call the method that transfers the taint
	// from the receiver `fromCharData335` to the result `intoCharData699`
	// (`intoCharData699` is now tainted).
	intoCharData699 := fromCharData335.Copy()

	// Return the tainted `intoCharData699`:
	return intoCharData699
}

func TaintStepTest_EncodingXmlCommentCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromComment566` into `intoComment511`.

	// Assume that `sourceCQL` has the underlying type of `fromComment566`:
	fromComment566 := sourceCQL.(xml.Comment)

	// Call the method that transfers the taint
	// from the receiver `fromComment566` to the result `intoComment511`
	// (`intoComment511` is now tainted).
	intoComment511 := fromComment566.Copy()

	// Return the tainted `intoComment511`:
	return intoComment511
}

func TaintStepTest_EncodingXmlDecoderDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder544` into `intoInterface465`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder544`:
	fromDecoder544 := sourceCQL.(xml.Decoder)

	// Declare `intoInterface465` variable:
	var intoInterface465 interface{}

	// Call the method that transfers the taint
	// from the receiver `fromDecoder544` to the argument `intoInterface465`
	// (`intoInterface465` is now tainted).
	fromDecoder544.Decode(intoInterface465)

	// Return the tainted `intoInterface465`:
	return intoInterface465
}

func TaintStepTest_EncodingXmlDecoderDecodeElement_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder869` into `intoInterface830`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder869`:
	fromDecoder869 := sourceCQL.(xml.Decoder)

	// Declare `intoInterface830` variable:
	var intoInterface830 interface{}

	// Call the method that transfers the taint
	// from the receiver `fromDecoder869` to the argument `intoInterface830`
	// (`intoInterface830` is now tainted).
	fromDecoder869.DecodeElement(intoInterface830, nil)

	// Return the tainted `intoInterface830`:
	return intoInterface830
}

func TaintStepTest_EncodingXmlDecoderRawToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder914` into `intoToken343`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder914`:
	fromDecoder914 := sourceCQL.(xml.Decoder)

	// Call the method that transfers the taint
	// from the receiver `fromDecoder914` to the result `intoToken343`
	// (`intoToken343` is now tainted).
	intoToken343, _ := fromDecoder914.RawToken()

	// Return the tainted `intoToken343`:
	return intoToken343
}

func TaintStepTest_EncodingXmlDecoderToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder974` into `intoToken254`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder974`:
	fromDecoder974 := sourceCQL.(xml.Decoder)

	// Call the method that transfers the taint
	// from the receiver `fromDecoder974` to the result `intoToken254`
	// (`intoToken254` is now tainted).
	intoToken254, _ := fromDecoder974.Token()

	// Return the tainted `intoToken254`:
	return intoToken254
}

func TaintStepTest_EncodingXmlDirectiveCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDirective792` into `intoDirective953`.

	// Assume that `sourceCQL` has the underlying type of `fromDirective792`:
	fromDirective792 := sourceCQL.(xml.Directive)

	// Call the method that transfers the taint
	// from the receiver `fromDirective792` to the result `intoDirective953`
	// (`intoDirective953` is now tainted).
	intoDirective953 := fromDirective792.Copy()

	// Return the tainted `intoDirective953`:
	return intoDirective953
}

func TaintStepTest_EncodingXmlEncoderEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface783` into `intoEncoder967`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface783`:
	fromInterface783 := sourceCQL.(interface{})

	// Declare `intoEncoder967` variable:
	var intoEncoder967 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromInterface783` to the receiver `intoEncoder967`
	// (`intoEncoder967` is now tainted).
	intoEncoder967.Encode(fromInterface783)

	// Return the tainted `intoEncoder967`:
	return intoEncoder967
}

func TaintStepTest_EncodingXmlEncoderEncodeElement_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface403` into `intoEncoder467`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface403`:
	fromInterface403 := sourceCQL.(interface{})

	// Declare `intoEncoder467` variable:
	var intoEncoder467 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromInterface403` to the receiver `intoEncoder467`
	// (`intoEncoder467` is now tainted).
	intoEncoder467.EncodeElement(fromInterface403, xml.StartElement{})

	// Return the tainted `intoEncoder467`:
	return intoEncoder467
}

func TaintStepTest_EncodingXmlEncoderEncodeToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromToken759` into `intoEncoder637`.

	// Assume that `sourceCQL` has the underlying type of `fromToken759`:
	fromToken759 := sourceCQL.(xml.Token)

	// Declare `intoEncoder637` variable:
	var intoEncoder637 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromToken759` to the receiver `intoEncoder637`
	// (`intoEncoder637` is now tainted).
	intoEncoder637.EncodeToken(fromToken759)

	// Return the tainted `intoEncoder637`:
	return intoEncoder637
}

func TaintStepTest_EncodingXmlEncoderIndent_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString659` into `intoEncoder831`.

	// Assume that `sourceCQL` has the underlying type of `fromString659`:
	fromString659 := sourceCQL.(string)

	// Declare `intoEncoder831` variable:
	var intoEncoder831 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromString659` to the receiver `intoEncoder831`
	// (`intoEncoder831` is now tainted).
	intoEncoder831.Indent(fromString659, "")

	// Return the tainted `intoEncoder831`:
	return intoEncoder831
}

func TaintStepTest_EncodingXmlEncoderIndent_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString616` into `intoEncoder795`.

	// Assume that `sourceCQL` has the underlying type of `fromString616`:
	fromString616 := sourceCQL.(string)

	// Declare `intoEncoder795` variable:
	var intoEncoder795 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromString616` to the receiver `intoEncoder795`
	// (`intoEncoder795` is now tainted).
	intoEncoder795.Indent("", fromString616)

	// Return the tainted `intoEncoder795`:
	return intoEncoder795
}

func TaintStepTest_EncodingXmlProcInstCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromProcInst864` into `intoProcInst887`.

	// Assume that `sourceCQL` has the underlying type of `fromProcInst864`:
	fromProcInst864 := sourceCQL.(xml.ProcInst)

	// Call the method that transfers the taint
	// from the receiver `fromProcInst864` to the result `intoProcInst887`
	// (`intoProcInst887` is now tainted).
	intoProcInst887 := fromProcInst864.Copy()

	// Return the tainted `intoProcInst887`:
	return intoProcInst887
}

func TaintStepTest_EncodingXmlStartElementCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStartElement640` into `intoStartElement683`.

	// Assume that `sourceCQL` has the underlying type of `fromStartElement640`:
	fromStartElement640 := sourceCQL.(xml.StartElement)

	// Call the method that transfers the taint
	// from the receiver `fromStartElement640` to the result `intoStartElement683`
	// (`intoStartElement683` is now tainted).
	intoStartElement683 := fromStartElement640.Copy()

	// Return the tainted `intoStartElement683`:
	return intoStartElement683
}

func TaintStepTest_EncodingXmlMarshalerMarshalXML_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMarshaler308` into `intoEncoder383`.

	// Assume that `sourceCQL` has the underlying type of `fromMarshaler308`:
	fromMarshaler308 := sourceCQL.(xml.Marshaler)

	// Declare `intoEncoder383` variable:
	var intoEncoder383 *xml.Encoder

	// Call the method that transfers the taint
	// from the receiver `fromMarshaler308` to the argument `intoEncoder383`
	// (`intoEncoder383` is now tainted).
	fromMarshaler308.MarshalXML(intoEncoder383, xml.StartElement{})

	// Return the tainted `intoEncoder383`:
	return intoEncoder383
}

func TaintStepTest_EncodingXmlTokenReaderToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTokenReader951` into `intoToken425`.

	// Assume that `sourceCQL` has the underlying type of `fromTokenReader951`:
	fromTokenReader951 := sourceCQL.(xml.TokenReader)

	// Call the method that transfers the taint
	// from the receiver `fromTokenReader951` to the result `intoToken425`
	// (`intoToken425` is now tainted).
	intoToken425, _ := fromTokenReader951.Token()

	// Return the tainted `intoToken425`:
	return intoToken425
}

func TaintStepTest_EncodingXmlUnmarshalerUnmarshalXML_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder948` into `intoUnmarshaler763`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder948`:
	fromDecoder948 := sourceCQL.(*xml.Decoder)

	// Declare `intoUnmarshaler763` variable:
	var intoUnmarshaler763 xml.Unmarshaler

	// Call the method that transfers the taint
	// from the parameter `fromDecoder948` to the receiver `intoUnmarshaler763`
	// (`intoUnmarshaler763` is now tainted).
	intoUnmarshaler763.UnmarshalXML(fromDecoder948, xml.StartElement{})

	// Return the tainted `intoUnmarshaler763`:
	return intoUnmarshaler763
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
