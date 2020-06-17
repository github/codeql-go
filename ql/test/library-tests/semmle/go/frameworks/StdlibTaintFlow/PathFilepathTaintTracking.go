package main

import "path/filepath"

func TaintStepTest_PathFilepathAbs_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString626` into `intoString562`.

	// Assume that `sourceCQL` has the underlying type of `fromString626`:
	fromString626 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString626` to result `intoString562`
	// (`intoString562` is now tainted).
	intoString562, _ := filepath.Abs(fromString626)

	// Return the tainted `intoString562`:
	return intoString562
}

func TaintStepTest_PathFilepathBase_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString317` into `intoString295`.

	// Assume that `sourceCQL` has the underlying type of `fromString317`:
	fromString317 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString317` to result `intoString295`
	// (`intoString295` is now tainted).
	intoString295 := filepath.Base(fromString317)

	// Return the tainted `intoString295`:
	return intoString295
}

func TaintStepTest_PathFilepathClean_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString349` into `intoString663`.

	// Assume that `sourceCQL` has the underlying type of `fromString349`:
	fromString349 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString349` to result `intoString663`
	// (`intoString663` is now tainted).
	intoString663 := filepath.Clean(fromString349)

	// Return the tainted `intoString663`:
	return intoString663
}

func TaintStepTest_PathFilepathDir_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString523` into `intoString315`.

	// Assume that `sourceCQL` has the underlying type of `fromString523`:
	fromString523 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString523` to result `intoString315`
	// (`intoString315` is now tainted).
	intoString315 := filepath.Dir(fromString523)

	// Return the tainted `intoString315`:
	return intoString315
}

func TaintStepTest_PathFilepathEvalSymlinks_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString179` into `intoString827`.

	// Assume that `sourceCQL` has the underlying type of `fromString179`:
	fromString179 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString179` to result `intoString827`
	// (`intoString827` is now tainted).
	intoString827, _ := filepath.EvalSymlinks(fromString179)

	// Return the tainted `intoString827`:
	return intoString827
}

func TaintStepTest_PathFilepathExt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString263` into `intoString356`.

	// Assume that `sourceCQL` has the underlying type of `fromString263`:
	fromString263 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString263` to result `intoString356`
	// (`intoString356` is now tainted).
	intoString356 := filepath.Ext(fromString263)

	// Return the tainted `intoString356`:
	return intoString356
}

func TaintStepTest_PathFilepathFromSlash_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString156` into `intoString235`.

	// Assume that `sourceCQL` has the underlying type of `fromString156`:
	fromString156 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString156` to result `intoString235`
	// (`intoString235` is now tainted).
	intoString235 := filepath.FromSlash(fromString156)

	// Return the tainted `intoString235`:
	return intoString235
}

func TaintStepTest_PathFilepathJoin_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString531` into `intoString243`.

	// Assume that `sourceCQL` has the underlying type of `fromString531`:
	fromString531 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString531` to result `intoString243`
	// (`intoString243` is now tainted).
	intoString243 := filepath.Join(fromString531)

	// Return the tainted `intoString243`:
	return intoString243
}

func TaintStepTest_PathFilepathRel_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString727` into `intoString640`.

	// Assume that `sourceCQL` has the underlying type of `fromString727`:
	fromString727 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString727` to result `intoString640`
	// (`intoString640` is now tainted).
	intoString640, _ := filepath.Rel(fromString727, "")

	// Return the tainted `intoString640`:
	return intoString640
}

func TaintStepTest_PathFilepathRel_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString528` into `intoString985`.

	// Assume that `sourceCQL` has the underlying type of `fromString528`:
	fromString528 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString528` to result `intoString985`
	// (`intoString985` is now tainted).
	intoString985, _ := filepath.Rel("", fromString528)

	// Return the tainted `intoString985`:
	return intoString985
}

func TaintStepTest_PathFilepathSplit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString189` into `intoString838`.

	// Assume that `sourceCQL` has the underlying type of `fromString189`:
	fromString189 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString189` to result `intoString838`
	// (`intoString838` is now tainted).
	intoString838, _ := filepath.Split(fromString189)

	// Return the tainted `intoString838`:
	return intoString838
}

func TaintStepTest_PathFilepathSplit_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString992` into `intoString742`.

	// Assume that `sourceCQL` has the underlying type of `fromString992`:
	fromString992 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString992` to result `intoString742`
	// (`intoString742` is now tainted).
	_, intoString742 := filepath.Split(fromString992)

	// Return the tainted `intoString742`:
	return intoString742
}

func TaintStepTest_PathFilepathSplitList_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString789` into `intoString904`.

	// Assume that `sourceCQL` has the underlying type of `fromString789`:
	fromString789 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString789` to result `intoString904`
	// (`intoString904` is now tainted).
	intoString904 := filepath.SplitList(fromString789)

	// Return the tainted `intoString904`:
	return intoString904
}

func TaintStepTest_PathFilepathToSlash_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString782` into `intoString196`.

	// Assume that `sourceCQL` has the underlying type of `fromString782`:
	fromString782 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString782` to result `intoString196`
	// (`intoString196` is now tainted).
	intoString196 := filepath.ToSlash(fromString782)

	// Return the tainted `intoString196`:
	return intoString196
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
