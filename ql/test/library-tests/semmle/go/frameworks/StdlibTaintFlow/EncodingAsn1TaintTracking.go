package main

import "encoding/asn1"

func TaintStepTest_EncodingAsn1Marshal_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface970` into `intoByte681`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface970`:
	fromInterface970 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface970` to result `intoByte681`
	// (`intoByte681` is now tainted).
	intoByte681, _ := asn1.Marshal(fromInterface970)

	// Sink the tainted `intoByte681`:
	sink(intoByte681)
}

func TaintStepTest_EncodingAsn1MarshalWithParams_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface539` into `intoByte154`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface539`:
	fromInterface539 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface539` to result `intoByte154`
	// (`intoByte154` is now tainted).
	intoByte154, _ := asn1.MarshalWithParams(fromInterface539, "")

	// Sink the tainted `intoByte154`:
	sink(intoByte154)
}

func TaintStepTest_EncodingAsn1MarshalWithParams_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString161` into `intoByte953`.

	// Assume that `sourceCQL` has the underlying type of `fromString161`:
	fromString161 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString161` to result `intoByte953`
	// (`intoByte953` is now tainted).
	intoByte953, _ := asn1.MarshalWithParams(nil, fromString161)

	// Sink the tainted `intoByte953`:
	sink(intoByte953)
}

func TaintStepTest_EncodingAsn1Unmarshal_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte430` into `intoInterface401`.

	// Assume that `sourceCQL` has the underlying type of `fromByte430`:
	fromByte430 := sourceCQL.([]byte)

	// Declare `intoInterface401` variable:
	var intoInterface401 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromByte430` to parameter `intoInterface401`;
	// `intoInterface401` is now tainted.
	asn1.Unmarshal(fromByte430, intoInterface401)

	// Sink the tainted `intoInterface401`:
	sink(intoInterface401)
}

func TaintStepTest_EncodingAsn1Unmarshal_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromByte949` into `intoByte692`.

	// Assume that `sourceCQL` has the underlying type of `fromByte949`:
	fromByte949 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte949` to result `intoByte692`
	// (`intoByte692` is now tainted).
	intoByte692, _ := asn1.Unmarshal(fromByte949, nil)

	// Sink the tainted `intoByte692`:
	sink(intoByte692)
}

func TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte527` into `intoInterface214`.

	// Assume that `sourceCQL` has the underlying type of `fromByte527`:
	fromByte527 := sourceCQL.([]byte)

	// Declare `intoInterface214` variable:
	var intoInterface214 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromByte527` to parameter `intoInterface214`;
	// `intoInterface214` is now tainted.
	asn1.UnmarshalWithParams(fromByte527, intoInterface214, "")

	// Sink the tainted `intoInterface214`:
	sink(intoInterface214)
}

func TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromByte414` into `intoByte247`.

	// Assume that `sourceCQL` has the underlying type of `fromByte414`:
	fromByte414 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte414` to result `intoByte247`
	// (`intoByte247` is now tainted).
	intoByte247, _ := asn1.UnmarshalWithParams(fromByte414, nil, "")

	// Sink the tainted `intoByte247`:
	sink(intoByte247)
}

func TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString788` into `intoInterface731`.

	// Assume that `sourceCQL` has the underlying type of `fromString788`:
	fromString788 := sourceCQL.(string)

	// Declare `intoInterface731` variable:
	var intoInterface731 interface{}

	// Call the function that transfers the taint
	// from the parameter `fromString788` to parameter `intoInterface731`;
	// `intoInterface731` is now tainted.
	asn1.UnmarshalWithParams(nil, intoInterface731, fromString788)

	// Sink the tainted `intoInterface731`:
	sink(intoInterface731)
}

func TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I1O1(sourceCQL interface{}) {
	// The flow is from `fromString145` into `intoByte251`.

	// Assume that `sourceCQL` has the underlying type of `fromString145`:
	fromString145 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString145` to result `intoByte251`
	// (`intoByte251` is now tainted).
	intoByte251, _ := asn1.UnmarshalWithParams(nil, nil, fromString145)

	// Sink the tainted `intoByte251`:
	sink(intoByte251)
}

func TaintStepTest_EncodingAsn1ObjectIdentifierString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromObjectIdentifier556` into `intoString157`.

	// Assume that `sourceCQL` has the underlying type of `fromObjectIdentifier556`:
	fromObjectIdentifier556 := sourceCQL.(asn1.ObjectIdentifier)

	// Call the method that transfers the taint
	// from the receiver `fromObjectIdentifier556` to the result `intoString157`
	// (`intoString157` is now tainted).
	intoString157 := fromObjectIdentifier556.String()

	// Sink the tainted `intoString157`:
	sink(intoString157)
}

func RunAllTaints_EncodingAsn1() {
	{
		source := newSource()
		TaintStepTest_EncodingAsn1Marshal_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAsn1MarshalWithParams_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAsn1MarshalWithParams_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAsn1Unmarshal_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAsn1Unmarshal_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAsn1UnmarshalWithParams_B0I1O1(source)
	}
	{
		source := newSource()
		TaintStepTest_EncodingAsn1ObjectIdentifierString_B0I0O0(source)
	}
}
