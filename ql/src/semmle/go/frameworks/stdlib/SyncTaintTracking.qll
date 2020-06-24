/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `sync` package. */
module SyncTaintTracking {
  private class MethodAndInterfaceTaintTracking extends TaintTracking::FunctionModel, Method {
    FunctionInput inp;
    FunctionOutput outp;

    MethodAndInterfaceTaintTracking() {
      // Methods:
      // signature: func (*Map).Load(key interface{}) (value interface{}, ok bool)
      this.(Method).hasQualifiedName("sync", "Map", "Load") and
      (inp.isReceiver() and outp.isResult(0))
      or
      // signature: func (*Map).LoadOrStore(key interface{}, value interface{}) (actual interface{}, loaded bool)
      this.(Method).hasQualifiedName("sync", "Map", "LoadOrStore") and
      (inp.isReceiver() and outp.isResult(0))
      or
      inp.isParameter(1) and
      (outp.isReceiver() or outp.isResult(0))
      or
      // signature: func (*Map).Range(f func(key interface{}, value interface{}) bool)
      this.(Method).hasQualifiedName("sync", "Map", "Range") and
      (inp.isReceiver() and outp.isParameter(0))
      or
      // signature: func (*Map).Store(key interface{}, value interface{})
      this.(Method).hasQualifiedName("sync", "Map", "Store") and
      (inp.isParameter(_) and outp.isReceiver())
      or
      // signature: func (*Pool).Get() interface{}
      this.(Method).hasQualifiedName("sync", "Pool", "Get") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*Pool).Put(x interface{})
      this.(Method).hasQualifiedName("sync", "Pool", "Put") and
      (inp.isParameter(0) and outp.isReceiver())
      // Interfaces:
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
