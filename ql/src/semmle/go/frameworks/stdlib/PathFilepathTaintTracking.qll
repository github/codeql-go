/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `path/filepath` package. */
module PathFilepathTaintTracking {
  private class Abs extends TaintTracking::FunctionModel {
    // signature: func Abs(path string) (string, error)
    Abs() { hasQualifiedName("path/filepath", "Abs") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class Base extends TaintTracking::FunctionModel {
    // signature: func Base(path string) string
    Base() { hasQualifiedName("path/filepath", "Base") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class Clean extends TaintTracking::FunctionModel {
    // signature: func Clean(path string) string
    Clean() { hasQualifiedName("path/filepath", "Clean") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class Dir extends TaintTracking::FunctionModel {
    // signature: func Dir(path string) string
    Dir() { hasQualifiedName("path/filepath", "Dir") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class EvalSymlinks extends TaintTracking::FunctionModel {
    // signature: func EvalSymlinks(path string) (string, error)
    EvalSymlinks() { hasQualifiedName("path/filepath", "EvalSymlinks") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class Ext extends TaintTracking::FunctionModel {
    // signature: func Ext(path string) string
    Ext() { hasQualifiedName("path/filepath", "Ext") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class FromSlash extends TaintTracking::FunctionModel {
    // signature: func FromSlash(path string) string
    FromSlash() { hasQualifiedName("path/filepath", "FromSlash") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class Join extends TaintTracking::FunctionModel {
    // signature: func Join(elem ...string) string
    Join() { hasQualifiedName("path/filepath", "Join") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class Rel extends TaintTracking::FunctionModel {
    // signature: func Rel(basepath string, targpath string) (string, error)
    Rel() { hasQualifiedName("path/filepath", "Rel") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult(0))
    }
  }

  private class Split extends TaintTracking::FunctionModel {
    // signature: func Split(path string) (dir string, file string)
    Split() { hasQualifiedName("path/filepath", "Split") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(_))
    }
  }

  private class SplitList extends TaintTracking::FunctionModel {
    // signature: func SplitList(path string) []string
    SplitList() { hasQualifiedName("path/filepath", "SplitList") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class ToSlash extends TaintTracking::FunctionModel {
    // signature: func ToSlash(path string) string
    ToSlash() { hasQualifiedName("path/filepath", "ToSlash") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }
}
