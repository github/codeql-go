/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `crypto` package. */
module CryptoTaintTracking {
  private class MethodAndInterfaceTaintTracking extends TaintTracking::FunctionModel, Method {
    FunctionInput inp;
    FunctionOutput outp;

    MethodAndInterfaceTaintTracking() {
      // Methods:
      // Interfaces:
      // signature: func (Decrypter).Decrypt(rand io.Reader, msg []byte, opts DecrypterOpts) (plaintext []byte, err error)
      this.implements("crypto", "Decrypter", "Decrypt") and
      (inp.isParameter(1) and outp.isResult(0))
      or
      // signature: func (Signer).Sign(rand io.Reader, digest []byte, opts SignerOpts) (signature []byte, err error)
      this.implements("crypto", "Signer", "Sign") and
      (inp.isParameter(1) and outp.isResult(0))
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
