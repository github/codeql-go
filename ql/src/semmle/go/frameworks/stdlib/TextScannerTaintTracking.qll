/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `text/scanner` package. */
module TextScannerTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func TokenString(tok rune) string
      hasQualifiedName("text/scanner", "TokenString") and
      (inp.isParameter(0) and outp.isResult())
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }

  private class MethodAndInterfaceTaintTracking extends TaintTracking::FunctionModel, Method {
    FunctionInput inp;
    FunctionOutput outp;

    MethodAndInterfaceTaintTracking() {
      // Methods:
      // signature: func (*Scanner).Init(src io.Reader) *Scanner
      this.(Method).hasQualifiedName("text/scanner", "Scanner", "Init") and
      (
        inp.isParameter(0) and
        (outp.isReceiver() or outp.isResult())
      )
      or
      // signature: func (*Scanner).Next() rune
      this.(Method).hasQualifiedName("text/scanner", "Scanner", "Next") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*Scanner).Peek() rune
      this.(Method).hasQualifiedName("text/scanner", "Scanner", "Peek") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*Scanner).Scan() rune
      this.(Method).hasQualifiedName("text/scanner", "Scanner", "Scan") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*Scanner).TokenText() string
      this.(Method).hasQualifiedName("text/scanner", "Scanner", "TokenText") and
      (inp.isReceiver() and outp.isResult())
      // Interfaces:
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
