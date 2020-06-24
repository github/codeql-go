/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `reflect` package. */
module ReflectTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func Append(s Value, x ...Value) Value
      hasQualifiedName("reflect", "Append") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func AppendSlice(s Value, t Value) Value
      hasQualifiedName("reflect", "AppendSlice") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func Copy(dst Value, src Value) int
      hasQualifiedName("reflect", "Copy") and
      (inp.isParameter(1) and outp.isParameter(0))
      or
      // signature: func Indirect(v Value) Value
      hasQualifiedName("reflect", "Indirect") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func ValueOf(i interface{}) Value
      hasQualifiedName("reflect", "ValueOf") and
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
      // signature: func (*MapIter).Key() Value
      this.(Method).hasQualifiedName("reflect", "MapIter", "Key") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*MapIter).Value() Value
      this.(Method).hasQualifiedName("reflect", "MapIter", "Value") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (StructTag).Get(key string) string
      this.(Method).hasQualifiedName("reflect", "StructTag", "Get") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (StructTag).Lookup(key string) (value string, ok bool)
      this.(Method).hasQualifiedName("reflect", "StructTag", "Lookup") and
      (inp.isReceiver() and outp.isResult(0))
      or
      // signature: func (Value).Addr() Value
      this.(Method).hasQualifiedName("reflect", "Value", "Addr") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).Bool() bool
      this.(Method).hasQualifiedName("reflect", "Value", "Bool") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).Bytes() []byte
      this.(Method).hasQualifiedName("reflect", "Value", "Bytes") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).Complex() complex128
      this.(Method).hasQualifiedName("reflect", "Value", "Complex") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).Convert(t Type) Value
      this.(Method).hasQualifiedName("reflect", "Value", "Convert") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).Elem() Value
      this.(Method).hasQualifiedName("reflect", "Value", "Elem") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).Field(i int) Value
      this.(Method).hasQualifiedName("reflect", "Value", "Field") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).FieldByIndex(index []int) Value
      this.(Method).hasQualifiedName("reflect", "Value", "FieldByIndex") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).FieldByName(name string) Value
      this.(Method).hasQualifiedName("reflect", "Value", "FieldByName") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).FieldByNameFunc(match func(string) bool) Value
      this.(Method).hasQualifiedName("reflect", "Value", "FieldByNameFunc") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).Float() float64
      this.(Method).hasQualifiedName("reflect", "Value", "Float") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).Index(i int) Value
      this.(Method).hasQualifiedName("reflect", "Value", "Index") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).Int() int64
      this.(Method).hasQualifiedName("reflect", "Value", "Int") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).Interface() (i interface{})
      this.(Method).hasQualifiedName("reflect", "Value", "Interface") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).InterfaceData() [2]uintptr
      this.(Method).hasQualifiedName("reflect", "Value", "InterfaceData") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).MapIndex(key Value) Value
      this.(Method).hasQualifiedName("reflect", "Value", "MapIndex") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).MapKeys() []Value
      this.(Method).hasQualifiedName("reflect", "Value", "MapKeys") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).MapRange() *MapIter
      this.(Method).hasQualifiedName("reflect", "Value", "MapRange") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).Method(i int) Value
      this.(Method).hasQualifiedName("reflect", "Value", "Method") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).MethodByName(name string) Value
      this.(Method).hasQualifiedName("reflect", "Value", "MethodByName") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).Pointer() uintptr
      this.(Method).hasQualifiedName("reflect", "Value", "Pointer") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).Recv() (x Value, ok bool)
      this.(Method).hasQualifiedName("reflect", "Value", "Recv") and
      (inp.isReceiver() and outp.isResult(0))
      or
      // signature: func (Value).Send(x Value)
      this.(Method).hasQualifiedName("reflect", "Value", "Send") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (Value).Set(x Value)
      this.(Method).hasQualifiedName("reflect", "Value", "Set") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (Value).SetBool(x bool)
      this.(Method).hasQualifiedName("reflect", "Value", "SetBool") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (Value).SetBytes(x []byte)
      this.(Method).hasQualifiedName("reflect", "Value", "SetBytes") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (Value).SetComplex(x complex128)
      this.(Method).hasQualifiedName("reflect", "Value", "SetComplex") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (Value).SetFloat(x float64)
      this.(Method).hasQualifiedName("reflect", "Value", "SetFloat") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (Value).SetInt(x int64)
      this.(Method).hasQualifiedName("reflect", "Value", "SetInt") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (Value).SetMapIndex(key Value, elem Value)
      this.(Method).hasQualifiedName("reflect", "Value", "SetMapIndex") and
      (inp.isParameter(_) and outp.isReceiver())
      or
      // signature: func (Value).SetPointer(x unsafe.Pointer)
      this.(Method).hasQualifiedName("reflect", "Value", "SetPointer") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (Value).SetString(x string)
      this.(Method).hasQualifiedName("reflect", "Value", "SetString") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (Value).SetUint(x uint64)
      this.(Method).hasQualifiedName("reflect", "Value", "SetUint") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (Value).Slice(i int, j int) Value
      this.(Method).hasQualifiedName("reflect", "Value", "Slice") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).Slice3(i int, j int, k int) Value
      this.(Method).hasQualifiedName("reflect", "Value", "Slice3") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).String() string
      this.(Method).hasQualifiedName("reflect", "Value", "String") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).TryRecv() (x Value, ok bool)
      this.(Method).hasQualifiedName("reflect", "Value", "TryRecv") and
      (inp.isReceiver() and outp.isResult(0))
      or
      // signature: func (Value).TrySend(x Value) bool
      this.(Method).hasQualifiedName("reflect", "Value", "TrySend") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (Value).Uint() uint64
      this.(Method).hasQualifiedName("reflect", "Value", "Uint") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Value).UnsafeAddr() uintptr
      this.(Method).hasQualifiedName("reflect", "Value", "UnsafeAddr") and
      (inp.isReceiver() and outp.isResult())
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
