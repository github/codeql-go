/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `archive/tar` package. */
module ArchiveTarTaintTracking {
  private class FileInfoHeader extends TaintTracking::FunctionModel {
    // signature: func FileInfoHeader(fi os.FileInfo, link string) (*Header, error)
    FileInfoHeader() { hasQualifiedName("archive/tar", "FileInfoHeader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class NewReader extends TaintTracking::FunctionModel {
    // signature: func NewReader(r io.Reader) *Reader
    NewReader() { hasQualifiedName("archive/tar", "NewReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class NewWriter extends TaintTracking::FunctionModel {
    // signature: func NewWriter(w io.Writer) *Writer
    NewWriter() { hasQualifiedName("archive/tar", "NewWriter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isResult() and outp.isParameter(0))
    }
  }

  private class HeaderFileInfo extends TaintTracking::FunctionModel, Method {
    // signature: func (*Header).FileInfo() os.FileInfo
    HeaderFileInfo() { this.(Method).hasQualifiedName("archive/tar", "Header", "FileInfo") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class ReaderNext extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).Next() (*Header, error)
    ReaderNext() { this.(Method).hasQualifiedName("archive/tar", "Reader", "Next") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }

  private class ReaderRead extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).Read(b []byte) (int, error)
    ReaderRead() { this.(Method).hasQualifiedName("archive/tar", "Reader", "Read") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class WriterWrite extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).Write(b []byte) (int, error)
    WriterWrite() { this.(Method).hasQualifiedName("archive/tar", "Writer", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class WriterWriteHeader extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).WriteHeader(hdr *Header) error
    WriterWriteHeader() { this.(Method).hasQualifiedName("archive/tar", "Writer", "WriteHeader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }
}
