/**
 * @name: Use of insecure HostKeyCallback implementation
 * @description Detects insecure SSL client configuration when assigned an implementation of the `HostKeyCallback` that exclusively returns `nil`.
 * @kind path-problem
 * @problem.severity error
 * @precision very-high
 * @id go/security/crypto/ssl/insecure-configuration-client-config
 * @tags security
 */

import go
import DataFlow::PathGraph

class InsecureIgnoreHostKey extends Function {
  InsecureIgnoreHostKey() {
    this.hasQualifiedName("golang.org/x/crypto/ssh", "InsecureIgnoreHostKey")
  }
}

/** Gets a function definition corresponding to the given function node, if any. */
FuncDef getFuncDef(DataFlow::FunctionNode f) {
  result = f.(DataFlow::GlobalFunctionNode).getFunction().(DeclaredFunction).getFuncDecl() or
  result = f.(DataFlow::FuncLitNode).getExpr()
}

/** Gets a value returned by the given function via a return statement, if any. */
DataFlow::ResultNode getAResultNode(DataFlow::FunctionNode f) { result.getRoot() = getFuncDef(f) }

/** A callback function value that is insecure when used as a `HostKeyCallback`, because it always returns `nil`. */
class InsecureHostKeyCallbackFunc extends DataFlow::Node {
  InsecureHostKeyCallbackFunc() {
    // Either a call to InsecureIgnoreHostKey(), which we know returns an insecure callback.
    this = any(InsecureIgnoreHostKey f).getACall()
    or
    // Or a callback function in the source code  (named or anonymous) that always returns nil.
    forex(DataFlow::ResultNode returnValue | returnValue = getAResultNode(this) |
      returnValue = Builtin::nil().getARead()
    )
  }
}

class HostKeyCallbackField extends Field {
  HostKeyCallbackField() {
    this.hasQualifiedName("golang.org/x/crypto/ssh", "ClientConfig", "HostKeyCallback")
  }
}

class HostKeyCallbackAssignmentConfig extends DataFlow::Configuration {
  HostKeyCallbackAssignmentConfig() { this = "HostKeyCallbackAssignmentConfig" }

  override predicate isSource(DataFlow::Node source) {
    source instanceof InsecureHostKeyCallbackFunc
  }

  override predicate isSink(DataFlow::Node sink) {
    exists(HostKeyCallbackField f | sink = f.getAWrite().getRhs())
  }
}

from HostKeyCallbackAssignmentConfig config, DataFlow::PathNode source, DataFlow::PathNode sink
where config.hasFlowPath(source, sink)
select sink, source, sink,
  "Configuring SSH ClientConfig with insecure HostKeyCallback implementation from $@.",
  source.getNode(), "this source"
