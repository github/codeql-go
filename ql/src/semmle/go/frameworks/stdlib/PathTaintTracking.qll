/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `path` package. */
module PathTaintTracking {
  private class Base extends TaintTracking::FunctionModel {
    // signature: func Base(path string) string
    Base() { hasQualifiedName("path", "Base") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class Clean extends TaintTracking::FunctionModel {
    // signature: func Clean(path string) string
    Clean() { hasQualifiedName("path", "Clean") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class Dir extends TaintTracking::FunctionModel {
    // signature: func Dir(path string) string
    Dir() { hasQualifiedName("path", "Dir") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class Ext extends TaintTracking::FunctionModel {
    // signature: func Ext(path string) string
    Ext() { hasQualifiedName("path", "Ext") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class Join extends TaintTracking::FunctionModel {
    // signature: func Join(elem ...string) string
    Join() { hasQualifiedName("path", "Join") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isResult()
    }
  }

  private class Split extends TaintTracking::FunctionModel {
    // signature: func Split(path string) (dir string, file string)
    Split() { hasQualifiedName("path", "Split") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(_)
    }
  }
}
