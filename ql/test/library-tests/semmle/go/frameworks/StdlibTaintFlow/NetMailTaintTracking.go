package main

import (
	"io"
	"net/mail"
)

func TaintStepTest_NetMailParseAddress_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString673` into `intoAddress485`.

	// Assume that `sourceCQL` has the underlying type of `fromString673`:
	fromString673 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString673` to result `intoAddress485`
	// (`intoAddress485` is now tainted).
	intoAddress485, _ := mail.ParseAddress(fromString673)

	// Sink the tainted `intoAddress485`:
	sink(intoAddress485)
}

func TaintStepTest_NetMailParseAddressList_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString797` into `intoAddress889`.

	// Assume that `sourceCQL` has the underlying type of `fromString797`:
	fromString797 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString797` to result `intoAddress889`
	// (`intoAddress889` is now tainted).
	intoAddress889, _ := mail.ParseAddressList(fromString797)

	// Sink the tainted `intoAddress889`:
	sink(intoAddress889)
}

func TaintStepTest_NetMailReadMessage_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader677` into `intoMessage855`.

	// Assume that `sourceCQL` has the underlying type of `fromReader677`:
	fromReader677 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader677` to result `intoMessage855`
	// (`intoMessage855` is now tainted).
	intoMessage855, _ := mail.ReadMessage(fromReader677)

	// Sink the tainted `intoMessage855`:
	sink(intoMessage855)
}

func TaintStepTest_NetMailAddressParserParse_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString458` into `intoAddress866`.

	// Assume that `sourceCQL` has the underlying type of `fromString458`:
	fromString458 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mail.AddressParser

	// Call the method that transfers the taint
	// from the parameter `fromString458` to the result `intoAddress866`
	// (`intoAddress866` is now tainted).
	intoAddress866, _ := mediumObjCQL.Parse(fromString458)

	// Sink the tainted `intoAddress866`:
	sink(intoAddress866)
}

func TaintStepTest_NetMailAddressParserParseList_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString515` into `intoAddress683`.

	// Assume that `sourceCQL` has the underlying type of `fromString515`:
	fromString515 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mail.AddressParser

	// Call the method that transfers the taint
	// from the parameter `fromString515` to the result `intoAddress683`
	// (`intoAddress683` is now tainted).
	intoAddress683, _ := mediumObjCQL.ParseList(fromString515)

	// Sink the tainted `intoAddress683`:
	sink(intoAddress683)
}

func TaintStepTest_NetMailHeaderGet_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromHeader206` into `intoString258`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader206`:
	fromHeader206 := sourceCQL.(mail.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader206` to the result `intoString258`
	// (`intoString258` is now tainted).
	intoString258 := fromHeader206.Get("")

	// Sink the tainted `intoString258`:
	sink(intoString258)
}

func RunAllTaints_NetMail() {
	{
		source := newSource()
		TaintStepTest_NetMailParseAddress_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetMailParseAddressList_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetMailReadMessage_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetMailAddressParserParse_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetMailAddressParserParseList_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetMailHeaderGet_B0I0O0(source)
	}
}
