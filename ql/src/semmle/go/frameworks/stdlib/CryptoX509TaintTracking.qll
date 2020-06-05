/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `crypto/x509` package. */
module CryptoX509TaintTracking {
  private class DecryptPEMBlock extends TaintTracking::FunctionModel {
    // signature: func DecryptPEMBlock(b *encoding/pem.Block, password []byte) ([]byte, error)
    DecryptPEMBlock() { hasQualifiedName("crypto/x509", "DecryptPEMBlock") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class EncryptPEMBlock extends TaintTracking::FunctionModel {
    // signature: func EncryptPEMBlock(rand io.Reader, blockType string, data []byte, password []byte, alg PEMCipher) (*encoding/pem.Block, error)
    EncryptPEMBlock() { hasQualifiedName("crypto/x509", "EncryptPEMBlock") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(2) and outp.isResult(0)
    }
  }
}
