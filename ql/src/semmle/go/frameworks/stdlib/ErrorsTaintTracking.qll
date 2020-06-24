/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `errors` package. */
module ErrorsTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func As(err error, target interface{}) bool
      hasQualifiedName("errors", "As") and
      (inp.isParameter(0) and outp.isParameter(1))
      or
      // signature: func New(text string) error
      hasQualifiedName("errors", "New") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func Unwrap(err error) error
      hasQualifiedName("errors", "Unwrap") and
      (inp.isParameter(0) and outp.isResult())
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
