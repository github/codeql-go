package main

import (
	"io"
	"log"
)

func TaintStepTest_LogNew_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromLogger136` into `intoWriter742`.

	// Assume that `sourceCQL` has the underlying type of `fromLogger136`:
	fromLogger136 := sourceCQL.(*log.Logger)

	// Declare `intoWriter742` variable:
	var intoWriter742 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter742`:
	intermediateCQL := log.New(intoWriter742, "", 0)

	// Extra step (`fromLogger136` taints `intermediateCQL`, which taints `intoWriter742`:
	link(fromLogger136, intermediateCQL)

	// Return the tainted `intoWriter742`:
	return intoWriter742
}

func TaintStepTest_LogLoggerFatal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface706` into `intoLogger111`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface706`:
	fromInterface706 := sourceCQL.(interface{})

	// Declare `intoLogger111` variable:
	var intoLogger111 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface706` to the receiver `intoLogger111`
	// (`intoLogger111` is now tainted).
	intoLogger111.Fatal(fromInterface706)

	// Return the tainted `intoLogger111`:
	return intoLogger111
}

func TaintStepTest_LogLoggerFatalf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString484` into `intoLogger209`.

	// Assume that `sourceCQL` has the underlying type of `fromString484`:
	fromString484 := sourceCQL.(string)

	// Declare `intoLogger209` variable:
	var intoLogger209 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromString484` to the receiver `intoLogger209`
	// (`intoLogger209` is now tainted).
	intoLogger209.Fatalf(fromString484, nil)

	// Return the tainted `intoLogger209`:
	return intoLogger209
}

func TaintStepTest_LogLoggerFatalf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface345` into `intoLogger374`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface345`:
	fromInterface345 := sourceCQL.(interface{})

	// Declare `intoLogger374` variable:
	var intoLogger374 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface345` to the receiver `intoLogger374`
	// (`intoLogger374` is now tainted).
	intoLogger374.Fatalf("", fromInterface345)

	// Return the tainted `intoLogger374`:
	return intoLogger374
}

func TaintStepTest_LogLoggerFatalln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface587` into `intoLogger651`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface587`:
	fromInterface587 := sourceCQL.(interface{})

	// Declare `intoLogger651` variable:
	var intoLogger651 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface587` to the receiver `intoLogger651`
	// (`intoLogger651` is now tainted).
	intoLogger651.Fatalln(fromInterface587)

	// Return the tainted `intoLogger651`:
	return intoLogger651
}

func TaintStepTest_LogLoggerPanic_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface782` into `intoLogger653`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface782`:
	fromInterface782 := sourceCQL.(interface{})

	// Declare `intoLogger653` variable:
	var intoLogger653 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface782` to the receiver `intoLogger653`
	// (`intoLogger653` is now tainted).
	intoLogger653.Panic(fromInterface782)

	// Return the tainted `intoLogger653`:
	return intoLogger653
}

func TaintStepTest_LogLoggerPanicf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString879` into `intoLogger909`.

	// Assume that `sourceCQL` has the underlying type of `fromString879`:
	fromString879 := sourceCQL.(string)

	// Declare `intoLogger909` variable:
	var intoLogger909 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromString879` to the receiver `intoLogger909`
	// (`intoLogger909` is now tainted).
	intoLogger909.Panicf(fromString879, nil)

	// Return the tainted `intoLogger909`:
	return intoLogger909
}

func TaintStepTest_LogLoggerPanicf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface595` into `intoLogger293`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface595`:
	fromInterface595 := sourceCQL.(interface{})

	// Declare `intoLogger293` variable:
	var intoLogger293 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface595` to the receiver `intoLogger293`
	// (`intoLogger293` is now tainted).
	intoLogger293.Panicf("", fromInterface595)

	// Return the tainted `intoLogger293`:
	return intoLogger293
}

func TaintStepTest_LogLoggerPanicln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface269` into `intoLogger747`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface269`:
	fromInterface269 := sourceCQL.(interface{})

	// Declare `intoLogger747` variable:
	var intoLogger747 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface269` to the receiver `intoLogger747`
	// (`intoLogger747` is now tainted).
	intoLogger747.Panicln(fromInterface269)

	// Return the tainted `intoLogger747`:
	return intoLogger747
}

func TaintStepTest_LogLoggerPrint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface961` into `intoLogger219`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface961`:
	fromInterface961 := sourceCQL.(interface{})

	// Declare `intoLogger219` variable:
	var intoLogger219 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface961` to the receiver `intoLogger219`
	// (`intoLogger219` is now tainted).
	intoLogger219.Print(fromInterface961)

	// Return the tainted `intoLogger219`:
	return intoLogger219
}

func TaintStepTest_LogLoggerPrintf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString998` into `intoLogger834`.

	// Assume that `sourceCQL` has the underlying type of `fromString998`:
	fromString998 := sourceCQL.(string)

	// Declare `intoLogger834` variable:
	var intoLogger834 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromString998` to the receiver `intoLogger834`
	// (`intoLogger834` is now tainted).
	intoLogger834.Printf(fromString998, nil)

	// Return the tainted `intoLogger834`:
	return intoLogger834
}

func TaintStepTest_LogLoggerPrintf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface796` into `intoLogger235`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface796`:
	fromInterface796 := sourceCQL.(interface{})

	// Declare `intoLogger235` variable:
	var intoLogger235 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface796` to the receiver `intoLogger235`
	// (`intoLogger235` is now tainted).
	intoLogger235.Printf("", fromInterface796)

	// Return the tainted `intoLogger235`:
	return intoLogger235
}

func TaintStepTest_LogLoggerPrintln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface469` into `intoLogger213`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface469`:
	fromInterface469 := sourceCQL.(interface{})

	// Declare `intoLogger213` variable:
	var intoLogger213 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface469` to the receiver `intoLogger213`
	// (`intoLogger213` is now tainted).
	intoLogger213.Println(fromInterface469)

	// Return the tainted `intoLogger213`:
	return intoLogger213
}

func TaintStepTest_LogLoggerSetOutput_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromLogger914` into `intoWriter856`.

	// Assume that `sourceCQL` has the underlying type of `fromLogger914`:
	fromLogger914 := sourceCQL.(log.Logger)

	// Declare `intoWriter856` variable:
	var intoWriter856 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromLogger914` to the argument `intoWriter856`
	// (`intoWriter856` is now tainted).
	fromLogger914.SetOutput(intoWriter856)

	// Return the tainted `intoWriter856`:
	return intoWriter856
}

func TaintStepTest_LogLoggerSetPrefix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString421` into `intoLogger476`.

	// Assume that `sourceCQL` has the underlying type of `fromString421`:
	fromString421 := sourceCQL.(string)

	// Declare `intoLogger476` variable:
	var intoLogger476 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromString421` to the receiver `intoLogger476`
	// (`intoLogger476` is now tainted).
	intoLogger476.SetPrefix(fromString421)

	// Return the tainted `intoLogger476`:
	return intoLogger476
}

func TaintStepTest_LogLoggerWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter182` into `intoLogger144`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter182`:
	fromWriter182 := sourceCQL.(io.Writer)

	// Declare `intoLogger144` variable:
	var intoLogger144 log.Logger

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoLogger144`:
	intermediateCQL := intoLogger144.Writer()

	// Extra step (`fromWriter182` taints `intermediateCQL`, which taints `intoLogger144`:
	link(fromWriter182, intermediateCQL)

	// Return the tainted `intoLogger144`:
	return intoLogger144
}

func TaintStepTest_LogLoggerWriter_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromLogger725` into `intoWriter849`.

	// Assume that `sourceCQL` has the underlying type of `fromLogger725`:
	fromLogger725 := sourceCQL.(log.Logger)

	// Call the method that transfers the taint
	// from the receiver `fromLogger725` to the result `intoWriter849`
	// (`intoWriter849` is now tainted).
	intoWriter849 := fromLogger725.Writer()

	// Return the tainted `intoWriter849`:
	return intoWriter849
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
