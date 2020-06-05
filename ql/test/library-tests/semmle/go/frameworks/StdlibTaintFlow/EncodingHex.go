package main

import (
	"encoding/hex"
	"io"
)

func TaintStepTest_EncodingHexDecode(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare `dst` variable:
	var dst []byte

	// Call the function that transfers the taint
	// from the parameter `src` to parameter `dst`;
	// `dst` is now tainted.
	hex.Decode(dst, src)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_EncodingHexDecodeString(sourceCQL interface{}) {
	// The flow is from `s` into `into653`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into653`
	// (`into653` is now tainted).
	into653, _ := hex.DecodeString(s)

	// Sink the tainted `into653`:
	sink(into653)
}

func TaintStepTest_EncodingHexDump(sourceCQL interface{}) {
	// The flow is from `data` into `into603`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `data` to result `into603`
	// (`into603` is now tainted).
	into603 := hex.Dump(data)

	// Sink the tainted `into603`:
	sink(into603)
}

func TaintStepTest_EncodingHexDumper(sourceCQL interface{}) {
	// The flow is from `from934` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from934`:
	from934 := sourceCQL.(io.WriteCloser)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := hex.Dumper(w)

	// Extra step (`from934` taints `intermediateCQL`, which taints `w`:
	link(from934, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_EncodingHexEncode(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare `dst` variable:
	var dst []byte

	// Call the function that transfers the taint
	// from the parameter `src` to parameter `dst`;
	// `dst` is now tainted.
	hex.Encode(dst, src)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_EncodingHexEncodeToString(sourceCQL interface{}) {
	// The flow is from `src` into `into644`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `src` to result `into644`
	// (`into644` is now tainted).
	into644 := hex.EncodeToString(src)

	// Sink the tainted `into644`:
	sink(into644)
}

func TaintStepTest_EncodingHexNewDecoder(sourceCQL interface{}) {
	// The flow is from `r` into `into435`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into435`
	// (`into435` is now tainted).
	into435 := hex.NewDecoder(r)

	// Sink the tainted `into435`:
	sink(into435)
}

func TaintStepTest_EncodingHexNewEncoder(sourceCQL interface{}) {
	// The flow is from `from696` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from696`:
	from696 := sourceCQL.(io.Writer)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := hex.NewEncoder(w)

	// Extra step (`from696` taints `intermediateCQL`, which taints `w`:
	link(from696, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func RunAllTaints_EncodingHex(v interface{}) {
	{
		source := newSource()
		TaintStepTest_EncodingHexDecode(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingHexDecodeString(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingHexDump(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingHexDumper(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingHexEncode(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingHexEncodeToString(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingHexNewDecoder(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingHexNewEncoder(source)
	}
}
