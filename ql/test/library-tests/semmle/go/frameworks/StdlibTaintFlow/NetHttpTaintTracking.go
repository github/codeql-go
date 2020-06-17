package main

import (
	"bufio"
	"context"
	"io"
	"net/http"
)

func TaintStepTest_NetHttpCanonicalHeaderKey_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString887` into `intoString520`.

	// Assume that `sourceCQL` has the underlying type of `fromString887`:
	fromString887 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString887` to result `intoString520`
	// (`intoString520` is now tainted).
	intoString520 := http.CanonicalHeaderKey(fromString887)

	// Return the tainted `intoString520`:
	return intoString520
}

func TaintStepTest_NetHttpDetectContentType_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte423` into `intoString505`.

	// Assume that `sourceCQL` has the underlying type of `fromByte423`:
	fromByte423 := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte423` to result `intoString505`
	// (`intoString505` is now tainted).
	intoString505 := http.DetectContentType(fromByte423)

	// Return the tainted `intoString505`:
	return intoString505
}

func TaintStepTest_NetHttpError_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString297` into `intoResponseWriter187`.

	// Assume that `sourceCQL` has the underlying type of `fromString297`:
	fromString297 := sourceCQL.(string)

	// Declare `intoResponseWriter187` variable:
	var intoResponseWriter187 http.ResponseWriter

	// Call the function that transfers the taint
	// from the parameter `fromString297` to parameter `intoResponseWriter187`;
	// `intoResponseWriter187` is now tainted.
	http.Error(intoResponseWriter187, fromString297, 0)

	// Return the tainted `intoResponseWriter187`:
	return intoResponseWriter187
}

func TaintStepTest_NetHttpMaxBytesReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReadCloser140` into `intoReadCloser694`.

	// Assume that `sourceCQL` has the underlying type of `fromReadCloser140`:
	fromReadCloser140 := sourceCQL.(io.ReadCloser)

	// Call the function that transfers the taint
	// from the parameter `fromReadCloser140` to result `intoReadCloser694`
	// (`intoReadCloser694` is now tainted).
	intoReadCloser694 := http.MaxBytesReader(nil, fromReadCloser140, 0)

	// Return the tainted `intoReadCloser694`:
	return intoReadCloser694
}

func TaintStepTest_NetHttpNewRequest_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString296` into `intoRequest389`.

	// Assume that `sourceCQL` has the underlying type of `fromString296`:
	fromString296 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString296` to result `intoRequest389`
	// (`intoRequest389` is now tainted).
	intoRequest389, _ := http.NewRequest(fromString296, "", nil)

	// Return the tainted `intoRequest389`:
	return intoRequest389
}

func TaintStepTest_NetHttpNewRequest_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString813` into `intoRequest128`.

	// Assume that `sourceCQL` has the underlying type of `fromString813`:
	fromString813 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString813` to result `intoRequest128`
	// (`intoRequest128` is now tainted).
	intoRequest128, _ := http.NewRequest("", fromString813, nil)

	// Return the tainted `intoRequest128`:
	return intoRequest128
}

func TaintStepTest_NetHttpNewRequest_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader818` into `intoRequest745`.

	// Assume that `sourceCQL` has the underlying type of `fromReader818`:
	fromReader818 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader818` to result `intoRequest745`
	// (`intoRequest745` is now tainted).
	intoRequest745, _ := http.NewRequest("", "", fromReader818)

	// Return the tainted `intoRequest745`:
	return intoRequest745
}

func TaintStepTest_NetHttpNewRequestWithContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromContext461` into `intoRequest238`.

	// Assume that `sourceCQL` has the underlying type of `fromContext461`:
	fromContext461 := sourceCQL.(context.Context)

	// Call the function that transfers the taint
	// from the parameter `fromContext461` to result `intoRequest238`
	// (`intoRequest238` is now tainted).
	intoRequest238, _ := http.NewRequestWithContext(fromContext461, "", "", nil)

	// Return the tainted `intoRequest238`:
	return intoRequest238
}

func TaintStepTest_NetHttpNewRequestWithContext_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString155` into `intoRequest927`.

	// Assume that `sourceCQL` has the underlying type of `fromString155`:
	fromString155 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString155` to result `intoRequest927`
	// (`intoRequest927` is now tainted).
	intoRequest927, _ := http.NewRequestWithContext(nil, fromString155, "", nil)

	// Return the tainted `intoRequest927`:
	return intoRequest927
}

func TaintStepTest_NetHttpNewRequestWithContext_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString298` into `intoRequest523`.

	// Assume that `sourceCQL` has the underlying type of `fromString298`:
	fromString298 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString298` to result `intoRequest523`
	// (`intoRequest523` is now tainted).
	intoRequest523, _ := http.NewRequestWithContext(nil, "", fromString298, nil)

	// Return the tainted `intoRequest523`:
	return intoRequest523
}

func TaintStepTest_NetHttpNewRequestWithContext_B0I3O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader780` into `intoRequest648`.

	// Assume that `sourceCQL` has the underlying type of `fromReader780`:
	fromReader780 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader780` to result `intoRequest648`
	// (`intoRequest648` is now tainted).
	intoRequest648, _ := http.NewRequestWithContext(nil, "", "", fromReader780)

	// Return the tainted `intoRequest648`:
	return intoRequest648
}

func TaintStepTest_NetHttpReadRequest_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader284` into `intoRequest431`.

	// Assume that `sourceCQL` has the underlying type of `fromReader284`:
	fromReader284 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader284` to result `intoRequest431`
	// (`intoRequest431` is now tainted).
	intoRequest431, _ := http.ReadRequest(fromReader284)

	// Return the tainted `intoRequest431`:
	return intoRequest431
}

func TaintStepTest_NetHttpReadResponse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader900` into `intoResponse266`.

	// Assume that `sourceCQL` has the underlying type of `fromReader900`:
	fromReader900 := sourceCQL.(*bufio.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader900` to result `intoResponse266`
	// (`intoResponse266` is now tainted).
	intoResponse266, _ := http.ReadResponse(fromReader900, nil)

	// Return the tainted `intoResponse266`:
	return intoResponse266
}

func TaintStepTest_NetHttpSetCookie_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromCookie701` into `intoResponseWriter133`.

	// Assume that `sourceCQL` has the underlying type of `fromCookie701`:
	fromCookie701 := sourceCQL.(*http.Cookie)

	// Declare `intoResponseWriter133` variable:
	var intoResponseWriter133 http.ResponseWriter

	// Call the function that transfers the taint
	// from the parameter `fromCookie701` to parameter `intoResponseWriter133`;
	// `intoResponseWriter133` is now tainted.
	http.SetCookie(intoResponseWriter133, fromCookie701)

	// Return the tainted `intoResponseWriter133`:
	return intoResponseWriter133
}

func TaintStepTest_NetHttpHeaderAdd_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString691` into `intoHeader205`.

	// Assume that `sourceCQL` has the underlying type of `fromString691`:
	fromString691 := sourceCQL.(string)

	// Declare `intoHeader205` variable:
	var intoHeader205 http.Header

	// Call the method that transfers the taint
	// from the parameter `fromString691` to the receiver `intoHeader205`
	// (`intoHeader205` is now tainted).
	intoHeader205.Add(fromString691, "")

	// Return the tainted `intoHeader205`:
	return intoHeader205
}

func TaintStepTest_NetHttpHeaderAdd_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString985` into `intoHeader553`.

	// Assume that `sourceCQL` has the underlying type of `fromString985`:
	fromString985 := sourceCQL.(string)

	// Declare `intoHeader553` variable:
	var intoHeader553 http.Header

	// Call the method that transfers the taint
	// from the parameter `fromString985` to the receiver `intoHeader553`
	// (`intoHeader553` is now tainted).
	intoHeader553.Add("", fromString985)

	// Return the tainted `intoHeader553`:
	return intoHeader553
}

func TaintStepTest_NetHttpHeaderClone_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader469` into `intoHeader699`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader469`:
	fromHeader469 := sourceCQL.(http.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader469` to the result `intoHeader699`
	// (`intoHeader699` is now tainted).
	intoHeader699 := fromHeader469.Clone()

	// Return the tainted `intoHeader699`:
	return intoHeader699
}

func TaintStepTest_NetHttpHeaderGet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader720` into `intoString253`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader720`:
	fromHeader720 := sourceCQL.(http.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader720` to the result `intoString253`
	// (`intoString253` is now tainted).
	intoString253 := fromHeader720.Get("")

	// Return the tainted `intoString253`:
	return intoString253
}

func TaintStepTest_NetHttpHeaderSet_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString278` into `intoHeader313`.

	// Assume that `sourceCQL` has the underlying type of `fromString278`:
	fromString278 := sourceCQL.(string)

	// Declare `intoHeader313` variable:
	var intoHeader313 http.Header

	// Call the method that transfers the taint
	// from the parameter `fromString278` to the receiver `intoHeader313`
	// (`intoHeader313` is now tainted).
	intoHeader313.Set(fromString278, "")

	// Return the tainted `intoHeader313`:
	return intoHeader313
}

func TaintStepTest_NetHttpHeaderSet_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString450` into `intoHeader831`.

	// Assume that `sourceCQL` has the underlying type of `fromString450`:
	fromString450 := sourceCQL.(string)

	// Declare `intoHeader831` variable:
	var intoHeader831 http.Header

	// Call the method that transfers the taint
	// from the parameter `fromString450` to the receiver `intoHeader831`
	// (`intoHeader831` is now tainted).
	intoHeader831.Set("", fromString450)

	// Return the tainted `intoHeader831`:
	return intoHeader831
}

func TaintStepTest_NetHttpHeaderValues_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader568` into `intoString472`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader568`:
	fromHeader568 := sourceCQL.(http.Header)

	// Call the method that transfers the taint
	// from the receiver `fromHeader568` to the result `intoString472`
	// (`intoString472` is now tainted).
	intoString472 := fromHeader568.Values("")

	// Return the tainted `intoString472`:
	return intoString472
}

func TaintStepTest_NetHttpHeaderWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader468` into `intoWriter441`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader468`:
	fromHeader468 := sourceCQL.(http.Header)

	// Declare `intoWriter441` variable:
	var intoWriter441 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromHeader468` to the argument `intoWriter441`
	// (`intoWriter441` is now tainted).
	fromHeader468.Write(intoWriter441)

	// Return the tainted `intoWriter441`:
	return intoWriter441
}

func TaintStepTest_NetHttpHeaderWriteSubset_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHeader420` into `intoWriter879`.

	// Assume that `sourceCQL` has the underlying type of `fromHeader420`:
	fromHeader420 := sourceCQL.(http.Header)

	// Declare `intoWriter879` variable:
	var intoWriter879 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromHeader420` to the argument `intoWriter879`
	// (`intoWriter879` is now tainted).
	fromHeader420.WriteSubset(intoWriter879, nil)

	// Return the tainted `intoWriter879`:
	return intoWriter879
}

func TaintStepTest_NetHttpRequestAddCookie_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromCookie729` into `intoRequest736`.

	// Assume that `sourceCQL` has the underlying type of `fromCookie729`:
	fromCookie729 := sourceCQL.(*http.Cookie)

	// Declare `intoRequest736` variable:
	var intoRequest736 http.Request

	// Call the method that transfers the taint
	// from the parameter `fromCookie729` to the receiver `intoRequest736`
	// (`intoRequest736` is now tainted).
	intoRequest736.AddCookie(fromCookie729)

	// Return the tainted `intoRequest736`:
	return intoRequest736
}

func TaintStepTest_NetHttpRequestClone_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRequest774` into `intoRequest864`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest774`:
	fromRequest774 := sourceCQL.(http.Request)

	// Call the method that transfers the taint
	// from the receiver `fromRequest774` to the result `intoRequest864`
	// (`intoRequest864` is now tainted).
	intoRequest864 := fromRequest774.Clone(nil)

	// Return the tainted `intoRequest864`:
	return intoRequest864
}

func TaintStepTest_NetHttpRequestWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRequest746` into `intoWriter920`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest746`:
	fromRequest746 := sourceCQL.(http.Request)

	// Declare `intoWriter920` variable:
	var intoWriter920 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromRequest746` to the argument `intoWriter920`
	// (`intoWriter920` is now tainted).
	fromRequest746.Write(intoWriter920)

	// Return the tainted `intoWriter920`:
	return intoWriter920
}

func TaintStepTest_NetHttpRequestWriteProxy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRequest530` into `intoWriter315`.

	// Assume that `sourceCQL` has the underlying type of `fromRequest530`:
	fromRequest530 := sourceCQL.(http.Request)

	// Declare `intoWriter315` variable:
	var intoWriter315 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromRequest530` to the argument `intoWriter315`
	// (`intoWriter315` is now tainted).
	fromRequest530.WriteProxy(intoWriter315)

	// Return the tainted `intoWriter315`:
	return intoWriter315
}

func TaintStepTest_NetHttpResponseWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromResponse313` into `intoWriter444`.

	// Assume that `sourceCQL` has the underlying type of `fromResponse313`:
	fromResponse313 := sourceCQL.(http.Response)

	// Declare `intoWriter444` variable:
	var intoWriter444 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromResponse313` to the argument `intoWriter444`
	// (`intoWriter444` is now tainted).
	fromResponse313.Write(intoWriter444)

	// Return the tainted `intoWriter444`:
	return intoWriter444
}

func TaintStepTest_NetHttpHijackerHijack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHijacker127` into `intoConn593`.

	// Assume that `sourceCQL` has the underlying type of `fromHijacker127`:
	fromHijacker127 := sourceCQL.(http.Hijacker)

	// Call the method that transfers the taint
	// from the receiver `fromHijacker127` to the result `intoConn593`
	// (`intoConn593` is now tainted).
	intoConn593, _, _ := fromHijacker127.Hijack()

	// Return the tainted `intoConn593`:
	return intoConn593
}

func TaintStepTest_NetHttpHijackerHijack_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromHijacker372` into `intoReadWriter910`.

	// Assume that `sourceCQL` has the underlying type of `fromHijacker372`:
	fromHijacker372 := sourceCQL.(http.Hijacker)

	// Call the method that transfers the taint
	// from the receiver `fromHijacker372` to the result `intoReadWriter910`
	// (`intoReadWriter910` is now tainted).
	_, intoReadWriter910, _ := fromHijacker372.Hijack()

	// Return the tainted `intoReadWriter910`:
	return intoReadWriter910
}

func TaintStepTest_NetHttpResponseWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte550` into `intoResponseWriter279`.

	// Assume that `sourceCQL` has the underlying type of `fromByte550`:
	fromByte550 := sourceCQL.([]byte)

	// Declare `intoResponseWriter279` variable:
	var intoResponseWriter279 http.ResponseWriter

	// Call the method that transfers the taint
	// from the parameter `fromByte550` to the receiver `intoResponseWriter279`
	// (`intoResponseWriter279` is now tainted).
	intoResponseWriter279.Write(fromByte550)

	// Return the tainted `intoResponseWriter279`:
	return intoResponseWriter279
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
