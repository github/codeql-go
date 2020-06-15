/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `container/heap` package. */
module ContainerHeapTaintTracking {
  private class Pop extends TaintTracking::FunctionModel {
    // signature: func Pop(h Interface) interface{}
    Pop() { hasQualifiedName("container/heap", "Pop") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class Push extends TaintTracking::FunctionModel {
    // signature: func Push(h Interface, x interface{})
    Push() { hasQualifiedName("container/heap", "Push") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isParameter(0))
    }
  }

  private class Remove extends TaintTracking::FunctionModel {
    // signature: func Remove(h Interface, i int) interface{}
    Remove() { hasQualifiedName("container/heap", "Remove") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class InterfacePop extends TaintTracking::FunctionModel, Method {
    // signature: func (Interface).Pop() interface{}
    InterfacePop() { this.implements("container/heap", "Interface", "Pop") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class InterfacePush extends TaintTracking::FunctionModel, Method {
    // signature: func (Interface).Push(x interface{})
    InterfacePush() { this.implements("container/heap", "Interface", "Push") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }
}
