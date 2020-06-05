package main

import (
	"encoding/ascii85"
	"io"
)

func TaintStepTest_EncodingAscii85Decode(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare `dst` variable:
	var dst []byte

	// Call the function that transfers the taint
	// from the parameter `src` to parameter `dst`;
	// `dst` is now tainted.
	ascii85.Decode(dst, src, false)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_EncodingAscii85Encode(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.([]byte)

	// Declare `dst` variable:
	var dst []byte

	// Call the function that transfers the taint
	// from the parameter `src` to parameter `dst`;
	// `dst` is now tainted.
	ascii85.Encode(dst, src)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_EncodingAscii85NewDecoder(sourceCQL interface{}) {
	// The flow is from `r` into `into497`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into497`
	// (`into497` is now tainted).
	into497 := ascii85.NewDecoder(r)

	// Sink the tainted `into497`:
	sink(into497)
}

func TaintStepTest_EncodingAscii85NewEncoder(sourceCQL interface{}) {
	// The flow is from `into357` into `w`.

	// Assume that `sourceCQL` has the underlying type of `into357`:
	into357 := sourceCQL.(io.WriteCloser)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `w`:
	intermediateCQL := ascii85.NewEncoder(w)

	// Extra step (`into357` taints `intermediateCQL`, which taints `w`:
	link(into357, intermediateCQL)

	// Sink the tainted `w`:
	sink(w)
}

func RunAllTaints_EncodingAscii85(v interface{}) {
	{
		source := newSource()
		TaintStepTest_EncodingAscii85Decode(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAscii85Encode(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAscii85NewDecoder(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAscii85NewEncoder(source)
	}
}
