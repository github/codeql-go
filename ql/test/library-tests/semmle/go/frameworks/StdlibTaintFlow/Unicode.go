package main

import "unicode"

func TaintStepTest_UnicodeSimpleFold(sourceCQL interface{}) {
	// The flow is from `r` into `into115`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into115`
	// (`into115` is now tainted).
	into115 := unicode.SimpleFold(r)

	// Sink the tainted `into115`:
	sink(into115)
}

func TaintStepTest_UnicodeTo(sourceCQL interface{}) {
	// The flow is from `r` into `into316`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into316`
	// (`into316` is now tainted).
	into316 := unicode.To(0, r)

	// Sink the tainted `into316`:
	sink(into316)
}

func TaintStepTest_UnicodeToLower(sourceCQL interface{}) {
	// The flow is from `r` into `into945`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into945`
	// (`into945` is now tainted).
	into945 := unicode.ToLower(r)

	// Sink the tainted `into945`:
	sink(into945)
}

func TaintStepTest_UnicodeToTitle(sourceCQL interface{}) {
	// The flow is from `r` into `into345`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into345`
	// (`into345` is now tainted).
	into345 := unicode.ToTitle(r)

	// Sink the tainted `into345`:
	sink(into345)
}

func TaintStepTest_UnicodeToUpper(sourceCQL interface{}) {
	// The flow is from `r` into `into231`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into231`
	// (`into231` is now tainted).
	into231 := unicode.ToUpper(r)

	// Sink the tainted `into231`:
	sink(into231)
}

func TaintStepTest_UnicodeSpecialCaseToLower(sourceCQL interface{}) {
	// The flow is from `r` into `into909`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(rune)

	// Declare medium object/interface:
	var mediumObjCQL unicode.SpecialCase

	// Call the method that transfers the taint
	// from the parameter `r` to the result `into909`
	// (`into909` is now tainted).
	into909 := mediumObjCQL.ToLower(r)

	// Sink the tainted `into909`:
	sink(into909)
}

func TaintStepTest_UnicodeSpecialCaseToTitle(sourceCQL interface{}) {
	// The flow is from `r` into `into214`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(rune)

	// Declare medium object/interface:
	var mediumObjCQL unicode.SpecialCase

	// Call the method that transfers the taint
	// from the parameter `r` to the result `into214`
	// (`into214` is now tainted).
	into214 := mediumObjCQL.ToTitle(r)

	// Sink the tainted `into214`:
	sink(into214)
}

func TaintStepTest_UnicodeSpecialCaseToUpper(sourceCQL interface{}) {
	// The flow is from `r` into `into775`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(rune)

	// Declare medium object/interface:
	var mediumObjCQL unicode.SpecialCase

	// Call the method that transfers the taint
	// from the parameter `r` to the result `into775`
	// (`into775` is now tainted).
	into775 := mediumObjCQL.ToUpper(r)

	// Sink the tainted `into775`:
	sink(into775)
}

func RunAllTaints_Unicode(v interface{}) {
	{
		source := newSource()
		TaintStepTest_UnicodeSimpleFold(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeTo(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeToLower(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeToTitle(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeToUpper(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeSpecialCaseToLower(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeSpecialCaseToTitle(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeSpecialCaseToUpper(source)
	}
}
