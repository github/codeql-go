package main

import (
	"io"
	"regexp"
)

func TaintStepTest_RegexpCompile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString352` into `intoRegexp640`.

	// Assume that `sourceCQL` has the underlying type of `fromString352`:
	fromString352 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString352` to result `intoRegexp640`
	// (`intoRegexp640` is now tainted).
	intoRegexp640, _ := regexp.Compile(fromString352)

	// Return the tainted `intoRegexp640`:
	return intoRegexp640
}

func TaintStepTest_RegexpCompilePOSIX_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString827` into `intoRegexp712`.

	// Assume that `sourceCQL` has the underlying type of `fromString827`:
	fromString827 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString827` to result `intoRegexp712`
	// (`intoRegexp712` is now tainted).
	intoRegexp712, _ := regexp.CompilePOSIX(fromString827)

	// Return the tainted `intoRegexp712`:
	return intoRegexp712
}

func TaintStepTest_RegexpMustCompile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString517` into `intoRegexp531`.

	// Assume that `sourceCQL` has the underlying type of `fromString517`:
	fromString517 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString517` to result `intoRegexp531`
	// (`intoRegexp531` is now tainted).
	intoRegexp531 := regexp.MustCompile(fromString517)

	// Return the tainted `intoRegexp531`:
	return intoRegexp531
}

func TaintStepTest_RegexpMustCompilePOSIX_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString956` into `intoRegexp980`.

	// Assume that `sourceCQL` has the underlying type of `fromString956`:
	fromString956 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString956` to result `intoRegexp980`
	// (`intoRegexp980` is now tainted).
	intoRegexp980 := regexp.MustCompilePOSIX(fromString956)

	// Return the tainted `intoRegexp980`:
	return intoRegexp980
}

func TaintStepTest_RegexpQuoteMeta_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString600` into `intoString900`.

	// Assume that `sourceCQL` has the underlying type of `fromString600`:
	fromString600 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString600` to result `intoString900`
	// (`intoString900` is now tainted).
	intoString900 := regexp.QuoteMeta(fromString600)

	// Return the tainted `intoString900`:
	return intoString900
}

func TaintStepTest_RegexpRegexpCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRegexp289` into `intoRegexp472`.

	// Assume that `sourceCQL` has the underlying type of `fromRegexp289`:
	fromRegexp289 := sourceCQL.(regexp.Regexp)

	// Call the method that transfers the taint
	// from the receiver `fromRegexp289` to the result `intoRegexp472`
	// (`intoRegexp472` is now tainted).
	intoRegexp472 := fromRegexp289.Copy()

	// Return the tainted `intoRegexp472`:
	return intoRegexp472
}

func TaintStepTest_RegexpRegexpExpand_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte698` into `intoByte303`.

	// Assume that `sourceCQL` has the underlying type of `fromByte698`:
	fromByte698 := sourceCQL.([]byte)

	// Declare `intoByte303` variable:
	var intoByte303 []byte

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte698` to the parameter `intoByte303`
	// (`intoByte303` is now tainted).
	mediumObjCQL.Expand(intoByte303, fromByte698, nil, nil)

	// Return the tainted `intoByte303`:
	return intoByte303
}

func TaintStepTest_RegexpRegexpExpand_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte780` into `intoByte659`.

	// Assume that `sourceCQL` has the underlying type of `fromByte780`:
	fromByte780 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte780` to the result `intoByte659`
	// (`intoByte659` is now tainted).
	intoByte659 := mediumObjCQL.Expand(nil, fromByte780, nil, nil)

	// Return the tainted `intoByte659`:
	return intoByte659
}

func TaintStepTest_RegexpRegexpExpand_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte420` into `intoByte142`.

	// Assume that `sourceCQL` has the underlying type of `fromByte420`:
	fromByte420 := sourceCQL.([]byte)

	// Declare `intoByte142` variable:
	var intoByte142 []byte

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte420` to the parameter `intoByte142`
	// (`intoByte142` is now tainted).
	mediumObjCQL.Expand(intoByte142, nil, fromByte420, nil)

	// Return the tainted `intoByte142`:
	return intoByte142
}

func TaintStepTest_RegexpRegexpExpand_B0I1O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte983` into `intoByte281`.

	// Assume that `sourceCQL` has the underlying type of `fromByte983`:
	fromByte983 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte983` to the result `intoByte281`
	// (`intoByte281` is now tainted).
	intoByte281 := mediumObjCQL.Expand(nil, nil, fromByte983, nil)

	// Return the tainted `intoByte281`:
	return intoByte281
}

func TaintStepTest_RegexpRegexpExpandString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString664` into `intoByte601`.

	// Assume that `sourceCQL` has the underlying type of `fromString664`:
	fromString664 := sourceCQL.(string)

	// Declare `intoByte601` variable:
	var intoByte601 []byte

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString664` to the parameter `intoByte601`
	// (`intoByte601` is now tainted).
	mediumObjCQL.ExpandString(intoByte601, fromString664, "", nil)

	// Return the tainted `intoByte601`:
	return intoByte601
}

func TaintStepTest_RegexpRegexpExpandString_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString383` into `intoByte928`.

	// Assume that `sourceCQL` has the underlying type of `fromString383`:
	fromString383 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString383` to the result `intoByte928`
	// (`intoByte928` is now tainted).
	intoByte928 := mediumObjCQL.ExpandString(nil, fromString383, "", nil)

	// Return the tainted `intoByte928`:
	return intoByte928
}

func TaintStepTest_RegexpRegexpExpandString_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString328` into `intoByte113`.

	// Assume that `sourceCQL` has the underlying type of `fromString328`:
	fromString328 := sourceCQL.(string)

	// Declare `intoByte113` variable:
	var intoByte113 []byte

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString328` to the parameter `intoByte113`
	// (`intoByte113` is now tainted).
	mediumObjCQL.ExpandString(intoByte113, "", fromString328, nil)

	// Return the tainted `intoByte113`:
	return intoByte113
}

func TaintStepTest_RegexpRegexpExpandString_B0I1O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString613` into `intoByte619`.

	// Assume that `sourceCQL` has the underlying type of `fromString613`:
	fromString613 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString613` to the result `intoByte619`
	// (`intoByte619` is now tainted).
	intoByte619 := mediumObjCQL.ExpandString(nil, "", fromString613, nil)

	// Return the tainted `intoByte619`:
	return intoByte619
}

func TaintStepTest_RegexpRegexpFind_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte138` into `intoByte226`.

	// Assume that `sourceCQL` has the underlying type of `fromByte138`:
	fromByte138 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte138` to the result `intoByte226`
	// (`intoByte226` is now tainted).
	intoByte226 := mediumObjCQL.Find(fromByte138)

	// Return the tainted `intoByte226`:
	return intoByte226
}

func TaintStepTest_RegexpRegexpFindAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte327` into `intoByte479`.

	// Assume that `sourceCQL` has the underlying type of `fromByte327`:
	fromByte327 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte327` to the result `intoByte479`
	// (`intoByte479` is now tainted).
	intoByte479 := mediumObjCQL.FindAll(fromByte327, 0)

	// Return the tainted `intoByte479`:
	return intoByte479
}

func TaintStepTest_RegexpRegexpFindAllIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte329` into `intoInt819`.

	// Assume that `sourceCQL` has the underlying type of `fromByte329`:
	fromByte329 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte329` to the result `intoInt819`
	// (`intoInt819` is now tainted).
	intoInt819 := mediumObjCQL.FindAllIndex(fromByte329, 0)

	// Return the tainted `intoInt819`:
	return intoInt819
}

func TaintStepTest_RegexpRegexpFindAllString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString957` into `intoString959`.

	// Assume that `sourceCQL` has the underlying type of `fromString957`:
	fromString957 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString957` to the result `intoString959`
	// (`intoString959` is now tainted).
	intoString959 := mediumObjCQL.FindAllString(fromString957, 0)

	// Return the tainted `intoString959`:
	return intoString959
}

func TaintStepTest_RegexpRegexpFindAllStringIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString722` into `intoInt437`.

	// Assume that `sourceCQL` has the underlying type of `fromString722`:
	fromString722 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString722` to the result `intoInt437`
	// (`intoInt437` is now tainted).
	intoInt437 := mediumObjCQL.FindAllStringIndex(fromString722, 0)

	// Return the tainted `intoInt437`:
	return intoInt437
}

func TaintStepTest_RegexpRegexpFindAllStringSubmatch_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString808` into `intoString798`.

	// Assume that `sourceCQL` has the underlying type of `fromString808`:
	fromString808 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString808` to the result `intoString798`
	// (`intoString798` is now tainted).
	intoString798 := mediumObjCQL.FindAllStringSubmatch(fromString808, 0)

	// Return the tainted `intoString798`:
	return intoString798
}

func TaintStepTest_RegexpRegexpFindAllStringSubmatchIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString131` into `intoInt654`.

	// Assume that `sourceCQL` has the underlying type of `fromString131`:
	fromString131 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString131` to the result `intoInt654`
	// (`intoInt654` is now tainted).
	intoInt654 := mediumObjCQL.FindAllStringSubmatchIndex(fromString131, 0)

	// Return the tainted `intoInt654`:
	return intoInt654
}

func TaintStepTest_RegexpRegexpFindAllSubmatch_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte321` into `intoByte632`.

	// Assume that `sourceCQL` has the underlying type of `fromByte321`:
	fromByte321 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte321` to the result `intoByte632`
	// (`intoByte632` is now tainted).
	intoByte632 := mediumObjCQL.FindAllSubmatch(fromByte321, 0)

	// Return the tainted `intoByte632`:
	return intoByte632
}

func TaintStepTest_RegexpRegexpFindAllSubmatchIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte570` into `intoInt251`.

	// Assume that `sourceCQL` has the underlying type of `fromByte570`:
	fromByte570 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte570` to the result `intoInt251`
	// (`intoInt251` is now tainted).
	intoInt251 := mediumObjCQL.FindAllSubmatchIndex(fromByte570, 0)

	// Return the tainted `intoInt251`:
	return intoInt251
}

func TaintStepTest_RegexpRegexpFindIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte339` into `intoInt274`.

	// Assume that `sourceCQL` has the underlying type of `fromByte339`:
	fromByte339 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte339` to the result `intoInt274`
	// (`intoInt274` is now tainted).
	intoInt274 := mediumObjCQL.FindIndex(fromByte339)

	// Return the tainted `intoInt274`:
	return intoInt274
}

func TaintStepTest_RegexpRegexpFindReaderIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRuneReader825` into `intoInt287`.

	// Assume that `sourceCQL` has the underlying type of `fromRuneReader825`:
	fromRuneReader825 := sourceCQL.(io.RuneReader)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromRuneReader825` to the result `intoInt287`
	// (`intoInt287` is now tainted).
	intoInt287 := mediumObjCQL.FindReaderIndex(fromRuneReader825)

	// Return the tainted `intoInt287`:
	return intoInt287
}

func TaintStepTest_RegexpRegexpFindReaderSubmatchIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRuneReader319` into `intoInt703`.

	// Assume that `sourceCQL` has the underlying type of `fromRuneReader319`:
	fromRuneReader319 := sourceCQL.(io.RuneReader)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromRuneReader319` to the result `intoInt703`
	// (`intoInt703` is now tainted).
	intoInt703 := mediumObjCQL.FindReaderSubmatchIndex(fromRuneReader319)

	// Return the tainted `intoInt703`:
	return intoInt703
}

func TaintStepTest_RegexpRegexpFindString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString606` into `intoString867`.

	// Assume that `sourceCQL` has the underlying type of `fromString606`:
	fromString606 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString606` to the result `intoString867`
	// (`intoString867` is now tainted).
	intoString867 := mediumObjCQL.FindString(fromString606)

	// Return the tainted `intoString867`:
	return intoString867
}

func TaintStepTest_RegexpRegexpFindStringIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString312` into `intoInt240`.

	// Assume that `sourceCQL` has the underlying type of `fromString312`:
	fromString312 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString312` to the result `intoInt240`
	// (`intoInt240` is now tainted).
	intoInt240 := mediumObjCQL.FindStringIndex(fromString312)

	// Return the tainted `intoInt240`:
	return intoInt240
}

func TaintStepTest_RegexpRegexpFindStringSubmatch_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString718` into `intoString432`.

	// Assume that `sourceCQL` has the underlying type of `fromString718`:
	fromString718 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString718` to the result `intoString432`
	// (`intoString432` is now tainted).
	intoString432 := mediumObjCQL.FindStringSubmatch(fromString718)

	// Return the tainted `intoString432`:
	return intoString432
}

func TaintStepTest_RegexpRegexpFindStringSubmatchIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString339` into `intoInt530`.

	// Assume that `sourceCQL` has the underlying type of `fromString339`:
	fromString339 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString339` to the result `intoInt530`
	// (`intoInt530` is now tainted).
	intoInt530 := mediumObjCQL.FindStringSubmatchIndex(fromString339)

	// Return the tainted `intoInt530`:
	return intoInt530
}

func TaintStepTest_RegexpRegexpFindSubmatch_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte776` into `intoByte975`.

	// Assume that `sourceCQL` has the underlying type of `fromByte776`:
	fromByte776 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte776` to the result `intoByte975`
	// (`intoByte975` is now tainted).
	intoByte975 := mediumObjCQL.FindSubmatch(fromByte776)

	// Return the tainted `intoByte975`:
	return intoByte975
}

func TaintStepTest_RegexpRegexpFindSubmatchIndex_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte644` into `intoInt671`.

	// Assume that `sourceCQL` has the underlying type of `fromByte644`:
	fromByte644 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte644` to the result `intoInt671`
	// (`intoInt671` is now tainted).
	intoInt671 := mediumObjCQL.FindSubmatchIndex(fromByte644)

	// Return the tainted `intoInt671`:
	return intoInt671
}

func TaintStepTest_RegexpRegexpReplaceAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte460` into `intoByte420`.

	// Assume that `sourceCQL` has the underlying type of `fromByte460`:
	fromByte460 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte460` to the result `intoByte420`
	// (`intoByte420` is now tainted).
	intoByte420 := mediumObjCQL.ReplaceAll(fromByte460, nil)

	// Return the tainted `intoByte420`:
	return intoByte420
}

func TaintStepTest_RegexpRegexpReplaceAll_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte610` into `intoByte357`.

	// Assume that `sourceCQL` has the underlying type of `fromByte610`:
	fromByte610 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte610` to the result `intoByte357`
	// (`intoByte357` is now tainted).
	intoByte357 := mediumObjCQL.ReplaceAll(nil, fromByte610)

	// Return the tainted `intoByte357`:
	return intoByte357
}

func TaintStepTest_RegexpRegexpReplaceAllFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte537` into `intoByte715`.

	// Assume that `sourceCQL` has the underlying type of `fromByte537`:
	fromByte537 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte537` to the result `intoByte715`
	// (`intoByte715` is now tainted).
	intoByte715 := mediumObjCQL.ReplaceAllFunc(fromByte537, nil)

	// Return the tainted `intoByte715`:
	return intoByte715
}

func TaintStepTest_RegexpRegexpReplaceAllFunc_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFuncbytebyte746` into `intoByte112`.

	// Assume that `sourceCQL` has the underlying type of `fromFuncbytebyte746`:
	fromFuncbytebyte746 := sourceCQL.(func([]byte) []byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromFuncbytebyte746` to the result `intoByte112`
	// (`intoByte112` is now tainted).
	intoByte112 := mediumObjCQL.ReplaceAllFunc(nil, fromFuncbytebyte746)

	// Return the tainted `intoByte112`:
	return intoByte112
}

func TaintStepTest_RegexpRegexpReplaceAllLiteral_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte801` into `intoByte913`.

	// Assume that `sourceCQL` has the underlying type of `fromByte801`:
	fromByte801 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte801` to the result `intoByte913`
	// (`intoByte913` is now tainted).
	intoByte913 := mediumObjCQL.ReplaceAllLiteral(fromByte801, nil)

	// Return the tainted `intoByte913`:
	return intoByte913
}

func TaintStepTest_RegexpRegexpReplaceAllLiteral_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte596` into `intoByte330`.

	// Assume that `sourceCQL` has the underlying type of `fromByte596`:
	fromByte596 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte596` to the result `intoByte330`
	// (`intoByte330` is now tainted).
	intoByte330 := mediumObjCQL.ReplaceAllLiteral(nil, fromByte596)

	// Return the tainted `intoByte330`:
	return intoByte330
}

func TaintStepTest_RegexpRegexpReplaceAllLiteralString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString878` into `intoString308`.

	// Assume that `sourceCQL` has the underlying type of `fromString878`:
	fromString878 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString878` to the result `intoString308`
	// (`intoString308` is now tainted).
	intoString308 := mediumObjCQL.ReplaceAllLiteralString(fromString878, "")

	// Return the tainted `intoString308`:
	return intoString308
}

func TaintStepTest_RegexpRegexpReplaceAllLiteralString_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString661` into `intoString756`.

	// Assume that `sourceCQL` has the underlying type of `fromString661`:
	fromString661 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString661` to the result `intoString756`
	// (`intoString756` is now tainted).
	intoString756 := mediumObjCQL.ReplaceAllLiteralString("", fromString661)

	// Return the tainted `intoString756`:
	return intoString756
}

func TaintStepTest_RegexpRegexpReplaceAllString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString789` into `intoString905`.

	// Assume that `sourceCQL` has the underlying type of `fromString789`:
	fromString789 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString789` to the result `intoString905`
	// (`intoString905` is now tainted).
	intoString905 := mediumObjCQL.ReplaceAllString(fromString789, "")

	// Return the tainted `intoString905`:
	return intoString905
}

func TaintStepTest_RegexpRegexpReplaceAllString_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString635` into `intoString746`.

	// Assume that `sourceCQL` has the underlying type of `fromString635`:
	fromString635 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString635` to the result `intoString746`
	// (`intoString746` is now tainted).
	intoString746 := mediumObjCQL.ReplaceAllString("", fromString635)

	// Return the tainted `intoString746`:
	return intoString746
}

func TaintStepTest_RegexpRegexpReplaceAllStringFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString649` into `intoString546`.

	// Assume that `sourceCQL` has the underlying type of `fromString649`:
	fromString649 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString649` to the result `intoString546`
	// (`intoString546` is now tainted).
	intoString546 := mediumObjCQL.ReplaceAllStringFunc(fromString649, nil)

	// Return the tainted `intoString546`:
	return intoString546
}

func TaintStepTest_RegexpRegexpReplaceAllStringFunc_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFuncstringString269` into `intoString611`.

	// Assume that `sourceCQL` has the underlying type of `fromFuncstringString269`:
	fromFuncstringString269 := sourceCQL.(func(string) string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromFuncstringString269` to the result `intoString611`
	// (`intoString611` is now tainted).
	intoString611 := mediumObjCQL.ReplaceAllStringFunc("", fromFuncstringString269)

	// Return the tainted `intoString611`:
	return intoString611
}

func TaintStepTest_RegexpRegexpSplit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString813` into `intoString590`.

	// Assume that `sourceCQL` has the underlying type of `fromString813`:
	fromString813 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString813` to the result `intoString590`
	// (`intoString590` is now tainted).
	intoString590 := mediumObjCQL.Split(fromString813, 0)

	// Return the tainted `intoString590`:
	return intoString590
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
