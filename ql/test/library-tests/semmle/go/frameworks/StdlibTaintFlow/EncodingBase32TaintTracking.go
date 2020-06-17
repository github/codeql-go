package main

import (
	"encoding/base32"
	"io"
)

func TaintStepTest_EncodingBase32NewDecoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader307` into `intoReader630`.

	// Assume that `sourceCQL` has the underlying type of `fromReader307`:
	fromReader307 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader307` to result `intoReader630`
	// (`intoReader630` is now tainted).
	intoReader630 := base32.NewDecoder(nil, fromReader307)

	// Return the tainted `intoReader630`:
	return intoReader630
}

func TaintStepTest_EncodingBase32NewEncoder_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriteCloser225` into `intoWriter823`.

	// Assume that `sourceCQL` has the underlying type of `fromWriteCloser225`:
	fromWriteCloser225 := sourceCQL.(io.WriteCloser)

	// Declare `intoWriter823` variable:
	var intoWriter823 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter823`:
	intermediateCQL := base32.NewEncoder(nil, intoWriter823)

	// Extra step (`fromWriteCloser225` taints `intermediateCQL`, which taints `intoWriter823`:
	link(fromWriteCloser225, intermediateCQL)

	// Return the tainted `intoWriter823`:
	return intoWriter823
}

func TaintStepTest_EncodingBase32EncodingDecode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte933` into `intoByte756`.

	// Assume that `sourceCQL` has the underlying type of `fromByte933`:
	fromByte933 := sourceCQL.([]byte)

	// Declare `intoByte756` variable:
	var intoByte756 []byte

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte933` to the parameter `intoByte756`
	// (`intoByte756` is now tainted).
	mediumObjCQL.Decode(intoByte756, fromByte933)

	// Return the tainted `intoByte756`:
	return intoByte756
}

func TaintStepTest_EncodingBase32EncodingDecodeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString514` into `intoByte631`.

	// Assume that `sourceCQL` has the underlying type of `fromString514`:
	fromString514 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromString514` to the result `intoByte631`
	// (`intoByte631` is now tainted).
	intoByte631, _ := mediumObjCQL.DecodeString(fromString514)

	// Return the tainted `intoByte631`:
	return intoByte631
}

func TaintStepTest_EncodingBase32EncodingEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte412` into `intoByte778`.

	// Assume that `sourceCQL` has the underlying type of `fromByte412`:
	fromByte412 := sourceCQL.([]byte)

	// Declare `intoByte778` variable:
	var intoByte778 []byte

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte412` to the parameter `intoByte778`
	// (`intoByte778` is now tainted).
	mediumObjCQL.Encode(intoByte778, fromByte412)

	// Return the tainted `intoByte778`:
	return intoByte778
}

func TaintStepTest_EncodingBase32EncodingEncodeToString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte508` into `intoString687`.

	// Assume that `sourceCQL` has the underlying type of `fromByte508`:
	fromByte508 := sourceCQL.([]byte)

	// Declare medium object/interface:
	var mediumObjCQL base32.Encoding

	// Call the method that transfers the taint
	// from the parameter `fromByte508` to the result `intoString687`
	// (`intoString687` is now tainted).
	intoString687 := mediumObjCQL.EncodeToString(fromByte508)

	// Return the tainted `intoString687`:
	return intoString687
}

func RunAllTaints_EncodingBase32() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase32NewDecoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase32NewEncoder_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase32EncodingDecode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase32EncodingDecodeString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase32EncodingEncode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_EncodingBase32EncodingEncodeToString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
