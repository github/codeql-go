/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `fmt` package. */
module FmtTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func Errorf(format string, a ...interface{}) error
      hasQualifiedName("fmt", "Errorf") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func Fprint(w io.Writer, a ...interface{}) (n int, err error)
      hasQualifiedName("fmt", "Fprint") and
      (inp.isParameter(any(int i | i >= 1)) and outp.isParameter(0))
      or
      // signature: func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
      hasQualifiedName("fmt", "Fprintf") and
      (inp.isParameter([1, any(int i | i >= 2)]) and outp.isParameter(0))
      or
      // signature: func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
      hasQualifiedName("fmt", "Fprintln") and
      (inp.isParameter(any(int i | i >= 1)) and outp.isParameter(0))
      or
      // signature: func Fscan(r io.Reader, a ...interface{}) (n int, err error)
      hasQualifiedName("fmt", "Fscan") and
      (inp.isParameter(0) and outp.isParameter(any(int i | i >= 1)))
      or
      // signature: func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
      hasQualifiedName("fmt", "Fscanf") and
      (inp.isParameter([0, 1]) and outp.isParameter(any(int i | i >= 2)))
      or
      // signature: func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
      hasQualifiedName("fmt", "Fscanln") and
      (inp.isParameter(0) and outp.isParameter(any(int i | i >= 1)))
      or
      // signature: func Sprint(a ...interface{}) string
      hasQualifiedName("fmt", "Sprint") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func Sprintf(format string, a ...interface{}) string
      hasQualifiedName("fmt", "Sprintf") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func Sprintln(a ...interface{}) string
      hasQualifiedName("fmt", "Sprintln") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func Sscan(str string, a ...interface{}) (n int, err error)
      hasQualifiedName("fmt", "Sscan") and
      (inp.isParameter(0) and outp.isParameter(any(int i | i >= 1)))
      or
      // signature: func Sscanf(str string, format string, a ...interface{}) (n int, err error)
      hasQualifiedName("fmt", "Sscanf") and
      (inp.isParameter([0, 1]) and outp.isParameter(any(int i | i >= 2)))
      or
      // signature: func Sscanln(str string, a ...interface{}) (n int, err error)
      hasQualifiedName("fmt", "Sscanln") and
      (inp.isParameter(0) and outp.isParameter(any(int i | i >= 1)))
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
      // Interfaces:
      // signature: func (GoStringer).GoString() string
      this.implements("fmt", "GoStringer", "GoString") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (ScanState).Read(buf []byte) (n int, err error)
      this.implements("fmt", "ScanState", "Read") and
      (inp.isReceiver() and outp.isParameter(0))
      or
      // signature: func (ScanState).ReadRune() (r rune, size int, err error)
      this.implements("fmt", "ScanState", "ReadRune") and
      (inp.isReceiver() and outp.isResult(0))
      or
      // signature: func (Stringer).String() string
      this.implements("fmt", "Stringer", "String") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (ScanState).Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
      this.implements("fmt", "ScanState", "Token") and
      (inp.isReceiver() and outp.isResult(0))
      or
      // signature: func (State).Write(b []byte) (n int, err error)
      this.implements("fmt", "State", "Write") and
      (inp.isParameter(0) and outp.isReceiver())
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
