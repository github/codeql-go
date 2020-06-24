/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `io/ioutil` package. */
module IoIoutilTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func NopCloser(r io.Reader) io.ReadCloser
      hasQualifiedName("io/ioutil", "NopCloser") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func ReadAll(r io.Reader) ([]byte, error)
      hasQualifiedName("io/ioutil", "ReadAll") and
      (inp.isParameter(0) and outp.isResult(0))
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
