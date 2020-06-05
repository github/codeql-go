/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `encoding/csv` package. */
module EncodingCsvTaintTracking {
  private class NewReader extends TaintTracking::FunctionModel {
    // signature: func NewReader(r io.Reader) *Reader
    NewReader() { hasQualifiedName("encoding/csv", "NewReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class NewWriter extends TaintTracking::FunctionModel {
    // signature: func NewWriter(w io.Writer) *Writer
    NewWriter() { hasQualifiedName("encoding/csv", "NewWriter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult() and outp.isParameter(0)
    }
  }

  private class ReaderReadAll extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadAll() (records [][]string, err error)
    ReaderReadAll() { this.(Method).hasQualifiedName("encoding/csv", "Reader", "ReadAll") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class WriterWrite extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).Write(record []string) error
    WriterWrite() { this.(Method).hasQualifiedName("encoding/csv", "Writer", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class WriterWriteAll extends TaintTracking::FunctionModel, Method {
    // signature: func (*Writer).WriteAll(records [][]string) error
    WriterWriteAll() { this.(Method).hasQualifiedName("encoding/csv", "Writer", "WriteAll") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }
}
