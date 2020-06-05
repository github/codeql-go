/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `net/textproto` package. */
module NetTextprotoTaintTracking {
  private class CanonicalMIMEHeaderKey extends TaintTracking::FunctionModel {
    // signature: func CanonicalMIMEHeaderKey(s string) string
    CanonicalMIMEHeaderKey() { hasQualifiedName("net/textproto", "CanonicalMIMEHeaderKey") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class NewReader extends TaintTracking::FunctionModel {
    // signature: func NewReader(r *bufio.Reader) *Reader
    NewReader() { hasQualifiedName("net/textproto", "NewReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class NewWriter extends TaintTracking::FunctionModel {
    // signature: func NewWriter(w *bufio.Writer) *Writer
    NewWriter() { hasQualifiedName("net/textproto", "NewWriter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult() and outp.isParameter(0)
    }
  }

  private class TrimBytes extends TaintTracking::FunctionModel {
    // signature: func TrimBytes(b []byte) []byte
    TrimBytes() { hasQualifiedName("net/textproto", "TrimBytes") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class TrimString extends TaintTracking::FunctionModel {
    // signature: func TrimString(s string) string
    TrimString() { hasQualifiedName("net/textproto", "TrimString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class MIMEHeaderAdd extends TaintTracking::FunctionModel, Method {
    // signature: func (MIMEHeader).Add(key string, value string)
    MIMEHeaderAdd() { this.(Method).hasQualifiedName("net/textproto", "MIMEHeader", "Add") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isReceiver()
    }
  }

  private class MIMEHeaderGet extends TaintTracking::FunctionModel, Method {
    // signature: func (MIMEHeader).Get(key string) string
    MIMEHeaderGet() { this.(Method).hasQualifiedName("net/textproto", "MIMEHeader", "Get") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class MIMEHeaderSet extends TaintTracking::FunctionModel, Method {
    // signature: func (MIMEHeader).Set(key string, value string)
    MIMEHeaderSet() { this.(Method).hasQualifiedName("net/textproto", "MIMEHeader", "Set") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isReceiver()
    }
  }

  private class MIMEHeaderValues extends TaintTracking::FunctionModel, Method {
    // signature: func (MIMEHeader).Values(key string) []string
    MIMEHeaderValues() { this.(Method).hasQualifiedName("net/textproto", "MIMEHeader", "Values") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ReaderDotReader extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).DotReader() io.Reader
    ReaderDotReader() { this.(Method).hasQualifiedName("net/textproto", "Reader", "DotReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ReaderReadCodeLine extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadCodeLine(expectCode int) (code int, message string, err error)
    ReaderReadCodeLine() {
      this.(Method).hasQualifiedName("net/textproto", "Reader", "ReadCodeLine")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(1)
    }
  }

  private class ReaderReadContinuedLine extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadContinuedLine() (string, error)
    ReaderReadContinuedLine() {
      this.(Method).hasQualifiedName("net/textproto", "Reader", "ReadContinuedLine")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class ReaderReadContinuedLineBytes extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadContinuedLineBytes() ([]byte, error)
    ReaderReadContinuedLineBytes() {
      this.(Method).hasQualifiedName("net/textproto", "Reader", "ReadContinuedLineBytes")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class ReaderReadDotBytes extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadDotBytes() ([]byte, error)
    ReaderReadDotBytes() {
      this.(Method).hasQualifiedName("net/textproto", "Reader", "ReadDotBytes")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class ReaderReadDotLines extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadDotLines() ([]string, error)
    ReaderReadDotLines() {
      this.(Method).hasQualifiedName("net/textproto", "Reader", "ReadDotLines")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class ReaderReadLine extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadLine() (string, error)
    ReaderReadLine() { this.(Method).hasQualifiedName("net/textproto", "Reader", "ReadLine") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class ReaderReadLineBytes extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadLineBytes() ([]byte, error)
    ReaderReadLineBytes() {
      this.(Method).hasQualifiedName("net/textproto", "Reader", "ReadLineBytes")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class ReaderReadMIMEHeader extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadMIMEHeader() (MIMEHeader, error)
    ReaderReadMIMEHeader() {
      this.(Method).hasQualifiedName("net/textproto", "Reader", "ReadMIMEHeader")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class ReaderReadResponse extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadResponse(expectCode int) (code int, message string, err error)
    ReaderReadResponse() {
      this.(Method).hasQualifiedName("net/textproto", "Reader", "ReadResponse")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(1)
    }
  }

  private class WriterDotWriter extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).DotWriter() io.WriteCloser
    WriterDotWriter() { this.(Method).hasQualifiedName("net/textproto", "Writer", "DotWriter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult() and outp.isReceiver()
    }
  }

  private class WriterPrintfLine extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).PrintfLine(format string, args ...interface{}) error
    WriterPrintfLine() { this.(Method).hasQualifiedName("net/textproto", "Writer", "PrintfLine") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isReceiver()
    }
  }
}
