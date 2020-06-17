// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"log"
)

func TaintStepTest_LogNew_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromLogger656` into `intoWriter414`.

	// Assume that `sourceCQL` has the underlying type of `fromLogger656`:
	fromLogger656 := sourceCQL.(*log.Logger)

	// Declare `intoWriter414` variable:
	var intoWriter414 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter414`:
	intermediateCQL := log.New(intoWriter414, "", 0)

	// Extra step (`fromLogger656` taints `intermediateCQL`, which taints `intoWriter414`:
	link(fromLogger656, intermediateCQL)

	// Return the tainted `intoWriter414`:
	return intoWriter414
}

func TaintStepTest_LogLoggerFatal_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface518` into `intoLogger650`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface518`:
	fromInterface518 := sourceCQL.(interface{})

	// Declare `intoLogger650` variable:
	var intoLogger650 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface518` to the receiver `intoLogger650`
	// (`intoLogger650` is now tainted).
	intoLogger650.Fatal(fromInterface518)

	// Return the tainted `intoLogger650`:
	return intoLogger650
}

func TaintStepTest_LogLoggerFatalf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString784` into `intoLogger957`.

	// Assume that `sourceCQL` has the underlying type of `fromString784`:
	fromString784 := sourceCQL.(string)

	// Declare `intoLogger957` variable:
	var intoLogger957 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromString784` to the receiver `intoLogger957`
	// (`intoLogger957` is now tainted).
	intoLogger957.Fatalf(fromString784, nil)

	// Return the tainted `intoLogger957`:
	return intoLogger957
}

func TaintStepTest_LogLoggerFatalf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface520` into `intoLogger443`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface520`:
	fromInterface520 := sourceCQL.(interface{})

	// Declare `intoLogger443` variable:
	var intoLogger443 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface520` to the receiver `intoLogger443`
	// (`intoLogger443` is now tainted).
	intoLogger443.Fatalf("", fromInterface520)

	// Return the tainted `intoLogger443`:
	return intoLogger443
}

func TaintStepTest_LogLoggerFatalln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface127` into `intoLogger483`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface127`:
	fromInterface127 := sourceCQL.(interface{})

	// Declare `intoLogger483` variable:
	var intoLogger483 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface127` to the receiver `intoLogger483`
	// (`intoLogger483` is now tainted).
	intoLogger483.Fatalln(fromInterface127)

	// Return the tainted `intoLogger483`:
	return intoLogger483
}

func TaintStepTest_LogLoggerPanic_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface989` into `intoLogger982`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface989`:
	fromInterface989 := sourceCQL.(interface{})

	// Declare `intoLogger982` variable:
	var intoLogger982 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface989` to the receiver `intoLogger982`
	// (`intoLogger982` is now tainted).
	intoLogger982.Panic(fromInterface989)

	// Return the tainted `intoLogger982`:
	return intoLogger982
}

func TaintStepTest_LogLoggerPanicf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString417` into `intoLogger584`.

	// Assume that `sourceCQL` has the underlying type of `fromString417`:
	fromString417 := sourceCQL.(string)

	// Declare `intoLogger584` variable:
	var intoLogger584 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromString417` to the receiver `intoLogger584`
	// (`intoLogger584` is now tainted).
	intoLogger584.Panicf(fromString417, nil)

	// Return the tainted `intoLogger584`:
	return intoLogger584
}

func TaintStepTest_LogLoggerPanicf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface991` into `intoLogger881`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface991`:
	fromInterface991 := sourceCQL.(interface{})

	// Declare `intoLogger881` variable:
	var intoLogger881 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface991` to the receiver `intoLogger881`
	// (`intoLogger881` is now tainted).
	intoLogger881.Panicf("", fromInterface991)

	// Return the tainted `intoLogger881`:
	return intoLogger881
}

func TaintStepTest_LogLoggerPanicln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface186` into `intoLogger284`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface186`:
	fromInterface186 := sourceCQL.(interface{})

	// Declare `intoLogger284` variable:
	var intoLogger284 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface186` to the receiver `intoLogger284`
	// (`intoLogger284` is now tainted).
	intoLogger284.Panicln(fromInterface186)

	// Return the tainted `intoLogger284`:
	return intoLogger284
}

func TaintStepTest_LogLoggerPrint_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface908` into `intoLogger137`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface908`:
	fromInterface908 := sourceCQL.(interface{})

	// Declare `intoLogger137` variable:
	var intoLogger137 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface908` to the receiver `intoLogger137`
	// (`intoLogger137` is now tainted).
	intoLogger137.Print(fromInterface908)

	// Return the tainted `intoLogger137`:
	return intoLogger137
}

func TaintStepTest_LogLoggerPrintf_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString494` into `intoLogger873`.

	// Assume that `sourceCQL` has the underlying type of `fromString494`:
	fromString494 := sourceCQL.(string)

	// Declare `intoLogger873` variable:
	var intoLogger873 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromString494` to the receiver `intoLogger873`
	// (`intoLogger873` is now tainted).
	intoLogger873.Printf(fromString494, nil)

	// Return the tainted `intoLogger873`:
	return intoLogger873
}

func TaintStepTest_LogLoggerPrintf_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface599` into `intoLogger409`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface599`:
	fromInterface599 := sourceCQL.(interface{})

	// Declare `intoLogger409` variable:
	var intoLogger409 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface599` to the receiver `intoLogger409`
	// (`intoLogger409` is now tainted).
	intoLogger409.Printf("", fromInterface599)

	// Return the tainted `intoLogger409`:
	return intoLogger409
}

func TaintStepTest_LogLoggerPrintln_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface246` into `intoLogger898`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface246`:
	fromInterface246 := sourceCQL.(interface{})

	// Declare `intoLogger898` variable:
	var intoLogger898 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface246` to the receiver `intoLogger898`
	// (`intoLogger898` is now tainted).
	intoLogger898.Println(fromInterface246)

	// Return the tainted `intoLogger898`:
	return intoLogger898
}

func TaintStepTest_LogLoggerSetOutput_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromLogger598` into `intoWriter631`.

	// Assume that `sourceCQL` has the underlying type of `fromLogger598`:
	fromLogger598 := sourceCQL.(log.Logger)

	// Declare `intoWriter631` variable:
	var intoWriter631 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromLogger598` to the argument `intoWriter631`
	// (`intoWriter631` is now tainted).
	fromLogger598.SetOutput(intoWriter631)

	// Return the tainted `intoWriter631`:
	return intoWriter631
}

func TaintStepTest_LogLoggerSetPrefix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString165` into `intoLogger150`.

	// Assume that `sourceCQL` has the underlying type of `fromString165`:
	fromString165 := sourceCQL.(string)

	// Declare `intoLogger150` variable:
	var intoLogger150 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromString165` to the receiver `intoLogger150`
	// (`intoLogger150` is now tainted).
	intoLogger150.SetPrefix(fromString165)

	// Return the tainted `intoLogger150`:
	return intoLogger150
}

func TaintStepTest_LogLoggerWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter340` into `intoLogger471`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter340`:
	fromWriter340 := sourceCQL.(io.Writer)

	// Declare `intoLogger471` variable:
	var intoLogger471 log.Logger

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoLogger471`:
	intermediateCQL := intoLogger471.Writer()

	// Extra step (`fromWriter340` taints `intermediateCQL`, which taints `intoLogger471`:
	link(fromWriter340, intermediateCQL)

	// Return the tainted `intoLogger471`:
	return intoLogger471
}

func TaintStepTest_LogLoggerWriter_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromLogger290` into `intoWriter758`.

	// Assume that `sourceCQL` has the underlying type of `fromLogger290`:
	fromLogger290 := sourceCQL.(log.Logger)

	// Call the method that transfers the taint
	// from the receiver `fromLogger290` to the result `intoWriter758`
	// (`intoWriter758` is now tainted).
	intoWriter758 := fromLogger290.Writer()

	// Return the tainted `intoWriter758`:
	return intoWriter758
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
