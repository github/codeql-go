import go

class FatalFunction extends NoreturnAnnotatedFunction {
  FatalFunction() { getName() = ["Fatal", "Fatalf"] }
}

class OsExitFunction extends NoreturnAnnotatedFunction {
  OsExitFunction() { hasQualifiedName("os", "Exit") }
}

class GinkgoFailFunction extends NoreturnAnnotatedFunction {
  GinkgoFailFunction() { hasQualifiedName("github.com/onsi/ginkgo", "Fail") }
}

abstract class NoreturnAnnotatedFunction extends Function {
  override predicate mayReturnNormally() { none() }
}

class UserDefinedFunction extends DeclaredFunction {
  UserDefinedFunction() { not this instanceof NoreturnAnnotatedFunction and exists(getBody()) }

  override predicate mayReturnNormally() { ControlFlow::mayReturnNormally(getFuncDecl()) }
}

/**
 * Matches if `node` is a reference to the `nil` builtin constant.
 */
predicate isNil(DataFlow::Node node) {
  exists(Expr nilRef | nilRef = Builtin::nil().getAReference() | node.asExpr() = nilRef)
}

/**
 * Holds if `n1` and `n2` are both in a reachable basic block, and `n1` dominates `n2`.
 */
predicate dominates(ControlFlow::Node n1, ControlFlow::Node n2) {
  exists(BasicBlock bb, int i, int j |
    bb = n1.getBasicBlock() and
    bb = n2.getBasicBlock() and
    bb.getNode(i) = n1 and
    bb.getNode(j) = n2 and
    i <= j
  )
  or
  n1.getBasicBlock().(ReachableBasicBlock).strictlyDominates(n2.getBasicBlock())
}

/**
 * Matches if `value` is certainly consumed as the parameter to some function call before `useSite`.
 */
predicate valueConsumedByFunctionBefore(DataFlow::SsaNode value, DataFlow::Node useSite) {
  exists(DataFlow::CallNode checkCall |
    value.getAUse() = checkCall.getAnArgument() and
    dominates(checkCall.asInstruction(), useSite.asInstruction())
  )
}

/**
 * Matches if any of the return-sites in `call` may return a non-nil pointer alongside an error value
 */
predicate calleeMayReturnNonNilWithError(DataFlow::CallNode call) {
  exists(
    FuncDef callee, IR::ReturnInstruction ret, DataFlow::Node ptrReturn, DataFlow::Node errReturn
  |
    callee = call.getACallee() and
    callee = ret.getRoot() and
    ptrReturn = DataFlow::instructionNode(ret.getResult(0)) and
    errReturn = DataFlow::instructionNode(ret.getResult(1)) and
    not isNil(ptrReturn) and
    not isNil(errReturn)
  )
}

from
  DataFlow::SsaNode ptr, DataFlow::SsaNode err, DataFlow::CallNode call,
  DataFlow::PointerDereferenceNode deref
where
  ptr.getAPredecessor() = call.getResult(0) and
  err.getAPredecessor() = call.getResult(1) and
  exists(Type t | t = ptr.getType().getUnderlyingType() |
    t instanceof PointerType or t instanceof SliceType or t instanceof InterfaceType
  ) and
  err.getType().getEntity() = Builtin::error() and
  deref.getOperand() = ptr.getAUse() and
  not calleeMayReturnNonNilWithError(call) and
  not exists(ControlFlow::ConditionGuardNode cgn |
    cgn.ensuresEq(err.getAUse(), Builtin::nil().getARead()) or
    cgn.ensuresNeq(ptr.getAUse(), Builtin::nil().getARead())
  |
    cgn.dominates(deref.getBasicBlock())
  ) and
  not valueConsumedByFunctionBefore(err, deref) and
  not valueConsumedByFunctionBefore(ptr, deref)
select deref.getOperand(),
  ptr.getSourceVariable().getName() + " may be nil here, because $@ has not been checked.", err,
  err.getSourceVariable().getName()
