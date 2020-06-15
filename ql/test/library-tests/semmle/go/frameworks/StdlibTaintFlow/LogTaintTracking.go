package main

import (
	"io"
	"log"
)

func TaintStepTest_LogNew_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromLogger163` into `intoWriter449`.

	// Assume that `sourceCQL` has the underlying type of `fromLogger163`:
	fromLogger163 := sourceCQL.(*log.Logger)

	// Declare `intoWriter449` variable:
	var intoWriter449 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter449`:
	intermediateCQL := log.New(intoWriter449, "", 0)

	// Extra step (`fromLogger163` taints `intermediateCQL`, which taints `intoWriter449`:
	link(fromLogger163, intermediateCQL)

	// Sink the tainted `intoWriter449`:
	sink(intoWriter449)
}

func TaintStepTest_LogLoggerFatal_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface300` into `intoLogger536`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface300`:
	fromInterface300 := sourceCQL.([]interface{})

	// Declare `intoLogger536` variable:
	var intoLogger536 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface300` to the receiver `intoLogger536`
	// (`intoLogger536` is now tainted).
	intoLogger536.Fatal(fromInterface300)

	// Sink the tainted `intoLogger536`:
	sink(intoLogger536)
}

func TaintStepTest_LogLoggerFatalf_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString132` into `intoLogger837`.

	// Assume that `sourceCQL` has the underlying type of `fromString132`:
	fromString132 := sourceCQL.(string)

	// Declare `intoLogger837` variable:
	var intoLogger837 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromString132` to the receiver `intoLogger837`
	// (`intoLogger837` is now tainted).
	intoLogger837.Fatalf(fromString132, nil)

	// Sink the tainted `intoLogger837`:
	sink(intoLogger837)
}

func TaintStepTest_LogLoggerFatalf_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromInterface873` into `intoLogger218`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface873`:
	fromInterface873 := sourceCQL.([]interface{})

	// Declare `intoLogger218` variable:
	var intoLogger218 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface873` to the receiver `intoLogger218`
	// (`intoLogger218` is now tainted).
	intoLogger218.Fatalf("", fromInterface873)

	// Sink the tainted `intoLogger218`:
	sink(intoLogger218)
}

func TaintStepTest_LogLoggerFatalln_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface603` into `intoLogger835`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface603`:
	fromInterface603 := sourceCQL.([]interface{})

	// Declare `intoLogger835` variable:
	var intoLogger835 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface603` to the receiver `intoLogger835`
	// (`intoLogger835` is now tainted).
	intoLogger835.Fatalln(fromInterface603)

	// Sink the tainted `intoLogger835`:
	sink(intoLogger835)
}

func TaintStepTest_LogLoggerPanic_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface639` into `intoLogger291`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface639`:
	fromInterface639 := sourceCQL.([]interface{})

	// Declare `intoLogger291` variable:
	var intoLogger291 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface639` to the receiver `intoLogger291`
	// (`intoLogger291` is now tainted).
	intoLogger291.Panic(fromInterface639)

	// Sink the tainted `intoLogger291`:
	sink(intoLogger291)
}

func TaintStepTest_LogLoggerPanicf_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString893` into `intoLogger710`.

	// Assume that `sourceCQL` has the underlying type of `fromString893`:
	fromString893 := sourceCQL.(string)

	// Declare `intoLogger710` variable:
	var intoLogger710 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromString893` to the receiver `intoLogger710`
	// (`intoLogger710` is now tainted).
	intoLogger710.Panicf(fromString893, nil)

	// Sink the tainted `intoLogger710`:
	sink(intoLogger710)
}

func TaintStepTest_LogLoggerPanicf_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromInterface771` into `intoLogger828`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface771`:
	fromInterface771 := sourceCQL.([]interface{})

	// Declare `intoLogger828` variable:
	var intoLogger828 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface771` to the receiver `intoLogger828`
	// (`intoLogger828` is now tainted).
	intoLogger828.Panicf("", fromInterface771)

	// Sink the tainted `intoLogger828`:
	sink(intoLogger828)
}

func TaintStepTest_LogLoggerPanicln_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface339` into `intoLogger757`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface339`:
	fromInterface339 := sourceCQL.([]interface{})

	// Declare `intoLogger757` variable:
	var intoLogger757 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface339` to the receiver `intoLogger757`
	// (`intoLogger757` is now tainted).
	intoLogger757.Panicln(fromInterface339)

	// Sink the tainted `intoLogger757`:
	sink(intoLogger757)
}

func TaintStepTest_LogLoggerPrint_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface694` into `intoLogger493`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface694`:
	fromInterface694 := sourceCQL.([]interface{})

	// Declare `intoLogger493` variable:
	var intoLogger493 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface694` to the receiver `intoLogger493`
	// (`intoLogger493` is now tainted).
	intoLogger493.Print(fromInterface694)

	// Sink the tainted `intoLogger493`:
	sink(intoLogger493)
}

func TaintStepTest_LogLoggerPrintf_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString572` into `intoLogger892`.

	// Assume that `sourceCQL` has the underlying type of `fromString572`:
	fromString572 := sourceCQL.(string)

	// Declare `intoLogger892` variable:
	var intoLogger892 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromString572` to the receiver `intoLogger892`
	// (`intoLogger892` is now tainted).
	intoLogger892.Printf(fromString572, nil)

	// Sink the tainted `intoLogger892`:
	sink(intoLogger892)
}

func TaintStepTest_LogLoggerPrintf_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromInterface570` into `intoLogger853`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface570`:
	fromInterface570 := sourceCQL.([]interface{})

	// Declare `intoLogger853` variable:
	var intoLogger853 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface570` to the receiver `intoLogger853`
	// (`intoLogger853` is now tainted).
	intoLogger853.Printf("", fromInterface570)

	// Sink the tainted `intoLogger853`:
	sink(intoLogger853)
}

func TaintStepTest_LogLoggerPrintln_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface770` into `intoLogger236`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface770`:
	fromInterface770 := sourceCQL.([]interface{})

	// Declare `intoLogger236` variable:
	var intoLogger236 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromInterface770` to the receiver `intoLogger236`
	// (`intoLogger236` is now tainted).
	intoLogger236.Println(fromInterface770)

	// Sink the tainted `intoLogger236`:
	sink(intoLogger236)
}

func TaintStepTest_LogLoggerSetOutput_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromLogger159` into `intoWriter681`.

	// Assume that `sourceCQL` has the underlying type of `fromLogger159`:
	fromLogger159 := sourceCQL.(log.Logger)

	// Declare `intoWriter681` variable:
	var intoWriter681 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromLogger159` to the argument `intoWriter681`
	// (`intoWriter681` is now tainted).
	fromLogger159.SetOutput(intoWriter681)

	// Sink the tainted `intoWriter681`:
	sink(intoWriter681)
}

func TaintStepTest_LogLoggerSetPrefix_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString194` into `intoLogger953`.

	// Assume that `sourceCQL` has the underlying type of `fromString194`:
	fromString194 := sourceCQL.(string)

	// Declare `intoLogger953` variable:
	var intoLogger953 log.Logger

	// Call the method that transfers the taint
	// from the parameter `fromString194` to the receiver `intoLogger953`
	// (`intoLogger953` is now tainted).
	intoLogger953.SetPrefix(fromString194)

	// Sink the tainted `intoLogger953`:
	sink(intoLogger953)
}

func TaintStepTest_LogLoggerWriter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter736` into `intoLogger987`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter736`:
	fromWriter736 := sourceCQL.(io.Writer)

	// Declare `intoLogger987` variable:
	var intoLogger987 log.Logger

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoLogger987`:
	intermediateCQL := intoLogger987.Writer()

	// Extra step (`fromWriter736` taints `intermediateCQL`, which taints `intoLogger987`:
	link(fromWriter736, intermediateCQL)

	// Sink the tainted `intoLogger987`:
	sink(intoLogger987)
}

func TaintStepTest_LogLoggerWriter_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromLogger715` into `intoWriter300`.

	// Assume that `sourceCQL` has the underlying type of `fromLogger715`:
	fromLogger715 := sourceCQL.(log.Logger)

	// Call the method that transfers the taint
	// from the receiver `fromLogger715` to the result `intoWriter300`
	// (`intoWriter300` is now tainted).
	intoWriter300 := fromLogger715.Writer()

	// Sink the tainted `intoWriter300`:
	sink(intoWriter300)
}

func RunAllTaints_Log() {
	{
		source := newSource()
		TaintStepTest_LogNew_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerFatal_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerFatalf_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerFatalf_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerFatalln_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerPanic_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerPanicf_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerPanicf_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerPanicln_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerPrint_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerPrintf_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerPrintf_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerPrintln_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerSetOutput_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerSetPrefix_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerWriter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_LogLoggerWriter_B1I0O0(source)
	}
}
