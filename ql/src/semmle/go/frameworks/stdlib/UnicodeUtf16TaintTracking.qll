/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `unicode/utf16` package. */
module UnicodeUtf16TaintTracking {
  private class Decode extends TaintTracking::FunctionModel {
    // signature: func Decode(s []uint16) []rune
    Decode() { hasQualifiedName("unicode/utf16", "Decode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class DecodeRune extends TaintTracking::FunctionModel {
    // signature: func DecodeRune(r1 rune, r2 rune) rune
    DecodeRune() { hasQualifiedName("unicode/utf16", "DecodeRune") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isResult()
    }
  }

  private class Encode extends TaintTracking::FunctionModel {
    // signature: func Encode(s []rune) []uint16
    Encode() { hasQualifiedName("unicode/utf16", "Encode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class EncodeRune extends TaintTracking::FunctionModel {
    // signature: func EncodeRune(r rune) (r1 rune, r2 rune)
    EncodeRune() { hasQualifiedName("unicode/utf16", "EncodeRune") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(_)
    }
  }
}
