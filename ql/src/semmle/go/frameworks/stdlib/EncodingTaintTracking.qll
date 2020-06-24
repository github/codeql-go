/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `encoding` package. */
module EncodingTaintTracking {
  private class MethodAndInterfaceTaintTracking extends TaintTracking::FunctionModel, Method {
    FunctionInput inp;
    FunctionOutput outp;

    MethodAndInterfaceTaintTracking() {
      // Methods:
      // Interfaces:
      // signature: func (BinaryMarshaler).MarshalBinary() (data []byte, err error)
      this.implements("encoding", "BinaryMarshaler", "MarshalBinary") and
      (inp.isReceiver() and outp.isResult(0))
      or
      // signature: func (TextMarshaler).MarshalText() (text []byte, err error)
      this.implements("encoding", "TextMarshaler", "MarshalText") and
      (inp.isReceiver() and outp.isResult(0))
      or
      // signature: func (BinaryUnmarshaler).UnmarshalBinary(data []byte) error
      this.implements("encoding", "BinaryUnmarshaler", "UnmarshalBinary") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (TextUnmarshaler).UnmarshalText(text []byte) error
      this.implements("encoding", "TextUnmarshaler", "UnmarshalText") and
      (inp.isParameter(0) and outp.isReceiver())
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
