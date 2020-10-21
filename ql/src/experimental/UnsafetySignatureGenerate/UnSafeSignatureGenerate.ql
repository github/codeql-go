/**
 * @name Unsafe signature calculation
 * @description Unsafety Signature that means the developer calculate signatures by just adding it like "a+b" which would bypass easily
 * @kind problem
 * @problem.severity recommendation
 */

import go

from
  DataFlow::CallNode c1, DataFlow::Method m, DataFlow::EqualityTestNode b, Assignment a,
  Assignment a1, DataFlow::CallNode c
where
  (
    m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetString") or
    m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetStrings") or
    m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetInt") or
    m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetFloat") or
    m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetBool")
  ) and
  c = m.getACall() and
  b.getOperator() = "==" and
  a.getRhs() = m.getACall().asExpr() and
  c1.getCalleeName() = "EncodeToString" and
  c1.asExpr() = a1.getRhs() and
  DataFlow::localFlow(m.getACall(), b.getAnOperand()) and
  DataFlow::localFlow(c1, b.getAnOperand())
select b, "Comparison between $@ and $@.", c, "user input", c1, "hash result"
