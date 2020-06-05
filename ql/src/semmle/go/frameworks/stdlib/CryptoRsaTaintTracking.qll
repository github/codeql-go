/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `crypto/rsa` package. */
module CryptoRsaTaintTracking {
  private class DecryptOAEP extends TaintTracking::FunctionModel {
    // signature: func DecryptOAEP(hash hash.Hash, random io.Reader, priv *PrivateKey, ciphertext []byte, label []byte) ([]byte, error)
    DecryptOAEP() { hasQualifiedName("crypto/rsa", "DecryptOAEP") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(3) and outp.isResult(0)
    }
  }

  private class DecryptPKCS1V15 extends TaintTracking::FunctionModel {
    // signature: func DecryptPKCS1v15(rand io.Reader, priv *PrivateKey, ciphertext []byte) ([]byte, error)
    DecryptPKCS1V15() { hasQualifiedName("crypto/rsa", "DecryptPKCS1v15") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(2) and outp.isResult(0)
    }
  }

  private class DecryptPKCS1V15SessionKey extends TaintTracking::FunctionModel {
    // signature: func DecryptPKCS1v15SessionKey(rand io.Reader, priv *PrivateKey, ciphertext []byte, key []byte) error
    DecryptPKCS1V15SessionKey() { hasQualifiedName("crypto/rsa", "DecryptPKCS1v15SessionKey") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(2) and outp.isParameter(3)
    }
  }

  private class EncryptOAEP extends TaintTracking::FunctionModel {
    // signature: func EncryptOAEP(hash hash.Hash, random io.Reader, pub *PublicKey, msg []byte, label []byte) ([]byte, error)
    EncryptOAEP() { hasQualifiedName("crypto/rsa", "EncryptOAEP") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(3) and outp.isResult(0)
    }
  }

  private class EncryptPKCS1V15 extends TaintTracking::FunctionModel {
    // signature: func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) ([]byte, error)
    EncryptPKCS1V15() { hasQualifiedName("crypto/rsa", "EncryptPKCS1v15") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(2) and outp.isResult(0)
    }
  }

  private class SignPKCS1V15 extends TaintTracking::FunctionModel {
    // signature: func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error)
    SignPKCS1V15() { hasQualifiedName("crypto/rsa", "SignPKCS1v15") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(3) and outp.isResult(0)
    }
  }

  private class PrivateKeyDecrypt extends TaintTracking::FunctionModel, Method {
    // signature: func (*PrivateKey).Decrypt(rand io.Reader, ciphertext []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error)
    PrivateKeyDecrypt() { this.(Method).hasQualifiedName("crypto/rsa", "PrivateKey", "Decrypt") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isResult(0)
    }
  }

  private class PrivateKeySign extends TaintTracking::FunctionModel, Method {
    // signature: func (*PrivateKey).Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)
    PrivateKeySign() { this.(Method).hasQualifiedName("crypto/rsa", "PrivateKey", "Sign") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isResult(0)
    }
  }
}
