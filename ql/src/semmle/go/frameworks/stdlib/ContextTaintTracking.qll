/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `context` package. */
module ContextTaintTracking {
  private class WithCancel extends TaintTracking::FunctionModel {
    // signature: func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
    WithCancel() { hasQualifiedName("context", "WithCancel") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class WithDeadline extends TaintTracking::FunctionModel {
    // signature: func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
    WithDeadline() { hasQualifiedName("context", "WithDeadline") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class WithTimeout extends TaintTracking::FunctionModel {
    // signature: func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
    WithTimeout() { hasQualifiedName("context", "WithTimeout") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class WithValue extends TaintTracking::FunctionModel {
    // signature: func WithValue(parent Context, key interface{}, val interface{}) Context
    WithValue() { hasQualifiedName("context", "WithValue") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class ContextValue extends TaintTracking::FunctionModel, Method {
    // signature: func (Context).Value(key interface{}) interface{}
    ContextValue() { this.implements("context", "Context", "Value") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }
}
