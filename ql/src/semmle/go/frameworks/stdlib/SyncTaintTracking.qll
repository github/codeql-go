/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `sync` package. */
module SyncTaintTracking {
  private class MapLoad extends TaintTracking::FunctionModel, Method {
    // signature: func (*Map).Load(key interface{}) (value interface{}, ok bool)
    MapLoad() { this.(Method).hasQualifiedName("sync", "Map", "Load") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }

  private class MapLoadOrStore extends TaintTracking::FunctionModel, Method {
    // signature: func (*Map).LoadOrStore(key interface{}, value interface{}) (actual interface{}, loaded bool)
    MapLoadOrStore() { this.(Method).hasQualifiedName("sync", "Map", "LoadOrStore") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
      or
      inp.isParameter(1) and
      (outp.isReceiver() or outp.isResult(0))
    }
  }

  private class MapRange extends TaintTracking::FunctionModel, Method {
    // signature: func (*Map).Range(f func(key interface{}, value interface{}) bool)
    MapRange() { this.(Method).hasQualifiedName("sync", "Map", "Range") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class MapStore extends TaintTracking::FunctionModel, Method {
    // signature: func (*Map).Store(key interface{}, value interface{})
    MapStore() { this.(Method).hasQualifiedName("sync", "Map", "Store") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isReceiver())
    }
  }

  private class PoolGet extends TaintTracking::FunctionModel, Method {
    // signature: func (*Pool).Get() interface{}
    PoolGet() { this.(Method).hasQualifiedName("sync", "Pool", "Get") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class PoolPut extends TaintTracking::FunctionModel, Method {
    // signature: func (*Pool).Put(x interface{})
    PoolPut() { this.(Method).hasQualifiedName("sync", "Pool", "Put") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }
}
