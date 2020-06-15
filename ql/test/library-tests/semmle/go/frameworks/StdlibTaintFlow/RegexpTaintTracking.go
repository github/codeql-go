package main

import (
	"io"
	"regexp"
)

func TaintStepTest_RegexpCompile_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString466` into `intoRegexp564`.

	// Assume that `sourceCQL` has the underlying type of `fromString466`:
	fromString466 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString466` to result `intoRegexp564`
	// (`intoRegexp564` is now tainted).
	intoRegexp564, _ := regexp.Compile(fromString466)

	// Sink the tainted `intoRegexp564`:
	sink(intoRegexp564)
}

func TaintStepTest_RegexpCompilePOSIX_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString783` into `intoRegexp183`.

	// Assume that `sourceCQL` has the underlying type of `fromString783`:
	fromString783 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString783` to result `intoRegexp183`
	// (`intoRegexp183` is now tainted).
	intoRegexp183, _ := regexp.CompilePOSIX(fromString783)

	// Sink the tainted `intoRegexp183`:
	sink(intoRegexp183)
}

func TaintStepTest_RegexpMustCompile_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString498` into `intoRegexp169`.

	// Assume that `sourceCQL` has the underlying type of `fromString498`:
	fromString498 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString498` to result `intoRegexp169`
	// (`intoRegexp169` is now tainted).
	intoRegexp169 := regexp.MustCompile(fromString498)

	// Sink the tainted `intoRegexp169`:
	sink(intoRegexp169)
}

func TaintStepTest_RegexpMustCompilePOSIX_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString353` into `intoRegexp391`.

	// Assume that `sourceCQL` has the underlying type of `fromString353`:
	fromString353 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString353` to result `intoRegexp391`
	// (`intoRegexp391` is now tainted).
	intoRegexp391 := regexp.MustCompilePOSIX(fromString353)

	// Sink the tainted `intoRegexp391`:
	sink(intoRegexp391)
}

func TaintStepTest_RegexpQuoteMeta_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString804` into `intoString699`.

	// Assume that `sourceCQL` has the underlying type of `fromString804`:
	fromString804 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString804` to result `intoString699`
	// (`intoString699` is now tainted).
	intoString699 := regexp.QuoteMeta(fromString804)

	// Sink the tainted `intoString699`:
	sink(intoString699)
}

func TaintStepTest_RegexpRegexpCopy_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRegexp396` into `intoRegexp238`.

	// Assume that `sourceCQL` has the underlying type of `fromRegexp396`:
	fromRegexp396 := sourceCQL.(regexp.Regexp)

	// Call the method that transfers the taint
	// from the receiver `fromRegexp396` to the result `intoRegexp238`
	// (`intoRegexp238` is now tainted).
	intoRegexp238 := fromRegexp396.Copy()

	// Sink the tainted `intoRegexp238`:
	sink(intoRegexp238)
}

func TaintStepTest_RegexpRegexpExpand_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte401` into `intoByte670`.

	// Assume that `sourceCQL` has the underlying type of `fromByte401`:
	fromByte401 := sourceCQL.([]byte)

	// Declare `intoByte670` variable:
	var intoByte670 []byte

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte401` to the parameter `intoByte670`
	// (`intoByte670` is now tainted).
	mediumObjCQL.Expand(intoByte670, fromByte401, nil, nil)

	// Sink the tainted `intoByte670`:
	sink(intoByte670)
}

func TaintStepTest_RegexpRegexpExpand_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromByte389` into `intoByte506`.

	// Assume that `sourceCQL` has the underlying type of `fromByte389`:
	fromByte389 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte389` to the result `intoByte506`
	// (`intoByte506` is now tainted).
	intoByte506 := mediumObjCQL.Expand(nil, fromByte389, nil, nil)

	// Sink the tainted `intoByte506`:
	sink(intoByte506)
}

func TaintStepTest_RegexpRegexpExpand_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromByte912` into `intoByte313`.

	// Assume that `sourceCQL` has the underlying type of `fromByte912`:
	fromByte912 := sourceCQL.([]byte)

	// Declare `intoByte313` variable:
	var intoByte313 []byte

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte912` to the parameter `intoByte313`
	// (`intoByte313` is now tainted).
	mediumObjCQL.Expand(intoByte313, nil, fromByte912, nil)

	// Sink the tainted `intoByte313`:
	sink(intoByte313)
}

func TaintStepTest_RegexpRegexpExpand_B0I1O1(sourceCQL interface{}) {
	// The flow is from `fromByte777` into `intoByte568`.

	// Assume that `sourceCQL` has the underlying type of `fromByte777`:
	fromByte777 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte777` to the result `intoByte568`
	// (`intoByte568` is now tainted).
	intoByte568 := mediumObjCQL.Expand(nil, nil, fromByte777, nil)

	// Sink the tainted `intoByte568`:
	sink(intoByte568)
}

func TaintStepTest_RegexpRegexpExpandString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString369` into `intoByte683`.

	// Assume that `sourceCQL` has the underlying type of `fromString369`:
	fromString369 := sourceCQL.(string)

	// Declare `intoByte683` variable:
	var intoByte683 []byte

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString369` to the parameter `intoByte683`
	// (`intoByte683` is now tainted).
	mediumObjCQL.ExpandString(intoByte683, fromString369, "", nil)

	// Sink the tainted `intoByte683`:
	sink(intoByte683)
}

func TaintStepTest_RegexpRegexpExpandString_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromString285` into `intoByte847`.

	// Assume that `sourceCQL` has the underlying type of `fromString285`:
	fromString285 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString285` to the result `intoByte847`
	// (`intoByte847` is now tainted).
	intoByte847 := mediumObjCQL.ExpandString(nil, fromString285, "", nil)

	// Sink the tainted `intoByte847`:
	sink(intoByte847)
}

func TaintStepTest_RegexpRegexpExpandString_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString643` into `intoByte942`.

	// Assume that `sourceCQL` has the underlying type of `fromString643`:
	fromString643 := sourceCQL.(string)

	// Declare `intoByte942` variable:
	var intoByte942 []byte

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString643` to the parameter `intoByte942`
	// (`intoByte942` is now tainted).
	mediumObjCQL.ExpandString(intoByte942, "", fromString643, nil)

	// Sink the tainted `intoByte942`:
	sink(intoByte942)
}

func TaintStepTest_RegexpRegexpExpandString_B0I1O1(sourceCQL interface{}) {
	// The flow is from `fromString964` into `intoByte659`.

	// Assume that `sourceCQL` has the underlying type of `fromString964`:
	fromString964 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString964` to the result `intoByte659`
	// (`intoByte659` is now tainted).
	intoByte659 := mediumObjCQL.ExpandString(nil, "", fromString964, nil)

	// Sink the tainted `intoByte659`:
	sink(intoByte659)
}

func TaintStepTest_RegexpRegexpFind_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte888` into `intoByte627`.

	// Assume that `sourceCQL` has the underlying type of `fromByte888`:
	fromByte888 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte888` to the result `intoByte627`
	// (`intoByte627` is now tainted).
	intoByte627 := mediumObjCQL.Find(fromByte888)

	// Sink the tainted `intoByte627`:
	sink(intoByte627)
}

func TaintStepTest_RegexpRegexpFindAll_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte848` into `intoByte285`.

	// Assume that `sourceCQL` has the underlying type of `fromByte848`:
	fromByte848 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte848` to the result `intoByte285`
	// (`intoByte285` is now tainted).
	intoByte285 := mediumObjCQL.FindAll(fromByte848, 0)

	// Sink the tainted `intoByte285`:
	sink(intoByte285)
}

func TaintStepTest_RegexpRegexpFindAllIndex_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte417` into `intoInt653`.

	// Assume that `sourceCQL` has the underlying type of `fromByte417`:
	fromByte417 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte417` to the result `intoInt653`
	// (`intoInt653` is now tainted).
	intoInt653 := mediumObjCQL.FindAllIndex(fromByte417, 0)

	// Sink the tainted `intoInt653`:
	sink(intoInt653)
}

func TaintStepTest_RegexpRegexpFindAllString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString150` into `intoString882`.

	// Assume that `sourceCQL` has the underlying type of `fromString150`:
	fromString150 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString150` to the result `intoString882`
	// (`intoString882` is now tainted).
	intoString882 := mediumObjCQL.FindAllString(fromString150, 0)

	// Sink the tainted `intoString882`:
	sink(intoString882)
}

func TaintStepTest_RegexpRegexpFindAllStringIndex_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString660` into `intoInt559`.

	// Assume that `sourceCQL` has the underlying type of `fromString660`:
	fromString660 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString660` to the result `intoInt559`
	// (`intoInt559` is now tainted).
	intoInt559 := mediumObjCQL.FindAllStringIndex(fromString660, 0)

	// Sink the tainted `intoInt559`:
	sink(intoInt559)
}

func TaintStepTest_RegexpRegexpFindAllStringSubmatch_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString423` into `intoString376`.

	// Assume that `sourceCQL` has the underlying type of `fromString423`:
	fromString423 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString423` to the result `intoString376`
	// (`intoString376` is now tainted).
	intoString376 := mediumObjCQL.FindAllStringSubmatch(fromString423, 0)

	// Sink the tainted `intoString376`:
	sink(intoString376)
}

func TaintStepTest_RegexpRegexpFindAllStringSubmatchIndex_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString862` into `intoInt502`.

	// Assume that `sourceCQL` has the underlying type of `fromString862`:
	fromString862 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString862` to the result `intoInt502`
	// (`intoInt502` is now tainted).
	intoInt502 := mediumObjCQL.FindAllStringSubmatchIndex(fromString862, 0)

	// Sink the tainted `intoInt502`:
	sink(intoInt502)
}

func TaintStepTest_RegexpRegexpFindAllSubmatch_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte644` into `intoByte181`.

	// Assume that `sourceCQL` has the underlying type of `fromByte644`:
	fromByte644 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte644` to the result `intoByte181`
	// (`intoByte181` is now tainted).
	intoByte181 := mediumObjCQL.FindAllSubmatch(fromByte644, 0)

	// Sink the tainted `intoByte181`:
	sink(intoByte181)
}

func TaintStepTest_RegexpRegexpFindAllSubmatchIndex_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte895` into `intoInt537`.

	// Assume that `sourceCQL` has the underlying type of `fromByte895`:
	fromByte895 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte895` to the result `intoInt537`
	// (`intoInt537` is now tainted).
	intoInt537 := mediumObjCQL.FindAllSubmatchIndex(fromByte895, 0)

	// Sink the tainted `intoInt537`:
	sink(intoInt537)
}

func TaintStepTest_RegexpRegexpFindIndex_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte927` into `intoInt271`.

	// Assume that `sourceCQL` has the underlying type of `fromByte927`:
	fromByte927 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte927` to the result `intoInt271`
	// (`intoInt271` is now tainted).
	intoInt271 := mediumObjCQL.FindIndex(fromByte927)

	// Sink the tainted `intoInt271`:
	sink(intoInt271)
}

func TaintStepTest_RegexpRegexpFindReaderIndex_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRuneReader729` into `intoInt578`.

	// Assume that `sourceCQL` has the underlying type of `fromRuneReader729`:
	fromRuneReader729 := sourceCQL.(io.RuneReader)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromRuneReader729` to the result `intoInt578`
	// (`intoInt578` is now tainted).
	intoInt578 := mediumObjCQL.FindReaderIndex(fromRuneReader729)

	// Sink the tainted `intoInt578`:
	sink(intoInt578)
}

func TaintStepTest_RegexpRegexpFindReaderSubmatchIndex_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRuneReader950` into `intoInt200`.

	// Assume that `sourceCQL` has the underlying type of `fromRuneReader950`:
	fromRuneReader950 := sourceCQL.(io.RuneReader)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromRuneReader950` to the result `intoInt200`
	// (`intoInt200` is now tainted).
	intoInt200 := mediumObjCQL.FindReaderSubmatchIndex(fromRuneReader950)

	// Sink the tainted `intoInt200`:
	sink(intoInt200)
}

func TaintStepTest_RegexpRegexpFindString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString879` into `intoString644`.

	// Assume that `sourceCQL` has the underlying type of `fromString879`:
	fromString879 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString879` to the result `intoString644`
	// (`intoString644` is now tainted).
	intoString644 := mediumObjCQL.FindString(fromString879)

	// Sink the tainted `intoString644`:
	sink(intoString644)
}

func TaintStepTest_RegexpRegexpFindStringIndex_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString313` into `intoInt288`.

	// Assume that `sourceCQL` has the underlying type of `fromString313`:
	fromString313 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString313` to the result `intoInt288`
	// (`intoInt288` is now tainted).
	intoInt288 := mediumObjCQL.FindStringIndex(fromString313)

	// Sink the tainted `intoInt288`:
	sink(intoInt288)
}

func TaintStepTest_RegexpRegexpFindStringSubmatch_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString232` into `intoString281`.

	// Assume that `sourceCQL` has the underlying type of `fromString232`:
	fromString232 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString232` to the result `intoString281`
	// (`intoString281` is now tainted).
	intoString281 := mediumObjCQL.FindStringSubmatch(fromString232)

	// Sink the tainted `intoString281`:
	sink(intoString281)
}

func TaintStepTest_RegexpRegexpFindStringSubmatchIndex_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString720` into `intoInt905`.

	// Assume that `sourceCQL` has the underlying type of `fromString720`:
	fromString720 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString720` to the result `intoInt905`
	// (`intoInt905` is now tainted).
	intoInt905 := mediumObjCQL.FindStringSubmatchIndex(fromString720)

	// Sink the tainted `intoInt905`:
	sink(intoInt905)
}

func TaintStepTest_RegexpRegexpFindSubmatch_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte684` into `intoByte344`.

	// Assume that `sourceCQL` has the underlying type of `fromByte684`:
	fromByte684 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte684` to the result `intoByte344`
	// (`intoByte344` is now tainted).
	intoByte344 := mediumObjCQL.FindSubmatch(fromByte684)

	// Sink the tainted `intoByte344`:
	sink(intoByte344)
}

func TaintStepTest_RegexpRegexpFindSubmatchIndex_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte588` into `intoInt850`.

	// Assume that `sourceCQL` has the underlying type of `fromByte588`:
	fromByte588 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte588` to the result `intoInt850`
	// (`intoInt850` is now tainted).
	intoInt850 := mediumObjCQL.FindSubmatchIndex(fromByte588)

	// Sink the tainted `intoInt850`:
	sink(intoInt850)
}

func TaintStepTest_RegexpRegexpReplaceAll_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte507` into `intoByte966`.

	// Assume that `sourceCQL` has the underlying type of `fromByte507`:
	fromByte507 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte507` to the result `intoByte966`
	// (`intoByte966` is now tainted).
	intoByte966 := mediumObjCQL.ReplaceAll(fromByte507, nil)

	// Sink the tainted `intoByte966`:
	sink(intoByte966)
}

func TaintStepTest_RegexpRegexpReplaceAll_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromByte413` into `intoByte641`.

	// Assume that `sourceCQL` has the underlying type of `fromByte413`:
	fromByte413 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte413` to the result `intoByte641`
	// (`intoByte641` is now tainted).
	intoByte641 := mediumObjCQL.ReplaceAll(nil, fromByte413)

	// Sink the tainted `intoByte641`:
	sink(intoByte641)
}

func TaintStepTest_RegexpRegexpReplaceAllFunc_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte838` into `intoByte445`.

	// Assume that `sourceCQL` has the underlying type of `fromByte838`:
	fromByte838 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte838` to the result `intoByte445`
	// (`intoByte445` is now tainted).
	intoByte445 := mediumObjCQL.ReplaceAllFunc(fromByte838, nil)

	// Sink the tainted `intoByte445`:
	sink(intoByte445)
}

func TaintStepTest_RegexpRegexpReplaceAllFunc_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromFuncbytebyte636` into `intoByte235`.

	// Assume that `sourceCQL` has the underlying type of `fromFuncbytebyte636`:
	fromFuncbytebyte636 := sourceCQL.(func([]byte) []byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromFuncbytebyte636` to the result `intoByte235`
	// (`intoByte235` is now tainted).
	intoByte235 := mediumObjCQL.ReplaceAllFunc(nil, fromFuncbytebyte636)

	// Sink the tainted `intoByte235`:
	sink(intoByte235)
}

func TaintStepTest_RegexpRegexpReplaceAllLiteral_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte761` into `intoByte335`.

	// Assume that `sourceCQL` has the underlying type of `fromByte761`:
	fromByte761 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte761` to the result `intoByte335`
	// (`intoByte335` is now tainted).
	intoByte335 := mediumObjCQL.ReplaceAllLiteral(fromByte761, nil)

	// Sink the tainted `intoByte335`:
	sink(intoByte335)
}

func TaintStepTest_RegexpRegexpReplaceAllLiteral_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromByte882` into `intoByte192`.

	// Assume that `sourceCQL` has the underlying type of `fromByte882`:
	fromByte882 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromByte882` to the result `intoByte192`
	// (`intoByte192` is now tainted).
	intoByte192 := mediumObjCQL.ReplaceAllLiteral(nil, fromByte882)

	// Sink the tainted `intoByte192`:
	sink(intoByte192)
}

func TaintStepTest_RegexpRegexpReplaceAllLiteralString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString696` into `intoString935`.

	// Assume that `sourceCQL` has the underlying type of `fromString696`:
	fromString696 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString696` to the result `intoString935`
	// (`intoString935` is now tainted).
	intoString935 := mediumObjCQL.ReplaceAllLiteralString(fromString696, "")

	// Sink the tainted `intoString935`:
	sink(intoString935)
}

func TaintStepTest_RegexpRegexpReplaceAllLiteralString_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString905` into `intoString503`.

	// Assume that `sourceCQL` has the underlying type of `fromString905`:
	fromString905 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString905` to the result `intoString503`
	// (`intoString503` is now tainted).
	intoString503 := mediumObjCQL.ReplaceAllLiteralString("", fromString905)

	// Sink the tainted `intoString503`:
	sink(intoString503)
}

func TaintStepTest_RegexpRegexpReplaceAllString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString579` into `intoString860`.

	// Assume that `sourceCQL` has the underlying type of `fromString579`:
	fromString579 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString579` to the result `intoString860`
	// (`intoString860` is now tainted).
	intoString860 := mediumObjCQL.ReplaceAllString(fromString579, "")

	// Sink the tainted `intoString860`:
	sink(intoString860)
}

func TaintStepTest_RegexpRegexpReplaceAllString_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString891` into `intoString115`.

	// Assume that `sourceCQL` has the underlying type of `fromString891`:
	fromString891 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString891` to the result `intoString115`
	// (`intoString115` is now tainted).
	intoString115 := mediumObjCQL.ReplaceAllString("", fromString891)

	// Sink the tainted `intoString115`:
	sink(intoString115)
}

func TaintStepTest_RegexpRegexpReplaceAllStringFunc_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString382` into `intoString782`.

	// Assume that `sourceCQL` has the underlying type of `fromString382`:
	fromString382 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString382` to the result `intoString782`
	// (`intoString782` is now tainted).
	intoString782 := mediumObjCQL.ReplaceAllStringFunc(fromString382, nil)

	// Sink the tainted `intoString782`:
	sink(intoString782)
}

func TaintStepTest_RegexpRegexpReplaceAllStringFunc_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromFuncstringString396` into `intoString969`.

	// Assume that `sourceCQL` has the underlying type of `fromFuncstringString396`:
	fromFuncstringString396 := sourceCQL.(func(string) string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromFuncstringString396` to the result `intoString969`
	// (`intoString969` is now tainted).
	intoString969 := mediumObjCQL.ReplaceAllStringFunc("", fromFuncstringString396)

	// Sink the tainted `intoString969`:
	sink(intoString969)
}

func TaintStepTest_RegexpRegexpSplit_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString382` into `intoString420`.

	// Assume that `sourceCQL` has the underlying type of `fromString382`:
	fromString382 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `fromString382` to the result `intoString420`
	// (`intoString420` is now tainted).
	intoString420 := mediumObjCQL.Split(fromString382, 0)

	// Sink the tainted `intoString420`:
	sink(intoString420)
}

func RunAllTaints_Regexp() {
	{
		source := newSource()
		TaintStepTest_RegexpCompile_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpCompilePOSIX_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpMustCompile_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpMustCompilePOSIX_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpQuoteMeta_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpCopy_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpExpand_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpExpand_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpExpand_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpExpand_B0I1O1(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpExpandString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpExpandString_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpExpandString_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpExpandString_B0I1O1(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFind_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindAll_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindAllIndex_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindAllString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindAllStringIndex_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindAllStringSubmatch_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindAllStringSubmatchIndex_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindAllSubmatch_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindAllSubmatchIndex_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindIndex_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindReaderIndex_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindReaderSubmatchIndex_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindStringIndex_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindStringSubmatch_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindStringSubmatchIndex_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindSubmatch_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindSubmatchIndex_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAll_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAll_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAllFunc_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAllFunc_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAllLiteral_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAllLiteral_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAllLiteralString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAllLiteralString_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAllString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAllString_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAllStringFunc_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAllStringFunc_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpSplit_B0I0O0(source)
	}
}
