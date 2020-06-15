package main

import (
	"bytes"
	"io"
)

func TaintStepTest_BytesFields_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte705` into `intoByte280`.

	// Assume that `sourceCQL` has the underlying type of `fromByte705`:
	fromByte705 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte705` to result `intoByte280`
	// (`intoByte280` is now tainted).
	intoByte280 := bytes.Fields(fromByte705)

	// Sink the tainted `intoByte280`:
	sink(intoByte280)
}

func TaintStepTest_BytesFieldsFunc_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte857` into `intoByte358`.

	// Assume that `sourceCQL` has the underlying type of `fromByte857`:
	fromByte857 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte857` to result `intoByte358`
	// (`intoByte358` is now tainted).
	intoByte358 := bytes.FieldsFunc(fromByte857, nil)

	// Sink the tainted `intoByte358`:
	sink(intoByte358)
}

func TaintStepTest_BytesJoin_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte737` into `intoByte394`.

	// Assume that `sourceCQL` has the underlying type of `fromByte737`:
	fromByte737 := sourceCQL.([][]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte737` to result `intoByte394`
	// (`intoByte394` is now tainted).
	intoByte394 := bytes.Join(fromByte737, nil)

	// Sink the tainted `intoByte394`:
	sink(intoByte394)
}

func TaintStepTest_BytesJoin_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromByte230` into `intoByte226`.

	// Assume that `sourceCQL` has the underlying type of `fromByte230`:
	fromByte230 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte230` to result `intoByte226`
	// (`intoByte226` is now tainted).
	intoByte226 := bytes.Join(nil, fromByte230)

	// Sink the tainted `intoByte226`:
	sink(intoByte226)
}

func TaintStepTest_BytesMap_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte210` into `intoByte877`.

	// Assume that `sourceCQL` has the underlying type of `fromByte210`:
	fromByte210 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte210` to result `intoByte877`
	// (`intoByte877` is now tainted).
	intoByte877 := bytes.Map(nil, fromByte210)

	// Sink the tainted `intoByte877`:
	sink(intoByte877)
}

func TaintStepTest_BytesNewBuffer_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte385` into `intoBuffer742`.

	// Assume that `sourceCQL` has the underlying type of `fromByte385`:
	fromByte385 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte385` to result `intoBuffer742`
	// (`intoBuffer742` is now tainted).
	intoBuffer742 := bytes.NewBuffer(fromByte385)

	// Sink the tainted `intoBuffer742`:
	sink(intoBuffer742)
}

func TaintStepTest_BytesNewBufferString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString907` into `intoBuffer764`.

	// Assume that `sourceCQL` has the underlying type of `fromString907`:
	fromString907 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString907` to result `intoBuffer764`
	// (`intoBuffer764` is now tainted).
	intoBuffer764 := bytes.NewBufferString(fromString907)

	// Sink the tainted `intoBuffer764`:
	sink(intoBuffer764)
}

func TaintStepTest_BytesNewReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte905` into `intoReader532`.

	// Assume that `sourceCQL` has the underlying type of `fromByte905`:
	fromByte905 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte905` to result `intoReader532`
	// (`intoReader532` is now tainted).
	intoReader532 := bytes.NewReader(fromByte905)

	// Sink the tainted `intoReader532`:
	sink(intoReader532)
}

func TaintStepTest_BytesRepeat_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte394` into `intoByte973`.

	// Assume that `sourceCQL` has the underlying type of `fromByte394`:
	fromByte394 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte394` to result `intoByte973`
	// (`intoByte973` is now tainted).
	intoByte973 := bytes.Repeat(fromByte394, 0)

	// Sink the tainted `intoByte973`:
	sink(intoByte973)
}

func TaintStepTest_BytesReplace_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte783` into `intoByte223`.

	// Assume that `sourceCQL` has the underlying type of `fromByte783`:
	fromByte783 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte783` to result `intoByte223`
	// (`intoByte223` is now tainted).
	intoByte223 := bytes.Replace(fromByte783, nil, nil, 0)

	// Sink the tainted `intoByte223`:
	sink(intoByte223)
}

func TaintStepTest_BytesReplace_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromByte366` into `intoByte481`.

	// Assume that `sourceCQL` has the underlying type of `fromByte366`:
	fromByte366 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte366` to result `intoByte481`
	// (`intoByte481` is now tainted).
	intoByte481 := bytes.Replace(nil, nil, fromByte366, 0)

	// Sink the tainted `intoByte481`:
	sink(intoByte481)
}

func TaintStepTest_BytesReplaceAll_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte529` into `intoByte696`.

	// Assume that `sourceCQL` has the underlying type of `fromByte529`:
	fromByte529 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte529` to result `intoByte696`
	// (`intoByte696` is now tainted).
	intoByte696 := bytes.ReplaceAll(fromByte529, nil, nil)

	// Sink the tainted `intoByte696`:
	sink(intoByte696)
}

func TaintStepTest_BytesReplaceAll_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromByte780` into `intoByte457`.

	// Assume that `sourceCQL` has the underlying type of `fromByte780`:
	fromByte780 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte780` to result `intoByte457`
	// (`intoByte457` is now tainted).
	intoByte457 := bytes.ReplaceAll(nil, nil, fromByte780)

	// Sink the tainted `intoByte457`:
	sink(intoByte457)
}

func TaintStepTest_BytesRunes_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte158` into `intoRune995`.

	// Assume that `sourceCQL` has the underlying type of `fromByte158`:
	fromByte158 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte158` to result `intoRune995`
	// (`intoRune995` is now tainted).
	intoRune995 := bytes.Runes(fromByte158)

	// Sink the tainted `intoRune995`:
	sink(intoRune995)
}

func TaintStepTest_BytesSplit_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte816` into `intoByte773`.

	// Assume that `sourceCQL` has the underlying type of `fromByte816`:
	fromByte816 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte816` to result `intoByte773`
	// (`intoByte773` is now tainted).
	intoByte773 := bytes.Split(fromByte816, nil)

	// Sink the tainted `intoByte773`:
	sink(intoByte773)
}

func TaintStepTest_BytesSplitAfter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte789` into `intoByte187`.

	// Assume that `sourceCQL` has the underlying type of `fromByte789`:
	fromByte789 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte789` to result `intoByte187`
	// (`intoByte187` is now tainted).
	intoByte187 := bytes.SplitAfter(fromByte789, nil)

	// Sink the tainted `intoByte187`:
	sink(intoByte187)
}

func TaintStepTest_BytesSplitAfterN_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte997` into `intoByte148`.

	// Assume that `sourceCQL` has the underlying type of `fromByte997`:
	fromByte997 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte997` to result `intoByte148`
	// (`intoByte148` is now tainted).
	intoByte148 := bytes.SplitAfterN(fromByte997, nil, 0)

	// Sink the tainted `intoByte148`:
	sink(intoByte148)
}

func TaintStepTest_BytesSplitN_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte426` into `intoByte643`.

	// Assume that `sourceCQL` has the underlying type of `fromByte426`:
	fromByte426 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte426` to result `intoByte643`
	// (`intoByte643` is now tainted).
	intoByte643 := bytes.SplitN(fromByte426, nil, 0)

	// Sink the tainted `intoByte643`:
	sink(intoByte643)
}

func TaintStepTest_BytesTitle_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte910` into `intoByte186`.

	// Assume that `sourceCQL` has the underlying type of `fromByte910`:
	fromByte910 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte910` to result `intoByte186`
	// (`intoByte186` is now tainted).
	intoByte186 := bytes.Title(fromByte910)

	// Sink the tainted `intoByte186`:
	sink(intoByte186)
}

func TaintStepTest_BytesToLower_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte621` into `intoByte524`.

	// Assume that `sourceCQL` has the underlying type of `fromByte621`:
	fromByte621 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte621` to result `intoByte524`
	// (`intoByte524` is now tainted).
	intoByte524 := bytes.ToLower(fromByte621)

	// Sink the tainted `intoByte524`:
	sink(intoByte524)
}

func TaintStepTest_BytesToLowerSpecial_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte339` into `intoByte659`.

	// Assume that `sourceCQL` has the underlying type of `fromByte339`:
	fromByte339 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte339` to result `intoByte659`
	// (`intoByte659` is now tainted).
	intoByte659 := bytes.ToLowerSpecial(nil, fromByte339)

	// Sink the tainted `intoByte659`:
	sink(intoByte659)
}

func TaintStepTest_BytesToTitle_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte204` into `intoByte929`.

	// Assume that `sourceCQL` has the underlying type of `fromByte204`:
	fromByte204 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte204` to result `intoByte929`
	// (`intoByte929` is now tainted).
	intoByte929 := bytes.ToTitle(fromByte204)

	// Sink the tainted `intoByte929`:
	sink(intoByte929)
}

func TaintStepTest_BytesToTitleSpecial_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte477` into `intoByte961`.

	// Assume that `sourceCQL` has the underlying type of `fromByte477`:
	fromByte477 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte477` to result `intoByte961`
	// (`intoByte961` is now tainted).
	intoByte961 := bytes.ToTitleSpecial(nil, fromByte477)

	// Sink the tainted `intoByte961`:
	sink(intoByte961)
}

func TaintStepTest_BytesToUpper_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte416` into `intoByte387`.

	// Assume that `sourceCQL` has the underlying type of `fromByte416`:
	fromByte416 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte416` to result `intoByte387`
	// (`intoByte387` is now tainted).
	intoByte387 := bytes.ToUpper(fromByte416)

	// Sink the tainted `intoByte387`:
	sink(intoByte387)
}

func TaintStepTest_BytesToUpperSpecial_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte716` into `intoByte373`.

	// Assume that `sourceCQL` has the underlying type of `fromByte716`:
	fromByte716 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte716` to result `intoByte373`
	// (`intoByte373` is now tainted).
	intoByte373 := bytes.ToUpperSpecial(nil, fromByte716)

	// Sink the tainted `intoByte373`:
	sink(intoByte373)
}

func TaintStepTest_BytesToValidUTF8_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte366` into `intoByte518`.

	// Assume that `sourceCQL` has the underlying type of `fromByte366`:
	fromByte366 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte366` to result `intoByte518`
	// (`intoByte518` is now tainted).
	intoByte518 := bytes.ToValidUTF8(fromByte366, nil)

	// Sink the tainted `intoByte518`:
	sink(intoByte518)
}

func TaintStepTest_BytesToValidUTF8_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromByte297` into `intoByte253`.

	// Assume that `sourceCQL` has the underlying type of `fromByte297`:
	fromByte297 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte297` to result `intoByte253`
	// (`intoByte253` is now tainted).
	intoByte253 := bytes.ToValidUTF8(nil, fromByte297)

	// Sink the tainted `intoByte253`:
	sink(intoByte253)
}

func TaintStepTest_BytesTrim_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte520` into `intoByte132`.

	// Assume that `sourceCQL` has the underlying type of `fromByte520`:
	fromByte520 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte520` to result `intoByte132`
	// (`intoByte132` is now tainted).
	intoByte132 := bytes.Trim(fromByte520, "")

	// Sink the tainted `intoByte132`:
	sink(intoByte132)
}

func TaintStepTest_BytesTrimFunc_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte297` into `intoByte774`.

	// Assume that `sourceCQL` has the underlying type of `fromByte297`:
	fromByte297 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte297` to result `intoByte774`
	// (`intoByte774` is now tainted).
	intoByte774 := bytes.TrimFunc(fromByte297, nil)

	// Sink the tainted `intoByte774`:
	sink(intoByte774)
}

func TaintStepTest_BytesTrimLeft_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte662` into `intoByte333`.

	// Assume that `sourceCQL` has the underlying type of `fromByte662`:
	fromByte662 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte662` to result `intoByte333`
	// (`intoByte333` is now tainted).
	intoByte333 := bytes.TrimLeft(fromByte662, "")

	// Sink the tainted `intoByte333`:
	sink(intoByte333)
}

func TaintStepTest_BytesTrimLeftFunc_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte682` into `intoByte548`.

	// Assume that `sourceCQL` has the underlying type of `fromByte682`:
	fromByte682 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte682` to result `intoByte548`
	// (`intoByte548` is now tainted).
	intoByte548 := bytes.TrimLeftFunc(fromByte682, nil)

	// Sink the tainted `intoByte548`:
	sink(intoByte548)
}

func TaintStepTest_BytesTrimPrefix_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte240` into `intoByte696`.

	// Assume that `sourceCQL` has the underlying type of `fromByte240`:
	fromByte240 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte240` to result `intoByte696`
	// (`intoByte696` is now tainted).
	intoByte696 := bytes.TrimPrefix(fromByte240, nil)

	// Sink the tainted `intoByte696`:
	sink(intoByte696)
}

func TaintStepTest_BytesTrimRight_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte762` into `intoByte338`.

	// Assume that `sourceCQL` has the underlying type of `fromByte762`:
	fromByte762 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte762` to result `intoByte338`
	// (`intoByte338` is now tainted).
	intoByte338 := bytes.TrimRight(fromByte762, "")

	// Sink the tainted `intoByte338`:
	sink(intoByte338)
}

func TaintStepTest_BytesTrimRightFunc_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte291` into `intoByte281`.

	// Assume that `sourceCQL` has the underlying type of `fromByte291`:
	fromByte291 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte291` to result `intoByte281`
	// (`intoByte281` is now tainted).
	intoByte281 := bytes.TrimRightFunc(fromByte291, nil)

	// Sink the tainted `intoByte281`:
	sink(intoByte281)
}

func TaintStepTest_BytesTrimSpace_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte595` into `intoByte327`.

	// Assume that `sourceCQL` has the underlying type of `fromByte595`:
	fromByte595 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte595` to result `intoByte327`
	// (`intoByte327` is now tainted).
	intoByte327 := bytes.TrimSpace(fromByte595)

	// Sink the tainted `intoByte327`:
	sink(intoByte327)
}

func TaintStepTest_BytesTrimSuffix_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte646` into `intoByte252`.

	// Assume that `sourceCQL` has the underlying type of `fromByte646`:
	fromByte646 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte646` to result `intoByte252`
	// (`intoByte252` is now tainted).
	intoByte252 := bytes.TrimSuffix(fromByte646, nil)

	// Sink the tainted `intoByte252`:
	sink(intoByte252)
}

func TaintStepTest_BytesBufferBytes_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBuffer807` into `intoByte872`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer807`:
	fromBuffer807 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer807` to the result `intoByte872`
	// (`intoByte872` is now tainted).
	intoByte872 := fromBuffer807.Bytes()

	// Sink the tainted `intoByte872`:
	sink(intoByte872)
}

func TaintStepTest_BytesBufferNext_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBuffer671` into `intoByte978`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer671`:
	fromBuffer671 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer671` to the result `intoByte978`
	// (`intoByte978` is now tainted).
	intoByte978 := fromBuffer671.Next(0)

	// Sink the tainted `intoByte978`:
	sink(intoByte978)
}

func TaintStepTest_BytesBufferRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBuffer565` into `intoByte753`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer565`:
	fromBuffer565 := sourceCQL.(bytes.Buffer)

	// Declare `intoByte753` variable:
	var intoByte753 []byte

	// Call the method that transfers the taint
	// from the receiver `fromBuffer565` to the argument `intoByte753`
	// (`intoByte753` is now tainted).
	fromBuffer565.Read(intoByte753)

	// Sink the tainted `intoByte753`:
	sink(intoByte753)
}

func TaintStepTest_BytesBufferReadByte_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBuffer740` into `intoByte153`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer740`:
	fromBuffer740 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer740` to the result `intoByte153`
	// (`intoByte153` is now tainted).
	intoByte153, _ := fromBuffer740.ReadByte()

	// Sink the tainted `intoByte153`:
	sink(intoByte153)
}

func TaintStepTest_BytesBufferReadBytes_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBuffer978` into `intoByte532`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer978`:
	fromBuffer978 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer978` to the result `intoByte532`
	// (`intoByte532` is now tainted).
	intoByte532, _ := fromBuffer978.ReadBytes(0)

	// Sink the tainted `intoByte532`:
	sink(intoByte532)
}

func TaintStepTest_BytesBufferReadFrom_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader602` into `intoBuffer314`.

	// Assume that `sourceCQL` has the underlying type of `fromReader602`:
	fromReader602 := sourceCQL.(io.Reader)

	// Declare `intoBuffer314` variable:
	var intoBuffer314 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromReader602` to the receiver `intoBuffer314`
	// (`intoBuffer314` is now tainted).
	intoBuffer314.ReadFrom(fromReader602)

	// Sink the tainted `intoBuffer314`:
	sink(intoBuffer314)
}

func TaintStepTest_BytesBufferReadRune_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBuffer829` into `intoRune979`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer829`:
	fromBuffer829 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer829` to the result `intoRune979`
	// (`intoRune979` is now tainted).
	intoRune979, _, _ := fromBuffer829.ReadRune()

	// Sink the tainted `intoRune979`:
	sink(intoRune979)
}

func TaintStepTest_BytesBufferReadString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBuffer709` into `intoString985`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer709`:
	fromBuffer709 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer709` to the result `intoString985`
	// (`intoString985` is now tainted).
	intoString985, _ := fromBuffer709.ReadString(0)

	// Sink the tainted `intoString985`:
	sink(intoString985)
}

func TaintStepTest_BytesBufferString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBuffer809` into `intoString177`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer809`:
	fromBuffer809 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer809` to the result `intoString177`
	// (`intoString177` is now tainted).
	intoString177 := fromBuffer809.String()

	// Sink the tainted `intoString177`:
	sink(intoString177)
}

func TaintStepTest_BytesBufferWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte407` into `intoBuffer300`.

	// Assume that `sourceCQL` has the underlying type of `fromByte407`:
	fromByte407 := sourceCQL.([]byte)

	// Declare `intoBuffer300` variable:
	var intoBuffer300 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromByte407` to the receiver `intoBuffer300`
	// (`intoBuffer300` is now tainted).
	intoBuffer300.Write(fromByte407)

	// Sink the tainted `intoBuffer300`:
	sink(intoBuffer300)
}

func TaintStepTest_BytesBufferWriteByte_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte847` into `intoBuffer798`.

	// Assume that `sourceCQL` has the underlying type of `fromByte847`:
	fromByte847 := sourceCQL.(byte)

	// Declare `intoBuffer798` variable:
	var intoBuffer798 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromByte847` to the receiver `intoBuffer798`
	// (`intoBuffer798` is now tainted).
	intoBuffer798.WriteByte(fromByte847)

	// Sink the tainted `intoBuffer798`:
	sink(intoBuffer798)
}

func TaintStepTest_BytesBufferWriteRune_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune352` into `intoBuffer385`.

	// Assume that `sourceCQL` has the underlying type of `fromRune352`:
	fromRune352 := sourceCQL.(rune)

	// Declare `intoBuffer385` variable:
	var intoBuffer385 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromRune352` to the receiver `intoBuffer385`
	// (`intoBuffer385` is now tainted).
	intoBuffer385.WriteRune(fromRune352)

	// Sink the tainted `intoBuffer385`:
	sink(intoBuffer385)
}

func TaintStepTest_BytesBufferWriteString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString386` into `intoBuffer760`.

	// Assume that `sourceCQL` has the underlying type of `fromString386`:
	fromString386 := sourceCQL.(string)

	// Declare `intoBuffer760` variable:
	var intoBuffer760 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromString386` to the receiver `intoBuffer760`
	// (`intoBuffer760` is now tainted).
	intoBuffer760.WriteString(fromString386)

	// Sink the tainted `intoBuffer760`:
	sink(intoBuffer760)
}

func TaintStepTest_BytesBufferWriteTo_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBuffer380` into `intoWriter290`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer380`:
	fromBuffer380 := sourceCQL.(bytes.Buffer)

	// Declare `intoWriter290` variable:
	var intoWriter290 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromBuffer380` to the argument `intoWriter290`
	// (`intoWriter290` is now tainted).
	fromBuffer380.WriteTo(intoWriter290)

	// Sink the tainted `intoWriter290`:
	sink(intoWriter290)
}

func TaintStepTest_BytesReaderRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader256` into `intoByte207`.

	// Assume that `sourceCQL` has the underlying type of `fromReader256`:
	fromReader256 := sourceCQL.(bytes.Reader)

	// Declare `intoByte207` variable:
	var intoByte207 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader256` to the argument `intoByte207`
	// (`intoByte207` is now tainted).
	fromReader256.Read(intoByte207)

	// Sink the tainted `intoByte207`:
	sink(intoByte207)
}

func TaintStepTest_BytesReaderReadAt_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader770` into `intoByte175`.

	// Assume that `sourceCQL` has the underlying type of `fromReader770`:
	fromReader770 := sourceCQL.(bytes.Reader)

	// Declare `intoByte175` variable:
	var intoByte175 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader770` to the argument `intoByte175`
	// (`intoByte175` is now tainted).
	fromReader770.ReadAt(intoByte175, 0)

	// Sink the tainted `intoByte175`:
	sink(intoByte175)
}

func TaintStepTest_BytesReaderReadByte_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader474` into `intoByte696`.

	// Assume that `sourceCQL` has the underlying type of `fromReader474`:
	fromReader474 := sourceCQL.(bytes.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader474` to the result `intoByte696`
	// (`intoByte696` is now tainted).
	intoByte696, _ := fromReader474.ReadByte()

	// Sink the tainted `intoByte696`:
	sink(intoByte696)
}

func TaintStepTest_BytesReaderReadRune_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader662` into `intoRune395`.

	// Assume that `sourceCQL` has the underlying type of `fromReader662`:
	fromReader662 := sourceCQL.(bytes.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader662` to the result `intoRune395`
	// (`intoRune395` is now tainted).
	intoRune395, _, _ := fromReader662.ReadRune()

	// Sink the tainted `intoRune395`:
	sink(intoRune395)
}

func TaintStepTest_BytesReaderReset_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte522` into `intoReader483`.

	// Assume that `sourceCQL` has the underlying type of `fromByte522`:
	fromByte522 := sourceCQL.([]byte)

	// Declare `intoReader483` variable:
	var intoReader483 bytes.Reader

	// Call the method that transfers the taint
	// from the parameter `fromByte522` to the receiver `intoReader483`
	// (`intoReader483` is now tainted).
	intoReader483.Reset(fromByte522)

	// Sink the tainted `intoReader483`:
	sink(intoReader483)
}

func TaintStepTest_BytesReaderWriteTo_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader884` into `intoWriter616`.

	// Assume that `sourceCQL` has the underlying type of `fromReader884`:
	fromReader884 := sourceCQL.(bytes.Reader)

	// Declare `intoWriter616` variable:
	var intoWriter616 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromReader884` to the argument `intoWriter616`
	// (`intoWriter616` is now tainted).
	fromReader884.WriteTo(intoWriter616)

	// Sink the tainted `intoWriter616`:
	sink(intoWriter616)
}

func RunAllTaints_Bytes() {
	{
		source := newSource()
		TaintStepTest_BytesFields_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesFieldsFunc_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesJoin_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesJoin_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesMap_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesNewBuffer_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesNewBufferString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesNewReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesRepeat_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReplace_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReplace_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReplaceAll_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReplaceAll_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesRunes_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesSplit_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesSplitAfter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesSplitAfterN_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesSplitN_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTitle_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesToLower_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesToLowerSpecial_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesToTitle_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesToTitleSpecial_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesToUpper_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesToUpperSpecial_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesToValidUTF8_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesToValidUTF8_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrim_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrimFunc_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrimLeft_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrimLeftFunc_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrimPrefix_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrimRight_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrimRightFunc_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrimSpace_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrimSuffix_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferBytes_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferNext_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferReadByte_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferReadBytes_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferReadFrom_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferReadRune_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferReadString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferWrite_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferWriteByte_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferWriteRune_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferWriteString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferWriteTo_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReaderRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReaderReadAt_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReaderReadByte_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReaderReadRune_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReaderReset_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReaderWriteTo_B0I0O0(source)
	}
}
