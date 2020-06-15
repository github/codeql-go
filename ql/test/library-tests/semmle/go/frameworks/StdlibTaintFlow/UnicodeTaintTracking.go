package main

import "unicode"

func TaintStepTest_UnicodeSimpleFold_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune453` into `intoRune253`.

	// Assume that `sourceCQL` has the underlying type of `fromRune453`:
	fromRune453 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune453` to result `intoRune253`
	// (`intoRune253` is now tainted).
	intoRune253 := unicode.SimpleFold(fromRune453)

	// Sink the tainted `intoRune253`:
	sink(intoRune253)
}

func TaintStepTest_UnicodeTo_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune314` into `intoRune977`.

	// Assume that `sourceCQL` has the underlying type of `fromRune314`:
	fromRune314 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune314` to result `intoRune977`
	// (`intoRune977` is now tainted).
	intoRune977 := unicode.To(0, fromRune314)

	// Sink the tainted `intoRune977`:
	sink(intoRune977)
}

func TaintStepTest_UnicodeToLower_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune550` into `intoRune283`.

	// Assume that `sourceCQL` has the underlying type of `fromRune550`:
	fromRune550 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune550` to result `intoRune283`
	// (`intoRune283` is now tainted).
	intoRune283 := unicode.ToLower(fromRune550)

	// Sink the tainted `intoRune283`:
	sink(intoRune283)
}

func TaintStepTest_UnicodeToTitle_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune506` into `intoRune274`.

	// Assume that `sourceCQL` has the underlying type of `fromRune506`:
	fromRune506 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune506` to result `intoRune274`
	// (`intoRune274` is now tainted).
	intoRune274 := unicode.ToTitle(fromRune506)

	// Sink the tainted `intoRune274`:
	sink(intoRune274)
}

func TaintStepTest_UnicodeToUpper_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune393` into `intoRune408`.

	// Assume that `sourceCQL` has the underlying type of `fromRune393`:
	fromRune393 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune393` to result `intoRune408`
	// (`intoRune408` is now tainted).
	intoRune408 := unicode.ToUpper(fromRune393)

	// Sink the tainted `intoRune408`:
	sink(intoRune408)
}

func TaintStepTest_UnicodeSpecialCaseToLower_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune577` into `intoRune154`.

	// Assume that `sourceCQL` has the underlying type of `fromRune577`:
	fromRune577 := sourceCQL.(rune)

	// Declare medium object/interface:
	var mediumObjCQL unicode.SpecialCase

	// Call the method that transfers the taint
	// from the parameter `fromRune577` to the result `intoRune154`
	// (`intoRune154` is now tainted).
	intoRune154 := mediumObjCQL.ToLower(fromRune577)

	// Sink the tainted `intoRune154`:
	sink(intoRune154)
}

func TaintStepTest_UnicodeSpecialCaseToTitle_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune185` into `intoRune374`.

	// Assume that `sourceCQL` has the underlying type of `fromRune185`:
	fromRune185 := sourceCQL.(rune)

	// Declare medium object/interface:
	var mediumObjCQL unicode.SpecialCase

	// Call the method that transfers the taint
	// from the parameter `fromRune185` to the result `intoRune374`
	// (`intoRune374` is now tainted).
	intoRune374 := mediumObjCQL.ToTitle(fromRune185)

	// Sink the tainted `intoRune374`:
	sink(intoRune374)
}

func TaintStepTest_UnicodeSpecialCaseToUpper_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune344` into `intoRune481`.

	// Assume that `sourceCQL` has the underlying type of `fromRune344`:
	fromRune344 := sourceCQL.(rune)

	// Declare medium object/interface:
	var mediumObjCQL unicode.SpecialCase

	// Call the method that transfers the taint
	// from the parameter `fromRune344` to the result `intoRune481`
	// (`intoRune481` is now tainted).
	intoRune481 := mediumObjCQL.ToUpper(fromRune344)

	// Sink the tainted `intoRune481`:
	sink(intoRune481)
}

func RunAllTaints_Unicode() {
	{
		source := newSource()
		TaintStepTest_UnicodeSimpleFold_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeTo_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeToLower_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeToTitle_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeToUpper_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeSpecialCaseToLower_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeSpecialCaseToTitle_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeSpecialCaseToUpper_B0I0O0(source)
	}
}
