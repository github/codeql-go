/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `crypto/ed25519` package. */
module CryptoEd25519TaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func Sign(privateKey PrivateKey, message []byte) []byte
      hasQualifiedName("crypto/ed25519", "Sign") and
      (inp.isParameter(1) and outp.isResult())
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
      // signature: func (PrivateKey).Sign(rand io.Reader, message []byte, opts crypto.SignerOpts) (signature []byte, err error)
      this.(Method).hasQualifiedName("crypto/ed25519", "PrivateKey", "Sign") and
      (inp.isParameter(1) and outp.isResult(0))
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
