/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `strconv` package. */
module StrconvTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func AppendQuote(dst []byte, s string) []byte
      hasQualifiedName("strconv", "AppendQuote") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func AppendQuoteRune(dst []byte, r rune) []byte
      hasQualifiedName("strconv", "AppendQuoteRune") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func AppendQuoteRuneToASCII(dst []byte, r rune) []byte
      hasQualifiedName("strconv", "AppendQuoteRuneToASCII") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte
      hasQualifiedName("strconv", "AppendQuoteRuneToGraphic") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func AppendQuoteToASCII(dst []byte, s string) []byte
      hasQualifiedName("strconv", "AppendQuoteToASCII") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func AppendQuoteToGraphic(dst []byte, s string) []byte
      hasQualifiedName("strconv", "AppendQuoteToGraphic") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func Quote(s string) string
      hasQualifiedName("strconv", "Quote") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func QuoteRune(r rune) string
      hasQualifiedName("strconv", "QuoteRune") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func QuoteRuneToASCII(r rune) string
      hasQualifiedName("strconv", "QuoteRuneToASCII") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func QuoteRuneToGraphic(r rune) string
      hasQualifiedName("strconv", "QuoteRuneToGraphic") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func QuoteToASCII(s string) string
      hasQualifiedName("strconv", "QuoteToASCII") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func QuoteToGraphic(s string) string
      hasQualifiedName("strconv", "QuoteToGraphic") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func Unquote(s string) (string, error)
      hasQualifiedName("strconv", "Unquote") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)
      hasQualifiedName("strconv", "UnquoteChar") and
      (inp.isParameter(0) and outp.isResult(0))
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
