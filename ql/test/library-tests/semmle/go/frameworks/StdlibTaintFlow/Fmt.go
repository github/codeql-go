package main

import (
	"fmt"
	"io"
)

func TaintStepTest_FmtErrorf(sourceCQL interface{}) {
	// The flow is from `a` into `into565`.

	// Assume that `sourceCQL` has the underlying type of `a`:
	a := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `a` to result `into565`
	// (`into565` is now tainted).
	into565 := fmt.Errorf("", a)

	// Sink the tainted `into565`:
	sink(into565)
}

func TaintStepTest_FmtFprint(sourceCQL interface{}) {
	// The flow is from `a` into `w`.

	// Assume that `sourceCQL` has the underlying type of `a`:
	a := sourceCQL.([]interface{})

	// Declare `w` variable:
	var w io.Writer

	// Call the function that transfers the taint
	// from the parameter `a` to parameter `w`;
	// `w` is now tainted.
	fmt.Fprint(w, a)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_FmtFprintf(sourceCQL interface{}) {
	// The flow is from `a` into `w`.

	// Assume that `sourceCQL` has the underlying type of `a`:
	a := sourceCQL.([]interface{})

	// Declare `w` variable:
	var w io.Writer

	// Call the function that transfers the taint
	// from the parameter `a` to parameter `w`;
	// `w` is now tainted.
	fmt.Fprintf(w, "", a)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_FmtFprintln(sourceCQL interface{}) {
	// The flow is from `a` into `w`.

	// Assume that `sourceCQL` has the underlying type of `a`:
	a := sourceCQL.([]interface{})

	// Declare `w` variable:
	var w io.Writer

	// Call the function that transfers the taint
	// from the parameter `a` to parameter `w`;
	// `w` is now tainted.
	fmt.Fprintln(w, a)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_FmtFscan(sourceCQL interface{}) {
	// The flow is from `r` into `a`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Declare `a` variable:
	var a []interface{}

	// Call the function that transfers the taint
	// from the parameter `r` to parameter `a`;
	// `a` is now tainted.
	fmt.Fscan(r, a)

	// Sink the tainted `a`:
	sink(a)
}

func TaintStepTest_FmtFscanf(sourceCQL interface{}) {
	// The flow is from `r` into `a`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Declare `a` variable:
	var a []interface{}

	// Call the function that transfers the taint
	// from the parameter `r` to parameter `a`;
	// `a` is now tainted.
	fmt.Fscanf(r, "", a)

	// Sink the tainted `a`:
	sink(a)
}

func TaintStepTest_FmtFscanln(sourceCQL interface{}) {
	// The flow is from `r` into `a`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Declare `a` variable:
	var a []interface{}

	// Call the function that transfers the taint
	// from the parameter `r` to parameter `a`;
	// `a` is now tainted.
	fmt.Fscanln(r, a)

	// Sink the tainted `a`:
	sink(a)
}

func TaintStepTest_FmtSprint(sourceCQL interface{}) {
	// The flow is from `a` into `into818`.

	// Assume that `sourceCQL` has the underlying type of `a`:
	a := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `a` to result `into818`
	// (`into818` is now tainted).
	into818 := fmt.Sprint(a)

	// Sink the tainted `into818`:
	sink(into818)
}

func TaintStepTest_FmtSprintf(sourceCQL interface{}) {
	// The flow is from `a` into `into442`.

	// Assume that `sourceCQL` has the underlying type of `a`:
	a := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `a` to result `into442`
	// (`into442` is now tainted).
	into442 := fmt.Sprintf("", a)

	// Sink the tainted `into442`:
	sink(into442)
}

func TaintStepTest_FmtSprintln(sourceCQL interface{}) {
	// The flow is from `a` into `into402`.

	// Assume that `sourceCQL` has the underlying type of `a`:
	a := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `a` to result `into402`
	// (`into402` is now tainted).
	into402 := fmt.Sprintln(a)

	// Sink the tainted `into402`:
	sink(into402)
}

func TaintStepTest_FmtSscan(sourceCQL interface{}) {
	// The flow is from `str` into `a`.

	// Assume that `sourceCQL` has the underlying type of `str`:
	str := sourceCQL.(string)

	// Declare `a` variable:
	var a []interface{}

	// Call the function that transfers the taint
	// from the parameter `str` to parameter `a`;
	// `a` is now tainted.
	fmt.Sscan(str, a)

	// Sink the tainted `a`:
	sink(a)
}

func TaintStepTest_FmtSscanf(sourceCQL interface{}) {
	// The flow is from `str` into `a`.

	// Assume that `sourceCQL` has the underlying type of `str`:
	str := sourceCQL.(string)

	// Declare `a` variable:
	var a []interface{}

	// Call the function that transfers the taint
	// from the parameter `str` to parameter `a`;
	// `a` is now tainted.
	fmt.Sscanf(str, "", a)

	// Sink the tainted `a`:
	sink(a)
}

func TaintStepTest_FmtSscanln(sourceCQL interface{}) {
	// The flow is from `str` into `a`.

	// Assume that `sourceCQL` has the underlying type of `str`:
	str := sourceCQL.(string)

	// Declare `a` variable:
	var a []interface{}

	// Call the function that transfers the taint
	// from the parameter `str` to parameter `a`;
	// `a` is now tainted.
	fmt.Sscanln(str, a)

	// Sink the tainted `a`:
	sink(a)
}

func TaintStepTest_FmtGoStringerGoString(sourceCQL interface{}) {
	// The flow is from `from962` into `into570`.

	// Assume that `sourceCQL` has the underlying type of `from962`:
	from962 := sourceCQL.(fmt.GoStringer)

	// Call the method that transfers the taint
	// from the receiver `from962` to the result `into570`
	// (`into570` is now tainted).
	into570 := from962.GoString()

	// Sink the tainted `into570`:
	sink(into570)
}

func TaintStepTest_FmtScanStateRead(sourceCQL interface{}) {
	// The flow is from `from127` into `buf`.

	// Assume that `sourceCQL` has the underlying type of `from127`:
	from127 := sourceCQL.(fmt.ScanState)

	// Declare `buf` variable:
	var buf []byte

	// Call the method that transfers the taint
	// from the receiver `from127` to the argument `buf`
	// (`buf` is now tainted).
	from127.Read(buf)

	// Sink the tainted `buf`:
	sink(buf)
}

func TaintStepTest_FmtScanStateReadRune(sourceCQL interface{}) {
	// The flow is from `from370` into `r`.

	// Assume that `sourceCQL` has the underlying type of `from370`:
	from370 := sourceCQL.(fmt.ScanState)

	// Call the method that transfers the taint
	// from the receiver `from370` to the result `r`
	// (`r` is now tainted).
	r, _, _ := from370.ReadRune()

	// Sink the tainted `r`:
	sink(r)
}

func TaintStepTest_FmtScanStateToken(sourceCQL interface{}) {
	// The flow is from `from265` into `token`.

	// Assume that `sourceCQL` has the underlying type of `from265`:
	from265 := sourceCQL.(fmt.ScanState)

	// Call the method that transfers the taint
	// from the receiver `from265` to the result `token`
	// (`token` is now tainted).
	token, _ := from265.Token(false, nil)

	// Sink the tainted `token`:
	sink(token)
}

func TaintStepTest_FmtStateWrite(sourceCQL interface{}) {
	// The flow is from `b` into `into947`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare `into947` variable:
	var into947 fmt.State

	// Call the method that transfers the taint
	// from the parameter `b` to the receiver `into947`
	// (`into947` is now tainted).
	into947.Write(b)

	// Sink the tainted `into947`:
	sink(into947)
}

func TaintStepTest_FmtStringerString(sourceCQL interface{}) {
	// The flow is from `from201` into `into239`.

	// Assume that `sourceCQL` has the underlying type of `from201`:
	from201 := sourceCQL.(fmt.Stringer)

	// Call the method that transfers the taint
	// from the receiver `from201` to the result `into239`
	// (`into239` is now tainted).
	into239 := from201.String()

	// Sink the tainted `into239`:
	sink(into239)
}

func RunAllTaints_Fmt(v interface{}) {
	{
		source := newSource()
		TaintStepTest_FmtErrorf(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtFprint(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtFprintf(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtFprintln(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtFscan(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtFscanf(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtFscanln(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtSprint(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtSprintf(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtSprintln(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtSscan(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtSscanf(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtSscanln(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtGoStringerGoString(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtScanStateRead(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtScanStateReadRune(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtScanStateToken(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtStateWrite(source)
	}
	{
		source := newSource()
		TaintStepTest_FmtStringerString(source)
	}
}
