/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `crypto/cipher` package. */
module CryptoCipherTaintTracking {
  private class MethodAndInterfaceTaintTracking extends TaintTracking::FunctionModel, Method {
    FunctionInput inp;
    FunctionOutput outp;

    MethodAndInterfaceTaintTracking() {
      // Methods:
      // signature: func (StreamReader).Read(dst []byte) (n int, err error)
      this.(Method).hasQualifiedName("crypto/cipher", "StreamReader", "Read") and
      (inp.isReceiver() and outp.isParameter(0))
      or
      // signature: func (StreamWriter).Write(src []byte) (n int, err error)
      this.(Method).hasQualifiedName("crypto/cipher", "StreamWriter", "Write") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // Interfaces:
      // signature: func (BlockMode).CryptBlocks(dst []byte, src []byte)
      this.implements("crypto/cipher", "BlockMode", "CryptBlocks") and
      (inp.isParameter(1) and outp.isParameter(0))
      or
      // signature: func (Block).Decrypt(dst []byte, src []byte)
      this.implements("crypto/cipher", "Block", "Decrypt") and
      (inp.isParameter(1) and outp.isParameter(0))
      or
      // signature: func (Block).Encrypt(dst []byte, src []byte)
      this.implements("crypto/cipher", "Block", "Encrypt") and
      (inp.isParameter(1) and outp.isParameter(0))
      or
      // signature: func (AEAD).Open(dst []byte, nonce []byte, ciphertext []byte, additionalData []byte) ([]byte, error)
      this.implements("crypto/cipher", "AEAD", "Open") and
      (
        inp.isParameter(2) and
        (outp.isParameter(0) or outp.isResult(0))
      )
      or
      // signature: func (AEAD).Seal(dst []byte, nonce []byte, plaintext []byte, additionalData []byte) []byte
      this.implements("crypto/cipher", "AEAD", "Seal") and
      (
        inp.isParameter(2) and
        (outp.isParameter(0) or outp.isResult())
      )
      or
      // signature: func (Stream).XORKeyStream(dst []byte, src []byte)
      this.implements("crypto/cipher", "Stream", "XORKeyStream") and
      (inp.isParameter(1) and outp.isParameter(0))
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
