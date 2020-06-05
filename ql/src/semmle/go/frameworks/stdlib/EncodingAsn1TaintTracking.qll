/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `encoding/asn1` package. */
module EncodingAsn1TaintTracking {
  private class Marshal extends TaintTracking::FunctionModel {
    // signature: func Marshal(val interface{}) ([]byte, error)
    Marshal() { hasQualifiedName("encoding/asn1", "Marshal") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class MarshalWithParams extends TaintTracking::FunctionModel {
    // signature: func MarshalWithParams(val interface{}, params string) ([]byte, error)
    MarshalWithParams() { hasQualifiedName("encoding/asn1", "MarshalWithParams") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class Unmarshal extends TaintTracking::FunctionModel {
    // signature: func Unmarshal(b []byte, val interface{}) (rest []byte, err error)
    Unmarshal() { hasQualifiedName("encoding/asn1", "Unmarshal") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isParameter(1)
    }
  }

  private class UnmarshalWithParams extends TaintTracking::FunctionModel {
    // signature: func UnmarshalWithParams(b []byte, val interface{}, params string) (rest []byte, err error)
    UnmarshalWithParams() { hasQualifiedName("encoding/asn1", "UnmarshalWithParams") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isParameter(1)
    }
  }

  private class ObjectIdentifierString extends TaintTracking::FunctionModel, Method {
    // signature: func (ObjectIdentifier).String() string
    ObjectIdentifierString() {
      this.(Method).hasQualifiedName("encoding/asn1", "ObjectIdentifier", "String")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }
}
