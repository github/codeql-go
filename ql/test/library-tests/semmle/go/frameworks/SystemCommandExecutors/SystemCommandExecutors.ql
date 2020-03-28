/**
 * @name Command built from user-controlled sources
 * @description Building a system command from user-controlled sources is vulnerable to insertion of
 *              malicious code by the user.
 * @kind path-problem
 * @problem.severity error
 * @id go/command-injection
 * @tags security
 *       external/cwe/cwe-078
 */

import go
import semmle.go.security.CommandInjection
import DataFlow::PathGraph

// "net/http".Request sources:
class NethttpRequestSource extends UntrustedFlowSource::Range, DataFlow::Node {
  string thisPackagePath;
  string thisTypeName;

  NethttpRequestSource() {
    thisPackagePath = "net/http" and
    thisTypeName = "Request" and
    (
      // method calls:
      exists(DataFlow::MethodCallNode call, string methodName, int resultIndex |
        call.getTarget().hasQualifiedName(thisPackagePath, thisTypeName, methodName) and
        (
          methodName = "BasicAuth" and
          (
            //### 0 string string
            resultIndex = 0
            or
            //### 1 string string
            resultIndex = 1
          )
          or
          methodName = "FormValue" and
          //### 0 string string
          resultIndex = 0
          or
          methodName = "UserAgent" and
          //### 0 string string
          resultIndex = 0
          or
          methodName = "Referer" and
          //### 0 string string
          resultIndex = 0
          or
          methodName = "PostFormValue" and
          //### 0 string string
          resultIndex = 0
        )
      |
        resultIndex = 0 and this = call.getResult()
        or
        this = call.getResult(resultIndex)
      )
      or
      // field reads:
      exists(DataFlow::FieldReadNode fieldRead, DataFlow::Field fld, string fieldName |
        fieldRead = fld.getARead() and
        (
          fld.hasQualifiedName(thisPackagePath, thisTypeName, fieldName) and
          (
            //### string string
            fieldName = "Host"
            or
            //### string string
            fieldName = "Method"
            or
            //### string string
            fieldName = "Proto"
            or
            //### string string
            fieldName = "RequestURI"
            or
            //### string string
            fieldName = "RemoteAddr"
            or
            //### slice []string
            fieldName = "TransferEncoding"
          )
        )
      |
        this = fieldRead
      )
    )
  }
}

// "net/http".Cookie sources:
class NethttpCookieSource extends UntrustedFlowSource::Range, DataFlow::Node {
  string thisPackagePath;
  string thisTypeName;

  NethttpCookieSource() {
    thisPackagePath = "net/http" and
    thisTypeName = "Cookie" and
    (
      // method calls:
      exists(DataFlow::MethodCallNode call, string methodName, int resultIndex |
        call.getTarget().hasQualifiedName(thisPackagePath, thisTypeName, methodName) and
        (
          methodName = "String" and
          //### 0 string string
          resultIndex = 0
        )
      |
        resultIndex = 0 and this = call.getResult()
        or
        this = call.getResult(resultIndex)
      )
      or
      // field reads:
      exists(DataFlow::FieldReadNode fieldRead, DataFlow::Field fld, string fieldName |
        fieldRead = fld.getARead() and
        (
          fld.hasQualifiedName(thisPackagePath, thisTypeName, fieldName) and
          (
            //### slice []string
            fieldName = "Unparsed"
            or
            //### string string
            fieldName = "Value"
            or
            //### string string
            fieldName = "Path"
            or
            //### string string
            fieldName = "Domain"
            or
            //### string string
            fieldName = "RawExpires"
            or
            //### string string
            fieldName = "Raw"
            or
            //### string string
            fieldName = "Name"
          )
        )
      |
        this = fieldRead
      )
    )
  }
}

/**
 * ShellOrSudoExecution extends system command execution sinks to include
 * arguments passed to command interpreter binaries like shells, sudo,
 * or programming language interpreters.
 */
class ShellOrSudoExecution extends SystemCommandExecution::Range, DataFlow::CallNode {
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

class SystemCommandExecutors extends SystemCommandExecution::Range, DataFlow::CallNode {
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
class GoShCommandExecution extends SystemCommandExecution::Range, DataFlow::CallNode {
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
class SshCommandExecution extends SystemCommandExecution::Range, DataFlow::CallNode {
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
class ShellLike extends DataFlow::Node {
  ShellLike() {
    isSudoOrSimilar(this) or
    isShell(this) or
    isProgrammingLanguageCli(this) or
    isSsh(this)
  }
}

string getASudoCommand() {
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
predicate isSudoOrSimilar(DataFlow::Node node) {
  exists(string regex |
    regex = ".*(^|/)(" + concat(string cmd | cmd = getASudoCommand() | cmd, "|") + ")"
  |
    node.getStringValue().regexpMatch(regex)
  )
}

string getAShellCommand() {
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
predicate isShell(DataFlow::Node node) {
  exists(string regex |
    regex = ".*(^|/)(" + concat(string cmd | cmd = getAShellCommand() | cmd, "|") + ")"
  |
    node.getStringValue().regexpMatch(regex)
  )
}

string getAnInterpreterName() {
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
predicate isProgrammingLanguageCli(DataFlow::Node node) {
  // NOTE: we can enounter cases like /usr/bin/python3.1
  node.getStringValue().matches("%/python%")
  or
  exists(string regex |
    regex =
      ".*(^|/)(" + concat(string cmd | cmd = getAnInterpreterName() | cmd + "[\\d.\\-v]*", "|") + ")"
  |
    node.getStringValue().regexpMatch(regex)
  )
}

string getASshCommand() { result = "ssh" or result = "putty.exe" or result = "kitty.exe" }

/**
 * isSsh defines a string value node whose value is a ssh client or similar,
 * whose arguments can be commands that will be executed on the remote host.
 */
predicate isSsh(DataFlow::Node node) {
  exists(string regex |
    regex = ".*(^|/)(" + concat(string cmd | cmd = getASshCommand() | cmd, "|") + ")"
  |
    node.getStringValue().regexpMatch(regex)
  )
}

from CommandInjection::Configuration cfg, DataFlow::PathNode source, DataFlow::PathNode sink
where cfg.hasFlowPath(source, sink)
select sink.getNode(), source, sink, "This command depends on $@.", source.getNode(),
  "a user-provided value"
