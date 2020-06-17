// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "net/url"

func TaintStepTest_NetUrlParse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString656` into `intoURL414`.

	// Assume that `sourceCQL` has the underlying type of `fromString656`:
	fromString656 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString656` to result `intoURL414`
	// (`intoURL414` is now tainted).
	intoURL414, _ := url.Parse(fromString656)

	// Return the tainted `intoURL414`:
	return intoURL414
}

func TaintStepTest_NetUrlParseQuery_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString518` into `intoValues650`.

	// Assume that `sourceCQL` has the underlying type of `fromString518`:
	fromString518 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString518` to result `intoValues650`
	// (`intoValues650` is now tainted).
	intoValues650, _ := url.ParseQuery(fromString518)

	// Return the tainted `intoValues650`:
	return intoValues650
}

func TaintStepTest_NetUrlParseRequestURI_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString784` into `intoURL957`.

	// Assume that `sourceCQL` has the underlying type of `fromString784`:
	fromString784 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString784` to result `intoURL957`
	// (`intoURL957` is now tainted).
	intoURL957, _ := url.ParseRequestURI(fromString784)

	// Return the tainted `intoURL957`:
	return intoURL957
}

func TaintStepTest_NetUrlPathEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString520` into `intoString443`.

	// Assume that `sourceCQL` has the underlying type of `fromString520`:
	fromString520 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString520` to result `intoString443`
	// (`intoString443` is now tainted).
	intoString443 := url.PathEscape(fromString520)

	// Return the tainted `intoString443`:
	return intoString443
}

func TaintStepTest_NetUrlPathUnescape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString127` into `intoString483`.

	// Assume that `sourceCQL` has the underlying type of `fromString127`:
	fromString127 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString127` to result `intoString483`
	// (`intoString483` is now tainted).
	intoString483, _ := url.PathUnescape(fromString127)

	// Return the tainted `intoString483`:
	return intoString483
}

func TaintStepTest_NetUrlQueryEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString989` into `intoString982`.

	// Assume that `sourceCQL` has the underlying type of `fromString989`:
	fromString989 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString989` to result `intoString982`
	// (`intoString982` is now tainted).
	intoString982 := url.QueryEscape(fromString989)

	// Return the tainted `intoString982`:
	return intoString982
}

func TaintStepTest_NetUrlQueryUnescape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString417` into `intoString584`.

	// Assume that `sourceCQL` has the underlying type of `fromString417`:
	fromString417 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString417` to result `intoString584`
	// (`intoString584` is now tainted).
	intoString584, _ := url.QueryUnescape(fromString417)

	// Return the tainted `intoString584`:
	return intoString584
}

func TaintStepTest_NetUrlUser_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString991` into `intoUserinfo881`.

	// Assume that `sourceCQL` has the underlying type of `fromString991`:
	fromString991 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString991` to result `intoUserinfo881`
	// (`intoUserinfo881` is now tainted).
	intoUserinfo881 := url.User(fromString991)

	// Return the tainted `intoUserinfo881`:
	return intoUserinfo881
}

func TaintStepTest_NetUrlUserPassword_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString186` into `intoUserinfo284`.

	// Assume that `sourceCQL` has the underlying type of `fromString186`:
	fromString186 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString186` to result `intoUserinfo284`
	// (`intoUserinfo284` is now tainted).
	intoUserinfo284 := url.UserPassword(fromString186, "")

	// Return the tainted `intoUserinfo284`:
	return intoUserinfo284
}

func TaintStepTest_NetUrlUserPassword_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString908` into `intoUserinfo137`.

	// Assume that `sourceCQL` has the underlying type of `fromString908`:
	fromString908 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString908` to result `intoUserinfo137`
	// (`intoUserinfo137` is now tainted).
	intoUserinfo137 := url.UserPassword("", fromString908)

	// Return the tainted `intoUserinfo137`:
	return intoUserinfo137
}

func TaintStepTest_NetUrlURLEscapedPath_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL494` into `intoString873`.

	// Assume that `sourceCQL` has the underlying type of `fromURL494`:
	fromURL494 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL494` to the result `intoString873`
	// (`intoString873` is now tainted).
	intoString873 := fromURL494.EscapedPath()

	// Return the tainted `intoString873`:
	return intoString873
}

func TaintStepTest_NetUrlURLHostname_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL599` into `intoString409`.

	// Assume that `sourceCQL` has the underlying type of `fromURL599`:
	fromURL599 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL599` to the result `intoString409`
	// (`intoString409` is now tainted).
	intoString409 := fromURL599.Hostname()

	// Return the tainted `intoString409`:
	return intoString409
}

func TaintStepTest_NetUrlURLMarshalBinary_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL246` into `intoByte898`.

	// Assume that `sourceCQL` has the underlying type of `fromURL246`:
	fromURL246 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL246` to the result `intoByte898`
	// (`intoByte898` is now tainted).
	intoByte898, _ := fromURL246.MarshalBinary()

	// Return the tainted `intoByte898`:
	return intoByte898
}

func TaintStepTest_NetUrlURLParse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL598` into `intoURL631`.

	// Assume that `sourceCQL` has the underlying type of `fromURL598`:
	fromURL598 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL598` to the result `intoURL631`
	// (`intoURL631` is now tainted).
	intoURL631, _ := fromURL598.Parse("")

	// Return the tainted `intoURL631`:
	return intoURL631
}

func TaintStepTest_NetUrlURLParse_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString165` into `intoURL150`.

	// Assume that `sourceCQL` has the underlying type of `fromString165`:
	fromString165 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL url.URL

	// Call the method that transfers the taint
	// from the parameter `fromString165` to the result `intoURL150`
	// (`intoURL150` is now tainted).
	intoURL150, _ := mediumObjCQL.Parse(fromString165)

	// Return the tainted `intoURL150`:
	return intoURL150
}

func TaintStepTest_NetUrlURLPort_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL340` into `intoString471`.

	// Assume that `sourceCQL` has the underlying type of `fromURL340`:
	fromURL340 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL340` to the result `intoString471`
	// (`intoString471` is now tainted).
	intoString471 := fromURL340.Port()

	// Return the tainted `intoString471`:
	return intoString471
}

func TaintStepTest_NetUrlURLQuery_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL290` into `intoValues758`.

	// Assume that `sourceCQL` has the underlying type of `fromURL290`:
	fromURL290 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL290` to the result `intoValues758`
	// (`intoValues758` is now tainted).
	intoValues758 := fromURL290.Query()

	// Return the tainted `intoValues758`:
	return intoValues758
}

func TaintStepTest_NetUrlURLRequestURI_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL396` into `intoString707`.

	// Assume that `sourceCQL` has the underlying type of `fromURL396`:
	fromURL396 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL396` to the result `intoString707`
	// (`intoString707` is now tainted).
	intoString707 := fromURL396.RequestURI()

	// Return the tainted `intoString707`:
	return intoString707
}

func TaintStepTest_NetUrlURLResolveReference_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL912` into `intoURL718`.

	// Assume that `sourceCQL` has the underlying type of `fromURL912`:
	fromURL912 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL912` to the result `intoURL718`
	// (`intoURL718` is now tainted).
	intoURL718 := fromURL912.ResolveReference(nil)

	// Return the tainted `intoURL718`:
	return intoURL718
}

func TaintStepTest_NetUrlURLResolveReference_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL972` into `intoURL633`.

	// Assume that `sourceCQL` has the underlying type of `fromURL972`:
	fromURL972 := sourceCQL.(*url.URL)

	// Declare medium object/interface:
	var mediumObjCQL url.URL

	// Call the method that transfers the taint
	// from the parameter `fromURL972` to the result `intoURL633`
	// (`intoURL633` is now tainted).
	intoURL633 := mediumObjCQL.ResolveReference(fromURL972)

	// Return the tainted `intoURL633`:
	return intoURL633
}

func TaintStepTest_NetUrlURLString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL316` into `intoString145`.

	// Assume that `sourceCQL` has the underlying type of `fromURL316`:
	fromURL316 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL316` to the result `intoString145`
	// (`intoString145` is now tainted).
	intoString145 := fromURL316.String()

	// Return the tainted `intoString145`:
	return intoString145
}

func TaintStepTest_NetUrlURLUnmarshalBinary_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte817` into `intoURL474`.

	// Assume that `sourceCQL` has the underlying type of `fromByte817`:
	fromByte817 := sourceCQL.([]byte)

	// Declare `intoURL474` variable:
	var intoURL474 url.URL

	// Call the method that transfers the taint
	// from the parameter `fromByte817` to the receiver `intoURL474`
	// (`intoURL474` is now tainted).
	intoURL474.UnmarshalBinary(fromByte817)

	// Return the tainted `intoURL474`:
	return intoURL474
}

func TaintStepTest_NetUrlUserinfoPassword_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUserinfo832` into `intoString378`.

	// Assume that `sourceCQL` has the underlying type of `fromUserinfo832`:
	fromUserinfo832 := sourceCQL.(url.Userinfo)

	// Call the method that transfers the taint
	// from the receiver `fromUserinfo832` to the result `intoString378`
	// (`intoString378` is now tainted).
	intoString378, _ := fromUserinfo832.Password()

	// Return the tainted `intoString378`:
	return intoString378
}

func TaintStepTest_NetUrlUserinfoString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUserinfo541` into `intoString139`.

	// Assume that `sourceCQL` has the underlying type of `fromUserinfo541`:
	fromUserinfo541 := sourceCQL.(url.Userinfo)

	// Call the method that transfers the taint
	// from the receiver `fromUserinfo541` to the result `intoString139`
	// (`intoString139` is now tainted).
	intoString139 := fromUserinfo541.String()

	// Return the tainted `intoString139`:
	return intoString139
}

func TaintStepTest_NetUrlUserinfoUsername_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUserinfo814` into `intoString768`.

	// Assume that `sourceCQL` has the underlying type of `fromUserinfo814`:
	fromUserinfo814 := sourceCQL.(url.Userinfo)

	// Call the method that transfers the taint
	// from the receiver `fromUserinfo814` to the result `intoString768`
	// (`intoString768` is now tainted).
	intoString768 := fromUserinfo814.Username()

	// Return the tainted `intoString768`:
	return intoString768
}

func TaintStepTest_NetUrlValuesAdd_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString468` into `intoValues736`.

	// Assume that `sourceCQL` has the underlying type of `fromString468`:
	fromString468 := sourceCQL.(string)

	// Declare `intoValues736` variable:
	var intoValues736 url.Values

	// Call the method that transfers the taint
	// from the parameter `fromString468` to the receiver `intoValues736`
	// (`intoValues736` is now tainted).
	intoValues736.Add(fromString468, "")

	// Return the tainted `intoValues736`:
	return intoValues736
}

func TaintStepTest_NetUrlValuesAdd_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString516` into `intoValues246`.

	// Assume that `sourceCQL` has the underlying type of `fromString516`:
	fromString516 := sourceCQL.(string)

	// Declare `intoValues246` variable:
	var intoValues246 url.Values

	// Call the method that transfers the taint
	// from the parameter `fromString516` to the receiver `intoValues246`
	// (`intoValues246` is now tainted).
	intoValues246.Add("", fromString516)

	// Return the tainted `intoValues246`:
	return intoValues246
}

func TaintStepTest_NetUrlValuesEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValues679` into `intoString736`.

	// Assume that `sourceCQL` has the underlying type of `fromValues679`:
	fromValues679 := sourceCQL.(url.Values)

	// Call the method that transfers the taint
	// from the receiver `fromValues679` to the result `intoString736`
	// (`intoString736` is now tainted).
	intoString736 := fromValues679.Encode()

	// Return the tainted `intoString736`:
	return intoString736
}

func TaintStepTest_NetUrlValuesGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValues839` into `intoString273`.

	// Assume that `sourceCQL` has the underlying type of `fromValues839`:
	fromValues839 := sourceCQL.(url.Values)

	// Call the method that transfers the taint
	// from the receiver `fromValues839` to the result `intoString273`
	// (`intoString273` is now tainted).
	intoString273 := fromValues839.Get("")

	// Return the tainted `intoString273`:
	return intoString273
}

func TaintStepTest_NetUrlValuesSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString982` into `intoValues458`.

	// Assume that `sourceCQL` has the underlying type of `fromString982`:
	fromString982 := sourceCQL.(string)

	// Declare `intoValues458` variable:
	var intoValues458 url.Values

	// Call the method that transfers the taint
	// from the parameter `fromString982` to the receiver `intoValues458`
	// (`intoValues458` is now tainted).
	intoValues458.Set(fromString982, "")

	// Return the tainted `intoValues458`:
	return intoValues458
}

func TaintStepTest_NetUrlValuesSet_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString506` into `intoValues213`.

	// Assume that `sourceCQL` has the underlying type of `fromString506`:
	fromString506 := sourceCQL.(string)

	// Declare `intoValues213` variable:
	var intoValues213 url.Values

	// Call the method that transfers the taint
	// from the parameter `fromString506` to the receiver `intoValues213`
	// (`intoValues213` is now tainted).
	intoValues213.Set("", fromString506)

	// Return the tainted `intoValues213`:
	return intoValues213
}

func RunAllTaints_NetUrl() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlParse_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlParseQuery_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlParseRequestURI_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlPathEscape_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlPathUnescape_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlQueryEscape_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlQueryUnescape_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlUser_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlUserPassword_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlUserPassword_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlURLEscapedPath_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlURLHostname_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlURLMarshalBinary_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlURLParse_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlURLParse_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlURLPort_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlURLQuery_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlURLRequestURI_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlURLResolveReference_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlURLResolveReference_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlURLString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlURLUnmarshalBinary_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlUserinfoPassword_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlUserinfoString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlUserinfoUsername_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlValuesAdd_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlValuesAdd_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlValuesEncode_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlValuesGet_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlValuesSet_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUrlValuesSet_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
