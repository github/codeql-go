// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "path"

func TaintStepTest_PathBase_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString656` into `intoString414`.

	// Assume that `sourceCQL` has the underlying type of `fromString656`:
	fromString656 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString656` to result `intoString414`
	// (`intoString414` is now tainted).
	intoString414 := path.Base(fromString656)

	// Return the tainted `intoString414`:
	return intoString414
}

func TaintStepTest_PathClean_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString518` into `intoString650`.

	// Assume that `sourceCQL` has the underlying type of `fromString518`:
	fromString518 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString518` to result `intoString650`
	// (`intoString650` is now tainted).
	intoString650 := path.Clean(fromString518)

	// Return the tainted `intoString650`:
	return intoString650
}

func TaintStepTest_PathDir_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString784` into `intoString957`.

	// Assume that `sourceCQL` has the underlying type of `fromString784`:
	fromString784 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString784` to result `intoString957`
	// (`intoString957` is now tainted).
	intoString957 := path.Dir(fromString784)

	// Return the tainted `intoString957`:
	return intoString957
}

func TaintStepTest_PathExt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString520` into `intoString443`.

	// Assume that `sourceCQL` has the underlying type of `fromString520`:
	fromString520 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString520` to result `intoString443`
	// (`intoString443` is now tainted).
	intoString443 := path.Ext(fromString520)

	// Return the tainted `intoString443`:
	return intoString443
}

func TaintStepTest_PathJoin_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString127` into `intoString483`.

	// Assume that `sourceCQL` has the underlying type of `fromString127`:
	fromString127 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString127` to result `intoString483`
	// (`intoString483` is now tainted).
	intoString483 := path.Join(fromString127)

	// Return the tainted `intoString483`:
	return intoString483
}

func TaintStepTest_PathSplit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString989` into `intoString982`.

	// Assume that `sourceCQL` has the underlying type of `fromString989`:
	fromString989 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString989` to result `intoString982`
	// (`intoString982` is now tainted).
	intoString982, _ := path.Split(fromString989)

	// Return the tainted `intoString982`:
	return intoString982
}

func TaintStepTest_PathSplit_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString417` into `intoString584`.

	// Assume that `sourceCQL` has the underlying type of `fromString417`:
	fromString417 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString417` to result `intoString584`
	// (`intoString584` is now tainted).
	_, intoString584 := path.Split(fromString417)

	// Return the tainted `intoString584`:
	return intoString584
}

func RunAllTaints_Path() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathBase_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathClean_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathDir_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathExt_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathJoin_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathSplit_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathSplit_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
