/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `encoding/json` package. */
module EncodingJsonTaintTracking {
  private class Compact extends TaintTracking::FunctionModel {
    // signature: func Compact(dst *bytes.Buffer, src []byte) error
    Compact() { hasQualifiedName("encoding/json", "Compact") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class HTMLEscape extends TaintTracking::FunctionModel {
    // signature: func HTMLEscape(dst *bytes.Buffer, src []byte)
    HTMLEscape() { hasQualifiedName("encoding/json", "HTMLEscape") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class Indent extends TaintTracking::FunctionModel {
    // signature: func Indent(dst *bytes.Buffer, src []byte, prefix string, indent string) error
    Indent() { hasQualifiedName("encoding/json", "Indent") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class Marshal extends TaintTracking::FunctionModel {
    // signature: func Marshal(v interface{}) ([]byte, error)
    Marshal() { hasQualifiedName("encoding/json", "Marshal") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class MarshalIndent extends TaintTracking::FunctionModel {
    // signature: func MarshalIndent(v interface{}, prefix string, indent string) ([]byte, error)
    MarshalIndent() { hasQualifiedName("encoding/json", "MarshalIndent") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class NewDecoder extends TaintTracking::FunctionModel {
    // signature: func NewDecoder(r io.Reader) *Decoder
    NewDecoder() { hasQualifiedName("encoding/json", "NewDecoder") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class NewEncoder extends TaintTracking::FunctionModel {
    // signature: func NewEncoder(w io.Writer) *Encoder
    NewEncoder() { hasQualifiedName("encoding/json", "NewEncoder") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult() and outp.isParameter(0)
    }
  }

  private class Unmarshal extends TaintTracking::FunctionModel {
    // signature: func Unmarshal(data []byte, v interface{}) error
    Unmarshal() { hasQualifiedName("encoding/json", "Unmarshal") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isParameter(1)
    }
  }

  private class DecoderBuffered extends TaintTracking::FunctionModel, Method {
    // signature: func (*Decoder).Buffered() io.Reader
    DecoderBuffered() { this.(Method).hasQualifiedName("encoding/json", "Decoder", "Buffered") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class DecoderDecode extends TaintTracking::FunctionModel, Method {
    // signature: func (*Decoder).Decode(v interface{}) error
    DecoderDecode() { this.(Method).hasQualifiedName("encoding/json", "Decoder", "Decode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isParameter(0)
    }
  }

  private class DecoderToken extends TaintTracking::FunctionModel, Method {
    // signature: func (*Decoder).Token() (Token, error)
    DecoderToken() { this.(Method).hasQualifiedName("encoding/json", "Decoder", "Token") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class EncoderEncode extends TaintTracking::FunctionModel, Method {
    // signature: func (*Encoder).Encode(v interface{}) error
    EncoderEncode() { this.(Method).hasQualifiedName("encoding/json", "Encoder", "Encode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class RawMessageMarshalJSON extends TaintTracking::FunctionModel, Method {
    // signature: func (RawMessage).MarshalJSON() ([]byte, error)
    RawMessageMarshalJSON() {
      this.(Method).hasQualifiedName("encoding/json", "RawMessage", "MarshalJSON")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class RawMessageUnmarshalJSON extends TaintTracking::FunctionModel, Method {
    // signature: func (*RawMessage).UnmarshalJSON(data []byte) error
    RawMessageUnmarshalJSON() {
      this.(Method).hasQualifiedName("encoding/json", "RawMessage", "UnmarshalJSON")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class MarshalerMarshalJSON extends TaintTracking::FunctionModel, Method {
    // signature: func (Marshaler).MarshalJSON() ([]byte, error)
    MarshalerMarshalJSON() { this.implements("encoding/json", "Marshaler", "MarshalJSON") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class UnmarshalerUnmarshalJSON extends TaintTracking::FunctionModel, Method {
    // signature: func (Unmarshaler).UnmarshalJSON([]byte) error
    UnmarshalerUnmarshalJSON() { this.implements("encoding/json", "Unmarshaler", "UnmarshalJSON") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }
}
