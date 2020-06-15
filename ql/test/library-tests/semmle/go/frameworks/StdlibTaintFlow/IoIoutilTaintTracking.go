package main

import (
	"io"
	"io/ioutil"
)

func TaintStepTest_IoIoutilNopCloser_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader765` into `intoReadCloser911`.

	// Assume that `sourceCQL` has the underlying type of `fromReader765`:
	fromReader765 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader765` to result `intoReadCloser911`
	// (`intoReadCloser911` is now tainted).
	intoReadCloser911 := ioutil.NopCloser(fromReader765)

	// Sink the tainted `intoReadCloser911`:
	sink(intoReadCloser911)
}

func TaintStepTest_IoIoutilReadAll_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader837` into `intoByte198`.

	// Assume that `sourceCQL` has the underlying type of `fromReader837`:
	fromReader837 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader837` to result `intoByte198`
	// (`intoByte198` is now tainted).
	intoByte198, _ := ioutil.ReadAll(fromReader837)

	// Sink the tainted `intoByte198`:
	sink(intoByte198)
}

func RunAllTaints_IoIoutil() {
	{
		source := newSource()
		TaintStepTest_IoIoutilNopCloser_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoIoutilReadAll_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoWriterWrite_B0I0O0(source)
	}
}
