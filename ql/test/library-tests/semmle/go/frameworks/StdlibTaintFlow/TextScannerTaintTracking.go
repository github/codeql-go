package main

import (
	"io"
	"text/scanner"
)

func TaintStepTest_TextScannerTokenString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRune591` into `intoString844`.

	// Assume that `sourceCQL` has the underlying type of `fromRune591`:
	fromRune591 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune591` to result `intoString844`
	// (`intoString844` is now tainted).
	intoString844 := scanner.TokenString(fromRune591)

	// Sink the tainted `intoString844`:
	sink(intoString844)
}

func TaintStepTest_TextScannerScannerInit_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader179` into `intoScanner990`.

	// Assume that `sourceCQL` has the underlying type of `fromReader179`:
	fromReader179 := sourceCQL.(io.Reader)

	// Declare `intoScanner990` variable:
	var intoScanner990 scanner.Scanner

	// Call the method that transfers the taint
	// from the parameter `fromReader179` to the receiver `intoScanner990`
	// (`intoScanner990` is now tainted).
	intoScanner990.Init(fromReader179)

	// Sink the tainted `intoScanner990`:
	sink(intoScanner990)
}

func TaintStepTest_TextScannerScannerInit_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromReader493` into `intoScanner818`.

	// Assume that `sourceCQL` has the underlying type of `fromReader493`:
	fromReader493 := sourceCQL.(io.Reader)

	// Declare medium object/interface:
	var mediumObjCQL scanner.Scanner

	// Call the method that transfers the taint
	// from the parameter `fromReader493` to the result `intoScanner818`
	// (`intoScanner818` is now tainted).
	intoScanner818 := mediumObjCQL.Init(fromReader493)

	// Sink the tainted `intoScanner818`:
	sink(intoScanner818)
}

func TaintStepTest_TextScannerScannerNext_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromScanner962` into `intoRune672`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner962`:
	fromScanner962 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner962` to the result `intoRune672`
	// (`intoRune672` is now tainted).
	intoRune672 := fromScanner962.Next()

	// Sink the tainted `intoRune672`:
	sink(intoRune672)
}

func TaintStepTest_TextScannerScannerPeek_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromScanner818` into `intoRune182`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner818`:
	fromScanner818 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner818` to the result `intoRune182`
	// (`intoRune182` is now tainted).
	intoRune182 := fromScanner818.Peek()

	// Sink the tainted `intoRune182`:
	sink(intoRune182)
}

func TaintStepTest_TextScannerScannerScan_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromScanner211` into `intoRune255`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner211`:
	fromScanner211 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner211` to the result `intoRune255`
	// (`intoRune255` is now tainted).
	intoRune255 := fromScanner211.Scan()

	// Sink the tainted `intoRune255`:
	sink(intoRune255)
}

func TaintStepTest_TextScannerScannerTokenText_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromScanner171` into `intoString659`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner171`:
	fromScanner171 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner171` to the result `intoString659`
	// (`intoString659` is now tainted).
	intoString659 := fromScanner171.TokenText()

	// Sink the tainted `intoString659`:
	sink(intoString659)
}

func RunAllTaints_TextScanner() {
	{
		source := newSource()
		TaintStepTest_TextScannerTokenString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextScannerScannerInit_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextScannerScannerInit_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_TextScannerScannerNext_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextScannerScannerPeek_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextScannerScannerScan_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextScannerScannerTokenText_B0I0O0(source)
	}
}
