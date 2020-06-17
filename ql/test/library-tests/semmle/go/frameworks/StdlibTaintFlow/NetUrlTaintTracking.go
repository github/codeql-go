// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "net/url"

func TaintStepTest_NetUrlParse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString915` into `intoURL900`.

	// Assume that `sourceCQL` has the underlying type of `fromString915`:
	fromString915 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString915` to result `intoURL900`
	// (`intoURL900` is now tainted).
	intoURL900, _ := url.Parse(fromString915)

	// Return the tainted `intoURL900`:
	return intoURL900
}

func TaintStepTest_NetUrlParseQuery_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString344` into `intoValues472`.

	// Assume that `sourceCQL` has the underlying type of `fromString344`:
	fromString344 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString344` to result `intoValues472`
	// (`intoValues472` is now tainted).
	intoValues472, _ := url.ParseQuery(fromString344)

	// Return the tainted `intoValues472`:
	return intoValues472
}

func TaintStepTest_NetUrlParseRequestURI_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString744` into `intoURL323`.

	// Assume that `sourceCQL` has the underlying type of `fromString744`:
	fromString744 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString744` to result `intoURL323`
	// (`intoURL323` is now tainted).
	intoURL323, _ := url.ParseRequestURI(fromString744)

	// Return the tainted `intoURL323`:
	return intoURL323
}

func TaintStepTest_NetUrlPathEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString456` into `intoString857`.

	// Assume that `sourceCQL` has the underlying type of `fromString456`:
	fromString456 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString456` to result `intoString857`
	// (`intoString857` is now tainted).
	intoString857 := url.PathEscape(fromString456)

	// Return the tainted `intoString857`:
	return intoString857
}

func TaintStepTest_NetUrlPathUnescape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString493` into `intoString997`.

	// Assume that `sourceCQL` has the underlying type of `fromString493`:
	fromString493 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString493` to result `intoString997`
	// (`intoString997` is now tainted).
	intoString997, _ := url.PathUnescape(fromString493)

	// Return the tainted `intoString997`:
	return intoString997
}

func TaintStepTest_NetUrlQueryEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString226` into `intoString308`.

	// Assume that `sourceCQL` has the underlying type of `fromString226`:
	fromString226 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString226` to result `intoString308`
	// (`intoString308` is now tainted).
	intoString308 := url.QueryEscape(fromString226)

	// Return the tainted `intoString308`:
	return intoString308
}

func TaintStepTest_NetUrlQueryUnescape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString159` into `intoString940`.

	// Assume that `sourceCQL` has the underlying type of `fromString159`:
	fromString159 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString159` to result `intoString940`
	// (`intoString940` is now tainted).
	intoString940, _ := url.QueryUnescape(fromString159)

	// Return the tainted `intoString940`:
	return intoString940
}

func TaintStepTest_NetUrlUser_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString954` into `intoUserinfo162`.

	// Assume that `sourceCQL` has the underlying type of `fromString954`:
	fromString954 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString954` to result `intoUserinfo162`
	// (`intoUserinfo162` is now tainted).
	intoUserinfo162 := url.User(fromString954)

	// Return the tainted `intoUserinfo162`:
	return intoUserinfo162
}

func TaintStepTest_NetUrlUserPassword_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString494` into `intoUserinfo739`.

	// Assume that `sourceCQL` has the underlying type of `fromString494`:
	fromString494 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString494` to result `intoUserinfo739`
	// (`intoUserinfo739` is now tainted).
	intoUserinfo739 := url.UserPassword(fromString494, "")

	// Return the tainted `intoUserinfo739`:
	return intoUserinfo739
}

func TaintStepTest_NetUrlUserPassword_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString135` into `intoUserinfo613`.

	// Assume that `sourceCQL` has the underlying type of `fromString135`:
	fromString135 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString135` to result `intoUserinfo613`
	// (`intoUserinfo613` is now tainted).
	intoUserinfo613 := url.UserPassword("", fromString135)

	// Return the tainted `intoUserinfo613`:
	return intoUserinfo613
}

func TaintStepTest_NetUrlURLEscapedPath_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL567` into `intoString332`.

	// Assume that `sourceCQL` has the underlying type of `fromURL567`:
	fromURL567 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL567` to the result `intoString332`
	// (`intoString332` is now tainted).
	intoString332 := fromURL567.EscapedPath()

	// Return the tainted `intoString332`:
	return intoString332
}

func TaintStepTest_NetUrlURLHostname_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL843` into `intoString117`.

	// Assume that `sourceCQL` has the underlying type of `fromURL843`:
	fromURL843 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL843` to the result `intoString117`
	// (`intoString117` is now tainted).
	intoString117 := fromURL843.Hostname()

	// Return the tainted `intoString117`:
	return intoString117
}

func TaintStepTest_NetUrlURLMarshalBinary_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL281` into `intoByte938`.

	// Assume that `sourceCQL` has the underlying type of `fromURL281`:
	fromURL281 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL281` to the result `intoByte938`
	// (`intoByte938` is now tainted).
	intoByte938, _ := fromURL281.MarshalBinary()

	// Return the tainted `intoByte938`:
	return intoByte938
}

func TaintStepTest_NetUrlURLParse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL451` into `intoURL239`.

	// Assume that `sourceCQL` has the underlying type of `fromURL451`:
	fromURL451 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL451` to the result `intoURL239`
	// (`intoURL239` is now tainted).
	intoURL239, _ := fromURL451.Parse("")

	// Return the tainted `intoURL239`:
	return intoURL239
}

func TaintStepTest_NetUrlURLParse_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString531` into `intoURL502`.

	// Assume that `sourceCQL` has the underlying type of `fromString531`:
	fromString531 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL url.URL

	// Call the method that transfers the taint
	// from the parameter `fromString531` to the result `intoURL502`
	// (`intoURL502` is now tainted).
	intoURL502, _ := mediumObjCQL.Parse(fromString531)

	// Return the tainted `intoURL502`:
	return intoURL502
}

func TaintStepTest_NetUrlURLPort_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL497` into `intoString391`.

	// Assume that `sourceCQL` has the underlying type of `fromURL497`:
	fromURL497 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL497` to the result `intoString391`
	// (`intoString391` is now tainted).
	intoString391 := fromURL497.Port()

	// Return the tainted `intoString391`:
	return intoString391
}

func TaintStepTest_NetUrlURLQuery_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL850` into `intoValues352`.

	// Assume that `sourceCQL` has the underlying type of `fromURL850`:
	fromURL850 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL850` to the result `intoValues352`
	// (`intoValues352` is now tainted).
	intoValues352 := fromURL850.Query()

	// Return the tainted `intoValues352`:
	return intoValues352
}

func TaintStepTest_NetUrlURLRequestURI_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL353` into `intoString322`.

	// Assume that `sourceCQL` has the underlying type of `fromURL353`:
	fromURL353 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL353` to the result `intoString322`
	// (`intoString322` is now tainted).
	intoString322 := fromURL353.RequestURI()

	// Return the tainted `intoString322`:
	return intoString322
}

func TaintStepTest_NetUrlURLResolveReference_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL691` into `intoURL703`.

	// Assume that `sourceCQL` has the underlying type of `fromURL691`:
	fromURL691 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL691` to the result `intoURL703`
	// (`intoURL703` is now tainted).
	intoURL703 := fromURL691.ResolveReference(nil)

	// Return the tainted `intoURL703`:
	return intoURL703
}

func TaintStepTest_NetUrlURLResolveReference_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL818` into `intoURL210`.

	// Assume that `sourceCQL` has the underlying type of `fromURL818`:
	fromURL818 := sourceCQL.(*url.URL)

	// Declare medium object/interface:
	var mediumObjCQL url.URL

	// Call the method that transfers the taint
	// from the parameter `fromURL818` to the result `intoURL210`
	// (`intoURL210` is now tainted).
	intoURL210 := mediumObjCQL.ResolveReference(fromURL818)

	// Return the tainted `intoURL210`:
	return intoURL210
}

func TaintStepTest_NetUrlURLString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL314` into `intoString857`.

	// Assume that `sourceCQL` has the underlying type of `fromURL314`:
	fromURL314 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL314` to the result `intoString857`
	// (`intoString857` is now tainted).
	intoString857 := fromURL314.String()

	// Return the tainted `intoString857`:
	return intoString857
}

func TaintStepTest_NetUrlURLUnmarshalBinary_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte780` into `intoURL250`.

	// Assume that `sourceCQL` has the underlying type of `fromByte780`:
	fromByte780 := sourceCQL.([]byte)

	// Declare `intoURL250` variable:
	var intoURL250 url.URL

	// Call the method that transfers the taint
	// from the parameter `fromByte780` to the receiver `intoURL250`
	// (`intoURL250` is now tainted).
	intoURL250.UnmarshalBinary(fromByte780)

	// Return the tainted `intoURL250`:
	return intoURL250
}

func TaintStepTest_NetUrlUserinfoPassword_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUserinfo632` into `intoString597`.

	// Assume that `sourceCQL` has the underlying type of `fromUserinfo632`:
	fromUserinfo632 := sourceCQL.(url.Userinfo)

	// Call the method that transfers the taint
	// from the receiver `fromUserinfo632` to the result `intoString597`
	// (`intoString597` is now tainted).
	intoString597, _ := fromUserinfo632.Password()

	// Return the tainted `intoString597`:
	return intoString597
}

func TaintStepTest_NetUrlUserinfoString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUserinfo508` into `intoString849`.

	// Assume that `sourceCQL` has the underlying type of `fromUserinfo508`:
	fromUserinfo508 := sourceCQL.(url.Userinfo)

	// Call the method that transfers the taint
	// from the receiver `fromUserinfo508` to the result `intoString849`
	// (`intoString849` is now tainted).
	intoString849 := fromUserinfo508.String()

	// Return the tainted `intoString849`:
	return intoString849
}

func TaintStepTest_NetUrlUserinfoUsername_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUserinfo634` into `intoString183`.

	// Assume that `sourceCQL` has the underlying type of `fromUserinfo634`:
	fromUserinfo634 := sourceCQL.(url.Userinfo)

	// Call the method that transfers the taint
	// from the receiver `fromUserinfo634` to the result `intoString183`
	// (`intoString183` is now tainted).
	intoString183 := fromUserinfo634.Username()

	// Return the tainted `intoString183`:
	return intoString183
}

func TaintStepTest_NetUrlValuesAdd_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString572` into `intoValues388`.

	// Assume that `sourceCQL` has the underlying type of `fromString572`:
	fromString572 := sourceCQL.(string)

	// Declare `intoValues388` variable:
	var intoValues388 url.Values

	// Call the method that transfers the taint
	// from the parameter `fromString572` to the receiver `intoValues388`
	// (`intoValues388` is now tainted).
	intoValues388.Add(fromString572, "")

	// Return the tainted `intoValues388`:
	return intoValues388
}

func TaintStepTest_NetUrlValuesAdd_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString277` into `intoValues690`.

	// Assume that `sourceCQL` has the underlying type of `fromString277`:
	fromString277 := sourceCQL.(string)

	// Declare `intoValues690` variable:
	var intoValues690 url.Values

	// Call the method that transfers the taint
	// from the parameter `fromString277` to the receiver `intoValues690`
	// (`intoValues690` is now tainted).
	intoValues690.Add("", fromString277)

	// Return the tainted `intoValues690`:
	return intoValues690
}

func TaintStepTest_NetUrlValuesEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValues332` into `intoString184`.

	// Assume that `sourceCQL` has the underlying type of `fromValues332`:
	fromValues332 := sourceCQL.(url.Values)

	// Call the method that transfers the taint
	// from the receiver `fromValues332` to the result `intoString184`
	// (`intoString184` is now tainted).
	intoString184 := fromValues332.Encode()

	// Return the tainted `intoString184`:
	return intoString184
}

func TaintStepTest_NetUrlValuesGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValues610` into `intoString922`.

	// Assume that `sourceCQL` has the underlying type of `fromValues610`:
	fromValues610 := sourceCQL.(url.Values)

	// Call the method that transfers the taint
	// from the receiver `fromValues610` to the result `intoString922`
	// (`intoString922` is now tainted).
	intoString922 := fromValues610.Get("")

	// Return the tainted `intoString922`:
	return intoString922
}

func TaintStepTest_NetUrlValuesSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString962` into `intoValues506`.

	// Assume that `sourceCQL` has the underlying type of `fromString962`:
	fromString962 := sourceCQL.(string)

	// Declare `intoValues506` variable:
	var intoValues506 url.Values

	// Call the method that transfers the taint
	// from the parameter `fromString962` to the receiver `intoValues506`
	// (`intoValues506` is now tainted).
	intoValues506.Set(fromString962, "")

	// Return the tainted `intoValues506`:
	return intoValues506
}

func TaintStepTest_NetUrlValuesSet_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString413` into `intoValues120`.

	// Assume that `sourceCQL` has the underlying type of `fromString413`:
	fromString413 := sourceCQL.(string)

	// Declare `intoValues120` variable:
	var intoValues120 url.Values

	// Call the method that transfers the taint
	// from the parameter `fromString413` to the receiver `intoValues120`
	// (`intoValues120` is now tainted).
	intoValues120.Set("", fromString413)

	// Return the tainted `intoValues120`:
	return intoValues120
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
