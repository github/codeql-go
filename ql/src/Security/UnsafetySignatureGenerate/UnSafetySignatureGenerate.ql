/*
@name Unsafety Signature Generate checking
@description Unsafety Signature that means the developer calculate signatures by just adding it like "a+b" which would bypass easily
@kind authoraztion
*/
import go
import semmle.go.dataflow.DataFlow

class UnsafetySignatureGenerate extends DataFlow::Configuration {
    UnsafetySignatureGenerate() {this = "UnsafetySignatureGenerate"}
    override predicate isSource(DataFlow::Node source) {
        exists(Function f1, Function f2, DataFlow::CallNode call1, DataFlow::CallNode call2, Assignment a1 |
            (
                f1.hasQualifiedName("github.com/astaxie/beego.Controller", "GetString") or
                f1.hasQualifiedName("github.com/astaxie/beego.Controller", "GetStrings") or
                f1.hasQualifiedName("github.com/astaxie/beego.Controller", "GetInt") or
                f1.hasQualifiedName("github.com/astaxie/beego.Controller", "GetFloat") or
                f1.hasQualifiedName("github.com/astaxie/beego.Controller", "GetBool")
            )
            and 
            f2.hasQualifiedName("crypto/sha256", "New")
            and
            call1 = f2.getACall()
            and 
            call1.toString().matches(a1.getRhs().toString())
            and
            call2.getReceiver().toString() = a1.getLhs().toString()
            and
            call2.getCalleeName() = "Write"
            and 
            source.asExpr() = f1.getACall().asExpr()
            )
    }

    override predicate isSink(DataFlow::Node sink) {
        exists(Assignment a1, DataFlow::CallNode call1, DataFlow::BinaryOperationNode b1|
            call1.getCalleeName() = "EncodeToString"
            and 
            a1.getRhs() = call1.asExpr()
            and
            (
                b1.getLeftOperand().asExpr().toString() = a1.getLhs().toString() or 
                b1.getRightOperand().asExpr().toString() = a1.getLhs().toString())
            )
    }
}

from UnsafetySignatureGenerate config, DataFlow::Node source, DataFlow::Node sink
where config.hasFlow(source, sink)
select "Patameter come from:", source, "at line: ", source.getStartLine(), "Correspond usage: ", sink, "at line: ", sink.getStartLine()