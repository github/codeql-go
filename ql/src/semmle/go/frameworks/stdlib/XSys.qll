/**
 * Provides classes modeling security-relevant aspects of the `golang.org/x/sys` package.
 */

import go

/** Provides models of commonly used functions in the `golang.org/x/sys` package. */
module XSys {
  /** Name of the `golang.org/x/sys/unix` package. */
  string unixPackage() { result = "golang.org/x/sys/unix" }

  /** Name of the `golang.org/x/sys/windows` package. */
  string windowsPackage() { result = "golang.org/x/sys/windows" }

  private class UnixOrWindowsFsAccess extends FileSystemAccess::Range, DataFlow::CallNode {
    int pathArgumentIndex;

    UnixOrWindowsFsAccess() {
      exists(string fn | getTarget().hasQualifiedName([unixPackage(), windowsPackage()], fn) |
        fn = "Chmod" and pathArgumentIndex = 0
        or
        fn = "Chown" and pathArgumentIndex = 0
        or
        fn = "Lchown" and pathArgumentIndex = 0
        or
        fn = "Link" and pathArgumentIndex = [0, 1]
        or
        fn = "Mkdir" and pathArgumentIndex = 0
        or
        fn = "Readlink" and pathArgumentIndex = 0
        or
        fn = "Rename" and pathArgumentIndex = [0, 1]
        or
        fn = "Rmdir" and pathArgumentIndex = 0
        or
        fn = "Symlink" and pathArgumentIndex = [0, 1]
        or
        fn = "Unlink" and pathArgumentIndex = 0
      )
    }

    override DataFlow::Node getAPathArgument() { result = getArgument(pathArgumentIndex) }
  }

  private class UnixFsAccess extends FileSystemAccess::Range, DataFlow::CallNode {
    int pathArgumentIndex;

    UnixFsAccess() {
      exists(string fn | getTarget().hasQualifiedName(unixPackage(), fn) |
        fn = "Access" and pathArgumentIndex = 0
        or
        fn = "Chdir" and pathArgumentIndex = 0
        or
        fn = "Chroot" and pathArgumentIndex = 0
        or
        fn = "Creat" and pathArgumentIndex = 0
        or
        fn = "Getxattr" and pathArgumentIndex = 0
        or
        fn = "Lgetxattr" and pathArgumentIndex = 0
        or
        fn = "Linkat" and pathArgumentIndex = [1, 3]
        or
        fn = "Lremovexattr" and pathArgumentIndex = 0
        or
        fn = "Lsetxattr" and pathArgumentIndex = 0
        or
        fn = "Lstat" and pathArgumentIndex = 0
        or
        fn = "Lutimes" and pathArgumentIndex = 0
        or
        fn = "Mkdirat" and pathArgumentIndex = 1
        or
        fn = "Mkfifo" and pathArgumentIndex = 0
        or
        fn = "Mkfifoat" and pathArgumentIndex = 1
        or
        fn = "Mknod" and pathArgumentIndex = 0
        or
        fn = "Mknodat" and pathArgumentIndex = 1
        or
        fn = "Open" and pathArgumentIndex = 0
        or
        fn = "Openat" and pathArgumentIndex = 1
        or
        fn = "Openat2" and pathArgumentIndex = 1
        or
        fn = "Readlinkat" and pathArgumentIndex = 1
        or
        fn = "Removexattr" and pathArgumentIndex = 0
        or
        fn = "Renameat" and pathArgumentIndex = [1, 3]
        or
        fn = "Renameat2" and pathArgumentIndex = [1, 3]
        or
        fn = "Stat" and pathArgumentIndex = 0
        or
        fn = "Statfs" and pathArgumentIndex = 0
        or
        fn = "Statx" and pathArgumentIndex = 1
        or
        fn = "Symlinkat" and pathArgumentIndex = [0, 2]
        or
        fn = "Truncate" and pathArgumentIndex = 0
        or
        fn = "Unlinkat" and pathArgumentIndex = 1
        or
        fn = "Utime" and pathArgumentIndex = 0
        or
        fn = "Utimes" and pathArgumentIndex = 0
        or
        fn = "UtimesNano" and pathArgumentIndex = 0
        or
        fn = "UtimesNanoAt" and pathArgumentIndex = 1
      )
    }

    override DataFlow::Node getAPathArgument() { result = getArgument(pathArgumentIndex) }
  }

  private class WindowsFsAccess extends FileSystemAccess::Range, DataFlow::CallNode {
    int pathArgumentIndex;

    WindowsFsAccess() {
      exists(string fn | getTarget().hasQualifiedName(windowsPackage(), fn) |
        fn = "CreateDirectory" and pathArgumentIndex = 0
        or
        fn = "CreateHardLink" and pathArgumentIndex = [0, 1]
        or
        fn = "CreateSymbolicLink" and pathArgumentIndex = [0, 1]
        or
        fn = "DeleteFile" and pathArgumentIndex = 0
        or
        fn = "GetFileAttributes" and pathArgumentIndex = 0
        or
        fn = "GetFileAttributesEx" and pathArgumentIndex = 0
        or
        fn = "MoveFile" and pathArgumentIndex = [0, 1]
        or
        fn = "MoveFileEx" and pathArgumentIndex = [0, 1]
        or
        fn = "RemoveDirectory" and pathArgumentIndex = 0
        or
        fn = "SetCurrentDirectory" and pathArgumentIndex = 0
        or
        fn = "SetFileAttributes" and pathArgumentIndex = 0
      )
    }

    override DataFlow::Node getAPathArgument() { result = getArgument(pathArgumentIndex) }
  }

  private class ProcessExecutors extends SystemCommandExecution::Range, DataFlow::CallNode {
    int cmdArg;

    ProcessExecutors() {
      exists(string pkg, string name | this.getTarget().hasQualifiedName(pkg, name) |
        pkg = unixPackage() and name = "Exec" and cmdArg = 0
        or
        pkg = windowsPackage() and name = "CreateProcess" and cmdArg = 1
        or
        pkg = windowsPackage() and name = "ShellExecute" and cmdArg = [2, 4]
      )
    }

    override DataFlow::Node getCommandName() { result = getArgument(cmdArg) }
  }

  private class WindowsStringMethods extends TaintTracking::FunctionModel {
    FunctionInput input;
    FunctionOutput output;
    string fname;

    WindowsStringMethods() {
      this.hasQualifiedName(windowsPackage(), fname) and
      (
        fname = "GetFullPathName" and input.isParameter(0) and output.isParameter([2, 3])
        or
        fname = "GetLongPathName" and input.isParameter(0) and output.isParameter(1)
        or
        fname = "MultiByteToWideChar" and input.isParameter(2) and output.isParameter(4)
        or
        fname =
          ["FullPath", "StringToUTF16", "StringToUTF16Ptr", "UTF16FromString", "UTF16PtrFromString",
              "UTF16PtrToString", "UTF16ToString"] and
        input.isParameter(0) and
        output.isResult(0)
      )
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp = input and outp = output
    }
  }
}
