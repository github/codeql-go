/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `encoding/csv` package. */
module EncodingCsvTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func NewReader(r io.Reader) *Reader
      hasQualifiedName("encoding/csv", "NewReader") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func NewWriter(w io.Writer) *Writer
      hasQualifiedName("encoding/csv", "NewWriter") and
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
      // signature: func (*Reader).Read() (record []string, err error)
      this.(Method).hasQualifiedName("encoding/csv", "Reader", "Read") and
      (inp.isReceiver() and outp.isResult(0))
      or
      // signature: func (*Reader).ReadAll() (records [][]string, err error)
      this.(Method).hasQualifiedName("encoding/csv", "Reader", "ReadAll") and
      (inp.isReceiver() and outp.isResult(0))
      or
      // signature: func (*Writer).Write(record []string) error
      this.(Method).hasQualifiedName("encoding/csv", "Writer", "Write") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (*Writer).WriteAll(records [][]string) error
      this.(Method).hasQualifiedName("encoding/csv", "Writer", "WriteAll") and
      (inp.isParameter(0) and outp.isReceiver())
      // Interfaces:
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
