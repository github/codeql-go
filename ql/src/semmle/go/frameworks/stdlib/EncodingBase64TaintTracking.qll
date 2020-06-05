/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `encoding/base64` package. */
module EncodingBase64TaintTracking {
  private class NewDecoder extends TaintTracking::FunctionModel {
    // signature: func NewDecoder(enc *Encoding, r io.Reader) io.Reader
    NewDecoder() { hasQualifiedName("encoding/base64", "NewDecoder") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isResult()
    }
  }

  private class NewEncoder extends TaintTracking::FunctionModel {
    // signature: func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser
    NewEncoder() { hasQualifiedName("encoding/base64", "NewEncoder") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult() and outp.isParameter(1)
    }
  }

  private class EncodingDecode extends TaintTracking::FunctionModel, Method {
    // signature: func (*Encoding).Decode(dst []byte, src []byte) (n int, err error)
    EncodingDecode() { this.(Method).hasQualifiedName("encoding/base64", "Encoding", "Decode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class EncodingDecodeString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Encoding).DecodeString(s string) ([]byte, error)
    EncodingDecodeString() {
      this.(Method).hasQualifiedName("encoding/base64", "Encoding", "DecodeString")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class EncodingEncode extends TaintTracking::FunctionModel, Method {
    // signature: func (*Encoding).Encode(dst []byte, src []byte)
    EncodingEncode() { this.(Method).hasQualifiedName("encoding/base64", "Encoding", "Encode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class EncodingEncodeToString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Encoding).EncodeToString(src []byte) string
    EncodingEncodeToString() {
      this.(Method).hasQualifiedName("encoding/base64", "Encoding", "EncodeToString")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }
}
