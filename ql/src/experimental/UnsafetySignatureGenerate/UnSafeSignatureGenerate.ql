import go

from
  DataFlow::CallNode c1, DataFlow::Method m, DataFlow::BinaryOperationNode b, Assignment a,
  Assignment a1
where
  (
    m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetString") or
    m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetStrings") or
    m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetInt") or
    m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetFloat") or
    m.hasQualifiedName("github.com/astaxie/beego.Controller", "GetBool")
  ) and
  b.getOperator() = "==" and
  a.getRhs() = m.getACall().asExpr() and
  c1.getCalleeName() = "EncodeToString" and
  c1.asExpr() = a1.getRhs() and
  (
    b.getAnOperand().toString() = a1.getAnLhs().toString() and
    b.getAnOperand().toString() = a.getAnLhs().toString()
  )
select "Get input here: ", m.getACall(), "Get hash algorithm result here: ", c1,
  "Web's hash result compare with developer result here: ", b
