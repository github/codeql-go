/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `unicode` package. */
module UnicodeTaintTracking {
  private class SimpleFold extends TaintTracking::FunctionModel {
    // signature: func SimpleFold(r rune) rune
    SimpleFold() { hasQualifiedName("unicode", "SimpleFold") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class To extends TaintTracking::FunctionModel {
    // signature: func To(_case int, r rune) rune
    To() { hasQualifiedName("unicode", "To") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isResult())
    }
  }

  private class ToLower extends TaintTracking::FunctionModel {
    // signature: func ToLower(r rune) rune
    ToLower() { hasQualifiedName("unicode", "ToLower") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class ToTitle extends TaintTracking::FunctionModel {
    // signature: func ToTitle(r rune) rune
    ToTitle() { hasQualifiedName("unicode", "ToTitle") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class ToUpper extends TaintTracking::FunctionModel {
    // signature: func ToUpper(r rune) rune
    ToUpper() { hasQualifiedName("unicode", "ToUpper") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class SpecialCaseToLower extends TaintTracking::FunctionModel, Method {
    // signature: func (SpecialCase).ToLower(r rune) rune
    SpecialCaseToLower() { this.(Method).hasQualifiedName("unicode", "SpecialCase", "ToLower") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class SpecialCaseToTitle extends TaintTracking::FunctionModel, Method {
    // signature: func (SpecialCase).ToTitle(r rune) rune
    SpecialCaseToTitle() { this.(Method).hasQualifiedName("unicode", "SpecialCase", "ToTitle") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class SpecialCaseToUpper extends TaintTracking::FunctionModel, Method {
    // signature: func (SpecialCase).ToUpper(r rune) rune
    SpecialCaseToUpper() { this.(Method).hasQualifiedName("unicode", "SpecialCase", "ToUpper") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }
}
