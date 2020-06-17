// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"bufio"
	"context"
	"io"
	"net/http"
)

func TaintStepTest_NetHttpCanonicalHeaderKey_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString487` into `intoString830`.

	// Assume that `sourceCQL` has the underlying type of `fromString487`:
	fromString487 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString487` to result `intoString830`
	// (`intoString830` is now tainted).
	intoString830 := http.CanonicalHeaderKey(fromString487)

	// Return the tainted `intoString830`:
	return intoString830
}

func TaintStepTest_NetHttpDetectContentType_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte915` into `intoString410`.

	// Assume that `sourceCQL` has the underlying type of `fromByte915`:
	fromByte915 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte915` to result `intoString410`
	// (`intoString410` is now tainted).
	intoString410 := http.DetectContentType(fromByte915)

	// Return the tainted `intoString410`:
	return intoString410
}

func TaintStepTest_NetHttpError_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString559` into `intoResponseWriter942`.

	// Assume that `sourceCQL` has the underlying type of `fromString559`:
	fromString559 := sourceCQL.(string)

	// Declare `intoResponseWriter942` variable:
	var intoResponseWriter942 http.ResponseWriter

	// Call the function that transfers the taint
	// from the parameter `fromString559` to parameter `intoResponseWriter942`;
	// `intoResponseWriter942` is now tainted.
	http.Error(intoResponseWriter942, fromString559, 0)

	// Return the tainted `intoResponseWriter942`:
	return intoResponseWriter942
}

func TaintStepTest_NetHttpMaxBytesReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReadCloser332` into `intoReadCloser726`.

	// Assume that `sourceCQL` has the underlying type of `fromReadCloser332`:
	fromReadCloser332 := sourceCQL.(io.ReadCloser)

	// Call the function that transfers the taint
	// from the parameter `fromReadCloser332` to result `intoReadCloser726`
	// (`intoReadCloser726` is now tainted).
	intoReadCloser726 := http.MaxBytesReader(nil, fromReadCloser332, 0)

	// Return the tainted `intoReadCloser726`:
	return intoReadCloser726
}

func TaintStepTest_NetHttpNewRequest_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString519` into `intoRequest450`.

	// Assume that `sourceCQL` has the underlying type of `fromString519`:
	fromString519 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString519` to result `intoRequest450`
	// (`intoRequest450` is now tainted).
	intoRequest450, _ := http.NewRequest(fromString519, "", nil)

	// Return the tainted `intoRequest450`:
	return intoRequest450
}

func TaintStepTest_NetHttpNewRequest_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString560` into `intoRequest916`.

	// Assume that `sourceCQL` has the underlying type of `fromString560`:
	fromString560 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString560` to result `intoRequest916`
	// (`intoRequest916` is now tainted).
	intoRequest916, _ := http.NewRequest("", fromString560, nil)

	// Return the tainted `intoRequest916`:
	return intoRequest916
}

func TaintStepTest_NetHttpNewRequest_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader989` into `intoRequest222`.

	// Assume that `sourceCQL` has the underlying type of `fromReader989`:
	fromReader989 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader989` to result `intoRequest222`
	// (`intoRequest222` is now tainted).
	intoRequest222, _ := http.NewRequest("", "", fromReader989)

	// Return the tainted `intoRequest222`:
	return intoRequest222
}

func TaintStepTest_NetHttpNewRequestWithContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext801` into `intoRequest450`.

	// Assume that `sourceCQL` has the underlying type of `fromContext801`:
	fromContext801 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext801` to result `intoRequest450`
	// (`intoRequest450` is now tainted).
	intoRequest450, _ := http.NewRequestWithContext(fromContext801, "", "", nil)

	// Return the tainted `intoRequest450`:
	return intoRequest450
}

func TaintStepTest_NetHttpNewRequestWithContext_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString460` into `intoRequest189`.

	// Assume that `sourceCQL` has the underlying type of `fromString460`:
	fromString460 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString460` to result `intoRequest189`
	// (`intoRequest189` is now tainted).
	intoRequest189, _ := http.NewRequestWithContext(nil, fromString460, "", nil)

	// Return the tainted `intoRequest189`:
	return intoRequest189
}

func TaintStepTest_NetHttpNewRequestWithContext_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString178` into `intoRequest341`.

	// Assume that `sourceCQL` has the underlying type of `fromString178`:
	fromString178 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString178` to result `intoRequest341`
	// (`intoRequest341` is now tainted).
	intoRequest341, _ := http.NewRequestWithContext(nil, "", fromString178, nil)

	// Return the tainted `intoRequest341`:
	return intoRequest341
}

func TaintStepTest_NetHttpNewRequestWithContext_B0I3O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader250` into `intoRequest637`.

	// Assume that `sourceCQL` has the underlying type of `fromReader250`:
	fromReader250 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader250` to result `intoRequest637`
	// (`intoRequest637` is now tainted).
	intoRequest637, _ := http.NewRequestWithContext(nil, "", "", fromReader250)

	// Return the tainted `intoRequest637`:
	return intoRequest637
}

func TaintStepTest_NetHttpReadRequest_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader831` into `intoRequest382`.

	// Assume that `sourceCQL` has the underlying type of `fromReader831`:
	fromReader831 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader831` to result `intoRequest382`
	// (`intoRequest382` is now tainted).
	intoRequest382, _ := http.ReadRequest(fromReader831)

	// Return the tainted `intoRequest382`:
	return intoRequest382
}

func TaintStepTest_NetHttpReadResponse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader804` into `intoResponse509`.

	// Assume that `sourceCQL` has the underlying type of `fromReader804`:
	fromReader804 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader804` to result `intoResponse509`
	// (`intoResponse509` is now tainted).
	intoResponse509, _ := http.ReadResponse(fromReader804, nil)

	// Return the tainted `intoResponse509`:
	return intoResponse509
}

func TaintStepTest_NetHttpSetCookie_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromCookie637` into `intoResponseWriter912`.

	// Assume that `sourceCQL` has the underlying type of `fromCookie637`:
	fromCookie637 := sourceCQL.(*http.Cookie)

	// Declare `intoResponseWriter912` variable:
	var intoResponseWriter912 http.ResponseWriter

	// Call the function that transfers the taint
	// from the parameter `fromCookie637` to parameter `intoResponseWriter912`;
	// `intoResponseWriter912` is now tainted.
	http.SetCookie(intoResponseWriter912, fromCookie637)

	// Return the tainted `intoResponseWriter912`:
	return intoResponseWriter912
}

func TaintStepTest_NetHttpHeaderAdd_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString348` into `intoHeader525`.

	// Assume that `sourceCQL` has the underlying type of `fromString348`:
	fromString348 := sourceCQL.(string)

	// Declare `intoHeader525` variable:
	var intoHeader525 http.Header

	// Call the method that transfers the taint
	// from the parameter `fromString348` to the receiver `intoHeader525`
	// (`intoHeader525` is now tainted).
	intoHeader525.Add(fromString348, "")

	// Return the tainted `intoHeader525`:
	return intoHeader525
}

func TaintStepTest_NetHttpHeaderAdd_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString410` into `intoHeader660`.

	// Assume that `sourceCQL` has the underlying type of `fromString410`:
	fromString410 := sourceCQL.(string)

	// Declare `intoHeader660` variable:
	var intoHeader660 http.Header

	// Call the method that transfers the taint
	// from the parameter `fromString410` to the receiver `intoHeader660`
	// (`intoHeader660` is now tainted).
	intoHeader660.Add("", fromString410)

	// Return the tainted `intoHeader660`:
	return intoHeader660
}

func TaintStepTest_NetHttpHeaderClone_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader146` into `intoHeader128`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader146`:
	fromHeader146 := sourceCQL.(http.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader146` to the result `intoHeader128`
	// (`intoHeader128` is now tainted).
	intoHeader128 := fromHeader146.Clone()

	// Return the tainted `intoHeader128`:
	return intoHeader128
}

func TaintStepTest_NetHttpHeaderGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader944` into `intoString836`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader944`:
	fromHeader944 := sourceCQL.(http.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader944` to the result `intoString836`
	// (`intoString836` is now tainted).
	intoString836 := fromHeader944.Get("")

	// Return the tainted `intoString836`:
	return intoString836
}

func TaintStepTest_NetHttpHeaderSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString430` into `intoHeader703`.

	// Assume that `sourceCQL` has the underlying type of `fromString430`:
	fromString430 := sourceCQL.(string)

	// Declare `intoHeader703` variable:
	var intoHeader703 http.Header

	// Call the method that transfers the taint
	// from the parameter `fromString430` to the receiver `intoHeader703`
	// (`intoHeader703` is now tainted).
	intoHeader703.Set(fromString430, "")

	// Return the tainted `intoHeader703`:
	return intoHeader703
}

func TaintStepTest_NetHttpHeaderSet_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString689` into `intoHeader589`.

	// Assume that `sourceCQL` has the underlying type of `fromString689`:
	fromString689 := sourceCQL.(string)

	// Declare `intoHeader589` variable:
	var intoHeader589 http.Header

	// Call the method that transfers the taint
	// from the parameter `fromString689` to the receiver `intoHeader589`
	// (`intoHeader589` is now tainted).
	intoHeader589.Set("", fromString689)

	// Return the tainted `intoHeader589`:
	return intoHeader589
}

func TaintStepTest_NetHttpHeaderValues_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader455` into `intoString331`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader455`:
	fromHeader455 := sourceCQL.(http.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader455` to the result `intoString331`
	// (`intoString331` is now tainted).
	intoString331 := fromHeader455.Values("")

	// Return the tainted `intoString331`:
	return intoString331
}

func TaintStepTest_NetHttpHeaderWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader246` into `intoWriter327`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader246`:
	fromHeader246 := sourceCQL.(http.Header)

	// Declare `intoWriter327` variable:
	var intoWriter327 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromHeader246` to the argument `intoWriter327`
	// (`intoWriter327` is now tainted).
	fromHeader246.Write(intoWriter327)

	// Return the tainted `intoWriter327`:
	return intoWriter327
}

func TaintStepTest_NetHttpHeaderWriteSubset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader151` into `intoWriter422`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader151`:
	fromHeader151 := sourceCQL.(http.Header)

	// Declare `intoWriter422` variable:
	var intoWriter422 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromHeader151` to the argument `intoWriter422`
	// (`intoWriter422` is now tainted).
	fromHeader151.WriteSubset(intoWriter422, nil)

	// Return the tainted `intoWriter422`:
	return intoWriter422
}

func TaintStepTest_NetHttpRequestAddCookie_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromCookie491` into `intoRequest593`.

	// Assume that `sourceCQL` has the underlying type of `fromCookie491`:
	fromCookie491 := sourceCQL.(*http.Cookie)

	// Declare `intoRequest593` variable:
	var intoRequest593 http.Request

	// Call the method that transfers the taint
	// from the parameter `fromCookie491` to the receiver `intoRequest593`
	// (`intoRequest593` is now tainted).
	intoRequest593.AddCookie(fromCookie491)

	// Return the tainted `intoRequest593`:
	return intoRequest593
}

func TaintStepTest_NetHttpRequestClone_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRequest706` into `intoRequest834`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest706`:
	fromRequest706 := sourceCQL.(http.Request)

	// Call the method that transfers the taint
	// from the receiver `fromRequest706` to the result `intoRequest834`
	// (`intoRequest834` is now tainted).
	intoRequest834 := fromRequest706.Clone(nil)

	// Return the tainted `intoRequest834`:
	return intoRequest834
}

func TaintStepTest_NetHttpRequestWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRequest995` into `intoWriter735`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest995`:
	fromRequest995 := sourceCQL.(http.Request)

	// Declare `intoWriter735` variable:
	var intoWriter735 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromRequest995` to the argument `intoWriter735`
	// (`intoWriter735` is now tainted).
	fromRequest995.Write(intoWriter735)

	// Return the tainted `intoWriter735`:
	return intoWriter735
}

func TaintStepTest_NetHttpRequestWriteProxy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRequest880` into `intoWriter122`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest880`:
	fromRequest880 := sourceCQL.(http.Request)

	// Declare `intoWriter122` variable:
	var intoWriter122 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromRequest880` to the argument `intoWriter122`
	// (`intoWriter122` is now tainted).
	fromRequest880.WriteProxy(intoWriter122)

	// Return the tainted `intoWriter122`:
	return intoWriter122
}

func TaintStepTest_NetHttpResponseWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromResponse233` into `intoWriter277`.

	// Assume that `sourceCQL` has the underlying type of `fromResponse233`:
	fromResponse233 := sourceCQL.(http.Response)

	// Declare `intoWriter277` variable:
	var intoWriter277 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromResponse233` to the argument `intoWriter277`
	// (`intoWriter277` is now tainted).
	fromResponse233.Write(intoWriter277)

	// Return the tainted `intoWriter277`:
	return intoWriter277
}

func TaintStepTest_NetHttpHijackerHijack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHijacker536` into `intoConn906`.

	// Assume that `sourceCQL` has the underlying type of `fromHijacker536`:
	fromHijacker536 := sourceCQL.(http.Hijacker)

	// Call the method that transfers the taint
	// from the receiver `fromHijacker536` to the result `intoConn906`
	// (`intoConn906` is now tainted).
	intoConn906, _, _ := fromHijacker536.Hijack()

	// Return the tainted `intoConn906`:
	return intoConn906
}

func TaintStepTest_NetHttpHijackerHijack_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromHijacker155` into `intoReadWriter304`.

	// Assume that `sourceCQL` has the underlying type of `fromHijacker155`:
	fromHijacker155 := sourceCQL.(http.Hijacker)

	// Call the method that transfers the taint
	// from the receiver `fromHijacker155` to the result `intoReadWriter304`
	// (`intoReadWriter304` is now tainted).
	_, intoReadWriter304, _ := fromHijacker155.Hijack()

	// Return the tainted `intoReadWriter304`:
	return intoReadWriter304
}

func TaintStepTest_NetHttpResponseWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte797` into `intoResponseWriter923`.

	// Assume that `sourceCQL` has the underlying type of `fromByte797`:
	fromByte797 := sourceCQL.([]byte)

	// Declare `intoResponseWriter923` variable:
	var intoResponseWriter923 http.ResponseWriter

	// Call the method that transfers the taint
	// from the parameter `fromByte797` to the receiver `intoResponseWriter923`
	// (`intoResponseWriter923` is now tainted).
	intoResponseWriter923.Write(fromByte797)

	// Return the tainted `intoResponseWriter923`:
	return intoResponseWriter923
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
