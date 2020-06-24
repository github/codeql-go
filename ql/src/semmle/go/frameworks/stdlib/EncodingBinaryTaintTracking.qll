/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `encoding/binary` package. */
module EncodingBinaryTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func PutUvarint(buf []byte, x uint64) int
      hasQualifiedName("encoding/binary", "PutUvarint") and
      (inp.isParameter(1) and outp.isParameter(0))
      or
      // signature: func Read(r io.Reader, order ByteOrder, data interface{}) error
      hasQualifiedName("encoding/binary", "Read") and
      (inp.isParameter(0) and outp.isParameter(2))
      or
      // signature: func ReadUvarint(r io.ByteReader) (uint64, error)
      hasQualifiedName("encoding/binary", "ReadUvarint") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func Write(w io.Writer, order ByteOrder, data interface{}) error
      hasQualifiedName("encoding/binary", "Write") and
      (inp.isParameter(2) and outp.isParameter(0))
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
      // Interfaces:
      // signature: func (ByteOrder).String() string
      this.implements("encoding/binary", "ByteOrder", "String") and
      (inp.isReceiver() and outp.isResult())
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
