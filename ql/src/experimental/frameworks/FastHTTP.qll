/**
 * Provides classes for working with concepts from the [`github.com/valyala/fasthttp@v1.23.0`](https://pkg.go.dev/github.com/valyala/fasthttp@v1.23.0) package.
 */

import go

/**
 * Provides classes for working with concepts from the [`github.com/valyala/fasthttp@v1.23.0`](https://pkg.go.dev/github.com/valyala/fasthttp@v1.23.0) package.
 */
private module FastHTTP {
  string packagePath() { result = package("github.com/valyala/fasthttp", "") }

  /**
   * Models taint-tracking through functions.
   */
  private class TaintTrackingFunctionModels extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput out;

    TaintTrackingFunctionModels() {
      // Taint-tracking models for package: github.com/valyala/fasthttp@v1.23.0
      (
        // signature: func AppendBrotliBytes(dst []byte, src []byte) []byte
        this.hasQualifiedName(packagePath(), "AppendBrotliBytes") and
        (
          inp.isParameter(1) and
          (
            out.isParameter(0) or
            out.isResult()
          )
          or
          inp.isParameter(0) and
          out.isResult()
        )
        or
        // signature: func AppendBrotliBytesLevel(dst []byte, src []byte, level int) []byte
        this.hasQualifiedName(packagePath(), "AppendBrotliBytesLevel") and
        (
          inp.isParameter(1) and
          (
            out.isParameter(0) or
            out.isResult()
          )
          or
          inp.isParameter(0) and
          out.isResult()
        )
        or
        // signature: func AppendDeflateBytes(dst []byte, src []byte) []byte
        this.hasQualifiedName(packagePath(), "AppendDeflateBytes") and
        (
          inp.isParameter(1) and
          (
            out.isParameter(0) or
            out.isResult()
          )
          or
          inp.isParameter(0) and
          out.isResult()
        )
        or
        // signature: func AppendDeflateBytesLevel(dst []byte, src []byte, level int) []byte
        this.hasQualifiedName(packagePath(), "AppendDeflateBytesLevel") and
        (
          inp.isParameter(1) and
          (
            out.isParameter(0) or
            out.isResult()
          )
          or
          inp.isParameter(0) and
          out.isResult()
        )
        or
        // signature: func AppendGunzipBytes(dst []byte, src []byte) ([]byte, error)
        this.hasQualifiedName(packagePath(), "AppendGunzipBytes") and
        (
          inp.isParameter(1) and
          (
            out.isParameter(0) or
            out.isResult(0)
          )
          or
          inp.isParameter(0) and
          out.isResult(0)
        )
        or
        // signature: func AppendGzipBytes(dst []byte, src []byte) []byte
        this.hasQualifiedName(packagePath(), "AppendGzipBytes") and
        (
          inp.isParameter(1) and
          (
            out.isParameter(0) or
            out.isResult()
          )
          or
          inp.isParameter(0) and
          out.isResult()
        )
        or
        // signature: func AppendGzipBytesLevel(dst []byte, src []byte, level int) []byte
        this.hasQualifiedName(packagePath(), "AppendGzipBytesLevel") and
        (
          inp.isParameter(1) and
          (
            out.isParameter(0) or
            out.isResult()
          )
          or
          inp.isParameter(0) and
          out.isResult()
        )
        or
        // signature: func AppendHTMLEscape(dst []byte, s string) []byte
        this.hasQualifiedName(packagePath(), "AppendHTMLEscape") and
        (
          inp.isParameter(1) and
          (
            out.isParameter(0) or
            out.isResult()
          )
          or
          inp.isParameter(0) and
          out.isResult()
        )
        or
        // signature: func AppendHTMLEscapeBytes(dst []byte, s []byte) []byte
        this.hasQualifiedName(packagePath(), "AppendHTMLEscapeBytes") and
        (
          inp.isParameter(1) and
          (
            out.isParameter(0) or
            out.isResult()
          )
          or
          inp.isParameter(0) and
          out.isResult()
        )
        or
        // signature: func AppendInflateBytes(dst []byte, src []byte) ([]byte, error)
        this.hasQualifiedName(packagePath(), "AppendInflateBytes") and
        (
          inp.isParameter(1) and
          (
            out.isParameter(0) or
            out.isResult(0)
          )
          or
          inp.isParameter(0) and
          out.isResult(0)
        )
        or
        // signature: func AppendNormalizedHeaderKey(dst []byte, key string) []byte
        this.hasQualifiedName(packagePath(), "AppendNormalizedHeaderKey") and
        (
          inp.isParameter(1) and
          (
            out.isParameter(0) or
            out.isResult()
          )
          or
          inp.isParameter(0) and
          out.isResult()
        )
        or
        // signature: func AppendNormalizedHeaderKeyBytes(dst []byte, key []byte) []byte
        this.hasQualifiedName(packagePath(), "AppendNormalizedHeaderKeyBytes") and
        (
          inp.isParameter(1) and
          (
            out.isParameter(0) or
            out.isResult()
          )
          or
          inp.isParameter(0) and
          out.isResult()
        )
        or
        // signature: func AppendQuotedArg(dst []byte, src []byte) []byte
        this.hasQualifiedName(packagePath(), "AppendQuotedArg") and
        (
          inp.isParameter(1) and
          (
            out.isParameter(0) or
            out.isResult()
          )
          or
          inp.isParameter(0) and
          out.isResult()
        )
        or
        // signature: func AppendUnbrotliBytes(dst []byte, src []byte) ([]byte, error)
        this.hasQualifiedName(packagePath(), "AppendUnbrotliBytes") and
        (
          inp.isParameter(1) and
          (
            out.isParameter(0) or
            out.isResult(0)
          )
          or
          inp.isParameter(0) and
          out.isResult(0)
        )
        or
        // signature: func AppendUnquotedArg(dst []byte, src []byte) []byte
        this.hasQualifiedName(packagePath(), "AppendUnquotedArg") and
        (
          inp.isParameter(1) and
          (
            out.isParameter(0) or
            out.isResult()
          )
          or
          inp.isParameter(0) and
          out.isResult()
        )
        or
        // signature: func WriteBrotli(w io.Writer, p []byte) (int, error)
        this.hasQualifiedName(packagePath(), "WriteBrotli") and
        inp.isParameter(1) and
        out.isParameter(0)
        or
        // signature: func WriteBrotliLevel(w io.Writer, p []byte, level int) (int, error)
        this.hasQualifiedName(packagePath(), "WriteBrotliLevel") and
        inp.isParameter(1) and
        out.isParameter(0)
        or
        // signature: func WriteDeflate(w io.Writer, p []byte) (int, error)
        this.hasQualifiedName(packagePath(), "WriteDeflate") and
        inp.isParameter(1) and
        out.isParameter(0)
        or
        // signature: func WriteDeflateLevel(w io.Writer, p []byte, level int) (int, error)
        this.hasQualifiedName(packagePath(), "WriteDeflateLevel") and
        inp.isParameter(1) and
        out.isParameter(0)
        or
        // signature: func WriteGunzip(w io.Writer, p []byte) (int, error)
        this.hasQualifiedName(packagePath(), "WriteGunzip") and
        inp.isParameter(1) and
        out.isParameter(0)
        or
        // signature: func WriteGzip(w io.Writer, p []byte) (int, error)
        this.hasQualifiedName(packagePath(), "WriteGzip") and
        inp.isParameter(1) and
        out.isParameter(0)
        or
        // signature: func WriteGzipLevel(w io.Writer, p []byte, level int) (int, error)
        this.hasQualifiedName(packagePath(), "WriteGzipLevel") and
        inp.isParameter(1) and
        out.isParameter(0)
        or
        // signature: func WriteInflate(w io.Writer, p []byte) (int, error)
        this.hasQualifiedName(packagePath(), "WriteInflate") and
        inp.isParameter(1) and
        out.isParameter(0)
        or
        // signature: func WriteMultipartForm(w io.Writer, f *mime/multipart.Form, boundary string) error
        this.hasQualifiedName(packagePath(), "WriteMultipartForm") and
        inp.isParameter([1, 2]) and
        out.isParameter(0)
        or
        // signature: func WriteUnbrotli(w io.Writer, p []byte) (int, error)
        this.hasQualifiedName(packagePath(), "WriteUnbrotli") and
        inp.isParameter(1) and
        out.isParameter(0)
      )
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = out
    }
  }

  /**
   * Models taint-tracking through method calls.
   */
  private class TaintTrackingMethodModels extends TaintTracking::FunctionModel, Method {
    FunctionInput inp;
    FunctionOutput out;

    TaintTrackingMethodModels() {
      // Taint-tracking models for package: github.com/valyala/fasthttp@v1.23.0
      (
        // Receiver type: Args
        (
          // signature: func (*Args).Add(key string, value string)
          this.hasQualifiedName(packagePath(), "Args", "Add") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*Args).AddBytesK(key []byte, value string)
          this.hasQualifiedName(packagePath(), "Args", "AddBytesK") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*Args).AddBytesKNoValue(key []byte)
          this.hasQualifiedName(packagePath(), "Args", "AddBytesKNoValue") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Args).AddBytesKV(key []byte, value []byte)
          this.hasQualifiedName(packagePath(), "Args", "AddBytesKV") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*Args).AddBytesV(key string, value []byte)
          this.hasQualifiedName(packagePath(), "Args", "AddBytesV") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*Args).AddNoValue(key string)
          this.hasQualifiedName(packagePath(), "Args", "AddNoValue") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Args).AppendBytes(dst []byte) []byte
          this.hasQualifiedName(packagePath(), "Args", "AppendBytes") and
          (
            inp.isReceiver() and
            (
              out.isParameter(0) or
              out.isResult()
            )
            or
            inp.isParameter(0) and
            out.isResult()
          )
          or
          // signature: func (*Args).CopyTo(dst *Args)
          this.hasQualifiedName(packagePath(), "Args", "CopyTo") and
          inp.isReceiver() and
          out.isParameter(0)
          or
          // signature: func (*Args).Parse(s string)
          this.hasQualifiedName(packagePath(), "Args", "Parse") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Args).ParseBytes(b []byte)
          this.hasQualifiedName(packagePath(), "Args", "ParseBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Args).Peek(key string) []byte
          this.hasQualifiedName(packagePath(), "Args", "Peek") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Args).PeekBytes(key []byte) []byte
          this.hasQualifiedName(packagePath(), "Args", "PeekBytes") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Args).PeekMulti(key string) [][]byte
          this.hasQualifiedName(packagePath(), "Args", "PeekMulti") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Args).PeekMultiBytes(key []byte) [][]byte
          this.hasQualifiedName(packagePath(), "Args", "PeekMultiBytes") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Args).QueryString() []byte
          this.hasQualifiedName(packagePath(), "Args", "QueryString") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Args).Set(key string, value string)
          this.hasQualifiedName(packagePath(), "Args", "Set") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*Args).SetBytesK(key []byte, value string)
          this.hasQualifiedName(packagePath(), "Args", "SetBytesK") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*Args).SetBytesKNoValue(key []byte)
          this.hasQualifiedName(packagePath(), "Args", "SetBytesKNoValue") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Args).SetBytesKV(key []byte, value []byte)
          this.hasQualifiedName(packagePath(), "Args", "SetBytesKV") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*Args).SetBytesV(key string, value []byte)
          this.hasQualifiedName(packagePath(), "Args", "SetBytesV") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*Args).SetNoValue(key string)
          this.hasQualifiedName(packagePath(), "Args", "SetNoValue") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Args).String() string
          this.hasQualifiedName(packagePath(), "Args", "String") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Args).WriteTo(w io.Writer) (int64, error)
          this.hasQualifiedName(packagePath(), "Args", "WriteTo") and
          inp.isReceiver() and
          out.isParameter(0)
        )
        or
        // Receiver type: Cookie
        (
          // signature: func (*Cookie).AppendBytes(dst []byte) []byte
          this.hasQualifiedName(packagePath(), "Cookie", "AppendBytes") and
          (
            inp.isReceiver() and
            (
              out.isParameter(0) or
              out.isResult()
            )
            or
            inp.isParameter(0) and
            out.isResult()
          )
          or
          // signature: func (*Cookie).Cookie() []byte
          this.hasQualifiedName(packagePath(), "Cookie", "Cookie") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Cookie).CopyTo(src *Cookie)
          this.hasQualifiedName(packagePath(), "Cookie", "CopyTo") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Cookie).Domain() []byte
          this.hasQualifiedName(packagePath(), "Cookie", "Domain") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Cookie).Key() []byte
          this.hasQualifiedName(packagePath(), "Cookie", "Key") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Cookie).Parse(src string) error
          this.hasQualifiedName(packagePath(), "Cookie", "Parse") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Cookie).ParseBytes(src []byte) error
          this.hasQualifiedName(packagePath(), "Cookie", "ParseBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Cookie).Path() []byte
          this.hasQualifiedName(packagePath(), "Cookie", "Path") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Cookie).SetDomain(domain string)
          this.hasQualifiedName(packagePath(), "Cookie", "SetDomain") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Cookie).SetDomainBytes(domain []byte)
          this.hasQualifiedName(packagePath(), "Cookie", "SetDomainBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Cookie).SetKey(key string)
          this.hasQualifiedName(packagePath(), "Cookie", "SetKey") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Cookie).SetKeyBytes(key []byte)
          this.hasQualifiedName(packagePath(), "Cookie", "SetKeyBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Cookie).SetPath(path string)
          this.hasQualifiedName(packagePath(), "Cookie", "SetPath") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Cookie).SetPathBytes(path []byte)
          this.hasQualifiedName(packagePath(), "Cookie", "SetPathBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Cookie).SetValue(value string)
          this.hasQualifiedName(packagePath(), "Cookie", "SetValue") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Cookie).SetValueBytes(value []byte)
          this.hasQualifiedName(packagePath(), "Cookie", "SetValueBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Cookie).String() string
          this.hasQualifiedName(packagePath(), "Cookie", "String") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Cookie).Value() []byte
          this.hasQualifiedName(packagePath(), "Cookie", "Value") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Cookie).WriteTo(w io.Writer) (int64, error)
          this.hasQualifiedName(packagePath(), "Cookie", "WriteTo") and
          inp.isReceiver() and
          out.isParameter(0)
        )
        or
        // Receiver type: Request
        (
          // signature: func (*Request).AppendBody(p []byte)
          this.hasQualifiedName(packagePath(), "Request", "AppendBody") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Request).AppendBodyString(s string)
          this.hasQualifiedName(packagePath(), "Request", "AppendBodyString") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Request).Body() []byte
          this.hasQualifiedName(packagePath(), "Request", "Body") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Request).BodyGunzip() ([]byte, error)
          this.hasQualifiedName(packagePath(), "Request", "BodyGunzip") and
          inp.isReceiver() and
          out.isResult(0)
          or
          // signature: func (*Request).BodyInflate() ([]byte, error)
          this.hasQualifiedName(packagePath(), "Request", "BodyInflate") and
          inp.isReceiver() and
          out.isResult(0)
          or
          // signature: func (*Request).BodyUnbrotli() ([]byte, error)
          this.hasQualifiedName(packagePath(), "Request", "BodyUnbrotli") and
          inp.isReceiver() and
          out.isResult(0)
          or
          // signature: func (*Request).BodyWriteTo(w io.Writer) error
          this.hasQualifiedName(packagePath(), "Request", "BodyWriteTo") and
          inp.isReceiver() and
          out.isParameter(0)
          or
          // signature: func (*Request).BodyWriter() io.Writer
          this.hasQualifiedName(packagePath(), "Request", "BodyWriter") and
          inp.isResult() and
          out.isReceiver()
          or
          // signature: func (*Request).ContinueReadBody(r *bufio.Reader, maxBodySize int, preParseMultipartForm ...bool) error
          this.hasQualifiedName(packagePath(), "Request", "ContinueReadBody") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*Request).ContinueReadBodyStream(r *bufio.Reader, maxBodySize int, preParseMultipartForm ...bool) error
          this.hasQualifiedName(packagePath(), "Request", "ContinueReadBodyStream") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*Request).CopyTo(dst *Request)
          this.hasQualifiedName(packagePath(), "Request", "CopyTo") and
          inp.isReceiver() and
          out.isParameter(0)
          or
          // signature: func (*Request).Host() []byte
          this.hasQualifiedName(packagePath(), "Request", "Host") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Request).MultipartForm() (*mime/multipart.Form, error)
          this.hasQualifiedName(packagePath(), "Request", "MultipartForm") and
          inp.isReceiver() and
          out.isResult(0)
          or
          // signature: func (*Request).PostArgs() *Args
          this.hasQualifiedName(packagePath(), "Request", "PostArgs") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Request).Read(r *bufio.Reader) error
          this.hasQualifiedName(packagePath(), "Request", "Read") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Request).ReadLimitBody(r *bufio.Reader, maxBodySize int) error
          this.hasQualifiedName(packagePath(), "Request", "ReadLimitBody") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Request).RequestURI() []byte
          this.hasQualifiedName(packagePath(), "Request", "RequestURI") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Request).SetBody(body []byte)
          this.hasQualifiedName(packagePath(), "Request", "SetBody") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Request).SetBodyRaw(body []byte)
          this.hasQualifiedName(packagePath(), "Request", "SetBodyRaw") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Request).SetBodyStream(bodyStream io.Reader, bodySize int)
          this.hasQualifiedName(packagePath(), "Request", "SetBodyStream") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Request).SetBodyString(body string)
          this.hasQualifiedName(packagePath(), "Request", "SetBodyString") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Request).SetHost(host string)
          this.hasQualifiedName(packagePath(), "Request", "SetHost") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Request).SetHostBytes(host []byte)
          this.hasQualifiedName(packagePath(), "Request", "SetHostBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Request).SetRequestURI(requestURI string)
          this.hasQualifiedName(packagePath(), "Request", "SetRequestURI") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Request).SetRequestURIBytes(requestURI []byte)
          this.hasQualifiedName(packagePath(), "Request", "SetRequestURIBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Request).String() string
          this.hasQualifiedName(packagePath(), "Request", "String") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Request).SwapBody(body []byte) []byte
          this.hasQualifiedName(packagePath(), "Request", "SwapBody") and
          (
            inp.isParameter(0) and
            out.isReceiver()
            or
            inp.isReceiver() and
            out.isResult()
          )
          or
          // signature: func (*Request).URI() *URI
          this.hasQualifiedName(packagePath(), "Request", "URI") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Request).Write(w *bufio.Writer) error
          this.hasQualifiedName(packagePath(), "Request", "Write") and
          inp.isReceiver() and
          out.isParameter(0)
          or
          // signature: func (*Request).WriteTo(w io.Writer) (int64, error)
          this.hasQualifiedName(packagePath(), "Request", "WriteTo") and
          inp.isReceiver() and
          out.isParameter(0)
        )
        or
        // Receiver type: RequestCtx
        (
          // signature: func (*RequestCtx).Conn() net.Conn
          this.hasQualifiedName(packagePath(), "RequestCtx", "Conn") and
          (
            inp.isReceiver() and
            out.isResult()
            or
            inp.isResult() and
            out.isReceiver()
          )
          or
          // signature: func (*RequestCtx).FormFile(key string) (*mime/multipart.FileHeader, error)
          this.hasQualifiedName(packagePath(), "RequestCtx", "FormFile") and
          inp.isReceiver() and
          out.isResult(0)
          or
          // signature: func (*RequestCtx).FormValue(key string) []byte
          this.hasQualifiedName(packagePath(), "RequestCtx", "FormValue") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestCtx).Host() []byte
          this.hasQualifiedName(packagePath(), "RequestCtx", "Host") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestCtx).Method() []byte
          this.hasQualifiedName(packagePath(), "RequestCtx", "Method") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestCtx).MultipartForm() (*mime/multipart.Form, error)
          this.hasQualifiedName(packagePath(), "RequestCtx", "MultipartForm") and
          inp.isReceiver() and
          out.isResult(0)
          or
          // signature: func (*RequestCtx).Path() []byte
          this.hasQualifiedName(packagePath(), "RequestCtx", "Path") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestCtx).PostArgs() *Args
          this.hasQualifiedName(packagePath(), "RequestCtx", "PostArgs") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestCtx).PostBody() []byte
          this.hasQualifiedName(packagePath(), "RequestCtx", "PostBody") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestCtx).QueryArgs() *Args
          this.hasQualifiedName(packagePath(), "RequestCtx", "QueryArgs") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestCtx).Referer() []byte
          this.hasQualifiedName(packagePath(), "RequestCtx", "Referer") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestCtx).RequestBodyStream() io.Reader
          this.hasQualifiedName(packagePath(), "RequestCtx", "RequestBodyStream") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestCtx).RequestURI() []byte
          this.hasQualifiedName(packagePath(), "RequestCtx", "RequestURI") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestCtx).SetBody(body []byte)
          this.hasQualifiedName(packagePath(), "RequestCtx", "SetBody") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestCtx).SetBodyStream(bodyStream io.Reader, bodySize int)
          this.hasQualifiedName(packagePath(), "RequestCtx", "SetBodyStream") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestCtx).SetBodyString(body string)
          this.hasQualifiedName(packagePath(), "RequestCtx", "SetBodyString") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestCtx).String() string
          this.hasQualifiedName(packagePath(), "RequestCtx", "String") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestCtx).URI() *URI
          this.hasQualifiedName(packagePath(), "RequestCtx", "URI") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestCtx).UserAgent() []byte
          this.hasQualifiedName(packagePath(), "RequestCtx", "UserAgent") and
          inp.isReceiver() and
          out.isResult()
        )
        or
        // Receiver type: RequestHeader
        (
          // signature: func (*RequestHeader).Add(key string, value string)
          this.hasQualifiedName(packagePath(), "RequestHeader", "Add") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).AddBytesK(key []byte, value string)
          this.hasQualifiedName(packagePath(), "RequestHeader", "AddBytesK") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).AddBytesKV(key []byte, value []byte)
          this.hasQualifiedName(packagePath(), "RequestHeader", "AddBytesKV") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).AddBytesV(key string, value []byte)
          this.hasQualifiedName(packagePath(), "RequestHeader", "AddBytesV") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).AppendBytes(dst []byte) []byte
          this.hasQualifiedName(packagePath(), "RequestHeader", "AppendBytes") and
          (
            inp.isReceiver() and
            (
              out.isParameter(0) or
              out.isResult()
            )
            or
            inp.isParameter(0) and
            out.isResult()
          )
          or
          // signature: func (*RequestHeader).ContentType() []byte
          this.hasQualifiedName(packagePath(), "RequestHeader", "ContentType") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestHeader).Cookie(key string) []byte
          this.hasQualifiedName(packagePath(), "RequestHeader", "Cookie") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestHeader).CookieBytes(key []byte) []byte
          this.hasQualifiedName(packagePath(), "RequestHeader", "CookieBytes") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestHeader).CopyTo(dst *RequestHeader)
          this.hasQualifiedName(packagePath(), "RequestHeader", "CopyTo") and
          inp.isReceiver() and
          out.isParameter(0)
          or
          // signature: func (*RequestHeader).Header() []byte
          this.hasQualifiedName(packagePath(), "RequestHeader", "Header") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestHeader).Host() []byte
          this.hasQualifiedName(packagePath(), "RequestHeader", "Host") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestHeader).Method() []byte
          this.hasQualifiedName(packagePath(), "RequestHeader", "Method") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestHeader).MultipartFormBoundary() []byte
          this.hasQualifiedName(packagePath(), "RequestHeader", "MultipartFormBoundary") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestHeader).Peek(key string) []byte
          this.hasQualifiedName(packagePath(), "RequestHeader", "Peek") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestHeader).PeekBytes(key []byte) []byte
          this.hasQualifiedName(packagePath(), "RequestHeader", "PeekBytes") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestHeader).Protocol() []byte
          this.hasQualifiedName(packagePath(), "RequestHeader", "Protocol") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestHeader).RawHeaders() []byte
          this.hasQualifiedName(packagePath(), "RequestHeader", "RawHeaders") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestHeader).Read(r *bufio.Reader) error
          this.hasQualifiedName(packagePath(), "RequestHeader", "Read") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).Referer() []byte
          this.hasQualifiedName(packagePath(), "RequestHeader", "Referer") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestHeader).RequestURI() []byte
          this.hasQualifiedName(packagePath(), "RequestHeader", "RequestURI") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestHeader).Set(key string, value string)
          this.hasQualifiedName(packagePath(), "RequestHeader", "Set") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetBytesK(key []byte, value string)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetBytesK") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetBytesKV(key []byte, value []byte)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetBytesKV") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetBytesV(key string, value []byte)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetBytesV") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetCanonical(key []byte, value []byte)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetCanonical") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetContentType(contentType string)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetContentType") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetContentTypeBytes(contentType []byte)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetContentTypeBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetCookie(key string, value string)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetCookie") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetCookieBytesK(key []byte, value string)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetCookieBytesK") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetCookieBytesKV(key []byte, value []byte)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetCookieBytesKV") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetHost(host string)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetHost") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetHostBytes(host []byte)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetHostBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetMethod(method string)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetMethod") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetMethodBytes(method []byte)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetMethodBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetMultipartFormBoundary(boundary string)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetMultipartFormBoundary") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetMultipartFormBoundaryBytes(boundary []byte)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetMultipartFormBoundaryBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetProtocol(method string)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetProtocol") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetProtocolBytes(method []byte)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetProtocolBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetReferer(referer string)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetReferer") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetRefererBytes(referer []byte)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetRefererBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetRequestURI(requestURI string)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetRequestURI") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetRequestURIBytes(requestURI []byte)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetRequestURIBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetUserAgent(userAgent string)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetUserAgent") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).SetUserAgentBytes(userAgent []byte)
          this.hasQualifiedName(packagePath(), "RequestHeader", "SetUserAgentBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*RequestHeader).String() string
          this.hasQualifiedName(packagePath(), "RequestHeader", "String") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestHeader).UserAgent() []byte
          this.hasQualifiedName(packagePath(), "RequestHeader", "UserAgent") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*RequestHeader).Write(w *bufio.Writer) error
          this.hasQualifiedName(packagePath(), "RequestHeader", "Write") and
          inp.isReceiver() and
          out.isParameter(0)
          or
          // signature: func (*RequestHeader).WriteTo(w io.Writer) (int64, error)
          this.hasQualifiedName(packagePath(), "RequestHeader", "WriteTo") and
          inp.isReceiver() and
          out.isParameter(0)
        )
        or
        // Receiver type: Response
        (
          // signature: func (*Response).AppendBody(p []byte)
          this.hasQualifiedName(packagePath(), "Response", "AppendBody") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Response).AppendBodyString(s string)
          this.hasQualifiedName(packagePath(), "Response", "AppendBodyString") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Response).Body() []byte
          this.hasQualifiedName(packagePath(), "Response", "Body") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Response).BodyGunzip() ([]byte, error)
          this.hasQualifiedName(packagePath(), "Response", "BodyGunzip") and
          inp.isReceiver() and
          out.isResult(0)
          or
          // signature: func (*Response).BodyInflate() ([]byte, error)
          this.hasQualifiedName(packagePath(), "Response", "BodyInflate") and
          inp.isReceiver() and
          out.isResult(0)
          or
          // signature: func (*Response).BodyUnbrotli() ([]byte, error)
          this.hasQualifiedName(packagePath(), "Response", "BodyUnbrotli") and
          inp.isReceiver() and
          out.isResult(0)
          or
          // signature: func (*Response).BodyWriteTo(w io.Writer) error
          this.hasQualifiedName(packagePath(), "Response", "BodyWriteTo") and
          inp.isReceiver() and
          out.isParameter(0)
          or
          // signature: func (*Response).BodyWriter() io.Writer
          this.hasQualifiedName(packagePath(), "Response", "BodyWriter") and
          inp.isResult() and
          out.isReceiver()
          or
          // signature: func (*Response).CopyTo(dst *Response)
          this.hasQualifiedName(packagePath(), "Response", "CopyTo") and
          inp.isReceiver() and
          out.isParameter(0)
          or
          // signature: func (*Response).Read(r *bufio.Reader) error
          this.hasQualifiedName(packagePath(), "Response", "Read") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Response).ReadLimitBody(r *bufio.Reader, maxBodySize int) error
          this.hasQualifiedName(packagePath(), "Response", "ReadLimitBody") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Response).SetBody(body []byte)
          this.hasQualifiedName(packagePath(), "Response", "SetBody") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Response).SetBodyRaw(body []byte)
          this.hasQualifiedName(packagePath(), "Response", "SetBodyRaw") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Response).SetBodyStream(bodyStream io.Reader, bodySize int)
          this.hasQualifiedName(packagePath(), "Response", "SetBodyStream") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Response).SetBodyStreamWriter(sw StreamWriter)
          this.hasQualifiedName(packagePath(), "Response", "SetBodyStreamWriter") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Response).SetBodyString(body string)
          this.hasQualifiedName(packagePath(), "Response", "SetBodyString") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*Response).String() string
          this.hasQualifiedName(packagePath(), "Response", "String") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*Response).SwapBody(body []byte) []byte
          this.hasQualifiedName(packagePath(), "Response", "SwapBody") and
          (
            inp.isParameter(0) and
            out.isReceiver()
            or
            inp.isReceiver() and
            out.isResult()
          )
          or
          // signature: func (*Response).Write(w *bufio.Writer) error
          this.hasQualifiedName(packagePath(), "Response", "Write") and
          inp.isReceiver() and
          out.isParameter(0)
          or
          // signature: func (*Response).WriteDeflate(w *bufio.Writer) error
          this.hasQualifiedName(packagePath(), "Response", "WriteDeflate") and
          inp.isReceiver() and
          out.isParameter(0)
          or
          // signature: func (*Response).WriteDeflateLevel(w *bufio.Writer, level int) error
          this.hasQualifiedName(packagePath(), "Response", "WriteDeflateLevel") and
          inp.isReceiver() and
          out.isParameter(0)
          or
          // signature: func (*Response).WriteGzip(w *bufio.Writer) error
          this.hasQualifiedName(packagePath(), "Response", "WriteGzip") and
          inp.isReceiver() and
          out.isParameter(0)
          or
          // signature: func (*Response).WriteGzipLevel(w *bufio.Writer, level int) error
          this.hasQualifiedName(packagePath(), "Response", "WriteGzipLevel") and
          inp.isReceiver() and
          out.isParameter(0)
          or
          // signature: func (*Response).WriteTo(w io.Writer) (int64, error)
          this.hasQualifiedName(packagePath(), "Response", "WriteTo") and
          inp.isReceiver() and
          out.isParameter(0)
        )
        or
        // Receiver type: ResponseHeader
        (
          // signature: func (*ResponseHeader).Add(key string, value string)
          this.hasQualifiedName(packagePath(), "ResponseHeader", "Add") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*ResponseHeader).AddBytesK(key []byte, value string)
          this.hasQualifiedName(packagePath(), "ResponseHeader", "AddBytesK") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*ResponseHeader).AddBytesKV(key []byte, value []byte)
          this.hasQualifiedName(packagePath(), "ResponseHeader", "AddBytesKV") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*ResponseHeader).AddBytesV(key string, value []byte)
          this.hasQualifiedName(packagePath(), "ResponseHeader", "AddBytesV") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*ResponseHeader).AppendBytes(dst []byte) []byte
          this.hasQualifiedName(packagePath(), "ResponseHeader", "AppendBytes") and
          (
            inp.isReceiver() and
            (
              out.isParameter(0) or
              out.isResult()
            )
            or
            inp.isParameter(0) and
            out.isResult()
          )
          or
          // signature: func (*ResponseHeader).ContentType() []byte
          this.hasQualifiedName(packagePath(), "ResponseHeader", "ContentType") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*ResponseHeader).Cookie(cookie *Cookie) bool
          this.hasQualifiedName(packagePath(), "ResponseHeader", "Cookie") and
          inp.isReceiver() and
          out.isParameter(0)
          or
          // signature: func (*ResponseHeader).CopyTo(dst *ResponseHeader)
          this.hasQualifiedName(packagePath(), "ResponseHeader", "CopyTo") and
          inp.isReceiver() and
          out.isParameter(0)
          or
          // signature: func (*ResponseHeader).Header() []byte
          this.hasQualifiedName(packagePath(), "ResponseHeader", "Header") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*ResponseHeader).Peek(key string) []byte
          this.hasQualifiedName(packagePath(), "ResponseHeader", "Peek") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*ResponseHeader).PeekBytes(key []byte) []byte
          this.hasQualifiedName(packagePath(), "ResponseHeader", "PeekBytes") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*ResponseHeader).PeekCookie(key string) []byte
          this.hasQualifiedName(packagePath(), "ResponseHeader", "PeekCookie") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*ResponseHeader).Read(r *bufio.Reader) error
          this.hasQualifiedName(packagePath(), "ResponseHeader", "Read") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*ResponseHeader).Server() []byte
          this.hasQualifiedName(packagePath(), "ResponseHeader", "Server") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*ResponseHeader).Set(key string, value string)
          this.hasQualifiedName(packagePath(), "ResponseHeader", "Set") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*ResponseHeader).SetBytesK(key []byte, value string)
          this.hasQualifiedName(packagePath(), "ResponseHeader", "SetBytesK") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*ResponseHeader).SetBytesKV(key []byte, value []byte)
          this.hasQualifiedName(packagePath(), "ResponseHeader", "SetBytesKV") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*ResponseHeader).SetBytesV(key string, value []byte)
          this.hasQualifiedName(packagePath(), "ResponseHeader", "SetBytesV") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*ResponseHeader).SetCanonical(key []byte, value []byte)
          this.hasQualifiedName(packagePath(), "ResponseHeader", "SetCanonical") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*ResponseHeader).SetContentType(contentType string)
          this.hasQualifiedName(packagePath(), "ResponseHeader", "SetContentType") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*ResponseHeader).SetContentTypeBytes(contentType []byte)
          this.hasQualifiedName(packagePath(), "ResponseHeader", "SetContentTypeBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*ResponseHeader).SetCookie(cookie *Cookie)
          this.hasQualifiedName(packagePath(), "ResponseHeader", "SetCookie") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*ResponseHeader).SetServer(server string)
          this.hasQualifiedName(packagePath(), "ResponseHeader", "SetServer") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*ResponseHeader).SetServerBytes(server []byte)
          this.hasQualifiedName(packagePath(), "ResponseHeader", "SetServerBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*ResponseHeader).String() string
          this.hasQualifiedName(packagePath(), "ResponseHeader", "String") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*ResponseHeader).Write(w *bufio.Writer) error
          this.hasQualifiedName(packagePath(), "ResponseHeader", "Write") and
          inp.isReceiver() and
          out.isParameter(0)
          or
          // signature: func (*ResponseHeader).WriteTo(w io.Writer) (int64, error)
          this.hasQualifiedName(packagePath(), "ResponseHeader", "WriteTo") and
          inp.isReceiver() and
          out.isParameter(0)
        )
        or
        // Receiver type: URI
        (
          // signature: func (*URI).AppendBytes(dst []byte) []byte
          this.hasQualifiedName(packagePath(), "URI", "AppendBytes") and
          (
            inp.isReceiver() and
            (
              out.isParameter(0) or
              out.isResult()
            )
            or
            inp.isParameter(0) and
            out.isResult()
          )
          or
          // signature: func (*URI).CopyTo(dst *URI)
          this.hasQualifiedName(packagePath(), "URI", "CopyTo") and
          inp.isReceiver() and
          out.isParameter(0)
          or
          // signature: func (*URI).FullURI() []byte
          this.hasQualifiedName(packagePath(), "URI", "FullURI") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*URI).Hash() []byte
          this.hasQualifiedName(packagePath(), "URI", "Hash") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*URI).Host() []byte
          this.hasQualifiedName(packagePath(), "URI", "Host") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*URI).LastPathSegment() []byte
          this.hasQualifiedName(packagePath(), "URI", "LastPathSegment") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*URI).Parse(host []byte, uri []byte) error
          this.hasQualifiedName(packagePath(), "URI", "Parse") and
          inp.isParameter(_) and
          out.isReceiver()
          or
          // signature: func (*URI).Password() []byte
          this.hasQualifiedName(packagePath(), "URI", "Password") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*URI).Path() []byte
          this.hasQualifiedName(packagePath(), "URI", "Path") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*URI).PathOriginal() []byte
          this.hasQualifiedName(packagePath(), "URI", "PathOriginal") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*URI).QueryArgs() *Args
          this.hasQualifiedName(packagePath(), "URI", "QueryArgs") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*URI).QueryString() []byte
          this.hasQualifiedName(packagePath(), "URI", "QueryString") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*URI).RequestURI() []byte
          this.hasQualifiedName(packagePath(), "URI", "RequestURI") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*URI).Scheme() []byte
          this.hasQualifiedName(packagePath(), "URI", "Scheme") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*URI).SetHash(hash string)
          this.hasQualifiedName(packagePath(), "URI", "SetHash") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*URI).SetHashBytes(hash []byte)
          this.hasQualifiedName(packagePath(), "URI", "SetHashBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*URI).SetHost(host string)
          this.hasQualifiedName(packagePath(), "URI", "SetHost") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*URI).SetHostBytes(host []byte)
          this.hasQualifiedName(packagePath(), "URI", "SetHostBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*URI).SetPassword(password string)
          this.hasQualifiedName(packagePath(), "URI", "SetPassword") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*URI).SetPasswordBytes(password []byte)
          this.hasQualifiedName(packagePath(), "URI", "SetPasswordBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*URI).SetPath(path string)
          this.hasQualifiedName(packagePath(), "URI", "SetPath") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*URI).SetPathBytes(path []byte)
          this.hasQualifiedName(packagePath(), "URI", "SetPathBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*URI).SetQueryString(queryString string)
          this.hasQualifiedName(packagePath(), "URI", "SetQueryString") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*URI).SetQueryStringBytes(queryString []byte)
          this.hasQualifiedName(packagePath(), "URI", "SetQueryStringBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*URI).SetScheme(scheme string)
          this.hasQualifiedName(packagePath(), "URI", "SetScheme") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*URI).SetSchemeBytes(scheme []byte)
          this.hasQualifiedName(packagePath(), "URI", "SetSchemeBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*URI).SetUsername(username string)
          this.hasQualifiedName(packagePath(), "URI", "SetUsername") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*URI).SetUsernameBytes(username []byte)
          this.hasQualifiedName(packagePath(), "URI", "SetUsernameBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*URI).String() string
          this.hasQualifiedName(packagePath(), "URI", "String") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*URI).Update(newURI string)
          this.hasQualifiedName(packagePath(), "URI", "Update") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*URI).UpdateBytes(newURI []byte)
          this.hasQualifiedName(packagePath(), "URI", "UpdateBytes") and
          inp.isParameter(0) and
          out.isReceiver()
          or
          // signature: func (*URI).Username() []byte
          this.hasQualifiedName(packagePath(), "URI", "Username") and
          inp.isReceiver() and
          out.isResult()
          or
          // signature: func (*URI).WriteTo(w io.Writer) (int64, error)
          this.hasQualifiedName(packagePath(), "URI", "WriteTo") and
          inp.isReceiver() and
          out.isParameter(0)
        )
        or
        // Receiver interface: Logger
        // signature: func (Logger).Printf(format string, args ...interface{})
        this.implements(packagePath(), "Logger", "Printf") and
        inp.isParameter(_) and
        out.isReceiver()
      )
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = out
    }
  }

  /**
   * Models HTTP redirects.
   */
  private class Redirect extends HTTP::Redirect::Range, DataFlow::CallNode {
    string package;
    DataFlow::Node urlNode;

    Redirect() {
      // HTTP redirect models for package: github.com/valyala/fasthttp@v1.23.0
      package = packagePath() and
      // Receiver type: RequestCtx
      (
        // signature: func (*RequestCtx).Redirect(uri string, statusCode int)
        this = any(Method m | m.hasQualifiedName(package, "RequestCtx", "Redirect")).getACall() and
        urlNode = this.getArgument(0)
        or
        // signature: func (*RequestCtx).RedirectBytes(uri []byte, statusCode int)
        this = any(Method m | m.hasQualifiedName(package, "RequestCtx", "RedirectBytes")).getACall() and
        urlNode = this.getArgument(0)
      )
    }

    override DataFlow::Node getUrl() { result = urlNode }

    override HTTP::ResponseWriter getResponseWriter() { result.getANode() = this.getReceiver() }
  }

  /**
   * Models HTTP header writers.
   * The write is done with a call where you can set both the key and the value of the header.
   */
  private class HeaderWrite extends HTTP::HeaderWrite::Range, DataFlow::CallNode {
    DataFlow::Node receiverNode;
    DataFlow::Node headerNameNode;
    DataFlow::Node headerValueNode;

    HeaderWrite() {
      setsHeaderDynamicKeyValue(_, _, this, headerNameNode, headerValueNode, receiverNode)
    }

    override DataFlow::Node getName() { result = headerNameNode }

    override DataFlow::Node getValue() { result = headerValueNode }

    override HTTP::ResponseWriter getResponseWriter() { result.getANode() = receiverNode }
  }

  // Holds for a call that sets a header with a key-value combination.
  private predicate setsHeaderDynamicKeyValue(
    string package, string receiverName, DataFlow::CallNode headerSetterCall,
    DataFlow::Node headerNameNode, DataFlow::Node headerValueNode, DataFlow::Node receiverNode
  ) {
    exists(string methodName, Method met |
      met.hasQualifiedName(package, receiverName, methodName) and
      headerSetterCall = met.getACall() and
      receiverNode = headerSetterCall.getReceiver()
    |
      package = packagePath() and
      (
        // Receiver type: Args
        receiverName = "Args" and
        (
          // signature: func (*Args).Add(key string, value string)
          methodName = "Add" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
        ) // Receiver type: RequestHeader
        or
        receiverName = "RequestHeader" and
        (
          // signature: func (*RequestHeader).Add(key string, value string)
          methodName = "Add" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
          or
          // signature: func (*RequestHeader).AddBytesK(key []byte, value string)
          methodName = "AddBytesK" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
          or
          // signature: func (*RequestHeader).AddBytesKV(key []byte, value []byte)
          methodName = "AddBytesKV" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
          or
          // signature: func (*RequestHeader).AddBytesV(key string, value []byte)
          methodName = "AddBytesV" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
          or
          // signature: func (*RequestHeader).Set(key string, value string)
          methodName = "Set" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
          or
          // signature: func (*RequestHeader).SetBytesK(key []byte, value string)
          methodName = "SetBytesK" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
          or
          // signature: func (*RequestHeader).SetBytesKV(key []byte, value []byte)
          methodName = "SetBytesKV" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
          or
          // signature: func (*RequestHeader).SetBytesV(key string, value []byte)
          methodName = "SetBytesV" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
          or
          // signature: func (*RequestHeader).SetCanonical(key []byte, value []byte)
          methodName = "SetCanonical" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
        ) // Receiver type: ResponseHeader
        or
        receiverName = "ResponseHeader" and
        (
          // signature: func (*ResponseHeader).Add(key string, value string)
          methodName = "Add" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
          or
          // signature: func (*ResponseHeader).AddBytesK(key []byte, value string)
          methodName = "AddBytesK" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
          or
          // signature: func (*ResponseHeader).AddBytesKV(key []byte, value []byte)
          methodName = "AddBytesKV" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
          or
          // signature: func (*ResponseHeader).AddBytesV(key string, value []byte)
          methodName = "AddBytesV" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
          or
          // signature: func (*ResponseHeader).Set(key string, value string)
          methodName = "Set" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
          or
          // signature: func (*ResponseHeader).SetBytesK(key []byte, value string)
          methodName = "SetBytesK" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
          or
          // signature: func (*ResponseHeader).SetBytesKV(key []byte, value []byte)
          methodName = "SetBytesKV" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
          or
          // signature: func (*ResponseHeader).SetBytesV(key string, value []byte)
          methodName = "SetBytesV" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
          or
          // signature: func (*ResponseHeader).SetCanonical(key []byte, value []byte)
          methodName = "SetCanonical" and
          headerNameNode = headerSetterCall.getArgument(0) and
          headerValueNode = headerSetterCall.getArgument(1)
        )
      )
    )
  }

  /**
   * Models an HTTP dynamic `content-type` header setter.
   */
  private class DynamicContentTypeHeaderSetter extends HTTP::HeaderWrite::Range, DataFlow::CallNode {
    DataFlow::Node receiverNode;
    DataFlow::Node valueNode;

    DynamicContentTypeHeaderSetter() {
      setsDynamicHeaderContentType(_, _, this, valueNode, receiverNode)
    }

    override string getHeaderName() { result = "content-type" }

    override DataFlow::Node getName() { none() }

    override DataFlow::Node getValue() { result = valueNode }

    override HTTP::ResponseWriter getResponseWriter() { result.getANode() = receiverNode }
  }

  // Holds for a call that sets the `content-type` header via a parameter.
  private predicate setsDynamicHeaderContentType(
    string package, string receiverName, DataFlow::CallNode setterCall, DataFlow::Node valueNode,
    DataFlow::Node receiverNode
  ) {
    exists(string methodName, Method met |
      met.hasQualifiedName(package, receiverName, methodName) and
      setterCall = met.getACall() and
      receiverNode = setterCall.getReceiver()
    |
      package = packagePath() and
      (
        // Receiver type: RequestCtx
        receiverName = "RequestCtx" and
        (
          // signature: func (*RequestCtx).SetContentType(contentType string)
          methodName = "SetContentType" and
          valueNode = setterCall.getArgument(0)
          or
          // signature: func (*RequestCtx).SetContentTypeBytes(contentType []byte)
          methodName = "SetContentTypeBytes" and
          valueNode = setterCall.getArgument(0)
        ) // Receiver type: RequestHeader
        or
        receiverName = "RequestHeader" and
        (
          // signature: func (*RequestHeader).SetContentType(contentType string)
          methodName = "SetContentType" and
          valueNode = setterCall.getArgument(0)
          or
          // signature: func (*RequestHeader).SetContentTypeBytes(contentType []byte)
          methodName = "SetContentTypeBytes" and
          valueNode = setterCall.getArgument(0)
        ) // Receiver type: ResponseHeader
        or
        receiverName = "ResponseHeader" and
        (
          // signature: func (*ResponseHeader).SetContentType(contentType string)
          methodName = "SetContentType" and
          valueNode = setterCall.getArgument(0)
          or
          // signature: func (*ResponseHeader).SetContentTypeBytes(contentType []byte)
          methodName = "SetContentTypeBytes" and
          valueNode = setterCall.getArgument(0)
        )
      )
    )
  }

  /**
   * Models HTTP ResponseBody where the content-type is static and non-modifiable.
   */
  private class ResponseBodyStaticContentType extends HTTP::ResponseBody::Range {
    string contentTypeString;
    DataFlow::Node receiverNode;

    ResponseBodyStaticContentType() {
      exists(string package, string receiverName |
        setsBodyAndStaticContentType(package, receiverName, this, contentTypeString, receiverNode)
      )
    }

    override string getAContentType() { result = contentTypeString }

    override HTTP::ResponseWriter getResponseWriter() { result.getANode() = receiverNode }
  }

  // Holds for a call that sets the body; the content-type is implicitly set.
  private predicate setsBodyAndStaticContentType(
    string package, string receiverName, DataFlow::Node bodyNode, string contentTypeString,
    DataFlow::Node receiverNode
  ) {
    exists(string methodName, Method met, DataFlow::CallNode bodySetterCall |
      met.hasQualifiedName(package, receiverName, methodName) and
      bodySetterCall = met.getACall() and
      receiverNode = bodySetterCall.getReceiver()
    |
      package = packagePath() and
      (
        // Receiver type: RequestCtx
        receiverName = "RequestCtx" and
        (
          // signature: func (*RequestCtx).Error(msg string, statusCode int)
          methodName = "Error" and
          bodyNode = bodySetterCall.getArgument(0) and
          contentTypeString = "text/plain"
        )
      )
    )
  }

  /**
   * Models HTTP ResponseBody where the content-type can be dynamically set by the caller.
   */
  private class ResponseBodyDynamicContentType extends HTTP::ResponseBody::Range {
    DataFlow::Node contentTypeNode;
    DataFlow::Node receiverNode;

    ResponseBodyDynamicContentType() {
      exists(string package, string receiverName |
        setsBodyAndDynamicContentType(package, receiverName, this, contentTypeNode, receiverNode)
      )
    }

    override DataFlow::Node getAContentTypeNode() { result = contentTypeNode }

    override HTTP::ResponseWriter getResponseWriter() { result.getANode() = receiverNode }
  }

  // Holds for a call that sets the body; the content-type is a parameter.
  // Both body and content-type are parameters in the same func call.
  private predicate setsBodyAndDynamicContentType(
    string package, string receiverName, DataFlow::Node bodyNode, DataFlow::Node contentTypeNode,
    DataFlow::Node receiverNode
  ) {
    exists(string methodName, Method met, DataFlow::CallNode bodySetterCall |
      met.hasQualifiedName(package, receiverName, methodName) and
      bodySetterCall = met.getACall() and
      receiverNode = bodySetterCall.getReceiver()
    |
      package = packagePath() and
      (
        // Receiver type: RequestCtx
        receiverName = "RequestCtx" and
        (
          // signature: func (*RequestCtx).Success(contentType string, body []byte)
          methodName = "Success" and
          bodyNode = bodySetterCall.getArgument(1) and
          contentTypeNode = bodySetterCall.getArgument(0)
          or
          // signature: func (*RequestCtx).SuccessString(contentType string, body string)
          methodName = "SuccessString" and
          bodyNode = bodySetterCall.getArgument(1) and
          contentTypeNode = bodySetterCall.getArgument(0)
        )
      )
    )
  }

  /**
   * Models HTTP ResponseBody where only the body is set.
   */
  private class ResponseBodyNoContentType extends HTTP::ResponseBody::Range {
    DataFlow::Node receiverNode;

    ResponseBodyNoContentType() {
      exists(string package, string receiverName |
        setsBody(package, receiverName, receiverNode, this)
      )
    }

    override HTTP::ResponseWriter getResponseWriter() { result.getANode() = receiverNode }
  }

  // Holds for a call that sets the body. The content-type is not defined.
  private predicate setsBody(
    string package, string receiverName, DataFlow::Node receiverNode, DataFlow::Node bodyNode
  ) {
    exists(string methodName, Method met, DataFlow::CallNode bodySetterCall |
      met.hasQualifiedName(package, receiverName, methodName) and
      bodySetterCall = met.getACall() and
      receiverNode = bodySetterCall.getReceiver()
    |
      package = packagePath() and
      (
        // Receiver type: Request
        receiverName = "Request" and
        (
          // signature: func (*Request).AppendBody(p []byte)
          methodName = "AppendBody" and
          bodyNode = bodySetterCall.getArgument(0)
          or
          // signature: func (*Request).AppendBodyString(s string)
          methodName = "AppendBodyString" and
          bodyNode = bodySetterCall.getArgument(0)
        ) // Receiver type: RequestCtx
        or
        receiverName = "RequestCtx" and
        (
          // signature: func (*RequestCtx).SetBody(body []byte)
          methodName = "SetBody" and
          bodyNode = bodySetterCall.getArgument(0)
          or
          // signature: func (*RequestCtx).SetBodyStream(bodyStream io.Reader, bodySize int)
          methodName = "SetBodyStream" and
          bodyNode = bodySetterCall.getArgument(0)
          or
          // signature: func (*RequestCtx).SetBodyStreamWriter(sw StreamWriter)
          methodName = "SetBodyStreamWriter" and
          bodyNode = bodySetterCall.getArgument(0)
          or
          // signature: func (*RequestCtx).SetBodyString(body string)
          methodName = "SetBodyString" and
          bodyNode = bodySetterCall.getArgument(0)
          or
          // signature: func (*RequestCtx).Write(p []byte) (int, error)
          methodName = "Write" and
          bodyNode = bodySetterCall.getArgument(0)
          or
          // signature: func (*RequestCtx).WriteString(s string) (int, error)
          methodName = "WriteString" and
          bodyNode = bodySetterCall.getArgument(0)
        ) // Receiver type: Response
        or
        receiverName = "Response" and
        (
          // signature: func (*Response).AppendBody(p []byte)
          methodName = "AppendBody" and
          bodyNode = bodySetterCall.getArgument(0)
          or
          // signature: func (*Response).AppendBodyString(s string)
          methodName = "AppendBodyString" and
          bodyNode = bodySetterCall.getArgument(0)
        )
      )
    )
  }

  /**
   * Provides models of untrusted flow sources.
   */
  private class UntrustedFlowSources extends UntrustedFlowSource::Range {
    UntrustedFlowSources() {
      // Methods on types of package: github.com/valyala/fasthttp@v1.23.0
      exists(string receiverName, string methodName, Method mtd, FunctionOutput out |
        this = out.getExitNode(mtd.getACall()) and
        mtd.hasQualifiedName(packagePath(), receiverName, methodName)
      |
        receiverName = "Args" and
        (
          // signature: func (*Args).Peek(key string) []byte
          methodName = "Peek" and
          out.isResult()
          or
          // signature: func (*Args).PeekBytes(key []byte) []byte
          methodName = "PeekBytes" and
          out.isResult()
          or
          // signature: func (*Args).PeekMulti(key string) [][]byte
          methodName = "PeekMulti" and
          out.isResult()
          or
          // signature: func (*Args).PeekMultiBytes(key []byte) [][]byte
          methodName = "PeekMultiBytes" and
          out.isResult()
          or
          // signature: func (*Args).QueryString() []byte
          methodName = "QueryString" and
          out.isResult()
          or
          // signature: func (*Args).String() string
          methodName = "String" and
          out.isResult()
        )
        or
        receiverName = "Cookie" and
        (
          // signature: func (*Cookie).Cookie() []byte
          methodName = "Cookie" and
          out.isResult()
          or
          // signature: func (*Cookie).Domain() []byte
          methodName = "Domain" and
          out.isResult()
          or
          // signature: func (*Cookie).Key() []byte
          methodName = "Key" and
          out.isResult()
          or
          // signature: func (*Cookie).Path() []byte
          methodName = "Path" and
          out.isResult()
          or
          // signature: func (*Cookie).String() string
          methodName = "String" and
          out.isResult()
          or
          // signature: func (*Cookie).Value() []byte
          methodName = "Value" and
          out.isResult()
        )
        or
        receiverName = "Request" and
        (
          // signature: func (*Request).Body() []byte
          methodName = "Body" and
          out.isResult()
          or
          // signature: func (*Request).BodyGunzip() ([]byte, error)
          methodName = "BodyGunzip" and
          out.isResult(0)
          or
          // signature: func (*Request).BodyInflate() ([]byte, error)
          methodName = "BodyInflate" and
          out.isResult(0)
          or
          // signature: func (*Request).BodyUnbrotli() ([]byte, error)
          methodName = "BodyUnbrotli" and
          out.isResult(0)
          or
          // signature: func (*Request).BodyWriteTo(w io.Writer) error
          methodName = "BodyWriteTo" and
          out.isParameter(0)
          or
          // signature: func (*Request).Host() []byte
          methodName = "Host" and
          out.isResult()
          or
          // signature: func (*Request).MultipartForm() (*mime/multipart.Form, error)
          methodName = "MultipartForm" and
          out.isResult(0)
          or
          // signature: func (*Request).PostArgs() *Args
          methodName = "PostArgs" and
          out.isResult()
          or
          // signature: func (*Request).RequestURI() []byte
          methodName = "RequestURI" and
          out.isResult()
          or
          // signature: func (*Request).String() string
          methodName = "String" and
          out.isResult()
          or
          // signature: func (*Request).URI() *URI
          methodName = "URI" and
          out.isResult()
        )
        or
        receiverName = "RequestCtx" and
        (
          // signature: func (*RequestCtx).FormFile(key string) (*mime/multipart.FileHeader, error)
          methodName = "FormFile" and
          out.isResult(0)
          or
          // signature: func (*RequestCtx).FormValue(key string) []byte
          methodName = "FormValue" and
          out.isResult()
          or
          // signature: func (*RequestCtx).Host() []byte
          methodName = "Host" and
          out.isResult()
          or
          // signature: func (*RequestCtx).Method() []byte
          methodName = "Method" and
          out.isResult()
          or
          // signature: func (*RequestCtx).MultipartForm() (*mime/multipart.Form, error)
          methodName = "MultipartForm" and
          out.isResult(0)
          or
          // signature: func (*RequestCtx).Path() []byte
          methodName = "Path" and
          out.isResult()
          or
          // signature: func (*RequestCtx).PostArgs() *Args
          methodName = "PostArgs" and
          out.isResult()
          or
          // signature: func (*RequestCtx).PostBody() []byte
          methodName = "PostBody" and
          out.isResult()
          or
          // signature: func (*RequestCtx).QueryArgs() *Args
          methodName = "QueryArgs" and
          out.isResult()
          or
          // signature: func (*RequestCtx).Referer() []byte
          methodName = "Referer" and
          out.isResult()
          or
          // signature: func (*RequestCtx).RequestURI() []byte
          methodName = "RequestURI" and
          out.isResult()
          or
          // signature: func (*RequestCtx).String() string
          methodName = "String" and
          out.isResult()
          or
          // signature: func (*RequestCtx).UserAgent() []byte
          methodName = "UserAgent" and
          out.isResult()
        )
        or
        receiverName = "RequestHeader" and
        (
          // signature: func (*RequestHeader).ContentType() []byte
          methodName = "ContentType" and
          out.isResult()
          or
          // signature: func (*RequestHeader).Header() []byte
          methodName = "Header" and
          out.isResult()
          or
          // signature: func (*RequestHeader).Host() []byte
          methodName = "Host" and
          out.isResult()
          or
          // signature: func (*RequestHeader).Method() []byte
          methodName = "Method" and
          out.isResult()
          or
          // signature: func (*RequestHeader).MultipartFormBoundary() []byte
          methodName = "MultipartFormBoundary" and
          out.isResult()
          or
          // signature: func (*RequestHeader).Peek(key string) []byte
          methodName = "Peek" and
          out.isResult()
          or
          // signature: func (*RequestHeader).PeekBytes(key []byte) []byte
          methodName = "PeekBytes" and
          out.isResult()
          or
          // signature: func (*RequestHeader).RawHeaders() []byte
          methodName = "RawHeaders" and
          out.isResult()
          or
          // signature: func (*RequestHeader).Referer() []byte
          methodName = "Referer" and
          out.isResult()
          or
          // signature: func (*RequestHeader).RequestURI() []byte
          methodName = "RequestURI" and
          out.isResult()
          or
          // signature: func (*RequestHeader).String() string
          methodName = "String" and
          out.isResult()
          or
          // signature: func (*RequestHeader).UserAgent() []byte
          methodName = "UserAgent" and
          out.isResult()
        )
        or
        receiverName = "ResponseHeader" and
        (
          // signature: func (*ResponseHeader).ContentType() []byte
          methodName = "ContentType" and
          out.isResult()
          or
          // signature: func (*ResponseHeader).Header() []byte
          methodName = "Header" and
          out.isResult()
          or
          // signature: func (*ResponseHeader).Server() []byte
          methodName = "Server" and
          out.isResult()
          or
          // signature: func (*ResponseHeader).String() string
          methodName = "String" and
          out.isResult()
        )
        or
        receiverName = "URI" and
        (
          // signature: func (*URI).FullURI() []byte
          methodName = "FullURI" and
          out.isResult()
          or
          // signature: func (*URI).Hash() []byte
          methodName = "Hash" and
          out.isResult()
          or
          // signature: func (*URI).Host() []byte
          methodName = "Host" and
          out.isResult()
          or
          // signature: func (*URI).LastPathSegment() []byte
          methodName = "LastPathSegment" and
          out.isResult()
          or
          // signature: func (*URI).Password() []byte
          methodName = "Password" and
          out.isResult()
          or
          // signature: func (*URI).Path() []byte
          methodName = "Path" and
          out.isResult()
          or
          // signature: func (*URI).PathOriginal() []byte
          methodName = "PathOriginal" and
          out.isResult()
          or
          // signature: func (*URI).QueryArgs() *Args
          methodName = "QueryArgs" and
          out.isResult()
          or
          // signature: func (*URI).QueryString() []byte
          methodName = "QueryString" and
          out.isResult()
          or
          // signature: func (*URI).RequestURI() []byte
          methodName = "RequestURI" and
          out.isResult()
          or
          // signature: func (*URI).Scheme() []byte
          methodName = "Scheme" and
          out.isResult()
          or
          // signature: func (*URI).String() string
          methodName = "String" and
          out.isResult()
          or
          // signature: func (*URI).Username() []byte
          methodName = "Username" and
          out.isResult()
        )
      )
      or
      // Structs of package: github.com/valyala/fasthttp@v1.23.0
      exists(string structName, string fields, DataFlow::Field fld |
        this = fld.getARead() and
        fld.hasQualifiedName(packagePath(), structName, fields)
      |
        structName = "Request" and
        fields = "Header"
        or
        structName = "RequestCtx" and
        fields = "Request"
      )
      or
      // Types of package: github.com/valyala/fasthttp@v1.23.0
      exists(ValueEntity v | v.getType().hasQualifiedName(packagePath(), ["Args", "Cookie"]) |
        this = v.getARead()
      )
    }
  }
}
