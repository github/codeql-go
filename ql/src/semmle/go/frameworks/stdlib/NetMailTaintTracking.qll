/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `net/mail` package. */
module NetMailTaintTracking {
  private class ParseAddress extends TaintTracking::FunctionModel {
    // signature: func ParseAddress(address string) (*Address, error)
    ParseAddress() { hasQualifiedName("net/mail", "ParseAddress") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class ParseAddressList extends TaintTracking::FunctionModel {
    // signature: func ParseAddressList(list string) ([]*Address, error)
    ParseAddressList() { hasQualifiedName("net/mail", "ParseAddressList") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class ReadMessage extends TaintTracking::FunctionModel {
    // signature: func ReadMessage(r io.Reader) (msg *Message, err error)
    ReadMessage() { hasQualifiedName("net/mail", "ReadMessage") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class AddressParserParse extends TaintTracking::FunctionModel, Method {
    // signature: func (*AddressParser).Parse(address string) (*Address, error)
    AddressParserParse() { this.(Method).hasQualifiedName("net/mail", "AddressParser", "Parse") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class AddressParserParseList extends TaintTracking::FunctionModel, Method {
    // signature: func (*AddressParser).ParseList(list string) ([]*Address, error)
    AddressParserParseList() {
      this.(Method).hasQualifiedName("net/mail", "AddressParser", "ParseList")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class HeaderGet extends TaintTracking::FunctionModel, Method {
    // signature: func (Header).Get(key string) string
    HeaderGet() { this.(Method).hasQualifiedName("net/mail", "Header", "Get") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }
}
