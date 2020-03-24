/**
 * Provides concrete classes for data-flow nodes that execute an
 * operating system command, for instance by spawning a new process.
 */

import go

private class SystemCommandExecutors extends SystemCommandExecution::Range, DataFlow::CallNode {
  string packagePath;
  string functionName;
  int cmdArgIndex;
  DataFlow::ExprNode argumentNode;

  SystemCommandExecutors() {
    (
      packagePath = "os/exec" and
      (
        functionName = "Command" and cmdArgIndex = 0
        or
        functionName = "CommandContext" and cmdArgIndex = 1
      )
      or
      packagePath = "os" and
      functionName = "StartProcess" and
      cmdArgIndex = 0
    ) and
    exists(Function fn |
      fn.hasQualifiedName(packagePath, functionName) and
      (
        // Either there is another argument to this call that is a sudo/shell ...
        exists(ShellLike shellOrSudoArg |
          shellOrSudoArg = this.getAnArgument().getAPredecessor*()
        |
          argumentNode = this.getAnArgument()
        )
        or
        // ... or just use the default cmdArgIndex, which is different for each function:
        argumentNode = this.getArgument(cmdArgIndex)
      )
    |
      this = fn.getACall()
    )
  }

  override DataFlow::Node getCommandName() { result = argumentNode }
}

/**
 * A call to the `Command` function, or `Call` or `Command` methods on a `Session` object
 * from the [go-sh](https://github.com/codeskyblue/go-sh) package, viewed as a
 * system-command execution.
 */
private class GoShCommandExecution extends SystemCommandExecution::Range, DataFlow::CallNode {
  string packagePath;
  DataFlow::ExprNode argumentNode;

  GoShCommandExecution() {
    packagePath = "github.com/codeskyblue/go-sh" and
    (
      // Catch method calls on the `Session` object:
      exists(DataFlow::MethodCallNode call, Method method |
        call = method.getACall() and
        (
          method.hasQualifiedName(packagePath, "Session", "Call")
          or
          method.hasQualifiedName(packagePath, "Session", "Command")
        ) and
        (
          // Either there is another argument to this call that is a sudo/shell ...
          exists(ShellLike shellOrSudoArg |
            shellOrSudoArg = this.getAnArgument().getAPredecessor*()
          |
            argumentNode = this.getAnArgument()
          )
          or
          // ... or just use the first argument:
          argumentNode = this.getArgument(0)
        )
      |
        this = call and
        argumentNode = call.getAnArgument()
      )
      or
      // Catch calls to the `Command` function:
      exists(Function fn |
        fn.hasQualifiedName(packagePath, "Command") and
        (
          // Either there is another argument to this call that is a sudo/shell ...
          exists(ShellLike shellOrSudoArg |
            shellOrSudoArg = this.getAnArgument().getAPredecessor*()
          |
            argumentNode = this.getAnArgument()
          )
          or
          // ... or just use the first argument:
          argumentNode = this.getArgument(0)
        )
      |
        this = fn.getACall()
      )
    )
  }

  override DataFlow::Node getCommandName() { result = argumentNode }
}

/**
 * A call to a method on a `Session` object from the [ssh](golang.org/x/crypto/ssh)
 * package, viewed as a system-command execution.
 */
private class SSHCommandExecution extends SystemCommandExecution::Range, DataFlow::CallNode {
  DataFlow::ExprNode argumentNode;

  SSHCommandExecution() {
    // Catch method calls on the `Session` object:
    exists(DataFlow::MethodCallNode call, Method method |
      call = method.getACall() and
      exists(string methodName |
        method.hasQualifiedName("golang.org/x/crypto/ssh", "Session", methodName)
      |
        methodName = "CombinedOutput"
        or
        methodName = "Output"
        or
        methodName = "Run"
        or
        methodName = "Start"
      )
    |
      this = call and
      argumentNode = call.getArgument(0)
    )
  }

  override DataFlow::Node getCommandName() { result = argumentNode }
}

private class ShellLike extends DataFlow::Node {
  ShellLike() {
    isSudoOrSimilar(this) or
    isShell(this) or
    isProgrammingLanguageCLI(this) or
    isSSH(this)
  }
}

predicate isSudoOrSimilar(DataFlow::Node node) {
  node.getStringValue().matches("%/sudo")
  or
  node.getStringValue() = "sudo"
  or
  node.getStringValue().matches("%/su")
  or
  node.getStringValue() = "su"
  or
  node.getStringValue().matches("%/doas")
  or
  node.getStringValue() = "doas"
  or
  node.getStringValue().matches("%/access")
  or
  node.getStringValue() = "access"
  or
  node.getStringValue().matches("%/vsys")
  or
  node.getStringValue() = "vsys"
  or
  node.getStringValue().matches("%/userv")
  or
  node.getStringValue() = "userv"
  or
  node.getStringValue().matches("%/sus")
  or
  node.getStringValue() = "sus"
  or
  node.getStringValue().matches("%/super")
  or
  node.getStringValue() = "super"
  or
  node.getStringValue().matches("%/priv")
  or
  node.getStringValue() = "priv"
  or
  node.getStringValue().matches("%/calife")
  or
  node.getStringValue() = "calife"
  or
  node.getStringValue().matches("%/ssu")
  or
  node.getStringValue() = "ssu"
  or
  node.getStringValue().matches("%/su1")
  or
  node.getStringValue() = "su1"
  or
  node.getStringValue().matches("%/op")
  or
  node.getStringValue() = "op"
  or
  node.getStringValue().matches("%/sudowin")
  or
  node.getStringValue() = "sudowin"
  or
  node.getStringValue().matches("%/sudown")
  or
  node.getStringValue() = "sudown"
}

predicate isShell(DataFlow::Node node) {
  node.getStringValue().matches("%/bash")
  or
  node.getStringValue() = "bash"
  or
  node.getStringValue().matches("%/sh")
  or
  node.getStringValue() = "sh"
  or
  node.getStringValue().matches("%/rbash")
  or
  node.getStringValue() = "rbash"
  or
  node.getStringValue().matches("%/dash")
  or
  node.getStringValue() = "dash"
  or
  node.getStringValue().matches("%/zsh")
  or
  node.getStringValue() = "zsh"
  or
  node.getStringValue().matches("%/csh")
  or
  node.getStringValue() = "csh"
  or
  node.getStringValue().matches("%/tcsh")
  or
  node.getStringValue() = "tcsh"
  or
  node.getStringValue().matches("%/ksh")
  or
  node.getStringValue() = "ksh"
  or
  node.getStringValue().matches("%/fish")
  or
  node.getStringValue() = "fish"
  or
  node.getStringValue().matches("%/dash")
  or
  node.getStringValue() = "dash"
  or
  node.getStringValue().matches("%/pwsh")
  or
  node.getStringValue() = "pwsh"
  or
  node.getStringValue().matches("%/elvish")
  or
  node.getStringValue() = "elvish"
  or
  node.getStringValue().matches("%/oh")
  or
  node.getStringValue() = "oh"
  or
  node.getStringValue().matches("%/ion")
  or
  node.getStringValue() = "ion"
  or
  node.getStringValue().matches("%/ksh")
  or
  node.getStringValue() = "ksh"
  or
  node.getStringValue().matches("%/rksh")
  or
  node.getStringValue() = "rksh"
  or
  node.getStringValue().matches("%/tksh")
  or
  node.getStringValue() = "tksh"
  or
  node.getStringValue().matches("%/mksh")
  or
  node.getStringValue() = "mksh"
  or
  node.getStringValue().matches("%/nu")
  or
  node.getStringValue() = "nu"
  or
  node.getStringValue().matches("%/oksh")
  or
  node.getStringValue() = "oksh"
  or
  node.getStringValue().matches("%/osh")
  or
  node.getStringValue() = "osh"
  or
  node.getStringValue().matches("%/shpp")
  or
  node.getStringValue() = "shpp"
  or
  node.getStringValue().matches("%/xiki")
  or
  node.getStringValue() = "xiki"
  or
  node.getStringValue().matches("%/xonsh")
  or
  node.getStringValue() = "xonsh"
  or
  node.getStringValue().matches("%/yash")
  or
  node.getStringValue() = "yash"
}

predicate isProgrammingLanguageCLI(DataFlow::Node node) {
  node.getStringValue().matches("%/python%")
  or
  node.getStringValue() = "python"
  or
  node.getStringValue().matches("%/php%")
  or
  node.getStringValue() = "php"
  or
  node.getStringValue().matches("%/ruby")
  or
  node.getStringValue() = "ruby"
}

predicate isSSH(DataFlow::Node node) {
  node.getStringValue().matches("%/ssh")
  or
  node.getStringValue() = "ssh"
}
