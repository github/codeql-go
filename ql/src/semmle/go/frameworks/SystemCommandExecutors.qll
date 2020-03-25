/**
 * Provides concrete classes for data-flow nodes that execute an
 * operating system command, for instance by spawning a new process.
 */

import go

private class ShellOrSudoExecution extends SystemCommandExecution::Range, DataFlow::CallNode {
  DataFlow::ExprNode argumentNode;

  ShellOrSudoExecution() {
    exists(SystemCommandExecution exec |
      // Either there is another argument to this call that is a sudo/shell ...
      exists(ShellLike shellOrSudoArg |
        shellOrSudoArg = exec.(DataFlow::CallNode).getAnArgument().getAPredecessor*()
      |
        argumentNode = exec.(DataFlow::CallNode).getAnArgument()
      )
      or
      // ... or just use the default getCommandName:
      argumentNode = exec.getCommandName()
    |
      this = exec
    )
  }

  override DataFlow::Node getCommandName() { result = argumentNode }
}

private class SystemCommandExecutors extends SystemCommandExecution::Range, DataFlow::CallNode {
  DataFlow::ExprNode argumentNode;

  SystemCommandExecutors() {
    exists(Function fn, string packagePath, string functionName, int cmdArgIndex |
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
      fn.hasQualifiedName(packagePath, functionName)
    |
      this = fn.getACall() and
      argumentNode = this.getArgument(cmdArgIndex)
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
  DataFlow::ExprNode argumentNode;

  GoShCommandExecution() {
    exists(string packagePath |
      packagePath = "github.com/codeskyblue/go-sh" and
      (
        // Catch method calls on the `Session` object:
        exists(Method method |
          method.hasQualifiedName(packagePath, "Session", "Call")
          or
          method.hasQualifiedName(packagePath, "Session", "Command")
        |
          this = method.getACall()
        )
        or
        // Catch calls to the `Command` function:
        getTarget().hasQualifiedName(packagePath, "Command")
      )
    |
      argumentNode = this.getArgument(0)
    )
  }

  override DataFlow::Node getCommandName() { result = argumentNode }
}

/**
 * A call to a method on a `Session` object from the [ssh](golang.org/x/crypto/ssh)
 * package, viewed as a system-command execution.
 */
private class SshCommandExecution extends SystemCommandExecution::Range, DataFlow::CallNode {
  DataFlow::ExprNode argumentNode;

  SshCommandExecution() {
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
    isProgrammingLanguageCli(this) or
    isSsh(this)
  }
}

private string getASudoCommand() {
  result = "sudo" or
  result = "su" or
  result = "doas" or
  result = "access" or
  result = "vsys" or
  result = "userv" or
  result = "sus" or
  result = "super" or
  result = "priv" or
  result = "calife" or
  result = "ssu" or
  result = "su1" or
  result = "op" or
  result = "sudowin" or
  result = "sudown"
}

private predicate isSudoOrSimilar(DataFlow::Node node) {
  exists(string regex |
    regex = ".*(^|/)(" + concat(string cmd | cmd = getASudoCommand() | cmd, "|") + ")"
  |
    node.getStringValue().regexpMatch(regex)
  )
}

private string getAShellCommand() {
  result = "bash" or
  result = "sh" or
  result = "rbash" or
  result = "dash" or
  result = "zsh" or
  result = "csh" or
  result = "tcsh" or
  result = "fish" or
  result = "pwsh" or
  result = "elvish" or
  result = "oh" or
  result = "ion" or
  result = "ksh" or
  result = "rksh" or
  result = "tksh" or
  result = "mksh" or
  result = "nu" or
  result = "oksh" or
  result = "osh" or
  result = "shpp" or
  result = "xiki" or
  result = "xonsh" or
  result = "yash"
}

private predicate isShell(DataFlow::Node node) {
  exists(string regex |
    regex = ".*(^|/)(" + concat(string cmd | cmd = getAShellCommand() | cmd, "|") + ")"
  |
    node.getStringValue().regexpMatch(regex)
  )
}

private string getAProgrammingLanguageCliCommand() {
  result = "python" or
  result = "php" or
  result = "ruby" or
  result = "perl"
}

private predicate isProgrammingLanguageCli(DataFlow::Node node) {
  // NOTE: we can enounter cases like /usr/bin/python3.1
  node.getStringValue().matches("%/python%")
  or
  exists(string regex |
    regex =
      ".*(^|/)(" + concat(string cmd | cmd = getAProgrammingLanguageCliCommand() | cmd, "|") + ")"
  |
    node.getStringValue().regexpMatch(regex)
  )
}

private string getASshCommand() { result = "ssh" }

private predicate isSsh(DataFlow::Node node) {
  exists(string regex |
    regex = ".*(^|/)(" + concat(string cmd | cmd = getASshCommand() | cmd, "|") + ")"
  |
    node.getStringValue().regexpMatch(regex)
  )
}
