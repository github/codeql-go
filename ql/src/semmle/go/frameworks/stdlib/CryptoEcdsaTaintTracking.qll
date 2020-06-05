/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `crypto/ecdsa` package. */
module CryptoEcdsaTaintTracking {
  private class PrivateKeySign extends TaintTracking::FunctionModel, Method {
    // signature: func (*PrivateKey).Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)
    PrivateKeySign() { this.(Method).hasQualifiedName("crypto/ecdsa", "PrivateKey", "Sign") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isResult(0)
    }
  }
}
