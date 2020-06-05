/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `mime/multipart` package. */
module MimeMultipartTaintTracking {
  private class NewReader extends TaintTracking::FunctionModel {
    // signature: func NewReader(r io.Reader, boundary string) *Reader
    NewReader() { hasQualifiedName("mime/multipart", "NewReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class NewWriter extends TaintTracking::FunctionModel {
    // signature: func NewWriter(w io.Writer) *Writer
    NewWriter() { hasQualifiedName("mime/multipart", "NewWriter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult() and outp.isParameter(0)
    }
  }

  private class FileHeaderOpen extends TaintTracking::FunctionModel, Method {
    // signature: func (*FileHeader).Open() (File, error)
    FileHeaderOpen() { this.(Method).hasQualifiedName("mime/multipart", "FileHeader", "Open") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class ReaderNextPart extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).NextPart() (*Part, error)
    ReaderNextPart() { this.(Method).hasQualifiedName("mime/multipart", "Reader", "NextPart") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class ReaderNextRawPart extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).NextRawPart() (*Part, error)
    ReaderNextRawPart() {
      this.(Method).hasQualifiedName("mime/multipart", "Reader", "NextRawPart")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class ReaderReadForm extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadForm(maxMemory int64) (*Form, error)
    ReaderReadForm() { this.(Method).hasQualifiedName("mime/multipart", "Reader", "ReadForm") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class WriterCreateFormField extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).CreateFormField(fieldname string) (io.Writer, error)
    WriterCreateFormField() {
      this.(Method).hasQualifiedName("mime/multipart", "Writer", "CreateFormField")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult(0) and outp.isReceiver()
    }
  }

  private class WriterCreateFormFile extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).CreateFormFile(fieldname string, filename string) (io.Writer, error)
    WriterCreateFormFile() {
      this.(Method).hasQualifiedName("mime/multipart", "Writer", "CreateFormFile")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult(0) and outp.isReceiver()
    }
  }

  private class WriterCreatePart extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).CreatePart(header net/textproto.MIMEHeader) (io.Writer, error)
    WriterCreatePart() { this.(Method).hasQualifiedName("mime/multipart", "Writer", "CreatePart") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult(0) and outp.isReceiver()
    }
  }

  private class WriterWriteField extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).WriteField(fieldname string, value string) error
    WriterWriteField() { this.(Method).hasQualifiedName("mime/multipart", "Writer", "WriteField") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isReceiver()
    }
  }
}
