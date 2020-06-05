/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `compress/gzip` package. */
module CompressGzipTaintTracking {
  private class NewReader extends TaintTracking::FunctionModel {
    // signature: func NewReader(r io.Reader) (*Reader, error)
    NewReader() { hasQualifiedName("compress/gzip", "NewReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class NewWriter extends TaintTracking::FunctionModel {
    // signature: func NewWriter(w io.Writer) *Writer
    NewWriter() { hasQualifiedName("compress/gzip", "NewWriter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult() and outp.isParameter(0)
    }
  }

  private class NewWriterLevel extends TaintTracking::FunctionModel {
    // signature: func NewWriterLevel(w io.Writer, level int) (*Writer, error)
    NewWriterLevel() { hasQualifiedName("compress/gzip", "NewWriterLevel") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult(0) and outp.isParameter(0)
    }
  }

  private class ReaderReset extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).Reset(r io.Reader) error
    ReaderReset() { this.(Method).hasQualifiedName("compress/gzip", "Reader", "Reset") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class WriterReset extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).Reset(w io.Writer)
    WriterReset() { this.(Method).hasQualifiedName("compress/gzip", "Writer", "Reset") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isParameter(0)
    }
  }
}
