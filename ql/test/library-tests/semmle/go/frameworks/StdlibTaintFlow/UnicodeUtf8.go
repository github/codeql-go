package main

import "unicode/utf8"

func TaintStepTest_UnicodeUtf8DecodeLastRune(sourceCQL interface{}) {
	// The flow is from `p` into `r`.

	// Assume that `sourceCQL` has the underlying type of `p`:
	p := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `p` to result `r`
	// (`r` is now tainted).
	r, _ := utf8.DecodeLastRune(p)

	// Sink the tainted `r`:
	sink(r)
}

func TaintStepTest_UnicodeUtf8DecodeLastRuneInString(sourceCQL interface{}) {
	// The flow is from `s` into `r`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `r`
	// (`r` is now tainted).
	r, _ := utf8.DecodeLastRuneInString(s)

	// Sink the tainted `r`:
	sink(r)
}

func TaintStepTest_UnicodeUtf8DecodeRune(sourceCQL interface{}) {
	// The flow is from `p` into `r`.

	// Assume that `sourceCQL` has the underlying type of `p`:
	p := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `p` to result `r`
	// (`r` is now tainted).
	r, _ := utf8.DecodeRune(p)

	// Sink the tainted `r`:
	sink(r)
}

func TaintStepTest_UnicodeUtf8DecodeRuneInString(sourceCQL interface{}) {
	// The flow is from `s` into `r`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `r`
	// (`r` is now tainted).
	r, _ := utf8.DecodeRuneInString(s)

	// Sink the tainted `r`:
	sink(r)
}

func TaintStepTest_UnicodeUtf8EncodeRune(sourceCQL interface{}) {
	// The flow is from `r` into `p`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(rune)

	// Declare `p` variable:
	var p []byte

	// Call the function that transfers the taint
	// from the parameter `r` to parameter `p`;
	// `p` is now tainted.
	utf8.EncodeRune(p, r)

	// Sink the tainted `p`:
	sink(p)
}

func RunAllTaints_UnicodeUtf8(v interface{}) {
	{
		source := newSource()
		TaintStepTest_UnicodeUtf8DecodeLastRune(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeUtf8DecodeLastRuneInString(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeUtf8DecodeRune(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeUtf8DecodeRuneInString(source)
	}
	{
		source := newSource()
		TaintStepTest_UnicodeUtf8EncodeRune(source)
	}
}
