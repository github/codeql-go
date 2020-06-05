/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `encoding/xml` package. */
module EncodingXmlTaintTracking {
  private class CopyToken extends TaintTracking::FunctionModel {
    // signature: func CopyToken(t Token) Token
    CopyToken() { hasQualifiedName("encoding/xml", "CopyToken") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class Escape extends TaintTracking::FunctionModel {
    // signature: func Escape(w io.Writer, s []byte)
    Escape() { hasQualifiedName("encoding/xml", "Escape") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class EscapeText extends TaintTracking::FunctionModel {
    // signature: func EscapeText(w io.Writer, s []byte) error
    EscapeText() { hasQualifiedName("encoding/xml", "EscapeText") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class Marshal extends TaintTracking::FunctionModel {
    // signature: func Marshal(v interface{}) ([]byte, error)
    Marshal() { hasQualifiedName("encoding/xml", "Marshal") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class MarshalIndent extends TaintTracking::FunctionModel {
    // signature: func MarshalIndent(v interface{}, prefix string, indent string) ([]byte, error)
    MarshalIndent() { hasQualifiedName("encoding/xml", "MarshalIndent") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class NewDecoder extends TaintTracking::FunctionModel {
    // signature: func NewDecoder(r io.Reader) *Decoder
    NewDecoder() { hasQualifiedName("encoding/xml", "NewDecoder") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class NewEncoder extends TaintTracking::FunctionModel {
    // signature: func NewEncoder(w io.Writer) *Encoder
    NewEncoder() { hasQualifiedName("encoding/xml", "NewEncoder") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult() and outp.isParameter(0)
    }
  }

  private class NewTokenDecoder extends TaintTracking::FunctionModel {
    // signature: func NewTokenDecoder(t TokenReader) *Decoder
    NewTokenDecoder() { hasQualifiedName("encoding/xml", "NewTokenDecoder") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class Unmarshal extends TaintTracking::FunctionModel {
    // signature: func Unmarshal(data []byte, v interface{}) error
    Unmarshal() { hasQualifiedName("encoding/xml", "Unmarshal") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isParameter(1)
    }
  }

  private class CharDataCopy extends TaintTracking::FunctionModel, Method {
    // signature: func (CharData).Copy() CharData
    CharDataCopy() { this.(Method).hasQualifiedName("encoding/xml", "CharData", "Copy") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class CommentCopy extends TaintTracking::FunctionModel, Method {
    // signature: func (Comment).Copy() Comment
    CommentCopy() { this.(Method).hasQualifiedName("encoding/xml", "Comment", "Copy") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class DecoderDecode extends TaintTracking::FunctionModel, Method {
    // signature: func (*Decoder).Decode(v interface{}) error
    DecoderDecode() { this.(Method).hasQualifiedName("encoding/xml", "Decoder", "Decode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isParameter(0)
    }
  }

  private class DecoderDecodeElement extends TaintTracking::FunctionModel, Method {
    // signature: func (*Decoder).DecodeElement(v interface{}, start *StartElement) error
    DecoderDecodeElement() {
      this.(Method).hasQualifiedName("encoding/xml", "Decoder", "DecodeElement")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isParameter(0)
    }
  }

  private class DecoderRawToken extends TaintTracking::FunctionModel, Method {
    // signature: func (*Decoder).RawToken() (Token, error)
    DecoderRawToken() { this.(Method).hasQualifiedName("encoding/xml", "Decoder", "RawToken") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class DecoderToken extends TaintTracking::FunctionModel, Method {
    // signature: func (*Decoder).Token() (Token, error)
    DecoderToken() { this.(Method).hasQualifiedName("encoding/xml", "Decoder", "Token") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class DirectiveCopy extends TaintTracking::FunctionModel, Method {
    // signature: func (Directive).Copy() Directive
    DirectiveCopy() { this.(Method).hasQualifiedName("encoding/xml", "Directive", "Copy") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class EncoderEncode extends TaintTracking::FunctionModel, Method {
    // signature: func (*Encoder).Encode(v interface{}) error
    EncoderEncode() { this.(Method).hasQualifiedName("encoding/xml", "Encoder", "Encode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class EncoderEncodeElement extends TaintTracking::FunctionModel, Method {
    // signature: func (*Encoder).EncodeElement(v interface{}, start StartElement) error
    EncoderEncodeElement() {
      this.(Method).hasQualifiedName("encoding/xml", "Encoder", "EncodeElement")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class EncoderEncodeToken extends TaintTracking::FunctionModel, Method {
    // signature: func (*Encoder).EncodeToken(t Token) error
    EncoderEncodeToken() {
      this.(Method).hasQualifiedName("encoding/xml", "Encoder", "EncodeToken")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ProcInstCopy extends TaintTracking::FunctionModel, Method {
    // signature: func (ProcInst).Copy() ProcInst
    ProcInstCopy() { this.(Method).hasQualifiedName("encoding/xml", "ProcInst", "Copy") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class StartElementCopy extends TaintTracking::FunctionModel, Method {
    // signature: func (StartElement).Copy() StartElement
    StartElementCopy() { this.(Method).hasQualifiedName("encoding/xml", "StartElement", "Copy") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class MarshalerMarshalXML extends TaintTracking::FunctionModel, Method {
    // signature: func (Marshaler).MarshalXML(e *Encoder, start StartElement) error
    MarshalerMarshalXML() { this.implements("encoding/xml", "Marshaler", "MarshalXML") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isParameter(0)
    }
  }

  private class TokenReaderToken extends TaintTracking::FunctionModel, Method {
    // signature: func (TokenReader).Token() (Token, error)
    TokenReaderToken() { this.implements("encoding/xml", "TokenReader", "Token") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class UnmarshalerUnmarshalXML extends TaintTracking::FunctionModel, Method {
    // signature: func (Unmarshaler).UnmarshalXML(d *Decoder, start StartElement) error
    UnmarshalerUnmarshalXML() { this.implements("encoding/xml", "Unmarshaler", "UnmarshalXML") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }
}
