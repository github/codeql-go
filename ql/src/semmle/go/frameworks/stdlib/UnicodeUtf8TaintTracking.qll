/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `unicode/utf8` package. */
module UnicodeUtf8TaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func DecodeLastRune(p []byte) (r rune, size int)
      hasQualifiedName("unicode/utf8", "DecodeLastRune") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func DecodeLastRuneInString(s string) (r rune, size int)
      hasQualifiedName("unicode/utf8", "DecodeLastRuneInString") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func DecodeRune(p []byte) (r rune, size int)
      hasQualifiedName("unicode/utf8", "DecodeRune") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func DecodeRuneInString(s string) (r rune, size int)
      hasQualifiedName("unicode/utf8", "DecodeRuneInString") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func EncodeRune(p []byte, r rune) int
      hasQualifiedName("unicode/utf8", "EncodeRune") and
      (inp.isParameter(1) and outp.isParameter(0))
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
