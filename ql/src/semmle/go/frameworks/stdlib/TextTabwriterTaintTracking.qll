/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `text/tabwriter` package. */
module TextTabwriterTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func NewWriter(output io.Writer, minwidth int, tabwidth int, padding int, padchar byte, flags uint) *Writer
      hasQualifiedName("text/tabwriter", "NewWriter") and
      (inp.isResult() and outp.isParameter(0))
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
      // signature: func (*Writer).Init(output io.Writer, minwidth int, tabwidth int, padding int, padchar byte, flags uint) *Writer
      this.(Method).hasQualifiedName("text/tabwriter", "Writer", "Init") and
      (
        (inp.isReceiver() or inp.isResult()) and
        outp.isParameter(0)
      )
      or
      // signature: func (*Writer).Write(buf []byte) (n int, err error)
      this.(Method).hasQualifiedName("text/tabwriter", "Writer", "Write") and
      (inp.isParameter(0) and outp.isReceiver())
      // Interfaces:
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
