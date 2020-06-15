package main

import (
	"fmt"
	"io"
)

func TaintStepTest_FmtErrorf_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString342` into `intoError527`.

	// Assume that `sourceCQL` has the underlying type of `fromString342`:
	fromString342 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString342` to result `intoError527`
	// (`intoError527` is now tainted).
	intoError527 := fmt.Errorf(fromString342, nil)

	// Sink the tainted `intoError527`:
	sink(intoError527)
}

func TaintStepTest_FmtErrorf_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromInterface499` into `intoError440`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface499`:
	fromInterface499 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface499` to result `intoError440`
	// (`intoError440` is now tainted).
	intoError440 := fmt.Errorf("", fromInterface499)

	// Sink the tainted `intoError440`:
	sink(intoError440)
}

func TaintStepTest_FmtFprint_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface616` into `intoWriter683`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface616`:
	fromInterface616 := sourceCQL.([]interface{})

	// Declare `intoWriter683` variable:
	var intoWriter683 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromInterface616` to parameter `intoWriter683`;
	// `intoWriter683` is now tainted.
	fmt.Fprint(intoWriter683, fromInterface616)

	// Sink the tainted `intoWriter683`:
	sink(intoWriter683)
}

func TaintStepTest_FmtFprintf_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString167` into `intoWriter152`.

	// Assume that `sourceCQL` has the underlying type of `fromString167`:
	fromString167 := sourceCQL.(string)

	// Declare `intoWriter152` variable:
	var intoWriter152 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromString167` to parameter `intoWriter152`;
	// `intoWriter152` is now tainted.
	fmt.Fprintf(intoWriter152, fromString167, nil)

	// Sink the tainted `intoWriter152`:
	sink(intoWriter152)
}

func TaintStepTest_FmtFprintf_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromInterface801` into `intoWriter669`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface801`:
	fromInterface801 := sourceCQL.([]interface{})

	// Declare `intoWriter669` variable:
	var intoWriter669 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromInterface801` to parameter `intoWriter669`;
	// `intoWriter669` is now tainted.
	fmt.Fprintf(intoWriter669, "", fromInterface801)

	// Sink the tainted `intoWriter669`:
	sink(intoWriter669)
}

func TaintStepTest_FmtFprintln_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface976` into `intoWriter630`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface976`:
	fromInterface976 := sourceCQL.([]interface{})

	// Declare `intoWriter630` variable:
	var intoWriter630 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromInterface976` to parameter `intoWriter630`;
	// `intoWriter630` is now tainted.
	fmt.Fprintln(intoWriter630, fromInterface976)

	// Sink the tainted `intoWriter630`:
	sink(intoWriter630)
}

func TaintStepTest_FmtFscan_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader402` into `intoInterface784`.

	// Assume that `sourceCQL` has the underlying type of `fromReader402`:
	fromReader402 := sourceCQL.(io.Reader)

	// Declare `intoInterface784` variable:
	var intoInterface784 []interface{}

	// Call the function that transfers the taint
	// from the parameter `fromReader402` to parameter `intoInterface784`;
	// `intoInterface784` is now tainted.
	fmt.Fscan(fromReader402, intoInterface784)

	// Sink the tainted `intoInterface784`:
	sink(intoInterface784)
}

func TaintStepTest_FmtFscanf_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader679` into `intoInterface127`.

	// Assume that `sourceCQL` has the underlying type of `fromReader679`:
	fromReader679 := sourceCQL.(io.Reader)

	// Declare `intoInterface127` variable:
	var intoInterface127 []interface{}

	// Call the function that transfers the taint
	// from the parameter `fromReader679` to parameter `intoInterface127`;
	// `intoInterface127` is now tainted.
	fmt.Fscanf(fromReader679, "", intoInterface127)

	// Sink the tainted `intoInterface127`:
	sink(intoInterface127)
}

func TaintStepTest_FmtFscanf_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString208` into `intoInterface274`.

	// Assume that `sourceCQL` has the underlying type of `fromString208`:
	fromString208 := sourceCQL.(string)

	// Declare `intoInterface274` variable:
	var intoInterface274 []interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString208` to parameter `intoInterface274`;
	// `intoInterface274` is now tainted.
	fmt.Fscanf(nil, fromString208, intoInterface274)

	// Sink the tainted `intoInterface274`:
	sink(intoInterface274)
}

func TaintStepTest_FmtFscanln_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader332` into `intoInterface177`.

	// Assume that `sourceCQL` has the underlying type of `fromReader332`:
	fromReader332 := sourceCQL.(io.Reader)

	// Declare `intoInterface177` variable:
	var intoInterface177 []interface{}

	// Call the function that transfers the taint
	// from the parameter `fromReader332` to parameter `intoInterface177`;
	// `intoInterface177` is now tainted.
	fmt.Fscanln(fromReader332, intoInterface177)

	// Sink the tainted `intoInterface177`:
	sink(intoInterface177)
}

func TaintStepTest_FmtSprint_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface841` into `intoString300`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface841`:
	fromInterface841 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface841` to result `intoString300`
	// (`intoString300` is now tainted).
	intoString300 := fmt.Sprint(fromInterface841)

	// Sink the tainted `intoString300`:
	sink(intoString300)
}

func TaintStepTest_FmtSprintf_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString235` into `intoString140`.

	// Assume that `sourceCQL` has the underlying type of `fromString235`:
	fromString235 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString235` to result `intoString140`
	// (`intoString140` is now tainted).
	intoString140 := fmt.Sprintf(fromString235, nil)

	// Sink the tainted `intoString140`:
	sink(intoString140)
}

func TaintStepTest_FmtSprintf_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromInterface506` into `intoString833`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface506`:
	fromInterface506 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface506` to result `intoString833`
	// (`intoString833` is now tainted).
	intoString833 := fmt.Sprintf("", fromInterface506)

	// Sink the tainted `intoString833`:
	sink(intoString833)
}

func TaintStepTest_FmtSprintln_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface375` into `intoString181`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface375`:
	fromInterface375 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface375` to result `intoString181`
	// (`intoString181` is now tainted).
	intoString181 := fmt.Sprintln(fromInterface375)

	// Sink the tainted `intoString181`:
	sink(intoString181)
}

func TaintStepTest_FmtSscan_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString620` into `intoInterface355`.

	// Assume that `sourceCQL` has the underlying type of `fromString620`:
	fromString620 := sourceCQL.(string)

	// Declare `intoInterface355` variable:
	var intoInterface355 []interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString620` to parameter `intoInterface355`;
	// `intoInterface355` is now tainted.
	fmt.Sscan(fromString620, intoInterface355)

	// Sink the tainted `intoInterface355`:
	sink(intoInterface355)
}

func TaintStepTest_FmtSscanf_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString433` into `intoInterface453`.

	// Assume that `sourceCQL` has the underlying type of `fromString433`:
	fromString433 := sourceCQL.(string)

	// Declare `intoInterface453` variable:
	var intoInterface453 []interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString433` to parameter `intoInterface453`;
	// `intoInterface453` is now tainted.
	fmt.Sscanf(fromString433, "", intoInterface453)

	// Sink the tainted `intoInterface453`:
	sink(intoInterface453)
}

func TaintStepTest_FmtSscanf_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString868` into `intoInterface624`.

	// Assume that `sourceCQL` has the underlying type of `fromString868`:
	fromString868 := sourceCQL.(string)

	// Declare `intoInterface624` variable:
	var intoInterface624 []interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString868` to parameter `intoInterface624`;
	// `intoInterface624` is now tainted.
	fmt.Sscanf("", fromString868, intoInterface624)

	// Sink the tainted `intoInterface624`:
	sink(intoInterface624)
}

func TaintStepTest_FmtSscanln_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString559` into `intoInterface873`.

	// Assume that `sourceCQL` has the underlying type of `fromString559`:
	fromString559 := sourceCQL.(string)

	// Declare `intoInterface873` variable:
	var intoInterface873 []interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString559` to parameter `intoInterface873`;
	// `intoInterface873` is now tainted.
	fmt.Sscanln(fromString559, intoInterface873)

	// Sink the tainted `intoInterface873`:
	sink(intoInterface873)
}

func TaintStepTest_FmtGoStringerGoString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromGoStringer326` into `intoString624`.

	// Assume that `sourceCQL` has the underlying type of `fromGoStringer326`:
	fromGoStringer326 := sourceCQL.(fmt.GoStringer)

	// Call the method that transfers the taint
	// from the receiver `fromGoStringer326` to the result `intoString624`
	// (`intoString624` is now tainted).
	intoString624 := fromGoStringer326.GoString()

	// Sink the tainted `intoString624`:
	sink(intoString624)
}

func TaintStepTest_FmtScanStateRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromScanState821` into `intoByte113`.

	// Assume that `sourceCQL` has the underlying type of `fromScanState821`:
	fromScanState821 := sourceCQL.(fmt.ScanState)

	// Declare `intoByte113` variable:
	var intoByte113 []byte

	// Call the method that transfers the taint
	// from the receiver `fromScanState821` to the argument `intoByte113`
	// (`intoByte113` is now tainted).
	fromScanState821.Read(intoByte113)

	// Sink the tainted `intoByte113`:
	sink(intoByte113)
}

func TaintStepTest_FmtScanStateReadRune_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromScanState869` into `intoRune281`.

	// Assume that `sourceCQL` has the underlying type of `fromScanState869`:
	fromScanState869 := sourceCQL.(fmt.ScanState)

	// Call the method that transfers the taint
	// from the receiver `fromScanState869` to the result `intoRune281`
	// (`intoRune281` is now tainted).
	intoRune281, _, _ := fromScanState869.ReadRune()

	// Sink the tainted `intoRune281`:
	sink(intoRune281)
}

func TaintStepTest_FmtStringerString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromStringer874` into `intoString880`.

	// Assume that `sourceCQL` has the underlying type of `fromStringer874`:
	fromStringer874 := sourceCQL.(fmt.Stringer)

	// Call the method that transfers the taint
	// from the receiver `fromStringer874` to the result `intoString880`
	// (`intoString880` is now tainted).
	intoString880 := fromStringer874.String()

	// Sink the tainted `intoString880`:
	sink(intoString880)
}

func TaintStepTest_FmtScanStateToken_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromScanState144` into `intoByte918`.

	// Assume that `sourceCQL` has the underlying type of `fromScanState144`:
	fromScanState144 := sourceCQL.(fmt.ScanState)

	// Call the method that transfers the taint
	// from the receiver `fromScanState144` to the result `intoByte918`
	// (`intoByte918` is now tainted).
	intoByte918, _ := fromScanState144.Token(false, nil)

	// Sink the tainted `intoByte918`:
	sink(intoByte918)
}

func TaintStepTest_FmtStateWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte722` into `intoState194`.

	// Assume that `sourceCQL` has the underlying type of `fromByte722`:
	fromByte722 := sourceCQL.([]byte)

	// Declare `intoState194` variable:
	var intoState194 fmt.State

	// Call the method that transfers the taint
	// from the parameter `fromByte722` to the receiver `intoState194`
	// (`intoState194` is now tainted).
	intoState194.Write(fromByte722)

	// Sink the tainted `intoState194`:
	sink(intoState194)
}

func RunAllTaints_Fmt() {
	{
		source := newSource()
		TaintStepTest_FmtErrorf_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtErrorf_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtFprint_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtFprintf_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtFprintf_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtFprintln_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtFscan_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtFscanf_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtFscanf_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtFscanln_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtSprint_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtSprintf_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtSprintf_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtSprintln_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtSscan_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtSscanf_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtSscanf_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtSscanln_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtGoStringerGoString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtScanStateRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtScanStateReadRune_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtStringerString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtScanStateToken_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtStateWrite_B0I0O0(source)
	}
}
