package main

import "path"

func TaintStepTest_PathBase(sourceCQL interface{}) {
	// The flow is from `path` into `into276`.

	// Assume that `sourceCQL` has the underlying type of `path`:
	path123 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `path` to result `into276`
	// (`into276` is now tainted).
	into276 := path.Base(path123)

	// Sink the tainted `into276`:
	sink(into276)
}

func TaintStepTest_PathClean(sourceCQL interface{}) {
	// The flow is from `path` into `into591`.

	// Assume that `sourceCQL` has the underlying type of `path`:
	path123 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `path` to result `into591`
	// (`into591` is now tainted).
	into591 := path.Clean(path123)

	// Sink the tainted `into591`:
	sink(into591)
}

func TaintStepTest_PathDir(sourceCQL interface{}) {
	// The flow is from `path` into `into469`.

	// Assume that `sourceCQL` has the underlying type of `path`:
	path123 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `path` to result `into469`
	// (`into469` is now tainted).
	into469 := path.Dir(path123)

	// Sink the tainted `into469`:
	sink(into469)
}

func TaintStepTest_PathExt(sourceCQL interface{}) {
	// The flow is from `path` into `into761`.

	// Assume that `sourceCQL` has the underlying type of `path`:
	path123 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `path` to result `into761`
	// (`into761` is now tainted).
	into761 := path.Ext(path123)

	// Sink the tainted `into761`:
	sink(into761)
}

func TaintStepTest_PathJoin(sourceCQL interface{}) {
	// The flow is from `elem` into `into140`.

	// Assume that `sourceCQL` has the underlying type of `elem`:
	elem := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `elem` to result `into140`
	// (`into140` is now tainted).
	into140 := path.Join(elem)

	// Sink the tainted `into140`:
	sink(into140)
}

func TaintStepTest_PathSplit(sourceCQL interface{}) {
	// The flow is from `path` into `dir`.

	// Assume that `sourceCQL` has the underlying type of `path`:
	path123 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `path` to result `dir`
	// (`dir` is now tainted).
	dir, _ := path.Split(path123)

	// Sink the tainted `dir`:
	sink(dir)
}

func RunAllTaints_Path(v interface{}) {
	{
		source := newSource()
		TaintStepTest_PathBase(source)
	}
	{
		source := newSource()
		TaintStepTest_PathClean(source)
	}
	{
		source := newSource()
		TaintStepTest_PathDir(source)
	}
	{
		source := newSource()
		TaintStepTest_PathExt(source)
	}
	{
		source := newSource()
		TaintStepTest_PathJoin(source)
	}
	{
		source := newSource()
		TaintStepTest_PathSplit(source)
	}
}
