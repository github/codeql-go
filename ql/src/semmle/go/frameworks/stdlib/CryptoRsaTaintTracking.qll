/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `crypto/rsa` package. */
module CryptoRsaTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func DecryptOAEP(hash hash.Hash, random io.Reader, priv *PrivateKey, ciphertext []byte, label []byte) ([]byte, error)
      hasQualifiedName("crypto/rsa", "DecryptOAEP") and
      (inp.isParameter(3) and outp.isResult(0))
      or
      // signature: func DecryptPKCS1v15(rand io.Reader, priv *PrivateKey, ciphertext []byte) ([]byte, error)
      hasQualifiedName("crypto/rsa", "DecryptPKCS1v15") and
      (inp.isParameter(2) and outp.isResult(0))
      or
      // signature: func EncryptOAEP(hash hash.Hash, random io.Reader, pub *PublicKey, msg []byte, label []byte) ([]byte, error)
      hasQualifiedName("crypto/rsa", "EncryptOAEP") and
      (inp.isParameter(3) and outp.isResult(0))
      or
      // signature: func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) ([]byte, error)
      hasQualifiedName("crypto/rsa", "EncryptPKCS1v15") and
      (inp.isParameter(2) and outp.isResult(0))
      or
      // signature: func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error)
      hasQualifiedName("crypto/rsa", "SignPKCS1v15") and
      (inp.isParameter(3) and outp.isResult(0))
      or
      // signature: func SignPSS(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte, opts *PSSOptions) ([]byte, error)
      hasQualifiedName("crypto/rsa", "SignPSS") and
      (inp.isParameter(3) and outp.isResult(0))
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
      // signature: func (*PrivateKey).Decrypt(rand io.Reader, ciphertext []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error)
      this.(Method).hasQualifiedName("crypto/rsa", "PrivateKey", "Decrypt") and
      (inp.isParameter(1) and outp.isResult(0))
      or
      // signature: func (*PrivateKey).Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)
      this.(Method).hasQualifiedName("crypto/rsa", "PrivateKey", "Sign") and
      (inp.isParameter(1) and outp.isResult(0))
      // Interfaces:
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
