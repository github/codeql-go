/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `crypto/x509` package. */
module CryptoX509TaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func DecryptPEMBlock(b *encoding/pem.Block, password []byte) ([]byte, error)
      hasQualifiedName("crypto/x509", "DecryptPEMBlock") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func EncryptPEMBlock(rand io.Reader, blockType string, data []byte, password []byte, alg PEMCipher) (*encoding/pem.Block, error)
      hasQualifiedName("crypto/x509", "EncryptPEMBlock") and
      (inp.isParameter(2) and outp.isResult(0))
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
      // signature: func (*CertPool).AddCert(cert *Certificate)
      this.(Method).hasQualifiedName("crypto/x509", "CertPool", "AddCert") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (*CertPool).AppendCertsFromPEM(pemCerts []byte) (ok bool)
      this.(Method).hasQualifiedName("crypto/x509", "CertPool", "AppendCertsFromPEM") and
      (inp.isParameter(0) and outp.isReceiver())
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
