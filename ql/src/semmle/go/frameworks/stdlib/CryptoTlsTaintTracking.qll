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

  private class Dial extends TaintTracking::FunctionModel {
    // signature: func Dial(network string, addr string, config *Config) (*Conn, error)
    Dial() { hasQualifiedName("crypto/tls", "Dial") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isResult(0)
    }
  }

  private class DialWithDialer extends TaintTracking::FunctionModel {
    // signature: func DialWithDialer(dialer *net.Dialer, network string, addr string, config *Config) (*Conn, error)
    DialWithDialer() { hasQualifiedName("crypto/tls", "DialWithDialer") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(2) and outp.isResult(0)
    }
  }

  private class NewListener extends TaintTracking::FunctionModel {
    // signature: func NewListener(inner net.Listener, config *Config) net.Listener
    NewListener() { hasQualifiedName("crypto/tls", "NewListener") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
      or
      inp.isResult() and outp.isParameter(0)
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
}
