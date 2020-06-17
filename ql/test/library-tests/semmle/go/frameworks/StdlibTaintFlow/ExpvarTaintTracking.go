// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "expvar"

func TaintStepTest_ExpvarGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString809` into `intoVar140`.

	// Assume that `sourceCQL` has the underlying type of `fromString809`:
	fromString809 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString809` to result `intoVar140`
	// (`intoVar140` is now tainted).
	intoVar140 := expvar.Get(fromString809)

	// Return the tainted `intoVar140`:
	return intoVar140
}

func TaintStepTest_ExpvarFloatString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFloat929` into `intoString715`.

	// Assume that `sourceCQL` has the underlying type of `fromFloat929`:
	fromFloat929 := sourceCQL.(expvar.Float)

	// Call the method that transfers the taint
	// from the receiver `fromFloat929` to the result `intoString715`
	// (`intoString715` is now tainted).
	intoString715 := fromFloat929.String()

	// Return the tainted `intoString715`:
	return intoString715
}

func TaintStepTest_ExpvarFuncString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFunc759` into `intoString456`.

	// Assume that `sourceCQL` has the underlying type of `fromFunc759`:
	fromFunc759 := sourceCQL.(expvar.Func)

	// Call the method that transfers the taint
	// from the receiver `fromFunc759` to the result `intoString456`
	// (`intoString456` is now tainted).
	intoString456 := fromFunc759.String()

	// Return the tainted `intoString456`:
	return intoString456
}

func TaintStepTest_ExpvarFuncValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFunc679` into `intoInterface495`.

	// Assume that `sourceCQL` has the underlying type of `fromFunc679`:
	fromFunc679 := sourceCQL.(expvar.Func)

	// Call the method that transfers the taint
	// from the receiver `fromFunc679` to the result `intoInterface495`
	// (`intoInterface495` is now tainted).
	intoInterface495 := fromFunc679.Value()

	// Return the tainted `intoInterface495`:
	return intoInterface495
}

func TaintStepTest_ExpvarIntString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInt552` into `intoString791`.

	// Assume that `sourceCQL` has the underlying type of `fromInt552`:
	fromInt552 := sourceCQL.(expvar.Int)

	// Call the method that transfers the taint
	// from the receiver `fromInt552` to the result `intoString791`
	// (`intoString791` is now tainted).
	intoString791 := fromInt552.String()

	// Return the tainted `intoString791`:
	return intoString791
}

func TaintStepTest_ExpvarMapDo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFuncexpvarKeyValue531` into `intoMap860`.

	// Assume that `sourceCQL` has the underlying type of `fromFuncexpvarKeyValue531`:
	fromFuncexpvarKeyValue531 := sourceCQL.(func(expvar.KeyValue))

	// Declare `intoMap860` variable:
	var intoMap860 expvar.Map

	// Call the method that transfers the taint
	// from the parameter `fromFuncexpvarKeyValue531` to the receiver `intoMap860`
	// (`intoMap860` is now tainted).
	intoMap860.Do(fromFuncexpvarKeyValue531)

	// Return the tainted `intoMap860`:
	return intoMap860
}

func TaintStepTest_ExpvarMapGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMap288` into `intoVar904`.

	// Assume that `sourceCQL` has the underlying type of `fromMap288`:
	fromMap288 := sourceCQL.(expvar.Map)

	// Call the method that transfers the taint
	// from the receiver `fromMap288` to the result `intoVar904`
	// (`intoVar904` is now tainted).
	intoVar904 := fromMap288.Get("")

	// Return the tainted `intoVar904`:
	return intoVar904
}

func TaintStepTest_ExpvarMapSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString704` into `intoMap267`.

	// Assume that `sourceCQL` has the underlying type of `fromString704`:
	fromString704 := sourceCQL.(string)

	// Declare `intoMap267` variable:
	var intoMap267 expvar.Map

	// Call the method that transfers the taint
	// from the parameter `fromString704` to the receiver `intoMap267`
	// (`intoMap267` is now tainted).
	intoMap267.Set(fromString704, nil)

	// Return the tainted `intoMap267`:
	return intoMap267
}

func TaintStepTest_ExpvarMapSet_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromVar449` into `intoMap846`.

	// Assume that `sourceCQL` has the underlying type of `fromVar449`:
	fromVar449 := sourceCQL.(expvar.Var)

	// Declare `intoMap846` variable:
	var intoMap846 expvar.Map

	// Call the method that transfers the taint
	// from the parameter `fromVar449` to the receiver `intoMap846`
	// (`intoMap846` is now tainted).
	intoMap846.Set("", fromVar449)

	// Return the tainted `intoMap846`:
	return intoMap846
}

func TaintStepTest_ExpvarMapString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMap507` into `intoString782`.

	// Assume that `sourceCQL` has the underlying type of `fromMap507`:
	fromMap507 := sourceCQL.(expvar.Map)

	// Call the method that transfers the taint
	// from the receiver `fromMap507` to the result `intoString782`
	// (`intoString782` is now tainted).
	intoString782 := fromMap507.String()

	// Return the tainted `intoString782`:
	return intoString782
}

func TaintStepTest_ExpvarStringSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString914` into `intoString270`.

	// Assume that `sourceCQL` has the underlying type of `fromString914`:
	fromString914 := sourceCQL.(string)

	// Declare `intoString270` variable:
	var intoString270 expvar.String

	// Call the method that transfers the taint
	// from the parameter `fromString914` to the receiver `intoString270`
	// (`intoString270` is now tainted).
	intoString270.Set(fromString914)

	// Return the tainted `intoString270`:
	return intoString270
}

func TaintStepTest_ExpvarStringString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString163` into `intoString504`.

	// Assume that `sourceCQL` has the underlying type of `fromString163`:
	fromString163 := sourceCQL.(expvar.String)

	// Call the method that transfers the taint
	// from the receiver `fromString163` to the result `intoString504`
	// (`intoString504` is now tainted).
	intoString504 := fromString163.String()

	// Return the tainted `intoString504`:
	return intoString504
}

func TaintStepTest_ExpvarStringValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString516` into `intoString797`.

	// Assume that `sourceCQL` has the underlying type of `fromString516`:
	fromString516 := sourceCQL.(expvar.String)

	// Call the method that transfers the taint
	// from the receiver `fromString516` to the result `intoString797`
	// (`intoString797` is now tainted).
	intoString797 := fromString516.Value()

	// Return the tainted `intoString797`:
	return intoString797
}

func TaintStepTest_ExpvarVarString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromVar181` into `intoString341`.

	// Assume that `sourceCQL` has the underlying type of `fromVar181`:
	fromVar181 := sourceCQL.(expvar.Var)

	// Call the method that transfers the taint
	// from the receiver `fromVar181` to the result `intoString341`
	// (`intoString341` is now tainted).
	intoString341 := fromVar181.String()

	// Return the tainted `intoString341`:
	return intoString341
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
