/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `expvar` package. */
module ExpvarTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func Get(name string) Var
      hasQualifiedName("expvar", "Get") and
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
      // signature: func (*Float).String() string
      this.(Method).hasQualifiedName("expvar", "Float", "String") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Func).String() string
      this.(Method).hasQualifiedName("expvar", "Func", "String") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Func).Value() interface{}
      this.(Method).hasQualifiedName("expvar", "Func", "Value") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*Int).String() string
      this.(Method).hasQualifiedName("expvar", "Int", "String") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*Map).Do(f func(KeyValue))
      this.(Method).hasQualifiedName("expvar", "Map", "Do") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (*Map).Get(key string) Var
      this.(Method).hasQualifiedName("expvar", "Map", "Get") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*Map).Set(key string, av Var)
      this.(Method).hasQualifiedName("expvar", "Map", "Set") and
      (inp.isParameter(_) and outp.isReceiver())
      or
      // signature: func (*Map).String() string
      this.(Method).hasQualifiedName("expvar", "Map", "String") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*String).Set(value string)
      this.(Method).hasQualifiedName("expvar", "String", "Set") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (*String).String() string
      this.(Method).hasQualifiedName("expvar", "String", "String") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*String).Value() string
      this.(Method).hasQualifiedName("expvar", "String", "Value") and
      (inp.isReceiver() and outp.isResult())
      or
      // Interfaces:
      // signature: func (Var).String() string
      this.implements("expvar", "Var", "String") and
      (inp.isReceiver() and outp.isResult())
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
