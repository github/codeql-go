/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `compress/flate` package. */
module CompressFlateTaintTracking {
  private class NewReader extends TaintTracking::FunctionModel {
    // signature: func NewReader(r io.Reader) io.ReadCloser
    NewReader() { hasQualifiedName("compress/flate", "NewReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class NewReaderDict extends TaintTracking::FunctionModel {
    // signature: func NewReaderDict(r io.Reader, dict []byte) io.ReadCloser
    NewReaderDict() { hasQualifiedName("compress/flate", "NewReaderDict") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class NewWriter extends TaintTracking::FunctionModel {
    // signature: func NewWriter(w io.Writer, level int) (*Writer, error)
    NewWriter() { hasQualifiedName("compress/flate", "NewWriter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isResult(0) and outp.isParameter(0))
    }
  }

  private class NewWriterDict extends TaintTracking::FunctionModel {
    // signature: func NewWriterDict(w io.Writer, level int, dict []byte) (*Writer, error)
    NewWriterDict() { hasQualifiedName("compress/flate", "NewWriterDict") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isResult(0) and outp.isParameter(0))
    }
  }

  private class WriterReset extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).Reset(dst io.Writer)
    WriterReset() { this.(Method).hasQualifiedName("compress/flate", "Writer", "Reset") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class WriterWrite extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).Write(data []byte) (n int, err error)
    WriterWrite() { this.(Method).hasQualifiedName("compress/flate", "Writer", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class ResetterReset extends TaintTracking::FunctionModel, Method {
    // signature: func (Resetter).Reset(r io.Reader, dict []byte) error
    ResetterReset() { this.implements("compress/flate", "Resetter", "Reset") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }
}
