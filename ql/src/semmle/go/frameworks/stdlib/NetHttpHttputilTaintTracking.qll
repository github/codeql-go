/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `net/http/httputil` package. */
module NetHttpHttputilTaintTracking {
  private class DumpRequest extends TaintTracking::FunctionModel {
    // signature: func DumpRequest(req *net/http.Request, body bool) ([]byte, error)
    DumpRequest() { hasQualifiedName("net/http/httputil", "DumpRequest") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class DumpRequestOut extends TaintTracking::FunctionModel {
    // signature: func DumpRequestOut(req *net/http.Request, body bool) ([]byte, error)
    DumpRequestOut() { hasQualifiedName("net/http/httputil", "DumpRequestOut") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class DumpResponse extends TaintTracking::FunctionModel {
    // signature: func DumpResponse(resp *net/http.Response, body bool) ([]byte, error)
    DumpResponse() { hasQualifiedName("net/http/httputil", "DumpResponse") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class NewChunkedReader extends TaintTracking::FunctionModel {
    // signature: func NewChunkedReader(r io.Reader) io.Reader
    NewChunkedReader() { hasQualifiedName("net/http/httputil", "NewChunkedReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class NewChunkedWriter extends TaintTracking::FunctionModel {
    // signature: func NewChunkedWriter(w io.Writer) io.WriteCloser
    NewChunkedWriter() { hasQualifiedName("net/http/httputil", "NewChunkedWriter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult() and outp.isParameter(0)
    }
  }

  private class BufferPoolGet extends TaintTracking::FunctionModel, Method {
    // signature: func (BufferPool).Get() []byte
    BufferPoolGet() { this.implements("net/http/httputil", "BufferPool", "Get") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class BufferPoolPut extends TaintTracking::FunctionModel, Method {
    // signature: func (BufferPool).Put([]byte)
    BufferPoolPut() { this.implements("net/http/httputil", "BufferPool", "Put") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }
}
