/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `mime` package. */
module MimeTaintTracking {
  private class FormatMediaType extends TaintTracking::FunctionModel {
    // signature: func FormatMediaType(t string, param map[string]string) string
    FormatMediaType() { hasQualifiedName("mime", "FormatMediaType") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class ParseMediaType extends TaintTracking::FunctionModel {
    // signature: func ParseMediaType(v string) (mediatype string, params map[string]string, err error)
    ParseMediaType() { hasQualifiedName("mime", "ParseMediaType") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult([0, 1]))
    }
  }

  private class WordDecoderDecode extends TaintTracking::FunctionModel, Method {
    // signature: func (*WordDecoder).Decode(word string) (string, error)
    WordDecoderDecode() { this.(Method).hasQualifiedName("mime", "WordDecoder", "Decode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class WordDecoderDecodeHeader extends TaintTracking::FunctionModel, Method {
    // signature: func (*WordDecoder).DecodeHeader(header string) (string, error)
    WordDecoderDecodeHeader() {
      this.(Method).hasQualifiedName("mime", "WordDecoder", "DecodeHeader")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class WordEncoderEncode extends TaintTracking::FunctionModel, Method {
    // signature: func (WordEncoder).Encode(charset string, s string) string
    WordEncoderEncode() { this.(Method).hasQualifiedName("mime", "WordEncoder", "Encode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isResult())
    }
  }
}
