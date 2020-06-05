package main

import (
	"bufio"
	"io"
	"net/http"
)

func TaintStepTest_NetHttpCanonicalHeaderKey(sourceCQL interface{}) {
	// The flow is from `s` into `into648`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into648`
	// (`into648` is now tainted).
	into648 := http.CanonicalHeaderKey(s)

	// Sink the tainted `into648`:
	sink(into648)
}

func TaintStepTest_NetHttpDetectContentType(sourceCQL interface{}) {
	// The flow is from `data` into `into210`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `data` to result `into210`
	// (`into210` is now tainted).
	into210 := http.DetectContentType(data)

	// Sink the tainted `into210`:
	sink(into210)
}

func TaintStepTest_NetHttpError(sourceCQL interface{}) {
	// The flow is from `error` into `w`.

	// Assume that `sourceCQL` has the underlying type of `error`:
	error := sourceCQL.(string)

	// Declare `w` variable:
	var w http.ResponseWriter

	// Call the function that transfers the taint
	// from the parameter `error` to parameter `w`;
	// `w` is now tainted.
	http.Error(w, error, 0)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_NetHttpMaxBytesReader(sourceCQL interface{}) {
	// The flow is from `r` into `into857`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.ReadCloser)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into857`
	// (`into857` is now tainted).
	into857 := http.MaxBytesReader(nil, r, 0)

	// Sink the tainted `into857`:
	sink(into857)
}

func TaintStepTest_NetHttpNewRequest(sourceCQL interface{}) {
	// The flow is from `body` into `into865`.

	// Assume that `sourceCQL` has the underlying type of `body`:
	body := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `body` to result `into865`
	// (`into865` is now tainted).
	into865, _ := http.NewRequest("", "", body)

	// Sink the tainted `into865`:
	sink(into865)
}

func TaintStepTest_NetHttpNewRequestWithContext(sourceCQL interface{}) {
	// The flow is from `body` into `into314`.

	// Assume that `sourceCQL` has the underlying type of `body`:
	body := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `body` to result `into314`
	// (`into314` is now tainted).
	into314, _ := http.NewRequestWithContext(nil, "", "", body)

	// Sink the tainted `into314`:
	sink(into314)
}

func TaintStepTest_NetHttpReadRequest(sourceCQL interface{}) {
	// The flow is from `b` into `into746`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `b` to result `into746`
	// (`into746` is now tainted).
	into746, _ := http.ReadRequest(b)

	// Sink the tainted `into746`:
	sink(into746)
}

func TaintStepTest_NetHttpReadResponse(sourceCQL interface{}) {
	// The flow is from `r` into `into253`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into253`
	// (`into253` is now tainted).
	into253, _ := http.ReadResponse(r, nil)

	// Sink the tainted `into253`:
	sink(into253)
}

func TaintStepTest_NetHttpSetCookie(sourceCQL interface{}) {
	// The flow is from `cookie` into `w`.

	// Assume that `sourceCQL` has the underlying type of `cookie`:
	cookie := sourceCQL.(*http.Cookie)

	// Declare `w` variable:
	var w http.ResponseWriter

	// Call the function that transfers the taint
	// from the parameter `cookie` to parameter `w`;
	// `w` is now tainted.
	http.SetCookie(w, cookie)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_NetHttpHeaderAdd(sourceCQL interface{}) {
	// The flow is from `value` into `into627`.

	// Assume that `sourceCQL` has the underlying type of `value`:
	value := sourceCQL.(string)

	// Declare `into627` variable:
	var into627 http.Header

	// Call the method that transfers the taint
	// from the parameter `value` to the receiver `into627`
	// (`into627` is now tainted).
	into627.Add("", value)

	// Sink the tainted `into627`:
	sink(into627)
}

func TaintStepTest_NetHttpHeaderClone(sourceCQL interface{}) {
	// The flow is from `from735` into `into549`.

	// Assume that `sourceCQL` has the underlying type of `from735`:
	from735 := sourceCQL.(http.Header)

	// Call the method that transfers the taint
	// from the receiver `from735` to the result `into549`
	// (`into549` is now tainted).
	into549 := from735.Clone()

	// Sink the tainted `into549`:
	sink(into549)
}

func TaintStepTest_NetHttpHeaderGet(sourceCQL interface{}) {
	// The flow is from `from188` into `into761`.

	// Assume that `sourceCQL` has the underlying type of `from188`:
	from188 := sourceCQL.(http.Header)

	// Call the method that transfers the taint
	// from the receiver `from188` to the result `into761`
	// (`into761` is now tainted).
	into761 := from188.Get("")

	// Sink the tainted `into761`:
	sink(into761)
}

func TaintStepTest_NetHttpHeaderSet(sourceCQL interface{}) {
	// The flow is from `value` into `into875`.

	// Assume that `sourceCQL` has the underlying type of `value`:
	value := sourceCQL.(string)

	// Declare `into875` variable:
	var into875 http.Header

	// Call the method that transfers the taint
	// from the parameter `value` to the receiver `into875`
	// (`into875` is now tainted).
	into875.Set("", value)

	// Sink the tainted `into875`:
	sink(into875)
}

func TaintStepTest_NetHttpHeaderValues(sourceCQL interface{}) {
	// The flow is from `from785` into `into896`.

	// Assume that `sourceCQL` has the underlying type of `from785`:
	from785 := sourceCQL.(http.Header)

	// Call the method that transfers the taint
	// from the receiver `from785` to the result `into896`
	// (`into896` is now tainted).
	into896 := from785.Values("")

	// Sink the tainted `into896`:
	sink(into896)
}

func TaintStepTest_NetHttpHeaderWrite(sourceCQL interface{}) {
	// The flow is from `from169` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from169`:
	from169 := sourceCQL.(http.Header)

	// Declare `w` variable:
	var w io.Writer

	// Call the method that transfers the taint
	// from the receiver `from169` to the argument `w`
	// (`w` is now tainted).
	from169.Write(w)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_NetHttpHeaderWriteSubset(sourceCQL interface{}) {
	// The flow is from `from831` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from831`:
	from831 := sourceCQL.(http.Header)

	// Declare `w` variable:
	var w io.Writer

	// Call the method that transfers the taint
	// from the receiver `from831` to the argument `w`
	// (`w` is now tainted).
	from831.WriteSubset(w, nil)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_NetHttpRequestAddCookie(sourceCQL interface{}) {
	// The flow is from `c` into `into383`.

	// Assume that `sourceCQL` has the underlying type of `c`:
	c := sourceCQL.(*http.Cookie)

	// Declare `into383` variable:
	var into383 http.Request

	// Call the method that transfers the taint
	// from the parameter `c` to the receiver `into383`
	// (`into383` is now tainted).
	into383.AddCookie(c)

	// Sink the tainted `into383`:
	sink(into383)
}

func TaintStepTest_NetHttpRequestClone(sourceCQL interface{}) {
	// The flow is from `from816` into `into294`.

	// Assume that `sourceCQL` has the underlying type of `from816`:
	from816 := sourceCQL.(http.Request)

	// Call the method that transfers the taint
	// from the receiver `from816` to the result `into294`
	// (`into294` is now tainted).
	into294 := from816.Clone(nil)

	// Sink the tainted `into294`:
	sink(into294)
}

func TaintStepTest_NetHttpRequestWrite(sourceCQL interface{}) {
	// The flow is from `from769` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from769`:
	from769 := sourceCQL.(http.Request)

	// Declare `w` variable:
	var w io.Writer

	// Call the method that transfers the taint
	// from the receiver `from769` to the argument `w`
	// (`w` is now tainted).
	from769.Write(w)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_NetHttpRequestWriteProxy(sourceCQL interface{}) {
	// The flow is from `from296` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from296`:
	from296 := sourceCQL.(http.Request)

	// Declare `w` variable:
	var w io.Writer

	// Call the method that transfers the taint
	// from the receiver `from296` to the argument `w`
	// (`w` is now tainted).
	from296.WriteProxy(w)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_NetHttpResponseWrite(sourceCQL interface{}) {
	// The flow is from `from747` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from747`:
	from747 := sourceCQL.(http.Response)

	// Declare `w` variable:
	var w io.Writer

	// Call the method that transfers the taint
	// from the receiver `from747` to the argument `w`
	// (`w` is now tainted).
	from747.Write(w)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_NetHttpResponseWriterWrite(sourceCQL interface{}) {
	// The flow is from `from705` into `into521`.

	// Assume that `sourceCQL` has the underlying type of `from705`:
	from705 := sourceCQL.([]byte)

	// Declare `into521` variable:
	var into521 http.ResponseWriter

	// Call the method that transfers the taint
	// from the parameter `from705` to the receiver `into521`
	// (`into521` is now tainted).
	into521.Write(from705)

	// Sink the tainted `into521`:
	sink(into521)
}

func RunAllTaints_NetHttp(v interface{}) {
	{
		source := newSource()
		TaintStepTest_NetHttpCanonicalHeaderKey(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpDetectContentType(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpError(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpMaxBytesReader(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpNewRequest(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpNewRequestWithContext(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpReadRequest(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpReadResponse(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpSetCookie(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHeaderAdd(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHeaderClone(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHeaderGet(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHeaderSet(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHeaderValues(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHeaderWrite(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHeaderWriteSubset(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpRequestAddCookie(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpRequestClone(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpRequestWrite(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpRequestWriteProxy(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpResponseWrite(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpResponseWriterWrite(source)
	}
}
