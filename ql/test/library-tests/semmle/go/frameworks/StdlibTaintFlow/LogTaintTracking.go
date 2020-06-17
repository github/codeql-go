// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"log"
)

func TaintStepTest_LogNew_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromLogger346` into `intoWriter806`.

	// Assume that `sourceCQL` has the underlying type of `fromLogger346`:
	fromLogger346 := sourceCQL.(*log.Logger)

	// Declare `intoWriter806` variable:
	var intoWriter806 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter806`:
	intermediateCQL := log.New(intoWriter806, "", 0)

	// Extra step (`fromLogger346` taints `intermediateCQL`, which taints `intoWriter806`:
	link(fromLogger346, intermediateCQL)

	// Return the tainted `intoWriter806`:
	return intoWriter806
}

func TaintStepTest_LogLoggerFatal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface748` into `intoLogger897`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface748`:
	fromInterface748 := sourceCQL.(interface{})

	// Declare `intoLogger897` variable:
	var intoLogger897 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface748` to the receiver `intoLogger897`
	// (`intoLogger897` is now tainted).
	intoLogger897.Fatal(fromInterface748)

	// Return the tainted `intoLogger897`:
	return intoLogger897
}

func TaintStepTest_LogLoggerFatalf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString324` into `intoLogger448`.

	// Assume that `sourceCQL` has the underlying type of `fromString324`:
	fromString324 := sourceCQL.(string)

	// Declare `intoLogger448` variable:
	var intoLogger448 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromString324` to the receiver `intoLogger448`
	// (`intoLogger448` is now tainted).
	intoLogger448.Fatalf(fromString324, nil)

	// Return the tainted `intoLogger448`:
	return intoLogger448
}

func TaintStepTest_LogLoggerFatalf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface940` into `intoLogger581`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface940`:
	fromInterface940 := sourceCQL.(interface{})

	// Declare `intoLogger581` variable:
	var intoLogger581 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface940` to the receiver `intoLogger581`
	// (`intoLogger581` is now tainted).
	intoLogger581.Fatalf("", fromInterface940)

	// Return the tainted `intoLogger581`:
	return intoLogger581
}

func TaintStepTest_LogLoggerFatalln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface499` into `intoLogger508`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface499`:
	fromInterface499 := sourceCQL.(interface{})

	// Declare `intoLogger508` variable:
	var intoLogger508 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface499` to the receiver `intoLogger508`
	// (`intoLogger508` is now tainted).
	intoLogger508.Fatalln(fromInterface499)

	// Return the tainted `intoLogger508`:
	return intoLogger508
}

func TaintStepTest_LogLoggerPanic_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface565` into `intoLogger438`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface565`:
	fromInterface565 := sourceCQL.(interface{})

	// Declare `intoLogger438` variable:
	var intoLogger438 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface565` to the receiver `intoLogger438`
	// (`intoLogger438` is now tainted).
	intoLogger438.Panic(fromInterface565)

	// Return the tainted `intoLogger438`:
	return intoLogger438
}

func TaintStepTest_LogLoggerPanicf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString618` into `intoLogger115`.

	// Assume that `sourceCQL` has the underlying type of `fromString618`:
	fromString618 := sourceCQL.(string)

	// Declare `intoLogger115` variable:
	var intoLogger115 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromString618` to the receiver `intoLogger115`
	// (`intoLogger115` is now tainted).
	intoLogger115.Panicf(fromString618, nil)

	// Return the tainted `intoLogger115`:
	return intoLogger115
}

func TaintStepTest_LogLoggerPanicf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface534` into `intoLogger483`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface534`:
	fromInterface534 := sourceCQL.(interface{})

	// Declare `intoLogger483` variable:
	var intoLogger483 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface534` to the receiver `intoLogger483`
	// (`intoLogger483` is now tainted).
	intoLogger483.Panicf("", fromInterface534)

	// Return the tainted `intoLogger483`:
	return intoLogger483
}

func TaintStepTest_LogLoggerPanicln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface848` into `intoLogger783`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface848`:
	fromInterface848 := sourceCQL.(interface{})

	// Declare `intoLogger783` variable:
	var intoLogger783 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface848` to the receiver `intoLogger783`
	// (`intoLogger783` is now tainted).
	intoLogger783.Panicln(fromInterface848)

	// Return the tainted `intoLogger783`:
	return intoLogger783
}

func TaintStepTest_LogLoggerPrint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface655` into `intoLogger498`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface655`:
	fromInterface655 := sourceCQL.(interface{})

	// Declare `intoLogger498` variable:
	var intoLogger498 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface655` to the receiver `intoLogger498`
	// (`intoLogger498` is now tainted).
	intoLogger498.Print(fromInterface655)

	// Return the tainted `intoLogger498`:
	return intoLogger498
}

func TaintStepTest_LogLoggerPrintf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString309` into `intoLogger274`.

	// Assume that `sourceCQL` has the underlying type of `fromString309`:
	fromString309 := sourceCQL.(string)

	// Declare `intoLogger274` variable:
	var intoLogger274 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromString309` to the receiver `intoLogger274`
	// (`intoLogger274` is now tainted).
	intoLogger274.Printf(fromString309, nil)

	// Return the tainted `intoLogger274`:
	return intoLogger274
}

func TaintStepTest_LogLoggerPrintf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface540` into `intoLogger785`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface540`:
	fromInterface540 := sourceCQL.(interface{})

	// Declare `intoLogger785` variable:
	var intoLogger785 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface540` to the receiver `intoLogger785`
	// (`intoLogger785` is now tainted).
	intoLogger785.Printf("", fromInterface540)

	// Return the tainted `intoLogger785`:
	return intoLogger785
}

func TaintStepTest_LogLoggerPrintln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface179` into `intoLogger852`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface179`:
	fromInterface179 := sourceCQL.(interface{})

	// Declare `intoLogger852` variable:
	var intoLogger852 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface179` to the receiver `intoLogger852`
	// (`intoLogger852` is now tainted).
	intoLogger852.Println(fromInterface179)

	// Return the tainted `intoLogger852`:
	return intoLogger852
}

func TaintStepTest_LogLoggerSetOutput_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromLogger376` into `intoWriter181`.

	// Assume that `sourceCQL` has the underlying type of `fromLogger376`:
	fromLogger376 := sourceCQL.(log.Logger)

	// Declare `intoWriter181` variable:
	var intoWriter181 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromLogger376` to the argument `intoWriter181`
	// (`intoWriter181` is now tainted).
	fromLogger376.SetOutput(intoWriter181)

	// Return the tainted `intoWriter181`:
	return intoWriter181
}

func TaintStepTest_LogLoggerSetPrefix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString727` into `intoLogger168`.

	// Assume that `sourceCQL` has the underlying type of `fromString727`:
	fromString727 := sourceCQL.(string)

	// Declare `intoLogger168` variable:
	var intoLogger168 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromString727` to the receiver `intoLogger168`
	// (`intoLogger168` is now tainted).
	intoLogger168.SetPrefix(fromString727)

	// Return the tainted `intoLogger168`:
	return intoLogger168
}

func TaintStepTest_LogLoggerWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter784` into `intoLogger314`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter784`:
	fromWriter784 := sourceCQL.(io.Writer)

	// Declare `intoLogger314` variable:
	var intoLogger314 log.Logger

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoLogger314`:
	intermediateCQL := intoLogger314.Writer()

	// Extra step (`fromWriter784` taints `intermediateCQL`, which taints `intoLogger314`:
	link(fromWriter784, intermediateCQL)

	// Return the tainted `intoLogger314`:
	return intoLogger314
}

func TaintStepTest_LogLoggerWriter_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromLogger164` into `intoWriter460`.

	// Assume that `sourceCQL` has the underlying type of `fromLogger164`:
	fromLogger164 := sourceCQL.(log.Logger)

	// Call the method that transfers the taint
	// from the receiver `fromLogger164` to the result `intoWriter460`
	// (`intoWriter460` is now tainted).
	intoWriter460 := fromLogger164.Writer()

	// Return the tainted `intoWriter460`:
	return intoWriter460
}

func RunAllTaints_Log() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_LogNew_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_LogLoggerFatal_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_LogLoggerFatalf_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_LogLoggerFatalf_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_LogLoggerFatalln_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_LogLoggerPanic_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_LogLoggerPanicf_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_LogLoggerPanicf_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_LogLoggerPanicln_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_LogLoggerPrint_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_LogLoggerPrintf_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_LogLoggerPrintf_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_LogLoggerPrintln_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_LogLoggerSetOutput_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_LogLoggerSetPrefix_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_LogLoggerWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_LogLoggerWriter_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
