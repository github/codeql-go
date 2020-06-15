/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `container/ring` package. */
module ContainerRingTaintTracking {
  private class RingLink extends TaintTracking::FunctionModel, Method {
    // signature: func (*Ring).Link(s *Ring) *Ring
    RingLink() { this.(Method).hasQualifiedName("container/ring", "Ring", "Link") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RingMove extends TaintTracking::FunctionModel, Method {
    // signature: func (*Ring).Move(n int) *Ring
    RingMove() { this.(Method).hasQualifiedName("container/ring", "Ring", "Move") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class RingNext extends TaintTracking::FunctionModel, Method {
    // signature: func (*Ring).Next() *Ring
    RingNext() { this.(Method).hasQualifiedName("container/ring", "Ring", "Next") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class RingPrev extends TaintTracking::FunctionModel, Method {
    // signature: func (*Ring).Prev() *Ring
    RingPrev() { this.(Method).hasQualifiedName("container/ring", "Ring", "Prev") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class RingUnlink extends TaintTracking::FunctionModel, Method {
    // signature: func (*Ring).Unlink(n int) *Ring
    RingUnlink() { this.(Method).hasQualifiedName("container/ring", "Ring", "Unlink") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }
}
