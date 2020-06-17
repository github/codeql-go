package main

import (
	"bytes"
	"io"
)

func TaintStepTest_BytesFields_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte665` into `intoByte543`.

	// Assume that `sourceCQL` has the underlying type of `fromByte665`:
	fromByte665 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte665` to result `intoByte543`
	// (`intoByte543` is now tainted).
	intoByte543 := bytes.Fields(fromByte665)

	// Return the tainted `intoByte543`:
	return intoByte543
}

func TaintStepTest_BytesFieldsFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte985` into `intoByte455`.

	// Assume that `sourceCQL` has the underlying type of `fromByte985`:
	fromByte985 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte985` to result `intoByte455`
	// (`intoByte455` is now tainted).
	intoByte455 := bytes.FieldsFunc(fromByte985, nil)

	// Return the tainted `intoByte455`:
	return intoByte455
}

func TaintStepTest_BytesJoin_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte504` into `intoByte349`.

	// Assume that `sourceCQL` has the underlying type of `fromByte504`:
	fromByte504 := sourceCQL.([][]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte504` to result `intoByte349`
	// (`intoByte349` is now tainted).
	intoByte349 := bytes.Join(fromByte504, nil)

	// Return the tainted `intoByte349`:
	return intoByte349
}

func TaintStepTest_BytesJoin_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte324` into `intoByte901`.

	// Assume that `sourceCQL` has the underlying type of `fromByte324`:
	fromByte324 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte324` to result `intoByte901`
	// (`intoByte901` is now tainted).
	intoByte901 := bytes.Join(nil, fromByte324)

	// Return the tainted `intoByte901`:
	return intoByte901
}

func TaintStepTest_BytesMap_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte605` into `intoByte338`.

	// Assume that `sourceCQL` has the underlying type of `fromByte605`:
	fromByte605 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte605` to result `intoByte338`
	// (`intoByte338` is now tainted).
	intoByte338 := bytes.Map(nil, fromByte605)

	// Return the tainted `intoByte338`:
	return intoByte338
}

func TaintStepTest_BytesNewBuffer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte375` into `intoBuffer507`.

	// Assume that `sourceCQL` has the underlying type of `fromByte375`:
	fromByte375 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte375` to result `intoBuffer507`
	// (`intoBuffer507` is now tainted).
	intoBuffer507 := bytes.NewBuffer(fromByte375)

	// Return the tainted `intoBuffer507`:
	return intoBuffer507
}

func TaintStepTest_BytesNewBufferString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString168` into `intoBuffer818`.

	// Assume that `sourceCQL` has the underlying type of `fromString168`:
	fromString168 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString168` to result `intoBuffer818`
	// (`intoBuffer818` is now tainted).
	intoBuffer818 := bytes.NewBufferString(fromString168)

	// Return the tainted `intoBuffer818`:
	return intoBuffer818
}

func TaintStepTest_BytesNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte768` into `intoReader294`.

	// Assume that `sourceCQL` has the underlying type of `fromByte768`:
	fromByte768 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte768` to result `intoReader294`
	// (`intoReader294` is now tainted).
	intoReader294 := bytes.NewReader(fromByte768)

	// Return the tainted `intoReader294`:
	return intoReader294
}

func TaintStepTest_BytesRepeat_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte808` into `intoByte904`.

	// Assume that `sourceCQL` has the underlying type of `fromByte808`:
	fromByte808 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte808` to result `intoByte904`
	// (`intoByte904` is now tainted).
	intoByte904 := bytes.Repeat(fromByte808, 0)

	// Return the tainted `intoByte904`:
	return intoByte904
}

func TaintStepTest_BytesReplace_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte259` into `intoByte279`.

	// Assume that `sourceCQL` has the underlying type of `fromByte259`:
	fromByte259 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte259` to result `intoByte279`
	// (`intoByte279` is now tainted).
	intoByte279 := bytes.Replace(fromByte259, nil, nil, 0)

	// Return the tainted `intoByte279`:
	return intoByte279
}

func TaintStepTest_BytesReplace_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte345` into `intoByte861`.

	// Assume that `sourceCQL` has the underlying type of `fromByte345`:
	fromByte345 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte345` to result `intoByte861`
	// (`intoByte861` is now tainted).
	intoByte861 := bytes.Replace(nil, nil, fromByte345, 0)

	// Return the tainted `intoByte861`:
	return intoByte861
}

func TaintStepTest_BytesReplaceAll_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte410` into `intoByte223`.

	// Assume that `sourceCQL` has the underlying type of `fromByte410`:
	fromByte410 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte410` to result `intoByte223`
	// (`intoByte223` is now tainted).
	intoByte223 := bytes.ReplaceAll(fromByte410, nil, nil)

	// Return the tainted `intoByte223`:
	return intoByte223
}

func TaintStepTest_BytesReplaceAll_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte261` into `intoByte846`.

	// Assume that `sourceCQL` has the underlying type of `fromByte261`:
	fromByte261 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte261` to result `intoByte846`
	// (`intoByte846` is now tainted).
	intoByte846 := bytes.ReplaceAll(nil, nil, fromByte261)

	// Return the tainted `intoByte846`:
	return intoByte846
}

func TaintStepTest_BytesRunes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte749` into `intoRune560`.

	// Assume that `sourceCQL` has the underlying type of `fromByte749`:
	fromByte749 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte749` to result `intoRune560`
	// (`intoRune560` is now tainted).
	intoRune560 := bytes.Runes(fromByte749)

	// Return the tainted `intoRune560`:
	return intoRune560
}

func TaintStepTest_BytesSplit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte274` into `intoByte768`.

	// Assume that `sourceCQL` has the underlying type of `fromByte274`:
	fromByte274 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte274` to result `intoByte768`
	// (`intoByte768` is now tainted).
	intoByte768 := bytes.Split(fromByte274, nil)

	// Return the tainted `intoByte768`:
	return intoByte768
}

func TaintStepTest_BytesSplitAfter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte899` into `intoByte856`.

	// Assume that `sourceCQL` has the underlying type of `fromByte899`:
	fromByte899 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte899` to result `intoByte856`
	// (`intoByte856` is now tainted).
	intoByte856 := bytes.SplitAfter(fromByte899, nil)

	// Return the tainted `intoByte856`:
	return intoByte856
}

func TaintStepTest_BytesSplitAfterN_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte833` into `intoByte763`.

	// Assume that `sourceCQL` has the underlying type of `fromByte833`:
	fromByte833 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte833` to result `intoByte763`
	// (`intoByte763` is now tainted).
	intoByte763 := bytes.SplitAfterN(fromByte833, nil, 0)

	// Return the tainted `intoByte763`:
	return intoByte763
}

func TaintStepTest_BytesSplitN_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte124` into `intoByte538`.

	// Assume that `sourceCQL` has the underlying type of `fromByte124`:
	fromByte124 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte124` to result `intoByte538`
	// (`intoByte538` is now tainted).
	intoByte538 := bytes.SplitN(fromByte124, nil, 0)

	// Return the tainted `intoByte538`:
	return intoByte538
}

func TaintStepTest_BytesTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte805` into `intoByte580`.

	// Assume that `sourceCQL` has the underlying type of `fromByte805`:
	fromByte805 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte805` to result `intoByte580`
	// (`intoByte580` is now tainted).
	intoByte580 := bytes.Title(fromByte805)

	// Return the tainted `intoByte580`:
	return intoByte580
}

func TaintStepTest_BytesToLower_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte223` into `intoByte778`.

	// Assume that `sourceCQL` has the underlying type of `fromByte223`:
	fromByte223 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte223` to result `intoByte778`
	// (`intoByte778` is now tainted).
	intoByte778 := bytes.ToLower(fromByte223)

	// Return the tainted `intoByte778`:
	return intoByte778
}

func TaintStepTest_BytesToLowerSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte763` into `intoByte703`.

	// Assume that `sourceCQL` has the underlying type of `fromByte763`:
	fromByte763 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte763` to result `intoByte703`
	// (`intoByte703` is now tainted).
	intoByte703 := bytes.ToLowerSpecial(nil, fromByte763)

	// Return the tainted `intoByte703`:
	return intoByte703
}

func TaintStepTest_BytesToTitle_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte644` into `intoByte500`.

	// Assume that `sourceCQL` has the underlying type of `fromByte644`:
	fromByte644 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte644` to result `intoByte500`
	// (`intoByte500` is now tainted).
	intoByte500 := bytes.ToTitle(fromByte644)

	// Return the tainted `intoByte500`:
	return intoByte500
}

func TaintStepTest_BytesToTitleSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte708` into `intoByte428`.

	// Assume that `sourceCQL` has the underlying type of `fromByte708`:
	fromByte708 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte708` to result `intoByte428`
	// (`intoByte428` is now tainted).
	intoByte428 := bytes.ToTitleSpecial(nil, fromByte708)

	// Return the tainted `intoByte428`:
	return intoByte428
}

func TaintStepTest_BytesToUpper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte154` into `intoByte861`.

	// Assume that `sourceCQL` has the underlying type of `fromByte154`:
	fromByte154 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte154` to result `intoByte861`
	// (`intoByte861` is now tainted).
	intoByte861 := bytes.ToUpper(fromByte154)

	// Return the tainted `intoByte861`:
	return intoByte861
}

func TaintStepTest_BytesToUpperSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte496` into `intoByte769`.

	// Assume that `sourceCQL` has the underlying type of `fromByte496`:
	fromByte496 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte496` to result `intoByte769`
	// (`intoByte769` is now tainted).
	intoByte769 := bytes.ToUpperSpecial(nil, fromByte496)

	// Return the tainted `intoByte769`:
	return intoByte769
}

func TaintStepTest_BytesToValidUTF8_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte219` into `intoByte232`.

	// Assume that `sourceCQL` has the underlying type of `fromByte219`:
	fromByte219 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte219` to result `intoByte232`
	// (`intoByte232` is now tainted).
	intoByte232 := bytes.ToValidUTF8(fromByte219, nil)

	// Return the tainted `intoByte232`:
	return intoByte232
}

func TaintStepTest_BytesToValidUTF8_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte733` into `intoByte556`.

	// Assume that `sourceCQL` has the underlying type of `fromByte733`:
	fromByte733 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte733` to result `intoByte556`
	// (`intoByte556` is now tainted).
	intoByte556 := bytes.ToValidUTF8(nil, fromByte733)

	// Return the tainted `intoByte556`:
	return intoByte556
}

func TaintStepTest_BytesTrim_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte430` into `intoByte822`.

	// Assume that `sourceCQL` has the underlying type of `fromByte430`:
	fromByte430 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte430` to result `intoByte822`
	// (`intoByte822` is now tainted).
	intoByte822 := bytes.Trim(fromByte430, "")

	// Return the tainted `intoByte822`:
	return intoByte822
}

func TaintStepTest_BytesTrimFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte856` into `intoByte683`.

	// Assume that `sourceCQL` has the underlying type of `fromByte856`:
	fromByte856 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte856` to result `intoByte683`
	// (`intoByte683` is now tainted).
	intoByte683 := bytes.TrimFunc(fromByte856, nil)

	// Return the tainted `intoByte683`:
	return intoByte683
}

func TaintStepTest_BytesTrimLeft_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte212` into `intoByte868`.

	// Assume that `sourceCQL` has the underlying type of `fromByte212`:
	fromByte212 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte212` to result `intoByte868`
	// (`intoByte868` is now tainted).
	intoByte868 := bytes.TrimLeft(fromByte212, "")

	// Return the tainted `intoByte868`:
	return intoByte868
}

func TaintStepTest_BytesTrimLeftFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte249` into `intoByte658`.

	// Assume that `sourceCQL` has the underlying type of `fromByte249`:
	fromByte249 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte249` to result `intoByte658`
	// (`intoByte658` is now tainted).
	intoByte658 := bytes.TrimLeftFunc(fromByte249, nil)

	// Return the tainted `intoByte658`:
	return intoByte658
}

func TaintStepTest_BytesTrimPrefix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte407` into `intoByte329`.

	// Assume that `sourceCQL` has the underlying type of `fromByte407`:
	fromByte407 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte407` to result `intoByte329`
	// (`intoByte329` is now tainted).
	intoByte329 := bytes.TrimPrefix(fromByte407, nil)

	// Return the tainted `intoByte329`:
	return intoByte329
}

func TaintStepTest_BytesTrimRight_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte236` into `intoByte710`.

	// Assume that `sourceCQL` has the underlying type of `fromByte236`:
	fromByte236 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte236` to result `intoByte710`
	// (`intoByte710` is now tainted).
	intoByte710 := bytes.TrimRight(fromByte236, "")

	// Return the tainted `intoByte710`:
	return intoByte710
}

func TaintStepTest_BytesTrimRightFunc_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte454` into `intoByte967`.

	// Assume that `sourceCQL` has the underlying type of `fromByte454`:
	fromByte454 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte454` to result `intoByte967`
	// (`intoByte967` is now tainted).
	intoByte967 := bytes.TrimRightFunc(fromByte454, nil)

	// Return the tainted `intoByte967`:
	return intoByte967
}

func TaintStepTest_BytesTrimSpace_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte713` into `intoByte235`.

	// Assume that `sourceCQL` has the underlying type of `fromByte713`:
	fromByte713 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte713` to result `intoByte235`
	// (`intoByte235` is now tainted).
	intoByte235 := bytes.TrimSpace(fromByte713)

	// Return the tainted `intoByte235`:
	return intoByte235
}

func TaintStepTest_BytesTrimSuffix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte428` into `intoByte928`.

	// Assume that `sourceCQL` has the underlying type of `fromByte428`:
	fromByte428 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte428` to result `intoByte928`
	// (`intoByte928` is now tainted).
	intoByte928 := bytes.TrimSuffix(fromByte428, nil)

	// Return the tainted `intoByte928`:
	return intoByte928
}

func TaintStepTest_BytesBufferBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer856` into `intoByte626`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer856`:
	fromBuffer856 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer856` to the result `intoByte626`
	// (`intoByte626` is now tainted).
	intoByte626 := fromBuffer856.Bytes()

	// Return the tainted `intoByte626`:
	return intoByte626
}

func TaintStepTest_BytesBufferNext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer154` into `intoByte878`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer154`:
	fromBuffer154 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer154` to the result `intoByte878`
	// (`intoByte878` is now tainted).
	intoByte878 := fromBuffer154.Next(0)

	// Return the tainted `intoByte878`:
	return intoByte878
}

func TaintStepTest_BytesBufferRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer447` into `intoByte520`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer447`:
	fromBuffer447 := sourceCQL.(bytes.Buffer)

	// Declare `intoByte520` variable:
	var intoByte520 []byte

	// Call the method that transfers the taint
	// from the receiver `fromBuffer447` to the argument `intoByte520`
	// (`intoByte520` is now tainted).
	fromBuffer447.Read(intoByte520)

	// Return the tainted `intoByte520`:
	return intoByte520
}

func TaintStepTest_BytesBufferReadByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer522` into `intoByte913`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer522`:
	fromBuffer522 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer522` to the result `intoByte913`
	// (`intoByte913` is now tainted).
	intoByte913, _ := fromBuffer522.ReadByte()

	// Return the tainted `intoByte913`:
	return intoByte913
}

func TaintStepTest_BytesBufferReadBytes_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer634` into `intoByte881`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer634`:
	fromBuffer634 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer634` to the result `intoByte881`
	// (`intoByte881` is now tainted).
	intoByte881, _ := fromBuffer634.ReadBytes(0)

	// Return the tainted `intoByte881`:
	return intoByte881
}

func TaintStepTest_BytesBufferReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader976` into `intoBuffer954`.

	// Assume that `sourceCQL` has the underlying type of `fromReader976`:
	fromReader976 := sourceCQL.(io.Reader)

	// Declare `intoBuffer954` variable:
	var intoBuffer954 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromReader976` to the receiver `intoBuffer954`
	// (`intoBuffer954` is now tainted).
	intoBuffer954.ReadFrom(fromReader976)

	// Return the tainted `intoBuffer954`:
	return intoBuffer954
}

func TaintStepTest_BytesBufferReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer353` into `intoRune748`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer353`:
	fromBuffer353 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer353` to the result `intoRune748`
	// (`intoRune748` is now tainted).
	intoRune748, _, _ := fromBuffer353.ReadRune()

	// Return the tainted `intoRune748`:
	return intoRune748
}

func TaintStepTest_BytesBufferReadString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer983` into `intoString812`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer983`:
	fromBuffer983 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer983` to the result `intoString812`
	// (`intoString812` is now tainted).
	intoString812, _ := fromBuffer983.ReadString(0)

	// Return the tainted `intoString812`:
	return intoString812
}

func TaintStepTest_BytesBufferString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer235` into `intoString580`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer235`:
	fromBuffer235 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `fromBuffer235` to the result `intoString580`
	// (`intoString580` is now tainted).
	intoString580 := fromBuffer235.String()

	// Return the tainted `intoString580`:
	return intoString580
}

func TaintStepTest_BytesBufferWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte473` into `intoBuffer668`.

	// Assume that `sourceCQL` has the underlying type of `fromByte473`:
	fromByte473 := sourceCQL.([]byte)

	// Declare `intoBuffer668` variable:
	var intoBuffer668 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromByte473` to the receiver `intoBuffer668`
	// (`intoBuffer668` is now tainted).
	intoBuffer668.Write(fromByte473)

	// Return the tainted `intoBuffer668`:
	return intoBuffer668
}

func TaintStepTest_BytesBufferWriteByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte648` into `intoBuffer725`.

	// Assume that `sourceCQL` has the underlying type of `fromByte648`:
	fromByte648 := sourceCQL.(byte)

	// Declare `intoBuffer725` variable:
	var intoBuffer725 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromByte648` to the receiver `intoBuffer725`
	// (`intoBuffer725` is now tainted).
	intoBuffer725.WriteByte(fromByte648)

	// Return the tainted `intoBuffer725`:
	return intoBuffer725
}

func TaintStepTest_BytesBufferWriteRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune994` into `intoBuffer667`.

	// Assume that `sourceCQL` has the underlying type of `fromRune994`:
	fromRune994 := sourceCQL.(rune)

	// Declare `intoBuffer667` variable:
	var intoBuffer667 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromRune994` to the receiver `intoBuffer667`
	// (`intoBuffer667` is now tainted).
	intoBuffer667.WriteRune(fromRune994)

	// Return the tainted `intoBuffer667`:
	return intoBuffer667
}

func TaintStepTest_BytesBufferWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString903` into `intoBuffer399`.

	// Assume that `sourceCQL` has the underlying type of `fromString903`:
	fromString903 := sourceCQL.(string)

	// Declare `intoBuffer399` variable:
	var intoBuffer399 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `fromString903` to the receiver `intoBuffer399`
	// (`intoBuffer399` is now tainted).
	intoBuffer399.WriteString(fromString903)

	// Return the tainted `intoBuffer399`:
	return intoBuffer399
}

func TaintStepTest_BytesBufferWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffer932` into `intoWriter362`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffer932`:
	fromBuffer932 := sourceCQL.(bytes.Buffer)

	// Declare `intoWriter362` variable:
	var intoWriter362 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromBuffer932` to the argument `intoWriter362`
	// (`intoWriter362` is now tainted).
	fromBuffer932.WriteTo(intoWriter362)

	// Return the tainted `intoWriter362`:
	return intoWriter362
}

func TaintStepTest_BytesReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader620` into `intoByte530`.

	// Assume that `sourceCQL` has the underlying type of `fromReader620`:
	fromReader620 := sourceCQL.(bytes.Reader)

	// Declare `intoByte530` variable:
	var intoByte530 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader620` to the argument `intoByte530`
	// (`intoByte530` is now tainted).
	fromReader620.Read(intoByte530)

	// Return the tainted `intoByte530`:
	return intoByte530
}

func TaintStepTest_BytesReaderReadAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader202` into `intoByte735`.

	// Assume that `sourceCQL` has the underlying type of `fromReader202`:
	fromReader202 := sourceCQL.(bytes.Reader)

	// Declare `intoByte735` variable:
	var intoByte735 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader202` to the argument `intoByte735`
	// (`intoByte735` is now tainted).
	fromReader202.ReadAt(intoByte735, 0)

	// Return the tainted `intoByte735`:
	return intoByte735
}

func TaintStepTest_BytesReaderReadByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader920` into `intoByte215`.

	// Assume that `sourceCQL` has the underlying type of `fromReader920`:
	fromReader920 := sourceCQL.(bytes.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader920` to the result `intoByte215`
	// (`intoByte215` is now tainted).
	intoByte215, _ := fromReader920.ReadByte()

	// Return the tainted `intoByte215`:
	return intoByte215
}

func TaintStepTest_BytesReaderReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader530` into `intoRune985`.

	// Assume that `sourceCQL` has the underlying type of `fromReader530`:
	fromReader530 := sourceCQL.(bytes.Reader)

	// Call the method that transfers the taint
	// from the receiver `fromReader530` to the result `intoRune985`
	// (`intoRune985` is now tainted).
	intoRune985, _, _ := fromReader530.ReadRune()

	// Return the tainted `intoRune985`:
	return intoRune985
}

func TaintStepTest_BytesReaderReset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte299` into `intoReader777`.

	// Assume that `sourceCQL` has the underlying type of `fromByte299`:
	fromByte299 := sourceCQL.([]byte)

	// Declare `intoReader777` variable:
	var intoReader777 bytes.Reader

	// Call the method that transfers the taint
	// from the parameter `fromByte299` to the receiver `intoReader777`
	// (`intoReader777` is now tainted).
	intoReader777.Reset(fromByte299)

	// Return the tainted `intoReader777`:
	return intoReader777
}

func TaintStepTest_BytesReaderWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader159` into `intoWriter474`.

	// Assume that `sourceCQL` has the underlying type of `fromReader159`:
	fromReader159 := sourceCQL.(bytes.Reader)

	// Declare `intoWriter474` variable:
	var intoWriter474 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromReader159` to the argument `intoWriter474`
	// (`intoWriter474` is now tainted).
	fromReader159.WriteTo(intoWriter474)

	// Return the tainted `intoWriter474`:
	return intoWriter474
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
