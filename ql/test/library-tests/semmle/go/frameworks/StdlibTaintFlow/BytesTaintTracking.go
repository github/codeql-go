// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"bytes"
	"io"
)

func TaintStepTest_BytesFields_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte466` into `intoByte753`.

	// Assume that `sourceCQL` has the underlying type of `fromByte466`:
	fromByte466 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte466` to result `intoByte753`
	// (`intoByte753` is now tainted).
	intoByte753 := bytes.Fields(fromByte466)

	// Return the tainted `intoByte753`:
	return intoByte753
}

func TaintStepTest_BytesFieldsFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte242` into `intoByte935`.

	// Assume that `sourceCQL` has the underlying type of `fromByte242`:
	fromByte242 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte242` to result `intoByte935`
	// (`intoByte935` is now tainted).
	intoByte935 := bytes.FieldsFunc(fromByte242, nil)

	// Return the tainted `intoByte935`:
	return intoByte935
}

func TaintStepTest_BytesJoin_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte559` into `intoByte696`.

	// Assume that `sourceCQL` has the underlying type of `fromByte559`:
	fromByte559 := sourceCQL.([][]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte559` to result `intoByte696`
	// (`intoByte696` is now tainted).
	intoByte696 := bytes.Join(fromByte559, nil)

	// Return the tainted `intoByte696`:
	return intoByte696
}

func TaintStepTest_BytesJoin_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte773` into `intoByte867`.

	// Assume that `sourceCQL` has the underlying type of `fromByte773`:
	fromByte773 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte773` to result `intoByte867`
	// (`intoByte867` is now tainted).
	intoByte867 := bytes.Join(nil, fromByte773)

	// Return the tainted `intoByte867`:
	return intoByte867
}

func TaintStepTest_BytesMap_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte944` into `intoByte499`.

	// Assume that `sourceCQL` has the underlying type of `fromByte944`:
	fromByte944 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte944` to result `intoByte499`
	// (`intoByte499` is now tainted).
	intoByte499 := bytes.Map(nil, fromByte944)

	// Return the tainted `intoByte499`:
	return intoByte499
}

func TaintStepTest_BytesNewBuffer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte536` into `intoBuffer232`.

	// Assume that `sourceCQL` has the underlying type of `fromByte536`:
	fromByte536 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte536` to result `intoBuffer232`
	// (`intoBuffer232` is now tainted).
	intoBuffer232 := bytes.NewBuffer(fromByte536)

	// Return the tainted `intoBuffer232`:
	return intoBuffer232
}

func TaintStepTest_BytesNewBufferString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString142` into `intoBuffer142`.

	// Assume that `sourceCQL` has the underlying type of `fromString142`:
	fromString142 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString142` to result `intoBuffer142`
	// (`intoBuffer142` is now tainted).
	intoBuffer142 := bytes.NewBufferString(fromString142)

	// Return the tainted `intoBuffer142`:
	return intoBuffer142
}

func TaintStepTest_BytesNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte911` into `intoReader445`.

	// Assume that `sourceCQL` has the underlying type of `fromByte911`:
	fromByte911 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte911` to result `intoReader445`
	// (`intoReader445` is now tainted).
	intoReader445 := bytes.NewReader(fromByte911)

	// Return the tainted `intoReader445`:
	return intoReader445
}

func TaintStepTest_BytesRepeat_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte309` into `intoByte148`.

	// Assume that `sourceCQL` has the underlying type of `fromByte309`:
	fromByte309 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte309` to result `intoByte148`
	// (`intoByte148` is now tainted).
	intoByte148 := bytes.Repeat(fromByte309, 0)

	// Return the tainted `intoByte148`:
	return intoByte148
}

func TaintStepTest_BytesReplace_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte587` into `intoByte821`.

	// Assume that `sourceCQL` has the underlying type of `fromByte587`:
	fromByte587 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte587` to result `intoByte821`
	// (`intoByte821` is now tainted).
	intoByte821 := bytes.Replace(fromByte587, nil, nil, 0)

	// Return the tainted `intoByte821`:
	return intoByte821
}

func TaintStepTest_BytesReplace_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte798` into `intoByte431`.

	// Assume that `sourceCQL` has the underlying type of `fromByte798`:
	fromByte798 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte798` to result `intoByte431`
	// (`intoByte431` is now tainted).
	intoByte431 := bytes.Replace(nil, nil, fromByte798, 0)

	// Return the tainted `intoByte431`:
	return intoByte431
}

func TaintStepTest_BytesReplaceAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte793` into `intoByte617`.

	// Assume that `sourceCQL` has the underlying type of `fromByte793`:
	fromByte793 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte793` to result `intoByte617`
	// (`intoByte617` is now tainted).
	intoByte617 := bytes.ReplaceAll(fromByte793, nil, nil)

	// Return the tainted `intoByte617`:
	return intoByte617
}

func TaintStepTest_BytesReplaceAll_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte596` into `intoByte514`.

	// Assume that `sourceCQL` has the underlying type of `fromByte596`:
	fromByte596 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte596` to result `intoByte514`
	// (`intoByte514` is now tainted).
	intoByte514 := bytes.ReplaceAll(nil, nil, fromByte596)

	// Return the tainted `intoByte514`:
	return intoByte514
}

func TaintStepTest_BytesRunes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte958` into `intoRune252`.

	// Assume that `sourceCQL` has the underlying type of `fromByte958`:
	fromByte958 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte958` to result `intoRune252`
	// (`intoRune252` is now tainted).
	intoRune252 := bytes.Runes(fromByte958)

	// Return the tainted `intoRune252`:
	return intoRune252
}

func TaintStepTest_BytesSplit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte556` into `intoByte630`.

	// Assume that `sourceCQL` has the underlying type of `fromByte556`:
	fromByte556 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte556` to result `intoByte630`
	// (`intoByte630` is now tainted).
	intoByte630 := bytes.Split(fromByte556, nil)

	// Return the tainted `intoByte630`:
	return intoByte630
}

func TaintStepTest_BytesSplitAfter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte154` into `intoByte759`.

	// Assume that `sourceCQL` has the underlying type of `fromByte154`:
	fromByte154 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte154` to result `intoByte759`
	// (`intoByte759` is now tainted).
	intoByte759 := bytes.SplitAfter(fromByte154, nil)

	// Return the tainted `intoByte759`:
	return intoByte759
}

func TaintStepTest_BytesSplitAfterN_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte149` into `intoByte982`.

	// Assume that `sourceCQL` has the underlying type of `fromByte149`:
	fromByte149 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte149` to result `intoByte982`
	// (`intoByte982` is now tainted).
	intoByte982 := bytes.SplitAfterN(fromByte149, nil, 0)

	// Return the tainted `intoByte982`:
	return intoByte982
}

func TaintStepTest_BytesSplitN_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte485` into `intoByte995`.

	// Assume that `sourceCQL` has the underlying type of `fromByte485`:
	fromByte485 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte485` to result `intoByte995`
	// (`intoByte995` is now tainted).
	intoByte995 := bytes.SplitN(fromByte485, nil, 0)

	// Return the tainted `intoByte995`:
	return intoByte995
}

func TaintStepTest_BytesTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte342` into `intoByte147`.

	// Assume that `sourceCQL` has the underlying type of `fromByte342`:
	fromByte342 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte342` to result `intoByte147`
	// (`intoByte147` is now tainted).
	intoByte147 := bytes.Title(fromByte342)

	// Return the tainted `intoByte147`:
	return intoByte147
}

func TaintStepTest_BytesToLower_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte549` into `intoByte717`.

	// Assume that `sourceCQL` has the underlying type of `fromByte549`:
	fromByte549 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte549` to result `intoByte717`
	// (`intoByte717` is now tainted).
	intoByte717 := bytes.ToLower(fromByte549)

	// Return the tainted `intoByte717`:
	return intoByte717
}

func TaintStepTest_BytesToLowerSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte250` into `intoByte577`.

	// Assume that `sourceCQL` has the underlying type of `fromByte250`:
	fromByte250 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte250` to result `intoByte577`
	// (`intoByte577` is now tainted).
	intoByte577 := bytes.ToLowerSpecial(nil, fromByte250)

	// Return the tainted `intoByte577`:
	return intoByte577
}

func TaintStepTest_BytesToTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte406` into `intoByte904`.

	// Assume that `sourceCQL` has the underlying type of `fromByte406`:
	fromByte406 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte406` to result `intoByte904`
	// (`intoByte904` is now tainted).
	intoByte904 := bytes.ToTitle(fromByte406)

	// Return the tainted `intoByte904`:
	return intoByte904
}

func TaintStepTest_BytesToTitleSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte345` into `intoByte719`.

	// Assume that `sourceCQL` has the underlying type of `fromByte345`:
	fromByte345 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte345` to result `intoByte719`
	// (`intoByte719` is now tainted).
	intoByte719 := bytes.ToTitleSpecial(nil, fromByte345)

	// Return the tainted `intoByte719`:
	return intoByte719
}

func TaintStepTest_BytesToUpper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte804` into `intoByte203`.

	// Assume that `sourceCQL` has the underlying type of `fromByte804`:
	fromByte804 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte804` to result `intoByte203`
	// (`intoByte203` is now tainted).
	intoByte203 := bytes.ToUpper(fromByte804)

	// Return the tainted `intoByte203`:
	return intoByte203
}

func TaintStepTest_BytesToUpperSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte793` into `intoByte672`.

	// Assume that `sourceCQL` has the underlying type of `fromByte793`:
	fromByte793 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte793` to result `intoByte672`
	// (`intoByte672` is now tainted).
	intoByte672 := bytes.ToUpperSpecial(nil, fromByte793)

	// Return the tainted `intoByte672`:
	return intoByte672
}

func TaintStepTest_BytesToValidUTF8_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte699` into `intoByte256`.

	// Assume that `sourceCQL` has the underlying type of `fromByte699`:
	fromByte699 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte699` to result `intoByte256`
	// (`intoByte256` is now tainted).
	intoByte256 := bytes.ToValidUTF8(fromByte699, nil)

	// Return the tainted `intoByte256`:
	return intoByte256
}

func TaintStepTest_BytesToValidUTF8_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte422` into `intoByte850`.

	// Assume that `sourceCQL` has the underlying type of `fromByte422`:
	fromByte422 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte422` to result `intoByte850`
	// (`intoByte850` is now tainted).
	intoByte850 := bytes.ToValidUTF8(nil, fromByte422)

	// Return the tainted `intoByte850`:
	return intoByte850
}

func TaintStepTest_BytesTrim_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte119` into `intoByte522`.

	// Assume that `sourceCQL` has the underlying type of `fromByte119`:
	fromByte119 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte119` to result `intoByte522`
	// (`intoByte522` is now tainted).
	intoByte522 := bytes.Trim(fromByte119, "")

	// Return the tainted `intoByte522`:
	return intoByte522
}

func TaintStepTest_BytesTrimFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte689` into `intoByte911`.

	// Assume that `sourceCQL` has the underlying type of `fromByte689`:
	fromByte689 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte689` to result `intoByte911`
	// (`intoByte911` is now tainted).
	intoByte911 := bytes.TrimFunc(fromByte689, nil)

	// Return the tainted `intoByte911`:
	return intoByte911
}

func TaintStepTest_BytesTrimLeft_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte878` into `intoByte393`.

	// Assume that `sourceCQL` has the underlying type of `fromByte878`:
	fromByte878 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte878` to result `intoByte393`
	// (`intoByte393` is now tainted).
	intoByte393 := bytes.TrimLeft(fromByte878, "")

	// Return the tainted `intoByte393`:
	return intoByte393
}

func TaintStepTest_BytesTrimLeftFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte649` into `intoByte456`.

	// Assume that `sourceCQL` has the underlying type of `fromByte649`:
	fromByte649 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte649` to result `intoByte456`
	// (`intoByte456` is now tainted).
	intoByte456 := bytes.TrimLeftFunc(fromByte649, nil)

	// Return the tainted `intoByte456`:
	return intoByte456
}

func TaintStepTest_BytesTrimPrefix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte627` into `intoByte986`.

	// Assume that `sourceCQL` has the underlying type of `fromByte627`:
	fromByte627 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte627` to result `intoByte986`
	// (`intoByte986` is now tainted).
	intoByte986 := bytes.TrimPrefix(fromByte627, nil)

	// Return the tainted `intoByte986`:
	return intoByte986
}

func TaintStepTest_BytesTrimRight_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte191` into `intoByte312`.

	// Assume that `sourceCQL` has the underlying type of `fromByte191`:
	fromByte191 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte191` to result `intoByte312`
	// (`intoByte312` is now tainted).
	intoByte312 := bytes.TrimRight(fromByte191, "")

	// Return the tainted `intoByte312`:
	return intoByte312
}

func TaintStepTest_BytesTrimRightFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte601` into `intoByte147`.

	// Assume that `sourceCQL` has the underlying type of `fromByte601`:
	fromByte601 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte601` to result `intoByte147`
	// (`intoByte147` is now tainted).
	intoByte147 := bytes.TrimRightFunc(fromByte601, nil)

	// Return the tainted `intoByte147`:
	return intoByte147
}

func TaintStepTest_BytesTrimSpace_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte194` into `intoByte316`.

	// Assume that `sourceCQL` has the underlying type of `fromByte194`:
	fromByte194 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte194` to result `intoByte316`
	// (`intoByte316` is now tainted).
	intoByte316 := bytes.TrimSpace(fromByte194)

	// Return the tainted `intoByte316`:
	return intoByte316
}

func TaintStepTest_BytesTrimSuffix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte284` into `intoByte912`.

	// Assume that `sourceCQL` has the underlying type of `fromByte284`:
	fromByte284 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte284` to result `intoByte912`
	// (`intoByte912` is now tainted).
	intoByte912 := bytes.TrimSuffix(fromByte284, nil)

	// Return the tainted `intoByte912`:
	return intoByte912
}

func TaintStepTest_BytesBufferBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer779` into `intoByte425`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer779`:
	fromBuffer779 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer779` to the result `intoByte425`
	// (`intoByte425` is now tainted).
	intoByte425 := fromBuffer779.Bytes()

	// Return the tainted `intoByte425`:
	return intoByte425
}

func TaintStepTest_BytesBufferNext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer141` into `intoByte672`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer141`:
	fromBuffer141 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer141` to the result `intoByte672`
	// (`intoByte672` is now tainted).
	intoByte672 := fromBuffer141.Next(0)

	// Return the tainted `intoByte672`:
	return intoByte672
}

func TaintStepTest_BytesBufferRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer667` into `intoByte945`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer667`:
	fromBuffer667 := sourceCQL.(bytes.Buffer)

	// Declare `intoByte945` variable:
	var intoByte945 []byte

	// Call the method that transfers the taint
	// from the receiver `fromBuffer667` to the argument `intoByte945`
	// (`intoByte945` is now tainted).
	fromBuffer667.Read(intoByte945)

	// Return the tainted `intoByte945`:
	return intoByte945
}

func TaintStepTest_BytesBufferReadByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer183` into `intoByte218`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer183`:
	fromBuffer183 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer183` to the result `intoByte218`
	// (`intoByte218` is now tainted).
	intoByte218, _ := fromBuffer183.ReadByte()

	// Return the tainted `intoByte218`:
	return intoByte218
}

func TaintStepTest_BytesBufferReadBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer663` into `intoByte179`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer663`:
	fromBuffer663 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer663` to the result `intoByte179`
	// (`intoByte179` is now tainted).
	intoByte179, _ := fromBuffer663.ReadBytes(0)

	// Return the tainted `intoByte179`:
	return intoByte179
}

func TaintStepTest_BytesBufferReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader886` into `intoBuffer279`.

	// Assume that `sourceCQL` has the underlying type of `fromReader886`:
	fromReader886 := sourceCQL.(io.Reader)

	// Declare `intoBuffer279` variable:
	var intoBuffer279 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromReader886` to the receiver `intoBuffer279`
	// (`intoBuffer279` is now tainted).
	intoBuffer279.ReadFrom(fromReader886)

	// Return the tainted `intoBuffer279`:
	return intoBuffer279
}

func TaintStepTest_BytesBufferReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer876` into `intoRune356`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer876`:
	fromBuffer876 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer876` to the result `intoRune356`
	// (`intoRune356` is now tainted).
	intoRune356, _, _ := fromBuffer876.ReadRune()

	// Return the tainted `intoRune356`:
	return intoRune356
}

func TaintStepTest_BytesBufferReadString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer228` into `intoString525`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer228`:
	fromBuffer228 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer228` to the result `intoString525`
	// (`intoString525` is now tainted).
	intoString525, _ := fromBuffer228.ReadString(0)

	// Return the tainted `intoString525`:
	return intoString525
}

func TaintStepTest_BytesBufferString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer765` into `intoString279`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer765`:
	fromBuffer765 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer765` to the result `intoString279`
	// (`intoString279` is now tainted).
	intoString279 := fromBuffer765.String()

	// Return the tainted `intoString279`:
	return intoString279
}

func TaintStepTest_BytesBufferWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte479` into `intoBuffer601`.

	// Assume that `sourceCQL` has the underlying type of `fromByte479`:
	fromByte479 := sourceCQL.([]byte)

	// Declare `intoBuffer601` variable:
	var intoBuffer601 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromByte479` to the receiver `intoBuffer601`
	// (`intoBuffer601` is now tainted).
	intoBuffer601.Write(fromByte479)

	// Return the tainted `intoBuffer601`:
	return intoBuffer601
}

func TaintStepTest_BytesBufferWriteByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte801` into `intoBuffer623`.

	// Assume that `sourceCQL` has the underlying type of `fromByte801`:
	fromByte801 := sourceCQL.(byte)

	// Declare `intoBuffer623` variable:
	var intoBuffer623 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromByte801` to the receiver `intoBuffer623`
	// (`intoBuffer623` is now tainted).
	intoBuffer623.WriteByte(fromByte801)

	// Return the tainted `intoBuffer623`:
	return intoBuffer623
}

func TaintStepTest_BytesBufferWriteRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune496` into `intoBuffer143`.

	// Assume that `sourceCQL` has the underlying type of `fromRune496`:
	fromRune496 := sourceCQL.(rune)

	// Declare `intoBuffer143` variable:
	var intoBuffer143 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromRune496` to the receiver `intoBuffer143`
	// (`intoBuffer143` is now tainted).
	intoBuffer143.WriteRune(fromRune496)

	// Return the tainted `intoBuffer143`:
	return intoBuffer143
}

func TaintStepTest_BytesBufferWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString712` into `intoBuffer887`.

	// Assume that `sourceCQL` has the underlying type of `fromString712`:
	fromString712 := sourceCQL.(string)

	// Declare `intoBuffer887` variable:
	var intoBuffer887 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromString712` to the receiver `intoBuffer887`
	// (`intoBuffer887` is now tainted).
	intoBuffer887.WriteString(fromString712)

	// Return the tainted `intoBuffer887`:
	return intoBuffer887
}

func TaintStepTest_BytesBufferWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer749` into `intoWriter860`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer749`:
	fromBuffer749 := sourceCQL.(bytes.Buffer)

	// Declare `intoWriter860` variable:
	var intoWriter860 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromBuffer749` to the argument `intoWriter860`
	// (`intoWriter860` is now tainted).
	fromBuffer749.WriteTo(intoWriter860)

	// Return the tainted `intoWriter860`:
	return intoWriter860
}

func TaintStepTest_BytesReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader692` into `intoByte544`.

	// Assume that `sourceCQL` has the underlying type of `fromReader692`:
	fromReader692 := sourceCQL.(bytes.Reader)

	// Declare `intoByte544` variable:
	var intoByte544 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader692` to the argument `intoByte544`
	// (`intoByte544` is now tainted).
	fromReader692.Read(intoByte544)

	// Return the tainted `intoByte544`:
	return intoByte544
}

func TaintStepTest_BytesReaderReadAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader690` into `intoByte786`.

	// Assume that `sourceCQL` has the underlying type of `fromReader690`:
	fromReader690 := sourceCQL.(bytes.Reader)

	// Declare `intoByte786` variable:
	var intoByte786 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader690` to the argument `intoByte786`
	// (`intoByte786` is now tainted).
	fromReader690.ReadAt(intoByte786, 0)

	// Return the tainted `intoByte786`:
	return intoByte786
}

func TaintStepTest_BytesReaderReadByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader323` into `intoByte746`.

	// Assume that `sourceCQL` has the underlying type of `fromReader323`:
	fromReader323 := sourceCQL.(bytes.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader323` to the result `intoByte746`
	// (`intoByte746` is now tainted).
	intoByte746, _ := fromReader323.ReadByte()

	// Return the tainted `intoByte746`:
	return intoByte746
}

func TaintStepTest_BytesReaderReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader822` into `intoRune292`.

	// Assume that `sourceCQL` has the underlying type of `fromReader822`:
	fromReader822 := sourceCQL.(bytes.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader822` to the result `intoRune292`
	// (`intoRune292` is now tainted).
	intoRune292, _, _ := fromReader822.ReadRune()

	// Return the tainted `intoRune292`:
	return intoRune292
}

func TaintStepTest_BytesReaderReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte712` into `intoReader532`.

	// Assume that `sourceCQL` has the underlying type of `fromByte712`:
	fromByte712 := sourceCQL.([]byte)

	// Declare `intoReader532` variable:
	var intoReader532 bytes.Reader

	// Call the method that transfers the taint
	// from the parameter `fromByte712` to the receiver `intoReader532`
	// (`intoReader532` is now tainted).
	intoReader532.Reset(fromByte712)

	// Return the tainted `intoReader532`:
	return intoReader532
}

func TaintStepTest_BytesReaderWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader310` into `intoWriter448`.

	// Assume that `sourceCQL` has the underlying type of `fromReader310`:
	fromReader310 := sourceCQL.(bytes.Reader)

	// Declare `intoWriter448` variable:
	var intoWriter448 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromReader310` to the argument `intoWriter448`
	// (`intoWriter448` is now tainted).
	fromReader310.WriteTo(intoWriter448)

	// Return the tainted `intoWriter448`:
	return intoWriter448
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
