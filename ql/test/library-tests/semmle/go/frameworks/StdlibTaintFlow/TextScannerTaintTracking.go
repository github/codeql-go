// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"text/scanner"
)

func TaintStepTest_TextScannerTokenString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune952` into `intoString719`.

	// Assume that `sourceCQL` has the underlying type of `fromRune952`:
	fromRune952 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune952` to result `intoString719`
	// (`intoString719` is now tainted).
	intoString719 := scanner.TokenString(fromRune952)

	// Return the tainted `intoString719`:
	return intoString719
}

func TaintStepTest_TextScannerScannerInit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader807` into `intoScanner503`.

	// Assume that `sourceCQL` has the underlying type of `fromReader807`:
	fromReader807 := sourceCQL.(io.Reader)

	// Declare `intoScanner503` variable:
	var intoScanner503 scanner.Scanner

	// Call the method that transfers the taint
	// from the parameter `fromReader807` to the receiver `intoScanner503`
	// (`intoScanner503` is now tainted).
	intoScanner503.Init(fromReader807)

	// Return the tainted `intoScanner503`:
	return intoScanner503
}

func TaintStepTest_TextScannerScannerInit_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader701` into `intoScanner148`.

	// Assume that `sourceCQL` has the underlying type of `fromReader701`:
	fromReader701 := sourceCQL.(io.Reader)

	// Declare medium object/interface:
	var mediumObjCQL scanner.Scanner

	// Call the method that transfers the taint
	// from the parameter `fromReader701` to the result `intoScanner148`
	// (`intoScanner148` is now tainted).
	intoScanner148 := mediumObjCQL.Init(fromReader701)

	// Return the tainted `intoScanner148`:
	return intoScanner148
}

func TaintStepTest_TextScannerScannerNext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner309` into `intoRune492`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner309`:
	fromScanner309 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner309` to the result `intoRune492`
	// (`intoRune492` is now tainted).
	intoRune492 := fromScanner309.Next()

	// Return the tainted `intoRune492`:
	return intoRune492
}

func TaintStepTest_TextScannerScannerPeek_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner758` into `intoRune669`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner758`:
	fromScanner758 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner758` to the result `intoRune669`
	// (`intoRune669` is now tainted).
	intoRune669 := fromScanner758.Peek()

	// Return the tainted `intoRune669`:
	return intoRune669
}

func TaintStepTest_TextScannerScannerScan_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner584` into `intoRune306`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner584`:
	fromScanner584 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner584` to the result `intoRune306`
	// (`intoRune306` is now tainted).
	intoRune306 := fromScanner584.Scan()

	// Return the tainted `intoRune306`:
	return intoRune306
}

func TaintStepTest_TextScannerScannerTokenText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner403` into `intoString963`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner403`:
	fromScanner403 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner403` to the result `intoString963`
	// (`intoString963` is now tainted).
	intoString963 := fromScanner403.TokenText()

	// Return the tainted `intoString963`:
	return intoString963
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
