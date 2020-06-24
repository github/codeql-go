/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `unicode` package. */
module UnicodeTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func SimpleFold(r rune) rune
      hasQualifiedName("unicode", "SimpleFold") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func To(_case int, r rune) rune
      hasQualifiedName("unicode", "To") and
      (inp.isParameter(1) and outp.isResult())
      or
      // signature: func ToLower(r rune) rune
      hasQualifiedName("unicode", "ToLower") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func ToTitle(r rune) rune
      hasQualifiedName("unicode", "ToTitle") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func ToUpper(r rune) rune
      hasQualifiedName("unicode", "ToUpper") and
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
      // signature: func (SpecialCase).ToLower(r rune) rune
      this.(Method).hasQualifiedName("unicode", "SpecialCase", "ToLower") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (SpecialCase).ToTitle(r rune) rune
      this.(Method).hasQualifiedName("unicode", "SpecialCase", "ToTitle") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (SpecialCase).ToUpper(r rune) rune
      this.(Method).hasQualifiedName("unicode", "SpecialCase", "ToUpper") and
      (inp.isParameter(0) and outp.isResult())
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
