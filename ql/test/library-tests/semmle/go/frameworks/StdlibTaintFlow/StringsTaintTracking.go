package main

import (
	"io"
	"strings"
)

func TaintStepTest_StringsFields_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString821` into `intoString819`.

	// Assume that `sourceCQL` has the underlying type of `fromString821`:
	fromString821 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString821` to result `intoString819`
	// (`intoString819` is now tainted).
	intoString819 := strings.Fields(fromString821)

	// Return the tainted `intoString819`:
	return intoString819
}

func TaintStepTest_StringsFieldsFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString220` into `intoString296`.

	// Assume that `sourceCQL` has the underlying type of `fromString220`:
	fromString220 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString220` to result `intoString296`
	// (`intoString296` is now tainted).
	intoString296 := strings.FieldsFunc(fromString220, nil)

	// Return the tainted `intoString296`:
	return intoString296
}

func TaintStepTest_StringsJoin_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString656` into `intoString732`.

	// Assume that `sourceCQL` has the underlying type of `fromString656`:
	fromString656 := sourceCQL.([]string)

	// Call the function that transfers the taint
	// from the parameter `fromString656` to result `intoString732`
	// (`intoString732` is now tainted).
	intoString732 := strings.Join(fromString656, "")

	// Return the tainted `intoString732`:
	return intoString732
}

func TaintStepTest_StringsJoin_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString747` into `intoString907`.

	// Assume that `sourceCQL` has the underlying type of `fromString747`:
	fromString747 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString747` to result `intoString907`
	// (`intoString907` is now tainted).
	intoString907 := strings.Join(nil, fromString747)

	// Return the tainted `intoString907`:
	return intoString907
}

func TaintStepTest_StringsMap_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString985` into `intoString612`.

	// Assume that `sourceCQL` has the underlying type of `fromString985`:
	fromString985 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString985` to result `intoString612`
	// (`intoString612` is now tainted).
	intoString612 := strings.Map(nil, fromString985)

	// Return the tainted `intoString612`:
	return intoString612
}

func TaintStepTest_StringsNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString926` into `intoReader875`.

	// Assume that `sourceCQL` has the underlying type of `fromString926`:
	fromString926 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString926` to result `intoReader875`
	// (`intoReader875` is now tainted).
	intoReader875 := strings.NewReader(fromString926)

	// Return the tainted `intoReader875`:
	return intoReader875
}

func TaintStepTest_StringsNewReplacer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString704` into `intoReplacer554`.

	// Assume that `sourceCQL` has the underlying type of `fromString704`:
	fromString704 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString704` to result `intoReplacer554`
	// (`intoReplacer554` is now tainted).
	intoReplacer554 := strings.NewReplacer(fromString704)

	// Return the tainted `intoReplacer554`:
	return intoReplacer554
}

func TaintStepTest_StringsRepeat_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString356` into `intoString935`.

	// Assume that `sourceCQL` has the underlying type of `fromString356`:
	fromString356 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString356` to result `intoString935`
	// (`intoString935` is now tainted).
	intoString935 := strings.Repeat(fromString356, 0)

	// Return the tainted `intoString935`:
	return intoString935
}

func TaintStepTest_StringsReplace_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString591` into `intoString684`.

	// Assume that `sourceCQL` has the underlying type of `fromString591`:
	fromString591 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString591` to result `intoString684`
	// (`intoString684` is now tainted).
	intoString684 := strings.Replace(fromString591, "", "", 0)

	// Return the tainted `intoString684`:
	return intoString684
}

func TaintStepTest_StringsReplace_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString399` into `intoString857`.

	// Assume that `sourceCQL` has the underlying type of `fromString399`:
	fromString399 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString399` to result `intoString857`
	// (`intoString857` is now tainted).
	intoString857 := strings.Replace("", "", fromString399, 0)

	// Return the tainted `intoString857`:
	return intoString857
}

func TaintStepTest_StringsReplaceAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString853` into `intoString819`.

	// Assume that `sourceCQL` has the underlying type of `fromString853`:
	fromString853 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString853` to result `intoString819`
	// (`intoString819` is now tainted).
	intoString819 := strings.ReplaceAll(fromString853, "", "")

	// Return the tainted `intoString819`:
	return intoString819
}

func TaintStepTest_StringsReplaceAll_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString415` into `intoString285`.

	// Assume that `sourceCQL` has the underlying type of `fromString415`:
	fromString415 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString415` to result `intoString285`
	// (`intoString285` is now tainted).
	intoString285 := strings.ReplaceAll("", "", fromString415)

	// Return the tainted `intoString285`:
	return intoString285
}

func TaintStepTest_StringsSplit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString208` into `intoString208`.

	// Assume that `sourceCQL` has the underlying type of `fromString208`:
	fromString208 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString208` to result `intoString208`
	// (`intoString208` is now tainted).
	intoString208 := strings.Split(fromString208, "")

	// Return the tainted `intoString208`:
	return intoString208
}

func TaintStepTest_StringsSplitAfter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString362` into `intoString714`.

	// Assume that `sourceCQL` has the underlying type of `fromString362`:
	fromString362 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString362` to result `intoString714`
	// (`intoString714` is now tainted).
	intoString714 := strings.SplitAfter(fromString362, "")

	// Return the tainted `intoString714`:
	return intoString714
}

func TaintStepTest_StringsSplitAfterN_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString836` into `intoString456`.

	// Assume that `sourceCQL` has the underlying type of `fromString836`:
	fromString836 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString836` to result `intoString456`
	// (`intoString456` is now tainted).
	intoString456 := strings.SplitAfterN(fromString836, "", 0)

	// Return the tainted `intoString456`:
	return intoString456
}

func TaintStepTest_StringsSplitN_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString328` into `intoString167`.

	// Assume that `sourceCQL` has the underlying type of `fromString328`:
	fromString328 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString328` to result `intoString167`
	// (`intoString167` is now tainted).
	intoString167 := strings.SplitN(fromString328, "", 0)

	// Return the tainted `intoString167`:
	return intoString167
}

func TaintStepTest_StringsTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString714` into `intoString335`.

	// Assume that `sourceCQL` has the underlying type of `fromString714`:
	fromString714 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString714` to result `intoString335`
	// (`intoString335` is now tainted).
	intoString335 := strings.Title(fromString714)

	// Return the tainted `intoString335`:
	return intoString335
}

func TaintStepTest_StringsToLower_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString972` into `intoString535`.

	// Assume that `sourceCQL` has the underlying type of `fromString972`:
	fromString972 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString972` to result `intoString535`
	// (`intoString535` is now tainted).
	intoString535 := strings.ToLower(fromString972)

	// Return the tainted `intoString535`:
	return intoString535
}

func TaintStepTest_StringsToLowerSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString851` into `intoString246`.

	// Assume that `sourceCQL` has the underlying type of `fromString851`:
	fromString851 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString851` to result `intoString246`
	// (`intoString246` is now tainted).
	intoString246 := strings.ToLowerSpecial(nil, fromString851)

	// Return the tainted `intoString246`:
	return intoString246
}

func TaintStepTest_StringsToTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString660` into `intoString317`.

	// Assume that `sourceCQL` has the underlying type of `fromString660`:
	fromString660 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString660` to result `intoString317`
	// (`intoString317` is now tainted).
	intoString317 := strings.ToTitle(fromString660)

	// Return the tainted `intoString317`:
	return intoString317
}

func TaintStepTest_StringsToTitleSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString234` into `intoString817`.

	// Assume that `sourceCQL` has the underlying type of `fromString234`:
	fromString234 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString234` to result `intoString817`
	// (`intoString817` is now tainted).
	intoString817 := strings.ToTitleSpecial(nil, fromString234)

	// Return the tainted `intoString817`:
	return intoString817
}

func TaintStepTest_StringsToUpper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString914` into `intoString490`.

	// Assume that `sourceCQL` has the underlying type of `fromString914`:
	fromString914 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString914` to result `intoString490`
	// (`intoString490` is now tainted).
	intoString490 := strings.ToUpper(fromString914)

	// Return the tainted `intoString490`:
	return intoString490
}

func TaintStepTest_StringsToUpperSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString555` into `intoString463`.

	// Assume that `sourceCQL` has the underlying type of `fromString555`:
	fromString555 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString555` to result `intoString463`
	// (`intoString463` is now tainted).
	intoString463 := strings.ToUpperSpecial(nil, fromString555)

	// Return the tainted `intoString463`:
	return intoString463
}

func TaintStepTest_StringsToValidUTF8_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString291` into `intoString780`.

	// Assume that `sourceCQL` has the underlying type of `fromString291`:
	fromString291 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString291` to result `intoString780`
	// (`intoString780` is now tainted).
	intoString780 := strings.ToValidUTF8(fromString291, "")

	// Return the tainted `intoString780`:
	return intoString780
}

func TaintStepTest_StringsToValidUTF8_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString536` into `intoString905`.

	// Assume that `sourceCQL` has the underlying type of `fromString536`:
	fromString536 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString536` to result `intoString905`
	// (`intoString905` is now tainted).
	intoString905 := strings.ToValidUTF8("", fromString536)

	// Return the tainted `intoString905`:
	return intoString905
}

func TaintStepTest_StringsTrim_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString310` into `intoString576`.

	// Assume that `sourceCQL` has the underlying type of `fromString310`:
	fromString310 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString310` to result `intoString576`
	// (`intoString576` is now tainted).
	intoString576 := strings.Trim(fromString310, "")

	// Return the tainted `intoString576`:
	return intoString576
}

func TaintStepTest_StringsTrimFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString962` into `intoString936`.

	// Assume that `sourceCQL` has the underlying type of `fromString962`:
	fromString962 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString962` to result `intoString936`
	// (`intoString936` is now tainted).
	intoString936 := strings.TrimFunc(fromString962, nil)

	// Return the tainted `intoString936`:
	return intoString936
}

func TaintStepTest_StringsTrimLeft_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString921` into `intoString555`.

	// Assume that `sourceCQL` has the underlying type of `fromString921`:
	fromString921 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString921` to result `intoString555`
	// (`intoString555` is now tainted).
	intoString555 := strings.TrimLeft(fromString921, "")

	// Return the tainted `intoString555`:
	return intoString555
}

func TaintStepTest_StringsTrimLeftFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString977` into `intoString584`.

	// Assume that `sourceCQL` has the underlying type of `fromString977`:
	fromString977 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString977` to result `intoString584`
	// (`intoString584` is now tainted).
	intoString584 := strings.TrimLeftFunc(fromString977, nil)

	// Return the tainted `intoString584`:
	return intoString584
}

func TaintStepTest_StringsTrimPrefix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString906` into `intoString324`.

	// Assume that `sourceCQL` has the underlying type of `fromString906`:
	fromString906 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString906` to result `intoString324`
	// (`intoString324` is now tainted).
	intoString324 := strings.TrimPrefix(fromString906, "")

	// Return the tainted `intoString324`:
	return intoString324
}

func TaintStepTest_StringsTrimRight_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString889` into `intoString815`.

	// Assume that `sourceCQL` has the underlying type of `fromString889`:
	fromString889 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString889` to result `intoString815`
	// (`intoString815` is now tainted).
	intoString815 := strings.TrimRight(fromString889, "")

	// Return the tainted `intoString815`:
	return intoString815
}

func TaintStepTest_StringsTrimRightFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString423` into `intoString858`.

	// Assume that `sourceCQL` has the underlying type of `fromString423`:
	fromString423 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString423` to result `intoString858`
	// (`intoString858` is now tainted).
	intoString858 := strings.TrimRightFunc(fromString423, nil)

	// Return the tainted `intoString858`:
	return intoString858
}

func TaintStepTest_StringsTrimSpace_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString609` into `intoString739`.

	// Assume that `sourceCQL` has the underlying type of `fromString609`:
	fromString609 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString609` to result `intoString739`
	// (`intoString739` is now tainted).
	intoString739 := strings.TrimSpace(fromString609)

	// Return the tainted `intoString739`:
	return intoString739
}

func TaintStepTest_StringsTrimSuffix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString150` into `intoString501`.

	// Assume that `sourceCQL` has the underlying type of `fromString150`:
	fromString150 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString150` to result `intoString501`
	// (`intoString501` is now tainted).
	intoString501 := strings.TrimSuffix(fromString150, "")

	// Return the tainted `intoString501`:
	return intoString501
}

func TaintStepTest_StringsBuilderString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuilder575` into `intoString318`.

	// Assume that `sourceCQL` has the underlying type of `fromBuilder575`:
	fromBuilder575 := sourceCQL.(strings.Builder)

	// Call the method that transfers the taint
	// from the receiver `fromBuilder575` to the result `intoString318`
	// (`intoString318` is now tainted).
	intoString318 := fromBuilder575.String()

	// Return the tainted `intoString318`:
	return intoString318
}

func TaintStepTest_StringsBuilderWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte983` into `intoBuilder365`.

	// Assume that `sourceCQL` has the underlying type of `fromByte983`:
	fromByte983 := sourceCQL.([]byte)

	// Declare `intoBuilder365` variable:
	var intoBuilder365 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `fromByte983` to the receiver `intoBuilder365`
	// (`intoBuilder365` is now tainted).
	intoBuilder365.Write(fromByte983)

	// Return the tainted `intoBuilder365`:
	return intoBuilder365
}

func TaintStepTest_StringsBuilderWriteByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte356` into `intoBuilder342`.

	// Assume that `sourceCQL` has the underlying type of `fromByte356`:
	fromByte356 := sourceCQL.(byte)

	// Declare `intoBuilder342` variable:
	var intoBuilder342 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `fromByte356` to the receiver `intoBuilder342`
	// (`intoBuilder342` is now tainted).
	intoBuilder342.WriteByte(fromByte356)

	// Return the tainted `intoBuilder342`:
	return intoBuilder342
}

func TaintStepTest_StringsBuilderWriteRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune347` into `intoBuilder115`.

	// Assume that `sourceCQL` has the underlying type of `fromRune347`:
	fromRune347 := sourceCQL.(rune)

	// Declare `intoBuilder115` variable:
	var intoBuilder115 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `fromRune347` to the receiver `intoBuilder115`
	// (`intoBuilder115` is now tainted).
	intoBuilder115.WriteRune(fromRune347)

	// Return the tainted `intoBuilder115`:
	return intoBuilder115
}

func TaintStepTest_StringsBuilderWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString332` into `intoBuilder601`.

	// Assume that `sourceCQL` has the underlying type of `fromString332`:
	fromString332 := sourceCQL.(string)

	// Declare `intoBuilder601` variable:
	var intoBuilder601 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `fromString332` to the receiver `intoBuilder601`
	// (`intoBuilder601` is now tainted).
	intoBuilder601.WriteString(fromString332)

	// Return the tainted `intoBuilder601`:
	return intoBuilder601
}

func TaintStepTest_StringsReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader350` into `intoByte755`.

	// Assume that `sourceCQL` has the underlying type of `fromReader350`:
	fromReader350 := sourceCQL.(strings.Reader)

	// Declare `intoByte755` variable:
	var intoByte755 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader350` to the argument `intoByte755`
	// (`intoByte755` is now tainted).
	fromReader350.Read(intoByte755)

	// Return the tainted `intoByte755`:
	return intoByte755
}

func TaintStepTest_StringsReaderReadAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader865` into `intoByte898`.

	// Assume that `sourceCQL` has the underlying type of `fromReader865`:
	fromReader865 := sourceCQL.(strings.Reader)

	// Declare `intoByte898` variable:
	var intoByte898 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader865` to the argument `intoByte898`
	// (`intoByte898` is now tainted).
	fromReader865.ReadAt(intoByte898, 0)

	// Return the tainted `intoByte898`:
	return intoByte898
}

func TaintStepTest_StringsReaderReadByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader980` into `intoByte450`.

	// Assume that `sourceCQL` has the underlying type of `fromReader980`:
	fromReader980 := sourceCQL.(strings.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader980` to the result `intoByte450`
	// (`intoByte450` is now tainted).
	intoByte450, _ := fromReader980.ReadByte()

	// Return the tainted `intoByte450`:
	return intoByte450
}

func TaintStepTest_StringsReaderReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader620` into `intoRune525`.

	// Assume that `sourceCQL` has the underlying type of `fromReader620`:
	fromReader620 := sourceCQL.(strings.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader620` to the result `intoRune525`
	// (`intoRune525` is now tainted).
	intoRune525, _, _ := fromReader620.ReadRune()

	// Return the tainted `intoRune525`:
	return intoRune525
}

func TaintStepTest_StringsReaderReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString721` into `intoReader994`.

	// Assume that `sourceCQL` has the underlying type of `fromString721`:
	fromString721 := sourceCQL.(string)

	// Declare `intoReader994` variable:
	var intoReader994 strings.Reader

	// Call the method that transfers the taint
	// from the parameter `fromString721` to the receiver `intoReader994`
	// (`intoReader994` is now tainted).
	intoReader994.Reset(fromString721)

	// Return the tainted `intoReader994`:
	return intoReader994
}

func TaintStepTest_StringsReaderWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader208` into `intoWriter965`.

	// Assume that `sourceCQL` has the underlying type of `fromReader208`:
	fromReader208 := sourceCQL.(strings.Reader)

	// Declare `intoWriter965` variable:
	var intoWriter965 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromReader208` to the argument `intoWriter965`
	// (`intoWriter965` is now tainted).
	fromReader208.WriteTo(intoWriter965)

	// Return the tainted `intoWriter965`:
	return intoWriter965
}

func TaintStepTest_StringsReplacerReplace_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString342` into `intoString517`.

	// Assume that `sourceCQL` has the underlying type of `fromString342`:
	fromString342 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL strings.Replacer

	// Call the method that transfers the taint
	// from the parameter `fromString342` to the result `intoString517`
	// (`intoString517` is now tainted).
	intoString517 := mediumObjCQL.Replace(fromString342)

	// Return the tainted `intoString517`:
	return intoString517
}

func TaintStepTest_StringsReplacerWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString612` into `intoWriter295`.

	// Assume that `sourceCQL` has the underlying type of `fromString612`:
	fromString612 := sourceCQL.(string)

	// Declare `intoWriter295` variable:
	var intoWriter295 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL strings.Replacer

	// Call the method that transfers the taint
	// from the parameter `fromString612` to the parameter `intoWriter295`
	// (`intoWriter295` is now tainted).
	mediumObjCQL.WriteString(intoWriter295, fromString612)

	// Return the tainted `intoWriter295`:
	return intoWriter295
}

func RunAllTaints_Strings() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsFields_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsFieldsFunc_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsJoin_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsJoin_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsMap_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsNewReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsNewReplacer_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsRepeat_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsReplace_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsReplace_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsReplaceAll_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsReplaceAll_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsSplit_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsSplitAfter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsSplitAfterN_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsSplitN_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsTitle_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsToLower_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsToLowerSpecial_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsToTitle_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsToTitleSpecial_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsToUpper_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsToUpperSpecial_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsToValidUTF8_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsToValidUTF8_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsTrim_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsTrimFunc_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsTrimLeft_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsTrimLeftFunc_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsTrimPrefix_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsTrimRight_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsTrimRightFunc_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsTrimSpace_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsTrimSuffix_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsBuilderString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsBuilderWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsBuilderWriteByte_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsBuilderWriteRune_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsBuilderWriteString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsReaderRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsReaderReadAt_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsReaderReadByte_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsReaderReadRune_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsReaderReset_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsReaderWriteTo_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsReplacerReplace_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_StringsReplacerWriteString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
