// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "unicode"

func TaintStepTest_UnicodeSimpleFold_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune656` into `intoRune414`.

	// Assume that `sourceCQL` has the underlying type of `fromRune656`:
	fromRune656 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune656` to result `intoRune414`
	// (`intoRune414` is now tainted).
	intoRune414 := unicode.SimpleFold(fromRune656)

	// Return the tainted `intoRune414`:
	return intoRune414
}

func TaintStepTest_UnicodeTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune518` into `intoRune650`.

	// Assume that `sourceCQL` has the underlying type of `fromRune518`:
	fromRune518 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune518` to result `intoRune650`
	// (`intoRune650` is now tainted).
	intoRune650 := unicode.To(0, fromRune518)

	// Return the tainted `intoRune650`:
	return intoRune650
}

func TaintStepTest_UnicodeToLower_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune784` into `intoRune957`.

	// Assume that `sourceCQL` has the underlying type of `fromRune784`:
	fromRune784 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune784` to result `intoRune957`
	// (`intoRune957` is now tainted).
	intoRune957 := unicode.ToLower(fromRune784)

	// Return the tainted `intoRune957`:
	return intoRune957
}

func TaintStepTest_UnicodeToTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune520` into `intoRune443`.

	// Assume that `sourceCQL` has the underlying type of `fromRune520`:
	fromRune520 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune520` to result `intoRune443`
	// (`intoRune443` is now tainted).
	intoRune443 := unicode.ToTitle(fromRune520)

	// Return the tainted `intoRune443`:
	return intoRune443
}

func TaintStepTest_UnicodeToUpper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune127` into `intoRune483`.

	// Assume that `sourceCQL` has the underlying type of `fromRune127`:
	fromRune127 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune127` to result `intoRune483`
	// (`intoRune483` is now tainted).
	intoRune483 := unicode.ToUpper(fromRune127)

	// Return the tainted `intoRune483`:
	return intoRune483
}

func TaintStepTest_UnicodeSpecialCaseToLower_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune989` into `intoRune982`.

	// Assume that `sourceCQL` has the underlying type of `fromRune989`:
	fromRune989 := sourceCQL.(rune)

	// Declare medium object/interface:
	var mediumObjCQL unicode.SpecialCase

	// Call the method that transfers the taint
	// from the parameter `fromRune989` to the result `intoRune982`
	// (`intoRune982` is now tainted).
	intoRune982 := mediumObjCQL.ToLower(fromRune989)

	// Return the tainted `intoRune982`:
	return intoRune982
}

func TaintStepTest_UnicodeSpecialCaseToTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune417` into `intoRune584`.

	// Assume that `sourceCQL` has the underlying type of `fromRune417`:
	fromRune417 := sourceCQL.(rune)

	// Declare medium object/interface:
	var mediumObjCQL unicode.SpecialCase

	// Call the method that transfers the taint
	// from the parameter `fromRune417` to the result `intoRune584`
	// (`intoRune584` is now tainted).
	intoRune584 := mediumObjCQL.ToTitle(fromRune417)

	// Return the tainted `intoRune584`:
	return intoRune584
}

func TaintStepTest_UnicodeSpecialCaseToUpper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune991` into `intoRune881`.

	// Assume that `sourceCQL` has the underlying type of `fromRune991`:
	fromRune991 := sourceCQL.(rune)

	// Declare medium object/interface:
	var mediumObjCQL unicode.SpecialCase

	// Call the method that transfers the taint
	// from the parameter `fromRune991` to the result `intoRune881`
	// (`intoRune881` is now tainted).
	intoRune881 := mediumObjCQL.ToUpper(fromRune991)

	// Return the tainted `intoRune881`:
	return intoRune881
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
