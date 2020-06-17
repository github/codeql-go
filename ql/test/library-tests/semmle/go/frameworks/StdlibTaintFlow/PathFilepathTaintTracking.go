// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "path/filepath"

func TaintStepTest_PathFilepathAbs_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString656` into `intoString414`.

	// Assume that `sourceCQL` has the underlying type of `fromString656`:
	fromString656 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString656` to result `intoString414`
	// (`intoString414` is now tainted).
	intoString414, _ := filepath.Abs(fromString656)

	// Return the tainted `intoString414`:
	return intoString414
}

func TaintStepTest_PathFilepathBase_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString518` into `intoString650`.

	// Assume that `sourceCQL` has the underlying type of `fromString518`:
	fromString518 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString518` to result `intoString650`
	// (`intoString650` is now tainted).
	intoString650 := filepath.Base(fromString518)

	// Return the tainted `intoString650`:
	return intoString650
}

func TaintStepTest_PathFilepathClean_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString784` into `intoString957`.

	// Assume that `sourceCQL` has the underlying type of `fromString784`:
	fromString784 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString784` to result `intoString957`
	// (`intoString957` is now tainted).
	intoString957 := filepath.Clean(fromString784)

	// Return the tainted `intoString957`:
	return intoString957
}

func TaintStepTest_PathFilepathDir_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString520` into `intoString443`.

	// Assume that `sourceCQL` has the underlying type of `fromString520`:
	fromString520 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString520` to result `intoString443`
	// (`intoString443` is now tainted).
	intoString443 := filepath.Dir(fromString520)

	// Return the tainted `intoString443`:
	return intoString443
}

func TaintStepTest_PathFilepathEvalSymlinks_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString127` into `intoString483`.

	// Assume that `sourceCQL` has the underlying type of `fromString127`:
	fromString127 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString127` to result `intoString483`
	// (`intoString483` is now tainted).
	intoString483, _ := filepath.EvalSymlinks(fromString127)

	// Return the tainted `intoString483`:
	return intoString483
}

func TaintStepTest_PathFilepathExt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString989` into `intoString982`.

	// Assume that `sourceCQL` has the underlying type of `fromString989`:
	fromString989 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString989` to result `intoString982`
	// (`intoString982` is now tainted).
	intoString982 := filepath.Ext(fromString989)

	// Return the tainted `intoString982`:
	return intoString982
}

func TaintStepTest_PathFilepathFromSlash_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString417` into `intoString584`.

	// Assume that `sourceCQL` has the underlying type of `fromString417`:
	fromString417 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString417` to result `intoString584`
	// (`intoString584` is now tainted).
	intoString584 := filepath.FromSlash(fromString417)

	// Return the tainted `intoString584`:
	return intoString584
}

func TaintStepTest_PathFilepathJoin_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString991` into `intoString881`.

	// Assume that `sourceCQL` has the underlying type of `fromString991`:
	fromString991 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString991` to result `intoString881`
	// (`intoString881` is now tainted).
	intoString881 := filepath.Join(fromString991)

	// Return the tainted `intoString881`:
	return intoString881
}

func TaintStepTest_PathFilepathRel_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString186` into `intoString284`.

	// Assume that `sourceCQL` has the underlying type of `fromString186`:
	fromString186 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString186` to result `intoString284`
	// (`intoString284` is now tainted).
	intoString284, _ := filepath.Rel(fromString186, "")

	// Return the tainted `intoString284`:
	return intoString284
}

func TaintStepTest_PathFilepathRel_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString908` into `intoString137`.

	// Assume that `sourceCQL` has the underlying type of `fromString908`:
	fromString908 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString908` to result `intoString137`
	// (`intoString137` is now tainted).
	intoString137, _ := filepath.Rel("", fromString908)

	// Return the tainted `intoString137`:
	return intoString137
}

func TaintStepTest_PathFilepathSplit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString494` into `intoString873`.

	// Assume that `sourceCQL` has the underlying type of `fromString494`:
	fromString494 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString494` to result `intoString873`
	// (`intoString873` is now tainted).
	intoString873, _ := filepath.Split(fromString494)

	// Return the tainted `intoString873`:
	return intoString873
}

func TaintStepTest_PathFilepathSplit_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString599` into `intoString409`.

	// Assume that `sourceCQL` has the underlying type of `fromString599`:
	fromString599 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString599` to result `intoString409`
	// (`intoString409` is now tainted).
	_, intoString409 := filepath.Split(fromString599)

	// Return the tainted `intoString409`:
	return intoString409
}

func TaintStepTest_PathFilepathSplitList_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString246` into `intoString898`.

	// Assume that `sourceCQL` has the underlying type of `fromString246`:
	fromString246 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString246` to result `intoString898`
	// (`intoString898` is now tainted).
	intoString898 := filepath.SplitList(fromString246)

	// Return the tainted `intoString898`:
	return intoString898
}

func TaintStepTest_PathFilepathToSlash_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString598` into `intoString631`.

	// Assume that `sourceCQL` has the underlying type of `fromString598`:
	fromString598 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString598` to result `intoString631`
	// (`intoString631` is now tainted).
	intoString631 := filepath.ToSlash(fromString598)

	// Return the tainted `intoString631`:
	return intoString631
}

func RunAllTaints_PathFilepath() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathFilepathAbs_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathFilepathBase_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathFilepathClean_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathFilepathDir_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathFilepathEvalSymlinks_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathFilepathExt_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathFilepathFromSlash_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathFilepathJoin_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathFilepathRel_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathFilepathRel_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathFilepathSplit_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathFilepathSplit_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathFilepathSplitList_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_PathFilepathToSlash_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
