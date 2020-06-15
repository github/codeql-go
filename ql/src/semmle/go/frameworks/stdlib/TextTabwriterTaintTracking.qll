/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `text/tabwriter` package. */
module TextTabwriterTaintTracking {
  private class NewWriter extends TaintTracking::FunctionModel {
    // signature: func NewWriter(output io.Writer, minwidth int, tabwidth int, padding int, padchar byte, flags uint) *Writer
    NewWriter() { hasQualifiedName("text/tabwriter", "NewWriter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isResult() and outp.isParameter(0))
    }
  }

  private class WriterInit extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).Init(output io.Writer, minwidth int, tabwidth int, padding int, padchar byte, flags uint) *Writer
    WriterInit() { this.(Method).hasQualifiedName("text/tabwriter", "Writer", "Init") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (
        (inp.isReceiver() or inp.isResult()) and
        outp.isParameter(0)
      )
    }
  }

  private class WriterWrite extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).Write(buf []byte) (n int, err error)
    WriterWrite() { this.(Method).hasQualifiedName("text/tabwriter", "Writer", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }
}
