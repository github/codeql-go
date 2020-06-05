/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `crypto/ed25519` package. */
module CryptoEd25519TaintTracking {
  private class Sign extends TaintTracking::FunctionModel {
    // signature: func Sign(privateKey PrivateKey, message []byte) []byte
    Sign() { hasQualifiedName("crypto/ed25519", "Sign") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isResult()
    }
  }

  private class PrivateKeySign extends TaintTracking::FunctionModel, Method {
    // signature: func (PrivateKey).Sign(rand io.Reader, message []byte, opts crypto.SignerOpts) (signature []byte, err error)
    PrivateKeySign() { this.(Method).hasQualifiedName("crypto/ed25519", "PrivateKey", "Sign") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isResult(0)
    }
  }
}
