// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"bufio"
	"context"
	"io"
	"net/http"
)

func TaintStepTest_NetHttpCanonicalHeaderKey_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString656` into `intoString414`.

	// Assume that `sourceCQL` has the underlying type of `fromString656`:
	fromString656 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString656` to result `intoString414`
	// (`intoString414` is now tainted).
	intoString414 := http.CanonicalHeaderKey(fromString656)

	// Return the tainted `intoString414`:
	return intoString414
}

func TaintStepTest_NetHttpDetectContentType_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte518` into `intoString650`.

	// Assume that `sourceCQL` has the underlying type of `fromByte518`:
	fromByte518 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte518` to result `intoString650`
	// (`intoString650` is now tainted).
	intoString650 := http.DetectContentType(fromByte518)

	// Return the tainted `intoString650`:
	return intoString650
}

func TaintStepTest_NetHttpError_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString784` into `intoResponseWriter957`.

	// Assume that `sourceCQL` has the underlying type of `fromString784`:
	fromString784 := sourceCQL.(string)

	// Declare `intoResponseWriter957` variable:
	var intoResponseWriter957 http.ResponseWriter

	// Call the function that transfers the taint
	// from the parameter `fromString784` to parameter `intoResponseWriter957`;
	// `intoResponseWriter957` is now tainted.
	http.Error(intoResponseWriter957, fromString784, 0)

	// Return the tainted `intoResponseWriter957`:
	return intoResponseWriter957
}

func TaintStepTest_NetHttpMaxBytesReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReadCloser520` into `intoReadCloser443`.

	// Assume that `sourceCQL` has the underlying type of `fromReadCloser520`:
	fromReadCloser520 := sourceCQL.(io.ReadCloser)

	// Call the function that transfers the taint
	// from the parameter `fromReadCloser520` to result `intoReadCloser443`
	// (`intoReadCloser443` is now tainted).
	intoReadCloser443 := http.MaxBytesReader(nil, fromReadCloser520, 0)

	// Return the tainted `intoReadCloser443`:
	return intoReadCloser443
}

func TaintStepTest_NetHttpNewRequest_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString127` into `intoRequest483`.

	// Assume that `sourceCQL` has the underlying type of `fromString127`:
	fromString127 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString127` to result `intoRequest483`
	// (`intoRequest483` is now tainted).
	intoRequest483, _ := http.NewRequest(fromString127, "", nil)

	// Return the tainted `intoRequest483`:
	return intoRequest483
}

func TaintStepTest_NetHttpNewRequest_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString989` into `intoRequest982`.

	// Assume that `sourceCQL` has the underlying type of `fromString989`:
	fromString989 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString989` to result `intoRequest982`
	// (`intoRequest982` is now tainted).
	intoRequest982, _ := http.NewRequest("", fromString989, nil)

	// Return the tainted `intoRequest982`:
	return intoRequest982
}

func TaintStepTest_NetHttpNewRequest_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader417` into `intoRequest584`.

	// Assume that `sourceCQL` has the underlying type of `fromReader417`:
	fromReader417 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader417` to result `intoRequest584`
	// (`intoRequest584` is now tainted).
	intoRequest584, _ := http.NewRequest("", "", fromReader417)

	// Return the tainted `intoRequest584`:
	return intoRequest584
}

func TaintStepTest_NetHttpNewRequestWithContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext991` into `intoRequest881`.

	// Assume that `sourceCQL` has the underlying type of `fromContext991`:
	fromContext991 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext991` to result `intoRequest881`
	// (`intoRequest881` is now tainted).
	intoRequest881, _ := http.NewRequestWithContext(fromContext991, "", "", nil)

	// Return the tainted `intoRequest881`:
	return intoRequest881
}

func TaintStepTest_NetHttpNewRequestWithContext_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString186` into `intoRequest284`.

	// Assume that `sourceCQL` has the underlying type of `fromString186`:
	fromString186 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString186` to result `intoRequest284`
	// (`intoRequest284` is now tainted).
	intoRequest284, _ := http.NewRequestWithContext(nil, fromString186, "", nil)

	// Return the tainted `intoRequest284`:
	return intoRequest284
}

func TaintStepTest_NetHttpNewRequestWithContext_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString908` into `intoRequest137`.

	// Assume that `sourceCQL` has the underlying type of `fromString908`:
	fromString908 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString908` to result `intoRequest137`
	// (`intoRequest137` is now tainted).
	intoRequest137, _ := http.NewRequestWithContext(nil, "", fromString908, nil)

	// Return the tainted `intoRequest137`:
	return intoRequest137
}

func TaintStepTest_NetHttpNewRequestWithContext_B0I3O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader494` into `intoRequest873`.

	// Assume that `sourceCQL` has the underlying type of `fromReader494`:
	fromReader494 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader494` to result `intoRequest873`
	// (`intoRequest873` is now tainted).
	intoRequest873, _ := http.NewRequestWithContext(nil, "", "", fromReader494)

	// Return the tainted `intoRequest873`:
	return intoRequest873
}

func TaintStepTest_NetHttpReadRequest_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader599` into `intoRequest409`.

	// Assume that `sourceCQL` has the underlying type of `fromReader599`:
	fromReader599 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader599` to result `intoRequest409`
	// (`intoRequest409` is now tainted).
	intoRequest409, _ := http.ReadRequest(fromReader599)

	// Return the tainted `intoRequest409`:
	return intoRequest409
}

func TaintStepTest_NetHttpReadResponse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader246` into `intoResponse898`.

	// Assume that `sourceCQL` has the underlying type of `fromReader246`:
	fromReader246 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader246` to result `intoResponse898`
	// (`intoResponse898` is now tainted).
	intoResponse898, _ := http.ReadResponse(fromReader246, nil)

	// Return the tainted `intoResponse898`:
	return intoResponse898
}

func TaintStepTest_NetHttpSetCookie_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromCookie598` into `intoResponseWriter631`.

	// Assume that `sourceCQL` has the underlying type of `fromCookie598`:
	fromCookie598 := sourceCQL.(*http.Cookie)

	// Declare `intoResponseWriter631` variable:
	var intoResponseWriter631 http.ResponseWriter

	// Call the function that transfers the taint
	// from the parameter `fromCookie598` to parameter `intoResponseWriter631`;
	// `intoResponseWriter631` is now tainted.
	http.SetCookie(intoResponseWriter631, fromCookie598)

	// Return the tainted `intoResponseWriter631`:
	return intoResponseWriter631
}

func TaintStepTest_NetHttpHeaderAdd_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString165` into `intoHeader150`.

	// Assume that `sourceCQL` has the underlying type of `fromString165`:
	fromString165 := sourceCQL.(string)

	// Declare `intoHeader150` variable:
	var intoHeader150 http.Header

	// Call the method that transfers the taint
	// from the parameter `fromString165` to the receiver `intoHeader150`
	// (`intoHeader150` is now tainted).
	intoHeader150.Add(fromString165, "")

	// Return the tainted `intoHeader150`:
	return intoHeader150
}

func TaintStepTest_NetHttpHeaderAdd_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString340` into `intoHeader471`.

	// Assume that `sourceCQL` has the underlying type of `fromString340`:
	fromString340 := sourceCQL.(string)

	// Declare `intoHeader471` variable:
	var intoHeader471 http.Header

	// Call the method that transfers the taint
	// from the parameter `fromString340` to the receiver `intoHeader471`
	// (`intoHeader471` is now tainted).
	intoHeader471.Add("", fromString340)

	// Return the tainted `intoHeader471`:
	return intoHeader471
}

func TaintStepTest_NetHttpHeaderClone_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader290` into `intoHeader758`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader290`:
	fromHeader290 := sourceCQL.(http.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader290` to the result `intoHeader758`
	// (`intoHeader758` is now tainted).
	intoHeader758 := fromHeader290.Clone()

	// Return the tainted `intoHeader758`:
	return intoHeader758
}

func TaintStepTest_NetHttpHeaderGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader396` into `intoString707`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader396`:
	fromHeader396 := sourceCQL.(http.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader396` to the result `intoString707`
	// (`intoString707` is now tainted).
	intoString707 := fromHeader396.Get("")

	// Return the tainted `intoString707`:
	return intoString707
}

func TaintStepTest_NetHttpHeaderSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString912` into `intoHeader718`.

	// Assume that `sourceCQL` has the underlying type of `fromString912`:
	fromString912 := sourceCQL.(string)

	// Declare `intoHeader718` variable:
	var intoHeader718 http.Header

	// Call the method that transfers the taint
	// from the parameter `fromString912` to the receiver `intoHeader718`
	// (`intoHeader718` is now tainted).
	intoHeader718.Set(fromString912, "")

	// Return the tainted `intoHeader718`:
	return intoHeader718
}

func TaintStepTest_NetHttpHeaderSet_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString972` into `intoHeader633`.

	// Assume that `sourceCQL` has the underlying type of `fromString972`:
	fromString972 := sourceCQL.(string)

	// Declare `intoHeader633` variable:
	var intoHeader633 http.Header

	// Call the method that transfers the taint
	// from the parameter `fromString972` to the receiver `intoHeader633`
	// (`intoHeader633` is now tainted).
	intoHeader633.Set("", fromString972)

	// Return the tainted `intoHeader633`:
	return intoHeader633
}

func TaintStepTest_NetHttpHeaderValues_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader316` into `intoString145`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader316`:
	fromHeader316 := sourceCQL.(http.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader316` to the result `intoString145`
	// (`intoString145` is now tainted).
	intoString145 := fromHeader316.Values("")

	// Return the tainted `intoString145`:
	return intoString145
}

func TaintStepTest_NetHttpHeaderWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader817` into `intoWriter474`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader817`:
	fromHeader817 := sourceCQL.(http.Header)

	// Declare `intoWriter474` variable:
	var intoWriter474 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromHeader817` to the argument `intoWriter474`
	// (`intoWriter474` is now tainted).
	fromHeader817.Write(intoWriter474)

	// Return the tainted `intoWriter474`:
	return intoWriter474
}

func TaintStepTest_NetHttpHeaderWriteSubset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader832` into `intoWriter378`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader832`:
	fromHeader832 := sourceCQL.(http.Header)

	// Declare `intoWriter378` variable:
	var intoWriter378 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromHeader832` to the argument `intoWriter378`
	// (`intoWriter378` is now tainted).
	fromHeader832.WriteSubset(intoWriter378, nil)

	// Return the tainted `intoWriter378`:
	return intoWriter378
}

func TaintStepTest_NetHttpRequestAddCookie_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromCookie541` into `intoRequest139`.

	// Assume that `sourceCQL` has the underlying type of `fromCookie541`:
	fromCookie541 := sourceCQL.(*http.Cookie)

	// Declare `intoRequest139` variable:
	var intoRequest139 http.Request

	// Call the method that transfers the taint
	// from the parameter `fromCookie541` to the receiver `intoRequest139`
	// (`intoRequest139` is now tainted).
	intoRequest139.AddCookie(fromCookie541)

	// Return the tainted `intoRequest139`:
	return intoRequest139
}

func TaintStepTest_NetHttpRequestClone_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRequest814` into `intoRequest768`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest814`:
	fromRequest814 := sourceCQL.(http.Request)

	// Call the method that transfers the taint
	// from the receiver `fromRequest814` to the result `intoRequest768`
	// (`intoRequest768` is now tainted).
	intoRequest768 := fromRequest814.Clone(nil)

	// Return the tainted `intoRequest768`:
	return intoRequest768
}

func TaintStepTest_NetHttpRequestWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRequest468` into `intoWriter736`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest468`:
	fromRequest468 := sourceCQL.(http.Request)

	// Declare `intoWriter736` variable:
	var intoWriter736 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromRequest468` to the argument `intoWriter736`
	// (`intoWriter736` is now tainted).
	fromRequest468.Write(intoWriter736)

	// Return the tainted `intoWriter736`:
	return intoWriter736
}

func TaintStepTest_NetHttpRequestWriteProxy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRequest516` into `intoWriter246`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest516`:
	fromRequest516 := sourceCQL.(http.Request)

	// Declare `intoWriter246` variable:
	var intoWriter246 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromRequest516` to the argument `intoWriter246`
	// (`intoWriter246` is now tainted).
	fromRequest516.WriteProxy(intoWriter246)

	// Return the tainted `intoWriter246`:
	return intoWriter246
}

func TaintStepTest_NetHttpResponseWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromResponse679` into `intoWriter736`.

	// Assume that `sourceCQL` has the underlying type of `fromResponse679`:
	fromResponse679 := sourceCQL.(http.Response)

	// Declare `intoWriter736` variable:
	var intoWriter736 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromResponse679` to the argument `intoWriter736`
	// (`intoWriter736` is now tainted).
	fromResponse679.Write(intoWriter736)

	// Return the tainted `intoWriter736`:
	return intoWriter736
}

func TaintStepTest_NetHttpHijackerHijack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHijacker839` into `intoConn273`.

	// Assume that `sourceCQL` has the underlying type of `fromHijacker839`:
	fromHijacker839 := sourceCQL.(http.Hijacker)

	// Call the method that transfers the taint
	// from the receiver `fromHijacker839` to the result `intoConn273`
	// (`intoConn273` is now tainted).
	intoConn273, _, _ := fromHijacker839.Hijack()

	// Return the tainted `intoConn273`:
	return intoConn273
}

func TaintStepTest_NetHttpHijackerHijack_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromHijacker982` into `intoReadWriter458`.

	// Assume that `sourceCQL` has the underlying type of `fromHijacker982`:
	fromHijacker982 := sourceCQL.(http.Hijacker)

	// Call the method that transfers the taint
	// from the receiver `fromHijacker982` to the result `intoReadWriter458`
	// (`intoReadWriter458` is now tainted).
	_, intoReadWriter458, _ := fromHijacker982.Hijack()

	// Return the tainted `intoReadWriter458`:
	return intoReadWriter458
}

func TaintStepTest_NetHttpResponseWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte506` into `intoResponseWriter213`.

	// Assume that `sourceCQL` has the underlying type of `fromByte506`:
	fromByte506 := sourceCQL.([]byte)

	// Declare `intoResponseWriter213` variable:
	var intoResponseWriter213 http.ResponseWriter

	// Call the method that transfers the taint
	// from the parameter `fromByte506` to the receiver `intoResponseWriter213`
	// (`intoResponseWriter213` is now tainted).
	intoResponseWriter213.Write(fromByte506)

	// Return the tainted `intoResponseWriter213`:
	return intoResponseWriter213
}

func RunAllTaints_NetHttp() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpCanonicalHeaderKey_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpDetectContentType_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpError_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpMaxBytesReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpNewRequest_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpNewRequest_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpNewRequest_B0I2O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpNewRequestWithContext_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpNewRequestWithContext_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpNewRequestWithContext_B0I2O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpNewRequestWithContext_B0I3O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpReadRequest_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpReadResponse_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpSetCookie_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHeaderAdd_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHeaderAdd_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHeaderClone_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHeaderGet_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHeaderSet_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHeaderSet_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHeaderValues_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHeaderWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHeaderWriteSubset_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpRequestAddCookie_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpRequestClone_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpRequestWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpRequestWriteProxy_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpResponseWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHijackerHijack_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpHijackerHijack_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHttpResponseWriterWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
