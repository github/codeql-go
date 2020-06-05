package main

import (
	"encoding/base32"
	"io"
)

func TaintStepTest_EncodingBase32NewDecoder(sourceCQL interface{}) {
	// The flow is from `r` into `into537`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into537`
	// (`into537` is now tainted).
	into537 := base32.NewDecoder(nil, r)

	// Sink the tainted `into537`:
	sink(into537)
}

func TaintStepTest_EncodingBase32NewEncoder(sourceCQL interface{}) {
	// The flow is from `from945` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from945`:
	from945 := sourceCQL.(io.WriteCloser)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := base32.NewEncoder(nil, w)

	// Extra step (`from945` taints `intermediateCQL`, which taints `w`:
	link(from945, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_EncodingBase32EncodingDecode(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare `dst` variable:
	var dst []byte

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `src` to the parameter `dst`
	// (`dst` is now tainted).
	mediumObjCQL.Decode(dst, src)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_EncodingBase32EncodingDecodeString(sourceCQL interface{}) {
	// The flow is from `s` into `into130`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `s` to the result `into130`
	// (`into130` is now tainted).
	into130, _ := mediumObjCQL.DecodeString(s)

	// Sink the tainted `into130`:
	sink(into130)
}

func TaintStepTest_EncodingBase32EncodingEncode(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare `dst` variable:
	var dst []byte

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `src` to the parameter `dst`
	// (`dst` is now tainted).
	mediumObjCQL.Encode(dst, src)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_EncodingBase32EncodingEncodeToString(sourceCQL interface{}) {
	// The flow is from `src` into `into972`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `src` to the result `into972`
	// (`into972` is now tainted).
	into972 := mediumObjCQL.EncodeToString(src)

	// Sink the tainted `into972`:
	sink(into972)
}

func RunAllTaints_EncodingBase32(v interface{}) {
	{
		source := newSource()
		TaintStepTest_EncodingBase32NewDecoder(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase32NewEncoder(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase32EncodingDecode(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase32EncodingDecodeString(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase32EncodingEncode(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingBase32EncodingEncodeToString(source)
	}
}
