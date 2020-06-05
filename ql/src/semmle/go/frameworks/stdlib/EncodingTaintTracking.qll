/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `encoding` package. */
module EncodingTaintTracking {
  private class BinaryMarshalerMarshalBinary extends TaintTracking::FunctionModel, Method {
    // signature: func (BinaryMarshaler).MarshalBinary() (data []byte, err error)
    BinaryMarshalerMarshalBinary() {
      this.implements("encoding", "BinaryMarshaler", "MarshalBinary")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class BinaryUnmarshalerUnmarshalBinary extends TaintTracking::FunctionModel, Method {
    // signature: func (BinaryUnmarshaler).UnmarshalBinary(data []byte) error
    BinaryUnmarshalerUnmarshalBinary() {
      this.implements("encoding", "BinaryUnmarshaler", "UnmarshalBinary")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class TextMarshalerMarshalText extends TaintTracking::FunctionModel, Method {
    // signature: func (TextMarshaler).MarshalText() (text []byte, err error)
    TextMarshalerMarshalText() { this.implements("encoding", "TextMarshaler", "MarshalText") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class TextUnmarshalerUnmarshalText extends TaintTracking::FunctionModel, Method {
    // signature: func (TextUnmarshaler).UnmarshalText(text []byte) error
    TextUnmarshalerUnmarshalText() {
      this.implements("encoding", "TextUnmarshaler", "UnmarshalText")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }
}
