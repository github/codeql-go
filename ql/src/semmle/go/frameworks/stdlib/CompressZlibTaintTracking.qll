/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `compress/zlib` package. */
module CompressZlibTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func NewReader(r io.Reader) (io.ReadCloser, error)
      hasQualifiedName("compress/zlib", "NewReader") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func NewReaderDict(r io.Reader, dict []byte) (io.ReadCloser, error)
      hasQualifiedName("compress/zlib", "NewReaderDict") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func NewWriter(w io.Writer) *Writer
      hasQualifiedName("compress/zlib", "NewWriter") and
      (inp.isResult() and outp.isParameter(0))
      or
      // signature: func NewWriterLevel(w io.Writer, level int) (*Writer, error)
      hasQualifiedName("compress/zlib", "NewWriterLevel") and
      (inp.isResult(0) and outp.isParameter(0))
      or
      // signature: func NewWriterLevelDict(w io.Writer, level int, dict []byte) (*Writer, error)
      hasQualifiedName("compress/zlib", "NewWriterLevelDict") and
      (inp.isResult(0) and outp.isParameter(0))
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
      // signature: func (*Writer).Reset(w io.Writer)
      this.(Method).hasQualifiedName("compress/zlib", "Writer", "Reset") and
      (inp.isReceiver() and outp.isParameter(0))
      or
      // signature: func (*Writer).Write(p []byte) (n int, err error)
      this.(Method).hasQualifiedName("compress/zlib", "Writer", "Write") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // Interfaces:
      // signature: func (Resetter).Reset(r io.Reader, dict []byte) error
      this.implements("compress/zlib", "Resetter", "Reset") and
      (inp.isParameter(0) and outp.isReceiver())
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
