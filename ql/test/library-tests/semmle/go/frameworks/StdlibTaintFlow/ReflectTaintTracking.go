// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"reflect"
	"unsafe"
)

func TaintStepTest_ReflectAppend_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue200` into `intoValue838`.

	// Assume that `sourceCQL` has the underlying type of `fromValue200`:
	fromValue200 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue200` to result `intoValue838`
	// (`intoValue838` is now tainted).
	intoValue838 := reflect.Append(fromValue200, reflect.Value{})

	// Return the tainted `intoValue838`:
	return intoValue838
}

func TaintStepTest_ReflectAppend_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue810` into `intoValue812`.

	// Assume that `sourceCQL` has the underlying type of `fromValue810`:
	fromValue810 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue810` to result `intoValue812`
	// (`intoValue812` is now tainted).
	intoValue812 := reflect.Append(reflect.Value{}, fromValue810)

	// Return the tainted `intoValue812`:
	return intoValue812
}

func TaintStepTest_ReflectAppendSlice_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue734` into `intoValue498`.

	// Assume that `sourceCQL` has the underlying type of `fromValue734`:
	fromValue734 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue734` to result `intoValue498`
	// (`intoValue498` is now tainted).
	intoValue498 := reflect.AppendSlice(fromValue734, reflect.Value{})

	// Return the tainted `intoValue498`:
	return intoValue498
}

func TaintStepTest_ReflectAppendSlice_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue299` into `intoValue696`.

	// Assume that `sourceCQL` has the underlying type of `fromValue299`:
	fromValue299 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue299` to result `intoValue696`
	// (`intoValue696` is now tainted).
	intoValue696 := reflect.AppendSlice(reflect.Value{}, fromValue299)

	// Return the tainted `intoValue696`:
	return intoValue696
}

func TaintStepTest_ReflectCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue525` into `intoValue275`.

	// Assume that `sourceCQL` has the underlying type of `fromValue525`:
	fromValue525 := sourceCQL.(reflect.Value)

	// Declare `intoValue275` variable:
	var intoValue275 reflect.Value

	// Call the function that transfers the taint
	// from the parameter `fromValue525` to parameter `intoValue275`;
	// `intoValue275` is now tainted.
	reflect.Copy(intoValue275, fromValue525)

	// Return the tainted `intoValue275`:
	return intoValue275
}

func TaintStepTest_ReflectIndirect_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue249` into `intoValue186`.

	// Assume that `sourceCQL` has the underlying type of `fromValue249`:
	fromValue249 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue249` to result `intoValue186`
	// (`intoValue186` is now tainted).
	intoValue186 := reflect.Indirect(fromValue249)

	// Return the tainted `intoValue186`:
	return intoValue186
}

func TaintStepTest_ReflectValueOf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface112` into `intoValue484`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface112`:
	fromInterface112 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface112` to result `intoValue484`
	// (`intoValue484` is now tainted).
	intoValue484 := reflect.ValueOf(fromInterface112)

	// Return the tainted `intoValue484`:
	return intoValue484
}

func TaintStepTest_ReflectMapIterKey_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMapIter866` into `intoValue149`.

	// Assume that `sourceCQL` has the underlying type of `fromMapIter866`:
	fromMapIter866 := sourceCQL.(reflect.MapIter)

	// Call the method that transfers the taint
	// from the receiver `fromMapIter866` to the result `intoValue149`
	// (`intoValue149` is now tainted).
	intoValue149 := fromMapIter866.Key()

	// Return the tainted `intoValue149`:
	return intoValue149
}

func TaintStepTest_ReflectMapIterValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMapIter145` into `intoValue275`.

	// Assume that `sourceCQL` has the underlying type of `fromMapIter145`:
	fromMapIter145 := sourceCQL.(reflect.MapIter)

	// Call the method that transfers the taint
	// from the receiver `fromMapIter145` to the result `intoValue275`
	// (`intoValue275` is now tainted).
	intoValue275 := fromMapIter145.Value()

	// Return the tainted `intoValue275`:
	return intoValue275
}

func TaintStepTest_ReflectStructTagGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStructTag335` into `intoString834`.

	// Assume that `sourceCQL` has the underlying type of `fromStructTag335`:
	fromStructTag335 := sourceCQL.(reflect.StructTag)

	// Call the method that transfers the taint
	// from the receiver `fromStructTag335` to the result `intoString834`
	// (`intoString834` is now tainted).
	intoString834 := fromStructTag335.Get("")

	// Return the tainted `intoString834`:
	return intoString834
}

func TaintStepTest_ReflectStructTagLookup_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStructTag357` into `intoString233`.

	// Assume that `sourceCQL` has the underlying type of `fromStructTag357`:
	fromStructTag357 := sourceCQL.(reflect.StructTag)

	// Call the method that transfers the taint
	// from the receiver `fromStructTag357` to the result `intoString233`
	// (`intoString233` is now tainted).
	intoString233, _ := fromStructTag357.Lookup("")

	// Return the tainted `intoString233`:
	return intoString233
}

func TaintStepTest_ReflectValueAddr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue215` into `intoValue335`.

	// Assume that `sourceCQL` has the underlying type of `fromValue215`:
	fromValue215 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue215` to the result `intoValue335`
	// (`intoValue335` is now tainted).
	intoValue335 := fromValue215.Addr()

	// Return the tainted `intoValue335`:
	return intoValue335
}

func TaintStepTest_ReflectValueBool_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue123` into `intoBool365`.

	// Assume that `sourceCQL` has the underlying type of `fromValue123`:
	fromValue123 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue123` to the result `intoBool365`
	// (`intoBool365` is now tainted).
	intoBool365 := fromValue123.Bool()

	// Return the tainted `intoBool365`:
	return intoBool365
}

func TaintStepTest_ReflectValueBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue373` into `intoByte883`.

	// Assume that `sourceCQL` has the underlying type of `fromValue373`:
	fromValue373 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue373` to the result `intoByte883`
	// (`intoByte883` is now tainted).
	intoByte883 := fromValue373.Bytes()

	// Return the tainted `intoByte883`:
	return intoByte883
}

func TaintStepTest_ReflectValueComplex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue249` into `intoComplex128700`.

	// Assume that `sourceCQL` has the underlying type of `fromValue249`:
	fromValue249 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue249` to the result `intoComplex128700`
	// (`intoComplex128700` is now tainted).
	intoComplex128700 := fromValue249.Complex()

	// Return the tainted `intoComplex128700`:
	return intoComplex128700
}

func TaintStepTest_ReflectValueConvert_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue509` into `intoValue998`.

	// Assume that `sourceCQL` has the underlying type of `fromValue509`:
	fromValue509 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue509` to the result `intoValue998`
	// (`intoValue998` is now tainted).
	intoValue998 := fromValue509.Convert(nil)

	// Return the tainted `intoValue998`:
	return intoValue998
}

func TaintStepTest_ReflectValueElem_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue633` into `intoValue664`.

	// Assume that `sourceCQL` has the underlying type of `fromValue633`:
	fromValue633 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue633` to the result `intoValue664`
	// (`intoValue664` is now tainted).
	intoValue664 := fromValue633.Elem()

	// Return the tainted `intoValue664`:
	return intoValue664
}

func TaintStepTest_ReflectValueField_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue531` into `intoValue564`.

	// Assume that `sourceCQL` has the underlying type of `fromValue531`:
	fromValue531 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue531` to the result `intoValue564`
	// (`intoValue564` is now tainted).
	intoValue564 := fromValue531.Field(0)

	// Return the tainted `intoValue564`:
	return intoValue564
}

func TaintStepTest_ReflectValueFieldByIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue904` into `intoValue562`.

	// Assume that `sourceCQL` has the underlying type of `fromValue904`:
	fromValue904 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue904` to the result `intoValue562`
	// (`intoValue562` is now tainted).
	intoValue562 := fromValue904.FieldByIndex(nil)

	// Return the tainted `intoValue562`:
	return intoValue562
}

func TaintStepTest_ReflectValueFieldByName_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue263` into `intoValue551`.

	// Assume that `sourceCQL` has the underlying type of `fromValue263`:
	fromValue263 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue263` to the result `intoValue551`
	// (`intoValue551` is now tainted).
	intoValue551 := fromValue263.FieldByName("")

	// Return the tainted `intoValue551`:
	return intoValue551
}

func TaintStepTest_ReflectValueFieldByNameFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue322` into `intoValue797`.

	// Assume that `sourceCQL` has the underlying type of `fromValue322`:
	fromValue322 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue322` to the result `intoValue797`
	// (`intoValue797` is now tainted).
	intoValue797 := fromValue322.FieldByNameFunc(nil)

	// Return the tainted `intoValue797`:
	return intoValue797
}

func TaintStepTest_ReflectValueFloat_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue150` into `intoFloat64694`.

	// Assume that `sourceCQL` has the underlying type of `fromValue150`:
	fromValue150 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue150` to the result `intoFloat64694`
	// (`intoFloat64694` is now tainted).
	intoFloat64694 := fromValue150.Float()

	// Return the tainted `intoFloat64694`:
	return intoFloat64694
}

func TaintStepTest_ReflectValueIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue498` into `intoValue822`.

	// Assume that `sourceCQL` has the underlying type of `fromValue498`:
	fromValue498 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue498` to the result `intoValue822`
	// (`intoValue822` is now tainted).
	intoValue822 := fromValue498.Index(0)

	// Return the tainted `intoValue822`:
	return intoValue822
}

func TaintStepTest_ReflectValueInt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue495` into `intoInt64574`.

	// Assume that `sourceCQL` has the underlying type of `fromValue495`:
	fromValue495 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue495` to the result `intoInt64574`
	// (`intoInt64574` is now tainted).
	intoInt64574 := fromValue495.Int()

	// Return the tainted `intoInt64574`:
	return intoInt64574
}

func TaintStepTest_ReflectValueInterface_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue627` into `intoInterface543`.

	// Assume that `sourceCQL` has the underlying type of `fromValue627`:
	fromValue627 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue627` to the result `intoInterface543`
	// (`intoInterface543` is now tainted).
	intoInterface543 := fromValue627.Interface()

	// Return the tainted `intoInterface543`:
	return intoInterface543
}

func TaintStepTest_ReflectValueInterfaceData_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue299` into `intoUintptr303`.

	// Assume that `sourceCQL` has the underlying type of `fromValue299`:
	fromValue299 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue299` to the result `intoUintptr303`
	// (`intoUintptr303` is now tainted).
	intoUintptr303 := fromValue299.InterfaceData()

	// Return the tainted `intoUintptr303`:
	return intoUintptr303
}

func TaintStepTest_ReflectValueMapIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue855` into `intoValue313`.

	// Assume that `sourceCQL` has the underlying type of `fromValue855`:
	fromValue855 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue855` to the result `intoValue313`
	// (`intoValue313` is now tainted).
	intoValue313 := fromValue855.MapIndex(reflect.Value{})

	// Return the tainted `intoValue313`:
	return intoValue313
}

func TaintStepTest_ReflectValueMapKeys_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue483` into `intoValue259`.

	// Assume that `sourceCQL` has the underlying type of `fromValue483`:
	fromValue483 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue483` to the result `intoValue259`
	// (`intoValue259` is now tainted).
	intoValue259 := fromValue483.MapKeys()

	// Return the tainted `intoValue259`:
	return intoValue259
}

func TaintStepTest_ReflectValueMapRange_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue800` into `intoMapIter890`.

	// Assume that `sourceCQL` has the underlying type of `fromValue800`:
	fromValue800 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue800` to the result `intoMapIter890`
	// (`intoMapIter890` is now tainted).
	intoMapIter890 := fromValue800.MapRange()

	// Return the tainted `intoMapIter890`:
	return intoMapIter890
}

func TaintStepTest_ReflectValueMethod_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue957` into `intoValue816`.

	// Assume that `sourceCQL` has the underlying type of `fromValue957`:
	fromValue957 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue957` to the result `intoValue816`
	// (`intoValue816` is now tainted).
	intoValue816 := fromValue957.Method(0)

	// Return the tainted `intoValue816`:
	return intoValue816
}

func TaintStepTest_ReflectValueMethodByName_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue878` into `intoValue148`.

	// Assume that `sourceCQL` has the underlying type of `fromValue878`:
	fromValue878 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue878` to the result `intoValue148`
	// (`intoValue148` is now tainted).
	intoValue148 := fromValue878.MethodByName("")

	// Return the tainted `intoValue148`:
	return intoValue148
}

func TaintStepTest_ReflectValuePointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue408` into `intoUintptr312`.

	// Assume that `sourceCQL` has the underlying type of `fromValue408`:
	fromValue408 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue408` to the result `intoUintptr312`
	// (`intoUintptr312` is now tainted).
	intoUintptr312 := fromValue408.Pointer()

	// Return the tainted `intoUintptr312`:
	return intoUintptr312
}

func TaintStepTest_ReflectValueRecv_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue334` into `intoValue561`.

	// Assume that `sourceCQL` has the underlying type of `fromValue334`:
	fromValue334 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue334` to the result `intoValue561`
	// (`intoValue561` is now tainted).
	intoValue561, _ := fromValue334.Recv()

	// Return the tainted `intoValue561`:
	return intoValue561
}

func TaintStepTest_ReflectValueSend_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue851` into `intoValue650`.

	// Assume that `sourceCQL` has the underlying type of `fromValue851`:
	fromValue851 := sourceCQL.(reflect.Value)

	// Declare `intoValue650` variable:
	var intoValue650 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue851` to the receiver `intoValue650`
	// (`intoValue650` is now tainted).
	intoValue650.Send(fromValue851)

	// Return the tainted `intoValue650`:
	return intoValue650
}

func TaintStepTest_ReflectValueSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue624` into `intoValue985`.

	// Assume that `sourceCQL` has the underlying type of `fromValue624`:
	fromValue624 := sourceCQL.(reflect.Value)

	// Declare `intoValue985` variable:
	var intoValue985 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue624` to the receiver `intoValue985`
	// (`intoValue985` is now tainted).
	intoValue985.Set(fromValue624)

	// Return the tainted `intoValue985`:
	return intoValue985
}

func TaintStepTest_ReflectValueSetBool_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBool295` into `intoValue247`.

	// Assume that `sourceCQL` has the underlying type of `fromBool295`:
	fromBool295 := sourceCQL.(bool)

	// Declare `intoValue247` variable:
	var intoValue247 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromBool295` to the receiver `intoValue247`
	// (`intoValue247` is now tainted).
	intoValue247.SetBool(fromBool295)

	// Return the tainted `intoValue247`:
	return intoValue247
}

func TaintStepTest_ReflectValueSetBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte532` into `intoValue615`.

	// Assume that `sourceCQL` has the underlying type of `fromByte532`:
	fromByte532 := sourceCQL.([]byte)

	// Declare `intoValue615` variable:
	var intoValue615 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromByte532` to the receiver `intoValue615`
	// (`intoValue615` is now tainted).
	intoValue615.SetBytes(fromByte532)

	// Return the tainted `intoValue615`:
	return intoValue615
}

func TaintStepTest_ReflectValueSetComplex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromComplex128769` into `intoValue961`.

	// Assume that `sourceCQL` has the underlying type of `fromComplex128769`:
	fromComplex128769 := sourceCQL.(complex128)

	// Declare `intoValue961` variable:
	var intoValue961 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromComplex128769` to the receiver `intoValue961`
	// (`intoValue961` is now tainted).
	intoValue961.SetComplex(fromComplex128769)

	// Return the tainted `intoValue961`:
	return intoValue961
}

func TaintStepTest_ReflectValueSetFloat_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFloat64784` into `intoValue438`.

	// Assume that `sourceCQL` has the underlying type of `fromFloat64784`:
	fromFloat64784 := sourceCQL.(float64)

	// Declare `intoValue438` variable:
	var intoValue438 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromFloat64784` to the receiver `intoValue438`
	// (`intoValue438` is now tainted).
	intoValue438.SetFloat(fromFloat64784)

	// Return the tainted `intoValue438`:
	return intoValue438
}

func TaintStepTest_ReflectValueSetInt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInt64376` into `intoValue136`.

	// Assume that `sourceCQL` has the underlying type of `fromInt64376`:
	fromInt64376 := sourceCQL.(int64)

	// Declare `intoValue136` variable:
	var intoValue136 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromInt64376` to the receiver `intoValue136`
	// (`intoValue136` is now tainted).
	intoValue136.SetInt(fromInt64376)

	// Return the tainted `intoValue136`:
	return intoValue136
}

func TaintStepTest_ReflectValueSetMapIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue752` into `intoValue304`.

	// Assume that `sourceCQL` has the underlying type of `fromValue752`:
	fromValue752 := sourceCQL.(reflect.Value)

	// Declare `intoValue304` variable:
	var intoValue304 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue752` to the receiver `intoValue304`
	// (`intoValue304` is now tainted).
	intoValue304.SetMapIndex(fromValue752, reflect.Value{})

	// Return the tainted `intoValue304`:
	return intoValue304
}

func TaintStepTest_ReflectValueSetMapIndex_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue241` into `intoValue927`.

	// Assume that `sourceCQL` has the underlying type of `fromValue241`:
	fromValue241 := sourceCQL.(reflect.Value)

	// Declare `intoValue927` variable:
	var intoValue927 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue241` to the receiver `intoValue927`
	// (`intoValue927` is now tainted).
	intoValue927.SetMapIndex(reflect.Value{}, fromValue241)

	// Return the tainted `intoValue927`:
	return intoValue927
}

func TaintStepTest_ReflectValueSetPointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer309` into `intoValue143`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer309`:
	fromPointer309 := sourceCQL.(unsafe.Pointer)

	// Declare `intoValue143` variable:
	var intoValue143 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromPointer309` to the receiver `intoValue143`
	// (`intoValue143` is now tainted).
	intoValue143.SetPointer(fromPointer309)

	// Return the tainted `intoValue143`:
	return intoValue143
}

func TaintStepTest_ReflectValueSetString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString636` into `intoValue678`.

	// Assume that `sourceCQL` has the underlying type of `fromString636`:
	fromString636 := sourceCQL.(string)

	// Declare `intoValue678` variable:
	var intoValue678 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromString636` to the receiver `intoValue678`
	// (`intoValue678` is now tainted).
	intoValue678.SetString(fromString636)

	// Return the tainted `intoValue678`:
	return intoValue678
}

func TaintStepTest_ReflectValueSetUint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUint64258` into `intoValue448`.

	// Assume that `sourceCQL` has the underlying type of `fromUint64258`:
	fromUint64258 := sourceCQL.(uint64)

	// Declare `intoValue448` variable:
	var intoValue448 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromUint64258` to the receiver `intoValue448`
	// (`intoValue448` is now tainted).
	intoValue448.SetUint(fromUint64258)

	// Return the tainted `intoValue448`:
	return intoValue448
}

func TaintStepTest_ReflectValueSlice_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue853` into `intoValue117`.

	// Assume that `sourceCQL` has the underlying type of `fromValue853`:
	fromValue853 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue853` to the result `intoValue117`
	// (`intoValue117` is now tainted).
	intoValue117 := fromValue853.Slice(0, 0)

	// Return the tainted `intoValue117`:
	return intoValue117
}

func TaintStepTest_ReflectValueSlice3_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue477` into `intoValue805`.

	// Assume that `sourceCQL` has the underlying type of `fromValue477`:
	fromValue477 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue477` to the result `intoValue805`
	// (`intoValue805` is now tainted).
	intoValue805 := fromValue477.Slice3(0, 0, 0)

	// Return the tainted `intoValue805`:
	return intoValue805
}

func TaintStepTest_ReflectValueString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue838` into `intoString205`.

	// Assume that `sourceCQL` has the underlying type of `fromValue838`:
	fromValue838 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue838` to the result `intoString205`
	// (`intoString205` is now tainted).
	intoString205 := fromValue838.String()

	// Return the tainted `intoString205`:
	return intoString205
}

func TaintStepTest_ReflectValueTryRecv_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue629` into `intoValue519`.

	// Assume that `sourceCQL` has the underlying type of `fromValue629`:
	fromValue629 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue629` to the result `intoValue519`
	// (`intoValue519` is now tainted).
	intoValue519, _ := fromValue629.TryRecv()

	// Return the tainted `intoValue519`:
	return intoValue519
}

func TaintStepTest_ReflectValueTrySend_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue513` into `intoValue875`.

	// Assume that `sourceCQL` has the underlying type of `fromValue513`:
	fromValue513 := sourceCQL.(reflect.Value)

	// Declare `intoValue875` variable:
	var intoValue875 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue513` to the receiver `intoValue875`
	// (`intoValue875` is now tainted).
	intoValue875.TrySend(fromValue513)

	// Return the tainted `intoValue875`:
	return intoValue875
}

func TaintStepTest_ReflectValueUint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue447` into `intoUint64619`.

	// Assume that `sourceCQL` has the underlying type of `fromValue447`:
	fromValue447 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue447` to the result `intoUint64619`
	// (`intoUint64619` is now tainted).
	intoUint64619 := fromValue447.Uint()

	// Return the tainted `intoUint64619`:
	return intoUint64619
}

func TaintStepTest_ReflectValueUnsafeAddr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue794` into `intoUintptr827`.

	// Assume that `sourceCQL` has the underlying type of `fromValue794`:
	fromValue794 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue794` to the result `intoUintptr827`
	// (`intoUintptr827` is now tainted).
	intoUintptr827 := fromValue794.UnsafeAddr()

	// Return the tainted `intoUintptr827`:
	return intoUintptr827
}

func RunAllTaints_Reflect() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectAppend_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectAppend_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectAppendSlice_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectAppendSlice_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectCopy_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectIndirect_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueOf_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectMapIterKey_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectMapIterValue_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectStructTagGet_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectStructTagLookup_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueAddr_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueBool_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueBytes_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueComplex_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueConvert_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueElem_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueField_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueFieldByIndex_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueFieldByName_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueFieldByNameFunc_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueFloat_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueIndex_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueInt_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueInterface_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueInterfaceData_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueMapIndex_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueMapKeys_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueMapRange_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueMethod_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueMethodByName_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValuePointer_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueRecv_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueSend_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueSet_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueSetBool_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueSetBytes_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueSetComplex_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueSetFloat_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueSetInt_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueSetMapIndex_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueSetMapIndex_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueSetPointer_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueSetString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueSetUint_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueSlice_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueSlice3_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueTryRecv_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueTrySend_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueUint_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ReflectValueUnsafeAddr_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
