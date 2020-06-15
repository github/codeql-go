package main

import (
	"bufio"
	"context"
	"io"
	"net/http"
)

func TaintStepTest_NetHttpCanonicalHeaderKey_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString935` into `intoString364`.

	// Assume that `sourceCQL` has the underlying type of `fromString935`:
	fromString935 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString935` to result `intoString364`
	// (`intoString364` is now tainted).
	intoString364 := http.CanonicalHeaderKey(fromString935)

	// Sink the tainted `intoString364`:
	sink(intoString364)
}

func TaintStepTest_NetHttpDetectContentType_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte385` into `intoString640`.

	// Assume that `sourceCQL` has the underlying type of `fromByte385`:
	fromByte385 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte385` to result `intoString640`
	// (`intoString640` is now tainted).
	intoString640 := http.DetectContentType(fromByte385)

	// Sink the tainted `intoString640`:
	sink(intoString640)
}

func TaintStepTest_NetHttpError_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString919` into `intoResponseWriter882`.

	// Assume that `sourceCQL` has the underlying type of `fromString919`:
	fromString919 := sourceCQL.(string)

	// Declare `intoResponseWriter882` variable:
	var intoResponseWriter882 http.ResponseWriter

	// Call the function that transfers the taint
	// from the parameter `fromString919` to parameter `intoResponseWriter882`;
	// `intoResponseWriter882` is now tainted.
	http.Error(intoResponseWriter882, fromString919, 0)

	// Sink the tainted `intoResponseWriter882`:
	sink(intoResponseWriter882)
}

func TaintStepTest_NetHttpMaxBytesReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReadCloser301` into `intoReadCloser514`.

	// Assume that `sourceCQL` has the underlying type of `fromReadCloser301`:
	fromReadCloser301 := sourceCQL.(io.ReadCloser)

	// Call the function that transfers the taint
	// from the parameter `fromReadCloser301` to result `intoReadCloser514`
	// (`intoReadCloser514` is now tainted).
	intoReadCloser514 := http.MaxBytesReader(nil, fromReadCloser301, 0)

	// Sink the tainted `intoReadCloser514`:
	sink(intoReadCloser514)
}

func TaintStepTest_NetHttpNewRequest_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString344` into `intoRequest663`.

	// Assume that `sourceCQL` has the underlying type of `fromString344`:
	fromString344 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString344` to result `intoRequest663`
	// (`intoRequest663` is now tainted).
	intoRequest663, _ := http.NewRequest(fromString344, "", nil)

	// Sink the tainted `intoRequest663`:
	sink(intoRequest663)
}

func TaintStepTest_NetHttpNewRequest_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString414` into `intoRequest572`.

	// Assume that `sourceCQL` has the underlying type of `fromString414`:
	fromString414 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString414` to result `intoRequest572`
	// (`intoRequest572` is now tainted).
	intoRequest572, _ := http.NewRequest("", fromString414, nil)

	// Sink the tainted `intoRequest572`:
	sink(intoRequest572)
}

func TaintStepTest_NetHttpNewRequest_B0I2O0(sourceCQL interface{}) {
	// The flow is from `fromReader802` into `intoRequest998`.

	// Assume that `sourceCQL` has the underlying type of `fromReader802`:
	fromReader802 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader802` to result `intoRequest998`
	// (`intoRequest998` is now tainted).
	intoRequest998, _ := http.NewRequest("", "", fromReader802)

	// Sink the tainted `intoRequest998`:
	sink(intoRequest998)
}

func TaintStepTest_NetHttpNewRequestWithContext_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromContext924` into `intoRequest750`.

	// Assume that `sourceCQL` has the underlying type of `fromContext924`:
	fromContext924 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext924` to result `intoRequest750`
	// (`intoRequest750` is now tainted).
	intoRequest750, _ := http.NewRequestWithContext(fromContext924, "", "", nil)

	// Sink the tainted `intoRequest750`:
	sink(intoRequest750)
}

func TaintStepTest_NetHttpNewRequestWithContext_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString837` into `intoRequest349`.

	// Assume that `sourceCQL` has the underlying type of `fromString837`:
	fromString837 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString837` to result `intoRequest349`
	// (`intoRequest349` is now tainted).
	intoRequest349, _ := http.NewRequestWithContext(nil, fromString837, "", nil)

	// Sink the tainted `intoRequest349`:
	sink(intoRequest349)
}

func TaintStepTest_NetHttpNewRequestWithContext_B0I2O0(sourceCQL interface{}) {
	// The flow is from `fromString186` into `intoRequest145`.

	// Assume that `sourceCQL` has the underlying type of `fromString186`:
	fromString186 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString186` to result `intoRequest145`
	// (`intoRequest145` is now tainted).
	intoRequest145, _ := http.NewRequestWithContext(nil, "", fromString186, nil)

	// Sink the tainted `intoRequest145`:
	sink(intoRequest145)
}

func TaintStepTest_NetHttpNewRequestWithContext_B0I3O0(sourceCQL interface{}) {
	// The flow is from `fromReader648` into `intoRequest372`.

	// Assume that `sourceCQL` has the underlying type of `fromReader648`:
	fromReader648 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader648` to result `intoRequest372`
	// (`intoRequest372` is now tainted).
	intoRequest372, _ := http.NewRequestWithContext(nil, "", "", fromReader648)

	// Sink the tainted `intoRequest372`:
	sink(intoRequest372)
}

func TaintStepTest_NetHttpReadRequest_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader488` into `intoRequest186`.

	// Assume that `sourceCQL` has the underlying type of `fromReader488`:
	fromReader488 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader488` to result `intoRequest186`
	// (`intoRequest186` is now tainted).
	intoRequest186, _ := http.ReadRequest(fromReader488)

	// Sink the tainted `intoRequest186`:
	sink(intoRequest186)
}

func TaintStepTest_NetHttpReadResponse_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader912` into `intoResponse706`.

	// Assume that `sourceCQL` has the underlying type of `fromReader912`:
	fromReader912 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader912` to result `intoResponse706`
	// (`intoResponse706` is now tainted).
	intoResponse706, _ := http.ReadResponse(fromReader912, nil)

	// Sink the tainted `intoResponse706`:
	sink(intoResponse706)
}

func TaintStepTest_NetHttpSetCookie_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromCookie982` into `intoResponseWriter371`.

	// Assume that `sourceCQL` has the underlying type of `fromCookie982`:
	fromCookie982 := sourceCQL.(*http.Cookie)

	// Declare `intoResponseWriter371` variable:
	var intoResponseWriter371 http.ResponseWriter

	// Call the function that transfers the taint
	// from the parameter `fromCookie982` to parameter `intoResponseWriter371`;
	// `intoResponseWriter371` is now tainted.
	http.SetCookie(intoResponseWriter371, fromCookie982)

	// Sink the tainted `intoResponseWriter371`:
	sink(intoResponseWriter371)
}

func TaintStepTest_NetHttpHeaderAdd_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString314` into `intoHeader493`.

	// Assume that `sourceCQL` has the underlying type of `fromString314`:
	fromString314 := sourceCQL.(string)

	// Declare `intoHeader493` variable:
	var intoHeader493 http.Header

	// Call the method that transfers the taint
	// from the parameter `fromString314` to the receiver `intoHeader493`
	// (`intoHeader493` is now tainted).
	intoHeader493.Add(fromString314, "")

	// Sink the tainted `intoHeader493`:
	sink(intoHeader493)
}

func TaintStepTest_NetHttpHeaderAdd_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString847` into `intoHeader265`.

	// Assume that `sourceCQL` has the underlying type of `fromString847`:
	fromString847 := sourceCQL.(string)

	// Declare `intoHeader265` variable:
	var intoHeader265 http.Header

	// Call the method that transfers the taint
	// from the parameter `fromString847` to the receiver `intoHeader265`
	// (`intoHeader265` is now tainted).
	intoHeader265.Add("", fromString847)

	// Sink the tainted `intoHeader265`:
	sink(intoHeader265)
}

func TaintStepTest_NetHttpHeaderClone_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromHeader456` into `intoHeader418`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader456`:
	fromHeader456 := sourceCQL.(http.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader456` to the result `intoHeader418`
	// (`intoHeader418` is now tainted).
	intoHeader418 := fromHeader456.Clone()

	// Sink the tainted `intoHeader418`:
	sink(intoHeader418)
}

func TaintStepTest_NetHttpHeaderGet_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromHeader614` into `intoString285`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader614`:
	fromHeader614 := sourceCQL.(http.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader614` to the result `intoString285`
	// (`intoString285` is now tainted).
	intoString285 := fromHeader614.Get("")

	// Sink the tainted `intoString285`:
	sink(intoString285)
}

func TaintStepTest_NetHttpHeaderSet_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString339` into `intoHeader285`.

	// Assume that `sourceCQL` has the underlying type of `fromString339`:
	fromString339 := sourceCQL.(string)

	// Declare `intoHeader285` variable:
	var intoHeader285 http.Header

	// Call the method that transfers the taint
	// from the parameter `fromString339` to the receiver `intoHeader285`
	// (`intoHeader285` is now tainted).
	intoHeader285.Set(fromString339, "")

	// Sink the tainted `intoHeader285`:
	sink(intoHeader285)
}

func TaintStepTest_NetHttpHeaderSet_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString487` into `intoHeader727`.

	// Assume that `sourceCQL` has the underlying type of `fromString487`:
	fromString487 := sourceCQL.(string)

	// Declare `intoHeader727` variable:
	var intoHeader727 http.Header

	// Call the method that transfers the taint
	// from the parameter `fromString487` to the receiver `intoHeader727`
	// (`intoHeader727` is now tainted).
	intoHeader727.Set("", fromString487)

	// Sink the tainted `intoHeader727`:
	sink(intoHeader727)
}

func TaintStepTest_NetHttpHeaderValues_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromHeader402` into `intoString308`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader402`:
	fromHeader402 := sourceCQL.(http.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader402` to the result `intoString308`
	// (`intoString308` is now tainted).
	intoString308 := fromHeader402.Values("")

	// Sink the tainted `intoString308`:
	sink(intoString308)
}

func TaintStepTest_NetHttpHeaderWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromHeader126` into `intoWriter123`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader126`:
	fromHeader126 := sourceCQL.(http.Header)

	// Declare `intoWriter123` variable:
	var intoWriter123 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromHeader126` to the argument `intoWriter123`
	// (`intoWriter123` is now tainted).
	fromHeader126.Write(intoWriter123)

	// Sink the tainted `intoWriter123`:
	sink(intoWriter123)
}

func TaintStepTest_NetHttpHeaderWriteSubset_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromHeader743` into `intoWriter498`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader743`:
	fromHeader743 := sourceCQL.(http.Header)

	// Declare `intoWriter498` variable:
	var intoWriter498 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromHeader743` to the argument `intoWriter498`
	// (`intoWriter498` is now tainted).
	fromHeader743.WriteSubset(intoWriter498, nil)

	// Sink the tainted `intoWriter498`:
	sink(intoWriter498)
}

func TaintStepTest_NetHttpRequestAddCookie_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromCookie980` into `intoRequest820`.

	// Assume that `sourceCQL` has the underlying type of `fromCookie980`:
	fromCookie980 := sourceCQL.(*http.Cookie)

	// Declare `intoRequest820` variable:
	var intoRequest820 http.Request

	// Call the method that transfers the taint
	// from the parameter `fromCookie980` to the receiver `intoRequest820`
	// (`intoRequest820` is now tainted).
	intoRequest820.AddCookie(fromCookie980)

	// Sink the tainted `intoRequest820`:
	sink(intoRequest820)
}

func TaintStepTest_NetHttpRequestClone_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRequest648` into `intoRequest835`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest648`:
	fromRequest648 := sourceCQL.(http.Request)

	// Call the method that transfers the taint
	// from the receiver `fromRequest648` to the result `intoRequest835`
	// (`intoRequest835` is now tainted).
	intoRequest835 := fromRequest648.Clone(nil)

	// Sink the tainted `intoRequest835`:
	sink(intoRequest835)
}

func TaintStepTest_NetHttpRequestWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRequest512` into `intoWriter220`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest512`:
	fromRequest512 := sourceCQL.(http.Request)

	// Declare `intoWriter220` variable:
	var intoWriter220 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromRequest512` to the argument `intoWriter220`
	// (`intoWriter220` is now tainted).
	fromRequest512.Write(intoWriter220)

	// Sink the tainted `intoWriter220`:
	sink(intoWriter220)
}

func TaintStepTest_NetHttpRequestWriteProxy_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRequest871` into `intoWriter886`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest871`:
	fromRequest871 := sourceCQL.(http.Request)

	// Declare `intoWriter886` variable:
	var intoWriter886 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromRequest871` to the argument `intoWriter886`
	// (`intoWriter886` is now tainted).
	fromRequest871.WriteProxy(intoWriter886)

	// Sink the tainted `intoWriter886`:
	sink(intoWriter886)
}

func TaintStepTest_NetHttpResponseWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromResponse156` into `intoWriter568`.

	// Assume that `sourceCQL` has the underlying type of `fromResponse156`:
	fromResponse156 := sourceCQL.(http.Response)

	// Declare `intoWriter568` variable:
	var intoWriter568 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromResponse156` to the argument `intoWriter568`
	// (`intoWriter568` is now tainted).
	fromResponse156.Write(intoWriter568)

	// Sink the tainted `intoWriter568`:
	sink(intoWriter568)
}

func TaintStepTest_NetHttpHijackerHijack_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromHijacker901` into `intoConn830`.

	// Assume that `sourceCQL` has the underlying type of `fromHijacker901`:
	fromHijacker901 := sourceCQL.(http.Hijacker)

	// Call the method that transfers the taint
	// from the receiver `fromHijacker901` to the result `intoConn830`
	// (`intoConn830` is now tainted).
	intoConn830, _, _ := fromHijacker901.Hijack()

	// Sink the tainted `intoConn830`:
	sink(intoConn830)
}

func TaintStepTest_NetHttpHijackerHijack_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromHijacker504` into `intoReadWriter830`.

	// Assume that `sourceCQL` has the underlying type of `fromHijacker504`:
	fromHijacker504 := sourceCQL.(http.Hijacker)

	// Call the method that transfers the taint
	// from the receiver `fromHijacker504` to the result `intoReadWriter830`
	// (`intoReadWriter830` is now tainted).
	_, intoReadWriter830, _ := fromHijacker504.Hijack()

	// Sink the tainted `intoReadWriter830`:
	sink(intoReadWriter830)
}

func TaintStepTest_NetHttpResponseWriterWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte795` into `intoResponseWriter837`.

	// Assume that `sourceCQL` has the underlying type of `fromByte795`:
	fromByte795 := sourceCQL.([]byte)

	// Declare `intoResponseWriter837` variable:
	var intoResponseWriter837 http.ResponseWriter

	// Call the method that transfers the taint
	// from the parameter `fromByte795` to the receiver `intoResponseWriter837`
	// (`intoResponseWriter837` is now tainted).
	intoResponseWriter837.Write(fromByte795)

	// Sink the tainted `intoResponseWriter837`:
	sink(intoResponseWriter837)
}

func RunAllTaints_NetHttp() {
	{
		source := newSource()
		TaintStepTest_NetHttpCanonicalHeaderKey_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpDetectContentType_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpError_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpMaxBytesReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpNewRequest_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpNewRequest_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpNewRequest_B0I2O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpNewRequestWithContext_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpNewRequestWithContext_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpNewRequestWithContext_B0I2O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpNewRequestWithContext_B0I3O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpReadRequest_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpReadResponse_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpSetCookie_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHeaderAdd_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHeaderAdd_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHeaderClone_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHeaderGet_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHeaderSet_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHeaderSet_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHeaderValues_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHeaderWrite_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHeaderWriteSubset_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpRequestAddCookie_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpRequestClone_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpRequestWrite_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpRequestWriteProxy_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpResponseWrite_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHijackerHijack_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpHijackerHijack_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHttpResponseWriterWrite_B0I0O0(source)
	}
}
