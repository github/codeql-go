/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `expvar` package. */
module ExpvarTaintTracking {
  private class Get extends TaintTracking::FunctionModel {
    // signature: func Get(name string) Var
    Get() { hasQualifiedName("expvar", "Get") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class FloatString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Float).String() string
    FloatString() { this.(Method).hasQualifiedName("expvar", "Float", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class FuncString extends TaintTracking::FunctionModel, Method {
    // signature: func (Func).String() string
    FuncString() { this.(Method).hasQualifiedName("expvar", "Func", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class FuncValue extends TaintTracking::FunctionModel, Method {
    // signature: func (Func).Value() interface{}
    FuncValue() { this.(Method).hasQualifiedName("expvar", "Func", "Value") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class IntString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Int).String() string
    IntString() { this.(Method).hasQualifiedName("expvar", "Int", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class MapDo extends TaintTracking::FunctionModel, Method {
    // signature: func (*Map).Do(f func(KeyValue))
    MapDo() { this.(Method).hasQualifiedName("expvar", "Map", "Do") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class MapGet extends TaintTracking::FunctionModel, Method {
    // signature: func (*Map).Get(key string) Var
    MapGet() { this.(Method).hasQualifiedName("expvar", "Map", "Get") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class MapSet extends TaintTracking::FunctionModel, Method {
    // signature: func (*Map).Set(key string, av Var)
    MapSet() { this.(Method).hasQualifiedName("expvar", "Map", "Set") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isReceiver())
    }
  }

  private class MapString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Map).String() string
    MapString() { this.(Method).hasQualifiedName("expvar", "Map", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class StringSet extends TaintTracking::FunctionModel, Method {
    // signature: func (*String).Set(value string)
    StringSet() { this.(Method).hasQualifiedName("expvar", "String", "Set") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class StringString extends TaintTracking::FunctionModel, Method {
    // signature: func (*String).String() string
    StringString() { this.(Method).hasQualifiedName("expvar", "String", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class StringValue extends TaintTracking::FunctionModel, Method {
    // signature: func (*String).Value() string
    StringValue() { this.(Method).hasQualifiedName("expvar", "String", "Value") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class VarString extends TaintTracking::FunctionModel, Method {
    // signature: func (Var).String() string
    VarString() { this.implements("expvar", "Var", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }
}
