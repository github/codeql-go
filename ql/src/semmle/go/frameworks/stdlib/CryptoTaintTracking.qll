/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `crypto` package. */
module CryptoTaintTracking {
  private class HashHashFunc extends TaintTracking::FunctionModel, Method {
    // signature: func (Hash).HashFunc() Hash
    HashHashFunc() { this.(Method).hasQualifiedName("crypto", "Hash", "HashFunc") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class HashNew extends TaintTracking::FunctionModel, Method {
    // signature: func (Hash).New() hash.Hash
    HashNew() { this.(Method).hasQualifiedName("crypto", "Hash", "New") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class DecrypterDecrypt extends TaintTracking::FunctionModel, Method {
    // signature: func (Decrypter).Decrypt(rand io.Reader, msg []byte, opts DecrypterOpts) (plaintext []byte, err error)
    DecrypterDecrypt() { this.implements("crypto", "Decrypter", "Decrypt") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isResult(0)
    }
  }

  private class SignerPublic extends TaintTracking::FunctionModel, Method {
    // signature: func (Signer).Public() PublicKey
    SignerPublic() { this.implements("crypto", "Signer", "Public") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class SignerSign extends TaintTracking::FunctionModel, Method {
    // signature: func (Signer).Sign(rand io.Reader, digest []byte, opts SignerOpts) (signature []byte, err error)
    SignerSign() { this.implements("crypto", "Signer", "Sign") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isResult(0)
    }
  }
}
