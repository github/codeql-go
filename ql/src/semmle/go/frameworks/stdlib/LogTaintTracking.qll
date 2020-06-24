/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `log` package. */
module LogTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func New(out io.Writer, prefix string, flag int) *Logger
      hasQualifiedName("log", "New") and
      (inp.isResult() and outp.isParameter(0))
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
      // signature: func (*Logger).Fatal(v ...interface{})
      this.(Method).hasQualifiedName("log", "Logger", "Fatal") and
      (inp.isParameter(_) and outp.isReceiver())
      or
      // signature: func (*Logger).Fatalf(format string, v ...interface{})
      this.(Method).hasQualifiedName("log", "Logger", "Fatalf") and
      (inp.isParameter(_) and outp.isReceiver())
      or
      // signature: func (*Logger).Fatalln(v ...interface{})
      this.(Method).hasQualifiedName("log", "Logger", "Fatalln") and
      (inp.isParameter(_) and outp.isReceiver())
      or
      // signature: func (*Logger).Panic(v ...interface{})
      this.(Method).hasQualifiedName("log", "Logger", "Panic") and
      (inp.isParameter(_) and outp.isReceiver())
      or
      // signature: func (*Logger).Panicf(format string, v ...interface{})
      this.(Method).hasQualifiedName("log", "Logger", "Panicf") and
      (inp.isParameter(_) and outp.isReceiver())
      or
      // signature: func (*Logger).Panicln(v ...interface{})
      this.(Method).hasQualifiedName("log", "Logger", "Panicln") and
      (inp.isParameter(_) and outp.isReceiver())
      or
      // signature: func (*Logger).Print(v ...interface{})
      this.(Method).hasQualifiedName("log", "Logger", "Print") and
      (inp.isParameter(_) and outp.isReceiver())
      or
      // signature: func (*Logger).Printf(format string, v ...interface{})
      this.(Method).hasQualifiedName("log", "Logger", "Printf") and
      (inp.isParameter(_) and outp.isReceiver())
      or
      // signature: func (*Logger).Println(v ...interface{})
      this.(Method).hasQualifiedName("log", "Logger", "Println") and
      (inp.isParameter(_) and outp.isReceiver())
      or
      // signature: func (*Logger).SetOutput(w io.Writer)
      this.(Method).hasQualifiedName("log", "Logger", "SetOutput") and
      (inp.isReceiver() and outp.isParameter(0))
      or
      // signature: func (*Logger).SetPrefix(prefix string)
      this.(Method).hasQualifiedName("log", "Logger", "SetPrefix") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (*Logger).Writer() io.Writer
      this.(Method).hasQualifiedName("log", "Logger", "Writer") and
      (inp.isResult() and outp.isReceiver())
      or
      inp.isReceiver() and outp.isResult()
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
