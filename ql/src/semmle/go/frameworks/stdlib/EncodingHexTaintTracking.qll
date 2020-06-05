/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `encoding/hex` package. */
module EncodingHexTaintTracking {
  private class Decode extends TaintTracking::FunctionModel {
    // signature: func Decode(dst []byte, src []byte) (int, error)
    Decode() { hasQualifiedName("encoding/hex", "Decode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class DecodeString extends TaintTracking::FunctionModel {
    // signature: func DecodeString(s string) ([]byte, error)
    DecodeString() { hasQualifiedName("encoding/hex", "DecodeString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class Dump extends TaintTracking::FunctionModel {
    // signature: func Dump(data []byte) string
    Dump() { hasQualifiedName("encoding/hex", "Dump") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class Dumper extends TaintTracking::FunctionModel {
    // signature: func Dumper(w io.Writer) io.WriteCloser
    Dumper() { hasQualifiedName("encoding/hex", "Dumper") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult() and outp.isParameter(0)
    }
  }

  private class Encode extends TaintTracking::FunctionModel {
    // signature: func Encode(dst []byte, src []byte) int
    Encode() { hasQualifiedName("encoding/hex", "Encode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class EncodeToString extends TaintTracking::FunctionModel {
    // signature: func EncodeToString(src []byte) string
    EncodeToString() { hasQualifiedName("encoding/hex", "EncodeToString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class NewDecoder extends TaintTracking::FunctionModel {
    // signature: func NewDecoder(r io.Reader) io.Reader
    NewDecoder() { hasQualifiedName("encoding/hex", "NewDecoder") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class NewEncoder extends TaintTracking::FunctionModel {
    // signature: func NewEncoder(w io.Writer) io.Writer
    NewEncoder() { hasQualifiedName("encoding/hex", "NewEncoder") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult() and outp.isParameter(0)
    }
  }
}
