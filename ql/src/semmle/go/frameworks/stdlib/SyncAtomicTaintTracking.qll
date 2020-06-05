/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `sync/atomic` package. */
module SyncAtomicTaintTracking {
  private class AddUintptr extends TaintTracking::FunctionModel {
    // signature: func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)
    AddUintptr() { hasQualifiedName("sync/atomic", "AddUintptr") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
      or
      inp.isParameter(_) and outp.isResult(0)
    }
  }

  private class CompareAndSwapPointer extends TaintTracking::FunctionModel {
    // signature: func CompareAndSwapPointer(addr *unsafe.Pointer, old unsafe.Pointer, new unsafe.Pointer) (swapped bool)
    CompareAndSwapPointer() { hasQualifiedName("sync/atomic", "CompareAndSwapPointer") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(2) and outp.isParameter(0)
    }
  }

  private class CompareAndSwapUintptr extends TaintTracking::FunctionModel {
    // signature: func CompareAndSwapUintptr(addr *uintptr, old uintptr, new uintptr) (swapped bool)
    CompareAndSwapUintptr() { hasQualifiedName("sync/atomic", "CompareAndSwapUintptr") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(2) and outp.isParameter(0)
    }
  }

  private class LoadPointer extends TaintTracking::FunctionModel {
    // signature: func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
    LoadPointer() { hasQualifiedName("sync/atomic", "LoadPointer") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class LoadUintptr extends TaintTracking::FunctionModel {
    // signature: func LoadUintptr(addr *uintptr) (val uintptr)
    LoadUintptr() { hasQualifiedName("sync/atomic", "LoadUintptr") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class StorePointer extends TaintTracking::FunctionModel {
    // signature: func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
    StorePointer() { hasQualifiedName("sync/atomic", "StorePointer") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class StoreUintptr extends TaintTracking::FunctionModel {
    // signature: func StoreUintptr(addr *uintptr, val uintptr)
    StoreUintptr() { hasQualifiedName("sync/atomic", "StoreUintptr") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class SwapPointer extends TaintTracking::FunctionModel {
    // signature: func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)
    SwapPointer() { hasQualifiedName("sync/atomic", "SwapPointer") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
      or
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class SwapUintptr extends TaintTracking::FunctionModel {
    // signature: func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
    SwapUintptr() { hasQualifiedName("sync/atomic", "SwapUintptr") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
      or
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class ValueLoad extends TaintTracking::FunctionModel, Method {
    // signature: func (*Value).Load() (x interface{})
    ValueLoad() { this.(Method).hasQualifiedName("sync/atomic", "Value", "Load") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueStore extends TaintTracking::FunctionModel, Method {
    // signature: func (*Value).Store(x interface{})
    ValueStore() { this.(Method).hasQualifiedName("sync/atomic", "Value", "Store") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }
}
