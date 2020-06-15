package main

import "expvar"

func TaintStepTest_ExpvarGet_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString616` into `intoVar615`.

	// Assume that `sourceCQL` has the underlying type of `fromString616`:
	fromString616 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString616` to result `intoVar615`
	// (`intoVar615` is now tainted).
	intoVar615 := expvar.Get(fromString616)

	// Sink the tainted `intoVar615`:
	sink(intoVar615)
}

func TaintStepTest_ExpvarFloatString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFloat385` into `intoString513`.

	// Assume that `sourceCQL` has the underlying type of `fromFloat385`:
	fromFloat385 := sourceCQL.(expvar.Float)

	// Call the method that transfers the taint
	// from the receiver `fromFloat385` to the result `intoString513`
	// (`intoString513` is now tainted).
	intoString513 := fromFloat385.String()

	// Sink the tainted `intoString513`:
	sink(intoString513)
}

func TaintStepTest_ExpvarFuncString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFunc374` into `intoString536`.

	// Assume that `sourceCQL` has the underlying type of `fromFunc374`:
	fromFunc374 := sourceCQL.(expvar.Func)

	// Call the method that transfers the taint
	// from the receiver `fromFunc374` to the result `intoString536`
	// (`intoString536` is now tainted).
	intoString536 := fromFunc374.String()

	// Sink the tainted `intoString536`:
	sink(intoString536)
}

func TaintStepTest_ExpvarFuncValue_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFunc405` into `intoInterface591`.

	// Assume that `sourceCQL` has the underlying type of `fromFunc405`:
	fromFunc405 := sourceCQL.(expvar.Func)

	// Call the method that transfers the taint
	// from the receiver `fromFunc405` to the result `intoInterface591`
	// (`intoInterface591` is now tainted).
	intoInterface591 := fromFunc405.Value()

	// Sink the tainted `intoInterface591`:
	sink(intoInterface591)
}

func TaintStepTest_ExpvarIntString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInt924` into `intoString924`.

	// Assume that `sourceCQL` has the underlying type of `fromInt924`:
	fromInt924 := sourceCQL.(expvar.Int)

	// Call the method that transfers the taint
	// from the receiver `fromInt924` to the result `intoString924`
	// (`intoString924` is now tainted).
	intoString924 := fromInt924.String()

	// Sink the tainted `intoString924`:
	sink(intoString924)
}

func TaintStepTest_ExpvarMapDo_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFuncexpvarKeyValue972` into `intoMap387`.

	// Assume that `sourceCQL` has the underlying type of `fromFuncexpvarKeyValue972`:
	fromFuncexpvarKeyValue972 := sourceCQL.(func(expvar.KeyValue))

	// Declare `intoMap387` variable:
	var intoMap387 expvar.Map

	// Call the method that transfers the taint
	// from the parameter `fromFuncexpvarKeyValue972` to the receiver `intoMap387`
	// (`intoMap387` is now tainted).
	intoMap387.Do(fromFuncexpvarKeyValue972)

	// Sink the tainted `intoMap387`:
	sink(intoMap387)
}

func TaintStepTest_ExpvarMapGet_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromMap550` into `intoVar850`.

	// Assume that `sourceCQL` has the underlying type of `fromMap550`:
	fromMap550 := sourceCQL.(expvar.Map)

	// Call the method that transfers the taint
	// from the receiver `fromMap550` to the result `intoVar850`
	// (`intoVar850` is now tainted).
	intoVar850 := fromMap550.Get("")

	// Sink the tainted `intoVar850`:
	sink(intoVar850)
}

func TaintStepTest_ExpvarMapSet_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString292` into `intoMap569`.

	// Assume that `sourceCQL` has the underlying type of `fromString292`:
	fromString292 := sourceCQL.(string)

	// Declare `intoMap569` variable:
	var intoMap569 expvar.Map

	// Call the method that transfers the taint
	// from the parameter `fromString292` to the receiver `intoMap569`
	// (`intoMap569` is now tainted).
	intoMap569.Set(fromString292, nil)

	// Sink the tainted `intoMap569`:
	sink(intoMap569)
}

func TaintStepTest_ExpvarMapSet_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromVar455` into `intoMap194`.

	// Assume that `sourceCQL` has the underlying type of `fromVar455`:
	fromVar455 := sourceCQL.(expvar.Var)

	// Declare `intoMap194` variable:
	var intoMap194 expvar.Map

	// Call the method that transfers the taint
	// from the parameter `fromVar455` to the receiver `intoMap194`
	// (`intoMap194` is now tainted).
	intoMap194.Set("", fromVar455)

	// Sink the tainted `intoMap194`:
	sink(intoMap194)
}

func TaintStepTest_ExpvarMapString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromMap505` into `intoString878`.

	// Assume that `sourceCQL` has the underlying type of `fromMap505`:
	fromMap505 := sourceCQL.(expvar.Map)

	// Call the method that transfers the taint
	// from the receiver `fromMap505` to the result `intoString878`
	// (`intoString878` is now tainted).
	intoString878 := fromMap505.String()

	// Sink the tainted `intoString878`:
	sink(intoString878)
}

func TaintStepTest_ExpvarStringSet_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString914` into `intoString691`.

	// Assume that `sourceCQL` has the underlying type of `fromString914`:
	fromString914 := sourceCQL.(string)

	// Declare `intoString691` variable:
	var intoString691 expvar.String

	// Call the method that transfers the taint
	// from the parameter `fromString914` to the receiver `intoString691`
	// (`intoString691` is now tainted).
	intoString691.Set(fromString914)

	// Sink the tainted `intoString691`:
	sink(intoString691)
}

func TaintStepTest_ExpvarStringString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString905` into `intoString983`.

	// Assume that `sourceCQL` has the underlying type of `fromString905`:
	fromString905 := sourceCQL.(expvar.String)

	// Call the method that transfers the taint
	// from the receiver `fromString905` to the result `intoString983`
	// (`intoString983` is now tainted).
	intoString983 := fromString905.String()

	// Sink the tainted `intoString983`:
	sink(intoString983)
}

func TaintStepTest_ExpvarStringValue_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString950` into `intoString181`.

	// Assume that `sourceCQL` has the underlying type of `fromString950`:
	fromString950 := sourceCQL.(expvar.String)

	// Call the method that transfers the taint
	// from the receiver `fromString950` to the result `intoString181`
	// (`intoString181` is now tainted).
	intoString181 := fromString950.Value()

	// Sink the tainted `intoString181`:
	sink(intoString181)
}

func TaintStepTest_ExpvarVarString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromVar764` into `intoString331`.

	// Assume that `sourceCQL` has the underlying type of `fromVar764`:
	fromVar764 := sourceCQL.(expvar.Var)

	// Call the method that transfers the taint
	// from the receiver `fromVar764` to the result `intoString331`
	// (`intoString331` is now tainted).
	intoString331 := fromVar764.String()

	// Sink the tainted `intoString331`:
	sink(intoString331)
}

func RunAllTaints_Expvar() {
	{
		source := newSource()
		TaintStepTest_ExpvarGet_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarFloatString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarFuncString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarFuncValue_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarIntString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarMapDo_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarMapGet_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarMapSet_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarMapSet_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarMapString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarStringSet_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarStringString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarStringValue_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarVarString_B0I0O0(source)
	}
}
