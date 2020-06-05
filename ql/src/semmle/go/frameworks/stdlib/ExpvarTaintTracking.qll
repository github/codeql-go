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
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class FloatAdd extends TaintTracking::FunctionModel, Method {
    // signature: func (*Float).Add(delta float64)
    FloatAdd() { this.(Method).hasQualifiedName("expvar", "Float", "Add") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class FloatSet extends TaintTracking::FunctionModel, Method {
    // signature: func (*Float).Set(value float64)
    FloatSet() { this.(Method).hasQualifiedName("expvar", "Float", "Set") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class FloatString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Float).String() string
    FloatString() { this.(Method).hasQualifiedName("expvar", "Float", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class FloatValue extends TaintTracking::FunctionModel, Method {
    // signature: func (*Float).Value() float64
    FloatValue() { this.(Method).hasQualifiedName("expvar", "Float", "Value") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class FuncString extends TaintTracking::FunctionModel, Method {
    // signature: func (Func).String() string
    FuncString() { this.(Method).hasQualifiedName("expvar", "Func", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class FuncValue extends TaintTracking::FunctionModel, Method {
    // signature: func (Func).Value() interface{}
    FuncValue() { this.(Method).hasQualifiedName("expvar", "Func", "Value") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class IntAdd extends TaintTracking::FunctionModel, Method {
    // signature: func (*Int).Add(delta int64)
    IntAdd() { this.(Method).hasQualifiedName("expvar", "Int", "Add") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class IntSet extends TaintTracking::FunctionModel, Method {
    // signature: func (*Int).Set(value int64)
    IntSet() { this.(Method).hasQualifiedName("expvar", "Int", "Set") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class IntString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Int).String() string
    IntString() { this.(Method).hasQualifiedName("expvar", "Int", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class IntValue extends TaintTracking::FunctionModel, Method {
    // signature: func (*Int).Value() int64
    IntValue() { this.(Method).hasQualifiedName("expvar", "Int", "Value") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class MapAdd extends TaintTracking::FunctionModel, Method {
    // signature: func (*Map).Add(key string, delta int64)
    MapAdd() { this.(Method).hasQualifiedName("expvar", "Map", "Add") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isReceiver()
    }
  }

  private class MapAddFloat extends TaintTracking::FunctionModel, Method {
    // signature: func (*Map).AddFloat(key string, delta float64)
    MapAddFloat() { this.(Method).hasQualifiedName("expvar", "Map", "AddFloat") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isReceiver()
    }
  }

  private class MapDo extends TaintTracking::FunctionModel, Method {
    // signature: func (*Map).Do(f func(KeyValue))
    MapDo() { this.(Method).hasQualifiedName("expvar", "Map", "Do") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class MapGet extends TaintTracking::FunctionModel, Method {
    // signature: func (*Map).Get(key string) Var
    MapGet() { this.(Method).hasQualifiedName("expvar", "Map", "Get") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class MapSet extends TaintTracking::FunctionModel, Method {
    // signature: func (*Map).Set(key string, av Var)
    MapSet() { this.(Method).hasQualifiedName("expvar", "Map", "Set") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isReceiver()
    }
  }

  private class MapString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Map).String() string
    MapString() { this.(Method).hasQualifiedName("expvar", "Map", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class StringSet extends TaintTracking::FunctionModel, Method {
    // signature: func (*String).Set(value string)
    StringSet() { this.(Method).hasQualifiedName("expvar", "String", "Set") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class StringString extends TaintTracking::FunctionModel, Method {
    // signature: func (*String).String() string
    StringString() { this.(Method).hasQualifiedName("expvar", "String", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class StringValue extends TaintTracking::FunctionModel, Method {
    // signature: func (*String).Value() string
    StringValue() { this.(Method).hasQualifiedName("expvar", "String", "Value") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class VarString extends TaintTracking::FunctionModel, Method {
    // signature: func (Var).String() string
    VarString() { this.implements("expvar", "Var", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }
}
