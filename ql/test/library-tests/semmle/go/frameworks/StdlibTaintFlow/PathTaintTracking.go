package main

import "path"

func TaintStepTest_PathBase_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString137` into `intoString811`.

	// Assume that `sourceCQL` has the underlying type of `fromString137`:
	fromString137 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString137` to result `intoString811`
	// (`intoString811` is now tainted).
	intoString811 := path.Base(fromString137)

	// Return the tainted `intoString811`:
	return intoString811
}

func TaintStepTest_PathClean_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString243` into `intoString861`.

	// Assume that `sourceCQL` has the underlying type of `fromString243`:
	fromString243 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString243` to result `intoString861`
	// (`intoString861` is now tainted).
	intoString861 := path.Clean(fromString243)

	// Return the tainted `intoString861`:
	return intoString861
}

func TaintStepTest_PathDir_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString601` into `intoString944`.

	// Assume that `sourceCQL` has the underlying type of `fromString601`:
	fromString601 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString601` to result `intoString944`
	// (`intoString944` is now tainted).
	intoString944 := path.Dir(fromString601)

	// Return the tainted `intoString944`:
	return intoString944
}

func TaintStepTest_PathExt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString751` into `intoString855`.

	// Assume that `sourceCQL` has the underlying type of `fromString751`:
	fromString751 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString751` to result `intoString855`
	// (`intoString855` is now tainted).
	intoString855 := path.Ext(fromString751)

	// Return the tainted `intoString855`:
	return intoString855
}

func TaintStepTest_PathJoin_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString996` into `intoString393`.

	// Assume that `sourceCQL` has the underlying type of `fromString996`:
	fromString996 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString996` to result `intoString393`
	// (`intoString393` is now tainted).
	intoString393 := path.Join(fromString996)

	// Return the tainted `intoString393`:
	return intoString393
}

func TaintStepTest_PathSplit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString541` into `intoString823`.

	// Assume that `sourceCQL` has the underlying type of `fromString541`:
	fromString541 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString541` to result `intoString823`
	// (`intoString823` is now tainted).
	intoString823, _ := path.Split(fromString541)

	// Return the tainted `intoString823`:
	return intoString823
}

func TaintStepTest_PathSplit_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString508` into `intoString908`.

	// Assume that `sourceCQL` has the underlying type of `fromString508`:
	fromString508 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString508` to result `intoString908`
	// (`intoString908` is now tainted).
	_, intoString908 := path.Split(fromString508)

	// Return the tainted `intoString908`:
	return intoString908
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
