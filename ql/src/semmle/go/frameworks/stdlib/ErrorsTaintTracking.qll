/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `errors` package. */
module ErrorsTaintTracking {
  private class As extends TaintTracking::FunctionModel {
    // signature: func As(err error, target interface{}) bool
    As() { hasQualifiedName("errors", "As") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isParameter(1)
    }
  }

  private class New extends TaintTracking::FunctionModel {
    // signature: func New(text string) error
    New() { hasQualifiedName("errors", "New") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class Unwrap extends TaintTracking::FunctionModel {
    // signature: func Unwrap(err error) error
    Unwrap() { hasQualifiedName("errors", "Unwrap") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }
}
