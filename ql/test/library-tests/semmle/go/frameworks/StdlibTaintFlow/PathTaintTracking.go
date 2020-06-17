// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "path"

func TaintStepTest_PathBase_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString163` into `intoString797`.

	// Assume that `sourceCQL` has the underlying type of `fromString163`:
	fromString163 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString163` to result `intoString797`
	// (`intoString797` is now tainted).
	intoString797 := path.Base(fromString163)

	// Return the tainted `intoString797`:
	return intoString797
}

func TaintStepTest_PathClean_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString959` into `intoString727`.

	// Assume that `sourceCQL` has the underlying type of `fromString959`:
	fromString959 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString959` to result `intoString727`
	// (`intoString727` is now tainted).
	intoString727 := path.Clean(fromString959)

	// Return the tainted `intoString727`:
	return intoString727
}

func TaintStepTest_PathDir_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString280` into `intoString920`.

	// Assume that `sourceCQL` has the underlying type of `fromString280`:
	fromString280 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString280` to result `intoString920`
	// (`intoString920` is now tainted).
	intoString920 := path.Dir(fromString280)

	// Return the tainted `intoString920`:
	return intoString920
}

func TaintStepTest_PathExt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString827` into `intoString992`.

	// Assume that `sourceCQL` has the underlying type of `fromString827`:
	fromString827 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString827` to result `intoString992`
	// (`intoString992` is now tainted).
	intoString992 := path.Ext(fromString827)

	// Return the tainted `intoString992`:
	return intoString992
}

func TaintStepTest_PathJoin_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString977` into `intoString595`.

	// Assume that `sourceCQL` has the underlying type of `fromString977`:
	fromString977 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString977` to result `intoString595`
	// (`intoString595` is now tainted).
	intoString595 := path.Join(fromString977)

	// Return the tainted `intoString595`:
	return intoString595
}

func TaintStepTest_PathSplit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString691` into `intoString170`.

	// Assume that `sourceCQL` has the underlying type of `fromString691`:
	fromString691 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString691` to result `intoString170`
	// (`intoString170` is now tainted).
	intoString170, _ := path.Split(fromString691)

	// Return the tainted `intoString170`:
	return intoString170
}

func TaintStepTest_PathSplit_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString309` into `intoString439`.

	// Assume that `sourceCQL` has the underlying type of `fromString309`:
	fromString309 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString309` to result `intoString439`
	// (`intoString439` is now tainted).
	_, intoString439 := path.Split(fromString309)

	// Return the tainted `intoString439`:
	return intoString439
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
