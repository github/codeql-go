package main

import "unicode"

func TaintStepTest_UnicodeSimpleFold_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune997` into `intoRune571`.

	// Assume that `sourceCQL` has the underlying type of `fromRune997`:
	fromRune997 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune997` to result `intoRune571`
	// (`intoRune571` is now tainted).
	intoRune571 := unicode.SimpleFold(fromRune997)

	// Return the tainted `intoRune571`:
	return intoRune571
}

func TaintStepTest_UnicodeTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune597` into `intoRune134`.

	// Assume that `sourceCQL` has the underlying type of `fromRune597`:
	fromRune597 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune597` to result `intoRune134`
	// (`intoRune134` is now tainted).
	intoRune134 := unicode.To(0, fromRune597)

	// Return the tainted `intoRune134`:
	return intoRune134
}

func TaintStepTest_UnicodeToLower_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune744` into `intoRune368`.

	// Assume that `sourceCQL` has the underlying type of `fromRune744`:
	fromRune744 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune744` to result `intoRune368`
	// (`intoRune368` is now tainted).
	intoRune368 := unicode.ToLower(fromRune744)

	// Return the tainted `intoRune368`:
	return intoRune368
}

func TaintStepTest_UnicodeToTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune152` into `intoRune971`.

	// Assume that `sourceCQL` has the underlying type of `fromRune152`:
	fromRune152 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune152` to result `intoRune971`
	// (`intoRune971` is now tainted).
	intoRune971 := unicode.ToTitle(fromRune152)

	// Return the tainted `intoRune971`:
	return intoRune971
}

func TaintStepTest_UnicodeToUpper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune414` into `intoRune338`.

	// Assume that `sourceCQL` has the underlying type of `fromRune414`:
	fromRune414 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune414` to result `intoRune338`
	// (`intoRune338` is now tainted).
	intoRune338 := unicode.ToUpper(fromRune414)

	// Return the tainted `intoRune338`:
	return intoRune338
}

func TaintStepTest_UnicodeSpecialCaseToLower_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune677` into `intoRune338`.

	// Assume that `sourceCQL` has the underlying type of `fromRune677`:
	fromRune677 := sourceCQL.(rune)

	// Declare medium object/interface:
	var mediumObjCQL unicode.SpecialCase

	// Call the method that transfers the taint
	// from the parameter `fromRune677` to the result `intoRune338`
	// (`intoRune338` is now tainted).
	intoRune338 := mediumObjCQL.ToLower(fromRune677)

	// Return the tainted `intoRune338`:
	return intoRune338
}

func TaintStepTest_UnicodeSpecialCaseToTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune576` into `intoRune898`.

	// Assume that `sourceCQL` has the underlying type of `fromRune576`:
	fromRune576 := sourceCQL.(rune)

	// Declare medium object/interface:
	var mediumObjCQL unicode.SpecialCase

	// Call the method that transfers the taint
	// from the parameter `fromRune576` to the result `intoRune898`
	// (`intoRune898` is now tainted).
	intoRune898 := mediumObjCQL.ToTitle(fromRune576)

	// Return the tainted `intoRune898`:
	return intoRune898
}

func TaintStepTest_UnicodeSpecialCaseToUpper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune374` into `intoRune918`.

	// Assume that `sourceCQL` has the underlying type of `fromRune374`:
	fromRune374 := sourceCQL.(rune)

	// Declare medium object/interface:
	var mediumObjCQL unicode.SpecialCase

	// Call the method that transfers the taint
	// from the parameter `fromRune374` to the result `intoRune918`
	// (`intoRune918` is now tainted).
	intoRune918 := mediumObjCQL.ToUpper(fromRune374)

	// Return the tainted `intoRune918`:
	return intoRune918
}

func RunAllTaints_Unicode() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeSimpleFold_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeTo_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeToLower_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeToTitle_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeToUpper_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeSpecialCaseToLower_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeSpecialCaseToTitle_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_UnicodeSpecialCaseToUpper_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
