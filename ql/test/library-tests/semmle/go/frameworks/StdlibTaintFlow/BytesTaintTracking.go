// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"bytes"
	"io"
)

func TaintStepTest_BytesFields_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte656` into `intoByte414`.

	// Assume that `sourceCQL` has the underlying type of `fromByte656`:
	fromByte656 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte656` to result `intoByte414`
	// (`intoByte414` is now tainted).
	intoByte414 := bytes.Fields(fromByte656)

	// Return the tainted `intoByte414`:
	return intoByte414
}

func TaintStepTest_BytesFieldsFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte518` into `intoByte650`.

	// Assume that `sourceCQL` has the underlying type of `fromByte518`:
	fromByte518 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte518` to result `intoByte650`
	// (`intoByte650` is now tainted).
	intoByte650 := bytes.FieldsFunc(fromByte518, nil)

	// Return the tainted `intoByte650`:
	return intoByte650
}

func TaintStepTest_BytesJoin_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte784` into `intoByte957`.

	// Assume that `sourceCQL` has the underlying type of `fromByte784`:
	fromByte784 := sourceCQL.([][]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte784` to result `intoByte957`
	// (`intoByte957` is now tainted).
	intoByte957 := bytes.Join(fromByte784, nil)

	// Return the tainted `intoByte957`:
	return intoByte957
}

func TaintStepTest_BytesJoin_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte520` into `intoByte443`.

	// Assume that `sourceCQL` has the underlying type of `fromByte520`:
	fromByte520 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte520` to result `intoByte443`
	// (`intoByte443` is now tainted).
	intoByte443 := bytes.Join(nil, fromByte520)

	// Return the tainted `intoByte443`:
	return intoByte443
}

func TaintStepTest_BytesMap_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte127` into `intoByte483`.

	// Assume that `sourceCQL` has the underlying type of `fromByte127`:
	fromByte127 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte127` to result `intoByte483`
	// (`intoByte483` is now tainted).
	intoByte483 := bytes.Map(nil, fromByte127)

	// Return the tainted `intoByte483`:
	return intoByte483
}

func TaintStepTest_BytesNewBuffer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte989` into `intoBuffer982`.

	// Assume that `sourceCQL` has the underlying type of `fromByte989`:
	fromByte989 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte989` to result `intoBuffer982`
	// (`intoBuffer982` is now tainted).
	intoBuffer982 := bytes.NewBuffer(fromByte989)

	// Return the tainted `intoBuffer982`:
	return intoBuffer982
}

func TaintStepTest_BytesNewBufferString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString417` into `intoBuffer584`.

	// Assume that `sourceCQL` has the underlying type of `fromString417`:
	fromString417 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString417` to result `intoBuffer584`
	// (`intoBuffer584` is now tainted).
	intoBuffer584 := bytes.NewBufferString(fromString417)

	// Return the tainted `intoBuffer584`:
	return intoBuffer584
}

func TaintStepTest_BytesNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte991` into `intoReader881`.

	// Assume that `sourceCQL` has the underlying type of `fromByte991`:
	fromByte991 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte991` to result `intoReader881`
	// (`intoReader881` is now tainted).
	intoReader881 := bytes.NewReader(fromByte991)

	// Return the tainted `intoReader881`:
	return intoReader881
}

func TaintStepTest_BytesRepeat_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte186` into `intoByte284`.

	// Assume that `sourceCQL` has the underlying type of `fromByte186`:
	fromByte186 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte186` to result `intoByte284`
	// (`intoByte284` is now tainted).
	intoByte284 := bytes.Repeat(fromByte186, 0)

	// Return the tainted `intoByte284`:
	return intoByte284
}

func TaintStepTest_BytesReplace_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte908` into `intoByte137`.

	// Assume that `sourceCQL` has the underlying type of `fromByte908`:
	fromByte908 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte908` to result `intoByte137`
	// (`intoByte137` is now tainted).
	intoByte137 := bytes.Replace(fromByte908, nil, nil, 0)

	// Return the tainted `intoByte137`:
	return intoByte137
}

func TaintStepTest_BytesReplace_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte494` into `intoByte873`.

	// Assume that `sourceCQL` has the underlying type of `fromByte494`:
	fromByte494 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte494` to result `intoByte873`
	// (`intoByte873` is now tainted).
	intoByte873 := bytes.Replace(nil, nil, fromByte494, 0)

	// Return the tainted `intoByte873`:
	return intoByte873
}

func TaintStepTest_BytesReplaceAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte599` into `intoByte409`.

	// Assume that `sourceCQL` has the underlying type of `fromByte599`:
	fromByte599 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte599` to result `intoByte409`
	// (`intoByte409` is now tainted).
	intoByte409 := bytes.ReplaceAll(fromByte599, nil, nil)

	// Return the tainted `intoByte409`:
	return intoByte409
}

func TaintStepTest_BytesReplaceAll_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte246` into `intoByte898`.

	// Assume that `sourceCQL` has the underlying type of `fromByte246`:
	fromByte246 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte246` to result `intoByte898`
	// (`intoByte898` is now tainted).
	intoByte898 := bytes.ReplaceAll(nil, nil, fromByte246)

	// Return the tainted `intoByte898`:
	return intoByte898
}

func TaintStepTest_BytesRunes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte598` into `intoRune631`.

	// Assume that `sourceCQL` has the underlying type of `fromByte598`:
	fromByte598 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte598` to result `intoRune631`
	// (`intoRune631` is now tainted).
	intoRune631 := bytes.Runes(fromByte598)

	// Return the tainted `intoRune631`:
	return intoRune631
}

func TaintStepTest_BytesSplit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte165` into `intoByte150`.

	// Assume that `sourceCQL` has the underlying type of `fromByte165`:
	fromByte165 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte165` to result `intoByte150`
	// (`intoByte150` is now tainted).
	intoByte150 := bytes.Split(fromByte165, nil)

	// Return the tainted `intoByte150`:
	return intoByte150
}

func TaintStepTest_BytesSplitAfter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte340` into `intoByte471`.

	// Assume that `sourceCQL` has the underlying type of `fromByte340`:
	fromByte340 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte340` to result `intoByte471`
	// (`intoByte471` is now tainted).
	intoByte471 := bytes.SplitAfter(fromByte340, nil)

	// Return the tainted `intoByte471`:
	return intoByte471
}

func TaintStepTest_BytesSplitAfterN_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte290` into `intoByte758`.

	// Assume that `sourceCQL` has the underlying type of `fromByte290`:
	fromByte290 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte290` to result `intoByte758`
	// (`intoByte758` is now tainted).
	intoByte758 := bytes.SplitAfterN(fromByte290, nil, 0)

	// Return the tainted `intoByte758`:
	return intoByte758
}

func TaintStepTest_BytesSplitN_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte396` into `intoByte707`.

	// Assume that `sourceCQL` has the underlying type of `fromByte396`:
	fromByte396 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte396` to result `intoByte707`
	// (`intoByte707` is now tainted).
	intoByte707 := bytes.SplitN(fromByte396, nil, 0)

	// Return the tainted `intoByte707`:
	return intoByte707
}

func TaintStepTest_BytesTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte912` into `intoByte718`.

	// Assume that `sourceCQL` has the underlying type of `fromByte912`:
	fromByte912 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte912` to result `intoByte718`
	// (`intoByte718` is now tainted).
	intoByte718 := bytes.Title(fromByte912)

	// Return the tainted `intoByte718`:
	return intoByte718
}

func TaintStepTest_BytesToLower_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte972` into `intoByte633`.

	// Assume that `sourceCQL` has the underlying type of `fromByte972`:
	fromByte972 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte972` to result `intoByte633`
	// (`intoByte633` is now tainted).
	intoByte633 := bytes.ToLower(fromByte972)

	// Return the tainted `intoByte633`:
	return intoByte633
}

func TaintStepTest_BytesToLowerSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte316` into `intoByte145`.

	// Assume that `sourceCQL` has the underlying type of `fromByte316`:
	fromByte316 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte316` to result `intoByte145`
	// (`intoByte145` is now tainted).
	intoByte145 := bytes.ToLowerSpecial(nil, fromByte316)

	// Return the tainted `intoByte145`:
	return intoByte145
}

func TaintStepTest_BytesToTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte817` into `intoByte474`.

	// Assume that `sourceCQL` has the underlying type of `fromByte817`:
	fromByte817 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte817` to result `intoByte474`
	// (`intoByte474` is now tainted).
	intoByte474 := bytes.ToTitle(fromByte817)

	// Return the tainted `intoByte474`:
	return intoByte474
}

func TaintStepTest_BytesToTitleSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte832` into `intoByte378`.

	// Assume that `sourceCQL` has the underlying type of `fromByte832`:
	fromByte832 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte832` to result `intoByte378`
	// (`intoByte378` is now tainted).
	intoByte378 := bytes.ToTitleSpecial(nil, fromByte832)

	// Return the tainted `intoByte378`:
	return intoByte378
}

func TaintStepTest_BytesToUpper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte541` into `intoByte139`.

	// Assume that `sourceCQL` has the underlying type of `fromByte541`:
	fromByte541 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte541` to result `intoByte139`
	// (`intoByte139` is now tainted).
	intoByte139 := bytes.ToUpper(fromByte541)

	// Return the tainted `intoByte139`:
	return intoByte139
}

func TaintStepTest_BytesToUpperSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte814` into `intoByte768`.

	// Assume that `sourceCQL` has the underlying type of `fromByte814`:
	fromByte814 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte814` to result `intoByte768`
	// (`intoByte768` is now tainted).
	intoByte768 := bytes.ToUpperSpecial(nil, fromByte814)

	// Return the tainted `intoByte768`:
	return intoByte768
}

func TaintStepTest_BytesToValidUTF8_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte468` into `intoByte736`.

	// Assume that `sourceCQL` has the underlying type of `fromByte468`:
	fromByte468 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte468` to result `intoByte736`
	// (`intoByte736` is now tainted).
	intoByte736 := bytes.ToValidUTF8(fromByte468, nil)

	// Return the tainted `intoByte736`:
	return intoByte736
}

func TaintStepTest_BytesToValidUTF8_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte516` into `intoByte246`.

	// Assume that `sourceCQL` has the underlying type of `fromByte516`:
	fromByte516 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte516` to result `intoByte246`
	// (`intoByte246` is now tainted).
	intoByte246 := bytes.ToValidUTF8(nil, fromByte516)

	// Return the tainted `intoByte246`:
	return intoByte246
}

func TaintStepTest_BytesTrim_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte679` into `intoByte736`.

	// Assume that `sourceCQL` has the underlying type of `fromByte679`:
	fromByte679 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte679` to result `intoByte736`
	// (`intoByte736` is now tainted).
	intoByte736 := bytes.Trim(fromByte679, "")

	// Return the tainted `intoByte736`:
	return intoByte736
}

func TaintStepTest_BytesTrimFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte839` into `intoByte273`.

	// Assume that `sourceCQL` has the underlying type of `fromByte839`:
	fromByte839 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte839` to result `intoByte273`
	// (`intoByte273` is now tainted).
	intoByte273 := bytes.TrimFunc(fromByte839, nil)

	// Return the tainted `intoByte273`:
	return intoByte273
}

func TaintStepTest_BytesTrimLeft_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte982` into `intoByte458`.

	// Assume that `sourceCQL` has the underlying type of `fromByte982`:
	fromByte982 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte982` to result `intoByte458`
	// (`intoByte458` is now tainted).
	intoByte458 := bytes.TrimLeft(fromByte982, "")

	// Return the tainted `intoByte458`:
	return intoByte458
}

func TaintStepTest_BytesTrimLeftFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte506` into `intoByte213`.

	// Assume that `sourceCQL` has the underlying type of `fromByte506`:
	fromByte506 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte506` to result `intoByte213`
	// (`intoByte213` is now tainted).
	intoByte213 := bytes.TrimLeftFunc(fromByte506, nil)

	// Return the tainted `intoByte213`:
	return intoByte213
}

func TaintStepTest_BytesTrimPrefix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte468` into `intoByte219`.

	// Assume that `sourceCQL` has the underlying type of `fromByte468`:
	fromByte468 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte468` to result `intoByte219`
	// (`intoByte219` is now tainted).
	intoByte219 := bytes.TrimPrefix(fromByte468, nil)

	// Return the tainted `intoByte219`:
	return intoByte219
}

func TaintStepTest_BytesTrimRight_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte265` into `intoByte971`.

	// Assume that `sourceCQL` has the underlying type of `fromByte265`:
	fromByte265 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte265` to result `intoByte971`
	// (`intoByte971` is now tainted).
	intoByte971 := bytes.TrimRight(fromByte265, "")

	// Return the tainted `intoByte971`:
	return intoByte971
}

func TaintStepTest_BytesTrimRightFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte320` into `intoByte545`.

	// Assume that `sourceCQL` has the underlying type of `fromByte320`:
	fromByte320 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte320` to result `intoByte545`
	// (`intoByte545` is now tainted).
	intoByte545 := bytes.TrimRightFunc(fromByte320, nil)

	// Return the tainted `intoByte545`:
	return intoByte545
}

func TaintStepTest_BytesTrimSpace_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte566` into `intoByte497`.

	// Assume that `sourceCQL` has the underlying type of `fromByte566`:
	fromByte566 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte566` to result `intoByte497`
	// (`intoByte497` is now tainted).
	intoByte497 := bytes.TrimSpace(fromByte566)

	// Return the tainted `intoByte497`:
	return intoByte497
}

func TaintStepTest_BytesTrimSuffix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte274` into `intoByte783`.

	// Assume that `sourceCQL` has the underlying type of `fromByte274`:
	fromByte274 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte274` to result `intoByte783`
	// (`intoByte783` is now tainted).
	intoByte783 := bytes.TrimSuffix(fromByte274, nil)

	// Return the tainted `intoByte783`:
	return intoByte783
}

func TaintStepTest_BytesBufferBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer905` into `intoByte389`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer905`:
	fromBuffer905 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer905` to the result `intoByte389`
	// (`intoByte389` is now tainted).
	intoByte389 := fromBuffer905.Bytes()

	// Return the tainted `intoByte389`:
	return intoByte389
}

func TaintStepTest_BytesBufferNext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer198` into `intoByte477`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer198`:
	fromBuffer198 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer198` to the result `intoByte477`
	// (`intoByte477` is now tainted).
	intoByte477 := fromBuffer198.Next(0)

	// Return the tainted `intoByte477`:
	return intoByte477
}

func TaintStepTest_BytesBufferRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer544` into `intoByte382`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer544`:
	fromBuffer544 := sourceCQL.(bytes.Buffer)

	// Declare `intoByte382` variable:
	var intoByte382 []byte

	// Call the method that transfers the taint
	// from the receiver `fromBuffer544` to the argument `intoByte382`
	// (`intoByte382` is now tainted).
	fromBuffer544.Read(intoByte382)

	// Return the tainted `intoByte382`:
	return intoByte382
}

func TaintStepTest_BytesBufferReadByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer715` into `intoByte179`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer715`:
	fromBuffer715 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer715` to the result `intoByte179`
	// (`intoByte179` is now tainted).
	intoByte179, _ := fromBuffer715.ReadByte()

	// Return the tainted `intoByte179`:
	return intoByte179
}

func TaintStepTest_BytesBufferReadBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer366` into `intoByte648`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer366`:
	fromBuffer366 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer366` to the result `intoByte648`
	// (`intoByte648` is now tainted).
	intoByte648, _ := fromBuffer366.ReadBytes(0)

	// Return the tainted `intoByte648`:
	return intoByte648
}

func TaintStepTest_BytesBufferReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader544` into `intoBuffer484`.

	// Assume that `sourceCQL` has the underlying type of `fromReader544`:
	fromReader544 := sourceCQL.(io.Reader)

	// Declare `intoBuffer484` variable:
	var intoBuffer484 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromReader544` to the receiver `intoBuffer484`
	// (`intoBuffer484` is now tainted).
	intoBuffer484.ReadFrom(fromReader544)

	// Return the tainted `intoBuffer484`:
	return intoBuffer484
}

func TaintStepTest_BytesBufferReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer824` into `intoRune754`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer824`:
	fromBuffer824 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer824` to the result `intoRune754`
	// (`intoRune754` is now tainted).
	intoRune754, _, _ := fromBuffer824.ReadRune()

	// Return the tainted `intoRune754`:
	return intoRune754
}

func TaintStepTest_BytesBufferReadString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer680` into `intoString722`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer680`:
	fromBuffer680 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer680` to the result `intoString722`
	// (`intoString722` is now tainted).
	intoString722, _ := fromBuffer680.ReadString(0)

	// Return the tainted `intoString722`:
	return intoString722
}

func TaintStepTest_BytesBufferString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer506` into `intoString121`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer506`:
	fromBuffer506 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer506` to the result `intoString121`
	// (`intoString121` is now tainted).
	intoString121 := fromBuffer506.String()

	// Return the tainted `intoString121`:
	return intoString121
}

func TaintStepTest_BytesBufferWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte293` into `intoBuffer151`.

	// Assume that `sourceCQL` has the underlying type of `fromByte293`:
	fromByte293 := sourceCQL.([]byte)

	// Declare `intoBuffer151` variable:
	var intoBuffer151 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromByte293` to the receiver `intoBuffer151`
	// (`intoBuffer151` is now tainted).
	intoBuffer151.Write(fromByte293)

	// Return the tainted `intoBuffer151`:
	return intoBuffer151
}

func TaintStepTest_BytesBufferWriteByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte849` into `intoBuffer322`.

	// Assume that `sourceCQL` has the underlying type of `fromByte849`:
	fromByte849 := sourceCQL.(byte)

	// Declare `intoBuffer322` variable:
	var intoBuffer322 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromByte849` to the receiver `intoBuffer322`
	// (`intoBuffer322` is now tainted).
	intoBuffer322.WriteByte(fromByte849)

	// Return the tainted `intoBuffer322`:
	return intoBuffer322
}

func TaintStepTest_BytesBufferWriteRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune339` into `intoBuffer478`.

	// Assume that `sourceCQL` has the underlying type of `fromRune339`:
	fromRune339 := sourceCQL.(rune)

	// Declare `intoBuffer478` variable:
	var intoBuffer478 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromRune339` to the receiver `intoBuffer478`
	// (`intoBuffer478` is now tainted).
	intoBuffer478.WriteRune(fromRune339)

	// Return the tainted `intoBuffer478`:
	return intoBuffer478
}

func TaintStepTest_BytesBufferWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString399` into `intoBuffer426`.

	// Assume that `sourceCQL` has the underlying type of `fromString399`:
	fromString399 := sourceCQL.(string)

	// Declare `intoBuffer426` variable:
	var intoBuffer426 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromString399` to the receiver `intoBuffer426`
	// (`intoBuffer426` is now tainted).
	intoBuffer426.WriteString(fromString399)

	// Return the tainted `intoBuffer426`:
	return intoBuffer426
}

func TaintStepTest_BytesBufferWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer628` into `intoWriter197`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer628`:
	fromBuffer628 := sourceCQL.(bytes.Buffer)

	// Declare `intoWriter197` variable:
	var intoWriter197 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromBuffer628` to the argument `intoWriter197`
	// (`intoWriter197` is now tainted).
	fromBuffer628.WriteTo(intoWriter197)

	// Return the tainted `intoWriter197`:
	return intoWriter197
}

func TaintStepTest_BytesReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader216` into `intoByte742`.

	// Assume that `sourceCQL` has the underlying type of `fromReader216`:
	fromReader216 := sourceCQL.(bytes.Reader)

	// Declare `intoByte742` variable:
	var intoByte742 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader216` to the argument `intoByte742`
	// (`intoByte742` is now tainted).
	fromReader216.Read(intoByte742)

	// Return the tainted `intoByte742`:
	return intoByte742
}

func TaintStepTest_BytesReaderReadAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader906` into `intoByte620`.

	// Assume that `sourceCQL` has the underlying type of `fromReader906`:
	fromReader906 := sourceCQL.(bytes.Reader)

	// Declare `intoByte620` variable:
	var intoByte620 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader906` to the argument `intoByte620`
	// (`intoByte620` is now tainted).
	fromReader906.ReadAt(intoByte620, 0)

	// Return the tainted `intoByte620`:
	return intoByte620
}

func TaintStepTest_BytesReaderReadByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader158` into `intoByte353`.

	// Assume that `sourceCQL` has the underlying type of `fromReader158`:
	fromReader158 := sourceCQL.(bytes.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader158` to the result `intoByte353`
	// (`intoByte353` is now tainted).
	intoByte353, _ := fromReader158.ReadByte()

	// Return the tainted `intoByte353`:
	return intoByte353
}

func TaintStepTest_BytesReaderReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader625` into `intoRune340`.

	// Assume that `sourceCQL` has the underlying type of `fromReader625`:
	fromReader625 := sourceCQL.(bytes.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader625` to the result `intoRune340`
	// (`intoRune340` is now tainted).
	intoRune340, _, _ := fromReader625.ReadRune()

	// Return the tainted `intoRune340`:
	return intoRune340
}

func TaintStepTest_BytesReaderReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte741` into `intoReader199`.

	// Assume that `sourceCQL` has the underlying type of `fromByte741`:
	fromByte741 := sourceCQL.([]byte)

	// Declare `intoReader199` variable:
	var intoReader199 bytes.Reader

	// Call the method that transfers the taint
	// from the parameter `fromByte741` to the receiver `intoReader199`
	// (`intoReader199` is now tainted).
	intoReader199.Reset(fromByte741)

	// Return the tainted `intoReader199`:
	return intoReader199
}

func TaintStepTest_BytesReaderWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader873` into `intoWriter304`.

	// Assume that `sourceCQL` has the underlying type of `fromReader873`:
	fromReader873 := sourceCQL.(bytes.Reader)

	// Declare `intoWriter304` variable:
	var intoWriter304 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromReader873` to the argument `intoWriter304`
	// (`intoWriter304` is now tainted).
	fromReader873.WriteTo(intoWriter304)

	// Return the tainted `intoWriter304`:
	return intoWriter304
}

func RunAllTaints_Bytes() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesFields_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesFieldsFunc_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesJoin_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesJoin_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesMap_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesNewBuffer_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesNewBufferString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesNewReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesRepeat_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesReplace_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesReplace_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesReplaceAll_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesReplaceAll_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesRunes_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesSplit_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesSplitAfter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesSplitAfterN_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesSplitN_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesTitle_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesToLower_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesToLowerSpecial_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesToTitle_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesToTitleSpecial_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesToUpper_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesToUpperSpecial_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesToValidUTF8_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesToValidUTF8_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesTrim_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesTrimFunc_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesTrimLeft_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesTrimLeftFunc_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesTrimPrefix_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesTrimRight_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesTrimRightFunc_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesTrimSpace_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesTrimSuffix_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesBufferBytes_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesBufferNext_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesBufferRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesBufferReadByte_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesBufferReadBytes_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesBufferReadFrom_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesBufferReadRune_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesBufferReadString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesBufferString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesBufferWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesBufferWriteByte_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesBufferWriteRune_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesBufferWriteString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesBufferWriteTo_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesReaderRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesReaderReadAt_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesReaderReadByte_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesReaderReadRune_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesReaderReset_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_BytesReaderWriteTo_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
