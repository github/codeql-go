/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `net/mail` package. */
module NetMailTaintTracking {
  private class ReadMessage extends TaintTracking::FunctionModel {
    // signature: func ReadMessage(r io.Reader) (msg *Message, err error)
    ReadMessage() { hasQualifiedName("net/mail", "ReadMessage") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class HeaderGet extends TaintTracking::FunctionModel, Method {
    // signature: func (Header).Get(key string) string
    HeaderGet() { this.(Method).hasQualifiedName("net/mail", "Header", "Get") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }
}
