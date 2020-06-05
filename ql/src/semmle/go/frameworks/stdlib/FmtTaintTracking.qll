/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `fmt` package. */
module FmtTaintTracking {
  private class Errorf extends TaintTracking::FunctionModel {
    // signature: func Errorf(format string, a ...interface{}) error
    Errorf() { hasQualifiedName("fmt", "Errorf") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isResult()
    }
  }

  private class Fprint extends TaintTracking::FunctionModel {
    // signature: func Fprint(w io.Writer, a ...interface{}) (n int, err error)
    Fprint() { hasQualifiedName("fmt", "Fprint") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(any(int i | i >= 1)) and outp.isParameter(0)
    }
  }

  private class Fprintf extends TaintTracking::FunctionModel {
    // signature: func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
    Fprintf() { hasQualifiedName("fmt", "Fprintf") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(any(int i | i >= 1)) and outp.isParameter(0)
    }
  }

  private class Fprintln extends TaintTracking::FunctionModel {
    // signature: func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
    Fprintln() { hasQualifiedName("fmt", "Fprintln") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(any(int i | i >= 1)) and outp.isParameter(0)
    }
  }

  private class Fscan extends TaintTracking::FunctionModel {
    // signature: func Fscan(r io.Reader, a ...interface{}) (n int, err error)
    Fscan() { hasQualifiedName("fmt", "Fscan") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isParameter(any(int i | i >= 1))
    }
  }

  private class Fscanf extends TaintTracking::FunctionModel {
    // signature: func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
    Fscanf() { hasQualifiedName("fmt", "Fscanf") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isParameter(any(int i | i >= 2))
    }
  }

  private class Fscanln extends TaintTracking::FunctionModel {
    // signature: func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
    Fscanln() { hasQualifiedName("fmt", "Fscanln") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isParameter(any(int i | i >= 1))
    }
  }

  private class Sprint extends TaintTracking::FunctionModel {
    // signature: func Sprint(a ...interface{}) string
    Sprint() { hasQualifiedName("fmt", "Sprint") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isResult()
    }
  }

  private class Sprintf extends TaintTracking::FunctionModel {
    // signature: func Sprintf(format string, a ...interface{}) string
    Sprintf() { hasQualifiedName("fmt", "Sprintf") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isResult()
    }
  }

  private class Sprintln extends TaintTracking::FunctionModel {
    // signature: func Sprintln(a ...interface{}) string
    Sprintln() { hasQualifiedName("fmt", "Sprintln") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isResult()
    }
  }

  private class Sscan extends TaintTracking::FunctionModel {
    // signature: func Sscan(str string, a ...interface{}) (n int, err error)
    Sscan() { hasQualifiedName("fmt", "Sscan") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isParameter(any(int i | i >= 1))
    }
  }

  private class Sscanf extends TaintTracking::FunctionModel {
    // signature: func Sscanf(str string, format string, a ...interface{}) (n int, err error)
    Sscanf() { hasQualifiedName("fmt", "Sscanf") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isParameter(any(int i | i >= 2))
    }
  }

  private class Sscanln extends TaintTracking::FunctionModel {
    // signature: func Sscanln(str string, a ...interface{}) (n int, err error)
    Sscanln() { hasQualifiedName("fmt", "Sscanln") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isParameter(any(int i | i >= 1))
    }
  }

  private class GoStringerGoString extends TaintTracking::FunctionModel, Method {
    // signature: func (GoStringer).GoString() string
    GoStringerGoString() { this.implements("fmt", "GoStringer", "GoString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ScanStateReadRune extends TaintTracking::FunctionModel, Method {
    // signature: func (ScanState).ReadRune() (r rune, size int, err error)
    ScanStateReadRune() { this.implements("fmt", "ScanState", "ReadRune") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class ScanStateToken extends TaintTracking::FunctionModel, Method {
    // signature: func (ScanState).Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
    ScanStateToken() { this.implements("fmt", "ScanState", "Token") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class StringerString extends TaintTracking::FunctionModel, Method {
    // signature: func (Stringer).String() string
    StringerString() { this.implements("fmt", "Stringer", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }
}
