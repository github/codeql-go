package main

import (
	"reflect"
	"unsafe"
)

func TaintStepTest_ReflectAppend_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue995` into `intoValue401`.

	// Assume that `sourceCQL` has the underlying type of `fromValue995`:
	fromValue995 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue995` to result `intoValue401`
	// (`intoValue401` is now tainted).
	intoValue401 := reflect.Append(fromValue995, reflect.Value{})

	// Return the tainted `intoValue401`:
	return intoValue401
}

func TaintStepTest_ReflectAppend_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue371` into `intoValue111`.

	// Assume that `sourceCQL` has the underlying type of `fromValue371`:
	fromValue371 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue371` to result `intoValue111`
	// (`intoValue111` is now tainted).
	intoValue111 := reflect.Append(reflect.Value{}, fromValue371)

	// Return the tainted `intoValue111`:
	return intoValue111
}

func TaintStepTest_ReflectAppendSlice_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue207` into `intoValue714`.

	// Assume that `sourceCQL` has the underlying type of `fromValue207`:
	fromValue207 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue207` to result `intoValue714`
	// (`intoValue714` is now tainted).
	intoValue714 := reflect.AppendSlice(fromValue207, reflect.Value{})

	// Return the tainted `intoValue714`:
	return intoValue714
}

func TaintStepTest_ReflectAppendSlice_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue123` into `intoValue516`.

	// Assume that `sourceCQL` has the underlying type of `fromValue123`:
	fromValue123 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue123` to result `intoValue516`
	// (`intoValue516` is now tainted).
	intoValue516 := reflect.AppendSlice(reflect.Value{}, fromValue123)

	// Return the tainted `intoValue516`:
	return intoValue516
}

func TaintStepTest_ReflectCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue638` into `intoValue632`.

	// Assume that `sourceCQL` has the underlying type of `fromValue638`:
	fromValue638 := sourceCQL.(reflect.Value)

	// Declare `intoValue632` variable:
	var intoValue632 reflect.Value

	// Call the function that transfers the taint
	// from the parameter `fromValue638` to parameter `intoValue632`;
	// `intoValue632` is now tainted.
	reflect.Copy(intoValue632, fromValue638)

	// Return the tainted `intoValue632`:
	return intoValue632
}

func TaintStepTest_ReflectIndirect_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue249` into `intoValue871`.

	// Assume that `sourceCQL` has the underlying type of `fromValue249`:
	fromValue249 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue249` to result `intoValue871`
	// (`intoValue871` is now tainted).
	intoValue871 := reflect.Indirect(fromValue249)

	// Return the tainted `intoValue871`:
	return intoValue871
}

func TaintStepTest_ReflectValueOf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface900` into `intoValue517`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface900`:
	fromInterface900 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface900` to result `intoValue517`
	// (`intoValue517` is now tainted).
	intoValue517 := reflect.ValueOf(fromInterface900)

	// Return the tainted `intoValue517`:
	return intoValue517
}

func TaintStepTest_ReflectMapIterKey_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMapIter270` into `intoValue211`.

	// Assume that `sourceCQL` has the underlying type of `fromMapIter270`:
	fromMapIter270 := sourceCQL.(reflect.MapIter)

	// Call the method that transfers the taint
	// from the receiver `fromMapIter270` to the result `intoValue211`
	// (`intoValue211` is now tainted).
	intoValue211 := fromMapIter270.Key()

	// Return the tainted `intoValue211`:
	return intoValue211
}

func TaintStepTest_ReflectMapIterValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMapIter288` into `intoValue570`.

	// Assume that `sourceCQL` has the underlying type of `fromMapIter288`:
	fromMapIter288 := sourceCQL.(reflect.MapIter)

	// Call the method that transfers the taint
	// from the receiver `fromMapIter288` to the result `intoValue570`
	// (`intoValue570` is now tainted).
	intoValue570 := fromMapIter288.Value()

	// Return the tainted `intoValue570`:
	return intoValue570
}

func TaintStepTest_ReflectStructTagGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStructTag357` into `intoString363`.

	// Assume that `sourceCQL` has the underlying type of `fromStructTag357`:
	fromStructTag357 := sourceCQL.(reflect.StructTag)

	// Call the method that transfers the taint
	// from the receiver `fromStructTag357` to the result `intoString363`
	// (`intoString363` is now tainted).
	intoString363 := fromStructTag357.Get("")

	// Return the tainted `intoString363`:
	return intoString363
}

func TaintStepTest_ReflectStructTagLookup_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStructTag124` into `intoString883`.

	// Assume that `sourceCQL` has the underlying type of `fromStructTag124`:
	fromStructTag124 := sourceCQL.(reflect.StructTag)

	// Call the method that transfers the taint
	// from the receiver `fromStructTag124` to the result `intoString883`
	// (`intoString883` is now tainted).
	intoString883, _ := fromStructTag124.Lookup("")

	// Return the tainted `intoString883`:
	return intoString883
}

func TaintStepTest_ReflectValueAddr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue374` into `intoValue434`.

	// Assume that `sourceCQL` has the underlying type of `fromValue374`:
	fromValue374 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue374` to the result `intoValue434`
	// (`intoValue434` is now tainted).
	intoValue434 := fromValue374.Addr()

	// Return the tainted `intoValue434`:
	return intoValue434
}

func TaintStepTest_ReflectValueBool_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue692` into `intoBool739`.

	// Assume that `sourceCQL` has the underlying type of `fromValue692`:
	fromValue692 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue692` to the result `intoBool739`
	// (`intoBool739` is now tainted).
	intoBool739 := fromValue692.Bool()

	// Return the tainted `intoBool739`:
	return intoBool739
}

func TaintStepTest_ReflectValueBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue538` into `intoByte923`.

	// Assume that `sourceCQL` has the underlying type of `fromValue538`:
	fromValue538 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue538` to the result `intoByte923`
	// (`intoByte923` is now tainted).
	intoByte923 := fromValue538.Bytes()

	// Return the tainted `intoByte923`:
	return intoByte923
}

func TaintStepTest_ReflectValueComplex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue593` into `intoComplex128631`.

	// Assume that `sourceCQL` has the underlying type of `fromValue593`:
	fromValue593 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue593` to the result `intoComplex128631`
	// (`intoComplex128631` is now tainted).
	intoComplex128631 := fromValue593.Complex()

	// Return the tainted `intoComplex128631`:
	return intoComplex128631
}

func TaintStepTest_ReflectValueConvert_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue799` into `intoValue634`.

	// Assume that `sourceCQL` has the underlying type of `fromValue799`:
	fromValue799 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue799` to the result `intoValue634`
	// (`intoValue634` is now tainted).
	intoValue634 := fromValue799.Convert(nil)

	// Return the tainted `intoValue634`:
	return intoValue634
}

func TaintStepTest_ReflectValueElem_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue766` into `intoValue659`.

	// Assume that `sourceCQL` has the underlying type of `fromValue766`:
	fromValue766 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue766` to the result `intoValue659`
	// (`intoValue659` is now tainted).
	intoValue659 := fromValue766.Elem()

	// Return the tainted `intoValue659`:
	return intoValue659
}

func TaintStepTest_ReflectValueField_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue384` into `intoValue823`.

	// Assume that `sourceCQL` has the underlying type of `fromValue384`:
	fromValue384 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue384` to the result `intoValue823`
	// (`intoValue823` is now tainted).
	intoValue823 := fromValue384.Field(0)

	// Return the tainted `intoValue823`:
	return intoValue823
}

func TaintStepTest_ReflectValueFieldByIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue355` into `intoValue133`.

	// Assume that `sourceCQL` has the underlying type of `fromValue355`:
	fromValue355 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue355` to the result `intoValue133`
	// (`intoValue133` is now tainted).
	intoValue133 := fromValue355.FieldByIndex(nil)

	// Return the tainted `intoValue133`:
	return intoValue133
}

func TaintStepTest_ReflectValueFieldByName_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue703` into `intoValue595`.

	// Assume that `sourceCQL` has the underlying type of `fromValue703`:
	fromValue703 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue703` to the result `intoValue595`
	// (`intoValue595` is now tainted).
	intoValue595 := fromValue703.FieldByName("")

	// Return the tainted `intoValue595`:
	return intoValue595
}

func TaintStepTest_ReflectValueFieldByNameFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue216` into `intoValue926`.

	// Assume that `sourceCQL` has the underlying type of `fromValue216`:
	fromValue216 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue216` to the result `intoValue926`
	// (`intoValue926` is now tainted).
	intoValue926 := fromValue216.FieldByNameFunc(nil)

	// Return the tainted `intoValue926`:
	return intoValue926
}

func TaintStepTest_ReflectValueFloat_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue423` into `intoFloat64461`.

	// Assume that `sourceCQL` has the underlying type of `fromValue423`:
	fromValue423 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue423` to the result `intoFloat64461`
	// (`intoFloat64461` is now tainted).
	intoFloat64461 := fromValue423.Float()

	// Return the tainted `intoFloat64461`:
	return intoFloat64461
}

func TaintStepTest_ReflectValueIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue711` into `intoValue659`.

	// Assume that `sourceCQL` has the underlying type of `fromValue711`:
	fromValue711 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue711` to the result `intoValue659`
	// (`intoValue659` is now tainted).
	intoValue659 := fromValue711.Index(0)

	// Return the tainted `intoValue659`:
	return intoValue659
}

func TaintStepTest_ReflectValueInt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue724` into `intoInt64271`.

	// Assume that `sourceCQL` has the underlying type of `fromValue724`:
	fromValue724 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue724` to the result `intoInt64271`
	// (`intoInt64271` is now tainted).
	intoInt64271 := fromValue724.Int()

	// Return the tainted `intoInt64271`:
	return intoInt64271
}

func TaintStepTest_ReflectValueInterface_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue607` into `intoInterface579`.

	// Assume that `sourceCQL` has the underlying type of `fromValue607`:
	fromValue607 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue607` to the result `intoInterface579`
	// (`intoInterface579` is now tainted).
	intoInterface579 := fromValue607.Interface()

	// Return the tainted `intoInterface579`:
	return intoInterface579
}

func TaintStepTest_ReflectValueInterfaceData_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue944` into `intoUintptr931`.

	// Assume that `sourceCQL` has the underlying type of `fromValue944`:
	fromValue944 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue944` to the result `intoUintptr931`
	// (`intoUintptr931` is now tainted).
	intoUintptr931 := fromValue944.InterfaceData()

	// Return the tainted `intoUintptr931`:
	return intoUintptr931
}

func TaintStepTest_ReflectValueMapIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue452` into `intoValue977`.

	// Assume that `sourceCQL` has the underlying type of `fromValue452`:
	fromValue452 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue452` to the result `intoValue977`
	// (`intoValue977` is now tainted).
	intoValue977 := fromValue452.MapIndex(reflect.Value{})

	// Return the tainted `intoValue977`:
	return intoValue977
}

func TaintStepTest_ReflectValueMapKeys_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue324` into `intoValue813`.

	// Assume that `sourceCQL` has the underlying type of `fromValue324`:
	fromValue324 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue324` to the result `intoValue813`
	// (`intoValue813` is now tainted).
	intoValue813 := fromValue324.MapKeys()

	// Return the tainted `intoValue813`:
	return intoValue813
}

func TaintStepTest_ReflectValueMapRange_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue599` into `intoMapIter780`.

	// Assume that `sourceCQL` has the underlying type of `fromValue599`:
	fromValue599 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue599` to the result `intoMapIter780`
	// (`intoMapIter780` is now tainted).
	intoMapIter780 := fromValue599.MapRange()

	// Return the tainted `intoMapIter780`:
	return intoMapIter780
}

func TaintStepTest_ReflectValueMethod_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue181` into `intoValue267`.

	// Assume that `sourceCQL` has the underlying type of `fromValue181`:
	fromValue181 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue181` to the result `intoValue267`
	// (`intoValue267` is now tainted).
	intoValue267 := fromValue181.Method(0)

	// Return the tainted `intoValue267`:
	return intoValue267
}

func TaintStepTest_ReflectValueMethodByName_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue176` into `intoValue471`.

	// Assume that `sourceCQL` has the underlying type of `fromValue176`:
	fromValue176 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue176` to the result `intoValue471`
	// (`intoValue471` is now tainted).
	intoValue471 := fromValue176.MethodByName("")

	// Return the tainted `intoValue471`:
	return intoValue471
}

func TaintStepTest_ReflectValuePointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue580` into `intoUintptr720`.

	// Assume that `sourceCQL` has the underlying type of `fromValue580`:
	fromValue580 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue580` to the result `intoUintptr720`
	// (`intoUintptr720` is now tainted).
	intoUintptr720 := fromValue580.Pointer()

	// Return the tainted `intoUintptr720`:
	return intoUintptr720
}

func TaintStepTest_ReflectValueRecv_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue685` into `intoValue297`.

	// Assume that `sourceCQL` has the underlying type of `fromValue685`:
	fromValue685 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue685` to the result `intoValue297`
	// (`intoValue297` is now tainted).
	intoValue297, _ := fromValue685.Recv()

	// Return the tainted `intoValue297`:
	return intoValue297
}

func TaintStepTest_ReflectValueSend_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue256` into `intoValue933`.

	// Assume that `sourceCQL` has the underlying type of `fromValue256`:
	fromValue256 := sourceCQL.(reflect.Value)

	// Declare `intoValue933` variable:
	var intoValue933 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue256` to the receiver `intoValue933`
	// (`intoValue933` is now tainted).
	intoValue933.Send(fromValue256)

	// Return the tainted `intoValue933`:
	return intoValue933
}

func TaintStepTest_ReflectValueSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue708` into `intoValue639`.

	// Assume that `sourceCQL` has the underlying type of `fromValue708`:
	fromValue708 := sourceCQL.(reflect.Value)

	// Declare `intoValue639` variable:
	var intoValue639 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue708` to the receiver `intoValue639`
	// (`intoValue639` is now tainted).
	intoValue639.Set(fromValue708)

	// Return the tainted `intoValue639`:
	return intoValue639
}

func TaintStepTest_ReflectValueSetBool_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBool653` into `intoValue921`.

	// Assume that `sourceCQL` has the underlying type of `fromBool653`:
	fromBool653 := sourceCQL.(bool)

	// Declare `intoValue921` variable:
	var intoValue921 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromBool653` to the receiver `intoValue921`
	// (`intoValue921` is now tainted).
	intoValue921.SetBool(fromBool653)

	// Return the tainted `intoValue921`:
	return intoValue921
}

func TaintStepTest_ReflectValueSetBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte399` into `intoValue587`.

	// Assume that `sourceCQL` has the underlying type of `fromByte399`:
	fromByte399 := sourceCQL.([]byte)

	// Declare `intoValue587` variable:
	var intoValue587 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromByte399` to the receiver `intoValue587`
	// (`intoValue587` is now tainted).
	intoValue587.SetBytes(fromByte399)

	// Return the tainted `intoValue587`:
	return intoValue587
}

func TaintStepTest_ReflectValueSetComplex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromComplex128235` into `intoValue825`.

	// Assume that `sourceCQL` has the underlying type of `fromComplex128235`:
	fromComplex128235 := sourceCQL.(complex128)

	// Declare `intoValue825` variable:
	var intoValue825 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromComplex128235` to the receiver `intoValue825`
	// (`intoValue825` is now tainted).
	intoValue825.SetComplex(fromComplex128235)

	// Return the tainted `intoValue825`:
	return intoValue825
}

func TaintStepTest_ReflectValueSetFloat_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFloat64704` into `intoValue388`.

	// Assume that `sourceCQL` has the underlying type of `fromFloat64704`:
	fromFloat64704 := sourceCQL.(float64)

	// Declare `intoValue388` variable:
	var intoValue388 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromFloat64704` to the receiver `intoValue388`
	// (`intoValue388` is now tainted).
	intoValue388.SetFloat(fromFloat64704)

	// Return the tainted `intoValue388`:
	return intoValue388
}

func TaintStepTest_ReflectValueSetInt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInt64919` into `intoValue624`.

	// Assume that `sourceCQL` has the underlying type of `fromInt64919`:
	fromInt64919 := sourceCQL.(int64)

	// Declare `intoValue624` variable:
	var intoValue624 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromInt64919` to the receiver `intoValue624`
	// (`intoValue624` is now tainted).
	intoValue624.SetInt(fromInt64919)

	// Return the tainted `intoValue624`:
	return intoValue624
}

func TaintStepTest_ReflectValueSetMapIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue854` into `intoValue266`.

	// Assume that `sourceCQL` has the underlying type of `fromValue854`:
	fromValue854 := sourceCQL.(reflect.Value)

	// Declare `intoValue266` variable:
	var intoValue266 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue854` to the receiver `intoValue266`
	// (`intoValue266` is now tainted).
	intoValue266.SetMapIndex(fromValue854, reflect.Value{})

	// Return the tainted `intoValue266`:
	return intoValue266
}

func TaintStepTest_ReflectValueSetMapIndex_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue624` into `intoValue282`.

	// Assume that `sourceCQL` has the underlying type of `fromValue624`:
	fromValue624 := sourceCQL.(reflect.Value)

	// Declare `intoValue282` variable:
	var intoValue282 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue624` to the receiver `intoValue282`
	// (`intoValue282` is now tainted).
	intoValue282.SetMapIndex(reflect.Value{}, fromValue624)

	// Return the tainted `intoValue282`:
	return intoValue282
}

func TaintStepTest_ReflectValueSetPointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer799` into `intoValue775`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer799`:
	fromPointer799 := sourceCQL.(unsafe.Pointer)

	// Declare `intoValue775` variable:
	var intoValue775 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromPointer799` to the receiver `intoValue775`
	// (`intoValue775` is now tainted).
	intoValue775.SetPointer(fromPointer799)

	// Return the tainted `intoValue775`:
	return intoValue775
}

func TaintStepTest_ReflectValueSetString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString483` into `intoValue157`.

	// Assume that `sourceCQL` has the underlying type of `fromString483`:
	fromString483 := sourceCQL.(string)

	// Declare `intoValue157` variable:
	var intoValue157 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromString483` to the receiver `intoValue157`
	// (`intoValue157` is now tainted).
	intoValue157.SetString(fromString483)

	// Return the tainted `intoValue157`:
	return intoValue157
}

func TaintStepTest_ReflectValueSetUint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUint64217` into `intoValue457`.

	// Assume that `sourceCQL` has the underlying type of `fromUint64217`:
	fromUint64217 := sourceCQL.(uint64)

	// Declare `intoValue457` variable:
	var intoValue457 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromUint64217` to the receiver `intoValue457`
	// (`intoValue457` is now tainted).
	intoValue457.SetUint(fromUint64217)

	// Return the tainted `intoValue457`:
	return intoValue457
}

func TaintStepTest_ReflectValueSlice_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue578` into `intoValue695`.

	// Assume that `sourceCQL` has the underlying type of `fromValue578`:
	fromValue578 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue578` to the result `intoValue695`
	// (`intoValue695` is now tainted).
	intoValue695 := fromValue578.Slice(0, 0)

	// Return the tainted `intoValue695`:
	return intoValue695
}

func TaintStepTest_ReflectValueSlice3_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue924` into `intoValue767`.

	// Assume that `sourceCQL` has the underlying type of `fromValue924`:
	fromValue924 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue924` to the result `intoValue767`
	// (`intoValue767` is now tainted).
	intoValue767 := fromValue924.Slice3(0, 0, 0)

	// Return the tainted `intoValue767`:
	return intoValue767
}

func TaintStepTest_ReflectValueString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue796` into `intoString139`.

	// Assume that `sourceCQL` has the underlying type of `fromValue796`:
	fromValue796 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue796` to the result `intoString139`
	// (`intoString139` is now tainted).
	intoString139 := fromValue796.String()

	// Return the tainted `intoString139`:
	return intoString139
}

func TaintStepTest_ReflectValueTryRecv_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue707` into `intoValue699`.

	// Assume that `sourceCQL` has the underlying type of `fromValue707`:
	fromValue707 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue707` to the result `intoValue699`
	// (`intoValue699` is now tainted).
	intoValue699, _ := fromValue707.TryRecv()

	// Return the tainted `intoValue699`:
	return intoValue699
}

func TaintStepTest_ReflectValueTrySend_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue839` into `intoValue862`.

	// Assume that `sourceCQL` has the underlying type of `fromValue839`:
	fromValue839 := sourceCQL.(reflect.Value)

	// Declare `intoValue862` variable:
	var intoValue862 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue839` to the receiver `intoValue862`
	// (`intoValue862` is now tainted).
	intoValue862.TrySend(fromValue839)

	// Return the tainted `intoValue862`:
	return intoValue862
}

func TaintStepTest_ReflectValueUint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue561` into `intoUint64752`.

	// Assume that `sourceCQL` has the underlying type of `fromValue561`:
	fromValue561 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue561` to the result `intoUint64752`
	// (`intoUint64752` is now tainted).
	intoUint64752 := fromValue561.Uint()

	// Return the tainted `intoUint64752`:
	return intoUint64752
}

func TaintStepTest_ReflectValueUnsafeAddr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue220` into `intoUintptr608`.

	// Assume that `sourceCQL` has the underlying type of `fromValue220`:
	fromValue220 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue220` to the result `intoUintptr608`
	// (`intoUintptr608` is now tainted).
	intoUintptr608 := fromValue220.UnsafeAddr()

	// Return the tainted `intoUintptr608`:
	return intoUintptr608
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
