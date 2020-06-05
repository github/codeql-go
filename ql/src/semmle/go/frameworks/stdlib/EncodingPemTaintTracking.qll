/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `encoding/pem` package. */
module EncodingPemTaintTracking {
  private class Decode extends TaintTracking::FunctionModel {
    // signature: func Decode(data []byte) (p *Block, rest []byte)
    Decode() { hasQualifiedName("encoding/pem", "Decode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class Encode extends TaintTracking::FunctionModel {
    // signature: func Encode(out io.Writer, b *Block) error
    Encode() { hasQualifiedName("encoding/pem", "Encode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class EncodeToMemory extends TaintTracking::FunctionModel {
    // signature: func EncodeToMemory(b *Block) []byte
    EncodeToMemory() { hasQualifiedName("encoding/pem", "EncodeToMemory") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }
}
