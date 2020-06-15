/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `mime/quotedprintable` package. */
module MimeQuotedprintableTaintTracking {
  private class NewReader extends TaintTracking::FunctionModel {
    // signature: func NewReader(r io.Reader) *Reader
    NewReader() { hasQualifiedName("mime/quotedprintable", "NewReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class NewWriter extends TaintTracking::FunctionModel {
    // signature: func NewWriter(w io.Writer) *Writer
    NewWriter() { hasQualifiedName("mime/quotedprintable", "NewWriter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isResult() and outp.isParameter(0))
    }
  }

  private class ReaderRead extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).Read(p []byte) (n int, err error)
    ReaderRead() { this.(Method).hasQualifiedName("mime/quotedprintable", "Reader", "Read") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class WriterWrite extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).Write(p []byte) (n int, err error)
    WriterWrite() { this.(Method).hasQualifiedName("mime/quotedprintable", "Writer", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }
}
