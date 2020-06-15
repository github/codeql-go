/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `html` package. */
module HtmlTaintTracking {
  private class EscapeString extends TaintTracking::FunctionModel {
    // signature: func EscapeString(s string) string
    EscapeString() { hasQualifiedName("html", "EscapeString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class UnescapeString extends TaintTracking::FunctionModel {
    // signature: func UnescapeString(s string) string
    UnescapeString() { hasQualifiedName("html", "UnescapeString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }
}
