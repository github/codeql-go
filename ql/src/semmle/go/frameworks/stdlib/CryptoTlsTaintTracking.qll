/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `crypto/tls` package. */
module CryptoTlsTaintTracking {
  private class Client extends TaintTracking::FunctionModel {
    // signature: func Client(conn net.Conn, config *Config) *Conn
    Client() { hasQualifiedName("crypto/tls", "Client") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
      or
      inp.isResult() and outp.isParameter(0)
    }
  }

  private class NewListener extends TaintTracking::FunctionModel {
    // signature: func NewListener(inner net.Listener, config *Config) net.Listener
    NewListener() { hasQualifiedName("crypto/tls", "NewListener") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class Server extends TaintTracking::FunctionModel {
    // signature: func Server(conn net.Conn, config *Config) *Conn
    Server() { hasQualifiedName("crypto/tls", "Server") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
      or
      inp.isResult() and outp.isParameter(0)
    }
  }

  private class ConnRead extends TaintTracking::FunctionModel, Method {
    // signature: func (*Conn).Read(b []byte) (int, error)
    ConnRead() { this.(Method).hasQualifiedName("crypto/tls", "Conn", "Read") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class ConnWrite extends TaintTracking::FunctionModel, Method {
    // signature: func (*Conn).Write(b []byte) (int, error)
    ConnWrite() { this.(Method).hasQualifiedName("crypto/tls", "Conn", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }
}
