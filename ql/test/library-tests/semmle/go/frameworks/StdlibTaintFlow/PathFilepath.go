package main

import "path/filepath"

func TaintStepTest_PathFilepathAbs(sourceCQL interface{}) {
	// The flow is from `path` into `into873`.

	// Assume that `sourceCQL` has the underlying type of `path`:
	path := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `path` to result `into873`
	// (`into873` is now tainted).
	into873, _ := filepath.Abs(path)

	// Sink the tainted `into873`:
	sink(into873)
}

func TaintStepTest_PathFilepathBase(sourceCQL interface{}) {
	// The flow is from `path` into `into414`.

	// Assume that `sourceCQL` has the underlying type of `path`:
	path := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `path` to result `into414`
	// (`into414` is now tainted).
	into414 := filepath.Base(path)

	// Sink the tainted `into414`:
	sink(into414)
}

func TaintStepTest_PathFilepathClean(sourceCQL interface{}) {
	// The flow is from `path` into `into762`.

	// Assume that `sourceCQL` has the underlying type of `path`:
	path := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `path` to result `into762`
	// (`into762` is now tainted).
	into762 := filepath.Clean(path)

	// Sink the tainted `into762`:
	sink(into762)
}

func TaintStepTest_PathFilepathDir(sourceCQL interface{}) {
	// The flow is from `path` into `into742`.

	// Assume that `sourceCQL` has the underlying type of `path`:
	path := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `path` to result `into742`
	// (`into742` is now tainted).
	into742 := filepath.Dir(path)

	// Sink the tainted `into742`:
	sink(into742)
}

func TaintStepTest_PathFilepathEvalSymlinks(sourceCQL interface{}) {
	// The flow is from `path` into `into816`.

	// Assume that `sourceCQL` has the underlying type of `path`:
	path := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `path` to result `into816`
	// (`into816` is now tainted).
	into816, _ := filepath.EvalSymlinks(path)

	// Sink the tainted `into816`:
	sink(into816)
}

func TaintStepTest_PathFilepathExt(sourceCQL interface{}) {
	// The flow is from `path` into `into600`.

	// Assume that `sourceCQL` has the underlying type of `path`:
	path := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `path` to result `into600`
	// (`into600` is now tainted).
	into600 := filepath.Ext(path)

	// Sink the tainted `into600`:
	sink(into600)
}

func TaintStepTest_PathFilepathFromSlash(sourceCQL interface{}) {
	// The flow is from `path` into `into714`.

	// Assume that `sourceCQL` has the underlying type of `path`:
	path := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `path` to result `into714`
	// (`into714` is now tainted).
	into714 := filepath.FromSlash(path)

	// Sink the tainted `into714`:
	sink(into714)
}

func TaintStepTest_PathFilepathJoin(sourceCQL interface{}) {
	// The flow is from `elem` into `into772`.

	// Assume that `sourceCQL` has the underlying type of `elem`:
	elem := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `elem` to result `into772`
	// (`into772` is now tainted).
	into772 := filepath.Join(elem)

	// Sink the tainted `into772`:
	sink(into772)
}

func TaintStepTest_PathFilepathRel(sourceCQL interface{}) {
	// The flow is from `basepath` into `into189`.

	// Assume that `sourceCQL` has the underlying type of `basepath`:
	basepath := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `basepath` to result `into189`
	// (`into189` is now tainted).
	into189, _ := filepath.Rel(basepath, "")

	// Sink the tainted `into189`:
	sink(into189)
}

func TaintStepTest_PathFilepathSplit(sourceCQL interface{}) {
	// The flow is from `path` into `dir`.

	// Assume that `sourceCQL` has the underlying type of `path`:
	path := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `path` to result `dir`
	// (`dir` is now tainted).
	dir, _ := filepath.Split(path)

	// Sink the tainted `dir`:
	sink(dir)
}

func TaintStepTest_PathFilepathSplitList(sourceCQL interface{}) {
	// The flow is from `path` into `into441`.

	// Assume that `sourceCQL` has the underlying type of `path`:
	path := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `path` to result `into441`
	// (`into441` is now tainted).
	into441 := filepath.SplitList(path)

	// Sink the tainted `into441`:
	sink(into441)
}

func TaintStepTest_PathFilepathToSlash(sourceCQL interface{}) {
	// The flow is from `path` into `into305`.

	// Assume that `sourceCQL` has the underlying type of `path`:
	path := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `path` to result `into305`
	// (`into305` is now tainted).
	into305 := filepath.ToSlash(path)

	// Sink the tainted `into305`:
	sink(into305)
}

func RunAllTaints_PathFilepath(v interface{}) {
	{
		source := newSource()
		TaintStepTest_PathFilepathAbs(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathBase(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathClean(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathDir(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathEvalSymlinks(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathExt(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathFromSlash(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathJoin(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathRel(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathSplit(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathSplitList(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathToSlash(source)
	}
}
