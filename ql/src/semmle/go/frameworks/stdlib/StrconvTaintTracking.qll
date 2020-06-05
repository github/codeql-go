/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `strconv` package. */
module StrconvTaintTracking {
  private class AppendQuote extends TaintTracking::FunctionModel {
    // signature: func AppendQuote(dst []byte, s string) []byte
    AppendQuote() { hasQualifiedName("strconv", "AppendQuote") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isResult()
      or
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class AppendQuoteRune extends TaintTracking::FunctionModel {
    // signature: func AppendQuoteRune(dst []byte, r rune) []byte
    AppendQuoteRune() { hasQualifiedName("strconv", "AppendQuoteRune") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isResult()
      or
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class AppendQuoteRuneToASCII extends TaintTracking::FunctionModel {
    // signature: func AppendQuoteRuneToASCII(dst []byte, r rune) []byte
    AppendQuoteRuneToASCII() { hasQualifiedName("strconv", "AppendQuoteRuneToASCII") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isResult()
      or
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class AppendQuoteRuneToGraphic extends TaintTracking::FunctionModel {
    // signature: func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte
    AppendQuoteRuneToGraphic() { hasQualifiedName("strconv", "AppendQuoteRuneToGraphic") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isResult()
      or
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class AppendQuoteToASCII extends TaintTracking::FunctionModel {
    // signature: func AppendQuoteToASCII(dst []byte, s string) []byte
    AppendQuoteToASCII() { hasQualifiedName("strconv", "AppendQuoteToASCII") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isResult()
      or
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class AppendQuoteToGraphic extends TaintTracking::FunctionModel {
    // signature: func AppendQuoteToGraphic(dst []byte, s string) []byte
    AppendQuoteToGraphic() { hasQualifiedName("strconv", "AppendQuoteToGraphic") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isResult()
      or
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class Quote extends TaintTracking::FunctionModel {
    // signature: func Quote(s string) string
    Quote() { hasQualifiedName("strconv", "Quote") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class QuoteRune extends TaintTracking::FunctionModel {
    // signature: func QuoteRune(r rune) string
    QuoteRune() { hasQualifiedName("strconv", "QuoteRune") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class QuoteRuneToASCII extends TaintTracking::FunctionModel {
    // signature: func QuoteRuneToASCII(r rune) string
    QuoteRuneToASCII() { hasQualifiedName("strconv", "QuoteRuneToASCII") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class QuoteRuneToGraphic extends TaintTracking::FunctionModel {
    // signature: func QuoteRuneToGraphic(r rune) string
    QuoteRuneToGraphic() { hasQualifiedName("strconv", "QuoteRuneToGraphic") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class QuoteToASCII extends TaintTracking::FunctionModel {
    // signature: func QuoteToASCII(s string) string
    QuoteToASCII() { hasQualifiedName("strconv", "QuoteToASCII") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class QuoteToGraphic extends TaintTracking::FunctionModel {
    // signature: func QuoteToGraphic(s string) string
    QuoteToGraphic() { hasQualifiedName("strconv", "QuoteToGraphic") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class Unquote extends TaintTracking::FunctionModel {
    // signature: func Unquote(s string) (string, error)
    Unquote() { hasQualifiedName("strconv", "Unquote") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class UnquoteChar extends TaintTracking::FunctionModel {
    // signature: func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)
    UnquoteChar() { hasQualifiedName("strconv", "UnquoteChar") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult([0, 2])
    }
  }
}
