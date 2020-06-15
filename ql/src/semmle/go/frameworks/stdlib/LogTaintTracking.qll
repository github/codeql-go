/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `log` package. */
module LogTaintTracking {
  private class New extends TaintTracking::FunctionModel {
    // signature: func New(out io.Writer, prefix string, flag int) *Logger
    New() { hasQualifiedName("log", "New") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isResult() and outp.isParameter(0))
    }
  }

  private class LoggerFatal extends TaintTracking::FunctionModel, Method {
    // signature: func (*Logger).Fatal(v ...interface{})
    LoggerFatal() { this.(Method).hasQualifiedName("log", "Logger", "Fatal") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isReceiver())
    }
  }

  private class LoggerFatalf extends TaintTracking::FunctionModel, Method {
    // signature: func (*Logger).Fatalf(format string, v ...interface{})
    LoggerFatalf() { this.(Method).hasQualifiedName("log", "Logger", "Fatalf") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isReceiver())
    }
  }

  private class LoggerFatalln extends TaintTracking::FunctionModel, Method {
    // signature: func (*Logger).Fatalln(v ...interface{})
    LoggerFatalln() { this.(Method).hasQualifiedName("log", "Logger", "Fatalln") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isReceiver())
    }
  }

  private class LoggerPanic extends TaintTracking::FunctionModel, Method {
    // signature: func (*Logger).Panic(v ...interface{})
    LoggerPanic() { this.(Method).hasQualifiedName("log", "Logger", "Panic") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isReceiver())
    }
  }

  private class LoggerPanicf extends TaintTracking::FunctionModel, Method {
    // signature: func (*Logger).Panicf(format string, v ...interface{})
    LoggerPanicf() { this.(Method).hasQualifiedName("log", "Logger", "Panicf") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isReceiver())
    }
  }

  private class LoggerPanicln extends TaintTracking::FunctionModel, Method {
    // signature: func (*Logger).Panicln(v ...interface{})
    LoggerPanicln() { this.(Method).hasQualifiedName("log", "Logger", "Panicln") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isReceiver())
    }
  }

  private class LoggerPrint extends TaintTracking::FunctionModel, Method {
    // signature: func (*Logger).Print(v ...interface{})
    LoggerPrint() { this.(Method).hasQualifiedName("log", "Logger", "Print") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isReceiver())
    }
  }

  private class LoggerPrintf extends TaintTracking::FunctionModel, Method {
    // signature: func (*Logger).Printf(format string, v ...interface{})
    LoggerPrintf() { this.(Method).hasQualifiedName("log", "Logger", "Printf") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isReceiver())
    }
  }

  private class LoggerPrintln extends TaintTracking::FunctionModel, Method {
    // signature: func (*Logger).Println(v ...interface{})
    LoggerPrintln() { this.(Method).hasQualifiedName("log", "Logger", "Println") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isReceiver())
    }
  }

  private class LoggerSetOutput extends TaintTracking::FunctionModel, Method {
    // signature: func (*Logger).SetOutput(w io.Writer)
    LoggerSetOutput() { this.(Method).hasQualifiedName("log", "Logger", "SetOutput") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class LoggerSetPrefix extends TaintTracking::FunctionModel, Method {
    // signature: func (*Logger).SetPrefix(prefix string)
    LoggerSetPrefix() { this.(Method).hasQualifiedName("log", "Logger", "SetPrefix") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class LoggerWriter extends TaintTracking::FunctionModel, Method {
    // signature: func (*Logger).Writer() io.Writer
    LoggerWriter() { this.(Method).hasQualifiedName("log", "Logger", "Writer") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult() and outp.isReceiver()
      or
      inp.isReceiver() and outp.isResult()
    }
  }
}
