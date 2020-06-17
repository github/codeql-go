// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"regexp"
)

func TaintStepTest_RegexpCompile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString656` into `intoRegexp414`.

	// Assume that `sourceCQL` has the underlying type of `fromString656`:
	fromString656 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString656` to result `intoRegexp414`
	// (`intoRegexp414` is now tainted).
	intoRegexp414, _ := regexp.Compile(fromString656)

	// Return the tainted `intoRegexp414`:
	return intoRegexp414
}

func TaintStepTest_RegexpCompilePOSIX_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString518` into `intoRegexp650`.

	// Assume that `sourceCQL` has the underlying type of `fromString518`:
	fromString518 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString518` to result `intoRegexp650`
	// (`intoRegexp650` is now tainted).
	intoRegexp650, _ := regexp.CompilePOSIX(fromString518)

	// Return the tainted `intoRegexp650`:
	return intoRegexp650
}

func TaintStepTest_RegexpMustCompile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString784` into `intoRegexp957`.

	// Assume that `sourceCQL` has the underlying type of `fromString784`:
	fromString784 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString784` to result `intoRegexp957`
	// (`intoRegexp957` is now tainted).
	intoRegexp957 := regexp.MustCompile(fromString784)

	// Return the tainted `intoRegexp957`:
	return intoRegexp957
}

func TaintStepTest_RegexpMustCompilePOSIX_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString520` into `intoRegexp443`.

	// Assume that `sourceCQL` has the underlying type of `fromString520`:
	fromString520 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString520` to result `intoRegexp443`
	// (`intoRegexp443` is now tainted).
	intoRegexp443 := regexp.MustCompilePOSIX(fromString520)

	// Return the tainted `intoRegexp443`:
	return intoRegexp443
}

func TaintStepTest_RegexpQuoteMeta_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString127` into `intoString483`.

	// Assume that `sourceCQL` has the underlying type of `fromString127`:
	fromString127 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString127` to result `intoString483`
	// (`intoString483` is now tainted).
	intoString483 := regexp.QuoteMeta(fromString127)

	// Return the tainted `intoString483`:
	return intoString483
}

func TaintStepTest_RegexpRegexpCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRegexp989` into `intoRegexp982`.

	// Assume that `sourceCQL` has the underlying type of `fromRegexp989`:
	fromRegexp989 := sourceCQL.(regexp.Regexp)

	// Call the method that transfers the taint
	// from the receiver `fromRegexp989` to the result `intoRegexp982`
	// (`intoRegexp982` is now tainted).
	intoRegexp982 := fromRegexp989.Copy()

	// Return the tainted `intoRegexp982`:
	return intoRegexp982
}

func TaintStepTest_RegexpRegexpExpand_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte417` into `intoByte584`.

	// Assume that `sourceCQL` has the underlying type of `fromByte417`:
	fromByte417 := sourceCQL.([]byte)

	// Declare `intoByte584` variable:
	var intoByte584 []byte

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte417` to the parameter `intoByte584`
	// (`intoByte584` is now tainted).
	mediumObjCQL.Expand(intoByte584, fromByte417, nil, nil)

	// Return the tainted `intoByte584`:
	return intoByte584
}

func TaintStepTest_RegexpRegexpExpand_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte991` into `intoByte881`.

	// Assume that `sourceCQL` has the underlying type of `fromByte991`:
	fromByte991 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte991` to the result `intoByte881`
	// (`intoByte881` is now tainted).
	intoByte881 := mediumObjCQL.Expand(nil, fromByte991, nil, nil)

	// Return the tainted `intoByte881`:
	return intoByte881
}

func TaintStepTest_RegexpRegexpExpand_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte186` into `intoByte284`.

	// Assume that `sourceCQL` has the underlying type of `fromByte186`:
	fromByte186 := sourceCQL.([]byte)

	// Declare `intoByte284` variable:
	var intoByte284 []byte

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte186` to the parameter `intoByte284`
	// (`intoByte284` is now tainted).
	mediumObjCQL.Expand(intoByte284, nil, fromByte186, nil)

	// Return the tainted `intoByte284`:
	return intoByte284
}

func TaintStepTest_RegexpRegexpExpand_B0I1O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte908` into `intoByte137`.

	// Assume that `sourceCQL` has the underlying type of `fromByte908`:
	fromByte908 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte908` to the result `intoByte137`
	// (`intoByte137` is now tainted).
	intoByte137 := mediumObjCQL.Expand(nil, nil, fromByte908, nil)

	// Return the tainted `intoByte137`:
	return intoByte137
}

func TaintStepTest_RegexpRegexpExpandString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString494` into `intoByte873`.

	// Assume that `sourceCQL` has the underlying type of `fromString494`:
	fromString494 := sourceCQL.(string)

	// Declare `intoByte873` variable:
	var intoByte873 []byte

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString494` to the parameter `intoByte873`
	// (`intoByte873` is now tainted).
	mediumObjCQL.ExpandString(intoByte873, fromString494, "", nil)

	// Return the tainted `intoByte873`:
	return intoByte873
}

func TaintStepTest_RegexpRegexpExpandString_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString599` into `intoByte409`.

	// Assume that `sourceCQL` has the underlying type of `fromString599`:
	fromString599 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString599` to the result `intoByte409`
	// (`intoByte409` is now tainted).
	intoByte409 := mediumObjCQL.ExpandString(nil, fromString599, "", nil)

	// Return the tainted `intoByte409`:
	return intoByte409
}

func TaintStepTest_RegexpRegexpExpandString_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString246` into `intoByte898`.

	// Assume that `sourceCQL` has the underlying type of `fromString246`:
	fromString246 := sourceCQL.(string)

	// Declare `intoByte898` variable:
	var intoByte898 []byte

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString246` to the parameter `intoByte898`
	// (`intoByte898` is now tainted).
	mediumObjCQL.ExpandString(intoByte898, "", fromString246, nil)

	// Return the tainted `intoByte898`:
	return intoByte898
}

func TaintStepTest_RegexpRegexpExpandString_B0I1O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString598` into `intoByte631`.

	// Assume that `sourceCQL` has the underlying type of `fromString598`:
	fromString598 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString598` to the result `intoByte631`
	// (`intoByte631` is now tainted).
	intoByte631 := mediumObjCQL.ExpandString(nil, "", fromString598, nil)

	// Return the tainted `intoByte631`:
	return intoByte631
}

func TaintStepTest_RegexpRegexpFind_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte165` into `intoByte150`.

	// Assume that `sourceCQL` has the underlying type of `fromByte165`:
	fromByte165 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte165` to the result `intoByte150`
	// (`intoByte150` is now tainted).
	intoByte150 := mediumObjCQL.Find(fromByte165)

	// Return the tainted `intoByte150`:
	return intoByte150
}

func TaintStepTest_RegexpRegexpFindAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte340` into `intoByte471`.

	// Assume that `sourceCQL` has the underlying type of `fromByte340`:
	fromByte340 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte340` to the result `intoByte471`
	// (`intoByte471` is now tainted).
	intoByte471 := mediumObjCQL.FindAll(fromByte340, 0)

	// Return the tainted `intoByte471`:
	return intoByte471
}

func TaintStepTest_RegexpRegexpFindAllIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte290` into `intoInt758`.

	// Assume that `sourceCQL` has the underlying type of `fromByte290`:
	fromByte290 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte290` to the result `intoInt758`
	// (`intoInt758` is now tainted).
	intoInt758 := mediumObjCQL.FindAllIndex(fromByte290, 0)

	// Return the tainted `intoInt758`:
	return intoInt758
}

func TaintStepTest_RegexpRegexpFindAllString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString396` into `intoString707`.

	// Assume that `sourceCQL` has the underlying type of `fromString396`:
	fromString396 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString396` to the result `intoString707`
	// (`intoString707` is now tainted).
	intoString707 := mediumObjCQL.FindAllString(fromString396, 0)

	// Return the tainted `intoString707`:
	return intoString707
}

func TaintStepTest_RegexpRegexpFindAllStringIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString912` into `intoInt718`.

	// Assume that `sourceCQL` has the underlying type of `fromString912`:
	fromString912 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString912` to the result `intoInt718`
	// (`intoInt718` is now tainted).
	intoInt718 := mediumObjCQL.FindAllStringIndex(fromString912, 0)

	// Return the tainted `intoInt718`:
	return intoInt718
}

func TaintStepTest_RegexpRegexpFindAllStringSubmatch_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString972` into `intoString633`.

	// Assume that `sourceCQL` has the underlying type of `fromString972`:
	fromString972 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString972` to the result `intoString633`
	// (`intoString633` is now tainted).
	intoString633 := mediumObjCQL.FindAllStringSubmatch(fromString972, 0)

	// Return the tainted `intoString633`:
	return intoString633
}

func TaintStepTest_RegexpRegexpFindAllStringSubmatchIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString316` into `intoInt145`.

	// Assume that `sourceCQL` has the underlying type of `fromString316`:
	fromString316 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString316` to the result `intoInt145`
	// (`intoInt145` is now tainted).
	intoInt145 := mediumObjCQL.FindAllStringSubmatchIndex(fromString316, 0)

	// Return the tainted `intoInt145`:
	return intoInt145
}

func TaintStepTest_RegexpRegexpFindAllSubmatch_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte817` into `intoByte474`.

	// Assume that `sourceCQL` has the underlying type of `fromByte817`:
	fromByte817 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte817` to the result `intoByte474`
	// (`intoByte474` is now tainted).
	intoByte474 := mediumObjCQL.FindAllSubmatch(fromByte817, 0)

	// Return the tainted `intoByte474`:
	return intoByte474
}

func TaintStepTest_RegexpRegexpFindAllSubmatchIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte832` into `intoInt378`.

	// Assume that `sourceCQL` has the underlying type of `fromByte832`:
	fromByte832 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte832` to the result `intoInt378`
	// (`intoInt378` is now tainted).
	intoInt378 := mediumObjCQL.FindAllSubmatchIndex(fromByte832, 0)

	// Return the tainted `intoInt378`:
	return intoInt378
}

func TaintStepTest_RegexpRegexpFindIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte541` into `intoInt139`.

	// Assume that `sourceCQL` has the underlying type of `fromByte541`:
	fromByte541 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte541` to the result `intoInt139`
	// (`intoInt139` is now tainted).
	intoInt139 := mediumObjCQL.FindIndex(fromByte541)

	// Return the tainted `intoInt139`:
	return intoInt139
}

func TaintStepTest_RegexpRegexpFindReaderIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRuneReader814` into `intoInt768`.

	// Assume that `sourceCQL` has the underlying type of `fromRuneReader814`:
	fromRuneReader814 := sourceCQL.(io.RuneReader)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromRuneReader814` to the result `intoInt768`
	// (`intoInt768` is now tainted).
	intoInt768 := mediumObjCQL.FindReaderIndex(fromRuneReader814)

	// Return the tainted `intoInt768`:
	return intoInt768
}

func TaintStepTest_RegexpRegexpFindReaderSubmatchIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRuneReader468` into `intoInt736`.

	// Assume that `sourceCQL` has the underlying type of `fromRuneReader468`:
	fromRuneReader468 := sourceCQL.(io.RuneReader)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromRuneReader468` to the result `intoInt736`
	// (`intoInt736` is now tainted).
	intoInt736 := mediumObjCQL.FindReaderSubmatchIndex(fromRuneReader468)

	// Return the tainted `intoInt736`:
	return intoInt736
}

func TaintStepTest_RegexpRegexpFindString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString516` into `intoString246`.

	// Assume that `sourceCQL` has the underlying type of `fromString516`:
	fromString516 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString516` to the result `intoString246`
	// (`intoString246` is now tainted).
	intoString246 := mediumObjCQL.FindString(fromString516)

	// Return the tainted `intoString246`:
	return intoString246
}

func TaintStepTest_RegexpRegexpFindStringIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString679` into `intoInt736`.

	// Assume that `sourceCQL` has the underlying type of `fromString679`:
	fromString679 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString679` to the result `intoInt736`
	// (`intoInt736` is now tainted).
	intoInt736 := mediumObjCQL.FindStringIndex(fromString679)

	// Return the tainted `intoInt736`:
	return intoInt736
}

func TaintStepTest_RegexpRegexpFindStringSubmatch_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString839` into `intoString273`.

	// Assume that `sourceCQL` has the underlying type of `fromString839`:
	fromString839 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString839` to the result `intoString273`
	// (`intoString273` is now tainted).
	intoString273 := mediumObjCQL.FindStringSubmatch(fromString839)

	// Return the tainted `intoString273`:
	return intoString273
}

func TaintStepTest_RegexpRegexpFindStringSubmatchIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString982` into `intoInt458`.

	// Assume that `sourceCQL` has the underlying type of `fromString982`:
	fromString982 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString982` to the result `intoInt458`
	// (`intoInt458` is now tainted).
	intoInt458 := mediumObjCQL.FindStringSubmatchIndex(fromString982)

	// Return the tainted `intoInt458`:
	return intoInt458
}

func TaintStepTest_RegexpRegexpFindSubmatch_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte506` into `intoByte213`.

	// Assume that `sourceCQL` has the underlying type of `fromByte506`:
	fromByte506 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte506` to the result `intoByte213`
	// (`intoByte213` is now tainted).
	intoByte213 := mediumObjCQL.FindSubmatch(fromByte506)

	// Return the tainted `intoByte213`:
	return intoByte213
}

func TaintStepTest_RegexpRegexpFindSubmatchIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte468` into `intoInt219`.

	// Assume that `sourceCQL` has the underlying type of `fromByte468`:
	fromByte468 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte468` to the result `intoInt219`
	// (`intoInt219` is now tainted).
	intoInt219 := mediumObjCQL.FindSubmatchIndex(fromByte468)

	// Return the tainted `intoInt219`:
	return intoInt219
}

func TaintStepTest_RegexpRegexpReplaceAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte265` into `intoByte971`.

	// Assume that `sourceCQL` has the underlying type of `fromByte265`:
	fromByte265 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte265` to the result `intoByte971`
	// (`intoByte971` is now tainted).
	intoByte971 := mediumObjCQL.ReplaceAll(fromByte265, nil)

	// Return the tainted `intoByte971`:
	return intoByte971
}

func TaintStepTest_RegexpRegexpReplaceAll_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte320` into `intoByte545`.

	// Assume that `sourceCQL` has the underlying type of `fromByte320`:
	fromByte320 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte320` to the result `intoByte545`
	// (`intoByte545` is now tainted).
	intoByte545 := mediumObjCQL.ReplaceAll(nil, fromByte320)

	// Return the tainted `intoByte545`:
	return intoByte545
}

func TaintStepTest_RegexpRegexpReplaceAllFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte566` into `intoByte497`.

	// Assume that `sourceCQL` has the underlying type of `fromByte566`:
	fromByte566 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte566` to the result `intoByte497`
	// (`intoByte497` is now tainted).
	intoByte497 := mediumObjCQL.ReplaceAllFunc(fromByte566, nil)

	// Return the tainted `intoByte497`:
	return intoByte497
}

func TaintStepTest_RegexpRegexpReplaceAllFunc_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFuncbytebyte274` into `intoByte783`.

	// Assume that `sourceCQL` has the underlying type of `fromFuncbytebyte274`:
	fromFuncbytebyte274 := sourceCQL.(func([]byte) []byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromFuncbytebyte274` to the result `intoByte783`
	// (`intoByte783` is now tainted).
	intoByte783 := mediumObjCQL.ReplaceAllFunc(nil, fromFuncbytebyte274)

	// Return the tainted `intoByte783`:
	return intoByte783
}

func TaintStepTest_RegexpRegexpReplaceAllLiteral_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte905` into `intoByte389`.

	// Assume that `sourceCQL` has the underlying type of `fromByte905`:
	fromByte905 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte905` to the result `intoByte389`
	// (`intoByte389` is now tainted).
	intoByte389 := mediumObjCQL.ReplaceAllLiteral(fromByte905, nil)

	// Return the tainted `intoByte389`:
	return intoByte389
}

func TaintStepTest_RegexpRegexpReplaceAllLiteral_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte198` into `intoByte477`.

	// Assume that `sourceCQL` has the underlying type of `fromByte198`:
	fromByte198 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte198` to the result `intoByte477`
	// (`intoByte477` is now tainted).
	intoByte477 := mediumObjCQL.ReplaceAllLiteral(nil, fromByte198)

	// Return the tainted `intoByte477`:
	return intoByte477
}

func TaintStepTest_RegexpRegexpReplaceAllLiteralString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString544` into `intoString382`.

	// Assume that `sourceCQL` has the underlying type of `fromString544`:
	fromString544 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString544` to the result `intoString382`
	// (`intoString382` is now tainted).
	intoString382 := mediumObjCQL.ReplaceAllLiteralString(fromString544, "")

	// Return the tainted `intoString382`:
	return intoString382
}

func TaintStepTest_RegexpRegexpReplaceAllLiteralString_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString715` into `intoString179`.

	// Assume that `sourceCQL` has the underlying type of `fromString715`:
	fromString715 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString715` to the result `intoString179`
	// (`intoString179` is now tainted).
	intoString179 := mediumObjCQL.ReplaceAllLiteralString("", fromString715)

	// Return the tainted `intoString179`:
	return intoString179
}

func TaintStepTest_RegexpRegexpReplaceAllString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString366` into `intoString648`.

	// Assume that `sourceCQL` has the underlying type of `fromString366`:
	fromString366 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString366` to the result `intoString648`
	// (`intoString648` is now tainted).
	intoString648 := mediumObjCQL.ReplaceAllString(fromString366, "")

	// Return the tainted `intoString648`:
	return intoString648
}

func TaintStepTest_RegexpRegexpReplaceAllString_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString544` into `intoString484`.

	// Assume that `sourceCQL` has the underlying type of `fromString544`:
	fromString544 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString544` to the result `intoString484`
	// (`intoString484` is now tainted).
	intoString484 := mediumObjCQL.ReplaceAllString("", fromString544)

	// Return the tainted `intoString484`:
	return intoString484
}

func TaintStepTest_RegexpRegexpReplaceAllStringFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString824` into `intoString754`.

	// Assume that `sourceCQL` has the underlying type of `fromString824`:
	fromString824 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString824` to the result `intoString754`
	// (`intoString754` is now tainted).
	intoString754 := mediumObjCQL.ReplaceAllStringFunc(fromString824, nil)

	// Return the tainted `intoString754`:
	return intoString754
}

func TaintStepTest_RegexpRegexpReplaceAllStringFunc_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFuncstringString680` into `intoString722`.

	// Assume that `sourceCQL` has the underlying type of `fromFuncstringString680`:
	fromFuncstringString680 := sourceCQL.(func(string) string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromFuncstringString680` to the result `intoString722`
	// (`intoString722` is now tainted).
	intoString722 := mediumObjCQL.ReplaceAllStringFunc("", fromFuncstringString680)

	// Return the tainted `intoString722`:
	return intoString722
}

func TaintStepTest_RegexpRegexpSplit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString506` into `intoString121`.

	// Assume that `sourceCQL` has the underlying type of `fromString506`:
	fromString506 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString506` to the result `intoString121`
	// (`intoString121` is now tainted).
	intoString121 := mediumObjCQL.Split(fromString506, 0)

	// Return the tainted `intoString121`:
	return intoString121
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
