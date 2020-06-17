// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"reflect"
	"unsafe"
)

func TaintStepTest_ReflectAppend_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue656` into `intoValue414`.

	// Assume that `sourceCQL` has the underlying type of `fromValue656`:
	fromValue656 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue656` to result `intoValue414`
	// (`intoValue414` is now tainted).
	intoValue414 := reflect.Append(fromValue656, reflect.Value{})

	// Return the tainted `intoValue414`:
	return intoValue414
}

func TaintStepTest_ReflectAppend_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue518` into `intoValue650`.

	// Assume that `sourceCQL` has the underlying type of `fromValue518`:
	fromValue518 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue518` to result `intoValue650`
	// (`intoValue650` is now tainted).
	intoValue650 := reflect.Append(reflect.Value{}, fromValue518)

	// Return the tainted `intoValue650`:
	return intoValue650
}

func TaintStepTest_ReflectAppendSlice_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue784` into `intoValue957`.

	// Assume that `sourceCQL` has the underlying type of `fromValue784`:
	fromValue784 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue784` to result `intoValue957`
	// (`intoValue957` is now tainted).
	intoValue957 := reflect.AppendSlice(fromValue784, reflect.Value{})

	// Return the tainted `intoValue957`:
	return intoValue957
}

func TaintStepTest_ReflectAppendSlice_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue520` into `intoValue443`.

	// Assume that `sourceCQL` has the underlying type of `fromValue520`:
	fromValue520 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue520` to result `intoValue443`
	// (`intoValue443` is now tainted).
	intoValue443 := reflect.AppendSlice(reflect.Value{}, fromValue520)

	// Return the tainted `intoValue443`:
	return intoValue443
}

func TaintStepTest_ReflectCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue127` into `intoValue483`.

	// Assume that `sourceCQL` has the underlying type of `fromValue127`:
	fromValue127 := sourceCQL.(reflect.Value)

	// Declare `intoValue483` variable:
	var intoValue483 reflect.Value

	// Call the function that transfers the taint
	// from the parameter `fromValue127` to parameter `intoValue483`;
	// `intoValue483` is now tainted.
	reflect.Copy(intoValue483, fromValue127)

	// Return the tainted `intoValue483`:
	return intoValue483
}

func TaintStepTest_ReflectIndirect_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue989` into `intoValue982`.

	// Assume that `sourceCQL` has the underlying type of `fromValue989`:
	fromValue989 := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `fromValue989` to result `intoValue982`
	// (`intoValue982` is now tainted).
	intoValue982 := reflect.Indirect(fromValue989)

	// Return the tainted `intoValue982`:
	return intoValue982
}

func TaintStepTest_ReflectValueOf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface417` into `intoValue584`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface417`:
	fromInterface417 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface417` to result `intoValue584`
	// (`intoValue584` is now tainted).
	intoValue584 := reflect.ValueOf(fromInterface417)

	// Return the tainted `intoValue584`:
	return intoValue584
}

func TaintStepTest_ReflectMapIterKey_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMapIter991` into `intoValue881`.

	// Assume that `sourceCQL` has the underlying type of `fromMapIter991`:
	fromMapIter991 := sourceCQL.(reflect.MapIter)

	// Call the method that transfers the taint
	// from the receiver `fromMapIter991` to the result `intoValue881`
	// (`intoValue881` is now tainted).
	intoValue881 := fromMapIter991.Key()

	// Return the tainted `intoValue881`:
	return intoValue881
}

func TaintStepTest_ReflectMapIterValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMapIter186` into `intoValue284`.

	// Assume that `sourceCQL` has the underlying type of `fromMapIter186`:
	fromMapIter186 := sourceCQL.(reflect.MapIter)

	// Call the method that transfers the taint
	// from the receiver `fromMapIter186` to the result `intoValue284`
	// (`intoValue284` is now tainted).
	intoValue284 := fromMapIter186.Value()

	// Return the tainted `intoValue284`:
	return intoValue284
}

func TaintStepTest_ReflectStructTagGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStructTag908` into `intoString137`.

	// Assume that `sourceCQL` has the underlying type of `fromStructTag908`:
	fromStructTag908 := sourceCQL.(reflect.StructTag)

	// Call the method that transfers the taint
	// from the receiver `fromStructTag908` to the result `intoString137`
	// (`intoString137` is now tainted).
	intoString137 := fromStructTag908.Get("")

	// Return the tainted `intoString137`:
	return intoString137
}

func TaintStepTest_ReflectStructTagLookup_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStructTag494` into `intoString873`.

	// Assume that `sourceCQL` has the underlying type of `fromStructTag494`:
	fromStructTag494 := sourceCQL.(reflect.StructTag)

	// Call the method that transfers the taint
	// from the receiver `fromStructTag494` to the result `intoString873`
	// (`intoString873` is now tainted).
	intoString873, _ := fromStructTag494.Lookup("")

	// Return the tainted `intoString873`:
	return intoString873
}

func TaintStepTest_ReflectValueAddr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue599` into `intoValue409`.

	// Assume that `sourceCQL` has the underlying type of `fromValue599`:
	fromValue599 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue599` to the result `intoValue409`
	// (`intoValue409` is now tainted).
	intoValue409 := fromValue599.Addr()

	// Return the tainted `intoValue409`:
	return intoValue409
}

func TaintStepTest_ReflectValueBool_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue246` into `intoBool898`.

	// Assume that `sourceCQL` has the underlying type of `fromValue246`:
	fromValue246 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue246` to the result `intoBool898`
	// (`intoBool898` is now tainted).
	intoBool898 := fromValue246.Bool()

	// Return the tainted `intoBool898`:
	return intoBool898
}

func TaintStepTest_ReflectValueBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue598` into `intoByte631`.

	// Assume that `sourceCQL` has the underlying type of `fromValue598`:
	fromValue598 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue598` to the result `intoByte631`
	// (`intoByte631` is now tainted).
	intoByte631 := fromValue598.Bytes()

	// Return the tainted `intoByte631`:
	return intoByte631
}

func TaintStepTest_ReflectValueComplex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue165` into `intoComplex128150`.

	// Assume that `sourceCQL` has the underlying type of `fromValue165`:
	fromValue165 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue165` to the result `intoComplex128150`
	// (`intoComplex128150` is now tainted).
	intoComplex128150 := fromValue165.Complex()

	// Return the tainted `intoComplex128150`:
	return intoComplex128150
}

func TaintStepTest_ReflectValueConvert_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue340` into `intoValue471`.

	// Assume that `sourceCQL` has the underlying type of `fromValue340`:
	fromValue340 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue340` to the result `intoValue471`
	// (`intoValue471` is now tainted).
	intoValue471 := fromValue340.Convert(nil)

	// Return the tainted `intoValue471`:
	return intoValue471
}

func TaintStepTest_ReflectValueElem_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue290` into `intoValue758`.

	// Assume that `sourceCQL` has the underlying type of `fromValue290`:
	fromValue290 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue290` to the result `intoValue758`
	// (`intoValue758` is now tainted).
	intoValue758 := fromValue290.Elem()

	// Return the tainted `intoValue758`:
	return intoValue758
}

func TaintStepTest_ReflectValueField_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue396` into `intoValue707`.

	// Assume that `sourceCQL` has the underlying type of `fromValue396`:
	fromValue396 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue396` to the result `intoValue707`
	// (`intoValue707` is now tainted).
	intoValue707 := fromValue396.Field(0)

	// Return the tainted `intoValue707`:
	return intoValue707
}

func TaintStepTest_ReflectValueFieldByIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue912` into `intoValue718`.

	// Assume that `sourceCQL` has the underlying type of `fromValue912`:
	fromValue912 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue912` to the result `intoValue718`
	// (`intoValue718` is now tainted).
	intoValue718 := fromValue912.FieldByIndex(nil)

	// Return the tainted `intoValue718`:
	return intoValue718
}

func TaintStepTest_ReflectValueFieldByName_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue972` into `intoValue633`.

	// Assume that `sourceCQL` has the underlying type of `fromValue972`:
	fromValue972 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue972` to the result `intoValue633`
	// (`intoValue633` is now tainted).
	intoValue633 := fromValue972.FieldByName("")

	// Return the tainted `intoValue633`:
	return intoValue633
}

func TaintStepTest_ReflectValueFieldByNameFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue316` into `intoValue145`.

	// Assume that `sourceCQL` has the underlying type of `fromValue316`:
	fromValue316 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue316` to the result `intoValue145`
	// (`intoValue145` is now tainted).
	intoValue145 := fromValue316.FieldByNameFunc(nil)

	// Return the tainted `intoValue145`:
	return intoValue145
}

func TaintStepTest_ReflectValueFloat_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue817` into `intoFloat64474`.

	// Assume that `sourceCQL` has the underlying type of `fromValue817`:
	fromValue817 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue817` to the result `intoFloat64474`
	// (`intoFloat64474` is now tainted).
	intoFloat64474 := fromValue817.Float()

	// Return the tainted `intoFloat64474`:
	return intoFloat64474
}

func TaintStepTest_ReflectValueIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue832` into `intoValue378`.

	// Assume that `sourceCQL` has the underlying type of `fromValue832`:
	fromValue832 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue832` to the result `intoValue378`
	// (`intoValue378` is now tainted).
	intoValue378 := fromValue832.Index(0)

	// Return the tainted `intoValue378`:
	return intoValue378
}

func TaintStepTest_ReflectValueInt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue541` into `intoInt64139`.

	// Assume that `sourceCQL` has the underlying type of `fromValue541`:
	fromValue541 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue541` to the result `intoInt64139`
	// (`intoInt64139` is now tainted).
	intoInt64139 := fromValue541.Int()

	// Return the tainted `intoInt64139`:
	return intoInt64139
}

func TaintStepTest_ReflectValueInterface_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue814` into `intoInterface768`.

	// Assume that `sourceCQL` has the underlying type of `fromValue814`:
	fromValue814 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue814` to the result `intoInterface768`
	// (`intoInterface768` is now tainted).
	intoInterface768 := fromValue814.Interface()

	// Return the tainted `intoInterface768`:
	return intoInterface768
}

func TaintStepTest_ReflectValueInterfaceData_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue468` into `intoUintptr736`.

	// Assume that `sourceCQL` has the underlying type of `fromValue468`:
	fromValue468 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue468` to the result `intoUintptr736`
	// (`intoUintptr736` is now tainted).
	intoUintptr736 := fromValue468.InterfaceData()

	// Return the tainted `intoUintptr736`:
	return intoUintptr736
}

func TaintStepTest_ReflectValueMapIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue516` into `intoValue246`.

	// Assume that `sourceCQL` has the underlying type of `fromValue516`:
	fromValue516 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue516` to the result `intoValue246`
	// (`intoValue246` is now tainted).
	intoValue246 := fromValue516.MapIndex(reflect.Value{})

	// Return the tainted `intoValue246`:
	return intoValue246
}

func TaintStepTest_ReflectValueMapKeys_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue679` into `intoValue736`.

	// Assume that `sourceCQL` has the underlying type of `fromValue679`:
	fromValue679 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue679` to the result `intoValue736`
	// (`intoValue736` is now tainted).
	intoValue736 := fromValue679.MapKeys()

	// Return the tainted `intoValue736`:
	return intoValue736
}

func TaintStepTest_ReflectValueMapRange_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue839` into `intoMapIter273`.

	// Assume that `sourceCQL` has the underlying type of `fromValue839`:
	fromValue839 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue839` to the result `intoMapIter273`
	// (`intoMapIter273` is now tainted).
	intoMapIter273 := fromValue839.MapRange()

	// Return the tainted `intoMapIter273`:
	return intoMapIter273
}

func TaintStepTest_ReflectValueMethod_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue982` into `intoValue458`.

	// Assume that `sourceCQL` has the underlying type of `fromValue982`:
	fromValue982 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue982` to the result `intoValue458`
	// (`intoValue458` is now tainted).
	intoValue458 := fromValue982.Method(0)

	// Return the tainted `intoValue458`:
	return intoValue458
}

func TaintStepTest_ReflectValueMethodByName_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue506` into `intoValue213`.

	// Assume that `sourceCQL` has the underlying type of `fromValue506`:
	fromValue506 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue506` to the result `intoValue213`
	// (`intoValue213` is now tainted).
	intoValue213 := fromValue506.MethodByName("")

	// Return the tainted `intoValue213`:
	return intoValue213
}

func TaintStepTest_ReflectValuePointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue468` into `intoUintptr219`.

	// Assume that `sourceCQL` has the underlying type of `fromValue468`:
	fromValue468 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue468` to the result `intoUintptr219`
	// (`intoUintptr219` is now tainted).
	intoUintptr219 := fromValue468.Pointer()

	// Return the tainted `intoUintptr219`:
	return intoUintptr219
}

func TaintStepTest_ReflectValueRecv_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue265` into `intoValue971`.

	// Assume that `sourceCQL` has the underlying type of `fromValue265`:
	fromValue265 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue265` to the result `intoValue971`
	// (`intoValue971` is now tainted).
	intoValue971, _ := fromValue265.Recv()

	// Return the tainted `intoValue971`:
	return intoValue971
}

func TaintStepTest_ReflectValueSend_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue320` into `intoValue545`.

	// Assume that `sourceCQL` has the underlying type of `fromValue320`:
	fromValue320 := sourceCQL.(reflect.Value)

	// Declare `intoValue545` variable:
	var intoValue545 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue320` to the receiver `intoValue545`
	// (`intoValue545` is now tainted).
	intoValue545.Send(fromValue320)

	// Return the tainted `intoValue545`:
	return intoValue545
}

func TaintStepTest_ReflectValueSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue566` into `intoValue497`.

	// Assume that `sourceCQL` has the underlying type of `fromValue566`:
	fromValue566 := sourceCQL.(reflect.Value)

	// Declare `intoValue497` variable:
	var intoValue497 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue566` to the receiver `intoValue497`
	// (`intoValue497` is now tainted).
	intoValue497.Set(fromValue566)

	// Return the tainted `intoValue497`:
	return intoValue497
}

func TaintStepTest_ReflectValueSetBool_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBool274` into `intoValue783`.

	// Assume that `sourceCQL` has the underlying type of `fromBool274`:
	fromBool274 := sourceCQL.(bool)

	// Declare `intoValue783` variable:
	var intoValue783 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromBool274` to the receiver `intoValue783`
	// (`intoValue783` is now tainted).
	intoValue783.SetBool(fromBool274)

	// Return the tainted `intoValue783`:
	return intoValue783
}

func TaintStepTest_ReflectValueSetBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte905` into `intoValue389`.

	// Assume that `sourceCQL` has the underlying type of `fromByte905`:
	fromByte905 := sourceCQL.([]byte)

	// Declare `intoValue389` variable:
	var intoValue389 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromByte905` to the receiver `intoValue389`
	// (`intoValue389` is now tainted).
	intoValue389.SetBytes(fromByte905)

	// Return the tainted `intoValue389`:
	return intoValue389
}

func TaintStepTest_ReflectValueSetComplex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromComplex128198` into `intoValue477`.

	// Assume that `sourceCQL` has the underlying type of `fromComplex128198`:
	fromComplex128198 := sourceCQL.(complex128)

	// Declare `intoValue477` variable:
	var intoValue477 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromComplex128198` to the receiver `intoValue477`
	// (`intoValue477` is now tainted).
	intoValue477.SetComplex(fromComplex128198)

	// Return the tainted `intoValue477`:
	return intoValue477
}

func TaintStepTest_ReflectValueSetFloat_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFloat64544` into `intoValue382`.

	// Assume that `sourceCQL` has the underlying type of `fromFloat64544`:
	fromFloat64544 := sourceCQL.(float64)

	// Declare `intoValue382` variable:
	var intoValue382 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromFloat64544` to the receiver `intoValue382`
	// (`intoValue382` is now tainted).
	intoValue382.SetFloat(fromFloat64544)

	// Return the tainted `intoValue382`:
	return intoValue382
}

func TaintStepTest_ReflectValueSetInt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInt64715` into `intoValue179`.

	// Assume that `sourceCQL` has the underlying type of `fromInt64715`:
	fromInt64715 := sourceCQL.(int64)

	// Declare `intoValue179` variable:
	var intoValue179 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromInt64715` to the receiver `intoValue179`
	// (`intoValue179` is now tainted).
	intoValue179.SetInt(fromInt64715)

	// Return the tainted `intoValue179`:
	return intoValue179
}

func TaintStepTest_ReflectValueSetMapIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue366` into `intoValue648`.

	// Assume that `sourceCQL` has the underlying type of `fromValue366`:
	fromValue366 := sourceCQL.(reflect.Value)

	// Declare `intoValue648` variable:
	var intoValue648 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue366` to the receiver `intoValue648`
	// (`intoValue648` is now tainted).
	intoValue648.SetMapIndex(fromValue366, reflect.Value{})

	// Return the tainted `intoValue648`:
	return intoValue648
}

func TaintStepTest_ReflectValueSetMapIndex_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue544` into `intoValue484`.

	// Assume that `sourceCQL` has the underlying type of `fromValue544`:
	fromValue544 := sourceCQL.(reflect.Value)

	// Declare `intoValue484` variable:
	var intoValue484 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue544` to the receiver `intoValue484`
	// (`intoValue484` is now tainted).
	intoValue484.SetMapIndex(reflect.Value{}, fromValue544)

	// Return the tainted `intoValue484`:
	return intoValue484
}

func TaintStepTest_ReflectValueSetPointer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPointer824` into `intoValue754`.

	// Assume that `sourceCQL` has the underlying type of `fromPointer824`:
	fromPointer824 := sourceCQL.(unsafe.Pointer)

	// Declare `intoValue754` variable:
	var intoValue754 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromPointer824` to the receiver `intoValue754`
	// (`intoValue754` is now tainted).
	intoValue754.SetPointer(fromPointer824)

	// Return the tainted `intoValue754`:
	return intoValue754
}

func TaintStepTest_ReflectValueSetString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString680` into `intoValue722`.

	// Assume that `sourceCQL` has the underlying type of `fromString680`:
	fromString680 := sourceCQL.(string)

	// Declare `intoValue722` variable:
	var intoValue722 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromString680` to the receiver `intoValue722`
	// (`intoValue722` is now tainted).
	intoValue722.SetString(fromString680)

	// Return the tainted `intoValue722`:
	return intoValue722
}

func TaintStepTest_ReflectValueSetUint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUint64506` into `intoValue121`.

	// Assume that `sourceCQL` has the underlying type of `fromUint64506`:
	fromUint64506 := sourceCQL.(uint64)

	// Declare `intoValue121` variable:
	var intoValue121 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromUint64506` to the receiver `intoValue121`
	// (`intoValue121` is now tainted).
	intoValue121.SetUint(fromUint64506)

	// Return the tainted `intoValue121`:
	return intoValue121
}

func TaintStepTest_ReflectValueSlice_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue293` into `intoValue151`.

	// Assume that `sourceCQL` has the underlying type of `fromValue293`:
	fromValue293 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue293` to the result `intoValue151`
	// (`intoValue151` is now tainted).
	intoValue151 := fromValue293.Slice(0, 0)

	// Return the tainted `intoValue151`:
	return intoValue151
}

func TaintStepTest_ReflectValueSlice3_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue849` into `intoValue322`.

	// Assume that `sourceCQL` has the underlying type of `fromValue849`:
	fromValue849 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue849` to the result `intoValue322`
	// (`intoValue322` is now tainted).
	intoValue322 := fromValue849.Slice3(0, 0, 0)

	// Return the tainted `intoValue322`:
	return intoValue322
}

func TaintStepTest_ReflectValueString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue339` into `intoString478`.

	// Assume that `sourceCQL` has the underlying type of `fromValue339`:
	fromValue339 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue339` to the result `intoString478`
	// (`intoString478` is now tainted).
	intoString478 := fromValue339.String()

	// Return the tainted `intoString478`:
	return intoString478
}

func TaintStepTest_ReflectValueTryRecv_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue399` into `intoValue426`.

	// Assume that `sourceCQL` has the underlying type of `fromValue399`:
	fromValue399 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue399` to the result `intoValue426`
	// (`intoValue426` is now tainted).
	intoValue426, _ := fromValue399.TryRecv()

	// Return the tainted `intoValue426`:
	return intoValue426
}

func TaintStepTest_ReflectValueTrySend_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue628` into `intoValue197`.

	// Assume that `sourceCQL` has the underlying type of `fromValue628`:
	fromValue628 := sourceCQL.(reflect.Value)

	// Declare `intoValue197` variable:
	var intoValue197 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `fromValue628` to the receiver `intoValue197`
	// (`intoValue197` is now tainted).
	intoValue197.TrySend(fromValue628)

	// Return the tainted `intoValue197`:
	return intoValue197
}

func TaintStepTest_ReflectValueUint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue216` into `intoUint64742`.

	// Assume that `sourceCQL` has the underlying type of `fromValue216`:
	fromValue216 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue216` to the result `intoUint64742`
	// (`intoUint64742` is now tainted).
	intoUint64742 := fromValue216.Uint()

	// Return the tainted `intoUint64742`:
	return intoUint64742
}

func TaintStepTest_ReflectValueUnsafeAddr_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValue906` into `intoUintptr620`.

	// Assume that `sourceCQL` has the underlying type of `fromValue906`:
	fromValue906 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `fromValue906` to the result `intoUintptr620`
	// (`intoUintptr620` is now tainted).
	intoUintptr620 := fromValue906.UnsafeAddr()

	// Return the tainted `intoUintptr620`:
	return intoUintptr620
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
