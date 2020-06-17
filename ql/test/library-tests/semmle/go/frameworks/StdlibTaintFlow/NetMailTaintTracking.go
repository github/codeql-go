// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"net/mail"
)

func TaintStepTest_NetMailParseAddress_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString656` into `intoAddress414`.

	// Assume that `sourceCQL` has the underlying type of `fromString656`:
	fromString656 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString656` to result `intoAddress414`
	// (`intoAddress414` is now tainted).
	intoAddress414, _ := mail.ParseAddress(fromString656)

	// Return the tainted `intoAddress414`:
	return intoAddress414
}

func TaintStepTest_NetMailParseAddressList_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString518` into `intoAddress650`.

	// Assume that `sourceCQL` has the underlying type of `fromString518`:
	fromString518 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString518` to result `intoAddress650`
	// (`intoAddress650` is now tainted).
	intoAddress650, _ := mail.ParseAddressList(fromString518)

	// Return the tainted `intoAddress650`:
	return intoAddress650
}

func TaintStepTest_NetMailReadMessage_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader784` into `intoMessage957`.

	// Assume that `sourceCQL` has the underlying type of `fromReader784`:
	fromReader784 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader784` to result `intoMessage957`
	// (`intoMessage957` is now tainted).
	intoMessage957, _ := mail.ReadMessage(fromReader784)

	// Return the tainted `intoMessage957`:
	return intoMessage957
}

func TaintStepTest_NetMailAddressParserParse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString520` into `intoAddress443`.

	// Assume that `sourceCQL` has the underlying type of `fromString520`:
	fromString520 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mail.AddressParser

	// Call the method that transfers the taint
	// from the parameter `fromString520` to the result `intoAddress443`
	// (`intoAddress443` is now tainted).
	intoAddress443, _ := mediumObjCQL.Parse(fromString520)

	// Return the tainted `intoAddress443`:
	return intoAddress443
}

func TaintStepTest_NetMailAddressParserParseList_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString127` into `intoAddress483`.

	// Assume that `sourceCQL` has the underlying type of `fromString127`:
	fromString127 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mail.AddressParser

	// Call the method that transfers the taint
	// from the parameter `fromString127` to the result `intoAddress483`
	// (`intoAddress483` is now tainted).
	intoAddress483, _ := mediumObjCQL.ParseList(fromString127)

	// Return the tainted `intoAddress483`:
	return intoAddress483
}

func TaintStepTest_NetMailHeaderGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader989` into `intoString982`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader989`:
	fromHeader989 := sourceCQL.(mail.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader989` to the result `intoString982`
	// (`intoString982` is now tainted).
	intoString982 := fromHeader989.Get("")

	// Return the tainted `intoString982`:
	return intoString982
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
