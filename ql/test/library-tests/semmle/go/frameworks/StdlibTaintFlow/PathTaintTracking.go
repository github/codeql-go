package main

import "path"

func TaintStepTest_PathBase_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString315` into `intoString965`.

	// Assume that `sourceCQL` has the underlying type of `fromString315`:
	fromString315 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString315` to result `intoString965`
	// (`intoString965` is now tainted).
	intoString965 := path.Base(fromString315)

	// Sink the tainted `intoString965`:
	sink(intoString965)
}

func TaintStepTest_PathClean_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString607` into `intoString726`.

	// Assume that `sourceCQL` has the underlying type of `fromString607`:
	fromString607 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString607` to result `intoString726`
	// (`intoString726` is now tainted).
	intoString726 := path.Clean(fromString607)

	// Sink the tainted `intoString726`:
	sink(intoString726)
}

func TaintStepTest_PathDir_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString847` into `intoString677`.

	// Assume that `sourceCQL` has the underlying type of `fromString847`:
	fromString847 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString847` to result `intoString677`
	// (`intoString677` is now tainted).
	intoString677 := path.Dir(fromString847)

	// Sink the tainted `intoString677`:
	sink(intoString677)
}

func TaintStepTest_PathExt_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString321` into `intoString194`.

	// Assume that `sourceCQL` has the underlying type of `fromString321`:
	fromString321 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString321` to result `intoString194`
	// (`intoString194` is now tainted).
	intoString194 := path.Ext(fromString321)

	// Sink the tainted `intoString194`:
	sink(intoString194)
}

func TaintStepTest_PathJoin_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString487` into `intoString465`.

	// Assume that `sourceCQL` has the underlying type of `fromString487`:
	fromString487 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString487` to result `intoString465`
	// (`intoString465` is now tainted).
	intoString465 := path.Join(fromString487)

	// Sink the tainted `intoString465`:
	sink(intoString465)
}

func TaintStepTest_PathSplit_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString406` into `intoString222`.

	// Assume that `sourceCQL` has the underlying type of `fromString406`:
	fromString406 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString406` to result `intoString222`
	// (`intoString222` is now tainted).
	intoString222, _ := path.Split(fromString406)

	// Sink the tainted `intoString222`:
	sink(intoString222)
}

func TaintStepTest_PathSplit_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromString605` into `intoString398`.

	// Assume that `sourceCQL` has the underlying type of `fromString605`:
	fromString605 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString605` to result `intoString398`
	// (`intoString398` is now tainted).
	_, intoString398 := path.Split(fromString605)

	// Sink the tainted `intoString398`:
	sink(intoString398)
}

func RunAllTaints_Path() {
	{
		source := newSource()
		TaintStepTest_PathBase_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathClean_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathDir_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathExt_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathJoin_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathSplit_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathSplit_B0I0O1(source)
	}
}
