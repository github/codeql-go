/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `net/http` package. */
module NetHttpTaintTracking {
  private class CanonicalHeaderKey extends TaintTracking::FunctionModel {
    // signature: func CanonicalHeaderKey(s string) string
    CanonicalHeaderKey() { hasQualifiedName("net/http", "CanonicalHeaderKey") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class DetectContentType extends TaintTracking::FunctionModel {
    // signature: func DetectContentType(data []byte) string
    DetectContentType() { hasQualifiedName("net/http", "DetectContentType") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class NetHttpError extends TaintTracking::FunctionModel {
    // signature: func Error(w ResponseWriter, error string, code int)
    NetHttpError() { hasQualifiedName("net/http", "Error") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isParameter(0))
    }
  }

  private class MaxBytesReader extends TaintTracking::FunctionModel {
    // signature: func MaxBytesReader(w ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser
    MaxBytesReader() { hasQualifiedName("net/http", "MaxBytesReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isResult())
    }
  }

  private class NewRequest extends TaintTracking::FunctionModel {
    // signature: func NewRequest(method string, url string, body io.Reader) (*Request, error)
    NewRequest() { hasQualifiedName("net/http", "NewRequest") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult(0))
    }
  }

  private class NewRequestWithContext extends TaintTracking::FunctionModel {
    // signature: func NewRequestWithContext(ctx context.Context, method string, url string, body io.Reader) (*Request, error)
    NewRequestWithContext() { hasQualifiedName("net/http", "NewRequestWithContext") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult(0))
    }
  }

  private class ReadRequest extends TaintTracking::FunctionModel {
    // signature: func ReadRequest(b *bufio.Reader) (*Request, error)
    ReadRequest() { hasQualifiedName("net/http", "ReadRequest") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class ReadResponse extends TaintTracking::FunctionModel {
    // signature: func ReadResponse(r *bufio.Reader, req *Request) (*Response, error)
    ReadResponse() { hasQualifiedName("net/http", "ReadResponse") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class SetCookie extends TaintTracking::FunctionModel {
    // signature: func SetCookie(w ResponseWriter, cookie *Cookie)
    SetCookie() { hasQualifiedName("net/http", "SetCookie") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isParameter(0))
    }
  }

  private class HeaderAdd extends TaintTracking::FunctionModel, Method {
    // signature: func (Header).Add(key string, value string)
    HeaderAdd() { this.(Method).hasQualifiedName("net/http", "Header", "Add") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isReceiver())
    }
  }

  private class HeaderClone extends TaintTracking::FunctionModel, Method {
    // signature: func (Header).Clone() Header
    HeaderClone() { this.(Method).hasQualifiedName("net/http", "Header", "Clone") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class HeaderGet extends TaintTracking::FunctionModel, Method {
    // signature: func (Header).Get(key string) string
    HeaderGet() { this.(Method).hasQualifiedName("net/http", "Header", "Get") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class HeaderSet extends TaintTracking::FunctionModel, Method {
    // signature: func (Header).Set(key string, value string)
    HeaderSet() { this.(Method).hasQualifiedName("net/http", "Header", "Set") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isReceiver())
    }
  }

  private class HeaderValues extends TaintTracking::FunctionModel, Method {
    // signature: func (Header).Values(key string) []string
    HeaderValues() { this.(Method).hasQualifiedName("net/http", "Header", "Values") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class HeaderWrite extends TaintTracking::FunctionModel, Method {
    // signature: func (Header).Write(w io.Writer) error
    HeaderWrite() { this.(Method).hasQualifiedName("net/http", "Header", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class HeaderWriteSubset extends TaintTracking::FunctionModel, Method {
    // signature: func (Header).WriteSubset(w io.Writer, exclude map[string]bool) error
    HeaderWriteSubset() { this.(Method).hasQualifiedName("net/http", "Header", "WriteSubset") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class RequestAddCookie extends TaintTracking::FunctionModel, Method {
    // signature: func (*Request).AddCookie(c *Cookie)
    RequestAddCookie() { this.(Method).hasQualifiedName("net/http", "Request", "AddCookie") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class RequestClone extends TaintTracking::FunctionModel, Method {
    // signature: func (*Request).Clone(ctx context.Context) *Request
    RequestClone() { this.(Method).hasQualifiedName("net/http", "Request", "Clone") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class RequestWrite extends TaintTracking::FunctionModel, Method {
    // signature: func (*Request).Write(w io.Writer) error
    RequestWrite() { this.(Method).hasQualifiedName("net/http", "Request", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class RequestWriteProxy extends TaintTracking::FunctionModel, Method {
    // signature: func (*Request).WriteProxy(w io.Writer) error
    RequestWriteProxy() { this.(Method).hasQualifiedName("net/http", "Request", "WriteProxy") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class ResponseWrite extends TaintTracking::FunctionModel, Method {
    // signature: func (*Response).Write(w io.Writer) error
    ResponseWrite() { this.(Method).hasQualifiedName("net/http", "Response", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class HijackerHijack extends TaintTracking::FunctionModel, Method {
    // signature: func (Hijacker).Hijack() (net.Conn, *bufio.ReadWriter, error)
    HijackerHijack() { this.implements("net/http", "Hijacker", "Hijack") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult([0, 1]))
    }
  }

  private class ResponseWriterWrite extends TaintTracking::FunctionModel, Method {
    // signature: func (ResponseWriter).Write([]byte) (int, error)
    ResponseWriterWrite() { this.implements("net/http", "ResponseWriter", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }
}
