/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `unicode/utf8` package. */
module UnicodeUtf8TaintTracking {
  private class DecodeLastRune extends TaintTracking::FunctionModel {
    // signature: func DecodeLastRune(p []byte) (r rune, size int)
    DecodeLastRune() { hasQualifiedName("unicode/utf8", "DecodeLastRune") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class DecodeLastRuneInString extends TaintTracking::FunctionModel {
    // signature: func DecodeLastRuneInString(s string) (r rune, size int)
    DecodeLastRuneInString() { hasQualifiedName("unicode/utf8", "DecodeLastRuneInString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class DecodeRune extends TaintTracking::FunctionModel {
    // signature: func DecodeRune(p []byte) (r rune, size int)
    DecodeRune() { hasQualifiedName("unicode/utf8", "DecodeRune") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class DecodeRuneInString extends TaintTracking::FunctionModel {
    // signature: func DecodeRuneInString(s string) (r rune, size int)
    DecodeRuneInString() { hasQualifiedName("unicode/utf8", "DecodeRuneInString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class EncodeRune extends TaintTracking::FunctionModel {
    // signature: func EncodeRune(p []byte, r rune) int
    EncodeRune() { hasQualifiedName("unicode/utf8", "EncodeRune") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }
}
