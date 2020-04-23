/**
 * @name Bad redirect check
 * @description A redirect check that checks for a leading slash but not two
 *              leading slashes or a leading slash followed by a backslash is
 *              incomplete.
 * @kind path-problem
 * @problem.severity warning
 * @id go/bad-redirect-check
 * @tags security
 *       external/cwe/cwe-601
 * @precision high
 */

import go
import semmle.go.security.OpenUrlRedirectCustomizations
import DataFlow::PathGraph

// holds if `checked` is checked for a leading slash
StringOps::HasPrefix checkForLeadingSlash(DataFlow::Node checked) {
  exists(DataFlow::Node substr |
    result.getBaseString() = checked and result.getSubstring() = substr
  |
    substr.getStringValue() = "/"
  )
}

// holds if `checked` is checked for a slash in its second position
DataFlow::Node checkForSecondSlash(DataFlow::Node checked) {
  exists(StringOps::HasPrefix hp | result = hp and hp.getBaseString() = checked |
    hp.getSubstring().getStringValue() = "//"
  )
  or
  exists(DataFlow::EqualityTestNode eq, DataFlow::Node slash, DataFlow::ElementReadNode er |
    result = eq
  |
    slash.getStringValue() = "/" and
    er.getBase() = checked and
    er.getIndex().getIntValue() = 1 and
    eq.eq(_, er, slash)
  )
}

// holds if `checked` is checked for a backslash in its second position
DataFlow::Node checkForSecondBackslash(DataFlow::Node checked) {
  exists(StringOps::HasPrefix hp | result = hp and hp.getBaseString() = checked |
    hp.getSubstring().getStringValue() = "/\\"
  )
  or
  exists(DataFlow::EqualityTestNode eq, DataFlow::Node slash, DataFlow::ElementReadNode er |
    result = eq
  |
    slash.getStringValue() = "\\" and
    er.getBase() = checked and
    er.getIndex().getIntValue() = 1 and
    eq.eq(_, er, slash)
  )
}

class Configuration extends TaintTracking::Configuration {
  Configuration() { this = "BadRedirectCheck" }

  override predicate isSource(DataFlow::Node source) { this.isSource(source, _) }

  predicate isSource(DataFlow::Node source, BadRedirectCheckBarrier check) {
    source = check.getAGuardedNode()
    or
    exists(DataFlow::Node pred | pred.getASuccessor() = source | check.guardsFlowFrom(pred))
  }

  override predicate isAdditionalTaintStep(DataFlow::Node pred, DataFlow::Node succ) {
    // this is very over-approximate, because most filtering is done by the isSource predicate
    exists(Write w | w.writesField(succ, _, pred))
  }

  override predicate isSanitizerOut(DataFlow::Node node) {
    // assume this value is safe if something is prepended to it.
    exists(StringOps::Concatenation conc, int i, int j | i < j |
      node = conc.getOperand(j) and
      exists(conc.getOperand(i))
    )
  }

  override predicate isSink(DataFlow::Node sink) { sink instanceof OpenUrlRedirect::Sink }
}

// A barrier class that considers any URLs that come out of bad redirect check as "guarded"; this
// will be a requirement of the source.
class BadRedirectCheckBarrier extends DataFlow::BarrierGuard {
  Expr checked;

  BadRedirectCheckBarrier() {
    exists(SsaWithFields v |
      checked = v.getAUse().asExpr() and
      // a check for a leading slash
      this = checkForLeadingSlash(DataFlow::exprNode(checked)) and
      // where there does not exist a check for both a second slash and a second backslash
      not (
        exists(checkForSecondSlash(v.getAUse())) and exists(checkForSecondBackslash(v.getAUse()))
      )
    )
    or
    exists(
      DeclaredFunction f, FunctionInput inp, FunctionOutput outp,
      ControlFlow::ConditionGuardNode guard, BadRedirectCheckBarrier check
    |
      check.checks(inp.getExitNode(f.getFuncDecl()).getASuccessor*().asExpr(), _) and
      guard.ensures(outp.getEntryNode(f.getFuncDecl()), _) and
      TaintTracking::localTaint(check, outp.getEntryNode(f.getFuncDecl())) and
      this = f.getACall() and
      checked = inp.getNode(this).asExpr()
    )
  }

  override predicate checks(Expr e, boolean b) { e = checked and b = [false, true] }
}

from Configuration cfg, DataFlow::PathNode source, DataFlow::PathNode sink, DataFlow::Node check
where
  cfg.isSource(source.getNode(), check) and
  cfg.hasFlowPath(source, sink)
select check, source, sink,
  "This is a check that $@, which flows into a $@, has a leading slash, but not that does not have '/' or '\\' in its second position.",
  source.getNode(), "this value", sink.getNode(), "redirect"
