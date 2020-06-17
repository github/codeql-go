package main

import "net/url"

func TaintStepTest_NetUrlParse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString534` into `intoURL137`.

	// Assume that `sourceCQL` has the underlying type of `fromString534`:
	fromString534 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString534` to result `intoURL137`
	// (`intoURL137` is now tainted).
	intoURL137, _ := url.Parse(fromString534)

	// Return the tainted `intoURL137`:
	return intoURL137
}

func TaintStepTest_NetUrlParseQuery_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString361` into `intoValues248`.

	// Assume that `sourceCQL` has the underlying type of `fromString361`:
	fromString361 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString361` to result `intoValues248`
	// (`intoValues248` is now tainted).
	intoValues248, _ := url.ParseQuery(fromString361)

	// Return the tainted `intoValues248`:
	return intoValues248
}

func TaintStepTest_NetUrlParseRequestURI_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString262` into `intoURL612`.

	// Assume that `sourceCQL` has the underlying type of `fromString262`:
	fromString262 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString262` to result `intoURL612`
	// (`intoURL612` is now tainted).
	intoURL612, _ := url.ParseRequestURI(fromString262)

	// Return the tainted `intoURL612`:
	return intoURL612
}

func TaintStepTest_NetUrlPathEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString186` into `intoString794`.

	// Assume that `sourceCQL` has the underlying type of `fromString186`:
	fromString186 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString186` to result `intoString794`
	// (`intoString794` is now tainted).
	intoString794 := url.PathEscape(fromString186)

	// Return the tainted `intoString794`:
	return intoString794
}

func TaintStepTest_NetUrlPathUnescape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString880` into `intoString371`.

	// Assume that `sourceCQL` has the underlying type of `fromString880`:
	fromString880 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString880` to result `intoString371`
	// (`intoString371` is now tainted).
	intoString371, _ := url.PathUnescape(fromString880)

	// Return the tainted `intoString371`:
	return intoString371
}

func TaintStepTest_NetUrlQueryEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString657` into `intoString177`.

	// Assume that `sourceCQL` has the underlying type of `fromString657`:
	fromString657 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString657` to result `intoString177`
	// (`intoString177` is now tainted).
	intoString177 := url.QueryEscape(fromString657)

	// Return the tainted `intoString177`:
	return intoString177
}

func TaintStepTest_NetUrlQueryUnescape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString118` into `intoString813`.

	// Assume that `sourceCQL` has the underlying type of `fromString118`:
	fromString118 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString118` to result `intoString813`
	// (`intoString813` is now tainted).
	intoString813, _ := url.QueryUnescape(fromString118)

	// Return the tainted `intoString813`:
	return intoString813
}

func TaintStepTest_NetUrlUser_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString377` into `intoUserinfo720`.

	// Assume that `sourceCQL` has the underlying type of `fromString377`:
	fromString377 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString377` to result `intoUserinfo720`
	// (`intoUserinfo720` is now tainted).
	intoUserinfo720 := url.User(fromString377)

	// Return the tainted `intoUserinfo720`:
	return intoUserinfo720
}

func TaintStepTest_NetUrlUserPassword_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString157` into `intoUserinfo170`.

	// Assume that `sourceCQL` has the underlying type of `fromString157`:
	fromString157 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString157` to result `intoUserinfo170`
	// (`intoUserinfo170` is now tainted).
	intoUserinfo170 := url.UserPassword(fromString157, "")

	// Return the tainted `intoUserinfo170`:
	return intoUserinfo170
}

func TaintStepTest_NetUrlUserPassword_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString339` into `intoUserinfo839`.

	// Assume that `sourceCQL` has the underlying type of `fromString339`:
	fromString339 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString339` to result `intoUserinfo839`
	// (`intoUserinfo839` is now tainted).
	intoUserinfo839 := url.UserPassword("", fromString339)

	// Return the tainted `intoUserinfo839`:
	return intoUserinfo839
}

func TaintStepTest_NetUrlURLEscapedPath_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL850` into `intoString674`.

	// Assume that `sourceCQL` has the underlying type of `fromURL850`:
	fromURL850 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL850` to the result `intoString674`
	// (`intoString674` is now tainted).
	intoString674 := fromURL850.EscapedPath()

	// Return the tainted `intoString674`:
	return intoString674
}

func TaintStepTest_NetUrlURLHostname_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL285` into `intoString389`.

	// Assume that `sourceCQL` has the underlying type of `fromURL285`:
	fromURL285 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL285` to the result `intoString389`
	// (`intoString389` is now tainted).
	intoString389 := fromURL285.Hostname()

	// Return the tainted `intoString389`:
	return intoString389
}

func TaintStepTest_NetUrlURLMarshalBinary_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL519` into `intoByte490`.

	// Assume that `sourceCQL` has the underlying type of `fromURL519`:
	fromURL519 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL519` to the result `intoByte490`
	// (`intoByte490` is now tainted).
	intoByte490, _ := fromURL519.MarshalBinary()

	// Return the tainted `intoByte490`:
	return intoByte490
}

func TaintStepTest_NetUrlURLParse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL680` into `intoURL291`.

	// Assume that `sourceCQL` has the underlying type of `fromURL680`:
	fromURL680 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL680` to the result `intoURL291`
	// (`intoURL291` is now tainted).
	intoURL291, _ := fromURL680.Parse("")

	// Return the tainted `intoURL291`:
	return intoURL291
}

func TaintStepTest_NetUrlURLParse_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString230` into `intoURL338`.

	// Assume that `sourceCQL` has the underlying type of `fromString230`:
	fromString230 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL url.URL

	// Call the method that transfers the taint
	// from the parameter `fromString230` to the result `intoURL338`
	// (`intoURL338` is now tainted).
	intoURL338, _ := mediumObjCQL.Parse(fromString230)

	// Return the tainted `intoURL338`:
	return intoURL338
}

func TaintStepTest_NetUrlURLPort_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL373` into `intoString145`.

	// Assume that `sourceCQL` has the underlying type of `fromURL373`:
	fromURL373 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL373` to the result `intoString145`
	// (`intoString145` is now tainted).
	intoString145 := fromURL373.Port()

	// Return the tainted `intoString145`:
	return intoString145
}

func TaintStepTest_NetUrlURLQuery_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL609` into `intoValues757`.

	// Assume that `sourceCQL` has the underlying type of `fromURL609`:
	fromURL609 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL609` to the result `intoValues757`
	// (`intoValues757` is now tainted).
	intoValues757 := fromURL609.Query()

	// Return the tainted `intoValues757`:
	return intoValues757
}

func TaintStepTest_NetUrlURLRequestURI_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL158` into `intoString944`.

	// Assume that `sourceCQL` has the underlying type of `fromURL158`:
	fromURL158 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL158` to the result `intoString944`
	// (`intoString944` is now tainted).
	intoString944 := fromURL158.RequestURI()

	// Return the tainted `intoString944`:
	return intoString944
}

func TaintStepTest_NetUrlURLResolveReference_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL498` into `intoURL400`.

	// Assume that `sourceCQL` has the underlying type of `fromURL498`:
	fromURL498 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL498` to the result `intoURL400`
	// (`intoURL400` is now tainted).
	intoURL400 := fromURL498.ResolveReference(nil)

	// Return the tainted `intoURL400`:
	return intoURL400
}

func TaintStepTest_NetUrlURLResolveReference_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL440` into `intoURL750`.

	// Assume that `sourceCQL` has the underlying type of `fromURL440`:
	fromURL440 := sourceCQL.(*url.URL)

	// Declare medium object/interface:
	var mediumObjCQL url.URL

	// Call the method that transfers the taint
	// from the parameter `fromURL440` to the result `intoURL750`
	// (`intoURL750` is now tainted).
	intoURL750 := mediumObjCQL.ResolveReference(fromURL440)

	// Return the tainted `intoURL750`:
	return intoURL750
}

func TaintStepTest_NetUrlURLString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromURL949` into `intoString137`.

	// Assume that `sourceCQL` has the underlying type of `fromURL949`:
	fromURL949 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `fromURL949` to the result `intoString137`
	// (`intoString137` is now tainted).
	intoString137 := fromURL949.String()

	// Return the tainted `intoString137`:
	return intoString137
}

func TaintStepTest_NetUrlURLUnmarshalBinary_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte561` into `intoURL580`.

	// Assume that `sourceCQL` has the underlying type of `fromByte561`:
	fromByte561 := sourceCQL.([]byte)

	// Declare `intoURL580` variable:
	var intoURL580 url.URL

	// Call the method that transfers the taint
	// from the parameter `fromByte561` to the receiver `intoURL580`
	// (`intoURL580` is now tainted).
	intoURL580.UnmarshalBinary(fromByte561)

	// Return the tainted `intoURL580`:
	return intoURL580
}

func TaintStepTest_NetUrlUserinfoPassword_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUserinfo265` into `intoString378`.

	// Assume that `sourceCQL` has the underlying type of `fromUserinfo265`:
	fromUserinfo265 := sourceCQL.(url.Userinfo)

	// Call the method that transfers the taint
	// from the receiver `fromUserinfo265` to the result `intoString378`
	// (`intoString378` is now tainted).
	intoString378, _ := fromUserinfo265.Password()

	// Return the tainted `intoString378`:
	return intoString378
}

func TaintStepTest_NetUrlUserinfoString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUserinfo718` into `intoString306`.

	// Assume that `sourceCQL` has the underlying type of `fromUserinfo718`:
	fromUserinfo718 := sourceCQL.(url.Userinfo)

	// Call the method that transfers the taint
	// from the receiver `fromUserinfo718` to the result `intoString306`
	// (`intoString306` is now tainted).
	intoString306 := fromUserinfo718.String()

	// Return the tainted `intoString306`:
	return intoString306
}

func TaintStepTest_NetUrlUserinfoUsername_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUserinfo636` into `intoString753`.

	// Assume that `sourceCQL` has the underlying type of `fromUserinfo636`:
	fromUserinfo636 := sourceCQL.(url.Userinfo)

	// Call the method that transfers the taint
	// from the receiver `fromUserinfo636` to the result `intoString753`
	// (`intoString753` is now tainted).
	intoString753 := fromUserinfo636.Username()

	// Return the tainted `intoString753`:
	return intoString753
}

func TaintStepTest_NetUrlValuesAdd_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString759` into `intoValues969`.

	// Assume that `sourceCQL` has the underlying type of `fromString759`:
	fromString759 := sourceCQL.(string)

	// Declare `intoValues969` variable:
	var intoValues969 url.Values

	// Call the method that transfers the taint
	// from the parameter `fromString759` to the receiver `intoValues969`
	// (`intoValues969` is now tainted).
	intoValues969.Add(fromString759, "")

	// Return the tainted `intoValues969`:
	return intoValues969
}

func TaintStepTest_NetUrlValuesAdd_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString723` into `intoValues840`.

	// Assume that `sourceCQL` has the underlying type of `fromString723`:
	fromString723 := sourceCQL.(string)

	// Declare `intoValues840` variable:
	var intoValues840 url.Values

	// Call the method that transfers the taint
	// from the parameter `fromString723` to the receiver `intoValues840`
	// (`intoValues840` is now tainted).
	intoValues840.Add("", fromString723)

	// Return the tainted `intoValues840`:
	return intoValues840
}

func TaintStepTest_NetUrlValuesEncode_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValues947` into `intoString120`.

	// Assume that `sourceCQL` has the underlying type of `fromValues947`:
	fromValues947 := sourceCQL.(url.Values)

	// Call the method that transfers the taint
	// from the receiver `fromValues947` to the result `intoString120`
	// (`intoString120` is now tainted).
	intoString120 := fromValues947.Encode()

	// Return the tainted `intoString120`:
	return intoString120
}

func TaintStepTest_NetUrlValuesGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValues282` into `intoString809`.

	// Assume that `sourceCQL` has the underlying type of `fromValues282`:
	fromValues282 := sourceCQL.(url.Values)

	// Call the method that transfers the taint
	// from the receiver `fromValues282` to the result `intoString809`
	// (`intoString809` is now tainted).
	intoString809 := fromValues282.Get("")

	// Return the tainted `intoString809`:
	return intoString809
}

func TaintStepTest_NetUrlValuesSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString615` into `intoValues515`.

	// Assume that `sourceCQL` has the underlying type of `fromString615`:
	fromString615 := sourceCQL.(string)

	// Declare `intoValues515` variable:
	var intoValues515 url.Values

	// Call the method that transfers the taint
	// from the parameter `fromString615` to the receiver `intoValues515`
	// (`intoValues515` is now tainted).
	intoValues515.Set(fromString615, "")

	// Return the tainted `intoValues515`:
	return intoValues515
}

func TaintStepTest_NetUrlValuesSet_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString719` into `intoValues158`.

	// Assume that `sourceCQL` has the underlying type of `fromString719`:
	fromString719 := sourceCQL.(string)

	// Declare `intoValues158` variable:
	var intoValues158 url.Values

	// Call the method that transfers the taint
	// from the parameter `fromString719` to the receiver `intoValues158`
	// (`intoValues158` is now tainted).
	intoValues158.Set("", fromString719)

	// Return the tainted `intoValues158`:
	return intoValues158
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
