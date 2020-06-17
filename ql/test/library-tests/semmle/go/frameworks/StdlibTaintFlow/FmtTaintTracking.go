package main

import (
	"fmt"
	"io"
)

func TaintStepTest_FmtErrorf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString245` into `intoError593`.

	// Assume that `sourceCQL` has the underlying type of `fromString245`:
	fromString245 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString245` to result `intoError593`
	// (`intoError593` is now tainted).
	intoError593 := fmt.Errorf(fromString245, nil)

	// Return the tainted `intoError593`:
	return intoError593
}

func TaintStepTest_FmtErrorf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface328` into `intoError975`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface328`:
	fromInterface328 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface328` to result `intoError975`
	// (`intoError975` is now tainted).
	intoError975 := fmt.Errorf("", fromInterface328)

	// Return the tainted `intoError975`:
	return intoError975
}

func TaintStepTest_FmtFprint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface678` into `intoWriter202`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface678`:
	fromInterface678 := sourceCQL.(interface{})

	// Declare `intoWriter202` variable:
	var intoWriter202 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromInterface678` to parameter `intoWriter202`;
	// `intoWriter202` is now tainted.
	fmt.Fprint(intoWriter202, fromInterface678)

	// Return the tainted `intoWriter202`:
	return intoWriter202
}

func TaintStepTest_FmtFprintf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString563` into `intoWriter560`.

	// Assume that `sourceCQL` has the underlying type of `fromString563`:
	fromString563 := sourceCQL.(string)

	// Declare `intoWriter560` variable:
	var intoWriter560 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromString563` to parameter `intoWriter560`;
	// `intoWriter560` is now tainted.
	fmt.Fprintf(intoWriter560, fromString563, nil)

	// Return the tainted `intoWriter560`:
	return intoWriter560
}

func TaintStepTest_FmtFprintf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface560` into `intoWriter411`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface560`:
	fromInterface560 := sourceCQL.(interface{})

	// Declare `intoWriter411` variable:
	var intoWriter411 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromInterface560` to parameter `intoWriter411`;
	// `intoWriter411` is now tainted.
	fmt.Fprintf(intoWriter411, "", fromInterface560)

	// Return the tainted `intoWriter411`:
	return intoWriter411
}

func TaintStepTest_FmtFprintln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface850` into `intoWriter272`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface850`:
	fromInterface850 := sourceCQL.(interface{})

	// Declare `intoWriter272` variable:
	var intoWriter272 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromInterface850` to parameter `intoWriter272`;
	// `intoWriter272` is now tainted.
	fmt.Fprintln(intoWriter272, fromInterface850)

	// Return the tainted `intoWriter272`:
	return intoWriter272
}

func TaintStepTest_FmtFscan_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader517` into `intoInterface315`.

	// Assume that `sourceCQL` has the underlying type of `fromReader517`:
	fromReader517 := sourceCQL.(io.Reader)

	// Declare `intoInterface315` variable:
	var intoInterface315 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromReader517` to parameter `intoInterface315`;
	// `intoInterface315` is now tainted.
	fmt.Fscan(fromReader517, intoInterface315)

	// Return the tainted `intoInterface315`:
	return intoInterface315
}

func TaintStepTest_FmtFscanf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader910` into `intoInterface574`.

	// Assume that `sourceCQL` has the underlying type of `fromReader910`:
	fromReader910 := sourceCQL.(io.Reader)

	// Declare `intoInterface574` variable:
	var intoInterface574 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromReader910` to parameter `intoInterface574`;
	// `intoInterface574` is now tainted.
	fmt.Fscanf(fromReader910, "", intoInterface574)

	// Return the tainted `intoInterface574`:
	return intoInterface574
}

func TaintStepTest_FmtFscanf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString915` into `intoInterface408`.

	// Assume that `sourceCQL` has the underlying type of `fromString915`:
	fromString915 := sourceCQL.(string)

	// Declare `intoInterface408` variable:
	var intoInterface408 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString915` to parameter `intoInterface408`;
	// `intoInterface408` is now tainted.
	fmt.Fscanf(nil, fromString915, intoInterface408)

	// Return the tainted `intoInterface408`:
	return intoInterface408
}

func TaintStepTest_FmtFscanln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader694` into `intoInterface648`.

	// Assume that `sourceCQL` has the underlying type of `fromReader694`:
	fromReader694 := sourceCQL.(io.Reader)

	// Declare `intoInterface648` variable:
	var intoInterface648 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromReader694` to parameter `intoInterface648`;
	// `intoInterface648` is now tainted.
	fmt.Fscanln(fromReader694, intoInterface648)

	// Return the tainted `intoInterface648`:
	return intoInterface648
}

func TaintStepTest_FmtSprint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface324` into `intoString770`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface324`:
	fromInterface324 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface324` to result `intoString770`
	// (`intoString770` is now tainted).
	intoString770 := fmt.Sprint(fromInterface324)

	// Return the tainted `intoString770`:
	return intoString770
}

func TaintStepTest_FmtSprintf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString141` into `intoString895`.

	// Assume that `sourceCQL` has the underlying type of `fromString141`:
	fromString141 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString141` to result `intoString895`
	// (`intoString895` is now tainted).
	intoString895 := fmt.Sprintf(fromString141, nil)

	// Return the tainted `intoString895`:
	return intoString895
}

func TaintStepTest_FmtSprintf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface256` into `intoString730`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface256`:
	fromInterface256 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface256` to result `intoString730`
	// (`intoString730` is now tainted).
	intoString730 := fmt.Sprintf("", fromInterface256)

	// Return the tainted `intoString730`:
	return intoString730
}

func TaintStepTest_FmtSprintln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface539` into `intoString977`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface539`:
	fromInterface539 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface539` to result `intoString977`
	// (`intoString977` is now tainted).
	intoString977 := fmt.Sprintln(fromInterface539)

	// Return the tainted `intoString977`:
	return intoString977
}

func TaintStepTest_FmtSscan_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString282` into `intoInterface920`.

	// Assume that `sourceCQL` has the underlying type of `fromString282`:
	fromString282 := sourceCQL.(string)

	// Declare `intoInterface920` variable:
	var intoInterface920 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString282` to parameter `intoInterface920`;
	// `intoInterface920` is now tainted.
	fmt.Sscan(fromString282, intoInterface920)

	// Return the tainted `intoInterface920`:
	return intoInterface920
}

func TaintStepTest_FmtSscanf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString190` into `intoInterface631`.

	// Assume that `sourceCQL` has the underlying type of `fromString190`:
	fromString190 := sourceCQL.(string)

	// Declare `intoInterface631` variable:
	var intoInterface631 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString190` to parameter `intoInterface631`;
	// `intoInterface631` is now tainted.
	fmt.Sscanf(fromString190, "", intoInterface631)

	// Return the tainted `intoInterface631`:
	return intoInterface631
}

func TaintStepTest_FmtSscanf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString792` into `intoInterface515`.

	// Assume that `sourceCQL` has the underlying type of `fromString792`:
	fromString792 := sourceCQL.(string)

	// Declare `intoInterface515` variable:
	var intoInterface515 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString792` to parameter `intoInterface515`;
	// `intoInterface515` is now tainted.
	fmt.Sscanf("", fromString792, intoInterface515)

	// Return the tainted `intoInterface515`:
	return intoInterface515
}

func TaintStepTest_FmtSscanln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString832` into `intoInterface703`.

	// Assume that `sourceCQL` has the underlying type of `fromString832`:
	fromString832 := sourceCQL.(string)

	// Declare `intoInterface703` variable:
	var intoInterface703 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString832` to parameter `intoInterface703`;
	// `intoInterface703` is now tainted.
	fmt.Sscanln(fromString832, intoInterface703)

	// Return the tainted `intoInterface703`:
	return intoInterface703
}

func TaintStepTest_FmtGoStringerGoString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromGoStringer410` into `intoString686`.

	// Assume that `sourceCQL` has the underlying type of `fromGoStringer410`:
	fromGoStringer410 := sourceCQL.(fmt.GoStringer)

	// Call the method that transfers the taint
	// from the receiver `fromGoStringer410` to the result `intoString686`
	// (`intoString686` is now tainted).
	intoString686 := fromGoStringer410.GoString()

	// Return the tainted `intoString686`:
	return intoString686
}

func TaintStepTest_FmtScanStateRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanState852` into `intoByte390`.

	// Assume that `sourceCQL` has the underlying type of `fromScanState852`:
	fromScanState852 := sourceCQL.(fmt.ScanState)

	// Declare `intoByte390` variable:
	var intoByte390 []byte

	// Call the method that transfers the taint
	// from the receiver `fromScanState852` to the argument `intoByte390`
	// (`intoByte390` is now tainted).
	fromScanState852.Read(intoByte390)

	// Return the tainted `intoByte390`:
	return intoByte390
}

func TaintStepTest_FmtScanStateReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanState495` into `intoRune377`.

	// Assume that `sourceCQL` has the underlying type of `fromScanState495`:
	fromScanState495 := sourceCQL.(fmt.ScanState)

	// Call the method that transfers the taint
	// from the receiver `fromScanState495` to the result `intoRune377`
	// (`intoRune377` is now tainted).
	intoRune377, _, _ := fromScanState495.ReadRune()

	// Return the tainted `intoRune377`:
	return intoRune377
}

func TaintStepTest_FmtStringerString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStringer774` into `intoString360`.

	// Assume that `sourceCQL` has the underlying type of `fromStringer774`:
	fromStringer774 := sourceCQL.(fmt.Stringer)

	// Call the method that transfers the taint
	// from the receiver `fromStringer774` to the result `intoString360`
	// (`intoString360` is now tainted).
	intoString360 := fromStringer774.String()

	// Return the tainted `intoString360`:
	return intoString360
}

func TaintStepTest_FmtScanStateToken_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromScanState807` into `intoByte814`.

	// Assume that `sourceCQL` has the underlying type of `fromScanState807`:
	fromScanState807 := sourceCQL.(fmt.ScanState)

	// Call the method that transfers the taint
	// from the receiver `fromScanState807` to the result `intoByte814`
	// (`intoByte814` is now tainted).
	intoByte814, _ := fromScanState807.Token(false, nil)

	// Return the tainted `intoByte814`:
	return intoByte814
}

func TaintStepTest_FmtStateWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte754` into `intoState171`.

	// Assume that `sourceCQL` has the underlying type of `fromByte754`:
	fromByte754 := sourceCQL.([]byte)

	// Declare `intoState171` variable:
	var intoState171 fmt.State

	// Call the method that transfers the taint
	// from the parameter `fromByte754` to the receiver `intoState171`
	// (`intoState171` is now tainted).
	intoState171.Write(fromByte754)

	// Return the tainted `intoState171`:
	return intoState171
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
