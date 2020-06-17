// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"strings"
)

func TaintStepTest_StringsFields_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString656` into `intoString414`.

	// Assume that `sourceCQL` has the underlying type of `fromString656`:
	fromString656 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString656` to result `intoString414`
	// (`intoString414` is now tainted).
	intoString414 := strings.Fields(fromString656)

	// Return the tainted `intoString414`:
	return intoString414
}

func TaintStepTest_StringsFieldsFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString518` into `intoString650`.

	// Assume that `sourceCQL` has the underlying type of `fromString518`:
	fromString518 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString518` to result `intoString650`
	// (`intoString650` is now tainted).
	intoString650 := strings.FieldsFunc(fromString518, nil)

	// Return the tainted `intoString650`:
	return intoString650
}

func TaintStepTest_StringsJoin_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString784` into `intoString957`.

	// Assume that `sourceCQL` has the underlying type of `fromString784`:
	fromString784 := sourceCQL.([]string)

	// Call the function that transfers the taint
	// from the parameter `fromString784` to result `intoString957`
	// (`intoString957` is now tainted).
	intoString957 := strings.Join(fromString784, "")

	// Return the tainted `intoString957`:
	return intoString957
}

func TaintStepTest_StringsJoin_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString520` into `intoString443`.

	// Assume that `sourceCQL` has the underlying type of `fromString520`:
	fromString520 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString520` to result `intoString443`
	// (`intoString443` is now tainted).
	intoString443 := strings.Join(nil, fromString520)

	// Return the tainted `intoString443`:
	return intoString443
}

func TaintStepTest_StringsMap_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString127` into `intoString483`.

	// Assume that `sourceCQL` has the underlying type of `fromString127`:
	fromString127 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString127` to result `intoString483`
	// (`intoString483` is now tainted).
	intoString483 := strings.Map(nil, fromString127)

	// Return the tainted `intoString483`:
	return intoString483
}

func TaintStepTest_StringsNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString989` into `intoReader982`.

	// Assume that `sourceCQL` has the underlying type of `fromString989`:
	fromString989 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString989` to result `intoReader982`
	// (`intoReader982` is now tainted).
	intoReader982 := strings.NewReader(fromString989)

	// Return the tainted `intoReader982`:
	return intoReader982
}

func TaintStepTest_StringsNewReplacer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString417` into `intoReplacer584`.

	// Assume that `sourceCQL` has the underlying type of `fromString417`:
	fromString417 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString417` to result `intoReplacer584`
	// (`intoReplacer584` is now tainted).
	intoReplacer584 := strings.NewReplacer(fromString417)

	// Return the tainted `intoReplacer584`:
	return intoReplacer584
}

func TaintStepTest_StringsRepeat_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString991` into `intoString881`.

	// Assume that `sourceCQL` has the underlying type of `fromString991`:
	fromString991 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString991` to result `intoString881`
	// (`intoString881` is now tainted).
	intoString881 := strings.Repeat(fromString991, 0)

	// Return the tainted `intoString881`:
	return intoString881
}

func TaintStepTest_StringsReplace_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString186` into `intoString284`.

	// Assume that `sourceCQL` has the underlying type of `fromString186`:
	fromString186 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString186` to result `intoString284`
	// (`intoString284` is now tainted).
	intoString284 := strings.Replace(fromString186, "", "", 0)

	// Return the tainted `intoString284`:
	return intoString284
}

func TaintStepTest_StringsReplace_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString908` into `intoString137`.

	// Assume that `sourceCQL` has the underlying type of `fromString908`:
	fromString908 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString908` to result `intoString137`
	// (`intoString137` is now tainted).
	intoString137 := strings.Replace("", "", fromString908, 0)

	// Return the tainted `intoString137`:
	return intoString137
}

func TaintStepTest_StringsReplaceAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString494` into `intoString873`.

	// Assume that `sourceCQL` has the underlying type of `fromString494`:
	fromString494 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString494` to result `intoString873`
	// (`intoString873` is now tainted).
	intoString873 := strings.ReplaceAll(fromString494, "", "")

	// Return the tainted `intoString873`:
	return intoString873
}

func TaintStepTest_StringsReplaceAll_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString599` into `intoString409`.

	// Assume that `sourceCQL` has the underlying type of `fromString599`:
	fromString599 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString599` to result `intoString409`
	// (`intoString409` is now tainted).
	intoString409 := strings.ReplaceAll("", "", fromString599)

	// Return the tainted `intoString409`:
	return intoString409
}

func TaintStepTest_StringsSplit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString246` into `intoString898`.

	// Assume that `sourceCQL` has the underlying type of `fromString246`:
	fromString246 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString246` to result `intoString898`
	// (`intoString898` is now tainted).
	intoString898 := strings.Split(fromString246, "")

	// Return the tainted `intoString898`:
	return intoString898
}

func TaintStepTest_StringsSplitAfter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString598` into `intoString631`.

	// Assume that `sourceCQL` has the underlying type of `fromString598`:
	fromString598 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString598` to result `intoString631`
	// (`intoString631` is now tainted).
	intoString631 := strings.SplitAfter(fromString598, "")

	// Return the tainted `intoString631`:
	return intoString631
}

func TaintStepTest_StringsSplitAfterN_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString165` into `intoString150`.

	// Assume that `sourceCQL` has the underlying type of `fromString165`:
	fromString165 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString165` to result `intoString150`
	// (`intoString150` is now tainted).
	intoString150 := strings.SplitAfterN(fromString165, "", 0)

	// Return the tainted `intoString150`:
	return intoString150
}

func TaintStepTest_StringsSplitN_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString340` into `intoString471`.

	// Assume that `sourceCQL` has the underlying type of `fromString340`:
	fromString340 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString340` to result `intoString471`
	// (`intoString471` is now tainted).
	intoString471 := strings.SplitN(fromString340, "", 0)

	// Return the tainted `intoString471`:
	return intoString471
}

func TaintStepTest_StringsTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString290` into `intoString758`.

	// Assume that `sourceCQL` has the underlying type of `fromString290`:
	fromString290 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString290` to result `intoString758`
	// (`intoString758` is now tainted).
	intoString758 := strings.Title(fromString290)

	// Return the tainted `intoString758`:
	return intoString758
}

func TaintStepTest_StringsToLower_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString396` into `intoString707`.

	// Assume that `sourceCQL` has the underlying type of `fromString396`:
	fromString396 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString396` to result `intoString707`
	// (`intoString707` is now tainted).
	intoString707 := strings.ToLower(fromString396)

	// Return the tainted `intoString707`:
	return intoString707
}

func TaintStepTest_StringsToLowerSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString912` into `intoString718`.

	// Assume that `sourceCQL` has the underlying type of `fromString912`:
	fromString912 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString912` to result `intoString718`
	// (`intoString718` is now tainted).
	intoString718 := strings.ToLowerSpecial(nil, fromString912)

	// Return the tainted `intoString718`:
	return intoString718
}

func TaintStepTest_StringsToTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString972` into `intoString633`.

	// Assume that `sourceCQL` has the underlying type of `fromString972`:
	fromString972 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString972` to result `intoString633`
	// (`intoString633` is now tainted).
	intoString633 := strings.ToTitle(fromString972)

	// Return the tainted `intoString633`:
	return intoString633
}

func TaintStepTest_StringsToTitleSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString316` into `intoString145`.

	// Assume that `sourceCQL` has the underlying type of `fromString316`:
	fromString316 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString316` to result `intoString145`
	// (`intoString145` is now tainted).
	intoString145 := strings.ToTitleSpecial(nil, fromString316)

	// Return the tainted `intoString145`:
	return intoString145
}

func TaintStepTest_StringsToUpper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString817` into `intoString474`.

	// Assume that `sourceCQL` has the underlying type of `fromString817`:
	fromString817 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString817` to result `intoString474`
	// (`intoString474` is now tainted).
	intoString474 := strings.ToUpper(fromString817)

	// Return the tainted `intoString474`:
	return intoString474
}

func TaintStepTest_StringsToUpperSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString832` into `intoString378`.

	// Assume that `sourceCQL` has the underlying type of `fromString832`:
	fromString832 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString832` to result `intoString378`
	// (`intoString378` is now tainted).
	intoString378 := strings.ToUpperSpecial(nil, fromString832)

	// Return the tainted `intoString378`:
	return intoString378
}

func TaintStepTest_StringsToValidUTF8_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString541` into `intoString139`.

	// Assume that `sourceCQL` has the underlying type of `fromString541`:
	fromString541 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString541` to result `intoString139`
	// (`intoString139` is now tainted).
	intoString139 := strings.ToValidUTF8(fromString541, "")

	// Return the tainted `intoString139`:
	return intoString139
}

func TaintStepTest_StringsToValidUTF8_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString814` into `intoString768`.

	// Assume that `sourceCQL` has the underlying type of `fromString814`:
	fromString814 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString814` to result `intoString768`
	// (`intoString768` is now tainted).
	intoString768 := strings.ToValidUTF8("", fromString814)

	// Return the tainted `intoString768`:
	return intoString768
}

func TaintStepTest_StringsTrim_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString468` into `intoString736`.

	// Assume that `sourceCQL` has the underlying type of `fromString468`:
	fromString468 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString468` to result `intoString736`
	// (`intoString736` is now tainted).
	intoString736 := strings.Trim(fromString468, "")

	// Return the tainted `intoString736`:
	return intoString736
}

func TaintStepTest_StringsTrimFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString516` into `intoString246`.

	// Assume that `sourceCQL` has the underlying type of `fromString516`:
	fromString516 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString516` to result `intoString246`
	// (`intoString246` is now tainted).
	intoString246 := strings.TrimFunc(fromString516, nil)

	// Return the tainted `intoString246`:
	return intoString246
}

func TaintStepTest_StringsTrimLeft_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString679` into `intoString736`.

	// Assume that `sourceCQL` has the underlying type of `fromString679`:
	fromString679 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString679` to result `intoString736`
	// (`intoString736` is now tainted).
	intoString736 := strings.TrimLeft(fromString679, "")

	// Return the tainted `intoString736`:
	return intoString736
}

func TaintStepTest_StringsTrimLeftFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString839` into `intoString273`.

	// Assume that `sourceCQL` has the underlying type of `fromString839`:
	fromString839 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString839` to result `intoString273`
	// (`intoString273` is now tainted).
	intoString273 := strings.TrimLeftFunc(fromString839, nil)

	// Return the tainted `intoString273`:
	return intoString273
}

func TaintStepTest_StringsTrimPrefix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString982` into `intoString458`.

	// Assume that `sourceCQL` has the underlying type of `fromString982`:
	fromString982 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString982` to result `intoString458`
	// (`intoString458` is now tainted).
	intoString458 := strings.TrimPrefix(fromString982, "")

	// Return the tainted `intoString458`:
	return intoString458
}

func TaintStepTest_StringsTrimRight_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString506` into `intoString213`.

	// Assume that `sourceCQL` has the underlying type of `fromString506`:
	fromString506 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString506` to result `intoString213`
	// (`intoString213` is now tainted).
	intoString213 := strings.TrimRight(fromString506, "")

	// Return the tainted `intoString213`:
	return intoString213
}

func TaintStepTest_StringsTrimRightFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString468` into `intoString219`.

	// Assume that `sourceCQL` has the underlying type of `fromString468`:
	fromString468 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString468` to result `intoString219`
	// (`intoString219` is now tainted).
	intoString219 := strings.TrimRightFunc(fromString468, nil)

	// Return the tainted `intoString219`:
	return intoString219
}

func TaintStepTest_StringsTrimSpace_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString265` into `intoString971`.

	// Assume that `sourceCQL` has the underlying type of `fromString265`:
	fromString265 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString265` to result `intoString971`
	// (`intoString971` is now tainted).
	intoString971 := strings.TrimSpace(fromString265)

	// Return the tainted `intoString971`:
	return intoString971
}

func TaintStepTest_StringsTrimSuffix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString320` into `intoString545`.

	// Assume that `sourceCQL` has the underlying type of `fromString320`:
	fromString320 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString320` to result `intoString545`
	// (`intoString545` is now tainted).
	intoString545 := strings.TrimSuffix(fromString320, "")

	// Return the tainted `intoString545`:
	return intoString545
}

func TaintStepTest_StringsBuilderString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuilder566` into `intoString497`.

	// Assume that `sourceCQL` has the underlying type of `fromBuilder566`:
	fromBuilder566 := sourceCQL.(strings.Builder)

	// Call the method that transfers the taint
	// from the receiver `fromBuilder566` to the result `intoString497`
	// (`intoString497` is now tainted).
	intoString497 := fromBuilder566.String()

	// Return the tainted `intoString497`:
	return intoString497
}

func TaintStepTest_StringsBuilderWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte274` into `intoBuilder783`.

	// Assume that `sourceCQL` has the underlying type of `fromByte274`:
	fromByte274 := sourceCQL.([]byte)

	// Declare `intoBuilder783` variable:
	var intoBuilder783 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `fromByte274` to the receiver `intoBuilder783`
	// (`intoBuilder783` is now tainted).
	intoBuilder783.Write(fromByte274)

	// Return the tainted `intoBuilder783`:
	return intoBuilder783
}

func TaintStepTest_StringsBuilderWriteByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte905` into `intoBuilder389`.

	// Assume that `sourceCQL` has the underlying type of `fromByte905`:
	fromByte905 := sourceCQL.(byte)

	// Declare `intoBuilder389` variable:
	var intoBuilder389 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `fromByte905` to the receiver `intoBuilder389`
	// (`intoBuilder389` is now tainted).
	intoBuilder389.WriteByte(fromByte905)

	// Return the tainted `intoBuilder389`:
	return intoBuilder389
}

func TaintStepTest_StringsBuilderWriteRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune198` into `intoBuilder477`.

	// Assume that `sourceCQL` has the underlying type of `fromRune198`:
	fromRune198 := sourceCQL.(rune)

	// Declare `intoBuilder477` variable:
	var intoBuilder477 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `fromRune198` to the receiver `intoBuilder477`
	// (`intoBuilder477` is now tainted).
	intoBuilder477.WriteRune(fromRune198)

	// Return the tainted `intoBuilder477`:
	return intoBuilder477
}

func TaintStepTest_StringsBuilderWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString544` into `intoBuilder382`.

	// Assume that `sourceCQL` has the underlying type of `fromString544`:
	fromString544 := sourceCQL.(string)

	// Declare `intoBuilder382` variable:
	var intoBuilder382 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `fromString544` to the receiver `intoBuilder382`
	// (`intoBuilder382` is now tainted).
	intoBuilder382.WriteString(fromString544)

	// Return the tainted `intoBuilder382`:
	return intoBuilder382
}

func TaintStepTest_StringsReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader715` into `intoByte179`.

	// Assume that `sourceCQL` has the underlying type of `fromReader715`:
	fromReader715 := sourceCQL.(strings.Reader)

	// Declare `intoByte179` variable:
	var intoByte179 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader715` to the argument `intoByte179`
	// (`intoByte179` is now tainted).
	fromReader715.Read(intoByte179)

	// Return the tainted `intoByte179`:
	return intoByte179
}

func TaintStepTest_StringsReaderReadAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader366` into `intoByte648`.

	// Assume that `sourceCQL` has the underlying type of `fromReader366`:
	fromReader366 := sourceCQL.(strings.Reader)

	// Declare `intoByte648` variable:
	var intoByte648 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader366` to the argument `intoByte648`
	// (`intoByte648` is now tainted).
	fromReader366.ReadAt(intoByte648, 0)

	// Return the tainted `intoByte648`:
	return intoByte648
}

func TaintStepTest_StringsReaderReadByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader544` into `intoByte484`.

	// Assume that `sourceCQL` has the underlying type of `fromReader544`:
	fromReader544 := sourceCQL.(strings.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader544` to the result `intoByte484`
	// (`intoByte484` is now tainted).
	intoByte484, _ := fromReader544.ReadByte()

	// Return the tainted `intoByte484`:
	return intoByte484
}

func TaintStepTest_StringsReaderReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader824` into `intoRune754`.

	// Assume that `sourceCQL` has the underlying type of `fromReader824`:
	fromReader824 := sourceCQL.(strings.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader824` to the result `intoRune754`
	// (`intoRune754` is now tainted).
	intoRune754, _, _ := fromReader824.ReadRune()

	// Return the tainted `intoRune754`:
	return intoRune754
}

func TaintStepTest_StringsReaderReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString680` into `intoReader722`.

	// Assume that `sourceCQL` has the underlying type of `fromString680`:
	fromString680 := sourceCQL.(string)

	// Declare `intoReader722` variable:
	var intoReader722 strings.Reader

	// Call the method that transfers the taint
	// from the parameter `fromString680` to the receiver `intoReader722`
	// (`intoReader722` is now tainted).
	intoReader722.Reset(fromString680)

	// Return the tainted `intoReader722`:
	return intoReader722
}

func TaintStepTest_StringsReaderWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader506` into `intoWriter121`.

	// Assume that `sourceCQL` has the underlying type of `fromReader506`:
	fromReader506 := sourceCQL.(strings.Reader)

	// Declare `intoWriter121` variable:
	var intoWriter121 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromReader506` to the argument `intoWriter121`
	// (`intoWriter121` is now tainted).
	fromReader506.WriteTo(intoWriter121)

	// Return the tainted `intoWriter121`:
	return intoWriter121
}

func TaintStepTest_StringsReplacerReplace_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString293` into `intoString151`.

	// Assume that `sourceCQL` has the underlying type of `fromString293`:
	fromString293 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL strings.Replacer

	// Call the method that transfers the taint
	// from the parameter `fromString293` to the result `intoString151`
	// (`intoString151` is now tainted).
	intoString151 := mediumObjCQL.Replace(fromString293)

	// Return the tainted `intoString151`:
	return intoString151
}

func TaintStepTest_StringsReplacerWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString849` into `intoWriter322`.

	// Assume that `sourceCQL` has the underlying type of `fromString849`:
	fromString849 := sourceCQL.(string)

	// Declare `intoWriter322` variable:
	var intoWriter322 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL strings.Replacer

	// Call the method that transfers the taint
	// from the parameter `fromString849` to the parameter `intoWriter322`
	// (`intoWriter322` is now tainted).
	mediumObjCQL.WriteString(intoWriter322, fromString849)

	// Return the tainted `intoWriter322`:
	return intoWriter322
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
