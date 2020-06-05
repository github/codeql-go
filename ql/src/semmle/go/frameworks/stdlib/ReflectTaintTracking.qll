/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `reflect` package. */
module ReflectTaintTracking {
  private class Append extends TaintTracking::FunctionModel {
    // signature: func Append(s Value, x ...Value) Value
    Append() { hasQualifiedName("reflect", "Append") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isResult()
    }
  }

  private class AppendSlice extends TaintTracking::FunctionModel {
    // signature: func AppendSlice(s Value, t Value) Value
    AppendSlice() { hasQualifiedName("reflect", "AppendSlice") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isResult()
    }
  }

  private class Copy extends TaintTracking::FunctionModel {
    // signature: func Copy(dst Value, src Value) int
    Copy() { hasQualifiedName("reflect", "Copy") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class Indirect extends TaintTracking::FunctionModel {
    // signature: func Indirect(v Value) Value
    Indirect() { hasQualifiedName("reflect", "Indirect") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class ValueOf extends TaintTracking::FunctionModel {
    // signature: func ValueOf(i interface{}) Value
    ValueOf() { hasQualifiedName("reflect", "ValueOf") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class MapIterKey extends TaintTracking::FunctionModel, Method {
    // signature: func (*MapIter).Key() Value
    MapIterKey() { this.(Method).hasQualifiedName("reflect", "MapIter", "Key") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class MapIterValue extends TaintTracking::FunctionModel, Method {
    // signature: func (*MapIter).Value() Value
    MapIterValue() { this.(Method).hasQualifiedName("reflect", "MapIter", "Value") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class StructTagGet extends TaintTracking::FunctionModel, Method {
    // signature: func (StructTag).Get(key string) string
    StructTagGet() { this.(Method).hasQualifiedName("reflect", "StructTag", "Get") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class StructTagLookup extends TaintTracking::FunctionModel, Method {
    // signature: func (StructTag).Lookup(key string) (value string, ok bool)
    StructTagLookup() { this.(Method).hasQualifiedName("reflect", "StructTag", "Lookup") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class ValueAddr extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Addr() Value
    ValueAddr() { this.(Method).hasQualifiedName("reflect", "Value", "Addr") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueBool extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Bool() bool
    ValueBool() { this.(Method).hasQualifiedName("reflect", "Value", "Bool") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueBytes extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Bytes() []byte
    ValueBytes() { this.(Method).hasQualifiedName("reflect", "Value", "Bytes") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueComplex extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Complex() complex128
    ValueComplex() { this.(Method).hasQualifiedName("reflect", "Value", "Complex") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueConvert extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Convert(t Type) Value
    ValueConvert() { this.(Method).hasQualifiedName("reflect", "Value", "Convert") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueElem extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Elem() Value
    ValueElem() { this.(Method).hasQualifiedName("reflect", "Value", "Elem") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueField extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Field(i int) Value
    ValueField() { this.(Method).hasQualifiedName("reflect", "Value", "Field") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueFieldByIndex extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).FieldByIndex(index []int) Value
    ValueFieldByIndex() { this.(Method).hasQualifiedName("reflect", "Value", "FieldByIndex") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueFieldByName extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).FieldByName(name string) Value
    ValueFieldByName() { this.(Method).hasQualifiedName("reflect", "Value", "FieldByName") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueFieldByNameFunc extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).FieldByNameFunc(match func(string) bool) Value
    ValueFieldByNameFunc() { this.(Method).hasQualifiedName("reflect", "Value", "FieldByNameFunc") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueFloat extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Float() float64
    ValueFloat() { this.(Method).hasQualifiedName("reflect", "Value", "Float") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueIndex extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Index(i int) Value
    ValueIndex() { this.(Method).hasQualifiedName("reflect", "Value", "Index") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueInt extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Int() int64
    ValueInt() { this.(Method).hasQualifiedName("reflect", "Value", "Int") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueInterface extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Interface() (i interface{})
    ValueInterface() { this.(Method).hasQualifiedName("reflect", "Value", "Interface") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueInterfaceData extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).InterfaceData() [2]uintptr
    ValueInterfaceData() { this.(Method).hasQualifiedName("reflect", "Value", "InterfaceData") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueMapIndex extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).MapIndex(key Value) Value
    ValueMapIndex() { this.(Method).hasQualifiedName("reflect", "Value", "MapIndex") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueMapKeys extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).MapKeys() []Value
    ValueMapKeys() { this.(Method).hasQualifiedName("reflect", "Value", "MapKeys") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueMapRange extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).MapRange() *MapIter
    ValueMapRange() { this.(Method).hasQualifiedName("reflect", "Value", "MapRange") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueMethod extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Method(i int) Value
    ValueMethod() { this.(Method).hasQualifiedName("reflect", "Value", "Method") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueMethodByName extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).MethodByName(name string) Value
    ValueMethodByName() { this.(Method).hasQualifiedName("reflect", "Value", "MethodByName") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValuePointer extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Pointer() uintptr
    ValuePointer() { this.(Method).hasQualifiedName("reflect", "Value", "Pointer") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueRecv extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Recv() (x Value, ok bool)
    ValueRecv() { this.(Method).hasQualifiedName("reflect", "Value", "Recv") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class ValueSend extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Send(x Value)
    ValueSend() { this.(Method).hasQualifiedName("reflect", "Value", "Send") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ValueSet extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Set(x Value)
    ValueSet() { this.(Method).hasQualifiedName("reflect", "Value", "Set") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ValueSetBool extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).SetBool(x bool)
    ValueSetBool() { this.(Method).hasQualifiedName("reflect", "Value", "SetBool") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ValueSetBytes extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).SetBytes(x []byte)
    ValueSetBytes() { this.(Method).hasQualifiedName("reflect", "Value", "SetBytes") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ValueSetComplex extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).SetComplex(x complex128)
    ValueSetComplex() { this.(Method).hasQualifiedName("reflect", "Value", "SetComplex") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ValueSetFloat extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).SetFloat(x float64)
    ValueSetFloat() { this.(Method).hasQualifiedName("reflect", "Value", "SetFloat") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ValueSetInt extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).SetInt(x int64)
    ValueSetInt() { this.(Method).hasQualifiedName("reflect", "Value", "SetInt") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ValueSetMapIndex extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).SetMapIndex(key Value, elem Value)
    ValueSetMapIndex() { this.(Method).hasQualifiedName("reflect", "Value", "SetMapIndex") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isReceiver()
    }
  }

  private class ValueSetPointer extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).SetPointer(x unsafe.Pointer)
    ValueSetPointer() { this.(Method).hasQualifiedName("reflect", "Value", "SetPointer") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ValueSetString extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).SetString(x string)
    ValueSetString() { this.(Method).hasQualifiedName("reflect", "Value", "SetString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ValueSetUint extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).SetUint(x uint64)
    ValueSetUint() { this.(Method).hasQualifiedName("reflect", "Value", "SetUint") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ValueSlice extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Slice(i int, j int) Value
    ValueSlice() { this.(Method).hasQualifiedName("reflect", "Value", "Slice") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueSlice3 extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Slice3(i int, j int, k int) Value
    ValueSlice3() { this.(Method).hasQualifiedName("reflect", "Value", "Slice3") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueString extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).String() string
    ValueString() { this.(Method).hasQualifiedName("reflect", "Value", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueTryRecv extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).TryRecv() (x Value, ok bool)
    ValueTryRecv() { this.(Method).hasQualifiedName("reflect", "Value", "TryRecv") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class ValueTrySend extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).TrySend(x Value) bool
    ValueTrySend() { this.(Method).hasQualifiedName("reflect", "Value", "TrySend") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ValueUint extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).Uint() uint64
    ValueUint() { this.(Method).hasQualifiedName("reflect", "Value", "Uint") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ValueUnsafeAddr extends TaintTracking::FunctionModel, Method {
    // signature: func (Value).UnsafeAddr() uintptr
    ValueUnsafeAddr() { this.(Method).hasQualifiedName("reflect", "Value", "UnsafeAddr") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }
}
