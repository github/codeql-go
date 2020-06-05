/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `io` package. */
module IoTaintTracking {
  private class Copy extends TaintTracking::FunctionModel {
    // signature: func Copy(dst Writer, src Reader) (written int64, err error)
    Copy() { hasQualifiedName("io", "Copy") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class CopyBuffer extends TaintTracking::FunctionModel {
    // signature: func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
    CopyBuffer() { hasQualifiedName("io", "CopyBuffer") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class CopyN extends TaintTracking::FunctionModel {
    // signature: func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
    CopyN() { hasQualifiedName("io", "CopyN") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class LimitReader extends TaintTracking::FunctionModel {
    // signature: func LimitReader(r Reader, n int64) Reader
    LimitReader() { hasQualifiedName("io", "LimitReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class MultiReader extends TaintTracking::FunctionModel {
    // signature: func MultiReader(readers ...Reader) Reader
    MultiReader() { hasQualifiedName("io", "MultiReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isResult()
    }
  }

  private class MultiWriter extends TaintTracking::FunctionModel {
    // signature: func MultiWriter(writers ...Writer) Writer
    MultiWriter() { hasQualifiedName("io", "MultiWriter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult() and outp.isParameter(_)
    }
  }

  private class NewSectionReader extends TaintTracking::FunctionModel {
    // signature: func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader
    NewSectionReader() { hasQualifiedName("io", "NewSectionReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class Pipe extends TaintTracking::FunctionModel {
    // signature: func Pipe() (*PipeReader, *PipeWriter)
    Pipe() { hasQualifiedName("io", "Pipe") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult(1) and outp.isResult(0)
    }
  }

  private class ReadAtLeast extends TaintTracking::FunctionModel {
    // signature: func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
    ReadAtLeast() { hasQualifiedName("io", "ReadAtLeast") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isParameter(1)
    }
  }

  private class ReadFull extends TaintTracking::FunctionModel {
    // signature: func ReadFull(r Reader, buf []byte) (n int, err error)
    ReadFull() { hasQualifiedName("io", "ReadFull") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isParameter(1)
    }
  }

  private class TeeReader extends TaintTracking::FunctionModel {
    // signature: func TeeReader(r Reader, w Writer) Reader
    TeeReader() { hasQualifiedName("io", "TeeReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and
      (outp.isParameter(1) or outp.isResult())
    }
  }

  private class WriteString extends TaintTracking::FunctionModel {
    // signature: func WriteString(w Writer, s string) (n int, err error)
    WriteString() { hasQualifiedName("io", "WriteString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class ByteReaderReadByte extends TaintTracking::FunctionModel, Method {
    // signature: func (ByteReader).ReadByte() (byte, error)
    ByteReaderReadByte() { this.implements("io", "ByteReader", "ReadByte") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class ByteWriterWriteByte extends TaintTracking::FunctionModel, Method {
    // signature: func (ByteWriter).WriteByte(c byte) error
    ByteWriterWriteByte() { this.implements("io", "ByteWriter", "WriteByte") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ReaderFromReadFrom extends TaintTracking::FunctionModel, Method {
    // signature: func (ReaderFrom).ReadFrom(r Reader) (n int64, err error)
    ReaderFromReadFrom() { this.implements("io", "ReaderFrom", "ReadFrom") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class RuneReaderReadRune extends TaintTracking::FunctionModel, Method {
    // signature: func (RuneReader).ReadRune() (r rune, size int, err error)
    RuneReaderReadRune() { this.implements("io", "RuneReader", "ReadRune") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class StringWriterWriteString extends TaintTracking::FunctionModel, Method {
    // signature: func (StringWriter).WriteString(s string) (n int, err error)
    StringWriterWriteString() { this.implements("io", "StringWriter", "WriteString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class WriterAtWriteAt extends TaintTracking::FunctionModel, Method {
    // signature: func (WriterAt).WriteAt(p []byte, off int64) (n int, err error)
    WriterAtWriteAt() { this.implements("io", "WriterAt", "WriteAt") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class WriterToWriteTo extends TaintTracking::FunctionModel, Method {
    // signature: func (WriterTo).WriteTo(w Writer) (n int64, err error)
    WriterToWriteTo() { this.implements("io", "WriterTo", "WriteTo") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isParameter(0)
    }
  }
}
