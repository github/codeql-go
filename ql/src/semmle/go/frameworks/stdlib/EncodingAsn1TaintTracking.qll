/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `encoding/asn1` package. */
module EncodingAsn1TaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func Marshal(val interface{}) ([]byte, error)
      hasQualifiedName("encoding/asn1", "Marshal") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func MarshalWithParams(val interface{}, params string) ([]byte, error)
      hasQualifiedName("encoding/asn1", "MarshalWithParams") and
      (inp.isParameter(_) and outp.isResult(0))
      or
      // signature: func Unmarshal(b []byte, val interface{}) (rest []byte, err error)
      hasQualifiedName("encoding/asn1", "Unmarshal") and
      (
        inp.isParameter(0) and
        (outp.isParameter(1) or outp.isResult(0))
      )
      or
      // signature: func UnmarshalWithParams(b []byte, val interface{}, params string) (rest []byte, err error)
      hasQualifiedName("encoding/asn1", "UnmarshalWithParams") and
      (
        inp.isParameter([0, 2]) and
        (outp.isParameter(1) or outp.isResult(0))
      )
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
      // signature: func (ObjectIdentifier).String() string
      this.(Method).hasQualifiedName("encoding/asn1", "ObjectIdentifier", "String") and
      (inp.isReceiver() and outp.isResult())
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
