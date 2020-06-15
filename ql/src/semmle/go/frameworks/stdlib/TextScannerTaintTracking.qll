/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `text/scanner` package. */
module TextScannerTaintTracking {
  private class TokenString extends TaintTracking::FunctionModel {
    // signature: func TokenString(tok rune) string
    TokenString() { hasQualifiedName("text/scanner", "TokenString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class ScannerInit extends TaintTracking::FunctionModel, Method {
    // signature: func (*Scanner).Init(src io.Reader) *Scanner
    ScannerInit() { this.(Method).hasQualifiedName("text/scanner", "Scanner", "Init") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (
        inp.isParameter(0) and
        (outp.isReceiver() or outp.isResult())
      )
    }
  }

  private class ScannerNext extends TaintTracking::FunctionModel, Method {
    // signature: func (*Scanner).Next() rune
    ScannerNext() { this.(Method).hasQualifiedName("text/scanner", "Scanner", "Next") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class ScannerPeek extends TaintTracking::FunctionModel, Method {
    // signature: func (*Scanner).Peek() rune
    ScannerPeek() { this.(Method).hasQualifiedName("text/scanner", "Scanner", "Peek") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class ScannerScan extends TaintTracking::FunctionModel, Method {
    // signature: func (*Scanner).Scan() rune
    ScannerScan() { this.(Method).hasQualifiedName("text/scanner", "Scanner", "Scan") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class ScannerTokenText extends TaintTracking::FunctionModel, Method {
    // signature: func (*Scanner).TokenText() string
    ScannerTokenText() { this.(Method).hasQualifiedName("text/scanner", "Scanner", "TokenText") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }
}
