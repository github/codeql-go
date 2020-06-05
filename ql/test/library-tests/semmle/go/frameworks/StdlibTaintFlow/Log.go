package main

import (
	"io"
	"log"
)

func TaintStepTest_LogNew(sourceCQL interface{}) {
	// The flow is from `from569` into `out`.

	// Assume that `sourceCQL` has the underlying type of `from569`:
	from569 := sourceCQL.(*log.Logger)

	// Declare `out` variable:
	var out io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `out`:
	intermediateCQL := log.New(out, "", 0)

	// Extra step (`from569` taints `intermediateCQL`, which taints `out`:
	link(from569, intermediateCQL)

	// Sink the tainted `out`:
	sink(out)
}

func TaintStepTest_LogLoggerFatal(sourceCQL interface{}) {
	// The flow is from `v` into `from588`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.([]interface{})

	// Declare `from588` variable:
	var from588 log.Logger

	// Call the method that transfers the taint
	// from the parameter `v` to the receiver `from588`
	// (`from588` is now tainted).
	from588.Fatal(v)

	// Sink the tainted `from588`:
	sink(from588)
}

func TaintStepTest_LogLoggerFatalf(sourceCQL interface{}) {
	// The flow is from `v` into `from138`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.([]interface{})

	// Declare `from138` variable:
	var from138 log.Logger

	// Call the method that transfers the taint
	// from the parameter `v` to the receiver `from138`
	// (`from138` is now tainted).
	from138.Fatalf("", v)

	// Sink the tainted `from138`:
	sink(from138)
}

func TaintStepTest_LogLoggerFatalln(sourceCQL interface{}) {
	// The flow is from `v` into `into390`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.([]interface{})

	// Declare `into390` variable:
	var into390 log.Logger

	// Call the method that transfers the taint
	// from the parameter `v` to the receiver `into390`
	// (`into390` is now tainted).
	into390.Fatalln(v)

	// Sink the tainted `into390`:
	sink(into390)
}

func TaintStepTest_LogLoggerPanic(sourceCQL interface{}) {
	// The flow is from `v` into `into349`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.([]interface{})

	// Declare `into349` variable:
	var into349 log.Logger

	// Call the method that transfers the taint
	// from the parameter `v` to the receiver `into349`
	// (`into349` is now tainted).
	into349.Panic(v)

	// Sink the tainted `into349`:
	sink(into349)
}

func TaintStepTest_LogLoggerPanicf(sourceCQL interface{}) {
	// The flow is from `v` into `into819`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.([]interface{})

	// Declare `into819` variable:
	var into819 log.Logger

	// Call the method that transfers the taint
	// from the parameter `v` to the receiver `into819`
	// (`into819` is now tainted).
	into819.Panicf("", v)

	// Sink the tainted `into819`:
	sink(into819)
}

func TaintStepTest_LogLoggerPanicln(sourceCQL interface{}) {
	// The flow is from `v` into `into549`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.([]interface{})

	// Declare `into549` variable:
	var into549 log.Logger

	// Call the method that transfers the taint
	// from the parameter `v` to the receiver `into549`
	// (`into549` is now tainted).
	into549.Panicln(v)

	// Sink the tainted `into549`:
	sink(into549)
}

func TaintStepTest_LogLoggerPrint(sourceCQL interface{}) {
	// The flow is from `v` into `into288`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.([]interface{})

	// Declare `into288` variable:
	var into288 log.Logger

	// Call the method that transfers the taint
	// from the parameter `v` to the receiver `into288`
	// (`into288` is now tainted).
	into288.Print(v)

	// Sink the tainted `into288`:
	sink(into288)
}

func TaintStepTest_LogLoggerPrintf(sourceCQL interface{}) {
	// The flow is from `v` into `into545`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.([]interface{})

	// Declare `into545` variable:
	var into545 log.Logger

	// Call the method that transfers the taint
	// from the parameter `v` to the receiver `into545`
	// (`into545` is now tainted).
	into545.Printf("", v)

	// Sink the tainted `into545`:
	sink(into545)
}

func TaintStepTest_LogLoggerPrintln(sourceCQL interface{}) {
	// The flow is from `v` into `into329`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.([]interface{})

	// Declare `into329` variable:
	var into329 log.Logger

	// Call the method that transfers the taint
	// from the parameter `v` to the receiver `into329`
	// (`into329` is now tainted).
	into329.Println(v)

	// Sink the tainted `into329`:
	sink(into329)
}

func TaintStepTest_LogLoggerSetOutput(sourceCQL interface{}) {
	// The flow is from `from685` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from685`:
	from685 := sourceCQL.(log.Logger)

	// Declare `w` variable:
	var w io.Writer

	// Call the method that transfers the taint
	// from the receiver `from685` to the argument `w`
	// (`w` is now tainted).
	from685.SetOutput(w)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_LogLoggerSetPrefix(sourceCQL interface{}) {
	// The flow is from `prefix` into `into519`.

	// Assume that `sourceCQL` has the underlying type of `prefix`:
	prefix := sourceCQL.(string)

	// Declare `into519` variable:
	var into519 log.Logger

	// Call the method that transfers the taint
	// from the parameter `prefix` to the receiver `into519`
	// (`into519` is now tainted).
	into519.SetPrefix(prefix)

	// Sink the tainted `into519`:
	sink(into519)
}

func TaintStepTest_LogLoggerWriter(sourceCQL interface{}) {
	// The flow is from `from859` into `into534`.

	// Assume that `sourceCQL` has the underlying type of `from859`:
	from859 := sourceCQL.(log.Logger)

	// Call the method that transfers the taint
	// from the receiver `from859` to the result `into534`
	// (`into534` is now tainted).
	into534 := from859.Writer()

	// Sink the tainted `into534`:
	sink(into534)
}

func RunAllTaints_Log(v interface{}) {
	{
		source := newSource()
		TaintStepTest_LogNew(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerFatal(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerFatalf(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerFatalln(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerPanic(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerPanicf(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerPanicln(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerPrint(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerPrintf(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerPrintln(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerSetOutput(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerSetPrefix(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerWriter(source)
	}
}
