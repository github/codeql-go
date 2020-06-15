package main

import "net/url"

func TaintStepTest_NetUrlParse_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString662` into `intoURL555`.

	// Assume that `sourceCQL` has the underlying type of `fromString662`:
	fromString662 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString662` to result `intoURL555`
	// (`intoURL555` is now tainted).
	intoURL555, _ := url.Parse(fromString662)

	// Sink the tainted `intoURL555`:
	sink(intoURL555)
}

func TaintStepTest_NetUrlParseQuery_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString205` into `intoValues473`.

	// Assume that `sourceCQL` has the underlying type of `fromString205`:
	fromString205 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString205` to result `intoValues473`
	// (`intoValues473` is now tainted).
	intoValues473, _ := url.ParseQuery(fromString205)

	// Sink the tainted `intoValues473`:
	sink(intoValues473)
}

func TaintStepTest_NetUrlParseRequestURI_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString301` into `intoURL505`.

	// Assume that `sourceCQL` has the underlying type of `fromString301`:
	fromString301 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString301` to result `intoURL505`
	// (`intoURL505` is now tainted).
	intoURL505, _ := url.ParseRequestURI(fromString301)

	// Sink the tainted `intoURL505`:
	sink(intoURL505)
}

func TaintStepTest_NetUrlPathEscape_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString546` into `intoString236`.

	// Assume that `sourceCQL` has the underlying type of `fromString546`:
	fromString546 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString546` to result `intoString236`
	// (`intoString236` is now tainted).
	intoString236 := url.PathEscape(fromString546)

	// Sink the tainted `intoString236`:
	sink(intoString236)
}

func TaintStepTest_NetUrlPathUnescape_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString197` into `intoString706`.

	// Assume that `sourceCQL` has the underlying type of `fromString197`:
	fromString197 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString197` to result `intoString706`
	// (`intoString706` is now tainted).
	intoString706, _ := url.PathUnescape(fromString197)

	// Sink the tainted `intoString706`:
	sink(intoString706)
}

func TaintStepTest_NetUrlQueryEscape_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString555` into `intoString782`.

	// Assume that `sourceCQL` has the underlying type of `fromString555`:
	fromString555 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString555` to result `intoString782`
	// (`intoString782` is now tainted).
	intoString782 := url.QueryEscape(fromString555)

	// Sink the tainted `intoString782`:
	sink(intoString782)
}

func TaintStepTest_NetUrlQueryUnescape_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString934` into `intoString663`.

	// Assume that `sourceCQL` has the underlying type of `fromString934`:
	fromString934 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString934` to result `intoString663`
	// (`intoString663` is now tainted).
	intoString663, _ := url.QueryUnescape(fromString934)

	// Sink the tainted `intoString663`:
	sink(intoString663)
}

func TaintStepTest_NetUrlUser_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString768` into `intoUserinfo994`.

	// Assume that `sourceCQL` has the underlying type of `fromString768`:
	fromString768 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString768` to result `intoUserinfo994`
	// (`intoUserinfo994` is now tainted).
	intoUserinfo994 := url.User(fromString768)

	// Sink the tainted `intoUserinfo994`:
	sink(intoUserinfo994)
}

func TaintStepTest_NetUrlUserPassword_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString916` into `intoUserinfo472`.

	// Assume that `sourceCQL` has the underlying type of `fromString916`:
	fromString916 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString916` to result `intoUserinfo472`
	// (`intoUserinfo472` is now tainted).
	intoUserinfo472 := url.UserPassword(fromString916, "")

	// Sink the tainted `intoUserinfo472`:
	sink(intoUserinfo472)
}

func TaintStepTest_NetUrlUserPassword_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString992` into `intoUserinfo364`.

	// Assume that `sourceCQL` has the underlying type of `fromString992`:
	fromString992 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString992` to result `intoUserinfo364`
	// (`intoUserinfo364` is now tainted).
	intoUserinfo364 := url.UserPassword("", fromString992)

	// Sink the tainted `intoUserinfo364`:
	sink(intoUserinfo364)
}

func TaintStepTest_NetUrlURLEscapedPath_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromURL463` into `intoString405`.

	// Assume that `sourceCQL` has the underlying type of `fromURL463`:
	fromURL463 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL463` to the result `intoString405`
	// (`intoString405` is now tainted).
	intoString405 := fromURL463.EscapedPath()

	// Sink the tainted `intoString405`:
	sink(intoString405)
}

func TaintStepTest_NetUrlURLHostname_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromURL196` into `intoString256`.

	// Assume that `sourceCQL` has the underlying type of `fromURL196`:
	fromURL196 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL196` to the result `intoString256`
	// (`intoString256` is now tainted).
	intoString256 := fromURL196.Hostname()

	// Sink the tainted `intoString256`:
	sink(intoString256)
}

func TaintStepTest_NetUrlURLMarshalBinary_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromURL649` into `intoByte791`.

	// Assume that `sourceCQL` has the underlying type of `fromURL649`:
	fromURL649 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL649` to the result `intoByte791`
	// (`intoByte791` is now tainted).
	intoByte791, _ := fromURL649.MarshalBinary()

	// Sink the tainted `intoByte791`:
	sink(intoByte791)
}

func TaintStepTest_NetUrlURLParse_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromURL710` into `intoURL287`.

	// Assume that `sourceCQL` has the underlying type of `fromURL710`:
	fromURL710 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL710` to the result `intoURL287`
	// (`intoURL287` is now tainted).
	intoURL287, _ := fromURL710.Parse("")

	// Sink the tainted `intoURL287`:
	sink(intoURL287)
}

func TaintStepTest_NetUrlURLParse_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString493` into `intoURL163`.

	// Assume that `sourceCQL` has the underlying type of `fromString493`:
	fromString493 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL url.URL

	// Call the method that transfers the taint
	// from the parameter `fromString493` to the result `intoURL163`
	// (`intoURL163` is now tainted).
	intoURL163, _ := mediumObjCQL.Parse(fromString493)

	// Sink the tainted `intoURL163`:
	sink(intoURL163)
}

func TaintStepTest_NetUrlURLPort_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromURL234` into `intoString936`.

	// Assume that `sourceCQL` has the underlying type of `fromURL234`:
	fromURL234 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL234` to the result `intoString936`
	// (`intoString936` is now tainted).
	intoString936 := fromURL234.Port()

	// Sink the tainted `intoString936`:
	sink(intoString936)
}

func TaintStepTest_NetUrlURLQuery_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromURL630` into `intoValues177`.

	// Assume that `sourceCQL` has the underlying type of `fromURL630`:
	fromURL630 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL630` to the result `intoValues177`
	// (`intoValues177` is now tainted).
	intoValues177 := fromURL630.Query()

	// Sink the tainted `intoValues177`:
	sink(intoValues177)
}

func TaintStepTest_NetUrlURLRequestURI_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromURL686` into `intoString730`.

	// Assume that `sourceCQL` has the underlying type of `fromURL686`:
	fromURL686 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL686` to the result `intoString730`
	// (`intoString730` is now tainted).
	intoString730 := fromURL686.RequestURI()

	// Sink the tainted `intoString730`:
	sink(intoString730)
}

func TaintStepTest_NetUrlURLResolveReference_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromURL559` into `intoURL702`.

	// Assume that `sourceCQL` has the underlying type of `fromURL559`:
	fromURL559 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL559` to the result `intoURL702`
	// (`intoURL702` is now tainted).
	intoURL702 := fromURL559.ResolveReference(nil)

	// Sink the tainted `intoURL702`:
	sink(intoURL702)
}

func TaintStepTest_NetUrlURLResolveReference_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromURL989` into `intoURL112`.

	// Assume that `sourceCQL` has the underlying type of `fromURL989`:
	fromURL989 := sourceCQL.(*url.URL)

	// Declare medium object/interface:
	var mediumObjCQL url.URL

	// Call the method that transfers the taint
	// from the parameter `fromURL989` to the result `intoURL112`
	// (`intoURL112` is now tainted).
	intoURL112 := mediumObjCQL.ResolveReference(fromURL989)

	// Sink the tainted `intoURL112`:
	sink(intoURL112)
}

func TaintStepTest_NetUrlURLString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromURL463` into `intoString372`.

	// Assume that `sourceCQL` has the underlying type of `fromURL463`:
	fromURL463 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL463` to the result `intoString372`
	// (`intoString372` is now tainted).
	intoString372 := fromURL463.String()

	// Sink the tainted `intoString372`:
	sink(intoString372)
}

func TaintStepTest_NetUrlURLUnmarshalBinary_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte542` into `intoURL886`.

	// Assume that `sourceCQL` has the underlying type of `fromByte542`:
	fromByte542 := sourceCQL.([]byte)

	// Declare `intoURL886` variable:
	var intoURL886 url.URL

	// Call the method that transfers the taint
	// from the parameter `fromByte542` to the receiver `intoURL886`
	// (`intoURL886` is now tainted).
	intoURL886.UnmarshalBinary(fromByte542)

	// Sink the tainted `intoURL886`:
	sink(intoURL886)
}

func TaintStepTest_NetUrlUserinfoPassword_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUserinfo945` into `intoString947`.

	// Assume that `sourceCQL` has the underlying type of `fromUserinfo945`:
	fromUserinfo945 := sourceCQL.(url.Userinfo)

	// Call the method that transfers the taint
	// from the receiver `fromUserinfo945` to the result `intoString947`
	// (`intoString947` is now tainted).
	intoString947, _ := fromUserinfo945.Password()

	// Sink the tainted `intoString947`:
	sink(intoString947)
}

func TaintStepTest_NetUrlUserinfoString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUserinfo631` into `intoString664`.

	// Assume that `sourceCQL` has the underlying type of `fromUserinfo631`:
	fromUserinfo631 := sourceCQL.(url.Userinfo)

	// Call the method that transfers the taint
	// from the receiver `fromUserinfo631` to the result `intoString664`
	// (`intoString664` is now tainted).
	intoString664 := fromUserinfo631.String()

	// Sink the tainted `intoString664`:
	sink(intoString664)
}

func TaintStepTest_NetUrlUserinfoUsername_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUserinfo203` into `intoString819`.

	// Assume that `sourceCQL` has the underlying type of `fromUserinfo203`:
	fromUserinfo203 := sourceCQL.(url.Userinfo)

	// Call the method that transfers the taint
	// from the receiver `fromUserinfo203` to the result `intoString819`
	// (`intoString819` is now tainted).
	intoString819 := fromUserinfo203.Username()

	// Sink the tainted `intoString819`:
	sink(intoString819)
}

func TaintStepTest_NetUrlValuesAdd_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString291` into `intoValues156`.

	// Assume that `sourceCQL` has the underlying type of `fromString291`:
	fromString291 := sourceCQL.(string)

	// Declare `intoValues156` variable:
	var intoValues156 url.Values

	// Call the method that transfers the taint
	// from the parameter `fromString291` to the receiver `intoValues156`
	// (`intoValues156` is now tainted).
	intoValues156.Add(fromString291, "")

	// Sink the tainted `intoValues156`:
	sink(intoValues156)
}

func TaintStepTest_NetUrlValuesAdd_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString904` into `intoValues350`.

	// Assume that `sourceCQL` has the underlying type of `fromString904`:
	fromString904 := sourceCQL.(string)

	// Declare `intoValues350` variable:
	var intoValues350 url.Values

	// Call the method that transfers the taint
	// from the parameter `fromString904` to the receiver `intoValues350`
	// (`intoValues350` is now tainted).
	intoValues350.Add("", fromString904)

	// Sink the tainted `intoValues350`:
	sink(intoValues350)
}

func TaintStepTest_NetUrlValuesEncode_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValues213` into `intoString951`.

	// Assume that `sourceCQL` has the underlying type of `fromValues213`:
	fromValues213 := sourceCQL.(url.Values)

	// Call the method that transfers the taint
	// from the receiver `fromValues213` to the result `intoString951`
	// (`intoString951` is now tainted).
	intoString951 := fromValues213.Encode()

	// Sink the tainted `intoString951`:
	sink(intoString951)
}

func TaintStepTest_NetUrlValuesGet_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValues163` into `intoString113`.

	// Assume that `sourceCQL` has the underlying type of `fromValues163`:
	fromValues163 := sourceCQL.(url.Values)

	// Call the method that transfers the taint
	// from the receiver `fromValues163` to the result `intoString113`
	// (`intoString113` is now tainted).
	intoString113 := fromValues163.Get("")

	// Sink the tainted `intoString113`:
	sink(intoString113)
}

func TaintStepTest_NetUrlValuesSet_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString786` into `intoValues691`.

	// Assume that `sourceCQL` has the underlying type of `fromString786`:
	fromString786 := sourceCQL.(string)

	// Declare `intoValues691` variable:
	var intoValues691 url.Values

	// Call the method that transfers the taint
	// from the parameter `fromString786` to the receiver `intoValues691`
	// (`intoValues691` is now tainted).
	intoValues691.Set(fromString786, "")

	// Sink the tainted `intoValues691`:
	sink(intoValues691)
}

func TaintStepTest_NetUrlValuesSet_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString294` into `intoValues564`.

	// Assume that `sourceCQL` has the underlying type of `fromString294`:
	fromString294 := sourceCQL.(string)

	// Declare `intoValues564` variable:
	var intoValues564 url.Values

	// Call the method that transfers the taint
	// from the parameter `fromString294` to the receiver `intoValues564`
	// (`intoValues564` is now tainted).
	intoValues564.Set("", fromString294)

	// Sink the tainted `intoValues564`:
	sink(intoValues564)
}

func RunAllTaints_NetUrl() {
	{
		source := newSource()
		TaintStepTest_NetUrlParse_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlParseQuery_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlParseRequestURI_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlPathEscape_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlPathUnescape_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlQueryEscape_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlQueryUnescape_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlUser_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlUserPassword_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlUserPassword_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLEscapedPath_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLHostname_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLMarshalBinary_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLParse_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLParse_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLPort_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLQuery_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLRequestURI_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLResolveReference_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLResolveReference_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLUnmarshalBinary_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlUserinfoPassword_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlUserinfoString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlUserinfoUsername_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlValuesAdd_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlValuesAdd_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlValuesEncode_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlValuesGet_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlValuesSet_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlValuesSet_B0I1O0(source)
	}
}
