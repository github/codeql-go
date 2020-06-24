/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `encoding/base64` package. */
module EncodingBase64TaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func NewDecoder(enc *Encoding, r io.Reader) io.Reader
      hasQualifiedName("encoding/base64", "NewDecoder") and
      (inp.isParameter(1) and outp.isResult())
      or
      // signature: func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser
      hasQualifiedName("encoding/base64", "NewEncoder") and
      (inp.isResult() and outp.isParameter(1))
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }

  private class MethodAndInterfaceTaintTracking extends TaintTracking::FunctionModel, Method {
    FunctionInput inp;
    FunctionOutput outp;

    MethodAndInterfaceTaintTracking() {
      // Methods:
      // signature: func (*Encoding).Decode(dst []byte, src []byte) (n int, err error)
      this.(Method).hasQualifiedName("encoding/base64", "Encoding", "Decode") and
      (inp.isParameter(1) and outp.isParameter(0))
      or
      // signature: func (*Encoding).DecodeString(s string) ([]byte, error)
      this.(Method).hasQualifiedName("encoding/base64", "Encoding", "DecodeString") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func (*Encoding).Encode(dst []byte, src []byte)
      this.(Method).hasQualifiedName("encoding/base64", "Encoding", "Encode") and
      (inp.isParameter(1) and outp.isParameter(0))
      or
      // signature: func (*Encoding).EncodeToString(src []byte) string
      this.(Method).hasQualifiedName("encoding/base64", "Encoding", "EncodeToString") and
      (inp.isParameter(0) and outp.isResult())
      // Interfaces:
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
