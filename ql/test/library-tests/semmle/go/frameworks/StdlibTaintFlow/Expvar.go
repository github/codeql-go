package main

import "expvar"

func TaintStepTest_ExpvarGet(sourceCQL interface{}) {
	// The flow is from `name` into `into695`.

	// Assume that `sourceCQL` has the underlying type of `name`:
	name := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `name` to result `into695`
	// (`into695` is now tainted).
	into695 := expvar.Get(name)

	// Sink the tainted `into695`:
	sink(into695)
}

func TaintStepTest_ExpvarFloatAdd(sourceCQL interface{}) {
	// The flow is from `delta` into `into781`.

	// Assume that `sourceCQL` has the underlying type of `delta`:
	delta := sourceCQL.(float64)

	// Declare `into781` variable:
	var into781 expvar.Float

	// Call the method that transfers the taint
	// from the parameter `delta` to the receiver `into781`
	// (`into781` is now tainted).
	into781.Add(delta)

	// Sink the tainted `into781`:
	sink(into781)
}

func TaintStepTest_ExpvarFloatSet(sourceCQL interface{}) {
	// The flow is from `value` into `into816`.

	// Assume that `sourceCQL` has the underlying type of `value`:
	value := sourceCQL.(float64)

	// Declare `into816` variable:
	var into816 expvar.Float

	// Call the method that transfers the taint
	// from the parameter `value` to the receiver `into816`
	// (`into816` is now tainted).
	into816.Set(value)

	// Sink the tainted `into816`:
	sink(into816)
}

func TaintStepTest_ExpvarFloatString(sourceCQL interface{}) {
	// The flow is from `from960` into `into242`.

	// Assume that `sourceCQL` has the underlying type of `from960`:
	from960 := sourceCQL.(expvar.Float)

	// Call the method that transfers the taint
	// from the receiver `from960` to the result `into242`
	// (`into242` is now tainted).
	into242 := from960.String()

	// Sink the tainted `into242`:
	sink(into242)
}

func TaintStepTest_ExpvarFloatValue(sourceCQL interface{}) {
	// The flow is from `from737` into `into830`.

	// Assume that `sourceCQL` has the underlying type of `from737`:
	from737 := sourceCQL.(expvar.Float)

	// Call the method that transfers the taint
	// from the receiver `from737` to the result `into830`
	// (`into830` is now tainted).
	into830 := from737.Value()

	// Sink the tainted `into830`:
	sink(into830)
}

func TaintStepTest_ExpvarFuncString(sourceCQL interface{}) {
	// The flow is from `from610` into `into853`.

	// Assume that `sourceCQL` has the underlying type of `from610`:
	from610 := sourceCQL.(expvar.Func)

	// Call the method that transfers the taint
	// from the receiver `from610` to the result `into853`
	// (`into853` is now tainted).
	into853 := from610.String()

	// Sink the tainted `into853`:
	sink(into853)
}

func TaintStepTest_ExpvarFuncValue(sourceCQL interface{}) {
	// The flow is from `from919` into `into890`.

	// Assume that `sourceCQL` has the underlying type of `from919`:
	from919 := sourceCQL.(expvar.Func)

	// Call the method that transfers the taint
	// from the receiver `from919` to the result `into890`
	// (`into890` is now tainted).
	into890 := from919.Value()

	// Sink the tainted `into890`:
	sink(into890)
}

func TaintStepTest_ExpvarIntAdd(sourceCQL interface{}) {
	// The flow is from `delta` into `into478`.

	// Assume that `sourceCQL` has the underlying type of `delta`:
	delta := sourceCQL.(int64)

	// Declare `into478` variable:
	var into478 expvar.Int

	// Call the method that transfers the taint
	// from the parameter `delta` to the receiver `into478`
	// (`into478` is now tainted).
	into478.Add(delta)

	// Sink the tainted `into478`:
	sink(into478)
}

func TaintStepTest_ExpvarIntSet(sourceCQL interface{}) {
	// The flow is from `value` into `into331`.

	// Assume that `sourceCQL` has the underlying type of `value`:
	value := sourceCQL.(int64)

	// Declare `into331` variable:
	var into331 expvar.Int

	// Call the method that transfers the taint
	// from the parameter `value` to the receiver `into331`
	// (`into331` is now tainted).
	into331.Set(value)

	// Sink the tainted `into331`:
	sink(into331)
}

func TaintStepTest_ExpvarIntString(sourceCQL interface{}) {
	// The flow is from `from586` into `into349`.

	// Assume that `sourceCQL` has the underlying type of `from586`:
	from586 := sourceCQL.(expvar.Int)

	// Call the method that transfers the taint
	// from the receiver `from586` to the result `into349`
	// (`into349` is now tainted).
	into349 := from586.String()

	// Sink the tainted `into349`:
	sink(into349)
}

func TaintStepTest_ExpvarIntValue(sourceCQL interface{}) {
	// The flow is from `from709` into `into454`.

	// Assume that `sourceCQL` has the underlying type of `from709`:
	from709 := sourceCQL.(expvar.Int)

	// Call the method that transfers the taint
	// from the receiver `from709` to the result `into454`
	// (`into454` is now tainted).
	into454 := from709.Value()

	// Sink the tainted `into454`:
	sink(into454)
}

func TaintStepTest_ExpvarMapAdd(sourceCQL interface{}) {
	// The flow is from `delta` into `into795`.

	// Assume that `sourceCQL` has the underlying type of `delta`:
	delta := sourceCQL.(int64)

	// Declare `into795` variable:
	var into795 expvar.Map

	// Call the method that transfers the taint
	// from the parameter `delta` to the receiver `into795`
	// (`into795` is now tainted).
	into795.Add("", delta)

	// Sink the tainted `into795`:
	sink(into795)
}

func TaintStepTest_ExpvarMapAddFloat(sourceCQL interface{}) {
	// The flow is from `delta` into `into744`.

	// Assume that `sourceCQL` has the underlying type of `delta`:
	delta := sourceCQL.(float64)

	// Declare `into744` variable:
	var into744 expvar.Map

	// Call the method that transfers the taint
	// from the parameter `delta` to the receiver `into744`
	// (`into744` is now tainted).
	into744.AddFloat("", delta)

	// Sink the tainted `into744`:
	sink(into744)
}

func TaintStepTest_ExpvarMapDo(sourceCQL interface{}) {
	// The flow is from `f` into `into367`.

	// Assume that `sourceCQL` has the underlying type of `f`:
	f := sourceCQL.(func(expvar.KeyValue))

	// Declare `into367` variable:
	var into367 expvar.Map

	// Call the method that transfers the taint
	// from the parameter `f` to the receiver `into367`
	// (`into367` is now tainted).
	into367.Do(f)

	// Sink the tainted `into367`:
	sink(into367)
}

func TaintStepTest_ExpvarMapGet(sourceCQL interface{}) {
	// The flow is from `from549` into `into694`.

	// Assume that `sourceCQL` has the underlying type of `from549`:
	from549 := sourceCQL.(expvar.Map)

	// Call the method that transfers the taint
	// from the receiver `from549` to the result `into694`
	// (`into694` is now tainted).
	into694 := from549.Get("")

	// Sink the tainted `into694`:
	sink(into694)
}

func TaintStepTest_ExpvarMapSet(sourceCQL interface{}) {
	// The flow is from `av` into `into741`.

	// Assume that `sourceCQL` has the underlying type of `av`:
	av := sourceCQL.(expvar.Var)

	// Declare `into741` variable:
	var into741 expvar.Map

	// Call the method that transfers the taint
	// from the parameter `av` to the receiver `into741`
	// (`into741` is now tainted).
	into741.Set("", av)

	// Sink the tainted `into741`:
	sink(into741)
}

func TaintStepTest_ExpvarMapString(sourceCQL interface{}) {
	// The flow is from `from343` into `into801`.

	// Assume that `sourceCQL` has the underlying type of `from343`:
	from343 := sourceCQL.(expvar.Map)

	// Call the method that transfers the taint
	// from the receiver `from343` to the result `into801`
	// (`into801` is now tainted).
	into801 := from343.String()

	// Sink the tainted `into801`:
	sink(into801)
}

func TaintStepTest_ExpvarStringSet(sourceCQL interface{}) {
	// The flow is from `value` into `into552`.

	// Assume that `sourceCQL` has the underlying type of `value`:
	value := sourceCQL.(string)

	// Declare `into552` variable:
	var into552 expvar.String

	// Call the method that transfers the taint
	// from the parameter `value` to the receiver `into552`
	// (`into552` is now tainted).
	into552.Set(value)

	// Sink the tainted `into552`:
	sink(into552)
}

func TaintStepTest_ExpvarStringString(sourceCQL interface{}) {
	// The flow is from `from555` into `into162`.

	// Assume that `sourceCQL` has the underlying type of `from555`:
	from555 := sourceCQL.(expvar.String)

	// Call the method that transfers the taint
	// from the receiver `from555` to the result `into162`
	// (`into162` is now tainted).
	into162 := from555.String()

	// Sink the tainted `into162`:
	sink(into162)
}

func TaintStepTest_ExpvarStringValue(sourceCQL interface{}) {
	// The flow is from `from918` into `into962`.

	// Assume that `sourceCQL` has the underlying type of `from918`:
	from918 := sourceCQL.(expvar.String)

	// Call the method that transfers the taint
	// from the receiver `from918` to the result `into962`
	// (`into962` is now tainted).
	into962 := from918.Value()

	// Sink the tainted `into962`:
	sink(into962)
}

func TaintStepTest_ExpvarVarString(sourceCQL interface{}) {
	// The flow is from `from698` into `into181`.

	// Assume that `sourceCQL` has the underlying type of `from698`:
	from698 := sourceCQL.(expvar.Var)

	// Call the method that transfers the taint
	// from the receiver `from698` to the result `into181`
	// (`into181` is now tainted).
	into181 := from698.String()

	// Sink the tainted `into181`:
	sink(into181)
}

func RunAllTaints_Expvar(v interface{}) {
	{
		source := newSource()
		TaintStepTest_ExpvarGet(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarFloatAdd(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarFloatSet(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarFloatString(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarFloatValue(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarFuncString(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarFuncValue(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarIntAdd(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarIntSet(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarIntString(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarIntValue(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarMapAdd(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarMapAddFloat(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarMapDo(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarMapGet(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarMapSet(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarMapString(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarStringSet(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarStringString(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarStringValue(source)
	}
	{
		source := newSource()
		TaintStepTest_ExpvarVarString(source)
	}
}
