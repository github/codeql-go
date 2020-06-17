package main

import (
	"encoding/xml"
	"io"
)

func TaintStepTest_EncodingXmlCopyToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromToken787` into `intoToken618`.

	// Assume that `sourceCQL` has the underlying type of `fromToken787`:
	fromToken787 := sourceCQL.(xml.Token)

	// Call the function that transfers the taint
	// from the parameter `fromToken787` to result `intoToken618`
	// (`intoToken618` is now tainted).
	intoToken618 := xml.CopyToken(fromToken787)

	// Return the tainted `intoToken618`:
	return intoToken618
}

func TaintStepTest_EncodingXmlEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte249` into `intoWriter672`.

	// Assume that `sourceCQL` has the underlying type of `fromByte249`:
	fromByte249 := sourceCQL.([]byte)

	// Declare `intoWriter672` variable:
	var intoWriter672 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte249` to parameter `intoWriter672`;
	// `intoWriter672` is now tainted.
	xml.Escape(intoWriter672, fromByte249)

	// Return the tainted `intoWriter672`:
	return intoWriter672
}

func TaintStepTest_EncodingXmlEscapeText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte485` into `intoWriter626`.

	// Assume that `sourceCQL` has the underlying type of `fromByte485`:
	fromByte485 := sourceCQL.([]byte)

	// Declare `intoWriter626` variable:
	var intoWriter626 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte485` to parameter `intoWriter626`;
	// `intoWriter626` is now tainted.
	xml.EscapeText(intoWriter626, fromByte485)

	// Return the tainted `intoWriter626`:
	return intoWriter626
}

func TaintStepTest_EncodingXmlMarshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface672` into `intoByte984`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface672`:
	fromInterface672 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface672` to result `intoByte984`
	// (`intoByte984` is now tainted).
	intoByte984, _ := xml.Marshal(fromInterface672)

	// Return the tainted `intoByte984`:
	return intoByte984
}

func TaintStepTest_EncodingXmlMarshalIndent_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface112` into `intoByte287`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface112`:
	fromInterface112 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface112` to result `intoByte287`
	// (`intoByte287` is now tainted).
	intoByte287, _ := xml.MarshalIndent(fromInterface112, "", "")

	// Return the tainted `intoByte287`:
	return intoByte287
}

func TaintStepTest_EncodingXmlMarshalIndent_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString586` into `intoByte542`.

	// Assume that `sourceCQL` has the underlying type of `fromString586`:
	fromString586 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString586` to result `intoByte542`
	// (`intoByte542` is now tainted).
	intoByte542, _ := xml.MarshalIndent(nil, fromString586, "")

	// Return the tainted `intoByte542`:
	return intoByte542
}

func TaintStepTest_EncodingXmlMarshalIndent_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString367` into `intoByte420`.

	// Assume that `sourceCQL` has the underlying type of `fromString367`:
	fromString367 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString367` to result `intoByte420`
	// (`intoByte420` is now tainted).
	intoByte420, _ := xml.MarshalIndent(nil, "", fromString367)

	// Return the tainted `intoByte420`:
	return intoByte420
}

func TaintStepTest_EncodingXmlNewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader617` into `intoDecoder837`.

	// Assume that `sourceCQL` has the underlying type of `fromReader617`:
	fromReader617 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader617` to result `intoDecoder837`
	// (`intoDecoder837` is now tainted).
	intoDecoder837 := xml.NewDecoder(fromReader617)

	// Return the tainted `intoDecoder837`:
	return intoDecoder837
}

func TaintStepTest_EncodingXmlNewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromEncoder508` into `intoWriter215`.

	// Assume that `sourceCQL` has the underlying type of `fromEncoder508`:
	fromEncoder508 := sourceCQL.(*xml.Encoder)

	// Declare `intoWriter215` variable:
	var intoWriter215 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter215`:
	intermediateCQL := xml.NewEncoder(intoWriter215)

	// Extra step (`fromEncoder508` taints `intermediateCQL`, which taints `intoWriter215`:
	link(fromEncoder508, intermediateCQL)

	// Return the tainted `intoWriter215`:
	return intoWriter215
}

func TaintStepTest_EncodingXmlNewTokenDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTokenReader208` into `intoDecoder788`.

	// Assume that `sourceCQL` has the underlying type of `fromTokenReader208`:
	fromTokenReader208 := sourceCQL.(xml.TokenReader)

	// Call the function that transfers the taint
	// from the parameter `fromTokenReader208` to result `intoDecoder788`
	// (`intoDecoder788` is now tainted).
	intoDecoder788 := xml.NewTokenDecoder(fromTokenReader208)

	// Return the tainted `intoDecoder788`:
	return intoDecoder788
}

func TaintStepTest_EncodingXmlUnmarshal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte656` into `intoInterface599`.

	// Assume that `sourceCQL` has the underlying type of `fromByte656`:
	fromByte656 := sourceCQL.([]byte)

	// Declare `intoInterface599` variable:
	var intoInterface599 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromByte656` to parameter `intoInterface599`;
	// `intoInterface599` is now tainted.
	xml.Unmarshal(fromByte656, intoInterface599)

	// Return the tainted `intoInterface599`:
	return intoInterface599
}

func TaintStepTest_EncodingXmlCharDataCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromCharData391` into `intoCharData285`.

	// Assume that `sourceCQL` has the underlying type of `fromCharData391`:
	fromCharData391 := sourceCQL.(xml.CharData)

	// Call the method that transfers the taint
	// from the receiver `fromCharData391` to the result `intoCharData285`
	// (`intoCharData285` is now tainted).
	intoCharData285 := fromCharData391.Copy()

	// Return the tainted `intoCharData285`:
	return intoCharData285
}

func TaintStepTest_EncodingXmlCommentCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromComment947` into `intoComment613`.

	// Assume that `sourceCQL` has the underlying type of `fromComment947`:
	fromComment947 := sourceCQL.(xml.Comment)

	// Call the method that transfers the taint
	// from the receiver `fromComment947` to the result `intoComment613`
	// (`intoComment613` is now tainted).
	intoComment613 := fromComment947.Copy()

	// Return the tainted `intoComment613`:
	return intoComment613
}

func TaintStepTest_EncodingXmlDecoderDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder380` into `intoInterface842`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder380`:
	fromDecoder380 := sourceCQL.(xml.Decoder)

	// Declare `intoInterface842` variable:
	var intoInterface842 interface{}

	// Call the method that transfers the taint
	// from the receiver `fromDecoder380` to the argument `intoInterface842`
	// (`intoInterface842` is now tainted).
	fromDecoder380.Decode(intoInterface842)

	// Return the tainted `intoInterface842`:
	return intoInterface842
}

func TaintStepTest_EncodingXmlDecoderDecodeElement_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder794` into `intoInterface748`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder794`:
	fromDecoder794 := sourceCQL.(xml.Decoder)

	// Declare `intoInterface748` variable:
	var intoInterface748 interface{}

	// Call the method that transfers the taint
	// from the receiver `fromDecoder794` to the argument `intoInterface748`
	// (`intoInterface748` is now tainted).
	fromDecoder794.DecodeElement(intoInterface748, nil)

	// Return the tainted `intoInterface748`:
	return intoInterface748
}

func TaintStepTest_EncodingXmlDecoderRawToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder998` into `intoToken594`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder998`:
	fromDecoder998 := sourceCQL.(xml.Decoder)

	// Call the method that transfers the taint
	// from the receiver `fromDecoder998` to the result `intoToken594`
	// (`intoToken594` is now tainted).
	intoToken594, _ := fromDecoder998.RawToken()

	// Return the tainted `intoToken594`:
	return intoToken594
}

func TaintStepTest_EncodingXmlDecoderToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder141` into `intoToken687`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder141`:
	fromDecoder141 := sourceCQL.(xml.Decoder)

	// Call the method that transfers the taint
	// from the receiver `fromDecoder141` to the result `intoToken687`
	// (`intoToken687` is now tainted).
	intoToken687, _ := fromDecoder141.Token()

	// Return the tainted `intoToken687`:
	return intoToken687
}

func TaintStepTest_EncodingXmlDirectiveCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDirective403` into `intoDirective494`.

	// Assume that `sourceCQL` has the underlying type of `fromDirective403`:
	fromDirective403 := sourceCQL.(xml.Directive)

	// Call the method that transfers the taint
	// from the receiver `fromDirective403` to the result `intoDirective494`
	// (`intoDirective494` is now tainted).
	intoDirective494 := fromDirective403.Copy()

	// Return the tainted `intoDirective494`:
	return intoDirective494
}

func TaintStepTest_EncodingXmlEncoderEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface671` into `intoEncoder644`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface671`:
	fromInterface671 := sourceCQL.(interface{})

	// Declare `intoEncoder644` variable:
	var intoEncoder644 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromInterface671` to the receiver `intoEncoder644`
	// (`intoEncoder644` is now tainted).
	intoEncoder644.Encode(fromInterface671)

	// Return the tainted `intoEncoder644`:
	return intoEncoder644
}

func TaintStepTest_EncodingXmlEncoderEncodeElement_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface631` into `intoEncoder602`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface631`:
	fromInterface631 := sourceCQL.(interface{})

	// Declare `intoEncoder602` variable:
	var intoEncoder602 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromInterface631` to the receiver `intoEncoder602`
	// (`intoEncoder602` is now tainted).
	intoEncoder602.EncodeElement(fromInterface631, xml.StartElement{})

	// Return the tainted `intoEncoder602`:
	return intoEncoder602
}

func TaintStepTest_EncodingXmlEncoderEncodeToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromToken940` into `intoEncoder538`.

	// Assume that `sourceCQL` has the underlying type of `fromToken940`:
	fromToken940 := sourceCQL.(xml.Token)

	// Declare `intoEncoder538` variable:
	var intoEncoder538 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromToken940` to the receiver `intoEncoder538`
	// (`intoEncoder538` is now tainted).
	intoEncoder538.EncodeToken(fromToken940)

	// Return the tainted `intoEncoder538`:
	return intoEncoder538
}

func TaintStepTest_EncodingXmlEncoderIndent_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString126` into `intoEncoder230`.

	// Assume that `sourceCQL` has the underlying type of `fromString126`:
	fromString126 := sourceCQL.(string)

	// Declare `intoEncoder230` variable:
	var intoEncoder230 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromString126` to the receiver `intoEncoder230`
	// (`intoEncoder230` is now tainted).
	intoEncoder230.Indent(fromString126, "")

	// Return the tainted `intoEncoder230`:
	return intoEncoder230
}

func TaintStepTest_EncodingXmlEncoderIndent_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString211` into `intoEncoder714`.

	// Assume that `sourceCQL` has the underlying type of `fromString211`:
	fromString211 := sourceCQL.(string)

	// Declare `intoEncoder714` variable:
	var intoEncoder714 xml.Encoder

	// Call the method that transfers the taint
	// from the parameter `fromString211` to the receiver `intoEncoder714`
	// (`intoEncoder714` is now tainted).
	intoEncoder714.Indent("", fromString211)

	// Return the tainted `intoEncoder714`:
	return intoEncoder714
}

func TaintStepTest_EncodingXmlProcInstCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromProcInst471` into `intoProcInst450`.

	// Assume that `sourceCQL` has the underlying type of `fromProcInst471`:
	fromProcInst471 := sourceCQL.(xml.ProcInst)

	// Call the method that transfers the taint
	// from the receiver `fromProcInst471` to the result `intoProcInst450`
	// (`intoProcInst450` is now tainted).
	intoProcInst450 := fromProcInst471.Copy()

	// Return the tainted `intoProcInst450`:
	return intoProcInst450
}

func TaintStepTest_EncodingXmlStartElementCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStartElement146` into `intoStartElement760`.

	// Assume that `sourceCQL` has the underlying type of `fromStartElement146`:
	fromStartElement146 := sourceCQL.(xml.StartElement)

	// Call the method that transfers the taint
	// from the receiver `fromStartElement146` to the result `intoStartElement760`
	// (`intoStartElement760` is now tainted).
	intoStartElement760 := fromStartElement146.Copy()

	// Return the tainted `intoStartElement760`:
	return intoStartElement760
}

func TaintStepTest_EncodingXmlMarshalerMarshalXML_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMarshaler696` into `intoEncoder172`.

	// Assume that `sourceCQL` has the underlying type of `fromMarshaler696`:
	fromMarshaler696 := sourceCQL.(xml.Marshaler)

	// Declare `intoEncoder172` variable:
	var intoEncoder172 *xml.Encoder

	// Call the method that transfers the taint
	// from the receiver `fromMarshaler696` to the argument `intoEncoder172`
	// (`intoEncoder172` is now tainted).
	fromMarshaler696.MarshalXML(intoEncoder172, xml.StartElement{})

	// Return the tainted `intoEncoder172`:
	return intoEncoder172
}

func TaintStepTest_EncodingXmlTokenReaderToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTokenReader232` into `intoToken880`.

	// Assume that `sourceCQL` has the underlying type of `fromTokenReader232`:
	fromTokenReader232 := sourceCQL.(xml.TokenReader)

	// Call the method that transfers the taint
	// from the receiver `fromTokenReader232` to the result `intoToken880`
	// (`intoToken880` is now tainted).
	intoToken880, _ := fromTokenReader232.Token()

	// Return the tainted `intoToken880`:
	return intoToken880
}

func TaintStepTest_EncodingXmlUnmarshalerUnmarshalXML_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromDecoder513` into `intoUnmarshaler248`.

	// Assume that `sourceCQL` has the underlying type of `fromDecoder513`:
	fromDecoder513 := sourceCQL.(*xml.Decoder)

	// Declare `intoUnmarshaler248` variable:
	var intoUnmarshaler248 xml.Unmarshaler

	// Call the method that transfers the taint
	// from the parameter `fromDecoder513` to the receiver `intoUnmarshaler248`
	// (`intoUnmarshaler248` is now tainted).
	intoUnmarshaler248.UnmarshalXML(fromDecoder513, xml.StartElement{})

	// Return the tainted `intoUnmarshaler248`:
	return intoUnmarshaler248
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
