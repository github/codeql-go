// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"fmt"
	"io"
)

func TaintStepTest_FmtErrorf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString656` into `intoError414`.

	// Assume that `sourceCQL` has the underlying type of `fromString656`:
	fromString656 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString656` to result `intoError414`
	// (`intoError414` is now tainted).
	intoError414 := fmt.Errorf(fromString656, nil)

	// Return the tainted `intoError414`:
	return intoError414
}

func TaintStepTest_FmtErrorf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface518` into `intoError650`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface518`:
	fromInterface518 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface518` to result `intoError650`
	// (`intoError650` is now tainted).
	intoError650 := fmt.Errorf("", fromInterface518)

	// Return the tainted `intoError650`:
	return intoError650
}

func TaintStepTest_FmtFprint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface784` into `intoWriter957`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface784`:
	fromInterface784 := sourceCQL.(interface{})

	// Declare `intoWriter957` variable:
	var intoWriter957 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromInterface784` to parameter `intoWriter957`;
	// `intoWriter957` is now tainted.
	fmt.Fprint(intoWriter957, fromInterface784)

	// Return the tainted `intoWriter957`:
	return intoWriter957
}

func TaintStepTest_FmtFprintf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString520` into `intoWriter443`.

	// Assume that `sourceCQL` has the underlying type of `fromString520`:
	fromString520 := sourceCQL.(string)

	// Declare `intoWriter443` variable:
	var intoWriter443 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromString520` to parameter `intoWriter443`;
	// `intoWriter443` is now tainted.
	fmt.Fprintf(intoWriter443, fromString520, nil)

	// Return the tainted `intoWriter443`:
	return intoWriter443
}

func TaintStepTest_FmtFprintf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface127` into `intoWriter483`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface127`:
	fromInterface127 := sourceCQL.(interface{})

	// Declare `intoWriter483` variable:
	var intoWriter483 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromInterface127` to parameter `intoWriter483`;
	// `intoWriter483` is now tainted.
	fmt.Fprintf(intoWriter483, "", fromInterface127)

	// Return the tainted `intoWriter483`:
	return intoWriter483
}

func TaintStepTest_FmtFprintln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface989` into `intoWriter982`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface989`:
	fromInterface989 := sourceCQL.(interface{})

	// Declare `intoWriter982` variable:
	var intoWriter982 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromInterface989` to parameter `intoWriter982`;
	// `intoWriter982` is now tainted.
	fmt.Fprintln(intoWriter982, fromInterface989)

	// Return the tainted `intoWriter982`:
	return intoWriter982
}

func TaintStepTest_FmtFscan_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader417` into `intoInterface584`.

	// Assume that `sourceCQL` has the underlying type of `fromReader417`:
	fromReader417 := sourceCQL.(io.Reader)

	// Declare `intoInterface584` variable:
	var intoInterface584 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromReader417` to parameter `intoInterface584`;
	// `intoInterface584` is now tainted.
	fmt.Fscan(fromReader417, intoInterface584)

	// Return the tainted `intoInterface584`:
	return intoInterface584
}

func TaintStepTest_FmtFscanf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader991` into `intoInterface881`.

	// Assume that `sourceCQL` has the underlying type of `fromReader991`:
	fromReader991 := sourceCQL.(io.Reader)

	// Declare `intoInterface881` variable:
	var intoInterface881 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromReader991` to parameter `intoInterface881`;
	// `intoInterface881` is now tainted.
	fmt.Fscanf(fromReader991, "", intoInterface881)

	// Return the tainted `intoInterface881`:
	return intoInterface881
}

func TaintStepTest_FmtFscanf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString186` into `intoInterface284`.

	// Assume that `sourceCQL` has the underlying type of `fromString186`:
	fromString186 := sourceCQL.(string)

	// Declare `intoInterface284` variable:
	var intoInterface284 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString186` to parameter `intoInterface284`;
	// `intoInterface284` is now tainted.
	fmt.Fscanf(nil, fromString186, intoInterface284)

	// Return the tainted `intoInterface284`:
	return intoInterface284
}

func TaintStepTest_FmtFscanln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader908` into `intoInterface137`.

	// Assume that `sourceCQL` has the underlying type of `fromReader908`:
	fromReader908 := sourceCQL.(io.Reader)

	// Declare `intoInterface137` variable:
	var intoInterface137 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromReader908` to parameter `intoInterface137`;
	// `intoInterface137` is now tainted.
	fmt.Fscanln(fromReader908, intoInterface137)

	// Return the tainted `intoInterface137`:
	return intoInterface137
}

func TaintStepTest_FmtSprint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface494` into `intoString873`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface494`:
	fromInterface494 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface494` to result `intoString873`
	// (`intoString873` is now tainted).
	intoString873 := fmt.Sprint(fromInterface494)

	// Return the tainted `intoString873`:
	return intoString873
}

func TaintStepTest_FmtSprintf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString599` into `intoString409`.

	// Assume that `sourceCQL` has the underlying type of `fromString599`:
	fromString599 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString599` to result `intoString409`
	// (`intoString409` is now tainted).
	intoString409 := fmt.Sprintf(fromString599, nil)

	// Return the tainted `intoString409`:
	return intoString409
}

func TaintStepTest_FmtSprintf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface246` into `intoString898`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface246`:
	fromInterface246 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface246` to result `intoString898`
	// (`intoString898` is now tainted).
	intoString898 := fmt.Sprintf("", fromInterface246)

	// Return the tainted `intoString898`:
	return intoString898
}

func TaintStepTest_FmtSprintln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface598` into `intoString631`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface598`:
	fromInterface598 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface598` to result `intoString631`
	// (`intoString631` is now tainted).
	intoString631 := fmt.Sprintln(fromInterface598)

	// Return the tainted `intoString631`:
	return intoString631
}

func TaintStepTest_FmtSscan_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString165` into `intoInterface150`.

	// Assume that `sourceCQL` has the underlying type of `fromString165`:
	fromString165 := sourceCQL.(string)

	// Declare `intoInterface150` variable:
	var intoInterface150 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString165` to parameter `intoInterface150`;
	// `intoInterface150` is now tainted.
	fmt.Sscan(fromString165, intoInterface150)

	// Return the tainted `intoInterface150`:
	return intoInterface150
}

func TaintStepTest_FmtSscanf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString340` into `intoInterface471`.

	// Assume that `sourceCQL` has the underlying type of `fromString340`:
	fromString340 := sourceCQL.(string)

	// Declare `intoInterface471` variable:
	var intoInterface471 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString340` to parameter `intoInterface471`;
	// `intoInterface471` is now tainted.
	fmt.Sscanf(fromString340, "", intoInterface471)

	// Return the tainted `intoInterface471`:
	return intoInterface471
}

func TaintStepTest_FmtSscanf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString290` into `intoInterface758`.

	// Assume that `sourceCQL` has the underlying type of `fromString290`:
	fromString290 := sourceCQL.(string)

	// Declare `intoInterface758` variable:
	var intoInterface758 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString290` to parameter `intoInterface758`;
	// `intoInterface758` is now tainted.
	fmt.Sscanf("", fromString290, intoInterface758)

	// Return the tainted `intoInterface758`:
	return intoInterface758
}

func TaintStepTest_FmtSscanln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString396` into `intoInterface707`.

	// Assume that `sourceCQL` has the underlying type of `fromString396`:
	fromString396 := sourceCQL.(string)

	// Declare `intoInterface707` variable:
	var intoInterface707 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString396` to parameter `intoInterface707`;
	// `intoInterface707` is now tainted.
	fmt.Sscanln(fromString396, intoInterface707)

	// Return the tainted `intoInterface707`:
	return intoInterface707
}

func TaintStepTest_FmtGoStringerGoString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromGoStringer912` into `intoString718`.

	// Assume that `sourceCQL` has the underlying type of `fromGoStringer912`:
	fromGoStringer912 := sourceCQL.(fmt.GoStringer)

	// Call the method that transfers the taint
	// from the receiver `fromGoStringer912` to the result `intoString718`
	// (`intoString718` is now tainted).
	intoString718 := fromGoStringer912.GoString()

	// Return the tainted `intoString718`:
	return intoString718
}

func TaintStepTest_FmtScanStateRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanState972` into `intoByte633`.

	// Assume that `sourceCQL` has the underlying type of `fromScanState972`:
	fromScanState972 := sourceCQL.(fmt.ScanState)

	// Declare `intoByte633` variable:
	var intoByte633 []byte

	// Call the method that transfers the taint
	// from the receiver `fromScanState972` to the argument `intoByte633`
	// (`intoByte633` is now tainted).
	fromScanState972.Read(intoByte633)

	// Return the tainted `intoByte633`:
	return intoByte633
}

func TaintStepTest_FmtScanStateReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanState316` into `intoRune145`.

	// Assume that `sourceCQL` has the underlying type of `fromScanState316`:
	fromScanState316 := sourceCQL.(fmt.ScanState)

	// Call the method that transfers the taint
	// from the receiver `fromScanState316` to the result `intoRune145`
	// (`intoRune145` is now tainted).
	intoRune145, _, _ := fromScanState316.ReadRune()

	// Return the tainted `intoRune145`:
	return intoRune145
}

func TaintStepTest_FmtStringerString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStringer817` into `intoString474`.

	// Assume that `sourceCQL` has the underlying type of `fromStringer817`:
	fromStringer817 := sourceCQL.(fmt.Stringer)

	// Call the method that transfers the taint
	// from the receiver `fromStringer817` to the result `intoString474`
	// (`intoString474` is now tainted).
	intoString474 := fromStringer817.String()

	// Return the tainted `intoString474`:
	return intoString474
}

func TaintStepTest_FmtScanStateToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanState832` into `intoByte378`.

	// Assume that `sourceCQL` has the underlying type of `fromScanState832`:
	fromScanState832 := sourceCQL.(fmt.ScanState)

	// Call the method that transfers the taint
	// from the receiver `fromScanState832` to the result `intoByte378`
	// (`intoByte378` is now tainted).
	intoByte378, _ := fromScanState832.Token(false, nil)

	// Return the tainted `intoByte378`:
	return intoByte378
}

func TaintStepTest_FmtStateWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte541` into `intoState139`.

	// Assume that `sourceCQL` has the underlying type of `fromByte541`:
	fromByte541 := sourceCQL.([]byte)

	// Declare `intoState139` variable:
	var intoState139 fmt.State

	// Call the method that transfers the taint
	// from the parameter `fromByte541` to the receiver `intoState139`
	// (`intoState139` is now tainted).
	intoState139.Write(fromByte541)

	// Return the tainted `intoState139`:
	return intoState139
}

func RunAllTaints_Fmt() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtErrorf_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtErrorf_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtFprint_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtFprintf_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtFprintf_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtFprintln_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtFscan_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtFscanf_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtFscanf_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtFscanln_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtSprint_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtSprintf_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtSprintf_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtSprintln_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtSscan_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtSscanf_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtSscanf_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtSscanln_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtGoStringerGoString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtScanStateRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtScanStateReadRune_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtStringerString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtScanStateToken_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_FmtStateWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
