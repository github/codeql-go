/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `crypto/ecdsa` package. */
module CryptoEcdsaTaintTracking {
  private class MethodAndInterfaceTaintTracking extends TaintTracking::FunctionModel, Method {
    FunctionInput inp;
    FunctionOutput outp;

    MethodAndInterfaceTaintTracking() {
      // Methods:
      // signature: func (*PrivateKey).Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)
      this.(Method).hasQualifiedName("crypto/ecdsa", "PrivateKey", "Sign") and
      (inp.isParameter(1) and outp.isResult(0))
      // Interfaces:
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
