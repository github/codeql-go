package main

import (
	"encoding/ascii85"
	"io"
)

func TaintStepTest_EncodingAscii85Decode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte209` into `intoByte255`.

	// Assume that `sourceCQL` has the underlying type of `fromByte209`:
	fromByte209 := sourceCQL.([]byte)

	// Declare `intoByte255` variable:
	var intoByte255 []byte

	// Call the function that transfers the taint
	// from the parameter `fromByte209` to parameter `intoByte255`;
	// `intoByte255` is now tainted.
	ascii85.Decode(intoByte255, fromByte209, false)

	// Sink the tainted `intoByte255`:
	sink(intoByte255)
}

func TaintStepTest_EncodingAscii85Encode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte923` into `intoByte725`.

	// Assume that `sourceCQL` has the underlying type of `fromByte923`:
	fromByte923 := sourceCQL.([]byte)

	// Declare `intoByte725` variable:
	var intoByte725 []byte

	// Call the function that transfers the taint
	// from the parameter `fromByte923` to parameter `intoByte725`;
	// `intoByte725` is now tainted.
	ascii85.Encode(intoByte725, fromByte923)

	// Sink the tainted `intoByte725`:
	sink(intoByte725)
}

func TaintStepTest_EncodingAscii85NewDecoder_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader868` into `intoReader353`.

	// Assume that `sourceCQL` has the underlying type of `fromReader868`:
	fromReader868 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader868` to result `intoReader353`
	// (`intoReader353` is now tainted).
	intoReader353 := ascii85.NewDecoder(fromReader868)

	// Sink the tainted `intoReader353`:
	sink(intoReader353)
}

func TaintStepTest_EncodingAscii85NewEncoder_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriteCloser624` into `intoWriter894`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser624`:
	fromWriteCloser624 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter894` variable:
	var intoWriter894 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter894`:
	intermediateCQL := ascii85.NewEncoder(intoWriter894)

	// Extra step (`fromWriteCloser624` taints `intermediateCQL`, which taints `intoWriter894`:
	link(fromWriteCloser624, intermediateCQL)

	// Sink the tainted `intoWriter894`:
	sink(intoWriter894)
}

func RunAllTaints_EncodingAscii85() {
	{
		source := newSource()
		TaintStepTest_EncodingAscii85Decode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAscii85Encode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAscii85NewDecoder_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAscii85NewEncoder_B0I0O0(source)
	}
}
