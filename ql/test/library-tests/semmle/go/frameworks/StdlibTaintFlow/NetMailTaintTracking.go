// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"net/mail"
)

func TaintStepTest_NetMailParseAddress_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString869` into `intoAddress222`.

	// Assume that `sourceCQL` has the underlying type of `fromString869`:
	fromString869 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString869` to result `intoAddress222`
	// (`intoAddress222` is now tainted).
	intoAddress222, _ := mail.ParseAddress(fromString869)

	// Return the tainted `intoAddress222`:
	return intoAddress222
}

func TaintStepTest_NetMailParseAddressList_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString518` into `intoAddress228`.

	// Assume that `sourceCQL` has the underlying type of `fromString518`:
	fromString518 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString518` to result `intoAddress228`
	// (`intoAddress228` is now tainted).
	intoAddress228, _ := mail.ParseAddressList(fromString518)

	// Return the tainted `intoAddress228`:
	return intoAddress228
}

func TaintStepTest_NetMailReadMessage_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader465` into `intoMessage496`.

	// Assume that `sourceCQL` has the underlying type of `fromReader465`:
	fromReader465 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader465` to result `intoMessage496`
	// (`intoMessage496` is now tainted).
	intoMessage496, _ := mail.ReadMessage(fromReader465)

	// Return the tainted `intoMessage496`:
	return intoMessage496
}

func TaintStepTest_NetMailAddressParserParse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString859` into `intoAddress877`.

	// Assume that `sourceCQL` has the underlying type of `fromString859`:
	fromString859 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mail.AddressParser

	// Call the method that transfers the taint
	// from the parameter `fromString859` to the result `intoAddress877`
	// (`intoAddress877` is now tainted).
	intoAddress877, _ := mediumObjCQL.Parse(fromString859)

	// Return the tainted `intoAddress877`:
	return intoAddress877
}

func TaintStepTest_NetMailAddressParserParseList_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString814` into `intoAddress529`.

	// Assume that `sourceCQL` has the underlying type of `fromString814`:
	fromString814 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mail.AddressParser

	// Call the method that transfers the taint
	// from the parameter `fromString814` to the result `intoAddress529`
	// (`intoAddress529` is now tainted).
	intoAddress529, _ := mediumObjCQL.ParseList(fromString814)

	// Return the tainted `intoAddress529`:
	return intoAddress529
}

func TaintStepTest_NetMailHeaderGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader646` into `intoString366`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader646`:
	fromHeader646 := sourceCQL.(mail.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader646` to the result `intoString366`
	// (`intoString366` is now tainted).
	intoString366 := fromHeader646.Get("")

	// Return the tainted `intoString366`:
	return intoString366
}

func RunAllTaints_NetMail() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetMailParseAddress_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetMailParseAddressList_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetMailReadMessage_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetMailAddressParserParse_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetMailAddressParserParseList_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetMailHeaderGet_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
