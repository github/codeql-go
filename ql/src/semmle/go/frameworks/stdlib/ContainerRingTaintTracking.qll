/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `container/ring` package. */
module ContainerRingTaintTracking {
  private class MethodAndInterfaceTaintTracking extends TaintTracking::FunctionModel, Method {
    FunctionInput inp;
    FunctionOutput outp;

    MethodAndInterfaceTaintTracking() {
      // Methods:
      // signature: func (*Ring).Link(s *Ring) *Ring
      this.(Method).hasQualifiedName("container/ring", "Ring", "Link") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Ring).Move(n int) *Ring
      this.(Method).hasQualifiedName("container/ring", "Ring", "Move") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*Ring).Next() *Ring
      this.(Method).hasQualifiedName("container/ring", "Ring", "Next") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*Ring).Prev() *Ring
      this.(Method).hasQualifiedName("container/ring", "Ring", "Prev") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*Ring).Unlink(n int) *Ring
      this.(Method).hasQualifiedName("container/ring", "Ring", "Unlink") and
      (inp.isReceiver() and outp.isResult())
      // Interfaces:
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
