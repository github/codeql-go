/**
 * @kind problem
 */

import go

class SwitchNodeOverride extends ControlFlow::Node {
  override ControlFlow::Node getASuccessor() {
    result = super.getASuccessor() and
    forall(ExpressionSwitchStmt s | this.isFirstNodeOf(s) |
      exists(s.getExpr()) or
      s.getNumCase() = 0 or
      result.isFirstNodeOf(s.getACase().getAnExpr()) or
      result.isFirstNodeOf(s.getACase().getStmt(0))
    )
  }
}

/**
 * Matches if `node` is a reference to the `nil` builtin constant.
 */
predicate isNil(DataFlow::Node node) {
  exists(Expr nilRef | nilRef = Builtin::nil().getAReference() | node.asExpr() = nilRef)
}

/**
 * Matches if any of the return-sites in `call` may return a nil pointer alongside an error value.
 *
 * This is both an over- and under-estimate: over in that opaque functions may use this convention,
 * and under in that functions with bodies are only recognised if they use a literal `nil` for the
 * pointer return value at some return site.
 */
predicate calleeMayReturnNilWithError(DataFlow::CallNode call) {
  not exists(call.getACallee())
  or
  exists(FuncDef callee | callee = call.getACallee() |
    not exists(callee.getBody())
    or
    exists(IR::ReturnInstruction ret, DataFlow::Node ptrReturn, DataFlow::Node errReturn |
      callee = call.getACallee() and
      callee = ret.getRoot() and
      ptrReturn = DataFlow::instructionNode(ret.getResult(0)) and
      errReturn = DataFlow::instructionNode(ret.getResult(1)) and
      isNil(ptrReturn) and
      not isNil(errReturn)
    )
  )
}

predicate isCallReturningPointerAndError(
  DataFlow::CallNode call, DataFlow::SsaNode ptr, DataFlow::SsaNode err
) {
  calleeMayReturnNilWithError(call) and
  exists(Type t | t = ptr.getType().getUnderlyingType() |
    t instanceof PointerType or t instanceof SliceType or t instanceof InterfaceType
  ) and
  err.getType().getEntity() = Builtin::error()
}

predicate checksValue(IR::Instruction instruction, DataFlow::SsaNode value) {
  exists(DataFlow::InstructionNode instNode | instNode.asInstruction() = instruction |
    instNode.(DataFlow::CallNode).getAnArgument() = value.getAUse() or
    instNode.(DataFlow::EqualityTestNode).getAnOperand() = value.getAUse()
  )
  or
  value.getAUse().asInstruction() = instruction and
  (
    exists(ExpressionSwitchStmt s | instruction.(IR::EvalInstruction).getExpr() = s.getExpr())
    or
    // This case accounts for both a type-switch or cast used to check `value`
    exists(TypeAssertExpr e | instruction.(IR::EvalInstruction).getExpr() = e.getExpr())
  )
}

predicate returnUncheckedAtNode(
  DataFlow::CallNode call, ControlFlow::Node node, DataFlow::SsaNode ptr, DataFlow::SsaNode err
) {
  ptr.getAPredecessor() = call.getResult(0) and
  err.getAPredecessor() = call.getResult(1) and
  (
    isCallReturningPointerAndError(call, ptr, err) and call.asInstruction() = node
    or
    returnUncheckedAtNode(call, node.getAPredecessor(), ptr, err) and
    not exists(DataFlow::SsaNode checked | DataFlow::localFlow(ptr, checked) |
      checksValue(node, checked)
    ) and
    not exists(DataFlow::SsaNode checked | DataFlow::localFlow(err, checked) |
      checksValue(node, checked)
    )
    // Alternate more-generous implementation, considering any use to be a check,
    // including e.g. copying into a parameter or other variable:
    //not exists(DataFlow::InstructionNode i | i = ptr.getAUse() and i.asInstruction() = node) and
    //not exists(DataFlow::InstructionNode i | i = err.getAUse() and i.asInstruction() = node))
  )
}

from
  DataFlow::CallNode call, DataFlow::SsaNode ptr, DataFlow::SsaNode err,
  DataFlow::PointerDereferenceNode deref, ControlFlow::Node derefNodePred
where
  deref.getOperand().asInstruction().getAPredecessor() = derefNodePred and
  returnUncheckedAtNode(call, derefNodePred, ptr, err) and
  deref.getOperand() = ptr.getAUse()
select deref.getOperand(),
  ptr.getSourceVariable().getName() + " may be nil here, because $@ may not been checked.", err,
  err.getSourceVariable().getName()
