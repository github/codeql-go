// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "path/filepath"

func TaintStepTest_PathFilepathAbs_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString530` into `intoString198`.

	// Assume that `sourceCQL` has the underlying type of `fromString530`:
	fromString530 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString530` to result `intoString198`
	// (`intoString198` is now tainted).
	intoString198, _ := filepath.Abs(fromString530)

	// Return the tainted `intoString198`:
	return intoString198
}

func TaintStepTest_PathFilepathBase_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString278` into `intoString628`.

	// Assume that `sourceCQL` has the underlying type of `fromString278`:
	fromString278 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString278` to result `intoString628`
	// (`intoString628` is now tainted).
	intoString628 := filepath.Base(fromString278)

	// Return the tainted `intoString628`:
	return intoString628
}

func TaintStepTest_PathFilepathClean_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString856` into `intoString847`.

	// Assume that `sourceCQL` has the underlying type of `fromString856`:
	fromString856 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString856` to result `intoString847`
	// (`intoString847` is now tainted).
	intoString847 := filepath.Clean(fromString856)

	// Return the tainted `intoString847`:
	return intoString847
}

func TaintStepTest_PathFilepathDir_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString467` into `intoString701`.

	// Assume that `sourceCQL` has the underlying type of `fromString467`:
	fromString467 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString467` to result `intoString701`
	// (`intoString701` is now tainted).
	intoString701 := filepath.Dir(fromString467)

	// Return the tainted `intoString701`:
	return intoString701
}

func TaintStepTest_PathFilepathEvalSymlinks_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString152` into `intoString718`.

	// Assume that `sourceCQL` has the underlying type of `fromString152`:
	fromString152 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString152` to result `intoString718`
	// (`intoString718` is now tainted).
	intoString718, _ := filepath.EvalSymlinks(fromString152)

	// Return the tainted `intoString718`:
	return intoString718
}

func TaintStepTest_PathFilepathExt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString312` into `intoString203`.

	// Assume that `sourceCQL` has the underlying type of `fromString312`:
	fromString312 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString312` to result `intoString203`
	// (`intoString203` is now tainted).
	intoString203 := filepath.Ext(fromString312)

	// Return the tainted `intoString203`:
	return intoString203
}

func TaintStepTest_PathFilepathFromSlash_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString523` into `intoString680`.

	// Assume that `sourceCQL` has the underlying type of `fromString523`:
	fromString523 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString523` to result `intoString680`
	// (`intoString680` is now tainted).
	intoString680 := filepath.FromSlash(fromString523)

	// Return the tainted `intoString680`:
	return intoString680
}

func TaintStepTest_PathFilepathJoin_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString935` into `intoString773`.

	// Assume that `sourceCQL` has the underlying type of `fromString935`:
	fromString935 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString935` to result `intoString773`
	// (`intoString773` is now tainted).
	intoString773 := filepath.Join(fromString935)

	// Return the tainted `intoString773`:
	return intoString773
}

func TaintStepTest_PathFilepathRel_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString289` into `intoString807`.

	// Assume that `sourceCQL` has the underlying type of `fromString289`:
	fromString289 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString289` to result `intoString807`
	// (`intoString807` is now tainted).
	intoString807, _ := filepath.Rel(fromString289, "")

	// Return the tainted `intoString807`:
	return intoString807
}

func TaintStepTest_PathFilepathRel_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString267` into `intoString896`.

	// Assume that `sourceCQL` has the underlying type of `fromString267`:
	fromString267 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString267` to result `intoString896`
	// (`intoString896` is now tainted).
	intoString896, _ := filepath.Rel("", fromString267)

	// Return the tainted `intoString896`:
	return intoString896
}

func TaintStepTest_PathFilepathSplit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString848` into `intoString258`.

	// Assume that `sourceCQL` has the underlying type of `fromString848`:
	fromString848 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString848` to result `intoString258`
	// (`intoString258` is now tainted).
	intoString258, _ := filepath.Split(fromString848)

	// Return the tainted `intoString258`:
	return intoString258
}

func TaintStepTest_PathFilepathSplit_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString584` into `intoString732`.

	// Assume that `sourceCQL` has the underlying type of `fromString584`:
	fromString584 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString584` to result `intoString732`
	// (`intoString732` is now tainted).
	_, intoString732 := filepath.Split(fromString584)

	// Return the tainted `intoString732`:
	return intoString732
}

func TaintStepTest_PathFilepathSplitList_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString451` into `intoString313`.

	// Assume that `sourceCQL` has the underlying type of `fromString451`:
	fromString451 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString451` to result `intoString313`
	// (`intoString313` is now tainted).
	intoString313 := filepath.SplitList(fromString451)

	// Return the tainted `intoString313`:
	return intoString313
}

func TaintStepTest_PathFilepathToSlash_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString623` into `intoString223`.

	// Assume that `sourceCQL` has the underlying type of `fromString623`:
	fromString623 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString623` to result `intoString223`
	// (`intoString223` is now tainted).
	intoString223 := filepath.ToSlash(fromString623)

	// Return the tainted `intoString223`:
	return intoString223
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
