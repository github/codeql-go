package main

import (
	"io"
	"strings"
)

func TaintStepTest_StringsFields_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString293` into `intoString127`.

	// Assume that `sourceCQL` has the underlying type of `fromString293`:
	fromString293 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString293` to result `intoString127`
	// (`intoString127` is now tainted).
	intoString127 := strings.Fields(fromString293)

	// Sink the tainted `intoString127`:
	sink(intoString127)
}

func TaintStepTest_StringsFieldsFunc_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString566` into `intoString649`.

	// Assume that `sourceCQL` has the underlying type of `fromString566`:
	fromString566 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString566` to result `intoString649`
	// (`intoString649` is now tainted).
	intoString649 := strings.FieldsFunc(fromString566, nil)

	// Sink the tainted `intoString649`:
	sink(intoString649)
}

func TaintStepTest_StringsJoin_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString946` into `intoString349`.

	// Assume that `sourceCQL` has the underlying type of `fromString946`:
	fromString946 := sourceCQL.([]string)

	// Call the function that transfers the taint
	// from the parameter `fromString946` to result `intoString349`
	// (`intoString349` is now tainted).
	intoString349 := strings.Join(fromString946, "")

	// Sink the tainted `intoString349`:
	sink(intoString349)
}

func TaintStepTest_StringsJoin_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString657` into `intoString508`.

	// Assume that `sourceCQL` has the underlying type of `fromString657`:
	fromString657 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString657` to result `intoString508`
	// (`intoString508` is now tainted).
	intoString508 := strings.Join(nil, fromString657)

	// Sink the tainted `intoString508`:
	sink(intoString508)
}

func TaintStepTest_StringsMap_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString982` into `intoString179`.

	// Assume that `sourceCQL` has the underlying type of `fromString982`:
	fromString982 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString982` to result `intoString179`
	// (`intoString179` is now tainted).
	intoString179 := strings.Map(nil, fromString982)

	// Sink the tainted `intoString179`:
	sink(intoString179)
}

func TaintStepTest_StringsNewReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString186` into `intoReader423`.

	// Assume that `sourceCQL` has the underlying type of `fromString186`:
	fromString186 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString186` to result `intoReader423`
	// (`intoReader423` is now tainted).
	intoReader423 := strings.NewReader(fromString186)

	// Sink the tainted `intoReader423`:
	sink(intoReader423)
}

func TaintStepTest_StringsNewReplacer_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString662` into `intoReplacer910`.

	// Assume that `sourceCQL` has the underlying type of `fromString662`:
	fromString662 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString662` to result `intoReplacer910`
	// (`intoReplacer910` is now tainted).
	intoReplacer910 := strings.NewReplacer(fromString662)

	// Sink the tainted `intoReplacer910`:
	sink(intoReplacer910)
}

func TaintStepTest_StringsRepeat_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString659` into `intoString301`.

	// Assume that `sourceCQL` has the underlying type of `fromString659`:
	fromString659 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString659` to result `intoString301`
	// (`intoString301` is now tainted).
	intoString301 := strings.Repeat(fromString659, 0)

	// Sink the tainted `intoString301`:
	sink(intoString301)
}

func TaintStepTest_StringsReplace_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString562` into `intoString979`.

	// Assume that `sourceCQL` has the underlying type of `fromString562`:
	fromString562 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString562` to result `intoString979`
	// (`intoString979` is now tainted).
	intoString979 := strings.Replace(fromString562, "", "", 0)

	// Sink the tainted `intoString979`:
	sink(intoString979)
}

func TaintStepTest_StringsReplace_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString813` into `intoString815`.

	// Assume that `sourceCQL` has the underlying type of `fromString813`:
	fromString813 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString813` to result `intoString815`
	// (`intoString815` is now tainted).
	intoString815 := strings.Replace("", "", fromString813, 0)

	// Sink the tainted `intoString815`:
	sink(intoString815)
}

func TaintStepTest_StringsReplaceAll_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString245` into `intoString726`.

	// Assume that `sourceCQL` has the underlying type of `fromString245`:
	fromString245 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString245` to result `intoString726`
	// (`intoString726` is now tainted).
	intoString726 := strings.ReplaceAll(fromString245, "", "")

	// Sink the tainted `intoString726`:
	sink(intoString726)
}

func TaintStepTest_StringsReplaceAll_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString875` into `intoString216`.

	// Assume that `sourceCQL` has the underlying type of `fromString875`:
	fromString875 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString875` to result `intoString216`
	// (`intoString216` is now tainted).
	intoString216 := strings.ReplaceAll("", "", fromString875)

	// Sink the tainted `intoString216`:
	sink(intoString216)
}

func TaintStepTest_StringsSplit_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString120` into `intoString564`.

	// Assume that `sourceCQL` has the underlying type of `fromString120`:
	fromString120 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString120` to result `intoString564`
	// (`intoString564` is now tainted).
	intoString564 := strings.Split(fromString120, "")

	// Sink the tainted `intoString564`:
	sink(intoString564)
}

func TaintStepTest_StringsSplitAfter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString391` into `intoString140`.

	// Assume that `sourceCQL` has the underlying type of `fromString391`:
	fromString391 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString391` to result `intoString140`
	// (`intoString140` is now tainted).
	intoString140 := strings.SplitAfter(fromString391, "")

	// Sink the tainted `intoString140`:
	sink(intoString140)
}

func TaintStepTest_StringsSplitAfterN_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString562` into `intoString840`.

	// Assume that `sourceCQL` has the underlying type of `fromString562`:
	fromString562 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString562` to result `intoString840`
	// (`intoString840` is now tainted).
	intoString840 := strings.SplitAfterN(fromString562, "", 0)

	// Sink the tainted `intoString840`:
	sink(intoString840)
}

func TaintStepTest_StringsSplitN_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString456` into `intoString428`.

	// Assume that `sourceCQL` has the underlying type of `fromString456`:
	fromString456 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString456` to result `intoString428`
	// (`intoString428` is now tainted).
	intoString428 := strings.SplitN(fromString456, "", 0)

	// Sink the tainted `intoString428`:
	sink(intoString428)
}

func TaintStepTest_StringsTitle_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString843` into `intoString496`.

	// Assume that `sourceCQL` has the underlying type of `fromString843`:
	fromString843 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString843` to result `intoString496`
	// (`intoString496` is now tainted).
	intoString496 := strings.Title(fromString843)

	// Sink the tainted `intoString496`:
	sink(intoString496)
}

func TaintStepTest_StringsToLower_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString261` into `intoString593`.

	// Assume that `sourceCQL` has the underlying type of `fromString261`:
	fromString261 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString261` to result `intoString593`
	// (`intoString593` is now tainted).
	intoString593 := strings.ToLower(fromString261)

	// Sink the tainted `intoString593`:
	sink(intoString593)
}

func TaintStepTest_StringsToLowerSpecial_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString433` into `intoString759`.

	// Assume that `sourceCQL` has the underlying type of `fromString433`:
	fromString433 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString433` to result `intoString759`
	// (`intoString759` is now tainted).
	intoString759 := strings.ToLowerSpecial(nil, fromString433)

	// Sink the tainted `intoString759`:
	sink(intoString759)
}

func TaintStepTest_StringsToTitle_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString394` into `intoString281`.

	// Assume that `sourceCQL` has the underlying type of `fromString394`:
	fromString394 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString394` to result `intoString281`
	// (`intoString281` is now tainted).
	intoString281 := strings.ToTitle(fromString394)

	// Sink the tainted `intoString281`:
	sink(intoString281)
}

func TaintStepTest_StringsToTitleSpecial_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString207` into `intoString142`.

	// Assume that `sourceCQL` has the underlying type of `fromString207`:
	fromString207 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString207` to result `intoString142`
	// (`intoString142` is now tainted).
	intoString142 := strings.ToTitleSpecial(nil, fromString207)

	// Sink the tainted `intoString142`:
	sink(intoString142)
}

func TaintStepTest_StringsToUpper_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString632` into `intoString702`.

	// Assume that `sourceCQL` has the underlying type of `fromString632`:
	fromString632 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString632` to result `intoString702`
	// (`intoString702` is now tainted).
	intoString702 := strings.ToUpper(fromString632)

	// Sink the tainted `intoString702`:
	sink(intoString702)
}

func TaintStepTest_StringsToUpperSpecial_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString209` into `intoString722`.

	// Assume that `sourceCQL` has the underlying type of `fromString209`:
	fromString209 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString209` to result `intoString722`
	// (`intoString722` is now tainted).
	intoString722 := strings.ToUpperSpecial(nil, fromString209)

	// Sink the tainted `intoString722`:
	sink(intoString722)
}

func TaintStepTest_StringsToValidUTF8_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString309` into `intoString910`.

	// Assume that `sourceCQL` has the underlying type of `fromString309`:
	fromString309 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString309` to result `intoString910`
	// (`intoString910` is now tainted).
	intoString910 := strings.ToValidUTF8(fromString309, "")

	// Sink the tainted `intoString910`:
	sink(intoString910)
}

func TaintStepTest_StringsToValidUTF8_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString660` into `intoString888`.

	// Assume that `sourceCQL` has the underlying type of `fromString660`:
	fromString660 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString660` to result `intoString888`
	// (`intoString888` is now tainted).
	intoString888 := strings.ToValidUTF8("", fromString660)

	// Sink the tainted `intoString888`:
	sink(intoString888)
}

func TaintStepTest_StringsTrim_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString693` into `intoString901`.

	// Assume that `sourceCQL` has the underlying type of `fromString693`:
	fromString693 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString693` to result `intoString901`
	// (`intoString901` is now tainted).
	intoString901 := strings.Trim(fromString693, "")

	// Sink the tainted `intoString901`:
	sink(intoString901)
}

func TaintStepTest_StringsTrimFunc_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString115` into `intoString598`.

	// Assume that `sourceCQL` has the underlying type of `fromString115`:
	fromString115 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString115` to result `intoString598`
	// (`intoString598` is now tainted).
	intoString598 := strings.TrimFunc(fromString115, nil)

	// Sink the tainted `intoString598`:
	sink(intoString598)
}

func TaintStepTest_StringsTrimLeft_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString732` into `intoString944`.

	// Assume that `sourceCQL` has the underlying type of `fromString732`:
	fromString732 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString732` to result `intoString944`
	// (`intoString944` is now tainted).
	intoString944 := strings.TrimLeft(fromString732, "")

	// Sink the tainted `intoString944`:
	sink(intoString944)
}

func TaintStepTest_StringsTrimLeftFunc_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString208` into `intoString543`.

	// Assume that `sourceCQL` has the underlying type of `fromString208`:
	fromString208 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString208` to result `intoString543`
	// (`intoString543` is now tainted).
	intoString543 := strings.TrimLeftFunc(fromString208, nil)

	// Sink the tainted `intoString543`:
	sink(intoString543)
}

func TaintStepTest_StringsTrimPrefix_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString452` into `intoString433`.

	// Assume that `sourceCQL` has the underlying type of `fromString452`:
	fromString452 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString452` to result `intoString433`
	// (`intoString433` is now tainted).
	intoString433 := strings.TrimPrefix(fromString452, "")

	// Sink the tainted `intoString433`:
	sink(intoString433)
}

func TaintStepTest_StringsTrimRight_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString759` into `intoString889`.

	// Assume that `sourceCQL` has the underlying type of `fromString759`:
	fromString759 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString759` to result `intoString889`
	// (`intoString889` is now tainted).
	intoString889 := strings.TrimRight(fromString759, "")

	// Sink the tainted `intoString889`:
	sink(intoString889)
}

func TaintStepTest_StringsTrimRightFunc_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString821` into `intoString292`.

	// Assume that `sourceCQL` has the underlying type of `fromString821`:
	fromString821 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString821` to result `intoString292`
	// (`intoString292` is now tainted).
	intoString292 := strings.TrimRightFunc(fromString821, nil)

	// Sink the tainted `intoString292`:
	sink(intoString292)
}

func TaintStepTest_StringsTrimSpace_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString321` into `intoString447`.

	// Assume that `sourceCQL` has the underlying type of `fromString321`:
	fromString321 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString321` to result `intoString447`
	// (`intoString447` is now tainted).
	intoString447 := strings.TrimSpace(fromString321)

	// Sink the tainted `intoString447`:
	sink(intoString447)
}

func TaintStepTest_StringsTrimSuffix_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString171` into `intoString455`.

	// Assume that `sourceCQL` has the underlying type of `fromString171`:
	fromString171 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString171` to result `intoString455`
	// (`intoString455` is now tainted).
	intoString455 := strings.TrimSuffix(fromString171, "")

	// Sink the tainted `intoString455`:
	sink(intoString455)
}

func TaintStepTest_StringsBuilderString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBuilder835` into `intoString351`.

	// Assume that `sourceCQL` has the underlying type of `fromBuilder835`:
	fromBuilder835 := sourceCQL.(strings.Builder)

	// Call the method that transfers the taint
	// from the receiver `fromBuilder835` to the result `intoString351`
	// (`intoString351` is now tainted).
	intoString351 := fromBuilder835.String()

	// Sink the tainted `intoString351`:
	sink(intoString351)
}

func TaintStepTest_StringsBuilderWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte949` into `intoBuilder958`.

	// Assume that `sourceCQL` has the underlying type of `fromByte949`:
	fromByte949 := sourceCQL.([]byte)

	// Declare `intoBuilder958` variable:
	var intoBuilder958 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `fromByte949` to the receiver `intoBuilder958`
	// (`intoBuilder958` is now tainted).
	intoBuilder958.Write(fromByte949)

	// Sink the tainted `intoBuilder958`:
	sink(intoBuilder958)
}

func TaintStepTest_StringsBuilderWriteByte_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte309` into `intoBuilder876`.

	// Assume that `sourceCQL` has the underlying type of `fromByte309`:
	fromByte309 := sourceCQL.(byte)

	// Declare `intoBuilder876` variable:
	var intoBuilder876 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `fromByte309` to the receiver `intoBuilder876`
	// (`intoBuilder876` is now tainted).
	intoBuilder876.WriteByte(fromByte309)

	// Sink the tainted `intoBuilder876`:
	sink(intoBuilder876)
}

func TaintStepTest_StringsBuilderWriteRune_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune563` into `intoBuilder663`.

	// Assume that `sourceCQL` has the underlying type of `fromRune563`:
	fromRune563 := sourceCQL.(rune)

	// Declare `intoBuilder663` variable:
	var intoBuilder663 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `fromRune563` to the receiver `intoBuilder663`
	// (`intoBuilder663` is now tainted).
	intoBuilder663.WriteRune(fromRune563)

	// Sink the tainted `intoBuilder663`:
	sink(intoBuilder663)
}

func TaintStepTest_StringsBuilderWriteString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString997` into `intoBuilder145`.

	// Assume that `sourceCQL` has the underlying type of `fromString997`:
	fromString997 := sourceCQL.(string)

	// Declare `intoBuilder145` variable:
	var intoBuilder145 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `fromString997` to the receiver `intoBuilder145`
	// (`intoBuilder145` is now tainted).
	intoBuilder145.WriteString(fromString997)

	// Sink the tainted `intoBuilder145`:
	sink(intoBuilder145)
}

func TaintStepTest_StringsReaderRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader136` into `intoByte606`.

	// Assume that `sourceCQL` has the underlying type of `fromReader136`:
	fromReader136 := sourceCQL.(strings.Reader)

	// Declare `intoByte606` variable:
	var intoByte606 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader136` to the argument `intoByte606`
	// (`intoByte606` is now tainted).
	fromReader136.Read(intoByte606)

	// Sink the tainted `intoByte606`:
	sink(intoByte606)
}

func TaintStepTest_StringsReaderReadAt_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader493` into `intoByte612`.

	// Assume that `sourceCQL` has the underlying type of `fromReader493`:
	fromReader493 := sourceCQL.(strings.Reader)

	// Declare `intoByte612` variable:
	var intoByte612 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader493` to the argument `intoByte612`
	// (`intoByte612` is now tainted).
	fromReader493.ReadAt(intoByte612, 0)

	// Sink the tainted `intoByte612`:
	sink(intoByte612)
}

func TaintStepTest_StringsReaderReadByte_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader111` into `intoByte462`.

	// Assume that `sourceCQL` has the underlying type of `fromReader111`:
	fromReader111 := sourceCQL.(strings.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader111` to the result `intoByte462`
	// (`intoByte462` is now tainted).
	intoByte462, _ := fromReader111.ReadByte()

	// Sink the tainted `intoByte462`:
	sink(intoByte462)
}

func TaintStepTest_StringsReaderReadRune_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader660` into `intoRune231`.

	// Assume that `sourceCQL` has the underlying type of `fromReader660`:
	fromReader660 := sourceCQL.(strings.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader660` to the result `intoRune231`
	// (`intoRune231` is now tainted).
	intoRune231, _, _ := fromReader660.ReadRune()

	// Sink the tainted `intoRune231`:
	sink(intoRune231)
}

func TaintStepTest_StringsReaderReset_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString204` into `intoReader796`.

	// Assume that `sourceCQL` has the underlying type of `fromString204`:
	fromString204 := sourceCQL.(string)

	// Declare `intoReader796` variable:
	var intoReader796 strings.Reader

	// Call the method that transfers the taint
	// from the parameter `fromString204` to the receiver `intoReader796`
	// (`intoReader796` is now tainted).
	intoReader796.Reset(fromString204)

	// Sink the tainted `intoReader796`:
	sink(intoReader796)
}

func TaintStepTest_StringsReaderWriteTo_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader783` into `intoWriter145`.

	// Assume that `sourceCQL` has the underlying type of `fromReader783`:
	fromReader783 := sourceCQL.(strings.Reader)

	// Declare `intoWriter145` variable:
	var intoWriter145 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromReader783` to the argument `intoWriter145`
	// (`intoWriter145` is now tainted).
	fromReader783.WriteTo(intoWriter145)

	// Sink the tainted `intoWriter145`:
	sink(intoWriter145)
}

func TaintStepTest_StringsReplacerReplace_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString533` into `intoString121`.

	// Assume that `sourceCQL` has the underlying type of `fromString533`:
	fromString533 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL strings.Replacer

	// Call the method that transfers the taint
	// from the parameter `fromString533` to the result `intoString121`
	// (`intoString121` is now tainted).
	intoString121 := mediumObjCQL.Replace(fromString533)

	// Sink the tainted `intoString121`:
	sink(intoString121)
}

func TaintStepTest_StringsReplacerWriteString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString663` into `intoWriter145`.

	// Assume that `sourceCQL` has the underlying type of `fromString663`:
	fromString663 := sourceCQL.(string)

	// Declare `intoWriter145` variable:
	var intoWriter145 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL strings.Replacer

	// Call the method that transfers the taint
	// from the parameter `fromString663` to the parameter `intoWriter145`
	// (`intoWriter145` is now tainted).
	mediumObjCQL.WriteString(intoWriter145, fromString663)

	// Sink the tainted `intoWriter145`:
	sink(intoWriter145)
}

func RunAllTaints_Strings() {
	{
		source := newSource()
		TaintStepTest_StringsFields_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsFieldsFunc_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsJoin_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsJoin_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsMap_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsNewReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsNewReplacer_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsRepeat_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReplace_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReplace_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReplaceAll_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReplaceAll_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsSplit_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsSplitAfter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsSplitAfterN_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsSplitN_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTitle_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsToLower_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsToLowerSpecial_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsToTitle_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsToTitleSpecial_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsToUpper_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsToUpperSpecial_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsToValidUTF8_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsToValidUTF8_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrim_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrimFunc_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrimLeft_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrimLeftFunc_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrimPrefix_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrimRight_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrimRightFunc_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrimSpace_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsTrimSuffix_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsBuilderString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsBuilderWrite_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsBuilderWriteByte_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsBuilderWriteRune_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsBuilderWriteString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReaderRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReaderReadAt_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReaderReadByte_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReaderReadRune_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReaderReset_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReaderWriteTo_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReplacerReplace_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_StringsReplacerWriteString_B0I0O0(source)
	}
}
