package main

import (
	"encoding/pem"
	"io"
)

func TaintStepTest_EncodingPemDecode(sourceCQL interface{}) {
	// The flow is from `data` into `p`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `data` to result `p`
	// (`p` is now tainted).
	p, _ := pem.Decode(data)

	// Sink the tainted `p`:
	sink(p)
}

func TaintStepTest_EncodingPemEncode(sourceCQL interface{}) {
	// The flow is from `b` into `out`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.(*pem.Block)

	// Declare `out` variable:
	var out io.Writer

	// Call the function that transfers the taint
	// from the parameter `b` to parameter `out`;
	// `out` is now tainted.
	pem.Encode(out, b)

	// Sink the tainted `out`:
	sink(out)
}

func TaintStepTest_EncodingPemEncodeToMemory(sourceCQL interface{}) {
	// The flow is from `b` into `into478`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.(*pem.Block)

	// Call the function that transfers the taint
	// from the parameter `b` to result `into478`
	// (`into478` is now tainted).
	into478 := pem.EncodeToMemory(b)

	// Sink the tainted `into478`:
	sink(into478)
}

func RunAllTaints_EncodingPem(v interface{}) {
	{
		source := newSource()
		TaintStepTest_EncodingPemDecode(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingPemEncode(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingPemEncodeToMemory(source)
	}
}
