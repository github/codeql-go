package main

import (
	"io"
	"strings"
)

func TaintStepTest_StringsFields(sourceCQL interface{}) {
	// The flow is from `s` into `into397`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into397`
	// (`into397` is now tainted).
	into397 := strings.Fields(s)

	// Sink the tainted `into397`:
	sink(into397)
}

func TaintStepTest_StringsFieldsFunc(sourceCQL interface{}) {
	// The flow is from `s` into `into334`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into334`
	// (`into334` is now tainted).
	into334 := strings.FieldsFunc(s, nil)

	// Sink the tainted `into334`:
	sink(into334)
}

func TaintStepTest_StringsJoin(sourceCQL interface{}) {
	// The flow is from `elems` into `into455`.

	// Assume that `sourceCQL` has the underlying type of `elems`:
	elems := sourceCQL.([]string)

	// Call the function that transfers the taint
	// from the parameter `elems` to result `into455`
	// (`into455` is now tainted).
	into455 := strings.Join(elems, "")

	// Sink the tainted `into455`:
	sink(into455)
}

func TaintStepTest_StringsMap(sourceCQL interface{}) {
	// The flow is from `s` into `into286`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into286`
	// (`into286` is now tainted).
	into286 := strings.Map(nil, s)

	// Sink the tainted `into286`:
	sink(into286)
}

func TaintStepTest_StringsNewReader(sourceCQL interface{}) {
	// The flow is from `s` into `into290`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into290`
	// (`into290` is now tainted).
	into290 := strings.NewReader(s)

	// Sink the tainted `into290`:
	sink(into290)
}

func TaintStepTest_StringsNewReplacer(sourceCQL interface{}) {
	// The flow is from `oldnew` into `into628`.

	// Assume that `sourceCQL` has the underlying type of `oldnew`:
	oldnew := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `oldnew` to result `into628`
	// (`into628` is now tainted).
	into628 := strings.NewReplacer(oldnew)

	// Sink the tainted `into628`:
	sink(into628)
}

func TaintStepTest_StringsRepeat(sourceCQL interface{}) {
	// The flow is from `s` into `into903`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into903`
	// (`into903` is now tainted).
	into903 := strings.Repeat(s, 0)

	// Sink the tainted `into903`:
	sink(into903)
}

func TaintStepTest_StringsReplace(sourceCQL interface{}) {
	// The flow is from `s` into `into293`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into293`
	// (`into293` is now tainted).
	into293 := strings.Replace(s, "", "", 0)

	// Sink the tainted `into293`:
	sink(into293)
}

func TaintStepTest_StringsReplaceAll(sourceCQL interface{}) {
	// The flow is from `s` into `into830`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into830`
	// (`into830` is now tainted).
	into830 := strings.ReplaceAll(s, "", "")

	// Sink the tainted `into830`:
	sink(into830)
}

func TaintStepTest_StringsSplit(sourceCQL interface{}) {
	// The flow is from `s` into `into348`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into348`
	// (`into348` is now tainted).
	into348 := strings.Split(s, "")

	// Sink the tainted `into348`:
	sink(into348)
}

func TaintStepTest_StringsSplitAfter(sourceCQL interface{}) {
	// The flow is from `s` into `into867`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into867`
	// (`into867` is now tainted).
	into867 := strings.SplitAfter(s, "")

	// Sink the tainted `into867`:
	sink(into867)
}

func TaintStepTest_StringsSplitAfterN(sourceCQL interface{}) {
	// The flow is from `s` into `into193`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into193`
	// (`into193` is now tainted).
	into193 := strings.SplitAfterN(s, "", 0)

	// Sink the tainted `into193`:
	sink(into193)
}

func TaintStepTest_StringsSplitN(sourceCQL interface{}) {
	// The flow is from `s` into `into128`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into128`
	// (`into128` is now tainted).
	into128 := strings.SplitN(s, "", 0)

	// Sink the tainted `into128`:
	sink(into128)
}

func TaintStepTest_StringsTitle(sourceCQL interface{}) {
	// The flow is from `s` into `into763`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into763`
	// (`into763` is now tainted).
	into763 := strings.Title(s)

	// Sink the tainted `into763`:
	sink(into763)
}

func TaintStepTest_StringsToLower(sourceCQL interface{}) {
	// The flow is from `s` into `into915`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into915`
	// (`into915` is now tainted).
	into915 := strings.ToLower(s)

	// Sink the tainted `into915`:
	sink(into915)
}

func TaintStepTest_StringsToLowerSpecial(sourceCQL interface{}) {
	// The flow is from `s` into `into680`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into680`
	// (`into680` is now tainted).
	into680 := strings.ToLowerSpecial(nil, s)

	// Sink the tainted `into680`:
	sink(into680)
}

func TaintStepTest_StringsToTitle(sourceCQL interface{}) {
	// The flow is from `s` into `into878`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into878`
	// (`into878` is now tainted).
	into878 := strings.ToTitle(s)

	// Sink the tainted `into878`:
	sink(into878)
}

func TaintStepTest_StringsToTitleSpecial(sourceCQL interface{}) {
	// The flow is from `s` into `into761`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into761`
	// (`into761` is now tainted).
	into761 := strings.ToTitleSpecial(nil, s)

	// Sink the tainted `into761`:
	sink(into761)
}

func TaintStepTest_StringsToUpper(sourceCQL interface{}) {
	// The flow is from `s` into `into835`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into835`
	// (`into835` is now tainted).
	into835 := strings.ToUpper(s)

	// Sink the tainted `into835`:
	sink(into835)
}

func TaintStepTest_StringsToUpperSpecial(sourceCQL interface{}) {
	// The flow is from `s` into `into364`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into364`
	// (`into364` is now tainted).
	into364 := strings.ToUpperSpecial(nil, s)

	// Sink the tainted `into364`:
	sink(into364)
}

func TaintStepTest_StringsToValidUTF8(sourceCQL interface{}) {
	// The flow is from `s` into `into789`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into789`
	// (`into789` is now tainted).
	into789 := strings.ToValidUTF8(s, "")

	// Sink the tainted `into789`:
	sink(into789)
}

func TaintStepTest_StringsTrim(sourceCQL interface{}) {
	// The flow is from `s` into `into768`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into768`
	// (`into768` is now tainted).
	into768 := strings.Trim(s, "")

	// Sink the tainted `into768`:
	sink(into768)
}

func TaintStepTest_StringsTrimFunc(sourceCQL interface{}) {
	// The flow is from `s` into `into984`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into984`
	// (`into984` is now tainted).
	into984 := strings.TrimFunc(s, nil)

	// Sink the tainted `into984`:
	sink(into984)
}

func TaintStepTest_StringsTrimLeft(sourceCQL interface{}) {
	// The flow is from `s` into `into758`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into758`
	// (`into758` is now tainted).
	into758 := strings.TrimLeft(s, "")

	// Sink the tainted `into758`:
	sink(into758)
}

func TaintStepTest_StringsTrimLeftFunc(sourceCQL interface{}) {
	// The flow is from `s` into `into177`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into177`
	// (`into177` is now tainted).
	into177 := strings.TrimLeftFunc(s, nil)

	// Sink the tainted `into177`:
	sink(into177)
}

func TaintStepTest_StringsTrimPrefix(sourceCQL interface{}) {
	// The flow is from `s` into `into698`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into698`
	// (`into698` is now tainted).
	into698 := strings.TrimPrefix(s, "")

	// Sink the tainted `into698`:
	sink(into698)
}

func TaintStepTest_StringsTrimRight(sourceCQL interface{}) {
	// The flow is from `s` into `into268`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into268`
	// (`into268` is now tainted).
	into268 := strings.TrimRight(s, "")

	// Sink the tainted `into268`:
	sink(into268)
}

func TaintStepTest_StringsTrimRightFunc(sourceCQL interface{}) {
	// The flow is from `s` into `into386`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into386`
	// (`into386` is now tainted).
	into386 := strings.TrimRightFunc(s, nil)

	// Sink the tainted `into386`:
	sink(into386)
}

func TaintStepTest_StringsTrimSpace(sourceCQL interface{}) {
	// The flow is from `s` into `into764`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into764`
	// (`into764` is now tainted).
	into764 := strings.TrimSpace(s)

	// Sink the tainted `into764`:
	sink(into764)
}

func TaintStepTest_StringsTrimSuffix(sourceCQL interface{}) {
	// The flow is from `s` into `into526`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into526`
	// (`into526` is now tainted).
	into526 := strings.TrimSuffix(s, "")

	// Sink the tainted `into526`:
	sink(into526)
}

func TaintStepTest_StringsBuilderString(sourceCQL interface{}) {
	// The flow is from `from695` into `into343`.

	// Assume that `sourceCQL` has the underlying type of `from695`:
	from695 := sourceCQL.(strings.Builder)

	// Call the method that transfers the taint
	// from the receiver `from695` to the result `into343`
	// (`into343` is now tainted).
	into343 := from695.String()

	// Sink the tainted `into343`:
	sink(into343)
}

func TaintStepTest_StringsBuilderWrite(sourceCQL interface{}) {
	// The flow is from `p` into `into385`.

	// Assume that `sourceCQL` has the underlying type of `p`:
	p := sourceCQL.([]byte)

	// Declare `into385` variable:
	var into385 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `p` to the receiver `into385`
	// (`into385` is now tainted).
	into385.Write(p)

	// Sink the tainted `into385`:
	sink(into385)
}

func TaintStepTest_StringsBuilderWriteByte(sourceCQL interface{}) {
	// The flow is from `c` into `into789`.

	// Assume that `sourceCQL` has the underlying type of `c`:
	c := sourceCQL.(byte)

	// Declare `into789` variable:
	var into789 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `c` to the receiver `into789`
	// (`into789` is now tainted).
	into789.WriteByte(c)

	// Sink the tainted `into789`:
	sink(into789)
}

func TaintStepTest_StringsBuilderWriteRune(sourceCQL interface{}) {
	// The flow is from `r` into `into256`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(rune)

	// Declare `into256` variable:
	var into256 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `r` to the receiver `into256`
	// (`into256` is now tainted).
	into256.WriteRune(r)

	// Sink the tainted `into256`:
	sink(into256)
}

func TaintStepTest_StringsBuilderWriteString(sourceCQL interface{}) {
	// The flow is from `s` into `into112`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare `into112` variable:
	var into112 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `s` to the receiver `into112`
	// (`into112` is now tainted).
	into112.WriteString(s)

	// Sink the tainted `into112`:
	sink(into112)
}

func TaintStepTest_StringsReaderRead(sourceCQL interface{}) {
	// The flow is from `from715` into `b`.

	// Assume that `sourceCQL` has the underlying type of `from715`:
	from715 := sourceCQL.(strings.Reader)

	// Declare `b` variable:
	var b []byte

	// Call the method that transfers the taint
	// from the receiver `from715` to the argument `b`
	// (`b` is now tainted).
	from715.Read(b)

	// Sink the tainted `b`:
	sink(b)
}

func TaintStepTest_StringsReaderReadAt(sourceCQL interface{}) {
	// The flow is from `from968` into `b`.

	// Assume that `sourceCQL` has the underlying type of `from968`:
	from968 := sourceCQL.(strings.Reader)

	// Declare `b` variable:
	var b []byte

	// Call the method that transfers the taint
	// from the receiver `from968` to the argument `b`
	// (`b` is now tainted).
	from968.ReadAt(b, 0)

	// Sink the tainted `b`:
	sink(b)
}

func TaintStepTest_StringsReaderReadByte(sourceCQL interface{}) {
	// The flow is from `from208` into `into433`.

	// Assume that `sourceCQL` has the underlying type of `from208`:
	from208 := sourceCQL.(strings.Reader)

	// Call the method that transfers the taint
	// from the receiver `from208` to the result `into433`
	// (`into433` is now tainted).
	into433, _ := from208.ReadByte()

	// Sink the tainted `into433`:
	sink(into433)
}

func TaintStepTest_StringsReaderReadRune(sourceCQL interface{}) {
	// The flow is from `from617` into `ch`.

	// Assume that `sourceCQL` has the underlying type of `from617`:
	from617 := sourceCQL.(strings.Reader)

	// Call the method that transfers the taint
	// from the receiver `from617` to the result `ch`
	// (`ch` is now tainted).
	ch, _, _ := from617.ReadRune()

	// Sink the tainted `ch`:
	sink(ch)
}

func TaintStepTest_StringsReaderReset(sourceCQL interface{}) {
	// The flow is from `s` into `into272`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare `into272` variable:
	var into272 strings.Reader

	// Call the method that transfers the taint
	// from the parameter `s` to the receiver `into272`
	// (`into272` is now tainted).
	into272.Reset(s)

	// Sink the tainted `into272`:
	sink(into272)
}

func TaintStepTest_StringsReaderWriteTo(sourceCQL interface{}) {
	// The flow is from `from360` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from360`:
	from360 := sourceCQL.(strings.Reader)

	// Declare `w` variable:
	var w io.Writer

	// Call the method that transfers the taint
	// from the receiver `from360` to the argument `w`
	// (`w` is now tainted).
	from360.WriteTo(w)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_StringsReplacerReplace(sourceCQL interface{}) {
	// The flow is from `s` into `into185`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL strings.Replacer

	// Call the method that transfers the taint
	// from the parameter `s` to the result `into185`
	// (`into185` is now tainted).
	into185 := mediumObjCQL.Replace(s)

	// Sink the tainted `into185`:
	sink(into185)
}

func TaintStepTest_StringsReplacerWriteString(sourceCQL interface{}) {
	// The flow is from `s` into `w`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare `w` variable:
	var w io.Writer

	// Declare medium object/interface:
	var mediumObjCQL strings.Replacer

	// Call the method that transfers the taint
	// from the parameter `s` to the parameter `w`
	// (`w` is now tainted).
	mediumObjCQL.WriteString(w, s)

	// Sink the tainted `w`:
	sink(w)
}

func RunAllTaints_Strings(v interface{}) {
	{
		source := newSource()
		TaintStepTest_StringsFields(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsFieldsFunc(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsJoin(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsMap(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsNewReader(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsNewReplacer(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsRepeat(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReplace(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReplaceAll(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsSplit(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsSplitAfter(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsSplitAfterN(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsSplitN(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTitle(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsToLower(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsToLowerSpecial(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsToTitle(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsToTitleSpecial(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsToUpper(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsToUpperSpecial(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsToValidUTF8(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrim(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrimFunc(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrimLeft(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrimLeftFunc(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrimPrefix(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrimRight(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrimRightFunc(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrimSpace(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrimSuffix(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsBuilderString(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsBuilderWrite(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsBuilderWriteByte(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsBuilderWriteRune(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsBuilderWriteString(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReaderRead(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReaderReadAt(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReaderReadByte(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReaderReadRune(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReaderReset(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReaderWriteTo(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReplacerReplace(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReplacerWriteString(source)
	}
}
