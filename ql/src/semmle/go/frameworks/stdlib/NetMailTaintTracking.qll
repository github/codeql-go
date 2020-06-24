/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `net/mail` package. */
module NetMailTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func ParseAddress(address string) (*Address, error)
      hasQualifiedName("net/mail", "ParseAddress") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func ParseAddressList(list string) ([]*Address, error)
      hasQualifiedName("net/mail", "ParseAddressList") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func ReadMessage(r io.Reader) (msg *Message, err error)
      hasQualifiedName("net/mail", "ReadMessage") and
      (inp.isParameter(0) and outp.isResult(0))
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }

  private class MethodAndInterfaceTaintTracking extends TaintTracking::FunctionModel, Method {
    FunctionInput inp;
    FunctionOutput outp;

    MethodAndInterfaceTaintTracking() {
      // Methods:
      // signature: func (*AddressParser).Parse(address string) (*Address, error)
      this.(Method).hasQualifiedName("net/mail", "AddressParser", "Parse") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func (*AddressParser).ParseList(list string) ([]*Address, error)
      this.(Method).hasQualifiedName("net/mail", "AddressParser", "ParseList") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func (Header).Get(key string) string
      this.(Method).hasQualifiedName("net/mail", "Header", "Get") and
      (inp.isReceiver() and outp.isResult())
      // Interfaces:
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
