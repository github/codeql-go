package main

import (
	"io"
	"text/scanner"
)

func TaintStepTest_TextScannerTokenString(sourceCQL interface{}) {
	// The flow is from `tok` into `into407`.

	// Assume that `sourceCQL` has the underlying type of `tok`:
	tok := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `tok` to result `into407`
	// (`into407` is now tainted).
	into407 := scanner.TokenString(tok)

	// Sink the tainted `into407`:
	sink(into407)
}

func TaintStepTest_TextScannerScannerInit(sourceCQL interface{}) {
	// The flow is from `src` into `into868`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.(io.Reader)

	// Declare `into868` variable:
	var into868 scanner.Scanner

	// Call the method that transfers the taint
	// from the parameter `src` to the receiver `into868`
	// (`into868` is now tainted).
	into868.Init(src)

	// Sink the tainted `into868`:
	sink(into868)
}

func TaintStepTest_TextScannerScannerNext(sourceCQL interface{}) {
	// The flow is from `from370` into `into796`.

	// Assume that `sourceCQL` has the underlying type of `from370`:
	from370 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `from370` to the result `into796`
	// (`into796` is now tainted).
	into796 := from370.Next()

	// Sink the tainted `into796`:
	sink(into796)
}

func TaintStepTest_TextScannerScannerPeek(sourceCQL interface{}) {
	// The flow is from `from597` into `into876`.

	// Assume that `sourceCQL` has the underlying type of `from597`:
	from597 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `from597` to the result `into876`
	// (`into876` is now tainted).
	into876 := from597.Peek()

	// Sink the tainted `into876`:
	sink(into876)
}

func TaintStepTest_TextScannerScannerScan(sourceCQL interface{}) {
	// The flow is from `from958` into `into514`.

	// Assume that `sourceCQL` has the underlying type of `from958`:
	from958 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `from958` to the result `into514`
	// (`into514` is now tainted).
	into514 := from958.Scan()

	// Sink the tainted `into514`:
	sink(into514)
}

func TaintStepTest_TextScannerScannerTokenText(sourceCQL interface{}) {
	// The flow is from `from936` into `into448`.

	// Assume that `sourceCQL` has the underlying type of `from936`:
	from936 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `from936` to the result `into448`
	// (`into448` is now tainted).
	into448 := from936.TokenText()

	// Sink the tainted `into448`:
	sink(into448)
}

func RunAllTaints_TextScanner(v interface{}) {
	{
		source := newSource()
		TaintStepTest_TextScannerTokenString(source)
	}
	{
		source := newSource()
		TaintStepTest_TextScannerScannerInit(source)
	}
	{
		source := newSource()
		TaintStepTest_TextScannerScannerNext(source)
	}
	{
		source := newSource()
		TaintStepTest_TextScannerScannerPeek(source)
	}
	{
		source := newSource()
		TaintStepTest_TextScannerScannerScan(source)
	}
	{
		source := newSource()
		TaintStepTest_TextScannerScannerTokenText(source)
	}
}
