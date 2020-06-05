package main

import (
	"reflect"
	"unsafe"
)

func TaintStepTest_ReflectAppend(sourceCQL interface{}) {
	// The flow is from `x` into `into877`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `x` to result `into877`
	// (`into877` is now tainted).
	into877 := reflect.Append(reflect.Value{}, x)

	// Sink the tainted `into877`:
	sink(into877)
}

func TaintStepTest_ReflectAppendSlice(sourceCQL interface{}) {
	// The flow is from `t` into `into642`.

	// Assume that `sourceCQL` has the underlying type of `t`:
	t := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `t` to result `into642`
	// (`into642` is now tainted).
	into642 := reflect.AppendSlice(reflect.Value{}, t)

	// Sink the tainted `into642`:
	sink(into642)
}

func TaintStepTest_ReflectCopy(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.(reflect.Value)

	// Declare `dst` variable:
	var dst reflect.Value

	// Call the function that transfers the taint
	// from the parameter `src` to parameter `dst`;
	// `dst` is now tainted.
	reflect.Copy(dst, src)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_ReflectIndirect(sourceCQL interface{}) {
	// The flow is from `v` into `into374`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.(reflect.Value)

	// Call the function that transfers the taint
	// from the parameter `v` to result `into374`
	// (`into374` is now tainted).
	into374 := reflect.Indirect(v)

	// Sink the tainted `into374`:
	sink(into374)
}

func TaintStepTest_ReflectValueOf(sourceCQL interface{}) {
	// The flow is from `i` into `into426`.

	// Assume that `sourceCQL` has the underlying type of `i`:
	i := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `i` to result `into426`
	// (`into426` is now tainted).
	into426 := reflect.ValueOf(i)

	// Sink the tainted `into426`:
	sink(into426)
}

func TaintStepTest_ReflectMapIterKey(sourceCQL interface{}) {
	// The flow is from `from994` into `into559`.

	// Assume that `sourceCQL` has the underlying type of `from994`:
	from994 := sourceCQL.(reflect.MapIter)

	// Call the method that transfers the taint
	// from the receiver `from994` to the result `into559`
	// (`into559` is now tainted).
	into559 := from994.Key()

	// Sink the tainted `into559`:
	sink(into559)
}

func TaintStepTest_ReflectMapIterValue(sourceCQL interface{}) {
	// The flow is from `from127` into `into979`.

	// Assume that `sourceCQL` has the underlying type of `from127`:
	from127 := sourceCQL.(reflect.MapIter)

	// Call the method that transfers the taint
	// from the receiver `from127` to the result `into979`
	// (`into979` is now tainted).
	into979 := from127.Value()

	// Sink the tainted `into979`:
	sink(into979)
}

func TaintStepTest_ReflectStructTagGet(sourceCQL interface{}) {
	// The flow is from `from818` into `into896`.

	// Assume that `sourceCQL` has the underlying type of `from818`:
	from818 := sourceCQL.(reflect.StructTag)

	// Call the method that transfers the taint
	// from the receiver `from818` to the result `into896`
	// (`into896` is now tainted).
	into896 := from818.Get("")

	// Sink the tainted `into896`:
	sink(into896)
}

func TaintStepTest_ReflectStructTagLookup(sourceCQL interface{}) {
	// The flow is from `from704` into `value`.

	// Assume that `sourceCQL` has the underlying type of `from704`:
	from704 := sourceCQL.(reflect.StructTag)

	// Call the method that transfers the taint
	// from the receiver `from704` to the result `value`
	// (`value` is now tainted).
	value, _ := from704.Lookup("")

	// Sink the tainted `value`:
	sink(value)
}

func TaintStepTest_ReflectValueAddr(sourceCQL interface{}) {
	// The flow is from `from603` into `into747`.

	// Assume that `sourceCQL` has the underlying type of `from603`:
	from603 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from603` to the result `into747`
	// (`into747` is now tainted).
	into747 := from603.Addr()

	// Sink the tainted `into747`:
	sink(into747)
}

func TaintStepTest_ReflectValueBool(sourceCQL interface{}) {
	// The flow is from `from661` into `into522`.

	// Assume that `sourceCQL` has the underlying type of `from661`:
	from661 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from661` to the result `into522`
	// (`into522` is now tainted).
	into522 := from661.Bool()

	// Sink the tainted `into522`:
	sink(into522)
}

func TaintStepTest_ReflectValueBytes(sourceCQL interface{}) {
	// The flow is from `from557` into `into406`.

	// Assume that `sourceCQL` has the underlying type of `from557`:
	from557 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from557` to the result `into406`
	// (`into406` is now tainted).
	into406 := from557.Bytes()

	// Sink the tainted `into406`:
	sink(into406)
}

func TaintStepTest_ReflectValueComplex(sourceCQL interface{}) {
	// The flow is from `from325` into `into571`.

	// Assume that `sourceCQL` has the underlying type of `from325`:
	from325 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from325` to the result `into571`
	// (`into571` is now tainted).
	into571 := from325.Complex()

	// Sink the tainted `into571`:
	sink(into571)
}

func TaintStepTest_ReflectValueConvert(sourceCQL interface{}) {
	// The flow is from `from321` into `into822`.

	// Assume that `sourceCQL` has the underlying type of `from321`:
	from321 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from321` to the result `into822`
	// (`into822` is now tainted).
	into822 := from321.Convert(nil)

	// Sink the tainted `into822`:
	sink(into822)
}

func TaintStepTest_ReflectValueElem(sourceCQL interface{}) {
	// The flow is from `from375` into `into570`.

	// Assume that `sourceCQL` has the underlying type of `from375`:
	from375 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from375` to the result `into570`
	// (`into570` is now tainted).
	into570 := from375.Elem()

	// Sink the tainted `into570`:
	sink(into570)
}

func TaintStepTest_ReflectValueField(sourceCQL interface{}) {
	// The flow is from `from540` into `into336`.

	// Assume that `sourceCQL` has the underlying type of `from540`:
	from540 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from540` to the result `into336`
	// (`into336` is now tainted).
	into336 := from540.Field(0)

	// Sink the tainted `into336`:
	sink(into336)
}

func TaintStepTest_ReflectValueFieldByIndex(sourceCQL interface{}) {
	// The flow is from `from238` into `into405`.

	// Assume that `sourceCQL` has the underlying type of `from238`:
	from238 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from238` to the result `into405`
	// (`into405` is now tainted).
	into405 := from238.FieldByIndex(nil)

	// Sink the tainted `into405`:
	sink(into405)
}

func TaintStepTest_ReflectValueFieldByName(sourceCQL interface{}) {
	// The flow is from `from609` into `into602`.

	// Assume that `sourceCQL` has the underlying type of `from609`:
	from609 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from609` to the result `into602`
	// (`into602` is now tainted).
	into602 := from609.FieldByName("")

	// Sink the tainted `into602`:
	sink(into602)
}

func TaintStepTest_ReflectValueFieldByNameFunc(sourceCQL interface{}) {
	// The flow is from `from298` into `into461`.

	// Assume that `sourceCQL` has the underlying type of `from298`:
	from298 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from298` to the result `into461`
	// (`into461` is now tainted).
	into461 := from298.FieldByNameFunc(nil)

	// Sink the tainted `into461`:
	sink(into461)
}

func TaintStepTest_ReflectValueFloat(sourceCQL interface{}) {
	// The flow is from `from717` into `into488`.

	// Assume that `sourceCQL` has the underlying type of `from717`:
	from717 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from717` to the result `into488`
	// (`into488` is now tainted).
	into488 := from717.Float()

	// Sink the tainted `into488`:
	sink(into488)
}

func TaintStepTest_ReflectValueIndex(sourceCQL interface{}) {
	// The flow is from `from757` into `into270`.

	// Assume that `sourceCQL` has the underlying type of `from757`:
	from757 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from757` to the result `into270`
	// (`into270` is now tainted).
	into270 := from757.Index(0)

	// Sink the tainted `into270`:
	sink(into270)
}

func TaintStepTest_ReflectValueInt(sourceCQL interface{}) {
	// The flow is from `from958` into `into226`.

	// Assume that `sourceCQL` has the underlying type of `from958`:
	from958 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from958` to the result `into226`
	// (`into226` is now tainted).
	into226 := from958.Int()

	// Sink the tainted `into226`:
	sink(into226)
}

func TaintStepTest_ReflectValueInterface(sourceCQL interface{}) {
	// The flow is from `from425` into `i`.

	// Assume that `sourceCQL` has the underlying type of `from425`:
	from425 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from425` to the result `i`
	// (`i` is now tainted).
	i := from425.Interface()

	// Sink the tainted `i`:
	sink(i)
}

func TaintStepTest_ReflectValueInterfaceData(sourceCQL interface{}) {
	// The flow is from `from882` into `into936`.

	// Assume that `sourceCQL` has the underlying type of `from882`:
	from882 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from882` to the result `into936`
	// (`into936` is now tainted).
	into936 := from882.InterfaceData()

	// Sink the tainted `into936`:
	sink(into936)
}

func TaintStepTest_ReflectValueMapIndex(sourceCQL interface{}) {
	// The flow is from `from118` into `into293`.

	// Assume that `sourceCQL` has the underlying type of `from118`:
	from118 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from118` to the result `into293`
	// (`into293` is now tainted).
	into293 := from118.MapIndex(reflect.Value{})

	// Sink the tainted `into293`:
	sink(into293)
}

func TaintStepTest_ReflectValueMapKeys(sourceCQL interface{}) {
	// The flow is from `from947` into `into874`.

	// Assume that `sourceCQL` has the underlying type of `from947`:
	from947 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from947` to the result `into874`
	// (`into874` is now tainted).
	into874 := from947.MapKeys()

	// Sink the tainted `into874`:
	sink(into874)
}

func TaintStepTest_ReflectValueMapRange(sourceCQL interface{}) {
	// The flow is from `from708` into `into786`.

	// Assume that `sourceCQL` has the underlying type of `from708`:
	from708 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from708` to the result `into786`
	// (`into786` is now tainted).
	into786 := from708.MapRange()

	// Sink the tainted `into786`:
	sink(into786)
}

func TaintStepTest_ReflectValueMethod(sourceCQL interface{}) {
	// The flow is from `from159` into `into876`.

	// Assume that `sourceCQL` has the underlying type of `from159`:
	from159 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from159` to the result `into876`
	// (`into876` is now tainted).
	into876 := from159.Method(0)

	// Sink the tainted `into876`:
	sink(into876)
}

func TaintStepTest_ReflectValueMethodByName(sourceCQL interface{}) {
	// The flow is from `from657` into `into574`.

	// Assume that `sourceCQL` has the underlying type of `from657`:
	from657 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from657` to the result `into574`
	// (`into574` is now tainted).
	into574 := from657.MethodByName("")

	// Sink the tainted `into574`:
	sink(into574)
}

func TaintStepTest_ReflectValuePointer(sourceCQL interface{}) {
	// The flow is from `from282` into `into155`.

	// Assume that `sourceCQL` has the underlying type of `from282`:
	from282 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from282` to the result `into155`
	// (`into155` is now tainted).
	into155 := from282.Pointer()

	// Sink the tainted `into155`:
	sink(into155)
}

func TaintStepTest_ReflectValueRecv(sourceCQL interface{}) {
	// The flow is from `from951` into `x`.

	// Assume that `sourceCQL` has the underlying type of `from951`:
	from951 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from951` to the result `x`
	// (`x` is now tainted).
	x, _ := from951.Recv()

	// Sink the tainted `x`:
	sink(x)
}

func TaintStepTest_ReflectValueSend(sourceCQL interface{}) {
	// The flow is from `x` into `into866`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.(reflect.Value)

	// Declare `into866` variable:
	var into866 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `x` to the receiver `into866`
	// (`into866` is now tainted).
	into866.Send(x)

	// Sink the tainted `into866`:
	sink(into866)
}

func TaintStepTest_ReflectValueSet(sourceCQL interface{}) {
	// The flow is from `x` into `into501`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.(reflect.Value)

	// Declare `into501` variable:
	var into501 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `x` to the receiver `into501`
	// (`into501` is now tainted).
	into501.Set(x)

	// Sink the tainted `into501`:
	sink(into501)
}

func TaintStepTest_ReflectValueSetBool(sourceCQL interface{}) {
	// The flow is from `x` into `into439`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.(bool)

	// Declare `into439` variable:
	var into439 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `x` to the receiver `into439`
	// (`into439` is now tainted).
	into439.SetBool(x)

	// Sink the tainted `into439`:
	sink(into439)
}

func TaintStepTest_ReflectValueSetBytes(sourceCQL interface{}) {
	// The flow is from `x` into `into130`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.([]byte)

	// Declare `into130` variable:
	var into130 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `x` to the receiver `into130`
	// (`into130` is now tainted).
	into130.SetBytes(x)

	// Sink the tainted `into130`:
	sink(into130)
}

func TaintStepTest_ReflectValueSetComplex(sourceCQL interface{}) {
	// The flow is from `x` into `into521`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.(complex128)

	// Declare `into521` variable:
	var into521 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `x` to the receiver `into521`
	// (`into521` is now tainted).
	into521.SetComplex(x)

	// Sink the tainted `into521`:
	sink(into521)
}

func TaintStepTest_ReflectValueSetFloat(sourceCQL interface{}) {
	// The flow is from `x` into `into795`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.(float64)

	// Declare `into795` variable:
	var into795 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `x` to the receiver `into795`
	// (`into795` is now tainted).
	into795.SetFloat(x)

	// Sink the tainted `into795`:
	sink(into795)
}

func TaintStepTest_ReflectValueSetInt(sourceCQL interface{}) {
	// The flow is from `x` into `into791`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.(int64)

	// Declare `into791` variable:
	var into791 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `x` to the receiver `into791`
	// (`into791` is now tainted).
	into791.SetInt(x)

	// Sink the tainted `into791`:
	sink(into791)
}

func TaintStepTest_ReflectValueSetMapIndex(sourceCQL interface{}) {
	// The flow is from `key` into `into247`.

	// Assume that `sourceCQL` has the underlying type of `key`:
	key := sourceCQL.(reflect.Value)

	// Declare `into247` variable:
	var into247 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `key` to the receiver `into247`
	// (`into247` is now tainted).
	into247.SetMapIndex(key, reflect.Value{})

	// Sink the tainted `into247`:
	sink(into247)
}

func TaintStepTest_ReflectValueSetPointer(sourceCQL interface{}) {
	// The flow is from `x` into `into659`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.(unsafe.Pointer)

	// Declare `into659` variable:
	var into659 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `x` to the receiver `into659`
	// (`into659` is now tainted).
	into659.SetPointer(x)

	// Sink the tainted `into659`:
	sink(into659)
}

func TaintStepTest_ReflectValueSetString(sourceCQL interface{}) {
	// The flow is from `x` into `into233`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.(string)

	// Declare `into233` variable:
	var into233 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `x` to the receiver `into233`
	// (`into233` is now tainted).
	into233.SetString(x)

	// Sink the tainted `into233`:
	sink(into233)
}

func TaintStepTest_ReflectValueSetUint(sourceCQL interface{}) {
	// The flow is from `x` into `into232`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.(uint64)

	// Declare `into232` variable:
	var into232 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `x` to the receiver `into232`
	// (`into232` is now tainted).
	into232.SetUint(x)

	// Sink the tainted `into232`:
	sink(into232)
}

func TaintStepTest_ReflectValueSlice(sourceCQL interface{}) {
	// The flow is from `from181` into `into466`.

	// Assume that `sourceCQL` has the underlying type of `from181`:
	from181 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from181` to the result `into466`
	// (`into466` is now tainted).
	into466 := from181.Slice(0, 0)

	// Sink the tainted `into466`:
	sink(into466)
}

func TaintStepTest_ReflectValueSlice3(sourceCQL interface{}) {
	// The flow is from `from483` into `into140`.

	// Assume that `sourceCQL` has the underlying type of `from483`:
	from483 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from483` to the result `into140`
	// (`into140` is now tainted).
	into140 := from483.Slice3(0, 0, 0)

	// Sink the tainted `into140`:
	sink(into140)
}

func TaintStepTest_ReflectValueString(sourceCQL interface{}) {
	// The flow is from `from838` into `into991`.

	// Assume that `sourceCQL` has the underlying type of `from838`:
	from838 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from838` to the result `into991`
	// (`into991` is now tainted).
	into991 := from838.String()

	// Sink the tainted `into991`:
	sink(into991)
}

func TaintStepTest_ReflectValueTryRecv(sourceCQL interface{}) {
	// The flow is from `from414` into `x`.

	// Assume that `sourceCQL` has the underlying type of `from414`:
	from414 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from414` to the result `x`
	// (`x` is now tainted).
	x, _ := from414.TryRecv()

	// Sink the tainted `x`:
	sink(x)
}

func TaintStepTest_ReflectValueTrySend(sourceCQL interface{}) {
	// The flow is from `x` into `into868`.

	// Assume that `sourceCQL` has the underlying type of `x`:
	x := sourceCQL.(reflect.Value)

	// Declare `into868` variable:
	var into868 reflect.Value

	// Call the method that transfers the taint
	// from the parameter `x` to the receiver `into868`
	// (`into868` is now tainted).
	into868.TrySend(x)

	// Sink the tainted `into868`:
	sink(into868)
}

func TaintStepTest_ReflectValueUint(sourceCQL interface{}) {
	// The flow is from `from947` into `into736`.

	// Assume that `sourceCQL` has the underlying type of `from947`:
	from947 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from947` to the result `into736`
	// (`into736` is now tainted).
	into736 := from947.Uint()

	// Sink the tainted `into736`:
	sink(into736)
}

func TaintStepTest_ReflectValueUnsafeAddr(sourceCQL interface{}) {
	// The flow is from `from842` into `into833`.

	// Assume that `sourceCQL` has the underlying type of `from842`:
	from842 := sourceCQL.(reflect.Value)

	// Call the method that transfers the taint
	// from the receiver `from842` to the result `into833`
	// (`into833` is now tainted).
	into833 := from842.UnsafeAddr()

	// Sink the tainted `into833`:
	sink(into833)
}

func RunAllTaints_Reflect(v interface{}) {
	{
		source := newSource()
		TaintStepTest_ReflectAppend(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectAppendSlice(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectCopy(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectIndirect(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueOf(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectMapIterKey(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectMapIterValue(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectStructTagGet(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectStructTagLookup(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueAddr(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueBool(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueBytes(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueComplex(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueConvert(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueElem(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueField(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueFieldByIndex(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueFieldByName(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueFieldByNameFunc(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueFloat(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueIndex(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueInt(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueInterface(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueInterfaceData(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueMapIndex(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueMapKeys(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueMapRange(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueMethod(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueMethodByName(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValuePointer(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueRecv(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSend(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSet(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetBool(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetBytes(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetComplex(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetFloat(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetInt(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetMapIndex(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetPointer(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetString(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSetUint(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSlice(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueSlice3(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueString(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueTryRecv(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueTrySend(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueUint(source)
	}
	{
		source := newSource()
		TaintStepTest_ReflectValueUnsafeAddr(source)
	}
}
