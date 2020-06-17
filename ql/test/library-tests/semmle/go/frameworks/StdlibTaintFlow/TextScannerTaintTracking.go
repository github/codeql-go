// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"text/scanner"
)

func TaintStepTest_TextScannerTokenString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune656` into `intoString414`.

	// Assume that `sourceCQL` has the underlying type of `fromRune656`:
	fromRune656 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune656` to result `intoString414`
	// (`intoString414` is now tainted).
	intoString414 := scanner.TokenString(fromRune656)

	// Return the tainted `intoString414`:
	return intoString414
}

func TaintStepTest_TextScannerScannerInit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader518` into `intoScanner650`.

	// Assume that `sourceCQL` has the underlying type of `fromReader518`:
	fromReader518 := sourceCQL.(io.Reader)

	// Declare `intoScanner650` variable:
	var intoScanner650 scanner.Scanner

	// Call the method that transfers the taint
	// from the parameter `fromReader518` to the receiver `intoScanner650`
	// (`intoScanner650` is now tainted).
	intoScanner650.Init(fromReader518)

	// Return the tainted `intoScanner650`:
	return intoScanner650
}

func TaintStepTest_TextScannerScannerInit_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader784` into `intoScanner957`.

	// Assume that `sourceCQL` has the underlying type of `fromReader784`:
	fromReader784 := sourceCQL.(io.Reader)

	// Declare medium object/interface:
	var mediumObjCQL scanner.Scanner

	// Call the method that transfers the taint
	// from the parameter `fromReader784` to the result `intoScanner957`
	// (`intoScanner957` is now tainted).
	intoScanner957 := mediumObjCQL.Init(fromReader784)

	// Return the tainted `intoScanner957`:
	return intoScanner957
}

func TaintStepTest_TextScannerScannerNext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner520` into `intoRune443`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner520`:
	fromScanner520 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner520` to the result `intoRune443`
	// (`intoRune443` is now tainted).
	intoRune443 := fromScanner520.Next()

	// Return the tainted `intoRune443`:
	return intoRune443
}

func TaintStepTest_TextScannerScannerPeek_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner127` into `intoRune483`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner127`:
	fromScanner127 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner127` to the result `intoRune483`
	// (`intoRune483` is now tainted).
	intoRune483 := fromScanner127.Peek()

	// Return the tainted `intoRune483`:
	return intoRune483
}

func TaintStepTest_TextScannerScannerScan_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner989` into `intoRune982`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner989`:
	fromScanner989 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner989` to the result `intoRune982`
	// (`intoRune982` is now tainted).
	intoRune982 := fromScanner989.Scan()

	// Return the tainted `intoRune982`:
	return intoRune982
}

func TaintStepTest_TextScannerScannerTokenText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner417` into `intoString584`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner417`:
	fromScanner417 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner417` to the result `intoString584`
	// (`intoString584` is now tainted).
	intoString584 := fromScanner417.TokenText()

	// Return the tainted `intoString584`:
	return intoString584
}

func RunAllTaints_TextScanner() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextScannerTokenString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextScannerScannerInit_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextScannerScannerInit_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextScannerScannerNext_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextScannerScannerPeek_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextScannerScannerScan_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextScannerScannerTokenText_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
