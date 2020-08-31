/*
* @name Unsafety Signature Generate checking
* @description Unsafety Signature that means the developer calculate signatures by just adding it like "a+b" which would bypass easily
* @kind problem
* @problem.severity recommendation
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
            a1.getRhs() = call1.asExpr()
            and
            DataFlow::localFlow(call1, call2.getReceiver())
            and
            call2.getCalleeName() = "Write"
            and
            source = f1.getACall()
            )
    }

    override predicate isSink(DataFlow::Node sink) {
        exists(Assignment a1, DataFlow::CallNode call1, DataFlow::BinaryOperationNode b1|
            call1.getCalleeName() = "EncodeToString"
            and 
            a1.getRhs() = call1.asExpr()
            and
            (
                DataFlow::localFlow(call1, b1.getLeftOperand())
                or
                DataFlow::localFlow(call1, b1.getRightOperand())
            )
        )
    }
}

from UnsafetySignatureGenerate config, DataFlow::Node source, DataFlow::Node sink
where config.hasFlow(source, sink)
select source, sink