/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `compress/bzip2` package. */
module CompressBzip2TaintTracking {
  private class NewReader extends TaintTracking::FunctionModel {
    // signature: func NewReader(r io.Reader) io.Reader
    NewReader() { hasQualifiedName("compress/bzip2", "NewReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }
}
