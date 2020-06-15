/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `bufio` package. */
module BufioTaintTracking {
  private class NewReadWriter extends TaintTracking::FunctionModel {
    // signature: func NewReadWriter(r *Reader, w *Writer) *ReadWriter
    NewReadWriter() { hasQualifiedName("bufio", "NewReadWriter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
      or
      inp.isResult() and outp.isParameter(1)
    }
  }

  private class NewReader extends TaintTracking::FunctionModel {
    // signature: func NewReader(rd io.Reader) *Reader
    NewReader() { hasQualifiedName("bufio", "NewReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class NewReaderSize extends TaintTracking::FunctionModel {
    // signature: func NewReaderSize(rd io.Reader, size int) *Reader
    NewReaderSize() { hasQualifiedName("bufio", "NewReaderSize") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class NewScanner extends TaintTracking::FunctionModel {
    // signature: func NewScanner(r io.Reader) *Scanner
    NewScanner() { hasQualifiedName("bufio", "NewScanner") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class NewWriter extends TaintTracking::FunctionModel {
    // signature: func NewWriter(w io.Writer) *Writer
    NewWriter() { hasQualifiedName("bufio", "NewWriter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isResult() and outp.isParameter(0))
    }
  }

  private class NewWriterSize extends TaintTracking::FunctionModel {
    // signature: func NewWriterSize(w io.Writer, size int) *Writer
    NewWriterSize() { hasQualifiedName("bufio", "NewWriterSize") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isResult() and outp.isParameter(0))
    }
  }

  private class ScanBytes extends TaintTracking::FunctionModel {
    // signature: func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
    ScanBytes() { hasQualifiedName("bufio", "ScanBytes") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(1))
    }
  }

  private class ScanLines extends TaintTracking::FunctionModel {
    // signature: func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)
    ScanLines() { hasQualifiedName("bufio", "ScanLines") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(1))
    }
  }

  private class ScanRunes extends TaintTracking::FunctionModel {
    // signature: func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)
    ScanRunes() { hasQualifiedName("bufio", "ScanRunes") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(1))
    }
  }

  private class ScanWords extends TaintTracking::FunctionModel {
    // signature: func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
    ScanWords() { hasQualifiedName("bufio", "ScanWords") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(1))
    }
  }

  private class ReaderPeek extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).Peek(n int) ([]byte, error)
    ReaderPeek() { this.(Method).hasQualifiedName("bufio", "Reader", "Peek") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }

  private class ReaderRead extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).Read(p []byte) (n int, err error)
    ReaderRead() { this.(Method).hasQualifiedName("bufio", "Reader", "Read") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class ReaderReadByte extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadByte() (byte, error)
    ReaderReadByte() { this.(Method).hasQualifiedName("bufio", "Reader", "ReadByte") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }

  private class ReaderReadBytes extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadBytes(delim byte) ([]byte, error)
    ReaderReadBytes() { this.(Method).hasQualifiedName("bufio", "Reader", "ReadBytes") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }

  private class ReaderReadLine extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadLine() (line []byte, isPrefix bool, err error)
    ReaderReadLine() { this.(Method).hasQualifiedName("bufio", "Reader", "ReadLine") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }

  private class ReaderReadRune extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadRune() (r rune, size int, err error)
    ReaderReadRune() { this.(Method).hasQualifiedName("bufio", "Reader", "ReadRune") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }

  private class ReaderReadSlice extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadSlice(delim byte) (line []byte, err error)
    ReaderReadSlice() { this.(Method).hasQualifiedName("bufio", "Reader", "ReadSlice") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }

  private class ReaderReadString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadString(delim byte) (string, error)
    ReaderReadString() { this.(Method).hasQualifiedName("bufio", "Reader", "ReadString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }

  private class ReaderReset extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).Reset(r io.Reader)
    ReaderReset() { this.(Method).hasQualifiedName("bufio", "Reader", "Reset") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class ReaderWriteTo extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).WriteTo(w io.Writer) (n int64, err error)
    ReaderWriteTo() { this.(Method).hasQualifiedName("bufio", "Reader", "WriteTo") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class ScannerBytes extends TaintTracking::FunctionModel, Method {
    // signature: func (*Scanner).Bytes() []byte
    ScannerBytes() { this.(Method).hasQualifiedName("bufio", "Scanner", "Bytes") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class ScannerText extends TaintTracking::FunctionModel, Method {
    // signature: func (*Scanner).Text() string
    ScannerText() { this.(Method).hasQualifiedName("bufio", "Scanner", "Text") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class WriterReadFrom extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).ReadFrom(r io.Reader) (n int64, err error)
    WriterReadFrom() { this.(Method).hasQualifiedName("bufio", "Writer", "ReadFrom") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class WriterReset extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).Reset(w io.Writer)
    WriterReset() { this.(Method).hasQualifiedName("bufio", "Writer", "Reset") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class WriterWrite extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).Write(p []byte) (nn int, err error)
    WriterWrite() { this.(Method).hasQualifiedName("bufio", "Writer", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class WriterWriteByte extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).WriteByte(c byte) error
    WriterWriteByte() { this.(Method).hasQualifiedName("bufio", "Writer", "WriteByte") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class WriterWriteRune extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).WriteRune(r rune) (size int, err error)
    WriterWriteRune() { this.(Method).hasQualifiedName("bufio", "Writer", "WriteRune") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class WriterWriteString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).WriteString(s string) (int, error)
    WriterWriteString() { this.(Method).hasQualifiedName("bufio", "Writer", "WriteString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }
}
