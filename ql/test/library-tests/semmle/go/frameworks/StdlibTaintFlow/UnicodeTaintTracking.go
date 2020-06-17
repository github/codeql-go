// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "unicode"

func TaintStepTest_UnicodeSimpleFold_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune934` into `intoRune937`.

	// Assume that `sourceCQL` has the underlying type of `fromRune934`:
	fromRune934 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune934` to result `intoRune937`
	// (`intoRune937` is now tainted).
	intoRune937 := unicode.SimpleFold(fromRune934)

	// Return the tainted `intoRune937`:
	return intoRune937
}

func TaintStepTest_UnicodeTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune681` into `intoRune797`.

	// Assume that `sourceCQL` has the underlying type of `fromRune681`:
	fromRune681 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune681` to result `intoRune797`
	// (`intoRune797` is now tainted).
	intoRune797 := unicode.To(0, fromRune681)

	// Return the tainted `intoRune797`:
	return intoRune797
}

func TaintStepTest_UnicodeToLower_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune454` into `intoRune720`.

	// Assume that `sourceCQL` has the underlying type of `fromRune454`:
	fromRune454 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune454` to result `intoRune720`
	// (`intoRune720` is now tainted).
	intoRune720 := unicode.ToLower(fromRune454)

	// Return the tainted `intoRune720`:
	return intoRune720
}

func TaintStepTest_UnicodeToTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune239` into `intoRune925`.

	// Assume that `sourceCQL` has the underlying type of `fromRune239`:
	fromRune239 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune239` to result `intoRune925`
	// (`intoRune925` is now tainted).
	intoRune925 := unicode.ToTitle(fromRune239)

	// Return the tainted `intoRune925`:
	return intoRune925
}

func TaintStepTest_UnicodeToUpper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune972` into `intoRune769`.

	// Assume that `sourceCQL` has the underlying type of `fromRune972`:
	fromRune972 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune972` to result `intoRune769`
	// (`intoRune769` is now tainted).
	intoRune769 := unicode.ToUpper(fromRune972)

	// Return the tainted `intoRune769`:
	return intoRune769
}

func TaintStepTest_UnicodeSpecialCaseToLower_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune914` into `intoRune703`.

	// Assume that `sourceCQL` has the underlying type of `fromRune914`:
	fromRune914 := sourceCQL.(rune)

	// Declare medium object/interface:
	var mediumObjCQL unicode.SpecialCase

	// Call the method that transfers the taint
	// from the parameter `fromRune914` to the result `intoRune703`
	// (`intoRune703` is now tainted).
	intoRune703 := mediumObjCQL.ToLower(fromRune914)

	// Return the tainted `intoRune703`:
	return intoRune703
}

func TaintStepTest_UnicodeSpecialCaseToTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune476` into `intoRune305`.

	// Assume that `sourceCQL` has the underlying type of `fromRune476`:
	fromRune476 := sourceCQL.(rune)

	// Declare medium object/interface:
	var mediumObjCQL unicode.SpecialCase

	// Call the method that transfers the taint
	// from the parameter `fromRune476` to the result `intoRune305`
	// (`intoRune305` is now tainted).
	intoRune305 := mediumObjCQL.ToTitle(fromRune476)

	// Return the tainted `intoRune305`:
	return intoRune305
}

func TaintStepTest_UnicodeSpecialCaseToUpper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune726` into `intoRune417`.

	// Assume that `sourceCQL` has the underlying type of `fromRune726`:
	fromRune726 := sourceCQL.(rune)

	// Declare medium object/interface:
	var mediumObjCQL unicode.SpecialCase

	// Call the method that transfers the taint
	// from the parameter `fromRune726` to the result `intoRune417`
	// (`intoRune417` is now tainted).
	intoRune417 := mediumObjCQL.ToUpper(fromRune726)

	// Return the tainted `intoRune417`:
	return intoRune417
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
