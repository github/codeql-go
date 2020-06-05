/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `syscall` package. */
module SyscallTaintTracking {
  private class BytePtrFromString extends TaintTracking::FunctionModel {
    // signature: func BytePtrFromString(s string) (*byte, error)
    BytePtrFromString() { hasQualifiedName("syscall", "BytePtrFromString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class ByteSliceFromString extends TaintTracking::FunctionModel {
    // signature: func ByteSliceFromString(s string) ([]byte, error)
    ByteSliceFromString() { hasQualifiedName("syscall", "ByteSliceFromString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class Pread extends TaintTracking::FunctionModel {
    // signature: func Pread(fd int, p []byte, offset int64) (n int, err error)
    Pread() { hasQualifiedName("syscall", "Pread") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isParameter(1)
    }
  }

  private class Pwrite extends TaintTracking::FunctionModel {
    // signature: func Pwrite(fd int, p []byte, offset int64) (n int, err error)
    Pwrite() { hasQualifiedName("syscall", "Pwrite") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class SyscallRead extends TaintTracking::FunctionModel {
    // signature: func Read(fd int, p []byte) (n int, err error)
    SyscallRead() { hasQualifiedName("syscall", "Read") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isParameter(1)
    }
  }

  private class ReadDirent extends TaintTracking::FunctionModel {
    // signature: func ReadDirent(fd int, buf []byte) (n int, err error)
    ReadDirent() { hasQualifiedName("syscall", "ReadDirent") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isParameter(1)
    }
  }

  private class SlicePtrFromStrings extends TaintTracking::FunctionModel {
    // signature: func SlicePtrFromStrings(ss []string) ([]*byte, error)
    SlicePtrFromStrings() { hasQualifiedName("syscall", "SlicePtrFromStrings") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
    }
  }

  private class StringBytePtr extends TaintTracking::FunctionModel {
    // signature: func StringBytePtr(s string) *byte
    StringBytePtr() { hasQualifiedName("syscall", "StringBytePtr") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class StringByteSlice extends TaintTracking::FunctionModel {
    // signature: func StringByteSlice(s string) []byte
    StringByteSlice() { hasQualifiedName("syscall", "StringByteSlice") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class StringSlicePtr extends TaintTracking::FunctionModel {
    // signature: func StringSlicePtr(ss []string) []*byte
    StringSlicePtr() { hasQualifiedName("syscall", "StringSlicePtr") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class Tee extends TaintTracking::FunctionModel {
    // signature: func Tee(rfd int, wfd int, len int, flags int) (n int64, err error)
    Tee() { hasQualifiedName("syscall", "Tee") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class UnixCredentials extends TaintTracking::FunctionModel {
    // signature: func UnixCredentials(ucred *Ucred) []byte
    UnixCredentials() { hasQualifiedName("syscall", "UnixCredentials") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class UnixRights extends TaintTracking::FunctionModel {
    // signature: func UnixRights(fds ...int) []byte
    UnixRights() { hasQualifiedName("syscall", "UnixRights") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(_) and outp.isResult()
    }
  }

  private class SyscallWrite extends TaintTracking::FunctionModel {
    // signature: func Write(fd int, p []byte) (n int, err error)
    SyscallWrite() { hasQualifiedName("syscall", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isParameter(0)
    }
  }

  private class RawConnRead extends TaintTracking::FunctionModel, Method {
    // signature: func (RawConn).Read(f func(fd uintptr) (done bool)) error
    RawConnRead() { this.implements("syscall", "RawConn", "Read") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isParameter(0)
    }
  }

  private class RawConnWrite extends TaintTracking::FunctionModel, Method {
    // signature: func (RawConn).Write(f func(fd uintptr) (done bool)) error
    RawConnWrite() { this.implements("syscall", "RawConn", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }
}
