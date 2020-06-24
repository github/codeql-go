/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `html` package. */
module HtmlTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func EscapeString(s string) string
      hasQualifiedName("html", "EscapeString") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func UnescapeString(s string) string
      hasQualifiedName("html", "UnescapeString") and
      (inp.isParameter(0) and outp.isResult())
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
