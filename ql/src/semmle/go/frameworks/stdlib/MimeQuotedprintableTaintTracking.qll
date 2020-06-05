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
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class NewWriter extends TaintTracking::FunctionModel {
    // signature: func NewWriter(w io.Writer) *Writer
    NewWriter() { hasQualifiedName("mime/quotedprintable", "NewWriter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult() and outp.isParameter(0)
    }
  }
}
