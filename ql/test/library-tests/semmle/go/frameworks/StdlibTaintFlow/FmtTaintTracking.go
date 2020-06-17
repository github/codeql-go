// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"fmt"
	"io"
)

func TaintStepTest_FmtErrorf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString472` into `intoError196`.

	// Assume that `sourceCQL` has the underlying type of `fromString472`:
	fromString472 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString472` to result `intoError196`
	// (`intoError196` is now tainted).
	intoError196 := fmt.Errorf(fromString472, nil)

	// Return the tainted `intoError196`:
	return intoError196
}

func TaintStepTest_FmtErrorf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface736` into `intoError872`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface736`:
	fromInterface736 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface736` to result `intoError872`
	// (`intoError872` is now tainted).
	intoError872 := fmt.Errorf("", fromInterface736)

	// Return the tainted `intoError872`:
	return intoError872
}

func TaintStepTest_FmtFprint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface504` into `intoWriter389`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface504`:
	fromInterface504 := sourceCQL.(interface{})

	// Declare `intoWriter389` variable:
	var intoWriter389 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromInterface504` to parameter `intoWriter389`;
	// `intoWriter389` is now tainted.
	fmt.Fprint(intoWriter389, fromInterface504)

	// Return the tainted `intoWriter389`:
	return intoWriter389
}

func TaintStepTest_FmtFprintf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString825` into `intoWriter759`.

	// Assume that `sourceCQL` has the underlying type of `fromString825`:
	fromString825 := sourceCQL.(string)

	// Declare `intoWriter759` variable:
	var intoWriter759 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromString825` to parameter `intoWriter759`;
	// `intoWriter759` is now tainted.
	fmt.Fprintf(intoWriter759, fromString825, nil)

	// Return the tainted `intoWriter759`:
	return intoWriter759
}

func TaintStepTest_FmtFprintf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface720` into `intoWriter562`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface720`:
	fromInterface720 := sourceCQL.(interface{})

	// Declare `intoWriter562` variable:
	var intoWriter562 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromInterface720` to parameter `intoWriter562`;
	// `intoWriter562` is now tainted.
	fmt.Fprintf(intoWriter562, "", fromInterface720)

	// Return the tainted `intoWriter562`:
	return intoWriter562
}

func TaintStepTest_FmtFprintln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface137` into `intoWriter112`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface137`:
	fromInterface137 := sourceCQL.(interface{})

	// Declare `intoWriter112` variable:
	var intoWriter112 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromInterface137` to parameter `intoWriter112`;
	// `intoWriter112` is now tainted.
	fmt.Fprintln(intoWriter112, fromInterface137)

	// Return the tainted `intoWriter112`:
	return intoWriter112
}

func TaintStepTest_FmtFscan_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader838` into `intoInterface974`.

	// Assume that `sourceCQL` has the underlying type of `fromReader838`:
	fromReader838 := sourceCQL.(io.Reader)

	// Declare `intoInterface974` variable:
	var intoInterface974 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromReader838` to parameter `intoInterface974`;
	// `intoInterface974` is now tainted.
	fmt.Fscan(fromReader838, intoInterface974)

	// Return the tainted `intoInterface974`:
	return intoInterface974
}

func TaintStepTest_FmtFscanf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader941` into `intoInterface314`.

	// Assume that `sourceCQL` has the underlying type of `fromReader941`:
	fromReader941 := sourceCQL.(io.Reader)

	// Declare `intoInterface314` variable:
	var intoInterface314 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromReader941` to parameter `intoInterface314`;
	// `intoInterface314` is now tainted.
	fmt.Fscanf(fromReader941, "", intoInterface314)

	// Return the tainted `intoInterface314`:
	return intoInterface314
}

func TaintStepTest_FmtFscanf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString961` into `intoInterface483`.

	// Assume that `sourceCQL` has the underlying type of `fromString961`:
	fromString961 := sourceCQL.(string)

	// Declare `intoInterface483` variable:
	var intoInterface483 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString961` to parameter `intoInterface483`;
	// `intoInterface483` is now tainted.
	fmt.Fscanf(nil, fromString961, intoInterface483)

	// Return the tainted `intoInterface483`:
	return intoInterface483
}

func TaintStepTest_FmtFscanln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader664` into `intoInterface230`.

	// Assume that `sourceCQL` has the underlying type of `fromReader664`:
	fromReader664 := sourceCQL.(io.Reader)

	// Declare `intoInterface230` variable:
	var intoInterface230 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromReader664` to parameter `intoInterface230`;
	// `intoInterface230` is now tainted.
	fmt.Fscanln(fromReader664, intoInterface230)

	// Return the tainted `intoInterface230`:
	return intoInterface230
}

func TaintStepTest_FmtSprint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface605` into `intoString538`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface605`:
	fromInterface605 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface605` to result `intoString538`
	// (`intoString538` is now tainted).
	intoString538 := fmt.Sprint(fromInterface605)

	// Return the tainted `intoString538`:
	return intoString538
}

func TaintStepTest_FmtSprintf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString854` into `intoString580`.

	// Assume that `sourceCQL` has the underlying type of `fromString854`:
	fromString854 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString854` to result `intoString580`
	// (`intoString580` is now tainted).
	intoString580 := fmt.Sprintf(fromString854, nil)

	// Return the tainted `intoString580`:
	return intoString580
}

func TaintStepTest_FmtSprintf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface414` into `intoString346`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface414`:
	fromInterface414 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface414` to result `intoString346`
	// (`intoString346` is now tainted).
	intoString346 := fmt.Sprintf("", fromInterface414)

	// Return the tainted `intoString346`:
	return intoString346
}

func TaintStepTest_FmtSprintln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface184` into `intoString761`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface184`:
	fromInterface184 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface184` to result `intoString761`
	// (`intoString761` is now tainted).
	intoString761 := fmt.Sprintln(fromInterface184)

	// Return the tainted `intoString761`:
	return intoString761
}

func TaintStepTest_FmtSscan_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString205` into `intoInterface525`.

	// Assume that `sourceCQL` has the underlying type of `fromString205`:
	fromString205 := sourceCQL.(string)

	// Declare `intoInterface525` variable:
	var intoInterface525 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString205` to parameter `intoInterface525`;
	// `intoInterface525` is now tainted.
	fmt.Sscan(fromString205, intoInterface525)

	// Return the tainted `intoInterface525`:
	return intoInterface525
}

func TaintStepTest_FmtSscanf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString899` into `intoInterface822`.

	// Assume that `sourceCQL` has the underlying type of `fromString899`:
	fromString899 := sourceCQL.(string)

	// Declare `intoInterface822` variable:
	var intoInterface822 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString899` to parameter `intoInterface822`;
	// `intoInterface822` is now tainted.
	fmt.Sscanf(fromString899, "", intoInterface822)

	// Return the tainted `intoInterface822`:
	return intoInterface822
}

func TaintStepTest_FmtSscanf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString592` into `intoInterface343`.

	// Assume that `sourceCQL` has the underlying type of `fromString592`:
	fromString592 := sourceCQL.(string)

	// Declare `intoInterface343` variable:
	var intoInterface343 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString592` to parameter `intoInterface343`;
	// `intoInterface343` is now tainted.
	fmt.Sscanf("", fromString592, intoInterface343)

	// Return the tainted `intoInterface343`:
	return intoInterface343
}

func TaintStepTest_FmtSscanln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString544` into `intoInterface116`.

	// Assume that `sourceCQL` has the underlying type of `fromString544`:
	fromString544 := sourceCQL.(string)

	// Declare `intoInterface116` variable:
	var intoInterface116 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString544` to parameter `intoInterface116`;
	// `intoInterface116` is now tainted.
	fmt.Sscanln(fromString544, intoInterface116)

	// Return the tainted `intoInterface116`:
	return intoInterface116
}

func TaintStepTest_FmtGoStringerGoString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromGoStringer149` into `intoString433`.

	// Assume that `sourceCQL` has the underlying type of `fromGoStringer149`:
	fromGoStringer149 := sourceCQL.(fmt.GoStringer)

	// Call the method that transfers the taint
	// from the receiver `fromGoStringer149` to the result `intoString433`
	// (`intoString433` is now tainted).
	intoString433 := fromGoStringer149.GoString()

	// Return the tainted `intoString433`:
	return intoString433
}

func TaintStepTest_FmtScanStateRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanState121` into `intoByte953`.

	// Assume that `sourceCQL` has the underlying type of `fromScanState121`:
	fromScanState121 := sourceCQL.(fmt.ScanState)

	// Declare `intoByte953` variable:
	var intoByte953 []byte

	// Call the method that transfers the taint
	// from the receiver `fromScanState121` to the argument `intoByte953`
	// (`intoByte953` is now tainted).
	fromScanState121.Read(intoByte953)

	// Return the tainted `intoByte953`:
	return intoByte953
}

func TaintStepTest_FmtScanStateReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanState851` into `intoRune876`.

	// Assume that `sourceCQL` has the underlying type of `fromScanState851`:
	fromScanState851 := sourceCQL.(fmt.ScanState)

	// Call the method that transfers the taint
	// from the receiver `fromScanState851` to the result `intoRune876`
	// (`intoRune876` is now tainted).
	intoRune876, _, _ := fromScanState851.ReadRune()

	// Return the tainted `intoRune876`:
	return intoRune876
}

func TaintStepTest_FmtStringerString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStringer235` into `intoString127`.

	// Assume that `sourceCQL` has the underlying type of `fromStringer235`:
	fromStringer235 := sourceCQL.(fmt.Stringer)

	// Call the method that transfers the taint
	// from the receiver `fromStringer235` to the result `intoString127`
	// (`intoString127` is now tainted).
	intoString127 := fromStringer235.String()

	// Return the tainted `intoString127`:
	return intoString127
}

func TaintStepTest_FmtScanStateToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanState777` into `intoByte946`.

	// Assume that `sourceCQL` has the underlying type of `fromScanState777`:
	fromScanState777 := sourceCQL.(fmt.ScanState)

	// Call the method that transfers the taint
	// from the receiver `fromScanState777` to the result `intoByte946`
	// (`intoByte946` is now tainted).
	intoByte946, _ := fromScanState777.Token(false, nil)

	// Return the tainted `intoByte946`:
	return intoByte946
}

func TaintStepTest_FmtStateWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte690` into `intoState763`.

	// Assume that `sourceCQL` has the underlying type of `fromByte690`:
	fromByte690 := sourceCQL.([]byte)

	// Declare `intoState763` variable:
	var intoState763 fmt.State

	// Call the method that transfers the taint
	// from the parameter `fromByte690` to the receiver `intoState763`
	// (`intoState763` is now tainted).
	intoState763.Write(fromByte690)

	// Return the tainted `intoState763`:
	return intoState763
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
