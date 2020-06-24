/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `compress/gzip` package. */
module CompressGzipTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func NewReader(r io.Reader) (*Reader, error)
      hasQualifiedName("compress/gzip", "NewReader") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func NewWriter(w io.Writer) *Writer
      hasQualifiedName("compress/gzip", "NewWriter") and
      (inp.isResult() and outp.isParameter(0))
      or
      // signature: func NewWriterLevel(w io.Writer, level int) (*Writer, error)
      hasQualifiedName("compress/gzip", "NewWriterLevel") and
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
      // signature: func (*Reader).Read(p []byte) (n int, err error)
      this.(Method).hasQualifiedName("compress/gzip", "Reader", "Read") and
      (inp.isReceiver() and outp.isParameter(0))
      or
      // signature: func (*Reader).Reset(r io.Reader) error
      this.(Method).hasQualifiedName("compress/gzip", "Reader", "Reset") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (*Writer).Reset(w io.Writer)
      this.(Method).hasQualifiedName("compress/gzip", "Writer", "Reset") and
      (inp.isReceiver() and outp.isParameter(0))
      or
      // signature: func (*Writer).Write(p []byte) (int, error)
      this.(Method).hasQualifiedName("compress/gzip", "Writer", "Write") and
      (inp.isParameter(0) and outp.isReceiver())
      // Interfaces:
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
