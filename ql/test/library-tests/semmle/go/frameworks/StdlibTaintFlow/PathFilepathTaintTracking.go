package main

import "path/filepath"

func TaintStepTest_PathFilepathAbs_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString637` into `intoString625`.

	// Assume that `sourceCQL` has the underlying type of `fromString637`:
	fromString637 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString637` to result `intoString625`
	// (`intoString625` is now tainted).
	intoString625, _ := filepath.Abs(fromString637)

	// Sink the tainted `intoString625`:
	sink(intoString625)
}

func TaintStepTest_PathFilepathBase_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString270` into `intoString935`.

	// Assume that `sourceCQL` has the underlying type of `fromString270`:
	fromString270 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString270` to result `intoString935`
	// (`intoString935` is now tainted).
	intoString935 := filepath.Base(fromString270)

	// Sink the tainted `intoString935`:
	sink(intoString935)
}

func TaintStepTest_PathFilepathClean_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString963` into `intoString487`.

	// Assume that `sourceCQL` has the underlying type of `fromString963`:
	fromString963 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString963` to result `intoString487`
	// (`intoString487` is now tainted).
	intoString487 := filepath.Clean(fromString963)

	// Sink the tainted `intoString487`:
	sink(intoString487)
}

func TaintStepTest_PathFilepathDir_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString730` into `intoString889`.

	// Assume that `sourceCQL` has the underlying type of `fromString730`:
	fromString730 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString730` to result `intoString889`
	// (`intoString889` is now tainted).
	intoString889 := filepath.Dir(fromString730)

	// Sink the tainted `intoString889`:
	sink(intoString889)
}

func TaintStepTest_PathFilepathEvalSymlinks_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString205` into `intoString574`.

	// Assume that `sourceCQL` has the underlying type of `fromString205`:
	fromString205 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString205` to result `intoString574`
	// (`intoString574` is now tainted).
	intoString574, _ := filepath.EvalSymlinks(fromString205)

	// Sink the tainted `intoString574`:
	sink(intoString574)
}

func TaintStepTest_PathFilepathExt_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString161` into `intoString141`.

	// Assume that `sourceCQL` has the underlying type of `fromString161`:
	fromString161 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString161` to result `intoString141`
	// (`intoString141` is now tainted).
	intoString141 := filepath.Ext(fromString161)

	// Sink the tainted `intoString141`:
	sink(intoString141)
}

func TaintStepTest_PathFilepathFromSlash_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString690` into `intoString880`.

	// Assume that `sourceCQL` has the underlying type of `fromString690`:
	fromString690 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString690` to result `intoString880`
	// (`intoString880` is now tainted).
	intoString880 := filepath.FromSlash(fromString690)

	// Sink the tainted `intoString880`:
	sink(intoString880)
}

func TaintStepTest_PathFilepathJoin_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString649` into `intoString995`.

	// Assume that `sourceCQL` has the underlying type of `fromString649`:
	fromString649 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString649` to result `intoString995`
	// (`intoString995` is now tainted).
	intoString995 := filepath.Join(fromString649)

	// Sink the tainted `intoString995`:
	sink(intoString995)
}

func TaintStepTest_PathFilepathRel_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString397` into `intoString599`.

	// Assume that `sourceCQL` has the underlying type of `fromString397`:
	fromString397 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString397` to result `intoString599`
	// (`intoString599` is now tainted).
	intoString599, _ := filepath.Rel(fromString397, "")

	// Sink the tainted `intoString599`:
	sink(intoString599)
}

func TaintStepTest_PathFilepathRel_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString589` into `intoString840`.

	// Assume that `sourceCQL` has the underlying type of `fromString589`:
	fromString589 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString589` to result `intoString840`
	// (`intoString840` is now tainted).
	intoString840, _ := filepath.Rel("", fromString589)

	// Sink the tainted `intoString840`:
	sink(intoString840)
}

func TaintStepTest_PathFilepathSplit_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString963` into `intoString282`.

	// Assume that `sourceCQL` has the underlying type of `fromString963`:
	fromString963 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString963` to result `intoString282`
	// (`intoString282` is now tainted).
	intoString282, _ := filepath.Split(fromString963)

	// Sink the tainted `intoString282`:
	sink(intoString282)
}

func TaintStepTest_PathFilepathSplit_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromString688` into `intoString336`.

	// Assume that `sourceCQL` has the underlying type of `fromString688`:
	fromString688 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString688` to result `intoString336`
	// (`intoString336` is now tainted).
	_, intoString336 := filepath.Split(fromString688)

	// Sink the tainted `intoString336`:
	sink(intoString336)
}

func TaintStepTest_PathFilepathSplitList_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString969` into `intoString147`.

	// Assume that `sourceCQL` has the underlying type of `fromString969`:
	fromString969 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString969` to result `intoString147`
	// (`intoString147` is now tainted).
	intoString147 := filepath.SplitList(fromString969)

	// Sink the tainted `intoString147`:
	sink(intoString147)
}

func TaintStepTest_PathFilepathToSlash_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString594` into `intoString993`.

	// Assume that `sourceCQL` has the underlying type of `fromString594`:
	fromString594 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString594` to result `intoString993`
	// (`intoString993` is now tainted).
	intoString993 := filepath.ToSlash(fromString594)

	// Sink the tainted `intoString993`:
	sink(intoString993)
}

func RunAllTaints_PathFilepath() {
	{
		source := newSource()
		TaintStepTest_PathFilepathAbs_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathBase_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathClean_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathDir_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathEvalSymlinks_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathExt_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathFromSlash_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathJoin_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathRel_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathRel_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathSplit_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathSplit_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathSplitList_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_PathFilepathToSlash_B0I0O0(source)
	}
}
