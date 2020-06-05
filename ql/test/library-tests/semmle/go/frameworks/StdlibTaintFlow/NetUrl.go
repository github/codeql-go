package main

import "net/url"

func TaintStepTest_NetUrlParse(sourceCQL interface{}) {
	// The flow is from `rawurl` into `into925`.

	// Assume that `sourceCQL` has the underlying type of `rawurl`:
	rawurl := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `rawurl` to result `into925`
	// (`into925` is now tainted).
	into925, _ := url.Parse(rawurl)

	// Sink the tainted `into925`:
	sink(into925)
}

func TaintStepTest_NetUrlParseQuery(sourceCQL interface{}) {
	// The flow is from `query` into `into648`.

	// Assume that `sourceCQL` has the underlying type of `query`:
	query := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `query` to result `into648`
	// (`into648` is now tainted).
	into648, _ := url.ParseQuery(query)

	// Sink the tainted `into648`:
	sink(into648)
}

func TaintStepTest_NetUrlParseRequestURI(sourceCQL interface{}) {
	// The flow is from `rawurl` into `into856`.

	// Assume that `sourceCQL` has the underlying type of `rawurl`:
	rawurl := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `rawurl` to result `into856`
	// (`into856` is now tainted).
	into856, _ := url.ParseRequestURI(rawurl)

	// Sink the tainted `into856`:
	sink(into856)
}

func TaintStepTest_NetUrlPathEscape(sourceCQL interface{}) {
	// The flow is from `s` into `into656`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into656`
	// (`into656` is now tainted).
	into656 := url.PathEscape(s)

	// Sink the tainted `into656`:
	sink(into656)
}

func TaintStepTest_NetUrlPathUnescape(sourceCQL interface{}) {
	// The flow is from `s` into `into765`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into765`
	// (`into765` is now tainted).
	into765, _ := url.PathUnescape(s)

	// Sink the tainted `into765`:
	sink(into765)
}

func TaintStepTest_NetUrlQueryEscape(sourceCQL interface{}) {
	// The flow is from `s` into `into393`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into393`
	// (`into393` is now tainted).
	into393 := url.QueryEscape(s)

	// Sink the tainted `into393`:
	sink(into393)
}

func TaintStepTest_NetUrlQueryUnescape(sourceCQL interface{}) {
	// The flow is from `s` into `into944`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into944`
	// (`into944` is now tainted).
	into944, _ := url.QueryUnescape(s)

	// Sink the tainted `into944`:
	sink(into944)
}

func TaintStepTest_NetUrlUser(sourceCQL interface{}) {
	// The flow is from `username` into `into132`.

	// Assume that `sourceCQL` has the underlying type of `username`:
	username := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `username` to result `into132`
	// (`into132` is now tainted).
	into132 := url.User(username)

	// Sink the tainted `into132`:
	sink(into132)
}

func TaintStepTest_NetUrlUserPassword(sourceCQL interface{}) {
	// The flow is from `username` into `into594`.

	// Assume that `sourceCQL` has the underlying type of `username`:
	username := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `username` to result `into594`
	// (`into594` is now tainted).
	into594 := url.UserPassword(username, "")

	// Sink the tainted `into594`:
	sink(into594)
}

func TaintStepTest_NetUrlURLEscapedPath(sourceCQL interface{}) {
	// The flow is from `from586` into `into508`.

	// Assume that `sourceCQL` has the underlying type of `from586`:
	from586 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `from586` to the result `into508`
	// (`into508` is now tainted).
	into508 := from586.EscapedPath()

	// Sink the tainted `into508`:
	sink(into508)
}

func TaintStepTest_NetUrlURLHostname(sourceCQL interface{}) {
	// The flow is from `from325` into `into531`.

	// Assume that `sourceCQL` has the underlying type of `from325`:
	from325 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `from325` to the result `into531`
	// (`into531` is now tainted).
	into531 := from325.Hostname()

	// Sink the tainted `into531`:
	sink(into531)
}

func TaintStepTest_NetUrlURLMarshalBinary(sourceCQL interface{}) {
	// The flow is from `from954` into `text`.

	// Assume that `sourceCQL` has the underlying type of `from954`:
	from954 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `from954` to the result `text`
	// (`text` is now tainted).
	text, _ := from954.MarshalBinary()

	// Sink the tainted `text`:
	sink(text)
}

func TaintStepTest_NetUrlURLParse(sourceCQL interface{}) {
	// The flow is from `ref` into `into661`.

	// Assume that `sourceCQL` has the underlying type of `ref`:
	ref := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL url.URL

	// Call the method that transfers the taint
	// from the parameter `ref` to the result `into661`
	// (`into661` is now tainted).
	into661, _ := mediumObjCQL.Parse(ref)

	// Sink the tainted `into661`:
	sink(into661)
}

func TaintStepTest_NetUrlURLPort(sourceCQL interface{}) {
	// The flow is from `from606` into `into288`.

	// Assume that `sourceCQL` has the underlying type of `from606`:
	from606 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `from606` to the result `into288`
	// (`into288` is now tainted).
	into288 := from606.Port()

	// Sink the tainted `into288`:
	sink(into288)
}

func TaintStepTest_NetUrlURLQuery(sourceCQL interface{}) {
	// The flow is from `from292` into `into884`.

	// Assume that `sourceCQL` has the underlying type of `from292`:
	from292 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `from292` to the result `into884`
	// (`into884` is now tainted).
	into884 := from292.Query()

	// Sink the tainted `into884`:
	sink(into884)
}

func TaintStepTest_NetUrlURLRequestURI(sourceCQL interface{}) {
	// The flow is from `from808` into `into543`.

	// Assume that `sourceCQL` has the underlying type of `from808`:
	from808 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `from808` to the result `into543`
	// (`into543` is now tainted).
	into543 := from808.RequestURI()

	// Sink the tainted `into543`:
	sink(into543)
}

func TaintStepTest_NetUrlURLResolveReference(sourceCQL interface{}) {
	// The flow is from `ref` into `into214`.

	// Assume that `sourceCQL` has the underlying type of `ref`:
	ref := sourceCQL.(*url.URL)

	// Declare medium object/interface:
	var mediumObjCQL url.URL

	// Call the method that transfers the taint
	// from the parameter `ref` to the result `into214`
	// (`into214` is now tainted).
	into214 := mediumObjCQL.ResolveReference(ref)

	// Sink the tainted `into214`:
	sink(into214)
}

func TaintStepTest_NetUrlURLString(sourceCQL interface{}) {
	// The flow is from `from336` into `into509`.

	// Assume that `sourceCQL` has the underlying type of `from336`:
	from336 := sourceCQL.(url.URL)

	// Call the method that transfers the taint
	// from the receiver `from336` to the result `into509`
	// (`into509` is now tainted).
	into509 := from336.String()

	// Sink the tainted `into509`:
	sink(into509)
}

func TaintStepTest_NetUrlURLUnmarshalBinary(sourceCQL interface{}) {
	// The flow is from `text` into `into620`.

	// Assume that `sourceCQL` has the underlying type of `text`:
	text := sourceCQL.([]byte)

	// Declare `into620` variable:
	var into620 url.URL

	// Call the method that transfers the taint
	// from the parameter `text` to the receiver `into620`
	// (`into620` is now tainted).
	into620.UnmarshalBinary(text)

	// Sink the tainted `into620`:
	sink(into620)
}

func TaintStepTest_NetUrlUserinfoPassword(sourceCQL interface{}) {
	// The flow is from `from795` into `into335`.

	// Assume that `sourceCQL` has the underlying type of `from795`:
	from795 := sourceCQL.(url.Userinfo)

	// Call the method that transfers the taint
	// from the receiver `from795` to the result `into335`
	// (`into335` is now tainted).
	into335, _ := from795.Password()

	// Sink the tainted `into335`:
	sink(into335)
}

func TaintStepTest_NetUrlUserinfoString(sourceCQL interface{}) {
	// The flow is from `from415` into `into401`.

	// Assume that `sourceCQL` has the underlying type of `from415`:
	from415 := sourceCQL.(url.Userinfo)

	// Call the method that transfers the taint
	// from the receiver `from415` to the result `into401`
	// (`into401` is now tainted).
	into401 := from415.String()

	// Sink the tainted `into401`:
	sink(into401)
}

func TaintStepTest_NetUrlUserinfoUsername(sourceCQL interface{}) {
	// The flow is from `from764` into `into792`.

	// Assume that `sourceCQL` has the underlying type of `from764`:
	from764 := sourceCQL.(url.Userinfo)

	// Call the method that transfers the taint
	// from the receiver `from764` to the result `into792`
	// (`into792` is now tainted).
	into792 := from764.Username()

	// Sink the tainted `into792`:
	sink(into792)
}

func TaintStepTest_NetUrlValuesAdd(sourceCQL interface{}) {
	// The flow is from `value` into `into452`.

	// Assume that `sourceCQL` has the underlying type of `value`:
	value := sourceCQL.(string)

	// Declare `into452` variable:
	var into452 url.Values

	// Call the method that transfers the taint
	// from the parameter `value` to the receiver `into452`
	// (`into452` is now tainted).
	into452.Add("", value)

	// Sink the tainted `into452`:
	sink(into452)
}

func TaintStepTest_NetUrlValuesEncode(sourceCQL interface{}) {
	// The flow is from `from831` into `into575`.

	// Assume that `sourceCQL` has the underlying type of `from831`:
	from831 := sourceCQL.(url.Values)

	// Call the method that transfers the taint
	// from the receiver `from831` to the result `into575`
	// (`into575` is now tainted).
	into575 := from831.Encode()

	// Sink the tainted `into575`:
	sink(into575)
}

func TaintStepTest_NetUrlValuesGet(sourceCQL interface{}) {
	// The flow is from `from789` into `into620`.

	// Assume that `sourceCQL` has the underlying type of `from789`:
	from789 := sourceCQL.(url.Values)

	// Call the method that transfers the taint
	// from the receiver `from789` to the result `into620`
	// (`into620` is now tainted).
	into620 := from789.Get("")

	// Sink the tainted `into620`:
	sink(into620)
}

func TaintStepTest_NetUrlValuesSet(sourceCQL interface{}) {
	// The flow is from `value` into `into434`.

	// Assume that `sourceCQL` has the underlying type of `value`:
	value := sourceCQL.(string)

	// Declare `into434` variable:
	var into434 url.Values

	// Call the method that transfers the taint
	// from the parameter `value` to the receiver `into434`
	// (`into434` is now tainted).
	into434.Set("", value)

	// Sink the tainted `into434`:
	sink(into434)
}

func RunAllTaints_NetUrl(v interface{}) {
	{
		source := newSource()
		TaintStepTest_NetUrlParse(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlParseQuery(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlParseRequestURI(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlPathEscape(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlPathUnescape(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlQueryEscape(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlQueryUnescape(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlUser(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlUserPassword(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLEscapedPath(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLHostname(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLMarshalBinary(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLParse(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLPort(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLQuery(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLRequestURI(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLResolveReference(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLString(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlURLUnmarshalBinary(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlUserinfoPassword(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlUserinfoString(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlUserinfoUsername(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlValuesAdd(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlValuesEncode(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlValuesGet(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUrlValuesSet(source)
	}
}
