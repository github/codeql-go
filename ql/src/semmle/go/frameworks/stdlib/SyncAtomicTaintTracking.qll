/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `sync/atomic` package. */
module SyncAtomicTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)
      hasQualifiedName("sync/atomic", "AddUintptr") and
      (
        inp.isParameter(1) and
        (outp.isParameter(0) or outp.isResult())
      )
      or
      // signature: func CompareAndSwapPointer(addr *unsafe.Pointer, old unsafe.Pointer, new unsafe.Pointer) (swapped bool)
      hasQualifiedName("sync/atomic", "CompareAndSwapPointer") and
      (inp.isParameter(2) and outp.isParameter(0))
      or
      // signature: func CompareAndSwapUintptr(addr *uintptr, old uintptr, new uintptr) (swapped bool)
      hasQualifiedName("sync/atomic", "CompareAndSwapUintptr") and
      (inp.isParameter(2) and outp.isParameter(0))
      or
      // signature: func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
      hasQualifiedName("sync/atomic", "LoadPointer") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func LoadUintptr(addr *uintptr) (val uintptr)
      hasQualifiedName("sync/atomic", "LoadUintptr") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
      hasQualifiedName("sync/atomic", "StorePointer") and
      (inp.isParameter(1) and outp.isParameter(0))
      or
      // signature: func StoreUintptr(addr *uintptr, val uintptr)
      hasQualifiedName("sync/atomic", "StoreUintptr") and
      (inp.isParameter(1) and outp.isParameter(0))
      or
      // signature: func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)
      hasQualifiedName("sync/atomic", "SwapPointer") and
      (inp.isParameter(1) and outp.isParameter(0))
      or
      inp.isParameter(0) and outp.isResult()
      or
      // signature: func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
      hasQualifiedName("sync/atomic", "SwapUintptr") and
      (inp.isParameter(1) and outp.isParameter(0))
      or
      inp.isParameter(0) and outp.isResult()
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
      // signature: func (*Value).Load() (x interface{})
      this.(Method).hasQualifiedName("sync/atomic", "Value", "Load") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*Value).Store(x interface{})
      this.(Method).hasQualifiedName("sync/atomic", "Value", "Store") and
      (inp.isParameter(0) and outp.isReceiver())
      // Interfaces:
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
