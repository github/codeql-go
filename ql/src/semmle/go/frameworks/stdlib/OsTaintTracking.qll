/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `os` package. */
module OsTaintTracking {
  private class Expand extends TaintTracking::FunctionModel {
    // signature: func Expand(s string, mapping func(string) string) string
    Expand() { hasQualifiedName("os", "Expand") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class ExpandEnv extends TaintTracking::FunctionModel {
    // signature: func ExpandEnv(s string) string
    ExpandEnv() { hasQualifiedName("os", "ExpandEnv") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class NewFile extends TaintTracking::FunctionModel {
    // signature: func NewFile(fd uintptr, name string) *File
    NewFile() { hasQualifiedName("os", "NewFile") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class Pipe extends TaintTracking::FunctionModel {
    // signature: func Pipe() (r *File, w *File, err error)
    Pipe() { hasQualifiedName("os", "Pipe") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isResult(1) and outp.isResult(0))
    }
  }

  private class FileFd extends TaintTracking::FunctionModel, Method {
    // signature: func (*File).Fd() uintptr
    FileFd() { this.(Method).hasQualifiedName("os", "File", "Fd") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class FileRead extends TaintTracking::FunctionModel, Method {
    // signature: func (*File).Read(b []byte) (n int, err error)
    FileRead() { this.(Method).hasQualifiedName("os", "File", "Read") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class FileReadAt extends TaintTracking::FunctionModel, Method {
    // signature: func (*File).ReadAt(b []byte, off int64) (n int, err error)
    FileReadAt() { this.(Method).hasQualifiedName("os", "File", "ReadAt") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class FileSyscallConn extends TaintTracking::FunctionModel, Method {
    // signature: func (*File).SyscallConn() (syscall.RawConn, error)
    FileSyscallConn() { this.(Method).hasQualifiedName("os", "File", "SyscallConn") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
      or
      inp.isResult(0) and outp.isReceiver()
    }
  }

  private class FileWrite extends TaintTracking::FunctionModel, Method {
    // signature: func (*File).Write(b []byte) (n int, err error)
    FileWrite() { this.(Method).hasQualifiedName("os", "File", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class FileWriteAt extends TaintTracking::FunctionModel, Method {
    // signature: func (*File).WriteAt(b []byte, off int64) (n int, err error)
    FileWriteAt() { this.(Method).hasQualifiedName("os", "File", "WriteAt") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class FileWriteString extends TaintTracking::FunctionModel, Method {
    // signature: func (*File).WriteString(s string) (n int, err error)
    FileWriteString() { this.(Method).hasQualifiedName("os", "File", "WriteString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }
}
