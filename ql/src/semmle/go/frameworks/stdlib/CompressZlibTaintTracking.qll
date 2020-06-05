/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `compress/zlib` package. */
module CompressZlibTaintTracking {
  private class NewReader extends TaintTracking::FunctionModel {
    // signature: func NewReader(r io.Reader) (io.ReadCloser, error)
    NewReader() { hasQualifiedName("compress/zlib", "NewReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class NewReaderDict extends TaintTracking::FunctionModel {
    // signature: func NewReaderDict(r io.Reader, dict []byte) (io.ReadCloser, error)
    NewReaderDict() { hasQualifiedName("compress/zlib", "NewReaderDict") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class NewWriter extends TaintTracking::FunctionModel {
    // signature: func NewWriter(w io.Writer) *Writer
    NewWriter() { hasQualifiedName("compress/zlib", "NewWriter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult() and outp.isParameter(0)
    }
  }

  private class NewWriterLevel extends TaintTracking::FunctionModel {
    // signature: func NewWriterLevel(w io.Writer, level int) (*Writer, error)
    NewWriterLevel() { hasQualifiedName("compress/zlib", "NewWriterLevel") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult(0) and outp.isParameter(0)
    }
  }

  private class NewWriterLevelDict extends TaintTracking::FunctionModel {
    // signature: func NewWriterLevelDict(w io.Writer, level int, dict []byte) (*Writer, error)
    NewWriterLevelDict() { hasQualifiedName("compress/zlib", "NewWriterLevelDict") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult(0) and outp.isParameter(0)
    }
  }

  private class WriterReset extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).Reset(w io.Writer)
    WriterReset() { this.(Method).hasQualifiedName("compress/zlib", "Writer", "Reset") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isParameter(0)
    }
  }

  private class WriterWrite extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).Write(p []byte) (n int, err error)
    WriterWrite() { this.(Method).hasQualifiedName("compress/zlib", "Writer", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ResetterReset extends TaintTracking::FunctionModel, Method {
    // signature: func (Resetter).Reset(r io.Reader, dict []byte) error
    ResetterReset() { this.implements("compress/zlib", "Resetter", "Reset") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }
}
