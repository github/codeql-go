package main

import (
	"io"
	"net/mail"
)

func TaintStepTest_NetMailParseAddress_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString853` into `intoAddress437`.

	// Assume that `sourceCQL` has the underlying type of `fromString853`:
	fromString853 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString853` to result `intoAddress437`
	// (`intoAddress437` is now tainted).
	intoAddress437, _ := mail.ParseAddress(fromString853)

	// Return the tainted `intoAddress437`:
	return intoAddress437
}

func TaintStepTest_NetMailParseAddressList_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString653` into `intoAddress475`.

	// Assume that `sourceCQL` has the underlying type of `fromString653`:
	fromString653 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString653` to result `intoAddress475`
	// (`intoAddress475` is now tainted).
	intoAddress475, _ := mail.ParseAddressList(fromString653)

	// Return the tainted `intoAddress475`:
	return intoAddress475
}

func TaintStepTest_NetMailReadMessage_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader164` into `intoMessage464`.

	// Assume that `sourceCQL` has the underlying type of `fromReader164`:
	fromReader164 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader164` to result `intoMessage464`
	// (`intoMessage464` is now tainted).
	intoMessage464, _ := mail.ReadMessage(fromReader164)

	// Return the tainted `intoMessage464`:
	return intoMessage464
}

func TaintStepTest_NetMailAddressParserParse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString445` into `intoAddress182`.

	// Assume that `sourceCQL` has the underlying type of `fromString445`:
	fromString445 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mail.AddressParser

	// Call the method that transfers the taint
	// from the parameter `fromString445` to the result `intoAddress182`
	// (`intoAddress182` is now tainted).
	intoAddress182, _ := mediumObjCQL.Parse(fromString445)

	// Return the tainted `intoAddress182`:
	return intoAddress182
}

func TaintStepTest_NetMailAddressParserParseList_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString759` into `intoAddress288`.

	// Assume that `sourceCQL` has the underlying type of `fromString759`:
	fromString759 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL mail.AddressParser

	// Call the method that transfers the taint
	// from the parameter `fromString759` to the result `intoAddress288`
	// (`intoAddress288` is now tainted).
	intoAddress288, _ := mediumObjCQL.ParseList(fromString759)

	// Return the tainted `intoAddress288`:
	return intoAddress288
}

func TaintStepTest_NetMailHeaderGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader667` into `intoString477`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader667`:
	fromHeader667 := sourceCQL.(mail.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader667` to the result `intoString477`
	// (`intoString477` is now tainted).
	intoString477 := fromHeader667.Get("")

	// Return the tainted `intoString477`:
	return intoString477
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
