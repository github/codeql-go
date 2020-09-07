/*
 * @name Unsafe signature calculation
 * @description Unsafety Signature that means the developer calculate signatures by just adding it like "a+b" which would bypass easily
 * @kind problem
 * @problem.severity recommendation
 */

import go
import semmle.go.dataflow.DataFlow

class UnsafetySignatureGenerate extends DataFlow::Configuration {
  UnsafetySignatureGenerate() { this = "UnsafetySignatureGenerate" }

  override predicate isSource(DataFlow::Node source) {
    exists(DataFlow::CallNode c1, DataFlow::Method m, DataFlow::ExprNode d |
      (
        m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetString") or
        m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetStrings") or
        m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetInt") or
        m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetFloat") or
        m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetBool")
      ) and
      c1.getCalleeName() = "Write" and
      DataFlow::localFlow(m.getACall(), d) and
      source = d.getASuccessor()
    )
  }

  override predicate isSink(DataFlow::Node sink) {
    exists(
      DataFlow::CallNode call1, DataFlow::ExprNode d, DataFlow::Method m,
      DataFlow::BinaryOperationNode b
    |
      call1.getCalleeName() = "EncodeToString" and
      (
        m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetString") or
        m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetStrings") or
        m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetInt") or
        m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetFloat") or
        m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetBool")
      ) and
      DataFlow::localFlow(m.getACall(), d) and
      b.getAnOperand() = d and
      b.getAnOperand() = call1.getASuccessor().getASuccessor() and
      sink = d
    )
  }
}

from UnsafetySignatureGenerate config, DataFlow::Node source, DataFlow::Node sink
where config.hasFlow(source, sink)
select source, sink