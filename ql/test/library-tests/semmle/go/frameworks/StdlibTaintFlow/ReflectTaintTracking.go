package main

import (
	"reflect"
	"unsafe"
)

func TaintStepTest_ReflectAppend_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue860` into `intoValue264`.

	// Assume that `sourceCQL` has the underlying type of `fromValue860`:
	fromValue860 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue860` to result `intoValue264`
	// (`intoValue264` is now tainted).
	intoValue264 := reflect.Append(fromValue860, reflect.Value{})

	// Sink the tainted `intoValue264`:
	sink(intoValue264)
}

func TaintStepTest_ReflectAppend_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromValue201` into `intoValue600`.

	// Assume that `sourceCQL` has the underlying type of `fromValue201`:
	fromValue201 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue201` to result `intoValue600`
	// (`intoValue600` is now tainted).
	intoValue600 := reflect.Append(reflect.Value{}, fromValue201)

	// Sink the tainted `intoValue600`:
	sink(intoValue600)
}

func TaintStepTest_ReflectAppendSlice_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue184` into `intoValue590`.

	// Assume that `sourceCQL` has the underlying type of `fromValue184`:
	fromValue184 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue184` to result `intoValue590`
	// (`intoValue590` is now tainted).
	intoValue590 := reflect.AppendSlice(fromValue184, reflect.Value{})

	// Sink the tainted `intoValue590`:
	sink(intoValue590)
}

func TaintStepTest_ReflectAppendSlice_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromValue174` into `intoValue574`.

	// Assume that `sourceCQL` has the underlying type of `fromValue174`:
	fromValue174 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue174` to result `intoValue574`
	// (`intoValue574` is now tainted).
	intoValue574 := reflect.AppendSlice(reflect.Value{}, fromValue174)

	// Sink the tainted `intoValue574`:
	sink(intoValue574)
}

func TaintStepTest_ReflectCopy_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue951` into `intoValue643`.

	// Assume that `sourceCQL` has the underlying type of `fromValue951`:
	fromValue951 := sourceCQL.(reflect.Value)

	// Declare `intoValue643` variable:
	var intoValue643 reflect.Value

	// Call the function that transfers the taint
	// from the parameter `fromValue951` to parameter `intoValue643`;
	// `intoValue643` is now tainted.
	reflect.Copy(intoValue643, fromValue951)

	// Sink the tainted `intoValue643`:
	sink(intoValue643)
}

func TaintStepTest_ReflectIndirect_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue228` into `intoValue774`.

	// Assume that `sourceCQL` has the underlying type of `fromValue228`:
	fromValue228 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue228` to result `intoValue774`
	// (`intoValue774` is now tainted).
	intoValue774 := reflect.Indirect(fromValue228)

	// Sink the tainted `intoValue774`:
	sink(intoValue774)
}

func TaintStepTest_ReflectValueOf_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface818` into `intoValue150`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface818`:
	fromInterface818 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface818` to result `intoValue150`
	// (`intoValue150` is now tainted).
	intoValue150 := reflect.ValueOf(fromInterface818)

	// Sink the tainted `intoValue150`:
	sink(intoValue150)
}

func TaintStepTest_ReflectMapIterKey_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromMapIter895` into `intoValue409`.

	// Assume that `sourceCQL` has the underlying type of `fromMapIter895`:
	fromMapIter895 := sourceCQL.(reflect.MapIter)

	// Call the method that transfers the taint
	// from the receiver `fromMapIter895` to the result `intoValue409`
	// (`intoValue409` is now tainted).
	intoValue409 := fromMapIter895.Key()

	// Sink the tainted `intoValue409`:
	sink(intoValue409)
}

func TaintStepTest_ReflectMapIterValue_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromMapIter968` into `intoValue925`.

	// Assume that `sourceCQL` has the underlying type of `fromMapIter968`:
	fromMapIter968 := sourceCQL.(reflect.MapIter)

	// Call the method that transfers the taint
	// from the receiver `fromMapIter968` to the result `intoValue925`
	// (`intoValue925` is now tainted).
	intoValue925 := fromMapIter968.Value()

	// Sink the tainted `intoValue925`:
	sink(intoValue925)
}

func TaintStepTest_ReflectStructTagGet_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromStructTag160` into `intoString360`.

	// Assume that `sourceCQL` has the underlying type of `fromStructTag160`:
	fromStructTag160 := sourceCQL.(reflect.StructTag)

	// Call the method that transfers the taint
	// from the receiver `fromStructTag160` to the result `intoString360`
	// (`intoString360` is now tainted).
	intoString360 := fromStructTag160.Get("")

	// Sink the tainted `intoString360`:
	sink(intoString360)
}

func TaintStepTest_ReflectStructTagLookup_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromStructTag302` into `intoString693`.

	// Assume that `sourceCQL` has the underlying type of `fromStructTag302`:
	fromStructTag302 := sourceCQL.(reflect.StructTag)

	// Call the method that transfers the taint
	// from the receiver `fromStructTag302` to the result `intoString693`
	// (`intoString693` is now tainted).
	intoString693, _ := fromStructTag302.Lookup("")

	// Sink the tainted `intoString693`:
	sink(intoString693)
}

func TaintStepTest_ReflectValueAddr_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue710` into `intoValue252`.

	// Assume that `sourceCQL` has the underlying type of `fromValue710`:
	fromValue710 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue710` to the result `intoValue252`
	// (`intoValue252` is now tainted).
	intoValue252 := fromValue710.Addr()

	// Sink the tainted `intoValue252`:
	sink(intoValue252)
}

func TaintStepTest_ReflectValueBool_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue160` into `intoBool951`.

	// Assume that `sourceCQL` has the underlying type of `fromValue160`:
	fromValue160 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue160` to the result `intoBool951`
	// (`intoBool951` is now tainted).
	intoBool951 := fromValue160.Bool()

	// Sink the tainted `intoBool951`:
	sink(intoBool951)
}

func TaintStepTest_ReflectValueBytes_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue744` into `intoByte718`.

	// Assume that `sourceCQL` has the underlying type of `fromValue744`:
	fromValue744 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue744` to the result `intoByte718`
	// (`intoByte718` is now tainted).
	intoByte718 := fromValue744.Bytes()

	// Sink the tainted `intoByte718`:
	sink(intoByte718)
}

func TaintStepTest_ReflectValueComplex_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue276` into `intoComplex128861`.

	// Assume that `sourceCQL` has the underlying type of `fromValue276`:
	fromValue276 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue276` to the result `intoComplex128861`
	// (`intoComplex128861` is now tainted).
	intoComplex128861 := fromValue276.Complex()

	// Sink the tainted `intoComplex128861`:
	sink(intoComplex128861)
}

func TaintStepTest_ReflectValueConvert_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue160` into `intoValue970`.

	// Assume that `sourceCQL` has the underlying type of `fromValue160`:
	fromValue160 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue160` to the result `intoValue970`
	// (`intoValue970` is now tainted).
	intoValue970 := fromValue160.Convert(nil)

	// Sink the tainted `intoValue970`:
	sink(intoValue970)
}

func TaintStepTest_ReflectValueElem_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue385` into `intoValue853`.

	// Assume that `sourceCQL` has the underlying type of `fromValue385`:
	fromValue385 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue385` to the result `intoValue853`
	// (`intoValue853` is now tainted).
	intoValue853 := fromValue385.Elem()

	// Sink the tainted `intoValue853`:
	sink(intoValue853)
}

func TaintStepTest_ReflectValueField_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue993` into `intoValue156`.

	// Assume that `sourceCQL` has the underlying type of `fromValue993`:
	fromValue993 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue993` to the result `intoValue156`
	// (`intoValue156` is now tainted).
	intoValue156 := fromValue993.Field(0)

	// Sink the tainted `intoValue156`:
	sink(intoValue156)
}

func TaintStepTest_ReflectValueFieldByIndex_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue194` into `intoValue750`.

	// Assume that `sourceCQL` has the underlying type of `fromValue194`:
	fromValue194 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue194` to the result `intoValue750`
	// (`intoValue750` is now tainted).
	intoValue750 := fromValue194.FieldByIndex(nil)

	// Sink the tainted `intoValue750`:
	sink(intoValue750)
}

func TaintStepTest_ReflectValueFieldByName_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue874` into `intoValue229`.

	// Assume that `sourceCQL` has the underlying type of `fromValue874`:
	fromValue874 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue874` to the result `intoValue229`
	// (`intoValue229` is now tainted).
	intoValue229 := fromValue874.FieldByName("")

	// Sink the tainted `intoValue229`:
	sink(intoValue229)
}

func TaintStepTest_ReflectValueFieldByNameFunc_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue253` into `intoValue368`.

	// Assume that `sourceCQL` has the underlying type of `fromValue253`:
	fromValue253 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue253` to the result `intoValue368`
	// (`intoValue368` is now tainted).
	intoValue368 := fromValue253.FieldByNameFunc(nil)

	// Sink the tainted `intoValue368`:
	sink(intoValue368)
}

func TaintStepTest_ReflectValueFloat_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue711` into `intoFloat64726`.

	// Assume that `sourceCQL` has the underlying type of `fromValue711`:
	fromValue711 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue711` to the result `intoFloat64726`
	// (`intoFloat64726` is now tainted).
	intoFloat64726 := fromValue711.Float()

	// Sink the tainted `intoFloat64726`:
	sink(intoFloat64726)
}

func TaintStepTest_ReflectValueIndex_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue522` into `intoValue816`.

	// Assume that `sourceCQL` has the underlying type of `fromValue522`:
	fromValue522 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue522` to the result `intoValue816`
	// (`intoValue816` is now tainted).
	intoValue816 := fromValue522.Index(0)

	// Sink the tainted `intoValue816`:
	sink(intoValue816)
}

func TaintStepTest_ReflectValueInt_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue427` into `intoInt64621`.

	// Assume that `sourceCQL` has the underlying type of `fromValue427`:
	fromValue427 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue427` to the result `intoInt64621`
	// (`intoInt64621` is now tainted).
	intoInt64621 := fromValue427.Int()

	// Sink the tainted `intoInt64621`:
	sink(intoInt64621)
}

func TaintStepTest_ReflectValueInterface_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue130` into `intoInterface753`.

	// Assume that `sourceCQL` has the underlying type of `fromValue130`:
	fromValue130 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue130` to the result `intoInterface753`
	// (`intoInterface753` is now tainted).
	intoInterface753 := fromValue130.Interface()

	// Sink the tainted `intoInterface753`:
	sink(intoInterface753)
}

func TaintStepTest_ReflectValueInterfaceData_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue975` into `intoUintptr794`.

	// Assume that `sourceCQL` has the underlying type of `fromValue975`:
	fromValue975 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue975` to the result `intoUintptr794`
	// (`intoUintptr794` is now tainted).
	intoUintptr794 := fromValue975.InterfaceData()

	// Sink the tainted `intoUintptr794`:
	sink(intoUintptr794)
}

func TaintStepTest_ReflectValueMapIndex_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue956` into `intoValue340`.

	// Assume that `sourceCQL` has the underlying type of `fromValue956`:
	fromValue956 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue956` to the result `intoValue340`
	// (`intoValue340` is now tainted).
	intoValue340 := fromValue956.MapIndex(reflect.Value{})

	// Sink the tainted `intoValue340`:
	sink(intoValue340)
}

func TaintStepTest_ReflectValueMapKeys_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue484` into `intoValue273`.

	// Assume that `sourceCQL` has the underlying type of `fromValue484`:
	fromValue484 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue484` to the result `intoValue273`
	// (`intoValue273` is now tainted).
	intoValue273 := fromValue484.MapKeys()

	// Sink the tainted `intoValue273`:
	sink(intoValue273)
}

func TaintStepTest_ReflectValueMapRange_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue584` into `intoMapIter471`.

	// Assume that `sourceCQL` has the underlying type of `fromValue584`:
	fromValue584 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue584` to the result `intoMapIter471`
	// (`intoMapIter471` is now tainted).
	intoMapIter471 := fromValue584.MapRange()

	// Sink the tainted `intoMapIter471`:
	sink(intoMapIter471)
}

func TaintStepTest_ReflectValueMethod_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue847` into `intoValue628`.

	// Assume that `sourceCQL` has the underlying type of `fromValue847`:
	fromValue847 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue847` to the result `intoValue628`
	// (`intoValue628` is now tainted).
	intoValue628 := fromValue847.Method(0)

	// Sink the tainted `intoValue628`:
	sink(intoValue628)
}

func TaintStepTest_ReflectValueMethodByName_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue194` into `intoValue400`.

	// Assume that `sourceCQL` has the underlying type of `fromValue194`:
	fromValue194 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue194` to the result `intoValue400`
	// (`intoValue400` is now tainted).
	intoValue400 := fromValue194.MethodByName("")

	// Sink the tainted `intoValue400`:
	sink(intoValue400)
}

func TaintStepTest_ReflectValuePointer_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue295` into `intoUintptr437`.

	// Assume that `sourceCQL` has the underlying type of `fromValue295`:
	fromValue295 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue295` to the result `intoUintptr437`
	// (`intoUintptr437` is now tainted).
	intoUintptr437 := fromValue295.Pointer()

	// Sink the tainted `intoUintptr437`:
	sink(intoUintptr437)
}

func TaintStepTest_ReflectValueRecv_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue839` into `intoValue792`.

	// Assume that `sourceCQL` has the underlying type of `fromValue839`:
	fromValue839 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue839` to the result `intoValue792`
	// (`intoValue792` is now tainted).
	intoValue792, _ := fromValue839.Recv()

	// Sink the tainted `intoValue792`:
	sink(intoValue792)
}

func TaintStepTest_ReflectValueSend_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue339` into `intoValue761`.

	// Assume that `sourceCQL` has the underlying type of `fromValue339`:
	fromValue339 := sourceCQL.(reflect.Value)

	// Declare `intoValue761` variable:
	var intoValue761 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue339` to the receiver `intoValue761`
	// (`intoValue761` is now tainted).
	intoValue761.Send(fromValue339)

	// Sink the tainted `intoValue761`:
	sink(intoValue761)
}

func TaintStepTest_ReflectValueSet_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue634` into `intoValue745`.

	// Assume that `sourceCQL` has the underlying type of `fromValue634`:
	fromValue634 := sourceCQL.(reflect.Value)

	// Declare `intoValue745` variable:
	var intoValue745 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue634` to the receiver `intoValue745`
	// (`intoValue745` is now tainted).
	intoValue745.Set(fromValue634)

	// Sink the tainted `intoValue745`:
	sink(intoValue745)
}

func TaintStepTest_ReflectValueSetBool_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBool494` into `intoValue800`.

	// Assume that `sourceCQL` has the underlying type of `fromBool494`:
	fromBool494 := sourceCQL.(bool)

	// Declare `intoValue800` variable:
	var intoValue800 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromBool494` to the receiver `intoValue800`
	// (`intoValue800` is now tainted).
	intoValue800.SetBool(fromBool494)

	// Sink the tainted `intoValue800`:
	sink(intoValue800)
}

func TaintStepTest_ReflectValueSetBytes_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte691` into `intoValue123`.

	// Assume that `sourceCQL` has the underlying type of `fromByte691`:
	fromByte691 := sourceCQL.([]byte)

	// Declare `intoValue123` variable:
	var intoValue123 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromByte691` to the receiver `intoValue123`
	// (`intoValue123` is now tainted).
	intoValue123.SetBytes(fromByte691)

	// Sink the tainted `intoValue123`:
	sink(intoValue123)
}

func TaintStepTest_ReflectValueSetComplex_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromComplex128924` into `intoValue882`.

	// Assume that `sourceCQL` has the underlying type of `fromComplex128924`:
	fromComplex128924 := sourceCQL.(complex128)

	// Declare `intoValue882` variable:
	var intoValue882 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromComplex128924` to the receiver `intoValue882`
	// (`intoValue882` is now tainted).
	intoValue882.SetComplex(fromComplex128924)

	// Sink the tainted `intoValue882`:
	sink(intoValue882)
}

func TaintStepTest_ReflectValueSetFloat_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFloat64918` into `intoValue728`.

	// Assume that `sourceCQL` has the underlying type of `fromFloat64918`:
	fromFloat64918 := sourceCQL.(float64)

	// Declare `intoValue728` variable:
	var intoValue728 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromFloat64918` to the receiver `intoValue728`
	// (`intoValue728` is now tainted).
	intoValue728.SetFloat(fromFloat64918)

	// Sink the tainted `intoValue728`:
	sink(intoValue728)
}

func TaintStepTest_ReflectValueSetInt_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInt64636` into `intoValue141`.

	// Assume that `sourceCQL` has the underlying type of `fromInt64636`:
	fromInt64636 := sourceCQL.(int64)

	// Declare `intoValue141` variable:
	var intoValue141 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromInt64636` to the receiver `intoValue141`
	// (`intoValue141` is now tainted).
	intoValue141.SetInt(fromInt64636)

	// Sink the tainted `intoValue141`:
	sink(intoValue141)
}

func TaintStepTest_ReflectValueSetMapIndex_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue952` into `intoValue498`.

	// Assume that `sourceCQL` has the underlying type of `fromValue952`:
	fromValue952 := sourceCQL.(reflect.Value)

	// Declare `intoValue498` variable:
	var intoValue498 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue952` to the receiver `intoValue498`
	// (`intoValue498` is now tainted).
	intoValue498.SetMapIndex(fromValue952, reflect.Value{})

	// Sink the tainted `intoValue498`:
	sink(intoValue498)
}

func TaintStepTest_ReflectValueSetMapIndex_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromValue123` into `intoValue966`.

	// Assume that `sourceCQL` has the underlying type of `fromValue123`:
	fromValue123 := sourceCQL.(reflect.Value)

	// Declare `intoValue966` variable:
	var intoValue966 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue123` to the receiver `intoValue966`
	// (`intoValue966` is now tainted).
	intoValue966.SetMapIndex(reflect.Value{}, fromValue123)

	// Sink the tainted `intoValue966`:
	sink(intoValue966)
}

func TaintStepTest_ReflectValueSetPointer_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromPointer584` into `intoValue671`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer584`:
	fromPointer584 := sourceCQL.(unsafe.Pointer)

	// Declare `intoValue671` variable:
	var intoValue671 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromPointer584` to the receiver `intoValue671`
	// (`intoValue671` is now tainted).
	intoValue671.SetPointer(fromPointer584)

	// Sink the tainted `intoValue671`:
	sink(intoValue671)
}

func TaintStepTest_ReflectValueSetString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString354` into `intoValue771`.

	// Assume that `sourceCQL` has the underlying type of `fromString354`:
	fromString354 := sourceCQL.(string)

	// Declare `intoValue771` variable:
	var intoValue771 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromString354` to the receiver `intoValue771`
	// (`intoValue771` is now tainted).
	intoValue771.SetString(fromString354)

	// Sink the tainted `intoValue771`:
	sink(intoValue771)
}

func TaintStepTest_ReflectValueSetUint_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUint64610` into `intoValue446`.

	// Assume that `sourceCQL` has the underlying type of `fromUint64610`:
	fromUint64610 := sourceCQL.(uint64)

	// Declare `intoValue446` variable:
	var intoValue446 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromUint64610` to the receiver `intoValue446`
	// (`intoValue446` is now tainted).
	intoValue446.SetUint(fromUint64610)

	// Sink the tainted `intoValue446`:
	sink(intoValue446)
}

func TaintStepTest_ReflectValueSlice_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue234` into `intoValue364`.

	// Assume that `sourceCQL` has the underlying type of `fromValue234`:
	fromValue234 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue234` to the result `intoValue364`
	// (`intoValue364` is now tainted).
	intoValue364 := fromValue234.Slice(0, 0)

	// Sink the tainted `intoValue364`:
	sink(intoValue364)
}

func TaintStepTest_ReflectValueSlice3_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue985` into `intoValue518`.

	// Assume that `sourceCQL` has the underlying type of `fromValue985`:
	fromValue985 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue985` to the result `intoValue518`
	// (`intoValue518` is now tainted).
	intoValue518 := fromValue985.Slice3(0, 0, 0)

	// Sink the tainted `intoValue518`:
	sink(intoValue518)
}

func TaintStepTest_ReflectValueString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue608` into `intoString489`.

	// Assume that `sourceCQL` has the underlying type of `fromValue608`:
	fromValue608 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue608` to the result `intoString489`
	// (`intoString489` is now tainted).
	intoString489 := fromValue608.String()

	// Sink the tainted `intoString489`:
	sink(intoString489)
}

func TaintStepTest_ReflectValueTryRecv_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue485` into `intoValue717`.

	// Assume that `sourceCQL` has the underlying type of `fromValue485`:
	fromValue485 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue485` to the result `intoValue717`
	// (`intoValue717` is now tainted).
	intoValue717, _ := fromValue485.TryRecv()

	// Sink the tainted `intoValue717`:
	sink(intoValue717)
}

func TaintStepTest_ReflectValueTrySend_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue678` into `intoValue142`.

	// Assume that `sourceCQL` has the underlying type of `fromValue678`:
	fromValue678 := sourceCQL.(reflect.Value)

	// Declare `intoValue142` variable:
	var intoValue142 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue678` to the receiver `intoValue142`
	// (`intoValue142` is now tainted).
	intoValue142.TrySend(fromValue678)

	// Sink the tainted `intoValue142`:
	sink(intoValue142)
}

func TaintStepTest_ReflectValueUint_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue127` into `intoUint64857`.

	// Assume that `sourceCQL` has the underlying type of `fromValue127`:
	fromValue127 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue127` to the result `intoUint64857`
	// (`intoUint64857` is now tainted).
	intoUint64857 := fromValue127.Uint()

	// Sink the tainted `intoUint64857`:
	sink(intoUint64857)
}

func TaintStepTest_ReflectValueUnsafeAddr_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValue731` into `intoUintptr606`.

	// Assume that `sourceCQL` has the underlying type of `fromValue731`:
	fromValue731 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue731` to the result `intoUintptr606`
	// (`intoUintptr606` is now tainted).
	intoUintptr606 := fromValue731.UnsafeAddr()

	// Sink the tainted `intoUintptr606`:
	sink(intoUintptr606)
}

func RunAllTaints_Reflect() {
	{
		source := newSource()
		TaintStepTest_ReflectAppend_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectAppend_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectAppendSlice_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectAppendSlice_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectCopy_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectIndirect_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueOf_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectMapIterKey_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectMapIterValue_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectStructTagGet_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectStructTagLookup_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueAddr_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueBool_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueBytes_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueComplex_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueConvert_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueElem_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueField_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueFieldByIndex_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueFieldByName_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueFieldByNameFunc_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueFloat_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueIndex_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueInt_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueInterface_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueInterfaceData_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueMapIndex_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueMapKeys_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueMapRange_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueMethod_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueMethodByName_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValuePointer_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueRecv_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSend_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSet_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetBool_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetBytes_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetComplex_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetFloat_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetInt_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetMapIndex_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetMapIndex_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetPointer_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetUint_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSlice_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSlice3_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueTryRecv_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueTrySend_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueUint_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueUnsafeAddr_B0I0O0(source)
	}
}
