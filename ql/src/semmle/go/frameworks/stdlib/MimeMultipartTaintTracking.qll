/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `mime/multipart` package. */
module MimeMultipartTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func NewReader(r io.Reader, boundary string) *Reader
      hasQualifiedName("mime/multipart", "NewReader") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func NewWriter(w io.Writer) *Writer
      hasQualifiedName("mime/multipart", "NewWriter") and
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
      // signature: func (*FileHeader).Open() (File, error)
      this.(Method).hasQualifiedName("mime/multipart", "FileHeader", "Open") and
      (inp.isReceiver() and outp.isResult(0))
      or
      // signature: func (*Part).Read(d []byte) (n int, err error)
      this.(Method).hasQualifiedName("mime/multipart", "Part", "Read") and
      (inp.isReceiver() and outp.isParameter(0))
      or
      // signature: func (*Reader).NextPart() (*Part, error)
      this.(Method).hasQualifiedName("mime/multipart", "Reader", "NextPart") and
      (inp.isReceiver() and outp.isResult(0))
      or
      // signature: func (*Reader).NextRawPart() (*Part, error)
      this.(Method).hasQualifiedName("mime/multipart", "Reader", "NextRawPart") and
      (inp.isReceiver() and outp.isResult(0))
      or
      // signature: func (*Reader).ReadForm(maxMemory int64) (*Form, error)
      this.(Method).hasQualifiedName("mime/multipart", "Reader", "ReadForm") and
      (inp.isReceiver() and outp.isResult(0))
      or
      // signature: func (*Writer).CreateFormField(fieldname string) (io.Writer, error)
      this.(Method).hasQualifiedName("mime/multipart", "Writer", "CreateFormField") and
      (inp.isResult(0) and outp.isReceiver())
      or
      // signature: func (*Writer).CreateFormFile(fieldname string, filename string) (io.Writer, error)
      this.(Method).hasQualifiedName("mime/multipart", "Writer", "CreateFormFile") and
      (inp.isResult(0) and outp.isReceiver())
      or
      // signature: func (*Writer).CreatePart(header net/textproto.MIMEHeader) (io.Writer, error)
      this.(Method).hasQualifiedName("mime/multipart", "Writer", "CreatePart") and
      (inp.isResult(0) and outp.isReceiver())
      or
      // signature: func (*Writer).WriteField(fieldname string, value string) error
      this.(Method).hasQualifiedName("mime/multipart", "Writer", "WriteField") and
      (inp.isParameter(_) and outp.isReceiver())
      // Interfaces:
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
