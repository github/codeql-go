// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"regexp"
)

func TaintStepTest_RegexpCompile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString748` into `intoRegexp210`.

	// Assume that `sourceCQL` has the underlying type of `fromString748`:
	fromString748 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString748` to result `intoRegexp210`
	// (`intoRegexp210` is now tainted).
	intoRegexp210, _ := regexp.Compile(fromString748)

	// Return the tainted `intoRegexp210`:
	return intoRegexp210
}

func TaintStepTest_RegexpCompilePOSIX_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString174` into `intoRegexp800`.

	// Assume that `sourceCQL` has the underlying type of `fromString174`:
	fromString174 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString174` to result `intoRegexp800`
	// (`intoRegexp800` is now tainted).
	intoRegexp800, _ := regexp.CompilePOSIX(fromString174)

	// Return the tainted `intoRegexp800`:
	return intoRegexp800
}

func TaintStepTest_RegexpMustCompile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString463` into `intoRegexp491`.

	// Assume that `sourceCQL` has the underlying type of `fromString463`:
	fromString463 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString463` to result `intoRegexp491`
	// (`intoRegexp491` is now tainted).
	intoRegexp491 := regexp.MustCompile(fromString463)

	// Return the tainted `intoRegexp491`:
	return intoRegexp491
}

func TaintStepTest_RegexpMustCompilePOSIX_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString754` into `intoRegexp284`.

	// Assume that `sourceCQL` has the underlying type of `fromString754`:
	fromString754 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString754` to result `intoRegexp284`
	// (`intoRegexp284` is now tainted).
	intoRegexp284 := regexp.MustCompilePOSIX(fromString754)

	// Return the tainted `intoRegexp284`:
	return intoRegexp284
}

func TaintStepTest_RegexpQuoteMeta_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString932` into `intoString919`.

	// Assume that `sourceCQL` has the underlying type of `fromString932`:
	fromString932 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString932` to result `intoString919`
	// (`intoString919` is now tainted).
	intoString919 := regexp.QuoteMeta(fromString932)

	// Return the tainted `intoString919`:
	return intoString919
}

func TaintStepTest_RegexpRegexpCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRegexp271` into `intoRegexp775`.

	// Assume that `sourceCQL` has the underlying type of `fromRegexp271`:
	fromRegexp271 := sourceCQL.(regexp.Regexp)

	// Call the method that transfers the taint
	// from the receiver `fromRegexp271` to the result `intoRegexp775`
	// (`intoRegexp775` is now tainted).
	intoRegexp775 := fromRegexp271.Copy()

	// Return the tainted `intoRegexp775`:
	return intoRegexp775
}

func TaintStepTest_RegexpRegexpExpand_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte886` into `intoByte943`.

	// Assume that `sourceCQL` has the underlying type of `fromByte886`:
	fromByte886 := sourceCQL.([]byte)

	// Declare `intoByte943` variable:
	var intoByte943 []byte

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte886` to the parameter `intoByte943`
	// (`intoByte943` is now tainted).
	mediumObjCQL.Expand(intoByte943, fromByte886, nil, nil)

	// Return the tainted `intoByte943`:
	return intoByte943
}

func TaintStepTest_RegexpRegexpExpand_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte560` into `intoByte302`.

	// Assume that `sourceCQL` has the underlying type of `fromByte560`:
	fromByte560 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte560` to the result `intoByte302`
	// (`intoByte302` is now tainted).
	intoByte302 := mediumObjCQL.Expand(nil, fromByte560, nil, nil)

	// Return the tainted `intoByte302`:
	return intoByte302
}

func TaintStepTest_RegexpRegexpExpand_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte136` into `intoByte246`.

	// Assume that `sourceCQL` has the underlying type of `fromByte136`:
	fromByte136 := sourceCQL.([]byte)

	// Declare `intoByte246` variable:
	var intoByte246 []byte

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte136` to the parameter `intoByte246`
	// (`intoByte246` is now tainted).
	mediumObjCQL.Expand(intoByte246, nil, fromByte136, nil)

	// Return the tainted `intoByte246`:
	return intoByte246
}

func TaintStepTest_RegexpRegexpExpand_B0I1O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte790` into `intoByte264`.

	// Assume that `sourceCQL` has the underlying type of `fromByte790`:
	fromByte790 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte790` to the result `intoByte264`
	// (`intoByte264` is now tainted).
	intoByte264 := mediumObjCQL.Expand(nil, nil, fromByte790, nil)

	// Return the tainted `intoByte264`:
	return intoByte264
}

func TaintStepTest_RegexpRegexpExpandString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString854` into `intoByte344`.

	// Assume that `sourceCQL` has the underlying type of `fromString854`:
	fromString854 := sourceCQL.(string)

	// Declare `intoByte344` variable:
	var intoByte344 []byte

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString854` to the parameter `intoByte344`
	// (`intoByte344` is now tainted).
	mediumObjCQL.ExpandString(intoByte344, fromString854, "", nil)

	// Return the tainted `intoByte344`:
	return intoByte344
}

func TaintStepTest_RegexpRegexpExpandString_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString556` into `intoByte654`.

	// Assume that `sourceCQL` has the underlying type of `fromString556`:
	fromString556 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString556` to the result `intoByte654`
	// (`intoByte654` is now tainted).
	intoByte654 := mediumObjCQL.ExpandString(nil, fromString556, "", nil)

	// Return the tainted `intoByte654`:
	return intoByte654
}

func TaintStepTest_RegexpRegexpExpandString_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString584` into `intoByte893`.

	// Assume that `sourceCQL` has the underlying type of `fromString584`:
	fromString584 := sourceCQL.(string)

	// Declare `intoByte893` variable:
	var intoByte893 []byte

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString584` to the parameter `intoByte893`
	// (`intoByte893` is now tainted).
	mediumObjCQL.ExpandString(intoByte893, "", fromString584, nil)

	// Return the tainted `intoByte893`:
	return intoByte893
}

func TaintStepTest_RegexpRegexpExpandString_B0I1O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString708` into `intoByte304`.

	// Assume that `sourceCQL` has the underlying type of `fromString708`:
	fromString708 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString708` to the result `intoByte304`
	// (`intoByte304` is now tainted).
	intoByte304 := mediumObjCQL.ExpandString(nil, "", fromString708, nil)

	// Return the tainted `intoByte304`:
	return intoByte304
}

func TaintStepTest_RegexpRegexpFind_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte599` into `intoByte129`.

	// Assume that `sourceCQL` has the underlying type of `fromByte599`:
	fromByte599 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte599` to the result `intoByte129`
	// (`intoByte129` is now tainted).
	intoByte129 := mediumObjCQL.Find(fromByte599)

	// Return the tainted `intoByte129`:
	return intoByte129
}

func TaintStepTest_RegexpRegexpFindAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte518` into `intoByte555`.

	// Assume that `sourceCQL` has the underlying type of `fromByte518`:
	fromByte518 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte518` to the result `intoByte555`
	// (`intoByte555` is now tainted).
	intoByte555 := mediumObjCQL.FindAll(fromByte518, 0)

	// Return the tainted `intoByte555`:
	return intoByte555
}

func TaintStepTest_RegexpRegexpFindAllIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte890` into `intoInt802`.

	// Assume that `sourceCQL` has the underlying type of `fromByte890`:
	fromByte890 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte890` to the result `intoInt802`
	// (`intoInt802` is now tainted).
	intoInt802 := mediumObjCQL.FindAllIndex(fromByte890, 0)

	// Return the tainted `intoInt802`:
	return intoInt802
}

func TaintStepTest_RegexpRegexpFindAllString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString165` into `intoString518`.

	// Assume that `sourceCQL` has the underlying type of `fromString165`:
	fromString165 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString165` to the result `intoString518`
	// (`intoString518` is now tainted).
	intoString518 := mediumObjCQL.FindAllString(fromString165, 0)

	// Return the tainted `intoString518`:
	return intoString518
}

func TaintStepTest_RegexpRegexpFindAllStringIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString807` into `intoInt945`.

	// Assume that `sourceCQL` has the underlying type of `fromString807`:
	fromString807 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString807` to the result `intoInt945`
	// (`intoInt945` is now tainted).
	intoInt945 := mediumObjCQL.FindAllStringIndex(fromString807, 0)

	// Return the tainted `intoInt945`:
	return intoInt945
}

func TaintStepTest_RegexpRegexpFindAllStringSubmatch_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString755` into `intoString730`.

	// Assume that `sourceCQL` has the underlying type of `fromString755`:
	fromString755 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString755` to the result `intoString730`
	// (`intoString730` is now tainted).
	intoString730 := mediumObjCQL.FindAllStringSubmatch(fromString755, 0)

	// Return the tainted `intoString730`:
	return intoString730
}

func TaintStepTest_RegexpRegexpFindAllStringSubmatchIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString828` into `intoInt297`.

	// Assume that `sourceCQL` has the underlying type of `fromString828`:
	fromString828 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString828` to the result `intoInt297`
	// (`intoInt297` is now tainted).
	intoInt297 := mediumObjCQL.FindAllStringSubmatchIndex(fromString828, 0)

	// Return the tainted `intoInt297`:
	return intoInt297
}

func TaintStepTest_RegexpRegexpFindAllSubmatch_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte915` into `intoByte401`.

	// Assume that `sourceCQL` has the underlying type of `fromByte915`:
	fromByte915 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte915` to the result `intoByte401`
	// (`intoByte401` is now tainted).
	intoByte401 := mediumObjCQL.FindAllSubmatch(fromByte915, 0)

	// Return the tainted `intoByte401`:
	return intoByte401
}

func TaintStepTest_RegexpRegexpFindAllSubmatchIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte555` into `intoInt338`.

	// Assume that `sourceCQL` has the underlying type of `fromByte555`:
	fromByte555 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte555` to the result `intoInt338`
	// (`intoInt338` is now tainted).
	intoInt338 := mediumObjCQL.FindAllSubmatchIndex(fromByte555, 0)

	// Return the tainted `intoInt338`:
	return intoInt338
}

func TaintStepTest_RegexpRegexpFindIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte810` into `intoInt511`.

	// Assume that `sourceCQL` has the underlying type of `fromByte810`:
	fromByte810 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte810` to the result `intoInt511`
	// (`intoInt511` is now tainted).
	intoInt511 := mediumObjCQL.FindIndex(fromByte810)

	// Return the tainted `intoInt511`:
	return intoInt511
}

func TaintStepTest_RegexpRegexpFindReaderIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRuneReader728` into `intoInt587`.

	// Assume that `sourceCQL` has the underlying type of `fromRuneReader728`:
	fromRuneReader728 := sourceCQL.(io.RuneReader)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromRuneReader728` to the result `intoInt587`
	// (`intoInt587` is now tainted).
	intoInt587 := mediumObjCQL.FindReaderIndex(fromRuneReader728)

	// Return the tainted `intoInt587`:
	return intoInt587
}

func TaintStepTest_RegexpRegexpFindReaderSubmatchIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRuneReader189` into `intoInt426`.

	// Assume that `sourceCQL` has the underlying type of `fromRuneReader189`:
	fromRuneReader189 := sourceCQL.(io.RuneReader)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromRuneReader189` to the result `intoInt426`
	// (`intoInt426` is now tainted).
	intoInt426 := mediumObjCQL.FindReaderSubmatchIndex(fromRuneReader189)

	// Return the tainted `intoInt426`:
	return intoInt426
}

func TaintStepTest_RegexpRegexpFindString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString830` into `intoString139`.

	// Assume that `sourceCQL` has the underlying type of `fromString830`:
	fromString830 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString830` to the result `intoString139`
	// (`intoString139` is now tainted).
	intoString139 := mediumObjCQL.FindString(fromString830)

	// Return the tainted `intoString139`:
	return intoString139
}

func TaintStepTest_RegexpRegexpFindStringIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString295` into `intoInt161`.

	// Assume that `sourceCQL` has the underlying type of `fromString295`:
	fromString295 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString295` to the result `intoInt161`
	// (`intoInt161` is now tainted).
	intoInt161 := mediumObjCQL.FindStringIndex(fromString295)

	// Return the tainted `intoInt161`:
	return intoInt161
}

func TaintStepTest_RegexpRegexpFindStringSubmatch_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString376` into `intoString496`.

	// Assume that `sourceCQL` has the underlying type of `fromString376`:
	fromString376 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString376` to the result `intoString496`
	// (`intoString496` is now tainted).
	intoString496 := mediumObjCQL.FindStringSubmatch(fromString376)

	// Return the tainted `intoString496`:
	return intoString496
}

func TaintStepTest_RegexpRegexpFindStringSubmatchIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString827` into `intoInt736`.

	// Assume that `sourceCQL` has the underlying type of `fromString827`:
	fromString827 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString827` to the result `intoInt736`
	// (`intoInt736` is now tainted).
	intoInt736 := mediumObjCQL.FindStringSubmatchIndex(fromString827)

	// Return the tainted `intoInt736`:
	return intoInt736
}

func TaintStepTest_RegexpRegexpFindSubmatch_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte960` into `intoByte345`.

	// Assume that `sourceCQL` has the underlying type of `fromByte960`:
	fromByte960 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte960` to the result `intoByte345`
	// (`intoByte345` is now tainted).
	intoByte345 := mediumObjCQL.FindSubmatch(fromByte960)

	// Return the tainted `intoByte345`:
	return intoByte345
}

func TaintStepTest_RegexpRegexpFindSubmatchIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte722` into `intoInt591`.

	// Assume that `sourceCQL` has the underlying type of `fromByte722`:
	fromByte722 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte722` to the result `intoInt591`
	// (`intoInt591` is now tainted).
	intoInt591 := mediumObjCQL.FindSubmatchIndex(fromByte722)

	// Return the tainted `intoInt591`:
	return intoInt591
}

func TaintStepTest_RegexpRegexpReplaceAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte366` into `intoByte499`.

	// Assume that `sourceCQL` has the underlying type of `fromByte366`:
	fromByte366 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte366` to the result `intoByte499`
	// (`intoByte499` is now tainted).
	intoByte499 := mediumObjCQL.ReplaceAll(fromByte366, nil)

	// Return the tainted `intoByte499`:
	return intoByte499
}

func TaintStepTest_RegexpRegexpReplaceAll_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte186` into `intoByte332`.

	// Assume that `sourceCQL` has the underlying type of `fromByte186`:
	fromByte186 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte186` to the result `intoByte332`
	// (`intoByte332` is now tainted).
	intoByte332 := mediumObjCQL.ReplaceAll(nil, fromByte186)

	// Return the tainted `intoByte332`:
	return intoByte332
}

func TaintStepTest_RegexpRegexpReplaceAllFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte382` into `intoByte447`.

	// Assume that `sourceCQL` has the underlying type of `fromByte382`:
	fromByte382 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte382` to the result `intoByte447`
	// (`intoByte447` is now tainted).
	intoByte447 := mediumObjCQL.ReplaceAllFunc(fromByte382, nil)

	// Return the tainted `intoByte447`:
	return intoByte447
}

func TaintStepTest_RegexpRegexpReplaceAllFunc_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFuncbytebyte366` into `intoByte317`.

	// Assume that `sourceCQL` has the underlying type of `fromFuncbytebyte366`:
	fromFuncbytebyte366 := sourceCQL.(func([]byte) []byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromFuncbytebyte366` to the result `intoByte317`
	// (`intoByte317` is now tainted).
	intoByte317 := mediumObjCQL.ReplaceAllFunc(nil, fromFuncbytebyte366)

	// Return the tainted `intoByte317`:
	return intoByte317
}

func TaintStepTest_RegexpRegexpReplaceAllLiteral_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte815` into `intoByte467`.

	// Assume that `sourceCQL` has the underlying type of `fromByte815`:
	fromByte815 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte815` to the result `intoByte467`
	// (`intoByte467` is now tainted).
	intoByte467 := mediumObjCQL.ReplaceAllLiteral(fromByte815, nil)

	// Return the tainted `intoByte467`:
	return intoByte467
}

func TaintStepTest_RegexpRegexpReplaceAllLiteral_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte133` into `intoByte429`.

	// Assume that `sourceCQL` has the underlying type of `fromByte133`:
	fromByte133 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte133` to the result `intoByte429`
	// (`intoByte429` is now tainted).
	intoByte429 := mediumObjCQL.ReplaceAllLiteral(nil, fromByte133)

	// Return the tainted `intoByte429`:
	return intoByte429
}

func TaintStepTest_RegexpRegexpReplaceAllLiteralString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString723` into `intoString896`.

	// Assume that `sourceCQL` has the underlying type of `fromString723`:
	fromString723 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString723` to the result `intoString896`
	// (`intoString896` is now tainted).
	intoString896 := mediumObjCQL.ReplaceAllLiteralString(fromString723, "")

	// Return the tainted `intoString896`:
	return intoString896
}

func TaintStepTest_RegexpRegexpReplaceAllLiteralString_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString474` into `intoString363`.

	// Assume that `sourceCQL` has the underlying type of `fromString474`:
	fromString474 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString474` to the result `intoString363`
	// (`intoString363` is now tainted).
	intoString363 := mediumObjCQL.ReplaceAllLiteralString("", fromString474)

	// Return the tainted `intoString363`:
	return intoString363
}

func TaintStepTest_RegexpRegexpReplaceAllString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString235` into `intoString121`.

	// Assume that `sourceCQL` has the underlying type of `fromString235`:
	fromString235 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString235` to the result `intoString121`
	// (`intoString121` is now tainted).
	intoString121 := mediumObjCQL.ReplaceAllString(fromString235, "")

	// Return the tainted `intoString121`:
	return intoString121
}

func TaintStepTest_RegexpRegexpReplaceAllString_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString640` into `intoString124`.

	// Assume that `sourceCQL` has the underlying type of `fromString640`:
	fromString640 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString640` to the result `intoString124`
	// (`intoString124` is now tainted).
	intoString124 := mediumObjCQL.ReplaceAllString("", fromString640)

	// Return the tainted `intoString124`:
	return intoString124
}

func TaintStepTest_RegexpRegexpReplaceAllStringFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString991` into `intoString559`.

	// Assume that `sourceCQL` has the underlying type of `fromString991`:
	fromString991 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString991` to the result `intoString559`
	// (`intoString559` is now tainted).
	intoString559 := mediumObjCQL.ReplaceAllStringFunc(fromString991, nil)

	// Return the tainted `intoString559`:
	return intoString559
}

func TaintStepTest_RegexpRegexpReplaceAllStringFunc_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFuncstringString654` into `intoString400`.

	// Assume that `sourceCQL` has the underlying type of `fromFuncstringString654`:
	fromFuncstringString654 := sourceCQL.(func(string) string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromFuncstringString654` to the result `intoString400`
	// (`intoString400` is now tainted).
	intoString400 := mediumObjCQL.ReplaceAllStringFunc("", fromFuncstringString654)

	// Return the tainted `intoString400`:
	return intoString400
}

func TaintStepTest_RegexpRegexpSplit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString551` into `intoString158`.

	// Assume that `sourceCQL` has the underlying type of `fromString551`:
	fromString551 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString551` to the result `intoString158`
	// (`intoString158` is now tainted).
	intoString158 := mediumObjCQL.Split(fromString551, 0)

	// Return the tainted `intoString158`:
	return intoString158
}

func RunAllTaints_Regexp() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpCompile_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpCompilePOSIX_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpMustCompile_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpMustCompilePOSIX_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpQuoteMeta_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpCopy_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpExpand_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpExpand_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpExpand_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpExpand_B0I1O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpExpandString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpExpandString_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpExpandString_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpExpandString_B0I1O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFind_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFindAll_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFindAllIndex_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFindAllString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFindAllStringIndex_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFindAllStringSubmatch_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFindAllStringSubmatchIndex_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFindAllSubmatch_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFindAllSubmatchIndex_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFindIndex_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFindReaderIndex_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFindReaderSubmatchIndex_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFindString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFindStringIndex_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFindStringSubmatch_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFindStringSubmatchIndex_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFindSubmatch_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpFindSubmatchIndex_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpReplaceAll_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpReplaceAll_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpReplaceAllFunc_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpReplaceAllFunc_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpReplaceAllLiteral_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpReplaceAllLiteral_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpReplaceAllLiteralString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpReplaceAllLiteralString_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpReplaceAllString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpReplaceAllString_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpReplaceAllStringFunc_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpReplaceAllStringFunc_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_RegexpRegexpSplit_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
