/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `encoding/gob` package. */
module EncodingGobTaintTracking {
  private class NewDecoder extends TaintTracking::FunctionModel {
    // signature: func NewDecoder(r io.Reader) *Decoder
    NewDecoder() { hasQualifiedName("encoding/gob", "NewDecoder") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class NewEncoder extends TaintTracking::FunctionModel {
    // signature: func NewEncoder(w io.Writer) *Encoder
    NewEncoder() { hasQualifiedName("encoding/gob", "NewEncoder") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isResult() and outp.isParameter(0))
    }
  }

  private class DecoderDecode extends TaintTracking::FunctionModel, Method {
    // signature: func (*Decoder).Decode(e interface{}) error
    DecoderDecode() { this.(Method).hasQualifiedName("encoding/gob", "Decoder", "Decode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class DecoderDecodeValue extends TaintTracking::FunctionModel, Method {
    // signature: func (*Decoder).DecodeValue(v reflect.Value) error
    DecoderDecodeValue() {
      this.(Method).hasQualifiedName("encoding/gob", "Decoder", "DecodeValue")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class EncoderEncode extends TaintTracking::FunctionModel, Method {
    // signature: func (*Encoder).Encode(e interface{}) error
    EncoderEncode() { this.(Method).hasQualifiedName("encoding/gob", "Encoder", "Encode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class EncoderEncodeValue extends TaintTracking::FunctionModel, Method {
    // signature: func (*Encoder).EncodeValue(value reflect.Value) error
    EncoderEncodeValue() {
      this.(Method).hasQualifiedName("encoding/gob", "Encoder", "EncodeValue")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class GobDecoderGobDecode extends TaintTracking::FunctionModel, Method {
    // signature: func (GobDecoder).GobDecode([]byte) error
    GobDecoderGobDecode() { this.implements("encoding/gob", "GobDecoder", "GobDecode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class GobEncoderGobEncode extends TaintTracking::FunctionModel, Method {
    // signature: func (GobEncoder).GobEncode() ([]byte, error)
    GobEncoderGobEncode() { this.implements("encoding/gob", "GobEncoder", "GobEncode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }
}
