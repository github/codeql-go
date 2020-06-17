package main

import (
	"io"
	"text/scanner"
)

func TaintStepTest_TextScannerTokenString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRune820` into `intoString981`.

	// Assume that `sourceCQL` has the underlying type of `fromRune820`:
	fromRune820 := sourceCQL.(rune)

	// Call the function that transfers the taint
	// from the parameter `fromRune820` to result `intoString981`
	// (`intoString981` is now tainted).
	intoString981 := scanner.TokenString(fromRune820)

	// Return the tainted `intoString981`:
	return intoString981
}

func TaintStepTest_TextScannerScannerInit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader385` into `intoScanner417`.

	// Assume that `sourceCQL` has the underlying type of `fromReader385`:
	fromReader385 := sourceCQL.(io.Reader)

	// Declare `intoScanner417` variable:
	var intoScanner417 scanner.Scanner

	// Call the method that transfers the taint
	// from the parameter `fromReader385` to the receiver `intoScanner417`
	// (`intoScanner417` is now tainted).
	intoScanner417.Init(fromReader385)

	// Return the tainted `intoScanner417`:
	return intoScanner417
}

func TaintStepTest_TextScannerScannerInit_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader475` into `intoScanner466`.

	// Assume that `sourceCQL` has the underlying type of `fromReader475`:
	fromReader475 := sourceCQL.(io.Reader)

	// Declare medium object/interface:
	var mediumObjCQL scanner.Scanner

	// Call the method that transfers the taint
	// from the parameter `fromReader475` to the result `intoScanner466`
	// (`intoScanner466` is now tainted).
	intoScanner466 := mediumObjCQL.Init(fromReader475)

	// Return the tainted `intoScanner466`:
	return intoScanner466
}

func TaintStepTest_TextScannerScannerNext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner304` into `intoRune640`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner304`:
	fromScanner304 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner304` to the result `intoRune640`
	// (`intoRune640` is now tainted).
	intoRune640 := fromScanner304.Next()

	// Return the tainted `intoRune640`:
	return intoRune640
}

func TaintStepTest_TextScannerScannerPeek_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner570` into `intoRune652`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner570`:
	fromScanner570 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner570` to the result `intoRune652`
	// (`intoRune652` is now tainted).
	intoRune652 := fromScanner570.Peek()

	// Return the tainted `intoRune652`:
	return intoRune652
}

func TaintStepTest_TextScannerScannerScan_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner711` into `intoRune935`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner711`:
	fromScanner711 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner711` to the result `intoRune935`
	// (`intoRune935` is now tainted).
	intoRune935 := fromScanner711.Scan()

	// Return the tainted `intoRune935`:
	return intoRune935
}

func TaintStepTest_TextScannerScannerTokenText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanner791` into `intoString804`.

	// Assume that `sourceCQL` has the underlying type of `fromScanner791`:
	fromScanner791 := sourceCQL.(scanner.Scanner)

	// Call the method that transfers the taint
	// from the receiver `fromScanner791` to the result `intoString804`
	// (`intoString804` is now tainted).
	intoString804 := fromScanner791.TokenText()

	// Return the tainted `intoString804`:
	return intoString804
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
