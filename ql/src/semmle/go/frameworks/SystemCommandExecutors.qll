/**
 * Provides concrete classes for data-flow nodes that execute an
 * operating system command, for instance by spawning a new process.
 */

import go

/**
 * ShellOrSudoExecution extends system command execution sinks to include
 * arguments passed to command interpreter binaries like shells, sudo,
 * or programming language interpreters.
 */
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
    |
      this = exec
    )
  }

  override DataFlow::Node getCommandName() { result = argumentNode }
}

private class SystemCommandExecutors extends SystemCommandExecution::Range, DataFlow::CallNode {
  int cmdArg;

  SystemCommandExecutors() {
    exists(string pkg, string name | this.getTarget().hasQualifiedName(pkg, name) |
      pkg = "os" and name = "StartProcess" and cmdArg = 0
      or
      // assume that if a `Cmd` is instantiated it will be run
      pkg = "os/exec" and name = "Command" and cmdArg = 0
      or
      pkg = "os/exec" and name = "CommandContext" and cmdArg = 1
    )
  }

  override DataFlow::Node getCommandName() { result = this.getArgument(cmdArg) }
}

/**
 * A call to the `Command` function, or `Call` or `Command` methods on a `Session` object
 * from the [go-sh](https://github.com/codeskyblue/go-sh) package, viewed as a
 * system-command execution.
 */
private class GoShCommandExecution extends SystemCommandExecution::Range, DataFlow::CallNode {
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
    )
  }

  override DataFlow::Node getCommandName() { result = this.getArgument(0) }
}

/**
 * A call to a method on a `Session` object from the [ssh](golang.org/x/crypto/ssh)
 * package, viewed as a system-command execution.
 */
private class SshCommandExecution extends SystemCommandExecution::Range, DataFlow::CallNode {
  SshCommandExecution() {
    // Catch method calls on the `Session` object:
    exists(Method method, string methodName |
      method.hasQualifiedName("golang.org/x/crypto/ssh", "Session", methodName) and
      methodName = "CombinedOutput"
      or
      methodName = "Output"
      or
      methodName = "Run"
      or
      methodName = "Start"
    |
      this = method.getACall()
    )
  }

  override DataFlow::Node getCommandName() { result = this.getArgument(0) }
}

/**
 * ShellLike defines a string value node whose value is any binary
 * (shells, sudo, programming language interpreters, ssh clients, etc.)
 * which executes arguments passed to as commands.
 */
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

/**
 * isSudoOrSimilar defines a string value node whose value is any binary
 * that work in a similar manner to `sudo`, whose arguments are
 * interpreted as system commands.
 */
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

/**
 * isShell defines a string value node whose value is any shell,
 * whose arguments are interpreted as system commands.
 */
private predicate isShell(DataFlow::Node node) {
  exists(string regex |
    regex = ".*(^|/)(" + concat(string cmd | cmd = getAShellCommand() | cmd, "|") + ")"
  |
    node.getStringValue().regexpMatch(regex)
  )
}

private string getAnInterpreterName() {
  result = "python" or
  result = "php" or
  result = "ruby" or
  result = "perl" or
  result = "node" or
  result = "nodejs"
}

/**
 * isShell defines a string value node whose value is any CLI
 * programming language interpreter, whose arguments are
 * executed as code.
 */
private predicate isProgrammingLanguageCli(DataFlow::Node node) {
  // NOTE: we can enounter cases like /usr/bin/python3.1
  node.getStringValue().matches("%/python%")
  or
  exists(string regex |
    regex = ".*(^|/)(" + concat(string cmd | cmd = getAnInterpreterName() | cmd, "|") + ")"
  |
    node.getStringValue().regexpMatch(regex)
  )
}

private string getASshCommand() { result = "ssh" or result = "putty.exe" or result = "kitty.exe" }

/**
 * isSsh defines a string value node whose value is a ssh client or similar,
 * whose arguments can be commands that will be executed on the remote host.
 */
private predicate isSsh(DataFlow::Node node) {
  exists(string regex |
    regex = ".*(^|/)(" + concat(string cmd | cmd = getASshCommand() | cmd, "|") + ")"
  |
    node.getStringValue().regexpMatch(regex)
  )
}
