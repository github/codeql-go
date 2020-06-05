/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `encoding/ascii85` package. */
module EncodingAscii85TaintTracking {
  private class Decode extends TaintTracking::FunctionModel {
    // signature: func Decode(dst []byte, src []byte, flush bool) (ndst int, nsrc int, err error)
    Decode() { hasQualifiedName("encoding/ascii85", "Decode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class Encode extends TaintTracking::FunctionModel {
    // signature: func Encode(dst []byte, src []byte) int
    Encode() { hasQualifiedName("encoding/ascii85", "Encode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class NewDecoder extends TaintTracking::FunctionModel {
    // signature: func NewDecoder(r io.Reader) io.Reader
    NewDecoder() { hasQualifiedName("encoding/ascii85", "NewDecoder") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class NewEncoder extends TaintTracking::FunctionModel {
    // signature: func NewEncoder(w io.Writer) io.WriteCloser
    NewEncoder() { hasQualifiedName("encoding/ascii85", "NewEncoder") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult() and outp.isParameter(0)
    }
  }
}
