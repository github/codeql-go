/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `sort` package. */
module SortTaintTracking {
  private class Reverse extends TaintTracking::FunctionModel {
    // signature: func Reverse(data Interface) Interface
    Reverse() { hasQualifiedName("sort", "Reverse") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }
}
