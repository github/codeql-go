// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"strings"
)

func TaintStepTest_StringsFields_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString825` into `intoString600`.

	// Assume that `sourceCQL` has the underlying type of `fromString825`:
	fromString825 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString825` to result `intoString600`
	// (`intoString600` is now tainted).
	intoString600 := strings.Fields(fromString825)

	// Return the tainted `intoString600`:
	return intoString600
}

func TaintStepTest_StringsFieldsFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString799` into `intoString505`.

	// Assume that `sourceCQL` has the underlying type of `fromString799`:
	fromString799 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString799` to result `intoString505`
	// (`intoString505` is now tainted).
	intoString505 := strings.FieldsFunc(fromString799, nil)

	// Return the tainted `intoString505`:
	return intoString505
}

func TaintStepTest_StringsJoin_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString374` into `intoString507`.

	// Assume that `sourceCQL` has the underlying type of `fromString374`:
	fromString374 := sourceCQL.([]string)

	// Call the function that transfers the taint
	// from the parameter `fromString374` to result `intoString507`
	// (`intoString507` is now tainted).
	intoString507 := strings.Join(fromString374, "")

	// Return the tainted `intoString507`:
	return intoString507
}

func TaintStepTest_StringsJoin_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString770` into `intoString897`.

	// Assume that `sourceCQL` has the underlying type of `fromString770`:
	fromString770 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString770` to result `intoString897`
	// (`intoString897` is now tainted).
	intoString897 := strings.Join(nil, fromString770)

	// Return the tainted `intoString897`:
	return intoString897
}

func TaintStepTest_StringsMap_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString654` into `intoString549`.

	// Assume that `sourceCQL` has the underlying type of `fromString654`:
	fromString654 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString654` to result `intoString549`
	// (`intoString549` is now tainted).
	intoString549 := strings.Map(nil, fromString654)

	// Return the tainted `intoString549`:
	return intoString549
}

func TaintStepTest_StringsNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString883` into `intoReader386`.

	// Assume that `sourceCQL` has the underlying type of `fromString883`:
	fromString883 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString883` to result `intoReader386`
	// (`intoReader386` is now tainted).
	intoReader386 := strings.NewReader(fromString883)

	// Return the tainted `intoReader386`:
	return intoReader386
}

func TaintStepTest_StringsNewReplacer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString843` into `intoReplacer846`.

	// Assume that `sourceCQL` has the underlying type of `fromString843`:
	fromString843 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString843` to result `intoReplacer846`
	// (`intoReplacer846` is now tainted).
	intoReplacer846 := strings.NewReplacer(fromString843)

	// Return the tainted `intoReplacer846`:
	return intoReplacer846
}

func TaintStepTest_StringsRepeat_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString490` into `intoString976`.

	// Assume that `sourceCQL` has the underlying type of `fromString490`:
	fromString490 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString490` to result `intoString976`
	// (`intoString976` is now tainted).
	intoString976 := strings.Repeat(fromString490, 0)

	// Return the tainted `intoString976`:
	return intoString976
}

func TaintStepTest_StringsReplace_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString596` into `intoString739`.

	// Assume that `sourceCQL` has the underlying type of `fromString596`:
	fromString596 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString596` to result `intoString739`
	// (`intoString739` is now tainted).
	intoString739 := strings.Replace(fromString596, "", "", 0)

	// Return the tainted `intoString739`:
	return intoString739
}

func TaintStepTest_StringsReplace_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString150` into `intoString384`.

	// Assume that `sourceCQL` has the underlying type of `fromString150`:
	fromString150 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString150` to result `intoString384`
	// (`intoString384` is now tainted).
	intoString384 := strings.Replace("", "", fromString150, 0)

	// Return the tainted `intoString384`:
	return intoString384
}

func TaintStepTest_StringsReplaceAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString560` into `intoString401`.

	// Assume that `sourceCQL` has the underlying type of `fromString560`:
	fromString560 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString560` to result `intoString401`
	// (`intoString401` is now tainted).
	intoString401 := strings.ReplaceAll(fromString560, "", "")

	// Return the tainted `intoString401`:
	return intoString401
}

func TaintStepTest_StringsReplaceAll_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString387` into `intoString714`.

	// Assume that `sourceCQL` has the underlying type of `fromString387`:
	fromString387 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString387` to result `intoString714`
	// (`intoString714` is now tainted).
	intoString714 := strings.ReplaceAll("", "", fromString387)

	// Return the tainted `intoString714`:
	return intoString714
}

func TaintStepTest_StringsSplit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString866` into `intoString806`.

	// Assume that `sourceCQL` has the underlying type of `fromString866`:
	fromString866 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString866` to result `intoString806`
	// (`intoString806` is now tainted).
	intoString806 := strings.Split(fromString866, "")

	// Return the tainted `intoString806`:
	return intoString806
}

func TaintStepTest_StringsSplitAfter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString953` into `intoString949`.

	// Assume that `sourceCQL` has the underlying type of `fromString953`:
	fromString953 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString953` to result `intoString949`
	// (`intoString949` is now tainted).
	intoString949 := strings.SplitAfter(fromString953, "")

	// Return the tainted `intoString949`:
	return intoString949
}

func TaintStepTest_StringsSplitAfterN_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString921` into `intoString307`.

	// Assume that `sourceCQL` has the underlying type of `fromString921`:
	fromString921 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString921` to result `intoString307`
	// (`intoString307` is now tainted).
	intoString307 := strings.SplitAfterN(fromString921, "", 0)

	// Return the tainted `intoString307`:
	return intoString307
}

func TaintStepTest_StringsSplitN_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString729` into `intoString716`.

	// Assume that `sourceCQL` has the underlying type of `fromString729`:
	fromString729 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString729` to result `intoString716`
	// (`intoString716` is now tainted).
	intoString716 := strings.SplitN(fromString729, "", 0)

	// Return the tainted `intoString716`:
	return intoString716
}

func TaintStepTest_StringsTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString410` into `intoString944`.

	// Assume that `sourceCQL` has the underlying type of `fromString410`:
	fromString410 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString410` to result `intoString944`
	// (`intoString944` is now tainted).
	intoString944 := strings.Title(fromString410)

	// Return the tainted `intoString944`:
	return intoString944
}

func TaintStepTest_StringsToLower_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString430` into `intoString996`.

	// Assume that `sourceCQL` has the underlying type of `fromString430`:
	fromString430 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString430` to result `intoString996`
	// (`intoString996` is now tainted).
	intoString996 := strings.ToLower(fromString430)

	// Return the tainted `intoString996`:
	return intoString996
}

func TaintStepTest_StringsToLowerSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString766` into `intoString848`.

	// Assume that `sourceCQL` has the underlying type of `fromString766`:
	fromString766 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString766` to result `intoString848`
	// (`intoString848` is now tainted).
	intoString848 := strings.ToLowerSpecial(nil, fromString766)

	// Return the tainted `intoString848`:
	return intoString848
}

func TaintStepTest_StringsToTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString490` into `intoString987`.

	// Assume that `sourceCQL` has the underlying type of `fromString490`:
	fromString490 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString490` to result `intoString987`
	// (`intoString987` is now tainted).
	intoString987 := strings.ToTitle(fromString490)

	// Return the tainted `intoString987`:
	return intoString987
}

func TaintStepTest_StringsToTitleSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString884` into `intoString338`.

	// Assume that `sourceCQL` has the underlying type of `fromString884`:
	fromString884 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString884` to result `intoString338`
	// (`intoString338` is now tainted).
	intoString338 := strings.ToTitleSpecial(nil, fromString884)

	// Return the tainted `intoString338`:
	return intoString338
}

func TaintStepTest_StringsToUpper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString625` into `intoString673`.

	// Assume that `sourceCQL` has the underlying type of `fromString625`:
	fromString625 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString625` to result `intoString673`
	// (`intoString673` is now tainted).
	intoString673 := strings.ToUpper(fromString625)

	// Return the tainted `intoString673`:
	return intoString673
}

func TaintStepTest_StringsToUpperSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString222` into `intoString262`.

	// Assume that `sourceCQL` has the underlying type of `fromString222`:
	fromString222 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString222` to result `intoString262`
	// (`intoString262` is now tainted).
	intoString262 := strings.ToUpperSpecial(nil, fromString222)

	// Return the tainted `intoString262`:
	return intoString262
}

func TaintStepTest_StringsToValidUTF8_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString485` into `intoString534`.

	// Assume that `sourceCQL` has the underlying type of `fromString485`:
	fromString485 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString485` to result `intoString534`
	// (`intoString534` is now tainted).
	intoString534 := strings.ToValidUTF8(fromString485, "")

	// Return the tainted `intoString534`:
	return intoString534
}

func TaintStepTest_StringsToValidUTF8_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString458` into `intoString310`.

	// Assume that `sourceCQL` has the underlying type of `fromString458`:
	fromString458 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString458` to result `intoString310`
	// (`intoString310` is now tainted).
	intoString310 := strings.ToValidUTF8("", fromString458)

	// Return the tainted `intoString310`:
	return intoString310
}

func TaintStepTest_StringsTrim_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString475` into `intoString552`.

	// Assume that `sourceCQL` has the underlying type of `fromString475`:
	fromString475 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString475` to result `intoString552`
	// (`intoString552` is now tainted).
	intoString552 := strings.Trim(fromString475, "")

	// Return the tainted `intoString552`:
	return intoString552
}

func TaintStepTest_StringsTrimFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString622` into `intoString390`.

	// Assume that `sourceCQL` has the underlying type of `fromString622`:
	fromString622 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString622` to result `intoString390`
	// (`intoString390` is now tainted).
	intoString390 := strings.TrimFunc(fromString622, nil)

	// Return the tainted `intoString390`:
	return intoString390
}

func TaintStepTest_StringsTrimLeft_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString820` into `intoString673`.

	// Assume that `sourceCQL` has the underlying type of `fromString820`:
	fromString820 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString820` to result `intoString673`
	// (`intoString673` is now tainted).
	intoString673 := strings.TrimLeft(fromString820, "")

	// Return the tainted `intoString673`:
	return intoString673
}

func TaintStepTest_StringsTrimLeftFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString197` into `intoString497`.

	// Assume that `sourceCQL` has the underlying type of `fromString197`:
	fromString197 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString197` to result `intoString497`
	// (`intoString497` is now tainted).
	intoString497 := strings.TrimLeftFunc(fromString197, nil)

	// Return the tainted `intoString497`:
	return intoString497
}

func TaintStepTest_StringsTrimPrefix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString323` into `intoString307`.

	// Assume that `sourceCQL` has the underlying type of `fromString323`:
	fromString323 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString323` to result `intoString307`
	// (`intoString307` is now tainted).
	intoString307 := strings.TrimPrefix(fromString323, "")

	// Return the tainted `intoString307`:
	return intoString307
}

func TaintStepTest_StringsTrimRight_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString532` into `intoString167`.

	// Assume that `sourceCQL` has the underlying type of `fromString532`:
	fromString532 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString532` to result `intoString167`
	// (`intoString167` is now tainted).
	intoString167 := strings.TrimRight(fromString532, "")

	// Return the tainted `intoString167`:
	return intoString167
}

func TaintStepTest_StringsTrimRightFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString638` into `intoString277`.

	// Assume that `sourceCQL` has the underlying type of `fromString638`:
	fromString638 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString638` to result `intoString277`
	// (`intoString277` is now tainted).
	intoString277 := strings.TrimRightFunc(fromString638, nil)

	// Return the tainted `intoString277`:
	return intoString277
}

func TaintStepTest_StringsTrimSpace_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString411` into `intoString846`.

	// Assume that `sourceCQL` has the underlying type of `fromString411`:
	fromString411 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString411` to result `intoString846`
	// (`intoString846` is now tainted).
	intoString846 := strings.TrimSpace(fromString411)

	// Return the tainted `intoString846`:
	return intoString846
}

func TaintStepTest_StringsTrimSuffix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString615` into `intoString862`.

	// Assume that `sourceCQL` has the underlying type of `fromString615`:
	fromString615 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString615` to result `intoString862`
	// (`intoString862` is now tainted).
	intoString862 := strings.TrimSuffix(fromString615, "")

	// Return the tainted `intoString862`:
	return intoString862
}

func TaintStepTest_StringsBuilderString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuilder843` into `intoString676`.

	// Assume that `sourceCQL` has the underlying type of `fromBuilder843`:
	fromBuilder843 := sourceCQL.(strings.Builder)

	// Call the method that transfers the taint
	// from the receiver `fromBuilder843` to the result `intoString676`
	// (`intoString676` is now tainted).
	intoString676 := fromBuilder843.String()

	// Return the tainted `intoString676`:
	return intoString676
}

func TaintStepTest_StringsBuilderWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte987` into `intoBuilder364`.

	// Assume that `sourceCQL` has the underlying type of `fromByte987`:
	fromByte987 := sourceCQL.([]byte)

	// Declare `intoBuilder364` variable:
	var intoBuilder364 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `fromByte987` to the receiver `intoBuilder364`
	// (`intoBuilder364` is now tainted).
	intoBuilder364.Write(fromByte987)

	// Return the tainted `intoBuilder364`:
	return intoBuilder364
}

func TaintStepTest_StringsBuilderWriteByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte860` into `intoBuilder515`.

	// Assume that `sourceCQL` has the underlying type of `fromByte860`:
	fromByte860 := sourceCQL.(byte)

	// Declare `intoBuilder515` variable:
	var intoBuilder515 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `fromByte860` to the receiver `intoBuilder515`
	// (`intoBuilder515` is now tainted).
	intoBuilder515.WriteByte(fromByte860)

	// Return the tainted `intoBuilder515`:
	return intoBuilder515
}

func TaintStepTest_StringsBuilderWriteRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune322` into `intoBuilder934`.

	// Assume that `sourceCQL` has the underlying type of `fromRune322`:
	fromRune322 := sourceCQL.(rune)

	// Declare `intoBuilder934` variable:
	var intoBuilder934 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `fromRune322` to the receiver `intoBuilder934`
	// (`intoBuilder934` is now tainted).
	intoBuilder934.WriteRune(fromRune322)

	// Return the tainted `intoBuilder934`:
	return intoBuilder934
}

func TaintStepTest_StringsBuilderWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString870` into `intoBuilder197`.

	// Assume that `sourceCQL` has the underlying type of `fromString870`:
	fromString870 := sourceCQL.(string)

	// Declare `intoBuilder197` variable:
	var intoBuilder197 strings.Builder

	// Call the method that transfers the taint
	// from the parameter `fromString870` to the receiver `intoBuilder197`
	// (`intoBuilder197` is now tainted).
	intoBuilder197.WriteString(fromString870)

	// Return the tainted `intoBuilder197`:
	return intoBuilder197
}

func TaintStepTest_StringsReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader183` into `intoByte254`.

	// Assume that `sourceCQL` has the underlying type of `fromReader183`:
	fromReader183 := sourceCQL.(strings.Reader)

	// Declare `intoByte254` variable:
	var intoByte254 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader183` to the argument `intoByte254`
	// (`intoByte254` is now tainted).
	fromReader183.Read(intoByte254)

	// Return the tainted `intoByte254`:
	return intoByte254
}

func TaintStepTest_StringsReaderReadAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader560` into `intoByte556`.

	// Assume that `sourceCQL` has the underlying type of `fromReader560`:
	fromReader560 := sourceCQL.(strings.Reader)

	// Declare `intoByte556` variable:
	var intoByte556 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader560` to the argument `intoByte556`
	// (`intoByte556` is now tainted).
	fromReader560.ReadAt(intoByte556, 0)

	// Return the tainted `intoByte556`:
	return intoByte556
}

func TaintStepTest_StringsReaderReadByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader331` into `intoByte131`.

	// Assume that `sourceCQL` has the underlying type of `fromReader331`:
	fromReader331 := sourceCQL.(strings.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader331` to the result `intoByte131`
	// (`intoByte131` is now tainted).
	intoByte131, _ := fromReader331.ReadByte()

	// Return the tainted `intoByte131`:
	return intoByte131
}

func TaintStepTest_StringsReaderReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader822` into `intoRune545`.

	// Assume that `sourceCQL` has the underlying type of `fromReader822`:
	fromReader822 := sourceCQL.(strings.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader822` to the result `intoRune545`
	// (`intoRune545` is now tainted).
	intoRune545, _, _ := fromReader822.ReadRune()

	// Return the tainted `intoRune545`:
	return intoRune545
}

func TaintStepTest_StringsReaderReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString477` into `intoReader276`.

	// Assume that `sourceCQL` has the underlying type of `fromString477`:
	fromString477 := sourceCQL.(string)

	// Declare `intoReader276` variable:
	var intoReader276 strings.Reader

	// Call the method that transfers the taint
	// from the parameter `fromString477` to the receiver `intoReader276`
	// (`intoReader276` is now tainted).
	intoReader276.Reset(fromString477)

	// Return the tainted `intoReader276`:
	return intoReader276
}

func TaintStepTest_StringsReaderWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader359` into `intoWriter417`.

	// Assume that `sourceCQL` has the underlying type of `fromReader359`:
	fromReader359 := sourceCQL.(strings.Reader)

	// Declare `intoWriter417` variable:
	var intoWriter417 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromReader359` to the argument `intoWriter417`
	// (`intoWriter417` is now tainted).
	fromReader359.WriteTo(intoWriter417)

	// Return the tainted `intoWriter417`:
	return intoWriter417
}

func TaintStepTest_StringsReplacerReplace_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString272` into `intoString873`.

	// Assume that `sourceCQL` has the underlying type of `fromString272`:
	fromString272 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL strings.Replacer

	// Call the method that transfers the taint
	// from the parameter `fromString272` to the result `intoString873`
	// (`intoString873` is now tainted).
	intoString873 := mediumObjCQL.Replace(fromString272)

	// Return the tainted `intoString873`:
	return intoString873
}

func TaintStepTest_StringsReplacerWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString165` into `intoWriter327`.

	// Assume that `sourceCQL` has the underlying type of `fromString165`:
	fromString165 := sourceCQL.(string)

	// Declare `intoWriter327` variable:
	var intoWriter327 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL strings.Replacer

	// Call the method that transfers the taint
	// from the parameter `fromString165` to the parameter `intoWriter327`
	// (`intoWriter327` is now tainted).
	mediumObjCQL.WriteString(intoWriter327, fromString165)

	// Return the tainted `intoWriter327`:
	return intoWriter327
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
