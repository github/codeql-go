// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "expvar"

func TaintStepTest_ExpvarGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString656` into `intoVar414`.

	// Assume that `sourceCQL` has the underlying type of `fromString656`:
	fromString656 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString656` to result `intoVar414`
	// (`intoVar414` is now tainted).
	intoVar414 := expvar.Get(fromString656)

	// Return the tainted `intoVar414`:
	return intoVar414
}

func TaintStepTest_ExpvarFloatString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFloat518` into `intoString650`.

	// Assume that `sourceCQL` has the underlying type of `fromFloat518`:
	fromFloat518 := sourceCQL.(expvar.Float)

	// Call the method that transfers the taint
	// from the receiver `fromFloat518` to the result `intoString650`
	// (`intoString650` is now tainted).
	intoString650 := fromFloat518.String()

	// Return the tainted `intoString650`:
	return intoString650
}

func TaintStepTest_ExpvarFuncString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFunc784` into `intoString957`.

	// Assume that `sourceCQL` has the underlying type of `fromFunc784`:
	fromFunc784 := sourceCQL.(expvar.Func)

	// Call the method that transfers the taint
	// from the receiver `fromFunc784` to the result `intoString957`
	// (`intoString957` is now tainted).
	intoString957 := fromFunc784.String()

	// Return the tainted `intoString957`:
	return intoString957
}

func TaintStepTest_ExpvarFuncValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFunc520` into `intoInterface443`.

	// Assume that `sourceCQL` has the underlying type of `fromFunc520`:
	fromFunc520 := sourceCQL.(expvar.Func)

	// Call the method that transfers the taint
	// from the receiver `fromFunc520` to the result `intoInterface443`
	// (`intoInterface443` is now tainted).
	intoInterface443 := fromFunc520.Value()

	// Return the tainted `intoInterface443`:
	return intoInterface443
}

func TaintStepTest_ExpvarIntString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInt127` into `intoString483`.

	// Assume that `sourceCQL` has the underlying type of `fromInt127`:
	fromInt127 := sourceCQL.(expvar.Int)

	// Call the method that transfers the taint
	// from the receiver `fromInt127` to the result `intoString483`
	// (`intoString483` is now tainted).
	intoString483 := fromInt127.String()

	// Return the tainted `intoString483`:
	return intoString483
}

func TaintStepTest_ExpvarMapDo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFuncexpvarKeyValue989` into `intoMap982`.

	// Assume that `sourceCQL` has the underlying type of `fromFuncexpvarKeyValue989`:
	fromFuncexpvarKeyValue989 := sourceCQL.(func(expvar.KeyValue))

	// Declare `intoMap982` variable:
	var intoMap982 expvar.Map

	// Call the method that transfers the taint
	// from the parameter `fromFuncexpvarKeyValue989` to the receiver `intoMap982`
	// (`intoMap982` is now tainted).
	intoMap982.Do(fromFuncexpvarKeyValue989)

	// Return the tainted `intoMap982`:
	return intoMap982
}

func TaintStepTest_ExpvarMapGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMap417` into `intoVar584`.

	// Assume that `sourceCQL` has the underlying type of `fromMap417`:
	fromMap417 := sourceCQL.(expvar.Map)

	// Call the method that transfers the taint
	// from the receiver `fromMap417` to the result `intoVar584`
	// (`intoVar584` is now tainted).
	intoVar584 := fromMap417.Get("")

	// Return the tainted `intoVar584`:
	return intoVar584
}

func TaintStepTest_ExpvarMapSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString991` into `intoMap881`.

	// Assume that `sourceCQL` has the underlying type of `fromString991`:
	fromString991 := sourceCQL.(string)

	// Declare `intoMap881` variable:
	var intoMap881 expvar.Map

	// Call the method that transfers the taint
	// from the parameter `fromString991` to the receiver `intoMap881`
	// (`intoMap881` is now tainted).
	intoMap881.Set(fromString991, nil)

	// Return the tainted `intoMap881`:
	return intoMap881
}

func TaintStepTest_ExpvarMapSet_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromVar186` into `intoMap284`.

	// Assume that `sourceCQL` has the underlying type of `fromVar186`:
	fromVar186 := sourceCQL.(expvar.Var)

	// Declare `intoMap284` variable:
	var intoMap284 expvar.Map

	// Call the method that transfers the taint
	// from the parameter `fromVar186` to the receiver `intoMap284`
	// (`intoMap284` is now tainted).
	intoMap284.Set("", fromVar186)

	// Return the tainted `intoMap284`:
	return intoMap284
}

func TaintStepTest_ExpvarMapString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMap908` into `intoString137`.

	// Assume that `sourceCQL` has the underlying type of `fromMap908`:
	fromMap908 := sourceCQL.(expvar.Map)

	// Call the method that transfers the taint
	// from the receiver `fromMap908` to the result `intoString137`
	// (`intoString137` is now tainted).
	intoString137 := fromMap908.String()

	// Return the tainted `intoString137`:
	return intoString137
}

func TaintStepTest_ExpvarStringSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString494` into `intoString873`.

	// Assume that `sourceCQL` has the underlying type of `fromString494`:
	fromString494 := sourceCQL.(string)

	// Declare `intoString873` variable:
	var intoString873 expvar.String

	// Call the method that transfers the taint
	// from the parameter `fromString494` to the receiver `intoString873`
	// (`intoString873` is now tainted).
	intoString873.Set(fromString494)

	// Return the tainted `intoString873`:
	return intoString873
}

func TaintStepTest_ExpvarStringString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString599` into `intoString409`.

	// Assume that `sourceCQL` has the underlying type of `fromString599`:
	fromString599 := sourceCQL.(expvar.String)

	// Call the method that transfers the taint
	// from the receiver `fromString599` to the result `intoString409`
	// (`intoString409` is now tainted).
	intoString409 := fromString599.String()

	// Return the tainted `intoString409`:
	return intoString409
}

func TaintStepTest_ExpvarStringValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString246` into `intoString898`.

	// Assume that `sourceCQL` has the underlying type of `fromString246`:
	fromString246 := sourceCQL.(expvar.String)

	// Call the method that transfers the taint
	// from the receiver `fromString246` to the result `intoString898`
	// (`intoString898` is now tainted).
	intoString898 := fromString246.Value()

	// Return the tainted `intoString898`:
	return intoString898
}

func TaintStepTest_ExpvarVarString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromVar598` into `intoString631`.

	// Assume that `sourceCQL` has the underlying type of `fromVar598`:
	fromVar598 := sourceCQL.(expvar.Var)

	// Call the method that transfers the taint
	// from the receiver `fromVar598` to the result `intoString631`
	// (`intoString631` is now tainted).
	intoString631 := fromVar598.String()

	// Return the tainted `intoString631`:
	return intoString631
}

func RunAllTaints_Expvar() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ExpvarGet_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ExpvarFloatString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ExpvarFuncString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ExpvarFuncValue_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ExpvarIntString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ExpvarMapDo_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ExpvarMapGet_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ExpvarMapSet_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ExpvarMapSet_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ExpvarMapString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ExpvarStringSet_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ExpvarStringString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ExpvarStringValue_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ExpvarVarString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
