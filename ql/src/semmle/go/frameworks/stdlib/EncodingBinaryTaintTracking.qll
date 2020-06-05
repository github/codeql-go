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
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class PutVarint extends TaintTracking::FunctionModel {
    // signature: func PutVarint(buf []byte, x int64) int
    PutVarint() { hasQualifiedName("encoding/binary", "PutVarint") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class EncodingBinaryRead extends TaintTracking::FunctionModel {
    // signature: func Read(r io.Reader, order ByteOrder, data interface{}) error
    EncodingBinaryRead() { hasQualifiedName("encoding/binary", "Read") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isParameter(2)
    }
  }

  private class ReadUvarint extends TaintTracking::FunctionModel {
    // signature: func ReadUvarint(r io.ByteReader) (uint64, error)
    ReadUvarint() { hasQualifiedName("encoding/binary", "ReadUvarint") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class ReadVarint extends TaintTracking::FunctionModel {
    // signature: func ReadVarint(r io.ByteReader) (int64, error)
    ReadVarint() { hasQualifiedName("encoding/binary", "ReadVarint") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class Uvarint extends TaintTracking::FunctionModel {
    // signature: func Uvarint(buf []byte) (uint64, int)
    Uvarint() { hasQualifiedName("encoding/binary", "Uvarint") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class Varint extends TaintTracking::FunctionModel {
    // signature: func Varint(buf []byte) (int64, int)
    Varint() { hasQualifiedName("encoding/binary", "Varint") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class EncodingBinaryWrite extends TaintTracking::FunctionModel {
    // signature: func Write(w io.Writer, order ByteOrder, data interface{}) error
    EncodingBinaryWrite() { hasQualifiedName("encoding/binary", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(2) and outp.isParameter(0)
    }
  }

  private class ByteOrderPutUint16 extends TaintTracking::FunctionModel, Method {
    // signature: func (ByteOrder).PutUint16([]byte, uint16)
    ByteOrderPutUint16() { this.implements("encoding/binary", "ByteOrder", "PutUint16") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class ByteOrderPutUint32 extends TaintTracking::FunctionModel, Method {
    // signature: func (ByteOrder).PutUint32([]byte, uint32)
    ByteOrderPutUint32() { this.implements("encoding/binary", "ByteOrder", "PutUint32") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class ByteOrderPutUint64 extends TaintTracking::FunctionModel, Method {
    // signature: func (ByteOrder).PutUint64([]byte, uint64)
    ByteOrderPutUint64() { this.implements("encoding/binary", "ByteOrder", "PutUint64") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class ByteOrderString extends TaintTracking::FunctionModel, Method {
    // signature: func (ByteOrder).String() string
    ByteOrderString() { this.implements("encoding/binary", "ByteOrder", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ByteOrderUint16 extends TaintTracking::FunctionModel, Method {
    // signature: func (ByteOrder).Uint16([]byte) uint16
    ByteOrderUint16() { this.implements("encoding/binary", "ByteOrder", "Uint16") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class ByteOrderUint32 extends TaintTracking::FunctionModel, Method {
    // signature: func (ByteOrder).Uint32([]byte) uint32
    ByteOrderUint32() { this.implements("encoding/binary", "ByteOrder", "Uint32") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class ByteOrderUint64 extends TaintTracking::FunctionModel, Method {
    // signature: func (ByteOrder).Uint64([]byte) uint64
    ByteOrderUint64() { this.implements("encoding/binary", "ByteOrder", "Uint64") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }
}
