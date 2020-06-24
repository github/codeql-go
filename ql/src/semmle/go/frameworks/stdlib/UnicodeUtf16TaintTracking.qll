/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `unicode/utf16` package. */
module UnicodeUtf16TaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func Decode(s []uint16) []rune
      hasQualifiedName("unicode/utf16", "Decode") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func DecodeRune(r1 rune, r2 rune) rune
      hasQualifiedName("unicode/utf16", "DecodeRune") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func Encode(s []rune) []uint16
      hasQualifiedName("unicode/utf16", "Encode") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func EncodeRune(r rune) (r1 rune, r2 rune)
      hasQualifiedName("unicode/utf16", "EncodeRune") and
      (inp.isParameter(0) and outp.isResult(_))
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
