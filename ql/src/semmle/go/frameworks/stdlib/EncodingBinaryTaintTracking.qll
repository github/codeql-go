/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `encoding/binary` package. */
module EncodingBinaryTaintTracking {
  private class PutUvarint extends TaintTracking::FunctionModel {
    // signature: func PutUvarint(buf []byte, x uint64) int
    PutUvarint() { hasQualifiedName("encoding/binary", "PutUvarint") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isParameter(0))
    }
  }

  private class EncodingBinaryRead extends TaintTracking::FunctionModel {
    // signature: func Read(r io.Reader, order ByteOrder, data interface{}) error
    EncodingBinaryRead() { hasQualifiedName("encoding/binary", "Read") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isParameter(2))
    }
  }

  private class ReadUvarint extends TaintTracking::FunctionModel {
    // signature: func ReadUvarint(r io.ByteReader) (uint64, error)
    ReadUvarint() { hasQualifiedName("encoding/binary", "ReadUvarint") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class EncodingBinaryWrite extends TaintTracking::FunctionModel {
    // signature: func Write(w io.Writer, order ByteOrder, data interface{}) error
    EncodingBinaryWrite() { hasQualifiedName("encoding/binary", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(2) and outp.isParameter(0))
    }
  }

  private class ByteOrderString extends TaintTracking::FunctionModel, Method {
    // signature: func (ByteOrder).String() string
    ByteOrderString() { this.implements("encoding/binary", "ByteOrder", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }
}
