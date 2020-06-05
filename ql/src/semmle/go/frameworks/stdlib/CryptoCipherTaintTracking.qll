/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `crypto/cipher` package. */
module CryptoCipherTaintTracking {
  private class AEADOpen extends TaintTracking::FunctionModel, Method {
    // signature: func (AEAD).Open(dst []byte, nonce []byte, ciphertext []byte, additionalData []byte) ([]byte, error)
    AEADOpen() { this.implements("crypto/cipher", "AEAD", "Open") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(2) and
      (outp.isResult(0) or outp.isParameter(0))
    }
  }

  private class BlockDecrypt extends TaintTracking::FunctionModel, Method {
    // signature: func (Block).Decrypt(dst []byte, src []byte)
    BlockDecrypt() { this.implements("crypto/cipher", "Block", "Decrypt") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class BlockEncrypt extends TaintTracking::FunctionModel, Method {
    // signature: func (Block).Encrypt(dst []byte, src []byte)
    BlockEncrypt() { this.implements("crypto/cipher", "Block", "Encrypt") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class BlockModeCryptBlocks extends TaintTracking::FunctionModel, Method {
    // signature: func (BlockMode).CryptBlocks(dst []byte, src []byte)
    BlockModeCryptBlocks() { this.implements("crypto/cipher", "BlockMode", "CryptBlocks") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class StreamXORKeyStream extends TaintTracking::FunctionModel, Method {
    // signature: func (Stream).XORKeyStream(dst []byte, src []byte)
    StreamXORKeyStream() { this.implements("crypto/cipher", "Stream", "XORKeyStream") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }
}
