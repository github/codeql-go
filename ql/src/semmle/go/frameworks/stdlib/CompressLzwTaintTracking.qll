/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `compress/lzw` package. */
module CompressLzwTaintTracking {
  private class NewReader extends TaintTracking::FunctionModel {
    // signature: func NewReader(r io.Reader, order Order, litWidth int) io.ReadCloser
    NewReader() { hasQualifiedName("compress/lzw", "NewReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class NewWriter extends TaintTracking::FunctionModel {
    // signature: func NewWriter(w io.Writer, order Order, litWidth int) io.WriteCloser
    NewWriter() { hasQualifiedName("compress/lzw", "NewWriter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult() and outp.isParameter(0)
    }
  }
}
