package main

import (
	"io"
	"regexp"
)

func TaintStepTest_RegexpQuoteMeta(sourceCQL interface{}) {
	// The flow is from `s` into `into273`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into273`
	// (`into273` is now tainted).
	into273 := regexp.QuoteMeta(s)

	// Sink the tainted `into273`:
	sink(into273)
}

func TaintStepTest_RegexpRegexpFind(sourceCQL interface{}) {
	// The flow is from `b` into `into538`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `b` to the result `into538`
	// (`into538` is now tainted).
	into538 := mediumObjCQL.Find(b)

	// Sink the tainted `into538`:
	sink(into538)
}

func TaintStepTest_RegexpRegexpFindAll(sourceCQL interface{}) {
	// The flow is from `b` into `into473`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `b` to the result `into473`
	// (`into473` is now tainted).
	into473 := mediumObjCQL.FindAll(b, 0)

	// Sink the tainted `into473`:
	sink(into473)
}

func TaintStepTest_RegexpRegexpFindAllIndex(sourceCQL interface{}) {
	// The flow is from `b` into `into404`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `b` to the result `into404`
	// (`into404` is now tainted).
	into404 := mediumObjCQL.FindAllIndex(b, 0)

	// Sink the tainted `into404`:
	sink(into404)
}

func TaintStepTest_RegexpRegexpFindAllString(sourceCQL interface{}) {
	// The flow is from `s` into `into644`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `s` to the result `into644`
	// (`into644` is now tainted).
	into644 := mediumObjCQL.FindAllString(s, 0)

	// Sink the tainted `into644`:
	sink(into644)
}

func TaintStepTest_RegexpRegexpFindAllStringIndex(sourceCQL interface{}) {
	// The flow is from `s` into `into482`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `s` to the result `into482`
	// (`into482` is now tainted).
	into482 := mediumObjCQL.FindAllStringIndex(s, 0)

	// Sink the tainted `into482`:
	sink(into482)
}

func TaintStepTest_RegexpRegexpFindAllStringSubmatch(sourceCQL interface{}) {
	// The flow is from `s` into `into579`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `s` to the result `into579`
	// (`into579` is now tainted).
	into579 := mediumObjCQL.FindAllStringSubmatch(s, 0)

	// Sink the tainted `into579`:
	sink(into579)
}

func TaintStepTest_RegexpRegexpFindAllStringSubmatchIndex(sourceCQL interface{}) {
	// The flow is from `s` into `into808`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `s` to the result `into808`
	// (`into808` is now tainted).
	into808 := mediumObjCQL.FindAllStringSubmatchIndex(s, 0)

	// Sink the tainted `into808`:
	sink(into808)
}

func TaintStepTest_RegexpRegexpFindAllSubmatch(sourceCQL interface{}) {
	// The flow is from `b` into `into322`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `b` to the result `into322`
	// (`into322` is now tainted).
	into322 := mediumObjCQL.FindAllSubmatch(b, 0)

	// Sink the tainted `into322`:
	sink(into322)
}

func TaintStepTest_RegexpRegexpFindAllSubmatchIndex(sourceCQL interface{}) {
	// The flow is from `b` into `into565`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `b` to the result `into565`
	// (`into565` is now tainted).
	into565 := mediumObjCQL.FindAllSubmatchIndex(b, 0)

	// Sink the tainted `into565`:
	sink(into565)
}

func TaintStepTest_RegexpRegexpFindIndex(sourceCQL interface{}) {
	// The flow is from `b` into `loc`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `b` to the result `loc`
	// (`loc` is now tainted).
	loc := mediumObjCQL.FindIndex(b)

	// Sink the tainted `loc`:
	sink(loc)
}

func TaintStepTest_RegexpRegexpFindReaderIndex(sourceCQL interface{}) {
	// The flow is from `r` into `loc`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.RuneReader)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `r` to the result `loc`
	// (`loc` is now tainted).
	loc := mediumObjCQL.FindReaderIndex(r)

	// Sink the tainted `loc`:
	sink(loc)
}

func TaintStepTest_RegexpRegexpFindReaderSubmatchIndex(sourceCQL interface{}) {
	// The flow is from `r` into `into208`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.RuneReader)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `r` to the result `into208`
	// (`into208` is now tainted).
	into208 := mediumObjCQL.FindReaderSubmatchIndex(r)

	// Sink the tainted `into208`:
	sink(into208)
}

func TaintStepTest_RegexpRegexpFindString(sourceCQL interface{}) {
	// The flow is from `s` into `into749`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `s` to the result `into749`
	// (`into749` is now tainted).
	into749 := mediumObjCQL.FindString(s)

	// Sink the tainted `into749`:
	sink(into749)
}

func TaintStepTest_RegexpRegexpFindStringIndex(sourceCQL interface{}) {
	// The flow is from `s` into `loc`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `s` to the result `loc`
	// (`loc` is now tainted).
	loc := mediumObjCQL.FindStringIndex(s)

	// Sink the tainted `loc`:
	sink(loc)
}

func TaintStepTest_RegexpRegexpFindStringSubmatch(sourceCQL interface{}) {
	// The flow is from `s` into `into216`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `s` to the result `into216`
	// (`into216` is now tainted).
	into216 := mediumObjCQL.FindStringSubmatch(s)

	// Sink the tainted `into216`:
	sink(into216)
}

func TaintStepTest_RegexpRegexpFindStringSubmatchIndex(sourceCQL interface{}) {
	// The flow is from `s` into `into440`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `s` to the result `into440`
	// (`into440` is now tainted).
	into440 := mediumObjCQL.FindStringSubmatchIndex(s)

	// Sink the tainted `into440`:
	sink(into440)
}

func TaintStepTest_RegexpRegexpFindSubmatch(sourceCQL interface{}) {
	// The flow is from `b` into `into184`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `b` to the result `into184`
	// (`into184` is now tainted).
	into184 := mediumObjCQL.FindSubmatch(b)

	// Sink the tainted `into184`:
	sink(into184)
}

func TaintStepTest_RegexpRegexpFindSubmatchIndex(sourceCQL interface{}) {
	// The flow is from `b` into `into989`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `b` to the result `into989`
	// (`into989` is now tainted).
	into989 := mediumObjCQL.FindSubmatchIndex(b)

	// Sink the tainted `into989`:
	sink(into989)
}

func TaintStepTest_RegexpRegexpReplaceAll(sourceCQL interface{}) {
	// The flow is from `src` into `into764`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `src` to the result `into764`
	// (`into764` is now tainted).
	into764 := mediumObjCQL.ReplaceAll(src, nil)

	// Sink the tainted `into764`:
	sink(into764)
}

func TaintStepTest_RegexpRegexpReplaceAllFunc(sourceCQL interface{}) {
	// The flow is from `src` into `into372`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `src` to the result `into372`
	// (`into372` is now tainted).
	into372 := mediumObjCQL.ReplaceAllFunc(src, nil)

	// Sink the tainted `into372`:
	sink(into372)
}

func TaintStepTest_RegexpRegexpReplaceAllLiteral(sourceCQL interface{}) {
	// The flow is from `src` into `into466`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `src` to the result `into466`
	// (`into466` is now tainted).
	into466 := mediumObjCQL.ReplaceAllLiteral(src, nil)

	// Sink the tainted `into466`:
	sink(into466)
}

func TaintStepTest_RegexpRegexpReplaceAllLiteralString(sourceCQL interface{}) {
	// The flow is from `src` into `into889`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `src` to the result `into889`
	// (`into889` is now tainted).
	into889 := mediumObjCQL.ReplaceAllLiteralString(src, "")

	// Sink the tainted `into889`:
	sink(into889)
}

func TaintStepTest_RegexpRegexpReplaceAllString(sourceCQL interface{}) {
	// The flow is from `src` into `into799`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `src` to the result `into799`
	// (`into799` is now tainted).
	into799 := mediumObjCQL.ReplaceAllString(src, "")

	// Sink the tainted `into799`:
	sink(into799)
}

func TaintStepTest_RegexpRegexpReplaceAllStringFunc(sourceCQL interface{}) {
	// The flow is from `src` into `into676`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `src` to the result `into676`
	// (`into676` is now tainted).
	into676 := mediumObjCQL.ReplaceAllStringFunc(src, nil)

	// Sink the tainted `into676`:
	sink(into676)
}

func TaintStepTest_RegexpRegexpSplit(sourceCQL interface{}) {
	// The flow is from `s` into `into361`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL regexp.Regexp

	// Call the method that transfers the taint
	// from the parameter `s` to the result `into361`
	// (`into361` is now tainted).
	into361 := mediumObjCQL.Split(s, 0)

	// Sink the tainted `into361`:
	sink(into361)
}

func RunAllTaints_Regexp(v interface{}) {
	{
		source := newSource()
		TaintStepTest_RegexpQuoteMeta(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFind(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindAll(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindAllIndex(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindAllString(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindAllStringIndex(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindAllStringSubmatch(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindAllStringSubmatchIndex(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindAllSubmatch(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindAllSubmatchIndex(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindIndex(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindReaderIndex(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindReaderSubmatchIndex(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindString(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindStringIndex(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindStringSubmatch(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindStringSubmatchIndex(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindSubmatch(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpFindSubmatchIndex(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAll(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAllFunc(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAllLiteral(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAllLiteralString(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAllString(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpReplaceAllStringFunc(source)
	}
	{
		source := newSource()
		TaintStepTest_RegexpRegexpSplit(source)
	}
}
