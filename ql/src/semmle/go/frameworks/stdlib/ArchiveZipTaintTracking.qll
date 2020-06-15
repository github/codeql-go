/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `archive/zip` package. */
module ArchiveZipTaintTracking {
  private class FileInfoHeader extends TaintTracking::FunctionModel {
    // signature: func FileInfoHeader(fi os.FileInfo) (*FileHeader, error)
    FileInfoHeader() { hasQualifiedName("archive/zip", "FileInfoHeader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class NewReader extends TaintTracking::FunctionModel {
    // signature: func NewReader(r io.ReaderAt, size int64) (*Reader, error)
    NewReader() { hasQualifiedName("archive/zip", "NewReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class NewWriter extends TaintTracking::FunctionModel {
    // signature: func NewWriter(w io.Writer) *Writer
    NewWriter() { hasQualifiedName("archive/zip", "NewWriter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isResult() and outp.isParameter(0))
    }
  }

  private class OpenReader extends TaintTracking::FunctionModel {
    // signature: func OpenReader(name string) (*ReadCloser, error)
    OpenReader() { hasQualifiedName("archive/zip", "OpenReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class FileOpen extends TaintTracking::FunctionModel, Method {
    // signature: func (*File).Open() (io.ReadCloser, error)
    FileOpen() { this.(Method).hasQualifiedName("archive/zip", "File", "Open") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }

  private class FileHeaderFileInfo extends TaintTracking::FunctionModel, Method {
    // signature: func (*FileHeader).FileInfo() os.FileInfo
    FileHeaderFileInfo() { this.(Method).hasQualifiedName("archive/zip", "FileHeader", "FileInfo") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class FileHeaderMode extends TaintTracking::FunctionModel, Method {
    // signature: func (*FileHeader).Mode() (mode os.FileMode)
    FileHeaderMode() { this.(Method).hasQualifiedName("archive/zip", "FileHeader", "Mode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class FileHeaderSetMode extends TaintTracking::FunctionModel, Method {
    // signature: func (*FileHeader).SetMode(mode os.FileMode)
    FileHeaderSetMode() { this.(Method).hasQualifiedName("archive/zip", "FileHeader", "SetMode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class ReaderRegisterDecompressor extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).RegisterDecompressor(method uint16, dcomp Decompressor)
    ReaderRegisterDecompressor() {
      this.(Method).hasQualifiedName("archive/zip", "Reader", "RegisterDecompressor")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isReceiver())
    }
  }

  private class WriterCreate extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).Create(name string) (io.Writer, error)
    WriterCreate() { this.(Method).hasQualifiedName("archive/zip", "Writer", "Create") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isResult(0) and outp.isReceiver())
    }
  }

  private class WriterCreateHeader extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).CreateHeader(fh *FileHeader) (io.Writer, error)
    WriterCreateHeader() { this.(Method).hasQualifiedName("archive/zip", "Writer", "CreateHeader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (
        (inp.isParameter(0) or inp.isResult(0)) and
        outp.isReceiver()
      )
    }
  }

  private class WriterRegisterCompressor extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).RegisterCompressor(method uint16, comp Compressor)
    WriterRegisterCompressor() {
      this.(Method).hasQualifiedName("archive/zip", "Writer", "RegisterCompressor")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isReceiver())
    }
  }

  private class WriterSetComment extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).SetComment(comment string) error
    WriterSetComment() { this.(Method).hasQualifiedName("archive/zip", "Writer", "SetComment") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }
}
