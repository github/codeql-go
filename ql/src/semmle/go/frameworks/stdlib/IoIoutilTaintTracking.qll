/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `io/ioutil` package. */
module IoIoutilTaintTracking {
  private class NopCloser extends TaintTracking::FunctionModel {
    // signature: func NopCloser(r io.Reader) io.ReadCloser
    NopCloser() { hasQualifiedName("io/ioutil", "NopCloser") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class ReadAll extends TaintTracking::FunctionModel {
    // signature: func ReadAll(r io.Reader) ([]byte, error)
    ReadAll() { hasQualifiedName("io/ioutil", "ReadAll") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class WriterWrite extends TaintTracking::FunctionModel, Method {
    // signature: func (Writer).Write(p []byte) (n int, err error)
    WriterWrite() { this.implements("io", "Writer", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }
}
