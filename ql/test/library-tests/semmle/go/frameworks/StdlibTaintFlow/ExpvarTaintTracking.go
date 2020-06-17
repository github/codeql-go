package main

import "expvar"

func TaintStepTest_ExpvarGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString612` into `intoVar261`.

	// Assume that `sourceCQL` has the underlying type of `fromString612`:
	fromString612 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString612` to result `intoVar261`
	// (`intoVar261` is now tainted).
	intoVar261 := expvar.Get(fromString612)

	// Return the tainted `intoVar261`:
	return intoVar261
}

func TaintStepTest_ExpvarFloatString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFloat422` into `intoString448`.

	// Assume that `sourceCQL` has the underlying type of `fromFloat422`:
	fromFloat422 := sourceCQL.(expvar.Float)

	// Call the method that transfers the taint
	// from the receiver `fromFloat422` to the result `intoString448`
	// (`intoString448` is now tainted).
	intoString448 := fromFloat422.String()

	// Return the tainted `intoString448`:
	return intoString448
}

func TaintStepTest_ExpvarFuncString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFunc590` into `intoString376`.

	// Assume that `sourceCQL` has the underlying type of `fromFunc590`:
	fromFunc590 := sourceCQL.(expvar.Func)

	// Call the method that transfers the taint
	// from the receiver `fromFunc590` to the result `intoString376`
	// (`intoString376` is now tainted).
	intoString376 := fromFunc590.String()

	// Return the tainted `intoString376`:
	return intoString376
}

func TaintStepTest_ExpvarFuncValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFunc415` into `intoInterface633`.

	// Assume that `sourceCQL` has the underlying type of `fromFunc415`:
	fromFunc415 := sourceCQL.(expvar.Func)

	// Call the method that transfers the taint
	// from the receiver `fromFunc415` to the result `intoInterface633`
	// (`intoInterface633` is now tainted).
	intoInterface633 := fromFunc415.Value()

	// Return the tainted `intoInterface633`:
	return intoInterface633
}

func TaintStepTest_ExpvarIntString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInt345` into `intoString251`.

	// Assume that `sourceCQL` has the underlying type of `fromInt345`:
	fromInt345 := sourceCQL.(expvar.Int)

	// Call the method that transfers the taint
	// from the receiver `fromInt345` to the result `intoString251`
	// (`intoString251` is now tainted).
	intoString251 := fromInt345.String()

	// Return the tainted `intoString251`:
	return intoString251
}

func TaintStepTest_ExpvarMapDo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFuncexpvarKeyValue720` into `intoMap501`.

	// Assume that `sourceCQL` has the underlying type of `fromFuncexpvarKeyValue720`:
	fromFuncexpvarKeyValue720 := sourceCQL.(func(expvar.KeyValue))

	// Declare `intoMap501` variable:
	var intoMap501 expvar.Map

	// Call the method that transfers the taint
	// from the parameter `fromFuncexpvarKeyValue720` to the receiver `intoMap501`
	// (`intoMap501` is now tainted).
	intoMap501.Do(fromFuncexpvarKeyValue720)

	// Return the tainted `intoMap501`:
	return intoMap501
}

func TaintStepTest_ExpvarMapGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMap332` into `intoVar684`.

	// Assume that `sourceCQL` has the underlying type of `fromMap332`:
	fromMap332 := sourceCQL.(expvar.Map)

	// Call the method that transfers the taint
	// from the receiver `fromMap332` to the result `intoVar684`
	// (`intoVar684` is now tainted).
	intoVar684 := fromMap332.Get("")

	// Return the tainted `intoVar684`:
	return intoVar684
}

func TaintStepTest_ExpvarMapSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString679` into `intoMap127`.

	// Assume that `sourceCQL` has the underlying type of `fromString679`:
	fromString679 := sourceCQL.(string)

	// Declare `intoMap127` variable:
	var intoMap127 expvar.Map

	// Call the method that transfers the taint
	// from the parameter `fromString679` to the receiver `intoMap127`
	// (`intoMap127` is now tainted).
	intoMap127.Set(fromString679, nil)

	// Return the tainted `intoMap127`:
	return intoMap127
}

func TaintStepTest_ExpvarMapSet_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromVar371` into `intoMap662`.

	// Assume that `sourceCQL` has the underlying type of `fromVar371`:
	fromVar371 := sourceCQL.(expvar.Var)

	// Declare `intoMap662` variable:
	var intoMap662 expvar.Map

	// Call the method that transfers the taint
	// from the parameter `fromVar371` to the receiver `intoMap662`
	// (`intoMap662` is now tainted).
	intoMap662.Set("", fromVar371)

	// Return the tainted `intoMap662`:
	return intoMap662
}

func TaintStepTest_ExpvarMapString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromMap383` into `intoString660`.

	// Assume that `sourceCQL` has the underlying type of `fromMap383`:
	fromMap383 := sourceCQL.(expvar.Map)

	// Call the method that transfers the taint
	// from the receiver `fromMap383` to the result `intoString660`
	// (`intoString660` is now tainted).
	intoString660 := fromMap383.String()

	// Return the tainted `intoString660`:
	return intoString660
}

func TaintStepTest_ExpvarStringSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString291` into `intoString534`.

	// Assume that `sourceCQL` has the underlying type of `fromString291`:
	fromString291 := sourceCQL.(string)

	// Declare `intoString534` variable:
	var intoString534 expvar.String

	// Call the method that transfers the taint
	// from the parameter `fromString291` to the receiver `intoString534`
	// (`intoString534` is now tainted).
	intoString534.Set(fromString291)

	// Return the tainted `intoString534`:
	return intoString534
}

func TaintStepTest_ExpvarStringString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString908` into `intoString793`.

	// Assume that `sourceCQL` has the underlying type of `fromString908`:
	fromString908 := sourceCQL.(expvar.String)

	// Call the method that transfers the taint
	// from the receiver `fromString908` to the result `intoString793`
	// (`intoString793` is now tainted).
	intoString793 := fromString908.String()

	// Return the tainted `intoString793`:
	return intoString793
}

func TaintStepTest_ExpvarStringValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString713` into `intoString740`.

	// Assume that `sourceCQL` has the underlying type of `fromString713`:
	fromString713 := sourceCQL.(expvar.String)

	// Call the method that transfers the taint
	// from the receiver `fromString713` to the result `intoString740`
	// (`intoString740` is now tainted).
	intoString740 := fromString713.Value()

	// Return the tainted `intoString740`:
	return intoString740
}

func TaintStepTest_ExpvarVarString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromVar602` into `intoString724`.

	// Assume that `sourceCQL` has the underlying type of `fromVar602`:
	fromVar602 := sourceCQL.(expvar.Var)

	// Call the method that transfers the taint
	// from the receiver `fromVar602` to the result `intoString724`
	// (`intoString724` is now tainted).
	intoString724 := fromVar602.String()

	// Return the tainted `intoString724`:
	return intoString724
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
