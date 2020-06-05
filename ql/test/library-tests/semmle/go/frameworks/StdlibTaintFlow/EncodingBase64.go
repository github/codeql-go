package main

import (
	"encoding/base64"
	"io"
)

func TaintStepTest_EncodingBase64NewDecoder(sourceCQL interface{}) {
	// The flow is from `r` into `into790`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into790`
	// (`into790` is now tainted).
	into790 := base64.NewDecoder(nil, r)

	// Sink the tainted `into790`:
	sink(into790)
}

func TaintStepTest_EncodingBase64NewEncoder(sourceCQL interface{}) {
	// The flow is from `from492` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from492`:
	from492 := sourceCQL.(io.WriteCloser)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := base64.NewEncoder(nil, w)

	// Extra step (`from492` taints `intermediateCQL`, which taints `w`:
	link(from492, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_EncodingBase64EncodingDecode(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare `dst` variable:
	var dst []byte

	// Declare medium object/interface:
	var mediumObjCQL base64.Encoding

	// Call the method that transfers the taint
	// from the parameter `src` to the parameter `dst`
	// (`dst` is now tainted).
	mediumObjCQL.Decode(dst, src)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_EncodingBase64EncodingDecodeString(sourceCQL interface{}) {
	// The flow is from `s` into `into424`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL base64.Encoding

	// Call the method that transfers the taint
	// from the parameter `s` to the result `into424`
	// (`into424` is now tainted).
	into424, _ := mediumObjCQL.DecodeString(s)

	// Sink the tainted `into424`:
	sink(into424)
}

func TaintStepTest_EncodingBase64EncodingEncode(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare `dst` variable:
	var dst []byte

	// Declare medium object/interface:
	var mediumObjCQL base64.Encoding

	// Call the method that transfers the taint
	// from the parameter `src` to the parameter `dst`
	// (`dst` is now tainted).
	mediumObjCQL.Encode(dst, src)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_EncodingBase64EncodingEncodeToString(sourceCQL interface{}) {
	// The flow is from `src` into `into314`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL base64.Encoding

	// Call the method that transfers the taint
	// from the parameter `src` to the result `into314`
	// (`into314` is now tainted).
	into314 := mediumObjCQL.EncodeToString(src)

	// Sink the tainted `into314`:
	sink(into314)
}

func RunAllTaints_EncodingBase64(v interface{}) {
	{
		source := newSource()
		TaintStepTest_EncodingBase64NewDecoder(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase64NewEncoder(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase64EncodingDecode(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase64EncodingDecodeString(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase64EncodingEncode(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase64EncodingEncodeToString(source)
	}
}
